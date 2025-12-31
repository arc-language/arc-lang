package irgen

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/pkg/compiler" // For Context
	"github.com/arc-language/arc-lang/pkg/semantics"
	"github.com/arc-language/arc-lang/pkg/symbol"
)

// Generator implements Pass 2 (IR Generation)
type Generator struct {
	*parser.BaseArcParserVisitor
	
	// State
	ctx      *compiler.Context
	analysis *semantics.AnalysisResult
	
	// Scoping (Traversing Pass 1 scopes)
	currentScope *symbol.Scope
	
	// Function Context
	deferStack *DeferStack
}

// Generate is the entry point for Pass 2
func Generate(tree parser.ICompilationUnitContext, moduleName string, analysis *semantics.AnalysisResult) *ir.Module {
	// Initialize the Builder Context
	// Note: We create a fresh Context just for building
	ctx := compiler.NewContext("", moduleName) 
	
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
	// We retrieve the specific scope created during Pass 1 for this node
	if scope, ok := g.analysis.Scopes[ctx]; ok {
		g.currentScope = scope
	}
}

func (g *Generator) exitScope() {
	if g.currentScope.Parent != nil {
		g.currentScope = g.currentScope.Parent
	}
}