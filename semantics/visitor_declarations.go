package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

func (a *Analyzer) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	savedPrefix := a.currentNamespacePrefix
	for _, ns := range ctx.AllNamespaceDecl() {
		if ns.IDENTIFIER() != nil {
			name := ns.IDENTIFIER().GetText()
			if a.currentNamespacePrefix == "" {
				a.currentNamespacePrefix = name
			} else {
				a.currentNamespacePrefix = a.currentNamespacePrefix + "." + name
			}
		}
	}
	for _, decl := range ctx.AllTopLevelDecl() {
		a.Visit(decl)
	}
	a.currentNamespacePrefix = savedPrefix
	return nil
}

func (a *Analyzer) VisitNamespaceDecl(ctx *parser.NamespaceDeclContext) interface{} {
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		if a.currentNamespacePrefix == "" {
			a.currentNamespacePrefix = name
		} else {
			a.currentNamespacePrefix = a.currentNamespacePrefix + "." + name
		}
	}
	return nil
}

func (a *Analyzer) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext) interface{} {
	if a.Phase == 0 {
		if ctx.StructDecl() != nil { return a.Visit(ctx.StructDecl()) }
		if ctx.ClassDecl() != nil { return a.Visit(ctx.ClassDecl()) }
		if ctx.EnumDecl() != nil { return a.Visit(ctx.EnumDecl()) }
		if ctx.ExternCppDecl() != nil { return a.Visit(ctx.ExternCppDecl()) }
		return nil
	}
	if a.Phase == 1 {
		if ctx.FunctionDecl() != nil { return a.Visit(ctx.FunctionDecl()) }
		if ctx.StructDecl() != nil { return a.Visit(ctx.StructDecl()) }
		if ctx.ClassDecl() != nil { return a.Visit(ctx.ClassDecl()) }
		if ctx.EnumDecl() != nil { return a.Visit(ctx.EnumDecl()) }
		if ctx.ExternCDecl() != nil { return a.Visit(ctx.ExternCDecl()) }
		if ctx.ExternCppDecl() != nil { return a.Visit(ctx.ExternCppDecl()) }
		if ctx.ConstDecl() != nil { return a.Visit(ctx.ConstDecl()) }
		if ctx.VariableDecl() != nil { return a.Visit(ctx.VariableDecl()) }
		if ctx.MutatingDecl() != nil { return a.Visit(ctx.MutatingDecl()) }
		return nil
	}
	if a.Phase == 2 {
		if ctx.FunctionDecl() != nil { return a.Visit(ctx.FunctionDecl()) }
		if ctx.MutatingDecl() != nil { return a.Visit(ctx.MutatingDecl()) }
		if ctx.StructDecl() != nil { return a.Visit(ctx.StructDecl()) }
		if ctx.ClassDecl() != nil { return a.Visit(ctx.ClassDecl()) }
		if ctx.VariableDecl() != nil { return a.Visit(ctx.VariableDecl()) }
		if ctx.ConstDecl() != nil { return a.Visit(ctx.ConstDecl()) }
		return nil
	}
	return nil
}

// --- Struct Declaration (FIXED) ---

func (a *Analyzer) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	if a.currentNamespacePrefix != "" {
		name = a.currentNamespacePrefix + "." + name
	}

	// 1. Definition / Registration
	// We do this check first regardless of Phase. If the symbol doesn't exist,
	// we create it. This ensures attributes are captured even if Phase 0 is skipped.
	if _, ok := a.currentScope.ResolveLocal(name); !ok {
		// Check for @packed attribute
		isPacked := false
		for _, attr := range ctx.AllAttribute() {
			if attr.IDENTIFIER().GetText() == "packed" {
				isPacked = true
			}
		}

		// Define StructType
		st := types.NewStruct(name, nil, isPacked)
		a.currentScope.Define(name, symbol.SymType, st)
	}

	// 2. Field Resolution (Phase 1)
	if a.Phase == 1 {
		sym, _ := a.currentScope.Resolve(name)
		if sym != nil {
			st := sym.Type.(*types.StructType)
			// Only parse fields if we haven't already
			if len(st.Fields) == 0 {
				var fields []types.Type
				indices := make(map[string]int)
				fieldCount := 0

				for _, member := range ctx.AllStructMember() {
					if f := member.StructField(); f != nil {
						fields = append(fields, a.resolveType(f.Type_()))
						indices[f.IDENTIFIER().GetText()] = fieldCount
						fieldCount++
					}
				}
				st.Fields = fields
				a.structIndices[name] = indices
			}
		}
	}

	// 3. Method Analysis (Phase 1 & 2)
	for _, member := range ctx.AllStructMember() {
		if m := member.FunctionDecl(); m != nil { a.Visit(m) }
		if m := member.MutatingDecl(); m != nil { a.Visit(m) }
	}
	return nil
}

