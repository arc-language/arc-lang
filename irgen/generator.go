package irgen

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/context"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/semantics"
	"github.com/arc-language/arc-lang/symbol"
)

// Generator implements Pass 2 (IR Generation)
type Generator struct {
	*parser.BaseArcParserVisitor
	ctx      *context.Context
	analysis *semantics.AnalysisResult
	currentScope *symbol.Scope
	deferStack   *DeferStack
	
	// Added loop stack
	loopStack []loopInfo
}

// Generate is the entry point for Pass 2
func Generate(tree parser.ICompilationUnitContext, moduleName string, analysis *semantics.AnalysisResult) *ir.Module {
	// Initialize the Builder Context using the new package
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

// --- Scope Management ---

func (g *Generator) enterScope(ctx antlr.ParserRuleContext) {
	// Retrieve the specific scope created during Pass 1 for this node
	if scope, ok := g.analysis.Scopes[ctx]; ok {
		g.currentScope = scope
	}
}

func (g *Generator) exitScope() {
	if g.currentScope.Parent != nil {
		g.currentScope = g.currentScope.Parent
	}
}