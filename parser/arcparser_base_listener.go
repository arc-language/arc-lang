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

// EnterGenericParams is called when production genericParams is entered.
func (s *BaseArcParserListener) EnterGenericParams(ctx *GenericParamsContext) {}

// ExitGenericParams is called when production genericParams is exited.
func (s *BaseArcParserListener) ExitGenericParams(ctx *GenericParamsContext) {}

// EnterGenericParamList is called when production genericParamList is entered.
func (s *BaseArcParserListener) EnterGenericParamList(ctx *GenericParamListContext) {}

// ExitGenericParamList is called when production genericParamList is exited.
func (s *BaseArcParserListener) ExitGenericParamList(ctx *GenericParamListContext) {}

// EnterGenericArgs is called when production genericArgs is entered.
func (s *BaseArcParserListener) EnterGenericArgs(ctx *GenericArgsContext) {}

// ExitGenericArgs is called when production genericArgs is exited.
func (s *BaseArcParserListener) ExitGenericArgs(ctx *GenericArgsContext) {}

// EnterGenericArgList is called when production genericArgList is entered.
func (s *BaseArcParserListener) EnterGenericArgList(ctx *GenericArgListContext) {}

// ExitGenericArgList is called when production genericArgList is exited.
func (s *BaseArcParserListener) ExitGenericArgList(ctx *GenericArgListContext) {}

// EnterGenericArg is called when production genericArg is entered.
func (s *BaseArcParserListener) EnterGenericArg(ctx *GenericArgContext) {}

// ExitGenericArg is called when production genericArg is exited.
func (s *BaseArcParserListener) ExitGenericArg(ctx *GenericArgContext) {}

// EnterFunctionDecl is called when production functionDecl is entered.
func (s *BaseArcParserListener) EnterFunctionDecl(ctx *FunctionDeclContext) {}

// ExitFunctionDecl is called when production functionDecl is exited.
func (s *BaseArcParserListener) ExitFunctionDecl(ctx *FunctionDeclContext) {}

// EnterReturnType is called when production returnType is entered.
func (s *BaseArcParserListener) EnterReturnType(ctx *ReturnTypeContext) {}

// ExitReturnType is called when production returnType is exited.
func (s *BaseArcParserListener) ExitReturnType(ctx *ReturnTypeContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseArcParserListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseArcParserListener) ExitTypeList(ctx *TypeListContext) {}

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

// EnterEnumDecl is called when production enumDecl is entered.
func (s *BaseArcParserListener) EnterEnumDecl(ctx *EnumDeclContext) {}

// ExitEnumDecl is called when production enumDecl is exited.
func (s *BaseArcParserListener) ExitEnumDecl(ctx *EnumDeclContext) {}

// EnterEnumMember is called when production enumMember is entered.
func (s *BaseArcParserListener) EnterEnumMember(ctx *EnumMemberContext) {}

// ExitEnumMember is called when production enumMember is exited.
func (s *BaseArcParserListener) ExitEnumMember(ctx *EnumMemberContext) {}

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

// EnterTuplePattern is called when production tuplePattern is entered.
func (s *BaseArcParserListener) EnterTuplePattern(ctx *TuplePatternContext) {}

// ExitTuplePattern is called when production tuplePattern is exited.
func (s *BaseArcParserListener) ExitTuplePattern(ctx *TuplePatternContext) {}

// EnterTupleType is called when production tupleType is entered.
func (s *BaseArcParserListener) EnterTupleType(ctx *TupleTypeContext) {}

// ExitTupleType is called when production tupleType is exited.
func (s *BaseArcParserListener) ExitTupleType(ctx *TupleTypeContext) {}

// EnterType is called when production type is entered.
func (s *BaseArcParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseArcParserListener) ExitType(ctx *TypeContext) {}

// EnterQualifiedType is called when production qualifiedType is entered.
func (s *BaseArcParserListener) EnterQualifiedType(ctx *QualifiedTypeContext) {}

