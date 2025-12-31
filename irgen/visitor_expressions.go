package irgen

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

// Helper to get operator token type at a specific index
func getBinaryOp(ctx antlr.ParserRuleContext, index int) int {
	// In a rule like `lhs op rhs`, the operator is at child index 1.
	// The child index is always 2*i + 1 relative to expression list, 
	// but context provides specific accessor methods usually.
	// For generic access:
	if child := ctx.GetChild(2*index + 1); child != nil {
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
		
		var op int
		if term, ok := ctx.GetChild(2*i + 1).(antlr.TerminalNode); ok {
			op = term.GetSymbol().GetTokenType()
		}

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
		
		var op int
		if term, ok := ctx.GetChild(2*i + 1).(antlr.TerminalNode); ok {
			op = term.GetSymbol().GetTokenType()
		}

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
		
		var op int
		if term, ok := ctx.GetChild(2*i + 1).(antlr.TerminalNode); ok {
			op = term.GetSymbol().GetTokenType()
		}

		if op == parser.ArcParserLT { // <<
			lhs = g.ctx.Builder.CreateShl(lhs, rhs, "")
		} else { // >>
			lhs = g.ctx.Builder.CreateAShr(lhs, rhs, "")
		}
	}
	return lhs
}

func (g *Generator) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} {
	// Range parsing is often handled by parent statements (loops), 
	// but if used as an expression, it visits the first operand.
	return g.Visit(ctx.AdditiveExpression(0))
}

func (g *Generator) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	lhs := g.Visit(ctx.MultiplicativeExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		rhs := g.Visit(ctx.MultiplicativeExpression(i)).(ir.Value)
		
		var op int
		if term, ok := ctx.GetChild(2*i + 1).(antlr.TerminalNode); ok {
			op = term.GetSymbol().GetTokenType()
		}

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
		
		var op int
		if term, ok := ctx.GetChild(2*i + 1).(antlr.TerminalNode); ok {
			op = term.GetSymbol().GetTokenType()
		}

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
	// 1. Recursive Unary Operators (prefix)
	if ctx.AMP() != nil {
		// Address-of (&x): Get L-Value address directly
		lval := g.getLValue(ctx.UnaryExpression())
		if lval != nil {
			return lval
		}
		fmt.Println("[IRGen] Error: Cannot take address of non-lvalue")
		return g.getZeroValue(types.I64)
	}

	if ctx.STAR() != nil {
		// Dereference (*ptr): Load value from pointer
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		if ptrType, ok := val.Type().(*types.PointerType); ok {
			return g.ctx.Builder.CreateLoad(ptrType.ElementType, val, "")
		}
		// Error case fallback
		return val
	}

	if ctx.MINUS() != nil {
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		if types.IsFloat(val.Type()) {
			return g.ctx.Builder.CreateFSub(g.getZeroValue(val.Type()), val, "")
		}
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

	// 2. Pre-increment/decrement (++x, --x)
	if ctx.INCREMENT() != nil || ctx.DECREMENT() != nil {
		ptr := g.getLValue(ctx.UnaryExpression())
		if ptr == nil {
			return g.getZeroValue(types.I64)
		}
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

	// 3. Fallthrough to PostfixExpression (e.g. literals, identifiers)
	if ctx.PostfixExpression() != nil {
		return g.Visit(ctx.PostfixExpression())
	}

	return g.getZeroValue(types.I64)
}

// --- Postfix Expressions ---

func (g *Generator) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	var currPtr ir.Value = g.getLValue(ctx.PrimaryExpression())
	var curr ir.Value

	// Determine initial R-Value
	if currPtr != nil {
		// If it's a pointer to memory (variable), load it.
		// BUT if it's a Function pointer, using it as a value doesn't require a load 
		// (unless it's a function pointer stored in a variable).
		if _, isFn := currPtr.(*ir.Function); isFn {
			curr = currPtr
			currPtr = nil // Functions are not assignable L-Values
		} else if ptrType, ok := currPtr.Type().(*types.PointerType); ok {
			curr = g.ctx.Builder.CreateLoad(ptrType.ElementType, currPtr, "")
		} else {
			// Register value
			curr = currPtr
		}
	} else {
		// Literal or temporary
		curr = g.Visit(ctx.PrimaryExpression()).(ir.Value)
	}

	for _, op := range ctx.AllPostfixOp() {
		// Function Call: f(...)
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
				currPtr = nil // Result is an R-Value
			} else {
				fmt.Printf("[IRGen] Warning: Attempted call on non-function value: %v\n", curr)
			}
			continue
		}

		// Member Access: .field
		if op.DOT() != nil {
			fieldName := op.IDENTIFIER().GetText()
			
			// Auto-dereference pointer to struct
			var structPtr ir.Value
			var structType *types.StructType

			// Case 1: curr is a pointer to struct (standard for allocas/params)
			if ptr, ok := curr.Type().(*types.PointerType); ok {
				if st, ok := ptr.ElementType.(*types.StructType); ok {
					structPtr = curr
					structType = st
				}
			} else if st, ok := curr.Type().(*types.StructType); ok {
				// Case 2: curr is a struct value (register)
				// We can't GEP a register value, so we extract
				if idx, ok := g.analysis.StructIndices[st.Name][fieldName]; ok {
					curr = g.ctx.Builder.CreateExtractValue(curr, []int{idx}, "")
					currPtr = nil
					continue
				}
			}

			if structType != nil {
				// Method Call Check (struct.Method)
				methodName := structType.Name + "_" + fieldName
				if fn := g.ctx.Module.GetFunction(methodName); fn != nil {
					curr = fn
					// Implicit 'self' logic would go here in a more advanced compiler,
					// injecting 'structPtr' into the arg list of the next call.
					continue
				}

				// Field Access
				if idx, ok := g.analysis.StructIndices[structType.Name][fieldName]; ok {
					gep := g.ctx.Builder.CreateStructGEP(structType, structPtr, idx, "")
					curr = g.ctx.Builder.CreateLoad(structType.Fields[idx], gep, "")
					currPtr = gep // Update L-Value pointer
				}
			}
			continue
		}

		// Indexing: [expr]
		if op.LBRACKET() != nil {
			idx := g.Visit(op.Expression()).(ir.Value)
			
			if ptr, ok := curr.Type().(*types.PointerType); ok {
				if _, isArray := ptr.ElementType.(*types.ArrayType); isArray {
					// Decay array pointer [N x T]* -> T*
					zero := g.ctx.Builder.ConstInt(types.I32, 0)
					currPtr = g.ctx.Builder.CreateInBoundsGEP(ptr.ElementType, curr, []ir.Value{zero, idx}, "")
				} else {
					// Standard pointer offset T* -> T*
					currPtr = g.ctx.Builder.CreateInBoundsGEP(ptr.ElementType, curr, []ir.Value{idx}, "")
				}
				// Load element
				curr = g.ctx.Builder.CreateLoad(currPtr.Type().(*types.PointerType).ElementType, currPtr, "")
			}
			continue
		}

		// Post-Increment/Decrement: x++, x--
		if op.INCREMENT() != nil || op.DECREMENT() != nil {
			if currPtr == nil {
				continue // Cannot mutate R-Value
			}
			
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
			currPtr = nil // Result matches original value (old value), is now R-Value
		}
	}
	return curr
}

