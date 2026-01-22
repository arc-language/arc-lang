package irgen

import (

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/irgen/mangler"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

func (g *Generator) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	g.currentNamespace = ""
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
	if ctx.MutatingDecl() != nil {
		return g.Visit(ctx.MutatingDecl())
	}

	if g.Phase == 1 {
		if ctx.VariableDecl() != nil {
			return g.Visit(ctx.VariableDecl())
		}
		if ctx.ConstDecl() != nil {
			return g.Visit(ctx.ConstDecl())
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
		if ctx.ExternCDecl() != nil {
			return g.Visit(ctx.ExternCDecl())
		}
		if ctx.ExternCppDecl() != nil {
			return g.Visit(ctx.ExternCppDecl())
		}
	}

	if g.Phase == 2 {
		if ctx.VariableDecl() != nil {
			return g.Visit(ctx.VariableDecl())
		}
		if ctx.StructDecl() != nil {
			return g.Visit(ctx.StructDecl())
		}
		if ctx.ClassDecl() != nil {
			return g.Visit(ctx.ClassDecl())
		}
	}
	return nil
}

func (g *Generator) VisitNamespaceDecl(ctx *parser.NamespaceDeclContext) interface{} {
	if ctx.IDENTIFIER() != nil {
		g.currentNamespace = ctx.IDENTIFIER().GetText()
	}
	return nil
}

func (g *Generator) VisitExternCDecl(ctx *parser.ExternCDeclContext) interface{} {
	if g.Phase != 1 {
		return nil
	}
	oldNs := g.currentNamespace
	for _, member := range ctx.AllExternCMember() {
		if fnDecl := member.ExternCFunctionDecl(); fnDecl != nil {
			g.visitExternCFunctionDecl(fnDecl.(*parser.ExternCFunctionDeclContext))
		}
	}
	g.currentNamespace = oldNs
	return nil
}

func (g *Generator) visitExternCFunctionDecl(ctx *parser.ExternCFunctionDeclContext) {
	name := ctx.IDENTIFIER().GetText()

	// Build lookup name with namespace
	lookupName := name
	if g.currentNamespace != "" {
		lookupName = g.currentNamespace + "." + name
	}

	// Get the external name (for linking)
	externalName := name
	if ctx.STRING_LITERAL() != nil {
		raw := ctx.STRING_LITERAL().GetText()
		if len(raw) >= 2 {
			externalName = raw[1 : len(raw)-1]
		}
	}

	// Resolve the symbol
	sym, ok := g.currentScope.Resolve(lookupName)
	if !ok {
		return
	}

	fnType := sym.Type.(*types.FunctionType)
	fn := g.ctx.Builder.DeclareFunction(externalName, fnType.ReturnType, fnType.ParamTypes, fnType.Variadic)

	if ctx.CCallingConvention() != nil {
		ccCtx := ctx.CCallingConvention()
		if ccCtx.STDCALL() != nil {
			g.ctx.Builder.SetCallConv(fn, ir.CC_StdCall)
		} else if ccCtx.FASTCALL() != nil {
			g.ctx.Builder.SetCallConv(fn, ir.CC_FastCall)
		}
	}

	if sym != nil {
		sym.IRValue = fn
	}
}

func (g *Generator) VisitExternCppDecl(ctx *parser.ExternCppDeclContext) interface{} {
	if g.Phase != 1 {
		return nil
	}
	for _, member := range ctx.AllExternCppMember() {
		g.visitExternCppMember(member)
	}
	return nil
}

func (g *Generator) visitExternCppMember(ctx parser.IExternCppMemberContext) {
	c := ctx.(*parser.ExternCppMemberContext)
	if fn := c.ExternCppFunctionDecl(); fn != nil {
		g.visitExternCppFunctionDecl(fn.(*parser.ExternCppFunctionDeclContext))
	} else if ns := c.ExternCppNamespaceDecl(); ns != nil {
		g.visitExternCppNamespace(ns.(*parser.ExternCppNamespaceDeclContext))
	} else if cl := c.ExternCppClassDecl(); cl != nil {
		g.visitExternCppClass(cl.(*parser.ExternCppClassDeclContext))
	}
}

func (g *Generator) visitExternCppNamespace(ctx *parser.ExternCppNamespaceDeclContext) {
	pathCtx := ctx.ExternNamespacePath()
	nsName := ""
	if pathCtx != nil {
		ids := pathCtx.AllIDENTIFIER()
		for i, id := range ids {
			if i > 0 {
				nsName += "."
			}
			nsName += id.GetText()
		}
	}
	prev := g.currentNamespace
	if g.currentNamespace == "" {
		g.currentNamespace = nsName
	} else {
		g.currentNamespace = g.currentNamespace + "." + nsName
	}
	for _, member := range ctx.AllExternCppMember() {
		g.visitExternCppMember(member)
	}
	g.currentNamespace = prev
}

func (g *Generator) visitExternCppClass(ctx *parser.ExternCppClassDeclContext) {
	className := ctx.IDENTIFIER().GetText()
	prev := g.currentNamespace
	if g.currentNamespace == "" {
		g.currentNamespace = className
	} else {
		g.currentNamespace = g.currentNamespace + "." + className
	}
	for _, member := range ctx.AllExternCppClassMember() {
		if method := member.ExternCppMethodDecl(); method != nil {
			g.visitExternCppMethod(method.(*parser.ExternCppMethodDeclContext))
		}
	}
	g.currentNamespace = prev
}

func (g *Generator) visitExternCppFunctionDecl(ctx *parser.ExternCppFunctionDeclContext) {
	name := ctx.IDENTIFIER().GetText()

	// Build lookup name with Arc namespace
	lookupName := name
	if g.currentNamespace != "" {
		lookupName = g.currentNamespace + "." + name
	}

	// Resolve symbol
	sym, ok := g.currentScope.Resolve(lookupName)
	if !ok {
		return
	}

	// Get external name (mangled or explicit)
	externalName := ""
	if ctx.STRING_LITERAL() != nil {
		raw := ctx.STRING_LITERAL().GetText()
		if len(raw) >= 2 {
			externalName = raw[1 : len(raw)-1]
		}
	} else {
		// Use Itanium mangling for C++ functions
		externalName = mangler.MangleItanium(sym, false)
	}

	fnType := sym.Type.(*types.FunctionType)
	fn := g.ctx.Builder.DeclareFunction(externalName, fnType.ReturnType, fnType.ParamTypes, fnType.Variadic)

	if ctx.CppCallingConvention() != nil {
		g.applyCppCallConv(fn, ctx.CppCallingConvention())
	}

	sym.IRValue = fn
}

func (g *Generator) visitExternCppMethod(ctx *parser.ExternCppMethodDeclContext) {
	name := ctx.IDENTIFIER().GetText()
	lookupName := g.currentNamespace + "_" + name
	sym, ok := g.currentScope.Resolve(lookupName)
	if !ok {
		return
	}
	if sym.IsVirtual {
		return
	}
	externalName := ""
	if ctx.STRING_LITERAL() != nil {
		raw := ctx.STRING_LITERAL().GetText()
		if len(raw) >= 2 {
			externalName = raw[1 : len(raw)-1]
		}
	} else {
		externalName = mangler.MangleItanium(sym, true)
	}
	fnType := sym.Type.(*types.FunctionType)
	fn := g.ctx.Builder.DeclareFunction(externalName, fnType.ReturnType, fnType.ParamTypes, fnType.Variadic)
	if ctx.CppCallingConvention() != nil {
		g.applyCppCallConv(fn, ctx.CppCallingConvention())
	}
	sym.IRValue = fn
}

func (g *Generator) applyCppCallConv(fn *ir.Function, ctx parser.ICppCallingConventionContext) {
	cc := ctx.(*parser.CppCallingConventionContext)
	if cc.STDCALL() != nil {
		g.ctx.Builder.SetCallConv(fn, ir.CC_StdCall)
	} else if cc.THISCALL() != nil {
		g.ctx.Builder.SetCallConv(fn, ir.CC_ThisCall)
	} else if cc.VECTORCALL() != nil {
		g.ctx.Builder.SetCallConv(fn, ir.CC_VectorCall)
	} else if cc.FASTCALL() != nil {
		g.ctx.Builder.SetCallConv(fn, ir.CC_FastCall)
	}
}

func (g *Generator) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	if g.Phase == 1 {
		name := ctx.IDENTIFIER().GetText()
		if sym, ok := g.currentScope.Resolve(name); ok {
			if st, ok := sym.Type.(*types.StructType); ok {
				g.ctx.Builder.DefineStruct(st)
			}
		}
	}
	for _, member := range ctx.AllStructMember() {
		if member.FunctionDecl() != nil {
			g.Visit(member.FunctionDecl())
		}
		if member.MutatingDecl() != nil {
			g.Visit(member.MutatingDecl())
		}
	}
	return nil
}

