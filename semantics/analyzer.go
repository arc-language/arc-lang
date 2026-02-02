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
	currentNamespacePrefix string
	
	// Phase control for multi-pass analysis
	// 0 = Single Pass (Legacy)
	// 1 = Declaration Scan (Types, Function Signatures)
	// 2 = Body Analysis (Statements, Expressions)
	Phase int 
	
	scopes        map[antlr.ParserRuleContext]*symbol.Scope
	nodeTypes     map[antlr.ParseTree]types.Type
	structIndices map[string]map[string]int
}

// NewAnalyzer creates an analyzer with an existing global scope.
func NewAnalyzer(globalScope *symbol.Scope, filename string, bag *diagnostic.Bag) *Analyzer {
	return &Analyzer{
		BaseArcParserVisitor: &parser.BaseArcParserVisitor{},
		file:                 filename,
		bag:                  bag,
		globalScope:          globalScope,
		currentScope:         globalScope,
		scopes:               make(map[antlr.ParserRuleContext]*symbol.Scope),
		nodeTypes:            make(map[antlr.ParseTree]types.Type),
		structIndices:        make(map[string]map[string]int),
	}
}

// Analyze performs semantic analysis on a single AST.
func (a *Analyzer) Analyze(tree parser.ICompilationUnitContext, result *AnalysisResult) {
	// Import state from shared result
	for k, v := range result.StructIndices {
		a.structIndices[k] = v
	}
	for k, v := range result.Scopes {
		a.scopes[k] = v
	}

	a.Visit(tree)

	// Export state back to shared result
	for k, v := range a.scopes {
		result.Scopes[k] = v
	}
	for k, v := range a.nodeTypes {
		result.NodeTypes[k] = v
	}
	for k, v := range a.structIndices {
		result.StructIndices[k] = v
	}
}

func (a *Analyzer) pushScope(ctx antlr.ParserRuleContext) {
	// Reuse existing scope if available (important for Phase 2)
	if s, ok := a.scopes[ctx]; ok {
		a.currentScope = s
		return
	}

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

// Visit dispatch
func (a *Analyzer) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil { return nil }

	switch ctx := tree.(type) {
	case *parser.CompilationUnitContext: return a.VisitCompilationUnit(ctx)
	case *parser.TopLevelDeclContext: return a.VisitTopLevelDecl(ctx)
	case *parser.NamespaceDeclContext: return a.VisitNamespaceDecl(ctx)
	case *parser.FunctionDeclContext: return a.VisitFunctionDecl(ctx)
	case *parser.VariableDeclContext: return a.VisitVariableDecl(ctx)
	case *parser.StructDeclContext: return a.VisitStructDecl(ctx)
	case *parser.ClassDeclContext: return a.VisitClassDecl(ctx)
	case *parser.EnumDeclContext: return a.VisitEnumDecl(ctx)
	case *parser.ConstDeclContext: return a.VisitConstDecl(ctx)
	case *parser.ExternCDeclContext: return a.VisitExternCDecl(ctx)
	case *parser.ExternCppDeclContext: return a.VisitExternCppDecl(ctx)
	// Added support for ObjC dispatch to match new grammar
	case *parser.ExternObjCDeclContext: return a.VisitExternObjCDecl(ctx)
	
	case *parser.BlockContext: return a.VisitBlock(ctx)
	case *parser.StatementContext: return a.VisitStatement(ctx)
	case *parser.ReturnStmtContext: return a.VisitReturnStmt(ctx)
	case *parser.IfStmtContext: return a.VisitIfStmt(ctx)
	case *parser.ForStmtContext: return a.VisitForStmt(ctx)
	case *parser.SwitchStmtContext: return a.VisitSwitchStmt(ctx)
	case *parser.ExpressionStmtContext: return a.VisitExpressionStmt(ctx)
	case *parser.BreakStmtContext: return a.VisitBreakStmt(ctx)
	case *parser.ContinueStmtContext: return a.VisitContinueStmt(ctx)
	case *parser.DeferStmtContext: return a.VisitDeferStmt(ctx)
	case *parser.AssignmentStmtContext: return a.VisitAssignmentStmt(ctx)

	case *parser.ExpressionContext: return a.VisitExpression(ctx)
	case *parser.AdditiveExpressionContext: return a.VisitAdditiveExpression(ctx)
	case *parser.MultiplicativeExpressionContext: return a.VisitMultiplicativeExpression(ctx)
	case *parser.UnaryExpressionContext: return a.VisitUnaryExpression(ctx)
	case *parser.PostfixExpressionContext: return a.VisitPostfixExpression(ctx)
	case *parser.PrimaryExpressionContext: return a.VisitPrimaryExpression(ctx)
	case *parser.LiteralContext: return a.VisitLiteral(ctx)
	case *parser.StructLiteralContext: return a.VisitStructLiteral(ctx)
	case *parser.InitializerListContext: return a.VisitInitializerList(ctx)
		
	default:
		return tree.Accept(a)
	}
}