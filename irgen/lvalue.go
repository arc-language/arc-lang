package irgen

import (
	"strings"

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
			// Explicit dereference: *ptr
			val := g.Visit(ctx.UnaryExpression()).(ir.Value)
			return val
		}
		if ctx.PostfixExpression() != nil {
			return g.getLValue(ctx.PostfixExpression())
		}

	case *parser.PrimaryExpressionContext:
		// Handle parenthesized L-Values: (x) = 1
		if ctx.Expression() != nil && ctx.LPAREN() != nil {
			return g.getLValue(ctx.Expression())
		}
		
		// Function calls (LPAREN but no Expression) are not L-Values
		if ctx.LPAREN() != nil { return nil }
		
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
				// Auto-dereference double pointers
				if ptrType, ok := addr.Type().(*types.PointerType); ok {
					if _, isPtrToPtr := ptrType.ElementType.(*types.PointerType); isPtrToPtr {
						addr = g.ctx.Builder.CreateLoad(ptrType.ElementType, addr, "")
					}
				}

				fieldName := ids[i].GetText()
				ptrType, isPtr := addr.Type().(*types.PointerType)
				if !isPtr { return nil }
				
				if st, ok := ptrType.ElementType.(*types.StructType); ok {
					indices, hasIndices := g.analysis.StructIndices[st.Name]
					if !hasIndices && g.currentNamespace != "" {
						indices, hasIndices = g.analysis.StructIndices[g.currentNamespace + "." + st.Name]
					}
					if !hasIndices && g.currentNamespace != "" {
						prefix := g.currentNamespace + "."
						if len(st.Name) > len(prefix) && strings.HasPrefix(st.Name, prefix) {
							indices, hasIndices = g.analysis.StructIndices[st.Name[len(prefix):]]
						}
					}

					if hasIndices {
						if idx, ok := indices[fieldName]; ok {
							physicalIndex := idx
							if st.IsClass {
								physicalIndex = idx + 1
							}
							addr = g.ctx.Builder.CreateStructGEP(st, addr, physicalIndex, "")
							continue
						}
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
		// Filter out function calls at this level too
		for _, op := range ctx.AllPostfixOp() {
			if op.LPAREN() != nil { return nil }
		}

		baseExpr := ctx.PrimaryExpression()
		addr := g.getLValue(baseExpr)
		
		if addr == nil { return nil }
		
		for _, op := range ctx.AllPostfixOp() {
			// Auto-dereference pointer-to-pointer
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
					indices, hasIndices := g.analysis.StructIndices[st.Name]
					if !hasIndices && g.currentNamespace != "" {
						indices, hasIndices = g.analysis.StructIndices[g.currentNamespace + "." + st.Name]
					}
					if !hasIndices && g.currentNamespace != "" {
						prefix := g.currentNamespace + "."
						if len(st.Name) > len(prefix) && strings.HasPrefix(st.Name, prefix) {
							indices, hasIndices = g.analysis.StructIndices[st.Name[len(prefix):]]
						}
					}

					if hasIndices {
						if idx, ok := indices[fieldName]; ok {
							physicalIndex := idx
							if st.IsClass {
								physicalIndex = idx + 1
							}
							addr = g.ctx.Builder.CreateStructGEP(st, addr, physicalIndex, "")
							continue
						}
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