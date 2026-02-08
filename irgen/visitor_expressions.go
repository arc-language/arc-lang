package irgen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

// Helper: Safely get operator token type
func getOp(ctx antlr.ParserRuleContext, i int) int {
	childIdx := 2*i - 1
	if childIdx >= 0 && childIdx < ctx.GetChildCount() {
		if term, ok := ctx.GetChild(childIdx).(antlr.TerminalNode); ok {
			return term.GetSymbol().GetTokenType()
		}
	}
	return 0
}

func (g *Generator) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return g.Visit(ctx.LogicalOrExpression())
}

func (g *Generator) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	lhs := g.Visit(ctx.LogicalAndExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllLogicalAndExpression()); i++ {
		rhs := g.Visit(ctx.LogicalAndExpression(i)).(ir.Value)
		
		if lConst, ok := lhs.(*ir.ConstantInt); ok {
			if rConst, ok := rhs.(*ir.ConstantInt); ok {
				res := lConst.Value
				if rConst.Value != 0 { res = 1 }
				lhs = g.ctx.Builder.ConstInt(lConst.Type().(*types.IntType), res)
				continue
			}
		}

		lhs = g.ctx.Builder.CreateOr(lhs, rhs, "")
	}
	return lhs
}

func (g *Generator) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	lhs := g.Visit(ctx.BitOrExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllBitOrExpression()); i++ {
		rhs := g.Visit(ctx.BitOrExpression(i)).(ir.Value)
		
		if lConst, ok := lhs.(*ir.ConstantInt); ok {
			if rConst, ok := rhs.(*ir.ConstantInt); ok {
				res := int64(0)
				if lConst.Value != 0 && rConst.Value != 0 { res = 1 }
				lhs = g.ctx.Builder.ConstInt(lConst.Type().(*types.IntType), res)
				continue
			}
		}

		lhs = g.ctx.Builder.CreateAnd(lhs, rhs, "")
	}
	return lhs
}

func (g *Generator) VisitBitOrExpression(ctx *parser.BitOrExpressionContext) interface{} {
	lhs := g.Visit(ctx.BitXorExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllBitXorExpression()); i++ {
		rhs := g.Visit(ctx.BitXorExpression(i)).(ir.Value)
		
		if lConst, ok := lhs.(*ir.ConstantInt); ok {
			if rConst, ok := rhs.(*ir.ConstantInt); ok {
				lhs = g.ctx.Builder.ConstInt(lConst.Type().(*types.IntType), lConst.Value | rConst.Value)
				continue
			}
		}

		lhs = g.ctx.Builder.CreateOr(lhs, rhs, "")
	}
	return lhs
}

func (g *Generator) VisitBitXorExpression(ctx *parser.BitXorExpressionContext) interface{} {
	lhs := g.Visit(ctx.BitAndExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllBitAndExpression()); i++ {
		rhs := g.Visit(ctx.BitAndExpression(i)).(ir.Value)
		
		if lConst, ok := lhs.(*ir.ConstantInt); ok {
			if rConst, ok := rhs.(*ir.ConstantInt); ok {
				lhs = g.ctx.Builder.ConstInt(lConst.Type().(*types.IntType), lConst.Value ^ rConst.Value)
				continue
			}
		}

		lhs = g.ctx.Builder.CreateXor(lhs, rhs, "")
	}
	return lhs
}

func (g *Generator) VisitBitAndExpression(ctx *parser.BitAndExpressionContext) interface{} {
	lhs := g.Visit(ctx.EqualityExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllEqualityExpression()); i++ {
		rhs := g.Visit(ctx.EqualityExpression(i)).(ir.Value)
		
		if lConst, ok := lhs.(*ir.ConstantInt); ok {
			if rConst, ok := rhs.(*ir.ConstantInt); ok {
				lhs = g.ctx.Builder.ConstInt(lConst.Type().(*types.IntType), lConst.Value & rConst.Value)
				continue
			}
		}

		lhs = g.ctx.Builder.CreateAnd(lhs, rhs, "")
	}
	return lhs
}

func (g *Generator) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	lhs := g.Visit(ctx.RelationalExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllRelationalExpression()); i++ {
		rhs := g.Visit(ctx.RelationalExpression(i)).(ir.Value)
		op := getOp(ctx, i)
		
		if lConst, ok := lhs.(*ir.ConstantInt); ok {
			if rConst, ok := rhs.(*ir.ConstantInt); ok {
				res := int64(0)
				if op == parser.ArcParserEQ {
					if lConst.Value == rConst.Value { res = 1 }
				} else {
					if lConst.Value != rConst.Value { res = 1 }
				}
				lhs = g.ctx.Builder.ConstInt(types.I1, res)
				continue
			}
		}

		if op == parser.ArcParserEQ {
			lhs = g.ctx.Builder.CreateICmpEQ(lhs, rhs, "")
		} else {
			lhs = g.ctx.Builder.CreateICmpNE(lhs, rhs, "")
		}

		// SAFETY FIX
		if inst, ok := lhs.(ir.Instruction); ok && inst.Parent() == nil {
			g.ctx.Builder.GetInsertBlock().AddInstruction(inst)
		}
	}
	return lhs
}

