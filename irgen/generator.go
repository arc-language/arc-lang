package irgen

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/context"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/semantics"
	"github.com/arc-language/arc-lang/symbol"
)

type Generator struct {
	*parser.BaseArcParserVisitor
	ctx              *context.Context
	analysis         *semantics.AnalysisResult
	currentScope     *symbol.Scope
	deferStack       *DeferStack
	loopStack        []loopInfo
	currentNamespace string // NEW: Track extern namespace
}

type loopInfo struct {
	breakBlock    *ir.BasicBlock
	continueBlock *ir.BasicBlock
}

func Generate(tree parser.ICompilationUnitContext, moduleName string, analysis *semantics.AnalysisResult) *ir.Module {
	ctx := context.NewContext(moduleName)
	gen := &Generator{
		BaseArcParserVisitor: &parser.BaseArcParserVisitor{},
		ctx:                  ctx,
		analysis:             analysis,
		currentScope:         analysis.GlobalScope,
		deferStack:           NewDeferStack(),
	}
	
	gen.Visit(tree)
	return ctx.Module
}

func (g *Generator) enterScope(ctx antlr.ParserRuleContext) {
	if s, ok := g.analysis.Scopes[ctx]; ok {
		g.currentScope = s
	}
}

func (g *Generator) exitScope() {
	if g.currentScope.Parent != nil {
		g.currentScope = g.currentScope.Parent
	}
}

// Visit manually dispatches to ensure correct IR generation methods are called
func (g *Generator) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil { return nil }

	switch ctx := tree.(type) {
	case *parser.CompilationUnitContext:
		return g.VisitCompilationUnit(ctx)
	case *parser.TopLevelDeclContext:
		return g.VisitTopLevelDecl(ctx)
	case *parser.FunctionDeclContext:
		return g.VisitFunctionDecl(ctx)
	case *parser.VariableDeclContext:
		return g.VisitVariableDecl(ctx)
	case *parser.ExternDeclContext:
		return g.VisitExternDecl(ctx) // NEW: Dispatch Extern
	case *parser.BlockContext:
		return g.VisitBlock(ctx)
	
	// Statements
	case *parser.StatementContext:
		if ctx.VariableDecl() != nil { return g.VisitVariableDecl(ctx.VariableDecl().(*parser.VariableDeclContext)) }
		if ctx.ReturnStmt() != nil { return g.VisitReturnStmt(ctx.ReturnStmt().(*parser.ReturnStmtContext)) }
		if ctx.IfStmt() != nil { return g.VisitIfStmt(ctx.IfStmt().(*parser.IfStmtContext)) }
		if ctx.ForStmt() != nil { return g.VisitForStmt(ctx.ForStmt().(*parser.ForStmtContext)) }
		if ctx.SwitchStmt() != nil { return g.VisitSwitchStmt(ctx.SwitchStmt().(*parser.SwitchStmtContext)) }
		if ctx.BreakStmt() != nil { return g.VisitBreakStmt(ctx.BreakStmt().(*parser.BreakStmtContext)) }
		if ctx.ContinueStmt() != nil { return g.VisitContinueStmt(ctx.ContinueStmt().(*parser.ContinueStmtContext)) }
		if ctx.DeferStmt() != nil { return g.VisitDeferStmt(ctx.DeferStmt().(*parser.DeferStmtContext)) }
		if ctx.ExpressionStmt() != nil { return g.Visit(ctx.ExpressionStmt().Expression()) }
		if ctx.Block() != nil { return g.VisitBlock(ctx.Block().(*parser.BlockContext)) }
		if ctx.TryStmt() != nil { return g.VisitTryStmt(ctx.TryStmt().(*parser.TryStmtContext)) }
		if ctx.ThrowStmt() != nil { return g.VisitThrowStmt(ctx.ThrowStmt().(*parser.ThrowStmtContext)) }
		return nil

	case *parser.ReturnStmtContext:
		return g.VisitReturnStmt(ctx)
	case *parser.IfStmtContext:
		return g.VisitIfStmt(ctx)
	case *parser.ForStmtContext:
		return g.VisitForStmt(ctx)
	case *parser.SwitchStmtContext:
		return g.VisitSwitchStmt(ctx)
	case *parser.BreakStmtContext:
		return g.VisitBreakStmt(ctx)
	case *parser.ContinueStmtContext:
		return g.VisitContinueStmt(ctx)
	case *parser.DeferStmtContext:
		return g.VisitDeferStmt(ctx)
	case *parser.TryStmtContext:
		return g.VisitTryStmt(ctx)
	case *parser.ThrowStmtContext:
		return g.VisitThrowStmt(ctx)

	// Expressions
	case *parser.ExpressionContext:
		return g.VisitExpression(ctx)
	case *parser.AdditiveExpressionContext:
		return g.VisitAdditiveExpression(ctx)
	case *parser.MultiplicativeExpressionContext:
		return g.VisitMultiplicativeExpression(ctx)
	case *parser.UnaryExpressionContext:
		return g.VisitUnaryExpression(ctx)
	case *parser.PostfixExpressionContext:
		return g.VisitPostfixExpression(ctx)
	case *parser.PrimaryExpressionContext:
		return g.VisitPrimaryExpression(ctx)
	case *parser.LiteralContext:
		return g.VisitLiteral(ctx)
		
	default:
		return g.BaseArcParserVisitor.Visit(tree)
	}
}