package semantics

import (
	"strconv"

	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/pkg/symbol"
)

// resolveType converts an AST Type Context into an internal Type definition
func (a *Analyzer) resolveType(ctx parser.ITypeContext) types.Type {
	tc := ctx.(*parser.TypeContext)

	// 1. Primitives
	if tc.PrimitiveType() != nil {
		name := tc.PrimitiveType().GetText()
		if sym, ok := a.currentScope.Resolve(name); ok && sym.Kind == symbol.SymType {
			return sym.Type
		}
		a.bag.Report(a.file, tc.GetStart().GetLine(), 0, "Unknown primitive type '%s'", name)
		return types.Void
	}

	// 2. Pointers
	if tc.PointerType() != nil {
		inner := a.resolveType(tc.PointerType().Type_())
		return types.NewPointer(inner)
	}

	// 3. Arrays
	if tc.ArrayType() != nil {
		inner := a.resolveType(tc.ArrayType().Type_())
		size := int64(0)
		if tc.ArrayType().ArraySize() != nil && tc.ArrayType().ArraySize().INTEGER_LITERAL() != nil {
			txt := tc.ArrayType().ArraySize().INTEGER_LITERAL().GetText()
			val, _ := strconv.ParseInt(txt, 10, 64)
			size = val
		}
		return types.NewArray(inner, size)
	}

	// 4. Named Types (Structs, Classes, Typedefs)
	if tc.IDENTIFIER() != nil {
		name := tc.IDENTIFIER().GetText()
		if sym, ok := a.currentScope.Resolve(name); ok {
			if sym.Kind != symbol.SymType {
				a.bag.Report(a.file, tc.GetStart().GetLine(), 0, "'%s' is not a type", name)
				return types.Void
			}
			return sym.Type
		}
		a.bag.Report(a.file, tc.GetStart().GetLine(), 0, "Unknown type '%s'", name)
	}

	return types.I64 // Fallback
}

// areTypesCompatible checks if src can be assigned to dest
func areTypesCompatible(src, dest types.Type) bool {
	// 1. Exact Match
	if src.Equal(dest) {
		return true
	}

	// 2. Implicit Integer Casting (e.g. i32 -> i64)
	if types.IsInteger(src) && types.IsInteger(dest) {
		if src.BitSize() <= dest.BitSize() {
			return true
		}
	}
	
	// 3. Void handling (always incompatible unless both void)
	if src == types.Void || dest == types.Void {
		return false
	}

	return false
}