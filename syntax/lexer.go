package syntax

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/parser"
)

// tokenSliceSource implements antlr.TokenSource backed by a pre-built token slice.
// It satisfies the interface required by antlr.CommonTokenStream.
type tokenSliceSource struct {
	tokens  []antlr.Token
	index   int
	factory antlr.TokenFactory
}

// NextToken returns the next token from the slice or EOF if exhausted.
func (s *tokenSliceSource) NextToken() antlr.Token {
	if s.index >= len(s.tokens) {
		return antlr.CommonTokenFactoryDEFAULT.Create(
			&antlr.TokenSourceCharStreamPair{},
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

// Required interface methods for antlr.TokenSource
func (s *tokenSliceSource) GetLine() int                        { return 0 }
func (s *tokenSliceSource) GetCharPositionInLine() int          { return 0 }
func (s *tokenSliceSource) GetInputStream() antlr.CharStream    { return nil }
func (s *tokenSliceSource) GetSourceName() string               { return "<slice>" }
func (s *tokenSliceSource) SetTokenFactory(f antlr.TokenFactory) { s.factory = f }
func (s *tokenSliceSource) GetTokenFactory() antlr.TokenFactory { return s.factory }

// More is a required method for the TokenSource interface in some ANTLR versions.
// For a slice playback, this is a no-op.
func (s *tokenSliceSource) More() {
	// No-op
}

// Skip is often required alongside More in Lexer interfaces. 
// We implement it as a no-op safety measure.
func (s *tokenSliceSource) Skip() {
	// No-op
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

	// 3. Build a stream from the original lexer (satisfies the Lexer type requirement),
	//    then swap the token source for our pre-built slice.
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	stream.SetTokenSource(&tokenSliceSource{
		tokens:  processedTokens,
		factory: antlr.CommonTokenFactoryDEFAULT,
	})
	return stream
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