package irgen

import (
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

// getLValue returns the memory address (pointer) of an expression.
// Returns nil if the expression is not an L-Value (e.g. a literal).
func (g *Generator) getLValue(ctx parser.IExpressionContext) ir.Value {
	// 1. Unary Unwrapping (ignore parens, etc)
	if p, ok := ctx.(*parser.PrimaryExpressionContext); ok {
		if p.Expression() != nil {
			return g.getLValue(p.Expression())
		}
		if p.IDENTIFIER() != nil {
			name := p.IDENTIFIER().GetText()
			if sym, ok := g.currentScope.Resolve(name); ok {
				// If it's an alloca (variable), that IS the address.
				if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
					return alloca
				}
				// If it's a global, that IS the address.
				if glob := g.ctx.Module.GetGlobal(name); glob != nil {
					return glob
				}
				// Function arguments are usually loaded into allocas in VisitFunctionDecl,
				// so resolving the symbol usually returns that alloca.
				return sym.IRValue
			}
		}
	}

	// 2. Dereference (*ptr)
	if u, ok := ctx.(*parser.UnaryExpressionContext); ok {
		if u.STAR() != nil {
			// The value of the sub-expression IS the address we want
			// e.g. *ptr = 5; -> we want the value of 'ptr'
			return g.Visit(u.UnaryExpression()).(ir.Value)
		}
		if u.PostfixExpression() != nil {
			return g.getLValue(u.PostfixExpression())
		}
	}

	// 3. Postfix Operations (Member access, Indexing)
	if post, ok := ctx.(*parser.PostfixExpressionContext); ok {
		// We need to manually walk the postfix chain to generate GEPs 
		// instead of Loads. This repeats some logic from VisitPostfixExpression
		// but stops short of the final Load.
		
		baseExpr := post.PrimaryExpression()
		addr := g.getLValue(baseExpr)
		if addr == nil {
			// If base isn't an LValue (e.g. function return), we can't assign to fields
			// unless we store it in a temp alloca (not implemented here)
			return nil
		}

		for _, op := range post.AllPostfixOp() {
			// .Field
			if op.DOT() != nil {
				fieldName := op.IDENTIFIER().GetText()
				
				// Resolve type of the pointer
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

			// [Index]
			if op.LBRACKET() != nil {
				idxVal := g.Visit(op.Expression()).(ir.Value)
				
				ptrType, isPtr := addr.Type().(*types.PointerType)
				if !isPtr { return nil }

				// If it's a pointer to an Array (e.g. [10 x i32]*), we need 0, i
				if _, isArray := ptrType.ElementType.(*types.ArrayType); isArray {
					zero := g.ctx.Builder.ConstInt(types.I32, 0)
					addr = g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, addr, []ir.Value{zero, idxVal}, "")
				} else {
					// Standard pointer indexing
					addr = g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, addr, []ir.Value{idxVal}, "")
				}
			}
		}
		return addr
	}

	return nil
}