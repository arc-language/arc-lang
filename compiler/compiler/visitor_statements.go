package compiler

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

func (v *IRVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	if ctx.VariableDecl() != nil { return v.Visit(ctx.VariableDecl()) }
	if ctx.ConstDecl() != nil { return v.Visit(ctx.ConstDecl()) }
	if ctx.AssignmentStmt() != nil { return v.Visit(ctx.AssignmentStmt()) }
	if ctx.ReturnStmt() != nil { return v.Visit(ctx.ReturnStmt()) }
	if ctx.IfStmt() != nil { return v.Visit(ctx.IfStmt()) }
	if ctx.ForStmt() != nil { return v.Visit(ctx.ForStmt()) }
	if ctx.SwitchStmt() != nil { return v.Visit(ctx.SwitchStmt()) }
	if ctx.TryStmt() != nil { return v.Visit(ctx.TryStmt()) }
	if ctx.BreakStmt() != nil { return v.Visit(ctx.BreakStmt()) }
	if ctx.ContinueStmt() != nil { return v.Visit(ctx.ContinueStmt()) }
	if ctx.DeferStmt() != nil { return v.Visit(ctx.DeferStmt()) }
	if ctx.ExpressionStmt() != nil { return v.Visit(ctx.ExpressionStmt()) }
	if ctx.Block() != nil { return v.Visit(ctx.Block()) }
	return nil
}

func (v *IRVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	stmts := ctx.AllStatement()
	v.ctx.PushScope()
	
	for i, stmt := range stmts {
		v.Visit(stmt)
		if v.ctx.currentBlock != nil && v.ctx.currentBlock.Terminator() != nil {
			v.logger.Debug("Hit terminator at statement %d in block, stopping", i)
			break
		}
	}
	
	v.ctx.PopScope()
	return nil
}

