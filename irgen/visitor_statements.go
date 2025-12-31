package irgen

import (
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (g *Generator) VisitBlock(ctx *parser.BlockContext) interface{} {
	// If semantics mapped a scope to this block, enter it
	if _, isMapped := g.analysis.Scopes[ctx]; isMapped {
		g.enterScope(ctx)
		defer g.exitScope()
	}

	for _, stmt := range ctx.AllStatement() {
		g.Visit(stmt)
		// Stop generation if block is terminated (dead code)
		if g.ctx.Builder.GetInsertBlock().Terminator() != nil {
			break
		}
	}
	return nil
}

func (g *Generator) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	// 1. Emit deferred statements (LIFO)
	g.deferStack.Emit(g)

	// 2. Handle Return Value
	var val ir.Value
	if ctx.Expression() != nil {
		val = g.Visit(ctx.Expression()).(ir.Value)
		
		// Implicit casting to return type
		if g.ctx.CurrentFunction != nil {
			expectedType := g.ctx.CurrentFunction.FuncType.ReturnType
			val = g.emitCast(val, expectedType)
		}
	}

	// 3. Create instruction
	if val != nil {
		g.ctx.Builder.CreateRet(val)
	} else {
		g.ctx.Builder.CreateRetVoid()
	}
	return nil
}

func (g *Generator) VisitAssignmentStmt(ctx *parser.AssignmentStmtContext) interface{} {
	lhsCtx := ctx.LeftHandSide()
	var destPtr ir.Value

	// --- L-Value Resolution ---
	
	// Case A: Simple Identifier (x = ...)
	if lhsCtx.IDENTIFIER() != nil && lhsCtx.DOT() == nil && lhsCtx.STAR() == nil && lhsCtx.LBRACKET() == nil {
		name := lhsCtx.IDENTIFIER().GetText()
		sym, ok := g.currentScope.Resolve(name)
		if ok {
			if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
				destPtr = alloca
			} else {
				// Globals or arguments
				destPtr = sym.IRValue
			}
		}
	} else if lhsCtx.STAR() != nil {
		// Case B: Dereference (*ptr = ...)
		// Visit the inner expression to get the pointer value
		destPtr = g.Visit(lhsCtx.PostfixExpression()).(ir.Value)
	} else if lhsCtx.LBRACKET() != nil {
		// Case C: Array Indexing (arr[i] = ...)
		base := g.Visit(lhsCtx.PostfixExpression()).(ir.Value)
		index := g.Visit(lhsCtx.Expression()).(ir.Value)
		
		if ptrType, ok := base.Type().(*types.PointerType); ok {
			destPtr = g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, base, []ir.Value{index}, "")
		}
	} else if lhsCtx.DOT() != nil {
		// Case D: Member Access (obj.field = ...)
		base := g.Visit(lhsCtx.PostfixExpression()).(ir.Value)
		fieldName := lhsCtx.IDENTIFIER().GetText()

		// Auto-dereference if base is a pointer
		var structType *types.StructType
		if ptr, ok := base.Type().(*types.PointerType); ok {
			if st, ok := ptr.ElementType.(*types.StructType); ok {
				structType = st
			}
		}

		if structType != nil {
			if idx, ok := g.analysis.StructIndices[structType.Name][fieldName]; ok {
				destPtr = g.ctx.Builder.CreateStructGEP(structType, base, idx, "")
			}
		}
	}

	if destPtr == nil {
		return nil
	}

	// --- RHS Evaluation ---
	rhs := g.Visit(ctx.Expression()).(ir.Value)
	
	// --- Compound Assignment Handling (+=, -=, etc) ---
	finalVal := rhs
	if ctx.AssignmentOp().ASSIGN() == nil {
		// Load current value
		ptrType := destPtr.Type().(*types.PointerType)
		currVal := g.ctx.Builder.CreateLoad(ptrType.ElementType, destPtr, "")
		
		op := ctx.AssignmentOp()
		if op.PLUS_ASSIGN() != nil {
			if types.IsFloat(currVal.Type()) {
				finalVal = g.ctx.Builder.CreateFAdd(currVal, rhs, "")
			} else {
				finalVal = g.ctx.Builder.CreateAdd(currVal, rhs, "")
			}
		} else if op.MINUS_ASSIGN() != nil {
			if types.IsFloat(currVal.Type()) {
				finalVal = g.ctx.Builder.CreateFSub(currVal, rhs, "")
			} else {
				finalVal = g.ctx.Builder.CreateSub(currVal, rhs, "")
			}
		} else if op.STAR_ASSIGN() != nil {
			if types.IsFloat(currVal.Type()) {
				finalVal = g.ctx.Builder.CreateFMul(currVal, rhs, "")
			} else {
				finalVal = g.ctx.Builder.CreateMul(currVal, rhs, "")
			}
		} else if op.SLASH_ASSIGN() != nil {
			if types.IsFloat(currVal.Type()) {
				finalVal = g.ctx.Builder.CreateFDiv(currVal, rhs, "")
			} else {
				finalVal = g.ctx.Builder.CreateSDiv(currVal, rhs, "")
			}
		}
	}

	// Ensure type match (Cast RHS to LHS type)
	ptrType := destPtr.Type().(*types.PointerType)
	finalVal = g.emitCast(finalVal, ptrType.ElementType)

	g.ctx.Builder.CreateStore(finalVal, destPtr)
	return nil
}