// --- Other Declarations (Unchanged but included for context) ---

func (a *Analyzer) VisitExternCDecl(ctx *parser.ExternCDeclContext) interface{} {
    if a.Phase == 1 {
        for _, member := range ctx.AllExternCMember() {
            if fn := member.ExternCFunctionDecl(); fn != nil {
                a.visitExternCFunction(fn.(*parser.ExternCFunctionDeclContext))
            }
        }
    }
    return nil
}

func (a *Analyzer) visitExternCFunction(ctx *parser.ExternCFunctionDeclContext) {
	name := ctx.IDENTIFIER().GetText()
	if a.currentNamespacePrefix != "" {
		name = a.currentNamespacePrefix + "." + name
	}
	if _, ok := a.currentScope.ResolveLocal(name); ok { return }

	var retType types.Type = types.Void
	if ctx.Type_() != nil { retType = a.resolveType(ctx.Type_()) }
	
	var paramTypes []types.Type
	variadic := false
	if pl := ctx.ExternCParameterList(); pl != nil {
		if pl.ELLIPSIS() != nil { variadic = true }
		for _, p := range pl.AllExternCParameter() {
			if pCtx, ok := p.(*parser.ExternCParameterContext); ok {
				paramTypes = append(paramTypes, a.resolveType(pCtx.Type_()))
			}
		}
	}

	fnType := types.NewFunction(retType, paramTypes, variadic)
	a.currentScope.Define(name, symbol.SymFunc, fnType)
}

func (a *Analyzer) VisitExternCppDecl(ctx *parser.ExternCppDeclContext) interface{} {
    if a.Phase == 0 {
        for _, member := range ctx.AllExternCppMember() {
            c := member.(*parser.ExternCppMemberContext)
            if c.ExternCppClassDecl() != nil {
                a.visitExternCppMember(member)
            } else if c.ExternCppOpaqueClassDecl() != nil {
                a.visitExternCppMember(member)
            } else if c.ExternCppNamespaceDecl() != nil {
                a.visitExternCppMember(member)
            }
        }
        return nil
    }
    if a.Phase == 1 {
        for _, member := range ctx.AllExternCppMember() {
            c := member.(*parser.ExternCppMemberContext)
            if c.ExternCppClassDecl() == nil && c.ExternCppOpaqueClassDecl() == nil {
                a.visitExternCppMember(member)
            } else if c.ExternCppClassDecl() != nil {
                 a.visitExternCppMember(member)
            }
        }
    }
    return nil
}

func (a *Analyzer) visitExternCppMember(ctx parser.IExternCppMemberContext) {
	if ctx == nil { return }
	c := ctx.(*parser.ExternCppMemberContext)

	if fn := c.ExternCppFunctionDecl(); fn != nil {
		a.visitExternCppFunction(fn.(*parser.ExternCppFunctionDeclContext))
	} else if ns := c.ExternCppNamespaceDecl(); ns != nil {
		a.visitExternCppNamespace(ns.(*parser.ExternCppNamespaceDeclContext))
	} else if cl := c.ExternCppClassDecl(); cl != nil {
		a.visitExternCppClass(cl.(*parser.ExternCppClassDeclContext))
	} else if op := c.ExternCppOpaqueClassDecl(); op != nil {
		a.visitExternCppOpaqueClass(op.(*parser.ExternCppOpaqueClassDeclContext))
	}
}

