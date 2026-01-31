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
	"github.com/arc-language/arc-lang/pkg"
	"github.com/arc-language/arc-lang/semantics"
	"github.com/arc-language/arc-lang/symbol"
	"github.com/arc-language/upkg" // Import upkg to access the Registry
)

// Compiler holds the context for the compilation process
type Compiler struct {
	logger         *context.Logger
	Importer       *Importer
	PackageManager *pkg.PackageManager
	
	// NativeLibs holds the list of libraries discovered from 'import c "..."'
	// e.g. ["ssl", "crypto", "sqlite3"]
	NativeLibs []string
}

func NewCompiler() *Compiler {
	return &Compiler{
		logger:         context.NewLogger("[Driver]"),
		Importer:       NewImporter(),
		PackageManager: pkg.NewPackageManager(),
		NativeLibs:     make([]string, 0),
	}
}

// CompileProject handles the frontend pipeline:
// 1. Discovery (Parsing & Downloading Dependencies & Registry Lookup)
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

			// 2. Detect Language Prefix (e.g., import c "sqlite", import cpp "...")
			lang := ""
			if decl.IDENTIFIER() != nil {
				lang = decl.IDENTIFIER().GetText()
			}

			// 3. Ensure Package is Installed
			downloadedPath, err := c.PackageManager.Ensure(lang, importPath)
			if err != nil {
				c.logger.Error("Failed to resolve package '%s': %v", importPath, err)
				return nil, fmt.Errorf("dependency resolution failed")
			}

			// 4. Handle Source Files (Only for Arc imports)
			if lang == "" {
				var absDir string

				if downloadedPath != "" {
					// It was a remote module (github, etc) downloaded to ~/.upkg/src/...
					absDir = downloadedPath
				} else {
					// It's likely a local relative import (e.g. import "./utils")
					absDir, err = c.Importer.ResolveImport(currentPath, importPath)
					if err != nil {
						c.logger.Error("Could not resolve local import '%s': %v", importPath, err)
						continue
					}
				}

				// Scan directory for .ax source files and add to queue
				if absDir != "" {
					sources, err := c.Importer.GetSourceFiles(absDir)
					if err != nil {
						c.logger.Debug("Warning: No sources found in resolved path %s", absDir)
					}
					for _, src := range sources {
						if !processed[src] {
							fileQueue = append(fileQueue, src)
						}
					}
				}
			} else {
				// 5. Handle Native Imports (C/C++) - Auto-Linker Logic
				c.logger.Debug("Resolved Native Dependency: %s (Language: %s)", importPath, lang)
				
				// Initialize upkg temporarily to query the registry
				// We need to know which libraries to link (e.g., "openssl" -> ["ssl", "crypto"])
				upkgCfg := upkg.DefaultConfig()
				mgr, err := upkg.NewManager(upkg.BackendAuto, upkgCfg)
				if err == nil {
					// Try to get registry info
					entry, err := mgr.GetRegistryEntry(importPath)
					if err == nil && len(entry.Libs) > 0 {
						c.logger.Info("Auto-detected libraries for '%s': %v", importPath, entry.Libs)
						c.NativeLibs = append(c.NativeLibs, entry.Libs...)
					} else {
						// Fallback: If no registry entry (maybe system installed manually), 
						// assume the package name is the library name.
						c.logger.Debug("No registry info for '%s', assuming library name match", importPath)
						c.NativeLibs = append(c.NativeLibs, importPath)
					}
					mgr.Close()
				} else {
					c.logger.Error("Failed to init upkg for registry lookup: %v", err)
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