package irgen

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/ir"
	"github.com/arc-language/arc-lang/builder/types"
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
	loopStack    []loopInfo
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
	if s, ok := g.analysis.Scopes[ctx]; ok { g.currentScope = s }
}

func (g *Generator) exitScope() {
	if g.currentScope.Parent != nil { g.currentScope = g.currentScope.Parent }
}

func (g *Generator) getZeroValue(t types.Type) ir.Value {
	if types.IsInteger(t) { return g.ctx.Builder.ConstInt(t.(*types.IntType), 0) }
	if types.IsFloat(t) { return g.ctx.Builder.ConstFloat(t.(*types.FloatType), 0.0) }
	return g.ctx.Builder.ConstZero(t)
}

func (g *Generator) emitCast(val ir.Value, target types.Type) ir.Value {
	src := val.Type()
	if src.Equal(target) { return val }
	if types.IsInteger(src) && types.IsInteger(target) {
		if src.BitSize() > target.BitSize() { return g.ctx.Builder.CreateTrunc(val, target, "") }
		return g.ctx.Builder.CreateSExt(val, target, "")
	}
	if types.IsFloat(src) && types.IsFloat(target) {
		if src.BitSize() > target.BitSize() { return g.ctx.Builder.CreateFPTrunc(val, target, "") }
		return g.ctx.Builder.CreateFPExt(val, target, "")
	}
	if types.IsInteger(src) && types.IsFloat(target) { return g.ctx.Builder.CreateSIToFP(val, target, "") }
	return val
}

func (g *Generator) resolveType(ctx parser.ITypeContext) types.Type {
	tc := ctx.(*parser.TypeContext)
	if tc.PrimitiveType() != nil {
		switch tc.PrimitiveType().GetText() {
		case "int": return types.I64
		case "float": return types.F64
		case "bool": return types.I1
		case "void": return types.Void
		}
	}
	if tc.PointerType() != nil {
		return types.NewPointer(g.resolveType(tc.PointerType().Type_()))
	}
	if tc.IDENTIFIER() != nil {
		// Lookup struct types created in Pass 1
		name := tc.IDENTIFIER().GetText()
		if s, ok := g.currentScope.Resolve(name); ok { return s.Type }
	}
	return types.I64
}