func (g *Generator) VisitClassDecl(ctx *parser.ClassDeclContext) interface{} {
	if g.Phase == 1 {
		name := ctx.IDENTIFIER().GetText()
		if sym, ok := g.currentScope.Resolve(name); ok {
			if st, ok := sym.Type.(*types.StructType); ok {
				// If it is a class, we do NOT use st.Fields directly for the IR definition.
				// We must prepend the Reference Count header (i64).
				if st.IsClass {
					// 1. Create a new slice for IR definition: [RefCount, ...Fields]
					irFields := make([]types.Type, len(st.Fields)+1)
					irFields[0] = types.I64 // Header
					copy(irFields[1:], st.Fields)
					
					// 2. Create a temporary type description for the backend definition
					// We keep the name so it registers as %ClassName
					defSt := types.NewStruct(st.Name, irFields, st.Packed)
					g.ctx.Builder.DefineStruct(defSt)
				} else {
					g.ctx.Builder.DefineStruct(st)
				}
			}
		}
	}
	
	for _, member := range ctx.AllClassMember() {
		if member.FunctionDecl() != nil {
			g.Visit(member.FunctionDecl())
		}
	}
	return nil
}

func (g *Generator) VisitEnumDecl(ctx *parser.EnumDeclContext) interface{} {
	if g.Phase == 1 {
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
	}
	return nil
}

