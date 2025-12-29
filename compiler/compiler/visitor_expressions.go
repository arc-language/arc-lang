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
	
	for i := 1; i < len(ctx.AllRangeExpression()); i++ {
		rhs := v.Visit(ctx.RangeExpression(i)).(ir.Value)
		
		// Determine operator type by checking tokens
		opIdx := i*2 - 1
		if opIdx < ctx.GetChildCount() {
			child := ctx.GetChild(opIdx)
			if termNode, ok := child.(antlr.TerminalNode); ok {
				tokenType := termNode.GetSymbol().GetTokenType()
				
				if tokenType == parser.ArcParserLT {
					// Check if next token is also LT (for <<)
					if opIdx+1 < ctx.GetChildCount() {
						nextChild := ctx.GetChild(opIdx + 1)
						if nextTerm, ok := nextChild.(antlr.TerminalNode); ok && nextTerm.GetSymbol().GetTokenType() == parser.ArcParserLT {
							// Left shift (<<)
							result = v.ctx.Builder.CreateShl(result, rhs, "")
							continue
						}
					}
				} else if tokenType == parser.ArcParserGT {
					// Check if next token is also GT (for >>)
					if opIdx+1 < ctx.GetChildCount() {
						nextChild := ctx.GetChild(opIdx + 1)
						if nextTerm, ok := nextChild.(antlr.TerminalNode); ok && nextTerm.GetSymbol().GetTokenType() == parser.ArcParserGT {
							// Right shift (>>)
							if intType, ok := result.Type().(*types.IntType); ok && intType.Signed {
								result = v.ctx.Builder.CreateAShr(result, rhs, "")
							} else {
								result = v.ctx.Builder.CreateLShr(result, rhs, "")
							}
							continue
						}
					}
				}
			}
		}
	}
	
	return result
}

