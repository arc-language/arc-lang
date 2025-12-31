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

// Helper to get operator token type at a specific index
func getBinaryOp(ctx antlr.ParserRuleContext, index int) int {
	if child := ctx.GetChild(2*index - 1); child != nil {
		if term, ok := child.(antlr.TerminalNode); ok {
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
		op := getBinaryOp(ctx, i)
		
		if op == parser.ArcParserEQ {
			lhs = g.ctx.Builder.CreateICmpEQ(lhs, rhs, "")
		} else if op == parser.ArcParserNEQ {
			lhs = g.ctx.Builder.CreateICmpNE(lhs, rhs, "")
		}
	}
	return lhs
}

func (g *Generator) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} {
	lhs := g.Visit(ctx.ShiftExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllShiftExpression()); i++ {
		rhs := g.Visit(ctx.ShiftExpression(i)).(ir.Value)
		op := getBinaryOp(ctx, i)
		
		switch op {
		case parser.ArcParserLT:
			lhs = g.ctx.Builder.CreateICmpSLT(lhs, rhs, "")
		case parser.ArcParserGT:
			lhs = g.ctx.Builder.CreateICmpSGT(lhs, rhs, "")
		case parser.ArcParserLE:
			lhs = g.ctx.Builder.CreateICmpSLE(lhs, rhs, "")
		case parser.ArcParserGE:
			lhs = g.ctx.Builder.CreateICmpSGE(lhs, rhs, "")
		}
	}
	return lhs
}

func (g *Generator) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} {
	lhs := g.Visit(ctx.RangeExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllRangeExpression()); i++ {
		rhs := g.Visit(ctx.RangeExpression(i)).(ir.Value)
		op := getBinaryOp(ctx, i)
		
		// Fallback logic for token checking if specific SHL/SHR constants are not direct matches
		// or if the lexer outputs LT LT for SHL. Assuming distinct tokens here.
		if op == parser.ArcParserSHL { 
			lhs = g.ctx.Builder.CreateShl(lhs, rhs, "")
		} else if op == parser.ArcParserSHR {
			lhs = g.ctx.Builder.CreateAShr(lhs, rhs, "")
		} else if op == parser.ArcParserLT {
			// Fallback: Parser might see '<<' as two LT tokens if not defined as SHL
			lhs = g.ctx.Builder.CreateShl(lhs, rhs, "") 
		} else if op == parser.ArcParserGT {
			lhs = g.ctx.Builder.CreateAShr(lhs, rhs, "")
		}
	}
	return lhs
}

func (g *Generator) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} {
	// Range Expressions (start..end) are typically handled in For Loops specifically.
	// As a general expression value, they might return a Struct{start, end}.
	// For now, pass through to additive.
	return g.Visit(ctx.AdditiveExpression(0))
}

