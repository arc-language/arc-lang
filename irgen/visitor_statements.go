package irgen

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (g *Generator) VisitBlock(ctx *parser.BlockContext) interface{} {
	shouldEnter := false
	if targetScope, isMapped := g.analysis.Scopes[ctx]; isMapped {
		if targetScope != g.currentScope {
			g.enterScope(ctx)
			shouldEnter = true
		}
	}

	if shouldEnter {
		defer g.exitScope()
	}

	for _, stmt := range ctx.AllStatement() {
		g.Visit(stmt)
		if g.ctx.Builder.GetInsertBlock().Terminator() != nil {
			break
		}
	}
	return nil
}

func (g *Generator) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	// Emit Deferred Actions BEFORE return
	if g.deferStack != nil {
		g.deferStack.Emit(g)
	}

	var val ir.Value
	if ctx.Expression() != nil {
		val = g.Visit(ctx.Expression()).(ir.Value)
	} 
	
	if ctx.TupleExpression() != nil {
		tupleCtx := ctx.TupleExpression()
		var fieldVals []ir.Value
		for _, expr := range tupleCtx.AllExpression() {
			v := g.Visit(expr).(ir.Value)
			fieldVals = append(fieldVals, v)
		}
		
		if g.ctx.CurrentFunction != nil {
			retType := g.ctx.CurrentFunction.FuncType.ReturnType
			if st, ok := retType.(*types.StructType); ok {
				var agg ir.Value = g.ctx.Builder.ConstZero(st)
				for i, v := range fieldVals {
					if i < len(st.Fields) {
						v = g.emitCast(v, st.Fields[i])
						agg = g.ctx.Builder.CreateInsertValue(agg, v, []int{i}, "")
					}
				}
				val = agg
			}
		}
	}

	if val != nil {
		if g.ctx.CurrentFunction != nil && ctx.TupleExpression() == nil {
			targetType := g.ctx.CurrentFunction.FuncType.ReturnType
			val = g.emitCast(val, targetType)
		}
		g.ctx.Builder.CreateRet(val)
	} else {
		g.ctx.Builder.CreateRetVoid()
	}
	return nil
}


func (g *Generator) VisitAssignmentStmt(ctx *parser.AssignmentStmtContext) interface{} {
	// Updated: Use UnaryExpression instead of LeftHandSide
	lhsCtx := ctx.UnaryExpression()
	
	// Resolve L-Value
	destPtr := g.getLValue(lhsCtx)

	if destPtr == nil {
		fmt.Printf("[IRGen] Error: Invalid assignment target at line %d\n", ctx.GetStart().GetLine())
		return nil
	}

	// SAFETY CHECK: Ensure assignment target is actually a pointer
	ptrType, isPtr := destPtr.Type().(*types.PointerType)
	if !isPtr {
		fmt.Printf("[IRGen] Error: Assignment target is not a pointer (type: %s) at line %d\n", destPtr.Type().String(), ctx.GetStart().GetLine())
		return nil
	}

	rhs := g.Visit(ctx.Expression()).(ir.Value)
	finalVal := rhs

	if ctx.AssignmentOp().ASSIGN() == nil {
		// Compound assignment: need to load current value
		currVal := g.ctx.Builder.CreateLoad(ptrType.ElementType, destPtr, "")
		op := ctx.AssignmentOp()
		
		if types.IsFloat(currVal.Type()) {
			if op.PLUS_ASSIGN() != nil {
				finalVal = g.ctx.Builder.CreateFAdd(currVal, rhs, "")
			} else if op.MINUS_ASSIGN() != nil {
				finalVal = g.ctx.Builder.CreateFSub(currVal, rhs, "")
			} else if op.STAR_ASSIGN() != nil {
				finalVal = g.ctx.Builder.CreateFMul(currVal, rhs, "")
			} else if op.SLASH_ASSIGN() != nil {
				finalVal = g.ctx.Builder.CreateFDiv(currVal, rhs, "")
			}
		} else {
			if op.PLUS_ASSIGN() != nil {
				finalVal = g.ctx.Builder.CreateAdd(currVal, rhs, "")
			} else if op.MINUS_ASSIGN() != nil {
				finalVal = g.ctx.Builder.CreateSub(currVal, rhs, "")
			} else if op.STAR_ASSIGN() != nil {
				finalVal = g.ctx.Builder.CreateMul(currVal, rhs, "")
			} else if op.SLASH_ASSIGN() != nil {
				finalVal = g.ctx.Builder.CreateSDiv(currVal, rhs, "")
			} else if op.PERCENT_ASSIGN() != nil {
				finalVal = g.ctx.Builder.CreateSRem(currVal, rhs, "")
			} else if op.BIT_AND_ASSIGN() != nil {
				finalVal = g.ctx.Builder.CreateAnd(currVal, rhs, "")
			} else if op.BIT_OR_ASSIGN() != nil {
				finalVal = g.ctx.Builder.CreateOr(currVal, rhs, "")
			} else if op.BIT_XOR_ASSIGN() != nil {
				finalVal = g.ctx.Builder.CreateXor(currVal, rhs, "")
			}
		}
	}

	// Ensure RHS type matches the element type of the pointer
	finalVal = g.emitCast(finalVal, ptrType.ElementType)
	g.ctx.Builder.CreateStore(finalVal, destPtr)

	return nil
}

