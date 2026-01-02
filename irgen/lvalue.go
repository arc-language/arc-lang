package irgen

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

// getLValue returns the memory address (pointer) of an expression.
// Returns nil if the expression is not an L-Value (e.g. a literal).
func (g *Generator) getLValue(tree antlr.ParseTree) ir.Value {
	switch ctx := tree.(type) {
	
	// --- Drill-down layers ---
	case *parser.ExpressionContext:
		return g.getLValue(ctx.LogicalOrExpression())
	case *parser.LogicalOrExpressionContext:
		if len(ctx.AllLogicalAndExpression()) == 1 {
			return g.getLValue(ctx.LogicalAndExpression(0))
		}
	case *parser.LogicalAndExpressionContext:
		if len(ctx.AllBitOrExpression()) == 1 {
			return g.getLValue(ctx.BitOrExpression(0))
		}
	case *parser.BitOrExpressionContext:
		if len(ctx.AllBitXorExpression()) == 1 {
			return g.getLValue(ctx.BitXorExpression(0))
		}
	case *parser.BitXorExpressionContext:
		if len(ctx.AllBitAndExpression()) == 1 {
			return g.getLValue(ctx.BitAndExpression(0))
		}
	case *parser.BitAndExpressionContext:
		if len(ctx.AllEqualityExpression()) == 1 {
			return g.getLValue(ctx.EqualityExpression(0))
		}
	case *parser.EqualityExpressionContext:
		if len(ctx.AllRelationalExpression()) == 1 {
			return g.getLValue(ctx.RelationalExpression(0))
		}
	case *parser.RelationalExpressionContext:
		if len(ctx.AllShiftExpression()) == 1 {
			return g.getLValue(ctx.ShiftExpression(0))
		}
	case *parser.ShiftExpressionContext:
		if len(ctx.AllRangeExpression()) == 1 {
			return g.getLValue(ctx.RangeExpression(0))
		}
	case *parser.RangeExpressionContext:
		if len(ctx.AllAdditiveExpression()) == 1 {
			return g.getLValue(ctx.AdditiveExpression(0))
		}
	case *parser.AdditiveExpressionContext:
		if len(ctx.AllMultiplicativeExpression()) == 1 {
			return g.getLValue(ctx.MultiplicativeExpression(0))
		}
	case *parser.MultiplicativeExpressionContext:
		if len(ctx.AllUnaryExpression()) == 1 {
			return g.getLValue(ctx.UnaryExpression(0))
		}

	// --- LeftHandSide ---
	case *parser.LeftHandSideContext:
		// IDENTIFIER
		if ctx.IDENTIFIER() != nil && ctx.DOT() == nil && ctx.STAR() == nil && ctx.LBRACKET() == nil {
			name := ctx.IDENTIFIER().GetText()
			if sym, ok := g.currentScope.Resolve(name); ok {
				if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
					return alloca
				}
				if glob := g.ctx.Module.GetGlobal(name); glob != nil {
					return glob
				}
				return sym.IRValue
			}
			return nil
		}

		// STAR postfixExpression
		if ctx.STAR() != nil {
			// Explicit dereference: *ptr
			val := g.Visit(ctx.PostfixExpression()).(ir.Value)
			// The value of 'ptr' IS the address we want to write to
			return val
		}

		// postfixExpression DOT IDENTIFIER
		if ctx.DOT() != nil {
			base := g.getLValue(ctx.PostfixExpression())
			if base == nil {
				// If base is not an l-value (e.g. function call return), we can't assign to field?
				// Actually, we might be able to if we loaded it? No, assignments require address.
				// Try visiting to get the pointer value if it's a pointer type?
				val := g.Visit(ctx.PostfixExpression()).(ir.Value)
				if types.IsPointer(val.Type()) {
					base = val
				}
			}
			
			if base != nil {
				fieldName := ctx.IDENTIFIER().GetText()
				ptrType := base.Type().(*types.PointerType)
				if st, ok := ptrType.ElementType.(*types.StructType); ok {
					if idx, ok := g.analysis.StructIndices[st.Name][fieldName]; ok {
						return g.ctx.Builder.CreateStructGEP(st, base, idx, "")
					}
				}
			}
			return nil
		}

		// postfixExpression LBRACKET expression RBRACKET
		if ctx.LBRACKET() != nil {
			idxVal := g.Visit(ctx.Expression()).(ir.Value)
			
			// Get the base address.
			// 1. Try getLValue (if it's an array variable)
			base := g.getLValue(ctx.PostfixExpression())
			
			// 2. If nil, try Visit (if it's a pointer value, e.g. from function call or pointer variable)
			if base == nil {
				val := g.Visit(ctx.PostfixExpression()).(ir.Value)
				if types.IsPointer(val.Type()) {
					// It's a pointer value, so that IS the base address
					base = val
				}
			} else {
				// If we got an l-value, we need to load it if it's a pointer variable (double pointer)
				// OR if it's an array, decay it.
				// `getLValue` returns the address of the variable.
				// If `arr` is `*i32` (variable), `getLValue` returns `**i32`.
				// If `arr` is `[5 x i32]`, `getLValue` returns `*[5 x i32]`.
				
				ptrType := base.Type().(*types.PointerType)
				
				// Case A: Variable holding a pointer (*int32)
				if ptrToPtr, ok := ptrType.ElementType.(*types.PointerType); ok {
					_ = ptrToPtr
					// Load the actual pointer value
					base = g.ctx.Builder.CreateLoad(ptrType.ElementType, base, "")
				}
				// Case B: Array (don't load, just GEP)
			}

			if base != nil {
				ptrType := base.Type().(*types.PointerType)
				
				if _, isArray := ptrType.ElementType.(*types.ArrayType); isArray {
					zero := g.ctx.Builder.ConstInt(types.I32, 0)
					return g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, base, []ir.Value{zero, idxVal}, "")
				} else {
					return g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, base, []ir.Value{idxVal}, "")
				}
			}
			return nil
		}

	// --- Actual L-Value Logic ---

	case *parser.PrimaryExpressionContext:
		// Function calls are NOT l-values. 
		if ctx.LPAREN() != nil {
			return nil
		}

		if ctx.Expression() != nil {
			return g.getLValue(ctx.Expression())
		}
		if ctx.IDENTIFIER() != nil {
			name := ctx.IDENTIFIER().GetText()
			if sym, ok := g.currentScope.Resolve(name); ok {
				if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
					return alloca
				}
				if glob := g.ctx.Module.GetGlobal(name); glob != nil {
					return glob
				}
				return sym.IRValue
			}
		}

	case *parser.UnaryExpressionContext:
		if ctx.STAR() != nil {
			return g.Visit(ctx.UnaryExpression()).(ir.Value)
		}
		if ctx.PostfixExpression() != nil {
			return g.getLValue(ctx.PostfixExpression())
		}

	case *parser.PostfixExpressionContext:
		// Postfix calls (foo()) are not l-values.
		for _, op := range ctx.AllPostfixOp() {
			if op.LPAREN() != nil {
				return nil
			}
		}

		baseExpr := ctx.PrimaryExpression()
		addr := g.getLValue(baseExpr)
		
		if addr == nil { return nil }

		for _, op := range ctx.AllPostfixOp() {
			if op.DOT() != nil {
				fieldName := op.IDENTIFIER().GetText()
				ptrType, isPtr := addr.Type().(*types.PointerType)
				if !isPtr { return nil }

				if st, ok := ptrType.ElementType.(*types.StructType); ok {
					if idx, ok := g.analysis.StructIndices[st.Name][fieldName]; ok {
						addr = g.ctx.Builder.CreateStructGEP(st, addr, idx, "")
						continue
					}
				}
				return nil
			}

			if op.LBRACKET() != nil {
				idxVal := g.Visit(op.Expression()).(ir.Value)
				ptrType, isPtr := addr.Type().(*types.PointerType)
				if !isPtr { return nil }

				if _, isArray := ptrType.ElementType.(*types.ArrayType); isArray {
					zero := g.ctx.Builder.ConstInt(types.I32, 0)
					addr = g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, addr, []ir.Value{zero, idxVal}, "")
				} else {
					addr = g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, addr, []ir.Value{idxVal}, "")
				}
			}
		}
		return addr
	}

	return nil
}