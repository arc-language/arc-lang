package elf

import (
	"bytes"
	"encoding/binary"
	"io"
)

// ELF64 Constants
const (
	// ELF Header
	EI_NIDENT   = 16
	EI_MAG0     = 0
	ELFMAG0     = 0x7f
	ELFMAG1     = 'E'
	ELFMAG2     = 'L'
	ELFMAG3     = 'F'
	EI_CLASS    = 4
	ELFCLASS64  = 2
	EI_DATA     = 5
	ELFDATA2LSB = 1
	EI_VERSION  = 6
	EV_CURRENT  = 1

	// Object file types
	ET_NONE = 0
	ET_REL  = 1
	ET_EXEC = 2
	ET_DYN  = 3
	ET_CORE = 4

	// Machine types
	EM_X86_64 = 62

	// Section types
	SHT_NULL     = 0
	SHT_PROGBITS = 1
	SHT_SYMTAB   = 2
	SHT_STRTAB   = 3
	SHT_RELA     = 4
	SHT_HASH     = 5
	SHT_DYNAMIC  = 6
	SHT_NOTE     = 7
	SHT_NOBITS   = 8
	SHT_REL      = 9

	// Section flags
	SHF_WRITE     = 0x1
	SHF_ALLOC     = 0x2
	SHF_EXECINSTR = 0x4
	SHF_MERGE     = 0x10
	SHF_STRINGS   = 0x20
	SHF_INFO_LINK = 0x40

	// Segment Types (Program Header)
	PT_NULL    = 0
	PT_LOAD    = 1
	PT_DYNAMIC = 2
	PT_INTERP  = 3
	PT_NOTE    = 4
	PT_SHLIB   = 5
	PT_PHDR    = 6

	// Segment Flags
	PF_X = 0x1 // Execute
	PF_W = 0x2 // Write
	PF_R = 0x4 // Read

	// Symbol binding
	STB_LOCAL  = 0
	STB_GLOBAL = 1
	STB_WEAK   = 2

	// Symbol types
	STT_NOTYPE  = 0
	STT_OBJECT  = 1
	STT_FUNC    = 2
	STT_SECTION = 3
	STT_FILE    = 4
	STT_COMMON  = 5
	STT_TLS     = 6

	// Symbol visibility
	STV_DEFAULT   = 0
	STV_INTERNAL  = 1
	STV_HIDDEN    = 2
	STV_PROTECTED = 3

	// Special section indices
	SHN_UNDEF = 0
	SHN_ABS   = 0xfff1

	// Relocation types for x86-64
	R_X86_64_NONE  = 0
	R_X86_64_64    = 1
	R_X86_64_PC32  = 2
	R_X86_64_GOT32 = 3
	R_X86_64_PLT32 = 4
	R_X86_64_COPY  = 5
	R_X86_64_32    = 10
	R_X86_64_32S   = 11
	R_X86_64_16    = 12
	R_X86_64_PC16  = 13
	R_X86_64_8     = 14
	R_X86_64_PC8   = 15
	R_X86_64_PC64  = 24
)

// File represents an ELF object file
type File struct {
	Type           uint16
	Sections       []*Section
	Symbols        []*Symbol
	ProgramHeaders []*ProgramHeader
	StrTab         *StringTable
	ShStrTab       *StringTable
	DataLayout     string
	Machine        uint16
	Entry          uint64 // Entry point virtual address
	RelaSections   []*Section
}

// Section represents an ELF section
type Section struct {
	Name      string
	Type      uint32
	Flags     uint64
	Addr      uint64
	Addralign uint64
	Entsize   uint64
	Link      uint32
	Info      uint32
	Content   []byte

	// Internal
	Index   uint16
	nameIdx uint32
	offset  uint64
	size    uint64
}

// ProgramHeader represents an ELF segment
type ProgramHeader struct {
	Type   uint32
	Flags  uint32
	Off    uint64
	Vaddr  uint64
	Paddr  uint64
	Filesz uint64
	Memsz  uint64
	Align  uint64
}

