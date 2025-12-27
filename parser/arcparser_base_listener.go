// Code generated from ArcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ArcParser

import "github.com/antlr4-go/antlr/v4"

// BaseArcParserListener is a complete listener for a parse tree produced by ArcParser.
type BaseArcParserListener struct{}

var _ ArcParserListener = &BaseArcParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseArcParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseArcParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseArcParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseArcParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseArcParserListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseArcParserListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BaseArcParserListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BaseArcParserListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *BaseArcParserListener) EnterImportSpec(ctx *ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *BaseArcParserListener) ExitImportSpec(ctx *ImportSpecContext) {}

// EnterNamespaceDecl is called when production namespaceDecl is entered.
func (s *BaseArcParserListener) EnterNamespaceDecl(ctx *NamespaceDeclContext) {}

// ExitNamespaceDecl is called when production namespaceDecl is exited.
func (s *BaseArcParserListener) ExitNamespaceDecl(ctx *NamespaceDeclContext) {}

// EnterTopLevelDecl is called when production topLevelDecl is entered.
func (s *BaseArcParserListener) EnterTopLevelDecl(ctx *TopLevelDeclContext) {}

// ExitTopLevelDecl is called when production topLevelDecl is exited.
func (s *BaseArcParserListener) ExitTopLevelDecl(ctx *TopLevelDeclContext) {}

// EnterExternDecl is called when production externDecl is entered.
func (s *BaseArcParserListener) EnterExternDecl(ctx *ExternDeclContext) {}

// ExitExternDecl is called when production externDecl is exited.
func (s *BaseArcParserListener) ExitExternDecl(ctx *ExternDeclContext) {}

// EnterExternMember is called when production externMember is entered.
func (s *BaseArcParserListener) EnterExternMember(ctx *ExternMemberContext) {}

// ExitExternMember is called when production externMember is exited.
func (s *BaseArcParserListener) ExitExternMember(ctx *ExternMemberContext) {}

// EnterExternFunctionDecl is called when production externFunctionDecl is entered.
func (s *BaseArcParserListener) EnterExternFunctionDecl(ctx *ExternFunctionDeclContext) {}

// ExitExternFunctionDecl is called when production externFunctionDecl is exited.
func (s *BaseArcParserListener) ExitExternFunctionDecl(ctx *ExternFunctionDeclContext) {}

// EnterExternParameterList is called when production externParameterList is entered.
func (s *BaseArcParserListener) EnterExternParameterList(ctx *ExternParameterListContext) {}

// ExitExternParameterList is called when production externParameterList is exited.
func (s *BaseArcParserListener) ExitExternParameterList(ctx *ExternParameterListContext) {}

// EnterFunctionDecl is called when production functionDecl is entered.
func (s *BaseArcParserListener) EnterFunctionDecl(ctx *FunctionDeclContext) {}

