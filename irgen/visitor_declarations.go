package irgen

import (
	"fmt"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (g *Generator) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	for _, ns := range ctx.AllNamespaceDecl() { g.Visit(ns) }
	for _, decl := range ctx.AllTopLevelDecl() { g.Visit(decl) }
	return nil
}

func (g *Generator) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext) interface{} {
	if ctx.FunctionDecl() != nil { return g.Visit(ctx.FunctionDecl()) }
	if ctx.VariableDecl() != nil { return g.Visit(ctx.VariableDecl()) }
	if ctx.ExternDecl() != nil { return g.Visit(ctx.ExternDecl()) }
	return nil
}

func (g *Generator) VisitNamespaceDecl(ctx *parser.NamespaceDeclContext) interface{} {
	if ctx.IDENTIFIER() != nil { g.currentNamespace = ctx.IDENTIFIER().GetText() }
	return nil
}

func (g *Generator) VisitExternDecl(ctx *parser.ExternDeclContext) interface{} {
	oldNs := g.currentNamespace
	if ctx.IDENTIFIER() != nil { g.currentNamespace = ctx.IDENTIFIER().GetText() }
	for _, member := range ctx.AllExternMember() {
		if fnDecl := member.ExternFunctionDecl(); fnDecl != nil {
			g.visitExternFunctionDecl(fnDecl.(*parser.ExternFunctionDeclContext))
		}
	}
	g.currentNamespace = oldNs
	return nil
}

func (g *Generator) visitExternFunctionDecl(ctx *parser.ExternFunctionDeclContext) {
	name := ctx.IDENTIFIER().GetText()
	externalName := name
	if ctx.STRING_LITERAL() != nil {
		raw := ctx.STRING_LITERAL().GetText()
		if len(raw) >= 2 { externalName = raw[1 : len(raw)-1] }
	}
	
	var retType types.Type = types.Void
	if ctx.Type_() != nil { retType = g.resolveType(ctx.Type_()) }
	
	var paramTypes []types.Type
	variadic := false
	if ctx.ExternParameterList() != nil {
		if ctx.ExternParameterList().ELLIPSIS() != nil { variadic = true }
		for _, t := range ctx.ExternParameterList().AllType_() {
			paramTypes = append(paramTypes, g.resolveType(t))
		}
	}
	
	fn := g.ctx.Builder.DeclareFunction(externalName, retType, paramTypes, variadic)
	
	// Register IRValue in the symbol table
	if sym, ok := g.currentScope.Resolve(name); ok {
		sym.IRValue = fn
	}
	
	// Also register under namespace if applicable
	if g.currentNamespace != "" {
		fullName := g.currentNamespace + "." + name
		if sym, ok := g.currentScope.Resolve(fullName); ok {
			sym.IRValue = fn
		}
	}
}

func (g *Generator) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	// Handle main vs namespaced functions
	irName := name
	if g.currentNamespace != "" && name != "main" {
		irName = g.currentNamespace + "_" + name
	}
	
	sym, _ := g.currentScope.Resolve(name)
	
	var paramTypes []types.Type
	var paramNames []string
	if ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			paramTypes = append(paramTypes, g.resolveType(param.Type_()))
			paramNames = append(paramNames, param.IDENTIFIER().GetText())
		}
	}
	
	var retType types.Type = types.Void
	if ctx.ReturnType() != nil { retType = g.resolveType(ctx.ReturnType().Type_()) }
	
	fn := g.ctx.Builder.CreateFunction(irName, retType, paramTypes, false)
	
	// Update symbol with IR function
	if sym != nil { sym.IRValue = fn }
	
	g.ctx.EnterFunction(fn)
	g.enterScope(ctx) // Enters function scope
	defer g.exitScope()
	
	// Create allocas for arguments
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
		g.Visit(ctx.Block())
	}
	
	// Implicit return
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
	
	// Critical Debug Check
	sym, ok := g.currentScope.Resolve(name)
	if !ok {
		fmt.Printf("[IRGen] Error: Variable '%s' was not defined in the symbol table (Semantic pass failure?)\n", name)
		return nil
	}
	
	alloca := g.ctx.Builder.CreateAlloca(sym.Type, name+".addr")
	sym.IRValue = alloca
	
	if ctx.Expression() != nil {
		val := g.Visit(ctx.Expression())
		if irVal, ok := val.(ir.Value); ok {
			irVal = g.emitCast(irVal, sym.Type)
			g.ctx.Builder.CreateStore(irVal, alloca)
		} else {
			// Fallback for failed expression gen
			g.ctx.Builder.CreateStore(g.getZeroValue(sym.Type), alloca)
		}
	} else {
		g.ctx.Builder.CreateStore(g.getZeroValue(sym.Type), alloca)
	}
	return nil
}