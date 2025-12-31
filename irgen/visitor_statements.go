package irgen

import (
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (g *Generator) VisitBlock(ctx *parser.BlockContext) interface{} {
	if _, isMapped := g.analysis.Scopes[ctx]; isMapped {
		g.enterScope(ctx)
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
	var val ir.Value
	if ctx.Expression() != nil {
		val = g.Visit(ctx.Expression()).(ir.Value)
		
		// Cast return value if needed (e.g. returning i32 in i64 func)
		if g.ctx.CurrentFunction != nil {
			retType := g.ctx.CurrentFunction.FuncType.ReturnType
			val = g.emitCast(val, retType)
		}
	}

	// EXECUTE DEFERS before returning
	g.deferStack.Emit(g)

	if val != nil {
		g.ctx.Builder.CreateRet(val)
	} else {
		g.ctx.Builder.CreateRetVoid()
	}
	return nil
}

func (g *Generator) VisitDeferStmt(ctx *parser.DeferStmtContext) interface{} {
	// Capture the AST node in a closure
	g.deferStack.Add(func(gen *Generator) {
		if ctx.Expression() != nil { gen.Visit(ctx.Expression()) }
		if ctx.AssignmentStmt() != nil { gen.Visit(ctx.AssignmentStmt()) }
		if ctx.Block() != nil { gen.Visit(ctx.Block()) }
	})
	return nil
}

func (g *Generator) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	// 1. Condition
	cond := g.Visit(ctx.Expression(0)).(ir.Value)
	
	// Ensure boolean
	if cond.Type().BitSize() > 1 {
		cond = g.ctx.Builder.CreateICmpNE(cond, g.ctx.Builder.ConstZero(cond.Type()), "")
	}

	// 2. Blocks
	thenBlock := g.ctx.Builder.CreateBlock("if.then")
	elseBlock := g.ctx.Builder.CreateBlock("if.else")
	mergeBlock := g.ctx.Builder.CreateBlock("if.end")
	
	// Optimization: If no else, jump to merge
	hasElse := len(ctx.AllBlock()) > 1
	targetElse := elseBlock
	if !hasElse {
		targetElse = mergeBlock
	}

	g.ctx.Builder.CreateCondBr(cond, thenBlock, targetElse)

	// 3. Emit Then
	g.ctx.SetInsertBlock(thenBlock)
	g.Visit(ctx.Block(0))
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(mergeBlock)
	}

	// 4. Emit Else
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