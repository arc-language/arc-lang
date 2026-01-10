package compiler

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/arc-language/arc-lang/diagnostic"
	"github.com/arc-language/arc-lang/parser" // The generated ANTLR code
)

// Parse reads a file and returns the ANTLR parse tree and any errors.
func Parse(path string) (parser.ISourceUnitContext, *diagnostic.Bag) {
	input, err := antlr.NewFileStream(path)
	if err != nil {
		errs := diagnostic.NewBag()
		errs.Report(diagnostic.Error{
			Msg:  "Failed to read file: " + err.Error(),
			File: path,
		})
		return nil, errs
	}

	// 1. Setup Lexer
	lexer := parser.NewArcLexer(input)
	lexer.RemoveErrorListeners()
	
	lexErrors := diagnostic.NewBag()
	lexerListener := diagnostic.NewErrorListener(path, lexErrors)
	lexer.AddErrorListener(lexerListener)

	stream := antlr.NewCommonTokenStream(lexer, 0)

	// 2. Setup Parser
	p := parser.NewArcParser(stream)
	p.RemoveErrorListeners()
	
	parseErrors := diagnostic.NewBag()
	parseListener := diagnostic.NewErrorListener(path, parseErrors)
	p.AddErrorListener(parseListener)

	// 3. Parse Entry Point (SourceUnit)
	tree := p.SourceUnit()

	// 4. Combine Errors
	allErrors := diagnostic.NewBag()
	allErrors.Errors = append(allErrors.Errors, lexErrors.Errors...)
	allErrors.Errors = append(allErrors.Errors, parseErrors.Errors...)

	return tree, allErrors
}