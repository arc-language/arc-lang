package elf

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

// Config allows customizing the target
type Config struct {
	Entry       string
	BaseAddr    uint64
	Interpreter string
}

// Linker holds the state of the linking process
type Linker struct {
	Config      Config
	Objects     []*InputObject
	SharedLibs  []*SharedObject
	GlobalTable map[string]*ResolvedSymbol

	DynStrTab  []byte
	DynSyms    []Elf64Sym
	RelaDyn    []Elf64Rela
	GotEntries []string

	InterpSect  []byte
	DynSect     []byte
	DynSymSect  []byte
	DynStrSect  []byte
	RelaDynSect []byte
	TextSection []byte
	DataSection []byte

	ShStrTab []byte

	TextAddr  uint64
	DataAddr  uint64
	BssAddr   uint64
	EntryAddr uint64
	DynAddr   uint64
	BssSize   uint64
}

type ResolvedSymbol struct {
	Name    string
	Value   uint64
	Section string
	Defined bool
}

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
		DynSyms:     []Elf64Sym{{}}, // Null Symbol
		DynStrTab:   []byte{0},
		ShStrTab:    []byte{0},
	}
}

func (l *Linker) AddObject(name string, data []byte) error {
	obj, err := LoadObject(name, data)
	if err != nil {
		return err
	}
	l.Objects = append(l.Objects, obj)
	return nil
}

func (l *Linker) AddArchive(path string) error {
	objs, err := LoadArchive(path)
	if err != nil {
		return err
	}
	l.Objects = append(l.Objects, objs...)
	return nil
}

func (l *Linker) AddSharedLib(path string, data []byte) error {
	so, err := LoadSharedObject(path, data)
	if err != nil {
		return err
	}
	l.SharedLibs = append(l.SharedLibs, so)
	return nil
}

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

	for _, obj := range l.Objects {
		for _, sym := range obj.Symbols {
			if sym.Bind == STB_GLOBAL {
				if sym.Section != nil {
					if existing, ok := l.GlobalTable[sym.Name]; ok && existing.Defined {
						return fmt.Errorf("duplicate symbol: %s in %s", sym.Name, obj.Name)
					}
					l.GlobalTable[sym.Name] = &ResolvedSymbol{
						Name: sym.Name, Defined: true, Section: "text",
					}
				} else {
					if _, ok := l.GlobalTable[sym.Name]; !ok {
						l.GlobalTable[sym.Name] = &ResolvedSymbol{Name: sym.Name, Defined: false}
					}
				}
			}
		}
	}

	if _, hasMain := l.GlobalTable["main"]; hasMain {
		if _, ok := l.GlobalTable["__libc_start_main"]; !ok {
			l.GlobalTable["__libc_start_main"] = &ResolvedSymbol{Name: "__libc_start_main", Defined: false}
		}
	}

	for _, so := range l.SharedLibs {
		for _, symName := range so.Symbols {
			cleanName := symName
			if idx := strings.Index(cleanName, "@"); idx != -1 {
				cleanName = cleanName[:idx]
			}

			if target, ok := l.GlobalTable[cleanName]; ok && !target.Defined {
				target.Defined = true
				target.Section = "dynamic"
				l.addDynamicSymbol(cleanName)
				l.GotEntries = append(l.GotEntries, cleanName)
			}
		}
	}

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

	l.DynSyms = append(l.DynSyms, Elf64Sym{
		Name: nameIdx, Info: 0x12, Shndx: 0, Value: 0, Size: 0,
	})
}

func (l *Linker) addShStr(s string) uint32 {
	idx := uint32(len(l.ShStrTab))
	l.ShStrTab = append(l.ShStrTab, []byte(s)...)
	l.ShStrTab = append(l.ShStrTab, 0)
	return idx
}

func align(val, align uint64) uint64 {
	if val%align != 0 {
		return val + (align - (val % align))
	}
	return val
}

