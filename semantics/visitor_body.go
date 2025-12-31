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
	return nil
}

// --- Expressions ---

func (a *Analyzer) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	if ctx.LogicalOrExpression() != nil {
		t := a.Visit(ctx.LogicalOrExpression()).(types.Type)
		// Store result for Pass 2
		a.nodeTypes[ctx] = t
		return t
	}
	return types.Void
}

func (a *Analyzer) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	// Start with the first term
	lhs := a.Visit(ctx.MultiplicativeExpression(0)).(types.Type)
	
	// Check subsequent terms
	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		rhs := a.Visit(ctx.MultiplicativeExpression(i)).(types.Type)
		
		if !areTypesCompatible(lhs, rhs) {
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0,
				"Operator mismatch: cannot add '%s' and '%s'", lhs.String(), rhs.String())
			return types.Void // Poison
		}
	}
	
	return lhs
}

func (a *Analyzer) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	// 1. Literals
	if ctx.Literal() != nil {
		return a.Visit(ctx.Literal())
	}
	
	// 2. Variables (Identifiers)
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		sym, ok := a.currentScope.Resolve(name)
		if !ok {
			a.bag.Report(a.file, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
				"Undefined variable '%s'", name)
			return types.Void
		}
		return sym.Type
	}
	
	// 3. Parenthesized Expression
	if ctx.Expression() != nil {
		return a.Visit(ctx.Expression())
	}
	
	return types.Void
}

func (a *Analyzer) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTEGER_LITERAL() != nil { return types.I64 }
	if ctx.FLOAT_LITERAL() != nil { return types.F64 }
	if ctx.BOOLEAN_LITERAL() != nil { return types.I1 }
	if ctx.STRING_LITERAL() != nil { return types.NewPointer(types.I8) }
	return types.Void
}

func (a *Analyzer) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	return a.Visit(ctx.LogicalAndExpression(0))
}
func (a *Analyzer) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	return a.Visit(ctx.BitOrExpression(0))
}
func (a *Analyzer) VisitBitOrExpression(ctx *parser.BitOrExpressionContext) interface{} {
	return a.Visit(ctx.BitXorExpression(0))
}
func (a *Analyzer) VisitBitXorExpression(ctx *parser.BitXorExpressionContext) interface{} {
	return a.Visit(ctx.BitAndExpression(0))
}
func (a *Analyzer) VisitBitAndExpression(ctx *parser.BitAndExpressionContext) interface{} {
	return a.Visit(ctx.EqualityExpression(0))
}
func (a *Analyzer) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	return a.Visit(ctx.RelationalExpression(0))
}
func (a *Analyzer) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} {
	return a.Visit(ctx.ShiftExpression(0))
}
func (a *Analyzer) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} {
	return a.Visit(ctx.RangeExpression(0))
}
func (a *Analyzer) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} {
	return a.Visit(ctx.AdditiveExpression(0))
}
func (a *Analyzer) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	return a.Visit(ctx.UnaryExpression(0))
}
func (a *Analyzer) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if ctx.PostfixExpression() != nil { return a.Visit(ctx.PostfixExpression()) }
	return a.Visit(ctx.UnaryExpression()) 
}
func (a *Analyzer) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	return a.Visit(ctx.PrimaryExpression())
}