func (g *Generator) VisitMutatingDecl(ctx *parser.MutatingDeclContext) interface{} {
	name := ctx.IDENTIFIER(0).GetText()

	var parentName string
	if parent := ctx.GetParent(); parent != nil {
		if _, ok := parent.(*parser.StructMemberContext); ok {
			if sd, ok := parent.GetParent().(*parser.StructDeclContext); ok {
				parentName = sd.IDENTIFIER().GetText()
			}
		}
	}

	// Fix: Handle Flat Mutating Methods (mutating func foo(self x: T))
	if parentName == "" && ctx.Type_() != nil {
		t := g.resolveType(ctx.Type_())
		if ptr, ok := t.(*types.PointerType); ok {
			t = ptr.ElementType
		}
		if st, ok := t.(*types.StructType); ok {
			parentName = st.Name
		}
	}

	irName := name
	if parentName != "" {
		irName = parentName + "_" + name
	}

	sym, _ := g.currentScope.Resolve(irName)

	if g.Phase == 1 {
		var retType types.Type = types.Void
		if ctx.ReturnType() != nil {
			if ctx.ReturnType().Type_() != nil {
				retType = g.resolveType(ctx.ReturnType().Type_())
			} else if ctx.ReturnType().TypeList() != nil {
				// Handle tuple return types
				var tupleTypes []types.Type
				for _, t := range ctx.ReturnType().TypeList().AllType_() {
					tupleTypes = append(tupleTypes, g.resolveType(t))
				}
				retType = types.NewStruct("", tupleTypes, false)
			}
		}
		var paramTypes []types.Type
		if ctx.Type_() != nil {
			paramTypes = append(paramTypes, g.resolveType(ctx.Type_()))
		}
		for _, param := range ctx.AllParameter() {
			paramTypes = append(paramTypes, g.resolveType(param.Type_()))
		}
		fn := g.ctx.Builder.CreateFunction(irName, retType, paramTypes, false)
		if sym != nil {
			sym.IRValue = fn
		}
		return nil
	}

	if g.Phase == 2 {
		if sym == nil || sym.IRValue == nil {
			return nil
		}
		fn := sym.IRValue.(*ir.Function)
		g.ctx.EnterFunction(fn)
		g.enterScope(ctx)
		defer g.exitScope()

		argIdx := 0
		selfName := ctx.IDENTIFIER(1).GetText()
		if argIdx < len(fn.Arguments) {
			arg := fn.Arguments[argIdx]
			arg.SetName(selfName)
			alloca := g.ctx.Builder.CreateAlloca(arg.Type(), selfName+".addr")
			g.ctx.Builder.CreateStore(arg, alloca)
			if s, ok := g.currentScope.Resolve(selfName); ok {
				s.IRValue = alloca
			}
			argIdx++
		}

		for _, param := range ctx.AllParameter() {
			if argIdx < len(fn.Arguments) {
				pName := param.IDENTIFIER().GetText()
				arg := fn.Arguments[argIdx]
				arg.SetName(pName)
				alloca := g.ctx.Builder.CreateAlloca(arg.Type(), pName+".addr")
				g.ctx.Builder.CreateStore(arg, alloca)
				if s, ok := g.currentScope.Resolve(pName); ok {
					s.IRValue = alloca
				}
				argIdx++
			}
		}

		if ctx.Block() != nil {
			g.deferStack = NewDeferStack()
			g.Visit(ctx.Block())
		}

		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			g.deferStack.Emit(g)
			if fn.FuncType.ReturnType == types.Void {
				g.ctx.Builder.CreateRetVoid()
			} else {
				g.ctx.Builder.CreateRet(g.getZeroValue(fn.FuncType.ReturnType))
			}
		}
		g.ctx.ExitFunction()
	}
	return nil
}

