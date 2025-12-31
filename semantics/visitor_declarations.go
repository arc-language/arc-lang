package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/pkg/symbol"
)

func (a *Analyzer) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	for _, decl := range ctx.AllTopLevelDecl() {
		a.Visit(decl)
	}
	return nil
}

func (a *Analyzer) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext) interface{} {
	if ctx.FunctionDecl() != nil { return a.Visit(ctx.FunctionDecl()) }
	if ctx.VariableDecl() != nil { return a.Visit(ctx.VariableDecl()) }
	if ctx.StructDecl() != nil { return a.Visit(ctx.StructDecl()) }
	return nil
}

func (a *Analyzer) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	// 1. Check Redeclaration
	if _, exists := a.currentScope.ResolveLocal(name); exists {
		a.bag.Report(a.file, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			"Redeclaration of variable '%s'", name)
		return nil
	}

	// 2. Resolve Declared Type (if exists)
	var declaredType types.Type
	if ctx.Type_() != nil {
		declaredType = a.resolveType(ctx.Type_())
	}

	// 3. Check Initializer Expression
	if ctx.Expression() != nil {
		exprType := a.Visit(ctx.Expression()).(types.Type)
		
		if declaredType == nil {
			// Type Inference
			declaredType = exprType
		} else {
			// Type Checking
			if !areTypesCompatible(exprType, declaredType) {
				a.bag.Report(a.file, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
					"Type Mismatch: Cannot assign '%s' to variable of type '%s'", 
					exprType.String(), declaredType.String())
			}
		}
	} else if declaredType == nil {
		a.bag.Report(a.file, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			"Variable '%s' must have a type annotation or an initializer", name)
		declaredType = types.Void // Prevent nil panic downstream
	}

	// 4. Define Symbol
	a.currentScope.Define(name, symbol.SymVar, declaredType)
	return nil
}

func (a *Analyzer) VisitFunctionDecl(ctx *parser.FunctionDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()

	// 1. Resolve Return Type
	// FIX: Explicitly define as interface type to avoid mismatch with *types.VoidType
	var retType types.Type = types.Void
	
	if ctx.ReturnType() != nil && ctx.ReturnType().Type_() != nil {
		retType = a.resolveType(ctx.ReturnType().Type_())
	}

	// 2. Create Function Symbol
	// FIX: Use underscore to ignore the returned symbol since we don't use it locally
	_ = a.currentScope.Define(name, symbol.SymFunc, retType)
	
	// 3. Enter Scope
	a.currentFuncRetType = retType
	a.pushScope(ctx)
	defer a.popScope()

	// 4. Process Parameters
	if ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			pName := param.IDENTIFIER().GetText()
			pType := a.resolveType(param.Type_())
			
			// Define param in function scope
			a.currentScope.Define(pName, symbol.SymVar, pType)
		}
	}

	// 5. Visit Body
	if ctx.Block() != nil {
		// Map the Block to the CURRENT function scope.
		a.scopes[ctx.Block()] = a.currentScope
		
		for _, stmt := range ctx.Block().AllStatement() {
			a.Visit(stmt)
		}
	}

	return nil
}

func (a *Analyzer) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	// Placeholder for struct type
	structType := types.NewStruct(name, nil, false)
	a.currentScope.Define(name, symbol.SymType, structType)
	
	return nil
}