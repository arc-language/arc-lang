package semantics

import (
	"strconv"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

func (a *Analyzer) resolveType(ctx parser.ITypeContext) types.Type {
	tc := ctx.(*parser.TypeContext)

	if tc.PrimitiveType() != nil {
		name := tc.PrimitiveType().GetText()
		if s, ok := a.currentScope.Resolve(name); ok && s.Kind == symbol.SymType {
			return s.Type
		}
		return types.I64
	}
	
	if tc.PointerType() != nil {
		return types.NewPointer(a.resolveType(tc.PointerType().Type_()))
	}

	if tc.ArrayType() != nil {
		elem := a.resolveType(tc.ArrayType().Type_())
		size := int64(0)
		if s := tc.ArrayType().ArraySize(); s != nil && s.INTEGER_LITERAL() != nil {
			size, _ = strconv.ParseInt(s.INTEGER_LITERAL().GetText(), 10, 64)
		}
		return types.NewArray(elem, size)
	}

	if tc.IDENTIFIER() != nil {
		name := tc.IDENTIFIER().GetText()
		if s, ok := a.currentScope.Resolve(name); ok && s.Kind == symbol.SymType {
			return s.Type
		}
	}
	
	return types.I64
}