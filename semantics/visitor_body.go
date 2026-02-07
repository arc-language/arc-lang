package semantics

import (
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
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
    
    // Auto-declare loop iterator if "IN" syntax is used: for x in ...
    for i, id := range ctx.AllIDENTIFIER() {
        if i == 0 && ctx.IN() != nil {
            name := id.GetText()
            // Assume I64 for range iterators for now. 
            // Proper type inference would look at RangeExpression type.
            a.currentScope.Define(name, symbol.SymVar, types.I64)
        }
    }

	for _, expr := range ctx.AllExpression() { a.Visit(expr) }
	for _, assign := range ctx.AllAssignmentStmt() { a.Visit(assign) }
	
	if ctx.Block() != nil {
		// Map block to the loop scope so var 'i' is visible inside
		a.scopes[ctx.Block()] = a.currentScope
		for _, s := range ctx.Block().AllStatement() { a.Visit(s) }
	}
	return nil
}

func (a *Analyzer) VisitDeferStmt(ctx *parser.DeferStmtContext) interface{} {
	return a.Visit(ctx.Statement())
}

// Stubs for remaining statements to ensure interface satisfaction
func (a *Analyzer) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} { return nil }
func (a *Analyzer) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} { return nil }
func (a *Analyzer) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} { return nil }
func (a *Analyzer) VisitAssignmentStmt(ctx *parser.AssignmentStmtContext) interface{} { return nil }