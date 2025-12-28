package compiler

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (v *IRVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return v.Visit(ctx.LogicalOrExpression())
}

func (v *IRVisitor) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	result := v.Visit(ctx.LogicalAndExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllLogicalAndExpression()); i++ {
		rhs := v.Visit(ctx.LogicalAndExpression(i)).(ir.Value)
		result = v.ctx.Builder.CreateOr(result, rhs, "")
	}
	return result
}

func (v *IRVisitor) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	result := v.Visit(ctx.BitOrExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllBitOrExpression()); i++ {
		rhs := v.Visit(ctx.BitOrExpression(i)).(ir.Value)
		result = v.ctx.Builder.CreateAnd(result, rhs, "")
	}
	return result
}

func (v *IRVisitor) VisitBitOrExpression(ctx *parser.BitOrExpressionContext) interface{} {
	result := v.Visit(ctx.BitXorExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllBitXorExpression()); i++ {
		rhs := v.Visit(ctx.BitXorExpression(i)).(ir.Value)
		result = v.ctx.Builder.CreateOr(result, rhs, "")
	}
	return result
}

func (v *IRVisitor) VisitBitXorExpression(ctx *parser.BitXorExpressionContext) interface{} {
	result := v.Visit(ctx.BitAndExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllBitAndExpression()); i++ {
		rhs := v.Visit(ctx.BitAndExpression(i)).(ir.Value)
		result = v.ctx.Builder.CreateXor(result, rhs, "")
	}
	return result
}

func (v *IRVisitor) VisitBitAndExpression(ctx *parser.BitAndExpressionContext) interface{} {
	result := v.Visit(ctx.EqualityExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllEqualityExpression()); i++ {
		rhs := v.Visit(ctx.EqualityExpression(i)).(ir.Value)
		result = v.ctx.Builder.CreateAnd(result, rhs, "")
	}
	return result
}

func (v *IRVisitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	result := v.Visit(ctx.RelationalExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllRelationalExpression()); i++ {
		rhs := v.Visit(ctx.RelationalExpression(i)).(ir.Value)
		if i-1 < len(ctx.AllEQ()) {
			result = v.ctx.Builder.CreateICmpEQ(result, rhs, "")
		} else {
			result = v.ctx.Builder.CreateICmpNE(result, rhs, "")
		}
	}
	return result
}

func (v *IRVisitor) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} {
	result := v.Visit(ctx.ShiftExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllShiftExpression()); i++ {
		rhs := v.Visit(ctx.ShiftExpression(i)).(ir.Value)
		if i-1 < len(ctx.AllLT()) {
			result = v.ctx.Builder.CreateICmpSLT(result, rhs, "")
		} else if i-1-len(ctx.AllLT()) < len(ctx.AllLE()) {
			result = v.ctx.Builder.CreateICmpSLE(result, rhs, "")
		} else if i-1-len(ctx.AllLT())-len(ctx.AllLE()) < len(ctx.AllGT()) {
			result = v.ctx.Builder.CreateICmpSGT(result, rhs, "")
		} else {
			result = v.ctx.Builder.CreateICmpSGE(result, rhs, "")
		}
	}
	return result
}

func (v *IRVisitor) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} {
	result := v.Visit(ctx.RangeExpression(0)).(ir.Value)
	
	count := len(ctx.AllRangeExpression())
	for i := 1; i < count; i++ {
		rhs := v.Visit(ctx.RangeExpression(i)).(ir.Value)
		
		opNode := ctx.GetChild(i*2 - 1).(antlr.TerminalNode)
		tokenType := opNode.GetSymbol().GetTokenType()
		
		if tokenType == parser.ArcLexerLSHIFT {
			result = v.ctx.Builder.CreateShl(result, rhs, "")
		} else {
			if intType, ok := result.Type().(*types.IntType); ok && intType.Signed {
				result = v.ctx.Builder.CreateAShr(result, rhs, "")
			} else {
				result = v.ctx.Builder.CreateLShr(result, rhs, "")
			}
		}
	}
	return result
}

