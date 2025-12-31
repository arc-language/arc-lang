package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

// ... (Arithmetic/Logic visitors same as before) ...

func (a *Analyzer) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.Literal() != nil { return a.Visit(ctx.Literal()) }
	
	if ctx.QualifiedIdentifier() != nil {
		q := ctx.QualifiedIdentifier()
		parts := ""
		for i, id := range q.AllIDENTIFIER() {
			if i > 0 { parts += "." }
			parts += id.GetText()
		}
		if sym, ok := a.currentScope.Resolve(parts); ok {
			return sym.Type
		}
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Undefined '%s'", parts)
		return types.Void
	}

	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		if s, ok := a.currentScope.Resolve(name); ok { return s.Type }
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Undefined %s", name)
	}
	if ctx.Expression() != nil { return a.Visit(ctx.Expression()) }
	return types.Void
}

// ... (Rest of expression visitors same as before) ...
func (a *Analyzer) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	if ctx.LogicalOrExpression() != nil {
		t := a.Visit(ctx.LogicalOrExpression()).(types.Type)
		a.nodeTypes[ctx] = t
		return t
	}
	return types.Void
}
func (a *Analyzer) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	lhs := a.Visit(ctx.MultiplicativeExpression(0)).(types.Type)
	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		a.Visit(ctx.MultiplicativeExpression(i))
	}
	return lhs
}
func (a *Analyzer) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	lhs := a.Visit(ctx.UnaryExpression(0)).(types.Type)
	for i := 1; i < len(ctx.AllUnaryExpression()); i++ {
		a.Visit(ctx.UnaryExpression(i))
	}
	return lhs
}
func (a *Analyzer) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if ctx.PostfixExpression() != nil { return a.Visit(ctx.PostfixExpression()) }
	val := a.Visit(ctx.UnaryExpression()).(types.Type)
	if ctx.AMP() != nil { return types.NewPointer(val) }
	if ctx.STAR() != nil {
		if ptr, ok := val.(*types.PointerType); ok { return ptr.ElementType }
	}
	return val
}
func (a *Analyzer) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	curr := a.Visit(ctx.PrimaryExpression()).(types.Type)
	for _, op := range ctx.AllPostfixOp() {
		if op.DOT() != nil && op.IDENTIFIER() != nil {
			fieldName := op.IDENTIFIER().GetText()
			if ptr, ok := curr.(*types.PointerType); ok { curr = ptr.ElementType }
			if st, ok := curr.(*types.StructType); ok {
				if indices, ok := a.structIndices[st.Name]; ok {
					if idx, ok := indices[fieldName]; ok {
						curr = st.Fields[idx]
						continue
					}
				}
			}
		}
		if op.LBRACKET() != nil {
			if ptr, ok := curr.(*types.PointerType); ok { curr = ptr.ElementType }
			else if arr, ok := curr.(*types.ArrayType); ok { curr = arr.ElementType }
		}
	}
	return curr
}
func (a *Analyzer) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTEGER_LITERAL() != nil { return types.I64 }
	if ctx.FLOAT_LITERAL() != nil { return types.F64 }
	if ctx.BOOLEAN_LITERAL() != nil { return types.I1 }
	if ctx.STRING_LITERAL() != nil { return types.NewPointer(types.I8) }
	return types.Void
}
func (a *Analyzer) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} { return a.Visit(ctx.LogicalAndExpression(0)) }
func (a *Analyzer) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} { return a.Visit(ctx.BitOrExpression(0)) }
func (a *Analyzer) VisitBitOrExpression(ctx *parser.BitOrExpressionContext) interface{} { return a.Visit(ctx.BitXorExpression(0)) }
func (a *Analyzer) VisitBitXorExpression(ctx *parser.BitXorExpressionContext) interface{} { return a.Visit(ctx.BitAndExpression(0)) }
func (a *Analyzer) VisitBitAndExpression(ctx *parser.BitAndExpressionContext) interface{} { return a.Visit(ctx.EqualityExpression(0)) }
func (a *Analyzer) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} { return a.Visit(ctx.RelationalExpression(0)) }
func (a *Analyzer) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} { return a.Visit(ctx.ShiftExpression(0)) }
func (a *Analyzer) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} { return a.Visit(ctx.RangeExpression(0)) }
func (a *Analyzer) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} { return a.Visit(ctx.AdditiveExpression(0)) }