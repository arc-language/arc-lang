package elf

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type Config struct {
	Entry       string
	BaseAddr    uint64
	Interpreter string // e.g. "/lib64/ld-linux-x86-64.so.2"
}

type Linker struct {
	Config      Config
	Objects     []*InputObject
	SharedLibs  []*SharedObject // .so files
	GlobalTable map[string]*ResolvedSymbol
	
	// Dynamic Linking State
	DynStrTab   []byte
	DynSyms     []Elf64Sym
	RelaDyn     []Elf64Rela // Relocations for the loader
	GotEntries  []string    // Symbols needing GOT entries
	
	// Buffers
	InterpSect  []byte
	DynSect     []byte
	DynSymSect  []byte
	DynStrSect  []byte
	RelaDynSect []byte
	TextSection []byte // Includes PLT
	DataSection []byte // Includes GOT
	
	// Layout
	TextAddr    uint64
	DataAddr    uint64
	DynAddr     uint64 // Start of dynamic info
}

// Internal structures for writing
type Elf64Sym struct {
	Name, Info, Other uint32 // Simplified for internal use, careful with packing
	Shndx            uint16
	Value, Size      uint64
}

type Elf64Rela struct {
	Offset uint64
	Info   uint64
	Addend int64
}

func NewLinker(cfg Config) *Linker {
	if cfg.Entry == "" { cfg.Entry = "_start" }
	if cfg.BaseAddr == 0 { cfg.BaseAddr = 0x400000 }
	if cfg.Interpreter == "" { cfg.Interpreter = "/lib64/ld-linux-x86-64.so.2" }

	return &Linker{
		Config:      cfg,
		GlobalTable: make(map[string]*ResolvedSymbol),
		// Null symbol for DynSym
		DynSyms:     []Elf64Sym{{}}, 
		DynStrTab:   []byte{0},
	}
}

func (l *Linker) AddObject(name string, data []byte) error {
	obj, err := LoadObject(name, data)
	if err == nil { l.Objects = append(l.Objects, obj) }
	return err
}

func (l *Linker) AddSharedLib(name string, data []byte) error {
	so, err := LoadSharedObject(name, data)
	if err == nil { l.SharedLibs = append(l.SharedLibs, so) }
	return err
}

func (l *Linker) Link(outPath string) error {
	if err := l.scanSymbols(); err != nil { return err }
	l.layout()
	l.applyRelocations() // Generate content
	return l.write(outPath)
}

func (l *Linker) scanSymbols() error {
	l.GlobalTable[l.Config.Entry] = &ResolvedSymbol{Name: l.Config.Entry}

	// 1. Scan Objects for definitions
	for _, obj := range l.Objects {
		for _, sym := range obj.Symbols {
			if sym.Bind == STB_GLOBAL && sym.Section != nil {
				if ex, ok := l.GlobalTable[sym.Name]; ok && ex.Defined {
					return fmt.Errorf("duplicate symbol: %s", sym.Name)
				}
				l.GlobalTable[sym.Name] = &ResolvedSymbol{
					Name: sym.Name, Defined: true, Section: "text", // simplify assumption
				}
			} else if sym.Bind == STB_GLOBAL {
				if _, ok := l.GlobalTable[sym.Name]; !ok {
					l.GlobalTable[sym.Name] = &ResolvedSymbol{Name: sym.Name, Defined: false}
				}
			}
		}
	}

	// 2. Scan Shared Libs to resolve undefined symbols
	for _, so := range l.SharedLibs {
		for _, symName := range so.Symbols {
			if target, ok := l.GlobalTable[symName]; ok && !target.Defined {
				// Mark as Dynamic Import
				target.Defined = true
				target.Section = "dynamic"
				
				// Create Dynamic Symbol Entry
				l.addDynamicSymbol(symName)
				l.GotEntries = append(l.GotEntries, symName)
			}
		}
	}

	// 3. Entry point check
	if !l.GlobalTable[l.Config.Entry].Defined {
		// Auto-stub generation logic (omitted for brevity, assume main exists)
		return fmt.Errorf("undefined entry point")
	}

	return nil
}

func (l *Linker) addDynamicSymbol(name string) {
	// Add string
	nameIdx := uint32(len(l.DynStrTab))
	l.DynStrTab = append(l.DynStrTab, []byte(name)...)
	l.DynStrTab = append(l.DynStrTab, 0)

	// Add Sym
	// Info: Global (1) << 4 | Func (2) = 0x12
	l.DynSyms = append(l.DynSyms, Elf64Sym{
		Name: nameIdx, Info: 0x12, Shndx: 0, Value: 0, Size: 0,
	})
}