func (l *Linker) layout() {
	headerSize := uint64(64 + 56*5)

	// --- Text Segment ---
	l.TextAddr = align(l.Config.BaseAddr+headerSize, 16)

	// .interp
	l.InterpSect = append([]byte(l.Config.Interpreter), 0)

	// PLT
	pltOffset := uint64(len(l.InterpSect))

	var objText []byte

	// Stub
	if l.GlobalTable[l.Config.Entry].Section == "stub" {
		l.GlobalTable[l.Config.Entry].Value = l.TextAddr + pltOffset
		objText = append(objText, make([]byte, 32)...)
	}

	for _, obj := range l.Objects {
		for _, sec := range obj.Sections {
			if sec.Flags&SHF_EXECINSTR != 0 {
				pad := (16 - (len(objText) % 16)) % 16
				for i := 0; i < int(pad); i++ {
					objText = append(objText, 0x90)
				}
				sec.OutputOffset = pltOffset + uint64(len(objText))
				sec.VirtualAddress = l.TextAddr + sec.OutputOffset
				objText = append(objText, sec.Data...)
			}
		}
	}

	// --- Data Segment ---
	estPltSize := uint64(len(l.GotEntries) * 6)
	totalTextSize := uint64(len(l.InterpSect)) + estPltSize + uint64(len(objText))

	l.DataAddr = align(l.TextAddr+totalTextSize, 4096)

	// Build Dynamic Data
	dynBuf := new(bytes.Buffer)
	symBuf := new(bytes.Buffer)
	for _, ds := range l.DynSyms {
		binary.Write(symBuf, Le, ds.Name); symBuf.WriteByte(ds.Info); symBuf.WriteByte(ds.Other)
		binary.Write(symBuf, Le, ds.Shndx); binary.Write(symBuf, Le, ds.Value); binary.Write(symBuf, Le, ds.Size)
	}
	l.DynSymSect = symBuf.Bytes()

	relaSize := uint64(len(l.GotEntries) * 24)
	
	// Layout & Alignment Tracking
	offset := uint64(0)
	offsetDyn := offset // Start of Dyn is 0 relative to DataAddr
	numDynEntries := len(l.SharedLibs) + 9
	sizeDyn := uint64(numDynEntries * 16)
	offset += sizeDyn

	offset = align(offset, 8)
	offsetSym := offset
	sizeSym := uint64(len(l.DynSymSect))
	offset += sizeSym

	offsetStr := offset
	sizeStr := uint64(len(l.DynStrTab))
	offset += sizeStr

	offset = align(offset, 8)
	offsetRela := offset
	offset += relaSize

	// CRITICAL: GOT ALIGNMENT
	offset = align(offset, 8)
	offsetGot := offset
	sizeGot := uint64(len(l.GotEntries) * 8)
	offset += sizeGot

	// Base Addresses
	// dynAddr (unused local removed)
	l.DynAddr = l.DataAddr + offsetDyn // Update struct field instead
	symAddr := l.DataAddr + offsetSym
	strAddr := l.DataAddr + offsetStr
	relaAddr := l.DataAddr + offsetRela
	gotAddr := l.DataAddr + offsetGot

	writeDyn := func(tag int64, val uint64) {
		binary.Write(dynBuf, Le, tag); binary.Write(dynBuf, Le, val)
	}
	for _, lib := range l.SharedLibs {
		l.DynStrTab = append(l.DynStrTab, []byte(lib.Name)...)
		l.DynStrTab = append(l.DynStrTab, 0)
		writeDyn(DT_NEEDED, uint64(len(l.DynStrTab)-len(lib.Name)-1))
	}
	
	writeDyn(DT_STRTAB, strAddr)
	writeDyn(DT_SYMTAB, symAddr)
	writeDyn(DT_STRSZ, uint64(len(l.DynStrTab)))
	writeDyn(DT_SYMENT, 24)
	writeDyn(DT_RELA, relaAddr)
	writeDyn(DT_RELASZ, relaSize)
	writeDyn(DT_RELAENT, 24)
	writeDyn(DT_NULL, 0)
	l.DynSect = dynBuf.Bytes()

	relaBuf := new(bytes.Buffer)
	for i := range l.GotEntries {
		rOff := gotAddr + uint64(i*8)
		rInfo := uint64((i+1)<<32) | uint64(R_X86_64_JMP_SLOT)
		binary.Write(relaBuf, Le, rOff)
		binary.Write(relaBuf, Le, rInfo)
		binary.Write(relaBuf, Le, int64(0))
	}
	l.RelaDynSect = relaBuf.Bytes()

	// Assemble Data Section with padding
	buf := new(bytes.Buffer)
	buf.Write(l.DynSect)
	for uint64(buf.Len()) < offsetSym { buf.WriteByte(0) }
	buf.Write(l.DynSymSect)
	for uint64(buf.Len()) < offsetStr { buf.WriteByte(0) }
	buf.Write(l.DynStrTab)
	for uint64(buf.Len()) < offsetRela { buf.WriteByte(0) }
	buf.Write(l.RelaDynSect)
	for uint64(buf.Len()) < offsetGot { buf.WriteByte(0) }
	buf.Write(make([]byte, sizeGot))
	l.DataSection = buf.Bytes()

	// Generate PLT
	pltBuf := new(bytes.Buffer)
	for i, symName := range l.GotEntries {
		targetGot := gotAddr + uint64(i*8)
		currentPC := l.TextAddr + pltOffset + uint64(pltBuf.Len())
		rel := int32(targetGot - (currentPC + 6))
		pltBuf.WriteByte(0xFF); pltBuf.WriteByte(0x25)
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
		objText = append(objText, make([]byte, 32)...)
	}
	
	for _, obj := range l.Objects {
		for _, sec := range obj.Sections {
			if sec.Flags&SHF_EXECINSTR != 0 {
				pad := (16 - ((currentTextLen + uint64(len(objText))) % 16)) % 16
				for i := 0; i < int(pad); i++ {
					objText = append(objText, 0x90)
				}
				sec.OutputOffset = currentTextLen + uint64(len(objText))
				sec.VirtualAddress = l.TextAddr + sec.OutputOffset
				objText = append(objText, sec.Data...)
			}
		}
	}
	l.TextSection = append(fullText, objText...)

	currentDataLen := uint64(len(l.DataSection))
	for _, obj := range l.Objects {
		for _, sec := range obj.Sections {
			if sec.Flags&SHF_WRITE != 0 && sec.Type == SHT_PROGBITS {
				pad := (8 - ((currentDataLen) % 8)) % 8
				l.DataSection = append(l.DataSection, make([]byte, int(pad))...)
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
	
	for _, obj := range l.Objects {
		for _, sym := range obj.Symbols {
			if sym.Section != nil {
				if gSym, ok := l.GlobalTable[sym.Name]; ok {
					gSym.Value = sym.Section.VirtualAddress + sym.Value
				}
			}
		}
	}
}

func (l *Linker) applyRelocations() error {
	if l.GlobalTable[l.Config.Entry].Section == "stub" {
		stub := []byte{
			0x31, 0xed,                   // xor rbp, rbp
			0x49, 0x89, 0xd1,             // mov r9, rdx
			0x5e,                         // pop rsi
			0x48, 0x89, 0xe2,             // mov rdx, rsp
			0x48, 0x83, 0xe4, 0xf0,       // and rsp, -16
			0x48, 0x8d, 0x3d, 0, 0, 0, 0, // lea rdi, [rip + offset]
			0x48, 0x31, 0xc9,             // xor rcx, rcx
			0x45, 0x31, 0xc0,             // xor r8d, r8d
			0xe8, 0, 0, 0, 0,             // call __libc_start_main
			0xf4,                         // hlt
		}

		stubOffset := l.GlobalTable[l.Config.Entry].Value - l.TextAddr
		leaPC := l.GlobalTable[l.Config.Entry].Value + 20
		mainAddr := l.GlobalTable["main"].Value
		binary.LittleEndian.PutUint32(stub[16:], uint32(int32(mainAddr - leaPC)))

		callPC := l.GlobalTable[l.Config.Entry].Value + 31
		libcStartAddr := uint64(0)
		if s, ok := l.GlobalTable["__libc_start_main"]; ok {
			libcStartAddr = s.Value
		} else {
			return fmt.Errorf("libc not found")
		}
		binary.LittleEndian.PutUint32(stub[27:], uint32(int32(libcStartAddr - callPC)))
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
				if bufOff >= uint64(len(buf)) { continue }

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

func (l *Linker) write(path string) error {
	f, err := os.Create(path)
	if err != nil { return err }
	defer f.Close()

	hdrSize := uint64(64 + 56*5)
	textOff := align(hdrSize, 16)
	textSize := uint64(len(l.TextSection))
	dataOff := align(textOff+textSize, 4096)
	dataSize := uint64(len(l.DataSection))

	l.addShStr("")
	idxInterp := l.addShStr(".interp")
	idxText := l.addShStr(".text")
	idxDyn := l.addShStr(".dynamic")
	idxData := l.addShStr(".data")
	idxBss := l.addShStr(".bss")
	idxShstr := l.addShStr(".shstrtab")

	shStrOff := dataOff + dataSize
	shTableOff := shStrOff + uint64(len(l.ShStrTab))
	shNum := uint16(7)

	ehdr := Header{
		Type: ET_EXEC, Machine: EM_X86_64, Version: EV_CURRENT,
		Entry: l.EntryAddr, Phoff: 64, Ehsize: 64, Phentsize: 56, Phnum: 5,
		Shoff: shTableOff, Shentsize: 64, Shnum: shNum, Shstrndx: 6,
	}
	ehdr.Ident[0] = 0x7F; ehdr.Ident[1] = 'E'; ehdr.Ident[2] = 'L'; ehdr.Ident[3] = 'F'
	ehdr.Ident[4] = ELFCLASS64; ehdr.Ident[5] = ELFDATA2LSB; ehdr.Ident[6] = EV_CURRENT; ehdr.Ident[7] = 3

	binary.Write(f, Le, ehdr)

	binary.Write(f, Le, ProgHeader{Type: PT_PHDR, Flags: PF_R, Off: 64, Vaddr: l.Config.BaseAddr + 64, Paddr: l.Config.BaseAddr + 64, Filesz: 56 * 5, Memsz: 56 * 5, Align: 8})
	binary.Write(f, Le, ProgHeader{Type: PT_INTERP, Flags: PF_R, Off: textOff, Vaddr: l.TextAddr, Paddr: l.TextAddr, Filesz: uint64(len(l.InterpSect)), Memsz: uint64(len(l.InterpSect)), Align: 1})
	binary.Write(f, Le, ProgHeader{Type: PT_LOAD, Flags: PF_R | PF_X, Off: 0, Vaddr: l.Config.BaseAddr, Paddr: l.Config.BaseAddr, Filesz: textOff + textSize, Memsz: textOff + textSize, Align: 4096})
	binary.Write(f, Le, ProgHeader{Type: PT_LOAD, Flags: PF_R | PF_W, Off: dataOff, Vaddr: l.DataAddr, Paddr: l.DataAddr, Filesz: dataSize, Memsz: dataSize + l.BssSize, Align: 4096})
	binary.Write(f, Le, ProgHeader{Type: PT_DYNAMIC, Flags: PF_R | PF_W, Off: dataOff, Vaddr: l.DataAddr, Paddr: l.DataAddr, Filesz: uint64(len(l.DynSect)), Memsz: uint64(len(l.DynSect)), Align: 8})

	cur := uint64(64 + 56*5)
	if textOff > cur { f.Write(make([]byte, textOff-cur)) }
	f.Write(l.TextSection)

	cur = textOff + textSize
	if dataOff > cur { f.Write(make([]byte, dataOff-cur)) }
	f.Write(l.DataSection)
	f.Write(l.ShStrTab)

	binary.Write(f, Le, SectionHeader{})
	binary.Write(f, Le, SectionHeader{Name: idxInterp, Type: SHT_PROGBITS, Flags: SHF_ALLOC, Addr: l.TextAddr, Offset: textOff, Size: uint64(len(l.InterpSect)), Addralign: 1})
	binary.Write(f, Le, SectionHeader{Name: idxText, Type: SHT_PROGBITS, Flags: SHF_ALLOC | SHF_EXECINSTR, Addr: l.TextAddr, Offset: textOff, Size: textSize, Addralign: 16})
	binary.Write(f, Le, SectionHeader{Name: idxDyn, Type: SHT_PROGBITS, Flags: SHF_ALLOC | SHF_WRITE, Addr: l.DataAddr, Offset: dataOff, Size: uint64(len(l.DynSect)), Link: 0, Addralign: 8})
	binary.Write(f, Le, SectionHeader{Name: idxData, Type: SHT_PROGBITS, Flags: SHF_ALLOC | SHF_WRITE, Addr: l.DataAddr, Offset: dataOff, Size: dataSize, Addralign: 4096})
	binary.Write(f, Le, SectionHeader{Name: idxBss, Type: SHT_NOBITS, Flags: SHF_ALLOC | SHF_WRITE, Addr: l.BssAddr, Offset: dataOff + dataSize, Size: l.BssSize, Addralign: 8})
	binary.Write(f, Le, SectionHeader{Name: idxShstr, Type: SHT_STRTAB, Flags: 0, Addr: 0, Offset: shStrOff, Size: uint64(len(l.ShStrTab)), Addralign: 1})

	f.Chmod(0755)
	return nil
}