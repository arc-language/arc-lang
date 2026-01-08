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

func (v *BaseArcParserVisitor) VisitExternCDecl(ctx *ExternCDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCMember(ctx *ExternCMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCFunctionDecl(ctx *ExternCFunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitCCallingConvention(ctx *CCallingConventionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCParameterList(ctx *ExternCParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCParameter(ctx *ExternCParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCConstDecl(ctx *ExternCConstDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCTypeAlias(ctx *ExternCTypeAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCOpaqueStructDecl(ctx *ExternCOpaqueStructDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppDecl(ctx *ExternCppDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppMember(ctx *ExternCppMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppNamespaceDecl(ctx *ExternCppNamespaceDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternNamespacePath(ctx *ExternNamespacePathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppFunctionDecl(ctx *ExternCppFunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitCppCallingConvention(ctx *CppCallingConventionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppParameterList(ctx *ExternCppParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppParameter(ctx *ExternCppParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppParamType(ctx *ExternCppParamTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppConstDecl(ctx *ExternCppConstDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppTypeAlias(ctx *ExternCppTypeAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppOpaqueClassDecl(ctx *ExternCppOpaqueClassDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppClassDecl(ctx *ExternCppClassDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppClassMember(ctx *ExternCppClassMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppConstructorDecl(ctx *ExternCppConstructorDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppDestructorDecl(ctx *ExternCppDestructorDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppMethodDecl(ctx *ExternCppMethodDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppMethodParams(ctx *ExternCppMethodParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExternCppSelfParam(ctx *ExternCppSelfParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitGenericParams(ctx *GenericParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitGenericParamList(ctx *GenericParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitGenericArgs(ctx *GenericArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitGenericArgList(ctx *GenericArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitGenericArg(ctx *GenericArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitReturnType(ctx *ReturnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
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

func (v *BaseArcParserVisitor) VisitComputeMarker(ctx *ComputeMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStructMember(ctx *StructMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStructField(ctx *StructFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitInitDecl(ctx *InitDeclContext) interface{} {
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

func (v *BaseArcParserVisitor) VisitEnumDecl(ctx *EnumDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitEnumMember(ctx *EnumMemberContext) interface{} {
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

func (v *BaseArcParserVisitor) VisitTuplePattern(ctx *TuplePatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitTupleType(ctx *TupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitQualifiedType(ctx *QualifiedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{} {
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

func (v *BaseArcParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitAssignmentStmt(ctx *AssignmentStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitAssignmentOp(ctx *AssignmentOpContext) interface{} {
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

func (v *BaseArcParserVisitor) VisitDeferStmt(ctx *DeferStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitThrowStmt(ctx *ThrowStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitSwitchCase(ctx *SwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitDefaultCase(ctx *DefaultCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitTryStmt(ctx *TryStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitExceptClause(ctx *ExceptClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitFinallyClause(ctx *FinallyClauseContext) interface{} {
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

func (v *BaseArcParserVisitor) VisitBitOrExpression(ctx *BitOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitBitXorExpression(ctx *BitXorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitBitAndExpression(ctx *BitAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitShiftExpression(ctx *ShiftExpressionContext) interface{} {
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

func (v *BaseArcParserVisitor) VisitComputeExpression(ctx *ComputeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitComputeContext(ctx *ComputeContextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitSizeofExpression(ctx *SizeofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitAlignofExpression(ctx *AlignofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitInitializerList(ctx *InitializerListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitInitializerEntry(ctx *InitializerEntryContext) interface{} {
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

func (v *BaseArcParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitLambdaExpression(ctx *LambdaExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitAnonymousFuncExpression(ctx *AnonymousFuncExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitLambdaParamList(ctx *LambdaParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitLambdaParam(ctx *LambdaParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArcParserVisitor) VisitTupleExpression(ctx *TupleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
