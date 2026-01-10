package elf

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

// Config allows customizing the target
type Config struct {
	Entry       string // Default "_start"
	BaseAddr    uint64 // Default 0x400000
	Interpreter string // e.g. "/lib64/ld-linux-x86-64.so.2"
}

// Linker holds the state of the linking process
type Linker struct {
	Config      Config
	Objects     []*InputObject
	SharedLibs  []*SharedObject
	GlobalTable map[string]*ResolvedSymbol

	// Dynamic Linking State
	DynStrTab  []byte
	DynSyms    []Elf64Sym
	RelaDyn    []Elf64Rela
	GotEntries []string

	// Output Buffers
	InterpSect  []byte
	DynSect     []byte
	DynSymSect  []byte
	DynStrSect  []byte
	RelaDynSect []byte
	TextSection []byte // Includes PLT
	DataSection []byte // Includes GOT

	// Section String Table (for objdump)
	ShStrTab []byte

	// Calculated Addresses
	TextAddr  uint64
	DataAddr  uint64
	BssAddr   uint64
	EntryAddr uint64
	DynAddr   uint64

	// Sizes
	BssSize uint64
}

// ResolvedSymbol represents a symbol's final location
type ResolvedSymbol struct {
	Name    string
	Value   uint64 // Absolute Virtual Address
	Section string // "text", "data", "bss", "plt", "dynamic"
	Defined bool
}

// NewLinker creates a new Linker instance
func NewLinker(cfg Config) *Linker {
	if cfg.Entry == "" {
		cfg.Entry = "_start"
	}
	if cfg.BaseAddr == 0 {
		cfg.BaseAddr = 0x400000
	}
	if cfg.Interpreter == "" {
		cfg.Interpreter = "/lib64/ld-linux-x86-64.so.2"
	}

	return &Linker{
		Config:      cfg,
		GlobalTable: make(map[string]*ResolvedSymbol),
		DynSyms:     []Elf64Sym{{}}, // Init with Null Symbol
		DynStrTab:   []byte{0},
		ShStrTab:    []byte{0}, // Init with Null Byte
	}
}

// AddObject adds an in-memory object file
func (l *Linker) AddObject(name string, data []byte) error {
	obj, err := LoadObject(name, data)
	if err != nil {
		return err
	}
	l.Objects = append(l.Objects, obj)
	return nil
}

// AddArchive adds a .a file from disk
func (l *Linker) AddArchive(path string) error {
	objs, err := LoadArchive(path)
	if err != nil {
		return err
	}
	l.Objects = append(l.Objects, objs...)
	return nil
}

// AddSharedLib parses a .so file for dynamic symbol resolution
func (l *Linker) AddSharedLib(path string, data []byte) error {
	so, err := LoadSharedObject(path, data)
	if err != nil {
		return err
	}
	l.SharedLibs = append(l.SharedLibs, so)
	return nil
}

// Link performs the linking process and writes the executable
func (l *Linker) Link(outPath string) error {
	if err := l.scanSymbols(); err != nil {
		return err
	}
	l.layout()
	if err := l.applyRelocations(); err != nil {
		return err
	}
	return l.write(outPath)
}

func (l *Linker) scanSymbols() error {
	l.GlobalTable[l.Config.Entry] = &ResolvedSymbol{Name: l.Config.Entry, Defined: false}

	// 1. Scan Objects
	for _, obj := range l.Objects {
		for _, sym := range obj.Symbols {
			if sym.Bind == STB_GLOBAL {
				if sym.Section != nil {
					// Definition
					if existing, ok := l.GlobalTable[sym.Name]; ok && existing.Defined {
						return fmt.Errorf("duplicate symbol: %s in %s", sym.Name, obj.Name)
					}
					l.GlobalTable[sym.Name] = &ResolvedSymbol{
						Name: sym.Name, Defined: true, Section: "text",
					}
				} else {
					// Reference
					if _, ok := l.GlobalTable[sym.Name]; !ok {
						l.GlobalTable[sym.Name] = &ResolvedSymbol{Name: sym.Name, Defined: false}
					}
				}
			}
		}
	}

	// 2. Scan Shared Libs
	for _, so := range l.SharedLibs {
		for _, symName := range so.Symbols {
			if target, ok := l.GlobalTable[symName]; ok && !target.Defined {
				target.Defined = true
				target.Section = "dynamic"
				l.addDynamicSymbol(symName)
				l.GotEntries = append(l.GotEntries, symName)
			}
		}
	}

	// 3. Entry Point Check
	entry := l.GlobalTable[l.Config.Entry]
	if !entry.Defined {
		if mainSym, ok := l.GlobalTable["main"]; ok && mainSym.Defined {
			entry.Defined = true
			entry.Section = "stub"
		} else {
			return fmt.Errorf("undefined entry point: %s", l.Config.Entry)
		}
	}
	return nil
}