func (g *Generator) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} {
	lhs := g.Visit(ctx.ShiftExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllShiftExpression()); i++ {
		rhs := g.Visit(ctx.ShiftExpression(i)).(ir.Value)
		op := getOp(ctx, i)
		switch op {
		case parser.ArcParserLT:
			lhs = g.ctx.Builder.CreateICmpSLT(lhs, rhs, "")
		case parser.ArcParserGT:
			lhs = g.ctx.Builder.CreateICmpSGT(lhs, rhs, "")
		case parser.ArcParserLE:
			lhs = g.ctx.Builder.CreateICmpSLE(lhs, rhs, "")
		case parser.ArcParserGE:
			lhs = g.ctx.Builder.CreateICmpSGE(lhs, rhs, "")
		}
		
		// SAFETY FIX
		if inst, ok := lhs.(ir.Instruction); ok && inst.Parent() == nil {
			g.ctx.Builder.GetInsertBlock().AddInstruction(inst)
		}
	}
	return lhs
}

func (g *Generator) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} {
	lhs := g.Visit(ctx.RangeExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllRangeExpression()); i++ {
		rhs := g.Visit(ctx.RangeExpression(i)).(ir.Value)
		op := getOp(ctx, i)
		
		if lConst, ok := lhs.(*ir.ConstantInt); ok {
			if rConst, ok := rhs.(*ir.ConstantInt); ok {
				if op == parser.ArcParserLT { // <<
					lhs = g.ctx.Builder.ConstInt(lConst.Type().(*types.IntType), lConst.Value << rConst.Value)
				} else { // >>
					lhs = g.ctx.Builder.ConstInt(lConst.Type().(*types.IntType), lConst.Value >> rConst.Value)
				}
				continue
			}
		}

		if op == parser.ArcParserLT {
			lhs = g.ctx.Builder.CreateShl(lhs, rhs, "")
		} else {
			lhs = g.ctx.Builder.CreateAShr(lhs, rhs, "")
		}
	}
	return lhs
}

func (g *Generator) VisitRangeExpression(ctx *parser.RangeExpressionContext) interface{} {
	return g.Visit(ctx.AdditiveExpression(0))
}

func (g *Generator) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	lhs := g.Visit(ctx.MultiplicativeExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		rhs := g.Visit(ctx.MultiplicativeExpression(i)).(ir.Value)
		op := getOp(ctx, i)

		if _, ok := lhs.Type().(*types.PointerType); ok {
			if op == parser.ArcParserPLUS {
				lhs = g.ctx.Builder.CreateInBoundsGEP(lhs.Type().(*types.PointerType).ElementType, lhs, []ir.Value{rhs}, "")
			} else if op == parser.ArcParserMINUS {
				negRhs := g.ctx.Builder.CreateSub(g.getZeroValue(rhs.Type()), rhs, "")
				lhs = g.ctx.Builder.CreateInBoundsGEP(lhs.Type().(*types.PointerType).ElementType, lhs, []ir.Value{negRhs}, "")
			}
			continue
		}

		if lConst, ok := lhs.(*ir.ConstantInt); ok {
			if rConst, ok := rhs.(*ir.ConstantInt); ok {
				if op == parser.ArcParserPLUS {
					lhs = g.ctx.Builder.ConstInt(lConst.Type().(*types.IntType), lConst.Value + rConst.Value)
				} else {
					lhs = g.ctx.Builder.ConstInt(lConst.Type().(*types.IntType), lConst.Value - rConst.Value)
				}
				continue
			}
		}

		if types.IsFloat(lhs.Type()) {
			if op == parser.ArcParserPLUS {
				lhs = g.ctx.Builder.CreateFAdd(lhs, rhs, "")
			} else {
				lhs = g.ctx.Builder.CreateFSub(lhs, rhs, "")
			}
		} else {
			if op == parser.ArcParserPLUS {
				lhs = g.ctx.Builder.CreateAdd(lhs, rhs, "")
			} else {
				lhs = g.ctx.Builder.CreateSub(lhs, rhs, "")
			}
		}
	}
	return lhs
}

func (g *Generator) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	lhs := g.Visit(ctx.UnaryExpression(0)).(ir.Value)
	for i := 1; i < len(ctx.AllUnaryExpression()); i++ {
		rhs := g.Visit(ctx.UnaryExpression(i)).(ir.Value)
		op := getOp(ctx, i)

		if lConst, ok := lhs.(*ir.ConstantInt); ok {
			if rConst, ok := rhs.(*ir.ConstantInt); ok {
				if op == parser.ArcParserSTAR {
					lhs = g.ctx.Builder.ConstInt(lConst.Type().(*types.IntType), lConst.Value * rConst.Value)
				} else if op == parser.ArcParserSLASH {
					lhs = g.ctx.Builder.ConstInt(lConst.Type().(*types.IntType), lConst.Value / rConst.Value)
				} else if op == parser.ArcParserPERCENT {
					lhs = g.ctx.Builder.ConstInt(lConst.Type().(*types.IntType), lConst.Value % rConst.Value)
				}
				continue
			}
		}

		if types.IsFloat(lhs.Type()) {
			if op == parser.ArcParserSTAR {
				lhs = g.ctx.Builder.CreateFMul(lhs, rhs, "")
			} else {
				lhs = g.ctx.Builder.CreateFDiv(lhs, rhs, "")
			}
		} else {
			if op == parser.ArcParserSTAR {
				lhs = g.ctx.Builder.CreateMul(lhs, rhs, "")
			} else if op == parser.ArcParserSLASH {
				lhs = g.ctx.Builder.CreateSDiv(lhs, rhs, "")
			} else if op == parser.ArcParserPERCENT {
				lhs = g.ctx.Builder.CreateSRem(lhs, rhs, "")
			}
		}
	}
	return lhs
}

