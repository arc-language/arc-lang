package lower

import (
	"github.com/arc-language/arc-lang/ast"
)

type arcInjector struct {
	vars []*ast.VarDecl // Track active 'var' declarations in scope
}

func (l *arcInjector) walk(node ast.Node) {
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

func (l *arcInjector) walkBlock(b *ast.BlockStmt) {
	parentVars := l.vars
	l.vars = make([]*ast.VarDecl, 0)

	var newStmts []ast.Stmt

	for _, stmt := range b.List {
		// 1. Check for new var declarations
		if declStmt, ok := stmt.(*ast.DeclStmt); ok {
			if declStmt.Decl.IsRef { // Only track 'var', not 'let'
				l.vars = append(l.vars, declStmt.Decl)
				// Inject 'incref' immediately after declaration? 
				// Usually declaration init sets ref=1 or steals reference. 
				// We assume init handles ref=1.
			}
		}

		// 2. Handle Returns (Inject decrefs before return)
		if ret, ok := stmt.(*ast.ReturnStmt); ok {
			l.injectDecrefs(&newStmts)
			newStmts = append(newStmts, ret)
			continue
		}

		// 3. Recurse
		if nestedBlock, ok := stmt.(*ast.BlockStmt); ok {
			l.walkBlock(nestedBlock)
		}
		
		newStmts = append(newStmts, stmt)
	}

	// 4. End of Block (Inject decrefs for vars falling out of scope)
	l.injectDecrefs(&newStmts)

	b.List = newStmts
	l.vars = parentVars
}

func (l *arcInjector) injectDecrefs(target *[]ast.Stmt) {
	// Decrefs run in LIFO order (mirroring destruction)
	for i := len(l.vars) - 1; i >= 0; i-- {
		v := l.vars[i]
		
		// Generate: decref(v)
		// Since 'decref' is an intrinsic/runtime call, we synthesize a CallExpr.
		// In a real compiler, this resolves to a specific runtime symbol.
		decrefCall := &ast.CallExpr{
			Fun: &ast.Ident{Name: "decref"}, // Magic compiler intrinsic
			Args: []ast.Expr{
				&ast.Ident{Name: v.Name},
			},
		}
		
		*target = append(*target, &ast.ExprStmt{X: decrefCall})
	}
}