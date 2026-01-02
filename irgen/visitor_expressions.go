package irgen

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

// Helper: Safely get operator token type
func getOp(ctx antlr.ParserRuleContext, i int) int {
	childIdx := 2*i - 1
	if childIdx >= 0 && childIdx < ctx.GetChildCount() {
		if term, ok := ctx.GetChild(childIdx).(antlr.TerminalNode); ok {
			return term.GetSymbol().GetTokenType()
		}
	}
	return 0
}

func (g *Generator) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return g.Visit(ctx.LogicalOrExpression())
}

// --- Binary Expressions ---

func (g *Generator) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	lhs := g.Visit(ctx.LogicalAndExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllLogicalAndExpression()); i++ {
		rhs := g.Visit(ctx.LogicalAndExpression(i)).(ir.Value)
		lhs = g.ctx.Builder.CreateOr(lhs, rhs, "")
	}
	return lhs
}

func (g *Generator) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	lhs := g.Visit(ctx.BitOrExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllBitOrExpression()); i++ {
		rhs := g.Visit(ctx.BitOrExpression(i)).(ir.Value)
		lhs = g.ctx.Builder.CreateAnd(lhs, rhs, "")
	}
	return lhs
}

func (g *Generator) VisitBitOrExpression(ctx *parser.BitOrExpressionContext) interface{} {
	lhs := g.Visit(ctx.BitXorExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllBitXorExpression()); i++ {
		rhs := g.Visit(ctx.BitXorExpression(i)).(ir.Value)
		lhs = g.ctx.Builder.CreateOr(lhs, rhs, "")
	}
	return lhs
}

func (g *Generator) VisitBitXorExpression(ctx *parser.BitXorExpressionContext) interface{} {
	lhs := g.Visit(ctx.BitAndExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllBitAndExpression()); i++ {
		rhs := g.Visit(ctx.BitAndExpression(i)).(ir.Value)
		lhs = g.ctx.Builder.CreateXor(lhs, rhs, "")
	}
	return lhs
}

func (g *Generator) VisitBitAndExpression(ctx *parser.BitAndExpressionContext) interface{} {
	lhs := g.Visit(ctx.EqualityExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllEqualityExpression()); i++ {
		rhs := g.Visit(ctx.EqualityExpression(i)).(ir.Value)
		lhs = g.ctx.Builder.CreateAnd(lhs, rhs, "")
	}
	return lhs
}

func (g *Generator) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	lhs := g.Visit(ctx.RelationalExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllRelationalExpression()); i++ {
		rhs := g.Visit(ctx.RelationalExpression(i)).(ir.Value)
		op := getOp(ctx, i)
		if op == parser.ArcParserEQ {
			lhs = g.ctx.Builder.CreateICmpEQ(lhs, rhs, "")
		} else {
			lhs = g.ctx.Builder.CreateICmpNE(lhs, rhs, "")
		}
	}
	return lhs
}

func (g *Generator) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} {
	lhs := g.Visit(ctx.ShiftExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllShiftExpression()); i++ {
		rhs := g.Visit(ctx.ShiftExpression(i)).(ir.Value)
		op := getOp(ctx, i)
		switch op {
		case parser.ArcParserLT: lhs = g.ctx.Builder.CreateICmpSLT(lhs, rhs, "")
		case parser.ArcParserGT: lhs = g.ctx.Builder.CreateICmpSGT(lhs, rhs, "")
		case parser.ArcParserLE: lhs = g.ctx.Builder.CreateICmpSLE(lhs, rhs, "")
		case parser.ArcParserGE: lhs = g.ctx.Builder.CreateICmpSGE(lhs, rhs, "")
		}
	}
	return lhs
}

func (g *Generator) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} {
	lhs := g.Visit(ctx.RangeExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllRangeExpression()); i++ {
		rhs := g.Visit(ctx.RangeExpression(i)).(ir.Value)
		op := getOp(ctx, i)
		if op == parser.ArcParserLT {
			lhs = g.ctx.Builder.CreateShl(lhs, rhs, "")
		} else {
			lhs = g.ctx.Builder.CreateAShr(lhs, rhs, "")
		}
	}
	return lhs
}

