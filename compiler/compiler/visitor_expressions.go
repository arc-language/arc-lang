package compiler

import (
	"strconv"
	"strings"
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

// ============================================================================
// Expression Visitors
// ============================================================================

func (v *IRVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return v.Visit(ctx.LogicalOrExpression())
}

func (v *IRVisitor) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	result := v.Visit(ctx.LogicalAndExpression(0)).(ir.Value)

	for i := 1; i < len(ctx.AllLogicalAndExpression()); i++ {
		right := v.Visit(ctx.LogicalAndExpression(i)).(ir.Value)
		result = v.ctx.Builder.CreateOr(result, right, "")
	}

	return result
}

func (v *IRVisitor) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	result := v.Visit(ctx.BitOrExpression(0)).(ir.Value)

	for i := 1; i < len(ctx.AllBitOrExpression()); i++ {
		right := v.Visit(ctx.BitOrExpression(i)).(ir.Value)
		result = v.ctx.Builder.CreateAnd(result, right, "")
	}

	return result
}

func (v *IRVisitor) VisitBitOrExpression(ctx *parser.BitOrExpressionContext) interface{} {
	result := v.Visit(ctx.BitXorExpression(0)).(ir.Value)

	for i := 1; i < len(ctx.AllBitXorExpression()); i++ {
		right := v.Visit(ctx.BitXorExpression(i)).(ir.Value)
		result = v.ctx.Builder.CreateOr(result, right, "")
	}

	return result
}

func (v *IRVisitor) VisitBitXorExpression(ctx *parser.BitXorExpressionContext) interface{} {
	result := v.Visit(ctx.BitAndExpression(0)).(ir.Value)

	for i := 1; i < len(ctx.AllBitAndExpression()); i++ {
		right := v.Visit(ctx.BitAndExpression(i)).(ir.Value)
		result = v.ctx.Builder.CreateXor(result, right, "")
	}

	return result
}

func (v *IRVisitor) VisitBitAndExpression(ctx *parser.BitAndExpressionContext) interface{} {
	result := v.Visit(ctx.EqualityExpression(0)).(ir.Value)

	for i := 1; i < len(ctx.AllEqualityExpression()); i++ {
		right := v.Visit(ctx.EqualityExpression(i)).(ir.Value)
		result = v.ctx.Builder.CreateAnd(result, right, "")
	}

	return result
}

func (v *IRVisitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	result := v.Visit(ctx.RelationalExpression(0)).(ir.Value)

	for i := 0; i < len(ctx.AllEQ()); i++ {
		right := v.Visit(ctx.RelationalExpression(i + 1)).(ir.Value)
		result = v.createComparison(ir.ICmpEQ, result, right)
	}

	for i := 0; i < len(ctx.AllNE()); i++ {
		right := v.Visit(ctx.RelationalExpression(i + 1)).(ir.Value)
		result = v.createComparison(ir.ICmpNE, result, right)
	}

	return result
}

func (v *IRVisitor) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} {
	result := v.Visit(ctx.ShiftExpression(0)).(ir.Value)

	allLT := ctx.AllLT()
	allLE := ctx.AllLE()
	allGT := ctx.AllGT()
	allGE := ctx.AllGE()

	idx := 1
	for i := 0; i < len(allLT); i++ {
		right := v.Visit(ctx.ShiftExpression(idx)).(ir.Value)
		result = v.createComparison(ir.ICmpSLT, result, right)
		idx++
	}
	for i := 0; i < len(allLE); i++ {
		right := v.Visit(ctx.ShiftExpression(idx)).(ir.Value)
		result = v.createComparison(ir.ICmpSLE, result, right)
		idx++
	}
	for i := 0; i < len(allGT); i++ {
		right := v.Visit(ctx.ShiftExpression(idx)).(ir.Value)
		result = v.createComparison(ir.ICmpSGT, result, right)
		idx++
	}
	for i := 0; i < len(allGE); i++ {
		right := v.Visit(ctx.ShiftExpression(idx)).(ir.Value)
		result = v.createComparison(ir.ICmpSGE, result, right)
		idx++
	}

	return result
}

