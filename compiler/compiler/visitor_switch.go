package compiler

import (
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/parser"
)

func (v *IRVisitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	token := ctx.GetStart()
	uniqueID := fmt.Sprintf("%d_%d", token.GetLine(), token.GetColumn())
	
	// Evaluate switch expression once
	switchVal := v.Visit(ctx.Expression()).(ir.Value)
	
	// Create blocks
	endBlock := v.ctx.Builder.CreateBlock("switch.end." + uniqueID)
	defaultBlock := endBlock // Default goes to end if not specified
	
	if ctx.DefaultCase() != nil {
		defaultBlock = v.ctx.Builder.CreateBlock("switch.default." + uniqueID)
	}
	
	// Process each case
	cases := ctx.AllSwitchCase()
	for i, caseCtx := range cases {
		caseBlock := v.ctx.Builder.CreateBlock(fmt.Sprintf("switch.case.%s.%d", uniqueID, i))
		nextCheck := defaultBlock
		
		if i+1 < len(cases) {
			nextCheck = v.ctx.Builder.CreateBlock(fmt.Sprintf("switch.check.%s.%d", uniqueID, i+1))
		}
		
		// Compare switch value with case value
		caseVal := v.Visit(caseCtx.Expression()).(ir.Value)
		cmp := v.ctx.Builder.CreateICmpEQ(switchVal, caseVal, "")
		v.ctx.Builder.CreateCondBr(cmp, caseBlock, nextCheck)
		
		// Generate case body
		v.ctx.SetInsertBlock(caseBlock)
		for _, stmt := range caseCtx.AllStatement() {
			v.Visit(stmt)
			if v.ctx.Builder.GetInsertBlock().Terminator() != nil {
				break
			}
		}
		
		if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
			v.ctx.Builder.CreateBr(endBlock)
		}
		
		if nextCheck != defaultBlock {
			v.ctx.SetInsertBlock(nextCheck)
		}
	}
	
	// Default case
	if ctx.DefaultCase() != nil {
		v.ctx.SetInsertBlock(defaultBlock)
		for _, stmt := range ctx.DefaultCase().AllStatement() {
			v.Visit(stmt)
			if v.ctx.Builder.GetInsertBlock().Terminator() != nil {
				break
			}
		}
		if v.ctx.Builder.GetInsertBlock().Terminator() == nil {
			v.ctx.Builder.CreateBr(endBlock)
		}
	} else {
		v.ctx.Builder.CreateBr(endBlock)
	}
	
	v.ctx.SetInsertBlock(endBlock)
	return nil
}