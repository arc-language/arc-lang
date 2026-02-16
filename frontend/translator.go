package frontend

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/ast"
	"github.com/arc-language/arc-lang/parser"
)

// Translate converts the raw CST into a clean AST.
func Translate(root parser.ICompilationUnitContext) *ast.File {
	t := &translator{}
	return t.file(root)
}

type translator struct{}

func (t *translator) pos(token antlr.Token) ast.Position {
	if token == nil {
		return ast.Position{}
	}
	return ast.Position{Line: token.GetLine(), Column: token.GetColumn()}
}

// --- Top Level ---

func (t *translator) file(ctx parser.ICompilationUnitContext) *ast.File {
	f := &ast.File{}

	// Namespace
	if ns := ctx.NamespaceDecl(); ns != nil {
		f.Namespace = ns.QualifiedName().GetText()
	}

	// Top Level Declarations
	for _, decl := range ctx.AllTopLevelDecl() {
		if d := t.topLevelDecl(decl); d != nil {
			switch node := d.(type) {
			case *ast.ImportDecl:
				f.Imports = append(f.Imports, node)
			default:
				f.Decls = append(f.Decls, node)
			}
		}
	}
	return f
}

func (t *translator) topLevelDecl(ctx parser.ITopLevelDeclContext) ast.Decl {
	if ctx.FuncDecl() != nil {
		return t.funcDecl(ctx.FuncDecl())
	}
	if ctx.ImportDecl() != nil {
		return t.importDecl(ctx.ImportDecl())
	}
	if ctx.TopLevelVarDecl() != nil {
		return t.varDecl(ctx.TopLevelVarDecl(), true) // isRef = true (var)
	}
	if ctx.TopLevelLetDecl() != nil {
		return t.letDecl(ctx.TopLevelLetDecl()) // isRef = false (let)
	}
	if ctx.InterfaceDecl() != nil {
		return t.interfaceDecl(ctx.InterfaceDecl())
	}
	return nil // TODO: Handle Enum, Extern, etc.
}

func (t *translator) importDecl(ctx parser.IImportDeclContext) *ast.ImportDecl {
	// Handles single line: import "fmt"
	// Note: Complex imports (grouped) are handled by the loop in file() essentially,
	// but strictly mapping 1:1 with grammar requires unwrapping the list.
	// For simplicity here, we assume single imports or map the first spec.
	spec := ctx.ImportSpec(0)
	path := strings.Trim(spec.STRING_LIT().GetText(), "\"")
	
	alias := ""
	if spec.ImportAlias() != nil {
		alias = spec.ImportAlias().GetText()
	}

	return &ast.ImportDecl{
		Alias: alias,
		Path:  path,
		Start: t.pos(ctx.GetStart()),
	}
}

// --- Declarations ---

func (t *translator) funcDecl(ctx parser.IFuncDeclContext) *ast.FuncDecl {
	fn := &ast.FuncDecl{
		Name:  ctx.IDENTIFIER().GetText(),
		Start: t.pos(ctx.GetStart()),
	}

	// Modifiers
	for _, mod := range ctx.AllFuncModifier() {
		if mod.ASYNC() != nil {
			fn.IsAsync = true
		}
		if mod.GPU() != nil {
			fn.IsGpu = true
		}
	}

	// Parameters
	if pl := ctx.ParamList(); pl != nil {
		for _, p := range pl.AllParam() {
			fn.Params = append(fn.Params, t.param(p))
		}
	}

	// Return Type
	if rt := ctx.ReturnType(); rt != nil {
		fn.ReturnType = t.typeRef(rt.TypeRef())
	}

	// Body
	if ctx.Block() != nil {
		fn.Body = t.blockStmt(ctx.Block())
	}

	return fn
}

func (t *translator) varDecl(ctx parser.ITopLevelVarDeclContext, isRef bool) *ast.VarDecl {
	v := &ast.VarDecl{
		Name:  ctx.IDENTIFIER().GetText(),
		IsRef: isRef,
		Start: t.pos(ctx.GetStart()),
	}

	if ctx.TypeRef() != nil {
		v.Type = t.typeRef(ctx.TypeRef())
	}

	if ctx.Expression() != nil {
		v.Value = t.expr(ctx.Expression())
	}

	return v
}

func (t *translator) letDecl(ctx parser.ITopLevelLetDeclContext) *ast.VarDecl {
	// Reuse VarDecl struct but IsRef = false
	v := &ast.VarDecl{
		Name:  ctx.IDENTIFIER().GetText(),
		IsRef: false,
		Start: t.pos(ctx.GetStart()),
	}
	if ctx.TypeRef() != nil {
		v.Type = t.typeRef(ctx.TypeRef())
	}
	if ctx.Expression() != nil {
		v.Value = t.expr(ctx.Expression())
	}
	return v
}

func (t *translator) interfaceDecl(ctx parser.IInterfaceDeclContext) *ast.InterfaceDecl {
	decl := &ast.InterfaceDecl{
		Name:  ctx.IDENTIFIER().GetText(),
		Start: t.pos(ctx.GetStart()),
	}
	
	for _, f := range ctx.AllInterfaceField() {
		decl.Fields = append(decl.Fields, &ast.Field{
			Name:  f.IDENTIFIER().GetText(),
			Type:  t.typeRef(f.TypeRef()),
			Start: t.pos(f.GetStart()),
		})
	}
	return decl
}

func (t *translator) param(ctx parser.IParamContext) *ast.Field {
	return &ast.Field{
		Name:  ctx.IDENTIFIER().GetText(),
		Type:  t.typeRef(ctx.ParamType().TypeRef()), // simplified, ignores mut for now
		Start: t.pos(ctx.GetStart()),
	}
}

