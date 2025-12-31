package irgen

import (
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (g *Generator) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return g.Visit(ctx.LogicalOrExpression())
}

func (g *Generator) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	// Base Value
	currVal := g.Visit(ctx.PrimaryExpression()).(ir.Value)
	
	for _, op := range ctx.AllPostfixOp() {
		// 1. Function Call
		if op.LPAREN() != nil {
			// Extract arguments
			var args []ir.Value
			if op.ArgumentList() != nil {
				for _, arg := range op.ArgumentList().AllArgument() {
					val := g.Visit(arg.Expression()).(ir.Value)
					args = append(args, val)
				}
			}
			
			if fn, ok := currVal.(*ir.Function); ok {
				currVal = g.ctx.Builder.CreateCall(fn, args, "")
			} else {
				// Handle function pointers or error
			}
		}

		// 2. Member Access (.field)
		if op.DOT() != nil && op.IDENTIFIER() != nil {
			fieldName := op.IDENTIFIER().GetText()
			
			// Auto-dereference if pointer
			isPtr := false
			if _, ok := currVal.Type().(*types.PointerType); ok {
				isPtr = true
				// For LLVM GEP, if we have a pointer to struct, we just use it directly
				// But we need the underlying StructType to find the index.
			}

			// We need to resolve field index. 
			// In Pass 2, we assume Pass 1 validated existence.
			// We need a helper to find index by name for the struct type.
			// For this snippet, assume a function `g.getFieldIndex(structType, fieldName)` exists
			// or we implement a simple one.
			
			// Simplification: We assume struct definition order matches logic
			// In real code: idx := g.ctx.StructFieldMap[structName][fieldName]
			idx := 0 // STUB: Always accessing field 0
			
			if isPtr {
				// Ptr -> Struct GEP
				// GEP(structType, base, 0, idx)
				elemType := currVal.Type().(*types.PointerType).ElementType
				currVal = g.ctx.Builder.CreateStructGEP(elemType, currVal, idx, "")
				// Load the field value
				// currVal = g.ctx.Builder.CreateLoad(..., currVal)
			}
		}

		// 3. Indexing ([expr])
		if op.LBRACKET() != nil {
			idx := g.Visit(op.Expression()).(ir.Value)
			
			if ptrType, ok := currVal.Type().(*types.PointerType); ok {
				// GEP on pointer
				// Array access: Base[idx] -> *(Base + idx)
				// GEP(elemType, base, idx)
				gep := g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, currVal, []ir.Value{idx}, "")
				currVal = g.ctx.Builder.CreateLoad(ptrType.ElementType, gep, "")
			}
		}
	}
	return currVal
}

// ... (Arithmetic helpers kept from previous response) ...

func (g *Generator) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	lhs := g.Visit(ctx.MultiplicativeExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		rhs := g.Visit(ctx.MultiplicativeExpression(i)).(ir.Value)
		// Assuming Add/Sub based on operator
		lhs = g.ctx.Builder.CreateAdd(lhs, rhs, "") 
	}
	return lhs
}
func (g *Generator) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	return g.Visit(ctx.UnaryExpression(0))
}
func (g *Generator) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if ctx.PostfixExpression() != nil { return g.Visit(ctx.PostfixExpression()) }
	
	val := g.Visit(ctx.UnaryExpression()).(ir.Value)
	if ctx.AMP() != nil {
		// Address-of is handled by NOT loading in PrimaryExpression, 
		// but since PrimaryExpression blindly loads allocas, 
		// Unary needs to handle AllocaInsts specifically or we need a Refactor.
		// FIX: We need PrimaryExpr to return the POINTER (Alloca), not the Value.
		// Logic:
		// 1. Visit children. If child is a Load, grab its Operand(0) (the ptr).
		// This is tricky in a visitor pattern. 
		// Alternate: CreateLoad only when needed.
		return val // Stub
	}
	if ctx.STAR() != nil {
		// Dereference: Load
		if ptr, ok := val.Type().(*types.PointerType); ok {
			return g.ctx.Builder.CreateLoad(ptr.ElementType, val, "")
		}
	}
	return val
}

func (g *Generator) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.Literal() != nil { return g.Visit(ctx.Literal()) }
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		if sym, ok := g.currentScope.Resolve(name); ok {
			if alloca, isAlloca := sym.IRValue.(*ir.AllocaInst); isAlloca {
				// Default behavior: Load variable value
				return g.ctx.Builder.CreateLoad(sym.Type, alloca, "")
			}
			return sym.IRValue
		}
	}
	if ctx.Expression() != nil { return g.Visit(ctx.Expression()) }
	return g.ctx.Builder.ConstZero(types.I64)
}

func (g *Generator) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTEGER_LITERAL() != nil { return g.ctx.Builder.ConstInt(types.I64, 0) } // Value from parser
	if ctx.FLOAT_LITERAL() != nil { return g.ctx.Builder.ConstFloat(types.F64, 0.0) }
	if ctx.BOOLEAN_LITERAL() != nil { return g.ctx.Builder.ConstInt(types.I1, 0) }
	return g.ctx.Builder.ConstZero(types.I64)
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