package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
)

// --- Statements ---

// Critical: We must override VisitStatement to prevent falling back to the BaseVisitor,
// which would lose our Analyzer context (scopes, etc).
func (a *Analyzer) VisitStatement(ctx *parser.StatementContext) interface{} {
	if ctx.VariableDecl() != nil { return a.Visit(ctx.VariableDecl()) }
	if ctx.ConstDecl() != nil { return a.Visit(ctx.ConstDecl()) }
	if ctx.AssignmentStmt() != nil { return a.Visit(ctx.AssignmentStmt()) }
	if ctx.ExpressionStmt() != nil { return a.Visit(ctx.ExpressionStmt()) }
	if ctx.ReturnStmt() != nil { return a.Visit(ctx.ReturnStmt()) }
	if ctx.IfStmt() != nil { return a.Visit(ctx.IfStmt()) }
	if ctx.ForStmt() != nil { return a.Visit(ctx.ForStmt()) }
	if ctx.SwitchStmt() != nil { return a.Visit(ctx.SwitchStmt()) }
	if ctx.TryStmt() != nil { return a.Visit(ctx.TryStmt()) }
	if ctx.ThrowStmt() != nil { return a.Visit(ctx.ThrowStmt()) }
	if ctx.BreakStmt() != nil { return a.Visit(ctx.BreakStmt()) }
	if ctx.ContinueStmt() != nil { return a.Visit(ctx.ContinueStmt()) }
	if ctx.DeferStmt() != nil { return a.Visit(ctx.DeferStmt()) }
	if ctx.Block() != nil { return a.Visit(ctx.Block()) }
	return nil
}

func (a *Analyzer) VisitBlock(ctx *parser.BlockContext) interface{} {
	shouldPush := true
	if _, exists := a.scopes[ctx]; exists { shouldPush = false }
	if shouldPush {
		a.pushScope(ctx)
		defer a.popScope()
	}
	for _, stmt := range ctx.AllStatement() { a.Visit(stmt) }
	return nil
}

func (a *Analyzer) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	if ctx.Expression() != nil {
		exprType := a.Visit(ctx.Expression()).(types.Type)
		
		if a.currentFuncRetType != nil && !areTypesCompatible(exprType, a.currentFuncRetType) {
			a.bag.Report(a.file, ctx.GetStart().GetLine(), 0, 
				"Return type mismatch: expected %s, got %s", 
				a.currentFuncRetType.String(), exprType.String())
		}
	}
	return nil
}

func (a *Analyzer) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	a.Visit(ctx.Expression(0))
	a.Visit(ctx.Block(0))
	if len(ctx.AllBlock()) > 1 { a.Visit(ctx.Block(1)) }
	return nil
}

func (a *Analyzer) VisitExpressionStmt(ctx *parser.ExpressionStmtContext) interface{} {
	a.Visit(ctx.Expression())
	return nil
}

func (a *Analyzer) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	a.pushScope(ctx)
	defer a.popScope()
	
	if ctx.VariableDecl() != nil { a.Visit(ctx.VariableDecl()) }
	for _, expr := range ctx.AllExpression() { a.Visit(expr) }
	for _, assign := range ctx.AllAssignmentStmt() { a.Visit(assign) }
	
	if ctx.Block() != nil {
		// Map block to the loop scope so var 'i' is visible inside
		a.scopes[ctx.Block()] = a.currentScope
		for _, s := range ctx.Block().AllStatement() { a.Visit(s) }
	}
	return nil
}