func (g *Generator) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	mergeBlock := g.ctx.Builder.CreateBlock("if.end")
	conditionCount := len(ctx.AllExpression())
	var nextCheckBlock *ir.BasicBlock
	
	for i := 0; i < conditionCount; i++ {
		if nextCheckBlock != nil {
			g.ctx.SetInsertBlock(nextCheckBlock)
		}
		
		cond := g.Visit(ctx.Expression(i)).(ir.Value)
		
		// DEBUG: Print what we got
		fmt.Printf("[DEBUG] If condition %d: type=%s, value=%v\n", i, cond.Type(), cond)
		
		// FIX: Ensure the condition evaluation instruction is attached to the block
		if inst, ok := cond.(ir.Instruction); ok && inst.Parent() == nil {
			g.ctx.Builder.GetInsertBlock().AddInstruction(inst)
		}

		if cond.Type().BitSize() > 1 {
			cond = g.ctx.Builder.CreateICmpNE(cond, g.ctx.Builder.ConstZero(cond.Type()), "")
			
			// FIX: Ensure the implicit boolean check is attached to the block
			if inst, ok := cond.(ir.Instruction); ok && inst.Parent() == nil {
				g.ctx.Builder.GetInsertBlock().AddInstruction(inst)
			}
			
			fmt.Printf("[DEBUG] After CreateICmpNE: type=%s, value=%v\n", cond.Type(), cond)
		}
		
		thenBlock := g.ctx.Builder.CreateBlock("if.then")
		
		if i < conditionCount - 1 {
			nextCheckBlock = g.ctx.Builder.CreateBlock("if.check")
		} else {
			if len(ctx.AllBlock()) > conditionCount {
				nextCheckBlock = g.ctx.Builder.CreateBlock("if.else")
			} else {
				nextCheckBlock = mergeBlock
			}
		}
		
		g.ctx.Builder.CreateCondBr(cond, thenBlock, nextCheckBlock)
		
		g.ctx.SetInsertBlock(thenBlock)
		g.Visit(ctx.Block(i))
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			g.ctx.Builder.CreateBr(mergeBlock)
		}
	}
	
	if len(ctx.AllBlock()) > conditionCount {
		g.ctx.SetInsertBlock(nextCheckBlock)
		g.Visit(ctx.Block(conditionCount))
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			g.ctx.Builder.CreateBr(mergeBlock)
		}
	}
	
	g.ctx.SetInsertBlock(mergeBlock)
	return nil
}

