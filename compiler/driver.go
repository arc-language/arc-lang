package compiler

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	//"strings"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/codegen"
	"github.com/arc-language/arc-lang/linker/elf"
)

// Run is the main entry point for the compiler library.
func (c *Compiler) Run(cfg Config) error {
	cfg.PostProcess()
	c.logger.Info("Compiling %s -> %s", cfg.InputFile, cfg.OutputFile)

	// 1. Compile to IR
	module, err := c.CompileProject(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("compilation failed")
	}

	// 2. Output Handling
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

func (c *Compiler) emitObject(m *ir.Module, path string) error {
	objData, err := codegen.GenerateObject(m)
	if err != nil {
		return err
	}
	return os.WriteFile(path, objData, 0644)
}

func (c *Compiler) emitExecutable(m *ir.Module, cfg Config) error {
	// Step 1: Generate Machine Code
	c.logger.Debug("Generating machine code...")
	objData, err := codegen.GenerateObject(m)
	if err != nil {
		return fmt.Errorf("codegen failed: %w", err)
	}

	// Step 2: Configure Linker
	c.logger.Debug("Linking...")
	linkConf := elf.Config{
		Entry:    "_start",
		BaseAddr: 0x400000,
	}
	linker := elf.NewLinker(linkConf)

	if err := linker.AddObject("main.o", objData); err != nil {
		return fmt.Errorf("linker failed to load internal object: %w", err)
	}

	// Step 3: Resolve External Libraries
	// Standard search paths including typical distros
	searchPaths := append(cfg.LibraryPaths,
		"/usr/lib/x86_64-linux-gnu",
		"/lib/x86_64-linux-gnu",
		"/usr/lib64",
		"/lib64",
		"/usr/lib",
		"/lib",
	)

	// Regex to find library path inside GNU ld script
	// Matches: GROUP ( /path/to/lib ... )
	ldScriptRegex := regexp.MustCompile(`(?:GROUP|INPUT)\s*\(\s*([^\s)]+)`)

	for _, lib := range cfg.Libraries {
		found := false

		for _, dir := range searchPaths {
			// Priority 1: Shared Library (.so)
			soPath := filepath.Join(dir, "lib"+lib+".so")
			if data, err := os.ReadFile(soPath); err == nil {
				c.logger.Debug("Found candidate library: %s", soPath)

				// CHECK: Is this a GNU Linker Script? (e.g. libc.so text file)
				// Magic: "/* GNU ld script"
				if len(data) > 8 && string(data[:8]) == "/* GNU l" {
					c.logger.Debug("Parsing linker script: %s", soPath)
					content := string(data)
					match := ldScriptRegex.FindStringSubmatch(content)
					if len(match) > 1 {
						realPath := match[1]
						// If path is relative, join with current dir, otherwise uses absolute
						if !filepath.IsAbs(realPath) {
							realPath = filepath.Join(dir, realPath)
						}

						// Load the REAL library (e.g. libc.so.6)
						if realData, err := os.ReadFile(realPath); err == nil {
							c.logger.Debug("Linking resolved library: %s", realPath)
							if err := linker.AddSharedLib(realPath, realData); err != nil {
								return fmt.Errorf("failed to link resolved lib %s: %w", realPath, err)
							}
							found = true
							break
						}
					}
				}

				// Standard Binary ELF Load
				if err := linker.AddSharedLib(soPath, data); err == nil {
					c.logger.Debug("Linking shared library: %s", soPath)
					found = true
					break
				}
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