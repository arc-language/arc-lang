package compiler

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/codegen/codegen"
	"github.com/arc-language/arc-lang/linker/elf"
	"github.com/arc-language/upkg" // Import upkg to find library paths
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

	// Get upkg configuration to find where packages are installed
	upkgConfig := upkg.DefaultConfig()
	installPath := upkgConfig.InstallPath

	// Build search paths: User Flags -> Upkg Paths -> System Paths
	searchPaths := cfg.LibraryPaths
	
	// Add upkg paths (check various common layouts: lib, usr/lib, lib64)
	searchPaths = append(searchPaths,
		filepath.Join(installPath, "lib"),
		filepath.Join(installPath, "usr", "lib"),
		filepath.Join(installPath, "usr", "lib64"),
		filepath.Join(installPath, "lib64"),
	)

	// Add standard system paths (Linux fallback)
	searchPaths = append(searchPaths,
		"/usr/lib/x86_64-linux-gnu",
		"/lib/x86_64-linux-gnu",
		"/usr/lib64",
		"/lib64",
		"/usr/lib",
		"/lib",
	)

	ldScriptRegex := regexp.MustCompile(`(?:GROUP|INPUT)\s*\(\s*([^\s)]+)`)

	for _, lib := range cfg.Libraries {
		found := false

		for _, dir := range searchPaths {
			// Priority 1: Shared Library (.so)
			soPath := filepath.Join(dir, "lib"+lib+".so")
			data, err := os.ReadFile(soPath)
			if err != nil {
				continue // Try next path
			}
			
			c.logger.Debug("Found candidate library: %s", soPath)

			// CHECK: Is this a GNU Linker Script?
			if len(data) > 8 && string(data[:8]) == "/* GNU l" {
				c.logger.Debug("Parsing linker script: %s", soPath)
				content := string(data)
				match := ldScriptRegex.FindStringSubmatch(content)
				if len(match) > 1 {
					realPath := match[1]
					if !filepath.IsAbs(realPath) {
						realPath = filepath.Join(dir, realPath)
					}

					// Load the REAL library
					if realData, err := os.ReadFile(realPath); err == nil {
						c.logger.Debug("Linking resolved library: %s", realPath)
						if err := linker.AddSharedLib(realPath, realData); err != nil {
							c.logger.Debug("Failed to link resolved lib %s: %v", realPath, err)
							continue
						}
						found = true
						break
					} else {
						c.logger.Debug("Failed to read real library %s: %v", realPath, err)
					}
				}
				// Linker script parsing failed or couldn't find real lib, try next search path
				continue
			}

			// Standard Binary ELF Load
			if err := linker.AddSharedLib(soPath, data); err == nil {
				c.logger.Debug("Linking shared library: %s", soPath)
				found = true
				break
			} else {
				c.logger.Debug("Failed to load %s as ELF: %v", soPath, err)
			}

			// Priority 2: Static Library (.a)
			aPath := filepath.Join(dir, "lib"+lib+".a")
			if _, err := os.Stat(aPath); err == nil {
				c.logger.Debug("Linking static library: %s", aPath)
				if err := linker.AddArchive(aPath); err != nil {
					c.logger.Debug("Failed to link archive %s: %v", aPath, err)
					continue
				}
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("library -l%s not found in search paths (checked upkg: %s)", lib, installPath)
		}
	}

	// Step 4: Write Final Binary
	if err := linker.Link(cfg.OutputFile); err != nil {
		return fmt.Errorf("linking failed: %w", err)
	}

	c.logger.Info("Success! Executable created: %s", cfg.OutputFile)
	return nil
}