package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/parser"
)

// CustomErrorListener implements ANTLR error listener for better error reporting
type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

func NewCustomErrorListener() *CustomErrorListener {
	return &CustomErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		Errors:               make([]string, 0),
	}
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	errorMsg := fmt.Sprintf("line %d:%d %s", line, column, msg)
	c.Errors = append(c.Errors, errorMsg)
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", errorMsg)
}

// ArcPrintListener is a comprehensive listener that prints all parse tree nodes
type ArcPrintListener struct {
	*parser.BaseArcParserListener
	indent int
}

func NewArcPrintListener() *ArcPrintListener {
	return &ArcPrintListener{
		BaseArcParserListener: &parser.BaseArcParserListener{},
		indent:                0,
	}
}

func (l *ArcPrintListener) printIndent(s string) {
	for i := 0; i < l.indent; i++ {
		fmt.Print("  ")
	}
	fmt.Println(s)
}

// ============================================================================
// COMPILATION UNIT
// ============================================================================

func (l *ArcPrintListener) EnterCompilationUnit(ctx *parser.CompilationUnitContext) {
	l.printIndent("╔═══ Compilation Unit ═══")
	l.indent++
}

func (l *ArcPrintListener) ExitCompilationUnit(ctx *parser.CompilationUnitContext) {
	l.indent--
	l.printIndent("╚═══ End Compilation Unit ═══")
}

// ============================================================================
// IMPORTS
// ============================================================================

func (l *ArcPrintListener) EnterImportDecl(ctx *parser.ImportDeclContext) {
	if ctx.STRING_LITERAL() != nil {
		l.printIndent(fmt.Sprintf("📦 Import: %s", ctx.STRING_LITERAL().GetText()))
	} else {
		l.printIndent("📦 Import (multiple):")
		l.indent++
	}
}

