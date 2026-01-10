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
	DynStrTab   []byte
	DynSyms     []Elf64Sym // Uses struct from const.go
	RelaDyn     []Elf64Rela // Uses struct from const.go
	GotEntries  []string
	
	// Output Buffers
	InterpSect  []byte
	DynSect     []byte
	DynSymSect  []byte
	DynStrSect  []byte
	RelaDynSect []byte
	TextSection []byte // Includes PLT
	DataSection []byte // Includes GOT
	
	// Calculated Addresses
	TextAddr    uint64
	DataAddr    uint64
	BssAddr     uint64
	EntryAddr   uint64
	DynAddr     uint64
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
	if cfg.Entry == "" { cfg.Entry = "_start" }
	if cfg.BaseAddr == 0 { cfg.BaseAddr = 0x400000 }
	if cfg.Interpreter == "" { cfg.Interpreter = "/lib64/ld-linux-x86-64.so.2" }

	return &Linker{
		Config:      cfg,
		GlobalTable: make(map[string]*ResolvedSymbol),
		// Init DynSyms with Null Symbol
		DynSyms:     []Elf64Sym{{}}, 
		DynStrTab:   []byte{0},
	}
}

// AddObject adds an in-memory object file
func (l *Linker) AddObject(name string, data []byte) error {
	obj, err := LoadObject(name, data)
	if err != nil { return err }
	l.Objects = append(l.Objects, obj)
	return nil
}

// AddArchive adds a .a file from disk
func (l *Linker) AddArchive(path string) error {
	objs, err := LoadArchive(path)
	if err != nil { return err }
	l.Objects = append(l.Objects, objs...)
	return nil
}

// AddSharedLib parses a .so file for dynamic symbol resolution
func (l *Linker) AddSharedLib(path string, data []byte) error {
	so, err := LoadSharedObject(path, data)
	if err != nil { return err }
	l.SharedLibs = append(l.SharedLibs, so)
	return nil
}

// Link performs the linking process and writes the executable
func (l *Linker) Link(outPath string) error {
	if err := l.scanSymbols(); err != nil { return err }
	l.layout()
	if err := l.applyRelocations(); err != nil { return err }
	return l.write(outPath)
}

func (l *Linker) scanSymbols() error {
	// Ensure entry point exists in table
	l.GlobalTable[l.Config.Entry] = &ResolvedSymbol{Name: l.Config.Entry, Defined: false}

	// 1. Scan Objects for definitions
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

	// 2. Scan Shared Libs to resolve undefined symbols
	for _, so := range l.SharedLibs {
		for _, symName := range so.Symbols {
			if target, ok := l.GlobalTable[symName]; ok && !target.Defined {
				// Found in shared lib -> Mark as Dynamic
				target.Defined = true
				target.Section = "dynamic"
				
				// Register for GOT/PLT generation
				l.addDynamicSymbol(symName)
				l.GotEntries = append(l.GotEntries, symName)
			}
		}
	}
	
	// 3. Entry Point Check
	// If _start is missing but main exists, we generate a stub later. 
	// If main is also missing, error.
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
	// Global(1) << 4 | Func(2) = 0x12
	l.DynSyms = append(l.DynSyms, Elf64Sym{
		Name: nameIdx, Info: 0x12, Shndx: 0, Value: 0, Size: 0,
	})
}

