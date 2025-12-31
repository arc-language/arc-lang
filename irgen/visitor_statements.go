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
		// Only enter if we aren't ALREADY in this scope
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
	// Execute deferred statements before returning
	g.deferStack.Emit(g)

	var val ir.Value
	if ctx.Expression() != nil {
		val = g.Visit(ctx.Expression()).(ir.Value)
		// Auto-cast return value if needed
		if g.ctx.CurrentFunction != nil {
			val = g.emitCast(val, g.ctx.CurrentFunction.FuncType.ReturnType)
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

	// 1. Resolve Destination
	if lhsCtx.IDENTIFIER() != nil && lhsCtx.DOT() == nil && lhsCtx.STAR() == nil && lhsCtx.LBRACKET() == nil {
		name := lhsCtx.IDENTIFIER().GetText()
		sym, ok := g.currentScope.Resolve(name)
		if ok && sym.IRValue != nil {
			if alloca, isAlloca := sym.IRValue.(*ir.AllocaInst); isAlloca {
				destPtr = alloca
			} else {
				// Global variable or other pointer
				destPtr = sym.IRValue
			}
		} else {
			fmt.Printf("[IRGen] Error: Cannot resolve assignment target '%s'\n", name)
			return nil
		}
	} else {
		// Handle complex lvalues (*ptr, arr[i], etc)
		destPtr = g.getLValue(lhsCtx)
	}

	if destPtr == nil {
		return nil
	}

	// 2. Evaluate RHS
	rhs := g.Visit(ctx.Expression()).(ir.Value)
	finalVal := rhs

	// 3. Handle Compound Assignments (+=, -=, etc)
	if ctx.AssignmentOp().ASSIGN() == nil {
		// Load current value
		ptrType := destPtr.Type().(*types.PointerType)
		currVal := g.ctx.Builder.CreateLoad(ptrType.ElementType, destPtr, "")
		
		op := ctx.AssignmentOp()
		isFloat := types.IsFloat(currVal.Type())

		if op.PLUS_ASSIGN() != nil {
			if isFloat { finalVal = g.ctx.Builder.CreateFAdd(currVal, rhs, "") } else { finalVal = g.ctx.Builder.CreateAdd(currVal, rhs, "") }
		} else if op.MINUS_ASSIGN() != nil {
			if isFloat { finalVal = g.ctx.Builder.CreateFSub(currVal, rhs, "") } else { finalVal = g.ctx.Builder.CreateSub(currVal, rhs, "") }
		} else if op.STAR_ASSIGN() != nil {
			if isFloat { finalVal = g.ctx.Builder.CreateFMul(currVal, rhs, "") } else { finalVal = g.ctx.Builder.CreateMul(currVal, rhs, "") }
		} else if op.SLASH_ASSIGN() != nil {
			if isFloat { finalVal = g.ctx.Builder.CreateFDiv(currVal, rhs, "") } else { finalVal = g.ctx.Builder.CreateSDiv(currVal, rhs, "") }
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
	
	// Ensure boolean
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

	// Gen Then
	g.ctx.SetInsertBlock(thenBlock)
	g.Visit(ctx.Block(0))
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(mergeBlock)
	}

	// Gen Else
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

// ... include other statement visitors (For, Switch, etc.) as they were ...
// (Omitting unmodified long visitors for brevity, copy them from your existing file)
func (g *Generator) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	// (Same as original)
	if ctx.IN() != nil { return g.visitForInLoop(ctx) }
	condBlock := g.ctx.Builder.CreateBlock("loop.cond")
	bodyBlock := g.ctx.Builder.CreateBlock("loop.body")
	postBlock := g.ctx.Builder.CreateBlock("loop.post")
	endBlock := g.ctx.Builder.CreateBlock("loop.end")
	if ctx.VariableDecl() != nil { g.Visit(ctx.VariableDecl()) } else if len(ctx.AllAssignmentStmt()) > 0 && len(ctx.AllSEMICOLON()) > 0 { g.Visit(ctx.AssignmentStmt(0)) }
	g.ctx.Builder.CreateBr(condBlock)
	g.ctx.SetInsertBlock(condBlock)
	var cond ir.Value = g.ctx.Builder.ConstInt(types.I1, 1)
	if len(ctx.AllExpression()) > 0 { cond = g.Visit(ctx.Expression(0)).(ir.Value) }
	g.ctx.Builder.CreateCondBr(cond, bodyBlock, endBlock)
	g.ctx.SetInsertBlock(bodyBlock)
	g.loopStack = append(g.loopStack, loopInfo{breakBlock: endBlock, continueBlock: postBlock})
	g.Visit(ctx.Block())
	g.loopStack = g.loopStack[:len(g.loopStack)-1]
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil { g.ctx.Builder.CreateBr(postBlock) }
	g.ctx.SetInsertBlock(postBlock)
	if len(ctx.AllAssignmentStmt()) > 0 { g.Visit(ctx.AssignmentStmt(len(ctx.AllAssignmentStmt()) - 1)) }
	g.ctx.Builder.CreateBr(condBlock)
	g.ctx.SetInsertBlock(endBlock)
	return nil
}

func (g *Generator) visitForInLoop(ctx *parser.ForStmtContext) interface{} {
    // (Same as original)
	return nil // placeholder, use original impl
}

func (g *Generator) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	if len(g.loopStack) > 0 { g.ctx.Builder.CreateBr(g.loopStack[len(g.loopStack)-1].breakBlock) }
	return nil
}

func (g *Generator) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {
	if len(g.loopStack) > 0 { g.ctx.Builder.CreateBr(g.loopStack[len(g.loopStack)-1].continueBlock) }
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
    // (Same as original)
    return nil // placeholder
}

func (g *Generator) VisitTryStmt(ctx *parser.TryStmtContext) interface{} {
    // (Same as original)
    return nil // placeholder
}

func (g *Generator) VisitThrowStmt(ctx *parser.ThrowStmtContext) interface{} {
	val := g.Visit(ctx.Expression()).(ir.Value)
	// (Same as original)
	g.ctx.Builder.CreateRet(g.ctx.Builder.ConstInt(types.I32, -1))
	return val
}