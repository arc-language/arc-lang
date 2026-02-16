package lower

import (
	"github.com/arc-language/arc-lang/ast"
)

type deferLowerer struct {
	defers []*ast.CallExpr // Stack of deferred calls in current scope
}

func (l *deferLowerer) walk(node ast.Node) {
	switch n := node.(type) {
	case *ast.File:
		for _, decl := range n.Decls {
			l.walk(decl)
		}
	case *ast.FuncDecl:
		if n.Body != nil {
			l.walkBlock(n.Body)
		}
	}
}

func (l *deferLowerer) walkBlock(b *ast.BlockStmt) {
	// Save parent state
	parentDefers := l.defers
	l.defers = make([]*ast.CallExpr, 0)

	var newStmts []ast.Stmt

	// 1. Scan statements
	for _, stmt := range b.List {
		switch s := stmt.(type) {
		case *ast.DeferStmt:
			// Push to stack (LIFO order is handled by iterating backwards later)
			l.defers = append(l.defers, s.Call)
			// Remove 'defer' stmt from the AST (it is now meta-data)
			continue

		case *ast.ReturnStmt:
			// Injection Point: Inject all current defers before the return
			l.injectDefers(&newStmts)
			newStmts = append(newStmts, s)

		case *ast.BlockStmt:
			l.walkBlock(s)
			newStmts = append(newStmts, s)
			
		case *ast.IfStmt:
			l.walkIf(s)
			newStmts = append(newStmts, s)

		default:
			newStmts = append(newStmts, s)
		}
	}

	// 2. End of Block Injection
	// If the block falls through (reaches the end), defers must run.
	// (Unless it ended with a return, but strictly appending is safe if return is terminal)
	l.injectDefers(&newStmts)

	b.List = newStmts
	
	// Restore parent state
	l.defers = parentDefers
}

func (l *deferLowerer) walkIf(s *ast.IfStmt) {
	l.walkBlock(s.Body)
	if s.Else != nil {
		switch e := s.Else.(type) {
		case *ast.BlockStmt:
			l.walkBlock(e)
		case *ast.IfStmt:
			l.walkIf(e)
		}
	}
}

func (l *deferLowerer) injectDefers(target *[]ast.Stmt) {
	// Defers run in LIFO (reverse) order
	for i := len(l.defers) - 1; i >= 0; i-- {
		call := l.defers[i]
		// Create a Stmt for the call
		*target = append(*target, &ast.ExprStmt{X: call})
	}
}