func (l *Linker) layout() {
	// Headers: Ehdr(64) + 5 Phdrs(56*5) = 344 bytes
	// Phdrs: INTERP, LOAD(Text), LOAD(Data), DYNAMIC, PHDR
	headerSize := uint64(64 + 56*5)

	// --- Text Segment (R+X) ---
	l.TextAddr = l.Config.BaseAddr + headerSize
	if l.TextAddr%16 != 0 {
		l.TextAddr += 16 - (l.TextAddr % 16)
	}

	// 1. .interp section (Fix: append 0 directly, not 0...)
	l.InterpSect = append([]byte(l.Config.Interpreter), 0)

	// 2. PLT Stubs (Placeholder for now, calculated after GOT address known)
	pltOffset := uint64(len(l.InterpSect))

	// 3. Object Text (Gathering)
	var objText []byte

	// Reserve space for Entry Stub if needed
	if l.GlobalTable[l.Config.Entry].Section == "stub" {
		objText = append(objText, make([]byte, 29)...) // 29 bytes for stub
		l.GlobalTable[l.Config.Entry].Value = l.TextAddr + pltOffset // Temp value
	}

	for _, obj := range l.Objects {
		for _, sec := range obj.Sections {
			if sec.Flags&SHF_EXECINSTR != 0 {
				// Align 16
				pad := (16 - (len(objText) % 16)) % 16
				objText = append(objText, make([]byte, pad)...)

				sec.OutputOffset = pltOffset + uint64(len(objText)) // Relative to TextStart
				sec.VirtualAddress = l.TextAddr + sec.OutputOffset
				objText = append(objText, sec.Data...)
			}
		}
	}

	// --- Data Segment (R+W) ---
	// Page Align from end of Text
	// We estimate PLT size = len(GotEntries) * 6
	estPltSize := uint64(len(l.GotEntries) * 6)
	totalTextSize := uint64(len(l.InterpSect)) + estPltSize + uint64(len(objText))

	l.DataAddr = l.TextAddr + totalTextSize
	if l.DataAddr%4096 != 0 {
		l.DataAddr += 4096 - (l.DataAddr % 4096)
	}

	// 4. Dynamic Sections
	// Structure: [Dynamic Tags] [SymTab] [StrTab] [Rela] [GOT]

	dynSize := 16 * 15 // ~15 tags
	symSize := len(l.DynSyms) * 24
	strSize := len(l.DynStrTab)
	relaSize := len(l.GotEntries) * 24

	gotOffset := uint64(dynSize + symSize + strSize + relaSize)
	gotAddr := l.DataAddr + gotOffset

	// Now generate PLT Stubs with known GOT address
	pltBuf := new(bytes.Buffer)
	for i, symName := range l.GotEntries {
		targetGot := gotAddr + uint64(i*8)
		currentPC := l.TextAddr + pltOffset + uint64(pltBuf.Len())

		// JMP [RIP + rel] -> FF 25 rel32
		rel := int32(targetGot - (currentPC + 6))

		pltBuf.WriteByte(0xFF)
		pltBuf.WriteByte(0x25)
		binary.Write(pltBuf, Le, rel)

		// Update symbol to point to this PLT stub
		l.GlobalTable[symName].Value = currentPC
		l.GlobalTable[symName].Section = "plt"
	}

	// Re-calculate Text Layout with actual PLT
	actualPlt := pltBuf.Bytes()
	fullText := append(l.InterpSect, actualPlt...)

	// Re-align text start for objects
	currentTextLen := uint64(len(fullText))

	// Re-layout objects
	objText = nil // reset

	// Stub
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

	// Re-align Data Address based on final Text size
	l.DataAddr = l.TextAddr + uint64(len(l.TextSection))
	if l.DataAddr%4096 != 0 {
		l.DataAddr += 4096 - (l.DataAddr % 4096)
	}

	// Recalculate GOT addr based on final DataAddr
	gotAddr = l.DataAddr + gotOffset // Offset inside data block is constant

	// Generate Dynamic Data
	dynBuf := new(bytes.Buffer)
	writeDyn := func(tag int64, val uint64) {
		binary.Write(dynBuf, Le, tag)
		binary.Write(dynBuf, Le, val)
	}

	// Dynamic Offsets (Virtual Addresses)
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

	// .dynsym
	symBuf := new(bytes.Buffer)
	for _, ds := range l.DynSyms {
		binary.Write(symBuf, Le, ds.Name)
		symBuf.WriteByte(uint8(ds.Info))
		symBuf.WriteByte(uint8(ds.Other))
		binary.Write(symBuf, Le, ds.Shndx)
		binary.Write(symBuf, Le, ds.Value)
		binary.Write(symBuf, Le, ds.Size)
	}
	l.DynSymSect = symBuf.Bytes()

	// .rela.dyn
	relaBuf := new(bytes.Buffer)
	for i := range l.GotEntries {
		offset := gotAddr + uint64(i*8)
		info := uint64((i+1)<<32) | uint64(R_X86_64_GLOB_DAT)
		binary.Write(relaBuf, Le, offset)
		binary.Write(relaBuf, Le, info)
		binary.Write(relaBuf, Le, int64(0))
	}
	l.RelaDynSect = relaBuf.Bytes()

	// .got
	gotData := make([]byte, len(l.GotEntries)*8)

	// Assemble Data Section
	l.DataSection = append(l.DynSect, l.DynSymSect...)
	l.DataSection = append(l.DataSection, l.DynStrTab...)
	l.DataSection = append(l.DataSection, l.RelaDynSect...)
	l.DataSection = append(l.DataSection, gotData...)

	// Append Object Data
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

	// Layout BSS
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
	// 1. Generate Stub if needed
	if l.GlobalTable[l.Config.Entry].Section == "stub" {
		stub := []byte{
			0xE8, 0, 0, 0, 0, // call main
			0x48, 0x31, 0xFF, // xor rdi, rdi
			0x48, 0xC7, 0xC0, 0x3C, 0x00, 0x00, 0x00, // mov rax, 60
			0x0F, 0x05, // syscall
		}
		
		// Entry Address is the start of the stub in objText (which is appended to TextSection)
		// We need to find where the stub resides in l.TextSection.
		// It was appended immediately after fullText (Interp+PLT).
		
		stubOffset := l.GlobalTable[l.Config.Entry].Value - l.TextAddr
		
		// Patch 'call main'
		pc := l.TextAddr + stubOffset + 5
		target := l.GlobalTable["main"].Value
		binary.LittleEndian.PutUint32(stub[1:], uint32(int32(target - pc)))
		
		// Copy into TextSection
		copy(l.TextSection[stubOffset:], stub)
	}

	// 2. Apply Relocs
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
				
				if sec.Flags & SHF_EXECINSTR != 0 {
					buf = l.TextSection
					// sec.OutputOffset is relative to l.TextAddr
					// buf index is simply OutputOffset
					bufOff = sec.OutputOffset + r.Offset
				} else if sec.Flags & SHF_WRITE != 0 {
					buf = l.DataSection
					bufOff = sec.OutputOffset + r.Offset
				} else {
					continue
				}

				if bufOff >= uint64(len(buf)) {
					// Should not happen if layout is correct
					continue 
				}

				switch r.Type {
				case R_X86_64_64:
					binary.LittleEndian.PutUint64(buf[bufOff:], uint64(int64(symVal) + r.Addend))
				case R_X86_64_32:
					binary.LittleEndian.PutUint32(buf[bufOff:], uint32(int64(symVal) + r.Addend))
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

	// 5 Phdrs
	hdrSize := uint64(64 + 56*5)
	
	// File Offsets
	textOff := hdrSize
	if textOff%16 != 0 { textOff += 16 - (textOff%16) }
	
	textSize := uint64(len(l.TextSection))
	dataOff := textOff + textSize
	if dataOff%4096 != 0 { dataOff += 4096 - (dataOff%4096) }

	ehdr := Header{
		Type: ET_EXEC, Machine: EM_X86_64, Version: EV_CURRENT, 
		Entry: l.EntryAddr, Phoff: 64, Ehsize: 64, Phentsize: 56, Phnum: 5, 
		Shoff: 0, Shentsize: 64, Shnum: 0, Shstrndx: 0,
	}
	ehdr.Ident[0]=0x7F; ehdr.Ident[1]='E'; ehdr.Ident[2]='L'; ehdr.Ident[3]='F'
	ehdr.Ident[4]=ELFCLASS64; ehdr.Ident[5]=ELFDATA2LSB; ehdr.Ident[6]=EV_CURRENT; ehdr.Ident[7]=3

	binary.Write(f, Le, ehdr)

	// 1. INTERP (Points to start of Text -> Interp string)
	binary.Write(f, Le, ProgHeader{
		Type: PT_INTERP, Flags: PF_R, Off: textOff, 
		Vaddr: l.TextAddr, Paddr: l.TextAddr, 
		Filesz: uint64(len(l.InterpSect)), Memsz: uint64(len(l.InterpSect)), Align: 1,
	})
	
	// 2. LOAD Text
	binary.Write(f, Le, ProgHeader{
		Type: PT_LOAD, Flags: PF_R | PF_X, Off: 0, 
		Vaddr: l.Config.BaseAddr, Paddr: l.Config.BaseAddr, 
		Filesz: textOff + textSize, Memsz: textOff + textSize, Align: 4096,
	})
	
	// 3. LOAD Data
	binary.Write(f, Le, ProgHeader{
		Type: PT_LOAD, Flags: PF_R | PF_W, Off: dataOff, 
		Vaddr: l.DataAddr, Paddr: l.DataAddr, 
		Filesz: uint64(len(l.DataSection)), Memsz: uint64(len(l.DataSection)) + l.BssSize, Align: 4096,
	})
	
	// 4. DYNAMIC (Points to start of Data -> Dynamic Section)
	binary.Write(f, Le, ProgHeader{
		Type: PT_DYNAMIC, Flags: PF_R | PF_W, Off: dataOff, 
		Vaddr: l.DataAddr, Paddr: l.DataAddr, 
		Filesz: uint64(len(l.DynSect)), Memsz: uint64(len(l.DynSect)), Align: 8,
	})
	
	// 5. PHDR (Self-Ref)
	binary.Write(f, Le, ProgHeader{
		Type: 6, Flags: PF_R, Off: 64, // PT_PHDR = 6
		Vaddr: l.Config.BaseAddr + 64, Paddr: l.Config.BaseAddr + 64,
		Filesz: 56*5, Memsz: 56*5, Align: 8,
	})

	// Pad and Write
	cur := uint64(64 + 56*5)
	if textOff > cur { f.Write(make([]byte, textOff - cur)) }
	f.Write(l.TextSection)
	
	cur = textOff + textSize
	if dataOff > cur { f.Write(make([]byte, dataOff - cur)) }
	f.Write(l.DataSection)
	
	f.Chmod(0755)
	return nil
}