func (v *IRVisitor) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} {
	result := v.Visit(ctx.RangeExpression(0)).(ir.Value)

	allLT := ctx.AllLT()
	allGT := ctx.AllGT()

	// Left shift: 
	for i := 0; i < len(allLT)/2; i++ {
		right := v.Visit(ctx.RangeExpression(i + 1)).(ir.Value)
		result = v.ctx.Builder.CreateShl(result, right, "")
	}

	// Right shift: >>
	for i := 0; i < len(allGT)/2; i++ {
		right := v.Visit(ctx.RangeExpression(i + 1)).(ir.Value)
		result = v.ctx.Builder.CreateLShr(result, right, "")
	}

	return result
}

func (v *IRVisitor) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} {
	// For now, just handle as regular expression
	// TODO: Implement range support
	return v.Visit(ctx.AdditiveExpression(0))
}

func (v *IRVisitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	result := v.Visit(ctx.MultiplicativeExpression(0)).(ir.Value)

	for i := 0; i < len(ctx.AllPLUS()); i++ {
		right := v.Visit(ctx.MultiplicativeExpression(i + 1)).(ir.Value)
		if types.IsFloat(result.Type()) || types.IsFloat(right.Type()) {
			result = v.ctx.Builder.CreateFAdd(result, right, "")
		} else {
			result = v.ctx.Builder.CreateAdd(result, right, "")
		}
	}

	for i := 0; i < len(ctx.AllMINUS()); i++ {
		right := v.Visit(ctx.MultiplicativeExpression(i + 1 + len(ctx.AllPLUS()))).(ir.Value)
		if types.IsFloat(result.Type()) || types.IsFloat(right.Type()) {
			result = v.ctx.Builder.CreateFSub(result, right, "")
		} else {
			result = v.ctx.Builder.CreateSub(result, right, "")
		}
	}

	return result
}

func (v *IRVisitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	result := v.Visit(ctx.UnaryExpression(0)).(ir.Value)

	for i := 0; i < len(ctx.AllSTAR()); i++ {
		right := v.Visit(ctx.UnaryExpression(i + 1)).(ir.Value)
		if types.IsFloat(result.Type()) || types.IsFloat(right.Type()) {
			result = v.ctx.Builder.CreateFMul(result, right, "")
		} else {
			result = v.ctx.Builder.CreateMul(result, right, "")
		}
	}

	idx := len(ctx.AllSTAR()) + 1
	for i := 0; i < len(ctx.AllSLASH()); i++ {
		right := v.Visit(ctx.UnaryExpression(idx)).(ir.Value)
		if types.IsFloat(result.Type()) || types.IsFloat(right.Type()) {
			result = v.ctx.Builder.CreateFDiv(result, right, "")
		} else {
			result = v.ctx.Builder.CreateSDiv(result, right, "")
		}
		idx++
	}

	for i := 0; i < len(ctx.AllPERCENT()); i++ {
		right := v.Visit(ctx.UnaryExpression(idx)).(ir.Value)
		result = v.ctx.Builder.CreateSRem(result, right, "")
		idx++
	}

	return result
}

