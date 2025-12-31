package irgen

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (g *Generator) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	for _, decl := range ctx.AllTopLevelDecl() {
		g.Visit(decl)
	}
	return nil
}

func (g *Generator) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext) interface{} {
	if ctx.FunctionDecl() != nil { return g.Visit(ctx.FunctionDecl()) }
	if ctx.VariableDecl() != nil { return g.Visit(ctx.VariableDecl()) }
	// Structs are types, usually handled in Pass 1, unless generating metadata
	return nil
}

func (g *Generator) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	// 1. Resolve Symbol to get signature
	sym, _ := g.currentScope.Resolve(name)
	
	// Note: In a robust implementation, SymFunc would hold the full FunctionType.
	// Here we reconstruct param types from AST or lookup helpers.
	var paramTypes []types.Type
	var paramNames []string
	
	if ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			// In Pass 2, resolving types should ideally just look up the Symbol for that type
			// but for now we re-resolve using the helper.
			pType := g.resolveType(param.Type_())
			paramTypes = append(paramTypes, pType)
			paramNames = append(paramNames, param.IDENTIFIER().GetText())
		}
	}

	retType := types.Void
	if ctx.ReturnType() != nil {
		retType = g.resolveType(ctx.ReturnType().Type_())
	}

	// 2. Create IR Function
	fn := g.ctx.Builder.CreateFunction(name, retType, paramTypes, false)
	g.ctx.EnterFunction(fn)

	// Update Symbol with the IR Function so calls can find it
	sym.IRValue = fn

	// 3. Enter Function Scope
	g.enterScope(ctx)
	defer g.exitScope()

	// 4. Create Allocas for Parameters
	// This makes parameters mutable local variables
	entryBlock := g.ctx.Builder.GetInsertBlock()
	g.ctx.SetInsertBlock(entryBlock)
	
	for i, arg := range fn.Arguments {
		arg.SetName(paramNames[i])
		
		// Create stack slot
		alloca := g.ctx.Builder.CreateAlloca(arg.Type(), paramNames[i]+".addr")
		g.ctx.Builder.CreateStore(arg, alloca)
		
		// LINKING: Find the parameter symbol created in Pass 1 and attach this Alloca
		if paramSym, ok := g.currentScope.Resolve(paramNames[i]); ok {
			paramSym.IRValue = alloca
		}
	}

	// 5. Generate Body
	if ctx.Block() != nil {
		// Reset defer stack for this function
		g.deferStack = NewDeferStack()
		
		for _, stmt := range ctx.Block().AllStatement() {
			g.Visit(stmt)
		}
	}

	// 6. Ensure Terminator
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		if retType == types.Void {
			g.ctx.Builder.CreateRetVoid()
		} else {
			// Should be caught by Pass 1, but safe fallback
			g.ctx.Builder.CreateRet(g.ctx.Builder.ConstZero(retType))
		}
	}

	g.ctx.ExitFunction()
	return nil
}

func (g *Generator) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	// 1. Get Symbol
	sym, ok := g.currentScope.Resolve(name)
	if !ok { return nil } // Should catch in Pass 1
	
	// 2. Create Alloca
	alloca := g.ctx.Builder.CreateAlloca(sym.Type, name+".addr")
	sym.IRValue = alloca // Link for future lookups
	
	// 3. Init
	if ctx.Expression() != nil {
		val := g.Visit(ctx.Expression()).(ir.Value)
		val = g.emitCast(val, sym.Type)
		g.ctx.Builder.CreateStore(val, alloca)
	} else {
		g.ctx.Builder.CreateStore(g.ctx.Builder.ConstZero(sym.Type), alloca)
	}
	
	return nil
}