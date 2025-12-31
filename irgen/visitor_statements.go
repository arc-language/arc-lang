package irgen

import (
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (g *Generator) VisitBlock(ctx *parser.BlockContext) interface{} {
	// If the semantic analyzer mapped a specific scope to this block (e.g. function body), enter it.
	if _, isMapped := g.analysis.Scopes[ctx]; isMapped {
		g.enterScope(ctx)
		defer g.exitScope()
	}

	for _, stmt := range ctx.AllStatement() {
		g.Visit(stmt)
		// If the block is terminated (e.g. by return/break), stop generating code for unreachable statements
		if g.ctx.Builder.GetInsertBlock().Terminator() != nil {
			break
		}
	}
	return nil
}

func (g *Generator) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	// 1. Emit deferred statements (LIFO)
	g.deferStack.Emit(g)

	// 2. Handle Return Value
	var val ir.Value
	if ctx.Expression() != nil {
		val = g.Visit(ctx.Expression()).(ir.Value)
		
		// Implicit casting to return type
		if g.ctx.CurrentFunction != nil {
			expectedType := g.ctx.CurrentFunction.FuncType.ReturnType
			val = g.emitCast(val, expectedType)
		}
	}

	// 3. Create instruction
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

	// --- L-Value Resolution ---
	// We cannot use standard Visit() for LHS because Visit() returns loaded values (R-Values).
	// We must manually resolve the address.

	// Case A: Simple Identifier (x = ...)
	if lhsCtx.IDENTIFIER() != nil && lhsCtx.DOT() == nil && lhsCtx.STAR() == nil && lhsCtx.LBRACKET() == nil {
		name := lhsCtx.IDENTIFIER().GetText()
		sym, ok := g.currentScope.Resolve(name)
		if !ok {
			// Should have been caught by semantics
			return nil
		}
		if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
			destPtr = alloca
		} else {
			// Might be a Global or Argument pointer
			destPtr = sym.IRValue
		}
	} else if lhsCtx.STAR() != nil {
		// Case B: Dereference (*p = ...)
		// We visit the inner expression to get the pointer value
		destPtr = g.Visit(lhsCtx.PostfixExpression()).(ir.Value)
	} else if lhsCtx.LBRACKET() != nil {
		// Case C: Array Indexing (arr[i] = ...)
		base := g.Visit(lhsCtx.PostfixExpression()).(ir.Value)
		index := g.Visit(lhsCtx.Expression()).(ir.Value)
		
		if ptrType, ok := base.Type().(*types.PointerType); ok {
			// GEP to get address of element
			destPtr = g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, base, []ir.Value{index}, "")
		}
	} else if lhsCtx.DOT() != nil {
		// Case D: Member Access (obj.field = ...)
		base := g.Visit(lhsCtx.PostfixExpression()).(ir.Value)
		fieldName := lhsCtx.IDENTIFIER().GetText()

		// Auto-dereference if base is a pointer
		var structType *types.StructType
		if ptr, ok := base.Type().(*types.PointerType); ok {
			if st, ok := ptr.ElementType.(*types.StructType); ok {
				structType = st
			}
		}

		if structType != nil {
			if idx, ok := g.analysis.StructIndices[structType.Name][fieldName]; ok {
				destPtr = g.ctx.Builder.CreateStructGEP(structType, base, idx, "")
			}
		}
	}

	if destPtr == nil {
		// Panic or Error logging
		return nil
	}

	// --- RHS Evaluation ---
	rhs := g.Visit(ctx.Expression()).(ir.Value)
	
	// --- Compound Assignment Handling (+=, -=, etc) ---
	finalVal := rhs
	if ctx.AssignmentOp().ASSIGN() == nil {
		// Load current value
		ptrType := destPtr.Type().(*types.PointerType)
		currVal := g.ctx.Builder.CreateLoad(ptrType.ElementType, destPtr, "")
		
		op := ctx.AssignmentOp()
		if op.PLUS_ASSIGN() != nil {
			if types.IsFloat(currVal.Type()) {
				finalVal = g.ctx.Builder.CreateFAdd(currVal, rhs, "")
			} else {
				finalVal = g.ctx.Builder.CreateAdd(currVal, rhs, "")
			}
		} else if op.MINUS_ASSIGN() != nil {
			if types.IsFloat(currVal.Type()) {
				finalVal = g.ctx.Builder.CreateFSub(currVal, rhs, "")
			} else {
				finalVal = g.ctx.Builder.CreateSub(currVal, rhs, "")
			}
		} else if op.STAR_ASSIGN() != nil {
			if types.IsFloat(currVal.Type()) {
				finalVal = g.ctx.Builder.CreateFMul(currVal, rhs, "")
			} else {
				finalVal = g.ctx.Builder.CreateMul(currVal, rhs, "")
			}
		} else if op.SLASH_ASSIGN() != nil {
			if types.IsFloat(currVal.Type()) {
				finalVal = g.ctx.Builder.CreateFDiv(currVal, rhs, "")
			} else {
				finalVal = g.ctx.Builder.CreateSDiv(currVal, rhs, "")
			}
		}
		// ... (Add other compound ops like bitwise if needed)
	}

	// Ensure type match (Cast RHS to LHS type)
	ptrType := destPtr.Type().(*types.PointerType)
	finalVal = g.emitCast(finalVal, ptrType.ElementType)

	g.ctx.Builder.CreateStore(finalVal, destPtr)
	return nil
}