// Symbol represents an ELF symbol
type Symbol struct {
	Name    string
	Info    byte // Binding (high 4 bits) | Type (low 4 bits)
	Other   byte // Visibility
	Section *Section
	Value   uint64
	Size    uint64

	// Internal
	nameIdx uint32
	symIdx  int // Index in final symbol table
}

// Relocation represents a relocation entry
type Relocation struct {
	Offset uint64
	Symbol *Symbol
	Type   uint32
	Addend int64
}

// StringTable manages string storage
type StringTable struct {
	Data []byte
	strs map[string]uint32 // Deduplication
}

func NewStringTable() *StringTable {
	return &StringTable{
		Data: []byte{0}, // Null string at index 0
		strs: make(map[string]uint32),
	}
}

func (st *StringTable) Add(s string) uint32 {
	if s == "" {
		return 0
	}

	// Check if already exists
	if idx, exists := st.strs[s]; exists {
		return idx
	}

	idx := uint32(len(st.Data))
	st.Data = append(st.Data, []byte(s)...)
	st.Data = append(st.Data, 0)
	st.strs[s] = idx
	return idx
}

// NewFile creates a new ELF object file
func NewFile() *File {
	f := &File{
		Type:     ET_REL, // Default to relocatable
		StrTab:   NewStringTable(),
		ShStrTab: NewStringTable(),
		Machine:  EM_X86_64,
	}

	// Section 0 is always the null section
	f.Sections = append(f.Sections, &Section{
		Name: "",
		Type: SHT_NULL,
	})

	return f
}

// AddSection adds a new section
func (f *File) AddSection(name string, typ uint32, flags uint64, content []byte) *Section {
	s := &Section{
		Name:    name,
		Type:    typ,
		Flags:   flags,
		Content: content,
		Index:   uint16(len(f.Sections)),
	}

	f.Sections = append(f.Sections, s)
	return s
}

// AddProgramHeader adds a new segment
func (f *File) AddProgramHeader(typ, flags uint32, off, vaddr, filesz, memsz, align uint64) *ProgramHeader {
	ph := &ProgramHeader{
		Type:   typ,
		Flags:  flags,
		Off:    off,
		Vaddr:  vaddr,
		Paddr:  vaddr, // Usually same as Vaddr on modern systems
		Filesz: filesz,
		Memsz:  memsz,
		Align:  align,
	}
	f.ProgramHeaders = append(f.ProgramHeaders, ph)
	return ph
}

// AddSymbol adds a new symbol
func (f *File) AddSymbol(name string, info byte, section *Section, value, size uint64) *Symbol {
	sym := &Symbol{
		Name:    name,
		Info:    info,
		Other:   STV_DEFAULT,
		Section: section,
		Value:   value,
		Size:    size,
		symIdx:  -1, // Will be set when writing
	}

	f.Symbols = append(f.Symbols, sym)
	return sym
}

