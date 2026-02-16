package syntax

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/parser"
)

// createTokenStream wraps the generated lexer to perform Automatic Semicolon Insertion (ASI).
// It returns a token stream that includes synthetic SEMI tokens where newlines imply termination.
func createTokenStream(input string) *antlr.CommonTokenStream {
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewArcLexer(inputStream)

	// Remove default error listeners to keep the console clean during lexing.
	lexer.RemoveErrorListeners()

	// 1. Fetch all raw tokens from the ANTLR lexer.
	allTokens := lexer.GetAllTokens()
	var processedTokens []antlr.Token

	// 2. Iterate through tokens and inject SEMI where appropriate.
	for i, t := range allTokens {
		// Always add the current token to the stream.
		processedTokens = append(processedTokens, t)

		// Check if we need to insert a semicolon after this token.
		if shouldInsertSemi(t, i, allTokens) {
			// Create a synthetic semicolon token.
			// We calculate its position relative to the current token to keep source maps sane.
			semi := antlr.NewCommonToken(
				lexer.GetTokenSource(),
				parser.ArcLexerSEMI,
				";",
				antlr.TokenDefaultChannel,
				-1, -1,
			)
			semi.SetLine(t.GetLine())
			semi.SetColumn(t.GetColumn() + len(t.GetText()))
			
			processedTokens = append(processedTokens, semi)
		}
	}

	// 3. Create a new stream from our modified token list.
	tokenSource := antlr.NewListTokenSource(processedTokens)
	return antlr.NewCommonTokenStream(tokenSource, antlr.TokenDefaultChannel)
}

// shouldInsertSemi determines if a SEMI token should be inserted after token t.
// Rule: Insert SEMI if the current token is a "statement ender" AND 
// the next token is on a new line (or if we are at EOF).
func shouldInsertSemi(t antlr.Token, index int, all []antlr.Token) bool {
	// Condition A: We must be at the effective end of a line.
	// This happens if we are the last token, or if the next token is on a different line.
	isLast := index >= len(all)-1
	if !isLast {
		next := all[index+1]
		if next.GetLine() == t.GetLine() {
			return false // Next token is on the same line; implicit termination does not apply.
		}
	}

	// Condition B: The current token must be one that is allowed to end a statement.
	// This list is derived from the "post-lexer semicolon insertion pass" note in your grammar.
	switch t.GetTokenType() {
	case parser.ArcLexerIDENTIFIER,
		parser.ArcLexerINT_LIT,
		parser.ArcLexerHEX_LIT,
		parser.ArcLexerFLOAT_LIT,
		parser.ArcLexerSTRING_LIT,
		parser.ArcLexerCHAR_LIT,
		parser.ArcLexerTRUE,
		parser.ArcLexerFALSE,
		parser.ArcLexerNULL,
		parser.ArcLexerRETURN,
		parser.ArcLexerBREAK,
		parser.ArcLexerCONTINUE,
		parser.ArcLexerRPAREN,   // )
		parser.ArcLexerRBRACKET, // ]
		parser.ArcLexerRBRACE,   // }
		parser.ArcLexerINC,      // ++
		parser.ArcLexerDEC:      // --
		return true

	// Primitive types can also end a statement (e.g. inside a type alias: `type X = int32`)
	case parser.ArcLexerINT8, parser.ArcLexerINT16, parser.ArcLexerINT32, parser.ArcLexerINT64,
		parser.ArcLexerUINT8, parser.ArcLexerUINT16, parser.ArcLexerUINT32, parser.ArcLexerUINT64,
		parser.ArcLexerFLOAT32, parser.ArcLexerFLOAT64,
		parser.ArcLexerBOOL, parser.ArcLexerSTRING, parser.ArcLexerBYTE, parser.ArcLexerCHAR,
		parser.ArcLexerUSIZE, parser.ArcLexerISIZE:
		return true
	}

	return false
}