func (g *Generator) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	cond := g.Visit(ctx.Expression(0)).(ir.Value)
	
	// Ensure boolean (i1)
	if cond.Type().BitSize() > 1 {
		cond = g.ctx.Builder.CreateICmpNE(cond, g.ctx.Builder.ConstZero(cond.Type()), "")
	}

	thenBlock := g.ctx.Builder.CreateBlock("if.then")
	mergeBlock := g.ctx.Builder.CreateBlock("if.end")
	elseBlock := mergeBlock
	
	if len(ctx.AllBlock()) > 1 {
		elseBlock = g.ctx.Builder.CreateBlock("if.else")
	}

	g.ctx.Builder.CreateCondBr(cond, thenBlock, elseBlock)

	// Generate Then
	g.ctx.SetInsertBlock(thenBlock)
	g.Visit(ctx.Block(0))
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(mergeBlock)
	}

	// Generate Else
	if elseBlock != mergeBlock {
		g.ctx.SetInsertBlock(elseBlock)
		g.Visit(ctx.Block(1))
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			g.ctx.Builder.CreateBr(mergeBlock)
		}
	}

	g.ctx.SetInsertBlock(mergeBlock)
	return nil
}

func (g *Generator) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	// Handle 'for var in ...' loops
	if ctx.IN() != nil {
		return g.visitForInLoop(ctx)
	}

	// Standard C-Style For Loop: for (init; cond; step)
	condBlock := g.ctx.Builder.CreateBlock("loop.cond")
	bodyBlock := g.ctx.Builder.CreateBlock("loop.body")
	postBlock := g.ctx.Builder.CreateBlock("loop.post")
	endBlock := g.ctx.Builder.CreateBlock("loop.end")

	// Init
	// Usually parsed outside or as first child if defined in grammar
	if ctx.VariableDecl() != nil {
		g.Visit(ctx.VariableDecl())
	} else if len(ctx.AllAssignmentStmt()) > 0 && len(ctx.AllSEMICOLON()) > 0 {
		// Heuristic: First assignment is init if followed by semicolons
		g.Visit(ctx.AssignmentStmt(0))
	}

	g.ctx.Builder.CreateBr(condBlock)
	g.ctx.SetInsertBlock(condBlock)

	// Condition
	var cond ir.Value = g.ctx.Builder.ConstInt(types.I1, 1) // Default true
	if len(ctx.AllExpression()) > 0 {
		cond = g.Visit(ctx.Expression(0)).(ir.Value)
	}
	g.ctx.Builder.CreateCondBr(cond, bodyBlock, endBlock)

	// Body
	g.ctx.SetInsertBlock(bodyBlock)
	g.loopStack = append(g.loopStack, loopInfo{breakBlock: endBlock, continueBlock: postBlock})
	g.Visit(ctx.Block())
	g.loopStack = g.loopStack[:len(g.loopStack)-1] // Pop

	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(postBlock)
	}

	// Post (Step)
	g.ctx.SetInsertBlock(postBlock)
	// Visit last assignment if it looks like a step
	if len(ctx.AllAssignmentStmt()) > 0 {
		// Logic depends on exact grammar structure
		// Assuming last assignment is the step
		g.Visit(ctx.AssignmentStmt(len(ctx.AllAssignmentStmt()) - 1))
	}
	g.ctx.Builder.CreateBr(condBlock)

	g.ctx.SetInsertBlock(endBlock)
	return nil
}

