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

// compileFileInternal handles the parsing and visiting of a single file
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

func (c *Compiler) compileFileInternal(filename string, isEntry bool) (*ir.Module, error) {
	input, err := antlr.NewFileStream(filename)
	if err != nil {
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

// Reset resets the compiler state
func (c *Compiler) Reset() {
	c.logger.Info("Resetting compiler state")
	c.context = NewContext(c.context.Importer.entryDir, c.context.Module.Name)
}