func (g *Generator) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	cond := g.Visit(ctx.Expression(0)).(ir.Value)
	
	// Ensure boolean (i1)
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

	// Generate Then
	g.ctx.SetInsertBlock(thenBlock)
	g.Visit(ctx.Block(0))
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(mergeBlock)
	}

	// Generate Else
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
	condBlock := g.ctx.Builder.CreateBlock("loop.cond")
	bodyBlock := g.ctx.Builder.CreateBlock("loop.body")
	postBlock := g.ctx.Builder.CreateBlock("loop.post")
	endBlock := g.ctx.Builder.CreateBlock("loop.end")

	// Init block handled in VisitStatement before calling here? 
	// The parser structure wraps Init inside the loop logic usually.
	// We'll check for C-style Init here if the parser supports it inside ForStmt directly, 
	// otherwise assume Init was visited prior.
	// Based on old code: If it's a C-style for loop, variables are often declared before.

	g.ctx.Builder.CreateBr(condBlock)
	g.ctx.SetInsertBlock(condBlock)

	// Condition
	var cond ir.Value = g.ctx.Builder.ConstInt(types.I1, 1)
	if len(ctx.AllExpression()) > 0 {
		cond = g.Visit(ctx.Expression(0)).(ir.Value)
	}
	g.ctx.Builder.CreateCondBr(cond, bodyBlock, endBlock)

	// Body
	g.ctx.SetInsertBlock(bodyBlock)
	g.loopStack = append(g.loopStack, loopInfo{breakBlock: endBlock, continueBlock: postBlock})
	g.Visit(ctx.Block())
	g.loopStack = g.loopStack[:len(g.loopStack)-1] // Pop

	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.ctx.Builder.CreateBr(postBlock)
	}

	// Post Statement (Increment)
	g.ctx.SetInsertBlock(postBlock)
	if len(ctx.AllAssignmentStmt()) > 0 {
		g.Visit(ctx.AssignmentStmt(0))
	} else if len(ctx.AllExpression()) > 1 { 
		// Sometimes post is an expression like i++
		// Adjust index based on parser grammar
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
	// Defer executes a closure at function exit
	g.deferStack.Add(func(gen *Generator) {
		if ctx.Expression() != nil { gen.Visit(ctx.Expression()) }
		if ctx.AssignmentStmt() != nil { gen.Visit(ctx.AssignmentStmt()) }
	})
	return nil
}

