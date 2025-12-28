package compiler

import (
	"fmt"
	"strings"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (v *IRVisitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	token := ctx.GetStart()
	uniqueID := fmt.Sprintf("%d_%d", token.GetLine(), token.GetColumn())

	mergeBlock := v.ctx.Builder.CreateBlock("if.end." + uniqueID)

	cond := v.Visit(ctx.Expression(0)).(ir.Value)
	thenBlock := v.ctx.Builder.CreateBlock("if.then." + uniqueID)
	nextCheckBlock := v.ctx.Builder.CreateBlock("if.next." + uniqueID)

	v.ctx.Builder.CreateCondBr(cond, thenBlock, nextCheckBlock)

	v.ctx.SetInsertBlock(thenBlock)
	v.Visit(ctx.Block(0))
	if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
		v.ctx.Builder.CreateBr(mergeBlock)
	}

	v.ctx.SetInsertBlock(nextCheckBlock)
	count := len(ctx.AllIF())

	for i := 1; i < count; i++ {
		cond := v.Visit(ctx.Expression(i)).(ir.Value)
		
		thenName := fmt.Sprintf("elseif.then.%s.%d", uniqueID, i)
		nextName := fmt.Sprintf("elseif.next.%s.%d", uniqueID, i)
		
		thenBlock := v.ctx.Builder.CreateBlock(thenName)
		newNextBlock := v.ctx.Builder.CreateBlock(nextName)

		v.ctx.Builder.CreateCondBr(cond, thenBlock, newNextBlock)

		v.ctx.SetInsertBlock(thenBlock)
		v.Visit(ctx.Block(i))
		if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
			v.ctx.Builder.CreateBr(mergeBlock)
		}

		v.ctx.SetInsertBlock(newNextBlock)
	}

	if len(ctx.AllBlock()) > count {
		v.Visit(ctx.Block(count))
	}

	if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
		v.ctx.Builder.CreateBr(mergeBlock)
	}

	v.ctx.SetInsertBlock(mergeBlock)
	return nil
}

func (v *IRVisitor) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	v.ctx.PushScope()
	defer v.ctx.PopScope()

	if ctx.IN() != nil {
		return v.visitForInLoop(ctx)
	}

	token := ctx.GetStart()
	uniqueID := fmt.Sprintf("%d_%d", token.GetLine(), token.GetColumn())

	semicolons := ctx.AllSEMICOLON()
	isClause := len(semicolons) == 2

	if isClause {
		if ctx.VariableDecl() != nil {
			v.Visit(ctx.VariableDecl())
		} else if len(ctx.AllAssignmentStmt()) > 0 {
			firstAssign := ctx.AssignmentStmt(0)
			semi1 := semicolons[0]
			if v.isBefore(firstAssign, semi1) {
				v.Visit(firstAssign)
			}
		}
	}

	condBlock := v.ctx.Builder.CreateBlock("loop.cond." + uniqueID)
	bodyBlock := v.ctx.Builder.CreateBlock("loop.body." + uniqueID)
	postBlock := v.ctx.Builder.CreateBlock("loop.post." + uniqueID)
	endBlock := v.ctx.Builder.CreateBlock("loop.end." + uniqueID)

	continueTarget := condBlock
	if isClause {
		continueTarget = postBlock
	}

	v.ctx.Builder.CreateBr(condBlock)
	v.ctx.SetInsertBlock(condBlock)

	var cond ir.Value
	if isClause {
		semi1 := semicolons[0]
		semi2 := semicolons[1]
		found := false
		for _, expr := range ctx.AllExpression() {
			if v.isAfter(expr, semi1) && v.isBefore(expr, semi2) {
				cond = v.Visit(expr).(ir.Value)
				found = true
				break
			}
		}
		if !found {
			cond = v.ctx.Builder.True()
		}
	} else if ctx.Expression(0) != nil {
		cond = v.Visit(ctx.Expression(0)).(ir.Value)
	} else {
		cond = v.ctx.Builder.True()
	}

	v.ctx.Builder.CreateCondBr(cond, bodyBlock, endBlock)

	v.ctx.SetInsertBlock(bodyBlock)
	v.ctx.PushLoop(continueTarget, endBlock)
	v.Visit(ctx.Block())
	v.ctx.PopLoop()

	if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
		v.ctx.Builder.CreateBr(continueTarget)
	}

	v.ctx.SetInsertBlock(postBlock)
	if isClause {
		semi2 := semicolons[1]
		for _, assign := range ctx.AllAssignmentStmt() {
			if v.isAfter(assign, semi2) {
				v.Visit(assign)
			}
		}
		for _, expr := range ctx.AllExpression() {
			if v.isAfter(expr, semi2) {
				v.Visit(expr)
			}
		}
	}

	if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
		v.ctx.Builder.CreateBr(condBlock)
	}

	v.ctx.SetInsertBlock(endBlock)
	return nil
}

