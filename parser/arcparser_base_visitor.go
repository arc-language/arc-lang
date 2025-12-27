// Code generated from ArcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ArcParser

import "github.com/antlr4-go/antlr/v4"

type BaseArcParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseArcParserVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitImportSpec(ctx *ImportSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitNamespaceDecl(ctx *NamespaceDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitTopLevelDecl(ctx *TopLevelDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternDecl(ctx *ExternDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternMember(ctx *ExternMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternFunctionDecl(ctx *ExternFunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternParameterList(ctx *ExternParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStructDecl(ctx *StructDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStructMember(ctx *StructMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStructField(ctx *StructFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitClassDecl(ctx *ClassDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitClassMember(ctx *ClassMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitClassField(ctx *ClassFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitMethodDecl(ctx *MethodDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitMutatingDecl(ctx *MutatingDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitDeinitDecl(ctx *DeinitDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitVariableDecl(ctx *VariableDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitConstDecl(ctx *ConstDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitPointerType(ctx *PointerTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitReferenceType(ctx *ReferenceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitVectorType(ctx *VectorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitAssignmentStmt(ctx *AssignmentStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitLeftHandSide(ctx *LeftHandSideContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExpressionStmt(ctx *ExpressionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitDeferStmt(ctx *DeferStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitRangeExpression(ctx *RangeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitPostfixOp(ctx *PostfixOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitVectorLiteral(ctx *VectorLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitMapLiteral(ctx *MapLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitMapEntry(ctx *MapEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStructLiteral(ctx *StructLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitFieldInit(ctx *FieldInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitCastExpression(ctx *CastExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitAllocaExpression(ctx *AllocaExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitSyscallExpression(ctx *SyscallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitIntrinsicExpression(ctx *IntrinsicExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
