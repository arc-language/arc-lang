package compiler

import (
	"fmt"

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
	if isTwoVar {
		// Two-variable for-in is only for maps, which are now library types
		v.ctx.Logger.Error("Two-variable for-in loops are only supported for map types (library feature)")
		return nil
	}
	
	expr := ctx.Expression(0)
	
	// Check for range expression (x..y)
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

	// Evaluate the collection expression
	collection := v.Visit(expr).(ir.Value)
	collectionType := collection.Type()
	
	// Check if it's an array type
	if arrType, ok := collectionType.(*types.ArrayType); ok {
		return v.visitForInArray(ctx, varName, collection, arrType)
	}
	
	// Check if collection is a pointer to an array
	if ptrType, ok := collectionType.(*types.PointerType); ok {
		if arrType, ok := ptrType.ElementType.(*types.ArrayType); ok {
			return v.visitForInArray(ctx, varName, collection, arrType)
		}
	}
	
	v.ctx.Logger.Error("for-in loop expects a range or array (vectors/maps are library types)")
	return nil
}

func (v *IRVisitor) visitForInRange(ctx *parser.ForStmtContext, varName string, rngCtx parser.IRangeExpressionContext) interface{} {
	// Evaluate Start and End
	startVal := v.Visit(rngCtx.AdditiveExpression(0)).(ir.Value)
	endVal := v.Visit(rngCtx.AdditiveExpression(1)).(ir.Value)

	// Basic type check
	if !startVal.Type().Equal(endVal.Type()) {
		v.logger.Warning("Range start and end types differ, may need implicit cast")
	}

	// Setup Loop Variable
	loopVarType := startVal.Type()
	loopVarPtr := v.ctx.Builder.CreateAlloca(loopVarType, varName+".addr")
	
	// Initialize loop variable
	v.ctx.Builder.CreateStore(startVal, loopVarPtr)
	v.ctx.currentScope.Define(varName, loopVarPtr)

	// Create Blocks
	token := ctx.GetStart()
	uniqueID := fmt.Sprintf("%d_%d", token.GetLine(), token.GetColumn())
	
	condBlock := v.ctx.Builder.CreateBlock("for.cond." + uniqueID)
	bodyBlock := v.ctx.Builder.CreateBlock("for.body." + uniqueID)
	stepBlock := v.ctx.Builder.CreateBlock("for.step." + uniqueID)
	endBlock := v.ctx.Builder.CreateBlock("for.end." + uniqueID)

	v.ctx.Builder.CreateBr(condBlock)

	// Condition Block: if x < end
	v.ctx.SetInsertBlock(condBlock)
	currVal := v.ctx.Builder.CreateLoad(loopVarType, loopVarPtr, "")
	
	cmp := v.ctx.Builder.CreateICmpSLT(currVal, endVal, "")
	v.ctx.Builder.CreateCondBr(cmp, bodyBlock, endBlock)

	// Body Block
	v.ctx.SetInsertBlock(bodyBlock)
	v.ctx.PushLoop(stepBlock, endBlock) 
	v.Visit(ctx.Block())
	v.ctx.PopLoop()

	if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
		v.ctx.Builder.CreateBr(stepBlock)
	}

	// Step Block: x = x + 1
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

	// End Block
	v.ctx.SetInsertBlock(endBlock)
	return nil
}

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
	
	// Determine if collection is pointer or value
	var elemPtr ir.Value
	if ptrType, ok := collection.Type().(*types.PointerType); ok {
		// Collection is a pointer to array
		index := v.ctx.Builder.CreateLoad(indexType, indexPtr, "")
		elemPtr = v.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, collection, []ir.Value{zero, index}, "")
	} else {
		// Collection is array value - need to extract
		v.ctx.Logger.Warning("Direct array value iteration not fully supported, use pointer to array")
		index := v.ctx.Builder.CreateLoad(indexType, indexPtr, "")
		elemPtr = v.ctx.Builder.CreateInBoundsGEP(arrType, collection, []ir.Value{zero, index}, "")
	}
	
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

