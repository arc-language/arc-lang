package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/arc-language/arc-lang/backend/backend"
	backendelf "github.com/arc-language/arc-lang/backend/format/elf"
	"github.com/arc-language/arc-lang/codegen"
	"github.com/arc-language/arc-lang/frontend"
	"github.com/arc-language/arc-lang/lower"
	"github.com/arc-language/arc-lang/syntax"
)

// multiFlag accumulates repeated flags, e.g. -L /usr/lib -L /lib
type multiFlag []string

func (f *multiFlag) String() string  { return strings.Join(*f, ", ") }
func (f *multiFlag) Set(v string) error {
	*f = append(*f, v)
	return nil
}

// emitMode controls what the compiler produces
type emitMode int

const (
	emitIR  emitMode = iota // -emit ir  → IR text
	emitObj                 // -emit obj → relocatable .o
	emitExe                 // -emit exe → static ELF executable
	emitBin                 // -emit bin → dynamic ELF executable
)

func main() {
	sourceFile := flag.String("src", "main.ax", "Source file to compile")
	outputFile := flag.String("o", "", "Output file (default derived from source name)")
	emitStr    := flag.String("emit", "ir", "Output type: ir | obj | exe | bin")
	debugAST   := flag.Bool("debug-ast", false, "Print the lowered AST before codegen")
	entryPoint := flag.String("entry", "_start", "Entry point symbol (bin mode only)")

	var libPaths multiFlag
	var linkLibs multiFlag
	flag.Var(&libPaths, "L", "Library search path (repeatable)")
	flag.Var(&linkLibs, "l", "Link against shared library, e.g. -l c (repeatable)")

	flag.Parse()

	// Resolve emit mode
	var mode emitMode
	switch strings.ToLower(*emitStr) {
	case "ir":
		mode = emitIR
	case "obj", "object":
		mode = emitObj
	case "exe", "exec", "static":
		mode = emitExe
	case "bin", "binary", "dynamic":
		mode = emitBin
	default:
		fatalf("unknown emit mode '%s' (want ir|obj|exe|bin)", *emitStr)
	}

	// Default output path
	ext        := filepath.Ext(*sourceFile)
	moduleName := strings.TrimSuffix(filepath.Base(*sourceFile), ext)

	if *outputFile == "" {
		switch mode {
		case emitIR:
			// stdout, leave empty
		case emitObj:
			*outputFile = moduleName + ".o"
		case emitExe, emitBin:
			*outputFile = moduleName
		}
	}

	// Read source
	code, err := os.ReadFile(*sourceFile)
	if err != nil {
		fatalf("cannot read '%s': %v", *sourceFile, err)
	}

	// Phase 1: Parse
	parseResult := syntax.Parse(string(code))
	if len(parseResult.Errors) > 0 {
		fmt.Fprintln(os.Stderr, "syntax errors:")
		for _, e := range parseResult.Errors {
			fmt.Fprintln(os.Stderr, " ", e)
		}
		os.Exit(1)
	}

	// Phase 2: Semantic analysis
	astFile  := frontend.Translate(parseResult.Root)
	analyzer := frontend.NewAnalyzer()
	if err := analyzer.Analyze(astFile); err != nil {
		fatalf("semantic error: %v", err)
	}

	// Phase 3: Lowering
	lower.NewLowerer(astFile).Apply()

	if *debugAST {
		fmt.Fprintln(os.Stderr, "── lowered AST ──────────────────────────────")
		fmt.Fprintf(os.Stderr, "%+v\n", astFile)
		fmt.Fprintln(os.Stderr, "─────────────────────────────────────────────")
	}

	// Phase 4: IR generation
	gen      := codegen.New(moduleName)
	irModule, err := gen.Generate(astFile)
	if err != nil {
		fatalf("codegen error: %v", err)
	}

	// Phase 5: Backend
	switch mode {

	case emitIR:
		ir := irModule.String()
		if *outputFile != "" {
			if err := os.WriteFile(*outputFile, []byte(ir), 0o644); err != nil {
				fatalf("cannot write '%s': %v", *outputFile, err)
			}
			log.Printf("compiled '%s'  →  %s  (IR)", *sourceFile, *outputFile)
		} else {
			fmt.Print(ir)
		}

	case emitObj:
		if err := backend.Generate(irModule); err != nil {
			fatalf("backend error: %v", err)
		}
		objBytes, err := backend.GenerateObject(irModule)
		if err != nil {
			fatalf("object generation failed: %v", err)
		}
		if err := writeExe(*outputFile, objBytes, 0o644); err != nil {
			fatalf("cannot write '%s': %v", *outputFile, err)
		}
		log.Printf("compiled '%s'  →  %s  (%d bytes)", *sourceFile, *outputFile, len(objBytes))

	case emitExe:
		if err := backend.Generate(irModule); err != nil {
			fatalf("backend error: %v", err)
		}
		exeBytes, err := backend.GenerateExecutable(irModule)
		if err != nil {
			fatalf("executable generation failed: %v", err)
		}
		if err := writeExe(*outputFile, exeBytes, 0o755); err != nil {
			fatalf("cannot write '%s': %v", *outputFile, err)
		}
		log.Printf("compiled '%s'  →  %s  (%d bytes, static)", *sourceFile, *outputFile, len(exeBytes))

	case emitBin:
		if err := backend.Generate(irModule); err != nil {
			fatalf("backend error: %v", err)
		}
		objBytes, err := backend.GenerateObject(irModule)
		if err != nil {
			fatalf("object generation failed: %v", err)
		}

		linker := backendelf.NewLinker(backendelf.Config{
			Entry:    *entryPoint,
			BaseAddr: 0x400000,
		})
		if err := linker.AddObject(moduleName+".o", objBytes); err != nil {
			fatalf("linker: failed to add object: %v", err)
		}
		if err := resolveSharedLibs(linker, linkLibs, libPaths); err != nil {
			fatalf("linker: %v", err)
		}
		if err := linker.Link(*outputFile); err != nil {
			fatalf("link failed: %v", err)
		}
		log.Printf("compiled '%s'  →  %s  (dynamic)", *sourceFile, *outputFile)
	}
}