func (g *Generator) visitForInLoop(ctx *parser.ForStmtContext) interface{} {
	varName := ctx.IDENTIFIER(0).GetText()
	iterableExpr := ctx.Expression(0)

	// Evaluate Iterable
	collection := g.Visit(iterableExpr).(ir.Value)
	colType := collection.Type()

	// Check for Array or Pointer to Array
	if types.IsArray(colType) || (types.IsPointer(colType) && types.IsArray(colType.(*types.PointerType).ElementType)) {
		return g.visitForInArray(ctx, varName, collection)
	}
	
	// TODO: Add support for Range Expression (0..10) here
	return nil
}

func (g *Generator) visitForInArray(ctx *parser.ForStmtContext, varName string, collection ir.Value) interface{} {
	// 1. Setup Index
	idxPtr := g.ctx.Builder.CreateAlloca(types.I64, "idx")
	g.ctx.Builder.CreateStore(g.ctx.Builder.ConstInt(types.I64, 0), idxPtr)

	// 2. Determine Length and Base Pointer
	var length int64
	var arrType *types.ArrayType
	
	if ptr, ok := collection.Type().(*types.PointerType); ok {
		arrType = ptr.ElementType.(*types.ArrayType)
	} else {
		// If it's a value, spill to stack to iterate via GEP
		arrType = collection.Type().(*types.ArrayType)
		temp := g.ctx.Builder.CreateAlloca(arrType, "arr.temp")
		g.ctx.Builder.CreateStore(collection, temp)
		collection = temp
	}
	length = arrType.Length

	// 3. Blocks
	condBlock := g.ctx.Builder.CreateBlock("for.cond")
	bodyBlock := g.ctx.Builder.CreateBlock("for.body")
	stepBlock := g.ctx.Builder.CreateBlock("for.step")
	endBlock := g.ctx.Builder.CreateBlock("for.end")

	g.ctx.Builder.CreateBr(condBlock)
	
	// 4. Condition
	g.ctx.SetInsertBlock(condBlock)
	currIdx := g.ctx.Builder.CreateLoad(types.I64, idxPtr, "")
	cmp := g.ctx.Builder.CreateICmpSLT(currIdx, g.ctx.Builder.ConstInt(types.I64, length), "")
	g.ctx.Builder.CreateCondBr(cmp, bodyBlock, endBlock)

	// 5. Body
	g.ctx.SetInsertBlock(bodyBlock)
	
	// Load Element: arr[idx]
	elemPtr := g.ctx.Builder.CreateInBoundsGEP(arrType, collection, []ir.Value{g.getZeroValue(types.I32), currIdx}, "")
	elemVal := g.ctx.Builder.CreateLoad(arrType.ElementType, elemPtr, "")

	// Loop Variable
	if sym, ok := g.currentScope.Resolve(varName); ok {
		loopVarAlloca := g.ctx.Builder.CreateAlloca(sym.Type, varName+".addr")
		g.ctx.Builder.CreateStore(elemVal, loopVarAlloca)
		sym.IRValue = loopVarAlloca
	}

	g.loopStack = append(g.loopStack, loopInfo{breakBlock: endBlock, continueBlock: stepBlock})
	g.Visit(ctx.Block())
	g.loopStack = g.loopStack[:len(g.loopStack)-1]

	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(stepBlock)
	}

	// 6. Step
	g.ctx.SetInsertBlock(stepBlock)
	currIdx = g.ctx.Builder.CreateLoad(types.I64, idxPtr, "")
	nextIdx := g.ctx.Builder.CreateAdd(currIdx, g.ctx.Builder.ConstInt(types.I64, 1), "")
	g.ctx.Builder.CreateStore(nextIdx, idxPtr)
	g.ctx.Builder.CreateBr(condBlock)

	g.ctx.SetInsertBlock(endBlock)
	return nil
}

func (g *Generator) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	if len(g.loopStack) > 0 {
		g.ctx.Builder.CreateBr(g.loopStack[len(g.loopStack)-1].breakBlock)
	}
	return nil
}

