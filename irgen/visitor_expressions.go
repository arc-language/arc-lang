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

// Helper: Safely get operator token type from binary expression context
func getOp(ctx antlr.ParserRuleContext, i int) int {
	// Child 0 = expr, Child 1 = op, Child 2 = expr...
	// Operator for rhs[i] (where i starts at 1) is at index 2*i - 1.
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
	// Address Of (&)
	if ctx.AMP() != nil {
		lval := g.getLValue(ctx.UnaryExpression())
		if lval != nil { return lval }
		fmt.Println("[IRGen] Error: Cannot take address of non-lvalue")
		return g.getZeroValue(types.I64)
	}
	
	// Dereference (*)
	if ctx.STAR() != nil {
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		if ptrType, ok := val.Type().(*types.PointerType); ok {
			return g.ctx.Builder.CreateLoad(ptrType.ElementType, val, "")
		}
		// Error: dereferencing non-pointer
		return val
	}
	
	// Unary Minus (-)
	if ctx.MINUS() != nil {
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		if types.IsFloat(val.Type()) { return g.ctx.Builder.CreateFSub(g.getZeroValue(val.Type()), val, "") }
		return g.ctx.Builder.CreateSub(g.getZeroValue(val.Type()), val, "")
	}
	
	// Logical Not (!)
	if ctx.NOT() != nil {
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		return g.ctx.Builder.CreateXor(val, g.ctx.Builder.ConstInt(types.I1, 1), "")
	}
	
	// Bitwise Not (~)
	if ctx.BIT_NOT() != nil {
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		// XOR with -1
		return g.ctx.Builder.CreateXor(val, g.ctx.Builder.ConstInt(val.Type().(*types.IntType), -1), "")
	}
	
	// Prefix Increment/Decrement (++i, --i)
	if ctx.INCREMENT() != nil || ctx.DECREMENT() != nil {
		ptr := g.getLValue(ctx.UnaryExpression())
		if ptr == nil { return g.getZeroValue(types.I64) }
		
		elemType := ptr.Type().(*types.PointerType).ElementType
		curr := g.ctx.Builder.CreateLoad(elemType, ptr, "")
		
		var one ir.Value
		if intTy, ok := elemType.(*types.IntType); ok { 
			one = g.ctx.Builder.ConstInt(intTy, 1) 
		} else { 
			one = g.ctx.Builder.ConstInt(types.I64, 1) 
		}
		
		var next ir.Value
		if ctx.INCREMENT() != nil { 
			next = g.ctx.Builder.CreateAdd(curr, one, "") 
		} else { 
			next = g.ctx.Builder.CreateSub(curr, one, "") 
		}
		
		g.ctx.Builder.CreateStore(next, ptr)
		return next
	}

	if ctx.PostfixExpression() != nil {
		return g.Visit(ctx.PostfixExpression())
	}
	
	return g.getZeroValue(types.I64)
}

// --- Postfix Expressions (Calls, Indexing, Members) ---