func (g *Generator) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if ctx.AWAIT() != nil {
		handle := g.Visit(ctx.UnaryExpression()).(ir.Value)
		resultType := g.analysis.NodeTypes[ctx]
		if resultType == nil || resultType == types.Void {
			resultType = types.I32
		}
		return g.ctx.Builder.CreateAwaitTask(handle, resultType, "")
	}

	if ctx.AMP() != nil {
		lval := g.getLValue(ctx.UnaryExpression())
		if lval != nil {
			return lval
		}
		fmt.Println("[IRGen] Error: Cannot take address of non-lvalue")
		return g.getZeroValue(types.I64)
	}

	if ctx.STAR() != nil {
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		if ptrType, ok := val.Type().(*types.PointerType); ok {
			return g.ctx.Builder.CreateLoad(ptrType.ElementType, val, "")
		}
		return val
	}

	if ctx.MINUS() != nil {
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		if c, ok := val.(*ir.ConstantInt); ok {
			return g.ctx.Builder.ConstInt(c.Type().(*types.IntType), -c.Value)
		}
		if types.IsFloat(val.Type()) {
			return g.ctx.Builder.CreateFSub(g.getZeroValue(val.Type()), val, "")
		}
		return g.ctx.Builder.CreateSub(g.getZeroValue(val.Type()), val, "")
	}

	if ctx.NOT() != nil {
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		return g.ctx.Builder.CreateXor(val, g.ctx.Builder.ConstInt(types.I1, 1), "")
	}

	if ctx.BIT_NOT() != nil {
		val := g.Visit(ctx.UnaryExpression()).(ir.Value)
		return g.ctx.Builder.CreateXor(val, g.ctx.Builder.ConstInt(val.Type().(*types.IntType), -1), "")
	}

	if ctx.INCREMENT() != nil || ctx.DECREMENT() != nil {
		ptr := g.getLValue(ctx.UnaryExpression())
		if ptr == nil {
			return g.getZeroValue(types.I64)
		}
		elemType := ptr.Type().(*types.PointerType).ElementType
		curr := g.ctx.Builder.CreateLoad(elemType, ptr, "")
		var one ir.Value
		if intTy, ok := elemType.(*types.IntType); ok {
			one = g.ctx.Builder.ConstInt(intTy, 1)
		} else {
			one = g.ctx.Builder.ConstInt(types.I64, 1)
		}
		var next ir.Value
		if ctx.INCREMENT() != nil {
			next = g.ctx.Builder.CreateAdd(curr, one, "")
		} else {
			next = g.ctx.Builder.CreateSub(curr, one, "")
		}
		g.ctx.Builder.CreateStore(next, ptr)
		return next
	}

	if ctx.PostfixExpression() != nil {
		return g.Visit(ctx.PostfixExpression())
	}

	return g.getZeroValue(types.I64)
}

