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

// EnterExternCDecl is called when production externCDecl is entered.
func (s *BaseArcParserListener) EnterExternCDecl(ctx *ExternCDeclContext) {}

// ExitExternCDecl is called when production externCDecl is exited.
func (s *BaseArcParserListener) ExitExternCDecl(ctx *ExternCDeclContext) {}

// EnterExternCMember is called when production externCMember is entered.
func (s *BaseArcParserListener) EnterExternCMember(ctx *ExternCMemberContext) {}

// ExitExternCMember is called when production externCMember is exited.
func (s *BaseArcParserListener) ExitExternCMember(ctx *ExternCMemberContext) {}

// EnterExternCFunctionDecl is called when production externCFunctionDecl is entered.
func (s *BaseArcParserListener) EnterExternCFunctionDecl(ctx *ExternCFunctionDeclContext) {}

// ExitExternCFunctionDecl is called when production externCFunctionDecl is exited.
func (s *BaseArcParserListener) ExitExternCFunctionDecl(ctx *ExternCFunctionDeclContext) {}

// EnterCCallingConvention is called when production cCallingConvention is entered.
func (s *BaseArcParserListener) EnterCCallingConvention(ctx *CCallingConventionContext) {}

// ExitCCallingConvention is called when production cCallingConvention is exited.
func (s *BaseArcParserListener) ExitCCallingConvention(ctx *CCallingConventionContext) {}

// EnterExternCParameterList is called when production externCParameterList is entered.
func (s *BaseArcParserListener) EnterExternCParameterList(ctx *ExternCParameterListContext) {}

// ExitExternCParameterList is called when production externCParameterList is exited.
func (s *BaseArcParserListener) ExitExternCParameterList(ctx *ExternCParameterListContext) {}

// EnterExternCParameter is called when production externCParameter is entered.
func (s *BaseArcParserListener) EnterExternCParameter(ctx *ExternCParameterContext) {}

// ExitExternCParameter is called when production externCParameter is exited.
func (s *BaseArcParserListener) ExitExternCParameter(ctx *ExternCParameterContext) {}

// EnterExternCConstDecl is called when production externCConstDecl is entered.
func (s *BaseArcParserListener) EnterExternCConstDecl(ctx *ExternCConstDeclContext) {}

// ExitExternCConstDecl is called when production externCConstDecl is exited.
func (s *BaseArcParserListener) ExitExternCConstDecl(ctx *ExternCConstDeclContext) {}

// EnterExternCTypeAlias is called when production externCTypeAlias is entered.
func (s *BaseArcParserListener) EnterExternCTypeAlias(ctx *ExternCTypeAliasContext) {}

// ExitExternCTypeAlias is called when production externCTypeAlias is exited.
func (s *BaseArcParserListener) ExitExternCTypeAlias(ctx *ExternCTypeAliasContext) {}

// EnterExternCOpaqueStructDecl is called when production externCOpaqueStructDecl is entered.
func (s *BaseArcParserListener) EnterExternCOpaqueStructDecl(ctx *ExternCOpaqueStructDeclContext) {}

// ExitExternCOpaqueStructDecl is called when production externCOpaqueStructDecl is exited.
func (s *BaseArcParserListener) ExitExternCOpaqueStructDecl(ctx *ExternCOpaqueStructDeclContext) {}

// EnterExternCppDecl is called when production externCppDecl is entered.
func (s *BaseArcParserListener) EnterExternCppDecl(ctx *ExternCppDeclContext) {}

// ExitExternCppDecl is called when production externCppDecl is exited.
func (s *BaseArcParserListener) ExitExternCppDecl(ctx *ExternCppDeclContext) {}

// EnterExternCppMember is called when production externCppMember is entered.
func (s *BaseArcParserListener) EnterExternCppMember(ctx *ExternCppMemberContext) {}

// ExitExternCppMember is called when production externCppMember is exited.
func (s *BaseArcParserListener) ExitExternCppMember(ctx *ExternCppMemberContext) {}

// EnterExternCppNamespaceDecl is called when production externCppNamespaceDecl is entered.
func (s *BaseArcParserListener) EnterExternCppNamespaceDecl(ctx *ExternCppNamespaceDeclContext) {}