// WriteTo writes the complete ELF file
func (f *File) WriteTo(w io.Writer) error {
	// 1. Add string table sections FIRST (before building string tables)
	shstrtabSec := f.AddSection(".shstrtab", SHT_STRTAB, 0, nil)
	strTabSec := f.AddSection(".strtab", SHT_STRTAB, 0, nil)
	strTabSec.Addralign = 1

	// 2. Build string tables for symbols
	for _, sym := range f.Symbols {
		sym.nameIdx = f.StrTab.Add(sym.Name)
	}

	// 3. Build symbol table with correct ordering
	symBuf := new(bytes.Buffer)
	orderedSymbols := make([]*Symbol, 0, len(f.Symbols)+1)

	// First symbol is always null
	nullSym := &Symbol{}
	f.writeSymbol(symBuf, nullSym)
	orderedSymbols = append(orderedSymbols, nullSym)

	// Track first global symbol index
	firstGlobal := 1

	// Write local symbols first
	for _, sym := range f.Symbols {
		binding := sym.Info >> 4
		if binding == STB_LOCAL {
			sym.symIdx = len(orderedSymbols)
			f.writeSymbol(symBuf, sym)
			orderedSymbols = append(orderedSymbols, sym)
		}
	}

	firstGlobal = len(orderedSymbols)

	// Write global symbols
	for _, sym := range f.Symbols {
		binding := sym.Info >> 4
		if binding != STB_LOCAL {
			sym.symIdx = len(orderedSymbols)
			f.writeSymbol(symBuf, sym)
			orderedSymbols = append(orderedSymbols, sym)
		}
	}

	symTabSec := f.AddSection(".symtab", SHT_SYMTAB, 0, symBuf.Bytes())
	symTabSec.Link = uint32(strTabSec.Index)
	symTabSec.Info = uint32(firstGlobal)
	symTabSec.Addralign = 8
	symTabSec.Entsize = 24

	// 4. Fix up relocation section links
	for _, relaSec := range f.RelaSections {
		relaSec.Link = uint32(symTabSec.Index)
	}

	// 5. Build section name string table
	for _, sec := range f.Sections {
		sec.nameIdx = f.ShStrTab.Add(sec.Name)
	}

	// Set content for string tables
	shstrtabSec.Content = f.ShStrTab.Data
	shstrtabSec.size = uint64(len(f.ShStrTab.Data))
	strTabSec.Content = f.StrTab.Data
	strTabSec.size = uint64(len(f.StrTab.Data))

	// 6. Calculate ELF Header and Program Header sizes
	headerSize := uint64(64) // sizeof(Elf64_Ehdr)
	phdrSize := uint64(56)   // sizeof(Elf64_Phdr)
	phOff := headerSize
	
	if len(f.ProgramHeaders) > 0 {
		// Program headers follow ELF header
		headerSize += uint64(len(f.ProgramHeaders)) * phdrSize
	}

	// 7. Calculate section offsets
	// If ET_EXEC, sections might be placed at specific addresses controlled by ProgramHeaders
	// For now, we assume sections are packed after headers unless manually addressed.
	currentOffset := headerSize

	for _, sec := range f.Sections {
		// If section has a fixed offset (from static linking), use it if valid
		// Otherwise calculate linear layout
		if sec.offset == 0 {
			if sec.Addralign > 0 {
				if currentOffset%sec.Addralign != 0 {
					currentOffset += sec.Addralign - (currentOffset % sec.Addralign)
				}
			}
			sec.offset = currentOffset
		} else {
			// Section offset already set (e.g. by linker)
			currentOffset = sec.offset
		}

		if sec.size == 0 {
			sec.size = uint64(len(sec.Content))
		}
		currentOffset += sec.size
	}

	shdrOffset := currentOffset

	// 8. Write ELF header
	if err := f.writeElfHeader(w, shdrOffset, phOff, shstrtabSec.Index); err != nil {
		return err
	}

	// 9. Write Program Headers
	for _, ph := range f.ProgramHeaders {
		if err := f.writeProgramHeader(w, ph); err != nil {
			return err
		}
	}

	// 10. Write section contents
	written := headerSize
	
	// Sort sections by offset to write sequentially? 
	// For simplicity, we assume Sections array order mostly matches offset order or we fill gaps.
	// A robust writer would sort. Here we assume simple append order.
	
	for _, sec := range f.Sections {
		if sec.Type == SHT_NULL || sec.Type == SHT_NOBITS {
			continue
		}

		// Add padding
		if sec.offset > written {
			padding := make([]byte, sec.offset-written)
			if _, err := w.Write(padding); err != nil {
				return err
			}
			written = sec.offset
		}

		if _, err := w.Write(sec.Content); err != nil {
			return err
		}
		written += sec.size
	}

	// Pad to Section Header Table
	if shdrOffset > written {
		padding := make([]byte, shdrOffset-written)
		w.Write(padding)
	}

	// 11. Write section headers
	for _, sec := range f.Sections {
		if err := f.writeSectionHeader(w, sec); err != nil {
			return err
		}
	}

	return nil
}