func (g *Generator) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	var currPtr ir.Value = g.getLValue(ctx.PrimaryExpression())
	var curr ir.Value

	if currPtr != nil {
		if _, isFn := currPtr.(*ir.Function); isFn {
			curr = currPtr
			currPtr = nil
		} else if ptrType, ok := currPtr.Type().(*types.PointerType); ok {
			curr = g.ctx.Builder.CreateLoad(ptrType.ElementType, currPtr, "")
		} else {
			curr = currPtr
		}
	} else {
		res := g.Visit(ctx.PrimaryExpression())
		if res != nil {
			curr = res.(ir.Value)
		}
	}

	var pendingFnType *types.FunctionType

	for _, op := range ctx.AllPostfixOp() {
		
		if op.LPAREN() != nil {
			var args []ir.Value
			if bm, ok := curr.(*BoundMethod); ok {
				curr = bm.Fn
				pendingFnType = bm.Fn.FuncType
				args = append(args, bm.This)
			}
			var targetType *types.FunctionType = pendingFnType
			if targetType == nil {
				if fn, ok := curr.(*ir.Function); ok {
					targetType = fn.FuncType
				} else if curr != nil {
					if ptr, ok := curr.Type().(*types.PointerType); ok {
						if ft, ok := ptr.ElementType.(*types.FunctionType); ok {
							targetType = ft
						}
					}
				}
			}
			if currPtr != nil && targetType != nil && len(targetType.ParamTypes) > 0 {
				if g.checkTypeMatch(currPtr, targetType.ParamTypes[0]) {
					args = append(args, currPtr)
				} else {
					if pt, ok := currPtr.Type().(*types.PointerType); ok && pt.ElementType.Equal(targetType.ParamTypes[0]) {
						loaded := g.ctx.Builder.CreateLoad(pt.ElementType, currPtr, "")
						args = append(args, loaded)
					}
				}
			}
			if op.ArgumentList() != nil {
				for _, arg := range op.ArgumentList().AllArgument() {
					if v := g.Visit(arg.Expression()); v != nil {
						val := v.(ir.Value)
						if targetType != nil {
							targetIdx := len(args)
							if targetIdx < len(targetType.ParamTypes) {
								expected := targetType.ParamTypes[targetIdx]
								val = g.emitCast(val, expected)
							}
						}
						args = append(args, val)
					}
				}
			}
			if curr != nil {
				if targetType != nil && targetType.IsAsync {
					if fn, ok := curr.(*ir.Function); ok {
						curr = g.ctx.Builder.CreateAsyncTask(fn, args, "")
					} else {
						call := g.ctx.Builder.CreateIndirectCall(curr, args, "")
						call.SetType(targetType.ReturnType)
						curr = call
					}
				} else if targetType != nil && targetType.IsProcess {
					if fn, ok := curr.(*ir.Function); ok {
						curr = g.ctx.Builder.CreateProcess(fn, args, "")
					}
				} else {
					if fn, ok := curr.(*ir.Function); ok {
						curr = g.ctx.Builder.CreateCall(fn, args, "")
					} else {
						call := g.ctx.Builder.CreateIndirectCall(curr, args, "")
						if targetType != nil {
							call.SetType(targetType.ReturnType)
						}
						curr = call
					}
				}
				currPtr = nil
				pendingFnType = nil
			}
			continue
		}

		if op.DOT() != nil {
			pendingFnType = nil
			fieldName := op.IDENTIFIER().GetText()
			var basePtr ir.Value = currPtr
			if basePtr == nil && curr != nil && types.IsPointer(curr.Type()) {
				basePtr = curr
			}

			if basePtr != nil {
				ptrType := basePtr.Type().(*types.PointerType)
				if _, isPtrToPtr := ptrType.ElementType.(*types.PointerType); isPtrToPtr {
					basePtr = g.ctx.Builder.CreateLoad(ptrType.ElementType, basePtr, "")
				}
				ptrType = basePtr.Type().(*types.PointerType)

				if st, ok := ptrType.ElementType.(*types.StructType); ok {
					indices, hasIndices := g.analysis.StructIndices[st.Name]
					if !hasIndices && g.currentNamespace != "" {
						indices, hasIndices = g.analysis.StructIndices[g.currentNamespace + "." + st.Name]
					}
					if !hasIndices && g.currentNamespace != "" {
						prefix := g.currentNamespace + "."
						if len(st.Name) > len(prefix) && strings.HasPrefix(st.Name, prefix) {
							indices, hasIndices = g.analysis.StructIndices[st.Name[len(prefix):]]
						}
					}
					if hasIndices {
						if idx, ok := indices[fieldName]; ok {
							physicalIndex := idx
							if st.IsClass {
								physicalIndex = idx + 1
							}
							currPtr = g.ctx.Builder.CreateStructGEP(st, basePtr, physicalIndex, "")
							curr = g.ctx.Builder.CreateLoad(st.Fields[idx], currPtr, "")
							continue
						}
					}
					methodName := st.Name + "_" + fieldName
					methodSym, ok := g.currentScope.Resolve(methodName)
					if !ok && g.currentNamespace != "" {
						methodSym, ok = g.currentScope.Resolve(g.currentNamespace + "." + methodName)
					}
					if ok {
						if ft, ok := methodSym.Type.(*types.FunctionType); ok {
							pendingFnType = ft
						}
						if methodSym.IRValue != nil {
							curr = &BoundMethod{
								Fn:   methodSym.IRValue.(*ir.Function),
								This: basePtr,
							}
							currPtr = basePtr
							continue
						}
					}
				}
			} else if curr != nil {
				if st, ok := curr.Type().(*types.StructType); ok {
					indices, hasIndices := g.analysis.StructIndices[st.Name]
					if !hasIndices && g.currentNamespace != "" {
						indices, hasIndices = g.analysis.StructIndices[g.currentNamespace + "." + st.Name]
					}
					if !hasIndices && g.currentNamespace != "" {
						prefix := g.currentNamespace + "."
						if len(st.Name) > len(prefix) && strings.HasPrefix(st.Name, prefix) {
							indices, hasIndices = g.analysis.StructIndices[st.Name[len(prefix):]]
						}
					}
					if hasIndices {
						if idx, ok := indices[fieldName]; ok {
							physicalIndex := idx
							if st.IsClass {
								physicalIndex = idx + 1
							}
							curr = g.ctx.Builder.CreateExtractValue(curr, []int{physicalIndex}, "")
							currPtr = nil
							continue
						}
					}
				}
			}
			continue
		}

		if op.LBRACKET() != nil {
			idx := g.Visit(op.Expression()).(ir.Value)
			var basePtr ir.Value
			if currPtr != nil {
				if pt, ok := currPtr.Type().(*types.PointerType); ok {
					if _, isArray := pt.ElementType.(*types.ArrayType); isArray {
						basePtr = currPtr
					}
				}
			}
			if basePtr == nil && curr != nil && types.IsPointer(curr.Type()) {
				basePtr = curr
			}
			if basePtr == nil { basePtr = currPtr }

			if basePtr != nil {
				ptrType := basePtr.Type().(*types.PointerType)
				var elemPtr ir.Value
                var loadType types.Type

				if arrType, isArray := ptrType.ElementType.(*types.ArrayType); isArray {
					zero := g.ctx.Builder.ConstInt(types.I32, 0)
					elemPtr = g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, basePtr, []ir.Value{zero, idx}, "")
                    loadType = arrType.ElementType
				} else {
					elemPtr = g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, basePtr, []ir.Value{idx}, "")
                    loadType = ptrType.ElementType
				}
				currPtr = elemPtr
				curr = g.ctx.Builder.CreateLoad(loadType, currPtr, "")
			}
			continue
		}

		if op.INCREMENT() != nil || op.DECREMENT() != nil {
			if currPtr != nil {
				oldVal := curr
				var newVal ir.Value
				if types.IsFloat(oldVal.Type()) {
					one := g.ctx.Builder.ConstFloat(oldVal.Type().(*types.FloatType), 1.0)
					if op.INCREMENT() != nil {
						newVal = g.ctx.Builder.CreateFAdd(oldVal, one, "")
					} else {
						newVal = g.ctx.Builder.CreateFSub(oldVal, one, "")
					}
				} else if ptrType, ok := oldVal.Type().(*types.PointerType); ok {
					var offset int64 = 1
					if op.DECREMENT() != nil {
						offset = -1
					}
					idx := g.ctx.Builder.ConstInt(types.I64, offset)
					newVal = g.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, oldVal, []ir.Value{idx}, "")
				} else {
					var one ir.Value
					if intTy, ok := oldVal.Type().(*types.IntType); ok {
						one = g.ctx.Builder.ConstInt(intTy, 1)
					} else {
						one = g.ctx.Builder.ConstInt(types.I64, 1)
					}
					if op.INCREMENT() != nil {
						newVal = g.ctx.Builder.CreateAdd(oldVal, one, "")
					} else {
						newVal = g.ctx.Builder.CreateSub(oldVal, one, "")
					}
				}
				g.ctx.Builder.CreateStore(newVal, currPtr)
				curr = oldVal
				currPtr = nil
			}
			continue
		}
	}
	return curr
}

