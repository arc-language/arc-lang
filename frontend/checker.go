package frontend

import (
	"fmt"
	"github.com/arc-language/arc-lang/ast"
)

// Analyzer manages the semantic analysis passes.
type Analyzer struct {
	GlobalScope *Scope
	Errors      []error
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{
		GlobalScope: NewScope(nil, ScopeGlobal),
	}
}

// Analyze runs Pass 1 (Symbol Discovery) and Pass 2 (Type Checking).
func (a *Analyzer) Analyze(file *ast.File) error {
	// Pass 1: Discovery
	// Register top-level symbols (functions, structs) so they can be used out-of-order.
	a.discoverSymbols(file)

	// Pass 2: Type Checking
	// Dive into function bodies to check types and resolve references.
	a.checkTypes(file)

	if len(a.Errors) > 0 {
		return fmt.Errorf("semantic errors found: %v", a.Errors)
	}
	return nil
}

func (a *Analyzer) discoverSymbols(file *ast.File) {
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			sym := &Symbol{Name: d.Name, Kind: "func", Decl: d}
			a.GlobalScope.Insert(d.Name, sym)
		// Handle Interface, Enum, Var...
		}
	}
}

func (a *Analyzer) checkTypes(file *ast.File) {
	// Example: just walking decls
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			// Create function scope
			fnScope := NewScope(a.GlobalScope, ScopeFunc)
			// Add params to scope
			for _, param := range d.Params {
				sym := &Symbol{Name: param.Name, Kind: "param", Decl: param}
				fnScope.Insert(param.Name, sym)
			}
			// Check body
			if d.Body != nil {
				a.checkBlock(d.Body, fnScope)
			}
		}
	}
}

func (a *Analyzer) checkBlock(b *ast.BlockStmt, parent *Scope) {
	scope := NewScope(parent, ScopeBlock)
	for _, stmt := range b.List {
		a.checkStmt(stmt, scope)
	}
}

func (a *Analyzer) checkStmt(stmt ast.Stmt, scope *Scope) {
	// TODO: Implement statement checking logic
}