package compiler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/codegen"
	"github.com/arc-language/arc-lang/linker/elf" // Our native linker
)

// Run is the main entry point for the compiler library.
func (c *Compiler) Run(cfg Config) error {
	cfg.PostProcess()
	c.logger.Info("Compiling %s -> %s", cfg.InputFile, cfg.OutputFile)

	// 1. Compile to IR (Parsing + Semantics + IRGen)
	module, err := c.CompileProject(cfg.InputFile)
	if err != nil {
		// Errors are already logged inside CompileProject via the diagnostic bag
		return fmt.Errorf("compilation failed")
	}

	// 2. Output Handling based on Type
	switch cfg.OutputType {
	case OutputIR:
		return os.WriteFile(cfg.OutputFile, []byte(module.String()), 0644)

	case OutputObject:
		return c.emitObject(module, cfg.OutputFile)

	case OutputExecutable:
		return c.emitExecutable(module, cfg)
	
	default:
		return fmt.Errorf("unknown output type")
	}
}

// emitObject generates the raw ELF .o file using the backend
func (c *Compiler) emitObject(m *ir.Module, path string) error {
	objData, err := codegen.GenerateObject(m)
	if err != nil {
		return err
	}
	return os.WriteFile(path, objData, 0644)
}

// emitExecutable compiles to object in-memory, then invokes our internal Linker
func (c *Compiler) emitExecutable(m *ir.Module, cfg Config) error {
	// Step 1: Generate Machine Code (The .o file in memory)
	c.logger.Debug("Generating machine code...")
	objData, err := codegen.GenerateObject(m)
	if err != nil {
		return fmt.Errorf("codegen failed: %w", err)
	}

	// Step 2: Configure the Native Linker
	c.logger.Debug("Linking...")
	
	linkConf := elf.Config{
		Entry:    "_start",   // Standard ELF entry point
		BaseAddr: 0x400000,   // Standard Executable Base
	}
	
	linker := elf.NewLinker(linkConf)

	// Add our compiled code as the main object
	// We call it "main.o" internally, but it exists only in memory
	if err := linker.AddObject("main.o", objData); err != nil {
		return fmt.Errorf("linker failed to load internal object: %w", err)
	}

	// Step 3: Resolve External Libraries
	// This logic mimics how ld finds libraries in -L paths
	searchPaths := append(cfg.LibraryPaths, "/usr/lib", "/lib64", "/usr/lib/x86_64-linux-gnu")

	for _, lib := range cfg.Libraries {
		found := false
		
		for _, dir := range searchPaths {
			// Priority 1: Shared Library (.so)
			soPath := filepath.Join(dir, "lib"+lib+".so")
			if data, err := os.ReadFile(soPath); err == nil {
				c.logger.Debug("Linking shared library: %s", soPath)
				if err := linker.AddSharedLib(soPath, data); err != nil {
					return fmt.Errorf("failed to link shared lib %s: %w", soPath, err)
				}
				found = true
				break
			}
			
			// Priority 2: Static Library (.a)
			aPath := filepath.Join(dir, "lib"+lib+".a")
			if _, err := os.Stat(aPath); err == nil {
				c.logger.Debug("Linking static library: %s", aPath)
				if err := linker.AddArchive(aPath); err != nil {
					return fmt.Errorf("failed to link archive %s: %w", aPath, err)
				}
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("library -l%s not found in search paths", lib)
		}
	}

	// Step 4: Write Final Binary
	if err := linker.Link(cfg.OutputFile); err != nil {
		return fmt.Errorf("linking failed: %w", err)
	}

	c.logger.Info("Success! Executable created: %s", cfg.OutputFile)
	return nil
}