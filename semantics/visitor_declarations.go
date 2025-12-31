package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

func (a *Analyzer) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	for _, decl := range ctx.AllTopLevelDecl() {
		if decl.StructDecl() != nil {
			name := decl.StructDecl().IDENTIFIER().GetText()
			st := types.NewStruct(name, nil, false)
			a.currentScope.Define(name, symbol.SymType, st)
		}
	}
	for _, decl := range ctx.AllTopLevelDecl() {
		a.Visit(decl)
	}
	return nil
}

func (a *Analyzer) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext) interface{} {
	if ctx.FunctionDecl() != nil { return a.Visit(ctx.FunctionDecl()) }
	if ctx.VariableDecl() != nil { return a.Visit(ctx.VariableDecl()) }
	if ctx.StructDecl() != nil { return a.Visit(ctx.StructDecl()) }
	if ctx.ExternDecl() != nil { return a.Visit(ctx.ExternDecl()) } // NEW
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
			
			lookupName := name
			if ns != "" {
				lookupName = ns + "." + name
			}
			
			a.currentScope.Define(lookupName, symbol.SymFunc, retType)
		}
	}
	return nil
}

// ... (Struct, Variable, and Function visitors remain the same as previous updates) ...
func (a *Analyzer) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	sym, _ := a.currentScope.Resolve(name)
	structType := sym.Type.(*types.StructType)
	var fieldTypes []types.Type
	fieldIndices := make(map[string]int)
	idx := 0
	for _, member := range ctx.AllStructMember() {
		if f := member.StructField(); f != nil {
			fName := f.IDENTIFIER().GetText()
			fType := a.resolveType(f.Type_())
			fieldTypes = append(fieldTypes, fType)
			fieldIndices[fName] = idx
			idx++
		}
	}
	structType.Fields = fieldTypes
	a.structIndices[name] = fieldIndices
	return nil
}

func (a *Analyzer) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	var typ types.Type
	if ctx.Type_() != nil { typ = a.resolveType(ctx.Type_()) }
	if ctx.Expression() != nil {
		exprType := a.Visit(ctx.Expression()).(types.Type)
		if typ == nil { typ = exprType }
	} else if typ == nil { typ = types.Void }
	a.currentScope.Define(name, symbol.SymVar, typ)
	return nil
}

func (a *Analyzer) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	var retType types.Type = types.Void
	if ctx.ReturnType() != nil && ctx.ReturnType().Type_() != nil {
		retType = a.resolveType(ctx.ReturnType().Type_())
	}
	a.currentScope.Define(name, symbol.SymFunc, retType)
	a.currentFuncRetType = retType
	a.pushScope(ctx)
	defer a.popScope()
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
	return nil
}