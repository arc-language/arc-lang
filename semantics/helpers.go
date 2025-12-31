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

	if tc.PrimitiveType() != nil {
		name := tc.PrimitiveType().GetText()
		if s, ok := a.currentScope.Resolve(name); ok && s.Kind == symbol.SymType {
			return s.Type
		}
		// Default to I64 if unknown to prevent crashes, but log error
		a.bag.Report(a.file, tc.GetStart().GetLine(), 0, "Unknown primitive type '%s'", name)
		return types.I64
	}
	
	if tc.PointerType() != nil {
		return types.NewPointer(a.resolveType(tc.PointerType().Type_()))
	}

	if tc.ArrayType() != nil {
		elem := a.resolveType(tc.ArrayType().Type_())
		size := int64(0)
		if s := tc.ArrayType().ArraySize(); s != nil {
			if s.INTEGER_LITERAL() != nil {
				size, _ = strconv.ParseInt(s.INTEGER_LITERAL().GetText(), 0, 64)
			}
		}
		return types.NewArray(elem, size)
	}

	if tc.IDENTIFIER() != nil {
		name := tc.IDENTIFIER().GetText()
		if s, ok := a.currentScope.Resolve(name); ok && s.Kind == symbol.SymType {
			return s.Type
		}
		a.bag.Report(a.file, tc.GetStart().GetLine(), 0, "Unknown type '%s'", name)
	}
	
	return types.I64
}

// areTypesCompatible checks if two types can be used together (e.g. assignment, math)
func areTypesCompatible(src, dest types.Type) bool {
	if src.Equal(dest) {
		return true
	}
	// Implicit int casting (e.g. int32 -> int64)
	if types.IsInteger(src) && types.IsInteger(dest) {
		return true 
	}
	// Implicit float casting
	if types.IsFloat(src) && types.IsFloat(dest) {
		return true
	}
	
	return false
}