func (f *File) writeElfHeader(w io.Writer, shoff, phoff uint64, shstrndx uint16) error {
	var hdr elfHeader

	hdr.Ident[EI_MAG0] = ELFMAG0
	hdr.Ident[1] = ELFMAG1
	hdr.Ident[2] = ELFMAG2
	hdr.Ident[3] = ELFMAG3
	hdr.Ident[EI_CLASS] = ELFCLASS64
	hdr.Ident[EI_DATA] = ELFDATA2LSB
	hdr.Ident[EI_VERSION] = EV_CURRENT
	hdr.Ident[7] = 0 // ABI (System V)

	hdr.Type = f.Type
	hdr.Machine = f.Machine
	hdr.Version = EV_CURRENT
	hdr.Entry = f.Entry
	hdr.Phoff = 0
	hdr.Phnum = 0
	
	if len(f.ProgramHeaders) > 0 {
		hdr.Phoff = phoff
		hdr.Phnum = uint16(len(f.ProgramHeaders))
	}

	hdr.Shoff = shoff
	hdr.Flags = 0
	hdr.Ehsize = 64
	hdr.Phentsize = 56
	hdr.Shentsize = 64
	hdr.Shnum = uint16(len(f.Sections))
	hdr.Shstrndx = shstrndx

	return binary.Write(w, binary.LittleEndian, hdr)
}

func (f *File) writeProgramHeader(w io.Writer, ph *ProgramHeader) error {
	var p elfProgramHeader
	p.Type = ph.Type
	p.Flags = ph.Flags
	p.Off = ph.Off
	p.Vaddr = ph.Vaddr
	p.Paddr = ph.Paddr
	p.Filesz = ph.Filesz
	p.Memsz = ph.Memsz
	p.Align = ph.Align
	return binary.Write(w, binary.LittleEndian, p)
}

func (f *File) writeSectionHeader(w io.Writer, sec *Section) error {
	var shdr elfSectionHeader

	shdr.Name = sec.nameIdx
	shdr.Type = sec.Type
	shdr.Flags = sec.Flags
	shdr.Addr = sec.Addr
	shdr.Offset = sec.offset
	shdr.Size = sec.size
	shdr.Link = sec.Link
	shdr.Info = sec.Info
	shdr.Addralign = sec.Addralign
	shdr.Entsize = sec.Entsize

	return binary.Write(w, binary.LittleEndian, shdr)
}

func (f *File) writeSymbol(w io.Writer, sym *Symbol) error {
	shndx := uint16(SHN_UNDEF)
	if sym.Section != nil {
		shndx = sym.Section.Index
	}

	binary.Write(w, binary.LittleEndian, sym.nameIdx)
	w.Write([]byte{sym.Info})
	w.Write([]byte{sym.Other})
	binary.Write(w, binary.LittleEndian, shndx)
	binary.Write(w, binary.LittleEndian, sym.Value)
	binary.Write(w, binary.LittleEndian, sym.Size)

	return nil
}

func MakeSymbolInfo(binding, typ byte) byte {
	return (binding << 4) | (typ & 0xf)
}

// ELF structures
type elfHeader struct {
	Ident     [EI_NIDENT]byte
	Type      uint16
	Machine   uint16
	Version   uint32
	Entry     uint64
	Phoff     uint64
	Shoff     uint64
	Flags     uint32
	Ehsize    uint16
	Phentsize uint16
	Phnum     uint16
	Shentsize uint16
	Shnum     uint16
	Shstrndx  uint16
}

type elfProgramHeader struct {
	Type   uint32
	Flags  uint32
	Off    uint64
	Vaddr  uint64
	Paddr  uint64
	Filesz uint64
	Memsz  uint64
	Align  uint64
}

type elfSectionHeader struct {
	Name      uint32
	Type      uint32
	Flags     uint64
	Addr      uint64
	Offset    uint64
	Size      uint64
	Link      uint32
	Info      uint32
	Addralign uint64
	Entsize   uint64
}