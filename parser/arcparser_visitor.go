// Code generated from ArcParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ArcParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ArcParser.
type ArcParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ArcParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by ArcParser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#importSpec.
	VisitImportSpec(ctx *ImportSpecContext) interface{}

	// Visit a parse tree produced by ArcParser#namespaceDecl.
	VisitNamespaceDecl(ctx *NamespaceDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#topLevelDecl.
	VisitTopLevelDecl(ctx *TopLevelDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by ArcParser#externCDecl.
	VisitExternCDecl(ctx *ExternCDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externCMember.
	VisitExternCMember(ctx *ExternCMemberContext) interface{}

	// Visit a parse tree produced by ArcParser#externCFunctionDecl.
	VisitExternCFunctionDecl(ctx *ExternCFunctionDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#cCallingConvention.
	VisitCCallingConvention(ctx *CCallingConventionContext) interface{}

	// Visit a parse tree produced by ArcParser#externCParameterList.
	VisitExternCParameterList(ctx *ExternCParameterListContext) interface{}

	// Visit a parse tree produced by ArcParser#externCParameter.
	VisitExternCParameter(ctx *ExternCParameterContext) interface{}

	// Visit a parse tree produced by ArcParser#externCConstDecl.
	VisitExternCConstDecl(ctx *ExternCConstDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externCTypeAlias.
	VisitExternCTypeAlias(ctx *ExternCTypeAliasContext) interface{}

	// Visit a parse tree produced by ArcParser#externCStructDecl.
	VisitExternCStructDecl(ctx *ExternCStructDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externCStructField.
	VisitExternCStructField(ctx *ExternCStructFieldContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppDecl.
	VisitExternCppDecl(ctx *ExternCppDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppMember.
	VisitExternCppMember(ctx *ExternCppMemberContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppNamespaceDecl.
	VisitExternCppNamespaceDecl(ctx *ExternCppNamespaceDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externNamespacePath.
	VisitExternNamespacePath(ctx *ExternNamespacePathContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppFunctionDecl.
	VisitExternCppFunctionDecl(ctx *ExternCppFunctionDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#cppCallingConvention.
	VisitCppCallingConvention(ctx *CppCallingConventionContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppParameterList.
	VisitExternCppParameterList(ctx *ExternCppParameterListContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppParameter.
	VisitExternCppParameter(ctx *ExternCppParameterContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppParamType.
	VisitExternCppParamType(ctx *ExternCppParamTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppConstDecl.
	VisitExternCppConstDecl(ctx *ExternCppConstDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppTypeAlias.
	VisitExternCppTypeAlias(ctx *ExternCppTypeAliasContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppClassDecl.
	VisitExternCppClassDecl(ctx *ExternCppClassDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppClassMember.
	VisitExternCppClassMember(ctx *ExternCppClassMemberContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppConstructorDecl.
	VisitExternCppConstructorDecl(ctx *ExternCppConstructorDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppDestructorDecl.
	VisitExternCppDestructorDecl(ctx *ExternCppDestructorDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppMethodDecl.
	VisitExternCppMethodDecl(ctx *ExternCppMethodDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppMethodParams.
	VisitExternCppMethodParams(ctx *ExternCppMethodParamsContext) interface{}

	// Visit a parse tree produced by ArcParser#externCppSelfParam.
	VisitExternCppSelfParam(ctx *ExternCppSelfParamContext) interface{}

	// Visit a parse tree produced by ArcParser#externType.
	VisitExternType(ctx *ExternTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#externPointerType.
	VisitExternPointerType(ctx *ExternPointerTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#externPrimitiveType.
	VisitExternPrimitiveType(ctx *ExternPrimitiveTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#externFunctionType.
	VisitExternFunctionType(ctx *ExternFunctionTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#externTypeList.
	VisitExternTypeList(ctx *ExternTypeListContext) interface{}

	// Visit a parse tree produced by ArcParser#genericParams.
	VisitGenericParams(ctx *GenericParamsContext) interface{}

	// Visit a parse tree produced by ArcParser#genericParamList.
	VisitGenericParamList(ctx *GenericParamListContext) interface{}

	// Visit a parse tree produced by ArcParser#genericParam.
	VisitGenericParam(ctx *GenericParamContext) interface{}

	// Visit a parse tree produced by ArcParser#genericArgs.
	VisitGenericArgs(ctx *GenericArgsContext) interface{}

	// Visit a parse tree produced by ArcParser#genericArgList.
	VisitGenericArgList(ctx *GenericArgListContext) interface{}

	// Visit a parse tree produced by ArcParser#genericArg.
	VisitGenericArg(ctx *GenericArgContext) interface{}

	// Visit a parse tree produced by ArcParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by ArcParser#collectionType.
	VisitCollectionType(ctx *CollectionTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#qualifiedType.
	VisitQualifiedType(ctx *QualifiedTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#functionType.
	VisitFunctionType(ctx *FunctionTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by ArcParser#returnType.
	VisitReturnType(ctx *ReturnTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#qualifiedIdentifier.
	VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{}

	// Visit a parse tree produced by ArcParser#functionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#methodDecl.
	VisitMethodDecl(ctx *MethodDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#deinitDecl.
	VisitDeinitDecl(ctx *DeinitDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#structDecl.
	VisitStructDecl(ctx *StructDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#structMember.
	VisitStructMember(ctx *StructMemberContext) interface{}

	// Visit a parse tree produced by ArcParser#structField.
	VisitStructField(ctx *StructFieldContext) interface{}

	// Visit a parse tree produced by ArcParser#classDecl.
	VisitClassDecl(ctx *ClassDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#classMember.
	VisitClassMember(ctx *ClassMemberContext) interface{}

	// Visit a parse tree produced by ArcParser#classField.
	VisitClassField(ctx *ClassFieldContext) interface{}

	// Visit a parse tree produced by ArcParser#enumDecl.
	VisitEnumDecl(ctx *EnumDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#enumMember.
	VisitEnumMember(ctx *EnumMemberContext) interface{}

	// Visit a parse tree produced by ArcParser#variableDecl.
	VisitVariableDecl(ctx *VariableDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#constDecl.
	VisitConstDecl(ctx *ConstDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#tuplePattern.
	VisitTuplePattern(ctx *TuplePatternContext) interface{}

	// Visit a parse tree produced by ArcParser#tupleType.
	VisitTupleType(ctx *TupleTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by ArcParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by ArcParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ArcParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ArcParser#assignmentStmt.
	VisitAssignmentStmt(ctx *AssignmentStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#assignmentOp.
	VisitAssignmentOp(ctx *AssignmentOpContext) interface{}

	// Visit a parse tree produced by ArcParser#expressionStmt.
	VisitExpressionStmt(ctx *ExpressionStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#deferStmt.
	VisitDeferStmt(ctx *DeferStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#continueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#switchStmt.
	VisitSwitchStmt(ctx *SwitchStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#switchCase.
	VisitSwitchCase(ctx *SwitchCaseContext) interface{}

	// Visit a parse tree produced by ArcParser#defaultCase.
	VisitDefaultCase(ctx *DefaultCaseContext) interface{}

	// Visit a parse tree produced by ArcParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#bitOrExpression.
	VisitBitOrExpression(ctx *BitOrExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#bitXorExpression.
	VisitBitXorExpression(ctx *BitXorExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#bitAndExpression.
	VisitBitAndExpression(ctx *BitAndExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#shiftExpression.
	VisitShiftExpression(ctx *ShiftExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#rangeExpression.
	VisitRangeExpression(ctx *RangeExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#postfixOp.
	VisitPostfixOp(ctx *PostfixOpContext) interface{}

	// Visit a parse tree produced by ArcParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#newExpression.
	VisitNewExpression(ctx *NewExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#deleteExpression.
	VisitDeleteExpression(ctx *DeleteExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#castExpression.
	VisitCastExpression(ctx *CastExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#castTargetType.
	VisitCastTargetType(ctx *CastTargetTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#builtinExpression.
	VisitBuiltinExpression(ctx *BuiltinExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#initializerList.
	VisitInitializerList(ctx *InitializerListContext) interface{}

	// Visit a parse tree produced by ArcParser#initializerEntry.
	VisitInitializerEntry(ctx *InitializerEntryContext) interface{}

	// Visit a parse tree produced by ArcParser#structLiteral.
	VisitStructLiteral(ctx *StructLiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#fieldInit.
	VisitFieldInit(ctx *FieldInitContext) interface{}

	// Visit a parse tree produced by ArcParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by ArcParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by ArcParser#lambdaExpression.
	VisitLambdaExpression(ctx *LambdaExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#anonymousFuncExpression.
	VisitAnonymousFuncExpression(ctx *AnonymousFuncExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#executionStrategy.
	VisitExecutionStrategy(ctx *ExecutionStrategyContext) interface{}

	// Visit a parse tree produced by ArcParser#lambdaParamList.
	VisitLambdaParamList(ctx *LambdaParamListContext) interface{}

	// Visit a parse tree produced by ArcParser#lambdaParam.
	VisitLambdaParam(ctx *LambdaParamContext) interface{}

	// Visit a parse tree produced by ArcParser#tupleExpression.
	VisitTupleExpression(ctx *TupleExpressionContext) interface{}
}