func (v *IRVisitor) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} {
	return v.Visit(ctx.AdditiveExpression(0))
}

func (v *IRVisitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	result := v.Visit(ctx.MultiplicativeExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		rhs := v.Visit(ctx.MultiplicativeExpression(i)).(ir.Value)
		if i-1 < len(ctx.AllPLUS()) {
			result = v.ctx.Builder.CreateAdd(result, rhs, "")
		} else {
			result = v.ctx.Builder.CreateSub(result, rhs, "")
		}
	}
	return result
}

func (v *IRVisitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	result := v.Visit(ctx.UnaryExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllUnaryExpression()); i++ {
		rhs := v.Visit(ctx.UnaryExpression(i)).(ir.Value)
		if i-1 < len(ctx.AllSTAR()) {
			result = v.ctx.Builder.CreateMul(result, rhs, "")
		} else if i-1-len(ctx.AllSTAR()) < len(ctx.AllSLASH()) {
			result = v.ctx.Builder.CreateSDiv(result, rhs, "")
		} else {
			result = v.ctx.Builder.CreateSRem(result, rhs, "")
		}
	}
	return result
}

func (v *IRVisitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if ctx.AWAIT() != nil {
		val := v.Visit(ctx.UnaryExpression()).(ir.Value)
		if val == nil { return v.ctx.Builder.ConstInt(types.I64, 0) }
		_ = v.ctx.Builder.CreateCoroSuspend(false, "")
		return val
	}
	if ctx.MINUS() != nil {
		val := v.Visit(ctx.UnaryExpression()).(ir.Value)
		if val == nil { return v.ctx.Builder.ConstInt(types.I64, 0) }
		zero := v.getZeroValue(val.Type())
		return v.ctx.Builder.CreateSub(zero, val, "")
	}
	if ctx.NOT() != nil {
		val := v.Visit(ctx.UnaryExpression()).(ir.Value)
		if val == nil { return v.ctx.Builder.ConstInt(types.I64, 0) }
		return v.ctx.Builder.CreateXor(val, v.ctx.Builder.ConstInt(types.I1, 1), "")
	}
	if ctx.BIT_NOT() != nil {
		val := v.Visit(ctx.UnaryExpression()).(ir.Value)
		if val == nil { return v.ctx.Builder.ConstInt(types.I64, 0) }
		var allOnes ir.Constant
		if intType, ok := val.Type().(*types.IntType); ok {
			allOnes = v.ctx.Builder.ConstInt(intType, -1)
		} else {
			allOnes = v.ctx.Builder.ConstInt(types.I64, -1)
		}
		return v.ctx.Builder.CreateXor(val, allOnes, "")
	}
	if ctx.STAR() != nil {
		ptr := v.Visit(ctx.UnaryExpression()).(ir.Value)
		if ptr == nil { return v.ctx.Builder.ConstInt(types.I64, 0) }
		ptrType, ok := ptr.Type().(*types.PointerType)
		if !ok {
			v.ctx.Logger.Error("Cannot dereference non-pointer type: %v", ptr.Type())
			return ptr
		}
		return v.ctx.Builder.CreateLoad(ptrType.ElementType, ptr, "")
	}
	if ctx.AMP() != nil {
		if childCtx, ok := ctx.UnaryExpression().(*parser.UnaryExpressionContext); ok {
			return v.getExpressionAddress(childCtx)
		}
		return v.ctx.Builder.ConstInt(types.I64, 0)
	}
	
	if ctx.PostfixExpression() != nil {
		return v.Visit(ctx.PostfixExpression())
	}
	
	return v.ctx.Builder.ConstInt(types.I64, 0)
}

