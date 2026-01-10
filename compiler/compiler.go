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
func (c *Compiler) CompileProject(entryFile string) (*ir.Module, error) {
	absEntry, err := filepath.Abs(entryFile)
	if err != nil {
		return nil, err
	}

	c.logger.Info("Compiling project starting at: %s", absEntry)

	// --- PHASE 1: DISCOVERY ---
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
		
		// FIX: Call local Parse function (defined in parse.go)
		tree, errs := Parse(currentPath) 
		
		if errs.HasErrors() {
			c.printDiagnostics(errs)
			return nil, fmt.Errorf("parsing failed in %s", currentPath)
		}
		
		units = append(units, &irgen.SourceUnit{Path: currentPath, Tree: tree})

		// Scan imports
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

	// --- PHASE 2: SEMANTIC ANALYSIS ---
	globalScope := symbol.NewScope(nil)
	symbol.InitGlobalScope(globalScope)

	analysisRes := &semantics.AnalysisResult{
		GlobalScope:   globalScope,
		Scopes:        make(map[antlr.ParserRuleContext]*symbol.Scope),
		NodeTypes:     make(map[antlr.ParseTree]types.Type),
		StructIndices: make(map[string]map[string]int),
	}
	semanticErrors := diagnostic.NewBag()

	// Pass 1: Declaration Scan
	c.logger.Debug("Phase 2.1: Semantic Declarations")
	for _, unit := range units {
		analyzer := semantics.NewAnalyzer(globalScope, unit.Path, semanticErrors)
		analyzer.Phase = 1
		analyzer.Analyze(unit.Tree, analysisRes)
	}

	// Pass 2: Body Analysis
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

	return module, nil
}

func (c *Compiler) printDiagnostics(bag *diagnostic.Bag) {
	for _, err := range bag.Errors {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
	}
}