package compiler

import (
	"strings"

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
	
	if ctx.ASSIGN() == nil { // It's a compound assignment (+=, -=, etc.)
		// Load current value from destination
		ptrType := destPtr.Type().(*types.PointerType)
		currVal := v.ctx.Builder.CreateLoad(ptrType.ElementType, destPtr, "")
		
		if ctx.PLUS_ASSIGN() != nil {
			finalValue = v.ctx.Builder.CreateAdd(currVal, rhs, "")
		} else if ctx.MINUS_ASSIGN() != nil {
			finalValue = v.ctx.Builder.CreateSub(currVal, rhs, "")
		} else if ctx.STAR_ASSIGN() != nil {
			finalValue = v.ctx.Builder.CreateMul(currVal, rhs, "")
		} else if ctx.SLASH_ASSIGN() != nil {
			if types.IsFloat(currVal.Type()) {
				finalValue = v.ctx.Builder.CreateFDiv(currVal, rhs, "")
			} else if intType, ok := currVal.Type().(*types.IntType); ok && !intType.Signed {
				finalValue = v.ctx.Builder.CreateUDiv(currVal, rhs, "")
			} else {
				finalValue = v.ctx.Builder.CreateSDiv(currVal, rhs, "")
			}
		} else if ctx.PERCENT_ASSIGN() != nil {
			if intType, ok := currVal.Type().(*types.IntType); ok && !intType.Signed {
				finalValue = v.ctx.Builder.CreateURem(currVal, rhs, "")
			} else {
				finalValue = v.ctx.Builder.CreateSRem(currVal, rhs, "")
			}
		}
	}
	
	// 4. Store Result
	v.ctx.Builder.CreateStore(finalValue, destPtr)
	
	return nil
}

func (v *IRVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	v.logger.Debug("Compiling return statement")
	
	deferred := v.ctx.GetDeferredStmts()
	for i := len(deferred) - 1; i >= 0; i-- {
		_ = deferred[i]
	}
	
	if ctx.Expression() != nil {
		retVal := v.Visit(ctx.Expression()).(ir.Value)
		if v.ctx.currentFunction != nil {
			expectedType := v.ctx.currentFunction.FuncType.ReturnType
			if !retVal.Type().Equal(expectedType) {
				retVal = v.castValue(retVal, expectedType)
			}
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
	if ctx.Expression() != nil {
		_ = v.Visit(ctx.Expression())
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