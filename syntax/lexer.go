package syntax

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/parser"
)

// tokenSliceSource implements antlr.TokenSource by embedding the original lexer.
// Embedding *parser.ArcLexer ensures we satisfy ALL interface requirements
// (including private methods like setTokenFactory) automatically.
type tokenSliceSource struct {
	*parser.ArcLexer
	tokens []antlr.Token
	index  int
}

// NextToken overrides the embedded lexer's method to serve tokens from our slice.
func (s *tokenSliceSource) NextToken() antlr.Token {
	if s.index >= len(s.tokens) {
		// Create a proper EOF token using the embedded lexer's source pair
		return antlr.CommonTokenFactoryDEFAULT.Create(
			s.GetTokenSourceCharStreamPair(),
			antlr.TokenEOF,
			"<EOF>",
			antlr.TokenDefaultChannel,
			-1, -1, 0, 0,
		)
	}
	t := s.tokens[s.index]
	s.index++
	return t
}

// createTokenStream wraps the generated lexer to perform Automatic Semicolon Insertion (ASI).
// It returns a token stream that includes synthetic SEMI tokens where newlines imply termination.
func createTokenStream(input string) *antlr.CommonTokenStream {
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewArcLexer(inputStream)
	lexer.RemoveErrorListeners()

	// 1. Fetch all raw tokens from the ANTLR lexer.
	allTokens := lexer.GetAllTokens()
	var processedTokens []antlr.Token

	// Grab the lexer's source pair once for minting synthetic tokens.
	sourcePair := lexer.GetTokenSourceCharStreamPair()

	// 2. Iterate through tokens and inject SEMI where appropriate.
	for i, t := range allTokens {
		processedTokens = append(processedTokens, t)

		if shouldInsertSemi(t, i, allTokens) {
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

	// 3. Create a wrapper that embeds the original lexer but serves the processed tokens.
	// This wrapper is now fully compatible with the TokenSource interface.
	wrapper := &tokenSliceSource{
		ArcLexer: lexer,
		tokens:   processedTokens,
		index:    0,
	}

	// 4. Create the stream using our wrapper as the source.
	return antlr.NewCommonTokenStream(wrapper, antlr.TokenDefaultChannel)
}

// shouldInsertSemi determines if a SEMI token should be inserted after token t.
func shouldInsertSemi(t antlr.Token, index int, all []antlr.Token) bool {
	isLast := index >= len(all)-1
	if !isLast {
		next := all[index+1]
		if next.GetLine() == t.GetLine() {
			return false
		}
	}

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