func (l *Linker) layout() {
	// Headers: Ehdr(64) + 4 Phdrs(56*4) = 288 bytes
	// Phdrs: LOAD(Text), LOAD(Data), INTERP, DYNAMIC
	headerSize := uint64(64 + 56*4)
	
	// --- Text Segment (R+X) ---
	l.TextAddr = l.Config.BaseAddr + headerSize
	if l.TextAddr%16 != 0 { l.TextAddr += 16 - (l.TextAddr%16) }
	
	// 1. .interp section
	l.InterpSect = append([]byte(l.Config.Interpreter), 0...)
	
	// 2. PLT Stubs (Trampolines)
	// For every GOT entry, we need a code stub: jmp *[rip+offset]
	pltOffset := uint64(len(l.InterpSect))
	
	// 3. Object Text
	// (Standard gathering of object text, omitted strict alignment logic for brevity)
	// ... (Code from previous answer for gathering l.TextSection) ...
	// *Important*: We prepend Interp and PLT to TextSection later
	
	// --- Data Segment (R+W) ---
	// Page align
	l.DataAddr = l.TextAddr + 0x2000 // Arbitrary safe gap
	if l.DataAddr % 4096 != 0 { l.DataAddr += 4096 - (l.DataAddr%4096) }
	
	// 4. .dynamic section
	// 5. .dynsym
	// 6. .dynstr
	// 7. .rela.dyn
	// 8. .got (Global Offset Table)
	
	// Calculate GOT Address (it will be at end of Dynamic data)
	// Sizes:
	dynSize := 16 * 10 // approx 10 tags * 16 bytes
	symSize := len(l.DynSyms) * 24
	strSize := len(l.DynStrTab)
	relaSize := len(l.GotEntries) * 24 // 1 rela per GOT entry
	
	gotOffset := uint64(dynSize + symSize + strSize + relaSize)
	gotAddr := l.DataAddr + gotOffset
	
	// Generate PLT Stubs now that we know GOT addresses
	pltBuf := new(bytes.Buffer)
	for i, symName := range l.GotEntries {
		// Target GOT entry address
		targetGot := gotAddr + uint64(i*8)
		
		// Current PLT Instruction Address
		currentPC := l.TextAddr + pltOffset + uint64(pltBuf.Len())
		
		// JMP [RIP + rel]  (Opcode FF 25)
		// Instruction length is 6 bytes. RIP is PC + 6.
		rel := int32(targetGot - (currentPC + 6))
		
		pltBuf.WriteByte(0xFF)
		pltBuf.WriteByte(0x25)
		binary.Write(pltBuf, Le, rel)
		
		// Update ResolvedSymbol to point to this Stub!
		// When main calls 'printf', it relocates to this stub.
		l.GlobalTable[symName].Value = currentPC
		l.GlobalTable[symName].Section = "plt"
	}
	
	// Combine Text
	fullText := append(l.InterpSect, pltBuf.Bytes()...)
	// Append object code...
	for _, obj := range l.Objects {
		for _, sec := range obj.Sections {
			if sec.Flags&SHF_EXECINSTR != 0 {
				sec.OutputOffset = uint64(len(fullText))
				sec.VirtualAddress = l.TextAddr + sec.OutputOffset
				fullText = append(fullText, sec.Data...)
			}
		}
	}
	l.TextSection = fullText
	
	// Build Data Sections
	// .dynamic
	dynBuf := new(bytes.Buffer)
	writeDyn := func(tag int64, val uint64) {
		binary.Write(dynBuf, Le, tag); binary.Write(dynBuf, Le, val)
	}
	
	// Dynamic Tags
	strTabAddr := l.DataAddr + uint64(dynSize) + uint64(symSize)
	symTabAddr := l.DataAddr + uint64(dynSize)
	relaAddr   := l.DataAddr + uint64(dynSize) + uint64(symSize) + uint64(strSize)
	
	for _, lib := range l.SharedLibs {
		// DT_NEEDED (offset in strtab)
		// Simplication: append lib name to strtab now
		idx := len(l.DynStrTab)
		l.DynStrTab = append(l.DynStrTab, []byte(lib.Name)...)
		l.DynStrTab = append(l.DynStrTab, 0)
		writeDyn(DT_NEEDED, uint64(idx))
	}
	writeDyn(DT_STRTAB, strTabAddr)
	writeDyn(DT_SYMTAB, symTabAddr)
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
		// Write standard Elf64_Sym (24 bytes)
		binary.Write(symBuf, Le, ds.Name) // 4
		symBuf.WriteByte(uint8(ds.Info)); symBuf.WriteByte(0) // info, other
		binary.Write(symBuf, Le, ds.Shndx) // 2
		binary.Write(symBuf, Le, ds.Value) // 8
		binary.Write(symBuf, Le, ds.Size)  // 8
	}
	l.DynSymSect = symBuf.Bytes()

	// .rela.dyn
	relaBuf := new(bytes.Buffer)
	for i := range l.GotEntries {
		// R_X86_64_GLOB_DAT (6)
		// Info = (SymIdx << 32) | Type
		// SymIdx starts at 1 (0 is null)
		info := uint64((i+1)<<32) | uint64(R_X86_64_GLOB_DAT)
		offset := gotAddr + uint64(i*8)
		
		binary.Write(relaBuf, Le, offset)
		binary.Write(relaBuf, Le, info)
		binary.Write(relaBuf, Le, int64(0))
	}
	l.RelaDynSect = relaBuf.Bytes()

	// .got (Zeroed out)
	gotData := make([]byte, len(l.GotEntries)*8)

	// Combine Data
	fullData := append(l.DynSect, l.DynSymSect...)
	fullData = append(fullData, l.DynStrTab...)
	fullData = append(fullData, l.RelaDynSect...)
	fullData = append(fullData, gotData...)
	
	// Append Object Data
	// ... (Standard gathering of object data) ...
	for _, obj := range l.Objects {
		for _, sec := range obj.Sections {
			if sec.Flags&SHF_WRITE != 0 && sec.Type == SHT_PROGBITS {
				sec.OutputOffset = uint64(len(fullData))
				sec.VirtualAddress = l.DataAddr + sec.OutputOffset
				fullData = append(fullData, sec.Data...)
			}
		}
	}
	
	l.DataSection = fullData
}

