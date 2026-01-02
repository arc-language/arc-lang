package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

// --- Expressions Entry Point ---

func (a *Analyzer) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	if ctx.LogicalOrExpression() != nil {
		return a.Visit(ctx.LogicalOrExpression())
	}
	return types.Void
}

// --- Binary Operations ---

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

// --- Unary Operations ---

func (a *Analyzer) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	// Fallthrough to Postfix
	if ctx.PostfixExpression() != nil {
		return a.Visit(ctx.PostfixExpression())
	}
	
	// Handle Prefix Ops (- ! ~ * & ++ --)
	if ctx.UnaryExpression() != nil {
		valType := a.Visit(ctx.UnaryExpression()).(types.Type)
		
		// Address Of (&)
		if ctx.AMP() != nil {
			return types.NewPointer(valType)
		}
		
		// Dereference (*)
		if ctx.STAR() != nil {
			if ptr, ok := valType.(*types.PointerType); ok {
				return ptr.ElementType
			}
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, 
				"Cannot dereference non-pointer type '%s'", valType.String())
			return types.Void
		}

		// Math/Logic Ops
		if ctx.MINUS() != nil || ctx.NOT() != nil || ctx.BIT_NOT() != nil {
			return valType
		}

		// Prefix Increment/Decrement (++i, --i)
		if ctx.INCREMENT() != nil || ctx.DECREMENT() != nil {
			if !types.IsInteger(valType) && !types.IsPointer(valType) {
				a.bag.Report(a.file, ctx.GetStart().GetLine(), 0,
					"Cannot increment/decrement type '%s'", valType.String())
			}
			return valType
		}

		return valType
	}
	return types.Void
}

// --- Postfix Expressions (Calls, Members, Indexing) ---

func (a *Analyzer) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	// 1. Resolve Base (Primary Expression)
	curr := a.Visit(ctx.PrimaryExpression()).(types.Type)

	// 2. Iterate Ops (Left-to-Right)
	for _, op := range ctx.AllPostfixOp() {
		
		// --- Function Call: foo(args) ---
		if op.LPAREN() != nil {
			// Validate Arguments
			if op.ArgumentList() != nil {
				for _, arg := range op.ArgumentList().AllArgument() {
					a.Visit(arg.Expression())
				}
			}
			
			// Resolve Return Type
			if fn, ok := curr.(*types.FunctionType); ok {
				curr = fn.ReturnType
			} else {
				// Special case: 'cast<T>' returns a function type constructed in VisitPrimaryExpression
				a.bag.Report(a.file, op.GetStart().GetLine(), 0, 
					"Cannot call non-function type '%s'", curr.String())
				curr = types.Void
			}
		}

		// --- Member Access: obj.field or obj.method ---
		if op.DOT() != nil {
			name := op.IDENTIFIER().GetText()
			
			// Auto-dereference pointers (obj->field in C, obj.field in Arc)
			if ptr, ok := curr.(*types.PointerType); ok {
				curr = ptr.ElementType
			}

			if st, ok := curr.(*types.StructType); ok {
				// 1. Check Fields
				if indices, ok := a.structIndices[st.Name]; ok {
					if idx, ok := indices[name]; ok {
						curr = st.Fields[idx]
						continue
					}
				}
				
				// 2. Check Methods (Name Mangling: StructName_MethodName)
				methodName := st.Name + "_" + name
				if sym, ok := a.globalScope.Resolve(methodName); ok {
					curr = sym.Type
					continue
				}
			}

			// Report Error (unless we already errored earlier)
			if curr != types.Void {
				a.bag.Report(a.file, op.GetStart().GetLine(), 0, 
					"Type '%s' has no field or method '%s'", curr.String(), name)
			}
			return types.Void
		}

		// --- Indexing: arr[i] ---
		if op.LBRACKET() != nil {
			// Validate Index Expression
			idxType := a.Visit(op.Expression()).(types.Type)
			if !types.IsInteger(idxType) {
				a.bag.Report(a.file, op.GetStart().GetLine(), 0, 
					"Index must be an integer, got '%s'", idxType.String())
			}

			// Dereference Collection
			if ptr, ok := curr.(*types.PointerType); ok {
				curr = ptr.ElementType
				// If it was a pointer to array (*[N]T), dereference to element T
				if arr, ok := curr.(*types.ArrayType); ok {
					curr = arr.ElementType
				}
			} else if arr, ok := curr.(*types.ArrayType); ok {
				curr = arr.ElementType
			} else {
				a.bag.Report(a.file, op.GetStart().GetLine(), 0, 
					"Type '%s' is not indexable", curr.String())
			}
		}

		// --- Postfix Increment/Decrement: i++ ---
		if op.INCREMENT() != nil || op.DECREMENT() != nil {
			if !types.IsInteger(curr) && !types.IsPointer(curr) {
				a.bag.Report(a.file, op.GetStart().GetLine(), 0,
					"Cannot increment/decrement type '%s'", curr.String())
			}
		}
	}
	return curr
}

// --- Primary Expressions ---

