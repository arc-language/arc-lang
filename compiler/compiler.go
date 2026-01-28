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
	"github.com/arc-language/arc-lang/optimizer"
	"github.com/arc-language/arc-lang/pkg" // Import the new pkg manager
	"github.com/arc-language/arc-lang/semantics"
	"github.com/arc-language/arc-lang/symbol"
)

// Compiler holds the context for the compilation process
type Compiler struct {
	logger         *context.Logger
	Importer       *Importer
	PackageManager *pkg.PackageManager // Handle package downloads
}

func NewCompiler() *Compiler {
	return &Compiler{
		logger:         context.NewLogger("[Driver]"),
		Importer:       NewImporter(),
		PackageManager: pkg.NewPackageManager(), // Initialize pointing to ~/.arc/
	}
}

// CompileProject handles the frontend pipeline:
// 1. Discovery (Parsing & Downloading Dependencies)
// 2. Semantics
// 3. IR Gen
// 4. Optimization
func (c *Compiler) CompileProject(entryFile string) (*ir.Module, error) {
	absEntry, err := filepath.Abs(entryFile)
	if err != nil {
		return nil, err
	}

	c.logger.Info("Compiling project starting at: %s", absEntry)

	// --- PHASE 1: DISCOVERY (Parsing & Downloading) ---
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

		// Process Imports
		for _, decl := range tree.AllImportDecl() {
			if decl.STRING_LITERAL() == nil {
				continue
			}

			// 1. Clean Import Path (remove quotes)
			rawImport := decl.STRING_LITERAL().GetText()
			importPath := rawImport
			if len(importPath) >= 2 {
				importPath = importPath[1 : len(importPath)-1]
			}

			// 2. Detect Language Prefix (e.g., import c "...", import cpp "...")
			lang := ""
			if decl.IDENTIFIER() != nil {
				lang = decl.IDENTIFIER().GetText()
			}

			// 3. Ensure Package is Downloaded (Git, Nix, Brew)
			// If it's a remote package, this returns the path to ~/.arc/dest
			// If it's local/system, it returns empty string
			downloadedPath, err := c.PackageManager.Ensure(lang, importPath)
			if err != nil {
				c.logger.Error("Failed to resolve package '%s': %v", importPath, err)
				return nil, fmt.Errorf("dependency resolution failed")
			}

			// 4. Determine Absolute Directory to scan
			var absDir string
			if downloadedPath != "" {
				// Use the path provided by the package manager
				absDir = downloadedPath
			} else {
				// Local path resolution (relative to current source file)
				absDir, err = c.Importer.ResolveImport(currentPath, importPath)
			}

			// 5. Scan directory for source files (.ax)
			if err == nil {
				sources, _ := c.Importer.GetSourceFiles(absDir)
				for _, src := range sources {
					if !processed[src] {
						fileQueue = append(fileQueue, src)
					}
				}
			} else {
				// Only error if it wasn't a system library import (like "c" lib) which might not have source files
				if downloadedPath == "" && !filepath.IsAbs(importPath) {
					c.logger.Debug("Could not resolve local import '%s' - assuming system library", importPath)
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

	// --- PHASE 2.0: TYPE DISCOVERY ---
	c.logger.Debug("Phase 2.0: Type Discovery")
	for _, unit := range units {
		analyzer := semantics.NewAnalyzer(globalScope, unit.Path, semanticErrors)
		analyzer.Phase = 0
		analyzer.Analyze(unit.Tree, analysisRes)
	}

	// --- PHASE 2.1: DECLARATION SCAN ---
	c.logger.Debug("Phase 2.1: Semantic Declarations")
	for _, unit := range units {
		analyzer := semantics.NewAnalyzer(globalScope, unit.Path, semanticErrors)
		analyzer.Phase = 1
		analyzer.Analyze(unit.Tree, analysisRes)
	}

	// --- PHASE 2.2: BODY ANALYSIS ---
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
	c.logger.Debug("Phase 4: Optimization")
	dce := optimizer.NewDCE()
	dce.Run(module)

	return module, nil
}

func (c *Compiler) printDiagnostics(bag *diagnostic.Bag) {
	for _, err := range bag.Errors {
		fmt.Fprintf(os.Stderr, "%s\n", err.String())
	}
}