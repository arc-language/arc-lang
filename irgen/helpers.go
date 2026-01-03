package irgen

import (
	"strconv"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (g *Generator) resolveType(ctx parser.ITypeContext) types.Type {
	tc := ctx.(*parser.TypeContext)
	if tc.PrimitiveType() != nil {
		text := tc.PrimitiveType().GetText()
		switch text {
		case "int": return types.I64
		case "int8": return types.I8
		case "int16": return types.I16
		case "int32": return types.I32
		case "int64": return types.I64
		case "isize": return types.I64 
		
		case "uint8", "byte": return types.U8
		case "uint16": return types.U16
		case "uint32": return types.U32
		case "uint64": return types.U64
		case "usize": return types.U64 
		
		case "float": return types.F64
		case "float32": return types.F32
		case "float64": return types.F64
		case "bool": return types.I1
		case "void": return types.Void
		case "string": return types.NewPointer(types.I8)
		}
	}
	if tc.PointerType() != nil {
		return types.NewPointer(g.resolveType(tc.PointerType().Type_()))
	}
	
	// Generics / Arrays
	if tc.IDENTIFIER() != nil {
		name := tc.IDENTIFIER().GetText()
		
		if name == "array" && tc.GenericArgs() != nil {
			args := tc.GenericArgs().GenericArgList().AllGenericArg()
			if len(args) == 2 {
				// Type
				var elemType types.Type
				if tCtx := args[0].Type_(); tCtx != nil {
					elemType = g.resolveType(tCtx)
				}
				// Size
				var length int64
				if exprCtx := args[1].Expression(); exprCtx != nil {
					if val, err := strconv.ParseInt(exprCtx.GetText(), 0, 64); err == nil {
						length = val
					}
				}
				if elemType != nil && length > 0 {
					return types.NewArray(elemType, length)
				}
			}
		}

		if s, ok := g.currentScope.Resolve(name); ok { return s.Type }
	}
	return types.I64
}

func (g *Generator) emitCast(val ir.Value, target types.Type) ir.Value {
	src := val.Type()
	if src.Equal(target) { return val }
	
	if types.IsInteger(src) && types.IsInteger(target) {
		if src.BitSize() > target.BitSize() { return g.ctx.Builder.CreateTrunc(val, target, "") }
		return g.ctx.Builder.CreateSExt(val, target, "")
	}
	if types.IsFloat(src) && types.IsFloat(target) {
		if src.BitSize() > target.BitSize() { return g.ctx.Builder.CreateFPTrunc(val, target, "") }
		return g.ctx.Builder.CreateFPExt(val, target, "")
	}
	if types.IsInteger(src) && types.IsFloat(target) { return g.ctx.Builder.CreateSIToFP(val, target, "") }
	if types.IsFloat(src) && types.IsInteger(target) { return g.ctx.Builder.CreateFPToSI(val, target, "") }
	
	return val
}

func (g *Generator) getZeroValue(t types.Type) ir.Value {
	if types.IsInteger(t) { return g.ctx.Builder.ConstInt(t.(*types.IntType), 0) }
	if types.IsFloat(t) { return g.ctx.Builder.ConstFloat(t.(*types.FloatType), 0.0) }
	if types.IsPointer(t) { return g.ctx.Builder.ConstNull(t.(*types.PointerType)) }
	return g.ctx.Builder.ConstZero(t)
}