func (g *Generator) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	// --- Phase 1: Global Variable Declarations ---
	if g.ctx.CurrentFunction == nil && g.Phase == 1 {
		name := ctx.IDENTIFIER().GetText()
		lookupName := name
		if g.currentNamespace != "" {
			lookupName = g.currentNamespace + "." + name
		}

		sym, ok := g.currentScope.Resolve(lookupName)
		if !ok {
			return nil
		}
		var init ir.Constant = g.ctx.Builder.ConstZero(sym.Type)
		glob := g.ctx.Builder.CreateGlobalVariable(name, sym.Type, init)
		sym.IRValue = glob
		return nil
	}

	// --- Phase 2: Local Variable Declarations ---
	if g.ctx.CurrentFunction != nil && g.Phase == 2 {
		// 1. Handle Tuple Destructuring
		if ctx.TuplePattern() != nil {
			if ctx.Expression() == nil {
				return nil
			}
			val := g.Visit(ctx.Expression()).(ir.Value)
			names := ctx.TuplePattern().AllIDENTIFIER()
			for i, idNode := range names {
				name := idNode.GetText()
				sym, ok := g.currentScope.Resolve(name)
				if !ok {
					continue
				}
				fieldVal := g.ctx.Builder.CreateExtractValue(val, []int{i}, "")
				alloca := g.ctx.Builder.CreateAlloca(sym.Type, name+".addr")
				g.ctx.Builder.CreateStore(fieldVal, alloca)
				sym.IRValue = alloca
			}
			return nil
		}

		// 2. Handle Standard Declaration: let x = val
		name := ctx.IDENTIFIER().GetText()
		sym, ok := g.currentScope.Resolve(name)
		if !ok {
			return nil
		}

		var initVal ir.Value
		if ctx.Expression() != nil {
			initVal = g.Visit(ctx.Expression()).(ir.Value)
			
			// Type Inference
			if ctx.Type_() == nil && initVal != nil {
				sym.Type = initVal.Type()
			}
		}

		if sym.Type == nil || sym.Type == types.Void {
			sym.Type = types.I64
		}

		// Determine Storage Type
		// Struct: %MyStruct
		// Class:  %MyClass* (Pointer to heap object)
		var storageType types.Type = sym.Type
		if st, ok := sym.Type.(*types.StructType); ok && st.IsClass {
			// Variable holds a pointer to the object
			storageType = types.NewPointer(sym.Type)
		}

		// Create stack allocation
		alloca := g.ctx.Builder.CreateAlloca(storageType, name+".addr")
		sym.IRValue = alloca

		// Initialize value
		if initVal != nil {
			if initVal.Type() == types.Void {
				initVal = g.getZeroValue(sym.Type)
			}
			
			// If it's a class, initVal comes from VisitStructLiteral which returns %MyClass*.
			// storageType is %MyClass*. This matches.
			initVal = g.emitCast(initVal, storageType)
			g.ctx.Builder.CreateStore(initVal, alloca)
		} else {
			// Default zero initialization (Null for classes)
			g.ctx.Builder.CreateStore(g.getZeroValue(storageType), alloca)
		}
	}
	return nil
}

func (g *Generator) VisitConstDecl(ctx *parser.ConstDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	lookupName := name
	if g.currentNamespace != "" && name != "main" {
		lookupName = g.currentNamespace + "." + name
	}

	sym, ok := g.currentScope.Resolve(lookupName)
	if !ok {
		return nil
	}

	if g.currentScope.Parent == nil && g.Phase == 1 {
		val := g.Visit(ctx.Expression())
		if val != nil {
			if constant, ok := val.(ir.Constant); ok {
				sym.IRValue = constant
			}
		}
	} else if g.currentScope.Parent != nil {
		val := g.Visit(ctx.Expression())
		if constant, ok := val.(ir.Constant); ok {
			sym.IRValue = constant
		}
	}
	return nil
}