func (g *Generator) VisitAnonymousFuncExpression(ctx *parser.AnonymousFuncExpressionContext) interface{} {
	name := fmt.Sprintf("lambda_%d", len(g.ctx.Module.Functions))
	if g.ctx.CurrentFunction != nil {
		name = fmt.Sprintf("%s_lambda_%d", g.ctx.CurrentFunction.Name(), len(g.ctx.Module.Functions))
	}

	var retType types.Type = types.Void
	if ctx.ReturnType() != nil {
		if ctx.ReturnType().Type_() != nil {
			retType = g.resolveType(ctx.ReturnType().Type_())
		}
	}

	var paramTypes []types.Type
	var paramNames []string
	if ctx.ParameterList() != nil {
		for _, param := range ctx.ParameterList().AllParameter() {
			paramTypes = append(paramTypes, g.resolveType(param.Type_()))
			paramNames = append(paramNames, param.IDENTIFIER().GetText())
		}
	}

	fn := g.ctx.Builder.CreateFunction(name, retType, paramTypes, false)

	if es := ctx.ExecutionStrategy(); es != nil {
		if es.ASYNC() != nil {
			fn.FuncType.IsAsync = true
		} else if es.PROCESS() != nil {
			fn.FuncType.IsProcess = true
		}
	}

	prevFunc := g.ctx.CurrentFunction
	prevBlock := g.ctx.Builder.GetInsertBlock()
	
	g.ctx.EnterFunction(fn)
	g.enterScope(ctx)

	entryBlock := g.ctx.Builder.CreateBlock("entry")
	g.ctx.Builder.SetInsertPoint(entryBlock)

	for i, arg := range fn.Arguments {
		arg.SetName(paramNames[i])
		alloca := g.ctx.Builder.CreateAlloca(arg.Type(), paramNames[i]+".addr")
		g.ctx.Builder.CreateStore(arg, alloca)
		if s, ok := g.currentScope.ResolveLocal(paramNames[i]); ok {
			s.IRValue = alloca
		} else {
			g.currentScope.Define(paramNames[i], symbol.SymVar, arg.Type()).IRValue = alloca
		}
	}

	if ctx.Block() != nil {
		outerDefer := g.deferStack
		g.deferStack = NewDeferStack()
		g.Visit(ctx.Block())

		if g.ctx.Builder.GetInsertBlock().Terminator() == nil {
			g.deferStack.Emit(g)
			if retType == types.Void {
				g.ctx.Builder.CreateRetVoid()
			} else {
				g.ctx.Builder.CreateRet(g.getZeroValue(retType))
			}
		}
		g.deferStack = outerDefer
	}

	g.exitScope()
	g.ctx.ExitFunction()
	
	if prevFunc != nil {
		g.ctx.CurrentFunction = prevFunc
		g.ctx.SetInsertBlock(prevBlock)
	}

	return fn
}

