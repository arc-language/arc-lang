package irgen

import (
	"fmt"
	"strconv"

	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol" // Added missing import
)

func (g *Generator) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return g.Visit(ctx.LogicalOrExpression())
}

// --- Binary Expressions ---

func (g *Generator) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	lhs := g.Visit(ctx.LogicalAndExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllLogicalAndExpression()); i++ {
		rhs := g.Visit(ctx.LogicalAndExpression(i)).(ir.Value)
		lhs = g.ctx.Builder.CreateOr(lhs, rhs, "")
	}
	return lhs
}

func (g *Generator) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	lhs := g.Visit(ctx.BitOrExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllBitOrExpression()); i++ {
		rhs := g.Visit(ctx.BitOrExpression(i)).(ir.Value)
		lhs = g.ctx.Builder.CreateAnd(lhs, rhs, "")
	}
	return lhs
}

func (g *Generator) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	lhs := g.Visit(ctx.MultiplicativeExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		rhs := g.Visit(ctx.MultiplicativeExpression(i)).(ir.Value)
		isAdd := i <= len(ctx.AllPLUS()) 
		
		if types.IsFloat(lhs.Type()) {
			if isAdd { lhs = g.ctx.Builder.CreateFAdd(lhs, rhs, "") } else { lhs = g.ctx.Builder.CreateFSub(lhs, rhs, "") }
		} else {
			if isAdd { lhs = g.ctx.Builder.CreateAdd(lhs, rhs, "") } else { lhs = g.ctx.Builder.CreateSub(lhs, rhs, "") }
		}
	}
	return lhs
}

func (g *Generator) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	lhs := g.Visit(ctx.UnaryExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllUnaryExpression()); i++ {
		rhs := g.Visit(ctx.UnaryExpression(i)).(ir.Value)
		if types.IsFloat(lhs.Type()) {
			lhs = g.ctx.Builder.CreateFMul(lhs, rhs, "")
		} else {
			lhs = g.ctx.Builder.CreateMul(lhs, rhs, "")
		}
	}
	return lhs
}

// --- Unary Expressions ---

func (g *Generator) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if ctx.PostfixExpression() != nil { return g.Visit(ctx.PostfixExpression()) }
	
	val := g.Visit(ctx.UnaryExpression()).(ir.Value)
	
	if ctx.MINUS() != nil {
		if types.IsFloat(val.Type()) { return g.ctx.Builder.CreateFSub(g.getZeroValue(val.Type()), val, "") }
		return g.ctx.Builder.CreateSub(g.getZeroValue(val.Type()), val, "")
	}
	if ctx.NOT() != nil {
		return g.ctx.Builder.CreateXor(val, g.ctx.Builder.ConstInt(types.I1, 1), "")
	}
	if ctx.BIT_NOT() != nil {
		return g.ctx.Builder.CreateXor(val, g.ctx.Builder.ConstInt(val.Type().(*types.IntType), -1), "")
	}
	if ctx.STAR() != nil { // Dereference (*ptr)
		if ptr, ok := val.Type().(*types.PointerType); ok {
			return g.ctx.Builder.CreateLoad(ptr.ElementType, val, "")
		}
	}
	if ctx.AMP() != nil { // Address Of (&var)
		// Returning 0 for now as placeholder for L-Value logic
		return g.getZeroValue(types.I64) 
	}
	return val
}

// --- Postfix Expressions ---

func (g *Generator) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	curr := g.Visit(ctx.PrimaryExpression()).(ir.Value)

	for _, op := range ctx.AllPostfixOp() {
		// 1. Function Call
		if op.LPAREN() != nil {
			var args []ir.Value
			if op.ArgumentList() != nil {
				for _, arg := range op.ArgumentList().AllArgument() {
					if argExpr := g.Visit(arg.Expression()); argExpr != nil {
						args = append(args, argExpr.(ir.Value))
					}
				}
			}
			if fn, ok := curr.(*ir.Function); ok {
				curr = g.ctx.Builder.CreateCall(fn, args, "")
			}
		}
		
		// 2. Member Access (.field)
		if op.DOT() != nil {
			fieldName := op.IDENTIFIER().GetText()
			
			var structType *types.StructType
			var basePtr ir.Value
			
			if ptr, ok := curr.Type().(*types.PointerType); ok {
				if st, ok := ptr.ElementType.(*types.StructType); ok {
					structType = st
					basePtr = curr
				}
			} else if st, ok := curr.Type().(*types.StructType); ok {
				structType = st
				basePtr = nil
			}

			if structType != nil {
				if idx, ok := g.analysis.StructIndices[structType.Name][fieldName]; ok {
					if basePtr != nil {
						gep := g.ctx.Builder.CreateStructGEP(structType, basePtr, idx, "")
						curr = g.ctx.Builder.CreateLoad(structType.Fields[idx], gep, "")
					} else {
						curr = g.ctx.Builder.CreateExtractValue(curr, []int{idx}, "")
					}
				}
			}
		}
		
		// 3. Array Indexing ([i])
		if op.LBRACKET() != nil {
			idx := g.Visit(op.Expression()).(ir.Value)
			
			if ptr, ok := curr.Type().(*types.PointerType); ok {
				gep := g.ctx.Builder.CreateInBoundsGEP(ptr.ElementType, curr, []ir.Value{idx}, "")
				curr = g.ctx.Builder.CreateLoad(ptr.ElementType, gep, "")
			}
		}
	}
	return curr
}

