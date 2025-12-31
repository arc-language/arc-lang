package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (a *Analyzer) VisitBlock(ctx *parser.BlockContext) interface{} {
	// Only push a new scope if this block isn't already mapped (e.g. Function body)
	shouldPush := true
	if _, exists := a.scopes[ctx]; exists {
		shouldPush = false
		// If mapped, we are already in the scope thanks to VisitFunctionDecl or we need to switch?
		// VisitFunctionDecl pushes scope, maps it, calls children, then pops.
		// So when we arrive here via a.Visit(stmt) -> VisitBlock, we are ALREADY in the scope.
	}

	if shouldPush {
		a.pushScope(ctx)
		defer a.popScope()
	}

	for _, stmt := range ctx.AllStatement() {
		a.Visit(stmt)
	}
	return nil
}

func (a *Analyzer) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	if ctx.Expression() != nil {
		exprType := a.Visit(ctx.Expression()).(types.Type)
		
		if a.currentFuncRetType != nil && !areTypesCompatible(exprType, a.currentFuncRetType) {
			// Warn or Error
		}
	}
	return nil
}

func (a *Analyzer) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	a.Visit(ctx.Expression(0))
	a.Visit(ctx.Block(0))
	if len(ctx.AllBlock()) > 1 {
		a.Visit(ctx.Block(1))
	}
	return nil
}