func (v *IRVisitor) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} {
	// Just return the first additive expression
	// Range operator (..) is handled in for-in loop context
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
	
	// Pre-Increment (++i)
	if ctx.INCREMENT() != nil {
		if childCtx, ok := ctx.UnaryExpression().(*parser.UnaryExpressionContext); ok {
			ptr := v.getExpressionAddress(childCtx)
			if ptr != nil && types.IsPointer(ptr.Type()) {
				valType := ptr.Type().(*types.PointerType).ElementType
				val := v.ctx.Builder.CreateLoad(valType, ptr, "")
				
				var one ir.Constant
				if intType, ok := valType.(*types.IntType); ok {
					one = v.ctx.Builder.ConstInt(intType, 1)
				} else {
					one = v.ctx.Builder.ConstInt(types.I64, 1)
				}
				newVal := v.ctx.Builder.CreateAdd(val, one, "")
				v.ctx.Builder.CreateStore(newVal, ptr)
				return newVal
			}
		}
		v.ctx.Logger.Error("Invalid operand for pre-increment")
		return v.ctx.Builder.ConstInt(types.I64, 0)
	}
	
	// Pre-Decrement (--i)
	if ctx.DECREMENT() != nil {
		if childCtx, ok := ctx.UnaryExpression().(*parser.UnaryExpressionContext); ok {
			ptr := v.getExpressionAddress(childCtx)
			if ptr != nil && types.IsPointer(ptr.Type()) {
				valType := ptr.Type().(*types.PointerType).ElementType
				val := v.ctx.Builder.CreateLoad(valType, ptr, "")
				
				var one ir.Constant
				if intType, ok := valType.(*types.IntType); ok {
					one = v.ctx.Builder.ConstInt(intType, 1)
				} else {
					one = v.ctx.Builder.ConstInt(types.I64, 1)
				}
				newVal := v.ctx.Builder.CreateSub(val, one, "")
				v.ctx.Builder.CreateStore(newVal, ptr)
				return newVal
			}
		}
		v.ctx.Logger.Error("Invalid operand for pre-decrement")
		return v.ctx.Builder.ConstInt(types.I64, 0)
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
	
	return v.Visit(ctx.PostfixExpression())
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
	v.logger.Debug("Visiting postfix expression")

	// Start with the primary expression
	result := v.Visit(ctx.PrimaryExpression()).(ir.Value)

	// Apply each postfix operation
	for _, opCtx := range ctx.AllPostfixOp() {
		result = v.applyPostfixOp(result, opCtx)
	}

	return result
}

func (v *IRVisitor) applyPostfixOp(base ir.Value, opCtx parser.IPostfixOpContext) ir.Value {
	op := opCtx.(*parser.PostfixOpContext)

	// Member access: base.member
	if op.DOT() != nil && op.IDENTIFIER() != nil {
		memberName := op.IDENTIFIER().GetText()
		v.logger.Debug("Member access: .%s", memberName)

		// Check if this is a method call (has LPAREN)
		if op.LPAREN() != nil {
			// This is a method call: base.method(args)
			return v.handleMethodCall(base, memberName, op)
		}

		// Regular member access
		return v.handleMemberAccess(base, memberName)
	}

	// Function call: base(args)
	if op.LPAREN() != nil && op.DOT() == nil {
		v.logger.Debug("Function call")
		return v.handleFunctionCall(base, op)
	}

	// Array/pointer indexing: base[index]
	if op.LBRACKET() != nil {
		v.logger.Debug("Array indexing")
		indexVal := v.Visit(op.Expression()).(ir.Value)
		return v.handleIndexing(base, indexVal)
	}

	// Post-increment: base++
	if op.INCREMENT() != nil {
		v.logger.Debug("Post-increment")
		return v.handlePostIncrement(base)
	}

	// Post-decrement: base--
	if op.DECREMENT() != nil {
		v.logger.Debug("Post-decrement")
		return v.handlePostDecrement(base)
	}

	v.ctx.Logger.Error("Unknown postfix operation")
	return base
}

func (v *IRVisitor) handleFunctionCall(funcValue ir.Value, op *parser.PostfixOpContext) ir.Value {
	v.logger.Debug("Handling function call")

	// Get the function type
	var funcType *types.FunctionType
	
	// If funcValue is a pointer to a function, dereference it
	if ptrType, ok := funcValue.Type().(*types.PointerType); ok {
		if ft, ok := ptrType.ElementType.(*types.FunctionType); ok {
			funcType = ft
		}
	} else if ft, ok := funcValue.Type().(*types.FunctionType); ok {
		funcType = ft
	}

	// Build arguments with proper type casting
	args := []ir.Value{}
	argList := op.ArgumentList()
	
	if argList != nil {
		allArgs := argList.(*parser.ArgumentListContext).AllArgument()
		
		for i, argCtx := range allArgs {
			// Visit the argument expression
			var argVal ir.Value
			
			argContext := argCtx.(*parser.ArgumentContext)
			if argContext.Expression() != nil {
				argVal = v.Visit(argContext.Expression()).(ir.Value)
			} else if argContext.LambdaExpression() != nil {
				argVal = v.Visit(argContext.LambdaExpression()).(ir.Value)
			} else {
				v.ctx.Logger.Error("Invalid argument type")
				continue
			}
			
			// Cast to expected parameter type if function type is known
			if funcType != nil && i < len(funcType.ParamTypes) {
				expectedType := funcType.ParamTypes[i]
				if !argVal.Type().Equal(expectedType) {
					v.logger.Debug("Casting argument %d from %v to %v", i, argVal.Type(), expectedType)
					argVal = v.castValue(argVal, expectedType)
				}
			}
			
			args = append(args, argVal)
		}
	}

	// Create the call instruction
	result := v.ctx.Builder.CreateCall(funcValue, args, "")
	
	v.logger.Debug("Function call created, return type: %v", result.Type())
	
	return result
}

func (v *IRVisitor) handleMethodCall(base ir.Value, methodName string, op *parser.PostfixOpContext) ir.Value {
	v.logger.Debug("Handling method call: %s", methodName)

	// Get the base type
	baseType := base.Type()
	
	// If base is a pointer, get the element type
	if ptrType, ok := baseType.(*types.PointerType); ok {
		baseType = ptrType.ElementType
	}

	// Look up the method
	structType, ok := baseType.(*types.StructType)
	if !ok {
		v.ctx.Logger.Error("Cannot call method on non-struct type: %v", baseType)
		return base
	}

	// Find the method in the struct's associated methods
	methodFunc := v.ctx.LookupMethod(structType.Name, methodName)
	if methodFunc == nil {
		v.ctx.Logger.Error("Method %s not found on type %s", methodName, structType.Name)
		return base
	}

	// Build arguments: self is the first argument
	args := []ir.Value{base}
	
	argList := op.ArgumentList()
	if argList != nil {
		allArgs := argList.(*parser.ArgumentListContext).AllArgument()
		
		// Get function type for parameter type checking
		funcType := methodFunc.FuncType
		
		for i, argCtx := range allArgs {
			// Visit the argument expression
			var argVal ir.Value
			
			argContext := argCtx.(*parser.ArgumentContext)
			if argContext.Expression() != nil {
				argVal = v.Visit(argContext.Expression()).(ir.Value)
			} else if argContext.LambdaExpression() != nil {
				argVal = v.Visit(argContext.LambdaExpression()).(ir.Value)
			} else {
				v.ctx.Logger.Error("Invalid argument type")
				continue
			}
			
			// Cast to expected parameter type (i+1 because self is param 0)
			paramIndex := i + 1
			if paramIndex < len(funcType.ParamTypes) {
				expectedType := funcType.ParamTypes[paramIndex]
				if !argVal.Type().Equal(expectedType) {
					v.logger.Debug("Casting method argument %d from %v to %v", i, argVal.Type(), expectedType)
					argVal = v.castValue(argVal, expectedType)
				}
			}
			
			args = append(args, argVal)
		}
	}

	// Create the call instruction
	result := v.ctx.Builder.CreateCall(methodFunc, args, "")
	
	return result
}

func (v *IRVisitor) handleMemberAccess(base ir.Value, memberName string) ir.Value {
	v.logger.Debug("Handling member access: .%s", memberName)

	baseType := base.Type()

	// If base is a pointer, we need to load it first or use GEP
	var structVal ir.Value
	var structType *types.StructType

	if ptrType, ok := baseType.(*types.PointerType); ok {
		// Base is a pointer to a struct
		if st, ok := ptrType.ElementType.(*types.StructType); ok {
			structType = st
			structVal = base // Keep as pointer for GEP
		} else {
			v.ctx.Logger.Error("Cannot access member of non-struct pointer type: %v", ptrType.ElementType)
			return base
		}
	} else if st, ok := baseType.(*types.StructType); ok {
		// Base is a struct value - need to get its address
		structType = st
		// For struct values, we need an address to do GEP
		// This might require storing to a temporary and getting its address
		temp := v.ctx.Builder.CreateAlloca(structType, "temp.struct")
		v.ctx.Builder.CreateStore(base, temp)
		structVal = temp
	} else {
		v.ctx.Logger.Error("Cannot access member of non-struct type: %v", baseType)
		return base
	}

	// Find the field index
	fieldIndex := -1
	for i, field := range structType.Fields {
		if structType.FieldNames != nil && i < len(structType.FieldNames) {
			if structType.FieldNames[i] == memberName {
				fieldIndex = i
				break
			}
		}
	}

	if fieldIndex == -1 {
		v.ctx.Logger.Error("Field %s not found in struct %s", memberName, structType.Name)
		return base
	}

	// Create GEP to get pointer to the field
	indices := []ir.Value{
		v.ctx.Builder.ConstInt(types.NewInt(32), 0),
		v.ctx.Builder.ConstInt(types.NewInt(32), int64(fieldIndex)),
	}

	fieldPtr := v.ctx.Builder.CreateGEP(structType, structVal, indices, "")

	// Load the field value
	fieldValue := v.ctx.Builder.CreateLoad(structType.Fields[fieldIndex], fieldPtr, "")

	return fieldValue
}

func (v *IRVisitor) handleIndexing(base ir.Value, index ir.Value) ir.Value {
	v.logger.Debug("Handling array indexing")

	baseType := base.Type()

	// Handle pointer types
	if ptrType, ok := baseType.(*types.PointerType); ok {
		elemType := ptrType.ElementType

		// Create GEP
		indices := []ir.Value{index}
		elemPtr := v.ctx.Builder.CreateGEP(elemType, base, indices, "")

		// Load the element
		return v.ctx.Builder.CreateLoad(elemType, elemPtr, "")
	}

	// Handle array types
	if arrayType, ok := baseType.(*types.ArrayType); ok {
		// Need to get pointer to array first
		arrayPtr := v.ctx.Builder.CreateAlloca(arrayType, "temp.array")
		v.ctx.Builder.CreateStore(base, arrayPtr)

		// Create GEP with two indices: [0][index]
		indices := []ir.Value{
			v.ctx.Builder.ConstInt(types.NewInt(32), 0),
			index,
		}
		elemPtr := v.ctx.Builder.CreateGEP(arrayType, arrayPtr, indices, "")

		// Load the element
		return v.ctx.Builder.CreateLoad(arrayType.ElementType, elemPtr, "")
	}

	v.ctx.Logger.Error("Cannot index non-pointer/non-array type: %v", baseType)
	return base
}

func (v *IRVisitor) handlePostIncrement(base ir.Value) ir.Value {
	v.logger.Debug("Handling post-increment")

	// Load the current value
	currentVal := base
	if ptrType, ok := base.Type().(*types.PointerType); ok {
		currentVal = v.ctx.Builder.CreateLoad(ptrType.ElementType, base, "")
	}

	// Create the increment
	one := v.ctx.Builder.ConstInt(currentVal.Type(), 1)
	newVal := v.ctx.Builder.CreateAdd(currentVal, one, "")

	// Store back
	if ptrType, ok := base.Type().(*types.PointerType); ok {
		v.ctx.Builder.CreateStore(newVal, base)
	}

	// Return the original value (post-increment returns old value)
	return currentVal
}

func (v *IRVisitor) handlePostDecrement(base ir.Value) ir.Value {
	v.logger.Debug("Handling post-decrement")

	// Load the current value
	currentVal := base
	if ptrType, ok := base.Type().(*types.PointerType); ok {
		currentVal = v.ctx.Builder.CreateLoad(ptrType.ElementType, base, "")
	}

	// Create the decrement
	one := v.ctx.Builder.ConstInt(currentVal.Type(), 1)
	newVal := v.ctx.Builder.CreateSub(currentVal, one, "")

	// Store back
	if ptrType, ok := base.Type().(*types.PointerType); ok {
		v.ctx.Builder.CreateStore(newVal, base)
	}

	// Return the original value (post-decrement returns old value)
	return currentVal
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

	// Post-Increment (i++)
	if ctx.INCREMENT() != nil {
		if baseIdentifier != "" {
			if sym, ok := v.ctx.currentScope.Lookup(baseIdentifier); ok {
				ptr := sym.Value
				if !types.IsPointer(ptr.Type()) {
					v.ctx.Logger.Error("Cannot increment non-lvalue")
					return base
				}
				
				var one ir.Constant
				if intType, ok := base.Type().(*types.IntType); ok {
					one = v.ctx.Builder.ConstInt(intType, 1)
				} else {
					one = v.ctx.Builder.ConstInt(types.I64, 1)
				}
				newVal := v.ctx.Builder.CreateAdd(base, one, "")
				v.ctx.Builder.CreateStore(newVal, ptr)
				
				// Post-inc returns OLD value (which is 'base')
				return base
			}
		}
		v.ctx.Logger.Error("Post-increment only supported on simple variables for now")
		return base
	}

	// Post-Decrement (i--)
	if ctx.DECREMENT() != nil {
		if baseIdentifier != "" {
			if sym, ok := v.ctx.currentScope.Lookup(baseIdentifier); ok {
				ptr := sym.Value
				if !types.IsPointer(ptr.Type()) {
					v.ctx.Logger.Error("Cannot decrement non-lvalue")
					return base
				}
				
				var one ir.Constant
				if intType, ok := base.Type().(*types.IntType); ok {
					one = v.ctx.Builder.ConstInt(intType, 1)
				} else {
					one = v.ctx.Builder.ConstInt(types.I64, 1)
				}
				newVal := v.ctx.Builder.CreateSub(base, one, "")
				v.ctx.Builder.CreateStore(newVal, ptr)
				return base
			}
		}
		v.ctx.Logger.Error("Post-decrement only supported on simple variables for now")
		return base
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
			
			mangledName := nsName + "_" + memberName
			if global := v.ctx.Module.GetGlobal(mangledName); global != nil {
				ptrType := global.Type().(*types.PointerType)
				return v.ctx.Builder.CreateLoad(ptrType.ElementType, global, "")
			}
			
			v.ctx.Logger.Error("Member '%s' not found in namespace '%s'", memberName, nsName)
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
				mangled := v.ctx.currentNamespace.Name + "_" + name
				if global := v.ctx.Module.GetGlobal(mangled); global != nil {
					ptrType := global.Type().(*types.PointerType)
					return v.ctx.Builder.CreateLoad(ptrType.ElementType, global, "")
				}
			}
			if global := v.ctx.Module.GetGlobal(name); global != nil {
				ptrType := global.Type().(*types.PointerType)
				return v.ctx.Builder.CreateLoad(ptrType.ElementType, global, "")
			}
			
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

		// Handle global constants stored in scope
		if sym.IsConst {
			// If it's a global pointer, load it
			if global, ok := sym.Value.(*ir.Global); ok {
				ptrType := global.Type().(*types.PointerType)
				return v.ctx.Builder.CreateLoad(ptrType.ElementType, global, "")
			}
			// Otherwise return the constant value directly
			return sym.Value
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
	
	if ctx.IDENTIFIER() != nil {
		name = ctx.IDENTIFIER().GetText()
		if v.ctx.currentNamespace != nil && v.ctx.currentNamespace.Name != "" {
			nsName := v.ctx.currentNamespace.Name + "_" + name
			if t, ok := v.ctx.GetType(nsName); ok {
				name = nsName
				structType = t.(*types.StructType)
			}
		}
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
			
			if ns, ok := v.ctx.NamespaceRegistry[nsName]; ok {
				if typ, ok := ns.Types[typeName]; ok {
					if st, ok := typ.(*types.StructType); ok {
						structType = st
						name = st.Name
					}
				}
			}
		}
	}

	if structType == nil {
		v.ctx.Logger.Error("Unknown struct/class type: %s", ctx.GetText())
		return v.ctx.Builder.ConstInt(types.I64, 0)
	}

	irName := structType.Name
	v.logger.Debug("Creating struct literal for type: %s", irName)

	if v.ctx.IsClassType(irName) {
		ptrToClass := v.ctx.Builder.CreateAlloca(structType, irName+".instance")
		
		for i := 0; i < len(structType.Fields); i++ {
			gep := v.ctx.Builder.CreateStructGEP(structType, ptrToClass, i, "")
			zero := v.getZeroValue(structType.Fields[i])
			v.ctx.Builder.CreateStore(zero, gep)
		}
		
		for _, field := range ctx.AllFieldInit() {
			fieldName := field.IDENTIFIER().GetText()
			fieldVal := v.Visit(field.Expression()).(ir.Value)
			
			var idx int = -1
			if idxVal, ok := v.ctx.ClassFieldIndices[irName][fieldName]; ok {
				idx = idxVal
			}
			
			if idx < 0 {
				v.ctx.Logger.Error("Class %s has no field %s", irName, fieldName)
				continue
			}
			
			// Cast field value to match field type
			expectedType := structType.Fields[idx]
			if !fieldVal.Type().Equal(expectedType) {
				fieldVal = v.castValue(fieldVal, expectedType)
			}
			
			gep := v.ctx.Builder.CreateStructGEP(structType, ptrToClass, idx, "")
			v.ctx.Builder.CreateStore(fieldVal, gep)
		}
		return ptrToClass
	}

	var agg ir.Value = v.ctx.Builder.ConstZero(structType)
	for _, field := range ctx.AllFieldInit() {
		fieldName := field.IDENTIFIER().GetText()
		fieldVal := v.Visit(field.Expression()).(ir.Value)
		
		var idx int = -1
		if idxVal, ok := v.ctx.StructFieldIndices[irName][fieldName]; ok {
			idx = idxVal
		}
		
		if idx < 0 {
			v.ctx.Logger.Error("Struct %s has no field %s", irName, fieldName)
			continue
		}
		
		// Cast field value to match field type
		expectedType := structType.Fields[idx]
		if !fieldVal.Type().Equal(expectedType) {
			fieldVal = v.castValue(fieldVal, expectedType)
		}
		
		agg = v.ctx.Builder.CreateInsertValue(agg, fieldVal, []int{idx}, "")
	}
	return agg
}

func (v *IRVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTEGER_LITERAL() != nil {
		text := ctx.INTEGER_LITERAL().GetText()
		val, _ := parseInt(text)
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
		return v.visitStringLiteral(ctx.STRING_LITERAL().GetText())
	}
	if ctx.CHAR_LITERAL() != nil {
		text := ctx.CHAR_LITERAL().GetText()
		// Remove single quotes
		if len(text) >= 2 {
			text = text[1 : len(text)-1]
		}
		
		var charValue rune
		if len(text) > 0 && text[0] == '\\' {
			// Escape sequence
			if len(text) >= 2 {
				switch text[1] {
				case 'n':
					charValue = '\n'
				case 't':
					charValue = '\t'
				case 'r':
					charValue = '\r'
				case '\\':
					charValue = '\\'
				case '\'':
					charValue = '\''
				case '0':
					charValue = '\000'
				default:
					charValue = rune(text[1])
				}
			}
		} else if len(text) > 0 {
			charValue = rune(text[0])
		}
		
		return v.ctx.Builder.ConstInt(types.U32, int64(charValue))
	}
	if ctx.NULL() != nil {
		return v.ctx.Builder.ConstNull(types.NewPointer(types.Void))
	}
	if ctx.InitializerList() != nil {
		return v.visitInitializerList(ctx.InitializerList())
	}
	return v.ctx.Builder.ConstInt(types.I64, 0)
}

func (v *IRVisitor) visitStringLiteral(rawText string) ir.Value {
	content, err := strconv.Unquote(rawText)
	if err != nil {
		if len(rawText) >= 2 {
			content = rawText[1 : len(rawText)-1]
		} else {
			content = rawText
		}
	}
	
	// Check for string interpolation \(expr)
	if containsInterpolation(content) {
		v.ctx.Logger.Warning("String interpolation not yet implemented, treating as raw string")
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

func (v *IRVisitor) visitInitializerList(ctx parser.IInitializerListContext) ir.Value {
	initCtx := ctx.(*parser.InitializerListContext)
	exprs := initCtx.AllExpression()
	
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
		if i == 0 { 
			elemType = val.Type() 
		}
		
		if constVal, ok := val.(ir.Constant); ok {
			elements[i] = constVal
		} else {
			elements[i] = v.ctx.Builder.ConstInt(types.I32, 0)
			if elemType == nil { 
				elemType = types.I32 
			}
		}
	}
	
	arrType := types.NewArray(elemType, int64(len(elements)))
	return &ir.ConstantArray{
		BaseValue: ir.BaseValue{ValType: arrType},
		Elements:  elements,
	}
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
	for _, argCtx := range ctx.AllArgument() {
		if expr := argCtx.(*parser.ArgumentContext).Expression(); expr != nil {
			if val, ok := v.Visit(expr).(ir.Value); ok {
				args = append(args, val)
			}
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
		return v.ctx.Builder.ConstInt(types.I64, 0)
	}
	if ctx.RAISE() != nil && len(args) == 1 {
		v.ctx.Builder.CreateRaise(args[0])
		v.ctx.Builder.CreateUnreachable()
		return v.ctx.Builder.ConstInt(types.I64, 0)
	}
	if ctx.BIT_CAST() != nil && len(args) == 1 {
		targetType := v.resolveType(ctx.Type_())
		// Bit cast reinterprets bits without conversion
		return v.ctx.Builder.CreateBitCast(args[0], targetType, "")
	}
	if ctx.SLICE() != nil && len(args) >= 1 {
		v.ctx.Logger.Warning("slice() intrinsic not fully implemented")
		// Would return (ptr, len) tuple
		return v.ctx.Builder.ConstInt(types.I64, 0)
	}
	
	return v.ctx.Builder.ConstInt(types.I64, 0)
}

// Helper function to check if string contains interpolation pattern
func containsInterpolation(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '\\' && s[i+1] == '(' {
			return true
		}
	}
	return false
}