package irgen

import (
	"fmt"

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
	g.deferStack.Emit(g)

	var val ir.Value
	if ctx.Expression() != nil {
		val = g.Visit(ctx.Expression()).(ir.Value)
		
		if g.ctx.CurrentFunction != nil {
			targetType := g.ctx.CurrentFunction.FuncType.ReturnType
			val = g.emitCast(val, targetType)
		}
	}

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

	rhs := g.Visit(ctx.Expression()).(ir.Value)
	finalVal := rhs

	if ctx.AssignmentOp().ASSIGN() == nil {
		ptrType, isPtr := destPtr.Type().(*types.PointerType)
		if !isPtr {
			return nil
		}
		
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

	ptrType := destPtr.Type().(*types.PointerType)
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

	if ctx.IN() != nil { return nil }

	condBlock := g.ctx.Builder.CreateBlock("loop.cond")
	bodyBlock := g.ctx.Builder.CreateBlock("loop.body")
	postBlock := g.ctx.Builder.CreateBlock("loop.post")
	endBlock := g.ctx.Builder.CreateBlock("loop.end")

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
		
		// If this is the last case AND there is no default case,
		// the "next check" is actually the end block (fallthrough/exit)
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
	// If we have a default case, prevBlock points to the last check failure.
	// If we don't have a default case, prevBlock IS endBlock (set in loop above).
	// We must only generate code if we are NOT already at endBlock.
	
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