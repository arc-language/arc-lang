package codegen

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/codegen/cpu/amd64"
	"github.com/arc-language/arc-lang/codegen/format/elf"
	"github.com/arc-language/arc-lang/codegen/gpu/nvidia"
	"github.com/arc-language/arc-lang/codegen/gpu/amd"
	"github.com/arc-language/arc-lang/codegen/tpu" // New Import
	"github.com/arc-language/arc-lang/context"
)

// Generate handles the backend code generation for the module.
// It automatically detects GPU/TPU kernels and invokes the appropriate backend.
func Generate(module *ir.Module) error {
	logger := context.NewLogger("[CodeGen]")

	// 1. Scan for Hardware Specific Functions
	hasGPUFunctions := false
	hasROCmFunctions := false
	hasTPUFunctions := false

	for _, fn := range module.Functions {
		if fn.CallConv == ir.CC_PTX {
			hasGPUFunctions = true
		} else if fn.CallConv == ir.CC_ROCM {
			hasROCmFunctions = true
		} else if fn.CallConv == ir.CC_TPU {
			hasTPUFunctions = true
		}
	}

	// 2. Generate NVIDIA PTX if needed
	if hasGPUFunctions {
		logger.Info("NVIDIA GPU functions detected (CC_PTX). Generating PTX Assembly...")

		ptxCode, err := nvidia.Generate(module)
		if err != nil {
			return fmt.Errorf("NVIDIA codegen failed: %w", err)
		}

		fmt.Println("\n// ==========================================")
		fmt.Println("// GENERATED NVIDIA PTX ASSEMBLY")
		fmt.Println("// ==========================================")
		fmt.Println(ptxCode)
		fmt.Println("// ==========================================\n")
	}

	// 3. Generate AMD ROCm (GCN) if needed
	if hasROCmFunctions {
		logger.Info("AMD GPU functions detected (CC_ROCM). Generating GCN Assembly...")

		gcnCode, err := amd.Generate(module)
		if err != nil {
			return fmt.Errorf("AMD codegen failed: %w", err)
		}

		fmt.Println("\n// ==========================================")
		fmt.Println("// GENERATED AMD GCN ASSEMBLY")
		fmt.Println("// ==========================================")
		fmt.Println(gcnCode)
		fmt.Println("// ==========================================\n")
	}

	// 4. Generate Google TPU HLO if needed
	if hasTPUFunctions {
		logger.Info("TPU functions detected (CC_TPU). Generating Google HLO...")

		hloCode, err := tpu.Generate(module)
		if err != nil {
			return fmt.Errorf("TPU codegen failed: %w", err)
		}

		fmt.Println("\n// ==========================================")
		fmt.Println("// GENERATED TPU HLO IR")
		fmt.Println("// ==========================================")
		fmt.Println(hloCode)
		fmt.Println("// ==========================================\n")
	}

	// 5. Always proceed with standard CPU Code Generation
	// The CPU backend (amd64) handles the host code.
	// Hardware kernel functions (PTX/ROCm/TPU) act as externs/stubs in the host ELF,
	// typically managed via the runtime (not shown here) loading the generated assembly above.
	logger.Info("Generating x86-64 CPU code...")

	objBytes, err := GenerateObject(module)
	if err != nil {
		return fmt.Errorf("CPU codegen failed: %w", err)
	}
	
	logger.Info("Generated %d bytes of x86-64 object code.", len(objBytes))

	return nil
}

