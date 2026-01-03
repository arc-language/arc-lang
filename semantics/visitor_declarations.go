package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

func (a *Analyzer) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	// Pass 1: Register Types
	for _, decl := range ctx.AllTopLevelDecl() {
		if decl.StructDecl() != nil {
			name := decl.StructDecl().IDENTIFIER().GetText()
			st := types.NewStruct(name, nil, false)
			a.currentScope.Define(name, symbol.SymType, st)
		} else if decl.ClassDecl() != nil {
			name := decl.ClassDecl().IDENTIFIER().GetText()
			st := types.NewStruct(name, nil, false)
			a.currentScope.Define(name, symbol.SymType, st)
		} else if decl.EnumDecl() != nil {
			name := decl.EnumDecl().IDENTIFIER().GetText()
			a.currentScope.Define(name, symbol.SymType, types.I32)
		}
	}

	// Pass 2: Declarations
	for _, decl := range ctx.AllTopLevelDecl() { a.Visit(decl) }
	for _, ns := range ctx.AllNamespaceDecl() { a.Visit(ns) }
	return nil
}

func (a *Analyzer) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext) interface{} {
	if ctx.FunctionDecl() != nil { return a.Visit(ctx.FunctionDecl()) }
	if ctx.VariableDecl() != nil { return a.Visit(ctx.VariableDecl()) }
	if ctx.StructDecl() != nil { return a.Visit(ctx.StructDecl()) }
	if ctx.ClassDecl() != nil { return a.Visit(ctx.ClassDecl()) }
	if ctx.ExternDecl() != nil { return a.Visit(ctx.ExternDecl()) }
	if ctx.ConstDecl() != nil { return a.Visit(ctx.ConstDecl()) }
	if ctx.MutatingDecl() != nil { return a.Visit(ctx.MutatingDecl()) }
	return nil
}

func (a *Analyzer) VisitExternDecl(ctx *parser.ExternDeclContext) interface{} {
	ns := ""
	if ctx.IDENTIFIER() != nil { ns = ctx.IDENTIFIER().GetText() }
	for _, member := range ctx.AllExternMember() {
		if fnCtx := member.ExternFunctionDecl(); fnCtx != nil {
			name := fnCtx.IDENTIFIER().GetText()
			
			// Resolve Return Type
			var retType types.Type = types.Void
			if fnCtx.Type_() != nil { retType = a.resolveType(fnCtx.Type_()) }
			
			// Resolve Parameter Types
			var paramTypes []types.Type
			variadic := false
			if fnCtx.ExternParameterList() != nil {
				if fnCtx.ExternParameterList().ELLIPSIS() != nil {
					variadic = true
				}
				for _, t := range fnCtx.ExternParameterList().AllType_() {
					paramTypes = append(paramTypes, a.resolveType(t))
				}
			}
			
			// Create Function Type
			fnType := types.NewFunction(retType, paramTypes, variadic)
			
			lookupName := name
			if ns != "" { lookupName = ns + "." + name }
			a.currentScope.Define(lookupName, symbol.SymFunc, fnType)
		}
	}
	return nil
}