func (g *Generator) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} {
	return g.Visit(ctx.AdditiveExpression(0))
}

func (g *Generator) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	lhs := g.Visit(ctx.MultiplicativeExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		rhs := g.Visit(ctx.MultiplicativeExpression(i)).(ir.Value)
		op := getOp(ctx, i)
		
		if types.IsFloat(lhs.Type()) {
			if op == parser.ArcParserPLUS {
				lhs = g.ctx.Builder.CreateFAdd(lhs, rhs, "")
			} else {
				lhs = g.ctx.Builder.CreateFSub(lhs, rhs, "")
			}
		} else {
			if op == parser.ArcParserPLUS {
				lhs = g.ctx.Builder.CreateAdd(lhs, rhs, "")
			} else {
				lhs = g.ctx.Builder.CreateSub(lhs, rhs, "")
			}
		}
	}
	return lhs
}

func (g *Generator) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	lhs := g.Visit(ctx.UnaryExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllUnaryExpression()); i++ {
		rhs := g.Visit(ctx.UnaryExpression(i)).(ir.Value)
		op := getOp(ctx, i)
		
		if types.IsFloat(lhs.Type()) {
			if op == parser.ArcParserSTAR {
				lhs = g.ctx.Builder.CreateFMul(lhs, rhs, "")
			} else {
				lhs = g.ctx.Builder.CreateFDiv(lhs, rhs, "")
			}
		} else {
			if op == parser.ArcParserSTAR {
				lhs = g.ctx.Builder.CreateMul(lhs, rhs, "")
			} else if op == parser.ArcParserSLASH {
				lhs = g.ctx.Builder.CreateSDiv(lhs, rhs, "")
			} else if op == parser.ArcParserPERCENT {
				lhs = g.ctx.Builder.CreateSRem(lhs, rhs, "")
			}
		}
	}
	return lhs
}

// --- Unary Expressions ---

func (g *Generator) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if ctx.AMP() != nil {
		lval := g.getLValue(ctx.UnaryExpression())
		if lval != nil { return lval }
		fmt.Println("[IRGen] Error: Cannot take address of non-lvalue")
		return g.getZeroValue(types.I64)
	}
	
	if ctx.STAR() != nil {
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		if ptrType, ok := val.Type().(*types.PointerType); ok {
			return g.ctx.Builder.CreateLoad(ptrType.ElementType, val, "")
		}
		return val
	}
	
	if ctx.MINUS() != nil {
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		if types.IsFloat(val.Type()) { return g.ctx.Builder.CreateFSub(g.getZeroValue(val.Type()), val, "") }
		return g.ctx.Builder.CreateSub(g.getZeroValue(val.Type()), val, "")
	}
	
	if ctx.NOT() != nil {
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		return g.ctx.Builder.CreateXor(val, g.ctx.Builder.ConstInt(types.I1, 1), "")
	}
	
	if ctx.BIT_NOT() != nil {
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		return g.ctx.Builder.CreateXor(val, g.ctx.Builder.ConstInt(val.Type().(*types.IntType), -1), "")
	}
	
	if ctx.INCREMENT() != nil || ctx.DECREMENT() != nil {
		ptr := g.getLValue(ctx.UnaryExpression())
		if ptr == nil { return g.getZeroValue(types.I64) }
		elemType := ptr.Type().(*types.PointerType).ElementType
		curr := g.ctx.Builder.CreateLoad(elemType, ptr, "")
		var one ir.Value
		if intTy, ok := elemType.(*types.IntType); ok { one = g.ctx.Builder.ConstInt(intTy, 1) } else { one = g.ctx.Builder.ConstInt(types.I64, 1) }
		var next ir.Value
		if ctx.INCREMENT() != nil { next = g.ctx.Builder.CreateAdd(curr, one, "") } else { next = g.ctx.Builder.CreateSub(curr, one, "") }
		g.ctx.Builder.CreateStore(next, ptr)
		return next
	}

	if ctx.PostfixExpression() != nil {
		return g.Visit(ctx.PostfixExpression())
	}
	
	return g.getZeroValue(types.I64)
}