// --- Primary Expressions ---

func (g *Generator) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.StructLiteral() != nil { return g.Visit(ctx.StructLiteral()) }
	if ctx.Literal() != nil { return g.Visit(ctx.Literal()) }
	if ctx.CastExpression() != nil { return g.Visit(ctx.CastExpression()) }
	if ctx.SyscallExpression() != nil { return g.Visit(ctx.SyscallExpression()) }
	if ctx.IntrinsicExpression() != nil { return g.Visit(ctx.IntrinsicExpression()) }
	
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		
		if sym, ok := g.currentScope.Resolve(name); ok {
			if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
				return g.ctx.Builder.CreateLoad(sym.Type, alloca, "")
			}
			return sym.IRValue
		}
		
		if glob := g.ctx.Module.GetGlobal(name); glob != nil {
			return g.ctx.Builder.CreateLoad(glob.Type().(*types.PointerType).ElementType, glob, "")
		}
	}
	
	if ctx.Expression() != nil { return g.Visit(ctx.Expression()) }
	
	return g.getZeroValue(types.I64)
}

func (g *Generator) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTEGER_LITERAL() != nil {
		txt := ctx.INTEGER_LITERAL().GetText()
		val, _ := strconv.ParseInt(txt, 0, 64)
		return g.ctx.Builder.ConstInt(types.I64, val)
	}
	if ctx.FLOAT_LITERAL() != nil {
		txt := ctx.FLOAT_LITERAL().GetText()
		val, _ := strconv.ParseFloat(txt, 64)
		return g.ctx.Builder.ConstFloat(types.F64, val)
	}
	if ctx.STRING_LITERAL() != nil {
		raw := ctx.STRING_LITERAL().GetText()
		if len(raw) >= 2 { raw = raw[1 : len(raw)-1] }
		content := raw + "\x00"
		
		strName := fmt.Sprintf(".str.%d", len(g.ctx.Module.Globals))
		arrType := types.NewArray(types.I8, int64(len(content)))
		var chars []ir.Constant
		for _, b := range []byte(content) {
			chars = append(chars, g.ctx.Builder.ConstInt(types.I8, int64(b)))
		}
		constArr := &ir.ConstantArray{BaseValue: ir.BaseValue{ValType: arrType}, Elements: chars}
		
		global := g.ctx.Builder.CreateGlobalConstant(strName, constArr)
		zero := g.ctx.Builder.ConstInt(types.I32, 0)
		return g.ctx.Builder.CreateInBoundsGEP(arrType, global, []ir.Value{zero, zero}, "")
	}
	return g.getZeroValue(types.I64)
}

func (g *Generator) VisitStructLiteral(ctx *parser.StructLiteralContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	
	sym, ok := g.currentScope.Resolve(name)
	if !ok || sym.Kind != symbol.SymType {
		return g.getZeroValue(types.I64)
	}
	
	structType := sym.Type.(*types.StructType)
	var agg ir.Value = g.ctx.Builder.ConstZero(structType)
	
	indices := g.analysis.StructIndices[structType.Name]
	for _, field := range ctx.AllFieldInit() {
		fName := field.IDENTIFIER().GetText()
		fVal := g.Visit(field.Expression()).(ir.Value)
		
		if idx, ok := indices[fName]; ok {
			fVal = g.emitCast(fVal, structType.Fields[idx])
			agg = g.ctx.Builder.CreateInsertValue(agg, fVal, []int{idx}, "")
		}
	}
	return agg
}

func (g *Generator) VisitCastExpression(ctx *parser.CastExpressionContext) interface{} {
	val := g.Visit(ctx.Expression()).(ir.Value)
	target := g.resolveType(ctx.Type_())
	return g.emitCast(val, target)
}

func (g *Generator) VisitSyscallExpression(ctx *parser.SyscallExpressionContext) interface{} {
	var args []ir.Value
	for _, expr := range ctx.AllExpression() {
		args = append(args, g.Visit(expr).(ir.Value))
	}
	return g.ctx.Builder.CreateSyscall(args)
}

func (g *Generator) VisitIntrinsicExpression(ctx *parser.IntrinsicExpressionContext) interface{} {
	if ctx.SIZEOF() != nil {
		t := g.resolveType(ctx.Type_())
		return g.ctx.Builder.CreateSizeOf(t, "")
	}
	return g.getZeroValue(types.I64)
}

// Boilerplate traversals
func (g *Generator) VisitBitOrExpression(ctx *parser.BitOrExpressionContext) interface{} { return g.Visit(ctx.BitXorExpression(0)) }
func (g *Generator) VisitBitXorExpression(ctx *parser.BitXorExpressionContext) interface{} { return g.Visit(ctx.BitAndExpression(0)) }
func (g *Generator) VisitBitAndExpression(ctx *parser.BitAndExpressionContext) interface{} { return g.Visit(ctx.EqualityExpression(0)) }
func (g *Generator) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} { return g.Visit(ctx.RelationalExpression(0)) }
func (g *Generator) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} { return g.Visit(ctx.ShiftExpression(0)) }
func (g *Generator) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} { return g.Visit(ctx.RangeExpression(0)) }
func (g *Generator) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} { return g.Visit(ctx.AdditiveExpression(0)) }