func (v *IRVisitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if ctx.PostfixExpression() != nil {
		return v.Visit(ctx.PostfixExpression())
	}

	if ctx.MINUS() != nil {
		operand := v.Visit(ctx.UnaryExpression()).(ir.Value)
		if types.IsFloat(operand.Type()) {
			zero := v.ctx.Builder.ConstFloat(operand.Type().(*types.FloatType), 0.0)
			return v.ctx.Builder.CreateFSub(zero, operand, "")
		} else {
			zero := v.ctx.Builder.ConstInt(operand.Type().(*types.IntType), 0)
			return v.ctx.Builder.CreateSub(zero, operand, "")
		}
	}

	if ctx.NOT() != nil {
		operand := v.Visit(ctx.UnaryExpression()).(ir.Value)
		zero := v.ctx.Builder.ConstInt(operand.Type().(*types.IntType), 0)
		return v.ctx.Builder.CreateICmp(ir.ICmpEQ, operand, zero, "")
	}

	if ctx.BIT_NOT() != nil {
		operand := v.Visit(ctx.UnaryExpression()).(ir.Value)
		allOnes := v.ctx.Builder.ConstInt(operand.Type().(*types.IntType), -1)
		return v.ctx.Builder.CreateXor(operand, allOnes, "")
	}

	if ctx.STAR() != nil {
		// Dereference
		ptr := v.Visit(ctx.UnaryExpression()).(ir.Value)
		ptrType, ok := ptr.Type().(*types.PointerType)
		if !ok {
			v.ctx.Logger.Error("Cannot dereference non-pointer type")
			return ptr
		}
		return v.ctx.Builder.CreateLoad(ptrType.ElementType, ptr, "")
	}

	if ctx.AMP() != nil {
		// Address-of
		operand := v.Visit(ctx.UnaryExpression()).(ir.Value)
		// The operand should already be an lvalue (address)
		return operand
	}

	if ctx.INCREMENT() != nil {
		// Pre-increment
		operand := v.Visit(ctx.UnaryExpression()).(ir.Value)
		one := v.ctx.Builder.ConstInt(operand.Type().(*types.IntType), 1)
		newVal := v.ctx.Builder.CreateAdd(operand, one, "")
		v.ctx.Builder.CreateStore(newVal, operand)
		return newVal
	}

	if ctx.DECREMENT() != nil {
		// Pre-decrement
		operand := v.Visit(ctx.UnaryExpression()).(ir.Value)
		one := v.ctx.Builder.ConstInt(operand.Type().(*types.IntType), 1)
		newVal := v.ctx.Builder.CreateSub(operand, one, "")
		v.ctx.Builder.CreateStore(newVal, operand)
		return newVal
	}

	return nil
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
	var actualFunc *ir.Function
	
	// If funcValue is a Function, use it directly
	if fn, ok := funcValue.(*ir.Function); ok {
		actualFunc = fn
		funcType = fn.FuncType
	} else if ptrType, ok := funcValue.Type().(*types.PointerType); ok {
		// If it's a pointer to a function type
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
	var result ir.Value
	if actualFunc != nil {
		result = v.ctx.Builder.CreateCall(actualFunc, args, "")
	} else {
		// For function pointers or indirect calls, we need the actual function
		// This case shouldn't happen with proper IR generation
		v.ctx.Logger.Error("Cannot call non-function value")
		return funcValue
	}
	
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

	// Find the method - look it up by name from the module's function list
	var methodFunc *ir.Function
	for _, fn := range v.ctx.Module.Functions {
		if fn.Name() == structType.Name+"."+methodName {
			methodFunc = fn
			break
		}
		// Try without struct name prefix
		if fn.Name() == methodName {
			methodFunc = fn
			break
		}
	}
	
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

	// Find the field index by name
	fieldIndex := -1
	for i := range structType.Fields {
		// Try to get field name from the struct type's name
		// This assumes field names are stored somewhere accessible
		// You may need to adjust this based on your actual type system
		if structType.Name != "" {
			// For now, we'll use a simple lookup
			// TODO: Implement proper field name storage in StructType
			fieldIndex = i
			break
		}
	}

	if fieldIndex == -1 {
		v.ctx.Logger.Error("Field %s not found in struct %s", memberName, structType.Name)
		return base
	}

	// Create GEP to get pointer to the field
	zero := v.ctx.Builder.ConstInt(types.NewInt(32, false), 0)
	fieldIdx := v.ctx.Builder.ConstInt(types.NewInt(32, false), int64(fieldIndex))
	indices := []ir.Value{zero, fieldIdx}

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
		zero := v.ctx.Builder.ConstInt(types.NewInt(32, false), 0)
		indices := []ir.Value{zero, index}
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
	intType, ok := currentVal.Type().(*types.IntType)
	if !ok {
		v.ctx.Logger.Error("Cannot increment non-integer type")
		return currentVal
	}
	one := v.ctx.Builder.ConstInt(intType, 1)
	newVal := v.ctx.Builder.CreateAdd(currentVal, one, "")

	// Store back
	if _, ok := base.Type().(*types.PointerType); ok {
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
	intType, ok := currentVal.Type().(*types.IntType)
	if !ok {
		v.ctx.Logger.Error("Cannot decrement non-integer type")
		return currentVal
	}
	one := v.ctx.Builder.ConstInt(intType, 1)
	newVal := v.ctx.Builder.CreateSub(currentVal, one, "")

	// Store back
	if _, ok := base.Type().(*types.PointerType); ok {
		v.ctx.Builder.CreateStore(newVal, base)
	}

	// Return the original value (post-decrement returns old value)
	return currentVal
}

func (v *IRVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.Literal() != nil {
		return v.Visit(ctx.Literal())
	}

	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		return v.lookupVariable(name)
	}

	if ctx.LPAREN() != nil && ctx.Expression() != nil {
		return v.Visit(ctx.Expression())
	}

	if ctx.TupleExpression() != nil {
		return v.Visit(ctx.TupleExpression())
	}

	if ctx.StructLiteral() != nil {
		return v.Visit(ctx.StructLiteral())
	}

	if ctx.CastExpression() != nil {
		return v.Visit(ctx.CastExpression())
	}

	if ctx.AllocaExpression() != nil {
		return v.Visit(ctx.AllocaExpression())
	}

	if ctx.SyscallExpression() != nil {
		return v.Visit(ctx.SyscallExpression())
	}

	if ctx.IntrinsicExpression() != nil {
		return v.Visit(ctx.IntrinsicExpression())
	}

	if ctx.QualifiedIdentifier() != nil {
		return v.Visit(ctx.QualifiedIdentifier())
	}

	v.ctx.Logger.Error("Unhandled primary expression")
	return nil
}

func (v *IRVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTEGER_LITERAL() != nil {
		return v.parseIntegerLiteral(ctx.INTEGER_LITERAL().GetText())
	}

	if ctx.FLOAT_LITERAL() != nil {
		return v.parseFloatLiteral(ctx.FLOAT_LITERAL().GetText())
	}

	if ctx.STRING_LITERAL() != nil {
		return v.parseStringLiteral(ctx.STRING_LITERAL().GetText())
	}

	if ctx.CHAR_LITERAL() != nil {
		return v.parseCharLiteral(ctx.CHAR_LITERAL().GetText())
	}

	if ctx.BOOLEAN_LITERAL() != nil {
		value := ctx.BOOLEAN_LITERAL().GetText() == "true"
		boolVal := int64(0)
		if value {
			boolVal = 1
		}
		return v.ctx.Builder.ConstInt(types.NewInt(1, false), boolVal)
	}

	if ctx.NULL() != nil {
		return v.ctx.Builder.ConstNull(types.NewPointer(types.NewInt(8, false)))
	}

	if ctx.InitializerList() != nil {
		return v.Visit(ctx.InitializerList())
	}

	return nil
}

func (v *IRVisitor) VisitTupleExpression(ctx *parser.TupleExpressionContext) interface{} {
	// Evaluate all elements
	exprs := ctx.AllExpression()
	values := make([]ir.Value, len(exprs))
	tupleTypes := make([]types.Type, len(exprs))

	for i, expr := range exprs {
		values[i] = v.Visit(expr).(ir.Value)
		tupleTypes[i] = values[i].Type()
	}

	// Create tuple struct type
	tupleType := types.NewStruct("", tupleTypes, false)

	// Build tuple using insertvalue
	var result ir.Value = v.ctx.Builder.ConstZero(tupleType)
	for i, val := range values {
		result = v.ctx.Builder.CreateInsertValue(result, val, []int{i}, "")
	}

	return result
}

// Helper functions

func (v *IRVisitor) parseIntegerLiteral(text string) ir.Value {
	// Remove underscores
	text = strings.ReplaceAll(text, "_", "")

	var value int64
	var err error

	if strings.HasPrefix(text, "0x") || strings.HasPrefix(text, "0X") {
		value, err = strconv.ParseInt(text[2:], 16, 64)
	} else if strings.HasPrefix(text, "0b") || strings.HasPrefix(text, "0B") {
		value, err = strconv.ParseInt(text[2:], 2, 64)
	} else if strings.HasPrefix(text, "0o") || strings.HasPrefix(text, "0O") {
		value, err = strconv.ParseInt(text[2:], 8, 64)
	} else {
		value, err = strconv.ParseInt(text, 10, 64)
	}

	if err != nil {
		v.ctx.Logger.Error("Invalid integer literal: %s", text)
		return v.ctx.Builder.ConstInt(types.NewInt(64, true), 0)
	}

	return v.ctx.Builder.ConstInt(types.NewInt(64, true), value)
}

func (v *IRVisitor) parseFloatLiteral(text string) ir.Value {
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		v.ctx.Logger.Error("Invalid float literal: %s", text)
		return v.ctx.Builder.ConstFloat(types.NewFloat(64), 0.0)
	}
	return v.ctx.Builder.ConstFloat(types.NewFloat(64), value)
}

