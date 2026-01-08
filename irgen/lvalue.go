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
	case *parser.UnaryExpressionContext:
		if ctx.STAR() != nil {
			// Explicit dereference: *ptr -> The value of the ptr expression IS the L-Value address
			val := g.Visit(ctx.UnaryExpression()).(ir.Value)
			return val
		}
		if ctx.PostfixExpression() != nil {
			return g.getLValue(ctx.PostfixExpression())
		}

	// --- LeftHandSide ---
	case *parser.LeftHandSideContext:
		// IDENTIFIER
		if ctx.IDENTIFIER() != nil && ctx.DOT() == nil && ctx.STAR() == nil && ctx.LBRACKET() == nil {
			name := ctx.IDENTIFIER().GetText()
			
			// Try resolve
			sym, ok := g.currentScope.Resolve(name)
			if !ok && g.currentNamespace != "" {
				sym, ok = g.currentScope.Resolve(g.currentNamespace + "." + name)
			}

			if ok {
				if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
					return alloca
				}
				if glob := g.ctx.Module.GetGlobal(name); glob != nil {
					return glob
				}
				if fn, ok := sym.IRValue.(*ir.Function); ok {
					return fn
				}
				return sym.IRValue
			}
			return nil
		}

		// STAR postfixExpression
		if ctx.STAR() != nil {
			val := g.Visit(ctx.PostfixExpression()).(ir.Value)
			return val
		}

		// postfixExpression DOT IDENTIFIER
		if ctx.DOT() != nil {
			base := g.getLValue(ctx.PostfixExpression())
			if base == nil {
				val := g.Visit(ctx.PostfixExpression()).(ir.Value)
				if types.IsPointer(val.Type()) {
					base = val
				}
			}
			
			if base != nil {
				if ptrType, ok := base.Type().(*types.PointerType); ok {
					if _, isPtrToPtr := ptrType.ElementType.(*types.PointerType); isPtrToPtr {
						base = g.ctx.Builder.CreateLoad(ptrType.ElementType, base, "")
					}
				}

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
			base := g.getLValue(ctx.PostfixExpression())
			if base == nil {
				val := g.Visit(ctx.PostfixExpression()).(ir.Value)
				if types.IsPointer(val.Type()) {
					base = val
				}
			} else {
				ptrType := base.Type().(*types.PointerType)
				if _, ok := ptrType.ElementType.(*types.PointerType); ok {
					base = g.ctx.Builder.CreateLoad(ptrType.ElementType, base, "")
				}
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
		if ctx.LPAREN() != nil { return nil }
		if ctx.Expression() != nil { return g.getLValue(ctx.Expression()) }
		
		// Qualified Identifier
		if ctx.QualifiedIdentifier() != nil {
			qCtx := ctx.QualifiedIdentifier().(*parser.QualifiedIdentifierContext)
			ids := qCtx.AllIDENTIFIER()
			baseName := ids[0].GetText()
			
			var addr ir.Value
			
			sym, ok := g.currentScope.Resolve(baseName)
			if !ok && g.currentNamespace != "" {
				sym, ok = g.currentScope.Resolve(g.currentNamespace + "." + baseName)
			}

			if ok {
				if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
					addr = alloca
				} else if glob := g.ctx.Module.GetGlobal(baseName); glob != nil {
					addr = glob
				} else {
					addr = sym.IRValue
				}
			}
			
			if addr == nil { return nil }

			for i := 1; i < len(ids); i++ {
				if ptrType, ok := addr.Type().(*types.PointerType); ok {
					if _, isPtrToPtr := ptrType.ElementType.(*types.PointerType); isPtrToPtr {
						addr = g.ctx.Builder.CreateLoad(ptrType.ElementType, addr, "")
					}
				}

				fieldName := ids[i].GetText()
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
			return addr
		}

		// Single Identifier
		if ctx.IDENTIFIER() != nil {
			name := ctx.IDENTIFIER().GetText()
			
			sym, ok := g.currentScope.Resolve(name)
			if !ok && g.currentNamespace != "" {
				sym, ok = g.currentScope.Resolve(g.currentNamespace + "." + name)
			}

			if ok {
				if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
					return alloca
				}
				if glob := g.ctx.Module.GetGlobal(name); glob != nil {
					return glob
				}
				if fn, ok := sym.IRValue.(*ir.Function); ok {
					return fn
				}
				return sym.IRValue
			}
		}

	case *parser.PostfixExpressionContext:
		for _, op := range ctx.AllPostfixOp() {
			if op.LPAREN() != nil { return nil }
		}

		baseExpr := ctx.PrimaryExpression()
		addr := g.getLValue(baseExpr)
		
		if addr == nil { return nil }
		
		for _, op := range ctx.AllPostfixOp() {
			if ptrType, ok := addr.Type().(*types.PointerType); ok {
				if _, isPtrToPtr := ptrType.ElementType.(*types.PointerType); isPtrToPtr {
					addr = g.ctx.Builder.CreateLoad(ptrType.ElementType, addr, "")
				}
			}

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