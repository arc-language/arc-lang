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
		
		// Resolve the generic type arguments inside [ ]
		var typeArgs []types.Type
		for _, tCtx := range ct.AllType_() {
			typeArgs = append(typeArgs, a.resolveType(tCtx))
		}

		// Simple mapping for standard collections
		if baseName == "vector" && len(typeArgs) == 1 {
			// Vector is dynamic array
			return types.NewSlice(typeArgs[0])
		}
		
		// Fallback: treat as a generic struct instantiation if defined
		if sym, ok := a.currentScope.Resolve(baseName); ok {
			return sym.Type
		}
		
		return types.Void
	}

	// 5. Generic/Qualified Types
	if tc.IDENTIFIER() != nil || tc.QualifiedType() != nil {
		var name string
		// var genericArgs parser.IGenericArgsContext // unused for now

		if tc.QualifiedType() != nil {
			qt := tc.QualifiedType()
			for i, id := range qt.AllIDENTIFIER() {
				if i > 0 { name += "." }
				name += id.GetText()
			}
			// genericArgs = qt.GenericArgs()
		} else {
			name = tc.IDENTIFIER().GetText()
			// genericArgs = tc.GenericArgs()
		}

		// Resolve base symbol
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
	// Pointer compatibility
	if srcPtr, sOk := src.(*types.PointerType); sOk {
		if destPtr, dOk := dest.(*types.PointerType); dOk {
			// void* compatibility
			if srcPtr.ElementType == types.Void || destPtr.ElementType == types.Void {
				return true
			}
			// String/Byte pointer compatibility (*i8 <-> *u8)
			if types.IsInteger(srcPtr.ElementType) && types.IsInteger(destPtr.ElementType) {
				if srcPtr.ElementType.BitSize() == 8 && destPtr.ElementType.BitSize() == 8 {
					return true
				}
			}
		}
	}
	// Array/Slice compatibility
	if srcArr, sOk := src.(*types.ArrayType); sOk {
		if destArr, dOk := dest.(*types.ArrayType); dOk {
			if srcArr.Length == destArr.Length {
				return areTypesCompatible(srcArr.ElementType, destArr.ElementType)
			}
		}
	}
	
	return false
}