func (g *Generator) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.StructLiteral() != nil {
		return g.Visit(ctx.StructLiteral())
	}
	if ctx.AnonymousFuncExpression() != nil {
		return g.Visit(ctx.AnonymousFuncExpression())
	}
	if ctx.Literal() != nil {
		return g.Visit(ctx.Literal())
	}
	if ctx.SizeofExpression() != nil {
		t := g.resolveType(ctx.SizeofExpression().Type_())
		return g.ctx.Builder.CreateSizeOf(t, "")
	}
	if ctx.AlignofExpression() != nil {
		t := g.resolveType(ctx.AlignofExpression().Type_())
		return g.ctx.Builder.CreateAlignOf(t, "")
	}
	if ctx.Expression() != nil && ctx.LPAREN() != nil && ctx.IDENTIFIER() == nil && ctx.QualifiedIdentifier() == nil {
		return g.Visit(ctx.Expression())
	}

	var name string
	var isQualified bool
	var qCtx *parser.QualifiedIdentifierContext

	if ctx.QualifiedIdentifier() != nil {
		qCtx = ctx.QualifiedIdentifier().(*parser.QualifiedIdentifierContext)
		for i, id := range qCtx.AllIDENTIFIER() {
			if i > 0 { name += "." }
			name += id.GetText()
		}
		isQualified = true
	} else if ctx.IDENTIFIER() != nil {
		name = ctx.IDENTIFIER().GetText()
	}

	if name != "" {
		isCall := ctx.LPAREN() != nil
		var args []ir.Value
		if isCall {
			if ctx.ArgumentList() != nil {
				for _, arg := range ctx.ArgumentList().AllArgument() {
					if v := g.Visit(arg.Expression()); v != nil {
						args = append(args, v.(ir.Value))
					}
				}
			}
		}

		var entity ir.Value
		var argsToPass []ir.Value = args
		var pendingFnType *types.FunctionType

		sym, ok := g.currentScope.Resolve(name)
		// Fix: Removed !isQualified check
		if !ok && g.currentNamespace != "" {
			sym, ok = g.currentScope.Resolve(g.currentNamespace + "." + name)
		}

		if ok {
			if ft, ok := sym.Type.(*types.FunctionType); ok {
				pendingFnType = ft
			}
			if sym.IRValue != nil {
				if constant, ok := sym.IRValue.(ir.Constant); ok {
					return constant
				}
				if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
					if !isCall {
						loaded := g.ctx.Builder.CreateLoad(sym.Type, alloca, "")
						
						// SAFETY FIX: Concrete pointer check
						if loaded.Parent() == nil {
							g.ctx.Builder.GetInsertBlock().AddInstruction(loaded)
						}
						return loaded
					}
					
					loadInst := g.ctx.Builder.CreateLoad(sym.Type, alloca, "")
					// SAFETY FIX
					if loadInst.Parent() == nil {
						g.ctx.Builder.GetInsertBlock().AddInstruction(loadInst)
					}
					entity = loadInst
				} else {
					entity = sym.IRValue
				}
			}
		} else if fn := g.ctx.Module.GetFunction(name); fn != nil {
			entity = fn
			pendingFnType = fn.FuncType
		} else if glob := g.ctx.Module.GetGlobal(name); glob != nil {
			if !isCall {
				return g.ctx.Builder.CreateLoad(glob.Type().(*types.PointerType).ElementType, glob, "")
			}
			entity = g.ctx.Builder.CreateLoad(glob.Type().(*types.PointerType).ElementType, glob, "")
		} else if isQualified {
			ids := qCtx.AllIDENTIFIER()
			baseName := ids[0].GetText()
			var basePtr ir.Value
			sym, ok := g.currentScope.Resolve(baseName)
			if !ok && g.currentNamespace != "" {
				sym, ok = g.currentScope.Resolve(g.currentNamespace + "." + baseName)
			}
			if ok {
				if constant, ok := sym.IRValue.(ir.Constant); ok {
					return constant
				}
				if alloca, ok := sym.IRValue.(*ir.AllocaInst); ok {
					basePtr = alloca
				} else {
					basePtr = sym.IRValue
				}
			}

			if basePtr != nil {
				currPtr := basePtr
				if ptrType, ok := currPtr.Type().(*types.PointerType); ok {
					if _, isPtrToPtr := ptrType.ElementType.(*types.PointerType); isPtrToPtr {
						currPtr = g.ctx.Builder.CreateLoad(ptrType.ElementType, currPtr, "")
					}
				}
				valid := true
				for i := 1; i < len(ids); i++ {
					fieldName := ids[i].GetText()
					ptrType, isPtr := currPtr.Type().(*types.PointerType)
					if !isPtr {
						valid = false
						break
					}
					if st, ok := ptrType.ElementType.(*types.StructType); ok {
						indices, hasIndices := g.analysis.StructIndices[st.Name]
						if !hasIndices && g.currentNamespace != "" {
							indices, hasIndices = g.analysis.StructIndices[g.currentNamespace + "." + st.Name]
						}
						if !hasIndices && g.currentNamespace != "" {
							prefix := g.currentNamespace + "."
							if len(st.Name) > len(prefix) && strings.HasPrefix(st.Name, prefix) {
								indices, hasIndices = g.analysis.StructIndices[st.Name[len(prefix):]]
							}
						}
						if hasIndices {
							if idx, ok := indices[fieldName]; ok {
								physicalIndex := idx
								if st.IsClass {
									physicalIndex = idx + 1
								}
								currPtr = g.ctx.Builder.CreateStructGEP(st, currPtr, physicalIndex, "")
								continue
							}
						}
						
						methodName := st.Name + "_" + fieldName
						methodSym, ok := g.currentScope.Resolve(methodName)
						if !ok && g.currentNamespace != "" {
							methodSym, ok = g.currentScope.Resolve(g.currentNamespace + "." + methodName)
						}
						if ok {
							if fn, ok := methodSym.IRValue.(*ir.Function); ok {
								thisArg := currPtr
								ft := fn.FuncType
								if len(ft.ParamTypes) > 0 && !ft.ParamTypes[0].Equal(thisArg.Type()) {
									if pt, isPtr := thisArg.Type().(*types.PointerType); isPtr && pt.ElementType.Equal(ft.ParamTypes[0]) {
										thisArg = g.ctx.Builder.CreateLoad(pt.ElementType, thisArg, "")
									}
								}
								if isCall && i == len(ids)-1 {
									pendingFnType = ft
									argsToPass = append([]ir.Value{thisArg}, argsToPass...)
									entity = methodSym.IRValue
									valid = true
									break
								} else if i == len(ids)-1 {
									return &BoundMethod{Fn: fn, This: thisArg}
								}
							}
						}
					}
					valid = false
					break
				}
				if valid && entity == nil {
					ptrType := currPtr.Type().(*types.PointerType)
					loadInst := g.ctx.Builder.CreateLoad(ptrType.ElementType, currPtr, "")
					
					// SAFETY FIX
					if loadInst.Parent() == nil {
						g.ctx.Builder.GetInsertBlock().AddInstruction(loadInst)
					}
					entity = loadInst
				}
			}
		}

		if entity == nil {
			if isCall && !isQualified {
				var typeArgs []types.Type
				if ctx.GenericArgs() != nil {
					gArgs := ctx.GenericArgs().GenericArgList()
					if gArgs != nil {
						for _, ga := range gArgs.AllGenericArg() {
							if ga.Type_() != nil {
								typeArgs = append(typeArgs, g.resolveType(ga.Type_()))
							}
						}
					}
				}
				if intrinsicVal := g.GenerateIntrinsicCall(name, args, typeArgs); intrinsicVal != nil {
					return intrinsicVal
				}
			}
			if isCall {
				fmt.Printf("[IRGen] Error: Call to undefined function '%s'\n", name)
			} else {
				fmt.Printf("[IRGen] Error: Undefined identifier '%s'\n", name)
			}
			return g.getZeroValue(types.I64)
		}

		if isCall {
			if pendingFnType != nil {
				for i, argVal := range argsToPass {
					if i < len(pendingFnType.ParamTypes) {
						expected := pendingFnType.ParamTypes[i]
						argsToPass[i] = g.emitCast(argVal, expected)
					}
				}
			}
			if pendingFnType != nil && pendingFnType.IsAsync {
				if fn, ok := entity.(*ir.Function); ok {
					return g.ctx.Builder.CreateAsyncTask(fn, argsToPass, "")
				}
				call := g.ctx.Builder.CreateIndirectCall(entity, argsToPass, "")
				call.SetType(pendingFnType.ReturnType)
				return call
			}
			if fn, ok := entity.(*ir.Function); ok {
				return g.ctx.Builder.CreateCall(fn, argsToPass, "")
			} else {
				call := g.ctx.Builder.CreateIndirectCall(entity, argsToPass, "")
				if pendingFnType != nil {
					call.SetType(pendingFnType.ReturnType)
				}
				return call
			}
		}
		return entity
	}
	return g.getZeroValue(types.I64)
}

