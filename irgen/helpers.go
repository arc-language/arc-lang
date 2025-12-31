package irgen

import (
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

// resolveType converts AST type to IR type
// In a full implementation, this should look up the SymbolTable to find Type Symbols
func (g *Generator) resolveType(ctx parser.ITypeContext) types.Type {
	tc := ctx.(*parser.TypeContext)
	
	if tc.PrimitiveType() != nil {
		txt := tc.PrimitiveType().GetText()
		switch txt {
		case "int": return types.I64
		case "float": return types.F64
		case "bool": return types.I1
		case "void": return types.Void
		}
	}
	// Fallback
	return types.I64
}

// emitCast handles safe casting between types
func (g *Generator) emitCast(val ir.Value, target types.Type) ir.Value {
	src := val.Type()
	
	// Identity cast
	if src.Equal(target) {
		return val
	}
	
	// Integer casts
	if types.IsInteger(src) && types.IsInteger(target) {
		if src.BitSize() > target.BitSize() {
			return g.ctx.Builder.CreateTrunc(val, target, "")
		} else if src.BitSize() < target.BitSize() {
			return g.ctx.Builder.CreateSExt(val, target, "")
		}
	}
	
	// Float casts
	if types.IsFloat(src) && types.IsFloat(target) {
		if src.BitSize() > target.BitSize() {
			return g.ctx.Builder.CreateFPTrunc(val, target, "")
		} else if src.BitSize() < target.BitSize() {
			return g.ctx.Builder.CreateFPExt(val, target, "")
		}
	}
	
	// Int <-> Float
	if types.IsInteger(src) && types.IsFloat(target) {
		return g.ctx.Builder.CreateSIToFP(val, target, "")
	}
	if types.IsFloat(src) && types.IsInteger(target) {
		return g.ctx.Builder.CreateFPToSI(val, target, "")
	}

	return val
}