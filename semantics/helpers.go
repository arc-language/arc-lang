package semantics

import (
	"strconv"

	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

// resolveType converts AST Type nodes into internal Type representations
func (a *Analyzer) resolveType(ctx parser.ITypeContext) types.Type {
	if ctx == nil {
		return types.Void
	}
	tc := ctx.(*parser.TypeContext)

	// 1. Primitives
	if tc.PrimitiveType() != nil {
		name := tc.PrimitiveType().GetText()
		if s, ok := a.currentScope.Resolve(name); ok && s.Kind == symbol.SymType {
			return s.Type
		}
		a.bag.Report(a.file, tc.GetStart().GetLine(), 0, "Unknown primitive type '%s'", name)
		return types.I64
	}
	
	// 2. Pointers
	if tc.PointerType() != nil {
		return types.NewPointer(a.resolveType(tc.PointerType().Type_()))
	}

	// 3. References
	if tc.ReferenceType() != nil {
		return types.NewPointer(a.resolveType(tc.ReferenceType().Type_()))
	}

	// 4. Generic/Qualified Types (vector<int>, map<k,v>, MyClass)
	if tc.IDENTIFIER() != nil || tc.QualifiedType() != nil {
		var name string
		var genericArgs parser.IGenericArgsContext

		if tc.QualifiedType() != nil {
			qt := tc.QualifiedType()
			for i, id := range qt.AllIDENTIFIER() {
				if i > 0 { name += "." }
				name += id.GetText()
			}
			genericArgs = qt.GenericArgs()
		} else {
			name = tc.IDENTIFIER().GetText()
			genericArgs = tc.GenericArgs()
		}

		// Resolve base symbol
		s, ok := a.currentScope.Resolve(name)
		if !ok || s.Kind != symbol.SymType {
			a.bag.Report(a.file, tc.GetStart().GetLine(), 0, "Unknown type '%s'", name)
			return types.I64
		}

		// Handle Generics
		if genericArgs != nil {
			if name == "array" {
				args := genericArgs.GenericArgList().AllGenericArg()
				if len(args) == 2 {
					var elemType types.Type
					if tCtx := args[0].Type_(); tCtx != nil {
						elemType = a.resolveType(tCtx)
					}
					var length int64
					if exprCtx := args[1].Expression(); exprCtx != nil {
						if lit := exprCtx.GetText(); lit != "" {
							if val, err := strconv.ParseInt(lit, 0, 64); err == nil {
								length = val
							}
						}
					}
					if elemType != nil && length > 0 {
						return types.NewArray(elemType, length)
					}
				}
			}
			return types.NewPointer(types.I8) 
		}

		return s.Type
	}
	
	return types.I64
}

// areTypesCompatible checks if two types can be used together (e.g. assignment, math)
func areTypesCompatible(src, dest types.Type) bool {
	if src == nil || dest == nil { return false }
	if src.Equal(dest) {
		return true
	}
	// Implicit int casting
	if types.IsInteger(src) && types.IsInteger(dest) {
		return true 
	}
	// Implicit float casting
	if types.IsFloat(src) && types.IsFloat(dest) {
		return true
	}
	// void* compatibility
	if srcPtr, sOk := src.(*types.PointerType); sOk {
		if destPtr, dOk := dest.(*types.PointerType); dOk {
			if srcPtr.ElementType == types.Void || destPtr.ElementType == types.Void {
				return true
			}
		}
	}
	// Array compatibility (recursive)
	if srcArr, sOk := src.(*types.ArrayType); sOk {
		if destArr, dOk := dest.(*types.ArrayType); dOk {
			if srcArr.Length == destArr.Length {
				return areTypesCompatible(srcArr.ElementType, destArr.ElementType)
			}
		}
	}
	
	return false
}