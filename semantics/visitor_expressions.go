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
	if ctx.UnaryExpression() != nil {
		valType := a.Visit(ctx.UnaryExpression()).(types.Type)
		if ctx.AMP() != nil { return types.NewPointer(valType) }
		if ctx.STAR() != nil {
			if ptr, ok := valType.(*types.PointerType); ok { return ptr.ElementType }
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Cannot dereference non-pointer type '%s'", valType.String())
			return types.Void
		}
		return valType
	}
	return types.Void
}

// --- Postfix (Members, Calls, Indexing) ---

func (a *Analyzer) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	curr := a.Visit(ctx.PrimaryExpression()).(types.Type)

	for _, op := range ctx.AllPostfixOp() {
		// 1. Member Access: .field or .method
		if op.DOT() != nil {
			name := op.IDENTIFIER().GetText()
			
			// Auto-dereference pointer to struct
			if ptr, ok := curr.(*types.PointerType); ok {
				curr = ptr.ElementType
			}

			if st, ok := curr.(*types.StructType); ok {
				// Check fields
				if indices, ok := a.structIndices[st.Name]; ok {
					if idx, ok := indices[name]; ok {
						curr = st.Fields[idx]
						continue
					}
				}
				
				// Check methods (Name Mangling: Struct_Method)
				methodName := st.Name + "_" + name
				if sym, ok := a.globalScope.Resolve(methodName); ok {
					// Method found, return function type
					curr = sym.Type
					continue
				}
			}
			// Don't report error if it's "void" (already failed previously)
			if curr != types.Void {
				a.bag.Report(a.file, op.GetStart().GetLine(), 0, "Type '%s' has no field or method '%s'", curr.String(), name)
			}
			return types.Void
		}

		// 2. Indexing: [expr]
		if op.LBRACKET() != nil {
			a.Visit(op.Expression()) // Validate index
			if ptr, ok := curr.(*types.PointerType); ok {
				curr = ptr.ElementType
			} else if arr, ok := curr.(*types.ArrayType); ok {
				curr = arr.ElementType
			} else {
				a.bag.Report(a.file, op.GetStart().GetLine(), 0, "Type '%s' is not indexable", curr.String())
			}
		}

		// 3. Function Call: (args)
		if op.LPAREN() != nil {
			if op.ArgumentList() != nil {
				for _, arg := range op.ArgumentList().AllArgument() {
					a.Visit(arg.Expression())
				}
			}
			if fn, ok := curr.(*types.FunctionType); ok {
				curr = fn.ReturnType
			}
		}
	}
	return curr
}

// --- Primary ---

func (a *Analyzer) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.Literal() != nil { return a.Visit(ctx.Literal()) }
	if ctx.StructLiteral() != nil { return a.Visit(ctx.StructLiteral()) }
	
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		if s, ok := a.currentScope.Resolve(name); ok { return s.Type }
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Undefined identifier '%s'", name)
		return types.Void
	}
	
	// Handle Qualified Identifier (e.g. io.print)
	if ctx.QualifiedIdentifier() != nil {
		q := ctx.QualifiedIdentifier()
		ids := q.AllIDENTIFIER()
		var name string
		for i, id := range ids {
			if i > 0 { name += "." }
			name += id.GetText()
		}
		if s, ok := a.currentScope.Resolve(name); ok { return s.Type }
		// Fallback for externs that might not be fully registered with types yet
		return types.Void 
	}

	if ctx.CastExpression() != nil { return a.Visit(ctx.CastExpression()) }
	if ctx.AllocaExpression() != nil { return a.Visit(ctx.AllocaExpression()) }
	if ctx.IntrinsicExpression() != nil { return a.Visit(ctx.IntrinsicExpression()) }
	if ctx.Expression() != nil { return a.Visit(ctx.Expression()) }
	
	return types.Void
}

func (a *Analyzer) VisitStructLiteral(ctx *parser.StructLiteralContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	if s, ok := a.currentScope.Resolve(name); ok {
		return s.Type
	}
	return types.Void
}

func (a *Analyzer) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTEGER_LITERAL() != nil { return types.I64 }
	if ctx.FLOAT_LITERAL() != nil { return types.F64 }
	if ctx.BOOLEAN_LITERAL() != nil { return types.I1 }
	if ctx.STRING_LITERAL() != nil { return types.NewPointer(types.I8) }
	if ctx.CHAR_LITERAL() != nil { return types.I32 }
	if ctx.NULL() != nil { return types.NewPointer(types.Void) }
	if ctx.InitializerList() != nil { return a.Visit(ctx.InitializerList()) }
	return types.Void
}

func (a *Analyzer) VisitInitializerList(ctx *parser.InitializerListContext) interface{} {
	if len(ctx.AllExpression()) == 0 { return types.Void }
	elemType := a.Visit(ctx.Expression(0)).(types.Type)
	// Validate all elements match
	for i := 1; i < len(ctx.AllExpression()); i++ {
		t := a.Visit(ctx.Expression(i)).(types.Type)
		if !areTypesCompatible(t, elemType) {
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Mixed types in array initializer")
		}
	}
	return types.NewArray(elemType, int64(len(ctx.AllExpression())))
}

// --- Casts & Intrinsics ---

func (a *Analyzer) VisitCastExpression(ctx *parser.CastExpressionContext) interface{} {
	a.Visit(ctx.Expression())
	return a.resolveType(ctx.Type_())
}

func (a *Analyzer) VisitAllocaExpression(ctx *parser.AllocaExpressionContext) interface{} {
	t := a.resolveType(ctx.Type_())
	if ctx.Expression() != nil { a.Visit(ctx.Expression()) }
	return types.NewPointer(t)
}

func (a *Analyzer) VisitIntrinsicExpression(ctx *parser.IntrinsicExpressionContext) interface{} {
	if ctx.SIZEOF() != nil || ctx.ALIGNOF() != nil { return types.U64 }
	if ctx.BIT_CAST() != nil {
		a.Visit(ctx.Expression(0))
		return a.resolveType(ctx.Type_())
	}
	for _, expr := range ctx.AllExpression() { a.Visit(expr) }
	if ctx.VA_START() != nil { return types.NewPointer(types.I8) }
	if ctx.VA_ARG() != nil { return a.resolveType(ctx.Type_()) }
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