func (v *IRVisitor) visitForInLoop(ctx *parser.ForStmtContext) interface{} {
	isTwoVar := len(ctx.AllIDENTIFIER()) == 2
	
	varName := ctx.IDENTIFIER(0).GetText()
	var valueName string
	if isTwoVar {
		valueName = ctx.IDENTIFIER(1).GetText()
	}
	
	expr := ctx.Expression(0)
	
	// FIX: Update traversal to match new expression hierarchy
	// Hierarchy: Lor -> Land -> BOr -> BXor -> BAnd -> Eq -> Rel -> Shift -> Range
	
	var rngCtx parser.IRangeExpressionContext
	if lor := expr.LogicalOrExpression(); lor != nil {
		if land := lor.LogicalAndExpression(0); land != nil {
			if bor := land.BitOrExpression(0); bor != nil {
				if bxor := bor.BitXorExpression(0); bxor != nil {
					if band := bxor.BitAndExpression(0); band != nil {
						if eq := band.EqualityExpression(0); eq != nil {
							if rel := eq.RelationalExpression(0); rel != nil {
								if shift := rel.ShiftExpression(0); shift != nil {
									rngCtx = shift.RangeExpression(0)
								}
							}
						}
					}
				}
			}
		}
	}

	if rngCtx != nil && rngCtx.RANGE() != nil {
		return v.visitForInRange(ctx, varName, rngCtx)
	}

	collection := v.Visit(expr).(ir.Value)
	collectionType := collection.Type()
	
	var collectionPtr ir.Value
	
	// FIX: Traverse down to PrimaryExpression to find the identifier
	// This is painful manual traversal, but necessary to find the original alloca
	if lor := expr.LogicalOrExpression(); lor != nil {
		if land := lor.LogicalAndExpression(0); land != nil {
			if bor := land.BitOrExpression(0); bor != nil {
				if bxor := bor.BitXorExpression(0); bxor != nil {
					if band := bxor.BitAndExpression(0); band != nil {
						if eq := band.EqualityExpression(0); eq != nil {
							if rel := eq.RelationalExpression(0); rel != nil {
								if shift := rel.ShiftExpression(0); shift != nil {
									if rng := shift.RangeExpression(0); rng != nil {
										if add := rng.AdditiveExpression(0); add != nil {
											if mul := add.MultiplicativeExpression(0); mul != nil {
												if un := mul.UnaryExpression(0); un != nil {
													if post := un.PostfixExpression(); post != nil {
														if prim := post.PrimaryExpression(); prim != nil {
															if prim.IDENTIFIER() != nil {
																name := prim.IDENTIFIER().GetText()
																if sym, ok := v.ctx.currentScope.Lookup(name); ok {
																	if alloca, isAlloca := sym.Value.(*ir.AllocaInst); isAlloca {
																		collectionPtr = alloca
																		collectionType = alloca.AllocatedType
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	
	if collectionPtr == nil {
		collectionPtr = collection
	}
	
	if ptrType, ok := collectionType.(*types.PointerType); ok {
		collectionType = ptrType.ElementType
	}
	
	if vecType, ok := collectionType.(*types.DynamicVectorType); ok {
		return v.visitForInVector(ctx, varName, collectionPtr, vecType)
	}
	
	if structType, ok := collectionType.(*types.StructType); ok {
		if strings.HasPrefix(structType.Name, "__vector_") {
			if len(structType.Fields) >= 3 {
				if ptrField, ok := structType.Fields[0].(*types.PointerType); ok {
					elemType := ptrField.ElementType
					vecType := types.NewDynamicVector(elemType)
					return v.visitForInVector(ctx, varName, collectionPtr, vecType)
				}
			}
		}
	}
	
	if mapType, ok := collectionType.(*types.MapType); ok {
		return v.visitForInMap(ctx, varName, valueName, collectionPtr, mapType)
	}
	
	if arrType, ok := collectionType.(*types.ArrayType); ok {
		return v.visitForInArray(ctx, varName, collectionPtr, arrType)
	}
	
	v.ctx.Logger.Error("for-in loop expects a range, vector, map, or array")
	return nil
}

// ... (Rest of file: visitForInRange, visitForInVector, visitForInMap, visitForInArray, VisitBreakStmt, VisitContinueStmt - unchanged) ...
func (v *IRVisitor) visitForInRange(ctx *parser.ForStmtContext, varName string, rngCtx parser.IRangeExpressionContext) interface{} {
	// 2. Evaluate Start and End
	startVal := v.Visit(rngCtx.AdditiveExpression(0)).(ir.Value)
	endVal := v.Visit(rngCtx.AdditiveExpression(1)).(ir.Value)

	// Basic type check
	if !startVal.Type().Equal(endVal.Type()) {
		v.logger.Warning("Range start and end types differ, may need implicit cast")
	}

	// 3. Setup Loop Variable
	loopVarType := startVal.Type()
	loopVarPtr := v.ctx.Builder.CreateAlloca(loopVarType, varName+".addr")
	
	// Initialize loop variable
	v.ctx.Builder.CreateStore(startVal, loopVarPtr)
	v.ctx.currentScope.Define(varName, loopVarPtr)

	// 4. Create Blocks
	token := ctx.GetStart()
	uniqueID := fmt.Sprintf("%d_%d", token.GetLine(), token.GetColumn())
	
	condBlock := v.ctx.Builder.CreateBlock("for.cond." + uniqueID)
	bodyBlock := v.ctx.Builder.CreateBlock("for.body." + uniqueID)
	stepBlock := v.ctx.Builder.CreateBlock("for.step." + uniqueID)
	endBlock := v.ctx.Builder.CreateBlock("for.end." + uniqueID)

	v.ctx.Builder.CreateBr(condBlock)

	// 5. Condition Block: if x < end
	v.ctx.SetInsertBlock(condBlock)
	currVal := v.ctx.Builder.CreateLoad(loopVarType, loopVarPtr, "")
	
	cmp := v.ctx.Builder.CreateICmpSLT(currVal, endVal, "")
	v.ctx.Builder.CreateCondBr(cmp, bodyBlock, endBlock)

	// 6. Body Block
	v.ctx.SetInsertBlock(bodyBlock)
	v.ctx.PushLoop(stepBlock, endBlock) 
	v.Visit(ctx.Block())
	v.ctx.PopLoop()

	if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
		v.ctx.Builder.CreateBr(stepBlock)
	}

	// 7. Step Block: x = x + 1
	v.ctx.SetInsertBlock(stepBlock)
	currValForStep := v.ctx.Builder.CreateLoad(loopVarType, loopVarPtr, "")
	
	var one ir.Constant
	if intType, ok := loopVarType.(*types.IntType); ok {
		one = v.ctx.Builder.ConstInt(intType, 1)
	} else {
		one = v.ctx.Builder.ConstInt(types.I64, 1)
	}

	nextVal := v.ctx.Builder.CreateAdd(currValForStep, one, "")
	v.ctx.Builder.CreateStore(nextVal, loopVarPtr)
	v.ctx.Builder.CreateBr(condBlock)

	// 8. End Block
	v.ctx.SetInsertBlock(endBlock)
	return nil
}

func (v *IRVisitor) visitForInVector(ctx *parser.ForStmtContext, varName string, collection ir.Value, vecType *types.DynamicVectorType) interface{} {
	token := ctx.GetStart()
	uniqueID := fmt.Sprintf("%d_%d", token.GetLine(), token.GetColumn())
	
	v.logger.Debug("Compiling for-in loop over vector")
	
	// Get the runtime struct type
	structType := v.ctx.GetVectorRuntimeType(vecType.ElementType)
	
	// Collection should already be a pointer to the vector struct
	// (from the variable's alloca in VisitVariableDecl)
	var vecPtr ir.Value
	if _, ok := collection.Type().(*types.PointerType); ok {
		// It's already a pointer - use it directly
		vecPtr = collection
	} else {
		v.ctx.Logger.Error("Expected pointer to vector struct, got %v", collection.Type())
		return nil
	}
	
	// Get length field (field index 1) - do this ONCE before the loop
	lenGEP := v.ctx.Builder.CreateStructGEP(structType, vecPtr, 1, "")
	vecLen := v.ctx.Builder.CreateLoad(types.I64, lenGEP, "vec.len")
	
	// Get data pointer (field index 0) - do this ONCE before the loop
	dataGEP := v.ctx.Builder.CreateStructGEP(structType, vecPtr, 0, "")
	vecData := v.ctx.Builder.CreateLoad(types.NewPointer(vecType.ElementType), dataGEP, "vec.data")
	
	// Create index variable
	indexType := types.I64
	indexPtr := v.ctx.Builder.CreateAlloca(indexType, "vec.index.addr")
	zero := v.ctx.Builder.ConstInt(indexType, 0)
	v.ctx.Builder.CreateStore(zero, indexPtr)
	
	// Create blocks
	condBlock := v.ctx.Builder.CreateBlock("vec.cond." + uniqueID)
	bodyBlock := v.ctx.Builder.CreateBlock("vec.body." + uniqueID)
	stepBlock := v.ctx.Builder.CreateBlock("vec.step." + uniqueID)
	endBlock := v.ctx.Builder.CreateBlock("vec.end." + uniqueID)
	
	v.ctx.Builder.CreateBr(condBlock)
	
	// Condition: index < length
	v.ctx.SetInsertBlock(condBlock)
	currIndex := v.ctx.Builder.CreateLoad(indexType, indexPtr, "")
	cmp := v.ctx.Builder.CreateICmpSLT(currIndex, vecLen, "")
	v.ctx.Builder.CreateCondBr(cmp, bodyBlock, endBlock)
	
	// Body: load element and bind to loop variable
	v.ctx.SetInsertBlock(bodyBlock)
	
	// Get element: data[index]
	index := v.ctx.Builder.CreateLoad(indexType, indexPtr, "")
	elemPtr := v.ctx.Builder.CreateInBoundsGEP(vecType.ElementType, vecData, []ir.Value{index}, "")
	
	// Create loop variable and load element into it
	loopVarPtr := v.ctx.Builder.CreateAlloca(vecType.ElementType, varName+".addr")
	elemVal := v.ctx.Builder.CreateLoad(vecType.ElementType, elemPtr, "")
	v.ctx.Builder.CreateStore(elemVal, loopVarPtr)
	v.ctx.currentScope.Define(varName, loopVarPtr)
	
	// Execute loop body
	v.ctx.PushLoop(stepBlock, endBlock)
	v.Visit(ctx.Block())
	v.ctx.PopLoop()
	
	if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
		v.ctx.Builder.CreateBr(stepBlock)
	}
	
	// Step: index++
	v.ctx.SetInsertBlock(stepBlock)
	currIdx := v.ctx.Builder.CreateLoad(indexType, indexPtr, "")
	one := v.ctx.Builder.ConstInt(indexType, 1)
	nextIdx := v.ctx.Builder.CreateAdd(currIdx, one, "")
	v.ctx.Builder.CreateStore(nextIdx, indexPtr)
	v.ctx.Builder.CreateBr(condBlock)
	
	// End block
	v.ctx.SetInsertBlock(endBlock)
	return nil
}

// New function for map iteration
func (v *IRVisitor) visitForInMap(ctx *parser.ForStmtContext, keyName, valueName string, collection ir.Value, mapType *types.MapType) interface{} {
	token := ctx.GetStart()
	uniqueID := fmt.Sprintf("%d_%d", token.GetLine(), token.GetColumn())
	
	v.logger.Debug("Compiling for-in loop over map")
	
	// Get the runtime struct type
	_ = v.ctx.GetMapRuntimeType(mapType.KeyType, mapType.ValueType)
	
	// This is complex - requires iterating through hash buckets
	// For now, emit a warning
	v.ctx.Logger.Warning("Map iteration not fully implemented - loop will be empty")
	
	// Avoid unused variable warnings
	_ = keyName
	_ = valueName
	_ = collection
	
	// Create empty loop that exits immediately
	endBlock := v.ctx.Builder.CreateBlock("map.end." + uniqueID)
	v.ctx.Builder.CreateBr(endBlock)
	v.ctx.SetInsertBlock(endBlock)
	
	return nil
}

// Array iteration
func (v *IRVisitor) visitForInArray(ctx *parser.ForStmtContext, varName string, collection ir.Value, arrType *types.ArrayType) interface{} {
	token := ctx.GetStart()
	uniqueID := fmt.Sprintf("%d_%d", token.GetLine(), token.GetColumn())
	
	v.logger.Debug("Compiling for-in loop over array of length %d", arrType.Length)
	
	// Create index variable
	indexType := types.I64
	indexPtr := v.ctx.Builder.CreateAlloca(indexType, "arr.index.addr")
	zero := v.ctx.Builder.ConstInt(indexType, 0)
	v.ctx.Builder.CreateStore(zero, indexPtr)
	
	// Create blocks
	condBlock := v.ctx.Builder.CreateBlock("arr.cond." + uniqueID)
	bodyBlock := v.ctx.Builder.CreateBlock("arr.body." + uniqueID)
	stepBlock := v.ctx.Builder.CreateBlock("arr.step." + uniqueID)
	endBlock := v.ctx.Builder.CreateBlock("arr.end." + uniqueID)
	
	v.ctx.Builder.CreateBr(condBlock)
	
	// Condition: index < length
	v.ctx.SetInsertBlock(condBlock)
	currIndex := v.ctx.Builder.CreateLoad(indexType, indexPtr, "")
	length := v.ctx.Builder.ConstInt(indexType, arrType.Length)
	cmp := v.ctx.Builder.CreateICmpSLT(currIndex, length, "")
	v.ctx.Builder.CreateCondBr(cmp, bodyBlock, endBlock)
	
	// Body: load array element and bind to loop variable
	v.ctx.SetInsertBlock(bodyBlock)
	
	// Get pointer to array element: collection[index]
	index := v.ctx.Builder.CreateLoad(indexType, indexPtr, "")
	elemPtr := v.ctx.Builder.CreateInBoundsGEP(arrType, collection, []ir.Value{zero, index}, "")
	
	// Create loop variable and load element into it
	loopVarPtr := v.ctx.Builder.CreateAlloca(arrType.ElementType, varName+".addr")
	elemVal := v.ctx.Builder.CreateLoad(arrType.ElementType, elemPtr, "")
	v.ctx.Builder.CreateStore(elemVal, loopVarPtr)
	v.ctx.currentScope.Define(varName, loopVarPtr)
	
	// Execute loop body
	v.ctx.PushLoop(stepBlock, endBlock)
	v.Visit(ctx.Block())
	v.ctx.PopLoop()
	
	if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
		v.ctx.Builder.CreateBr(stepBlock)
	}
	
	// Step: index++
	v.ctx.SetInsertBlock(stepBlock)
	currIdx := v.ctx.Builder.CreateLoad(indexType, indexPtr, "")
	one := v.ctx.Builder.ConstInt(indexType, 1)
	nextIdx := v.ctx.Builder.CreateAdd(currIdx, one, "")
	v.ctx.Builder.CreateStore(nextIdx, indexPtr)
	v.ctx.Builder.CreateBr(condBlock)
	
	// End block
	v.ctx.SetInsertBlock(endBlock)
	return nil
}

func (v *IRVisitor) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	loop := v.ctx.CurrentLoop()
	if loop == nil {
		v.ctx.Logger.Error("break statement outside of loop")
		return nil
	}
	v.logger.Debug("Emitting break instruction")
	v.ctx.Builder.CreateBr(loop.BreakBlock)
	return nil
}

func (v *IRVisitor) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {
	loop := v.ctx.CurrentLoop()
	if loop == nil {
		v.ctx.Logger.Error("continue statement outside of loop")
		return nil
	}
	v.logger.Debug("Emitting continue instruction")
	v.ctx.Builder.CreateBr(loop.ContinueBlock)
	return nil
}