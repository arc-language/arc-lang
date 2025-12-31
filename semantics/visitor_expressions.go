package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

// --- Expressions ---

func (a *Analyzer) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	if ctx.LogicalOrExpression() != nil {
		return a.Visit(ctx.LogicalOrExpression())
	}
	return types.Void
}

// --- Binary Ops ---

func (a *Analyzer) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	lhs := a.Visit(ctx.MultiplicativeExpression(0)).(types.Type)
	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		rhs := a.Visit(ctx.MultiplicativeExpression(i)).(types.Type)
		// Check compatibility (omitted for brevity, assume compatible)
		_ = rhs 
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

// --- Unary Ops ---

func (a *Analyzer) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if ctx.PostfixExpression() != nil {
		return a.Visit(ctx.PostfixExpression())
	}
	
	// Handle Unary Operators
	if ctx.UnaryExpression() != nil {
		valType := a.Visit(ctx.UnaryExpression()).(types.Type)

		if ctx.AMP() != nil {
			// Address-of: &T -> *T
			return types.NewPointer(valType)
		}
		
		if ctx.STAR() != nil {
			// Dereference: *T -> T
			if ptr, ok := valType.(*types.PointerType); ok {
				return ptr.ElementType
			}
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Cannot dereference non-pointer type '%s'", valType)
			return types.Void
		}
		
		return valType
	}
	
	return types.Void
}

// --- Casts ---

func (a *Analyzer) VisitCastExpression(ctx *parser.CastExpressionContext) interface{} {
	// Visit expression to ensure valid symbol usage
	a.Visit(ctx.Expression())
	// Return the explicit target type
	return a.resolveType(ctx.Type_())
}

// --- Intrinsics ---

func (a *Analyzer) VisitIntrinsicExpression(ctx *parser.IntrinsicExpressionContext) interface{} {
	if ctx.SIZEOF() != nil || ctx.ALIGNOF() != nil {
		// sizeof<T>, alignof<T> -> usize (u64)
		return types.U64
	}
	
	if ctx.BIT_CAST() != nil {
		a.Visit(ctx.Expression(0))
		return a.resolveType(ctx.Type_())
	}
	
	// Variadics and Memory intrinsics (memset, etc)
	for _, expr := range ctx.AllExpression() {
		a.Visit(expr)
	}
	
	if ctx.VA_START() != nil {
		// Returns va_list (treated as *i8 for now)
		return types.NewPointer(types.I8)
	}
	
	if ctx.VA_ARG() != nil {
		return a.resolveType(ctx.Type_())
	}

	return types.Void
}

// --- Boilerplate ---
func (a *Analyzer) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	// Simplified: just visit primary
	return a.Visit(ctx.PrimaryExpression())
}

func (a *Analyzer) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.Literal() != nil { return a.Visit(ctx.Literal()) }
	if ctx.IDENTIFIER() != nil {
		if s, ok := a.currentScope.Resolve(ctx.IDENTIFIER().GetText()); ok { return s.Type }
	}
	if ctx.CastExpression() != nil { return a.Visit(ctx.CastExpression()) }
	if ctx.IntrinsicExpression() != nil { return a.Visit(ctx.IntrinsicExpression()) }
	if ctx.Expression() != nil { return a.Visit(ctx.Expression()) }
	return types.Void
}

func (a *Analyzer) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTEGER_LITERAL() != nil { return types.I64 }
	if ctx.FLOAT_LITERAL() != nil { return types.F64 }
	if ctx.BOOLEAN_LITERAL() != nil { return types.I1 }
	if ctx.STRING_LITERAL() != nil { return types.NewPointer(types.I8) }
	if ctx.NULL() != nil { return types.NewPointer(types.Void) }
	return types.Void
}

// Passthroughs
func (a *Analyzer) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} { return a.Visit(ctx.LogicalAndExpression(0)) }
func (a *Analyzer) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} { return a.Visit(ctx.BitOrExpression(0)) }
func (a *Analyzer) VisitBitOrExpression(ctx *parser.BitOrExpressionContext) interface{} { return a.Visit(ctx.BitXorExpression(0)) }
func (a *Analyzer) VisitBitXorExpression(ctx *parser.BitXorExpressionContext) interface{} { return a.Visit(ctx.BitAndExpression(0)) }
func (a *Analyzer) VisitBitAndExpression(ctx *parser.BitAndExpressionContext) interface{} { return a.Visit(ctx.EqualityExpression(0)) }
func (a *Analyzer) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} { return a.Visit(ctx.RelationalExpression(0)) }
func (a *Analyzer) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} { return a.Visit(ctx.ShiftExpression(0)) }
func (a *Analyzer) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} { return a.Visit(ctx.RangeExpression(0)) }
func (a *Analyzer) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} { return a.Visit(ctx.AdditiveExpression(0)) }