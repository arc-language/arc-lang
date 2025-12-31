package symbol

import (
	"fmt"
	
	"github.com/arc-language/arc-lang/builder/types"
)

// Scope represents a lexical region (Block, Function, Global) that holds symbols.
// It forms a tree structure where children point to their parent.
type Scope struct {
	Parent  *Scope
	Symbols map[string]*Symbol
	
	// Optional: Name for debugging (e.g., "func main", "if.then")
	DebugName string
}

// NewScope creates a child scope. Pass nil for a Global scope.
func NewScope(parent *Scope) *Scope {
	return &Scope{
		Parent:  parent,
		Symbols: make(map[string]*Symbol),
	}
}

// Define creates a new symbol in the *current* scope.
// It does NOT check for shadowing (Pass 1 logic handles that check).
func (s *Scope) Define(name string, kind SymbolKind, typ types.Type) *Symbol {
	sym := &Symbol{
		Name:      name,
		Kind:      kind,
		Type:      typ,
		IsDefined: true,
	}
	s.Symbols[name] = sym
	return sym
}

// Resolve looks up a symbol starting from current scope up to global.
func (s *Scope) Resolve(name string) (*Symbol, bool) {
	// 1. Check current scope
	if sym, ok := s.Symbols[name]; ok {
		return sym, true
	}
	
	// 2. Check parent (if exists)
	if s.Parent != nil {
		return s.Parent.Resolve(name)
	}
	
	// 3. Not found
	return nil, false
}

// ResolveLocal looks up a symbol *only* in the current scope.
// Used by Pass 1 to detect "Redeclaration of variable 'x'" errors.
func (s *Scope) ResolveLocal(name string) (*Symbol, bool) {
	sym, ok := s.Symbols[name]
	return sym, ok
}

// DebugPrint prints the scope hierarchy to stdout (helper for compiler devs)
func (s *Scope) DebugPrint(indent int) {
	pad := ""
	for i := 0; i < indent; i++ {
		pad += "  "
	}
	
	name := "Scope"
	if s.DebugName != "" {
		name += " (" + s.DebugName + ")"
	}
	
	fmt.Printf("%s%s {\n", pad, name)
	for _, sym := range s.Symbols {
		fmt.Printf("%s  %s\n", pad, sym.String())
	}
	fmt.Printf("%s}\n", pad)
}