func (l *ArcPrintListener) ExitImportDecl(ctx *parser.ImportDeclContext) {
	if ctx.LPAREN() != nil {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterImportSpec(ctx *parser.ImportSpecContext) {
	l.printIndent(fmt.Sprintf("  - %s", ctx.STRING_LITERAL().GetText()))
}

// ============================================================================
// NAMESPACE
// ============================================================================

func (l *ArcPrintListener) EnterNamespaceDecl(ctx *parser.NamespaceDeclContext) {
	l.printIndent(fmt.Sprintf("🏛️ Namespace: %s", ctx.IDENTIFIER().GetText()))
}

// ============================================================================
// EXTERN DECLARATIONS
// ============================================================================

func (l *ArcPrintListener) EnterExternDecl(ctx *parser.ExternDeclContext) {
	l.printIndent(fmt.Sprintf("🔗 Extern Block: %s", ctx.IDENTIFIER().GetText()))
	l.indent++
}

func (l *ArcPrintListener) ExitExternDecl(ctx *parser.ExternDeclContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterExternFunctionDecl(ctx *parser.ExternFunctionDeclContext) {
	funcName := ctx.IDENTIFIER().GetText()
	alias := ""
	if ctx.STRING_LITERAL() != nil {
		alias = fmt.Sprintf(" (alias: %s)", ctx.STRING_LITERAL().GetText())
	}
	returnType := "void"
	if ctx.Type_() != nil {
		returnType = ctx.Type_().GetText()
	}
	params := ""
	if ctx.ExternParameterList() != nil {
		params = ctx.ExternParameterList().GetText()
	}
	l.printIndent(fmt.Sprintf("  ⚡ extern func %s%s(%s) -> %s", funcName, alias, params, returnType))
}

// ============================================================================
// FUNCTION DECLARATIONS
// ============================================================================

func (l *ArcPrintListener) EnterFunctionDecl(ctx *parser.FunctionDeclContext) {
	funcName := ctx.IDENTIFIER().GetText()
	returnType := "void"
	if ctx.Type_() != nil {
		returnType = ctx.Type_().GetText()
	}
	asyncPrefix := ""
	if ctx.ASYNC() != nil {
		asyncPrefix = "async "
	}
	l.printIndent(fmt.Sprintf("⚡ Function: %s%s -> %s", asyncPrefix, funcName, returnType))
	l.indent++
}

func (l *ArcPrintListener) ExitFunctionDecl(ctx *parser.FunctionDeclContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterParameter(ctx *parser.ParameterContext) {
	isSelf := ctx.SELF() != nil
	paramName := ctx.IDENTIFIER().GetText()
	paramType := ctx.Type_().GetText()
	selfPrefix := ""
	if isSelf {
		selfPrefix = "self "
	}
	l.printIndent(fmt.Sprintf("    - %s%s: %s", selfPrefix, paramName, paramType))
}

// ============================================================================
// STRUCT DECLARATIONS
// ============================================================================

func (l *ArcPrintListener) EnterStructDecl(ctx *parser.StructDeclContext) {
	structName := ctx.IDENTIFIER().GetText()
	l.printIndent(fmt.Sprintf("🏗️ Struct: %s", structName))
	l.indent++
}

func (l *ArcPrintListener) ExitStructDecl(ctx *parser.StructDeclContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterStructField(ctx *parser.StructFieldContext) {
	fieldName := ctx.IDENTIFIER().GetText()
	fieldType := ctx.Type_().GetText()
	l.printIndent(fmt.Sprintf("  • %s: %s", fieldName, fieldType))
}

// ============================================================================
// CLASS DECLARATIONS
// ============================================================================

func (l *ArcPrintListener) EnterClassDecl(ctx *parser.ClassDeclContext) {
	className := ctx.IDENTIFIER().GetText()
	l.printIndent(fmt.Sprintf("🏛️ Class: %s", className))
	l.indent++
}

func (l *ArcPrintListener) ExitClassDecl(ctx *parser.ClassDeclContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterClassField(ctx *parser.ClassFieldContext) {
	fieldName := ctx.IDENTIFIER().GetText()
	fieldType := ctx.Type_().GetText()
	l.printIndent(fmt.Sprintf("  • %s: %s", fieldName, fieldType))
}

// ============================================================================
// METHOD DECLARATIONS (FLAT)
// ============================================================================

func (l *ArcPrintListener) EnterMethodDecl(ctx *parser.MethodDeclContext) {
	methodName := ctx.IDENTIFIER(0).GetText()
	selfParam := ctx.IDENTIFIER(1).GetText()
	selfType := ctx.AllType_()[0].GetText()
	returnType := "void"
	if len(ctx.AllType_()) > 1 {
		returnType = ctx.AllType_()[len(ctx.AllType_())-1].GetText()
	}
	asyncPrefix := ""
	if ctx.ASYNC() != nil {
		asyncPrefix = "async "
	}
	l.printIndent(fmt.Sprintf("🔧 Flat Method: %s%s(self %s: %s) -> %s", asyncPrefix, methodName, selfParam, selfType, returnType))
	l.indent++
}

func (l *ArcPrintListener) ExitMethodDecl(ctx *parser.MethodDeclContext) {
	l.indent--
}

// ============================================================================
// MUTATING METHOD DECLARATIONS
// ============================================================================

func (l *ArcPrintListener) EnterMutatingDecl(ctx *parser.MutatingDeclContext) {
	methodName := ctx.IDENTIFIER(0).GetText()
	selfParam := ctx.IDENTIFIER(1).GetText()
	selfType := ctx.AllType_()[0].GetText()
	returnType := "void"
	if len(ctx.AllType_()) > 1 {
		returnType = ctx.AllType_()[len(ctx.AllType_())-1].GetText()
	}
	l.printIndent(fmt.Sprintf("🔄 Mutating Method: %s(self %s: %s) -> %s", methodName, selfParam, selfType, returnType))
	l.indent++
}

func (l *ArcPrintListener) ExitMutatingDecl(ctx *parser.MutatingDeclContext) {
	l.indent--
}

// ============================================================================
// DEINIT DECLARATIONS
// ============================================================================

func (l *ArcPrintListener) EnterDeinitDecl(ctx *parser.DeinitDeclContext) {
	selfParam := ctx.IDENTIFIER().GetText()
	selfType := ctx.Type_().GetText()
	l.printIndent(fmt.Sprintf("💀 Deinit(self %s: %s)", selfParam, selfType))
	l.indent++
}

func (l *ArcPrintListener) ExitDeinitDecl(ctx *parser.DeinitDeclContext) {
	l.indent--
}

// ============================================================================
// VARIABLE AND CONSTANT DECLARATIONS
// ============================================================================

func (l *ArcPrintListener) EnterVariableDecl(ctx *parser.VariableDeclContext) {
	varName := ctx.IDENTIFIER().GetText()
	typeInfo := "<inferred>"
	if ctx.Type_() != nil {
		typeInfo = ctx.Type_().GetText()
	}
	l.printIndent(fmt.Sprintf("📌 let %s: %s", varName, typeInfo))
	l.indent++
}

func (l *ArcPrintListener) ExitVariableDecl(ctx *parser.VariableDeclContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterConstDecl(ctx *parser.ConstDeclContext) {
	constName := ctx.IDENTIFIER().GetText()
	typeInfo := "<inferred>"
	if ctx.Type_() != nil {
		typeInfo = ctx.Type_().GetText()
	}
	l.printIndent(fmt.Sprintf("🔒 const %s: %s", constName, typeInfo))
	l.indent++
}

func (l *ArcPrintListener) ExitConstDecl(ctx *parser.ConstDeclContext) {
	l.indent--
}

// ============================================================================
// BLOCKS AND STATEMENTS
// ============================================================================

func (l *ArcPrintListener) EnterBlock(ctx *parser.BlockContext) {
	l.printIndent("{ Block")
	l.indent++
}

func (l *ArcPrintListener) ExitBlock(ctx *parser.BlockContext) {
	l.indent--
	l.printIndent("}")
}

func (l *ArcPrintListener) EnterAssignmentStmt(ctx *parser.AssignmentStmtContext) {
	lhs := ctx.LeftHandSide().GetText()
	op := "="
	if ctx.PLUS_ASSIGN() != nil {
		op = "+="
	} else if ctx.MINUS_ASSIGN() != nil {
		op = "-="
	} else if ctx.STAR_ASSIGN() != nil {
		op = "*="
	} else if ctx.SLASH_ASSIGN() != nil {
		op = "/="
	} else if ctx.PERCENT_ASSIGN() != nil {
		op = "%="
	}
	l.printIndent(fmt.Sprintf("✏️ Assignment: %s %s", lhs, op))
	l.indent++
}

func (l *ArcPrintListener) ExitAssignmentStmt(ctx *parser.AssignmentStmtContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterReturnStmt(ctx *parser.ReturnStmtContext) {
	if ctx.Expression() != nil {
		l.printIndent("↩️ Return:")
		l.indent++
	} else {
		l.printIndent("↩️ Return (void)")
	}
}

func (l *ArcPrintListener) ExitReturnStmt(ctx *parser.ReturnStmtContext) {
	if ctx.Expression() != nil {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterIfStmt(ctx *parser.IfStmtContext) {
	elseIfCount := len(ctx.AllIF()) - 1
	hasElse := len(ctx.AllELSE()) > elseIfCount
	info := ""
	if elseIfCount > 0 {
		info += fmt.Sprintf(" [%d else-if]", elseIfCount)
	}
	if hasElse {
		info += " [else]"
	}
	l.printIndent(fmt.Sprintf("🔀 If Statement%s", info))
	l.indent++
}

func (l *ArcPrintListener) ExitIfStmt(ctx *parser.IfStmtContext) {
	l.indent--
}

// ============================================================================
// FOR LOOP STATEMENTS
// ============================================================================

func (l *ArcPrintListener) EnterForStmt(ctx *parser.ForStmtContext) {
	loopType := "unknown"
	
	if len(ctx.AllSEMICOLON()) == 2 {
		loopType = "C-style"
	} else if ctx.IN() != nil {
		if ctx.COMMA() != nil {
			loopType = "for-in (key, value)"
		} else {
			loopType = "for-in"
		}
	} else if len(ctx.AllExpression()) == 1 {
		loopType = "condition-only"
	} else if len(ctx.AllExpression()) == 0 && ctx.VariableDecl() == nil {
		loopType = "infinite"
	}
	
	l.printIndent(fmt.Sprintf("🔁 For Loop [%s]", loopType))
	l.indent++
}

func (l *ArcPrintListener) ExitForStmt(ctx *parser.ForStmtContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterBreakStmt(ctx *parser.BreakStmtContext) {
	l.printIndent("🛑 Break")
}

func (l *ArcPrintListener) EnterContinueStmt(ctx *parser.ContinueStmtContext) {
	l.printIndent("⏭️ Continue")
}

func (l *ArcPrintListener) EnterDeferStmt(ctx *parser.DeferStmtContext) {
	l.printIndent("⏳ Defer Statement:")
	l.indent++
}

func (l *ArcPrintListener) ExitDeferStmt(ctx *parser.DeferStmtContext) {
	l.indent--
}

// ============================================================================
// EXPRESSIONS
// ============================================================================

func (l *ArcPrintListener) EnterRangeExpression(ctx *parser.RangeExpressionContext) {
	if ctx.RANGE() != nil {
		l.printIndent("📏 Range Expression (..)")
		l.indent++
	}
}

func (l *ArcPrintListener) ExitRangeExpression(ctx *parser.RangeExpressionContext) {
	if ctx.RANGE() != nil {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterAdditiveExpression(ctx *parser.AdditiveExpressionContext) {
	if len(ctx.AllMultiplicativeExpression()) > 1 {
		ops := make([]string, 0)
		for i := 0; i < len(ctx.AllPLUS()); i++ {
			ops = append(ops, "+")
		}
		for i := 0; i < len(ctx.AllMINUS()); i++ {
			ops = append(ops, "-")
		}
		l.printIndent(fmt.Sprintf("➕ Additive [%s]", strings.Join(ops, ", ")))
		l.indent++
	}
}

func (l *ArcPrintListener) ExitAdditiveExpression(ctx *parser.AdditiveExpressionContext) {
	if len(ctx.AllMultiplicativeExpression()) > 1 {
		l.indent--
	}
}

func (l *ArcPrintListener) EnterPostfixOp(ctx *parser.PostfixOpContext) {
	if ctx.DOT() != nil && ctx.IDENTIFIER() != nil {
		if ctx.LPAREN() != nil {
			l.printIndent(fmt.Sprintf("📞 Method Call: .%s()", ctx.IDENTIFIER().GetText()))
		} else {
			l.printIndent(fmt.Sprintf("🔍 Field Access: .%s", ctx.IDENTIFIER().GetText()))
		}
	} else if ctx.LPAREN() != nil && ctx.DOT() == nil {
		l.printIndent("📞 Function Call")
	} else if ctx.INCREMENT() != nil {
		l.printIndent("➕ Post-Increment (++)")
	} else if ctx.DECREMENT() != nil {
		l.printIndent("➖ Post-Decrement (--)")
	}
}

// ============================================================================
// LITERALS
// ============================================================================

func (l *ArcPrintListener) EnterLiteral(ctx *parser.LiteralContext) {
	if ctx.INTEGER_LITERAL() != nil {
		l.printIndent(fmt.Sprintf("🔢 Integer: %s", ctx.INTEGER_LITERAL().GetText()))
	} else if ctx.FLOAT_LITERAL() != nil {
		l.printIndent(fmt.Sprintf("🔢 Float: %s", ctx.FLOAT_LITERAL().GetText()))
	} else if ctx.STRING_LITERAL() != nil {
		l.printIndent(fmt.Sprintf("📝 String: %s", ctx.STRING_LITERAL().GetText()))
	} else if ctx.CHAR_LITERAL() != nil {
		l.printIndent(fmt.Sprintf("🔤 Char: %s", ctx.CHAR_LITERAL().GetText()))
	} else if ctx.BOOLEAN_LITERAL() != nil {
		l.printIndent(fmt.Sprintf("✓/✗ Boolean: %s", ctx.BOOLEAN_LITERAL().GetText()))
	} else if ctx.NULL() != nil {
		l.printIndent("∅ Null")
	}
}

func (l *ArcPrintListener) EnterPrimaryExpression(ctx *parser.PrimaryExpressionContext) {
	if ctx.IDENTIFIER() != nil {
		l.printIndent(fmt.Sprintf("🏷️ Identifier: %s", ctx.IDENTIFIER().GetText()))
	}
}

func (l *ArcPrintListener) EnterCastExpression(ctx *parser.CastExpressionContext) {
	targetType := ctx.Type_().GetText()
	l.printIndent(fmt.Sprintf("🔄 Cast to: %s", targetType))
	l.indent++
}

func (l *ArcPrintListener) ExitCastExpression(ctx *parser.CastExpressionContext) {
	l.indent--
}

func (l *ArcPrintListener) EnterIntrinsicExpression(ctx *parser.IntrinsicExpressionContext) {
	intrinsicName := ""
	if ctx.SIZEOF() != nil {
		intrinsicName = "sizeof"
	} else if ctx.ALIGNOF() != nil {
		intrinsicName = "alignof"
	} else if ctx.MEMSET() != nil {
		intrinsicName = "memset"
	} else if ctx.MEMCPY() != nil {
		intrinsicName = "memcpy"
	} else if ctx.MEMMOVE() != nil {
		intrinsicName = "memmove"
	} else if ctx.STRLEN() != nil {
		intrinsicName = "strlen"
	} else if ctx.MEMCHR() != nil {
		intrinsicName = "memchr"
	} else if ctx.VA_START() != nil {
		intrinsicName = "va_start"
	} else if ctx.VA_ARG() != nil {
		intrinsicName = "va_arg"
	} else if ctx.VA_END() != nil {
		intrinsicName = "va_end"
	} else if ctx.RAISE() != nil {
		intrinsicName = "raise"
	} else if ctx.MEMCMP() != nil {
		intrinsicName = "memcmp"
	} else if ctx.BIT_CAST() != nil {
		intrinsicName = "bit_cast"
	}
	l.printIndent(fmt.Sprintf("🔧 Intrinsic: %s", intrinsicName))
}

// ============================================================================
// PARSER FUNCTIONS
// ============================================================================

func ParseArcFile(filename string) (antlr.Tree, error) {
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}

	lexer := parser.NewArcLexer(input)
	lexer.RemoveErrorListeners()
	errorListener := NewCustomErrorListener()
	lexer.AddErrorListener(errorListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewArcParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	tree := p.CompilationUnit()

	if len(errorListener.Errors) > 0 {
		return tree, fmt.Errorf("parsing failed with %d error(s)", len(errorListener.Errors))
	}

	return tree, nil
}

func WalkParseTree(tree antlr.Tree, listener antlr.ParseTreeListener) {
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <arc-source-file>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s example.arc\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: File '%s' does not exist\n", filename)
		os.Exit(1)
	}

	fmt.Printf("╔═══════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║          ARC LANGUAGE PARSER - COMPREHENSIVE TEST             ║\n")
	fmt.Printf("╚═══════════════════════════════════════════════════════════════╝\n")
	fmt.Printf("\n📂 File: %s\n", filename)
	fmt.Println(strings.Repeat("═", 65))

	tree, err := ParseArcFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n❌ Parse error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n🌳 PARSE TREE STRUCTURE:")
	fmt.Println(strings.Repeat("─", 65))

	listener := NewArcPrintListener()
	WalkParseTree(tree, listener)

	fmt.Println(strings.Repeat("═", 65))
	fmt.Println("✅ Parsing completed successfully!")
	fmt.Println()
}