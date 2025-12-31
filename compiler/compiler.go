package compiler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/codegen/codegen" // Assumes your codegen logic is here
	"github.com/arc-language/arc-lang/diagnostic"
	"github.com/arc-language/arc-lang/irgen"
	"github.com/arc-language/arc-lang/semantics"
)

// Compiler coordinates the compilation pipeline
type Compiler struct {
	logger   *Logger   // Internal debug logger (Info/Debug)
	Importer *Importer // Handles file resolution
}

func NewCompiler() *Compiler {
	return &Compiler{
		logger:   NewLogger("[Driver]"),
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
	// Catches: Missing semicolons, bad keywords, syntax errors
	c.logger.Debug("Phase 1: Parsing")
	tree, syntaxErrors := Parse(absPath)
	
	if syntaxErrors.HasErrors() {
		c.printDiagnostics(syntaxErrors)
		return nil, fmt.Errorf("parsing failed with %d errors", len(syntaxErrors.Errors))
	}

	// --- PHASE 2: SEMANTIC ANALYSIS ---
	// Validates logic, types, and scopes
	// Catches: Undefined vars, type mismatches, const re-assignment
	c.logger.Debug("Phase 2: Semantic Analysis")
	
	// We create a new bag for semantic errors
	semanticErrors := diagnostic.NewBag()
	
	// Analyze returns an AnalysisResult containing the GlobalScope and TypeMaps
	analysis, err := semantics.Analyze(tree, absPath, semanticErrors)
	if err != nil || semanticErrors.HasErrors() {
		c.printDiagnostics(semanticErrors)
		return nil, fmt.Errorf("semantic analysis failed")
	}

	// --- PHASE 3: IR GENERATION ---
	// Translates validated AST -> LLVM IR
	// Should not produce user errors if Phase 2 worked correctly
	c.logger.Debug("Phase 3: IR Generation")
	
	// We pass the module name (usually filename) and the analysis data
	moduleName := filepath.Base(absPath)
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
		// In the future, add color codes here (Red for Error, Yellow for Warn)
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
	}
}