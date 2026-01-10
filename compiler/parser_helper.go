package compiler

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/parser" // The generated ANTLR parser
	"github.com/arc-language/arc-lang/diagnostic"
)

// syntaxListener bridges ANTLR errors to our Diagnostic system
type syntaxListener struct {
	*antlr.DefaultErrorListener
	file string
	bag  *diagnostic.Bag
}

func (l *syntaxListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// Add the error to our shared bag
	l.bag.Report(l.file, line, column, "Syntax Error: %s", msg)
}

// Parse handles reading the file and running the ANTLR parser
func Parse(filepath string) (parser.ICompilationUnitContext, *diagnostic.Bag) {
	bag := diagnostic.NewBag()

	// Setup Input Stream
	input, err := antlr.NewFileStream(filepath)
	if err != nil {
		bag.Report(filepath, 0, 0, "File IO Error: %v", err)
		return nil, bag
	}

	// Setup Lexer
	lexer := parser.NewArcLexer(input)
	lexer.RemoveErrorListeners() // Remove console printer

	// Setup Token Stream
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Setup Parser
	p := parser.NewArcParser(stream)
	p.RemoveErrorListeners() // Remove console printer
	
	// Add our custom listener
	listener := &syntaxListener{
		file: filepath,
		bag:  bag,
	}
	p.AddErrorListener(listener)

	// Parse
	// Note: This assumes your grammar's entry point is named 'compilationUnit'
	tree := p.CompilationUnit()
	
	return tree, bag
}