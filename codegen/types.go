package codegen

import (
	"github.com/arc-language/arc-lang/ast"
	"github.com/arc-language/arc-lang/builder/types"
)

type TypeGenerator struct {
	// Cache could go here
}

func NewTypeGenerator() *TypeGenerator {
	return &TypeGenerator{}
}

func (tg *TypeGenerator) GenType(t ast.TypeRef) types.Type {
	if t == nil {
		return types.Void
	}

	switch typ := t.(type) {
	case *ast.NamedType:
		switch typ.Name {
		case "int8", "byte":   return types.I8
		case "int16":  return types.I16
		case "int32":  return types.I32
		case "int64", "usize", "isize": return types.I64
		case "float32": return types.F32
		case "float64": return types.F64
		case "bool":   return types.I1
		case "void":   return types.Void
		case "char":   return types.U32
		default:
			// In a real compiler, lookup struct/typedef name in module
			return types.Void 
		}

	case *ast.PointerType:
		elem := tg.GenType(typ.Base)
		return types.NewPointer(elem)

	case *ast.ArrayType:
		elem := tg.GenType(typ.Elem)
		
		if typ.Kind == "array" {
			// [N]T -> ArrayType
			// We assume length is resolved or handle generic Expr later
			return types.NewArray(elem, 0) 
		} else if typ.Kind == "slice" {
			// []T -> { *T, i64 }
			return types.NewStruct("slice", []types.Type{
				types.NewPointer(elem),
				types.I64,
			}, false)
		} else if typ.Kind == "vector" {
			// vector[T] -> { *T, i64, i64 }
			// Note: No longer "Class", just a Struct. 
			// Usage via 'var' implies it will be passed by pointer.
			return types.NewStruct("vector", []types.Type{
				types.NewPointer(elem),
				types.I64,
				types.I64,
			}, false)
		}

	case *ast.FuncType:
		// Map parameters
		var params []types.Type
		for _, p := range typ.Params {
			params = append(params, tg.GenType(p))
		}
		
		// Map return
		ret := types.Type(types.Void)
		if len(typ.Results) > 0 {
			ret = tg.GenType(typ.Results[0])
		}
		
		return types.NewFunction(ret, params, false)
	}

	return types.Void
}