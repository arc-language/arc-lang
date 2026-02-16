package codegen

import (
	"fmt"
	"github.com/arc-language/arc-lang/ast"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
)

func (cg *Codegen) declareFunction(fn *ast.FuncDecl) {
	var paramTypes []types.Type
	for _, p := range fn.Params {
		paramTypes = append(paramTypes, cg.TypeGen.GenType(p.Type))
	}
	
	retType := cg.TypeGen.GenType(fn.ReturnType)
	
	// CreateFunction adds it to the module
	cg.Builder.CreateFunction(fn.Name, retType, paramTypes, false)
}

func (cg *Codegen) genFuncBody(fn *ast.FuncDecl) error {
	irFn := cg.Module.GetFunction(fn.Name)
	if irFn == nil {
		return fmt.Errorf("function %s not declared", fn.Name)
	}

	// 1. Entry Block
	entry := cg.Builder.CreateBlockInFunction("entry", irFn)
	cg.Builder.SetInsertPoint(entry)

	cg.pushScope()
	defer cg.popScope()

	// 2. Argument Allocation (Stack Shadowing)
	// We alloc stack slots for args so they are mutable L-values
	for i, param := range fn.Params {
		argVal := irFn.Arguments[i]
		
		// Create stack slot: %arg.addr = alloca Type
		alloca := cg.Builder.CreateAlloca(argVal.Type(), param.Name+".addr")
		
		// Store argument: store %arg, %arg.addr
		cg.Builder.CreateStore(argVal, alloca)
		
		// Register in scope
		cg.defineVar(param.Name, alloca)
	}

	// 3. Body
	if fn.Body != nil {
		if err := cg.genBlock(fn.Body); err != nil {
			return err
		}
	}

	// 4. Default Return
	if entry.Terminator() == nil {
		cg.Builder.CreateRetVoid()
	}

	return nil
}

func (cg *Codegen) genBlock(block *ast.BlockStmt) error {
	for _, stmt := range block.List {
		if err := cg.genStmt(stmt); err != nil {
			return err
		}
	}
	return nil
}

func (cg *Codegen) genStmt(stmt ast.Stmt) error {
	switch s := stmt.(type) {
	case *ast.DeclStmt:
		// var x: T = val
		// 1. Generate value
		initVal := cg.genExpr(s.Decl.Value)
		
		// 2. Allocate Stack Slot
		typ := initVal.Type()
		alloca := cg.Builder.CreateAlloca(typ, s.Decl.Name)
		
		// 3. Store
		cg.Builder.CreateStore(initVal, alloca)
		
		// 4. Track
		cg.defineVar(s.Decl.Name, alloca)

	case *ast.ReturnStmt:
		if len(s.Results) > 0 {
			val := cg.genExpr(s.Results[0])
			cg.Builder.CreateRet(val)
		} else {
			cg.Builder.CreateRetVoid()
		}

	case *ast.ExprStmt:
		cg.genExpr(s.X)

	case *ast.IfStmt:
		return cg.genIf(s)
	
	case *ast.ForStmt:
		return cg.genFor(s)
	}
	return nil
}

func (cg *Codegen) genIf(s *ast.IfStmt) error {
	cond := cg.genExpr(s.Cond)
	
	thenBlock := cg.Builder.CreateBlock("if.then")
	elseBlock := cg.Builder.CreateBlock("if.else")
	endBlock := cg.Builder.CreateBlock("if.end")

	targetElse := elseBlock
	if s.Else == nil {
		targetElse = endBlock
	}
	cg.Builder.CreateCondBr(cond, thenBlock, targetElse)

	// Then
	cg.Builder.SetInsertPoint(thenBlock)
	cg.genBlock(s.Body)
	if cg.Builder.CurrentBlock().Terminator() == nil {
		cg.Builder.CreateBr(endBlock)
	}

	// Else
	if s.Else != nil {
		cg.Builder.SetInsertPoint(elseBlock)
		if block, ok := s.Else.(*ast.BlockStmt); ok {
			cg.genBlock(block)
		} else if ifStmt, ok := s.Else.(*ast.IfStmt); ok {
			cg.genIf(ifStmt)
		}
		if cg.Builder.CurrentBlock().Terminator() == nil {
			cg.Builder.CreateBr(endBlock)
		}
	}

	cg.Builder.SetInsertPoint(endBlock)
	return nil
}

func (cg *Codegen) genFor(s *ast.ForStmt) error {
	cg.pushScope()
	defer cg.popScope()

	if s.Init != nil {
		cg.genStmt(s.Init)
	}

	condBlock := cg.Builder.CreateBlock("loop.cond")
	bodyBlock := cg.Builder.CreateBlock("loop.body")
	postBlock := cg.Builder.CreateBlock("loop.post")
	endBlock := cg.Builder.CreateBlock("loop.end")

	cg.Builder.CreateBr(condBlock)

	// Condition
	cg.Builder.SetInsertPoint(condBlock)
	if s.Cond != nil {
		cond := cg.genExpr(s.Cond)
		cg.Builder.CreateCondBr(cond, bodyBlock, endBlock)
	} else {
		cg.Builder.CreateBr(bodyBlock)
	}

	// Body
	cg.Builder.SetInsertPoint(bodyBlock)
	cg.genBlock(s.Body)
	if cg.Builder.CurrentBlock().Terminator() == nil {
		cg.Builder.CreateBr(postBlock)
	}

	// Post
	cg.Builder.SetInsertPoint(postBlock)
	if s.Post != nil {
		cg.genStmt(s.Post)
	}
	cg.Builder.CreateBr(condBlock)

	cg.Builder.SetInsertPoint(endBlock)
	return nil
}