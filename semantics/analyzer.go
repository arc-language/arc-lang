package semantics

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/diagnostic"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

type AnalysisResult struct {
	GlobalScope   *symbol.Scope
	Scopes        map[antlr.ParserRuleContext]*symbol.Scope
	NodeTypes     map[antlr.ParseTree]types.Type
	StructIndices map[string]map[string]int
}

type Analyzer struct {
	*parser.BaseArcParserVisitor
	file string
	bag  *diagnostic.Bag
	
	globalScope  *symbol.Scope
	currentScope *symbol.Scope
	
	currentFuncRetType types.Type
	
	scopes        map[antlr.ParserRuleContext]*symbol.Scope
	nodeTypes     map[antlr.ParseTree]types.Type
	structIndices map[string]map[string]int
}

func Analyze(tree parser.ICompilationUnitContext, filename string, bag *diagnostic.Bag) (*AnalysisResult, error) {
	global := symbol.NewScope(nil)
	symbol.InitGlobalScope(global)

	a := &Analyzer{
		BaseArcParserVisitor: &parser.BaseArcParserVisitor{},
		file:                 filename,
		bag:                  bag,
		globalScope:          global,
		currentScope:         global,
		scopes:               make(map[antlr.ParserRuleContext]*symbol.Scope),
		nodeTypes:            make(map[antlr.ParseTree]types.Type),
		structIndices:        make(map[string]map[string]int),
	}

	// Start visitation
	a.Visit(tree)

	return &AnalysisResult{
		GlobalScope:   a.globalScope,
		Scopes:        a.scopes,
		NodeTypes:     a.nodeTypes,
		StructIndices: a.structIndices,
	}, nil
}

func (a *Analyzer) pushScope(ctx antlr.ParserRuleContext) {
	s := symbol.NewScope(a.currentScope)
	if a.currentScope.Parent == nil {
		s.DebugName = "FunctionScope"
	} else {
		s.DebugName = "BlockScope"
	}
	a.scopes[ctx] = s
	a.currentScope = s
}

func (a *Analyzer) popScope() {
	if a.currentScope.Parent != nil {
		a.currentScope = a.currentScope.Parent
	}
}

// Visit implements manual dispatch to ensure our methods are called.
// CRITICAL FIX: We must use tree.Accept(a) so that 'a' (the Analyzer) 
// is passed as the visitor, not the embedded BaseArcParserVisitor.
func (a *Analyzer) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil { return nil }

	switch ctx := tree.(type) {
	case *parser.CompilationUnitContext:
		return a.VisitCompilationUnit(ctx)
	case *parser.TopLevelDeclContext:
		return a.VisitTopLevelDecl(ctx)
	case *parser.FunctionDeclContext:
		return a.VisitFunctionDecl(ctx)
	case *parser.VariableDeclContext:
		return a.VisitVariableDecl(ctx)
	case *parser.StructDeclContext:
		return a.VisitStructDecl(ctx)
	case *parser.ClassDeclContext:
		return a.VisitClassDecl(ctx)
	case *parser.BlockContext:
		return a.VisitBlock(ctx)
	case *parser.ReturnStmtContext:
		return a.VisitReturnStmt(ctx)
	case *parser.IfStmtContext:
		return a.VisitIfStmt(ctx)
	case *parser.ForStmtContext:
		// Add specific visit method if you implemented it, otherwise fallthrough
		return ctx.Accept(a)
		
	// Expressions
	case *parser.ExpressionContext:
		return a.VisitExpression(ctx)
	case *parser.AdditiveExpressionContext:
		return a.VisitAdditiveExpression(ctx)
	case *parser.MultiplicativeExpressionContext:
		return a.VisitMultiplicativeExpression(ctx)
	case *parser.UnaryExpressionContext:
		return a.VisitUnaryExpression(ctx)
	case *parser.PostfixExpressionContext:
		return a.VisitPostfixExpression(ctx)
	case *parser.PrimaryExpressionContext:
		return a.VisitPrimaryExpression(ctx)
	case *parser.LiteralContext:
		return a.VisitLiteral(ctx)
		
	default:
		// Universal fallback: call Accept on the node passing 'a' as the visitor.
		// This ensures that if ctx is a StatementContext, it calls a.VisitStatement(ctx),
		// which delegates to children using 'a'.
		return tree.Accept(a)
	}
}