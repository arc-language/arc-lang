package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

func (a *Analyzer) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	// --- Pass 1: Register Types (Structs, Classes, Enums) ---
	// This allows functions declared before structs to still refer to them.
	for _, decl := range ctx.AllTopLevelDecl() {
		if decl.StructDecl() != nil {
			name := decl.StructDecl().IDENTIFIER().GetText()
			st := types.NewStruct(name, nil, false)
			a.currentScope.Define(name, symbol.SymType, st)
		} else if decl.ClassDecl() != nil {
			name := decl.ClassDecl().IDENTIFIER().GetText()
			// Classes are effectively structs passed by reference
			st := types.NewStruct(name, nil, false)
			a.currentScope.Define(name, symbol.SymType, st)
		} else if decl.EnumDecl() != nil {
			name := decl.EnumDecl().IDENTIFIER().GetText()
			// Enums map to Int32 for now
			a.currentScope.Define(name, symbol.SymType, types.I32)
		}
	}

	// --- Pass 2: Process Declarations and Statements ---
	for _, decl := range ctx.AllTopLevelDecl() {
		a.Visit(decl)
	}
	
	// Visit Namespace declarations (if any)
	for _, ns := range ctx.AllNamespaceDecl() {
		a.Visit(ns)
	}

	return nil
}

func (a *Analyzer) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext) interface{} {
	if ctx.FunctionDecl() != nil { return a.Visit(ctx.FunctionDecl()) }
	if ctx.VariableDecl() != nil { return a.Visit(ctx.VariableDecl()) }
	if ctx.StructDecl() != nil { return a.Visit(ctx.StructDecl()) }
	if ctx.ClassDecl() != nil { return a.Visit(ctx.ClassDecl()) }
	if ctx.EnumDecl() != nil { return a.Visit(ctx.EnumDecl()) }
	if ctx.ExternDecl() != nil { return a.Visit(ctx.ExternDecl()) }
	return nil
}

func (a *Analyzer) VisitNamespaceDecl(ctx *parser.NamespaceDeclContext) interface{} {
	// In a full implementation, this might push a named scope.
	// For now, we rely on the IRGen to handle namespace mangling, 
	// but we must visit children to validate them.
	return nil
}

func (a *Analyzer) VisitExternDecl(ctx *parser.ExternDeclContext) interface{} {
	ns := ""
	if ctx.IDENTIFIER() != nil {
		ns = ctx.IDENTIFIER().GetText()
	}

	for _, member := range ctx.AllExternMember() {
		if fnCtx := member.ExternFunctionDecl(); fnCtx != nil {
			name := fnCtx.IDENTIFIER().GetText()
			
			var retType types.Type = types.Void
			if fnCtx.Type_() != nil {
				retType = a.resolveType(fnCtx.Type_())
			}

			// Construct lookup name (e.g., "io.printf")
			lookupName := name
			if ns != "" {
				lookupName = ns + "." + name
			}

			// Externs are defined in the current (global) scope
			a.currentScope.Define(lookupName, symbol.SymFunc, retType)
		}
	}
	return nil
}

func (a *Analyzer) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
    // (Ensure this matches previous fix, re-listing for completeness)
	name := ctx.IDENTIFIER().GetText()
	var retType types.Type = types.Void
	if ctx.ReturnType() != nil && ctx.ReturnType().Type_() != nil {
		retType = a.resolveType(ctx.ReturnType().Type_())
	}
	a.currentScope.Define(name, symbol.SymFunc, retType)
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
	name := ctx.IDENTIFIER().GetText()
	if _, ok := a.currentScope.ResolveLocal(name); ok {
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Redeclaration of '%s'", name)
		return nil
	}
	var typ types.Type
	if ctx.Type_() != nil { typ = a.resolveType(ctx.Type_()) }
	if ctx.Expression() != nil {
		exprType := a.Visit(ctx.Expression()).(types.Type)
		if typ == nil { typ = exprType }
	}
	if typ == nil { typ = types.I64 }
	a.currentScope.Define(name, symbol.SymVar, typ)
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
		a.currentScope.Define(name+"."+m.IDENTIFIER().GetText(), symbol.SymConst, types.I32)
	}
	return nil
}

func (a *Analyzer) VisitConstDecl(ctx *parser.ConstDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	if _, ok := a.currentScope.ResolveLocal(name); ok {
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Redeclaration of '%s'", name)
		return nil
	}
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