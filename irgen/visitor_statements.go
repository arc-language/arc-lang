package irgen

import (
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/parser"
)

func (g *Generator) VisitBlock(ctx *parser.BlockContext) interface{} {
	if _, isMapped := g.analysis.Scopes[ctx]; isMapped {
		g.enterScope(ctx)
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
	var val ir.Value
	if ctx.Expression() != nil {
		val = g.Visit(ctx.Expression()).(ir.Value)
		
		if g.ctx.CurrentFunction != nil {
			retType := g.ctx.CurrentFunction.FuncType.ReturnType
			val = g.emitCast(val, retType)
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

func (g *Generator) VisitDeferStmt(ctx *parser.DeferStmtContext) interface{} {
	g.deferStack.Add(func(gen *Generator) {
		if ctx.Expression() != nil { gen.Visit(ctx.Expression()) }
		if ctx.AssignmentStmt() != nil { gen.Visit(ctx.AssignmentStmt()) }
		// Removed ctx.Block() check as it is not part of the DeferStmtContext
	})
	return nil
}

func (g *Generator) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	cond := g.Visit(ctx.Expression(0)).(ir.Value)
	
	if cond.Type().BitSize() > 1 {
		cond = g.ctx.Builder.CreateICmpNE(cond, g.ctx.Builder.ConstZero(cond.Type()), "")
	}

	thenBlock := g.ctx.Builder.CreateBlock("if.then")
	elseBlock := g.ctx.Builder.CreateBlock("if.else")
	mergeBlock := g.ctx.Builder.CreateBlock("if.end")
	
	hasElse := len(ctx.AllBlock()) > 1
	targetElse := elseBlock
	if !hasElse {
		targetElse = mergeBlock
	}

	g.ctx.Builder.CreateCondBr(cond, thenBlock, targetElse)

	g.ctx.SetInsertBlock(thenBlock)
	g.Visit(ctx.Block(0))
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(mergeBlock)
	}

	if hasElse {
		g.ctx.SetInsertBlock(elseBlock)
		g.Visit(ctx.Block(1))
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			g.ctx.Builder.CreateBr(mergeBlock)
		}
	}

	g.ctx.SetInsertBlock(mergeBlock)
	return nil
}