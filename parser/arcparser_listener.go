// Code generated from ArcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ArcParser

import "github.com/antlr4-go/antlr/v4"

// ArcParserListener is a complete listener for a parse tree produced by ArcParser.
type ArcParserListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterImportDecl is called when entering the importDecl production.
	EnterImportDecl(c *ImportDeclContext)

	// EnterImportSpec is called when entering the importSpec production.
	EnterImportSpec(c *ImportSpecContext)

	// EnterNamespaceDecl is called when entering the namespaceDecl production.
	EnterNamespaceDecl(c *NamespaceDeclContext)

	// EnterTopLevelDecl is called when entering the topLevelDecl production.
	EnterTopLevelDecl(c *TopLevelDeclContext)

	// EnterExternDecl is called when entering the externDecl production.
	EnterExternDecl(c *ExternDeclContext)

	// EnterExternMember is called when entering the externMember production.
	EnterExternMember(c *ExternMemberContext)

	// EnterExternFunctionDecl is called when entering the externFunctionDecl production.
	EnterExternFunctionDecl(c *ExternFunctionDeclContext)

	// EnterExternParameterList is called when entering the externParameterList production.
	EnterExternParameterList(c *ExternParameterListContext)

	// EnterFunctionDecl is called when entering the functionDecl production.
	EnterFunctionDecl(c *FunctionDeclContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterStructDecl is called when entering the structDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterStructMember is called when entering the structMember production.
	EnterStructMember(c *StructMemberContext)

	// EnterStructField is called when entering the structField production.
	EnterStructField(c *StructFieldContext)

	// EnterClassDecl is called when entering the classDecl production.
	EnterClassDecl(c *ClassDeclContext)

	// EnterClassMember is called when entering the classMember production.
	EnterClassMember(c *ClassMemberContext)

	// EnterClassField is called when entering the classField production.
	EnterClassField(c *ClassFieldContext)

	// EnterMethodDecl is called when entering the methodDecl production.
	EnterMethodDecl(c *MethodDeclContext)

	// EnterMutatingDecl is called when entering the mutatingDecl production.
	EnterMutatingDecl(c *MutatingDeclContext)

	// EnterDeinitDecl is called when entering the deinitDecl production.
	EnterDeinitDecl(c *DeinitDeclContext)

	// EnterVariableDecl is called when entering the variableDecl production.
	EnterVariableDecl(c *VariableDeclContext)

	// EnterConstDecl is called when entering the constDecl production.
	EnterConstDecl(c *ConstDeclContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterPointerType is called when entering the pointerType production.
	EnterPointerType(c *PointerTypeContext)

	// EnterReferenceType is called when entering the referenceType production.
	EnterReferenceType(c *ReferenceTypeContext)

	// EnterVectorType is called when entering the vectorType production.
	EnterVectorType(c *VectorTypeContext)

	// EnterMapType is called when entering the mapType production.
	EnterMapType(c *MapTypeContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignmentStmt is called when entering the assignmentStmt production.
	EnterAssignmentStmt(c *AssignmentStmtContext)

	// EnterLeftHandSide is called when entering the leftHandSide production.
	EnterLeftHandSide(c *LeftHandSideContext)

	// EnterExpressionStmt is called when entering the expressionStmt production.
	EnterExpressionStmt(c *ExpressionStmtContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterForStmt is called when entering the forStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterBreakStmt is called when entering the breakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the continueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterDeferStmt is called when entering the deferStmt production.
	EnterDeferStmt(c *DeferStmtContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLogicalOrExpression is called when entering the logicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterLogicalAndExpression is called when entering the logicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterRangeExpression is called when entering the rangeExpression production.
	EnterRangeExpression(c *RangeExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterPostfixOp is called when entering the postfixOp production.
	EnterPostfixOp(c *PostfixOpContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterVectorLiteral is called when entering the vectorLiteral production.
	EnterVectorLiteral(c *VectorLiteralContext)

	// EnterMapLiteral is called when entering the mapLiteral production.
	EnterMapLiteral(c *MapLiteralContext)

	// EnterMapEntry is called when entering the mapEntry production.
	EnterMapEntry(c *MapEntryContext)

	// EnterStructLiteral is called when entering the structLiteral production.
	EnterStructLiteral(c *StructLiteralContext)

	// EnterFieldInit is called when entering the fieldInit production.
	EnterFieldInit(c *FieldInitContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterCastExpression is called when entering the castExpression production.
	EnterCastExpression(c *CastExpressionContext)

	// EnterAllocaExpression is called when entering the allocaExpression production.
	EnterAllocaExpression(c *AllocaExpressionContext)

	// EnterSyscallExpression is called when entering the syscallExpression production.
	EnterSyscallExpression(c *SyscallExpressionContext)

	// EnterIntrinsicExpression is called when entering the intrinsicExpression production.
	EnterIntrinsicExpression(c *IntrinsicExpressionContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitImportDecl is called when exiting the importDecl production.
	ExitImportDecl(c *ImportDeclContext)

	// ExitImportSpec is called when exiting the importSpec production.
	ExitImportSpec(c *ImportSpecContext)

	// ExitNamespaceDecl is called when exiting the namespaceDecl production.
	ExitNamespaceDecl(c *NamespaceDeclContext)

	// ExitTopLevelDecl is called when exiting the topLevelDecl production.
	ExitTopLevelDecl(c *TopLevelDeclContext)

	// ExitExternDecl is called when exiting the externDecl production.
	ExitExternDecl(c *ExternDeclContext)

	// ExitExternMember is called when exiting the externMember production.
	ExitExternMember(c *ExternMemberContext)

	// ExitExternFunctionDecl is called when exiting the externFunctionDecl production.
	ExitExternFunctionDecl(c *ExternFunctionDeclContext)

	// ExitExternParameterList is called when exiting the externParameterList production.
	ExitExternParameterList(c *ExternParameterListContext)

	// ExitFunctionDecl is called when exiting the functionDecl production.
	ExitFunctionDecl(c *FunctionDeclContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitStructDecl is called when exiting the structDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitStructMember is called when exiting the structMember production.
	ExitStructMember(c *StructMemberContext)

	// ExitStructField is called when exiting the structField production.
	ExitStructField(c *StructFieldContext)

	// ExitClassDecl is called when exiting the classDecl production.
	ExitClassDecl(c *ClassDeclContext)

	// ExitClassMember is called when exiting the classMember production.
	ExitClassMember(c *ClassMemberContext)

	// ExitClassField is called when exiting the classField production.
	ExitClassField(c *ClassFieldContext)

	// ExitMethodDecl is called when exiting the methodDecl production.
	ExitMethodDecl(c *MethodDeclContext)

	// ExitMutatingDecl is called when exiting the mutatingDecl production.
	ExitMutatingDecl(c *MutatingDeclContext)

	// ExitDeinitDecl is called when exiting the deinitDecl production.
	ExitDeinitDecl(c *DeinitDeclContext)

	// ExitVariableDecl is called when exiting the variableDecl production.
	ExitVariableDecl(c *VariableDeclContext)

	// ExitConstDecl is called when exiting the constDecl production.
	ExitConstDecl(c *ConstDeclContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitPointerType is called when exiting the pointerType production.
	ExitPointerType(c *PointerTypeContext)

	// ExitReferenceType is called when exiting the referenceType production.
	ExitReferenceType(c *ReferenceTypeContext)

	// ExitVectorType is called when exiting the vectorType production.
	ExitVectorType(c *VectorTypeContext)

	// ExitMapType is called when exiting the mapType production.
	ExitMapType(c *MapTypeContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignmentStmt is called when exiting the assignmentStmt production.
	ExitAssignmentStmt(c *AssignmentStmtContext)

	// ExitLeftHandSide is called when exiting the leftHandSide production.
	ExitLeftHandSide(c *LeftHandSideContext)

	// ExitExpressionStmt is called when exiting the expressionStmt production.
	ExitExpressionStmt(c *ExpressionStmtContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitForStmt is called when exiting the forStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitBreakStmt is called when exiting the breakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the continueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitDeferStmt is called when exiting the deferStmt production.
	ExitDeferStmt(c *DeferStmtContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLogicalOrExpression is called when exiting the logicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitLogicalAndExpression is called when exiting the logicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitRangeExpression is called when exiting the rangeExpression production.
	ExitRangeExpression(c *RangeExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitPostfixOp is called when exiting the postfixOp production.
	ExitPostfixOp(c *PostfixOpContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitVectorLiteral is called when exiting the vectorLiteral production.
	ExitVectorLiteral(c *VectorLiteralContext)

	// ExitMapLiteral is called when exiting the mapLiteral production.
	ExitMapLiteral(c *MapLiteralContext)

	// ExitMapEntry is called when exiting the mapEntry production.
	ExitMapEntry(c *MapEntryContext)

	// ExitStructLiteral is called when exiting the structLiteral production.
	ExitStructLiteral(c *StructLiteralContext)

	// ExitFieldInit is called when exiting the fieldInit production.
	ExitFieldInit(c *FieldInitContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitCastExpression is called when exiting the castExpression production.
	ExitCastExpression(c *CastExpressionContext)

	// ExitAllocaExpression is called when exiting the allocaExpression production.
	ExitAllocaExpression(c *AllocaExpressionContext)

	// ExitSyscallExpression is called when exiting the syscallExpression production.
	ExitSyscallExpression(c *SyscallExpressionContext)

	// ExitIntrinsicExpression is called when exiting the intrinsicExpression production.
	ExitIntrinsicExpression(c *IntrinsicExpressionContext)
}
