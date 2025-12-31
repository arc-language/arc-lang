package semantics

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/diagnostic"
	"github.com/arc-language/arc-lang/symbol"
)

// AnalysisResult contains everything Pass 2 needs to generate code
type AnalysisResult struct {
	GlobalScope *symbol.Scope
	
	// Scopes maps AST nodes (like Blocks or FunctionDecls) to the Scope they created.
	// Pass 2 uses this to enter the correct scope at the right time.
	Scopes map[antlr.ParserRuleContext]*symbol.Scope
	
	// NodeTypes maps Expression AST nodes to their resolved Type.
	// Pass 2 uses this so it doesn't have to re-calculate types.
	NodeTypes map[antlr.ParseTree]types.Type
}

// Analyzer maintains the state during the semantic pass
type Analyzer struct {
	*parser.BaseArcParserVisitor
	
	// Inputs
	file string
	bag  *diagnostic.Bag
	
	// State
	globalScope  *symbol.Scope
	currentScope *symbol.Scope
	
	// Context Tracking
	currentFuncRetType types.Type // To check 'return' statements
	inLoop             bool       // To check 'break/continue' statements
	
	// Outputs
	scopes    map[antlr.ParserRuleContext]*symbol.Scope
	nodeTypes map[antlr.ParseTree]types.Type
}

// Analyze is the entry point for Pass 1
func Analyze(tree parser.ICompilationUnitContext, filename string, bag *diagnostic.Bag) (*AnalysisResult, error) {
	// Initialize Global Scope with Primitives
	global := symbol.NewScope(nil)
	registerPrimitives(global)

	analyzer := &Analyzer{
		BaseArcParserVisitor: &parser.BaseArcParserVisitor{},
		file:                 filename,
		bag:                  bag,
		globalScope:          global,
		currentScope:         global,
		scopes:               make(map[antlr.ParserRuleContext]*symbol.Scope),
		nodeTypes:            make(map[antlr.ParseTree]types.Type),
	}

	// Start the walk
	analyzer.Visit(tree)

	return &AnalysisResult{
		GlobalScope: analyzer.globalScope,
		Scopes:      analyzer.scopes,
		NodeTypes:   analyzer.nodeTypes,
	}, nil
}

func registerPrimitives(s *symbol.Scope) {
	s.Define("int", symbol.SymType, types.I64)
	s.Define("float", symbol.SymType, types.F64)
	s.Define("bool", symbol.SymType, types.I1)
	s.Define("char", symbol.SymType, types.I32) // Rune
	s.Define("void", symbol.SymType, types.Void)
	// Simplified String for now
	s.Define("string", symbol.SymType, types.NewPointer(types.I8)) 
}

// --- Scope Helpers ---

func (a *Analyzer) pushScope(ctx antlr.ParserRuleContext) {
	newScope := symbol.NewScope(a.currentScope)
	a.scopes[ctx] = newScope // Record for Pass 2
	a.currentScope = newScope
}

func (a *Analyzer) popScope() {
	if a.currentScope.Parent != nil {
		a.currentScope = a.currentScope.Parent
	}
}