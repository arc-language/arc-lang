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