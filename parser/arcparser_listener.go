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

	// EnterExternCDecl is called when entering the externCDecl production.
	EnterExternCDecl(c *ExternCDeclContext)

	// EnterExternCMember is called when entering the externCMember production.
	EnterExternCMember(c *ExternCMemberContext)

	// EnterExternCFunctionDecl is called when entering the externCFunctionDecl production.
	EnterExternCFunctionDecl(c *ExternCFunctionDeclContext)

	// EnterCCallingConvention is called when entering the cCallingConvention production.
	EnterCCallingConvention(c *CCallingConventionContext)

	// EnterExternCParameterList is called when entering the externCParameterList production.
	EnterExternCParameterList(c *ExternCParameterListContext)

	// EnterExternCParameter is called when entering the externCParameter production.
	EnterExternCParameter(c *ExternCParameterContext)

	// EnterExternCConstDecl is called when entering the externCConstDecl production.
	EnterExternCConstDecl(c *ExternCConstDeclContext)

	// EnterExternCTypeAlias is called when entering the externCTypeAlias production.
	EnterExternCTypeAlias(c *ExternCTypeAliasContext)

	// EnterExternCOpaqueStructDecl is called when entering the externCOpaqueStructDecl production.
	EnterExternCOpaqueStructDecl(c *ExternCOpaqueStructDeclContext)

	// EnterExternCppDecl is called when entering the externCppDecl production.
	EnterExternCppDecl(c *ExternCppDeclContext)

	// EnterExternCppMember is called when entering the externCppMember production.
	EnterExternCppMember(c *ExternCppMemberContext)

	// EnterExternCppNamespaceDecl is called when entering the externCppNamespaceDecl production.
	EnterExternCppNamespaceDecl(c *ExternCppNamespaceDeclContext)

	// EnterExternNamespacePath is called when entering the externNamespacePath production.
	EnterExternNamespacePath(c *ExternNamespacePathContext)

	// EnterExternCppFunctionDecl is called when entering the externCppFunctionDecl production.
	EnterExternCppFunctionDecl(c *ExternCppFunctionDeclContext)

	// EnterCppCallingConvention is called when entering the cppCallingConvention production.
	EnterCppCallingConvention(c *CppCallingConventionContext)

	// EnterExternCppParameterList is called when entering the externCppParameterList production.
	EnterExternCppParameterList(c *ExternCppParameterListContext)

	// EnterExternCppParameter is called when entering the externCppParameter production.
	EnterExternCppParameter(c *ExternCppParameterContext)

	// EnterExternCppParamType is called when entering the externCppParamType production.
	EnterExternCppParamType(c *ExternCppParamTypeContext)

	// EnterExternCppConstDecl is called when entering the externCppConstDecl production.
	EnterExternCppConstDecl(c *ExternCppConstDeclContext)

	// EnterExternCppTypeAlias is called when entering the externCppTypeAlias production.
	EnterExternCppTypeAlias(c *ExternCppTypeAliasContext)

	// EnterExternCppOpaqueClassDecl is called when entering the externCppOpaqueClassDecl production.
	EnterExternCppOpaqueClassDecl(c *ExternCppOpaqueClassDeclContext)

	// EnterExternCppClassDecl is called when entering the externCppClassDecl production.
	EnterExternCppClassDecl(c *ExternCppClassDeclContext)

	// EnterExternCppClassMember is called when entering the externCppClassMember production.
	EnterExternCppClassMember(c *ExternCppClassMemberContext)

	// EnterExternCppConstructorDecl is called when entering the externCppConstructorDecl production.
	EnterExternCppConstructorDecl(c *ExternCppConstructorDeclContext)

	// EnterExternCppDestructorDecl is called when entering the externCppDestructorDecl production.
	EnterExternCppDestructorDecl(c *ExternCppDestructorDeclContext)

	// EnterExternCppMethodDecl is called when entering the externCppMethodDecl production.
	EnterExternCppMethodDecl(c *ExternCppMethodDeclContext)

	// EnterExternCppMethodParams is called when entering the externCppMethodParams production.
	EnterExternCppMethodParams(c *ExternCppMethodParamsContext)

	// EnterExternCppSelfParam is called when entering the externCppSelfParam production.
	EnterExternCppSelfParam(c *ExternCppSelfParamContext)

	// EnterGenericParams is called when entering the genericParams production.
	EnterGenericParams(c *GenericParamsContext)

	// EnterGenericParamList is called when entering the genericParamList production.
	EnterGenericParamList(c *GenericParamListContext)

	// EnterGenericArgs is called when entering the genericArgs production.
	EnterGenericArgs(c *GenericArgsContext)

	// EnterGenericArgList is called when entering the genericArgList production.
	EnterGenericArgList(c *GenericArgListContext)

	// EnterGenericArg is called when entering the genericArg production.
	EnterGenericArg(c *GenericArgContext)

	// EnterFunctionDecl is called when entering the functionDecl production.
	EnterFunctionDecl(c *FunctionDeclContext)

	// EnterReturnType is called when entering the returnType production.
	EnterReturnType(c *ReturnTypeContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterStructDecl is called when entering the structDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterComputeMarker is called when entering the computeMarker production.
	EnterComputeMarker(c *ComputeMarkerContext)

	// EnterStructMember is called when entering the structMember production.
	EnterStructMember(c *StructMemberContext)

	// EnterStructField is called when entering the structField production.
	EnterStructField(c *StructFieldContext)

	// EnterInitDecl is called when entering the initDecl production.
	EnterInitDecl(c *InitDeclContext)

	// EnterClassDecl is called when entering the classDecl production.
	EnterClassDecl(c *ClassDeclContext)

	// EnterClassMember is called when entering the classMember production.
	EnterClassMember(c *ClassMemberContext)

	// EnterClassField is called when entering the classField production.
	EnterClassField(c *ClassFieldContext)

	// EnterEnumDecl is called when entering the enumDecl production.
	EnterEnumDecl(c *EnumDeclContext)

	// EnterEnumMember is called when entering the enumMember production.
	EnterEnumMember(c *EnumMemberContext)

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

	// EnterTuplePattern is called when entering the tuplePattern production.
	EnterTuplePattern(c *TuplePatternContext)

	// EnterTupleType is called when entering the tupleType production.
	EnterTupleType(c *TupleTypeContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterQualifiedType is called when entering the qualifiedType production.
	EnterQualifiedType(c *QualifiedTypeContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterQualifiedIdentifier is called when entering the qualifiedIdentifier production.
	EnterQualifiedIdentifier(c *QualifiedIdentifierContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterPointerType is called when entering the pointerType production.
	EnterPointerType(c *PointerTypeContext)

	// EnterReferenceType is called when entering the referenceType production.
	EnterReferenceType(c *ReferenceTypeContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignmentStmt is called when entering the assignmentStmt production.
	EnterAssignmentStmt(c *AssignmentStmtContext)

	// EnterAssignmentOp is called when entering the assignmentOp production.
	EnterAssignmentOp(c *AssignmentOpContext)

	// EnterLeftHandSide is called when entering the leftHandSide production.
	EnterLeftHandSide(c *LeftHandSideContext)

	// EnterExpressionStmt is called when entering the expressionStmt production.
	EnterExpressionStmt(c *ExpressionStmtContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterDeferStmt is called when entering the deferStmt production.
	EnterDeferStmt(c *DeferStmtContext)

	// EnterBreakStmt is called when entering the breakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the continueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterThrowStmt is called when entering the throwStmt production.
	EnterThrowStmt(c *ThrowStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterForStmt is called when entering the forStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterSwitchStmt is called when entering the switchStmt production.
	EnterSwitchStmt(c *SwitchStmtContext)

	// EnterSwitchCase is called when entering the switchCase production.
	EnterSwitchCase(c *SwitchCaseContext)

	// EnterDefaultCase is called when entering the defaultCase production.
	EnterDefaultCase(c *DefaultCaseContext)

	// EnterTryStmt is called when entering the tryStmt production.
	EnterTryStmt(c *TryStmtContext)

	// EnterExceptClause is called when entering the exceptClause production.
	EnterExceptClause(c *ExceptClauseContext)

	// EnterFinallyClause is called when entering the finallyClause production.
	EnterFinallyClause(c *FinallyClauseContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLogicalOrExpression is called when entering the logicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterLogicalAndExpression is called when entering the logicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterBitOrExpression is called when entering the bitOrExpression production.
	EnterBitOrExpression(c *BitOrExpressionContext)

	// EnterBitXorExpression is called when entering the bitXorExpression production.
	EnterBitXorExpression(c *BitXorExpressionContext)

	// EnterBitAndExpression is called when entering the bitAndExpression production.
	EnterBitAndExpression(c *BitAndExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterShiftExpression is called when entering the shiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

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

	// EnterComputeExpression is called when entering the computeExpression production.
	EnterComputeExpression(c *ComputeExpressionContext)

	// EnterComputeContext is called when entering the computeContext production.
	EnterComputeContext(c *ComputeContextContext)

	// EnterSizeofExpression is called when entering the sizeofExpression production.
	EnterSizeofExpression(c *SizeofExpressionContext)

	// EnterAlignofExpression is called when entering the alignofExpression production.
	EnterAlignofExpression(c *AlignofExpressionContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterInitializerList is called when entering the initializerList production.
	EnterInitializerList(c *InitializerListContext)

	// EnterInitializerEntry is called when entering the initializerEntry production.
	EnterInitializerEntry(c *InitializerEntryContext)

	// EnterStructLiteral is called when entering the structLiteral production.
	EnterStructLiteral(c *StructLiteralContext)

	// EnterFieldInit is called when entering the fieldInit production.
	EnterFieldInit(c *FieldInitContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterLambdaExpression is called when entering the lambdaExpression production.
	EnterLambdaExpression(c *LambdaExpressionContext)

	// EnterAnonymousFuncExpression is called when entering the anonymousFuncExpression production.
	EnterAnonymousFuncExpression(c *AnonymousFuncExpressionContext)

	// EnterLambdaParamList is called when entering the lambdaParamList production.
	EnterLambdaParamList(c *LambdaParamListContext)

	// EnterLambdaParam is called when entering the lambdaParam production.
	EnterLambdaParam(c *LambdaParamContext)

	// EnterTupleExpression is called when entering the tupleExpression production.
	EnterTupleExpression(c *TupleExpressionContext)

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

	// ExitExternCDecl is called when exiting the externCDecl production.
	ExitExternCDecl(c *ExternCDeclContext)

	// ExitExternCMember is called when exiting the externCMember production.
	ExitExternCMember(c *ExternCMemberContext)

	// ExitExternCFunctionDecl is called when exiting the externCFunctionDecl production.
	ExitExternCFunctionDecl(c *ExternCFunctionDeclContext)

	// ExitCCallingConvention is called when exiting the cCallingConvention production.
	ExitCCallingConvention(c *CCallingConventionContext)

	// ExitExternCParameterList is called when exiting the externCParameterList production.
	ExitExternCParameterList(c *ExternCParameterListContext)

	// ExitExternCParameter is called when exiting the externCParameter production.
	ExitExternCParameter(c *ExternCParameterContext)

	// ExitExternCConstDecl is called when exiting the externCConstDecl production.
	ExitExternCConstDecl(c *ExternCConstDeclContext)

	// ExitExternCTypeAlias is called when exiting the externCTypeAlias production.
	ExitExternCTypeAlias(c *ExternCTypeAliasContext)

	// ExitExternCOpaqueStructDecl is called when exiting the externCOpaqueStructDecl production.
	ExitExternCOpaqueStructDecl(c *ExternCOpaqueStructDeclContext)

	// ExitExternCppDecl is called when exiting the externCppDecl production.
	ExitExternCppDecl(c *ExternCppDeclContext)

	// ExitExternCppMember is called when exiting the externCppMember production.
	ExitExternCppMember(c *ExternCppMemberContext)

	// ExitExternCppNamespaceDecl is called when exiting the externCppNamespaceDecl production.
	ExitExternCppNamespaceDecl(c *ExternCppNamespaceDeclContext)

	// ExitExternNamespacePath is called when exiting the externNamespacePath production.
	ExitExternNamespacePath(c *ExternNamespacePathContext)

	// ExitExternCppFunctionDecl is called when exiting the externCppFunctionDecl production.
	ExitExternCppFunctionDecl(c *ExternCppFunctionDeclContext)

	// ExitCppCallingConvention is called when exiting the cppCallingConvention production.
	ExitCppCallingConvention(c *CppCallingConventionContext)

	// ExitExternCppParameterList is called when exiting the externCppParameterList production.
	ExitExternCppParameterList(c *ExternCppParameterListContext)

	// ExitExternCppParameter is called when exiting the externCppParameter production.
	ExitExternCppParameter(c *ExternCppParameterContext)

	// ExitExternCppParamType is called when exiting the externCppParamType production.
	ExitExternCppParamType(c *ExternCppParamTypeContext)

	// ExitExternCppConstDecl is called when exiting the externCppConstDecl production.
	ExitExternCppConstDecl(c *ExternCppConstDeclContext)

	// ExitExternCppTypeAlias is called when exiting the externCppTypeAlias production.
	ExitExternCppTypeAlias(c *ExternCppTypeAliasContext)

	// ExitExternCppOpaqueClassDecl is called when exiting the externCppOpaqueClassDecl production.
	ExitExternCppOpaqueClassDecl(c *ExternCppOpaqueClassDeclContext)

	// ExitExternCppClassDecl is called when exiting the externCppClassDecl production.
	ExitExternCppClassDecl(c *ExternCppClassDeclContext)

	// ExitExternCppClassMember is called when exiting the externCppClassMember production.
	ExitExternCppClassMember(c *ExternCppClassMemberContext)

	// ExitExternCppConstructorDecl is called when exiting the externCppConstructorDecl production.
	ExitExternCppConstructorDecl(c *ExternCppConstructorDeclContext)

	// ExitExternCppDestructorDecl is called when exiting the externCppDestructorDecl production.
	ExitExternCppDestructorDecl(c *ExternCppDestructorDeclContext)

	// ExitExternCppMethodDecl is called when exiting the externCppMethodDecl production.
	ExitExternCppMethodDecl(c *ExternCppMethodDeclContext)

	// ExitExternCppMethodParams is called when exiting the externCppMethodParams production.
	ExitExternCppMethodParams(c *ExternCppMethodParamsContext)

	// ExitExternCppSelfParam is called when exiting the externCppSelfParam production.
	ExitExternCppSelfParam(c *ExternCppSelfParamContext)

	// ExitGenericParams is called when exiting the genericParams production.
	ExitGenericParams(c *GenericParamsContext)

	// ExitGenericParamList is called when exiting the genericParamList production.
	ExitGenericParamList(c *GenericParamListContext)

	// ExitGenericArgs is called when exiting the genericArgs production.
	ExitGenericArgs(c *GenericArgsContext)

	// ExitGenericArgList is called when exiting the genericArgList production.
	ExitGenericArgList(c *GenericArgListContext)

	// ExitGenericArg is called when exiting the genericArg production.
	ExitGenericArg(c *GenericArgContext)

	// ExitFunctionDecl is called when exiting the functionDecl production.
	ExitFunctionDecl(c *FunctionDeclContext)

	// ExitReturnType is called when exiting the returnType production.
	ExitReturnType(c *ReturnTypeContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitStructDecl is called when exiting the structDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitComputeMarker is called when exiting the computeMarker production.
	ExitComputeMarker(c *ComputeMarkerContext)

	// ExitStructMember is called when exiting the structMember production.
	ExitStructMember(c *StructMemberContext)

	// ExitStructField is called when exiting the structField production.
	ExitStructField(c *StructFieldContext)

	// ExitInitDecl is called when exiting the initDecl production.
	ExitInitDecl(c *InitDeclContext)

	// ExitClassDecl is called when exiting the classDecl production.
	ExitClassDecl(c *ClassDeclContext)

	// ExitClassMember is called when exiting the classMember production.
	ExitClassMember(c *ClassMemberContext)

	// ExitClassField is called when exiting the classField production.
	ExitClassField(c *ClassFieldContext)

	// ExitEnumDecl is called when exiting the enumDecl production.
	ExitEnumDecl(c *EnumDeclContext)

	// ExitEnumMember is called when exiting the enumMember production.
	ExitEnumMember(c *EnumMemberContext)

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

	// ExitTuplePattern is called when exiting the tuplePattern production.
	ExitTuplePattern(c *TuplePatternContext)

	// ExitTupleType is called when exiting the tupleType production.
	ExitTupleType(c *TupleTypeContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitQualifiedType is called when exiting the qualifiedType production.
	ExitQualifiedType(c *QualifiedTypeContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitQualifiedIdentifier is called when exiting the qualifiedIdentifier production.
	ExitQualifiedIdentifier(c *QualifiedIdentifierContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitPointerType is called when exiting the pointerType production.
	ExitPointerType(c *PointerTypeContext)

	// ExitReferenceType is called when exiting the referenceType production.
	ExitReferenceType(c *ReferenceTypeContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignmentStmt is called when exiting the assignmentStmt production.
	ExitAssignmentStmt(c *AssignmentStmtContext)

	// ExitAssignmentOp is called when exiting the assignmentOp production.
	ExitAssignmentOp(c *AssignmentOpContext)

	// ExitLeftHandSide is called when exiting the leftHandSide production.
	ExitLeftHandSide(c *LeftHandSideContext)

	// ExitExpressionStmt is called when exiting the expressionStmt production.
	ExitExpressionStmt(c *ExpressionStmtContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitDeferStmt is called when exiting the deferStmt production.
	ExitDeferStmt(c *DeferStmtContext)

	// ExitBreakStmt is called when exiting the breakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the continueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitThrowStmt is called when exiting the throwStmt production.
	ExitThrowStmt(c *ThrowStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitForStmt is called when exiting the forStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitSwitchStmt is called when exiting the switchStmt production.
	ExitSwitchStmt(c *SwitchStmtContext)

	// ExitSwitchCase is called when exiting the switchCase production.
	ExitSwitchCase(c *SwitchCaseContext)

	// ExitDefaultCase is called when exiting the defaultCase production.
	ExitDefaultCase(c *DefaultCaseContext)

	// ExitTryStmt is called when exiting the tryStmt production.
	ExitTryStmt(c *TryStmtContext)

	// ExitExceptClause is called when exiting the exceptClause production.
	ExitExceptClause(c *ExceptClauseContext)

	// ExitFinallyClause is called when exiting the finallyClause production.
	ExitFinallyClause(c *FinallyClauseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLogicalOrExpression is called when exiting the logicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitLogicalAndExpression is called when exiting the logicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitBitOrExpression is called when exiting the bitOrExpression production.
	ExitBitOrExpression(c *BitOrExpressionContext)

	// ExitBitXorExpression is called when exiting the bitXorExpression production.
	ExitBitXorExpression(c *BitXorExpressionContext)

	// ExitBitAndExpression is called when exiting the bitAndExpression production.
	ExitBitAndExpression(c *BitAndExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitShiftExpression is called when exiting the shiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

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

	// ExitComputeExpression is called when exiting the computeExpression production.
	ExitComputeExpression(c *ComputeExpressionContext)

	// ExitComputeContext is called when exiting the computeContext production.
	ExitComputeContext(c *ComputeContextContext)

	// ExitSizeofExpression is called when exiting the sizeofExpression production.
	ExitSizeofExpression(c *SizeofExpressionContext)

	// ExitAlignofExpression is called when exiting the alignofExpression production.
	ExitAlignofExpression(c *AlignofExpressionContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitInitializerList is called when exiting the initializerList production.
	ExitInitializerList(c *InitializerListContext)

	// ExitInitializerEntry is called when exiting the initializerEntry production.
	ExitInitializerEntry(c *InitializerEntryContext)

	// ExitStructLiteral is called when exiting the structLiteral production.
	ExitStructLiteral(c *StructLiteralContext)

	// ExitFieldInit is called when exiting the fieldInit production.
	ExitFieldInit(c *FieldInitContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitLambdaExpression is called when exiting the lambdaExpression production.
	ExitLambdaExpression(c *LambdaExpressionContext)

	// ExitAnonymousFuncExpression is called when exiting the anonymousFuncExpression production.
	ExitAnonymousFuncExpression(c *AnonymousFuncExpressionContext)

	// ExitLambdaParamList is called when exiting the lambdaParamList production.
	ExitLambdaParamList(c *LambdaParamListContext)

	// ExitLambdaParam is called when exiting the lambdaParam production.
	ExitLambdaParam(c *LambdaParamContext)

	// ExitTupleExpression is called when exiting the tupleExpression production.
	ExitTupleExpression(c *TupleExpressionContext)
}