func (a *Analyzer) visitExternCppNamespace(ctx *parser.ExternCppNamespaceDeclContext) {
	pathCtx := ctx.ExternNamespacePath()
	nsName := ""
	if pathCtx != nil {
		ids := pathCtx.AllIDENTIFIER()
		for i, id := range ids {
			if i > 0 { nsName += "." }
			nsName += id.GetText()
		}
	}

	prevPrefix := a.currentNamespacePrefix
	if a.currentNamespacePrefix == "" {
		a.currentNamespacePrefix = nsName
	} else {
		a.currentNamespacePrefix = a.currentNamespacePrefix + "." + nsName
	}

	for _, member := range ctx.AllExternCppMember() {
		a.visitExternCppMember(member)
	}

	a.currentNamespacePrefix = prevPrefix
}

func (a *Analyzer) visitExternCppClass(ctx *parser.ExternCppClassDeclContext) {
	className := ctx.IDENTIFIER().GetText()
	fullName := className
	if a.currentNamespacePrefix != "" {
		fullName = a.currentNamespacePrefix + "." + className
	}

	var st *types.StructType
	if sym, ok := a.currentScope.ResolveLocal(fullName); ok {
		st = sym.Type.(*types.StructType)
	} else {
		st = types.NewStruct(fullName, nil, false)
		a.currentScope.Define(fullName, symbol.SymType, st)
	}

	vtableIndex := 0
	
	prevPrefix := a.currentNamespacePrefix
	a.currentNamespacePrefix = fullName

	for _, member := range ctx.AllExternCppClassMember() {
		if method := member.ExternCppMethodDecl(); method != nil {
			a.visitExternCppMethod(method.(*parser.ExternCppMethodDeclContext), st, &vtableIndex)
		}
	}

	a.currentNamespacePrefix = prevPrefix
}

func (a *Analyzer) visitExternCppOpaqueClass(ctx *parser.ExternCppOpaqueClassDeclContext) {
	name := ctx.IDENTIFIER().GetText()
	fullName := name
	if a.currentNamespacePrefix != "" {
		fullName = a.currentNamespacePrefix + "." + name
	}
	if _, ok := a.currentScope.ResolveLocal(fullName); !ok {
		st := types.NewStruct(fullName, nil, false)
		a.currentScope.Define(fullName, symbol.SymType, st)
	}
}

func (a *Analyzer) visitExternCppFunction(ctx *parser.ExternCppFunctionDeclContext) {
	name := ctx.IDENTIFIER().GetText()
	fullName := name
	if a.currentNamespacePrefix != "" {
		fullName = a.currentNamespacePrefix + "." + name
	}
	if _, ok := a.currentScope.ResolveLocal(fullName); ok { return }

	var retType types.Type = types.Void
	if ctx.Type_() != nil { retType = a.resolveType(ctx.Type_()) }

	var paramTypes []types.Type
	variadic := false
	if pl := ctx.ExternCppParameterList(); pl != nil {
		if pl.ELLIPSIS() != nil { variadic = true }
		for _, p := range pl.AllExternCppParameter() {
			if pCtx, ok := p.(*parser.ExternCppParameterContext); ok {
				paramTypes = append(paramTypes, a.resolveType(pCtx.ExternCppParamType().Type_()))
			}
		}
	}

	fnType := types.NewFunction(retType, paramTypes, variadic)
	a.currentScope.Define(fullName, symbol.SymFunc, fnType)
}

func (a *Analyzer) visitExternCppMethod(ctx *parser.ExternCppMethodDeclContext, parentClass *types.StructType, vIndex *int) {
	methodName := ctx.IDENTIFIER().GetText()
	fullName := parentClass.Name + "_" + methodName
	if _, ok := a.currentScope.ResolveLocal(fullName); ok { return }

	var retType types.Type = types.Void
	if ctx.Type_() != nil { retType = a.resolveType(ctx.Type_()) }

	var paramTypes []types.Type
	paramsCtx := ctx.ExternCppMethodParams()
	
	if self := paramsCtx.ExternCppSelfParam(); self != nil {
		paramTypes = append(paramTypes, types.NewPointer(parentClass))
	}

	if pl := paramsCtx.ExternCppParameterList(); pl != nil {
		for _, p := range pl.AllExternCppParameter() {
			if pCtx, ok := p.(*parser.ExternCppParameterContext); ok {
				paramTypes = append(paramTypes, a.resolveType(pCtx.ExternCppParamType().Type_()))
			}
		}
	} else {
		for _, p := range paramsCtx.AllExternCppParameter() {
			if pCtx, ok := p.(*parser.ExternCppParameterContext); ok {
				paramTypes = append(paramTypes, a.resolveType(pCtx.ExternCppParamType().Type_()))
			}
		}
	}

	fnType := types.NewFunction(retType, paramTypes, false)
	sym := a.currentScope.Define(fullName, symbol.SymFunc, fnType)
	
	if ctx.VIRTUAL() != nil {
		sym.IsVirtual = true
		sym.VTableIndex = *vIndex
		*vIndex++
	}
}