// ExitFunctionDecl is called when production functionDecl is exited.
func (s *BaseArcParserListener) ExitFunctionDecl(ctx *FunctionDeclContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseArcParserListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseArcParserListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseArcParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseArcParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterStructDecl is called when production structDecl is entered.
func (s *BaseArcParserListener) EnterStructDecl(ctx *StructDeclContext) {}

// ExitStructDecl is called when production structDecl is exited.
func (s *BaseArcParserListener) ExitStructDecl(ctx *StructDeclContext) {}

// EnterStructMember is called when production structMember is entered.
func (s *BaseArcParserListener) EnterStructMember(ctx *StructMemberContext) {}

// ExitStructMember is called when production structMember is exited.
func (s *BaseArcParserListener) ExitStructMember(ctx *StructMemberContext) {}

// EnterStructField is called when production structField is entered.
func (s *BaseArcParserListener) EnterStructField(ctx *StructFieldContext) {}

// ExitStructField is called when production structField is exited.
func (s *BaseArcParserListener) ExitStructField(ctx *StructFieldContext) {}

// EnterClassDecl is called when production classDecl is entered.
func (s *BaseArcParserListener) EnterClassDecl(ctx *ClassDeclContext) {}

// ExitClassDecl is called when production classDecl is exited.
func (s *BaseArcParserListener) ExitClassDecl(ctx *ClassDeclContext) {}

// EnterClassMember is called when production classMember is entered.
func (s *BaseArcParserListener) EnterClassMember(ctx *ClassMemberContext) {}

// ExitClassMember is called when production classMember is exited.
func (s *BaseArcParserListener) ExitClassMember(ctx *ClassMemberContext) {}

// EnterClassField is called when production classField is entered.
func (s *BaseArcParserListener) EnterClassField(ctx *ClassFieldContext) {}

// ExitClassField is called when production classField is exited.
func (s *BaseArcParserListener) ExitClassField(ctx *ClassFieldContext) {}

// EnterMethodDecl is called when production methodDecl is entered.
func (s *BaseArcParserListener) EnterMethodDecl(ctx *MethodDeclContext) {}

// ExitMethodDecl is called when production methodDecl is exited.
func (s *BaseArcParserListener) ExitMethodDecl(ctx *MethodDeclContext) {}

// EnterMutatingDecl is called when production mutatingDecl is entered.
func (s *BaseArcParserListener) EnterMutatingDecl(ctx *MutatingDeclContext) {}

// ExitMutatingDecl is called when production mutatingDecl is exited.
func (s *BaseArcParserListener) ExitMutatingDecl(ctx *MutatingDeclContext) {}

// EnterDeinitDecl is called when production deinitDecl is entered.
func (s *BaseArcParserListener) EnterDeinitDecl(ctx *DeinitDeclContext) {}

// ExitDeinitDecl is called when production deinitDecl is exited.
func (s *BaseArcParserListener) ExitDeinitDecl(ctx *DeinitDeclContext) {}

// EnterVariableDecl is called when production variableDecl is entered.
func (s *BaseArcParserListener) EnterVariableDecl(ctx *VariableDeclContext) {}

// ExitVariableDecl is called when production variableDecl is exited.
func (s *BaseArcParserListener) ExitVariableDecl(ctx *VariableDeclContext) {}

// EnterConstDecl is called when production constDecl is entered.
func (s *BaseArcParserListener) EnterConstDecl(ctx *ConstDeclContext) {}

// ExitConstDecl is called when production constDecl is exited.
func (s *BaseArcParserListener) ExitConstDecl(ctx *ConstDeclContext) {}

// EnterType is called when production type is entered.
func (s *BaseArcParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseArcParserListener) ExitType(ctx *TypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseArcParserListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseArcParserListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterPointerType is called when production pointerType is entered.
func (s *BaseArcParserListener) EnterPointerType(ctx *PointerTypeContext) {}

// ExitPointerType is called when production pointerType is exited.
func (s *BaseArcParserListener) ExitPointerType(ctx *PointerTypeContext) {}

// EnterReferenceType is called when production referenceType is entered.
func (s *BaseArcParserListener) EnterReferenceType(ctx *ReferenceTypeContext) {}

// ExitReferenceType is called when production referenceType is exited.
func (s *BaseArcParserListener) ExitReferenceType(ctx *ReferenceTypeContext) {}

// EnterVectorType is called when production vectorType is entered.
func (s *BaseArcParserListener) EnterVectorType(ctx *VectorTypeContext) {}

// ExitVectorType is called when production vectorType is exited.
func (s *BaseArcParserListener) ExitVectorType(ctx *VectorTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseArcParserListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseArcParserListener) ExitMapType(ctx *MapTypeContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseArcParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseArcParserListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseArcParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseArcParserListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignmentStmt is called when production assignmentStmt is entered.
func (s *BaseArcParserListener) EnterAssignmentStmt(ctx *AssignmentStmtContext) {}

// ExitAssignmentStmt is called when production assignmentStmt is exited.
func (s *BaseArcParserListener) ExitAssignmentStmt(ctx *AssignmentStmtContext) {}

// EnterLeftHandSide is called when production leftHandSide is entered.
func (s *BaseArcParserListener) EnterLeftHandSide(ctx *LeftHandSideContext) {}

// ExitLeftHandSide is called when production leftHandSide is exited.
func (s *BaseArcParserListener) ExitLeftHandSide(ctx *LeftHandSideContext) {}

// EnterExpressionStmt is called when production expressionStmt is entered.
func (s *BaseArcParserListener) EnterExpressionStmt(ctx *ExpressionStmtContext) {}

// ExitExpressionStmt is called when production expressionStmt is exited.
func (s *BaseArcParserListener) ExitExpressionStmt(ctx *ExpressionStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseArcParserListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseArcParserListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseArcParserListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseArcParserListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseArcParserListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseArcParserListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseArcParserListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseArcParserListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BaseArcParserListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BaseArcParserListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterDeferStmt is called when production deferStmt is entered.
func (s *BaseArcParserListener) EnterDeferStmt(ctx *DeferStmtContext) {}

// ExitDeferStmt is called when production deferStmt is exited.
func (s *BaseArcParserListener) ExitDeferStmt(ctx *DeferStmtContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseArcParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseArcParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BaseArcParserListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BaseArcParserListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BaseArcParserListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BaseArcParserListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseArcParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseArcParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseArcParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseArcParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterRangeExpression is called when production rangeExpression is entered.
func (s *BaseArcParserListener) EnterRangeExpression(ctx *RangeExpressionContext) {}

// ExitRangeExpression is called when production rangeExpression is exited.
func (s *BaseArcParserListener) ExitRangeExpression(ctx *RangeExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseArcParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseArcParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseArcParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseArcParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseArcParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseArcParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseArcParserListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseArcParserListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterPostfixOp is called when production postfixOp is entered.
func (s *BaseArcParserListener) EnterPostfixOp(ctx *PostfixOpContext) {}

// ExitPostfixOp is called when production postfixOp is exited.
func (s *BaseArcParserListener) ExitPostfixOp(ctx *PostfixOpContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseArcParserListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseArcParserListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseArcParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseArcParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterVectorLiteral is called when production vectorLiteral is entered.
func (s *BaseArcParserListener) EnterVectorLiteral(ctx *VectorLiteralContext) {}

// ExitVectorLiteral is called when production vectorLiteral is exited.
func (s *BaseArcParserListener) ExitVectorLiteral(ctx *VectorLiteralContext) {}

// EnterMapLiteral is called when production mapLiteral is entered.
func (s *BaseArcParserListener) EnterMapLiteral(ctx *MapLiteralContext) {}

// ExitMapLiteral is called when production mapLiteral is exited.
func (s *BaseArcParserListener) ExitMapLiteral(ctx *MapLiteralContext) {}

// EnterMapEntry is called when production mapEntry is entered.
func (s *BaseArcParserListener) EnterMapEntry(ctx *MapEntryContext) {}

// ExitMapEntry is called when production mapEntry is exited.
func (s *BaseArcParserListener) ExitMapEntry(ctx *MapEntryContext) {}

// EnterStructLiteral is called when production structLiteral is entered.
func (s *BaseArcParserListener) EnterStructLiteral(ctx *StructLiteralContext) {}

// ExitStructLiteral is called when production structLiteral is exited.
func (s *BaseArcParserListener) ExitStructLiteral(ctx *StructLiteralContext) {}

// EnterFieldInit is called when production fieldInit is entered.
func (s *BaseArcParserListener) EnterFieldInit(ctx *FieldInitContext) {}

// ExitFieldInit is called when production fieldInit is exited.
func (s *BaseArcParserListener) ExitFieldInit(ctx *FieldInitContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseArcParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseArcParserListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *BaseArcParserListener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *BaseArcParserListener) ExitCastExpression(ctx *CastExpressionContext) {}

// EnterAllocaExpression is called when production allocaExpression is entered.
func (s *BaseArcParserListener) EnterAllocaExpression(ctx *AllocaExpressionContext) {}

// ExitAllocaExpression is called when production allocaExpression is exited.
func (s *BaseArcParserListener) ExitAllocaExpression(ctx *AllocaExpressionContext) {}

// EnterSyscallExpression is called when production syscallExpression is entered.
func (s *BaseArcParserListener) EnterSyscallExpression(ctx *SyscallExpressionContext) {}

// ExitSyscallExpression is called when production syscallExpression is exited.
func (s *BaseArcParserListener) ExitSyscallExpression(ctx *SyscallExpressionContext) {}

// EnterIntrinsicExpression is called when production intrinsicExpression is entered.
func (s *BaseArcParserListener) EnterIntrinsicExpression(ctx *IntrinsicExpressionContext) {}

// ExitIntrinsicExpression is called when production intrinsicExpression is exited.
func (s *BaseArcParserListener) ExitIntrinsicExpression(ctx *IntrinsicExpressionContext) {}