func (l *Linker) addDynamicSymbol(name string) {
	nameIdx := uint32(len(l.DynStrTab))
	l.DynStrTab = append(l.DynStrTab, []byte(name)...)
	l.DynStrTab = append(l.DynStrTab, 0)

	// ELF64 Sym: Info = (Bind << 4) + Type
	l.DynSyms = append(l.DynSyms, Elf64Sym{
		Name: nameIdx, Info: 0x12, Shndx: 0, Value: 0, Size: 0,
	})
}

// addShStr adds a string to the Section Header String Table and returns its index
func (l *Linker) addShStr(s string) uint32 {
	idx := uint32(len(l.ShStrTab))
	l.ShStrTab = append(l.ShStrTab, []byte(s)...)
	l.ShStrTab = append(l.ShStrTab, 0)
	return idx
}

func (l *Linker) layout() {
	headerSize := uint64(64 + 56*5)

	// --- Text Segment ---
	l.TextAddr = l.Config.BaseAddr + headerSize
	if l.TextAddr%16 != 0 {
		l.TextAddr += 16 - (l.TextAddr % 16)
	}

	// 1. .interp
	l.InterpSect = append([]byte(l.Config.Interpreter), 0)

	// 2. PLT
	pltOffset := uint64(len(l.InterpSect))

	// 3. Object Text
	var objText []byte
	if l.GlobalTable[l.Config.Entry].Section == "stub" {
		objText = append(objText, make([]byte, 29)...)
		l.GlobalTable[l.Config.Entry].Value = l.TextAddr + pltOffset
	}

	for _, obj := range l.Objects {
		for _, sec := range obj.Sections {
			if sec.Flags&SHF_EXECINSTR != 0 {
				pad := (16 - (len(objText) % 16)) % 16
				objText = append(objText, make([]byte, pad)...)
				sec.OutputOffset = pltOffset + uint64(len(objText))
				sec.VirtualAddress = l.TextAddr + sec.OutputOffset
				objText = append(objText, sec.Data...)
			}
		}
	}

	// --- Data Segment ---
	estPltSize := uint64(len(l.GotEntries) * 6)
	totalTextSize := uint64(len(l.InterpSect)) + estPltSize + uint64(len(objText))

	l.DataAddr = l.TextAddr + totalTextSize
	if l.DataAddr%4096 != 0 {
		l.DataAddr += 4096 - (l.DataAddr % 4096)
	}

	dynSize := 16 * 15
	symSize := len(l.DynSyms) * 24
	strSize := len(l.DynStrTab)
	relaSize := len(l.GotEntries) * 24

	gotOffset := uint64(dynSize + symSize + strSize + relaSize)
	gotAddr := l.DataAddr + gotOffset

	// Generate PLT
	pltBuf := new(bytes.Buffer)
	for i, symName := range l.GotEntries {
		targetGot := gotAddr + uint64(i*8)
		currentPC := l.TextAddr + pltOffset + uint64(pltBuf.Len())
		rel := int32(targetGot - (currentPC + 6))

		pltBuf.WriteByte(0xFF)
		pltBuf.WriteByte(0x25)
		binary.Write(pltBuf, Le, rel)

		l.GlobalTable[symName].Value = currentPC
		l.GlobalTable[symName].Section = "plt"
	}

	actualPlt := pltBuf.Bytes()
	fullText := append(l.InterpSect, actualPlt...)

	currentTextLen := uint64(len(fullText))
	objText = nil

	if l.GlobalTable[l.Config.Entry].Section == "stub" {
		l.GlobalTable[l.Config.Entry].Value = l.TextAddr + currentTextLen
		objText = append(objText, make([]byte, 29)...)
	}

	for _, obj := range l.Objects {
		for _, sec := range obj.Sections {
			if sec.Flags&SHF_EXECINSTR != 0 {
				pad := (16 - ((currentTextLen + uint64(len(objText))) % 16)) % 16
				objText = append(objText, make([]byte, pad)...)
				sec.OutputOffset = currentTextLen + uint64(len(objText))
				sec.VirtualAddress = l.TextAddr + sec.OutputOffset
				objText = append(objText, sec.Data...)
			}
		}
	}
	l.TextSection = append(fullText, objText...)

	// Re-align Data Addr
	l.DataAddr = l.TextAddr + uint64(len(l.TextSection))
	if l.DataAddr%4096 != 0 {
		l.DataAddr += 4096 - (l.DataAddr % 4096)
	}
	gotAddr = l.DataAddr + gotOffset

	// Generate Dynamic Data
	dynBuf := new(bytes.Buffer)
	writeDyn := func(tag int64, val uint64) {
		binary.Write(dynBuf, Le, tag)
		binary.Write(dynBuf, Le, val)
	}

	baseData := l.DataAddr
	symAddr := baseData + uint64(dynSize)
	strAddr := symAddr + uint64(symSize)
	relaAddr := strAddr + uint64(strSize)

	for _, lib := range l.SharedLibs {
		idx := len(l.DynStrTab)
		l.DynStrTab = append(l.DynStrTab, []byte(lib.Name)...)
		l.DynStrTab = append(l.DynStrTab, 0)
		writeDyn(DT_NEEDED, uint64(idx))
	}

	writeDyn(DT_STRTAB, strAddr)
	writeDyn(DT_SYMTAB, symAddr)
	writeDyn(DT_STRSZ, uint64(len(l.DynStrTab)))
	writeDyn(DT_SYMENT, 24)
	writeDyn(DT_RELA, relaAddr)
	writeDyn(DT_RELASZ, uint64(relaSize))
	writeDyn(DT_RELAENT, 24)
	writeDyn(DT_NULL, 0)

	l.DynSect = dynBuf.Bytes()

	symBuf := new(bytes.Buffer)
	for _, ds := range l.DynSyms {
		binary.Write(symBuf, Le, ds.Name)
		symBuf.WriteByte(ds.Info)
		symBuf.WriteByte(ds.Other)
		binary.Write(symBuf, Le, ds.Shndx)
		binary.Write(symBuf, Le, ds.Value)
		binary.Write(symBuf, Le, ds.Size)
	}
	l.DynSymSect = symBuf.Bytes()

	relaBuf := new(bytes.Buffer)
	for i := range l.GotEntries {
		offset := gotAddr + uint64(i*8)
		info := uint64((i+1)<<32) | uint64(R_X86_64_GLOB_DAT)
		binary.Write(relaBuf, Le, offset)
		binary.Write(relaBuf, Le, info)
		binary.Write(relaBuf, Le, int64(0))
	}
	l.RelaDynSect = relaBuf.Bytes()

	gotData := make([]byte, len(l.GotEntries)*8)

	l.DataSection = append(l.DynSect, l.DynSymSect...)
	l.DataSection = append(l.DataSection, l.DynStrTab...)
	l.DataSection = append(l.DataSection, l.RelaDynSect...)
	l.DataSection = append(l.DataSection, gotData...)

	currentDataLen := uint64(len(l.DataSection))
	for _, obj := range l.Objects {
		for _, sec := range obj.Sections {
			if sec.Flags&SHF_WRITE != 0 && sec.Type == SHT_PROGBITS {
				pad := (8 - ((currentDataLen) % 8)) % 8
				l.DataSection = append(l.DataSection, make([]byte, pad)...)
				currentDataLen += uint64(pad)
				sec.OutputOffset = currentDataLen
				sec.VirtualAddress = l.DataAddr + sec.OutputOffset
				l.DataSection = append(l.DataSection, sec.Data...)
				currentDataLen += uint64(len(sec.Data))
			}
		}
	}

	l.BssAddr = l.DataAddr + currentDataLen
	currBss := uint64(0)
	for _, obj := range l.Objects {
		for _, sec := range obj.Sections {
			if sec.Flags&SHF_WRITE != 0 && sec.Type == SHT_NOBITS {
				if currBss%8 != 0 {
					currBss += 8 - (currBss % 8)
				}
				sec.OutputOffset = currBss
				sec.VirtualAddress = l.BssAddr + currBss
				currBss += uint64(len(sec.Data))
			}
		}
	}
	l.BssSize = currBss
	l.EntryAddr = l.GlobalTable[l.Config.Entry].Value
}