func (a *Analyzer) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	rawName := ctx.IDENTIFIER().GetText()
	
	var fullName string
	var parentName string
	isMethod := false

	if parent := ctx.GetParent(); parent != nil {
		if _, ok := parent.(*parser.ClassMemberContext); ok {
			if cd, ok := parent.GetParent().(*parser.ClassDeclContext); ok {
				parentName = cd.IDENTIFIER().GetText()
				isMethod = true
			}
		} else if _, ok := parent.(*parser.StructMemberContext); ok {
			if sd, ok := parent.GetParent().(*parser.StructDeclContext); ok {
				parentName = sd.IDENTIFIER().GetText()
				isMethod = true
			}
		}
	}

	if !isMethod && ctx.ParameterList() != nil {
		params := ctx.ParameterList().AllParameter()
		if len(params) > 0 {
			if params[0].SELF() != nil {
				selfType := a.resolveType(params[0].Type_())
				if ptr, ok := selfType.(*types.PointerType); ok {
					selfType = ptr.ElementType
				}
				if st, ok := selfType.(*types.StructType); ok {
					parentName = st.Name
					isMethod = true
				}
			}
		}
	}

	if isMethod {
		nameToUse := parentName
		if a.currentNamespacePrefix != "" {
			// Check if parentName already starts with namespace prefix
			prefix := a.currentNamespacePrefix + "."
			if len(parentName) > len(prefix) && parentName[:len(prefix)] == prefix {
				nameToUse = parentName[len(prefix):]
			}
			fullName = a.currentNamespacePrefix + "." + nameToUse + "_" + rawName
		} else {
			fullName = parentName + "_" + rawName
		}
	} else {
		if a.currentNamespacePrefix != "" && rawName != "main" {
			fullName = a.currentNamespacePrefix + "." + rawName
		} else {
			fullName = rawName
		}
	}

	if a.Phase == 1 {
		var retType types.Type = types.Void
		if ctx.ReturnType() != nil {
			if ctx.ReturnType().Type_() != nil {
				retType = a.resolveType(ctx.ReturnType().Type_())
			} else if ctx.ReturnType().TypeList() != nil {
				var tupleTypes []types.Type
				for _, t := range ctx.ReturnType().TypeList().AllType_() {
					tupleTypes = append(tupleTypes, a.resolveType(t))
				}
				retType = types.NewStruct("", tupleTypes, false)
			}
		}

		var paramTypes []types.Type
		if ctx.ParameterList() != nil {
			for _, param := range ctx.ParameterList().AllParameter() {
				paramTypes = append(paramTypes, a.resolveType(param.Type_()))
			}
		}

		var fnType *types.FunctionType
		isAsync := false
		isProcess := false
		
		if es := ctx.ExecutionStrategy(); es != nil {
			if es.ASYNC() != nil {
				isAsync = true
			} else if es.PROCESS() != nil {
				isProcess = true
			}
		}

		if isAsync {
			fnType = types.NewAsyncFunction(retType, paramTypes, false)
		} else if isProcess {
			fnType = types.NewProcessFunction(retType, paramTypes, false)
		} else {
			fnType = types.NewFunction(retType, paramTypes, false)
		}

		if _, ok := a.currentScope.ResolveLocal(fullName); !ok {
			a.currentScope.Define(fullName, symbol.SymFunc, fnType)
		}
	}

	if a.Phase == 2 {
		sym, ok := a.currentScope.Resolve(fullName)
		var retType types.Type = types.Void
		if ok {
			if fn, ok := sym.Type.(*types.FunctionType); ok {
				retType = fn.ReturnType
			}
		}

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
			for _, stmt := range ctx.Block().AllStatement() {
				a.Visit(stmt)
			}
		}
	}
	return nil
}

