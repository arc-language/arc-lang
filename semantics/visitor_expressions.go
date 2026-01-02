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
		
		// Handle Pointer Arithmetic (ptr + int, ptr - int)
		if types.IsPointer(lhs) && types.IsInteger(rhs) {
			continue
		}
		
		// Handle Pointer Arithmetic (int + ptr)
		if types.IsInteger(lhs) && types.IsPointer(rhs) {
			lhs = rhs 
			continue
		}

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
	
	// Handle Prefix Ops
	if ctx.UnaryExpression() != nil {
		valType := a.Visit(ctx.UnaryExpression()).(types.Type)
		
		if ctx.AMP() != nil {
			return types.NewPointer(valType)
		}
		
		if ctx.STAR() != nil {
			if ptr, ok := valType.(*types.PointerType); ok {
				return ptr.ElementType
			}
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, 
				"Cannot dereference non-pointer type '%s'", valType.String())
			return types.Void
		}

		if ctx.MINUS() != nil || ctx.NOT() != nil || ctx.BIT_NOT() != nil {
			return valType
		}

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

// --- Postfix Expressions ---

func (a *Analyzer) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	curr := a.Visit(ctx.PrimaryExpression()).(types.Type)

	for _, op := range ctx.AllPostfixOp() {
		
		// Function Call via Postfix Op
		if op.LPAREN() != nil {
			if op.ArgumentList() != nil {
				for _, arg := range op.ArgumentList().AllArgument() {
					a.Visit(arg.Expression())
				}
			}
			
			if fn, ok := curr.(*types.FunctionType); ok {
				curr = fn.ReturnType
			} else {
				a.bag.Report(a.file, op.GetStart().GetLine(), 0, 
					"Cannot call non-function type '%s'", curr.String())
				curr = types.Void
			}
		}

		// Member Access
		if op.DOT() != nil {
			name := op.IDENTIFIER().GetText()
			
			if ptr, ok := curr.(*types.PointerType); ok {
				curr = ptr.ElementType
			}

			if st, ok := curr.(*types.StructType); ok {
				if indices, ok := a.structIndices[st.Name]; ok {
					if idx, ok := indices[name]; ok {
						curr = st.Fields[idx]
						continue
					}
				}
				
				methodName := st.Name + "_" + name
				if sym, ok := a.globalScope.Resolve(methodName); ok {
					curr = sym.Type
					continue
				}
			}

			if curr != types.Void {
				a.bag.Report(a.file, op.GetStart().GetLine(), 0, 
					"Type '%s' has no field or method '%s'", curr.String(), name)
			}
			return types.Void
		}

		// Indexing
		if op.LBRACKET() != nil {
			idxType := a.Visit(op.Expression()).(types.Type)
			if !types.IsInteger(idxType) {
				a.bag.Report(a.file, op.GetStart().GetLine(), 0, 
					"Index must be an integer, got '%s'", idxType.String())
			}

			if ptr, ok := curr.(*types.PointerType); ok {
				curr = ptr.ElementType
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

		// Increment/Decrement
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
	
	// 2. Keywords
	if ctx.SizeofExpression() != nil {
		a.resolveType(ctx.SizeofExpression().Type_())
		return types.U64
	}
	if ctx.AlignofExpression() != nil {
		a.resolveType(ctx.AlignofExpression().Type_())
		return types.U64
	}

	// 3. Identifiers / Qualified
	var name string
	var isQualified bool
	var hasParens bool = (ctx.LPAREN() != nil)
	var qCtx *parser.QualifiedIdentifierContext

	if ctx.QualifiedIdentifier() != nil {
		qCtx = ctx.QualifiedIdentifier().(*parser.QualifiedIdentifierContext)
		for i, id := range qCtx.AllIDENTIFIER() {
			if i > 0 { name += "." }
			name += id.GetText()
		}
		isQualified = true
	} else if ctx.IDENTIFIER() != nil {
		name = ctx.IDENTIFIER().GetText()
	}

	if name != "" {
		// Handle Special Forms (cast, alloca, etc.)
		if ctx.GenericArgs() != nil && !isQualified {
			if name == "cast" || name == "bit_cast" {
				gArgs := ctx.GenericArgs().GenericArgList()
				if gArgs != nil && len(gArgs.AllGenericArg()) > 0 {
					argCtx := gArgs.GenericArg(0)
					if argCtx.Type_() != nil {
						targetType := a.resolveType(argCtx.Type_())
						if hasParens {
							return targetType
						}
						return types.NewFunction(targetType, nil, true)
					}
				}
			}
			
			if name == "alloca" {
				gArgs := ctx.GenericArgs().GenericArgList()
				if gArgs != nil && len(gArgs.AllGenericArg()) > 0 {
					argCtx := gArgs.GenericArg(0)
					if argCtx.Type_() != nil {
						targetType := a.resolveType(argCtx.Type_())
						ptrType := types.NewPointer(targetType)
						if hasParens {
							return ptrType
						}
						return types.NewFunction(ptrType, nil, true)
					}
				}
			}
		}
		
		// Normal Symbol Resolution
		if s, ok := a.currentScope.Resolve(name); ok {
			typ := s.Type
			if hasParens {
				if fn, ok := typ.(*types.FunctionType); ok {
					return fn.ReturnType
				}
				a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Called non-function '%s'", name)
			}
			return typ
		} else if isQualified {
			// Handle Member Access Fallback: rect.width or counter.get
			ids := qCtx.AllIDENTIFIER()
			baseName := ids[0].GetText()
			
			if s, ok := a.currentScope.Resolve(baseName); ok {
				curr := s.Type
				valid := true
				
				for i := 1; i < len(ids); i++ {
					fieldName := ids[i].GetText()
					
					// Auto-deref
					if ptr, ok := curr.(*types.PointerType); ok {
						curr = ptr.ElementType
					}
					
					if st, ok := curr.(*types.StructType); ok {
						if indices, ok := a.structIndices[st.Name]; ok {
							if idx, ok := indices[fieldName]; ok {
								curr = st.Fields[idx]
								continue
							}
						}
						
						// Method check
						methodName := st.Name + "_" + fieldName
						if sym, ok := a.globalScope.Resolve(methodName); ok {
							curr = sym.Type
							continue
						}
						
						valid = false
						break
					} else {
						valid = false; break
					}
				}
				
				if valid {
					if hasParens {
						if fn, ok := curr.(*types.FunctionType); ok {
							return fn.ReturnType
						}
						a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Called non-function '%s'", name)
					}
					return curr
				}
			}
		}
		
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Undefined identifier '%s'", name)
		return types.Void
	}

	// 5. Parenthesized Expression ( expr )
	if ctx.Expression() != nil && hasParens && name == "" {
		return a.Visit(ctx.Expression())
	}
	
	return types.Void
}

// --- Literals & Structures ---

func (a *Analyzer) VisitStructLiteral(ctx *parser.StructLiteralContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
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
	if ctx.STRING_LITERAL() != nil { return types.NewPointer(types.I8) } 
	if ctx.CHAR_LITERAL() != nil { return types.I32 } 
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