// ExitExternCppNamespaceDecl is called when production externCppNamespaceDecl is exited.
func (s *BaseArcParserListener) ExitExternCppNamespaceDecl(ctx *ExternCppNamespaceDeclContext) {}

// EnterExternNamespacePath is called when production externNamespacePath is entered.
func (s *BaseArcParserListener) EnterExternNamespacePath(ctx *ExternNamespacePathContext) {}

// ExitExternNamespacePath is called when production externNamespacePath is exited.
func (s *BaseArcParserListener) ExitExternNamespacePath(ctx *ExternNamespacePathContext) {}

// EnterExternCppFunctionDecl is called when production externCppFunctionDecl is entered.
func (s *BaseArcParserListener) EnterExternCppFunctionDecl(ctx *ExternCppFunctionDeclContext) {}

// ExitExternCppFunctionDecl is called when production externCppFunctionDecl is exited.
func (s *BaseArcParserListener) ExitExternCppFunctionDecl(ctx *ExternCppFunctionDeclContext) {}

// EnterCppCallingConvention is called when production cppCallingConvention is entered.
func (s *BaseArcParserListener) EnterCppCallingConvention(ctx *CppCallingConventionContext) {}

// ExitCppCallingConvention is called when production cppCallingConvention is exited.
func (s *BaseArcParserListener) ExitCppCallingConvention(ctx *CppCallingConventionContext) {}

// EnterExternCppParameterList is called when production externCppParameterList is entered.
func (s *BaseArcParserListener) EnterExternCppParameterList(ctx *ExternCppParameterListContext) {}

// ExitExternCppParameterList is called when production externCppParameterList is exited.
func (s *BaseArcParserListener) ExitExternCppParameterList(ctx *ExternCppParameterListContext) {}

// EnterExternCppParameter is called when production externCppParameter is entered.
func (s *BaseArcParserListener) EnterExternCppParameter(ctx *ExternCppParameterContext) {}

// ExitExternCppParameter is called when production externCppParameter is exited.
func (s *BaseArcParserListener) ExitExternCppParameter(ctx *ExternCppParameterContext) {}

// EnterExternCppParamType is called when production externCppParamType is entered.
func (s *BaseArcParserListener) EnterExternCppParamType(ctx *ExternCppParamTypeContext) {}

// ExitExternCppParamType is called when production externCppParamType is exited.
func (s *BaseArcParserListener) ExitExternCppParamType(ctx *ExternCppParamTypeContext) {}

// EnterExternCppConstDecl is called when production externCppConstDecl is entered.
func (s *BaseArcParserListener) EnterExternCppConstDecl(ctx *ExternCppConstDeclContext) {}

// ExitExternCppConstDecl is called when production externCppConstDecl is exited.
func (s *BaseArcParserListener) ExitExternCppConstDecl(ctx *ExternCppConstDeclContext) {}

// EnterExternCppTypeAlias is called when production externCppTypeAlias is entered.
func (s *BaseArcParserListener) EnterExternCppTypeAlias(ctx *ExternCppTypeAliasContext) {}

// ExitExternCppTypeAlias is called when production externCppTypeAlias is exited.
func (s *BaseArcParserListener) ExitExternCppTypeAlias(ctx *ExternCppTypeAliasContext) {}

// EnterExternCppOpaqueClassDecl is called when production externCppOpaqueClassDecl is entered.
func (s *BaseArcParserListener) EnterExternCppOpaqueClassDecl(ctx *ExternCppOpaqueClassDeclContext) {}

// ExitExternCppOpaqueClassDecl is called when production externCppOpaqueClassDecl is exited.
func (s *BaseArcParserListener) ExitExternCppOpaqueClassDecl(ctx *ExternCppOpaqueClassDeclContext) {}

// EnterExternCppClassDecl is called when production externCppClassDecl is entered.
func (s *BaseArcParserListener) EnterExternCppClassDecl(ctx *ExternCppClassDeclContext) {}

// ExitExternCppClassDecl is called when production externCppClassDecl is exited.
func (s *BaseArcParserListener) ExitExternCppClassDecl(ctx *ExternCppClassDeclContext) {}

// EnterExternCppClassMember is called when production externCppClassMember is entered.
func (s *BaseArcParserListener) EnterExternCppClassMember(ctx *ExternCppClassMemberContext) {}