// ExitQualifiedType is called when production qualifiedType is exited.
func (s *BaseArcParserListener) ExitQualifiedType(ctx *QualifiedTypeContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BaseArcParserListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BaseArcParserListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterQualifiedIdentifier is called when production qualifiedIdentifier is entered.
func (s *BaseArcParserListener) EnterQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// ExitQualifiedIdentifier is called when production qualifiedIdentifier is exited.
func (s *BaseArcParserListener) ExitQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

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

// EnterAssignmentOp is called when production assignmentOp is entered.
func (s *BaseArcParserListener) EnterAssignmentOp(ctx *AssignmentOpContext) {}

// ExitAssignmentOp is called when production assignmentOp is exited.
func (s *BaseArcParserListener) ExitAssignmentOp(ctx *AssignmentOpContext) {}

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

// EnterDeferStmt is called when production deferStmt is entered.
func (s *BaseArcParserListener) EnterDeferStmt(ctx *DeferStmtContext) {}

// ExitDeferStmt is called when production deferStmt is exited.
func (s *BaseArcParserListener) ExitDeferStmt(ctx *DeferStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseArcParserListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseArcParserListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BaseArcParserListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BaseArcParserListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterThrowStmt is called when production throwStmt is entered.
func (s *BaseArcParserListener) EnterThrowStmt(ctx *ThrowStmtContext) {}

// ExitThrowStmt is called when production throwStmt is exited.
func (s *BaseArcParserListener) ExitThrowStmt(ctx *ThrowStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseArcParserListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseArcParserListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseArcParserListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseArcParserListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterSwitchStmt is called when production switchStmt is entered.
func (s *BaseArcParserListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production switchStmt is exited.
func (s *BaseArcParserListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterSwitchCase is called when production switchCase is entered.
func (s *BaseArcParserListener) EnterSwitchCase(ctx *SwitchCaseContext) {}

// ExitSwitchCase is called when production switchCase is exited.
func (s *BaseArcParserListener) ExitSwitchCase(ctx *SwitchCaseContext) {}

// EnterDefaultCase is called when production defaultCase is entered.
func (s *BaseArcParserListener) EnterDefaultCase(ctx *DefaultCaseContext) {}

// ExitDefaultCase is called when production defaultCase is exited.
func (s *BaseArcParserListener) ExitDefaultCase(ctx *DefaultCaseContext) {}

// EnterTryStmt is called when production tryStmt is entered.
func (s *BaseArcParserListener) EnterTryStmt(ctx *TryStmtContext) {}

// ExitTryStmt is called when production tryStmt is exited.
func (s *BaseArcParserListener) ExitTryStmt(ctx *TryStmtContext) {}

// EnterExceptClause is called when production exceptClause is entered.
func (s *BaseArcParserListener) EnterExceptClause(ctx *ExceptClauseContext) {}

// ExitExceptClause is called when production exceptClause is exited.
func (s *BaseArcParserListener) ExitExceptClause(ctx *ExceptClauseContext) {}

// EnterFinallyClause is called when production finallyClause is entered.
func (s *BaseArcParserListener) EnterFinallyClause(ctx *FinallyClauseContext) {}

// ExitFinallyClause is called when production finallyClause is exited.
func (s *BaseArcParserListener) ExitFinallyClause(ctx *FinallyClauseContext) {}

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

// EnterBitOrExpression is called when production bitOrExpression is entered.
func (s *BaseArcParserListener) EnterBitOrExpression(ctx *BitOrExpressionContext) {}

// ExitBitOrExpression is called when production bitOrExpression is exited.
func (s *BaseArcParserListener) ExitBitOrExpression(ctx *BitOrExpressionContext) {}

// EnterBitXorExpression is called when production bitXorExpression is entered.
func (s *BaseArcParserListener) EnterBitXorExpression(ctx *BitXorExpressionContext) {}

// ExitBitXorExpression is called when production bitXorExpression is exited.
func (s *BaseArcParserListener) ExitBitXorExpression(ctx *BitXorExpressionContext) {}

// EnterBitAndExpression is called when production bitAndExpression is entered.
func (s *BaseArcParserListener) EnterBitAndExpression(ctx *BitAndExpressionContext) {}

// ExitBitAndExpression is called when production bitAndExpression is exited.
func (s *BaseArcParserListener) ExitBitAndExpression(ctx *BitAndExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseArcParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseArcParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseArcParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseArcParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *BaseArcParserListener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *BaseArcParserListener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

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

// EnterSizeofExpression is called when production sizeofExpression is entered.
func (s *BaseArcParserListener) EnterSizeofExpression(ctx *SizeofExpressionContext) {}

// ExitSizeofExpression is called when production sizeofExpression is exited.
func (s *BaseArcParserListener) ExitSizeofExpression(ctx *SizeofExpressionContext) {}

// EnterAlignofExpression is called when production alignofExpression is entered.
func (s *BaseArcParserListener) EnterAlignofExpression(ctx *AlignofExpressionContext) {}

// ExitAlignofExpression is called when production alignofExpression is exited.
func (s *BaseArcParserListener) ExitAlignofExpression(ctx *AlignofExpressionContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseArcParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseArcParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterInitializerList is called when production initializerList is entered.
func (s *BaseArcParserListener) EnterInitializerList(ctx *InitializerListContext) {}

// ExitInitializerList is called when production initializerList is exited.
func (s *BaseArcParserListener) ExitInitializerList(ctx *InitializerListContext) {}

// EnterInitializerEntry is called when production initializerEntry is entered.
func (s *BaseArcParserListener) EnterInitializerEntry(ctx *InitializerEntryContext) {}

// ExitInitializerEntry is called when production initializerEntry is exited.
func (s *BaseArcParserListener) ExitInitializerEntry(ctx *InitializerEntryContext) {}

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

// EnterArgument is called when production argument is entered.
func (s *BaseArcParserListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseArcParserListener) ExitArgument(ctx *ArgumentContext) {}

// EnterLambdaExpression is called when production lambdaExpression is entered.
func (s *BaseArcParserListener) EnterLambdaExpression(ctx *LambdaExpressionContext) {}

// ExitLambdaExpression is called when production lambdaExpression is exited.
func (s *BaseArcParserListener) ExitLambdaExpression(ctx *LambdaExpressionContext) {}

// EnterAnonymousFuncExpression is called when production anonymousFuncExpression is entered.
func (s *BaseArcParserListener) EnterAnonymousFuncExpression(ctx *AnonymousFuncExpressionContext) {}

// ExitAnonymousFuncExpression is called when production anonymousFuncExpression is exited.
func (s *BaseArcParserListener) ExitAnonymousFuncExpression(ctx *AnonymousFuncExpressionContext) {}

// EnterLambdaParamList is called when production lambdaParamList is entered.
func (s *BaseArcParserListener) EnterLambdaParamList(ctx *LambdaParamListContext) {}

// ExitLambdaParamList is called when production lambdaParamList is exited.
func (s *BaseArcParserListener) ExitLambdaParamList(ctx *LambdaParamListContext) {}

// EnterLambdaParam is called when production lambdaParam is entered.
func (s *BaseArcParserListener) EnterLambdaParam(ctx *LambdaParamContext) {}

// ExitLambdaParam is called when production lambdaParam is exited.
func (s *BaseArcParserListener) ExitLambdaParam(ctx *LambdaParamContext) {}

// EnterTupleExpression is called when production tupleExpression is entered.
func (s *BaseArcParserListener) EnterTupleExpression(ctx *TupleExpressionContext) {}

// ExitTupleExpression is called when production tupleExpression is exited.
func (s *BaseArcParserListener) ExitTupleExpression(ctx *TupleExpressionContext) {}
