package ast

// BlockStmt represents "{ ... }".
type BlockStmt struct {
	List   []Stmt
	LBrace Position
	RBrace Position
}

func (s *BlockStmt) Pos() Position { return s.LBrace }
func (s *BlockStmt) End() Position { return s.RBrace }
func (s *BlockStmt) stmtNode()     {}

// ReturnStmt represents "return expr".
type ReturnStmt struct {
	Results []Expr
	Start   Position
}

func (s *ReturnStmt) Pos() Position { return s.Start }
func (s *ReturnStmt) End() Position { return s.Start }
func (s *ReturnStmt) stmtNode()     {}

// IfStmt represents "if Cond { Body } else { Else }".
type IfStmt struct {
	Cond  Expr
	Body  *BlockStmt
	Else  Stmt // Can be BlockStmt or IfStmt (else if)
	Start Position
}

func (s *IfStmt) Pos() Position { return s.Start }
func (s *IfStmt) End() Position { return s.Body.End() }
func (s *IfStmt) stmtNode()     {}

// ForStmt handles all loops (C-style, iterator, infinite).
type ForStmt struct {
	Init  Stmt // let i = 0
	Cond  Expr // i < 10
	Post  Stmt // i++
	Body  *BlockStmt
	Start Position
}

func (s *ForStmt) Pos() Position { return s.Start }
func (s *ForStmt) End() Position { return s.Body.End() }
func (s *ForStmt) stmtNode()     {}

// DeferStmt represents "defer call()".
type DeferStmt struct {
	Call  *CallExpr
	Start Position
}

func (s *DeferStmt) Pos() Position { return s.Start }
func (s *DeferStmt) End() Position { return s.Call.End() }
func (s *DeferStmt) stmtNode()     {}

// ExprStmt represents a standalone expression (like a function call).
type ExprStmt struct {
	X Expr
}

func (s *ExprStmt) Pos() Position { return s.X.Pos() }
func (s *ExprStmt) End() Position { return s.X.End() }
func (s *ExprStmt) stmtNode()     {}

// DeclStmt wrappers a VarDecl when it appears inside a function body.
type DeclStmt struct {
	Decl *VarDecl
}

func (s *DeclStmt) Pos() Position { return s.Decl.Pos() }
func (s *DeclStmt) End() Position { return s.Decl.End() }
func (s *DeclStmt) stmtNode()     {}