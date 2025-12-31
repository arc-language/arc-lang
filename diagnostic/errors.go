package diagnostic

import (
	"fmt"
)

// Error represents a specific issue found in the source code
type Error struct {
	File    string
	Line    int
	Column  int
	Message string
	// You could add 'Severity' (Error vs Warning) here in the future
}

// String formats the error in a standard compiler output format:
// /path/to/file.arc:10:5: Variable 'x' is undefined
func (e *Error) String() string {
	return fmt.Sprintf("%s:%d:%d: %s", e.File, e.Line, e.Column, e.Message)
}

// Bag is a collection of errors.
// It allows the compiler to collect multiple errors before stopping,
// giving the user more feedback in one run.
type Bag struct {
	Errors []*Error
}

func NewBag() *Bag {
	return &Bag{
		Errors: make([]*Error, 0),
	}
}

// Report adds a new error to the bag
func (b *Bag) Report(file string, line, column int, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	
	b.Errors = append(b.Errors, &Error{
		File:    file,
		Line:    line,
		Column:  column,
		Message: msg,
	})
}

// HasErrors returns true if the bag contains any errors
func (b *Bag) HasErrors() bool {
	return len(b.Errors) > 0
}

// Clear removes all errors (useful if resetting state)
func (b *Bag) Clear() {
	b.Errors = make([]*Error, 0)
}