package irgen

import (
	"log"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
    "github.com/arc-language/arc-lang/codegen/cpu/amd64"
)

func (g *Generator) resolveType(ctx parser.ITypeContext) types.Type {
	if ctx == nil {
		return types.Void
	}
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
	
	// Replaced PointerType with RawPointer check
	if tc.RAWPTR() != nil {
		return types.NewPointer(types.Void)
	}
	
	if tc.IDENTIFIER() != nil {
		name := tc.IDENTIFIER().GetText()
		
		// 1. Try resolving symbol directly
		if s, ok := g.currentScope.Resolve(name); ok { 
			// Classes are reference types, so they are inherently pointers in IR.
			if st, ok := s.Type.(*types.StructType); ok && st.IsClass {
				return types.NewPointer(s.Type)
			}
			return s.Type 
		} 
		
		// 2. Fallback: Check current namespace
		if g.currentNamespace != "" {
			if s, ok := g.currentScope.Resolve(g.currentNamespace + "." + name); ok {
				if st, ok := s.Type.(*types.StructType); ok && st.IsClass {
					return types.NewPointer(s.Type)
				}
				return s.Type
			}
		}

		log.Printf("[DEBUG] IRGen: resolveType failed to find symbol '%s'. Defaulting to I64.", name)
	}
	return types.I64
}

func (g *Generator) emitCast(val ir.Value, target types.Type) ir.Value {
	src := val.Type()
	if src.Equal(target) { return val }
	
	// Constant Folding
	if cInt, ok := val.(*ir.ConstantInt); ok {
		if tInt, ok := target.(*types.IntType); ok {
			return g.ctx.Builder.ConstInt(tInt, cInt.Value)
		}
		if tFloat, ok := target.(*types.FloatType); ok {
			return g.ctx.Builder.ConstFloat(tFloat, float64(cInt.Value))
		}
		// Integer constant to pointer.
		if types.IsPointer(target) {
			c := &ir.ConstantInt{Value: cInt.Value}
			c.SetType(target)
			return c
		}
		// Handle 0 initialization for Aggregates (Arrays/Structs)
		if types.IsAggregate(target) && cInt.Value == 0 {
			return g.ctx.Builder.ConstZero(target)
		}
	}
	
	if cFloat, ok := val.(*ir.ConstantFloat); ok {
		if tInt, ok := target.(*types.IntType); ok {
			return g.ctx.Builder.ConstInt(tInt, int64(cFloat.Value))
		}
		if tFloat, ok := target.(*types.FloatType); ok {
			return g.ctx.Builder.ConstFloat(tFloat, cFloat.Value)
		}
	}

	// Runtime Casting
	if types.IsInteger(src) && types.IsInteger(target) {
		if src.BitSize() > target.BitSize() { return g.ctx.Builder.CreateTrunc(val, target, "") }
		if src.Equal(types.U8) || src.Equal(types.U16) || src.Equal(types.U32) || src.Equal(types.U64) {
			return g.ctx.Builder.CreateZExt(val, target, "")
		}
		return g.ctx.Builder.CreateSExt(val, target, "")
	}
	if types.IsFloat(src) && types.IsFloat(target) {
		if src.BitSize() > target.BitSize() { return g.ctx.Builder.CreateFPTrunc(val, target, "") }
		return g.ctx.Builder.CreateFPExt(val, target, "")
	}
	if types.IsInteger(src) && types.IsFloat(target) { return g.ctx.Builder.CreateSIToFP(val, target, "") }
	if types.IsFloat(src) && types.IsInteger(target) { return g.ctx.Builder.CreateFPToSI(val, target, "") }

	// Pointer Casting
	if types.IsPointer(src) && types.IsPointer(target) {
		return g.ctx.Builder.CreateBitCast(val, target, "")
	}

	// Array Constant Casting (Recursive) with Padding support
	if srcArr, ok := src.(*types.ArrayType); ok {
		if targetArr, ok := target.(*types.ArrayType); ok {
			if srcArr.Length <= targetArr.Length {
				if cArr, ok := val.(*ir.ConstantArray); ok {
					var newElems []ir.Constant
					
					for _, elem := range cArr.Elements {
						casted := g.emitCast(elem, targetArr.ElementType)
						if cCasted, ok := casted.(ir.Constant); ok {
							newElems = append(newElems, cCasted)
						} else {
							panic("Cast of constant array element resulted in non-constant")
						}
					}
					
					if srcArr.Length < targetArr.Length {
						zero := g.getZeroValue(targetArr.ElementType)
						if zeroConst, ok := zero.(ir.Constant); ok {
							for int64(len(newElems)) < targetArr.Length {
								newElems = append(newElems, zeroConst)
							}
						}
					}
					
					return &ir.ConstantArray{
						BaseValue: ir.BaseValue{ValType: target},
						Elements:  newElems,
					}
				}
				
				if _, ok := val.(*ir.ConstantZero); ok {
					return g.ctx.Builder.ConstZero(target)
				}
			}
		}
	}
	
	return val
}

func (g *Generator) getZeroValue(t types.Type) ir.Value {
	if types.IsInteger(t) { return g.ctx.Builder.ConstInt(t.(*types.IntType), 0) }
	if types.IsFloat(t) { return g.ctx.Builder.ConstFloat(t.(*types.FloatType), 0.0) }
	if types.IsPointer(t) { return g.ctx.Builder.ConstNull(t.(*types.PointerType)) }
	return g.ctx.Builder.ConstZero(t)
}

func needsSret(t types.Type) bool {
    if t == nil || t.Kind() == types.VoidKind {
        return false
    }
    return amd64.SizeOf(t) > 16
}

func (g *Generator) emitCallWithSret(fn *ir.Function, args []ir.Value, name string) ir.Value {
    isSret := fn.FuncType.OriginalReturnType != nil && needsSret(fn.FuncType.OriginalReturnType)

    if isSret {
        retType := fn.FuncType.OriginalReturnType
        alloca := g.ctx.Builder.CreateAlloca(retType, name+".sret")

        sretArgs := append([]ir.Value{alloca}, args...)
        g.ctx.Builder.CreateCall(fn, sretArgs, "")

        return g.ctx.Builder.CreateLoad(retType, alloca, name+".result")
    }

    return g.ctx.Builder.CreateCall(fn, args, name)
}