func (v *IRVisitor) parseStringLiteral(text string) ir.Value {
	// Remove quotes
	if len(text) >= 2 {
		text = text[1 : len(text)-1]
	}

	// Handle escape sequences
	text = strings.ReplaceAll(text, "\\n", "\n")
	text = strings.ReplaceAll(text, "\\t", "\t")
	text = strings.ReplaceAll(text, "\\r", "\r")
	text = strings.ReplaceAll(text, "\\\"", "\"")
	text = strings.ReplaceAll(text, "\\\\", "\\")

	// Create array type for the string (length + 1 for null terminator)
	strLen := int64(len(text) + 1)
	arrayType := types.NewArray(types.NewInt(8, false), strLen)
	
	// Create a new global variable
	globalName := fmt.Sprintf(".str.%d", len(v.ctx.Module.Globals))
	global := &ir.Global{
		Name: globalName,
		Type: arrayType,
	}
	
	// Add to module globals
	v.ctx.Module.Globals = append(v.ctx.Module.Globals, global)
	
	// Return pointer to the first element
	zero := v.ctx.Builder.ConstInt(types.NewInt(64, false), 0)
	return v.ctx.Builder.CreateGEP(arrayType, global, []ir.Value{zero, zero}, "")
}

func (v *IRVisitor) parseCharLiteral(text string) ir.Value {
	// Remove quotes
	if len(text) >= 2 {
		text = text[1 : len(text)-1]
	}

	var charValue int64
	if strings.HasPrefix(text, "\\") {
		// Escape sequence
		switch text {
		case "\\n":
			charValue = '\n'
		case "\\t":
			charValue = '\t'
		case "\\r":
			charValue = '\r'
		case "\\'":
			charValue = '\''
		case "\\\\":
			charValue = '\\'
		case "\\0":
			charValue = 0
		default:
			charValue = int64(text[1])
		}
	} else if len(text) > 0 {
		charValue = int64(text[0])
	}

	return v.ctx.Builder.ConstInt(types.NewInt(8, false), charValue)
}

