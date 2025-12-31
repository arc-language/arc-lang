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

func (g *Generator) VisitExternDecl(ctx *parser.ExternDeclContext) interface{} {
	// Temporarily switch namespace if an identifier is provided
	oldNs := g.currentNamespace
	if ctx.IDENTIFIER() != nil {
		g.currentNamespace = ctx.IDENTIFIER().GetText()
	}

	for _, member := range ctx.AllExternMember() {
		if fnDecl := member.ExternFunctionDecl(); fnDecl != nil {
			// CAST: Interface -> Concrete Pointer
			g.visitExternFunctionDecl(fnDecl.(*parser.ExternFunctionDeclContext))
		}
	}

	g.currentNamespace = oldNs
	return nil
}

func (g *Generator) visitExternFunctionDecl(ctx *parser.ExternFunctionDeclContext) {
	name := ctx.IDENTIFIER().GetText()
	
	// Use explicit external name if provided (e.g. func print "printf")
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

	// Declare in LLVM
	fn := g.ctx.Builder.DeclareFunction(externalName, retType, paramTypes, variadic)
	
	// Register in Symbol Table so calls to 'name' use this function
	if sym, ok := g.currentScope.Resolve(name); ok {
		sym.IRValue = fn
	} else if g.currentNamespace != "" {
		// If using namespace, check for qualified symbol
		if sym, ok := g.currentScope.Resolve(g.currentNamespace + "." + name); ok {
			sym.IRValue = fn
		}
	}
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
			// Visit the expression to evaluate it
			exprVal := g.Visit(member.Expression())
			if irVal, ok := exprVal.(ir.Value); ok {
				if constInt, ok := irVal.(*ir.ConstantInt); ok {
					val = constInt.Value
				}
			}
		}
		
		// Create IR Constant
		constVal := g.ctx.Builder.ConstInt(types.I32, val)
		
		// Define Global Constant
		fullName := enumName + "_" + memName // Mangled name
		global := g.ctx.Builder.CreateGlobalConstant(fullName, constVal)
		
		// Update Symbol in Scope (Pass 2 update)
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
	irName := name
	
	// 1. Determine Context (Method vs Function)
	var parentName string
	isMethod := false
	
	// Check parent in parse tree to see if we are inside a Class or Struct
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

	// 2. Name Mangling
	if isMethod {
		irName = parentName + "_" + name
	} else if g.currentNamespace != "" && name != "main" {
		irName = g.currentNamespace + "_" + name
	}

	sym, _ := g.currentScope.Resolve(name)

	var paramTypes []types.Type
	var paramNames []string

	// 3. Parameter Parsing
	if ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			if param.SELF() != nil {
				// Handle 'self'
				if isMethod && parentName != "" {
					// Find the type of the parent struct/class
					if parentSym, ok := g.currentScope.Resolve(parentName); ok {
						// Pass by pointer
						structType := parentSym.Type
						paramTypes = append(paramTypes, types.NewPointer(structType))
					} else {
						// Fallback if semantic analysis failed
						paramTypes = append(paramTypes, types.NewPointer(types.I8))
					}
				} else {
					paramTypes = append(paramTypes, types.NewPointer(types.I8))
				}
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

	// 4. Create IR Function
	fn := g.ctx.Builder.CreateFunction(irName, retType, paramTypes, false)
	
	// Handle Async Attribute
	if ctx.ASYNC() != nil {
		fn.Attributes = append(fn.Attributes, ir.AttrCoroutine)
	}

	g.ctx.EnterFunction(fn)
	
	if sym != nil {
		sym.IRValue = fn
	}

	g.enterScope(ctx)
	defer g.exitScope()

	// 5. Argument Allocation
	for i, arg := range fn.Arguments {
		arg.SetName(paramNames[i])
		alloca := g.ctx.Builder.CreateAlloca(arg.Type(), paramNames[i]+".addr")
		g.ctx.Builder.CreateStore(arg, alloca)

		if s, ok := g.currentScope.Resolve(paramNames[i]); ok {
			s.IRValue = alloca
		}
	}

	// 6. Body Generation
	if ctx.Block() != nil {
		g.deferStack = NewDeferStack()
		for _, stmt := range ctx.Block().AllStatement() {
			g.Visit(stmt)
		}
	}

	// 7. Implicit Return
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