func (v *IRVisitor) getExpressionAddress(ctx *parser.UnaryExpressionContext) ir.Value {
	if post := ctx.PostfixExpression(); post != nil {
		if len(post.AllPostfixOp()) == 0 {
			if prim := post.PrimaryExpression(); prim != nil {
				if prim.IDENTIFIER() != nil {
					name := prim.IDENTIFIER().GetText()
					if sym, ok := v.ctx.currentScope.Lookup(name); ok {
						return sym.Value
					}
				}
			}
		}
	}
	v.ctx.Logger.Error("Invalid operand for address-of operator")
	return v.ctx.Builder.ConstInt(types.I64, 0)
}

func (v *IRVisitor) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	result := v.Visit(ctx.PrimaryExpression()).(ir.Value)
	
	var baseIdentifier string
	if primaryCtx := ctx.PrimaryExpression(); primaryCtx != nil {
		if primaryCtx.IDENTIFIER() != nil {
			baseIdentifier = primaryCtx.IDENTIFIER().GetText()
		}
	}
	
	for _, op := range ctx.AllPostfixOp() {
		result = v.visitPostfixOp(result, op.(*parser.PostfixOpContext), baseIdentifier)
		baseIdentifier = ""
	}
	return result
}

func (v *IRVisitor) visitPostfixOp(base ir.Value, ctx *parser.PostfixOpContext, baseIdentifier string) ir.Value {
	if ctx.LPAREN() != nil {
		var args []ir.Value
		if ctx.ArgumentList() != nil {
			argResult := v.Visit(ctx.ArgumentList())
			if argResult != nil {
				args = argResult.([]ir.Value)
			} else {
				args = []ir.Value{}
			}
		}
		
		if fn, ok := base.(*ir.Function); ok {
			if v.pendingMethodSelf != nil {
				args = append([]ir.Value{v.pendingMethodSelf}, args...)
				v.pendingMethodSelf = nil
			}
			return v.ctx.Builder.CreateCall(fn, args, "")
		}
		v.ctx.Logger.Error("Cannot call non-function")
		return base
	}
	
	if ctx.LBRACKET() != nil {
		if ctx.Expression() == nil { return base }
		indexExpr := v.Visit(ctx.Expression()).(ir.Value)
		
		if ptrType, ok := base.Type().(*types.PointerType); ok {
			elemPtr := v.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, base, []ir.Value{indexExpr}, "")
			return v.ctx.Builder.CreateLoad(ptrType.ElementType, elemPtr, "")
		}
		v.ctx.Logger.Error("Cannot index non-pointer type: %v", base.Type())
		return base
	}
	
	if ctx.DOT() != nil && ctx.IDENTIFIER() != nil {
		memberName := ctx.IDENTIFIER().GetText()
		v.pendingMethodSelf = nil
		
		if baseIdentifier != "" {
			if ns, ok := v.ctx.NamespaceRegistry[baseIdentifier]; ok {
				if fn, ok := ns.LookupFunction(memberName); ok {
					return fn
				}
			}
		}
		
		return v.handleMemberAccess(base, memberName)
	}
	
	return base
}

func (v *IRVisitor) handleMemberAccess(base ir.Value, memberName string) ir.Value {
	if ptrType, ok := base.Type().(*types.PointerType); ok {
		if structType, ok := ptrType.ElementType.(*types.StructType); ok {
			
			if v.ctx.IsClassType(structType.Name) {
				methodName := structType.Name + "_" + memberName
				if fn := v.ctx.Module.GetFunction(methodName); fn != nil {
					v.pendingMethodSelf = base
					return fn
				}
			}
			
			var fieldIdx int = -1
			if v.ctx.IsClassType(structType.Name) {
				if idx, ok := v.ctx.ClassFieldIndices[structType.Name][memberName]; ok {
					fieldIdx = idx
				}
			} else {
				fieldIdx = v.findFieldIndex(structType, memberName)
			}
			
			if fieldIdx >= 0 {
				gep := v.ctx.Builder.CreateStructGEP(structType, base, fieldIdx, "")
				return v.ctx.Builder.CreateLoad(structType.Fields[fieldIdx], gep, "")
			}
		}
	}
	
	if structType, ok := base.Type().(*types.StructType); ok {
		if v.ctx.IsClassType(structType.Name) {
			v.ctx.Logger.Error("Class instances must be accessed via pointer")
			return base
		}
		fieldIdx := v.findFieldIndex(structType, memberName)
		if fieldIdx >= 0 {
			return v.ctx.Builder.CreateExtractValue(base, []int{fieldIdx}, "")
		}
	}
	
	v.ctx.Logger.Error("Type '%v' has no member '%s'", base.Type(), memberName)
	return base
}

