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
		
		// Basic type compatibility check
		if !areTypesCompatible(lhs, rhs) {
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, 
				"Type mismatch in addition: %s and %s", lhs.String(), rhs.String())
		}
	}
	return lhs
}

func (a *Analyzer) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	lhs := a.Visit(ctx.UnaryExpression(0)).(types.Type)
	for i := 1; i < len(ctx.AllUnaryExpression()); i++ {
		rhs := a.Visit(ctx.UnaryExpression(i)).(types.Type)
		
		if !areTypesCompatible(lhs, rhs) {
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, 
				"Type mismatch in multiplication: %s and %s", lhs.String(), rhs.String())
		}
	}
	return lhs
}

// --- Unary Ops ---

func (a *Analyzer) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if ctx.PostfixExpression() != nil {
		return a.Visit(ctx.PostfixExpression())
	}
	
	// Handle Recursive Unary Operators
	if ctx.UnaryExpression() != nil {
		valType := a.Visit(ctx.UnaryExpression()).(types.Type)

		// Address-of: &T -> *T
		if ctx.AMP() != nil {
			return types.NewPointer(valType)
		}
		
		// Dereference: *T -> T
		if ctx.STAR() != nil {
			if ptr, ok := valType.(*types.PointerType); ok {
				return ptr.ElementType
			}
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, 
				"Cannot dereference non-pointer type '%s'", valType.String())
			return types.Void
		}
		
		return valType
	}
	
	return types.Void
}

// --- Casts ---

func (a *Analyzer) VisitCastExpression(ctx *parser.CastExpressionContext) interface{} {
	// Visit expression to ensure valid symbol usage / catch errors in the expression
	a.Visit(ctx.Expression())
	
	// Return the explicit target type specified in the cast
	return a.resolveType(ctx.Type_())
}

// --- Allocations ---

func (a *Analyzer) VisitAllocaExpression(ctx *parser.AllocaExpressionContext) interface{} {
	// alloca(Type, [Count]) -> returns *Type
	t := a.resolveType(ctx.Type_())
	
	if ctx.Expression() != nil {
		countType := a.Visit(ctx.Expression()).(types.Type)
		if !types.IsInteger(countType) {
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Alloca count must be an integer")
		}
	}
	
	return types.NewPointer(t)
}

// --- Intrinsics ---

func (a *Analyzer) VisitIntrinsicExpression(ctx *parser.IntrinsicExpressionContext) interface{} {
	// sizeof<T>, alignof<T> -> returns usize (u64)
	if ctx.SIZEOF() != nil || ctx.ALIGNOF() != nil {
		return types.U64
	}
	
	// bit_cast<T>(val) -> returns T
	if ctx.BIT_CAST() != nil {
		a.Visit(ctx.Expression(0))
		return a.resolveType(ctx.Type_())
	}
	
	// Visit arguments for other intrinsics to ensure they are valid
	for _, expr := range ctx.AllExpression() {
		a.Visit(expr)
	}
	
	// va_start -> returns va_list (treated as *i8 for now)
	if ctx.VA_START() != nil {
		return types.NewPointer(types.I8)
	}
	
	// va_arg(list) -> returns T (specified in template)
	if ctx.VA_ARG() != nil {
		return a.resolveType(ctx.Type_())
	}

	// memset, memcpy, etc. return void
	return types.Void
}

// --- Postfix & Primary ---

func (a *Analyzer) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	// TODO: Full member access (.field), indexing ([]), and call () validation.
	// For now, we delegate to Primary to resolve the base type.
	return a.Visit(ctx.PrimaryExpression())
}

func (a *Analyzer) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.Literal() != nil { return a.Visit(ctx.Literal()) }
	
	// Variable / Function lookup
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		if s, ok := a.currentScope.Resolve(name); ok {
			return s.Type
		}
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Undefined identifier '%s'", name)
		return types.Void
	}
	
	// Qualified Identifier (Namespace.Name)
	if ctx.QualifiedIdentifier() != nil {
		// Simplified: assumes valid if parsed, returns void type fallback
		// Real impl would look up in namespace map
		return types.Void 
	}

	if ctx.CastExpression() != nil { return a.Visit(ctx.CastExpression()) }
	if ctx.AllocaExpression() != nil { return a.Visit(ctx.AllocaExpression()) }
	if ctx.IntrinsicExpression() != nil { return a.Visit(ctx.IntrinsicExpression()) }
	if ctx.Expression() != nil { return a.Visit(ctx.Expression()) }
	
	return types.Void
}

func (a *Analyzer) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTEGER_LITERAL() != nil { return types.I64 }
	if ctx.FLOAT_LITERAL() != nil { return types.F64 }
	if ctx.BOOLEAN_LITERAL() != nil { return types.I1 }
	if ctx.STRING_LITERAL() != nil { return types.NewPointer(types.I8) }
	
	// Char Literal Support
	if ctx.CHAR_LITERAL() != nil { return types.I32 }
	
	if ctx.NULL() != nil { return types.NewPointer(types.Void) }
	return types.Void
}

// --- Passthroughs ---
// These ensure the visitor traverses down the grammar hierarchy

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