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
}

// ParseArcSource parses Arc source code from a string
func ParseArcSource(source string) error {
	// Create input stream from string
	input := antlr.NewInputStream(source)

	// Create the lexer
	lexer := parser.NewArcLexer(input)
	
	// Remove default error listeners and add custom one
	lexer.RemoveErrorListeners()
	errorListener := NewCustomErrorListener()
	lexer.AddErrorListener(errorListener)

	// Create token stream
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewArcParser(stream)
	
	// Remove default error listeners and add custom one
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	// Parse the compilation unit (starting rule)
	_ = p.CompilationUnit()

	// Check for parsing errors
	if len(errorListener.Errors) > 0 {
		return fmt.Errorf("parsing failed: %s", strings.Join(errorListener.Errors, "; "))
	}

	return nil
}

// TestCase represents a single test case
type TestCase struct {
	Name string
	Code func() string
}

// Global test registry
var AllTests []TestCase

// RegisterTest adds a test to the global registry
func RegisterTest(name string, codeFn func() string) {
	AllTests = append(AllTests, TestCase{Name: name, Code: codeFn})
}

func main() {
	// Run all tests
	passed := 0
	failed := 0

	fmt.Println("Running Arc Language Parser Tests")
	fmt.Println(strings.Repeat("=", 70))

	for i, tc := range AllTests {
		code := tc.Code()
		err := ParseArcSource(code)
		
		if err != nil {
			failed++
			fmt.Printf("❌ Test %2d FAILED: %-35s\n", i+1, tc.Name)
			fmt.Printf("   Error: %s\n", err)
			if len(code) < 200 {
				fmt.Printf("   Code: %s\n", strings.ReplaceAll(code, "\n", "\\n"))
			}
		} else {
			passed++
			fmt.Printf("✅ Test %2d PASSED: %-35s\n", i+1, tc.Name)
		}
	}

	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("\nResults: %d passed, %d failed out of %d tests\n", 
		passed, failed, len(AllTests))
	
	if failed > 0 {
		os.Exit(1)
	}
}