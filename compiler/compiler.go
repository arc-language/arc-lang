package compiler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/context"
	"github.com/arc-language/arc-lang/diagnostic"
	"github.com/arc-language/arc-lang/irgen"
	"github.com/arc-language/arc-lang/optimizer" // <--- Import the optimizer package
	"github.com/arc-language/arc-lang/semantics"
	"github.com/arc-language/arc-lang/symbol"
)

// Compiler holds the context for the compilation process
type Compiler struct {
	logger   *context.Logger
	Importer *Importer
}

func NewCompiler() *Compiler {
	return &Compiler{
		logger:   context.NewLogger("[Driver]"),
		Importer: NewImporter(),
	}
}

// CompileProject handles the frontend pipeline:
// 1. Discovery (Parsing)
// 2. Semantics
// 3. IR Gen
// 4. Optimization (NEW)
func (c *Compiler) CompileProject(entryFile string) (*ir.Module, error) {
	absEntry, err := filepath.Abs(entryFile)
	if err != nil {
		return nil, err
	}

	c.logger.Info("Compiling project starting at: %s", absEntry)

	// --- PHASE 1: DISCOVERY (Parsing) ---
	fileQueue := []string{absEntry}
	processed := make(map[string]bool)

	var units []*irgen.SourceUnit

	for i := 0; i < len(fileQueue); i++ {
		currentPath := fileQueue[i]
		if processed[currentPath] {
			continue
		}
		processed[currentPath] = true

		c.logger.Debug("Parsing: %s", currentPath)

		tree, errs := Parse(currentPath)

		if errs.HasErrors() {
			c.printDiagnostics(errs)
			return nil, fmt.Errorf("parsing failed in %s", currentPath)
		}

		units = append(units, &irgen.SourceUnit{Path: currentPath, Tree: tree})

		for _, decl := range tree.AllImportDecl() {
			if decl.STRING_LITERAL() == nil {
				continue
			}
			importStr := decl.STRING_LITERAL().GetText()
			if len(importStr) >= 2 {
				importStr = importStr[1 : len(importStr)-1]
			}
			absDir, err := c.Importer.ResolveImport(currentPath, importStr)
			if err == nil {
				sources, _ := c.Importer.GetSourceFiles(absDir)
				for _, src := range sources {
					if !processed[src] {
						fileQueue = append(fileQueue, src)
					}
				}
			}
		}
	}

	// --- SETUP SEMANTICS ---
	globalScope := symbol.NewScope(nil)
	symbol.InitGlobalScope(globalScope)

	analysisRes := &semantics.AnalysisResult{
		GlobalScope:   globalScope,
		Scopes:        make(map[antlr.ParserRuleContext]*symbol.Scope),
		NodeTypes:     make(map[antlr.ParseTree]types.Type),
		StructIndices: make(map[string]map[string]int),
	}
	semanticErrors := diagnostic.NewBag()

	// --- PHASE 2.0: TYPE DISCOVERY (NEW) ---
	// Just register Struct/Class/Enum names. No field resolution yet.
	c.logger.Debug("Phase 2.0: Type Discovery")
	for _, unit := range units {
		analyzer := semantics.NewAnalyzer(globalScope, unit.Path, semanticErrors)
		analyzer.Phase = 0
		analyzer.Analyze(unit.Tree, analysisRes)
	}

	// --- PHASE 2.1: DECLARATION SCAN ---
	// Resolve Struct Fields and Function Signatures.
	c.logger.Debug("Phase 2.1: Semantic Declarations")
	for _, unit := range units {
		analyzer := semantics.NewAnalyzer(globalScope, unit.Path, semanticErrors)
		analyzer.Phase = 1
		analyzer.Analyze(unit.Tree, analysisRes)
	}

	// --- PHASE 2.2: BODY ANALYSIS ---
	// Compile statements and expressions.
	c.logger.Debug("Phase 2.2: Semantic Bodies")
	for _, unit := range units {
		analyzer := semantics.NewAnalyzer(globalScope, unit.Path, semanticErrors)
		analyzer.Phase = 2
		analyzer.Analyze(unit.Tree, analysisRes)
	}

	if semanticErrors.HasErrors() {
		c.printDiagnostics(semanticErrors)
		return nil, fmt.Errorf("semantic analysis failed")
	}

	// --- PHASE 3: IR GENERATION ---
	c.logger.Debug("Phase 3: IR Generation")
	moduleName := filepath.Base(absEntry)

	module := irgen.GenerateProject(units, moduleName, analysisRes)

	// --- PHASE 4: OPTIMIZATION ---
	//c.logger.Debug("Phase 4: Optimization")
	//dce := optimizer.NewDCE()
	//dce.Run(module)

	return module, nil
}

func (c *Compiler) printDiagnostics(bag *diagnostic.Bag) {
	for _, err := range bag.Errors {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
	}
}