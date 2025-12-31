package semantics

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/diagnostic"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/pkg/symbol"
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
	initGlobalScope(global)

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

	a.Visit(tree)

	return &AnalysisResult{
		GlobalScope:   a.globalScope,
		Scopes:        a.scopes,
		NodeTypes:     a.nodeTypes,
		StructIndices: a.structIndices,
	}, nil
}

func initGlobalScope(s *symbol.Scope) {
	s.Define("int", symbol.SymType, types.I64)
	s.Define("int32", symbol.SymType, types.I32)
	s.Define("float", symbol.SymType, types.F64)
	s.Define("bool", symbol.SymType, types.I1)
	s.Define("void", symbol.SymType, types.Void)
	s.Define("string", symbol.SymType, types.NewPointer(types.I8))
}

func (a *Analyzer) pushScope(ctx antlr.ParserRuleContext) {
	s := symbol.NewScope(a.currentScope)
	a.scopes[ctx] = s
	a.currentScope = s
}

func (a *Analyzer) popScope() {
	if a.currentScope.Parent != nil {
		a.currentScope = a.currentScope.Parent
	}
}

// Visit implements manual dispatch to ensure our methods are called
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
	case *parser.BlockContext:
		return a.VisitBlock(ctx)
	case *parser.StatementContext:
		// Dispatch statement types manually if needed, or rely on children
		// But usually Statement just wraps one child
		return a.BaseArcParserVisitor.Visit(tree)
	case *parser.ReturnStmtContext:
		return a.VisitReturnStmt(ctx)
	case *parser.IfStmtContext:
		return a.VisitIfStmt(ctx)
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
		return a.BaseArcParserVisitor.Visit(tree)
	}
}