func (v *IRVisitor) VisitAssignmentStmt(ctx *parser.AssignmentStmtContext) interface{} {
	lhsCtx := ctx.LeftHandSide()
	var destPtr ir.Value
	
	// 1. Resolve Destination Pointer (L-Value)
	
	// Case A: Simple Variable (x = ...)
	if lhsCtx.IDENTIFIER() != nil && lhsCtx.DOT() == nil && lhsCtx.STAR() == nil && lhsCtx.LBRACKET() == nil {
		name := lhsCtx.IDENTIFIER().GetText()
		sym, ok := v.ctx.currentScope.Lookup(name)
		if !ok {
			v.ctx.Logger.Error("Undefined variable: %s", name)
			return nil
		}
		if sym.IsConst {
			v.ctx.Logger.Error("Cannot assign to constant '%s'", name)
			return nil
		}
		if ptr, isAlloca := sym.Value.(*ir.AllocaInst); isAlloca {
			destPtr = ptr
		} else {
			// Should not happen for mutable vars
			v.ctx.Logger.Error("Cannot assign to non-lvalue '%s'", name)
			return nil
		}
	} else if lhsCtx.STAR() != nil {
		// Case B: Pointer Dereference (*ptr = ...)
		destPtr = v.Visit(lhsCtx.PostfixExpression()).(ir.Value)
	} else if lhsCtx.LBRACKET() != nil && lhsCtx.PostfixExpression() != nil {
		// Case C: Indexing (arr[i] = ...)
		base := v.Visit(lhsCtx.PostfixExpression()).(ir.Value)
		indexExpr := v.Visit(lhsCtx.Expression()).(ir.Value)
		
		if ptrType, ok := base.Type().(*types.PointerType); ok {
			destPtr = v.ctx.Builder.CreateInBoundsGEP(ptrType.ElementType, base, []ir.Value{indexExpr}, "")
		} else {
			v.ctx.Logger.Error("Cannot index non-pointer type: %v", base.Type())
			return nil
		}
	} else if lhsCtx.DOT() != nil && lhsCtx.PostfixExpression() != nil {
		// Case D: Field Access (obj.field = ...)
		postfixCtx := lhsCtx.PostfixExpression()
		var basePtr ir.Value

		// Check if simple identifier to look up variable directly
		if postfixCtx.PrimaryExpression() != nil && postfixCtx.PrimaryExpression().IDENTIFIER() != nil {
			varName := postfixCtx.PrimaryExpression().IDENTIFIER().GetText()
			if sym, ok := v.ctx.currentScope.Lookup(varName); ok {
				if alloca, isAlloca := sym.Value.(*ir.AllocaInst); isAlloca {
					if _, isPtr := alloca.AllocatedType.(*types.PointerType); isPtr {
						basePtr = v.ctx.Builder.CreateLoad(alloca.AllocatedType, alloca, "")
					} else {
						basePtr = alloca
					}
				}
			}
		}
		
		if basePtr == nil {
			basePtr = v.Visit(postfixCtx).(ir.Value)
		}
		
		fieldName := lhsCtx.IDENTIFIER().GetText()
		
		if ptrType, ok := basePtr.Type().(*types.PointerType); ok {
			if structType, ok := ptrType.ElementType.(*types.StructType); ok {
				
				var fieldIdx int = -1
				if v.ctx.IsClassType(structType.Name) {
					if idx, ok := v.ctx.ClassFieldIndices[structType.Name][fieldName]; ok {
						fieldIdx = idx
					}
				} else {
					fieldIdx = v.findFieldIndex(structType, fieldName)
				}
				
				if fieldIdx >= 0 {
					destPtr = v.ctx.Builder.CreateStructGEP(structType, basePtr, fieldIdx, "")
				} else {
					v.ctx.Logger.Error("Struct/class '%s' has no field '%s'", structType.Name, fieldName)
					return nil
				}
			}
		}
		
		if destPtr == nil {
			v.ctx.Logger.Error("Invalid field assignment target")
			return nil
		}
	} else {
		v.ctx.Logger.Error("Unsupported assignment target")
		return nil
	}

	// 2. Evaluate RHS
	rhs := v.Visit(ctx.Expression()).(ir.Value)
	
	// 3. Handle Compound Assignment Logic
	finalValue := rhs
	
	if ctx.AssignmentOp() != nil {
		opCtx := ctx.AssignmentOp()
		
		// Simple assignment (=)
		if opCtx.ASSIGN() != nil {
			finalValue = rhs
		} else {
			// Compound assignment - load current value first
			ptrType := destPtr.Type().(*types.PointerType)
			currVal := v.ctx.Builder.CreateLoad(ptrType.ElementType, destPtr, "")
			
			if opCtx.PLUS_ASSIGN() != nil {
				finalValue = v.ctx.Builder.CreateAdd(currVal, rhs, "")
			} else if opCtx.MINUS_ASSIGN() != nil {
				finalValue = v.ctx.Builder.CreateSub(currVal, rhs, "")
			} else if opCtx.STAR_ASSIGN() != nil {
				finalValue = v.ctx.Builder.CreateMul(currVal, rhs, "")
			} else if opCtx.SLASH_ASSIGN() != nil {
				if types.IsFloat(currVal.Type()) {
					finalValue = v.ctx.Builder.CreateFDiv(currVal, rhs, "")
				} else if intType, ok := currVal.Type().(*types.IntType); ok && !intType.Signed {
					finalValue = v.ctx.Builder.CreateUDiv(currVal, rhs, "")
				} else {
					finalValue = v.ctx.Builder.CreateSDiv(currVal, rhs, "")
				}
			} else if opCtx.PERCENT_ASSIGN() != nil {
				if intType, ok := currVal.Type().(*types.IntType); ok && !intType.Signed {
					finalValue = v.ctx.Builder.CreateURem(currVal, rhs, "")
				} else {
					finalValue = v.ctx.Builder.CreateSRem(currVal, rhs, "")
				}
			} else if opCtx.BIT_OR_ASSIGN() != nil {
				finalValue = v.ctx.Builder.CreateOr(currVal, rhs, "")
			} else if opCtx.BIT_AND_ASSIGN() != nil {
				finalValue = v.ctx.Builder.CreateAnd(currVal, rhs, "")
			} else if opCtx.BIT_XOR_ASSIGN() != nil {
				finalValue = v.ctx.Builder.CreateXor(currVal, rhs, "")
			} else if opCtx.LT() != nil {
				// Left shift assignment (<<=)
				// Check for double LT
				finalValue = v.ctx.Builder.CreateShl(currVal, rhs, "")
			} else if opCtx.GT() != nil {
				// Right shift assignment (>>=)
				// Check for double GT
				if intType, ok := currVal.Type().(*types.IntType); ok && intType.Signed {
					finalValue = v.ctx.Builder.CreateAShr(currVal, rhs, "")
				} else {
					finalValue = v.ctx.Builder.CreateLShr(currVal, rhs, "")
				}
			}
		}
	}
	
	// 4. Store Result
	v.ctx.Builder.CreateStore(finalValue, destPtr)
	
	return nil
}