func (g *Generator) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	g.enterScope(ctx)
	defer g.exitScope()

	// Use CreateBlock to ensure blocks are attached to the current function in the builder context
	condBlock := g.ctx.Builder.CreateBlock("loop.cond")
	bodyBlock := g.ctx.Builder.CreateBlock("loop.body")
	postBlock := g.ctx.Builder.CreateBlock("loop.post")
	endBlock := g.ctx.Builder.CreateBlock("loop.end")

	// --- Range-based Loop (for i in 0..5) ---
	if ctx.IN() != nil {
		if len(ctx.AllExpression()) == 0 { return nil }
		
		var findRange func(antlr.ParseTree) parser.IRangeExpressionContext
		findRange = func(node antlr.ParseTree) parser.IRangeExpressionContext {
			if r, ok := node.(parser.IRangeExpressionContext); ok {
				return r
			}
			if node.GetChildCount() == 1 {
				child := node.GetChild(0)
				if pt, ok := child.(antlr.ParseTree); ok {
					return findRange(pt)
				}
			}
			return nil
		}
		
		rangeCtx := findRange(ctx.Expression(0))
		if rangeCtx == nil || rangeCtx.RANGE() == nil {
			return nil
		}

		startVal := g.Visit(rangeCtx.AdditiveExpression(0)).(ir.Value)
		endVal := g.Visit(rangeCtx.AdditiveExpression(1)).(ir.Value)
		endVal = g.emitCast(endVal, startVal.Type())

		varName := ctx.IDENTIFIER(0).GetText()
		sym, ok := g.currentScope.Resolve(varName)
		if !ok {
			return nil
		}
		
		if sym.IRValue == nil {
			alloca := g.ctx.Builder.CreateAlloca(startVal.Type(), varName+".addr")
			sym.IRValue = alloca
		}
		
		g.ctx.Builder.CreateStore(startVal, sym.IRValue)
		g.ctx.Builder.CreateBr(condBlock)

		// Condition Check
		g.ctx.Builder.SetInsertPoint(condBlock)
		currVal := g.ctx.Builder.CreateLoad(startVal.Type(), sym.IRValue, "")
		
		// FIX: currVal is *ir.LoadInst (concrete), check Parent directly
		if currVal.Parent() == nil {
			g.ctx.Builder.GetInsertBlock().AddInstruction(currVal)
		}
		
		var cmp ir.Value
		if types.IsFloat(startVal.Type()) {
			cmp = g.ctx.Builder.CreateFCmp(ir.FCmpOLT, currVal, endVal, "")
		} else {
			cmp = g.ctx.Builder.CreateICmpSLT(currVal, endVal, "")
		}

		// FIX: cmp is ir.Value (interface), type assertion required
		if inst, ok := cmp.(ir.Instruction); ok && inst.Parent() == nil {
			g.ctx.Builder.GetInsertBlock().AddInstruction(inst)
		}

		g.ctx.Builder.CreateCondBr(cmp, bodyBlock, endBlock)

		// Loop Body
		g.ctx.Builder.SetInsertPoint(bodyBlock)
		g.loopStack = append(g.loopStack, loopInfo{breakBlock: endBlock, continueBlock: postBlock})
		g.Visit(ctx.Block())
		g.loopStack = g.loopStack[:len(g.loopStack)-1]
		
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			g.ctx.Builder.CreateBr(postBlock)
		}

		// Post Step
		g.ctx.Builder.SetInsertPoint(postBlock)
		currVal = g.ctx.Builder.CreateLoad(startVal.Type(), sym.IRValue, "")
		// Note: We should probably check this load too, though usually Post blocks are simple
		if currVal.Parent() == nil {
			g.ctx.Builder.GetInsertBlock().AddInstruction(currVal)
		}
		
		var nextVal ir.Value
		if types.IsFloat(startVal.Type()) {
			one := g.ctx.Builder.ConstFloat(startVal.Type().(*types.FloatType), 1.0)
			nextVal = g.ctx.Builder.CreateFAdd(currVal, one, "")
		} else {
			one := g.ctx.Builder.ConstInt(startVal.Type().(*types.IntType), 1)
			nextVal = g.ctx.Builder.CreateAdd(currVal, one, "")
		}
		
		g.ctx.Builder.CreateStore(nextVal, sym.IRValue)
		g.ctx.Builder.CreateBr(condBlock)

		g.ctx.Builder.SetInsertPoint(endBlock)
		return nil
	}

	// --- C-style Loop (for var i=0; i<10; i++) ---
	// 1. Initialization
	if len(ctx.AllSEMICOLON()) >= 2 {
		if ctx.VariableDecl() != nil { g.Visit(ctx.VariableDecl()) }
		if len(ctx.AllAssignmentStmt()) > 0 && ctx.AssignmentStmt(0).GetStart().GetTokenIndex() < ctx.SEMICOLON(0).GetSymbol().GetTokenIndex() {
			g.Visit(ctx.AssignmentStmt(0))
		}
	}

	g.ctx.Builder.CreateBr(condBlock)

	// 2. Condition
	g.ctx.Builder.SetInsertPoint(condBlock)
	var cond ir.Value = g.ctx.Builder.True()
	
	if len(ctx.AllSEMICOLON()) >= 2 {
		semi1 := ctx.SEMICOLON(0).GetSymbol().GetTokenIndex()
		semi2 := ctx.SEMICOLON(1).GetSymbol().GetTokenIndex()
		for _, expr := range ctx.AllExpression() {
			if expr.GetStart().GetTokenIndex() > semi1 && expr.GetStart().GetTokenIndex() < semi2 {
				cond = g.Visit(expr).(ir.Value)
				break
			}
		}
	} else if len(ctx.AllExpression()) > 0 {
		cond = g.Visit(ctx.Expression(0)).(ir.Value)
	}

	// FIX: Ensure loop condition expression is attached
	if inst, ok := cond.(ir.Instruction); ok && inst.Parent() == nil {
		g.ctx.Builder.GetInsertBlock().AddInstruction(inst)
	}

	if cond.Type().BitSize() > 1 {
		cond = g.ctx.Builder.CreateICmpNE(cond, g.ctx.Builder.ConstZero(cond.Type()), "")
		// FIX: Ensure implicit boolean check is attached
		if inst, ok := cond.(ir.Instruction); ok && inst.Parent() == nil {
			g.ctx.Builder.GetInsertBlock().AddInstruction(inst)
		}
	}
	g.ctx.Builder.CreateCondBr(cond, bodyBlock, endBlock)

	// 3. Body
	g.ctx.Builder.SetInsertPoint(bodyBlock)
	g.loopStack = append(g.loopStack, loopInfo{breakBlock: endBlock, continueBlock: postBlock})
	g.Visit(ctx.Block())
	g.loopStack = g.loopStack[:len(g.loopStack)-1]
	
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(postBlock)
	}

	// 4. Post
	g.ctx.Builder.SetInsertPoint(postBlock)
	if len(ctx.AllSEMICOLON()) >= 2 {
		semi2 := ctx.SEMICOLON(1).GetSymbol().GetTokenIndex()
		for _, assign := range ctx.AllAssignmentStmt() {
			if assign.GetStart().GetTokenIndex() > semi2 { g.Visit(assign) }
		}
		for _, expr := range ctx.AllExpression() {
			if expr.GetStart().GetTokenIndex() > semi2 { g.Visit(expr) }
		}
	}
	g.ctx.Builder.CreateBr(condBlock)

	// 5. End
	g.ctx.Builder.SetInsertPoint(endBlock)
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
		if ctx.Expression() != nil {
			gen.Visit(ctx.Expression())
		}
		if ctx.AssignmentStmt() != nil {
			gen.Visit(ctx.AssignmentStmt())
		}
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

		// Handle multiple expressions in a single case
		currentCheck := prevBlock
		for j, expr := range c.AllExpression() {
			g.ctx.SetInsertBlock(currentCheck)
			
			val := g.Visit(expr).(ir.Value)
			cmp := g.ctx.Builder.CreateICmpEQ(cond, val, "")
			
			nextExprBlock := nextCheckBlock
			if j < len(c.AllExpression())-1 {
				nextExprBlock = g.ctx.Builder.CreateBlock(fmt.Sprintf("check.%d.%d", i, j))
			}
			
			g.ctx.Builder.CreateCondBr(cmp, caseBlock, nextExprBlock)
			currentCheck = nextExprBlock
		}
		
		// Generate Case Body
		g.ctx.SetInsertBlock(caseBlock)
		for _, s := range c.AllStatement() {
			g.Visit(s)
			if g.ctx.Builder.GetInsertBlock().Terminator() != nil {
				break
			}
		}
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			g.ctx.Builder.CreateBr(endBlock)
		}

		prevBlock = nextCheckBlock
	}

	// Handle Default
	if prevBlock != endBlock {
		g.ctx.SetInsertBlock(prevBlock)
		
		if ctx.DefaultCase() != nil {
			for _, s := range ctx.DefaultCase().AllStatement() {
				g.Visit(s)
				if g.ctx.Builder.GetInsertBlock().Terminator() != nil {
					break
				}
			}
		}
		
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			g.ctx.Builder.CreateBr(endBlock)
		}
	}

	g.ctx.SetInsertBlock(endBlock)
	return nil
}

func (g *Generator) VisitExpressionStmt(ctx *parser.ExpressionStmtContext) interface{} {
	return g.Visit(ctx.Expression())
}