func (v *IRVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.StructLiteral() != nil { return v.Visit(ctx.StructLiteral()) }
	if ctx.Literal() != nil { return v.Visit(ctx.Literal()) }
	if ctx.Expression() != nil { return v.Visit(ctx.Expression()) }
	if ctx.CastExpression() != nil { return v.Visit(ctx.CastExpression()) }
	if ctx.AllocaExpression() != nil { return v.Visit(ctx.AllocaExpression()) }
	if ctx.SyscallExpression() != nil { return v.Visit(ctx.SyscallExpression()) }
	if ctx.IntrinsicExpression() != nil { return v.Visit(ctx.IntrinsicExpression()) }
	
	if ctx.QualifiedIdentifier() != nil {
		qCtx := ctx.QualifiedIdentifier()
		parts := make([]string, len(qCtx.AllIDENTIFIER()))
		for i, node := range qCtx.AllIDENTIFIER() {
			parts[i] = node.GetText()
		}
		
		if len(parts) < 2 { return v.ctx.Builder.ConstInt(types.I64, 0) }
		firstPart := parts[0]

		if sym, ok := v.ctx.currentScope.Lookup(firstPart); ok {
			var currentVal ir.Value = sym.Value
			
			if ptr, isAlloca := currentVal.(*ir.AllocaInst); isAlloca {
				if _, isPtr := ptr.AllocatedType.(*types.PointerType); isPtr {
					currentVal = v.ctx.Builder.CreateLoad(ptr.AllocatedType, ptr, "")
				} else {
					currentVal = ptr
				}
			}
			
			for i := 1; i < len(parts); i++ {
				currentVal = v.handleMemberAccess(currentVal, parts[i])
			}
			return currentVal
		}

		nsName := firstPart
		memberName := parts[1]
		
		if ns, ok := v.ctx.NamespaceRegistry[nsName]; ok {
			if fn, ok := ns.LookupFunction(memberName); ok {
				return fn
			}
			v.ctx.Logger.Error("Function '%s' not found in namespace '%s'", memberName, nsName)
		} else {
			v.ctx.Logger.Error("Unknown namespace or variable: %s", nsName)
		}
		return v.ctx.Builder.ConstInt(types.I64, 0)
	}

	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		
		if _, isType := v.ctx.GetType(name); isType {
			v.ctx.Logger.Error("Type '%s' used as value", name)
			return v.ctx.Builder.ConstInt(types.I64, 0)
		}
		
		if _, isNamespace := v.ctx.NamespaceRegistry[name]; isNamespace {
			return v.ctx.Builder.ConstInt(types.I64, 0)
		}
		
		sym, ok := v.ctx.currentScope.Lookup(name)
		if !ok {
			if v.ctx.currentNamespace != nil {
				if fn, ok := v.ctx.currentNamespace.Functions[name]; ok {
					return fn
				}
			}
			if fn := v.ctx.Module.GetFunction(name); fn != nil {
				return fn
			}
			v.ctx.Logger.Error("Undefined: %s", name)
			return v.ctx.Builder.ConstInt(types.I64, 0)
		}

		if ptr, isAlloca := sym.Value.(*ir.AllocaInst); isAlloca {
			ptrType := ptr.Type().(*types.PointerType)
			return v.ctx.Builder.CreateLoad(ptrType.ElementType, ptr, "")
		}

		return sym.Value
	}
	
	return v.ctx.Builder.ConstInt(types.I64, 0)
}

