package compiler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/codegen/codegen"
	"github.com/arc-language/arc-lang/context"
	"github.com/arc-language/arc-lang/diagnostic"
	"github.com/arc-language/arc-lang/irgen"
	"github.com/arc-language/arc-lang/semantics"
)

// Compiler coordinates the compilation pipeline
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

// CompileFile is the main API to compile a single source file
func (c *Compiler) CompileFile(path string) (*ir.Module, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %v", err)
	}

	c.logger.Info("Compiling file: %s", absPath)

	// --- PHASE 1: PARSING ---
	// Converts source text -> AST
	c.logger.Debug("Phase 1: Parsing")
	tree, syntaxErrors := Parse(absPath)
	
	if syntaxErrors.HasErrors() {
		c.printDiagnostics(syntaxErrors)
		return nil, fmt.Errorf("parsing failed with %d errors", len(syntaxErrors.Errors))
	}

	// --- PHASE 2: SEMANTIC ANALYSIS ---
	// Validates logic, types, and scopes
	c.logger.Debug("Phase 2: Semantic Analysis")
	
	semanticErrors := diagnostic.NewBag()
	
	// Analyze returns an AnalysisResult containing the GlobalScope and TypeMaps
	analysis, err := semantics.Analyze(tree, absPath, semanticErrors)
	if err != nil || semanticErrors.HasErrors() {
		c.printDiagnostics(semanticErrors)
		return nil, fmt.Errorf("semantic analysis failed")
	}

	// --- PHASE 3: IR GENERATION ---
	// Translates validated AST -> LLVM IR
	c.logger.Debug("Phase 3: IR Generation")
	
	moduleName := filepath.Base(absPath)
	
	// Pass the analysis results to the generator
	module := irgen.Generate(tree, moduleName, analysis)

	c.logger.Info("Compilation successful. Module '%s' created.", module.Name)
	return module, nil
}

// CompileToExecutable wraps the pipeline and writes a binary
func (c *Compiler) CompileToExecutable(sourcePath, outputPath string) error {
	// 1. Run Pipeline
	mod, err := c.CompileFile(sourcePath)
	if err != nil {
		return err
	}

	// 2. Code Generation (Object -> Linked Exe)
	c.logger.Info("Generating executable: %s", outputPath)
	exeData, err := codegen.GenerateExecutable(mod)
	if err != nil {
		return fmt.Errorf("codegen failed: %v", err)
	}

	// 3. Write File
	if err := os.WriteFile(outputPath, exeData, 0755); err != nil {
		return fmt.Errorf("failed to write executable: %v", err)
	}

	return nil
}

// printDiagnostics formats and prints errors to stderr
func (c *Compiler) printDiagnostics(bag *diagnostic.Bag) {
	for _, err := range bag.Errors {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
	}
}