func (l *Linker) applyRelocations() error {
	if l.GlobalTable[l.Config.Entry].Section == "stub" {
		stub := []byte{
			0xE8, 0, 0, 0, 0, // call main
			0x48, 0x31, 0xFF, // xor rdi, rdi
			0x48, 0xC7, 0xC0, 0x3C, 0x00, 0x00, 0x00, // mov rax, 60
			0x0F, 0x05, // syscall
		}
		stubOffset := l.GlobalTable[l.Config.Entry].Value - l.TextAddr
		pc := l.TextAddr + stubOffset + 5
		target := l.GlobalTable["main"].Value
		binary.LittleEndian.PutUint32(stub[1:], uint32(int32(target-pc)))
		copy(l.TextSection[stubOffset:], stub)
	}

	for _, obj := range l.Objects {
		for _, sec := range obj.Sections {
			for _, r := range sec.Relocs {
				var symVal uint64
				if r.Sym.Bind == STB_LOCAL {
					symVal = r.Sym.Section.VirtualAddress + r.Sym.Value
				} else {
					global, ok := l.GlobalTable[r.Sym.Name]
					if !ok || !global.Defined {
						return fmt.Errorf("undefined symbol: %s", r.Sym.Name)
					}
					symVal = global.Value
				}
				P := sec.VirtualAddress + r.Offset
				var buf []byte
				var bufOff uint64
				if sec.Flags&SHF_EXECINSTR != 0 {
					buf = l.TextSection
					bufOff = sec.OutputOffset + r.Offset
				} else if sec.Flags&SHF_WRITE != 0 {
					buf = l.DataSection
					bufOff = sec.OutputOffset + r.Offset
				} else {
					continue
				}

				if bufOff >= uint64(len(buf)) {
					continue
				}

				switch r.Type {
				case R_X86_64_64:
					binary.LittleEndian.PutUint64(buf[bufOff:], uint64(int64(symVal)+r.Addend))
				case R_X86_64_32:
					binary.LittleEndian.PutUint32(buf[bufOff:], uint32(int64(symVal)+r.Addend))
				case R_X86_64_PC32, R_X86_64_PLT32:
					val := int32(int64(symVal) + r.Addend - int64(P))
					binary.LittleEndian.PutUint32(buf[bufOff:], uint32(val))
				}
			}
		}
	}
	return nil
}