func (v *IRVisitor) VisitStructLiteral(ctx *parser.StructLiteralContext) interface{} {
	var name string
	var structType *types.StructType
	
	// 1. Determine the name and resolve the type
	if ctx.IDENTIFIER() != nil {
		name = ctx.IDENTIFIER().GetText()
		// Try current namespace first
		if v.ctx.currentNamespace != nil && v.ctx.currentNamespace.Name != "" {
			nsName := v.ctx.currentNamespace.Name + "_" + name
			if t, ok := v.ctx.GetType(nsName); ok {
				name = nsName
				structType = t.(*types.StructType)
			}
		}
		// Try raw name if namespace failed
		if structType == nil {
			if t, ok := v.ctx.GetType(name); ok {
				structType = t.(*types.StructType)
			}
		}
	} else if ctx.QualifiedIdentifier() != nil {
		qCtx := ctx.QualifiedIdentifier()
		if len(qCtx.AllIDENTIFIER()) == 2 {
			nsName := qCtx.IDENTIFIER(0).GetText()
			typeName := qCtx.IDENTIFIER(1).GetText()
			
			// Resolve via namespace registry
			if ns, ok := v.ctx.NamespaceRegistry[nsName]; ok {
				if typ, ok := ns.Types[typeName]; ok {
					if st, ok := typ.(*types.StructType); ok {
						structType = st
						name = st.Name // Use the internal IR name (e.g., "net_Socket")
					}
				}
			}
		}
	}

	if structType == nil {
		v.ctx.Logger.Error("Unknown struct/class type: %s", ctx.GetText())
		return v.ctx.Builder.ConstInt(types.I64, 0)
	}

	// Use the resolved IR name
	irName := structType.Name
	v.logger.Debug("Creating struct literal for type: %s", irName)

	if v.ctx.IsClassType(irName) {
		ptrToClass := v.ctx.Builder.CreateAlloca(structType, irName+".instance")
		
		// Zero-initialize
		for i := 0; i < len(structType.Fields); i++ {
			gep := v.ctx.Builder.CreateStructGEP(structType, ptrToClass, i, "")
			zero := v.getZeroValue(structType.Fields[i])
			v.ctx.Builder.CreateStore(zero, gep)
		}
		
		// Initialize fields
		for _, field := range ctx.AllFieldInit() {
			fieldName := field.IDENTIFIER().GetText()
			fieldVal := v.Visit(field.Expression()).(ir.Value)
			
			var idx int = -1
			if fieldIndices, ok := v.ctx.ClassFieldIndices[irName]; ok {
				if fieldIdx, ok := fieldIndices[fieldName]; ok {
					idx = fieldIdx
				}
			}
			
			if idx < 0 {
				v.ctx.Logger.Error("Class %s has no field %s", irName, fieldName)
				continue
			}
			
			gep := v.ctx.Builder.CreateStructGEP(structType, ptrToClass, idx, "")
			v.ctx.Builder.CreateStore(fieldVal, gep)
		}
		
		return ptrToClass
	}

	// Regular struct (Value type)
	var agg ir.Value = v.ctx.Builder.ConstZero(structType)

	for _, field := range ctx.AllFieldInit() {
		fieldName := field.IDENTIFIER().GetText()
		fieldVal := v.Visit(field.Expression()).(ir.Value)
		
		var idx int = -1
		if fieldIndices, ok := v.ctx.StructFieldIndices[irName]; ok {
			if fieldIdx, ok := fieldIndices[fieldName]; ok {
				idx = fieldIdx
			}
		}
		
		if idx < 0 {
			v.ctx.Logger.Error("Struct %s has no field %s", irName, fieldName)
			continue
		}
		
		agg = v.ctx.Builder.CreateInsertValue(agg, fieldVal, []int{idx}, "")
	}
	
	return agg
}