func processEscapes(s string) string {
	var buf []byte
	i := 0
	for i < len(s) {
		if s[i] == '\\' && i+1 < len(s) {
			switch s[i+1] {
			case '0':
				buf = append(buf, 0)
				i += 2
				continue
			case 'n':
				buf = append(buf, '\n')
				i += 2
				continue
			case 't':
				buf = append(buf, '\t')
				i += 2
				continue
			case 'r':
				buf = append(buf, '\r')
				i += 2
				continue
			case '\\':
				buf = append(buf, '\\')
				i += 2
				continue
			case '"':
				buf = append(buf, '"')
				i += 2
				continue
			case 'x':
				if i+3 < len(s) {
					val, err := strconv.ParseUint(s[i+2:i+4], 16, 8)
					if err == nil {
						buf = append(buf, byte(val))
						i += 4
						continue
					}
				}
			}
		}
		buf = append(buf, s[i])
		i++
	}
	return string(buf)
}

func (g *Generator) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.InitializerList() != nil {
		return g.Visit(ctx.InitializerList())
	}
	txt := ctx.GetText()
	if ctx.NULL() != nil {
		return g.ctx.Builder.ConstNull(types.NewPointer(types.Void))
	}
	if ctx.CHAR_LITERAL() != nil {
		// Fix: Use strconv.Unquote to parse escape sequences (e.g. '\n' -> 10, not 92)
		val, err := strconv.Unquote(txt)
		if err == nil && len(val) > 0 {
			r := []rune(val)[0]
			return g.ctx.Builder.ConstInt(types.I32, int64(r))
		}
		
		// Fallback for simple escapes if Unquote fails (redundancy for safety)
		if len(txt) >= 3 && txt[0] == '\'' && txt[len(txt)-1] == '\'' {
			inner := txt[1 : len(txt)-1]
			if len(inner) == 2 && inner[0] == '\\' {
				switch inner[1] {
				case 'n': return g.ctx.Builder.ConstInt(types.I32, 10)
				case 't': return g.ctx.Builder.ConstInt(types.I32, 9)
				case 'r': return g.ctx.Builder.ConstInt(types.I32, 13)
				case '0': return g.ctx.Builder.ConstInt(types.I32, 0)
				case '\\': return g.ctx.Builder.ConstInt(types.I32, 92)
				case '\'': return g.ctx.Builder.ConstInt(types.I32, 39)
				case '"': return g.ctx.Builder.ConstInt(types.I32, 34)
				}
			}
		}
		return g.getZeroValue(types.I32)
	}
	if ctx.INTEGER_LITERAL() != nil {
		val, _ := strconv.ParseInt(txt, 0, 64)
		return g.ctx.Builder.ConstInt(types.I64, val)
	}
	if ctx.FLOAT_LITERAL() != nil {
		val, _ := strconv.ParseFloat(txt, 64)
		return g.ctx.Builder.ConstFloat(types.F64, val)
	}
	if ctx.BOOLEAN_LITERAL() != nil {
		val := int64(0)
		if txt == "true" {
			val = 1
		}
		return g.ctx.Builder.ConstInt(types.I1, val)
	}
	if ctx.STRING_LITERAL() != nil {
		unquoted, err := strconv.Unquote(txt)
		if err != nil {
			if len(txt) >= 2 {
				unquoted = txt[1 : len(txt)-1]
			}
		}
		// Process escape sequences that strconv.Unquote doesn't handle
		unquoted = processEscapes(unquoted)

		if len(unquoted) == 0 || unquoted[len(unquoted)-1] != 0 {
			unquoted += "\x00"
		}
		content := unquoted
		arrType := types.NewArray(types.I8, int64(len(content)))
		var chars []ir.Constant
		for _, b := range []byte(content) {
			chars = append(chars, g.ctx.Builder.ConstInt(types.I8, int64(b)))
		}
		strName := fmt.Sprintf(".str.%d", len(g.ctx.Module.Globals))
		global := g.ctx.Builder.CreateGlobalConstant(strName, &ir.ConstantArray{BaseValue: ir.BaseValue{ValType: arrType}, Elements: chars})
		zero := g.ctx.Builder.ConstInt(types.I32, 0)
		return g.ctx.Builder.CreateInBoundsGEP(arrType, global, []ir.Value{zero, zero}, "")
	}
	return g.getZeroValue(types.I64)
}

