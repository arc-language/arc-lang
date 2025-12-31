package irgen

import (
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (g *Generator) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return g.Visit(ctx.LogicalOrExpression())
}

// --- Arithmetic ---
func (g *Generator) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	lhs := g.Visit(ctx.MultiplicativeExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		rhs := g.Visit(ctx.MultiplicativeExpression(i)).(ir.Value)
		
		isAdd := i <= len(ctx.AllPLUS())
		if types.IsFloat(lhs.Type()) {
			if isAdd { lhs = g.ctx.Builder.CreateFAdd(lhs, rhs, "") }
			else { lhs = g.ctx.Builder.CreateFSub(lhs, rhs, "") }
		} else {
			if isAdd { lhs = g.ctx.Builder.CreateAdd(lhs, rhs, "") }
			else { lhs = g.ctx.Builder.CreateSub(lhs, rhs, "") }
		}
	}
	return lhs
}

func (g *Generator) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	lhs := g.Visit(ctx.UnaryExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllUnaryExpression()); i++ {
		rhs := g.Visit(ctx.UnaryExpression(i)).(ir.Value)
		
		isMult := i <= len(ctx.AllSTAR())
		if types.IsFloat(lhs.Type()) {
			if isMult { lhs = g.ctx.Builder.CreateFMul(lhs, rhs, "") }
			else { lhs = g.ctx.Builder.CreateFDiv(lhs, rhs, "") }
		} else {
			if isMult { lhs = g.ctx.Builder.CreateMul(lhs, rhs, "") }
			else { lhs = g.ctx.Builder.CreateSDiv(lhs, rhs, "") }
		}
	}
	return lhs
}

func (g *Generator) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if ctx.PostfixExpression() != nil { return g.Visit(ctx.PostfixExpression()) }
	
	val := g.Visit(ctx.UnaryExpression()).(ir.Value)
	if ctx.MINUS() != nil {
		if types.IsFloat(val.Type()) { return g.ctx.Builder.CreateFSub(g.getZeroValue(val.Type()), val, "") }
		return g.ctx.Builder.CreateSub(g.getZeroValue(val.Type()), val, "")
	}
	if ctx.STAR() != nil {
		if ptr, ok := val.Type().(*types.PointerType); ok {
			return g.ctx.Builder.CreateLoad(ptr.ElementType, val, "")
		}
	}
	return val
}

// --- Postfix (Calls, Members) ---
func (g *Generator) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	curr := g.Visit(ctx.PrimaryExpression()).(ir.Value)

	for _, op := range ctx.AllPostfixOp() {
		// 1. Calls
		if op.LPAREN() != nil {
			var args []ir.Value
			if op.ArgumentList() != nil {
				for _, arg := range op.ArgumentList().AllArgument() {
					args = append(args, g.Visit(arg.Expression()).(ir.Value))
				}
			}
			if fn, ok := curr.(*ir.Function); ok {
				curr = g.ctx.Builder.CreateCall(fn, args, "")
			}
		}

		// 2. Member Access (.field)
		if op.DOT() != nil {
			fieldName := op.IDENTIFIER().GetText()
			
			// Auto deref pointer
			isPtr := false
			if _, ok := curr.Type().(*types.PointerType); ok {
				isPtr = true
			}
			
			// Resolve underlying Struct Type
			var structType *types.StructType
			if isPtr {
				structType = curr.Type().(*types.PointerType).ElementType.(*types.StructType)
			} else {
				structType = curr.Type().(*types.StructType)
			}

			// Find Field Index using the Map from Pass 1
			idx := -1
			if indices, ok := g.analysis.StructIndices[structType.Name]; ok {
				if i, ok := indices[fieldName]; ok {
					idx = i
				}
			}
			
			if idx >= 0 {
				if isPtr {
					// GEP -> Load
					gep := g.ctx.Builder.CreateStructGEP(structType, curr, idx, "")
					curr = g.ctx.Builder.CreateLoad(structType.Fields[idx], gep, "")
				} else {
					// ExtractValue (if value type)
					curr = g.ctx.Builder.CreateExtractValue(curr, []int{idx}, "")
				}
			}
		}
		
		// 3. Indexing ([expr])
		if op.LBRACKET() != nil {
			idx := g.Visit(op.Expression()).(ir.Value)
			if ptr, ok := curr.Type().(*types.PointerType); ok {
				// GEP
				gep := g.ctx.Builder.CreateInBoundsGEP(ptr.ElementType, curr, []ir.Value{idx}, "")
				curr = g.ctx.Builder.CreateLoad(ptr.ElementType, gep, "")
			}
		}
	}
	return curr
}

func (g *Generator) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.Literal() != nil { return g.Visit(ctx.Literal()) }
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		if sym, ok := g.currentScope.Resolve(name); ok {
			// If it's a variable (Alloca), load it
			if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
				return g.ctx.Builder.CreateLoad(sym.Type, alloca, "")
			}
			return sym.IRValue
		}
	}
	if ctx.Expression() != nil { return g.Visit(ctx.Expression()) }
	return g.getZeroValue(types.I64)
}

func (g *Generator) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTEGER_LITERAL() != nil { return g.ctx.Builder.ConstInt(types.I64, 0) } 
	if ctx.FLOAT_LITERAL() != nil { return g.ctx.Builder.ConstFloat(types.F64, 0.0) }
	if ctx.BOOLEAN_LITERAL() != nil { return g.ctx.Builder.ConstInt(types.I1, 0) }
	return g.getZeroValue(types.I64)
}

// Boilerplate
func (g *Generator) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} { return g.Visit(ctx.LogicalAndExpression(0)) }
func (g *Generator) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} { return g.Visit(ctx.BitOrExpression(0)) }
func (g *Generator) VisitBitOrExpression(ctx *parser.BitOrExpressionContext) interface{} { return g.Visit(ctx.BitXorExpression(0)) }
func (g *Generator) VisitBitXorExpression(ctx *parser.BitXorExpressionContext) interface{} { return g.Visit(ctx.BitAndExpression(0)) }
func (g *Generator) VisitBitAndExpression(ctx *parser.BitAndExpressionContext) interface{} { return g.Visit(ctx.EqualityExpression(0)) }
func (g *Generator) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} { return g.Visit(ctx.RelationalExpression(0)) }
func (g *Generator) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} { return g.Visit(ctx.ShiftExpression(0)) }
func (g *Generator) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} { return g.Visit(ctx.RangeExpression(0)) }
func (g *Generator) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} { return g.Visit(ctx.AdditiveExpression(0)) }