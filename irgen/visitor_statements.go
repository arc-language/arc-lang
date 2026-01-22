package irgen

import (
	"fmt"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (g *Generator) VisitBlock(ctx *parser.BlockContext) interface{} {
	// 1. Scope Management
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

	// 2. Defer Stack Scope
	// We need to snapshot the defer stack state so we only release
	// variables declared *inside* this block when this block ends.
	// (Note: In a full implementation, you'd push a new DeferScope)
	// For now, assuming simple function-level defers or manually handling scoping logic.
	
	// Visit Statements
	for _, stmt := range ctx.AllStatement() {
		g.Visit(stmt)
		if g.ctx.Builder.GetInsertBlock().Terminator() != nil {
			break
		}
	}
	
	return nil
}

func (g *Generator) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	// 1. Emit Deferred Actions (ARC Release, Defer statements)
	// This generates the code to decrement RefCounts and free objects
	// BEFORE the actual return instruction.
	if g.deferStack != nil {
		g.deferStack.Emit(g)
	}

	var val ir.Value
	
	// 2. Handle single expression return
	if ctx.Expression() != nil {
		val = g.Visit(ctx.Expression()).(ir.Value)
	} 
	
	// 3. Handle tuple return
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

	// 4. Generate Ret Instruction
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
	lhsCtx := ctx.LeftHandSide()
	var destPtr ir.Value

	if lhsCtx.IDENTIFIER() != nil && lhsCtx.DOT() == nil && lhsCtx.STAR() == nil && lhsCtx.LBRACKET() == nil {
		name := lhsCtx.IDENTIFIER().GetText()
		sym, ok := g.currentScope.Resolve(name)
		if ok && sym.IRValue != nil {
			if alloca, isAlloca := sym.IRValue.(*ir.AllocaInst); isAlloca {
				destPtr = alloca
			} else {
				destPtr = sym.IRValue
			}
		} else {
			fmt.Printf("[IRGen] Error: Cannot resolve assignment target '%s'\n", name)
			return nil
		}
	} else {
		destPtr = g.getLValue(lhsCtx)
	}

	if destPtr == nil {
		return nil
	}

	// SAFETY CHECK: Ensure assignment target is actually a pointer
	ptrType, isPtr := destPtr.Type().(*types.PointerType)
	if !isPtr {
		// This happens if trying to assign to a Constant or non-lvalue
		// We ignore it or report error to prevent panic
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
		if cond.Type().BitSize() > 1 {
			cond = g.ctx.Builder.CreateICmpNE(cond, g.ctx.Builder.ConstZero(cond.Type()), "")
		}
		
		thenBlock := g.ctx.Builder.CreateBlock(fmt.Sprintf("if.then.%d", i))
		
		if i < conditionCount - 1 {
			nextCheckBlock = g.ctx.Builder.CreateBlock(fmt.Sprintf("if.check.%d", i+1))
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

	condBlock := g.ctx.Builder.CreateBlock("loop.cond")
	bodyBlock := g.ctx.Builder.CreateBlock("loop.body")
	postBlock := g.ctx.Builder.CreateBlock("loop.post")
	endBlock := g.ctx.Builder.CreateBlock("loop.end")

	// --- Range-based Loop (for i in 0..5) ---
	if ctx.IN() != nil {
		if len(ctx.AllExpression()) == 0 { return nil }
		
		// Manually drill down to find RangeExpression
		var rangeCtx parser.IRangeExpressionContext
		
		e := ctx.Expression(0)
		if e != nil {
			if l1 := e.LogicalOrExpression(); l1 != nil {
				if l2 := l1.LogicalAndExpression(0); l2 != nil {
					if l3 := l2.BitOrExpression(0); l3 != nil {
						if l4 := l3.BitXorExpression(0); l4 != nil {
							if l5 := l4.BitAndExpression(0); l5 != nil {
								if l6 := l5.EqualityExpression(0); l6 != nil {
									if l7 := l6.RelationalExpression(0); l7 != nil {
										if l8 := l7.ShiftExpression(0); l8 != nil {
											rangeCtx = l8.RangeExpression(0)
										}
									}
								}
							}
						}
					}
				}
			}
		}

		if rangeCtx == nil || rangeCtx.RANGE() == nil {
			return nil
		}

		startVal := g.Visit(rangeCtx.AdditiveExpression(0)).(ir.Value)
		endVal := g.Visit(rangeCtx.AdditiveExpression(1)).(ir.Value)

		// Define loop variable 'i'
		varName := ctx.IDENTIFIER(0).GetText()
		sym, ok := g.currentScope.Resolve(varName)
		if !ok {
			return nil
		}
		
		// Create allocas for loop var if not already existing
		if sym.IRValue == nil {
			alloca := g.ctx.Builder.CreateAlloca(startVal.Type(), varName+".addr")
			sym.IRValue = alloca
		}
		
		// Init loop var = start
		g.ctx.Builder.CreateStore(startVal, sym.IRValue)
		g.ctx.Builder.CreateBr(condBlock)

		// Condition Check (i < end)
		g.ctx.SetInsertBlock(condBlock)
		currVal := g.ctx.Builder.CreateLoad(startVal.Type(), sym.IRValue, "")
		
		cmp := g.ctx.Builder.CreateICmpSLT(currVal, endVal, "")
		g.ctx.Builder.CreateCondBr(cmp, bodyBlock, endBlock)

		// Loop Body
		g.ctx.SetInsertBlock(bodyBlock)
		g.loopStack = append(g.loopStack, loopInfo{breakBlock: endBlock, continueBlock: postBlock})
		g.Visit(ctx.Block())
		g.loopStack = g.loopStack[:len(g.loopStack)-1]
		
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			g.ctx.Builder.CreateBr(postBlock)
		}

		// Post Step (i++)
		g.ctx.SetInsertBlock(postBlock)
		currVal = g.ctx.Builder.CreateLoad(startVal.Type(), sym.IRValue, "")
		one := g.ctx.Builder.ConstInt(startVal.Type().(*types.IntType), 1)
		nextVal := g.ctx.Builder.CreateAdd(currVal, one, "")
		g.ctx.Builder.CreateStore(nextVal, sym.IRValue)
		g.ctx.Builder.CreateBr(condBlock)

		g.ctx.SetInsertBlock(endBlock)
		return nil
	}

	// --- C-style Loop (for var i=0; i<10; i++) ---
	if len(ctx.AllSEMICOLON()) >= 2 {
		if ctx.VariableDecl() != nil { g.Visit(ctx.VariableDecl()) }
		if len(ctx.AllAssignmentStmt()) > 0 && ctx.AssignmentStmt(0).GetStart().GetTokenIndex() < ctx.SEMICOLON(0).GetSymbol().GetTokenIndex() {
			g.Visit(ctx.AssignmentStmt(0))
		}
	}

	g.ctx.Builder.CreateBr(condBlock)

	g.ctx.SetInsertBlock(condBlock)
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

	if cond.Type().BitSize() > 1 {
		cond = g.ctx.Builder.CreateICmpNE(cond, g.ctx.Builder.ConstZero(cond.Type()), "")
	}
	g.ctx.Builder.CreateCondBr(cond, bodyBlock, endBlock)

	g.ctx.SetInsertBlock(bodyBlock)
	g.loopStack = append(g.loopStack, loopInfo{breakBlock: endBlock, continueBlock: postBlock})
	g.Visit(ctx.Block())
	g.loopStack = g.loopStack[:len(g.loopStack)-1]
	
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(postBlock)
	}

	g.ctx.SetInsertBlock(postBlock)
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

func (g *Generator) VisitTryStmt(ctx *parser.TryStmtContext) interface{} {
	tryBlock := g.ctx.Builder.CreateBlock("try.start")
	endBlock := g.ctx.Builder.CreateBlock("try.end")
	
	var catchBlock *ir.BasicBlock
	if len(ctx.AllExceptClause()) > 0 {
		catchBlock = g.ctx.Builder.CreateBlock("try.catch")
	}

	g.ctx.Builder.CreateBr(tryBlock)
	g.ctx.SetInsertBlock(tryBlock)
	g.Visit(ctx.Block())
	
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(endBlock)
	}

	if catchBlock != nil {
		g.ctx.SetInsertBlock(catchBlock)
		g.Visit(ctx.ExceptClause(0).Block())
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			g.ctx.Builder.CreateBr(endBlock)
		}
	}

	g.ctx.SetInsertBlock(endBlock)
	return nil
}

func (g *Generator) VisitThrowStmt(ctx *parser.ThrowStmtContext) interface{} {
	val := g.Visit(ctx.Expression()).(ir.Value)
	_ = val 
	g.ctx.Builder.CreateRet(g.ctx.Builder.ConstInt(types.I32, -1))
	return nil
}

func (g *Generator) VisitExpressionStmt(ctx *parser.ExpressionStmtContext) interface{} {
	return g.Visit(ctx.Expression())
}