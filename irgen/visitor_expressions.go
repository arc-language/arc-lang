package irgen

import (
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (g *Generator) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	// Bypass wrapper nodes
	return g.Visit(ctx.LogicalOrExpression())
}

// --- Arithmetic ---

func (g *Generator) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	left := g.Visit(ctx.MultiplicativeExpression(0)).(ir.Value)
	
	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		right := g.Visit(ctx.MultiplicativeExpression(i)).(ir.Value)
		
		// Determine operation
		isAdd := i <= len(ctx.AllPLUS())
		
		// Check types (Float vs Int)
		if types.IsFloat(left.Type()) {
			if isAdd {
				left = g.ctx.Builder.CreateFAdd(left, right, "")
			} else {
				left = g.ctx.Builder.CreateFSub(left, right, "")
			}
		} else {
			if isAdd {
				left = g.ctx.Builder.CreateAdd(left, right, "")
			} else {
				left = g.ctx.Builder.CreateSub(left, right, "")
			}
		}
	}
	return left
}

func (g *Generator) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	left := g.Visit(ctx.UnaryExpression(0)).(ir.Value)
	
	for i := 1; i < len(ctx.AllUnaryExpression()); i++ {
		right := g.Visit(ctx.UnaryExpression(i)).(ir.Value)
		
		if i <= len(ctx.AllSTAR()) {
			left = g.ctx.Builder.CreateMul(left, right, "")
		} else {
			// Division (assume signed for now)
			left = g.ctx.Builder.CreateSDiv(left, right, "")
		}
	}
	return left
}

func (g *Generator) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if ctx.PostfixExpression() != nil {
		return g.Visit(ctx.PostfixExpression())
	}
	// Handle -x, !x
	operand := g.Visit(ctx.UnaryExpression()).(ir.Value)
	if ctx.MINUS() != nil {
		return g.ctx.Builder.CreateSub(g.ctx.Builder.ConstZero(operand.Type()), operand, "")
	}
	return operand
}

// --- Terminals ---

func (g *Generator) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	return g.Visit(ctx.PrimaryExpression())
}

func (g *Generator) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	// Literals
	if ctx.Literal() != nil {
		return g.Visit(ctx.Literal())
	}
	
	// Variables
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		
		// Pass 1 guaranteed this exists
		if sym, ok := g.currentScope.Resolve(name); ok {
			// If it's an Alloca (variable), Load it
			if alloca, isAlloca := sym.IRValue.(*ir.AllocaInst); isAlloca {
				return g.ctx.Builder.CreateLoad(sym.Type, alloca, "")
			}
			// If it's a Global or Function, return the pointer
			return sym.IRValue
		}
	}
	
	// Parenthesis
	if ctx.Expression() != nil {
		return g.Visit(ctx.Expression())
	}
	
	return g.ctx.Builder.ConstZero(types.I64)
}

func (g *Generator) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTEGER_LITERAL() != nil {
		// Stub: assumes 64-bit int
		return g.ctx.Builder.ConstInt(types.I64, 0) 
	}
	if ctx.FLOAT_LITERAL() != nil {
		return g.ctx.Builder.ConstFloat(types.F64, 0.0)
	}
	if ctx.BOOLEAN_LITERAL() != nil {
		if ctx.BOOLEAN_LITERAL().GetText() == "true" {
			return g.ctx.Builder.True()
		}
		return g.ctx.Builder.False()
	}
	return g.ctx.Builder.ConstZero(types.I64)
}

// Boilerplate traversals
func (g *Generator) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	return g.Visit(ctx.LogicalAndExpression(0))
}
func (g *Generator) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	return g.Visit(ctx.BitOrExpression(0))
}
func (g *Generator) VisitBitOrExpression(ctx *parser.BitOrExpressionContext) interface{} {
	return g.Visit(ctx.BitXorExpression(0))
}
func (g *Generator) VisitBitXorExpression(ctx *parser.BitXorExpressionContext) interface{} {
	return g.Visit(ctx.BitAndExpression(0))
}
func (g *Generator) VisitBitAndExpression(ctx *parser.BitAndExpressionContext) interface{} {
	return g.Visit(ctx.EqualityExpression(0))
}
func (g *Generator) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	return g.Visit(ctx.RelationalExpression(0))
}
func (g *Generator) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} {
	return g.Visit(ctx.ShiftExpression(0))
}
func (g *Generator) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} {
	return g.Visit(ctx.RangeExpression(0))
}
func (g *Generator) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} {
	return g.Visit(ctx.AdditiveExpression(0))
}