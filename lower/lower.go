package lower

import (
	"github.com/arc-language/arc-lang/ast"
)

// Lowerer manages the transformation of the AST.
type Lowerer struct {
	File *ast.File
}

func NewLowerer(file *ast.File) *Lowerer {
	return &Lowerer{File: file}
}

// Apply runs all lowering passes.
// 1. Async Lowering (State Machine generation)
// 2. Defer Resolution (Injects logic at exit points)
// 3. ARC Injection (Injects memory management)
func (l *Lowerer) Apply() {
	// 1. Lower Async functions (transform to state machines)
	// This usually generates new structs and methods, so it must happen first.
	l.lowerAsync()

	// 2. Resolve 'defer' statements
	// This expands 'defer' into explicit calls at every return/break/scope-exit.
	l.lowerDefer()

	// 3. Inject ARC (Automatic Reference Counting)
	// This analyzes variable lifetimes and inserts decref() calls.
	l.injectARC()
}

func (l *Lowerer) lowerAsync() {
	visitor := &asyncLowerer{file: l.File}
	visitor.walk(l.File)
}

func (l *Lowerer) lowerDefer() {
	visitor := &deferLowerer{}
	visitor.walk(l.File)
}

func (l *Lowerer) injectARC() {
	visitor := &arcInjector{}
	visitor.walk(l.File)
}