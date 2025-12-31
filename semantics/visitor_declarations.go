// --- FILE: semantics/visitor_declarations.go ---

package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

func (a *Analyzer) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	// Pass 1.0: Register Types (Structs, Classes, Enums) to allow forward references
	for _, decl := range ctx.AllTopLevelDecl() {
		if decl.StructDecl() != nil {
			name := decl.StructDecl().IDENTIFIER().GetText()
			st := types.NewStruct(name, nil, false)
			a.currentScope.Define(name, symbol.SymType, st)
		} else if decl.ClassDecl() != nil {
			name := decl.ClassDecl().IDENTIFIER().GetText()
			// Classes are treated as structs implicitly passed by reference in many places
			st := types.NewStruct(name, nil, false) 
			a.currentScope.Define(name, symbol.SymType, st)
		} else if decl.EnumDecl() != nil {
			name := decl.EnumDecl().IDENTIFIER().GetText()
			// Enums default to Int32 for now
			a.currentScope.Define(name, symbol.SymType, types.I32)
		}
	}

	// Pass 1.1: Process everything else
	for _, decl := range ctx.AllTopLevelDecl() {
		a.Visit(decl)
	}
	
	// Process Namespace declarations if any (often at top level)
	for _, ns := range ctx.AllNamespaceDecl() {
		a.Visit(ns)
	}
	
	return nil
}

func (a *Analyzer) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext) interface{} {
	if ctx.FunctionDecl() != nil { return a.Visit(ctx.FunctionDecl()) }
	if ctx.VariableDecl() != nil { return a.Visit(ctx.VariableDecl()) }
	if ctx.StructDecl() != nil { return a.Visit(ctx.StructDecl()) }
	if ctx.ClassDecl() != nil { return a.Visit(ctx.ClassDecl()) } // NEW
	if ctx.EnumDecl() != nil { return a.Visit(ctx.EnumDecl()) }   // NEW
	if ctx.ExternDecl() != nil { return a.Visit(ctx.ExternDecl()) }
	return nil
}

func (a *Analyzer) VisitNamespaceDecl(ctx *parser.NamespaceDeclContext) interface{} {
	// Simple namespace handling: just prefixing symbols is handled in IRGen.
	// In semantics, we might want to push a scope or just allow the children to validate.
	// For now, we continue visiting children.
	// Note: Real namespacing would require Scope hierarchy adjustments.
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
		// Methods are processed as functions but could be scoped here
		if m := member.FunctionDecl(); m != nil {
			// Push struct scope/namespace if needed, or mangle name
			// For Pass 1, we just visit it to check semantics
			a.Visit(m)
		}
	}
	structType.Fields = fieldTypes
	a.structIndices[name] = fieldIndices
	return nil
}

func (a *Analyzer) VisitClassDecl(ctx *parser.ClassDeclContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	sym, _ := a.currentScope.Resolve(name)
	classType := sym.Type.(*types.StructType) // Backend representation is struct
	
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
			// Check methods
			a.Visit(m)
		}
	}
	classType.Fields = fieldTypes
	a.structIndices[name] = fieldIndices // Reuse struct indices for classes
	return nil
}

func (a *Analyzer) VisitEnumDecl(ctx *parser.EnumDeclContext) interface{} {
	// Define enum members as integer constants in the current scope
	// e.g. enum Color { Red, Green } -> defines Color.Red, Color.Green
	// Note: Simple implementation defines them as globals or scoped constants
	enumName := ctx.IDENTIFIER().GetText()
	
	for _, member := range ctx.AllEnumMember() {
		memName := member.IDENTIFIER().GetText()
		fullName := enumName + "." + memName // logical name
		a.currentScope.Define(fullName, symbol.SymConst, types.I32)
	}
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
	
	// Handle method naming (Class.Method) if inside class context
	// This simplified analyzer assumes flat structure or relies on unique names for now.
	// In a full implementation, we check parent context.
	
	a.currentScope.Define(name, symbol.SymFunc, retType)
	a.currentFuncRetType = retType
	
	a.pushScope(ctx)
	defer a.popScope()
	
	if ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			if param.SELF() != nil {
				// implicit self type
				continue 
			}
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