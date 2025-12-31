package irgen

import (
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

// resolveType converts AST type to IR type
func (g *Generator) resolveType(ctx parser.ITypeContext) types.Type {
	tc := ctx.(*parser.TypeContext)
	if tc.PrimitiveType() != nil {
		switch tc.PrimitiveType().GetText() {
		case "int": return types.I64
		case "int32": return types.I32
		case "float": return types.F64
		case "bool": return types.I1
		case "void": return types.Void
		case "string": return types.NewPointer(types.I8)
		}
	}
	if tc.PointerType() != nil {
		return types.NewPointer(g.resolveType(tc.PointerType().Type_()))
	}
	if tc.ArrayType() != nil {
		// Basic array support
		elem := g.resolveType(tc.ArrayType().Type_())
		// Note: Actual length parsing would happen here if we supported fixed-size arrays in IR directly
		// For now returning a pointer or slice equivalent is common, but let's stick to fixed
		return types.NewPointer(elem) 
	}
	if tc.IDENTIFIER() != nil {
		// Lookup struct types created in Pass 1
		name := tc.IDENTIFIER().GetText()
		if s, ok := g.currentScope.Resolve(name); ok { 
			return s.Type 
		}
	}
	return types.I64
}

// emitCast handles safe casting between types
func (g *Generator) emitCast(val ir.Value, target types.Type) ir.Value {
	src := val.Type()
	
	if src.Equal(target) {
		return val
	}
	
	// Integer casts
	if types.IsInteger(src) && types.IsInteger(target) {
		if src.BitSize() > target.BitSize() {
			return g.ctx.Builder.CreateTrunc(val, target, "")
		} else if src.BitSize() < target.BitSize() {
			// Defaulting to Sign Extension for safety
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

// getZeroValue returns a constant zero for a given type
func (g *Generator) getZeroValue(t types.Type) ir.Value {
	if types.IsInteger(t) {
		return g.ctx.Builder.ConstInt(t.(*types.IntType), 0)
	}
	if types.IsFloat(t) {
		return g.ctx.Builder.ConstFloat(t.(*types.FloatType), 0.0)
	}
	if types.IsPointer(t) {
		return g.ctx.Builder.ConstNull(t.(*types.PointerType))
	}
	return g.ctx.Builder.ConstZero(t)
}