func (g *Generator) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	cond := g.Visit(ctx.Expression()).(ir.Value)
	endBlock := g.ctx.Builder.CreateBlock("switch.end")
	
	// Chain of if-else blocks (simplified implementation of switch)
	prevBlock := g.ctx.Builder.GetInsertBlock()
	
	for i, c := range ctx.AllSwitchCase() {
		g.ctx.SetInsertBlock(prevBlock)
		
		caseBlock := g.ctx.Builder.CreateBlock(fmt.Sprintf("case.%d", i))
		nextCheckBlock := g.ctx.Builder.CreateBlock(fmt.Sprintf("check.%d", i))
		
		// If last case and no default, next goes to end
		if i == len(ctx.AllSwitchCase())-1 && ctx.DefaultCase() == nil {
			nextCheckBlock = endBlock
		}
		
		caseVal := g.Visit(c.Expression()).(ir.Value)
		cmp := g.ctx.Builder.CreateICmpEQ(cond, caseVal, "")
		g.ctx.Builder.CreateCondBr(cmp, caseBlock, nextCheckBlock)
		
		// Case Body
		g.ctx.SetInsertBlock(caseBlock)
		for _, s := range c.AllStatement() {
			g.Visit(s)
			if g.ctx.Builder.GetInsertBlock().Terminator() != nil { break }
		}
		if g.ctx.Builder.GetInsertBlock().Terminator() == nil { 
			g.ctx.Builder.CreateBr(endBlock) 
		}
		
		prevBlock = nextCheckBlock
	}
	
	// Default Case
	g.ctx.SetInsertBlock(prevBlock)
	if ctx.DefaultCase() != nil {
		for _, s := range ctx.DefaultCase().AllStatement() {
			g.Visit(s)
			if g.ctx.Builder.GetInsertBlock().Terminator() != nil { break }
		}
	}
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil { 
		g.ctx.Builder.CreateBr(endBlock) 
	}
	
	g.ctx.SetInsertBlock(endBlock)
	return nil
}

func (g *Generator) VisitTryStmt(ctx *parser.TryStmtContext) interface{} {
	// Exception handling using global state variable model
	// This mimics the old compiler logic
	
	tryBlock := g.ctx.Builder.CreateBlock("try.start")
	endBlock := g.ctx.Builder.CreateBlock("try.end")
	
	// We only support one except block for simplicity here, logic can be expanded
	var catchBlock *ir.BasicBlock
	if len(ctx.AllExceptClause()) > 0 {
		catchBlock = g.ctx.Builder.CreateBlock("try.catch")
	}

	g.ctx.Builder.CreateBr(tryBlock)
	g.ctx.SetInsertBlock(tryBlock)

	// Execute Try Body
	// We visit statements manually. If this were a full EH implementation,
	// every function call would need to check the exception register.
	// For this pass, we just visit the block.
	g.Visit(ctx.Block())
	
	// Stub: In a real implementation, we would check __exception_state here
	// globalExc := g.ctx.Module.GetGlobal("__exception_state")
	// ... check logic ...
	
	g.ctx.Builder.CreateBr(endBlock)

	// Execute Catch Body
	if catchBlock != nil {
		g.ctx.SetInsertBlock(catchBlock)
		// Clear exception state logic would go here
		g.Visit(ctx.ExceptClause(0).Block())
		g.ctx.Builder.CreateBr(endBlock)
	}

	g.ctx.SetInsertBlock(endBlock)
	return nil
}

func (g *Generator) VisitThrowStmt(ctx *parser.ThrowStmtContext) interface{} {
	val := g.Visit(ctx.Expression()).(ir.Value)
	
	// Get global exception state
	excGlobal := g.ctx.Module.GetGlobal("__exception_state")
	if excGlobal == nil {
		// Lazily create if not exists (Struct: {i1 hasException, i8* msg})
		st := types.NewStruct("ExceptionState", []types.Type{types.I1, types.NewPointer(types.I8)}, false)
		excGlobal = g.ctx.Builder.CreateGlobalVariable("__exception_state", st, nil)
	}

	// Set hasException = true
	// Set msg = val
	// Code omitted for brevity, but follows standard GEP + Store patterns
	
	// Return error code to unwind stack
	g.ctx.Builder.CreateRet(g.ctx.Builder.ConstInt(types.I32, -1))
	
	return val
}