func (a *Analyzer) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	// Resolve Return Type
	var retType types.Type = types.Void
	if ctx.ReturnType() != nil && ctx.ReturnType().Type_() != nil {
		retType = a.resolveType(ctx.ReturnType().Type_())
	} else if ctx.ReturnType() != nil && ctx.ReturnType().TypeList() != nil {
		// Tuple return type
		var tupleTypes []types.Type
		for _, t := range ctx.ReturnType().TypeList().AllType_() {
			tupleTypes = append(tupleTypes, a.resolveType(t))
		}
		// Represent tuple as anonymous struct
		retType = types.NewStruct("", tupleTypes, false)
	}

	// Resolve Parameter Types & Handle 'self'
	var paramTypes []types.Type
	var selfType types.Type
	
	if ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			pType := a.resolveType(param.Type_())
			if param.SELF() != nil {
				// Mark that we found self
				selfType = pType
			}
			paramTypes = append(paramTypes, pType)
		}
	}
	
	// Mangle name if it's a method (has self)
	if selfType != nil {
		// Extract struct name from self type
		// Self can be T or *T
		base := selfType
		if ptr, ok := base.(*types.PointerType); ok {
			base = ptr.ElementType
		}
		
		if st, ok := base.(*types.StructType); ok {
			name = st.Name + "_" + name
		}
	}
	
	// Create Function Type
	fnType := types.NewFunction(retType, paramTypes, false)

	a.currentScope.Define(name, symbol.SymFunc, fnType)
	
	// Scope Analysis
	a.currentFuncRetType = retType
	a.pushScope(ctx)
	defer func() { a.popScope(); a.currentFuncRetType = nil }()

	if ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			pName := param.IDENTIFIER().GetText()
			pType := a.resolveType(param.Type_())
			a.currentScope.Define(pName, symbol.SymVar, pType)
		}
	}
	if ctx.Block() != nil {
		a.scopes[ctx.Block()] = a.currentScope
		for _, stmt := range ctx.Block().AllStatement() { a.Visit(stmt) }
	}
	return nil
}

func (a *Analyzer) VisitMutatingDecl(ctx *parser.MutatingDeclContext) interface{} {
	// Syntax: mutating Name(self param: Type, p2: Type...) { ... }
	name := ctx.IDENTIFIER(0).GetText()
	
	// Resolve Return Type
	var retType types.Type = types.Void
	if ctx.ReturnType() != nil {
		if ctx.ReturnType().Type_() != nil {
			retType = a.resolveType(ctx.ReturnType().Type_())
		}
	}

	var paramTypes []types.Type
	
	// 1. Handle Self Param (Explicit in grammar)
	// Rule structure: IDENTIFIER(0)=name, IDENTIFIER(1)=selfName, Type_()=selfType
	selfParamName := ctx.IDENTIFIER(1).GetText()
	selfType := a.resolveType(ctx.Type_()) // The first type listed is self's type
	paramTypes = append(paramTypes, selfType)
	
	// 2. Handle Other Params
	for _, param := range ctx.AllParameter() {
		pType := a.resolveType(param.Type_())
		paramTypes = append(paramTypes, pType)
	}

	// 3. Name Mangling (Struct_Method)
	// Infer struct name from self type (expected to be *Struct)
	var structName string
	base := selfType
	if ptr, ok := base.(*types.PointerType); ok {
		base = ptr.ElementType
	}
	if st, ok := base.(*types.StructType); ok {
		structName = st.Name
	}
	
	fullName := name
	if structName != "" {
		fullName = structName + "_" + name
	}

	// 4. Define Function Symbol
	fnType := types.NewFunction(retType, paramTypes, false)
	a.currentScope.Define(fullName, symbol.SymFunc, fnType)

	// 5. Scope & Body Analysis
	a.currentFuncRetType = retType
	a.pushScope(ctx)
	defer func() { a.popScope(); a.currentFuncRetType = nil }()

	// Define Self
	a.currentScope.Define(selfParamName, symbol.SymVar, selfType)
	
	// Define Params
	for _, param := range ctx.AllParameter() {
		pName := param.IDENTIFIER().GetText()
		pType := a.resolveType(param.Type_())
		a.currentScope.Define(pName, symbol.SymVar, pType)
	}

	if ctx.Block() != nil {
		a.scopes[ctx.Block()] = a.currentScope
		for _, stmt := range ctx.Block().AllStatement() { a.Visit(stmt) }
	}

	return nil
}

