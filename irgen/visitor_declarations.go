package irgen

import (
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

func (g *Generator) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	// Pass 1 (Semantics) already defined types in the AnalysisResult.
	// We just need to ensure the LLVM types are created/registered in the Module.
	
	// Handle Namespace Declarations
	for _, ns := range ctx.AllNamespaceDecl() {
		g.Visit(ns)
	}
	
	// Handle Top Level Declarations
	for _, decl := range ctx.AllTopLevelDecl() {
		g.Visit(decl)
	}
	return nil
}

func (g *Generator) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext) interface{} {
	if ctx.FunctionDecl() != nil { return g.Visit(ctx.FunctionDecl()) }
	if ctx.VariableDecl() != nil { return g.Visit(ctx.VariableDecl()) }
	if ctx.ExternDecl() != nil { return g.Visit(ctx.ExternDecl()) }
	if ctx.StructDecl() != nil { return g.Visit(ctx.StructDecl()) }
	if ctx.ClassDecl() != nil { return g.Visit(ctx.ClassDecl()) }
	if ctx.EnumDecl() != nil { return g.Visit(ctx.EnumDecl()) }
	return nil
}

func (g *Generator) VisitNamespaceDecl(ctx *parser.NamespaceDeclContext) interface{} {
	if ctx.IDENTIFIER() != nil {
		g.currentNamespace = ctx.IDENTIFIER().GetText()
	}
	return nil
}

func (g *Generator) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	// Struct types are pre-calculated in analysis, but we might need to process methods here
	for _, member := range ctx.AllStructMember() {
		if member.FunctionDecl() != nil {
			g.Visit(member.FunctionDecl())
		}
	}
	return nil
}

func (g *Generator) VisitClassDecl(ctx *parser.ClassDeclContext) interface{} {
	// Similar to struct, process methods
	for _, member := range ctx.AllClassMember() {
		if member.FunctionDecl() != nil {
			g.Visit(member.FunctionDecl())
		}
	}
	return nil
}

func (g *Generator) VisitEnumDecl(ctx *parser.EnumDeclContext) interface{} {
	// Create constants for enum members
	// Enums are essentially global integer constants
	enumName := ctx.IDENTIFIER().GetText()
	val := int64(0)
	
	for _, member := range ctx.AllEnumMember() {
		memName := member.IDENTIFIER().GetText()
		
		// Determine value
		if member.Expression() != nil {
			// Visit the expression to evaluate it (e.g. "1", "1 + 2", etc.)
			// We expect the visitor to return an ir.Value (specifically a ConstantInt)
			exprVal := g.Visit(member.Expression())
			
			if irVal, ok := exprVal.(ir.Value); ok {
				if constInt, ok := irVal.(*ir.ConstantInt); ok {
					val = constInt.Value
				}
			}
		}
		
		// Create IR Constant
		constVal := g.ctx.Builder.ConstInt(types.I32, val)
		
		// Define Global
		fullName := enumName + "_" + memName // Mangled name
		global := g.ctx.Builder.CreateGlobalConstant(fullName, constVal)
		
		// Update Symbol in Scope (Pass 2 update)
		// We use the logical name (Enum.Member) to find the symbol created in Pass 1
		if sym, ok := g.currentScope.Resolve(enumName + "." + memName); ok {
			sym.IRValue = global
			sym.Kind = symbol.SymConst
		}
		
		val++
	}
	return nil
}

func (g *Generator) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	if ctx.GenericParams() != nil { return nil }

	name := ctx.IDENTIFIER().GetText()
	
	// Apply namespace mangling if applicable
	irName := name
	if g.currentNamespace != "" && name != "main" {
		irName = g.currentNamespace + "_" + name
	}
	
	// Method mangling (simplified check)
	if ctx.GetParent() != nil {
		if _, ok := ctx.GetParent().(*parser.ClassMemberContext); ok {
			// In a robust visitor, we'd pass parent context or track 'currentContainer'
			// This part would depend on how ClassMember contexts are structured in your parser
		}
	}

	sym, _ := g.currentScope.Resolve(name)

	var paramTypes []types.Type
	var paramNames []string

	if ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			if param.SELF() != nil {
				// Handle 'self'
				// For now, assume void* or specific struct ptr if we tracked it
				paramTypes = append(paramTypes, types.NewPointer(types.I8)) 
				paramNames = append(paramNames, "self")
				continue
			}
			pType := g.resolveType(param.Type_())
			paramTypes = append(paramTypes, pType)
			paramNames = append(paramNames, param.IDENTIFIER().GetText())
		}
	}

	var retType types.Type = types.Void
	if ctx.ReturnType() != nil {
		retType = g.resolveType(ctx.ReturnType().Type_())
	}

	// Create IR Function
	fn := g.ctx.Builder.CreateFunction(irName, retType, paramTypes, false)
	g.ctx.EnterFunction(fn)
	
	if sym != nil {
		sym.IRValue = fn
	}

	g.enterScope(ctx)
	defer g.exitScope()

	// Parameters
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
	sym, ok := g.currentScope.Resolve(name)
	if !ok { return nil }

	alloca := g.ctx.Builder.CreateAlloca(sym.Type, name+".addr")
	sym.IRValue = alloca

	if ctx.Expression() != nil {
		val := g.Visit(ctx.Expression())
		if irVal, ok := val.(ir.Value); ok {
			irVal = g.emitCast(irVal, sym.Type)
			g.ctx.Builder.CreateStore(irVal, alloca)
		}
	} else {
		g.ctx.Builder.CreateStore(g.getZeroValue(sym.Type), alloca)
	}
	return nil
}