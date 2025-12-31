package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

// --- Statements ---

func (a *Analyzer) VisitBlock(ctx *parser.BlockContext) interface{} {
	// If this block is already mapped (e.g. by FuncDecl), use existing scope
	if _, mapped := a.scopes[ctx]; !mapped {
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
		
		if a.currentFuncRetType == nil {
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Return statement outside of function")
			return nil
		}

		if !areTypesCompatible(exprType, a.currentFuncRetType) {
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0,
				"Return type mismatch: expected '%s', got '%s'", 
				a.currentFuncRetType.String(), exprType.String())
		}
	} else {
		// Check if we expected a return value
		if a.currentFuncRetType != nil && a.currentFuncRetType != types.Void {
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Missing return value")
		}
	}
	return nil
}

func (a *Analyzer) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	// Check Condition
	condType := a.Visit(ctx.Expression(0)).(types.Type)
	if !condType.Equal(types.I1) {
		// Optionally enforce boolean checks here
	}
	
	a.Visit(ctx.Block(0))
	
	if len(ctx.AllBlock()) > 1 {
		a.Visit(ctx.Block(1))
	}
	
	// Handle else-if blocks if your grammar supports them in the list
	// usually ctx.AllIf() or similar
	
	return nil
}