func (v *IRVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	v.logger.Debug("Compiling return statement")
	
	// Execute deferred statements in LIFO order
	deferred := v.ctx.GetDeferredStmts()
	for i := len(deferred) - 1; i >= 0; i-- {
		_ = deferred[i]
		// TODO: Actually emit deferred instruction execution
	}
	
	// Get expected return type
	var expectedRetType types.Type
	if v.ctx.currentFunction != nil {
		expectedRetType = v.ctx.currentFunction.FuncType.ReturnType
	}
	
	// Check for explicit tuple expression in return statement
	if ctx.TupleExpression() != nil {
		v.logger.Debug("Processing tuple expression return")
		tupleCtx := ctx.TupleExpression()
		exprs := tupleCtx.AllExpression()
		
		if len(exprs) == 0 {
			v.ctx.Builder.CreateRetVoid()
			return nil
		}
		
		// Evaluate all tuple element expressions
		values := make([]ir.Value, len(exprs))
		for i, expr := range exprs {
			values[i] = v.Visit(expr).(ir.Value)
		}
		
		// Check if function expects a struct/tuple return
		if structType, ok := expectedRetType.(*types.StructType); ok {
			v.logger.Debug("Building tuple return value with %d fields", len(structType.Fields))
			
			// Build the tuple struct
			var tuple ir.Value = v.ctx.Builder.ConstZero(structType)
			
			for i, val := range values {
				if i >= len(structType.Fields) {
					v.ctx.Logger.Error("Too many values in tuple return: expected %d, got %d", 
						len(structType.Fields), len(values))
					break
				}
				
				// Cast value to match expected field type if needed
				expectedFieldType := structType.Fields[i]
				if !val.Type().Equal(expectedFieldType) {
					v.logger.Debug("Casting tuple element %d from %v to %v", i, val.Type(), expectedFieldType)
					val = v.castValue(val, expectedFieldType)
				}
				
				tuple = v.ctx.Builder.CreateInsertValue(tuple, val, []int{i}, "")
			}
			
			v.ctx.Builder.CreateRet(tuple)
			return nil
		}
		
		// If single value and function expects scalar, just return the value
		if len(values) == 1 && expectedRetType != nil && expectedRetType.Kind() != types.VoidKind {
			retVal := values[0]
			if !retVal.Type().Equal(expectedRetType) {
				retVal = v.castValue(retVal, expectedRetType)
			}
			v.ctx.Builder.CreateRet(retVal)
			return nil
		}
		
		v.ctx.Logger.Error("Cannot return tuple when function doesn't expect tuple type")
		v.ctx.Builder.CreateRetVoid()
		return nil
	}
	
	// Handle regular expression return
	if ctx.Expression() != nil {
		retVal := v.Visit(ctx.Expression()).(ir.Value)
		
		// Cast to expected type if needed
		if expectedRetType != nil && !retVal.Type().Equal(expectedRetType) {
			retVal = v.castValue(retVal, expectedRetType)
		}
		
		v.ctx.Builder.CreateRet(retVal)
	} else {
		v.ctx.Builder.CreateRetVoid()
	}
	
	return nil
}

func (v *IRVisitor) VisitExpressionStmt(ctx *parser.ExpressionStmtContext) interface{} {
	v.Visit(ctx.Expression())
	return nil
}

func (v *IRVisitor) VisitDeferStmt(ctx *parser.DeferStmtContext) interface{} {
	// Defer executes a statement at function exit (in LIFO order)
	if ctx.Expression() != nil {
		_ = v.Visit(ctx.Expression())
	} else if ctx.AssignmentStmt() != nil {
		_ = v.Visit(ctx.AssignmentStmt())
	}
	
	v.ctx.Logger.Warning("defer statement is not fully implemented yet")
	return nil
}

// Helpers for token ordering
func (v *IRVisitor) isBefore(ctx antlr.ParserRuleContext, token antlr.TerminalNode) bool {
	if ctx == nil || token == nil { return false }
	return ctx.GetStop().GetTokenIndex() < token.GetSymbol().GetTokenIndex()
}

func (v *IRVisitor) isAfter(ctx antlr.ParserRuleContext, token antlr.TerminalNode) bool {
	if ctx == nil || token == nil { return false }
	return ctx.GetStart().GetTokenIndex() > token.GetSymbol().GetTokenIndex()
}