// --- Primary Expressions ---

func (g *Generator) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.StructLiteral() != nil { return g.Visit(ctx.StructLiteral()) }
	if ctx.Literal() != nil { return g.Visit(ctx.Literal()) }
	if ctx.CastExpression() != nil { return g.Visit(ctx.CastExpression()) }
	if ctx.AllocaExpression() != nil { return g.Visit(ctx.AllocaExpression()) }
	if ctx.SyscallExpression() != nil { return g.Visit(ctx.SyscallExpression()) }
	if ctx.IntrinsicExpression() != nil { return g.Visit(ctx.IntrinsicExpression()) }

	// Qualified IDs: Namespace.Var or Enum.Member
	if ctx.QualifiedIdentifier() != nil {
		q := ctx.QualifiedIdentifier()
		ids := q.AllIDENTIFIER()
		var name string
		
		if q.SYSCALL() != nil { name = "syscall" }
		
		for i, id := range ids {
			if i > 0 || name != "" { name += "." }
			name += id.GetText()
		}
		
		// 1. Try Scope Resolve
		if sym, ok := g.currentScope.Resolve(name); ok && sym.IRValue != nil {
			// If it's a variable (Alloca), load it. If it's a global/func, return it.
			if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
				return g.ctx.Builder.CreateLoad(sym.Type, alloca, "")
			}
			return sym.IRValue
		}
		
		// 2. Try Function lookup
		if fn := g.ctx.Module.GetFunction(name); fn != nil {
			return fn
		}
		
		return g.getZeroValue(types.I64)
	}

	// Simple Identifier
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		
		if sym, ok := g.currentScope.Resolve(name); ok {
			if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
				return g.ctx.Builder.CreateLoad(sym.Type, alloca, "")
			}
			if sym.IRValue != nil {
				return sym.IRValue
			}
		}
		
		// Global Function?
		if fn := g.ctx.Module.GetFunction(name); fn != nil {
			return fn
		}
		
		// Global Variable?
		if glob := g.ctx.Module.GetGlobal(name); glob != nil {
			return g.ctx.Builder.CreateLoad(glob.Type().(*types.PointerType).ElementType, glob, "")
		}
	}

	// Parenthesized Expression
	if ctx.Expression() != nil {
		return g.Visit(ctx.Expression())
	}

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
			// Basic escape handling
			if txt[1] == '\\' && len(txt) > 3 {
				switch txt[2] {
				case 'n': r = '\n'
				case 't': r = '\t'
				case 'r': r = '\r'
				case '0': r = 0
				case '\\': r = '\\'
				case '\'': r = '\''
				}
			}
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
	if !ok {
		return g.getZeroValue(types.I64)
	}
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
	srcType := val.Type()

	if types.IsPointer(srcType) && types.IsPointer(targetType) {
		return g.ctx.Builder.CreateBitCast(val, targetType, "")
	}
	if types.IsPointer(srcType) && types.IsInteger(targetType) {
		return g.ctx.Builder.CreatePtrToInt(val, targetType, "")
	}
	if types.IsInteger(srcType) && types.IsPointer(targetType) {
		return g.ctx.Builder.CreateIntToPtr(val, targetType, "")
	}

	return g.emitCast(val, targetType)
}

