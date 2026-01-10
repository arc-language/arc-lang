package elf

import (
	"bytes"
	stdelf "debug/elf" // Alias standard lib to avoid collision
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// InputObject represents a parsed .o file
type InputObject struct {
	Name     string
	Sections []*InputSection
	Symbols  []*InputSymbol
	File     *stdelf.File 
}

type InputSection struct {
	Name    string
	Type    stdelf.SectionType
	Flags   stdelf.SectionFlag
	Data    []byte
	Relocs  []InputReloc
	
	// Output mapping
	VirtualAddress uint64
	OutputOffset   uint64
}

type InputReloc struct {
	Offset uint64
	Type   uint32
	Addend int64
	Sym    *InputSymbol 
}

type InputSymbol struct {
	Name    string
	Type    stdelf.SymType
	Bind    stdelf.SymBind
	Section *InputSection // Nil if Undefined
	Value   uint64        
	Size    uint64
}

// LoadObject parses an ELF object from a byte slice
func LoadObject(name string, data []byte) (*InputObject, error) {
	f, err := stdelf.NewFile(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to parse elf %s: %w", name, err)
	}

	obj := &InputObject{Name: name, File: f}
	
	// 1. Load Sections (Only PROGBITS and NOBITS)
	sectionsByIndex := make(map[int]*InputSection)
	
	for i, sec := range f.Sections {
		if sec.Type == stdelf.SHT_PROGBITS || sec.Type == stdelf.SHT_NOBITS {
			data, _ := sec.Data()
			isec := &InputSection{
				Name:  sec.Name,
				Type:  sec.Type,
				Flags: sec.Flags,
				Data:  data,
			}
			obj.Sections = append(obj.Sections, isec)
			sectionsByIndex[i] = isec
		}
	}

	// 2. Load Symbols
	syms, err := f.Symbols()
	if err != nil && err != stdelf.ErrNoSymbols {
		return nil, err
	}

	for _, s := range syms {
		isym := &InputSymbol{
			Name:  s.Name,
			Type:  stdelf.SymType(s.Info & 0xf),
			Bind:  stdelf.SymBind(s.Info >> 4),
			Value: s.Value,
			Size:  s.Size,
		}

		if s.Section >= 0 && int(s.Section) < len(f.Sections) {
			if targetSec, ok := sectionsByIndex[int(s.Section)]; ok {
				isym.Section = targetSec
			}
		}
		obj.Symbols = append(obj.Symbols, isym)
	}

	// 3. Load Relocations
	for _, sec := range f.Sections {
		if sec.Type == stdelf.SHT_RELA {
			targetSec := sectionsByIndex[int(sec.Info)]
			if targetSec == nil { continue }

			rdata, _ := sec.Data()
			rreader := bytes.NewReader(rdata)
			numRelocs := uint64(len(rdata)) / 24 // sizeof(Elf64_Rela)

			for i := uint64(0); i < numRelocs; i++ {
				var rOff, rInfo uint64
				var rAddend int64
				binary.Read(rreader, Le, &rOff)
				binary.Read(rreader, Le, &rInfo)
				binary.Read(rreader, Le, &rAddend)
				
				symIdx := int(rInfo >> 32)
				rType := uint32(rInfo & 0xFFFFFFFF)

				if symIdx < len(obj.Symbols) {
					targetSec.Relocs = append(targetSec.Relocs, InputReloc{
						Offset: rOff,
						Type:   rType,
						Addend: rAddend,
						Sym:    obj.Symbols[symIdx],
					})
				}
			}
		}
	}

	return obj, nil
}

// LoadArchive iterates a .a file and returns all contained ELF objects.
func LoadArchive(path string) ([]*InputObject, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	magic := make([]byte, 8)
	if _, err := f.Read(magic); err != nil || string(magic) != "!<arch>\n" {
		return nil, fmt.Errorf("not a valid archive: %s", path)
	}

	var objects []*InputObject

	for {
		header := make([]byte, 60)
		if _, err := io.ReadFull(f, header); err != nil {
			if err == io.EOF { break }
			return nil, err
		}

		name := strings.TrimSpace(string(header[0:16]))
		sizeStr := strings.TrimSpace(string(header[48:58]))
		size, _ := strconv.ParseInt(sizeStr, 10, 64)

		content := make([]byte, size)
		if _, err := io.ReadFull(f, content); err != nil {
			return nil, err
		}

		if size%2 != 0 { f.Seek(1, io.SeekCurrent) }

		if name == "/" || name == "//" || name == "/SYM64/" { continue }

		// Heuristic: Check for ELF magic inside archive member
		if len(content) > 4 && string(content[1:4]) == "ELF" {
			obj, err := LoadObject(name, content)
			if err == nil {
				objects = append(objects, obj)
			}
		}
	}
	return objects, nil
}

// LoadSharedObject parses a .so file to find exported symbols
func LoadSharedObject(name string, data []byte) (*SharedObject, error) {
	f, err := stdelf.NewFile(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to parse so %s: %w", name, err)
	}

	so := &SharedObject{Name: name}

	// We look at SHT_DYNSYM (Dynamic Symbols) not SHT_SYMTAB
	syms, err := f.DynamicSymbols()
	if err != nil {
		// Some libs might strip dynsyms or handle them differently, 
		// but standard .so files have them.
		return nil, fmt.Errorf("no dynamic symbols in %s", name)
	}

	for _, s := range syms {
		// We only care about Defined Global Functions or Objects
		if s.Section != stdelf.SHN_UNDEF && (stdelf.ST_BIND(s.Info) == stdelf.STB_GLOBAL || stdelf.ST_BIND(s.Info) == stdelf.STB_WEAK) {
			so.Symbols = append(so.Symbols, s.Name)
		}
	}

	return so, nil
}