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
	if ctx.FunctionDecl() != nil {
		return g.Visit(ctx.FunctionDecl())
	}
	if ctx.VariableDecl() != nil {
		return g.Visit(ctx.VariableDecl())
	}
	return nil
}

func (g *Generator) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	if ctx.GenericParams() != nil {
		return nil
	}

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

	// FIX: Use explicit interface type to match resolveType return
	var retType types.Type = types.Void
	if ctx.ReturnType() != nil {
		retType = g.resolveType(ctx.ReturnType().Type_())
	}

	// Create IR Function
	fn := g.ctx.Builder.CreateFunction(name, retType, paramTypes, false)
	g.ctx.EnterFunction(fn)
	
	if sym != nil {
		sym.IRValue = fn
	}

	g.enterScope(ctx)
	defer g.exitScope()

	entry := g.ctx.Builder.GetInsertBlock()
	g.ctx.SetInsertBlock(entry)

	// Args
	for i, arg := range fn.Arguments {
		arg.SetName(paramNames[i])
		alloca := g.ctx.Builder.CreateAlloca(arg.Type(), paramNames[i]+".addr")
		g.ctx.Builder.CreateStore(arg, alloca)

		if s, ok := g.currentScope.Resolve(paramNames[i]); ok {
			s.IRValue = alloca
		}
	}

	if ctx.Block() != nil {
		g.deferStack = NewDeferStack()
		for _, stmt := range ctx.Block().AllStatement() {
			g.Visit(stmt)
		}
	}

	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		if retType == types.Void {
			g.ctx.Builder.CreateRetVoid()
		} else {
			g.ctx.Builder.CreateRet(g.getZeroValue(retType))
		}
	}

	g.ctx.ExitFunction()
	return nil
}

func (g *Generator) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	sym, ok := g.currentScope.Resolve(name)
	if !ok {
		return nil
	}

	alloca := g.ctx.Builder.CreateAlloca(sym.Type, name+".addr")
	sym.IRValue = alloca

	if ctx.Expression() != nil {
		val := g.Visit(ctx.Expression())
		
		// FIX: Correct type assertion for IR Value
		if irVal, ok := val.(ir.Value); ok {
			irVal = g.emitCast(irVal, sym.Type)
			g.ctx.Builder.CreateStore(irVal, alloca)
		}
	} else {
		g.ctx.Builder.CreateStore(g.getZeroValue(sym.Type), alloca)
	}
	return nil
}