func (g *Generator) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {
	if len(g.loopStack) > 0 {
		g.ctx.Builder.CreateBr(g.loopStack[len(g.loopStack)-1].continueBlock)
	}
	return nil
}

func (g *Generator) VisitDeferStmt(ctx *parser.DeferStmtContext) interface{} {
	g.deferStack.Add(func(gen *Generator) {
		if ctx.Expression() != nil { gen.Visit(ctx.Expression()) }
		if ctx.AssignmentStmt() != nil { gen.Visit(ctx.AssignmentStmt()) }
	})
	return nil
}

func (g *Generator) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	cond := g.Visit(ctx.Expression()).(ir.Value)
	endBlock := g.ctx.Builder.CreateBlock("switch.end")
	
	prevBlock := g.ctx.Builder.GetInsertBlock()
	
	for i, c := range ctx.AllSwitchCase() {
		g.ctx.SetInsertBlock(prevBlock)
		
		caseBlock := g.ctx.Builder.CreateBlock(fmt.Sprintf("case.%d", i))
		nextCheckBlock := g.ctx.Builder.CreateBlock(fmt.Sprintf("check.%d", i))
		
		if i == len(ctx.AllSwitchCase())-1 && ctx.DefaultCase() == nil {
			nextCheckBlock = endBlock
		}
		
		caseVal := g.Visit(c.Expression()).(ir.Value)
		cmp := g.ctx.Builder.CreateICmpEQ(cond, caseVal, "")
		g.ctx.Builder.CreateCondBr(cmp, caseBlock, nextCheckBlock)
		
		// Case Body
		g.ctx.SetInsertBlock(caseBlock)
		for _, s := range c.AllStatement() {
			g.Visit(s)
			if g.ctx.Builder.GetInsertBlock().Terminator() != nil { break }
		}
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil { 
			g.ctx.Builder.CreateBr(endBlock) 
		}
		
		prevBlock = nextCheckBlock
	}
	
	// Default Case
	g.ctx.SetInsertBlock(prevBlock)
	if ctx.DefaultCase() != nil {
		for _, s := range ctx.DefaultCase().AllStatement() {
			g.Visit(s)
			if g.ctx.Builder.GetInsertBlock().Terminator() != nil { break }
		}
	}
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil { 
		g.ctx.Builder.CreateBr(endBlock) 
	}
	
	g.ctx.SetInsertBlock(endBlock)
	return nil
}

func (g *Generator) VisitTryStmt(ctx *parser.TryStmtContext) interface{} {
	tryBlock := g.ctx.Builder.CreateBlock("try.start")
	endBlock := g.ctx.Builder.CreateBlock("try.end")
	
	var catchBlock *ir.BasicBlock
	if len(ctx.AllExceptClause()) > 0 {
		catchBlock = g.ctx.Builder.CreateBlock("try.catch")
	}

	g.ctx.Builder.CreateBr(tryBlock)
	g.ctx.SetInsertBlock(tryBlock)

	// Execute Try Body
	// Placeholder for EH checking after calls
	g.Visit(ctx.Block())
	
	// Check exception state here in full impl
	
	g.ctx.Builder.CreateBr(endBlock)

	if catchBlock != nil {
		g.ctx.SetInsertBlock(catchBlock)
		g.Visit(ctx.ExceptClause(0).Block())
		g.ctx.Builder.CreateBr(endBlock)
	}

	g.ctx.SetInsertBlock(endBlock)
	return nil
}

func (g *Generator) VisitThrowStmt(ctx *parser.ThrowStmtContext) interface{} {
	val := g.Visit(ctx.Expression()).(ir.Value)
	
	// Get global exception state or create it
	excGlobal := g.ctx.Module.GetGlobal("__exception_state")
	if excGlobal == nil {
		st := types.NewStruct("ExceptionState", []types.Type{types.I1, types.NewPointer(types.I8)}, false)
		excGlobal = g.ctx.Builder.CreateGlobalVariable("__exception_state", st, nil)
	}

	// Logic to store 'val' into exception state would go here (GEP + Store)
	
	// Return error code to unwind stack
	g.ctx.Builder.CreateRet(g.ctx.Builder.ConstInt(types.I32, -1))
	return val
}