func (g *Generator) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	lhs := g.Visit(ctx.MultiplicativeExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		rhs := g.Visit(ctx.MultiplicativeExpression(i)).(ir.Value)
		op := getBinaryOp(ctx, i)

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
		op := getBinaryOp(ctx, i)
		
		if types.IsFloat(lhs.Type()) {
			if op == parser.ArcParserSTAR {
				lhs = g.ctx.Builder.CreateFMul(lhs, rhs, "")
			} else if op == parser.ArcParserSLASH {
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
	// Handle Prefix Increment/Decrement (++i, --i)
	if ctx.INCREMENT() != nil || ctx.DECREMENT() != nil {
		ptr := g.getLValue(ctx.UnaryExpression())
		if ptr == nil { return g.getZeroValue(types.I64) }

		ptrType := ptr.Type().(*types.PointerType)
		elemType := ptrType.ElementType
		curr := g.ctx.Builder.CreateLoad(elemType, ptr, "")

		var next ir.Value
		if types.IsFloat(elemType) {
			oneF := g.ctx.Builder.ConstFloat(types.F64, 1.0)
			if ctx.INCREMENT() != nil { next = g.ctx.Builder.CreateFAdd(curr, oneF, "") } else { next = g.ctx.Builder.CreateFSub(curr, oneF, "") }
		} else {
			one := g.ctx.Builder.ConstInt(types.I64, 1)
			if intTy, ok := elemType.(*types.IntType); ok { one = g.ctx.Builder.ConstInt(intTy, 1) }
			if ctx.INCREMENT() != nil { next = g.ctx.Builder.CreateAdd(curr, one, "") } else { next = g.ctx.Builder.CreateSub(curr, one, "") }
		}

		g.ctx.Builder.CreateStore(next, ptr)
		return next
	}

	if ctx.PostfixExpression() != nil { return g.Visit(ctx.PostfixExpression()) }
	
	val := g.Visit(ctx.UnaryExpression()).(ir.Value)
	
	if ctx.MINUS() != nil {
		if types.IsFloat(val.Type()) { return g.ctx.Builder.CreateFSub(g.getZeroValue(val.Type()), val, "") }
		return g.ctx.Builder.CreateSub(g.getZeroValue(val.Type()), val, "")
	}
	if ctx.NOT() != nil {
		return g.ctx.Builder.CreateXor(val, g.ctx.Builder.ConstInt(types.I1, 1), "")
	}
	if ctx.BIT_NOT() != nil {
		return g.ctx.Builder.CreateXor(val, g.ctx.Builder.ConstInt(val.Type().(*types.IntType), -1), "")
	}
	if ctx.STAR() != nil { // Dereference (*ptr)
		if ptr, ok := val.Type().(*types.PointerType); ok {
			return g.ctx.Builder.CreateLoad(ptr.ElementType, val, "")
		}
	}
	if ctx.AMP() != nil { // Address Of (&var)
		// To implement '&', we need the address, not the value.
		// Visit() normally returns the R-Value (Load). 
		// We use getLValue to retrieve the address instead.
		if child := ctx.UnaryExpression(); child != nil {
			if lval := g.getLValue(child); lval != nil { return lval }
		}
		return g.getZeroValue(types.I64)
	}
	return val
}

// --- Postfix Expressions ---

func (g *Generator) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	// Initialize with the Primary Expression
	// We attempt to get an L-Value address first to support mutation chains (e.g. arr[i]++)
	var currPtr ir.Value = g.getLValue(ctx.PrimaryExpression())
	var curr ir.Value

	if currPtr != nil {
		// If it's an L-Value, load it to get the R-Value for calculation
		ptrType := currPtr.Type().(*types.PointerType)
		curr = g.ctx.Builder.CreateLoad(ptrType.ElementType, currPtr, "")
	} else {
		// Otherwise, visit normally to get the R-Value (e.g. function call return)
		curr = g.Visit(ctx.PrimaryExpression()).(ir.Value)
	}

	for _, op := range ctx.AllPostfixOp() {
		// 1. Function Call
		if op.LPAREN() != nil {
			var args []ir.Value
			if op.ArgumentList() != nil {
				for _, arg := range op.ArgumentList().AllArgument() {
					if argExpr := g.Visit(arg.Expression()); argExpr != nil {
						args = append(args, argExpr.(ir.Value))
					}
				}
			}
			if fn, ok := curr.(*ir.Function); ok {
				curr = g.ctx.Builder.CreateCall(fn, args, "")
				// Result of a function call is an R-Value
				currPtr = nil 
			}
			continue
		}
		
		// 2. Member Access (.field)
		if op.DOT() != nil {
			fieldName := op.IDENTIFIER().GetText()
			
			var structType *types.StructType
			var isPtr = false
			
			if ptr, ok := curr.Type().(*types.PointerType); ok {
				if st, ok := ptr.ElementType.(*types.StructType); ok {
					structType = st
					isPtr = true
				}
			} else if st, ok := curr.Type().(*types.StructType); ok {
				structType = st
			}

			if structType != nil {
				if idx, ok := g.analysis.StructIndices[structType.Name][fieldName]; ok {
					if isPtr && currPtr != nil {
						// L-Value access: GEP + Load
						// currPtr tracks the address of the struct pointer variable
						// We need to GEP from curr (the pointer value)
						gep := g.ctx.Builder.CreateStructGEP(structType, curr, idx, "")
						currPtr = gep
						curr = g.ctx.Builder.CreateLoad(structType.Fields[idx], gep, "")
					} else if isPtr {
						// R-Value pointer (e.g. from function return)
						gep := g.ctx.Builder.CreateStructGEP(structType, curr, idx, "")
						currPtr = gep
						curr = g.ctx.Builder.CreateLoad(structType.Fields[idx], gep, "")
					} else {
						// R-Value struct (value)
						curr = g.ctx.Builder.CreateExtractValue(curr, []int{idx}, "")
						currPtr = nil
					}
				}
			}
			continue
		}
		
		// 3. Array Indexing ([i])
		if op.LBRACKET() != nil {
			idx := g.Visit(op.Expression()).(ir.Value)
			
			if ptr, ok := curr.Type().(*types.PointerType); ok {
				// GEP the pointer
				gep := g.ctx.Builder.CreateInBoundsGEP(ptr.ElementType, curr, []ir.Value{idx}, "")
				currPtr = gep
				curr = g.ctx.Builder.CreateLoad(ptr.ElementType, gep, "")
			}
			continue
		}

		// 4. Increment / Decrement (++, --)
		if op.INCREMENT() != nil || op.DECREMENT() != nil {
			if currPtr == nil { continue }
			
			var next ir.Value
			elemType := curr.Type()
			
			if types.IsFloat(elemType) {
				oneF := g.ctx.Builder.ConstFloat(types.F64, 1.0)
				if op.INCREMENT() != nil { 
					next = g.ctx.Builder.CreateFAdd(curr, oneF, "") 
				} else { 
					next = g.ctx.Builder.CreateFSub(curr, oneF, "") 
				}
			} else {
				one := g.ctx.Builder.ConstInt(types.I64, 1)
				if intTy, ok := elemType.(*types.IntType); ok { one = g.ctx.Builder.ConstInt(intTy, 1) }
				
				if op.INCREMENT() != nil { 
					next = g.ctx.Builder.CreateAdd(curr, one, "") 
				} else { 
					next = g.ctx.Builder.CreateSub(curr, one, "") 
				}
			}
			
			g.ctx.Builder.CreateStore(next, currPtr)
			// Postfix returns old value, 'curr' is currently holding old value
			currPtr = nil
		}
	}
	return curr
}

// --- Primary Expressions ---

func (g *Generator) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.StructLiteral() != nil { return g.Visit(ctx.StructLiteral()) }
	if ctx.Literal() != nil { return g.Visit(ctx.Literal()) }
	if ctx.CastExpression() != nil { return g.Visit(ctx.CastExpression()) }
	if ctx.SyscallExpression() != nil { return g.Visit(ctx.SyscallExpression()) }
	if ctx.IntrinsicExpression() != nil { return g.Visit(ctx.IntrinsicExpression()) }
	
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		
		// 1. Resolve in Scope
		if sym, ok := g.currentScope.Resolve(name); ok {
			// If it's a variable (Alloca), load it
			if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
				return g.ctx.Builder.CreateLoad(sym.Type, alloca, "")
			}
			// If it's a constant or function, return directly
			return sym.IRValue
		}
		
		// 2. Resolve Global/Extern
		if glob := g.ctx.Module.GetGlobal(name); glob != nil {
			// Globals are pointers, load the value
			return g.ctx.Builder.CreateLoad(glob.Type().(*types.PointerType).ElementType, glob, "")
		}
	}
	
	if ctx.Expression() != nil { return g.Visit(ctx.Expression()) }
	
	return g.getZeroValue(types.I64)
}

