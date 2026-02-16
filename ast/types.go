package ast

// TypeRef is the interface for type nodes.
type TypeRef interface {
	Node
	typeNode()
}

// NamedType represents "int32", "string", "net.Socket".
type NamedType struct {
	Name         string // Could include package prefix "net.Socket"
	GenericArgs  []TypeRef
	Pos          Position
}

func (t *NamedType) Pos() Position { return t.Pos }
func (t *NamedType) End() Position { return t.Pos }
func (t *NamedType) typeNode()     {}

// PointerType represents "*T" (used in externs).
type PointerType struct {
	Base TypeRef
	Start Position
}

func (t *PointerType) Pos() Position { return t.Start }
func (t *PointerType) End() Position { return t.Base.End() }
func (t *PointerType) typeNode()     {}

// ArrayType represents "vector[T]" or "[]T" or "[N]T".
type ArrayType struct {
	Kind    string // "vector", "slice", "array"
	Elem    TypeRef
	Len     Expr   // Only for fixed array [N]T
	Start   Position
}

func (t *ArrayType) Pos() Position { return t.Start }
func (t *ArrayType) End() Position { return t.Elem.End() }
func (t *ArrayType) typeNode()     {}

// MapType represents "map[K]V".
type MapType struct {
	Key   TypeRef
	Value TypeRef
	Start Position
}

func (t *MapType) Pos() Position { return t.Start }
func (t *MapType) End() Position { return t.Value.End() }
func (t *MapType) typeNode()     {}

// FuncType represents "func(int, int) bool".
type FuncType struct {
	Params  []TypeRef
	Results []TypeRef
	Start   Position
}

func (t *FuncType) Pos() Position { return t.Start }
func (t *FuncType) End() Position { return Position{} } // Simplified
func (t *FuncType) typeNode()     {}