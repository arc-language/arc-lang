package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/arc-language/arc-lang/codegen"
	"github.com/arc-language/arc-lang/frontend"
	"github.com/arc-language/arc-lang/lower"
	"github.com/arc-language/arc-lang/syntax"
)

func main() {
	// 1. CLI Arguments
	sourceFile := flag.String("src", "main.ax", "Source file to compile")
	outputFile := flag.String("o", "", "Output IR file (default: stdout)")
	debugAST := flag.Bool("debug-ast", false, "Print the AST before generation")
	flag.Parse()

	// 2. Read Source Code
	code, err := os.ReadFile(*sourceFile)
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", *sourceFile, err)
		os.Exit(1)
	}

	fmt.Printf("Compiling %s...\n", *sourceFile)

	// =========================================================================
	// Phase 1: Syntax Analysis (Lexing & Parsing)
	// =========================================================================
	// syntax.Parse handles the "Auto-Semicolon Insertion" and ANTLR logic
	parseResult := syntax.Parse(string(code))

	if len(parseResult.Errors) > 0 {
		fmt.Println("Syntax Errors:")
		for _, e := range parseResult.Errors {
			fmt.Println(e)
		}
		os.Exit(1)
	}

	// =========================================================================
	// Phase 2: Frontend (Translation & Semantics)
	// =========================================================================
	// 2a. Translate CST (ANTLR tree) -> AST (Our clean structs)
	astFile := frontend.Translate(parseResult.Root)
	
	// 2b. Semantic Analysis (Symbol Table & Type Checking)
	analyzer := frontend.NewAnalyzer()
	if err := analyzer.Analyze(astFile); err != nil {
		fmt.Printf("Semantic Error: %v\n", err)
		os.Exit(1)
	}

	// =========================================================================
	// Phase 3: Lowering (Arc Magic)
	// =========================================================================
	// Transforms high-level features into explicit logic:
	// - Async -> State Machines / Thread Spawning
	// - Defer -> Explicit calls at exit points
	// - Var   -> Reference counting injection
	low := lower.NewLowerer(astFile)
	low.Apply()

	if *debugAST {
		// You would implement a pretty printer for the AST here
		fmt.Println("--- Lowered AST (Debug) ---")
		fmt.Printf("%+v\n", astFile)
		fmt.Println("---------------------------")
	}

	// =========================================================================
	// Phase 4: Code Generation (AST -> IR)
	// =========================================================================
	// Initialize the generator with the module name (filename without ext)
	moduleName := filepath.Base(*sourceFile)
	moduleName = moduleName[:len(moduleName)-len(filepath.Ext(moduleName))]
	
	gen := codegen.New(moduleName)
	irModule, err := gen.Generate(astFile)
	if err != nil {
		fmt.Printf("Codegen Error: %v\n", err)
		os.Exit(1)
	}

	// =========================================================================
	// Phase 5: Output
	// =========================================================================
	irOutput := irModule.String()

	if *outputFile != "" {
		if err := os.WriteFile(*outputFile, []byte(irOutput), 0644); err != nil {
			fmt.Printf("Error writing output: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully compiled to %s\n", *outputFile)
	} else {
		// Default: Print IR to stdout
		fmt.Println("--- Generated IR ---")
		fmt.Println(irOutput)
	}
}