func (g *Generator) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTEGER_LITERAL() != nil {
		txt := ctx.INTEGER_LITERAL().GetText()
		val, _ := strconv.ParseInt(txt, 0, 64)
		return g.ctx.Builder.ConstInt(types.I64, val)
	}
	if ctx.FLOAT_LITERAL() != nil {
		txt := ctx.FLOAT_LITERAL().GetText()
		val, _ := strconv.ParseFloat(txt, 64)
		return g.ctx.Builder.ConstFloat(types.F64, val)
	}
	if ctx.STRING_LITERAL() != nil {
		raw := ctx.STRING_LITERAL().GetText()
		if len(raw) >= 2 { raw = raw[1 : len(raw)-1] }
		content := raw + "\x00"
		
		strName := fmt.Sprintf(".str.%d", len(g.ctx.Module.Globals))
		arrType := types.NewArray(types.I8, int64(len(content)))
		var chars []ir.Constant
		for _, b := range []byte(content) {
			chars = append(chars, g.ctx.Builder.ConstInt(types.I8, int64(b)))
		}
		constArr := &ir.ConstantArray{BaseValue: ir.BaseValue{ValType: arrType}, Elements: chars}
		
		global := g.ctx.Builder.CreateGlobalConstant(strName, constArr)
		zero := g.ctx.Builder.ConstInt(types.I32, 0)
		return g.ctx.Builder.CreateInBoundsGEP(arrType, global, []ir.Value{zero, zero}, "")
	}
	return g.getZeroValue(types.I64)
}

