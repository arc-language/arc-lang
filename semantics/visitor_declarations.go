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
	name := ctx.IDENTIFIER().GetText()
	
	var retType types.Type = types.Void
	if ctx.ReturnType() != nil && ctx.ReturnType().Type_() != nil {
		retType = a.resolveType(ctx.ReturnType().Type_())
	}

	// 1. Define function symbol in the current (parent) scope
	a.currentScope.Define(name, symbol.SymFunc, retType)
	
	// Track return type for checking return statements inside
	prevRetType := a.currentFuncRetType
	a.currentFuncRetType = retType
	
	// 2. Create and Push Function Scope
	a.pushScope(ctx)
	defer func() {
		a.popScope()
		a.currentFuncRetType = prevRetType
	}()

	// 3. Define Parameters in Function Scope
	if ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			if param.SELF() != nil { continue }
			
			pName := param.IDENTIFIER().GetText()
			pType := a.resolveType(param.Type_())
			
			a.currentScope.Define(pName, symbol.SymVar, pType)
		}
	}

	// 4. Visit Body
	if ctx.Block() != nil {
		// CRITICAL: Map the block context to the *current* function scope.
		// This tells IRGen: "When you see this block, don't create a new scope, 
		// use the one created for the function."
		a.scopes[ctx.Block()] = a.currentScope

		// Visit statements manually to ensure we use 'a' (Analyzer) and not BaseVisitor
		for _, stmt := range ctx.Block().AllStatement() {
			a.Visit(stmt)
		}
	}

	return nil
}

func (a *Analyzer) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	// Check for redeclaration in local scope
	if _, ok := a.currentScope.ResolveLocal(name); ok {
		a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, "Redeclaration of variable '%s'", name)
		return nil
	}

	var typ types.Type

	// Explicit Type?
	if ctx.Type_() != nil {
		typ = a.resolveType(ctx.Type_())
	}

	// Initializer?
	if ctx.Expression() != nil {
		exprType := a.Visit(ctx.Expression()).(types.Type)
		
		// Type Inference
		if typ == nil {
			typ = exprType
		} else {
			// Type Checking
			if !areTypesCompatible(exprType, typ) {
				a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, 
					"Type mismatch in variable '%s': expected %s, got %s", 
					name, typ.String(), exprType.String())
			}
		}
	}

	// Fallback if semantic analysis fails to determine type
	if typ == nil {
		typ = types.I64 
	}

	// Define variable in current scope
	a.currentScope.Define(name, symbol.SymVar, typ)
	return nil
}

func (a *Analyzer) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	// Retrieve the type definition created in Pass 1
	sym, _ := a.currentScope.Resolve(name)
	if sym == nil { return nil } // Should not happen
	
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
		
		if m := member.FunctionDecl(); m != nil {
			// Methods would be visited here to validate body
			// You might push a scope representing 'self' here
			a.Visit(m)
		}
	}

	// Update the struct type with actual field info
	structType.Fields = fieldTypes
	a.structIndices[name] = fieldIndices
	
	return nil
}

func (a *Analyzer) VisitClassDecl(ctx *parser.ClassDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	sym, _ := a.currentScope.Resolve(name)
	if sym == nil { return nil }

	classType := sym.Type.(*types.StructType)
	
	var fieldTypes []types.Type
	fieldIndices := make(map[string]int)
	idx := 0

	for _, member := range ctx.AllClassMember() {
		if f := member.ClassField(); f != nil {
			fName := f.IDENTIFIER().GetText()
			fType := a.resolveType(f.Type_())
			
			fieldTypes = append(fieldTypes, fType)
			fieldIndices[fName] = idx
			idx++
		}
		if m := member.FunctionDecl(); m != nil {
			a.Visit(m)
		}
	}

	classType.Fields = fieldTypes
	a.structIndices[name] = fieldIndices
	return nil
}

func (a *Analyzer) VisitEnumDecl(ctx *parser.EnumDeclContext) interface{} {
	enumName := ctx.IDENTIFIER().GetText()
	
	// Define enum members as constants in the scope
	// e.g., Color.Red
	for _, member := range ctx.AllEnumMember() {
		memName := member.IDENTIFIER().GetText()
		
		// We use a dot notation for logical grouping in symbol table
		fullName := enumName + "." + memName
		
		// For now, enums are just Integers
		a.currentScope.Define(fullName, symbol.SymConst, types.I32)
	}
	return nil
}