func (v *IRVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
    if ctx.INTEGER_LITERAL() != nil {
        text := ctx.INTEGER_LITERAL().GetText()
        val, _ := strconv.ParseInt(text, 0, 64)
        return v.ctx.Builder.ConstInt(types.I64, val)
    }
    if ctx.FLOAT_LITERAL() != nil {
        text := ctx.FLOAT_LITERAL().GetText()
        val, _ := strconv.ParseFloat(text, 64)
        return v.ctx.Builder.ConstFloat(types.F64, val)
    }
    if ctx.BOOLEAN_LITERAL() != nil {
        if ctx.BOOLEAN_LITERAL().GetText() == "true" {
            return v.ctx.Builder.True()
        }
        return v.ctx.Builder.False()
    }
    if ctx.STRING_LITERAL() != nil {
        rawText := ctx.STRING_LITERAL().GetText()
        content, err := strconv.Unquote(rawText)
        if err != nil {
            if len(rawText) >= 2 {
                content = rawText[1 : len(rawText)-1]
            } else {
                content = rawText
            }
        }
        
        bytes := append([]byte(content), 0)
        elements := make([]ir.Constant, len(bytes))
        for i, b := range bytes {
            elements[i] = v.ctx.Builder.ConstInt(types.I8, int64(b))
        }
        
        arrType := types.NewArray(types.I8, int64(len(bytes)))
        constArr := &ir.ConstantArray{
            BaseValue: ir.BaseValue{ValType: arrType},
            Elements:  elements,
        }
        
        strName := fmt.Sprintf(".str.%d", len(v.ctx.Module.Globals))
        global := v.ctx.Builder.CreateGlobalConstant(strName, constArr)
        zero := v.ctx.Builder.ConstInt(types.I32, 0)
        
        return v.ctx.Builder.CreateInBoundsGEP(arrType, global, []ir.Value{zero, zero}, "")
    }
    if ctx.VectorLiteral() != nil {
        return v.Visit(ctx.VectorLiteral())
    }
    if ctx.MapLiteral() != nil {
        v.ctx.Logger.Warning("Map literals not yet implemented")
        return v.ctx.Builder.ConstInt(types.I64, 0)
    }
    return v.ctx.Builder.ConstInt(types.I64, 0)
}

func (v *IRVisitor) VisitCastExpression(ctx *parser.CastExpressionContext) interface{} {
	val := v.Visit(ctx.Expression()).(ir.Value)
	destType := v.resolveType(ctx.Type_())
	return v.castValue(val, destType)
}

func (v *IRVisitor) VisitAllocaExpression(ctx *parser.AllocaExpressionContext) interface{} {
	allocType := v.resolveType(ctx.Type_())
	if ctx.Expression() != nil {
		count := v.Visit(ctx.Expression()).(ir.Value)
		return v.ctx.Builder.CreateAllocaWithCount(allocType, count, "")
	}
	return v.ctx.Builder.CreateAlloca(allocType, "")
}

func (v *IRVisitor) VisitSyscallExpression(ctx *parser.SyscallExpressionContext) interface{} {
	exprs := ctx.AllExpression()
	if len(exprs) == 0 {
		return v.ctx.Builder.ConstInt(types.I64, -1)
	}
	args := make([]ir.Value, len(exprs))
	for i, expr := range exprs {
		val := v.Visit(expr).(ir.Value)
		if types.IsInteger(val.Type()) {
			if val.Type().BitSize() < 64 {
				val = v.ctx.Builder.CreateSExt(val, types.I64, "")
			}
		}
		args[i] = val
	}
	return v.ctx.Builder.CreateSyscall(args)
}

func (v *IRVisitor) VisitArgumentList(ctx *parser.ArgumentListContext) interface{} {
	var args []ir.Value
	for _, expr := range ctx.AllExpression() {
		if val, ok := v.Visit(expr).(ir.Value); ok {
			args = append(args, val)
		}
	}
	return args
}

func (v *IRVisitor) VisitLeftHandSide(ctx *parser.LeftHandSideContext) interface{} {
	if ctx.IDENTIFIER() != nil && ctx.DOT() == nil && ctx.STAR() == nil {
		name := ctx.IDENTIFIER().GetText()
		if sym, ok := v.ctx.currentScope.Lookup(name); ok {
			return sym.Value
		}
	}
	if ctx.STAR() != nil {
		return v.Visit(ctx.PostfixExpression())
	}
	if ctx.PostfixExpression() != nil {
		return v.Visit(ctx.PostfixExpression())
	}
	return v.ctx.Builder.ConstInt(types.I64, 0)
}

