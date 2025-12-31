package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/pkg/symbol"
)

func (a *Analyzer) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	// Pass 1.1: Register Types (Structs/Classes) first so they can reference each other
	for _, decl := range ctx.AllTopLevelDecl() {
		if decl.StructDecl() != nil {
			a.registerStructHead(decl.StructDecl())
		}
	}
	
	// Pass 1.2: Process everything else
	for _, decl := range ctx.AllTopLevelDecl() {
		a.Visit(decl)
	}
	return nil
}

func (a *Analyzer) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext) interface{} {
	if ctx.FunctionDecl() != nil { return a.Visit(ctx.FunctionDecl()) }
	if ctx.VariableDecl() != nil { return a.Visit(ctx.VariableDecl()) }
	// Structs handled in pre-pass, but we might visit members if needed
	if ctx.StructDecl() != nil { return a.Visit(ctx.StructDecl()) }
	return nil
}

// Pre-register struct name so we can handle recursive pointers
func (a *Analyzer) registerStructHead(ctx *parser.StructDeclContext) {
	name := ctx.IDENTIFIER().GetText()
	// Create empty struct type
	structType := types.NewStruct(name, nil, false)
	a.currentScope.Define(name, symbol.SymType, structType)
}

func (a *Analyzer) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	sym, _ := a.currentScope.Resolve(name)
	structType := sym.Type.(*types.StructType)
	
	// Populate fields
	var fieldTypes []types.Type
	
	for _, member := range ctx.AllStructMember() {
		if field := member.StructField(); field != nil {
			fieldName := field.IDENTIFIER().GetText()
			fieldType := a.resolveType(field.Type_())
			fieldTypes = append(fieldTypes, fieldType)
			
			// Note: We don't store field names in types.StructType directly in LLVM style,
			// but for Semantics, we need to map names to indices/types.
			// We can store this in the Symbol or a separate registry.
			// For simplicity here, we assume the Symbol holds a map or we rely on order.
			// In a real implementation, you'd extend Symbol or Scope to hold Field info.
			a.defineField(structType, fieldName, fieldType, len(fieldTypes)-1)
		}
	}
	
	// Update the type with fields
	structType.Fields = fieldTypes
	return nil
}

func (a *Analyzer) defineField(s *types.StructType, name string, t types.Type, idx int) {
	// In a full implementation, you'd store this in a lookup table:
	// a.structFields[s.Name][name] = FieldInfo{Type: t, Index: idx}
	// For now, we will assume standard LLVM struct behavior.
}

func (a *Analyzer) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	if _, exists := a.currentScope.ResolveLocal(name); exists {
		a.bag.Report(a.file, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "Redeclaration of '%s'", name)
		return nil
	}

	var declType types.Type
	if ctx.Type_() != nil {
		declType = a.resolveType(ctx.Type_())
	}

	if ctx.Expression() != nil {
		exprType := a.Visit(ctx.Expression()).(types.Type)
		if declType == nil {
			declType = exprType
		} else if !areTypesCompatible(exprType, declType) {
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Cannot assign %s to %s", exprType, declType)
		}
	} else if declType == nil {
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Variable '%s' missing type", name)
		declType = types.Void
	}

	a.currentScope.Define(name, symbol.SymVar, declType)
	return nil
}

func (a *Analyzer) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	var retType types.Type = types.Void
	if ctx.ReturnType() != nil && ctx.ReturnType().Type_() != nil {
		retType = a.resolveType(ctx.ReturnType().Type_())
	}

	// Define Function Symbol
	// Note: We should construct a proper FunctionType (args -> ret)
	// For now we just store RetType to allow recursion logic
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