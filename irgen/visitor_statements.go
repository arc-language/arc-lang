package irgen

import (
	"fmt"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (g *Generator) VisitBlock(ctx *parser.BlockContext) interface{} {
	// Check if this block defines a new scope.
	// We only enter the scope if we aren't ALREADY in it (e.g., from FunctionDecl).
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
		// Stop generating if we hit a terminator (ret, break, continue)
		if g.ctx.Builder.GetInsertBlock().Terminator() != nil {
			break
		}
	}
	return nil
}

func (g *Generator) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	// Execute deferred statements (LIFO) before the actual return
	g.deferStack.Emit(g)

	var val ir.Value
	if ctx.Expression() != nil {
		val = g.Visit(ctx.Expression()).(ir.Value)
		
		// Auto-cast return value to match function signature if needed
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

	// 1. Resolve Destination (L-Value)
	// Optimization for simple variable assignment
	if lhsCtx.IDENTIFIER() != nil && lhsCtx.DOT() == nil && lhsCtx.STAR() == nil && lhsCtx.LBRACKET() == nil {
		name := lhsCtx.IDENTIFIER().GetText()
		sym, ok := g.currentScope.Resolve(name)
		if ok && sym.IRValue != nil {
			// If it's an alloca (stack var), use it directly
			if alloca, isAlloca := sym.IRValue.(*ir.AllocaInst); isAlloca {
				destPtr = alloca
			} else {
				// Global variable or existing pointer value
				destPtr = sym.IRValue
			}
		} else {
			fmt.Printf("[IRGen] Error: Cannot resolve assignment target '%s'\n", name)
			return nil
		}
	} else {
		// Complex L-Value (pointers, fields, arrays)
		destPtr = g.getLValue(lhsCtx)
	}

	if destPtr == nil {
		return nil
	}

	// 2. Evaluate RHS
	rhs := g.Visit(ctx.Expression()).(ir.Value)
	finalVal := rhs

	// 3. Handle Compound Assignments (+=, -=, etc.)
	if ctx.AssignmentOp().ASSIGN() == nil {
		// For compound assignment, we must load the current value
		ptrType, isPtr := destPtr.Type().(*types.PointerType)
		if !isPtr {
			return nil
		}
		
		currVal := g.ctx.Builder.CreateLoad(ptrType.ElementType, destPtr, "")
		op := ctx.AssignmentOp()
		
		// Generate the arithmetic operation
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

	// 4. Store Result
	ptrType := destPtr.Type().(*types.PointerType)
	finalVal = g.emitCast(finalVal, ptrType.ElementType)
	g.ctx.Builder.CreateStore(finalVal, destPtr)

	return nil
}

func (g *Generator) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	cond := g.Visit(ctx.Expression(0)).(ir.Value)
	
	// Ensure condition is i1
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

	// Generate Then block
	g.ctx.SetInsertBlock(thenBlock)
	g.Visit(ctx.Block(0))
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(mergeBlock)
	}

	// Generate Else block (if exists)
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
	// Enter loop scope (defined in semantics)
	g.enterScope(ctx)
	defer g.exitScope()

	// Iterator loop placeholder
	if ctx.IN() != nil { return nil }

	// Blocks
	condBlock := g.ctx.Builder.CreateBlock("loop.cond")
	bodyBlock := g.ctx.Builder.CreateBlock("loop.body")
	postBlock := g.ctx.Builder.CreateBlock("loop.post")
	endBlock := g.ctx.Builder.CreateBlock("loop.end")

	// Init
	if len(ctx.AllSEMICOLON()) >= 2 {
		if ctx.VariableDecl() != nil { g.Visit(ctx.VariableDecl()) }
		if len(ctx.AllAssignmentStmt()) > 0 && ctx.AssignmentStmt(0).GetStart().GetTokenIndex() < ctx.SEMICOLON(0).GetSymbol().GetTokenIndex() {
			g.Visit(ctx.AssignmentStmt(0))
		}
	}

	g.ctx.Builder.CreateBr(condBlock)

	// Condition
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

	// Body
	g.ctx.SetInsertBlock(bodyBlock)
	g.loopStack = append(g.loopStack, loopInfo{breakBlock: endBlock, continueBlock: postBlock})
	g.Visit(ctx.Block())
	g.loopStack = g.loopStack[:len(g.loopStack)-1]
	
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(postBlock)
	}

	// Post
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
		// Create chain of if-else blocks
		g.ctx.SetInsertBlock(prevBlock)
		
		caseBlock := g.ctx.Builder.CreateBlock(fmt.Sprintf("case.%d", i))
		nextCheckBlock := g.ctx.Builder.CreateBlock(fmt.Sprintf("check.%d", i))
		
		// If last case and no default, jump to end
		if i == len(ctx.AllSwitchCase())-1 && ctx.DefaultCase() == nil {
			nextCheckBlock = endBlock
		}

		caseVal := g.Visit(c.Expression()).(ir.Value)
		cmp := g.ctx.Builder.CreateICmpEQ(cond, caseVal, "")
		g.ctx.Builder.CreateCondBr(cmp, caseBlock, nextCheckBlock)

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
	
	// Create global exception state if missing
	excGlobal := g.ctx.Module.GetGlobal("__exception_state")
	if excGlobal == nil {
		st := types.NewStruct("ExceptionState", []types.Type{types.I1, types.NewPointer(types.I8)}, false)
		excGlobal = g.ctx.Builder.CreateGlobalVariable("__exception_state", st, nil)
	}
	
	// Simple throw implementation: return -1
	// Real impl needs exception table/unwind info
	_ = val // store msg in global state in real impl
	g.ctx.Builder.CreateRet(g.ctx.Builder.ConstInt(types.I32, -1))
	return nil
}

func (g *Generator) VisitExpressionStmt(ctx *parser.ExpressionStmtContext) interface{} {
	return g.Visit(ctx.Expression())
}

// Helper to extract RangeExpression from expression tree
func getRangeExprContext(ctx parser.IExpressionContext) parser.IRangeExpressionContext {
	if ctx == nil { return nil }
	if lor := ctx.LogicalOrExpression(); lor != nil {
		if land := lor.LogicalAndExpression(0); land != nil {
			if bor := land.BitOrExpression(0); bor != nil {
				if bxor := bor.BitXorExpression(0); bxor != nil {
					if band := bxor.BitAndExpression(0); band != nil {
						if eq := band.EqualityExpression(0); eq != nil {
							if rel := eq.RelationalExpression(0); rel != nil {
								if sh := rel.ShiftExpression(0); sh != nil {
									return sh.RangeExpression(0)
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}