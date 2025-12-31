package compiler

import (
	"fmt"
	"path/filepath"

	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/parser"
)

// Compiler represents the Arc language compiler
type Compiler struct {
	context               *Context
	logger                *Logger
	asyncTransformEnabled bool
}

// NewCompiler creates a new compiler instance
func NewCompiler(moduleName string, entryFile string) *Compiler {
	logger := NewLogger(fmt.Sprintf("[Compiler:%s]", moduleName))
	logger.Info("Creating compiler for module '%s' with entry file '%s'", moduleName, entryFile)
	
	return &Compiler{
		context:               NewContext(entryFile, moduleName),
		logger:                logger,
		asyncTransformEnabled: true, // Async/Await enabled by default
	}
}

// syntaxErrorListener captures ANTLR syntax errors
type syntaxErrorListener struct {
	*antlr.DefaultErrorListener
	logger   *Logger
	filename string
}

func newSyntaxErrorListener(logger *Logger, filename string) *syntaxErrorListener {
	return &syntaxErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		logger:               logger,
		filename:             filename,
	}
}

func (l *syntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.logger.ErrorAt(l.filename, line, column, "%s", msg)
}

// CompileFile compiles an Arc source file to IR
func (c *Compiler) CompileFile(filename string) (*ir.Module, error) {
	c.logger.Info("Compiling file: %s", filename)
	
	absPath, err := filepath.Abs(filename)
	if err != nil {
		c.logger.Error("Failed to resolve path '%s': %v", filename, err)
		return nil, fmt.Errorf("failed to resolve path: %v", err)
	}

	return c.compileFileInternal(absPath, true)
}

// CompilePackage compiles all files in a directory as a single package
func (c *Compiler) CompilePackage(dirPath string) (*PackageInfo, error) {
	c.logger.Debug("Starting package compilation for directory: %s", dirPath)
	
	// 1. Check Cache
	if pkg, ok := c.context.Importer.GetPackage(dirPath); ok {
		if pkg.IsProcessing {
			c.logger.Error("Circular dependency detected importing '%s'", dirPath)
			return nil, fmt.Errorf("circular dependency detected importing %s", dirPath)
		}
		c.logger.Debug("Package '%s' found in cache", dirPath)
		return pkg, nil
	}

	// 2. Mark as processing
	pkgInfo := &PackageInfo{
		SourcePath:   dirPath,
		IsProcessing: true,
	}
	c.context.Importer.CachePackage(dirPath, pkgInfo)
	c.logger.Debug("Marked package '%s' as processing", dirPath)

	// 3. Find source files
	files, err := c.context.Importer.GetSourceFiles(dirPath)
	if err != nil {
		c.logger.Error("Failed to find source files in '%s': %v", dirPath, err)
		return nil, err
	}

	c.logger.Info("Compiling package at '%s' with %d file(s)", dirPath, len(files))

	// 4. Compile all files in directory
	var packageName string
	
	// Preserve current namespace to restore after compiling package
	prevNs := c.context.currentNamespace
	
	for i, file := range files {
		c.logger.Debug("Compiling file %d/%d: %s", i+1, len(files), file)
		
		// Reset namespace to root before parsing a new file in a package
		c.context.currentNamespace = c.context.rootNamespace
		
		_, err := c.compileFileInternal(file, false) 
		if err != nil {
			c.logger.Error("Compilation failed for file '%s': %v", file, err)
			return nil, err
		}
		
		// Validation: Verify package consistency
		currentNsName := c.context.currentNamespace.Name
		if currentNsName == "" {
			c.logger.Debug("File '%s' has no namespace declaration", file)
		} else {
			if packageName == "" {
				packageName = currentNsName
				c.logger.Debug("Package namespace set to '%s'", packageName)
			} else if currentNsName != packageName {
				c.logger.Error("File '%s' declares namespace '%s', expected '%s'", file, currentNsName, packageName)
				return nil, fmt.Errorf("file %s declares namespace '%s', expected '%s' (all files in a directory must belong to the same package)", 
					file, currentNsName, packageName)
			}
		}
	}

	// 5. Finalize
	pkgInfo.Name = packageName
	pkgInfo.Namespace = c.context.GetOrCreateNamespace(packageName)
	pkgInfo.IsProcessing = false
	
	// Restore namespace
	c.context.currentNamespace = prevNs
	
	c.logger.Info("Package '%s' compiled successfully (Namespace: %s)", dirPath, packageName)
	
	return pkgInfo, nil
}

// CompileString compiles Arc source code from a string
func (c *Compiler) CompileString(source string) (*ir.Module, error) {
	c.logger.Info("Compiling source string (%d bytes)", len(source))
	
	input := antlr.NewInputStream(source)
	lexer := parser.NewArcLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewArcParser(stream)
	
	p.RemoveErrorListeners()
	listener := newSyntaxErrorListener(c.context.Logger, "<string>")
	p.AddErrorListener(listener)
	
	tree := p.CompilationUnit()
	
	if c.context.Logger.HasErrors() {
		c.context.Logger.PrintSummary()
		return nil, fmt.Errorf("syntax errors found")
	}
	
	visitor := NewIRVisitor(c, "<string>")
	visitor.Visit(tree)
	
	if c.context.Logger.HasErrors() {
		c.context.Logger.PrintSummary()
		return nil, fmt.Errorf("compilation failed with %d error(s)", c.context.Logger.ErrorCount())
	}
	
	return c.context.Module, nil
}

// compileFileInternal handles the parsing and visiting of a single file
func (c *Compiler) compileFileInternal(filename string, isEntry bool) (*ir.Module, error) {
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		c.logger.Error("Failed to open file '%s': %v", filename, err)
		return nil, err
	}
	
	lexer := parser.NewArcLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewArcParser(stream)
	
	p.RemoveErrorListeners()
	p.AddErrorListener(newSyntaxErrorListener(c.context.Logger, filename))
	
	tree := p.CompilationUnit()
	
	if c.context.Logger.HasErrors() {
		if isEntry { c.context.Logger.PrintSummary() }
		return nil, fmt.Errorf("syntax errors in %s", filename)
	}
	
	NewIRVisitor(c, filename).Visit(tree)
	
	if c.context.Logger.HasErrors() {
		if isEntry { c.context.Logger.PrintSummary() }
		return nil, fmt.Errorf("compilation failed")
	}
	
	return c.context.Module, nil
}

// GetModule returns the compiled module
func (c *Compiler) GetModule() *ir.Module {
	return c.context.Module
}

// GetContext returns the compilation context
func (c *Compiler) GetContext() *Context {
	return c.context
}

// Reset resets the compiler state
func (c *Compiler) Reset() {
	c.logger.Info("Resetting compiler state")
	c.context = NewContext(c.context.Importer.entryDir, c.context.Module.Name)
}