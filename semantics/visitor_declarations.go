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
	if ctx.ConstDecl() != nil { return a.Visit(ctx.ConstDecl()) } // Added
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

	// Resolve Parameter Types
	var paramTypes []types.Type
	if ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			if param.SELF() != nil { 
				// Handle Self (placeholder for now, pointer to struct)
				paramTypes = append(paramTypes, types.NewPointer(types.Void)) 
				continue 
			}
			paramTypes = append(paramTypes, a.resolveType(param.Type_()))
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
			if param.SELF() != nil { continue }
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
	for i, member := range ctx.AllStructMember() {
		if f := member.StructField(); f != nil {
			fields = append(fields, a.resolveType(f.Type_()))
			indices[f.IDENTIFIER().GetText()] = i
		}
		if m := member.FunctionDecl(); m != nil { a.Visit(m) }
	}
	st.Fields = fields
	a.structIndices[name] = indices
	return nil
}

func (a *Analyzer) VisitClassDecl(ctx *parser.ClassDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	sym, _ := a.currentScope.Resolve(name)
	if sym == nil { return nil }
	st := sym.Type.(*types.StructType)
	var fields []types.Type
	indices := make(map[string]int)
	for i, member := range ctx.AllClassMember() {
		if f := member.ClassField(); f != nil {
			fields = append(fields, a.resolveType(f.Type_()))
			indices[f.IDENTIFIER().GetText()] = i
		}
		if m := member.FunctionDecl(); m != nil { a.Visit(m) }
	}
	st.Fields = fields
	a.structIndices[name] = indices
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
func (a *Analyzer) VisitMutatingDecl(ctx *parser.MutatingDeclContext) interface{} { return nil }
func (a *Analyzer) VisitDeinitDecl(ctx *parser.DeinitDeclContext) interface{} { return nil }