func (g *Generator) VisitAllocaExpression(ctx *parser.AllocaExpressionContext) interface{} {
	typ := g.resolveType(ctx.Type_())
	if ctx.Expression() != nil {
		count := g.Visit(ctx.Expression()).(ir.Value)
		return g.ctx.Builder.CreateAllocaWithCount(typ, count, "")
	}
	return g.ctx.Builder.CreateAlloca(typ, "")
}

func (g *Generator) VisitSyscallExpression(ctx *parser.SyscallExpressionContext) interface{} {
	var args []ir.Value
	for _, expr := range ctx.AllExpression() {
		args = append(args, g.Visit(expr).(ir.Value))
	}
	return g.ctx.Builder.CreateSyscall(args)
}

func (g *Generator) VisitIntrinsicExpression(ctx *parser.IntrinsicExpressionContext) interface{} {
	if ctx.SIZEOF() != nil {
		t := g.resolveType(ctx.Type_())
		return g.ctx.Builder.CreateSizeOf(t, "")
	}
	if ctx.ALIGNOF() != nil {
		t := g.resolveType(ctx.Type_())
		return g.ctx.Builder.CreateAlignOf(t, "")
	}
	if ctx.BIT_CAST() != nil {
		val := g.Visit(ctx.Expression(0)).(ir.Value)
		target := g.resolveType(ctx.Type_())
		return g.ctx.Builder.CreateBitCast(val, target, "")
	}

	var args []ir.Value
	for _, expr := range ctx.AllExpression() {
		args = append(args, g.Visit(expr).(ir.Value))
	}

	if ctx.MEMSET() != nil && len(args) == 3 {
		return g.ctx.Builder.CreateMemSet(args[0], args[1], args[2])
	}
	if ctx.MEMCPY() != nil && len(args) == 3 {
		return g.ctx.Builder.CreateMemCpy(args[0], args[1], args[2])
	}
	if ctx.MEMMOVE() != nil && len(args) == 3 {
		return g.ctx.Builder.CreateMemMove(args[0], args[1], args[2])
	}
	if ctx.STRLEN() != nil && len(args) == 1 {
		return g.ctx.Builder.CreateStrLen(args[0], "")
	}

	if ctx.VA_START() != nil {
		vaList := g.ctx.Builder.CreateAlloca(types.NewPointer(types.I8), "va_list")
		vaListArg := g.ctx.Builder.CreateBitCast(vaList, types.NewPointer(types.I8), "")
		g.ctx.Builder.CreateVaStart(vaListArg)
		return vaList
	}
	if ctx.VA_ARG() != nil && len(args) == 1 {
		target := g.resolveType(ctx.Type_())
		return g.ctx.Builder.CreateVaArg(args[0], target, "")
	}
	if ctx.VA_END() != nil && len(args) == 1 {
		g.ctx.Builder.CreateVaEnd(args[0])
		return g.getZeroValue(types.Void)
	}

	if ctx.RAISE() != nil && len(args) == 1 {
		g.ctx.Builder.CreateRaise(args[0])
	}

	return g.getZeroValue(types.I64)
}