// --- Statements ---

func (t *translator) stmt(ctx parser.IStatementContext) ast.Stmt {
	if ctx.LetStatement() != nil {
		return t.letStmt(ctx.LetStatement())
	}
	if ctx.ReturnStatement() != nil {
		return t.returnStmt(ctx.ReturnStatement())
	}
	if ctx.ExpressionStatement() != nil {
		return &ast.ExprStmt{X: t.expr(ctx.ExpressionStatement().Expression())}
	}
	if ctx.Block() != nil { // Implicit block statement?
		// grammar 'statement' doesn't usually hold block directly unless checked
	}
	if ctx.IfStatement() != nil {
		return t.ifStmt(ctx.IfStatement())
	}
    // ... Add For, Switch, Defer
	return nil
}

func (t *translator) blockStmt(ctx parser.IBlockContext) *ast.BlockStmt {
	b := &ast.BlockStmt{
		LBrace: t.pos(ctx.LBRACE().GetSymbol()),
		RBrace: t.pos(ctx.RBRACE().GetSymbol()),
	}
	for _, s := range ctx.AllStatement() {
		if stmt := t.stmt(s); stmt != nil {
			b.List = append(b.List, stmt)
		}
	}
	return b
}

func (t *translator) letStmt(ctx parser.ILetStatementContext) *ast.DeclStmt {
	// Local let is basically a declaration inside a statement wrapper
	decl := &ast.VarDecl{
		Name:  ctx.IDENTIFIER(0).GetText(),
		IsRef: false,
		Start: t.pos(ctx.GetStart()),
	}
	if ctx.TypeRef() != nil {
		decl.Type = t.typeRef(ctx.TypeRef())
	}
	if ctx.Expression() != nil {
		decl.Value = t.expr(ctx.Expression())
	}
	return &ast.DeclStmt{Decl: decl}
}

func (t *translator) returnStmt(ctx parser.IReturnStatementContext) *ast.ReturnStmt {
	ret := &ast.ReturnStmt{Start: t.pos(ctx.GetStart())}
	for _, e := range ctx.AllExpression() {
		ret.Results = append(ret.Results, t.expr(e))
	}
	return ret
}

func (t *translator) ifStmt(ctx parser.IIfStatementContext) *ast.IfStmt {
	// Logic for if/else if/else chains
	// Simplest case: One IF, One Block
	node := &ast.IfStmt{
		Cond:  t.expr(ctx.Expression(0)),
		Body:  t.blockStmt(ctx.Block(0)),
		Start: t.pos(ctx.GetStart()),
	}
	
	// Handle ELSE
	if len(ctx.AllELSE()) > 0 {
		// Check if it's "else if" (implies we have more expressions than blocks matched 1:1?)
		// The grammar usually nests them.
		// For simplicity in this snippet, we handle basic ELSE
		if len(ctx.AllBlock()) > 1 {
			node.Else = t.blockStmt(ctx.Block(1))
		}
	}
	return node
}

// --- Expressions ---

func (t *translator) expr(ctx parser.IExpressionContext) ast.Expr {
	// This requires unwrapping the specific context type because ANTLR 
	// generates different contexts for different alternatives (AddExpr, MulExpr, etc.)
	
	switch e := ctx.(type) {
	case *parser.PrimaryExprContext:
		return t.primary(e.Primary())
	
	case *parser.AddExprContext:
		return &ast.BinaryExpr{
			Left:  t.expr(e.Expression(0)),
			Op:    e.GetOp().GetText(),
			Right: t.expr(e.Expression(1)),
		}
	
	case *parser.MulExprContext:
		return &ast.BinaryExpr{
			Left:  t.expr(e.Expression(0)),
			Op:    e.GetOp().GetText(),
			Right: t.expr(e.Expression(1)),
		}

	case *parser.CallExprContext:
		call := &ast.CallExpr{
			Fun:   t.expr(e.Expression()),
			Start: t.pos(e.GetStart()),
            EndPos: t.pos(e.GetStop()),
		}
		if args := e.ArgumentList(); args != nil {
			for _, arg := range args.AllArgument() {
				call.Args = append(call.Args, t.expr(arg.Expression()))
			}
		}
		return call

	// ... Handle other expressions (Index, MemberAccess, etc.)
	}
	return nil
}

func (t *translator) primary(ctx parser.IPrimaryContext) ast.Expr {
	switch p := ctx.(type) {
	case *parser.IntLiteralContext:
		return &ast.BasicLit{
			Kind:  "INT",
			Value: p.GetText(),
			Pos:   t.pos(p.GetStart()),
		}
	case *parser.IdentExprContext:
		return &ast.Ident{
			Name: p.GetText(),
			Pos:  t.pos(p.GetStart()),
		}
	// ... Handle ParenExpr, etc.
	}
	return nil
}

// --- Types ---

func (t *translator) typeRef(ctx parser.ITypeRefContext) ast.TypeRef {
	if ctx == nil {
		return nil
	}
	if ctx.BaseType() != nil {
		// Check for primitives or named types
		bt := ctx.BaseType()
		if bt.PrimitiveType() != nil {
			return &ast.NamedType{
				Name: bt.PrimitiveType().GetText(),
				Pos:  t.pos(bt.GetStart()),
			}
		}
		if bt.QualifiedName() != nil {
			return &ast.NamedType{
				Name: bt.QualifiedName().GetText(),
				Pos:  t.pos(bt.GetStart()),
			}
		}
	}
	return nil
}