package context

import (
	"github.com/arc-language/arc-lang/builder/builder"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/symbol"
)

// Context holds the state shared during IR generation
type Context struct {
	Builder *builder.Builder
	Module  *ir.Module
	Logger  *Logger // Logger is defined in logging.go in this folder
	
	// Helper state for IR generation
	CurrentFunction *ir.Function
	CurrentBlock    *ir.BasicBlock
}

// NewContext creates a new compilation context
func NewContext(moduleName string) *Context {
	b := builder.New()
	mod := b.CreateModule(moduleName)
	
	return &Context{
		Builder: b,
		Module:  mod,
		Logger:  NewLogger("[Context]"),
	}
}

// EnterFunction sets up context for compiling a function
func (c *Context) EnterFunction(fn *ir.Function) {
	c.CurrentFunction = fn
	c.CurrentBlock = nil
	
	// Create entry block automatically if preferred, or handle in visitor
	entry := c.Builder.CreateBlock("entry")
	c.SetInsertBlock(entry)
}

// ExitFunction cleans up after compiling a function
func (c *Context) ExitFunction() {
	c.CurrentFunction = nil
	c.CurrentBlock = nil
}

// SetInsertBlock sets the current basic block for instruction insertion
func (c *Context) SetInsertBlock(block *ir.BasicBlock) {
	c.CurrentBlock = block
	c.Builder.SetInsertPoint(block)
}