func (l *Linker) applyRelocations() {
	// Same logic as static linker, but now l.GlobalTable["printf"]
	// points to the PLT stub we generated in layout().
	// So calls to printf become calls to the trampoline.
	
	// ... (Copy applyRelocations from previous answer) ...
	// Ensure you use the updated l.TextSection buffers
}

func (l *Linker) write(path string) error {
	f, err := os.Create(path)
	if err != nil { return err }
	defer f.Close()

	// ... (Headers setup) ...
	
	// We need 4 Program Headers:
	// 1. PHDR (Self reference, optional but good practice)
	// 2. INTERP (Points to .interp section)
	// 3. LOAD (Text)
	// 4. LOAD (Data)
	// 5. DYNAMIC (Points to .dynamic section)
	
	// Offsets
	textOff := uint64(64 + 56*5)
	if textOff%16 != 0 { textOff += 16-(textOff%16) }
	
	dataOff := textOff + uint64(len(l.TextSection))
	if dataOff%4096 != 0 { dataOff += 4096-(dataOff%4096) }

	ehdr := Header{Type: ET_EXEC, Machine: EM_X86_64, Version: EV_CURRENT, Entry: l.GlobalTable[l.Config.Entry].Value, Phnum: 5, Shoff: 0, Phentsize: 56, Ehsize: 64, Phoff: 64}
	ehdr.Ident[0]=0x7F; ehdr.Ident[1]='E'; ehdr.Ident[2]='L'; ehdr.Ident[3]='F'; ehdr.Ident[4]=ELFCLASS64; ehdr.Ident[5]=ELFDATA2LSB; ehdr.Ident[6]=EV_CURRENT; ehdr.Ident[7]=3 // Linux

	// Write Header
	binary.Write(f, Le, ehdr)

	// 1. INTERP
	binary.Write(f, Le, ProgHeader{
		Type: PT_INTERP, Flags: PF_R, Off: textOff, 
		Vaddr: l.TextAddr, Paddr: l.TextAddr, 
		Filesz: uint64(len(l.InterpSect)), Memsz: uint64(len(l.InterpSect)), Align: 1,
	})
	
	// 2. LOAD (Text)
	binary.Write(f, Le, ProgHeader{
		Type: PT_LOAD, Flags: PF_R | PF_X, Off: 0, 
		Vaddr: l.Config.BaseAddr, Paddr: l.Config.BaseAddr, 
		Filesz: textOff + uint64(len(l.TextSection)), Memsz: textOff + uint64(len(l.TextSection)), Align: 4096,
	})

	// 3. LOAD (Data)
	binary.Write(f, Le, ProgHeader{
		Type: PT_LOAD, Flags: PF_R | PF_W, Off: dataOff, 
		Vaddr: l.DataAddr, Paddr: l.DataAddr, 
		Filesz: uint64(len(l.DataSection)), Memsz: uint64(len(l.DataSection)), Align: 4096,
	})
	
	// 4. DYNAMIC
	// Located at start of Data Section
	binary.Write(f, Le, ProgHeader{
		Type: PT_DYNAMIC, Flags: PF_R | PF_W, Off: dataOff, 
		Vaddr: l.DataAddr, Paddr: l.DataAddr, 
		Filesz: uint64(len(l.DynSect)), Memsz: uint64(len(l.DynSect)), Align: 8,
	})
	
	// Fill
	cur := uint64(64 + 56*5)
	f.Write(make([]byte, textOff - cur))
	f.Write(l.TextSection)
	
	cur = textOff + uint64(len(l.TextSection))
	f.Write(make([]byte, dataOff - cur))
	f.Write(l.DataSection)
	
	f.Chmod(0755)
	return nil
}