// write generates the file with Section Headers to satisfy objdump
func (l *Linker) write(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Offsets
	hdrSize := uint64(64 + 56*5)
	textOff := hdrSize
	if textOff%16 != 0 {
		textOff += 16 - (textOff % 16)
	}

	textSize := uint64(len(l.TextSection))
	dataOff := textOff + textSize
	if dataOff%4096 != 0 {
		dataOff += 4096 - (dataOff % 4096)
	}

	dataSize := uint64(len(l.DataSection))

	// Define Sections for objdump
	// 0: NULL, 1: .interp, 2: .text, 3: .dynamic, 4: .data, 5: .bss, 6: .shstrtab

	l.addShStr("") // Add null string (index 0)
	idxInterp := l.addShStr(".interp")
	idxText := l.addShStr(".text")
	idxDyn := l.addShStr(".dynamic")
	idxData := l.addShStr(".data")
	idxBss := l.addShStr(".bss")
	idxShstr := l.addShStr(".shstrtab")

	// Append shstrtab to file content
	shStrOff := dataOff + dataSize

	// Calculate Table Offset
	shTableOff := shStrOff + uint64(len(l.ShStrTab))
	shNum := uint16(7)

	ehdr := Header{
		Type: ET_EXEC, Machine: EM_X86_64, Version: EV_CURRENT,
		Entry: l.EntryAddr, Phoff: 64, Ehsize: 64, Phentsize: 56, Phnum: 5,
		Shoff: shTableOff, Shentsize: 64, Shnum: shNum, Shstrndx: 6,
	}
	ehdr.Ident[0] = 0x7F
	ehdr.Ident[1] = 'E'
	ehdr.Ident[2] = 'L'
	ehdr.Ident[3] = 'F'
	ehdr.Ident[4] = ELFCLASS64
	ehdr.Ident[5] = ELFDATA2LSB
	ehdr.Ident[6] = EV_CURRENT
	ehdr.Ident[7] = 3

	binary.Write(f, Le, ehdr)

	// Segments
	binary.Write(f, Le, ProgHeader{Type: PT_INTERP, Flags: PF_R, Off: textOff, Vaddr: l.TextAddr, Paddr: l.TextAddr, Filesz: uint64(len(l.InterpSect)), Memsz: uint64(len(l.InterpSect)), Align: 1})
	binary.Write(f, Le, ProgHeader{Type: PT_LOAD, Flags: PF_R | PF_X, Off: 0, Vaddr: l.Config.BaseAddr, Paddr: l.Config.BaseAddr, Filesz: textOff + textSize, Memsz: textOff + textSize, Align: 4096})
	binary.Write(f, Le, ProgHeader{Type: PT_LOAD, Flags: PF_R | PF_W, Off: dataOff, Vaddr: l.DataAddr, Paddr: l.DataAddr, Filesz: dataSize, Memsz: dataSize + l.BssSize, Align: 4096})
	binary.Write(f, Le, ProgHeader{Type: PT_DYNAMIC, Flags: PF_R | PF_W, Off: dataOff, Vaddr: l.DataAddr, Paddr: l.DataAddr, Filesz: uint64(len(l.DynSect)), Memsz: uint64(len(l.DynSect)), Align: 8})
	binary.Write(f, Le, ProgHeader{Type: 6, Flags: PF_R, Off: 64, Vaddr: l.Config.BaseAddr + 64, Paddr: l.Config.BaseAddr + 64, Filesz: 56 * 5, Memsz: 56 * 5, Align: 8})

	// Body
	cur := uint64(64 + 56*5)
	if textOff > cur {
		f.Write(make([]byte, textOff-cur))
	}
	f.Write(l.TextSection)

	cur = textOff + textSize
	if dataOff > cur {
		f.Write(make([]byte, dataOff-cur))
	}
	f.Write(l.DataSection)

	// Write .shstrtab
	f.Write(l.ShStrTab)

	// Write Section Header Table
	// 0: NULL
	binary.Write(f, Le, SectionHeader{})

	// 1: .interp
	binary.Write(f, Le, SectionHeader{Name: idxInterp, Type: SHT_PROGBITS, Flags: SHF_ALLOC, Addr: l.TextAddr, Offset: textOff, Size: uint64(len(l.InterpSect)), Addralign: 1})

	// 2: .text
	binary.Write(f, Le, SectionHeader{Name: idxText, Type: SHT_PROGBITS, Flags: SHF_ALLOC | SHF_EXECINSTR, Addr: l.TextAddr, Offset: textOff, Size: textSize, Addralign: 16})

	// 3: .dynamic
	binary.Write(f, Le, SectionHeader{Name: idxDyn, Type: SHT_PROGBITS, Flags: SHF_ALLOC | SHF_WRITE, Addr: l.DataAddr, Offset: dataOff, Size: uint64(len(l.DynSect)), Link: 0, Addralign: 8})

	// 4: .data
	binary.Write(f, Le, SectionHeader{Name: idxData, Type: SHT_PROGBITS, Flags: SHF_ALLOC | SHF_WRITE, Addr: l.DataAddr, Offset: dataOff, Size: dataSize, Addralign: 4096})

	// 5: .bss
	binary.Write(f, Le, SectionHeader{Name: idxBss, Type: SHT_NOBITS, Flags: SHF_ALLOC | SHF_WRITE, Addr: l.BssAddr, Offset: dataOff + dataSize, Size: l.BssSize, Addralign: 8})

	// 6: .shstrtab
	binary.Write(f, Le, SectionHeader{Name: idxShstr, Type: SHT_STRTAB, Flags: 0, Addr: 0, Offset: shStrOff, Size: uint64(len(l.ShStrTab)), Addralign: 1})

	f.Chmod(0755)
	return nil
}