func (g *Generator) VisitStructLiteral(ctx *parser.StructLiteralContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	sym, ok := g.currentScope.Resolve(name)
	if !ok || sym.Kind != symbol.SymType { return g.getZeroValue(types.I64) }
	
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
	target := g.resolveType(ctx.Type_())
	return g.emitCast(val, target)
}

func (g *Generator) VisitSyscallExpression(ctx *parser.SyscallExpressionContext) interface{} {
	var args []ir.Value
	for _, expr := range ctx.AllExpression() {
		args = append(args, g.Visit(expr).(ir.Value))
	}
	return g.ctx.Builder.CreateSyscall(args)
}

func (g *Generator) VisitIntrinsicExpression(ctx *parser.IntrinsicExpressionContext) interface{} {
	// 1. SIZEOF
	if ctx.SIZEOF() != nil {
		t := g.resolveType(ctx.Type_())
		return g.ctx.Builder.CreateSizeOf(t, "")
	}

	// 2. ALIGNOF
	if ctx.ALIGNOF() != nil {
		t := g.resolveType(ctx.Type_())
		return g.ctx.Builder.CreateAlignOf(t, "")
	}

	// Helper to resolve arguments
	var args []ir.Value
	for _, expr := range ctx.AllExpression() {
		args = append(args, g.Visit(expr).(ir.Value))
	}

	// 3. Memory Intrinsics
	if ctx.MEMSET() != nil && len(args) == 3 {
		// memset(dest, val, size)
		return g.ctx.Builder.CreateMemSet(args[0], args[1], args[2])
	}
	if ctx.MEMCPY() != nil && len(args) == 3 {
		// memcpy(dest, src, size)
		return g.ctx.Builder.CreateMemCpy(args[0], args[1], args[2])
	}
	if ctx.MEMMOVE() != nil && len(args) == 3 {
		// memmove(dest, src, size)
		return g.ctx.Builder.CreateMemMove(args[0], args[1], args[2])
	}
	
	// 4. String Intrinsics
	if ctx.STRLEN() != nil && len(args) == 1 {
		return g.ctx.Builder.CreateStrLen(args[0], "")
	}

	// 5. Variadic Arguments (va_start, va_arg, va_end)
	if ctx.VA_START() != nil {
		vaListType := g.resolveType(ctx.Type_())
		if vaListType == types.Void || vaListType == types.I64 {
			vaListType = types.NewPointer(types.I8)
		}
		
		vaListPtr := g.ctx.Builder.CreateAlloca(vaListType, "va_list")
		i8PtrType := types.NewPointer(types.I8)
		vaStartArg := g.ctx.Builder.CreateBitCast(vaListPtr, i8PtrType, "")
		
		g.ctx.Builder.CreateVaStart(vaStartArg)
		return vaListPtr
	}
	if ctx.VA_ARG() != nil && len(args) == 1 {
		targetType := g.resolveType(ctx.Type_())
		return g.ctx.Builder.CreateVaArg(args[0], targetType, "")
	}
	if ctx.VA_END() != nil {
		return g.getZeroValue(types.I64)
	}

	// 6. Utility
	if ctx.BIT_CAST() != nil && len(args) == 1 {
		targetType := g.resolveType(ctx.Type_())
		return g.ctx.Builder.CreateBitCast(args[0], targetType, "")
	}
	
	// 7. Error Handling / Runtime
	if ctx.RAISE() != nil && len(args) == 1 {
		g.ctx.Builder.CreateRaise(args[0])
		return g.getZeroValue(types.I64)
	}

	return g.getZeroValue(types.I64)
}