// --- Postfix Expressions ---

func (g *Generator) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	var currPtr ir.Value = g.getLValue(ctx.PrimaryExpression())
	var curr ir.Value

	if currPtr != nil {
		if _, isFn := currPtr.(*ir.Function); isFn {
			curr = currPtr
			currPtr = nil
		} else if ptrType, ok := currPtr.Type().(*types.PointerType); ok {
			curr = g.ctx.Builder.CreateLoad(ptrType.ElementType, currPtr, "")
		} else {
			curr = currPtr
		}
	} else {
		res := g.Visit(ctx.PrimaryExpression())
		if res != nil {
			curr = res.(ir.Value)
		}
	}

	// If the primary expression was a call (e.g. foo()), curr might be the result.
	// We iterate ops like .field, [index], etc.
	for _, op := range ctx.AllPostfixOp() {
		// Function Call via postfix op (unlikely if grammar puts calls in PrimaryExpression, but possible for func pointers)
		if op.LPAREN() != nil {
			var args []ir.Value
			if op.ArgumentList() != nil {
				for _, arg := range op.ArgumentList().AllArgument() {
					if v := g.Visit(arg.Expression()); v != nil { 
						args = append(args, v.(ir.Value)) 
					}
				}
			}

			// Standard Call via function pointer or resolved entity
			if curr != nil {
				if fn, ok := curr.(*ir.Function); ok {
					// Cast arguments
					if len(args) == len(fn.FuncType.ParamTypes) || (fn.FuncType.Variadic && len(args) >= len(fn.FuncType.ParamTypes)) {
						for i, paramType := range fn.FuncType.ParamTypes {
							if i < len(args) {
								args[i] = g.emitCast(args[i], paramType)
							}
						}
					}
					curr = g.ctx.Builder.CreateCall(fn, args, "")
				} else {
					// Indirect Call attempt (function pointer)
					// Requires pointer-to-function type check
					panic(fmt.Sprintf("Indirect calls not fully supported here. Target: %v", curr))
				}
				currPtr = nil
			}
			continue
		}
		
		// Field Access
		if op.DOT() != nil {
			fieldName := op.IDENTIFIER().GetText()
			
			var basePtr ir.Value = currPtr
			if basePtr == nil && curr != nil && types.IsPointer(curr.Type()) {
				basePtr = curr
			}

			if basePtr != nil {
				ptrType := basePtr.Type().(*types.PointerType)
				if st, ok := ptrType.ElementType.(*types.StructType); ok {
					if idx, ok := g.analysis.StructIndices[st.Name][fieldName]; ok {
						currPtr = g.ctx.Builder.CreateStructGEP(st, basePtr, idx, "")
						curr = g.ctx.Builder.CreateLoad(st.Fields[idx], currPtr, "")
					}
				}
			}
			continue
		}
		
		// Indexing
		if op.LBRACKET() != nil {
			idx := g.Visit(op.Expression()).(ir.Value)
			
			var basePtr ir.Value = currPtr
			if basePtr == nil && curr != nil && types.IsPointer(curr.Type()) {
				basePtr = curr
			}

			if basePtr != nil {
				ptrType := basePtr.Type().(*types.PointerType)
				var elemPtr ir.Value
				
				if _, isArray := ptrType.ElementType.(*types.ArrayType); isArray {
					zero := g.ctx.Builder.ConstInt(types.I32, 0)
					elemPtr = g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, basePtr, []ir.Value{zero, idx}, "")
				} else {
					elemPtr = g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, basePtr, []ir.Value{idx}, "")
				}
				
				currPtr = elemPtr
				curr = g.ctx.Builder.CreateLoad(ptrType.ElementType, currPtr, "")
			}
			continue
		}

		if op.INCREMENT() != nil || op.DECREMENT() != nil {
			if currPtr != nil {
				one := g.ctx.Builder.ConstInt(types.I64, 1)
				if intTy, ok := curr.Type().(*types.IntType); ok { one = g.ctx.Builder.ConstInt(intTy, 1) }
				var next ir.Value
				if op.INCREMENT() != nil { next = g.ctx.Builder.CreateAdd(curr, one, "") } else { next = g.ctx.Builder.CreateSub(curr, one, "") }
				g.ctx.Builder.CreateStore(next, currPtr)
				currPtr = nil
			}
		}
	}
	return curr
}