// GenerateObject creates a relocatable ELF object file (.o)
// This file can be linked with gcc or ld (e.g., `gcc output.o -o main`)
func GenerateObject(m *ir.Module) ([]byte, error) {
	// 1. Compile IR to Machine Code
	artifact, err := amd64.Compile(m)
	if err != nil {
		return nil, err
	}

	// 2. Wrap in ELF container
	f := elf.NewFile()
	f.Type = elf.ET_REL

	// Create .text section
	textSec := f.AddSection(".text", elf.SHT_PROGBITS, elf.SHF_ALLOC|elf.SHF_EXECINSTR, artifact.Text)
	textSec.Addralign = 16

	// Create .data section
	dataSec := f.AddSection(".data", elf.SHT_PROGBITS, elf.SHF_ALLOC|elf.SHF_WRITE, artifact.Data)
	dataSec.Addralign = 8

	// 3. Add Symbols
	// We need a map to look up ELF symbols by name for relocations later
	symMap := make(map[string]*elf.Symbol)

	for _, symDef := range artifact.Symbols {
		var sec *elf.Section
		if symDef.IsFunc {
			sec = textSec
		} else {
			sec = dataSec
		}

		info := elf.MakeSymbolInfo(elf.STB_GLOBAL, elf.STT_OBJECT)
		if symDef.IsFunc {
			info = elf.MakeSymbolInfo(elf.STB_GLOBAL, elf.STT_FUNC)
		}

		es := f.AddSymbol(symDef.Name, info, sec, symDef.Offset, symDef.Size)
		symMap[symDef.Name] = es
	}

	// 4. Process Relocations
	// These tell the linker (ld) where to patch addresses in .text
	for _, reloc := range artifact.Relocs {
		sym, ok := symMap[reloc.Symbol]
		if !ok {
			// Symbol is external (e.g., printf, malloc), not defined in this module.
			// We must create an UNDEFINED global symbol for the linker to resolve.
			sym = f.AddSymbol(reloc.Symbol, elf.MakeSymbolInfo(elf.STB_GLOBAL, elf.STT_NOTYPE), nil, 0, 0)
			symMap[reloc.Symbol] = sym
		}

		// Currently, we only support relocations in .text
		// Cast RelocationType (int) to uint32 for the ELF writer
		f.AddRelocation(textSec, sym, uint64(reloc.Offset), uint32(reloc.Type), reloc.Addend)
	}

	// 5. Write to buffer
	buf := new(bytes.Buffer)
	if err := f.WriteTo(buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// GenerateExecutable creates a static ELF executable directly.
// It acts as a static linker, resolving internal relocations and adding an entry point.
func GenerateExecutable(m *ir.Module) ([]byte, error) {
	// 1. Compile to machine code
	artifact, err := amd64.Compile(m)
	if err != nil {
		return nil, err
	}

	// 2. Setup Memory Layout
	const (
		BaseAddr = 0x400000
		PageSize = 0x1000
	)

	// Entry stub: _start calls main and then exit(0)
	// We prepend this to the text buffer.
	// _start:
	//   call main      (E8 xx xx xx xx)
	//   xor rdi, rdi   (48 31 FF) - exit code 0
	//   mov rax, 60    (48 C7 C0 3C 00 00 00) - sys_exit
	//   syscall        (0F 05)
	entryStub := []byte{
		0xE8, 0x00, 0x00, 0x00, 0x00, // call main (placeholder)
		0x48, 0x31, 0xFF, // xor rdi, rdi
		0x48, 0xC7, 0xC0, 0x3C, 0x00, 0x00, 0x00, // mov rax, 60
		0x0F, 0x05, // syscall
	}

	stubSize := len(entryStub)
	finalText := append(entryStub, artifact.Text...)

	// Layout addresses
	textVAddr := uint64(BaseAddr + PageSize) // Text starts at 0x401000
	dataVAddr := textVAddr + uint64(len(finalText))

	// Align Data to page boundary
	if dataVAddr%PageSize != 0 {
		dataVAddr += PageSize - (dataVAddr % PageSize)
	}

	// 3. Resolve "main" for the stub
	var mainOffset uint64
	foundMain := false
	for _, sym := range artifact.Symbols {
		if sym.Name == "main" {
			mainOffset = sym.Offset // Offset within artifact.Text
			foundMain = true
			break
		}
	}

	if !foundMain {
		return nil, fmt.Errorf("entry point 'main' not found")
	}

	// Patch stub call to main
	// Target = (stubSize + mainOffset) relative to (PC + 5)
	// PC of next instruction is index 5
	rel := int32((stubSize + int(mainOffset)) - 5)
	binary.LittleEndian.PutUint32(finalText[1:], uint32(rel))

	// 4. Resolve Internal Relocations (Static Linking)
	for _, reloc := range artifact.Relocs {
		// Find symbol address
		var symAddr uint64
		found := false

		for _, sym := range artifact.Symbols {
			if sym.Name == reloc.Symbol {
				if sym.IsFunc {
					// Function in text section (offset shifted by stub)
					symAddr = textVAddr + uint64(stubSize) + sym.Offset
				} else {
					// Global in data section
					symAddr = dataVAddr + sym.Offset
				}
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("undefined symbol '%s' (static linking does not support external libs yet)", reloc.Symbol)
		}

		// Apply Patch
		// PC = VAddr of the instruction + Offset within text (shifted by stub)
		pc := textVAddr + uint64(stubSize) + uint64(reloc.Offset)

		// Calculate Value: Symbol - PC + Addend
		val := int32(int64(symAddr) - int64(pc) + reloc.Addend)

		// Write to finalText (offset shifted by stub)
		patchOffset := stubSize + int(reloc.Offset)
		binary.LittleEndian.PutUint32(finalText[patchOffset:], uint32(val))
	}

	// 5. Create ELF Container
	f := elf.NewFile()
	f.Type = elf.ET_EXEC
	f.Entry = textVAddr // Points to _start (beginning of stub)

	// Add Program Headers (Segments)
	// Segment 1: Text (Read+Exec)
	textSize := uint64(len(finalText))
	f.AddProgramHeader(elf.PT_LOAD, elf.PF_R|elf.PF_X, 0x1000, textVAddr, textSize, textSize, PageSize)

	// Segment 2: Data (Read+Write)
	dataSize := uint64(len(artifact.Data))
	if dataSize > 0 {
		// We calculate the file offset linearly.
		// Text is at 0x1000. Data is appended.
		// NOTE: A robust linker would align file offsets to pages too (0x2000),
		// but for small static binaries, packing often works if VAddr aligns.
		// Let's align file offset for safety.
		dataFileOff := uint64(0x1000) + textSize
		if dataFileOff%PageSize != 0 {
			dataFileOff += PageSize - (dataFileOff % PageSize)
		}

		f.AddProgramHeader(elf.PT_LOAD, elf.PF_R|elf.PF_W, dataFileOff, dataVAddr, dataSize, dataSize, PageSize)
	}

	// Add Sections
	t := f.AddSection(".text", elf.SHT_PROGBITS, elf.SHF_ALLOC|elf.SHF_EXECINSTR, finalText)
	t.Addr = textVAddr
	t.Addralign = PageSize

	d := f.AddSection(".data", elf.SHT_PROGBITS, elf.SHF_ALLOC|elf.SHF_WRITE, artifact.Data)
	d.Addr = dataVAddr
	d.Addralign = PageSize

	buf := new(bytes.Buffer)
	if err := f.WriteTo(buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}