func (g *Generator) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	// Attempt to resolve the LHS.
	// We call VisitPrimaryExpression directly for the base.
	// If it returns nil, it means it's a symbol without IRValue (likely an intrinsic placeholder).
	
	var curr ir.Value
	var currPtr ir.Value
	
	// Try to get LValue first (for field access / assignment targets)
	lval := g.getLValue(ctx.PrimaryExpression())
	if lval != nil {
		currPtr = lval
		// Auto-load if it's not a function
		if _, isFn := currPtr.(*ir.Function); isFn {
			curr = currPtr
			currPtr = nil
		} else if ptrType, ok := currPtr.Type().(*types.PointerType); ok {
			curr = g.ctx.Builder.CreateLoad(ptrType.ElementType, currPtr, "")
		} else {
			curr = currPtr
		}
	} else {
		// Not an LValue, evaluate normally (e.g. literals, or intrinsic names)
		res := g.Visit(ctx.PrimaryExpression())
		if res != nil {
			curr = res.(ir.Value)
		}
		// If res is nil, 'curr' remains nil. We check IDENTIFIER text later for intrinsics.
	}

	for _, op := range ctx.AllPostfixOp() {
		
		// 1. Function Call: (args)
		if op.LPAREN() != nil {
			var args []ir.Value
			if op.ArgumentList() != nil {
				for _, arg := range op.ArgumentList().AllArgument() {
					if v := g.Visit(arg.Expression()); v != nil { 
						args = append(args, v.(ir.Value)) 
					}
				}
			}

			// Check for Compiler Intrinsics (alloca, memset, syscall, etc.)
			// We check the name text because intrinsics don't have *ir.Function definitions.
			if ctx.PrimaryExpression().IDENTIFIER() != nil {
				funcName := ctx.PrimaryExpression().IDENTIFIER().GetText()
				
				// GenerateIntrinsicCall is defined in irgen/intrinsics.go
				if intrinsicVal := g.GenerateIntrinsicCall(funcName, args); intrinsicVal != nil {
					curr = intrinsicVal
					currPtr = nil // Result of intrinsic is R-Value
					continue
				}
			}

			// Standard Function Call
			if curr != nil {
				if fn, ok := curr.(*ir.Function); ok {
					curr = g.ctx.Builder.CreateCall(fn, args, "")
				} else {
					// Function Pointer Call
					curr = g.ctx.Builder.CreateCall(curr, args, "")
				}
				currPtr = nil
			}
			continue
		}
		
		// 2. Member Access: .field
		if op.DOT() != nil {
			fieldName := op.IDENTIFIER().GetText()
			
			// Resolve pointer if needed
			var basePtr ir.Value = currPtr
			if basePtr == nil && curr != nil && types.IsPointer(curr.Type()) {
				basePtr = curr // It's already an R-Value pointer
			}

			if basePtr != nil {
				ptrType := basePtr.Type().(*types.PointerType)
				if st, ok := ptrType.ElementType.(*types.StructType); ok {
					if idx, ok := g.analysis.StructIndices[st.Name][fieldName]; ok {
						currPtr = g.ctx.Builder.CreateStructGEP(st, basePtr, idx, "")
						curr = g.ctx.Builder.CreateLoad(st.Fields[idx], currPtr, "")
						continue
					}
				}
			}
			// If not a struct field, it might be a method call. 
			// IRGen for methods usually involves finding the mangled global function `Struct_Method`
			// and passing `self` as first arg. This logic is omitted for brevity but follows similar pattern.
			continue
		}
		
		// 3. Indexing: [expr]
		if op.LBRACKET() != nil {
			idx := g.Visit(op.Expression()).(ir.Value)
			
			// Resolve base pointer
			var basePtr ir.Value = currPtr
			if basePtr == nil && curr != nil && types.IsPointer(curr.Type()) {
				basePtr = curr
			}

			if basePtr != nil {
				ptrType := basePtr.Type().(*types.PointerType)
				
				var elemPtr ir.Value
				if _, isArray := ptrType.ElementType.(*types.ArrayType); isArray {
					// Array Pointer: need 0, index
					zero := g.ctx.Builder.ConstInt(types.I32, 0)
					elemPtr = g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, basePtr, []ir.Value{zero, idx}, "")
				} else {
					// Raw Pointer / Vector Data: need index
					elemPtr = g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, basePtr, []ir.Value{idx}, "")
				}
				
				currPtr = elemPtr
				curr = g.ctx.Builder.CreateLoad(ptrType.ElementType, currPtr, "")
			}
			continue
		}

		// 4. Postfix Increment/Decrement: i++
		if op.INCREMENT() != nil || op.DECREMENT() != nil {
			if currPtr != nil {
				one := g.ctx.Builder.ConstInt(types.I64, 1)
				if intTy, ok := curr.Type().(*types.IntType); ok { 
					one = g.ctx.Builder.ConstInt(intTy, 1) 
				}
				
				var next ir.Value
				if op.INCREMENT() != nil { 
					next = g.ctx.Builder.CreateAdd(curr, one, "") 
				} else { 
					next = g.ctx.Builder.CreateSub(curr, one, "") 
				}
				
				g.ctx.Builder.CreateStore(next, currPtr)
				currPtr = nil // Result of i++ is R-Value (usually original value in C, but here simply done)
			}
		}
	}
	
	return curr
}