// --- Primary Expressions ---

func (g *Generator) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.StructLiteral() != nil { return g.Visit(ctx.StructLiteral()) }
	if ctx.Literal() != nil { return g.Visit(ctx.Literal()) }

	if ctx.SizeofExpression() != nil {
		t := g.resolveType(ctx.SizeofExpression().Type_())
		return g.ctx.Builder.CreateSizeOf(t, "")
	}
	if ctx.AlignofExpression() != nil {
		t := g.resolveType(ctx.AlignofExpression().Type_())
		return g.ctx.Builder.CreateAlignOf(t, "")
	}

	// Parenthesized Expression: ( expr )
	// Grammar: LPAREN expression RPAREN (without args/identifier)
	// But grammar says: LPAREN expression RPAREN is a top-level alt.
	// Also qualified/IDENTIFIER alts have (LPAREN argumentList? RPAREN)?
	if ctx.Expression() != nil && ctx.LPAREN() != nil && ctx.IDENTIFIER() == nil && ctx.QualifiedIdentifier() == nil {
		return g.Visit(ctx.Expression())
	}

	// Handle Identifier or QualifiedIdentifier
	var name string
	var isQualified bool

	if ctx.QualifiedIdentifier() != nil {
		q := ctx.QualifiedIdentifier()
		for i, id := range q.AllIDENTIFIER() {
			if i > 0 { name += "." }
			name += id.GetText()
		}
		isQualified = true
	} else if ctx.IDENTIFIER() != nil {
		name = ctx.IDENTIFIER().GetText()
	}

	if name != "" {
		// 1. Check for function call syntax: IDENT(...) or QUAL.ID(...)
		// The grammar has optional parens at the end of the identifier rules
		isCall := ctx.LPAREN() != nil && (ctx.IDENTIFIER() != nil || ctx.QualifiedIdentifier() != nil)
		
		var args []ir.Value
		if isCall {
			if ctx.ArgumentList() != nil {
				for _, arg := range ctx.ArgumentList().AllArgument() {
					if v := g.Visit(arg.Expression()); v != nil { 
						args = append(args, v.(ir.Value)) 
					}
				}
			}

			// Handle intrinsics/casts (mostly for simple identifiers)
			if !isQualified {
				if name == "cast" && len(args) == 1 {
					if ga := ctx.GenericArgs(); ga != nil {
						if gl := ga.GenericArgList(); gl != nil && len(gl.AllGenericArg()) > 0 {
							targetType := g.resolveType(gl.GenericArg(0).Type_())
							return g.emitCast(args[0], targetType)
						}
					}
				}
				if intrinsicVal := g.GenerateIntrinsicCall(name, args); intrinsicVal != nil {
					return intrinsicVal
				}
			}
		}

		// 2. Resolve Symbol
		var entity ir.Value
		
		if sym, ok := g.currentScope.Resolve(name); ok {
			if sym.Kind == symbol.SymFunc && sym.IRValue == nil {
				// Intrinsic placeholder or undefined forward decl
				return nil
			}
			// If it's a variable and we are NOT calling it, load the value.
			// If we ARE calling it, we might be loading a func ptr, or calling the IR Function directly.
			if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
				if !isCall {
					return g.ctx.Builder.CreateLoad(sym.Type, alloca, "")
				}
				// Call to variable (func ptr)
				entity = g.ctx.Builder.CreateLoad(sym.Type, alloca, "")
			} else {
				// Likely *ir.Function or Global
				entity = sym.IRValue
			}
		} else if fn := g.ctx.Module.GetFunction(name); fn != nil {
			entity = fn
		} else if glob := g.ctx.Module.GetGlobal(name); glob != nil {
			if !isCall {
				return g.ctx.Builder.CreateLoad(glob.Type().(*types.PointerType).ElementType, glob, "")
			}
			// Global function pointer not loaded yet? 
			// If Global is a function declaration, it's covered by GetFunction usually.
			// If Global is a variable holding a func ptr:
			entity = g.ctx.Builder.CreateLoad(glob.Type().(*types.PointerType).ElementType, glob, "")
		} else {
			// Not found
			if isCall {
				fmt.Printf("[IRGen] Error: Call to undefined function '%s'\n", name)
			} else {
				fmt.Printf("[IRGen] Error: Undefined identifier '%s'\n", name)
			}
			return g.getZeroValue(types.I64)
		}

		// 3. Generate Call
		if isCall {
			if fn, ok := entity.(*ir.Function); ok {
				// Cast arguments to match function parameters
				if len(args) == len(fn.FuncType.ParamTypes) || (fn.FuncType.Variadic && len(args) >= len(fn.FuncType.ParamTypes)) {
					for i, paramType := range fn.FuncType.ParamTypes {
						if i < len(args) {
							args[i] = g.emitCast(args[i], paramType)
						}
					}
				}
				return g.ctx.Builder.CreateCall(fn, args, "")
			} else {
				// Indirect call (pointer to function)
				// Simplified: assume entity is the function pointer value
				// In a real implementation, we'd check the type of entity to ensure it's a function pointer
				// and extract the return type.
				// For this fix, we simply error as indirect calls need more builder support.
				panic(fmt.Sprintf("Indirect calls not fully implemented. Target: %s", name))
			}
		}

		return entity
	}

	return g.getZeroValue(types.I64)
}