func (g *Generator) VisitInitializerList(ctx *parser.InitializerListContext) interface{} {
	var values []ir.Value
	var elemType types.Type
	for _, expr := range ctx.AllExpression() {
		val := g.Visit(expr).(ir.Value)
		values = append(values, val)
		if elemType == nil {
			elemType = val.Type()
		}
	}
	if len(values) == 0 {
		return g.getZeroValue(types.I64)
	}
	arrType := types.NewArray(elemType, int64(len(values)))
	allConstants := true
	var constElems []ir.Constant
	for _, val := range values {
		if _, ok := val.(ir.Constant); !ok {
			allConstants = false
			break
		}
	}
	if allConstants {
		for _, val := range values {
			casted := g.emitCast(val, elemType)
			if c, ok := casted.(ir.Constant); ok {
				constElems = append(constElems, c)
			} else {
				allConstants = false
				break
			}
		}
	}
	if allConstants {
		return &ir.ConstantArray{
			BaseValue: ir.BaseValue{ValType: arrType},
			Elements:  constElems,
		}
	}
	var agg ir.Value = g.ctx.Builder.ConstZero(arrType)
	for i, val := range values {
		val = g.emitCast(val, elemType)
		agg = g.ctx.Builder.CreateInsertValue(agg, val, []int{i}, "")
	}
	return agg
}

func (g *Generator) VisitStructLiteral(ctx *parser.StructLiteralContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	if ctx.QualifiedIdentifier() != nil {
		qCtx := ctx.QualifiedIdentifier().(*parser.QualifiedIdentifierContext)
		name = ""
		for i, id := range qCtx.AllIDENTIFIER() {
			if i > 0 { name += "." }
			name += id.GetText()
		}
	}
    
	sym, ok := g.currentScope.Resolve(name)
	if !ok && g.currentNamespace != "" && ctx.QualifiedIdentifier() == nil {
		sym, ok = g.currentScope.Resolve(g.currentNamespace + "." + name)
	}

	if !ok { return g.getZeroValue(types.I64) }

	structType := sym.Type.(*types.StructType)
	indices := g.analysis.StructIndices[structType.Name]
	if indices == nil && g.currentNamespace != "" {
		indices = g.analysis.StructIndices[g.currentNamespace + "." + structType.Name]
	}
	if indices == nil && g.currentNamespace != "" {
		prefix := g.currentNamespace + "."
		if len(structType.Name) > len(prefix) && strings.HasPrefix(structType.Name, prefix) {
			indices = g.analysis.StructIndices[structType.Name[len(prefix):]]
		}
	}

	if !structType.IsClass {
		var agg ir.Value = g.ctx.Builder.ConstZero(structType)
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

	size := g.ctx.Builder.CreateSizeOf(structType, "class_size")
	mallocFunc := g.getOrCreateMalloc()
	rawPtr := g.ctx.Builder.CreateCall(mallocFunc, []ir.Value{size}, "raw_ptr")
	classPtr := g.ctx.Builder.CreateBitCast(rawPtr, types.NewPointer(structType), "obj_ptr")

	rcPtr := g.ctx.Builder.CreateStructGEP(structType, classPtr, 0, "rc_ptr")
	g.ctx.Builder.CreateStore(g.ctx.Builder.ConstInt(types.I64, 1), rcPtr)

	for _, field := range ctx.AllFieldInit() {
		fName := field.IDENTIFIER().GetText()
		fVal := g.Visit(field.Expression()).(ir.Value)
		if idx, ok := indices[fName]; ok {
			fVal = g.emitCast(fVal, structType.Fields[idx])
			physicalIndex := idx + 1
			fieldPtr := g.ctx.Builder.CreateStructGEP(structType, classPtr, physicalIndex, fName)
			g.ctx.Builder.CreateStore(fVal, fieldPtr)
		}
	}
	return classPtr
}

func (g *Generator) getOrCreateMalloc() *ir.Function {
	if fn := g.ctx.Module.GetFunction("malloc"); fn != nil {
		return fn
	}
	return g.ctx.Builder.DeclareFunction("malloc", types.NewPointer(types.I8), []types.Type{types.I64}, false)
}

func (g *Generator) checkTypeMatch(val ir.Value, expected types.Type) bool {
	if val == nil || expected == nil {
		return false
	}
	vType := val.Type()
	if vType.Equal(expected) {
		return true
	}
	return false
}

type BoundMethod struct {
	Fn   *ir.Function
	This ir.Value
}
func (b *BoundMethod) Type() types.Type { return b.Fn.Type() }
func (b *BoundMethod) Name() string     { return "bound_method" }
func (b *BoundMethod) SetName(n string) {}
func (b *BoundMethod) String() string   { return "bound_method" }