package frontend

import "github.com/arc-language/arc-lang/ast"

type ScopeKind int

const (
	ScopeGlobal ScopeKind = iota
	ScopeFile
	ScopeFunc
	ScopeBlock
)

// Symbol represents a named entity (var, func, type).
type Symbol struct {
	Name string
	Kind string      // "var", "func", "type", "const"
	Decl ast.Node    // The AST node where it was defined
	Type ast.TypeRef // The resolved type
}

// Scope tracks symbols in a specific region.
type Scope struct {
	Parent  *Scope
	Kind    ScopeKind
	Symbols map[string]*Symbol
}

func NewScope(parent *Scope, kind ScopeKind) *Scope {
	return &Scope{
		Parent:  parent,
		Kind:    kind,
		Symbols: make(map[string]*Symbol),
	}
}

func (s *Scope) Insert(name string, sym *Symbol) {
	s.Symbols[name] = sym
}

func (s *Scope) Lookup(name string) *Symbol {
	if sym, ok := s.Symbols[name]; ok {
		return sym
	}
	if s.Parent != nil {
		return s.Parent.Lookup(name)
	}
	return nil
}