func (v *IRVisitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	token := ctx.GetStart()
	uniqueID := fmt.Sprintf("%d_%d", token.GetLine(), token.GetColumn())
	
	// Evaluate switch expression once
	switchVal := v.Visit(ctx.Expression()).(ir.Value)
	
	// Create end block
	endBlock := v.ctx.Builder.CreateBlock("switch.end." + uniqueID)
	
	// Default block (goes to end if not specified)
	defaultBlock := endBlock
	if ctx.DefaultCase() != nil {
		defaultBlock = v.ctx.Builder.CreateBlock("switch.default." + uniqueID)
	}
	
	// Process each case
	cases := ctx.AllSwitchCase()
	caseBlocks := make([]*ir.BasicBlock, len(cases))
	
	// Create all case blocks first
	for i := range cases {
		caseBlocks[i] = v.ctx.Builder.CreateBlock(fmt.Sprintf("switch.case.%s.%d", uniqueID, i))
	}
	
	// Generate comparison chain
	currentCheckBlock := v.ctx.Builder.GetInsertBlock()
	
	for i, caseCtx := range cases {
		v.ctx.SetInsertBlock(currentCheckBlock)
		
		// Determine next check block
		var nextCheck *ir.BasicBlock
		if i+1 < len(cases) {
			nextCheck = v.ctx.Builder.CreateBlock(fmt.Sprintf("switch.check.%s.%d", uniqueID, i+1))
		} else {
			nextCheck = defaultBlock
		}
		
		// Compare switch value with case value
		caseVal := v.Visit(caseCtx.Expression()).(ir.Value)
		cmp := v.ctx.Builder.CreateICmpEQ(switchVal, caseVal, "")
		v.ctx.Builder.CreateCondBr(cmp, caseBlocks[i], nextCheck)
		
		// Generate case body
		v.ctx.SetInsertBlock(caseBlocks[i])
		for _, stmt := range caseCtx.AllStatement() {
			v.Visit(stmt)
			if v.ctx.Builder.GetInsertBlock().Terminator() != nil {
				break
			}
		}
		
		// Jump to end if no explicit terminator
		if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
			v.ctx.Builder.CreateBr(endBlock)
		}
		
		currentCheckBlock = nextCheck
	}
	
	// Default case
	if ctx.DefaultCase() != nil {
		v.ctx.SetInsertBlock(defaultBlock)
		for _, stmt := range ctx.DefaultCase().AllStatement() {
			v.Visit(stmt)
			if v.ctx.Builder.GetInsertBlock().Terminator() != nil {
				break
			}
		}
		if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
			v.ctx.Builder.CreateBr(endBlock)
		}
	} else {
		// If no default, the last nextCheck needs to jump to end
		v.ctx.SetInsertBlock(currentCheckBlock)
		if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
			v.ctx.Builder.CreateBr(endBlock)
		}
	}
	
	v.ctx.SetInsertBlock(endBlock)
	return nil
}

func (v *IRVisitor) VisitTryStmt(ctx *parser.TryStmtContext) interface{} {
	token := ctx.GetStart()
	uniqueID := fmt.Sprintf("%d_%d", token.GetLine(), token.GetColumn())
	
	tryBlock := v.ctx.Builder.CreateBlock("try." + uniqueID)
	endBlock := v.ctx.Builder.CreateBlock("try.end." + uniqueID)
	
	var catchBlock *ir.BasicBlock
	var finallyBlock *ir.BasicBlock
	
	if len(ctx.AllExceptClause()) > 0 {
		catchBlock = v.ctx.Builder.CreateBlock("catch." + uniqueID)
	}
	
	if ctx.FinallyClause() != nil {
		finallyBlock = v.ctx.Builder.CreateBlock("finally." + uniqueID)
	}
	
	v.ctx.Builder.CreateBr(tryBlock)
	
	// Try block
	v.ctx.SetInsertBlock(tryBlock)
	v.Visit(ctx.Block())
	
	if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
		if finallyBlock != nil {
			v.ctx.Builder.CreateBr(finallyBlock)
		} else {
			v.ctx.Builder.CreateBr(endBlock)
		}
	}
	
	// Except blocks
	if catchBlock != nil {
		v.ctx.SetInsertBlock(catchBlock)
		for _, except := range ctx.AllExceptClause() {
			// Generate catch block body
			if except.Block() != nil {
				v.Visit(except.Block())
			}
		}
		if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
			if finallyBlock != nil {
				v.ctx.Builder.CreateBr(finallyBlock)
			} else {
				v.ctx.Builder.CreateBr(endBlock)
			}
		}
	}
	
	// Finally block (always executes)
	if finallyBlock != nil {
		v.ctx.SetInsertBlock(finallyBlock)
		v.Visit(ctx.FinallyClause().Block())
		if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
			v.ctx.Builder.CreateBr(endBlock)
		}
	}
	
	v.ctx.SetInsertBlock(endBlock)
	
	v.ctx.Logger.Warning("Try-except-finally is not fully implemented - basic structure only")
	return nil
}