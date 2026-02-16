package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/arc-language/arc-lang/codegen"
	"github.com/arc-language/arc-lang/frontend"
	"github.com/arc-language/arc-lang/lower"
	"github.com/arc-language/arc-lang/syntax"
)

func main() {
	// ── CLI flags ─────────────────────────────────────────────────────────────
	sourceFile := flag.String("src", "main.ax", "Source file to compile")
	outputFile := flag.String("o", "", "Output IR file (default: stdout)")
	debugAST   := flag.Bool("debug-ast", false, "Print the lowered AST before codegen")
	flag.Parse()

	// ── Read source ───────────────────────────────────────────────────────────
	code, err := os.ReadFile(*sourceFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: cannot read '%s': %v\n", *sourceFile, err)
		os.Exit(1)
	}

	// ── Phase 1: Syntax (Lex + Parse) ─────────────────────────────────────────
	parseResult := syntax.Parse(string(code))
	if len(parseResult.Errors) > 0 {
		fmt.Fprintln(os.Stderr, "syntax errors:")
		for _, e := range parseResult.Errors {
			fmt.Fprintln(os.Stderr, " ", e)
		}
		os.Exit(1)
	}

	// ── Phase 2: Frontend (Translate CST → AST, then Semantic Analysis) ───────
	astFile := frontend.Translate(parseResult.Root)

	analyzer := frontend.NewAnalyzer()
	if err := analyzer.Analyze(astFile); err != nil {
		fmt.Fprintf(os.Stderr, "semantic error: %v\n", err)
		os.Exit(1)
	}

	// ── Phase 3: Lowering ─────────────────────────────────────────────────────
	// Rewrites high-level constructs into explicit IR-ready logic:
	//   async fn  → coroutine state-machine / thread spawn
	//   defer     → explicit calls at all exit points
	//   var       → reference-count increment/decrement injection
	lower.NewLowerer(astFile).Apply()

	if *debugAST {
		fmt.Fprintln(os.Stderr, "── lowered AST ──────────────────────────────")
		fmt.Fprintf(os.Stderr, "%+v\n", astFile)
		fmt.Fprintln(os.Stderr, "─────────────────────────────────────────────")
	}

	// ── Phase 4: Code Generation (AST → IR) ───────────────────────────────────
	ext        := filepath.Ext(*sourceFile)
	moduleName := strings.TrimSuffix(filepath.Base(*sourceFile), ext)

	gen := codegen.New(moduleName)
	irModule, err := gen.Generate(astFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "codegen error: %v\n", err)
		os.Exit(1)
	}

	// ── Phase 5: Emit IR ──────────────────────────────────────────────────────
	ir := irModule.String()

	if *outputFile != "" {
		if err := os.WriteFile(*outputFile, []byte(ir), 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "error: cannot write '%s': %v\n", *outputFile, err)
			os.Exit(1)
		}
		fmt.Printf("compiled '%s'  →  %s\n", *sourceFile, *outputFile)
	} else {
		fmt.Print(ir)
	}
}