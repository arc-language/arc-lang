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
	ctx          *context.Context
	analysis     *semantics.AnalysisResult
	currentScope *symbol.Scope
	deferStack   *DeferStack
	
	// Loop stack for break/continue
	loopStack []loopInfo
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