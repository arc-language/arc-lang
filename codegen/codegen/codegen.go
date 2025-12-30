package codegen

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/codegen/arch/amd64"
	"github.com/arc-language/arc-lang/codegen/elf"
)

// GenerateObject creates a relocatable ELF object file (.o)
func GenerateObject(m *ir.Module) ([]byte, error) {
	// Compile to machine code
	artifact, err := amd64.Compile(m)
	if err != nil {
		return nil, err
	}

	// Create ELF file
	f := elf.NewFile()
	f.Type = elf.ET_REL

	// Create .text section
	textSec := f.AddSection(".text", elf.SHT_PROGBITS, elf.SHF_ALLOC|elf.SHF_EXECINSTR, artifact.TextBuffer)
	textSec.Addralign = 16

	// Create .data section
	dataSec := f.AddSection(".data", elf.SHT_PROGBITS, elf.SHF_ALLOC|elf.SHF_WRITE, artifact.DataBuffer)
	dataSec.Addralign = 8

	// Add symbols
	for _, symDef := range artifact.Symbols {
		var sec *elf.Section
		if symDef.IsGlobal { // In our IR, globals are variables in .data
			sec = dataSec
		} else {
			sec = textSec
		}

		info := elf.MakeSymbolInfo(elf.STB_GLOBAL, elf.STT_FUNC)
		if symDef.IsGlobal {
			info = elf.MakeSymbolInfo(elf.STB_GLOBAL, elf.STT_OBJECT)
		}

		f.AddSymbol(symDef.Name, info, sec, symDef.Offset, symDef.Size)
	}

	// Write to buffer
	buf := new(bytes.Buffer)
	if err := f.WriteTo(buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// GenerateExecutable creates a static ELF executable directly
// It acts as a static linker, resolving relocations internally
func GenerateExecutable(m *ir.Module) ([]byte, error) {
	// 1. Compile to machine code
	artifact, err := amd64.Compile(m)
	if err != nil {
		return nil, err
	}

	// 2. Setup Memory Layout (Static Linking)
	// Base address for standard Linux executable
	const baseAddr = 0x400000
	const pageSize = 0x1000

	// Entry stub: _start calls main and then exit
	// We inject this at the beginning of the text section
	// _start:
	//   call main
	//   mov rdi, rax  (exit code)
	//   mov rax, 60   (sys_exit)
	//   syscall
	entryStub := []byte{
		0xE8, 0x00, 0x00, 0x00, 0x00, // call main (placeholder offset)
		0x48, 0x89, 0xC7,             // mov rdi, rax
		0xB8, 0x3C, 0x00, 0x00, 0x00, // mov eax, 60
		0x0F, 0x05,                   // syscall
	}
	
	// Prepend stub to text buffer
	finalText := append(entryStub, artifact.TextBuffer...)
	stubSize := uint64(len(entryStub))

	// Find 'main' symbol to resolve the stub call
	var mainOffset uint64
	foundMain := false
	for _, sym := range artifact.Symbols {
		if sym.Name == "main" {
			mainOffset = sym.Offset
			foundMain = true
			break
		}
	}

	if !foundMain {
		return nil, fmt.Errorf("entry point 'main' not found")
	}

	// Patch the call instruction in the stub
	// call rel32: target = mainOffset + stubSize (since main is after stub)
	// rel = target - (PC of next instr) = (stubSize + mainOffset) - 5
	rel := int32((stubSize + mainOffset) - 5)
	binary.LittleEndian.PutUint32(finalText[1:], uint32(rel))

	// Adjust symbol offsets because we shifted text by stubSize
	for i := range artifact.Symbols {
		if !artifact.Symbols[i].IsGlobal { // Functions are in text
			artifact.Symbols[i].Offset += stubSize
		}
	}
	// Adjust relocations
	for i := range artifact.Relocations {
		artifact.Relocations[i].Offset += stubSize
	}

	// Calculate Virtual Addresses
	// ELF Header (64) + 2 * Phdr (56) = 176 bytes
	// We'll align the text section to 0x400000 + 0x1000 to be safe and clean
	
	// Text Segment: R-X
	textVAddr := uint64(baseAddr + pageSize) 
	textSize := uint64(len(finalText))
	
	// Data Segment: RW- (aligned to page boundary after text)
	dataVAddr := textVAddr + textSize
	if dataVAddr%pageSize != 0 {
		dataVAddr += pageSize - (dataVAddr % pageSize)
	}
	dataSize := uint64(len(artifact.DataBuffer))

	// 3. Resolve Relocations
	// We need to patch the binary code based on where symbols ended up
	for _, reloc := range artifact.Relocations {
		// Target Symbol Address
		var symVAddr uint64
		found := false
		
		for _, sym := range artifact.Symbols {
			if sym.Name == reloc.SymbolName {
				if sym.IsGlobal {
					symVAddr = dataVAddr + sym.Offset
				} else {
					symVAddr = textVAddr + sym.Offset
				}
				found = true
				break
			}
		}
		
		if !found {
			// It might be a builtin like __exception_state if not explicitly in symbols
			// But compiler usually emits symbols. If missing, fail.
			return nil, fmt.Errorf("undefined symbol: %s", reloc.SymbolName)
		}

		// Apply relocation
		// PC = VAddr of the instruction + Offset within text
		pc := textVAddr + reloc.Offset
		
		switch reloc.Type {
		case amd64.R_X86_64_PC32, amd64.R_X86_64_PLT32:
			// Value = Symbol - PC + Addend
			val := int32(int64(symVAddr) - int64(pc) + reloc.Addend)
			
			// Write back 32-bit value
			// The relocation offset points to the end of the opcode, where the imm32 starts
			// Note: amd64.Compile emits placeholder bytes at these offsets
			binary.LittleEndian.PutUint32(finalText[reloc.Offset:], uint32(val))
			
		default:
			return nil, fmt.Errorf("unsupported relocation type for static link: %d", reloc.Type)
		}
	}

	// 4. Create ELF File
	f := elf.NewFile()
	f.Type = elf.ET_EXEC
	f.Entry = textVAddr // Points to _start stub

	// Create Program Headers (Segments)
	// 1. Load Text (R E)
	// We map the ELF header + Text into the first segment usually, but for simplicity
	// let's just map the sections.
	// Actually, for a valid executable, the headers usually need to be loaded.
	
	// Simplest static binary layout:
	// File Off 0: Headers
	// File Off 4096: .text
	// File Off ... : .data
	
	// We need to calculate file offsets first.
	// Headers Size: 64 + 2*56 = 176.
	// Let's force .text to file offset 4096.
	
	f.AddProgramHeader(elf.PT_LOAD, elf.PF_R|elf.PF_X, 0x1000, textVAddr, textSize, textSize, pageSize)
	
	if dataSize > 0 {
		// Calculate file offset for data
		// Since we write sequentially, we need to know where text ends in file.
		// Text starts at 0x1000. Ends at 0x1000 + len.
		// Data starts at dataVAddr - baseAddr in file? No.
		
		// In file:
		// [Headers ... padding ... ] [Text] [padding] [Data]
		
		dataFileOff := uint64(0x1000) + textSize
		// Align file offset to page size? Not strictly necessary for static binary if VAddr is aligned,
		// but good practice.
		if dataFileOff%pageSize != 0 {
			dataFileOff += pageSize - (dataFileOff % pageSize)
		}
		
		f.AddProgramHeader(elf.PT_LOAD, elf.PF_R|elf.PF_W, dataFileOff, dataVAddr, dataSize, dataSize, pageSize)
		
		// Create Sections (mapped to these file offsets)
		t := f.AddSection(".text", elf.SHT_PROGBITS, elf.SHF_ALLOC|elf.SHF_EXECINSTR, finalText)
		t.Addr = textVAddr
		// We manually set offset to match Phdr
		// Note: writer.go calculates offsets linearly. We might need padding.
		// Since writer.go logic is simple linear append, we must rely on its calculation
		// or ensure the content we pass includes padding.
		
		// To ensure writer.go places .text at 0x1000:
		// We insert a dummy section or rely on alignment?
		// writer.go: "currentOffset += sec.Addralign - ..."
		// If we set t.offset manually, writer respects it.
		t.Index = 1 // fixup internal index if needed
		// writer.go uses t.offset = currentOffset if 0.
		
		// HACK: To make writer.go produce the exact layout we want without modifying it too much,
		// we can create a "padding" section before text if needed, OR we just let writer calculate,
		// and we update Phdr offsets to match what writer WILL produce.
		
		// Let's predict writer output:
		// Header (64) + Phdr (112) = 176.
		// StrTabs come first. They are small.
		// Then .text.
		
		// Actually, to create a proper executable with fixed addresses using a simple writer,
		// the simplest way is to NOT use sections for execution, just Segments.
		// But `writer.go` is section-based.
		
		// Let's modify the approach:
		// We will set the `Addr` on sections.
		// We will add Phdrs that cover these sections.
		// We rely on `writer.go` to pack them.
		// For an executable, VAddr alignment matters.
		
		t.Addr = textVAddr
		t.Addralign = pageSize // Force alignment in file to 4096
		
		d := f.AddSection(".data", elf.SHT_PROGBITS, elf.SHF_ALLOC|elf.SHF_WRITE, artifact.DataBuffer)
		d.Addr = dataVAddr
		d.Addralign = pageSize
		
		// Re-make Phdrs based on section properties
		// Since we don't know exact file offsets yet, we can't fully fill Phdr.Off.
		// However, `writer.go` writes Phdrs BEFORE sections. It needs Phdr.Off.
		// This circular dependency (Phdr needs Section Offset, Section Offset calculated during write)
		// suggests we need a multi-pass or pre-calculation.
		
		// Pre-calculation logic:
		// Headers: ~200 bytes.
		// .shstrtab, .strtab, .symtab: ~few hundred bytes.
		// We can put all metadata at start, align .text to 0x1000.
		
		// Let's assume writer will align .text to 0x1000 because we set Addralign=4096.
		// Text Offset = 0x1000.
		// Data Offset = Text Offset + TextSize aligned to 4096.
		
		// Update Phdrs with these assumptions
		f.ProgramHeaders[0].Off = 0x1000
		f.ProgramHeaders[1].Off = dataFileOff
	} else {
		// Only text
		f.AddProgramHeader(elf.PT_LOAD, elf.PF_R|elf.PF_X, 0x1000, textVAddr, textSize, textSize, pageSize)
		t := f.AddSection(".text", elf.SHT_PROGBITS, elf.SHF_ALLOC|elf.SHF_EXECINSTR, finalText)
		t.Addr = textVAddr
		t.Addralign = pageSize
	}

	buf := new(bytes.Buffer)
	if err := f.WriteTo(buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}