func (a *Analyzer) VisitMutatingDecl(ctx *parser.MutatingDeclContext) interface{} {
	name := ctx.IDENTIFIER(0).GetText()
	
	var retType types.Type = types.Void
	if ctx.ReturnType() != nil {
		if ctx.ReturnType().Type_() != nil {
			retType = a.resolveType(ctx.ReturnType().Type_())
		}
	}

	var paramTypes []types.Type
	selfParamName := ctx.IDENTIFIER(1).GetText()
	selfType := a.resolveType(ctx.Type_()) 
	paramTypes = append(paramTypes, selfType)
	
	for _, param := range ctx.AllParameter() {
		pType := a.resolveType(param.Type_())
		paramTypes = append(paramTypes, pType)
	}

	var structName string
	base := selfType
	if ptr, ok := base.(*types.PointerType); ok { base = ptr.ElementType }
	if st, ok := base.(*types.StructType); ok { structName = st.Name }
	
	fullName := name
	if structName != "" { fullName = structName + "_" + name }

	if a.Phase == 1 || a.Phase == 0 {
		fnType := types.NewFunction(retType, paramTypes, false)
		if _, ok := a.currentScope.ResolveLocal(fullName); !ok {
			a.currentScope.Define(fullName, symbol.SymFunc, fnType)
		}
	}

	if a.Phase == 2 || a.Phase == 0 {
		a.currentFuncRetType = retType
		a.pushScope(ctx)
		defer func() { a.popScope(); a.currentFuncRetType = nil }()

		a.currentScope.Define(selfParamName, symbol.SymVar, selfType)
		for _, param := range ctx.AllParameter() {
			pName := param.IDENTIFIER().GetText()
			pType := a.resolveType(param.Type_())
			a.currentScope.Define(pName, symbol.SymVar, pType)
		}

		if ctx.Block() != nil {
			a.scopes[ctx.Block()] = a.currentScope
			for _, stmt := range ctx.Block().AllStatement() { a.Visit(stmt) }
		}
	}
	return nil
}

func (a *Analyzer) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	if ctx.TuplePattern() != nil {
		if a.Phase == 2 {
			var rhsType types.Type = types.Void
			if ctx.Expression() != nil {
				rhsType = a.Visit(ctx.Expression()).(types.Type)
			}
			st, isStruct := rhsType.(*types.StructType)
			ids := ctx.TuplePattern().AllIDENTIFIER()
			for i, id := range ids {
				name := id.GetText()
				var fieldType types.Type = types.I64
				if isStruct && i < len(st.Fields) {
					fieldType = st.Fields[i]
				}
				if _, ok := a.currentScope.ResolveLocal(name); !ok {
					a.currentScope.Define(name, symbol.SymVar, fieldType)
				}
			}
		}
		return nil
	}

	if ctx.IDENTIFIER() == nil { return nil }
	name := ctx.IDENTIFIER().GetText()

	if a.Phase == 1 && a.currentScope.Parent == nil {
		if _, ok := a.currentScope.ResolveLocal(name); !ok {
			var typ types.Type
			if ctx.Type_() != nil { typ = a.resolveType(ctx.Type_()) }
			if typ == nil { typ = types.I64 }
			a.currentScope.Define(name, symbol.SymVar, typ)
		}
		return nil
	}

	if a.Phase == 2 {
		sym, _ := a.currentScope.ResolveLocal(name)
		
		if sym == nil && ctx.Type_() != nil {
			typ := a.resolveType(ctx.Type_())
			sym = a.currentScope.Define(name, symbol.SymVar, typ)
		}

		var typ types.Type
		if sym != nil { typ = sym.Type }
		
		if ctx.Expression() != nil {
			exprType := a.Visit(ctx.Expression()).(types.Type)
			if typ == nil {
				typ = exprType
				if sym == nil { a.currentScope.Define(name, symbol.SymVar, typ) }
			} else {
				if exprType != types.Void && !areTypesCompatible(exprType, typ) {
					a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, 
						"Type mismatch in variable '%s'", name)
				}
			}
		}
		if sym == nil && typ == nil { a.currentScope.Define(name, symbol.SymVar, types.I64) }
	}
	return nil
}