// --- Primary Expressions ---

func (g *Generator) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.StructLiteral() != nil { return g.Visit(ctx.StructLiteral()) }
	if ctx.Literal() != nil { return g.Visit(ctx.Literal()) }
	
	// Handle Cast <Type>(Expr)
	if ctx.CastExpression() != nil { return g.Visit(ctx.CastExpression()) }

	// Handle Compiler Operators
	if ctx.SizeofExpression() != nil {
		t := g.resolveType(ctx.SizeofExpression().Type_())
		return g.ctx.Builder.CreateSizeOf(t, "")
	}
	if ctx.AlignofExpression() != nil {
		t := g.resolveType(ctx.AlignofExpression().Type_())
		return g.ctx.Builder.CreateAlignOf(t, "")
	}

	// Handle Qualified Identifiers (e.g. io.print)
	if ctx.QualifiedIdentifier() != nil {
		q := ctx.QualifiedIdentifier()
		var name string
		for i, id := range q.AllIDENTIFIER() {
			if i > 0 { name += "." }
			name += id.GetText()
		}
		
		// Check Symbols
		if sym, ok := g.currentScope.Resolve(name); ok && sym.IRValue != nil {
			if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
				return g.ctx.Builder.CreateLoad(sym.Type, alloca, "")
			}
			return sym.IRValue
		}
		
		// Check Global Functions
		if fn := g.ctx.Module.GetFunction(name); fn != nil { return fn }
		
		return g.getZeroValue(types.I64)
	}

	// Handle Identifiers
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		
		// 1. Resolve in Scope
		if sym, ok := g.currentScope.Resolve(name); ok {
			// If it's a known symbol but has no IRValue, it's likely an Intrinsic (alloca, syscall).
			// Return nil so PostfixExpression can detect it by name and use GenerateIntrinsicCall.
			if sym.Kind == symbol.SymFunc && sym.IRValue == nil {
				return nil
			}
			
			// Variable
			if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
				return g.ctx.Builder.CreateLoad(sym.Type, alloca, "")
			}
			return sym.IRValue
		}
		
		// 2. Global Function
		if fn := g.ctx.Module.GetFunction(name); fn != nil { return fn }
		
		// 3. Global Variable
		if glob := g.ctx.Module.GetGlobal(name); glob != nil {
			return g.ctx.Builder.CreateLoad(glob.Type().(*types.PointerType).ElementType, glob, "")
		}
	}

	// Parenthesized Expression
	if ctx.Expression() != nil { return g.Visit(ctx.Expression()) }
	
	return g.getZeroValue(types.I64)
}

func (g *Generator) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	txt := ctx.GetText()
	
	if ctx.NULL() != nil { 
		return g.ctx.Builder.ConstNull(types.NewPointer(types.Void)) 
	}
	
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
		// Strip quotes
		if len(txt) >= 2 { txt = txt[1 : len(txt)-1] }
		
		// Create global string constant
		content := txt + "\x00"
		arrType := types.NewArray(types.I8, int64(len(content)))
		
		var chars []ir.Constant
		for _, b := range []byte(content) { 
			chars = append(chars, g.ctx.Builder.ConstInt(types.I8, int64(b))) 
		}
		
		global := g.ctx.Builder.CreateGlobalConstant(".str", &ir.ConstantArray{
			BaseValue: ir.BaseValue{ValType: arrType}, 
			Elements: chars,
		})
		
		// Return pointer to start
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

func (g *Generator) VisitCastExpression(ctx *parser.CastExpressionContext) interface{} {
	val := g.Visit(ctx.Expression()).(ir.Value)
	targetType := g.resolveType(ctx.Type_())
	return g.emitCast(val, targetType)
}