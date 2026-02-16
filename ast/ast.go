package ast

// Position represents a location in the source file.
type Position struct {
	Line   int
	Column int
}

// Node is the base interface for all AST nodes.
type Node interface {
	Pos() Position
	End() Position
}

// Expr is the interface for all expression nodes.
type Expr interface {
	Node
	exprNode()
}

// Stmt is the interface for all statement nodes.
type Stmt interface {
	Node
	stmtNode()
}

// Decl is the interface for all declaration nodes.
type Decl interface {
	Node
	declNode()
}