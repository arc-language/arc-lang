package semantics

import (
	"strconv"
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/builder/types"
	"github.com/arc-language/arc-lang/diagnostic"
	"github.com/arc-language/arc-lang/parser"
	"github.com/arc-language/arc-lang/symbol"
)

type AnalysisResult struct {
	GlobalScope *symbol.Scope
	Scopes      map[antlr.ParserRuleContext]*symbol.Scope
	NodeTypes   map[antlr.ParseTree]types.Type
	// Map to store Struct Name -> Field Name -> Index
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

// resolveType from original visitor
func (a *Analyzer) resolveType(ctx parser.ITypeContext) types.Type {
	if ctx == nil { return types.Void }
	tc := ctx.(*parser.TypeContext)

	if tc.PrimitiveType() != nil {
		name := tc.PrimitiveType().GetText()
		if s, ok := a.currentScope.Resolve(name); ok && s.Kind == symbol.SymType {
			return s.Type
		}
		// Default to I64 if unknown (matching old logic)
		return types.I64
	}
	if tc.PointerType() != nil {
		return types.NewPointer(a.resolveType(tc.PointerType().Type_()))
	}
	if tc.ArrayType() != nil {
		elem := a.resolveType(tc.ArrayType().Type_())
		size := int64(0)
		if s := tc.ArrayType().ArraySize(); s != nil {
			if s.INTEGER_LITERAL() != nil {
				size, _ = strconv.ParseInt(s.INTEGER_LITERAL().GetText(), 0, 64)
			}
		}
		return types.NewArray(elem, size)
	}
	if tc.IDENTIFIER() != nil {
		name := tc.IDENTIFIER().GetText()
		if s, ok := a.currentScope.Resolve(name); ok && s.Kind == symbol.SymType {
			return s.Type
		}
	}
	return types.I64
}

// Compatibility check from old logic
func areTypesCompatible(src, dest types.Type) bool {
	if src.Equal(dest) { return true }
	if types.IsInteger(src) && types.IsInteger(dest) { return true } // Allow implicit int casting
	if types.IsFloat(src) && types.IsFloat(dest) { return true }
	// Pointer/Array checks can be added here
	return false
}