package ast

// BinaryExpr represents "x + y".
type BinaryExpr struct {
	Left  Expr
	Op    string // +, -, *, /, &&, ||
	Right Expr
}

func (e *BinaryExpr) Pos() Position { return e.Left.Pos() }
func (e *BinaryExpr) End() Position { return e.Right.End() }
func (e *BinaryExpr) exprNode()     {}

// CallExpr represents "fn(args)".
type CallExpr struct {
	Fun  Expr
	Args []Expr
	Start Position
	EndPos Position
}

func (e *CallExpr) Pos() Position { return e.Start }
func (e *CallExpr) End() Position { return e.EndPos }
func (e *CallExpr) exprNode()     {}

// Ident represents a variable name "x".
type Ident struct {
	Name string
	Pos  Position
}

func (e *Ident) Pos() Position { return e.Pos }
func (e *Ident) End() Position { return e.Pos } // Simplified
func (e *Ident) exprNode()     {}

// BasicLit represents a literal "123", "3.14", "hello".
type BasicLit struct {
	Kind  string // INT, FLOAT, STRING, CHAR
	Value string
	Pos   Position
}

func (e *BasicLit) Pos() Position { return e.Pos }
func (e *BasicLit) End() Position { return e.Pos }
func (e *BasicLit) exprNode()     {}

// CompositeLit represents struct/array init "Point{x: 1}".
type CompositeLit struct {
	Type   TypeRef // Optional (inferred if nil)
	Fields []Expr  // KeyValueExpr or just Expr
	LBrace Position
	RBrace Position
}

func (e *CompositeLit) Pos() Position { return e.LBrace }
func (e *CompositeLit) End() Position { return e.RBrace }
func (e *CompositeLit) exprNode()     {}

// KeyValueExpr represents "key: value" in composites.
type KeyValueExpr struct {
	Key   Expr
	Value Expr
}

func (e *KeyValueExpr) Pos() Position { return e.Key.Pos() }
func (e *KeyValueExpr) End() Position { return e.Value.End() }
func (e *KeyValueExpr) exprNode()     {}

// IndexExpr represents "arr[i]".
type IndexExpr struct {
	X      Expr
	Index  Expr
	LBrack Position
	RBrack Position
}

func (e *IndexExpr) Pos() Position { return e.X.Pos() }
func (e *IndexExpr) End() Position { return e.RBrack }
func (e *IndexExpr) exprNode()     {}