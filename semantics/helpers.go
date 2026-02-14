package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

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
	
	// 2. Raw Pointer (void*)
	if tc.RAWPTR() != nil {
		return types.NewPointer(types.Void)
	}

	// 3. Function Types
	if tc.FunctionType() != nil {
		ft := tc.FunctionType()
		var retType types.Type = types.Void
		if ft.ReturnType() != nil {
			retType = a.resolveType(ft.ReturnType().Type_())
		}
		
		var paramTypes []types.Type
		if ft.TypeList() != nil {
			for _, t := range ft.TypeList().AllType_() {
				paramTypes = append(paramTypes, a.resolveType(t))
			}
		}
		
		isAsync := false
		isProcess := false
		
		if es := ft.ExecutionStrategy(); es != nil {
			if es.ASYNC() != nil {
				isAsync = true
			} else if es.PROCESS() != nil {
				isProcess = true
			}
		}

		if isAsync {
			return types.NewAsyncFunction(retType, paramTypes, false)
		} else if isProcess {
			return types.NewProcessFunction(retType, paramTypes, false)
		}
		return types.NewFunction(retType, paramTypes, false)
	}

	// 4. Collection Types (vector[T], map[K]V, etc.)
	if tc.CollectionType() != nil {
		ct := tc.CollectionType()
		baseName := ct.IDENTIFIER().GetText()
		
		var typeArgs []types.Type
		for _, tCtx := range ct.AllType_() {
			typeArgs = append(typeArgs, a.resolveType(tCtx))
		}

		if baseName == "vector" && len(typeArgs) == 1 {
			// Vector mapped to dynamic array (length 0 usually implies dynamic in simple builders)
			return types.NewArray(typeArgs[0], 0)
		}
		
		if sym, ok := a.currentScope.Resolve(baseName); ok {
			return sym.Type
		}
		
		return types.Void
	}

	// 5. Generic/Qualified Types
	if tc.IDENTIFIER() != nil || tc.QualifiedType() != nil {
		var name string

		if tc.QualifiedType() != nil {
			qt := tc.QualifiedType()
			for i, id := range qt.AllIDENTIFIER() {
				if i > 0 { name += "." }
				name += id.GetText()
			}
		} else {
			name = tc.IDENTIFIER().GetText()
		}

		s, ok := a.currentScope.Resolve(name)
		
		if !ok && a.currentNamespacePrefix != "" && tc.QualifiedType() == nil {
			s, ok = a.currentScope.Resolve(a.currentNamespacePrefix + "." + name)
		}

		if !ok || s.Kind != symbol.SymType {
			a.bag.Report(a.file, tc.GetStart().GetLine(), 0, "Unknown type '%s'", name)
			return types.I64
		}

		return s.Type
	}
	
	return types.I64
}

func areTypesCompatible(src, dest types.Type) bool {
	if src == nil || dest == nil { return false }
	if src.Equal(dest) {
		return true
	}
	if types.IsInteger(src) && types.IsInteger(dest) {
		return true 
	}
	if types.IsFloat(src) && types.IsFloat(dest) {
		return true
	}
	if srcPtr, sOk := src.(*types.PointerType); sOk {
		if destPtr, dOk := dest.(*types.PointerType); dOk {
			if srcPtr.ElementType == types.Void || destPtr.ElementType == types.Void {
				return true
			}
			if types.IsInteger(srcPtr.ElementType) && types.IsInteger(destPtr.ElementType) {
				if srcPtr.ElementType.BitSize() == 8 && destPtr.ElementType.BitSize() == 8 {
					return true
				}
			}
		}
	}
	if srcArr, sOk := src.(*types.ArrayType); sOk {
		if destArr, dOk := dest.(*types.ArrayType); dOk {
			// Allow compatible arrays if lengths match (or both are dynamic/0)
			if srcArr.Length == destArr.Length {
				return areTypesCompatible(srcArr.ElementType, destArr.ElementType)
			}
		}
	}
	
	return false
}