// ExitExternCppClassMember is called when production externCppClassMember is exited.
func (s *BaseArcParserListener) ExitExternCppClassMember(ctx *ExternCppClassMemberContext) {}

// EnterExternCppConstructorDecl is called when production externCppConstructorDecl is entered.
func (s *BaseArcParserListener) EnterExternCppConstructorDecl(ctx *ExternCppConstructorDeclContext) {}

// ExitExternCppConstructorDecl is called when production externCppConstructorDecl is exited.
func (s *BaseArcParserListener) ExitExternCppConstructorDecl(ctx *ExternCppConstructorDeclContext) {}

// EnterExternCppDestructorDecl is called when production externCppDestructorDecl is entered.
func (s *BaseArcParserListener) EnterExternCppDestructorDecl(ctx *ExternCppDestructorDeclContext) {}

// ExitExternCppDestructorDecl is called when production externCppDestructorDecl is exited.
func (s *BaseArcParserListener) ExitExternCppDestructorDecl(ctx *ExternCppDestructorDeclContext) {}

// EnterExternCppMethodDecl is called when production externCppMethodDecl is entered.
func (s *BaseArcParserListener) EnterExternCppMethodDecl(ctx *ExternCppMethodDeclContext) {}

// ExitExternCppMethodDecl is called when production externCppMethodDecl is exited.
func (s *BaseArcParserListener) ExitExternCppMethodDecl(ctx *ExternCppMethodDeclContext) {}

// EnterExternCppMethodParams is called when production externCppMethodParams is entered.
func (s *BaseArcParserListener) EnterExternCppMethodParams(ctx *ExternCppMethodParamsContext) {}

// ExitExternCppMethodParams is called when production externCppMethodParams is exited.
func (s *BaseArcParserListener) ExitExternCppMethodParams(ctx *ExternCppMethodParamsContext) {}

// EnterExternCppSelfParam is called when production externCppSelfParam is entered.
func (s *BaseArcParserListener) EnterExternCppSelfParam(ctx *ExternCppSelfParamContext) {}

// ExitExternCppSelfParam is called when production externCppSelfParam is exited.
func (s *BaseArcParserListener) ExitExternCppSelfParam(ctx *ExternCppSelfParamContext) {}

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

// EnterComputeMarker is called when production computeMarker is entered.
func (s *BaseArcParserListener) EnterComputeMarker(ctx *ComputeMarkerContext) {}

// ExitComputeMarker is called when production computeMarker is exited.
func (s *BaseArcParserListener) ExitComputeMarker(ctx *ComputeMarkerContext) {}

// EnterStructMember is called when production structMember is entered.
func (s *BaseArcParserListener) EnterStructMember(ctx *StructMemberContext) {}

// ExitStructMember is called when production structMember is exited.
func (s *BaseArcParserListener) ExitStructMember(ctx *StructMemberContext) {}

// EnterStructField is called when production structField is entered.
func (s *BaseArcParserListener) EnterStructField(ctx *StructFieldContext) {}

// ExitStructField is called when production structField is exited.
func (s *BaseArcParserListener) ExitStructField(ctx *StructFieldContext) {}

// EnterInitDecl is called when production initDecl is entered.
func (s *BaseArcParserListener) EnterInitDecl(ctx *InitDeclContext) {}

// ExitInitDecl is called when production initDecl is exited.
func (s *BaseArcParserListener) ExitInitDecl(ctx *InitDeclContext) {}

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

// EnterArrayType is called when production arrayType is entered.
func (s *BaseArcParserListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseArcParserListener) ExitArrayType(ctx *ArrayTypeContext) {}

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

// EnterComputeExpression is called when production computeExpression is entered.
func (s *BaseArcParserListener) EnterComputeExpression(ctx *ComputeExpressionContext) {}

// ExitComputeExpression is called when production computeExpression is exited.
func (s *BaseArcParserListener) ExitComputeExpression(ctx *ComputeExpressionContext) {}

// EnterComputeContext is called when production computeContext is entered.
func (s *BaseArcParserListener) EnterComputeContext(ctx *ComputeContextContext) {}

// ExitComputeContext is called when production computeContext is exited.
func (s *BaseArcParserListener) ExitComputeContext(ctx *ComputeContextContext) {}

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