func (a *Analyzer) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	// Handle Tuple Destructuring: let (a, b) = pair
	if ctx.TuplePattern() != nil {
		var rhsType types.Type
		if ctx.Expression() != nil {
			rhsType = a.Visit(ctx.Expression()).(types.Type)
		}
		
		names := ctx.TuplePattern().AllIDENTIFIER()
		st, isStruct := rhsType.(*types.StructType)
		
		for i, idNode := range names {
			name := idNode.GetText()
			var fieldType types.Type = types.I64 // Fallback
			
			if isStruct && i < len(st.Fields) {
				fieldType = st.Fields[i]
			}
			a.currentScope.Define(name, symbol.SymVar, fieldType)
		}
		return nil
	}

	// Standard Declaration: let x = ...
	name := ctx.IDENTIFIER().GetText()
	if _, ok := a.currentScope.ResolveLocal(name); ok {
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Redeclaration of '%s'", name)
		return nil
	}

	var typ types.Type
	if ctx.Type_() != nil { typ = a.resolveType(ctx.Type_()) }
	
	if ctx.Expression() != nil {
		exprType := a.Visit(ctx.Expression()).(types.Type)
		if typ == nil {
			typ = exprType
		} else {
			// Allow Void expression to initialize Array/Struct (zero-init assumption)
			if exprType != types.Void && !areTypesCompatible(exprType, typ) {
				a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, 
					"Type mismatch in variable '%s': expected %s, got %s", 
					name, typ.String(), exprType.String())
			}
		}
	}
	if typ == nil { typ = types.I64 }
	a.currentScope.Define(name, symbol.SymVar, typ)
	return nil
}

func (a *Analyzer) VisitConstDecl(ctx *parser.ConstDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	var typ types.Type
	if ctx.Type_() != nil { typ = a.resolveType(ctx.Type_()) }
	if ctx.Expression() != nil {
		exprType := a.Visit(ctx.Expression()).(types.Type)
		if typ == nil { typ = exprType }
	}
	if typ == nil { typ = types.I64 }
	a.currentScope.Define(name, symbol.SymConst, typ)
	return nil
}

func (a *Analyzer) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	sym, _ := a.currentScope.Resolve(name)
	if sym == nil { return nil }
	st := sym.Type.(*types.StructType)
	
	var fields []types.Type
	indices := make(map[string]int)
	fieldCount := 0

	// Pass 1: Collect Fields
	for _, member := range ctx.AllStructMember() {
		if f := member.StructField(); f != nil {
			fields = append(fields, a.resolveType(f.Type_()))
			indices[f.IDENTIFIER().GetText()] = fieldCount
			fieldCount++
		}
	}
	
	// Update Struct definition immediately so methods can resolve fields
	st.Fields = fields
	a.structIndices[name] = indices
	
	// Pass 2: Analyze Methods
	for _, member := range ctx.AllStructMember() {
		if m := member.FunctionDecl(); m != nil { a.Visit(m) }
		if m := member.MutatingDecl(); m != nil { a.Visit(m) }
	}
	return nil
}

func (a *Analyzer) VisitClassDecl(ctx *parser.ClassDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	sym, _ := a.currentScope.Resolve(name)
	if sym == nil { return nil }
	st := sym.Type.(*types.StructType)
	
	var fields []types.Type
	indices := make(map[string]int)
	fieldCount := 0

	// Pass 1: Collect Fields
	for _, member := range ctx.AllClassMember() {
		if f := member.ClassField(); f != nil {
			fields = append(fields, a.resolveType(f.Type_()))
			indices[f.IDENTIFIER().GetText()] = fieldCount
			fieldCount++
		}
	}
	
	// Update Class/Struct definition immediately
	st.Fields = fields
	a.structIndices[name] = indices

	// Pass 2: Analyze Methods
	for _, member := range ctx.AllClassMember() {
		if m := member.FunctionDecl(); m != nil { a.Visit(m) }
		if d := member.DeinitDecl(); d != nil { a.Visit(d) }
	}
	return nil
}

func (a *Analyzer) VisitEnumDecl(ctx *parser.EnumDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	for _, m := range ctx.AllEnumMember() {
		// Define: Enum.Member
		a.currentScope.Define(name+"."+m.IDENTIFIER().GetText(), symbol.SymConst, types.I32)
	}
	return nil
}

// Stubs
func (a *Analyzer) VisitMethodDecl(ctx *parser.MethodDeclContext) interface{} { return nil }
func (a *Analyzer) VisitDeinitDecl(ctx *parser.DeinitDeclContext) interface{} { return nil }