// resolveSharedLibs finds and loads each -l library into the linker.
func resolveSharedLibs(linker *backendelf.Linker, libs, paths multiFlag) error {
	searchDirs := append([]string(paths), systemLibDirs()...)

	for _, lib := range libs {
		candidates := libCandidates(lib)
		found := false

		for _, dir := range searchDirs {
			for _, cand := range candidates {
				full := filepath.Join(dir, cand)
				data, err := os.ReadFile(full)
				if err != nil {
					continue
				}
				if err := linker.AddSharedLib(full, data); err != nil {
					return fmt.Errorf("loading '%s': %w", full, err)
				}
				log.Printf("linked shared library  %s", full)
				found = true
				break
			}
			if found {
				break
			}
		}

		if !found {
			return fmt.Errorf("shared library 'lib%s' not found in search paths", lib)
		}
	}
	return nil
}

func libCandidates(name string) []string {
	if strings.HasPrefix(name, "lib") &&
		(strings.HasSuffix(name, ".so") || strings.Contains(name, ".so.")) {
		return []string{name}
	}
	return []string{
		"lib" + name + ".so",
		"lib" + name + ".so.6",
		"lib" + name + ".so.2",
	}
}

func systemLibDirs() []string {
	return []string{
		"/lib/x86_64-linux-gnu",
		"/usr/lib/x86_64-linux-gnu",
		"/lib64",
		"/usr/lib64",
		"/lib",
		"/usr/lib",
	}
}

func writeExe(path string, data []byte, perm os.FileMode) error {
	if err := os.WriteFile(path, data, perm); err != nil {
		return err
	}
	return os.Chmod(path, perm)
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}