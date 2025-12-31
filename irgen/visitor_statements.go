package irgen

import (
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/parser"
)

// Loop info tracking
type loopInfo struct {
	breakBlock    *ir.BasicBlock
	continueBlock *ir.BasicBlock
}

func (g *Generator) VisitBlock(ctx *parser.BlockContext) interface{} {
	if _, isMapped := g.analysis.Scopes[ctx]; isMapped {
		g.enterScope(ctx)
		defer g.exitScope()
	}
	for _, stmt := range ctx.AllStatement() {
		g.Visit(stmt)
		if g.ctx.Builder.GetInsertBlock().Terminator() != nil { break }
	}
	return nil
}

func (g *Generator) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	// Setup Blocks
	condBlock := g.ctx.Builder.CreateBlock("for.cond")
	bodyBlock := g.ctx.Builder.CreateBlock("for.body")
	postBlock := g.ctx.Builder.CreateBlock("for.post") // For 'i++'
	endBlock := g.ctx.Builder.CreateBlock("for.end")

	// 1. Init (e.g. i := 0)
	// Handled by visiting variable decls or assignments before loop start if they exist
	// In parsing, the init stmt is usually separate or before the loop structure logic
	// If the grammar puts init inside ForStmt, visit it here.
	if ctx.VariableDecl() != nil { g.Visit(ctx.VariableDecl()) }

	g.ctx.Builder.CreateBr(condBlock)

	// 2. Condition
	g.ctx.SetInsertBlock(condBlock)
	// Parse condition expression (default true)
	var cond ir.Value = g.ctx.Builder.True()
	if len(ctx.AllExpression()) > 0 {
		cond = g.Visit(ctx.Expression(0)).(ir.Value)
	}
	g.ctx.Builder.CreateCondBr(cond, bodyBlock, endBlock)

	// 3. Body
	g.ctx.SetInsertBlock(bodyBlock)
	// Push Loop Info for Break/Continue
	g.loopStack = append(g.loopStack, loopInfo{breakBlock: endBlock, continueBlock: postBlock})
	g.Visit(ctx.Block())
	g.loopStack = g.loopStack[:len(g.loopStack)-1] // Pop

	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(postBlock)
	}

	// 4. Post (e.g. i++)
	g.ctx.SetInsertBlock(postBlock)
	// If grammar has post statement (assignment), visit it
	if len(ctx.AllAssignmentStmt()) > 0 {
		g.Visit(ctx.AssignmentStmt(0))
	}
	g.ctx.Builder.CreateBr(condBlock)

	// 5. End
	g.ctx.SetInsertBlock(endBlock)
	return nil
}

func (g *Generator) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	if len(g.loopStack) > 0 {
		info := g.loopStack[len(g.loopStack)-1]
		g.ctx.Builder.CreateBr(info.breakBlock)
	}
	return nil
}

func (g *Generator) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {
	if len(g.loopStack) > 0 {
		info := g.loopStack[len(g.loopStack)-1]
		g.ctx.Builder.CreateBr(info.continueBlock)
	}
	return nil
}

func (g *Generator) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	var val ir.Value
	if ctx.Expression() != nil {
		val = g.Visit(ctx.Expression()).(ir.Value)
		// Cast if necessary
		if g.ctx.CurrentFunction != nil {
			val = g.emitCast(val, g.ctx.CurrentFunction.FuncType.ReturnType)
		}
	}
	g.deferStack.Emit(g)
	if val != nil {
		g.ctx.Builder.CreateRet(val)
	} else {
		g.ctx.Builder.CreateRetVoid()
	}
	return nil
}

func (g *Generator) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	cond := g.Visit(ctx.Expression(0)).(ir.Value)
	thenBlock := g.ctx.Builder.CreateBlock("if.then")
	mergeBlock := g.ctx.Builder.CreateBlock("if.end")
	
	elseBlock := mergeBlock
	if len(ctx.AllBlock()) > 1 {
		elseBlock = g.ctx.Builder.CreateBlock("if.else")
	}

	g.ctx.Builder.CreateCondBr(cond, thenBlock, elseBlock)

	// Then
	g.ctx.SetInsertBlock(thenBlock)
	g.Visit(ctx.Block(0))
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil { g.ctx.Builder.CreateBr(mergeBlock) }

	// Else
	if elseBlock != mergeBlock {
		g.ctx.SetInsertBlock(elseBlock)
		g.Visit(ctx.Block(1))
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil { g.ctx.Builder.CreateBr(mergeBlock) }
	}

	g.ctx.SetInsertBlock(mergeBlock)
	return nil
}

func (g *Generator) VisitDeferStmt(ctx *parser.DeferStmtContext) interface{} {
	g.deferStack.Add(func(gen *Generator) {
		if ctx.Expression() != nil { gen.Visit(ctx.Expression()) }
		if ctx.AssignmentStmt() != nil { gen.Visit(ctx.AssignmentStmt()) }
		if ctx.Block() != nil { gen.Visit(ctx.Block()) }
	})
	return nil
}

func (g *Generator) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	token := ctx.GetStart()
	uniqueID := fmt.Sprintf("%d_%d", token.GetLine(), token.GetColumn())
	
	condVal := g.Visit(ctx.Expression()).(ir.Value)
	endBlock := g.ctx.Builder.CreateBlock("switch.end." + uniqueID)
	
	// Complex switches are often implemented as cascaded if-else in simple compilers
	// OR using the Switch instruction if operands are constant integers.
	// For robustness here, we'll do cascaded comparisons.
	
	prevCheckBlock := g.ctx.Builder.GetInsertBlock()
	
	for i, caseCtx := range ctx.AllSwitchCase() {
		g.ctx.SetInsertBlock(prevCheckBlock)
		
		caseBlock := g.ctx.Builder.CreateBlock(fmt.Sprintf("case.%d", i))
		nextCheckBlock := g.ctx.Builder.CreateBlock(fmt.Sprintf("check.%d", i+1))
		if i == len(ctx.AllSwitchCase())-1 && ctx.DefaultCase() == nil {
			nextCheckBlock = endBlock
		}

		// Compare
		caseVal := g.Visit(caseCtx.Expression()).(ir.Value)
		cmp := g.ctx.Builder.CreateICmpEQ(condVal, caseVal, "")
		g.ctx.Builder.CreateCondBr(cmp, caseBlock, nextCheckBlock)

		// Body
		g.ctx.SetInsertBlock(caseBlock)
		for _, stmt := range caseCtx.AllStatement() {
			g.Visit(stmt)
			if g.ctx.Builder.GetInsertBlock().Terminator() != nil { break }
		}
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			g.ctx.Builder.CreateBr(endBlock)
		}
		
		prevCheckBlock = nextCheckBlock
	}
	
	// Default
	g.ctx.SetInsertBlock(prevCheckBlock)
	if ctx.DefaultCase() != nil {
		for _, stmt := range ctx.DefaultCase().AllStatement() {
			g.Visit(stmt)
			if g.ctx.Builder.GetInsertBlock().Terminator() != nil { break }
		}
	}
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(endBlock)
	}

	g.ctx.SetInsertBlock(endBlock)
	return nil
}