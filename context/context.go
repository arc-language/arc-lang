package context

import (
	"github.com/arc-language/arc-lang/builder/builder"
	"github.com/arc-language/arc-lang/builder/ir"
)

// Context holds the state shared during IR generation
type Context struct {
	Builder *builder.Builder
	Module  *ir.Module
	Logger  *Logger
	
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
	
	// Create entry block and manually attach it to function
	entry := ir.NewBasicBlock("entry")
	entry.Parent = fn
	fn.Blocks = append(fn.Blocks, entry)
	
	// Now SetInsertPoint will correctly set builder's currentFunc
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