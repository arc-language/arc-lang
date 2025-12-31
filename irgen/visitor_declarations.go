package irgen

import (
	"github.com/arc-language/arc-lang/builder/ir"
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
	return nil
}

func (g *Generator) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	sym, _ := g.currentScope.Resolve(name)
	
	var paramTypes []types.Type
	var paramNames []string
	
	if ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			pType := g.resolveType(param.Type_())
			paramTypes = append(paramTypes, pType)
			paramNames = append(paramNames, param.IDENTIFIER().GetText())
		}
	}

	// FIX: Explicitly define variable type as interface to allow assignment
	var retType types.Type = types.Void
	if ctx.ReturnType() != nil {
		retType = g.resolveType(ctx.ReturnType().Type_())
	}

	// Create IR Function
	fn := g.ctx.Builder.CreateFunction(name, retType, paramTypes, false)
	g.ctx.EnterFunction(fn)

	// Update Symbol with the IR Function
	sym.IRValue = fn

	// Enter Function Scope
	g.enterScope(ctx)
	defer g.exitScope()

	// Create Allocas for Parameters
	entryBlock := g.ctx.Builder.GetInsertBlock()
	g.ctx.SetInsertBlock(entryBlock)
	
	for i, arg := range fn.Arguments {
		arg.SetName(paramNames[i])
		
		alloca := g.ctx.Builder.CreateAlloca(arg.Type(), paramNames[i]+".addr")
		g.ctx.Builder.CreateStore(arg, alloca)
		
		if paramSym, ok := g.currentScope.Resolve(paramNames[i]); ok {
			paramSym.IRValue = alloca
		}
	}

	// Generate Body
	if ctx.Block() != nil {
		g.deferStack = NewDeferStack()
		
		for _, stmt := range ctx.Block().AllStatement() {
			g.Visit(stmt)
		}
	}

	// Ensure Terminator
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		if retType == types.Void {
			g.ctx.Builder.CreateRetVoid()
		} else {
			g.ctx.Builder.CreateRet(g.ctx.Builder.ConstZero(retType))
		}
	}

	g.ctx.ExitFunction()
	return nil
}

func (g *Generator) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	sym, ok := g.currentScope.Resolve(name)
	if !ok { return nil }
	
	alloca := g.ctx.Builder.CreateAlloca(sym.Type, name+".addr")
	sym.IRValue = alloca 
	
	if ctx.Expression() != nil {
		// FIX: ir package is now imported, so ir.Value works
		val := g.Visit(ctx.Expression()).(ir.Value)
		val = g.emitCast(val, sym.Type)
		g.ctx.Builder.CreateStore(val, alloca)
	} else {
		g.ctx.Builder.CreateStore(g.ctx.Builder.ConstZero(sym.Type), alloca)
	}
	
	return nil
}