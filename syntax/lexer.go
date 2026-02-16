package syntax

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/parser"
)

// tokenSliceSource implements antlr.TokenSource backed by a pre-built slice.
type tokenSliceSource struct {
	tokens []antlr.Token
	index  int
}

func (s *tokenSliceSource) NextToken() antlr.Token {
	if s.index >= len(s.tokens) {
		// Return a bare EOF token.
		tok := antlr.CommonTokenFactoryDEFAULT.Create(
			&antlr.TokenSourceCharStreamPair{},
			antlr.TokenEOF,
			"<EOF>",
			antlr.TokenDefaultChannel,
			-1, -1, 0, 0,
		)
		return tok
	}
	t := s.tokens[s.index]
	s.index++
	return t
}

func (s *tokenSliceSource) GetSourceName() string      { return "<slice>" }
func (s *tokenSliceSource) GetInputStream() antlr.CharStream { return nil }

// createTokenStream wraps the generated lexer to perform Automatic Semicolon Insertion (ASI).
// It returns a token stream that includes synthetic SEMI tokens where newlines imply termination.
func createTokenStream(input string) *antlr.CommonTokenStream {
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewArcLexer(inputStream)
	lexer.RemoveErrorListeners()

	// 1. Fetch all raw tokens from the ANTLR lexer.
	allTokens := lexer.GetAllTokens()
	var processedTokens []antlr.Token

	// Grab the lexer's source pair once for use when minting synthetic tokens.
	sourcePair := lexer.GetTokenSourceCharStreamPair()

	// 2. Iterate through tokens and inject SEMI where appropriate.
	for i, t := range allTokens {
		processedTokens = append(processedTokens, t)

		if shouldInsertSemi(t, i, allTokens) {
			// Use the factory so the token is a proper CommonToken with all
			// fields set correctly. Line and column come from the real token.
			semi := antlr.CommonTokenFactoryDEFAULT.Create(
				sourcePair,
				parser.ArcLexerSEMI,
				";",
				antlr.TokenDefaultChannel,
				-1, -1,
				t.GetLine(),
				t.GetColumn()+len(t.GetText()),
			)
			processedTokens = append(processedTokens, semi)
		}
	}

	// 3. Build the stream from the original lexer, then swap the token source
	//    for our slice so the parser sees the modified token sequence.
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	stream.SetTokenSource(&tokenSliceSource{tokens: processedTokens})
	return stream
}

// shouldInsertSemi determines if a SEMI token should be inserted after token t.
// Rule: Insert SEMI if the current token is a "statement ender" AND
// the next token is on a new line (or if we are at EOF).
func shouldInsertSemi(t antlr.Token, index int, all []antlr.Token) bool {
	// Condition A: We must be at the effective end of a line.
	isLast := index >= len(all)-1
	if !isLast {
		next := all[index+1]
		if next.GetLine() == t.GetLine() {
			return false
		}
	}

	// Condition B: The current token must be one that is allowed to end a statement.
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
		parser.ArcLexerRPAREN,
		parser.ArcLexerRBRACKET,
		parser.ArcLexerRBRACE,
		parser.ArcLexerINC,
		parser.ArcLexerDEC:
		return true
	case parser.ArcLexerINT8, parser.ArcLexerINT16, parser.ArcLexerINT32, parser.ArcLexerINT64,
		parser.ArcLexerUINT8, parser.ArcLexerUINT16, parser.ArcLexerUINT32, parser.ArcLexerUINT64,
		parser.ArcLexerFLOAT32, parser.ArcLexerFLOAT64,
		parser.ArcLexerBOOL, parser.ArcLexerSTRING, parser.ArcLexerBYTE, parser.ArcLexerCHAR,
		parser.ArcLexerUSIZE, parser.ArcLexerISIZE:
		return true
	}
	return false
}