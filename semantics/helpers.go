package semantics

import (
	"strconv"

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
	
	// 2. Pointers
	if tc.PointerType() != nil {
		return types.NewPointer(a.resolveType(tc.PointerType().Type_()))
	}

	// 3. References
	if tc.ReferenceType() != nil {
		return types.NewPointer(a.resolveType(tc.ReferenceType().Type_()))
	}

	// 4. Function Types
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
		
		return types.NewFunction(retType, paramTypes, ft.ASYNC() != nil)
	}

	// 5. Generic/Qualified Types
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
		
		// Fix: Go-style namespace resolution.
		if !ok && a.currentNamespacePrefix != "" && tc.QualifiedType() == nil {
			s, ok = a.currentScope.Resolve(a.currentNamespacePrefix + "." + name)
		}

		if !ok || s.Kind != symbol.SymType {
			a.bag.Report(a.file, tc.GetStart().GetLine(), 0, "Unknown type '%s'", name)
			return types.I64
		}

		// Handle Generics (array, etc)
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

	// 6. FIXED: Array Types [Size]Type
	// This was missing, causing [2]string to fall through to I64
	if tc.ArrayType() != nil {
		at := tc.ArrayType()
		elemType := a.resolveType(at.Type_())
		
		var length int64 = 0
		if expr := at.Expression(); expr != nil {
			txt := expr.GetText()
			if val, err := strconv.ParseInt(txt, 0, 64); err == nil {
				length = val
			}
		}
		return types.NewArray(elemType, length)
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
			// Allows assigning string literals ("hello", *i8) to *byte (*u8)
			if types.IsInteger(srcPtr.ElementType) && types.IsInteger(destPtr.ElementType) {
				if srcPtr.ElementType.BitSize() == 8 && destPtr.ElementType.BitSize() == 8 {
					return true
				}
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