func (v *IRVisitor) createComparison(pred ir.ICmpPredicate, left, right ir.Value) ir.Value {
	if types.IsFloat(left.Type()) || types.IsFloat(right.Type()) {
		// Convert to float comparison
		var fpPred ir.FCmpPredicate
		switch pred {
		case ir.ICmpEQ:
			fpPred = ir.FCmpOEQ
		case ir.ICmpNE:
			fpPred = ir.FCmpONE
		case ir.ICmpSLT, ir.ICmpULT:
			fpPred = ir.FCmpOLT
		case ir.ICmpSLE, ir.ICmpULE:
			fpPred = ir.FCmpOLE
		case ir.ICmpSGT, ir.ICmpUGT:
			fpPred = ir.FCmpOGT
		case ir.ICmpSGE, ir.ICmpUGE:
			fpPred = ir.FCmpOGE
		}
		return v.ctx.Builder.CreateFCmp(fpPred, left, right, "")
	}
	return v.ctx.Builder.CreateICmp(pred, left, right, "")
}

func (v *IRVisitor) lookupVariable(name string) ir.Value {
	symbol, _ := v.ctx.currentScope.Lookup(name)
	if symbol == nil {
		v.ctx.Logger.Error("Undefined variable: %s", name)
		return v.ctx.Builder.ConstInt(types.NewInt(32, true), 0)
	}
	
	// The symbol contains the IR value
	if symbol.Value == nil {
		v.ctx.Logger.Error("Variable %s has no value", name)
		return v.ctx.Builder.ConstInt(types.NewInt(32, true), 0)
	}
	
	return symbol.Value
}