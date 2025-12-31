package irgen

import (
	"fmt"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

func (g *Generator) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	for _, ns := range ctx.AllNamespaceDecl() {
		g.Visit(ns)
	}
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
	if ctx.ExternDecl() != nil {
		return g.Visit(ctx.ExternDecl())
	}
	if ctx.StructDecl() != nil {
		return g.Visit(ctx.StructDecl())
	}
	if ctx.ClassDecl() != nil {
		return g.Visit(ctx.ClassDecl())
	}
	if ctx.EnumDecl() != nil {
		return g.Visit(ctx.EnumDecl())
	}
	if ctx.ConstDecl() != nil {
		return g.Visit(ctx.ConstDecl())
	}
	return nil
}

func (g *Generator) VisitNamespaceDecl(ctx *parser.NamespaceDeclContext) interface{} {
	if ctx.IDENTIFIER() != nil {
		g.currentNamespace = ctx.IDENTIFIER().GetText()
	}
	return nil
}

func (g *Generator) VisitExternDecl(ctx *parser.ExternDeclContext) interface{} {
	oldNs := g.currentNamespace
	if ctx.IDENTIFIER() != nil {
		g.currentNamespace = ctx.IDENTIFIER().GetText()
	}
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
		if len(raw) >= 2 {
			externalName = raw[1 : len(raw)-1]
		}
	}

	var retType types.Type = types.Void
	if ctx.Type_() != nil {
		retType = g.resolveType(ctx.Type_())
	}

	var paramTypes []types.Type
	variadic := false
	if ctx.ExternParameterList() != nil {
		if ctx.ExternParameterList().ELLIPSIS() != nil {
			variadic = true
		}
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

func (g *Generator) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	for _, member := range ctx.AllStructMember() {
		if member.FunctionDecl() != nil {
			g.Visit(member.FunctionDecl())
		}
	}
	return nil
}

func (g *Generator) VisitClassDecl(ctx *parser.ClassDeclContext) interface{} {
	for _, member := range ctx.AllClassMember() {
		if member.FunctionDecl() != nil {
			g.Visit(member.FunctionDecl())
		}
	}
	return nil
}

func (g *Generator) VisitEnumDecl(ctx *parser.EnumDeclContext) interface{} {
	enumName := ctx.IDENTIFIER().GetText()
	val := int64(0)
	for _, member := range ctx.AllEnumMember() {
		memName := member.IDENTIFIER().GetText()
		if member.Expression() != nil {
			exprVal := g.Visit(member.Expression())
			if irVal, ok := exprVal.(*ir.ConstantInt); ok {
				val = irVal.Value
			}
		}
		constVal := g.ctx.Builder.ConstInt(types.I32, val)
		
		fullName := enumName + "." + memName
		if sym, ok := g.currentScope.Resolve(fullName); ok {
			sym.IRValue = constVal
			sym.Kind = symbol.SymConst
		}
		val++
	}
	return nil
}

func (g *Generator) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	if ctx.GenericParams() != nil {
		// Generics not fully supported in IRGen yet
		return nil
	}

	name := ctx.IDENTIFIER().GetText()
	irName := name
	
	// Handle method naming (Struct_Method)
	var parentName string
	isMethod := false
	if parent := ctx.GetParent(); parent != nil {
		if _, ok := parent.(*parser.ClassMemberContext); ok {
			if classDecl, ok := parent.GetParent().(*parser.ClassDeclContext); ok {
				parentName = classDecl.IDENTIFIER().GetText()
				isMethod = true
			}
		} else if _, ok := parent.(*parser.StructMemberContext); ok {
			if structDecl, ok := parent.GetParent().(*parser.StructDeclContext); ok {
				parentName = structDecl.IDENTIFIER().GetText()
				isMethod = true
			}
		}
	}

	if isMethod {
		irName = parentName + "_" + name
	} else if g.currentNamespace != "" && name != "main" {
		irName = g.currentNamespace + "_" + name
	}

	sym, _ := g.currentScope.Resolve(name)
	
	var paramTypes []types.Type
	var paramNames []string
	
	if ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			if param.SELF() != nil {
				// Self handling: assume semantic pass resolved this to a pointer type
				if param.Type_() != nil {
					paramTypes = append(paramTypes, g.resolveType(param.Type_()))
				} else {
					paramTypes = append(paramTypes, types.NewPointer(types.I8)) // Fallback
				}
				paramNames = append(paramNames, "self")
				continue
			}
			paramTypes = append(paramTypes, g.resolveType(param.Type_()))
			paramNames = append(paramNames, param.IDENTIFIER().GetText())
		}
	}

	var retType types.Type = types.Void
	if ctx.ReturnType() != nil {
		if ctx.ReturnType().Type_() != nil {
			retType = g.resolveType(ctx.ReturnType().Type_())
		} else if ctx.ReturnType().TypeList() != nil {
			// Tuple Return: represented as anonymous struct
			var fields []types.Type
			for _, t := range ctx.ReturnType().TypeList().AllType_() {
				fields = append(fields, g.resolveType(t))
			}
			retType = types.NewStruct("", fields, false)
		}
	}

	fn := g.ctx.Builder.CreateFunction(irName, retType, paramTypes, false)
	
	if sym != nil {
		sym.IRValue = fn
	}

	g.ctx.EnterFunction(fn)
	g.enterScope(ctx) // Push function scope
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

	// Implicit return handling
	if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
		g.deferStack.Emit(g) // Run deferred actions
		
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
	// 1. Tuple Destructuring: let (a, b) = call()
	if ctx.TuplePattern() != nil {
		if ctx.Expression() == nil { return nil }
		
		// Evaluate RHS (returns struct/tuple value)
		val := g.Visit(ctx.Expression()).(ir.Value)
		names := ctx.TuplePattern().AllIDENTIFIER()
		
		for i, idNode := range names {
			name := idNode.GetText()
			sym, ok := g.currentScope.Resolve(name)
			if !ok { continue }
			
			// Extract field from tuple
			fieldVal := g.ctx.Builder.CreateExtractValue(val, []int{i}, "")
			
			// Create local variable
			alloca := g.ctx.Builder.CreateAlloca(sym.Type, name+".addr")
			g.ctx.Builder.CreateStore(fieldVal, alloca)
			sym.IRValue = alloca
		}
		return nil
	}

	// 2. Standard Declaration: let x = ...
	name := ctx.IDENTIFIER().GetText()
	sym, ok := g.currentScope.Resolve(name)
	if !ok {
		fmt.Printf("[IRGen] Error: Variable '%s' was not defined in the symbol table\n", name)
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
			g.ctx.Builder.CreateStore(g.getZeroValue(sym.Type), alloca)
		}
	} else {
		g.ctx.Builder.CreateStore(g.getZeroValue(sym.Type), alloca)
	}
	return nil
}

func (g *Generator) VisitConstDecl(ctx *parser.ConstDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	sym, ok := g.currentScope.Resolve(name)
	if !ok { return nil }

	val := g.Visit(ctx.Expression()).(ir.Value)
	
	// Global Constants
	if g.currentScope.Parent == nil {
		if constant, ok := val.(ir.Constant); ok {
			global := g.ctx.Builder.CreateGlobalConstant(name, constant)
			sym.IRValue = global
		}
	} else {
		// Local constants behave like values
		sym.IRValue = val
	}
	
	return nil
}