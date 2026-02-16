package ast

// File represents a complete Arc source file.
type File struct {
	Namespace string
	Imports   []*ImportDecl
	Decls     []Decl
}

func (f *File) Pos() Position { return Position{1, 1} }
func (f *File) End() Position { return Position{} } // Calculate from last decl

// ImportDecl represents "import foo 'path'".
type ImportDecl struct {
	Alias string // "_" or name or empty
	Path  string
	Start Position
}

func (d *ImportDecl) Pos() Position { return d.Start }
func (d *ImportDecl) End() Position { return d.Start } // Simplified
func (d *ImportDecl) declNode()     {}

// FuncDecl represents "async func name[T](...)".
type FuncDecl struct {
	Name          string
	IsAsync       bool
	IsGpu         bool
	GenericParams []*Field    // [T, U]
	Params        []*Field    // (a: int32, ...)
	ReturnType    TypeRef     // Optional
	Body          *BlockStmt  // nil if extern
	Start         Position
}

func (d *FuncDecl) Pos() Position { return d.Start }
func (d *FuncDecl) End() Position { return d.Body.End() }
func (d *FuncDecl) declNode()     {}

// VarDecl represents "var x: T = val" or "let x = val".
type VarDecl struct {
	Name     string
	IsRef    bool // true if 'var', false if 'let'
	Type     TypeRef
	Value    Expr
	IsExtern bool // true if defined in 'extern' block
	Start    Position
}

func (d *VarDecl) Pos() Position { return d.Start }
func (d *VarDecl) End() Position { return d.Start } // Simplified
func (d *VarDecl) declNode()     {}

// InterfaceDecl represents "interface Point { ... }".
type InterfaceDecl struct {
	Name          string
	GenericParams []*Field
	Methods       []*FuncDecl // Methods attached to this type
	Fields        []*Field
	Start         Position
}

func (d *InterfaceDecl) Pos() Position { return d.Start }
func (d *InterfaceDecl) End() Position { return d.Start } // Simplified
func (d *InterfaceDecl) declNode()     {}

// Field represents a name:type pair (used in params, structs, generics).
type Field struct {
	Name  string
	Type  TypeRef
	Start Position
}