func (a *Analyzer) VisitConstDecl(ctx *parser.ConstDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	if a.currentNamespacePrefix != "" && name != "main" {
		name = a.currentNamespacePrefix + "." + name
	}

	if a.Phase == 1 {
		if _, ok := a.currentScope.ResolveLocal(name); !ok {
			var typ types.Type
			if ctx.Type_() != nil { typ = a.resolveType(ctx.Type_()) }
			if typ == nil { typ = types.I64 }
			a.currentScope.Define(name, symbol.SymConst, typ)
		}
	}

	if a.Phase == 2 {
		sym, _ := a.currentScope.ResolveLocal(name)
		if ctx.Expression() != nil {
			exprType := a.Visit(ctx.Expression()).(types.Type)
			if sym != nil && sym.Type == types.I64 && ctx.Type_() == nil {
				sym.Type = exprType
			}
		}
	}
	return nil
}

func (a *Analyzer) VisitClassDecl(ctx *parser.ClassDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	if a.currentNamespacePrefix != "" {
		name = a.currentNamespacePrefix + "." + name
	}

	if _, ok := a.currentScope.ResolveLocal(name); !ok {
		st := types.NewClass(name, nil, false)
		a.currentScope.Define(name, symbol.SymType, st)
	}

	if a.Phase == 1 {
		sym, _ := a.currentScope.Resolve(name)
		if sym != nil {
			st := sym.Type.(*types.StructType)
			if len(st.Fields) == 0 {
				var fields []types.Type
				indices := make(map[string]int)
				fieldCount := 0

				for _, member := range ctx.AllClassMember() {
					if f := member.ClassField(); f != nil {
						fields = append(fields, a.resolveType(f.Type_()))
						indices[f.IDENTIFIER().GetText()] = fieldCount
						fieldCount++
					}
				}
				st.Fields = fields
				a.structIndices[name] = indices
			}
		}
	}

	for _, member := range ctx.AllClassMember() {
		if m := member.FunctionDecl(); m != nil { a.Visit(m) }
		if d := member.DeinitDecl(); d != nil { a.Visit(d) }
	}
	return nil
}

func (a *Analyzer) VisitEnumDecl(ctx *parser.EnumDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	if a.currentNamespacePrefix != "" {
		name = a.currentNamespacePrefix + "." + name
	}

	if _, ok := a.currentScope.ResolveLocal(name); !ok {
		a.currentScope.Define(name, symbol.SymType, types.I32)
	}

	if a.Phase == 1 {
		for _, m := range ctx.AllEnumMember() {
			memberName := name + "." + m.IDENTIFIER().GetText()
			if _, ok := a.currentScope.ResolveLocal(memberName); !ok {
				a.currentScope.Define(memberName, symbol.SymConst, types.I32)
			}
		}
	}
	return nil
}

func (a *Analyzer) VisitMethodDecl(ctx *parser.MethodDeclContext) interface{} { return nil }

func (a *Analyzer) VisitDeinitDecl(ctx *parser.DeinitDeclContext) interface{} {
	parentName := ""
	if classDecl, ok := ctx.GetParent().(*parser.ClassMemberContext).GetParent().(*parser.ClassDeclContext); ok {
		parentName = classDecl.IDENTIFIER().GetText()
	}

	methodName := parentName + "_deinit"
	fullName := methodName
	if a.currentNamespacePrefix != "" {
		fullName = a.currentNamespacePrefix + "." + methodName
	}

	var paramTypes []types.Type
	selfParam := ctx.IDENTIFIER()
	selfTypeCtx := ctx.Type_()
	
	selfType := a.resolveType(selfTypeCtx)
	paramTypes = append(paramTypes, selfType)

	fnType := types.NewFunction(types.Void, paramTypes, false)
	a.currentScope.Define(fullName, symbol.SymFunc, fnType)

	if a.Phase == 2 || a.Phase == 0 {
		a.pushScope(ctx)
		defer a.popScope()

		a.currentScope.Define(selfParam.GetText(), symbol.SymVar, selfType)

		if ctx.Block() != nil {
			a.scopes[ctx.Block()] = a.currentScope
			for _, stmt := range ctx.Block().AllStatement() {
				a.Visit(stmt)
			}
		}
	}

	return nil
}