func (g *Generator) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	txt := ctx.GetText()
	if ctx.NULL() != nil { return g.ctx.Builder.ConstNull(types.NewPointer(types.Void)) }
	
	if ctx.CHAR_LITERAL() != nil {
		if len(txt) >= 2 {
			r := []rune(txt)[1]
			return g.ctx.Builder.ConstInt(types.I32, int64(r))
		}
		return g.getZeroValue(types.I32)
	}

	if ctx.INTEGER_LITERAL() != nil {
		val, _ := strconv.ParseInt(txt, 0, 64)
		return g.ctx.Builder.ConstInt(types.I64, val)
	}
	if ctx.FLOAT_LITERAL() != nil {
		val, _ := strconv.ParseFloat(txt, 64)
		return g.ctx.Builder.ConstFloat(types.F64, val)
	}
	if ctx.BOOLEAN_LITERAL() != nil {
		val := int64(0)
		if txt == "true" { val = 1 }
		return g.ctx.Builder.ConstInt(types.I1, val)
	}
	if ctx.STRING_LITERAL() != nil {
		if len(txt) >= 2 { txt = txt[1 : len(txt)-1] }
		content := txt + "\x00"
		arrType := types.NewArray(types.I8, int64(len(content)))
		var chars []ir.Constant
		for _, b := range []byte(content) { chars = append(chars, g.ctx.Builder.ConstInt(types.I8, int64(b))) }
		global := g.ctx.Builder.CreateGlobalConstant(".str", &ir.ConstantArray{BaseValue: ir.BaseValue{ValType: arrType}, Elements: chars})
		zero := g.ctx.Builder.ConstInt(types.I32, 0)
		return g.ctx.Builder.CreateInBoundsGEP(arrType, global, []ir.Value{zero, zero}, "")
	}
	return g.getZeroValue(types.I64)
}

func (g *Generator) VisitStructLiteral(ctx *parser.StructLiteralContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	sym, ok := g.currentScope.Resolve(name)
	if !ok { return g.getZeroValue(types.I64) }
	structType := sym.Type.(*types.StructType)
	var agg ir.Value = g.ctx.Builder.ConstZero(structType)
	indices := g.analysis.StructIndices[structType.Name]
	for _, field := range ctx.AllFieldInit() {
		fName := field.IDENTIFIER().GetText()
		fVal := g.Visit(field.Expression()).(ir.Value)
		if idx, ok := indices[fName]; ok {
			fVal = g.emitCast(fVal, structType.Fields[idx])
			agg = g.ctx.Builder.CreateInsertValue(agg, fVal, []int{idx}, "")
		}
	}
	return agg
}