func (v *IRVisitor) VisitVectorLiteral(ctx *parser.VectorLiteralContext) interface{} {
    exprs := ctx.AllExpression()
    if len(exprs) == 0 {
        return &ir.ConstantArray{
            BaseValue: ir.BaseValue{ValType: types.NewArray(types.I32, 0)},
            Elements:  []ir.Constant{},
        }
    }
    
    elements := make([]ir.Constant, len(exprs))
    var elemType types.Type
    
    for i, expr := range exprs {
        val := v.Visit(expr).(ir.Value)
        if i == 0 { elemType = val.Type() }
        
        if constVal, ok := val.(ir.Constant); ok {
            elements[i] = constVal
        } else {
            elements[i] = v.ctx.Builder.ConstInt(types.I32, 0)
            if elemType == nil { elemType = types.I32 }
        }
    }
    
    arrType := types.NewArray(elemType, int64(len(elements)))
    return &ir.ConstantArray{
        BaseValue: ir.BaseValue{ValType: arrType},
        Elements:  elements,
    }
}

func (v *IRVisitor) VisitIntrinsicExpression(ctx *parser.IntrinsicExpressionContext) interface{} {
	if ctx.SIZEOF() != nil {
		typ := v.resolveType(ctx.Type_())
		return v.ctx.Builder.CreateSizeOf(typ, "")
	}
	if ctx.ALIGNOF() != nil {
		typ := v.resolveType(ctx.Type_())
		return v.ctx.Builder.CreateAlignOf(typ, "")
	}
	
	if ctx.VA_START() != nil {
		vaListType := v.ctx.namedTypes["va_list"]
		if vaListType == nil {
			vaListType = types.NewPointer(types.I8)
			v.ctx.namedTypes["va_list"] = vaListType
		}
		vaListPtr := v.ctx.Builder.CreateAlloca(vaListType, "va_list")
		v.ctx.Builder.CreateVaStart(vaListPtr)
		return vaListPtr
	}
	
	var args []ir.Value
	for _, expr := range ctx.AllExpression() {
		if val, ok := v.Visit(expr).(ir.Value); ok {
			args = append(args, val)
		}
	}
	
	if ctx.MEMSET() != nil && len(args) == 3 {
		return v.ctx.Builder.CreateMemSet(args[0], args[1], args[2])
	}
	if ctx.MEMCPY() != nil && len(args) == 3 {
		return v.ctx.Builder.CreateMemCpy(args[0], args[1], args[2])
	}
	if ctx.MEMMOVE() != nil && len(args) == 3 {
		return v.ctx.Builder.CreateMemMove(args[0], args[1], args[2])
	}
	if ctx.STRLEN() != nil && len(args) == 1 {
		return v.ctx.Builder.CreateStrLen(args[0], "")
	}
	if ctx.MEMCHR() != nil && len(args) == 3 {
		return v.ctx.Builder.CreateMemChr(args[0], args[1], args[2], "")
	}
	if ctx.MEMCMP() != nil && len(args) == 3 {
		return v.ctx.Builder.CreateMemCmp(args[0], args[1], args[2], "")
	}
	if ctx.VA_ARG() != nil && len(args) == 1 {
		targetType := v.resolveType(ctx.Type_())
		return v.ctx.Builder.CreateVaArg(args[0], targetType, "")
	}
	if ctx.VA_END() != nil {
		return v.ctx.Builder.CreateVaEnd(args[0])
	}
	if ctx.RAISE() != nil && len(args) == 1 {
		v.ctx.Builder.CreateRaise(args[0])
		v.ctx.Builder.CreateUnreachable()
		return v.ctx.Builder.ConstInt(types.I64, 0)
	}
	
	return v.ctx.Builder.ConstInt(types.I64, 0)
}