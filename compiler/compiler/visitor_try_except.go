package compiler

import (
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/parser"
)

func (v *IRVisitor) VisitTryStmt(ctx *parser.TryStmtContext) interface{} {
	token := ctx.GetStart()
	uniqueID := fmt.Sprintf("%d_%d", token.GetLine(), token.GetColumn())
	
	tryBlock := v.ctx.Builder.CreateBlock("try." + uniqueID)
	catchBlock := v.ctx.Builder.CreateBlock("catch." + uniqueID)
	finallyBlock := v.ctx.Builder.CreateBlock("finally." + uniqueID)
	endBlock := v.ctx.Builder.CreateBlock("try.end." + uniqueID)
	
	v.ctx.Builder.CreateBr(tryBlock)
	
	// Try block
	v.ctx.SetInsertBlock(tryBlock)
	v.Visit(ctx.Block())
	
	if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
		if ctx.FinallyClause() != nil {
			v.ctx.Builder.CreateBr(finallyBlock)
		} else {
			v.ctx.Builder.CreateBr(endBlock)
		}
	}
	
	// Except blocks
	if len(ctx.AllExceptClause()) > 0 {
		v.ctx.SetInsertBlock(catchBlock)
		for _, except := range ctx.AllExceptClause() {
			v.Visit(except)
		}
		if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
			if ctx.FinallyClause() != nil {
				v.ctx.Builder.CreateBr(finallyBlock)
			} else {
				v.ctx.Builder.CreateBr(endBlock)
			}
		}
	}
	
	// Finally block
	if ctx.FinallyClause() != nil {
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

func (v *IRVisitor) VisitExceptClause(ctx *parser.ExceptClauseContext) interface{} {
	v.Visit(ctx.Block())
	return nil
}