package compiler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml" // Needed to parse index.toml files
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/context"
	"github.com/arc-language/arc-lang/diagnostic"
	"github.com/arc-language/arc-lang/irgen"
	//"github.com/arc-language/arc-lang/optimizer"
	"github.com/arc-language/arc-lang/pkg"
	"github.com/arc-language/arc-lang/semantics"
	"github.com/arc-language/arc-lang/symbol"
	"github.com/arc-language/upkg"
)

// Compiler holds the context for the compilation process
type Compiler struct {
	logger         *context.Logger
	Importer       *Importer
	PackageManager *pkg.PackageManager
	
	// NativeLibs holds the list of libraries discovered from 'import c "..."'
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

// CompileProject handles the frontend pipeline
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

			rawImport := decl.STRING_LITERAL().GetText()
			importPath := rawImport
			if len(importPath) >= 2 {
				importPath = importPath[1 : len(importPath)-1]
			}

			lang := ""
			if decl.IDENTIFIER() != nil {
				lang = decl.IDENTIFIER().GetText()
			}

			// Ensure Package is Installed (Clone Git Repo or Check Registry)
			downloadedPath, err := c.PackageManager.Ensure(lang, importPath)
			if err != nil {
				c.logger.Error("Failed to resolve package '%s': %v", importPath, err)
				return nil, fmt.Errorf("dependency resolution failed")
			}

			// Handle Arc Imports (Source Scan)
			if lang == "" {
				var absDir string
				if downloadedPath != "" {
					absDir = downloadedPath
				} else {
					absDir, err = c.Importer.ResolveImport(currentPath, importPath)
					if err != nil {
						c.logger.Error("Could not resolve local import '%s': %v", importPath, err)
						continue
					}
				}

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
				// Handle Native Imports (C/C++)
				c.logger.Debug("Resolved Native Dependency: %s (Language: %s)", importPath, lang)
				
				// CASE A: Remote/Git Wrapper (contains "/" or ".")
				if strings.Contains(importPath, "/") || strings.Contains(importPath, ".") {
					// We look for index.toml in the downloaded directory
					tomlPath := filepath.Join(downloadedPath, "index.toml")
					
					if _, err := os.Stat(tomlPath); err == nil {
						var entry upkg.RegistryEntry
						if _, err := toml.DecodeFile(tomlPath, &entry); err != nil {
							c.logger.Error("Failed to parse index.toml in %s: %v", tomlPath, err)
						} else {
							if len(entry.Libs) > 0 {
								c.logger.Info("Loaded custom libs from %s: %v", importPath, entry.Libs)
								c.NativeLibs = append(c.NativeLibs, entry.Libs...)
							}
						}
					} else {
						c.logger.Error("Native import '%s' resolved to %s but no index.toml found", importPath, downloadedPath)
					}

				} else {
					// CASE B: System Package (e.g. "sqlite3") - Use Registry
					upkgCfg := upkg.DefaultConfig()
					mgr, err := upkg.NewManager(upkg.BackendAuto, upkgCfg)
					if err == nil {
						entry, err := mgr.GetRegistryEntry(importPath)
						if err == nil && len(entry.Libs) > 0 {
							c.logger.Info("Auto-detected libraries for '%s': %v", importPath, entry.Libs)
							c.NativeLibs = append(c.NativeLibs, entry.Libs...)
						} else {
							// Fallback: If no registry entry, assume package name == lib name
							c.logger.Debug("No registry info for '%s', assuming library name match", importPath)
							c.NativeLibs = append(c.NativeLibs, importPath)
						}
						mgr.Close()
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

	c.logger.Debug("Phase 2.0: Type Discovery")
	for _, unit := range units {
		analyzer := semantics.NewAnalyzer(globalScope, unit.Path, semanticErrors)
		analyzer.Phase = 0
		analyzer.Analyze(unit.Tree, analysisRes)
	}

	c.logger.Debug("Phase 2.1: Semantic Declarations")
	for _, unit := range units {
		analyzer := semantics.NewAnalyzer(globalScope, unit.Path, semanticErrors)
		analyzer.Phase = 1
		analyzer.Analyze(unit.Tree, analysisRes)
	}

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