func (a *Analyzer) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	// 1. Literals
	if ctx.Literal() != nil { return a.Visit(ctx.Literal()) }
	if ctx.StructLiteral() != nil { return a.Visit(ctx.StructLiteral()) }
	
	// 2. Compiler Keywords (Sizeof / Alignof)
	if ctx.SizeofExpression() != nil {
		a.resolveType(ctx.SizeofExpression().Type_()) // Check type validity
		return types.U64
	}
	if ctx.AlignofExpression() != nil {
		a.resolveType(ctx.AlignofExpression().Type_()) // Check type validity
		return types.U64
	}

	// 3. Identifiers (Variables, Functions, Intrinsics)
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()

		// Special handling for 'cast<T>' or 'bit_cast<T>'
		// These behave like functions that return T.
		if (name == "cast" || name == "bit_cast") && ctx.GenericArgs() != nil {
			// Extract target type T from generic args
			gArgs := ctx.GenericArgs().GenericArgList()
			if gArgs != nil && len(gArgs.AllGenericArg()) > 0 {
				argCtx := gArgs.GenericArg(0)
				// resolveType expects ITypeContext, usually in GenericArg
				if argCtx.Type_() != nil {
					targetType := a.resolveType(argCtx.Type_())
					// Return a synthetic function type: func(...) -> targetType
					// Params are checked loosely here or could be strict
					return types.NewFunction(targetType, nil, true)
				}
			}
		}
		
		// Look up in scope
		if s, ok := a.currentScope.Resolve(name); ok {
			return s.Type
		}
		
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Undefined identifier '%s'", name)
		return types.Void
	}
	
	// 4. Qualified Identifiers (e.g. io.print)
	if ctx.QualifiedIdentifier() != nil {
		q := ctx.QualifiedIdentifier()
		var name string
		for i, id := range q.AllIDENTIFIER() {
			if i > 0 { name += "." }
			name += id.GetText()
		}
		
		if s, ok := a.currentScope.Resolve(name); ok {
			return s.Type
		}
		
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Undefined qualified symbol '%s'", name)
		return types.Void 
	}

	// 5. Parenthesized Expression
	if ctx.Expression() != nil {
		return a.Visit(ctx.Expression())
	}
	
	return types.Void
}

// --- Literals & Structures ---

func (a *Analyzer) VisitStructLiteral(ctx *parser.StructLiteralContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	// Resolve Struct Type
	sym, ok := a.currentScope.Resolve(name)
	if !ok || sym.Kind != symbol.SymType {
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Unknown struct type '%s'", name)
		return types.Void
	}

	st, ok := sym.Type.(*types.StructType)
	if !ok {
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "'%s' is not a struct", name)
		return types.Void
	}

	indices, hasIndices := a.structIndices[name]
	if !hasIndices {
		return sym.Type
	}

	for _, field := range ctx.AllFieldInit() {
		fName := field.IDENTIFIER().GetText()
		idx, exists := indices[fName]
		if !exists {
			a.bag.Report(a.file, field.GetStart().GetLine(), 0, 
				"Struct '%s' has no field '%s'", name, fName)
			continue
		}

		exprType := a.Visit(field.Expression()).(types.Type)
		expectedType := st.Fields[idx]

		if !areTypesCompatible(exprType, expectedType) {
			a.bag.Report(a.file, field.GetStart().GetLine(), 0,
				"Field '%s' expects type %s, got %s", fName, expectedType.String(), exprType.String())
		}
	}

	return sym.Type
}

func (a *Analyzer) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTEGER_LITERAL() != nil { return types.I64 }
	if ctx.FLOAT_LITERAL() != nil { return types.F64 }
	if ctx.BOOLEAN_LITERAL() != nil { return types.I1 }
	if ctx.STRING_LITERAL() != nil { return types.NewPointer(types.I8) } // C-String
	if ctx.CHAR_LITERAL() != nil { return types.I32 } // Rune
	if ctx.NULL() != nil { return types.NewPointer(types.Void) }
	
	if ctx.InitializerList() != nil { 
		return a.Visit(ctx.InitializerList()) 
	}
	
	return types.Void
}

func (a *Analyzer) VisitInitializerList(ctx *parser.InitializerListContext) interface{} {
	if len(ctx.AllExpression()) == 0 { 
		return types.Void 
	}
	
	elemType := a.Visit(ctx.Expression(0)).(types.Type)
	
	for i := 1; i < len(ctx.AllExpression()); i++ {
		t := a.Visit(ctx.Expression(i)).(types.Type)
		if !areTypesCompatible(t, elemType) {
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, 
				"Mixed types in initializer list: %s vs %s", elemType.String(), t.String())
		}
	}
	
	return types.NewArray(elemType, int64(len(ctx.AllExpression())))
}

// --- Passthroughs for Precedence Rules ---
func (a *Analyzer) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} { return a.Visit(ctx.LogicalAndExpression(0)) }
func (a *Analyzer) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} { return a.Visit(ctx.BitOrExpression(0)) }
func (a *Analyzer) VisitBitOrExpression(ctx *parser.BitOrExpressionContext) interface{} { return a.Visit(ctx.BitXorExpression(0)) }
func (a *Analyzer) VisitBitXorExpression(ctx *parser.BitXorExpressionContext) interface{} { return a.Visit(ctx.BitAndExpression(0)) }
func (a *Analyzer) VisitBitAndExpression(ctx *parser.BitAndExpressionContext) interface{} { return a.Visit(ctx.EqualityExpression(0)) }
func (a *Analyzer) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} { return a.Visit(ctx.RelationalExpression(0)) }
func (a *Analyzer) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} { return a.Visit(ctx.ShiftExpression(0)) }
func (a *Analyzer) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} { return a.Visit(ctx.RangeExpression(0)) }
func (a *Analyzer) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} { return a.Visit(ctx.AdditiveExpression(0)) }