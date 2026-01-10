

CONSTANTS

const (
	ArcLexerIMPORT          = 1
	ArcLexerNAMESPACE       = 2
	ArcLexerLET             = 3
	ArcLexerCONST           = 4
	ArcLexerFUNC            = 5
	ArcLexerASYNC           = 6
	ArcLexerAWAIT           = 7
	ArcLexerPROCESS         = 8
	ArcLexerCONTAINER       = 9
	ArcLexerSTRUCT          = 10
	ArcLexerCLASS           = 11
	ArcLexerMUTATING        = 12
	ArcLexerINIT            = 13
	ArcLexerDEINIT          = 14
	ArcLexerRETURN          = 15
	ArcLexerIF              = 16
	ArcLexerELSE            = 17
	ArcLexerFOR             = 18
	ArcLexerIN              = 19
	ArcLexerBREAK           = 20
	ArcLexerCONTINUE        = 21
	ArcLexerDEFER           = 22
	ArcLexerSELF            = 23
	ArcLexerNULL            = 24
	ArcLexerSWITCH          = 25
	ArcLexerCASE            = 26
	ArcLexerDEFAULT         = 27
	ArcLexerTRY             = 28
	ArcLexerTHROW           = 29
	ArcLexerEXCEPT          = 30
	ArcLexerFINALLY         = 31
	ArcLexerENUM            = 32
	ArcLexerCOMPUTE         = 33
	ArcLexerEXTERN          = 34
	ArcLexerOPAQUE          = 35
	ArcLexerVIRTUAL         = 36
	ArcLexerSTATIC          = 37
	ArcLexerABSTRACT        = 38
	ArcLexerNEW             = 39
	ArcLexerDELETE          = 40
	ArcLexerTYPE            = 41
	ArcLexerSTDCALL         = 42
	ArcLexerCDECL           = 43
	ArcLexerFASTCALL        = 44
	ArcLexerVECTORCALL      = 45
	ArcLexerTHISCALL        = 46
	ArcLexerINT8            = 47
	ArcLexerINT16           = 48
	ArcLexerINT32           = 49
	ArcLexerINT64           = 50
	ArcLexerUINT8           = 51
	ArcLexerUINT16          = 52
	ArcLexerUINT32          = 53
	ArcLexerUINT64          = 54
	ArcLexerUSIZE           = 55
	ArcLexerISIZE           = 56
	ArcLexerFLOAT32         = 57
	ArcLexerFLOAT64         = 58
	ArcLexerBYTE            = 59
	ArcLexerBOOL            = 60
	ArcLexerCHAR            = 61
	ArcLexerSTRING          = 62
	ArcLexerVOID            = 63
	ArcLexerFLOAT16         = 64
	ArcLexerBFLOAT16        = 65
	ArcLexerSIZEOF          = 66
	ArcLexerALIGNOF         = 67
	ArcLexerARROW           = 68
	ArcLexerRANGE           = 69
	ArcLexerELLIPSIS        = 70
	ArcLexerEQ              = 71
	ArcLexerNE              = 72
	ArcLexerLE              = 73
	ArcLexerGE              = 74
	ArcLexerAND             = 75
	ArcLexerOR              = 76
	ArcLexerPLUS_ASSIGN     = 77
	ArcLexerMINUS_ASSIGN    = 78
	ArcLexerSTAR_ASSIGN     = 79
	ArcLexerSLASH_ASSIGN    = 80
	ArcLexerPERCENT_ASSIGN  = 81
	ArcLexerBIT_OR_ASSIGN   = 82
	ArcLexerBIT_AND_ASSIGN  = 83
	ArcLexerBIT_XOR_ASSIGN  = 84
	ArcLexerINCREMENT       = 85
	ArcLexerDECREMENT       = 86
	ArcLexerFAT_ARROW       = 87
	ArcLexerPLUS            = 88
	ArcLexerMINUS           = 89
	ArcLexerSTAR            = 90
	ArcLexerSLASH           = 91
	ArcLexerPERCENT         = 92
	ArcLexerLT              = 93
	ArcLexerGT              = 94
	ArcLexerNOT             = 95
	ArcLexerAMP             = 96
	ArcLexerBIT_OR          = 97
	ArcLexerBIT_XOR         = 98
	ArcLexerBIT_NOT         = 99
	ArcLexerAT              = 100
	ArcLexerASSIGN          = 101
	ArcLexerLPAREN          = 102
	ArcLexerRPAREN          = 103
	ArcLexerLBRACE          = 104
	ArcLexerRBRACE          = 105
	ArcLexerLBRACKET        = 106
	ArcLexerRBRACKET        = 107
	ArcLexerCOMMA           = 108
	ArcLexerCOLON           = 109
	ArcLexerSEMICOLON       = 110
	ArcLexerDOT             = 111
	ArcLexerUNDERSCORE      = 112
	ArcLexerBOOLEAN_LITERAL = 113
	ArcLexerINTEGER_LITERAL = 114
	ArcLexerFLOAT_LITERAL   = 115
	ArcLexerSTRING_LITERAL  = 116
	ArcLexerCHAR_LITERAL    = 117
	ArcLexerIDENTIFIER      = 118
	ArcLexerWS              = 119
	ArcLexerLINE_COMMENT    = 120
	ArcLexerBLOCK_COMMENT   = 121
	ArcLexerEXTERN_WS       = 122
	ArcLexerC_LANG          = 123
	ArcLexerCPP_LANG        = 124
)
    ArcLexer tokens.

const (
	ArcParserEOF             = antlr.TokenEOF
	ArcParserIMPORT          = 1
	ArcParserNAMESPACE       = 2
	ArcParserLET             = 3
	ArcParserCONST           = 4
	ArcParserFUNC            = 5
	ArcParserASYNC           = 6
	ArcParserAWAIT           = 7
	ArcParserPROCESS         = 8
	ArcParserCONTAINER       = 9
	ArcParserSTRUCT          = 10
	ArcParserCLASS           = 11
	ArcParserMUTATING        = 12
	ArcParserINIT            = 13
	ArcParserDEINIT          = 14
	ArcParserRETURN          = 15
	ArcParserIF              = 16
	ArcParserELSE            = 17
	ArcParserFOR             = 18
	ArcParserIN              = 19
	ArcParserBREAK           = 20
	ArcParserCONTINUE        = 21
	ArcParserDEFER           = 22
	ArcParserSELF            = 23
	ArcParserNULL            = 24
	ArcParserSWITCH          = 25
	ArcParserCASE            = 26
	ArcParserDEFAULT         = 27
	ArcParserTRY             = 28
	ArcParserTHROW           = 29
	ArcParserEXCEPT          = 30
	ArcParserFINALLY         = 31
	ArcParserENUM            = 32
	ArcParserCOMPUTE         = 33
	ArcParserEXTERN          = 34
	ArcParserOPAQUE          = 35
	ArcParserVIRTUAL         = 36
	ArcParserSTATIC          = 37
	ArcParserABSTRACT        = 38
	ArcParserNEW             = 39
	ArcParserDELETE          = 40
	ArcParserTYPE            = 41
	ArcParserSTDCALL         = 42
	ArcParserCDECL           = 43
	ArcParserFASTCALL        = 44
	ArcParserVECTORCALL      = 45
	ArcParserTHISCALL        = 46
	ArcParserINT8            = 47
	ArcParserINT16           = 48
	ArcParserINT32           = 49
	ArcParserINT64           = 50
	ArcParserUINT8           = 51
	ArcParserUINT16          = 52
	ArcParserUINT32          = 53
	ArcParserUINT64          = 54
	ArcParserUSIZE           = 55
	ArcParserISIZE           = 56
	ArcParserFLOAT32         = 57
	ArcParserFLOAT64         = 58
	ArcParserBYTE            = 59
	ArcParserBOOL            = 60
	ArcParserCHAR            = 61
	ArcParserSTRING          = 62
	ArcParserVOID            = 63
	ArcParserFLOAT16         = 64
	ArcParserBFLOAT16        = 65
	ArcParserSIZEOF          = 66
	ArcParserALIGNOF         = 67
	ArcParserARROW           = 68
	ArcParserRANGE           = 69
	ArcParserELLIPSIS        = 70
	ArcParserEQ              = 71
	ArcParserNE              = 72
	ArcParserLE              = 73
	ArcParserGE              = 74
	ArcParserAND             = 75
	ArcParserOR              = 76
	ArcParserPLUS_ASSIGN     = 77
	ArcParserMINUS_ASSIGN    = 78
	ArcParserSTAR_ASSIGN     = 79
	ArcParserSLASH_ASSIGN    = 80
	ArcParserPERCENT_ASSIGN  = 81
	ArcParserBIT_OR_ASSIGN   = 82
	ArcParserBIT_AND_ASSIGN  = 83
	ArcParserBIT_XOR_ASSIGN  = 84
	ArcParserINCREMENT       = 85
	ArcParserDECREMENT       = 86
	ArcParserFAT_ARROW       = 87
	ArcParserPLUS            = 88
	ArcParserMINUS           = 89
	ArcParserSTAR            = 90
	ArcParserSLASH           = 91
	ArcParserPERCENT         = 92
	ArcParserLT              = 93
	ArcParserGT              = 94
	ArcParserNOT             = 95
	ArcParserAMP             = 96
	ArcParserBIT_OR          = 97
	ArcParserBIT_XOR         = 98
	ArcParserBIT_NOT         = 99
	ArcParserAT              = 100
	ArcParserASSIGN          = 101
	ArcParserLPAREN          = 102
	ArcParserRPAREN          = 103
	ArcParserLBRACE          = 104
	ArcParserRBRACE          = 105
	ArcParserLBRACKET        = 106
	ArcParserRBRACKET        = 107
	ArcParserCOMMA           = 108
	ArcParserCOLON           = 109
	ArcParserSEMICOLON       = 110
	ArcParserDOT             = 111
	ArcParserUNDERSCORE      = 112
	ArcParserBOOLEAN_LITERAL = 113
	ArcParserINTEGER_LITERAL = 114
	ArcParserFLOAT_LITERAL   = 115
	ArcParserSTRING_LITERAL  = 116
	ArcParserCHAR_LITERAL    = 117
	ArcParserIDENTIFIER      = 118
	ArcParserWS              = 119
	ArcParserLINE_COMMENT    = 120
	ArcParserBLOCK_COMMENT   = 121
	ArcParserEXTERN_WS       = 122
	ArcParserC_LANG          = 123
	ArcParserCPP_LANG        = 124
)
    ArcParser tokens.

const (
	ArcParserRULE_compilationUnit          = 0
	ArcParserRULE_importDecl               = 1
	ArcParserRULE_importSpec               = 2
	ArcParserRULE_namespaceDecl            = 3
	ArcParserRULE_topLevelDecl             = 4
	ArcParserRULE_externCDecl              = 5
	ArcParserRULE_externCMember            = 6
	ArcParserRULE_externCFunctionDecl      = 7
	ArcParserRULE_cCallingConvention       = 8
	ArcParserRULE_externCParameterList     = 9
	ArcParserRULE_externCParameter         = 10
	ArcParserRULE_externCConstDecl         = 11
	ArcParserRULE_externCTypeAlias         = 12
	ArcParserRULE_externCOpaqueStructDecl  = 13
	ArcParserRULE_externCppDecl            = 14
	ArcParserRULE_externCppMember          = 15
	ArcParserRULE_externCppNamespaceDecl   = 16
	ArcParserRULE_externNamespacePath      = 17
	ArcParserRULE_externCppFunctionDecl    = 18
	ArcParserRULE_cppCallingConvention     = 19
	ArcParserRULE_externCppParameterList   = 20
	ArcParserRULE_externCppParameter       = 21
	ArcParserRULE_externCppParamType       = 22
	ArcParserRULE_externCppConstDecl       = 23
	ArcParserRULE_externCppTypeAlias       = 24
	ArcParserRULE_externCppOpaqueClassDecl = 25
	ArcParserRULE_externCppClassDecl       = 26
	ArcParserRULE_externCppClassMember     = 27
	ArcParserRULE_externCppConstructorDecl = 28
	ArcParserRULE_externCppDestructorDecl  = 29
	ArcParserRULE_externCppMethodDecl      = 30
	ArcParserRULE_externCppMethodParams    = 31
	ArcParserRULE_externCppSelfParam       = 32
	ArcParserRULE_genericParams            = 33
	ArcParserRULE_genericParamList         = 34
	ArcParserRULE_genericParam             = 35
	ArcParserRULE_genericArgs              = 36
	ArcParserRULE_genericArgList           = 37
	ArcParserRULE_genericArg               = 38
	ArcParserRULE_functionDecl             = 39
	ArcParserRULE_returnType               = 40
	ArcParserRULE_typeList                 = 41
	ArcParserRULE_parameterList            = 42
	ArcParserRULE_parameter                = 43
	ArcParserRULE_structDecl               = 44
	ArcParserRULE_computeMarker            = 45
	ArcParserRULE_structMember             = 46
	ArcParserRULE_structField              = 47
	ArcParserRULE_initDecl                 = 48
	ArcParserRULE_classDecl                = 49
	ArcParserRULE_classMember              = 50
	ArcParserRULE_classField               = 51
	ArcParserRULE_enumDecl                 = 52
	ArcParserRULE_enumMember               = 53
	ArcParserRULE_methodDecl               = 54
	ArcParserRULE_mutatingDecl             = 55
	ArcParserRULE_deinitDecl               = 56
	ArcParserRULE_variableDecl             = 57
	ArcParserRULE_constDecl                = 58
	ArcParserRULE_tuplePattern             = 59
	ArcParserRULE_tupleType                = 60
	ArcParserRULE_type                     = 61
	ArcParserRULE_qualifiedType            = 62
	ArcParserRULE_functionType             = 63
	ArcParserRULE_arrayType                = 64
	ArcParserRULE_qualifiedIdentifier      = 65
	ArcParserRULE_primitiveType            = 66
	ArcParserRULE_pointerType              = 67
	ArcParserRULE_referenceType            = 68
	ArcParserRULE_block                    = 69
	ArcParserRULE_statement                = 70
	ArcParserRULE_assignmentStmt           = 71
	ArcParserRULE_assignmentOp             = 72
	ArcParserRULE_leftHandSide             = 73
	ArcParserRULE_expressionStmt           = 74
	ArcParserRULE_returnStmt               = 75
	ArcParserRULE_deferStmt                = 76
	ArcParserRULE_breakStmt                = 77
	ArcParserRULE_continueStmt             = 78
	ArcParserRULE_throwStmt                = 79
	ArcParserRULE_ifStmt                   = 80
	ArcParserRULE_forStmt                  = 81
	ArcParserRULE_switchStmt               = 82
	ArcParserRULE_switchCase               = 83
	ArcParserRULE_defaultCase              = 84
	ArcParserRULE_tryStmt                  = 85
	ArcParserRULE_exceptClause             = 86
	ArcParserRULE_finallyClause            = 87
	ArcParserRULE_expression               = 88
	ArcParserRULE_logicalOrExpression      = 89
	ArcParserRULE_logicalAndExpression     = 90
	ArcParserRULE_bitOrExpression          = 91
	ArcParserRULE_bitXorExpression         = 92
	ArcParserRULE_bitAndExpression         = 93
	ArcParserRULE_equalityExpression       = 94
	ArcParserRULE_relationalExpression     = 95
	ArcParserRULE_shiftExpression          = 96
	ArcParserRULE_rangeExpression          = 97
	ArcParserRULE_additiveExpression       = 98
	ArcParserRULE_multiplicativeExpression = 99
	ArcParserRULE_unaryExpression          = 100
	ArcParserRULE_postfixExpression        = 101
	ArcParserRULE_postfixOp                = 102
	ArcParserRULE_primaryExpression        = 103
	ArcParserRULE_computeExpression        = 104
	ArcParserRULE_computeContext           = 105
	ArcParserRULE_sizeofExpression         = 106
	ArcParserRULE_alignofExpression        = 107
	ArcParserRULE_literal                  = 108
	ArcParserRULE_initializerList          = 109
	ArcParserRULE_initializerEntry         = 110
	ArcParserRULE_structLiteral            = 111
	ArcParserRULE_fieldInit                = 112
	ArcParserRULE_argumentList             = 113
	ArcParserRULE_argument                 = 114
	ArcParserRULE_lambdaExpression         = 115
	ArcParserRULE_anonymousFuncExpression  = 116
	ArcParserRULE_lambdaParamList          = 117
	ArcParserRULE_lambdaParam              = 118
	ArcParserRULE_tupleExpression          = 119
)
    ArcParser rules.

const ArcLexerEXTERN_LANG_MODE = 1
    ArcLexerEXTERN_LANG_MODE is the ArcLexer mode.


VARIABLES

var ArcLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}
var ArcParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

FUNCTIONS

func ArcLexerInit()
    ArcLexerInit initializes any static state used to implement ArcLexer. By
    default the static state used to implement the lexer is lazily initialized
    during the first call to NewArcLexer(). You can call this function if you
    wish to initialize the static state ahead of time.

func ArcParserInit()
    ArcParserInit initializes any static state used to implement ArcParser. By
    default the static state used to implement the parser is lazily initialized
    during the first call to NewArcParser(). You can call this function if you
    wish to initialize the static state ahead of time.

func InitEmptyAdditiveExpressionContext(p *AdditiveExpressionContext)
func InitEmptyAlignofExpressionContext(p *AlignofExpressionContext)
func InitEmptyAnonymousFuncExpressionContext(p *AnonymousFuncExpressionContext)
func InitEmptyArgumentContext(p *ArgumentContext)
func InitEmptyArgumentListContext(p *ArgumentListContext)
func InitEmptyArrayTypeContext(p *ArrayTypeContext)
func InitEmptyAssignmentOpContext(p *AssignmentOpContext)
func InitEmptyAssignmentStmtContext(p *AssignmentStmtContext)
func InitEmptyBitAndExpressionContext(p *BitAndExpressionContext)
func InitEmptyBitOrExpressionContext(p *BitOrExpressionContext)
func InitEmptyBitXorExpressionContext(p *BitXorExpressionContext)
func InitEmptyBlockContext(p *BlockContext)
func InitEmptyBreakStmtContext(p *BreakStmtContext)
func InitEmptyCCallingConventionContext(p *CCallingConventionContext)
func InitEmptyClassDeclContext(p *ClassDeclContext)
func InitEmptyClassFieldContext(p *ClassFieldContext)
func InitEmptyClassMemberContext(p *ClassMemberContext)
func InitEmptyCompilationUnitContext(p *CompilationUnitContext)
func InitEmptyComputeContextContext(p *ComputeContextContext)
func InitEmptyComputeExpressionContext(p *ComputeExpressionContext)
func InitEmptyComputeMarkerContext(p *ComputeMarkerContext)
func InitEmptyConstDeclContext(p *ConstDeclContext)
func InitEmptyContinueStmtContext(p *ContinueStmtContext)
func InitEmptyCppCallingConventionContext(p *CppCallingConventionContext)
func InitEmptyDefaultCaseContext(p *DefaultCaseContext)
func InitEmptyDeferStmtContext(p *DeferStmtContext)
func InitEmptyDeinitDeclContext(p *DeinitDeclContext)
func InitEmptyEnumDeclContext(p *EnumDeclContext)
func InitEmptyEnumMemberContext(p *EnumMemberContext)
func InitEmptyEqualityExpressionContext(p *EqualityExpressionContext)
func InitEmptyExceptClauseContext(p *ExceptClauseContext)
func InitEmptyExpressionContext(p *ExpressionContext)
func InitEmptyExpressionStmtContext(p *ExpressionStmtContext)
func InitEmptyExternCConstDeclContext(p *ExternCConstDeclContext)
func InitEmptyExternCDeclContext(p *ExternCDeclContext)
func InitEmptyExternCFunctionDeclContext(p *ExternCFunctionDeclContext)
func InitEmptyExternCMemberContext(p *ExternCMemberContext)
func InitEmptyExternCOpaqueStructDeclContext(p *ExternCOpaqueStructDeclContext)
func InitEmptyExternCParameterContext(p *ExternCParameterContext)
func InitEmptyExternCParameterListContext(p *ExternCParameterListContext)
func InitEmptyExternCTypeAliasContext(p *ExternCTypeAliasContext)
func InitEmptyExternCppClassDeclContext(p *ExternCppClassDeclContext)
func InitEmptyExternCppClassMemberContext(p *ExternCppClassMemberContext)
func InitEmptyExternCppConstDeclContext(p *ExternCppConstDeclContext)
func InitEmptyExternCppConstructorDeclContext(p *ExternCppConstructorDeclContext)
func InitEmptyExternCppDeclContext(p *ExternCppDeclContext)
func InitEmptyExternCppDestructorDeclContext(p *ExternCppDestructorDeclContext)
func InitEmptyExternCppFunctionDeclContext(p *ExternCppFunctionDeclContext)
func InitEmptyExternCppMemberContext(p *ExternCppMemberContext)
func InitEmptyExternCppMethodDeclContext(p *ExternCppMethodDeclContext)
func InitEmptyExternCppMethodParamsContext(p *ExternCppMethodParamsContext)
func InitEmptyExternCppNamespaceDeclContext(p *ExternCppNamespaceDeclContext)
func InitEmptyExternCppOpaqueClassDeclContext(p *ExternCppOpaqueClassDeclContext)
func InitEmptyExternCppParamTypeContext(p *ExternCppParamTypeContext)
func InitEmptyExternCppParameterContext(p *ExternCppParameterContext)
func InitEmptyExternCppParameterListContext(p *ExternCppParameterListContext)
func InitEmptyExternCppSelfParamContext(p *ExternCppSelfParamContext)
func InitEmptyExternCppTypeAliasContext(p *ExternCppTypeAliasContext)
func InitEmptyExternNamespacePathContext(p *ExternNamespacePathContext)
func InitEmptyFieldInitContext(p *FieldInitContext)
func InitEmptyFinallyClauseContext(p *FinallyClauseContext)
func InitEmptyForStmtContext(p *ForStmtContext)
func InitEmptyFunctionDeclContext(p *FunctionDeclContext)
func InitEmptyFunctionTypeContext(p *FunctionTypeContext)
func InitEmptyGenericArgContext(p *GenericArgContext)
func InitEmptyGenericArgListContext(p *GenericArgListContext)
func InitEmptyGenericArgsContext(p *GenericArgsContext)
func InitEmptyGenericParamContext(p *GenericParamContext)
func InitEmptyGenericParamListContext(p *GenericParamListContext)
func InitEmptyGenericParamsContext(p *GenericParamsContext)
func InitEmptyIfStmtContext(p *IfStmtContext)
func InitEmptyImportDeclContext(p *ImportDeclContext)
func InitEmptyImportSpecContext(p *ImportSpecContext)
func InitEmptyInitDeclContext(p *InitDeclContext)
func InitEmptyInitializerEntryContext(p *InitializerEntryContext)
func InitEmptyInitializerListContext(p *InitializerListContext)
func InitEmptyLambdaExpressionContext(p *LambdaExpressionContext)
func InitEmptyLambdaParamContext(p *LambdaParamContext)
func InitEmptyLambdaParamListContext(p *LambdaParamListContext)
func InitEmptyLeftHandSideContext(p *LeftHandSideContext)
func InitEmptyLiteralContext(p *LiteralContext)
func InitEmptyLogicalAndExpressionContext(p *LogicalAndExpressionContext)
func InitEmptyLogicalOrExpressionContext(p *LogicalOrExpressionContext)
func InitEmptyMethodDeclContext(p *MethodDeclContext)
func InitEmptyMultiplicativeExpressionContext(p *MultiplicativeExpressionContext)
func InitEmptyMutatingDeclContext(p *MutatingDeclContext)
func InitEmptyNamespaceDeclContext(p *NamespaceDeclContext)
func InitEmptyParameterContext(p *ParameterContext)
func InitEmptyParameterListContext(p *ParameterListContext)
func InitEmptyPointerTypeContext(p *PointerTypeContext)
func InitEmptyPostfixExpressionContext(p *PostfixExpressionContext)
func InitEmptyPostfixOpContext(p *PostfixOpContext)
func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext)
func InitEmptyPrimitiveTypeContext(p *PrimitiveTypeContext)
func InitEmptyQualifiedIdentifierContext(p *QualifiedIdentifierContext)
func InitEmptyQualifiedTypeContext(p *QualifiedTypeContext)
func InitEmptyRangeExpressionContext(p *RangeExpressionContext)
func InitEmptyReferenceTypeContext(p *ReferenceTypeContext)
func InitEmptyRelationalExpressionContext(p *RelationalExpressionContext)
func InitEmptyReturnStmtContext(p *ReturnStmtContext)
func InitEmptyReturnTypeContext(p *ReturnTypeContext)
func InitEmptyShiftExpressionContext(p *ShiftExpressionContext)
func InitEmptySizeofExpressionContext(p *SizeofExpressionContext)
func InitEmptyStatementContext(p *StatementContext)
func InitEmptyStructDeclContext(p *StructDeclContext)
func InitEmptyStructFieldContext(p *StructFieldContext)
func InitEmptyStructLiteralContext(p *StructLiteralContext)
func InitEmptyStructMemberContext(p *StructMemberContext)
func InitEmptySwitchCaseContext(p *SwitchCaseContext)
func InitEmptySwitchStmtContext(p *SwitchStmtContext)
func InitEmptyThrowStmtContext(p *ThrowStmtContext)
func InitEmptyTopLevelDeclContext(p *TopLevelDeclContext)
func InitEmptyTryStmtContext(p *TryStmtContext)
func InitEmptyTupleExpressionContext(p *TupleExpressionContext)
func InitEmptyTuplePatternContext(p *TuplePatternContext)
func InitEmptyTupleTypeContext(p *TupleTypeContext)
func InitEmptyTypeContext(p *TypeContext)
func InitEmptyTypeListContext(p *TypeListContext)
func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext)
func InitEmptyVariableDeclContext(p *VariableDeclContext)

TYPES

type AdditiveExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *AdditiveExpressionContext) AllMINUS() []antlr.TerminalNode

func (s *AdditiveExpressionContext) AllMultiplicativeExpression() []IMultiplicativeExpressionContext

func (s *AdditiveExpressionContext) AllPLUS() []antlr.TerminalNode

func (s *AdditiveExpressionContext) GetParser() antlr.Parser

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext

func (*AdditiveExpressionContext) IsAdditiveExpressionContext()

func (s *AdditiveExpressionContext) MINUS(i int) antlr.TerminalNode

func (s *AdditiveExpressionContext) MultiplicativeExpression(i int) IMultiplicativeExpressionContext

func (s *AdditiveExpressionContext) PLUS(i int) antlr.TerminalNode

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type AlignofExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewAlignofExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlignofExpressionContext

func NewEmptyAlignofExpressionContext() *AlignofExpressionContext

func (s *AlignofExpressionContext) ALIGNOF() antlr.TerminalNode

func (s *AlignofExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *AlignofExpressionContext) GT() antlr.TerminalNode

func (s *AlignofExpressionContext) GetParser() antlr.Parser

func (s *AlignofExpressionContext) GetRuleContext() antlr.RuleContext

func (*AlignofExpressionContext) IsAlignofExpressionContext()

func (s *AlignofExpressionContext) LT() antlr.TerminalNode

func (s *AlignofExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *AlignofExpressionContext) Type_() ITypeContext

type AnonymousFuncExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewAnonymousFuncExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnonymousFuncExpressionContext

func NewEmptyAnonymousFuncExpressionContext() *AnonymousFuncExpressionContext

func (s *AnonymousFuncExpressionContext) ASYNC() antlr.TerminalNode

func (s *AnonymousFuncExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *AnonymousFuncExpressionContext) Block() IBlockContext

func (s *AnonymousFuncExpressionContext) CONTAINER() antlr.TerminalNode

func (s *AnonymousFuncExpressionContext) FUNC() antlr.TerminalNode

func (s *AnonymousFuncExpressionContext) GenericParams() IGenericParamsContext

func (s *AnonymousFuncExpressionContext) GetParser() antlr.Parser

func (s *AnonymousFuncExpressionContext) GetRuleContext() antlr.RuleContext

func (*AnonymousFuncExpressionContext) IsAnonymousFuncExpressionContext()

func (s *AnonymousFuncExpressionContext) LPAREN() antlr.TerminalNode

func (s *AnonymousFuncExpressionContext) PROCESS() antlr.TerminalNode

func (s *AnonymousFuncExpressionContext) ParameterList() IParameterListContext

func (s *AnonymousFuncExpressionContext) RPAREN() antlr.TerminalNode

func (s *AnonymousFuncExpressionContext) ReturnType() IReturnTypeContext

func (s *AnonymousFuncExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ArcLexer struct {
	*antlr.BaseLexer

	// Has unexported fields.
}

func NewArcLexer(input antlr.CharStream) *ArcLexer
    NewArcLexer produces a new lexer instance for the optional input
    antlr.CharStream.

type ArcParser struct {
	*antlr.BaseParser
}

func NewArcParser(input antlr.TokenStream) *ArcParser
    NewArcParser produces a new parser instance for the optional input
    antlr.TokenStream.

func (p *ArcParser) AdditiveExpression() (localctx IAdditiveExpressionContext)

func (p *ArcParser) AlignofExpression() (localctx IAlignofExpressionContext)

func (p *ArcParser) AnonymousFuncExpression() (localctx IAnonymousFuncExpressionContext)

func (p *ArcParser) Argument() (localctx IArgumentContext)

func (p *ArcParser) ArgumentList() (localctx IArgumentListContext)

func (p *ArcParser) ArrayType() (localctx IArrayTypeContext)

func (p *ArcParser) AssignmentOp() (localctx IAssignmentOpContext)

func (p *ArcParser) AssignmentStmt() (localctx IAssignmentStmtContext)

func (p *ArcParser) BitAndExpression() (localctx IBitAndExpressionContext)

func (p *ArcParser) BitOrExpression() (localctx IBitOrExpressionContext)

func (p *ArcParser) BitXorExpression() (localctx IBitXorExpressionContext)

func (p *ArcParser) Block() (localctx IBlockContext)

func (p *ArcParser) BreakStmt() (localctx IBreakStmtContext)

func (p *ArcParser) CCallingConvention() (localctx ICCallingConventionContext)

func (p *ArcParser) ClassDecl() (localctx IClassDeclContext)

func (p *ArcParser) ClassField() (localctx IClassFieldContext)

func (p *ArcParser) ClassMember() (localctx IClassMemberContext)

func (p *ArcParser) CompilationUnit() (localctx ICompilationUnitContext)

func (p *ArcParser) ComputeContext() (localctx IComputeContextContext)

func (p *ArcParser) ComputeExpression() (localctx IComputeExpressionContext)

func (p *ArcParser) ComputeMarker() (localctx IComputeMarkerContext)

func (p *ArcParser) ConstDecl() (localctx IConstDeclContext)

func (p *ArcParser) ContinueStmt() (localctx IContinueStmtContext)

func (p *ArcParser) CppCallingConvention() (localctx ICppCallingConventionContext)

func (p *ArcParser) DefaultCase() (localctx IDefaultCaseContext)

func (p *ArcParser) DeferStmt() (localctx IDeferStmtContext)

func (p *ArcParser) DeinitDecl() (localctx IDeinitDeclContext)

func (p *ArcParser) EnumDecl() (localctx IEnumDeclContext)

func (p *ArcParser) EnumMember() (localctx IEnumMemberContext)

func (p *ArcParser) EqualityExpression() (localctx IEqualityExpressionContext)

func (p *ArcParser) ExceptClause() (localctx IExceptClauseContext)

func (p *ArcParser) Expression() (localctx IExpressionContext)

func (p *ArcParser) ExpressionStmt() (localctx IExpressionStmtContext)

func (p *ArcParser) ExternCConstDecl() (localctx IExternCConstDeclContext)

func (p *ArcParser) ExternCDecl() (localctx IExternCDeclContext)

func (p *ArcParser) ExternCFunctionDecl() (localctx IExternCFunctionDeclContext)

func (p *ArcParser) ExternCMember() (localctx IExternCMemberContext)

func (p *ArcParser) ExternCOpaqueStructDecl() (localctx IExternCOpaqueStructDeclContext)

func (p *ArcParser) ExternCParameter() (localctx IExternCParameterContext)

func (p *ArcParser) ExternCParameterList() (localctx IExternCParameterListContext)

func (p *ArcParser) ExternCTypeAlias() (localctx IExternCTypeAliasContext)

func (p *ArcParser) ExternCppClassDecl() (localctx IExternCppClassDeclContext)

func (p *ArcParser) ExternCppClassMember() (localctx IExternCppClassMemberContext)

func (p *ArcParser) ExternCppConstDecl() (localctx IExternCppConstDeclContext)

func (p *ArcParser) ExternCppConstructorDecl() (localctx IExternCppConstructorDeclContext)

func (p *ArcParser) ExternCppDecl() (localctx IExternCppDeclContext)

func (p *ArcParser) ExternCppDestructorDecl() (localctx IExternCppDestructorDeclContext)

func (p *ArcParser) ExternCppFunctionDecl() (localctx IExternCppFunctionDeclContext)

func (p *ArcParser) ExternCppMember() (localctx IExternCppMemberContext)

func (p *ArcParser) ExternCppMethodDecl() (localctx IExternCppMethodDeclContext)

func (p *ArcParser) ExternCppMethodParams() (localctx IExternCppMethodParamsContext)

func (p *ArcParser) ExternCppNamespaceDecl() (localctx IExternCppNamespaceDeclContext)

func (p *ArcParser) ExternCppOpaqueClassDecl() (localctx IExternCppOpaqueClassDeclContext)

func (p *ArcParser) ExternCppParamType() (localctx IExternCppParamTypeContext)

func (p *ArcParser) ExternCppParameter() (localctx IExternCppParameterContext)

func (p *ArcParser) ExternCppParameterList() (localctx IExternCppParameterListContext)

func (p *ArcParser) ExternCppSelfParam() (localctx IExternCppSelfParamContext)

func (p *ArcParser) ExternCppTypeAlias() (localctx IExternCppTypeAliasContext)

func (p *ArcParser) ExternNamespacePath() (localctx IExternNamespacePathContext)

func (p *ArcParser) FieldInit() (localctx IFieldInitContext)

func (p *ArcParser) FinallyClause() (localctx IFinallyClauseContext)

func (p *ArcParser) ForStmt() (localctx IForStmtContext)

func (p *ArcParser) FunctionDecl() (localctx IFunctionDeclContext)

func (p *ArcParser) FunctionType() (localctx IFunctionTypeContext)

func (p *ArcParser) GenericArg() (localctx IGenericArgContext)

func (p *ArcParser) GenericArgList() (localctx IGenericArgListContext)

func (p *ArcParser) GenericArgs() (localctx IGenericArgsContext)

func (p *ArcParser) GenericParam() (localctx IGenericParamContext)

func (p *ArcParser) GenericParamList() (localctx IGenericParamListContext)

func (p *ArcParser) GenericParams() (localctx IGenericParamsContext)

func (p *ArcParser) IfStmt() (localctx IIfStmtContext)

func (p *ArcParser) ImportDecl() (localctx IImportDeclContext)

func (p *ArcParser) ImportSpec() (localctx IImportSpecContext)

func (p *ArcParser) InitDecl() (localctx IInitDeclContext)

func (p *ArcParser) InitializerEntry() (localctx IInitializerEntryContext)

func (p *ArcParser) InitializerList() (localctx IInitializerListContext)

func (p *ArcParser) LambdaExpression() (localctx ILambdaExpressionContext)

func (p *ArcParser) LambdaParam() (localctx ILambdaParamContext)

func (p *ArcParser) LambdaParamList() (localctx ILambdaParamListContext)

func (p *ArcParser) LeftHandSide() (localctx ILeftHandSideContext)

func (p *ArcParser) Literal() (localctx ILiteralContext)

func (p *ArcParser) LogicalAndExpression() (localctx ILogicalAndExpressionContext)

func (p *ArcParser) LogicalOrExpression() (localctx ILogicalOrExpressionContext)

func (p *ArcParser) MethodDecl() (localctx IMethodDeclContext)

func (p *ArcParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext)

func (p *ArcParser) MutatingDecl() (localctx IMutatingDeclContext)

func (p *ArcParser) NamespaceDecl() (localctx INamespaceDeclContext)

func (p *ArcParser) Parameter() (localctx IParameterContext)

func (p *ArcParser) ParameterList() (localctx IParameterListContext)

func (p *ArcParser) PointerType() (localctx IPointerTypeContext)

func (p *ArcParser) PostfixExpression() (localctx IPostfixExpressionContext)

func (p *ArcParser) PostfixOp() (localctx IPostfixOpContext)

func (p *ArcParser) PrimaryExpression() (localctx IPrimaryExpressionContext)

func (p *ArcParser) PrimitiveType() (localctx IPrimitiveTypeContext)

func (p *ArcParser) QualifiedIdentifier() (localctx IQualifiedIdentifierContext)

func (p *ArcParser) QualifiedType() (localctx IQualifiedTypeContext)

func (p *ArcParser) RangeExpression() (localctx IRangeExpressionContext)

func (p *ArcParser) ReferenceType() (localctx IReferenceTypeContext)

func (p *ArcParser) RelationalExpression() (localctx IRelationalExpressionContext)

func (p *ArcParser) ReturnStmt() (localctx IReturnStmtContext)

func (p *ArcParser) ReturnType() (localctx IReturnTypeContext)

func (p *ArcParser) ShiftExpression() (localctx IShiftExpressionContext)

func (p *ArcParser) SizeofExpression() (localctx ISizeofExpressionContext)

func (p *ArcParser) Statement() (localctx IStatementContext)

func (p *ArcParser) StructDecl() (localctx IStructDeclContext)

func (p *ArcParser) StructField() (localctx IStructFieldContext)

func (p *ArcParser) StructLiteral() (localctx IStructLiteralContext)

func (p *ArcParser) StructMember() (localctx IStructMemberContext)

func (p *ArcParser) SwitchCase() (localctx ISwitchCaseContext)

func (p *ArcParser) SwitchStmt() (localctx ISwitchStmtContext)

func (p *ArcParser) ThrowStmt() (localctx IThrowStmtContext)

func (p *ArcParser) TopLevelDecl() (localctx ITopLevelDeclContext)

func (p *ArcParser) TryStmt() (localctx ITryStmtContext)

func (p *ArcParser) TupleExpression() (localctx ITupleExpressionContext)

func (p *ArcParser) TuplePattern() (localctx ITuplePatternContext)

func (p *ArcParser) TupleType() (localctx ITupleTypeContext)

func (p *ArcParser) TypeList() (localctx ITypeListContext)

func (p *ArcParser) Type_() (localctx ITypeContext)

func (p *ArcParser) UnaryExpression() (localctx IUnaryExpressionContext)

func (p *ArcParser) VariableDecl() (localctx IVariableDeclContext)

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
    ArcParserListener is a complete listener for a parse tree produced by
    ArcParser.

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

	// Visit a parse tree produced by ArcParser#externCOpaqueStructDecl.
	VisitExternCOpaqueStructDecl(ctx *ExternCOpaqueStructDeclContext) interface{}

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

	// Visit a parse tree produced by ArcParser#externCppOpaqueClassDecl.
	VisitExternCppOpaqueClassDecl(ctx *ExternCppOpaqueClassDeclContext) interface{}

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

	// Visit a parse tree produced by ArcParser#functionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#returnType.
	VisitReturnType(ctx *ReturnTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by ArcParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by ArcParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by ArcParser#structDecl.
	VisitStructDecl(ctx *StructDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#computeMarker.
	VisitComputeMarker(ctx *ComputeMarkerContext) interface{}

	// Visit a parse tree produced by ArcParser#structMember.
	VisitStructMember(ctx *StructMemberContext) interface{}

	// Visit a parse tree produced by ArcParser#structField.
	VisitStructField(ctx *StructFieldContext) interface{}

	// Visit a parse tree produced by ArcParser#initDecl.
	VisitInitDecl(ctx *InitDeclContext) interface{}

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

	// Visit a parse tree produced by ArcParser#methodDecl.
	VisitMethodDecl(ctx *MethodDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#mutatingDecl.
	VisitMutatingDecl(ctx *MutatingDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#deinitDecl.
	VisitDeinitDecl(ctx *DeinitDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#variableDecl.
	VisitVariableDecl(ctx *VariableDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#constDecl.
	VisitConstDecl(ctx *ConstDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#tuplePattern.
	VisitTuplePattern(ctx *TuplePatternContext) interface{}

	// Visit a parse tree produced by ArcParser#tupleType.
	VisitTupleType(ctx *TupleTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by ArcParser#qualifiedType.
	VisitQualifiedType(ctx *QualifiedTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#functionType.
	VisitFunctionType(ctx *FunctionTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#qualifiedIdentifier.
	VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{}

	// Visit a parse tree produced by ArcParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#pointerType.
	VisitPointerType(ctx *PointerTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#referenceType.
	VisitReferenceType(ctx *ReferenceTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ArcParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ArcParser#assignmentStmt.
	VisitAssignmentStmt(ctx *AssignmentStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#assignmentOp.
	VisitAssignmentOp(ctx *AssignmentOpContext) interface{}

	// Visit a parse tree produced by ArcParser#leftHandSide.
	VisitLeftHandSide(ctx *LeftHandSideContext) interface{}

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

	// Visit a parse tree produced by ArcParser#throwStmt.
	VisitThrowStmt(ctx *ThrowStmtContext) interface{}

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

	// Visit a parse tree produced by ArcParser#tryStmt.
	VisitTryStmt(ctx *TryStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#exceptClause.
	VisitExceptClause(ctx *ExceptClauseContext) interface{}

	// Visit a parse tree produced by ArcParser#finallyClause.
	VisitFinallyClause(ctx *FinallyClauseContext) interface{}

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

	// Visit a parse tree produced by ArcParser#computeExpression.
	VisitComputeExpression(ctx *ComputeExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#computeContext.
	VisitComputeContext(ctx *ComputeContextContext) interface{}

	// Visit a parse tree produced by ArcParser#sizeofExpression.
	VisitSizeofExpression(ctx *SizeofExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#alignofExpression.
	VisitAlignofExpression(ctx *AlignofExpressionContext) interface{}

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

	// Visit a parse tree produced by ArcParser#lambdaParamList.
	VisitLambdaParamList(ctx *LambdaParamListContext) interface{}

	// Visit a parse tree produced by ArcParser#lambdaParam.
	VisitLambdaParam(ctx *LambdaParamContext) interface{}

	// Visit a parse tree produced by ArcParser#tupleExpression.
	VisitTupleExpression(ctx *TupleExpressionContext) interface{}
}
    A complete Visitor for a parse tree produced by ArcParser.

type ArgumentContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext

func NewEmptyArgumentContext() *ArgumentContext

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ArgumentContext) AnonymousFuncExpression() IAnonymousFuncExpressionContext

func (s *ArgumentContext) Expression() IExpressionContext

func (s *ArgumentContext) GetParser() antlr.Parser

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext

func (*ArgumentContext) IsArgumentContext()

func (s *ArgumentContext) LambdaExpression() ILambdaExpressionContext

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext

func NewEmptyArgumentListContext() *ArgumentListContext

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ArgumentListContext) AllArgument() []IArgumentContext

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode

func (s *ArgumentListContext) Argument(i int) IArgumentContext

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode

func (s *ArgumentListContext) GetParser() antlr.Parser

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext

func (*ArgumentListContext) IsArgumentListContext()

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ArrayTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext

func NewEmptyArrayTypeContext() *ArrayTypeContext

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ArrayTypeContext) Expression() IExpressionContext

func (s *ArrayTypeContext) GetParser() antlr.Parser

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext

func (*ArrayTypeContext) IsArrayTypeContext()

func (s *ArrayTypeContext) LBRACKET() antlr.TerminalNode

func (s *ArrayTypeContext) RBRACKET() antlr.TerminalNode

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ArrayTypeContext) Type_() ITypeContext

type AssignmentOpContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewAssignmentOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOpContext

func NewEmptyAssignmentOpContext() *AssignmentOpContext

func (s *AssignmentOpContext) ASSIGN() antlr.TerminalNode

func (s *AssignmentOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *AssignmentOpContext) BIT_AND_ASSIGN() antlr.TerminalNode

func (s *AssignmentOpContext) BIT_OR_ASSIGN() antlr.TerminalNode

func (s *AssignmentOpContext) BIT_XOR_ASSIGN() antlr.TerminalNode

func (s *AssignmentOpContext) GE() antlr.TerminalNode

func (s *AssignmentOpContext) GT() antlr.TerminalNode

func (s *AssignmentOpContext) GetParser() antlr.Parser

func (s *AssignmentOpContext) GetRuleContext() antlr.RuleContext

func (*AssignmentOpContext) IsAssignmentOpContext()

func (s *AssignmentOpContext) LE() antlr.TerminalNode

func (s *AssignmentOpContext) LT() antlr.TerminalNode

func (s *AssignmentOpContext) MINUS_ASSIGN() antlr.TerminalNode

func (s *AssignmentOpContext) PERCENT_ASSIGN() antlr.TerminalNode

func (s *AssignmentOpContext) PLUS_ASSIGN() antlr.TerminalNode

func (s *AssignmentOpContext) SLASH_ASSIGN() antlr.TerminalNode

func (s *AssignmentOpContext) STAR_ASSIGN() antlr.TerminalNode

func (s *AssignmentOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type AssignmentStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewAssignmentStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStmtContext

func NewEmptyAssignmentStmtContext() *AssignmentStmtContext

func (s *AssignmentStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *AssignmentStmtContext) AssignmentOp() IAssignmentOpContext

func (s *AssignmentStmtContext) Expression() IExpressionContext

func (s *AssignmentStmtContext) GetParser() antlr.Parser

func (s *AssignmentStmtContext) GetRuleContext() antlr.RuleContext

func (*AssignmentStmtContext) IsAssignmentStmtContext()

func (s *AssignmentStmtContext) LeftHandSide() ILeftHandSideContext

func (s *AssignmentStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type BaseArcParserListener struct{}
    BaseArcParserListener is a complete listener for a parse tree produced by
    ArcParser.

func (s *BaseArcParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext)
    EnterAdditiveExpression is called when production additiveExpression is
    entered.

func (s *BaseArcParserListener) EnterAlignofExpression(ctx *AlignofExpressionContext)
    EnterAlignofExpression is called when production alignofExpression is
    entered.

func (s *BaseArcParserListener) EnterAnonymousFuncExpression(ctx *AnonymousFuncExpressionContext)
    EnterAnonymousFuncExpression is called when production
    anonymousFuncExpression is entered.

func (s *BaseArcParserListener) EnterArgument(ctx *ArgumentContext)
    EnterArgument is called when production argument is entered.

func (s *BaseArcParserListener) EnterArgumentList(ctx *ArgumentListContext)
    EnterArgumentList is called when production argumentList is entered.

func (s *BaseArcParserListener) EnterArrayType(ctx *ArrayTypeContext)
    EnterArrayType is called when production arrayType is entered.

func (s *BaseArcParserListener) EnterAssignmentOp(ctx *AssignmentOpContext)
    EnterAssignmentOp is called when production assignmentOp is entered.

func (s *BaseArcParserListener) EnterAssignmentStmt(ctx *AssignmentStmtContext)
    EnterAssignmentStmt is called when production assignmentStmt is entered.

func (s *BaseArcParserListener) EnterBitAndExpression(ctx *BitAndExpressionContext)
    EnterBitAndExpression is called when production bitAndExpression is entered.

func (s *BaseArcParserListener) EnterBitOrExpression(ctx *BitOrExpressionContext)
    EnterBitOrExpression is called when production bitOrExpression is entered.

func (s *BaseArcParserListener) EnterBitXorExpression(ctx *BitXorExpressionContext)
    EnterBitXorExpression is called when production bitXorExpression is entered.

func (s *BaseArcParserListener) EnterBlock(ctx *BlockContext)
    EnterBlock is called when production block is entered.

func (s *BaseArcParserListener) EnterBreakStmt(ctx *BreakStmtContext)
    EnterBreakStmt is called when production breakStmt is entered.

func (s *BaseArcParserListener) EnterCCallingConvention(ctx *CCallingConventionContext)
    EnterCCallingConvention is called when production cCallingConvention is
    entered.

func (s *BaseArcParserListener) EnterClassDecl(ctx *ClassDeclContext)
    EnterClassDecl is called when production classDecl is entered.

func (s *BaseArcParserListener) EnterClassField(ctx *ClassFieldContext)
    EnterClassField is called when production classField is entered.

func (s *BaseArcParserListener) EnterClassMember(ctx *ClassMemberContext)
    EnterClassMember is called when production classMember is entered.

func (s *BaseArcParserListener) EnterCompilationUnit(ctx *CompilationUnitContext)
    EnterCompilationUnit is called when production compilationUnit is entered.

func (s *BaseArcParserListener) EnterComputeContext(ctx *ComputeContextContext)
    EnterComputeContext is called when production computeContext is entered.

func (s *BaseArcParserListener) EnterComputeExpression(ctx *ComputeExpressionContext)
    EnterComputeExpression is called when production computeExpression is
    entered.

func (s *BaseArcParserListener) EnterComputeMarker(ctx *ComputeMarkerContext)
    EnterComputeMarker is called when production computeMarker is entered.

func (s *BaseArcParserListener) EnterConstDecl(ctx *ConstDeclContext)
    EnterConstDecl is called when production constDecl is entered.

func (s *BaseArcParserListener) EnterContinueStmt(ctx *ContinueStmtContext)
    EnterContinueStmt is called when production continueStmt is entered.

func (s *BaseArcParserListener) EnterCppCallingConvention(ctx *CppCallingConventionContext)
    EnterCppCallingConvention is called when production cppCallingConvention is
    entered.

func (s *BaseArcParserListener) EnterDefaultCase(ctx *DefaultCaseContext)
    EnterDefaultCase is called when production defaultCase is entered.

func (s *BaseArcParserListener) EnterDeferStmt(ctx *DeferStmtContext)
    EnterDeferStmt is called when production deferStmt is entered.

func (s *BaseArcParserListener) EnterDeinitDecl(ctx *DeinitDeclContext)
    EnterDeinitDecl is called when production deinitDecl is entered.

func (s *BaseArcParserListener) EnterEnumDecl(ctx *EnumDeclContext)
    EnterEnumDecl is called when production enumDecl is entered.

func (s *BaseArcParserListener) EnterEnumMember(ctx *EnumMemberContext)
    EnterEnumMember is called when production enumMember is entered.

func (s *BaseArcParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext)
    EnterEqualityExpression is called when production equalityExpression is
    entered.

func (s *BaseArcParserListener) EnterEveryRule(ctx antlr.ParserRuleContext)
    EnterEveryRule is called when any rule is entered.

func (s *BaseArcParserListener) EnterExceptClause(ctx *ExceptClauseContext)
    EnterExceptClause is called when production exceptClause is entered.

func (s *BaseArcParserListener) EnterExpression(ctx *ExpressionContext)
    EnterExpression is called when production expression is entered.

func (s *BaseArcParserListener) EnterExpressionStmt(ctx *ExpressionStmtContext)
    EnterExpressionStmt is called when production expressionStmt is entered.

func (s *BaseArcParserListener) EnterExternCConstDecl(ctx *ExternCConstDeclContext)
    EnterExternCConstDecl is called when production externCConstDecl is entered.

func (s *BaseArcParserListener) EnterExternCDecl(ctx *ExternCDeclContext)
    EnterExternCDecl is called when production externCDecl is entered.

func (s *BaseArcParserListener) EnterExternCFunctionDecl(ctx *ExternCFunctionDeclContext)
    EnterExternCFunctionDecl is called when production externCFunctionDecl is
    entered.

func (s *BaseArcParserListener) EnterExternCMember(ctx *ExternCMemberContext)
    EnterExternCMember is called when production externCMember is entered.

func (s *BaseArcParserListener) EnterExternCOpaqueStructDecl(ctx *ExternCOpaqueStructDeclContext)
    EnterExternCOpaqueStructDecl is called when production
    externCOpaqueStructDecl is entered.

func (s *BaseArcParserListener) EnterExternCParameter(ctx *ExternCParameterContext)
    EnterExternCParameter is called when production externCParameter is entered.

func (s *BaseArcParserListener) EnterExternCParameterList(ctx *ExternCParameterListContext)
    EnterExternCParameterList is called when production externCParameterList is
    entered.

func (s *BaseArcParserListener) EnterExternCTypeAlias(ctx *ExternCTypeAliasContext)
    EnterExternCTypeAlias is called when production externCTypeAlias is entered.

func (s *BaseArcParserListener) EnterExternCppClassDecl(ctx *ExternCppClassDeclContext)
    EnterExternCppClassDecl is called when production externCppClassDecl is
    entered.

func (s *BaseArcParserListener) EnterExternCppClassMember(ctx *ExternCppClassMemberContext)
    EnterExternCppClassMember is called when production externCppClassMember is
    entered.

func (s *BaseArcParserListener) EnterExternCppConstDecl(ctx *ExternCppConstDeclContext)
    EnterExternCppConstDecl is called when production externCppConstDecl is
    entered.

func (s *BaseArcParserListener) EnterExternCppConstructorDecl(ctx *ExternCppConstructorDeclContext)
    EnterExternCppConstructorDecl is called when production
    externCppConstructorDecl is entered.

func (s *BaseArcParserListener) EnterExternCppDecl(ctx *ExternCppDeclContext)
    EnterExternCppDecl is called when production externCppDecl is entered.

func (s *BaseArcParserListener) EnterExternCppDestructorDecl(ctx *ExternCppDestructorDeclContext)
    EnterExternCppDestructorDecl is called when production
    externCppDestructorDecl is entered.

func (s *BaseArcParserListener) EnterExternCppFunctionDecl(ctx *ExternCppFunctionDeclContext)
    EnterExternCppFunctionDecl is called when production externCppFunctionDecl
    is entered.

func (s *BaseArcParserListener) EnterExternCppMember(ctx *ExternCppMemberContext)
    EnterExternCppMember is called when production externCppMember is entered.

func (s *BaseArcParserListener) EnterExternCppMethodDecl(ctx *ExternCppMethodDeclContext)
    EnterExternCppMethodDecl is called when production externCppMethodDecl is
    entered.

func (s *BaseArcParserListener) EnterExternCppMethodParams(ctx *ExternCppMethodParamsContext)
    EnterExternCppMethodParams is called when production externCppMethodParams
    is entered.

func (s *BaseArcParserListener) EnterExternCppNamespaceDecl(ctx *ExternCppNamespaceDeclContext)
    EnterExternCppNamespaceDecl is called when production externCppNamespaceDecl
    is entered.

func (s *BaseArcParserListener) EnterExternCppOpaqueClassDecl(ctx *ExternCppOpaqueClassDeclContext)
    EnterExternCppOpaqueClassDecl is called when production
    externCppOpaqueClassDecl is entered.

func (s *BaseArcParserListener) EnterExternCppParamType(ctx *ExternCppParamTypeContext)
    EnterExternCppParamType is called when production externCppParamType is
    entered.

func (s *BaseArcParserListener) EnterExternCppParameter(ctx *ExternCppParameterContext)
    EnterExternCppParameter is called when production externCppParameter is
    entered.

func (s *BaseArcParserListener) EnterExternCppParameterList(ctx *ExternCppParameterListContext)
    EnterExternCppParameterList is called when production externCppParameterList
    is entered.

func (s *BaseArcParserListener) EnterExternCppSelfParam(ctx *ExternCppSelfParamContext)
    EnterExternCppSelfParam is called when production externCppSelfParam is
    entered.

func (s *BaseArcParserListener) EnterExternCppTypeAlias(ctx *ExternCppTypeAliasContext)
    EnterExternCppTypeAlias is called when production externCppTypeAlias is
    entered.

func (s *BaseArcParserListener) EnterExternNamespacePath(ctx *ExternNamespacePathContext)
    EnterExternNamespacePath is called when production externNamespacePath is
    entered.

func (s *BaseArcParserListener) EnterFieldInit(ctx *FieldInitContext)
    EnterFieldInit is called when production fieldInit is entered.

func (s *BaseArcParserListener) EnterFinallyClause(ctx *FinallyClauseContext)
    EnterFinallyClause is called when production finallyClause is entered.

func (s *BaseArcParserListener) EnterForStmt(ctx *ForStmtContext)
    EnterForStmt is called when production forStmt is entered.

func (s *BaseArcParserListener) EnterFunctionDecl(ctx *FunctionDeclContext)
    EnterFunctionDecl is called when production functionDecl is entered.

func (s *BaseArcParserListener) EnterFunctionType(ctx *FunctionTypeContext)
    EnterFunctionType is called when production functionType is entered.

func (s *BaseArcParserListener) EnterGenericArg(ctx *GenericArgContext)
    EnterGenericArg is called when production genericArg is entered.

func (s *BaseArcParserListener) EnterGenericArgList(ctx *GenericArgListContext)
    EnterGenericArgList is called when production genericArgList is entered.

func (s *BaseArcParserListener) EnterGenericArgs(ctx *GenericArgsContext)
    EnterGenericArgs is called when production genericArgs is entered.

func (s *BaseArcParserListener) EnterGenericParamList(ctx *GenericParamListContext)
    EnterGenericParamList is called when production genericParamList is entered.

func (s *BaseArcParserListener) EnterGenericParams(ctx *GenericParamsContext)
    EnterGenericParams is called when production genericParams is entered.

func (s *BaseArcParserListener) EnterIfStmt(ctx *IfStmtContext)
    EnterIfStmt is called when production ifStmt is entered.

func (s *BaseArcParserListener) EnterImportDecl(ctx *ImportDeclContext)
    EnterImportDecl is called when production importDecl is entered.

func (s *BaseArcParserListener) EnterImportSpec(ctx *ImportSpecContext)
    EnterImportSpec is called when production importSpec is entered.

func (s *BaseArcParserListener) EnterInitDecl(ctx *InitDeclContext)
    EnterInitDecl is called when production initDecl is entered.

func (s *BaseArcParserListener) EnterInitializerEntry(ctx *InitializerEntryContext)
    EnterInitializerEntry is called when production initializerEntry is entered.

func (s *BaseArcParserListener) EnterInitializerList(ctx *InitializerListContext)
    EnterInitializerList is called when production initializerList is entered.

func (s *BaseArcParserListener) EnterLambdaExpression(ctx *LambdaExpressionContext)
    EnterLambdaExpression is called when production lambdaExpression is entered.

func (s *BaseArcParserListener) EnterLambdaParam(ctx *LambdaParamContext)
    EnterLambdaParam is called when production lambdaParam is entered.

func (s *BaseArcParserListener) EnterLambdaParamList(ctx *LambdaParamListContext)
    EnterLambdaParamList is called when production lambdaParamList is entered.

func (s *BaseArcParserListener) EnterLeftHandSide(ctx *LeftHandSideContext)
    EnterLeftHandSide is called when production leftHandSide is entered.

func (s *BaseArcParserListener) EnterLiteral(ctx *LiteralContext)
    EnterLiteral is called when production literal is entered.

func (s *BaseArcParserListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext)
    EnterLogicalAndExpression is called when production logicalAndExpression is
    entered.

func (s *BaseArcParserListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext)
    EnterLogicalOrExpression is called when production logicalOrExpression is
    entered.

func (s *BaseArcParserListener) EnterMethodDecl(ctx *MethodDeclContext)
    EnterMethodDecl is called when production methodDecl is entered.

func (s *BaseArcParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext)
    EnterMultiplicativeExpression is called when production
    multiplicativeExpression is entered.

func (s *BaseArcParserListener) EnterMutatingDecl(ctx *MutatingDeclContext)
    EnterMutatingDecl is called when production mutatingDecl is entered.

func (s *BaseArcParserListener) EnterNamespaceDecl(ctx *NamespaceDeclContext)
    EnterNamespaceDecl is called when production namespaceDecl is entered.

func (s *BaseArcParserListener) EnterParameter(ctx *ParameterContext)
    EnterParameter is called when production parameter is entered.

func (s *BaseArcParserListener) EnterParameterList(ctx *ParameterListContext)
    EnterParameterList is called when production parameterList is entered.

func (s *BaseArcParserListener) EnterPointerType(ctx *PointerTypeContext)
    EnterPointerType is called when production pointerType is entered.

func (s *BaseArcParserListener) EnterPostfixExpression(ctx *PostfixExpressionContext)
    EnterPostfixExpression is called when production postfixExpression is
    entered.

func (s *BaseArcParserListener) EnterPostfixOp(ctx *PostfixOpContext)
    EnterPostfixOp is called when production postfixOp is entered.

func (s *BaseArcParserListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext)
    EnterPrimaryExpression is called when production primaryExpression is
    entered.

func (s *BaseArcParserListener) EnterPrimitiveType(ctx *PrimitiveTypeContext)
    EnterPrimitiveType is called when production primitiveType is entered.

func (s *BaseArcParserListener) EnterQualifiedIdentifier(ctx *QualifiedIdentifierContext)
    EnterQualifiedIdentifier is called when production qualifiedIdentifier is
    entered.

func (s *BaseArcParserListener) EnterQualifiedType(ctx *QualifiedTypeContext)
    EnterQualifiedType is called when production qualifiedType is entered.

func (s *BaseArcParserListener) EnterRangeExpression(ctx *RangeExpressionContext)
    EnterRangeExpression is called when production rangeExpression is entered.

func (s *BaseArcParserListener) EnterReferenceType(ctx *ReferenceTypeContext)
    EnterReferenceType is called when production referenceType is entered.

func (s *BaseArcParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext)
    EnterRelationalExpression is called when production relationalExpression is
    entered.

func (s *BaseArcParserListener) EnterReturnStmt(ctx *ReturnStmtContext)
    EnterReturnStmt is called when production returnStmt is entered.

func (s *BaseArcParserListener) EnterReturnType(ctx *ReturnTypeContext)
    EnterReturnType is called when production returnType is entered.

func (s *BaseArcParserListener) EnterShiftExpression(ctx *ShiftExpressionContext)
    EnterShiftExpression is called when production shiftExpression is entered.

func (s *BaseArcParserListener) EnterSizeofExpression(ctx *SizeofExpressionContext)
    EnterSizeofExpression is called when production sizeofExpression is entered.

func (s *BaseArcParserListener) EnterStatement(ctx *StatementContext)
    EnterStatement is called when production statement is entered.

func (s *BaseArcParserListener) EnterStructDecl(ctx *StructDeclContext)
    EnterStructDecl is called when production structDecl is entered.

func (s *BaseArcParserListener) EnterStructField(ctx *StructFieldContext)
    EnterStructField is called when production structField is entered.

func (s *BaseArcParserListener) EnterStructLiteral(ctx *StructLiteralContext)
    EnterStructLiteral is called when production structLiteral is entered.

func (s *BaseArcParserListener) EnterStructMember(ctx *StructMemberContext)
    EnterStructMember is called when production structMember is entered.

func (s *BaseArcParserListener) EnterSwitchCase(ctx *SwitchCaseContext)
    EnterSwitchCase is called when production switchCase is entered.

func (s *BaseArcParserListener) EnterSwitchStmt(ctx *SwitchStmtContext)
    EnterSwitchStmt is called when production switchStmt is entered.

func (s *BaseArcParserListener) EnterThrowStmt(ctx *ThrowStmtContext)
    EnterThrowStmt is called when production throwStmt is entered.

func (s *BaseArcParserListener) EnterTopLevelDecl(ctx *TopLevelDeclContext)
    EnterTopLevelDecl is called when production topLevelDecl is entered.

func (s *BaseArcParserListener) EnterTryStmt(ctx *TryStmtContext)
    EnterTryStmt is called when production tryStmt is entered.

func (s *BaseArcParserListener) EnterTupleExpression(ctx *TupleExpressionContext)
    EnterTupleExpression is called when production tupleExpression is entered.

func (s *BaseArcParserListener) EnterTuplePattern(ctx *TuplePatternContext)
    EnterTuplePattern is called when production tuplePattern is entered.

func (s *BaseArcParserListener) EnterTupleType(ctx *TupleTypeContext)
    EnterTupleType is called when production tupleType is entered.

func (s *BaseArcParserListener) EnterType(ctx *TypeContext)
    EnterType is called when production type is entered.

func (s *BaseArcParserListener) EnterTypeList(ctx *TypeListContext)
    EnterTypeList is called when production typeList is entered.

func (s *BaseArcParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext)
    EnterUnaryExpression is called when production unaryExpression is entered.

func (s *BaseArcParserListener) EnterVariableDecl(ctx *VariableDeclContext)
    EnterVariableDecl is called when production variableDecl is entered.

func (s *BaseArcParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext)
    ExitAdditiveExpression is called when production additiveExpression is
    exited.

func (s *BaseArcParserListener) ExitAlignofExpression(ctx *AlignofExpressionContext)
    ExitAlignofExpression is called when production alignofExpression is exited.

func (s *BaseArcParserListener) ExitAnonymousFuncExpression(ctx *AnonymousFuncExpressionContext)
    ExitAnonymousFuncExpression is called when production
    anonymousFuncExpression is exited.

func (s *BaseArcParserListener) ExitArgument(ctx *ArgumentContext)
    ExitArgument is called when production argument is exited.

func (s *BaseArcParserListener) ExitArgumentList(ctx *ArgumentListContext)
    ExitArgumentList is called when production argumentList is exited.

func (s *BaseArcParserListener) ExitArrayType(ctx *ArrayTypeContext)
    ExitArrayType is called when production arrayType is exited.

func (s *BaseArcParserListener) ExitAssignmentOp(ctx *AssignmentOpContext)
    ExitAssignmentOp is called when production assignmentOp is exited.

func (s *BaseArcParserListener) ExitAssignmentStmt(ctx *AssignmentStmtContext)
    ExitAssignmentStmt is called when production assignmentStmt is exited.

func (s *BaseArcParserListener) ExitBitAndExpression(ctx *BitAndExpressionContext)
    ExitBitAndExpression is called when production bitAndExpression is exited.

func (s *BaseArcParserListener) ExitBitOrExpression(ctx *BitOrExpressionContext)
    ExitBitOrExpression is called when production bitOrExpression is exited.

func (s *BaseArcParserListener) ExitBitXorExpression(ctx *BitXorExpressionContext)
    ExitBitXorExpression is called when production bitXorExpression is exited.

func (s *BaseArcParserListener) ExitBlock(ctx *BlockContext)
    ExitBlock is called when production block is exited.

func (s *BaseArcParserListener) ExitBreakStmt(ctx *BreakStmtContext)
    ExitBreakStmt is called when production breakStmt is exited.

func (s *BaseArcParserListener) ExitCCallingConvention(ctx *CCallingConventionContext)
    ExitCCallingConvention is called when production cCallingConvention is
    exited.

func (s *BaseArcParserListener) ExitClassDecl(ctx *ClassDeclContext)
    ExitClassDecl is called when production classDecl is exited.

func (s *BaseArcParserListener) ExitClassField(ctx *ClassFieldContext)
    ExitClassField is called when production classField is exited.

func (s *BaseArcParserListener) ExitClassMember(ctx *ClassMemberContext)
    ExitClassMember is called when production classMember is exited.

func (s *BaseArcParserListener) ExitCompilationUnit(ctx *CompilationUnitContext)
    ExitCompilationUnit is called when production compilationUnit is exited.

func (s *BaseArcParserListener) ExitComputeContext(ctx *ComputeContextContext)
    ExitComputeContext is called when production computeContext is exited.

func (s *BaseArcParserListener) ExitComputeExpression(ctx *ComputeExpressionContext)
    ExitComputeExpression is called when production computeExpression is exited.

func (s *BaseArcParserListener) ExitComputeMarker(ctx *ComputeMarkerContext)
    ExitComputeMarker is called when production computeMarker is exited.

func (s *BaseArcParserListener) ExitConstDecl(ctx *ConstDeclContext)
    ExitConstDecl is called when production constDecl is exited.

func (s *BaseArcParserListener) ExitContinueStmt(ctx *ContinueStmtContext)
    ExitContinueStmt is called when production continueStmt is exited.

func (s *BaseArcParserListener) ExitCppCallingConvention(ctx *CppCallingConventionContext)
    ExitCppCallingConvention is called when production cppCallingConvention is
    exited.

func (s *BaseArcParserListener) ExitDefaultCase(ctx *DefaultCaseContext)
    ExitDefaultCase is called when production defaultCase is exited.

func (s *BaseArcParserListener) ExitDeferStmt(ctx *DeferStmtContext)
    ExitDeferStmt is called when production deferStmt is exited.

func (s *BaseArcParserListener) ExitDeinitDecl(ctx *DeinitDeclContext)
    ExitDeinitDecl is called when production deinitDecl is exited.

func (s *BaseArcParserListener) ExitEnumDecl(ctx *EnumDeclContext)
    ExitEnumDecl is called when production enumDecl is exited.

func (s *BaseArcParserListener) ExitEnumMember(ctx *EnumMemberContext)
    ExitEnumMember is called when production enumMember is exited.

func (s *BaseArcParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext)
    ExitEqualityExpression is called when production equalityExpression is
    exited.

func (s *BaseArcParserListener) ExitEveryRule(ctx antlr.ParserRuleContext)
    ExitEveryRule is called when any rule is exited.

func (s *BaseArcParserListener) ExitExceptClause(ctx *ExceptClauseContext)
    ExitExceptClause is called when production exceptClause is exited.

func (s *BaseArcParserListener) ExitExpression(ctx *ExpressionContext)
    ExitExpression is called when production expression is exited.

func (s *BaseArcParserListener) ExitExpressionStmt(ctx *ExpressionStmtContext)
    ExitExpressionStmt is called when production expressionStmt is exited.

func (s *BaseArcParserListener) ExitExternCConstDecl(ctx *ExternCConstDeclContext)
    ExitExternCConstDecl is called when production externCConstDecl is exited.

func (s *BaseArcParserListener) ExitExternCDecl(ctx *ExternCDeclContext)
    ExitExternCDecl is called when production externCDecl is exited.

func (s *BaseArcParserListener) ExitExternCFunctionDecl(ctx *ExternCFunctionDeclContext)
    ExitExternCFunctionDecl is called when production externCFunctionDecl is
    exited.

func (s *BaseArcParserListener) ExitExternCMember(ctx *ExternCMemberContext)
    ExitExternCMember is called when production externCMember is exited.

func (s *BaseArcParserListener) ExitExternCOpaqueStructDecl(ctx *ExternCOpaqueStructDeclContext)
    ExitExternCOpaqueStructDecl is called when production
    externCOpaqueStructDecl is exited.

func (s *BaseArcParserListener) ExitExternCParameter(ctx *ExternCParameterContext)
    ExitExternCParameter is called when production externCParameter is exited.

func (s *BaseArcParserListener) ExitExternCParameterList(ctx *ExternCParameterListContext)
    ExitExternCParameterList is called when production externCParameterList is
    exited.

func (s *BaseArcParserListener) ExitExternCTypeAlias(ctx *ExternCTypeAliasContext)
    ExitExternCTypeAlias is called when production externCTypeAlias is exited.

func (s *BaseArcParserListener) ExitExternCppClassDecl(ctx *ExternCppClassDeclContext)
    ExitExternCppClassDecl is called when production externCppClassDecl is
    exited.

func (s *BaseArcParserListener) ExitExternCppClassMember(ctx *ExternCppClassMemberContext)
    ExitExternCppClassMember is called when production externCppClassMember is
    exited.

func (s *BaseArcParserListener) ExitExternCppConstDecl(ctx *ExternCppConstDeclContext)
    ExitExternCppConstDecl is called when production externCppConstDecl is
    exited.

func (s *BaseArcParserListener) ExitExternCppConstructorDecl(ctx *ExternCppConstructorDeclContext)
    ExitExternCppConstructorDecl is called when production
    externCppConstructorDecl is exited.

func (s *BaseArcParserListener) ExitExternCppDecl(ctx *ExternCppDeclContext)
    ExitExternCppDecl is called when production externCppDecl is exited.

func (s *BaseArcParserListener) ExitExternCppDestructorDecl(ctx *ExternCppDestructorDeclContext)
    ExitExternCppDestructorDecl is called when production
    externCppDestructorDecl is exited.

func (s *BaseArcParserListener) ExitExternCppFunctionDecl(ctx *ExternCppFunctionDeclContext)
    ExitExternCppFunctionDecl is called when production externCppFunctionDecl is
    exited.

func (s *BaseArcParserListener) ExitExternCppMember(ctx *ExternCppMemberContext)
    ExitExternCppMember is called when production externCppMember is exited.

func (s *BaseArcParserListener) ExitExternCppMethodDecl(ctx *ExternCppMethodDeclContext)
    ExitExternCppMethodDecl is called when production externCppMethodDecl is
    exited.

func (s *BaseArcParserListener) ExitExternCppMethodParams(ctx *ExternCppMethodParamsContext)
    ExitExternCppMethodParams is called when production externCppMethodParams is
    exited.

func (s *BaseArcParserListener) ExitExternCppNamespaceDecl(ctx *ExternCppNamespaceDeclContext)
    ExitExternCppNamespaceDecl is called when production externCppNamespaceDecl
    is exited.

func (s *BaseArcParserListener) ExitExternCppOpaqueClassDecl(ctx *ExternCppOpaqueClassDeclContext)
    ExitExternCppOpaqueClassDecl is called when production
    externCppOpaqueClassDecl is exited.

func (s *BaseArcParserListener) ExitExternCppParamType(ctx *ExternCppParamTypeContext)
    ExitExternCppParamType is called when production externCppParamType is
    exited.

func (s *BaseArcParserListener) ExitExternCppParameter(ctx *ExternCppParameterContext)
    ExitExternCppParameter is called when production externCppParameter is
    exited.

func (s *BaseArcParserListener) ExitExternCppParameterList(ctx *ExternCppParameterListContext)
    ExitExternCppParameterList is called when production externCppParameterList
    is exited.

func (s *BaseArcParserListener) ExitExternCppSelfParam(ctx *ExternCppSelfParamContext)
    ExitExternCppSelfParam is called when production externCppSelfParam is
    exited.

func (s *BaseArcParserListener) ExitExternCppTypeAlias(ctx *ExternCppTypeAliasContext)
    ExitExternCppTypeAlias is called when production externCppTypeAlias is
    exited.

func (s *BaseArcParserListener) ExitExternNamespacePath(ctx *ExternNamespacePathContext)
    ExitExternNamespacePath is called when production externNamespacePath is
    exited.

func (s *BaseArcParserListener) ExitFieldInit(ctx *FieldInitContext)
    ExitFieldInit is called when production fieldInit is exited.

func (s *BaseArcParserListener) ExitFinallyClause(ctx *FinallyClauseContext)
    ExitFinallyClause is called when production finallyClause is exited.

func (s *BaseArcParserListener) ExitForStmt(ctx *ForStmtContext)
    ExitForStmt is called when production forStmt is exited.

func (s *BaseArcParserListener) ExitFunctionDecl(ctx *FunctionDeclContext)
    ExitFunctionDecl is called when production functionDecl is exited.

func (s *BaseArcParserListener) ExitFunctionType(ctx *FunctionTypeContext)
    ExitFunctionType is called when production functionType is exited.

func (s *BaseArcParserListener) ExitGenericArg(ctx *GenericArgContext)
    ExitGenericArg is called when production genericArg is exited.

func (s *BaseArcParserListener) ExitGenericArgList(ctx *GenericArgListContext)
    ExitGenericArgList is called when production genericArgList is exited.

func (s *BaseArcParserListener) ExitGenericArgs(ctx *GenericArgsContext)
    ExitGenericArgs is called when production genericArgs is exited.

func (s *BaseArcParserListener) ExitGenericParamList(ctx *GenericParamListContext)
    ExitGenericParamList is called when production genericParamList is exited.

func (s *BaseArcParserListener) ExitGenericParams(ctx *GenericParamsContext)
    ExitGenericParams is called when production genericParams is exited.

func (s *BaseArcParserListener) ExitIfStmt(ctx *IfStmtContext)
    ExitIfStmt is called when production ifStmt is exited.

func (s *BaseArcParserListener) ExitImportDecl(ctx *ImportDeclContext)
    ExitImportDecl is called when production importDecl is exited.

func (s *BaseArcParserListener) ExitImportSpec(ctx *ImportSpecContext)
    ExitImportSpec is called when production importSpec is exited.

func (s *BaseArcParserListener) ExitInitDecl(ctx *InitDeclContext)
    ExitInitDecl is called when production initDecl is exited.

func (s *BaseArcParserListener) ExitInitializerEntry(ctx *InitializerEntryContext)
    ExitInitializerEntry is called when production initializerEntry is exited.

func (s *BaseArcParserListener) ExitInitializerList(ctx *InitializerListContext)
    ExitInitializerList is called when production initializerList is exited.

func (s *BaseArcParserListener) ExitLambdaExpression(ctx *LambdaExpressionContext)
    ExitLambdaExpression is called when production lambdaExpression is exited.

func (s *BaseArcParserListener) ExitLambdaParam(ctx *LambdaParamContext)
    ExitLambdaParam is called when production lambdaParam is exited.

func (s *BaseArcParserListener) ExitLambdaParamList(ctx *LambdaParamListContext)
    ExitLambdaParamList is called when production lambdaParamList is exited.

func (s *BaseArcParserListener) ExitLeftHandSide(ctx *LeftHandSideContext)
    ExitLeftHandSide is called when production leftHandSide is exited.

func (s *BaseArcParserListener) ExitLiteral(ctx *LiteralContext)
    ExitLiteral is called when production literal is exited.

func (s *BaseArcParserListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext)
    ExitLogicalAndExpression is called when production logicalAndExpression is
    exited.

func (s *BaseArcParserListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext)
    ExitLogicalOrExpression is called when production logicalOrExpression is
    exited.

func (s *BaseArcParserListener) ExitMethodDecl(ctx *MethodDeclContext)
    ExitMethodDecl is called when production methodDecl is exited.

func (s *BaseArcParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext)
    ExitMultiplicativeExpression is called when production
    multiplicativeExpression is exited.

func (s *BaseArcParserListener) ExitMutatingDecl(ctx *MutatingDeclContext)
    ExitMutatingDecl is called when production mutatingDecl is exited.

func (s *BaseArcParserListener) ExitNamespaceDecl(ctx *NamespaceDeclContext)
    ExitNamespaceDecl is called when production namespaceDecl is exited.

func (s *BaseArcParserListener) ExitParameter(ctx *ParameterContext)
    ExitParameter is called when production parameter is exited.

func (s *BaseArcParserListener) ExitParameterList(ctx *ParameterListContext)
    ExitParameterList is called when production parameterList is exited.

func (s *BaseArcParserListener) ExitPointerType(ctx *PointerTypeContext)
    ExitPointerType is called when production pointerType is exited.

func (s *BaseArcParserListener) ExitPostfixExpression(ctx *PostfixExpressionContext)
    ExitPostfixExpression is called when production postfixExpression is exited.

func (s *BaseArcParserListener) ExitPostfixOp(ctx *PostfixOpContext)
    ExitPostfixOp is called when production postfixOp is exited.

func (s *BaseArcParserListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext)
    ExitPrimaryExpression is called when production primaryExpression is exited.

func (s *BaseArcParserListener) ExitPrimitiveType(ctx *PrimitiveTypeContext)
    ExitPrimitiveType is called when production primitiveType is exited.

func (s *BaseArcParserListener) ExitQualifiedIdentifier(ctx *QualifiedIdentifierContext)
    ExitQualifiedIdentifier is called when production qualifiedIdentifier is
    exited.

func (s *BaseArcParserListener) ExitQualifiedType(ctx *QualifiedTypeContext)
    ExitQualifiedType is called when production qualifiedType is exited.

func (s *BaseArcParserListener) ExitRangeExpression(ctx *RangeExpressionContext)
    ExitRangeExpression is called when production rangeExpression is exited.

func (s *BaseArcParserListener) ExitReferenceType(ctx *ReferenceTypeContext)
    ExitReferenceType is called when production referenceType is exited.

func (s *BaseArcParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext)
    ExitRelationalExpression is called when production relationalExpression is
    exited.

func (s *BaseArcParserListener) ExitReturnStmt(ctx *ReturnStmtContext)
    ExitReturnStmt is called when production returnStmt is exited.

func (s *BaseArcParserListener) ExitReturnType(ctx *ReturnTypeContext)
    ExitReturnType is called when production returnType is exited.

func (s *BaseArcParserListener) ExitShiftExpression(ctx *ShiftExpressionContext)
    ExitShiftExpression is called when production shiftExpression is exited.

func (s *BaseArcParserListener) ExitSizeofExpression(ctx *SizeofExpressionContext)
    ExitSizeofExpression is called when production sizeofExpression is exited.

func (s *BaseArcParserListener) ExitStatement(ctx *StatementContext)
    ExitStatement is called when production statement is exited.

func (s *BaseArcParserListener) ExitStructDecl(ctx *StructDeclContext)
    ExitStructDecl is called when production structDecl is exited.

func (s *BaseArcParserListener) ExitStructField(ctx *StructFieldContext)
    ExitStructField is called when production structField is exited.

func (s *BaseArcParserListener) ExitStructLiteral(ctx *StructLiteralContext)
    ExitStructLiteral is called when production structLiteral is exited.

func (s *BaseArcParserListener) ExitStructMember(ctx *StructMemberContext)
    ExitStructMember is called when production structMember is exited.

func (s *BaseArcParserListener) ExitSwitchCase(ctx *SwitchCaseContext)
    ExitSwitchCase is called when production switchCase is exited.

func (s *BaseArcParserListener) ExitSwitchStmt(ctx *SwitchStmtContext)
    ExitSwitchStmt is called when production switchStmt is exited.

func (s *BaseArcParserListener) ExitThrowStmt(ctx *ThrowStmtContext)
    ExitThrowStmt is called when production throwStmt is exited.

func (s *BaseArcParserListener) ExitTopLevelDecl(ctx *TopLevelDeclContext)
    ExitTopLevelDecl is called when production topLevelDecl is exited.

func (s *BaseArcParserListener) ExitTryStmt(ctx *TryStmtContext)
    ExitTryStmt is called when production tryStmt is exited.

func (s *BaseArcParserListener) ExitTupleExpression(ctx *TupleExpressionContext)
    ExitTupleExpression is called when production tupleExpression is exited.

func (s *BaseArcParserListener) ExitTuplePattern(ctx *TuplePatternContext)
    ExitTuplePattern is called when production tuplePattern is exited.

func (s *BaseArcParserListener) ExitTupleType(ctx *TupleTypeContext)
    ExitTupleType is called when production tupleType is exited.

func (s *BaseArcParserListener) ExitType(ctx *TypeContext)
    ExitType is called when production type is exited.

func (s *BaseArcParserListener) ExitTypeList(ctx *TypeListContext)
    ExitTypeList is called when production typeList is exited.

func (s *BaseArcParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext)
    ExitUnaryExpression is called when production unaryExpression is exited.

func (s *BaseArcParserListener) ExitVariableDecl(ctx *VariableDeclContext)
    ExitVariableDecl is called when production variableDecl is exited.

func (s *BaseArcParserListener) VisitErrorNode(node antlr.ErrorNode)
    VisitErrorNode is called when an error node is visited.

func (s *BaseArcParserListener) VisitTerminal(node antlr.TerminalNode)
    VisitTerminal is called when a terminal node is visited.

type BaseArcParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseArcParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitAlignofExpression(ctx *AlignofExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitAnonymousFuncExpression(ctx *AnonymousFuncExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitArgument(ctx *ArgumentContext) interface{}

func (v *BaseArcParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{}

func (v *BaseArcParserVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitAssignmentOp(ctx *AssignmentOpContext) interface{}

func (v *BaseArcParserVisitor) VisitAssignmentStmt(ctx *AssignmentStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitBitAndExpression(ctx *BitAndExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitBitOrExpression(ctx *BitOrExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitBitXorExpression(ctx *BitXorExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitBlock(ctx *BlockContext) interface{}

func (v *BaseArcParserVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitCCallingConvention(ctx *CCallingConventionContext) interface{}

func (v *BaseArcParserVisitor) VisitClassDecl(ctx *ClassDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitClassField(ctx *ClassFieldContext) interface{}

func (v *BaseArcParserVisitor) VisitClassMember(ctx *ClassMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

func (v *BaseArcParserVisitor) VisitComputeContext(ctx *ComputeContextContext) interface{}

func (v *BaseArcParserVisitor) VisitComputeExpression(ctx *ComputeExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitComputeMarker(ctx *ComputeMarkerContext) interface{}

func (v *BaseArcParserVisitor) VisitConstDecl(ctx *ConstDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitCppCallingConvention(ctx *CppCallingConventionContext) interface{}

func (v *BaseArcParserVisitor) VisitDefaultCase(ctx *DefaultCaseContext) interface{}

func (v *BaseArcParserVisitor) VisitDeferStmt(ctx *DeferStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitDeinitDecl(ctx *DeinitDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitEnumDecl(ctx *EnumDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitEnumMember(ctx *EnumMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitExceptClause(ctx *ExceptClauseContext) interface{}

func (v *BaseArcParserVisitor) VisitExpression(ctx *ExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitExpressionStmt(ctx *ExpressionStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCConstDecl(ctx *ExternCConstDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCDecl(ctx *ExternCDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCFunctionDecl(ctx *ExternCFunctionDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCMember(ctx *ExternCMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCOpaqueStructDecl(ctx *ExternCOpaqueStructDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCParameter(ctx *ExternCParameterContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCParameterList(ctx *ExternCParameterListContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCTypeAlias(ctx *ExternCTypeAliasContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppClassDecl(ctx *ExternCppClassDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppClassMember(ctx *ExternCppClassMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppConstDecl(ctx *ExternCppConstDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppConstructorDecl(ctx *ExternCppConstructorDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppDecl(ctx *ExternCppDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppDestructorDecl(ctx *ExternCppDestructorDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppFunctionDecl(ctx *ExternCppFunctionDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppMember(ctx *ExternCppMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppMethodDecl(ctx *ExternCppMethodDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppMethodParams(ctx *ExternCppMethodParamsContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppNamespaceDecl(ctx *ExternCppNamespaceDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppOpaqueClassDecl(ctx *ExternCppOpaqueClassDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppParamType(ctx *ExternCppParamTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppParameter(ctx *ExternCppParameterContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppParameterList(ctx *ExternCppParameterListContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppSelfParam(ctx *ExternCppSelfParamContext) interface{}

func (v *BaseArcParserVisitor) VisitExternCppTypeAlias(ctx *ExternCppTypeAliasContext) interface{}

func (v *BaseArcParserVisitor) VisitExternNamespacePath(ctx *ExternNamespacePathContext) interface{}

func (v *BaseArcParserVisitor) VisitFieldInit(ctx *FieldInitContext) interface{}

func (v *BaseArcParserVisitor) VisitFinallyClause(ctx *FinallyClauseContext) interface{}

func (v *BaseArcParserVisitor) VisitForStmt(ctx *ForStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitGenericArg(ctx *GenericArgContext) interface{}

func (v *BaseArcParserVisitor) VisitGenericArgList(ctx *GenericArgListContext) interface{}

func (v *BaseArcParserVisitor) VisitGenericArgs(ctx *GenericArgsContext) interface{}

func (v *BaseArcParserVisitor) VisitGenericParam(ctx *GenericParamContext) interface{}

func (v *BaseArcParserVisitor) VisitGenericParamList(ctx *GenericParamListContext) interface{}

func (v *BaseArcParserVisitor) VisitGenericParams(ctx *GenericParamsContext) interface{}

func (v *BaseArcParserVisitor) VisitIfStmt(ctx *IfStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitImportSpec(ctx *ImportSpecContext) interface{}

func (v *BaseArcParserVisitor) VisitInitDecl(ctx *InitDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitInitializerEntry(ctx *InitializerEntryContext) interface{}

func (v *BaseArcParserVisitor) VisitInitializerList(ctx *InitializerListContext) interface{}

func (v *BaseArcParserVisitor) VisitLambdaExpression(ctx *LambdaExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitLambdaParam(ctx *LambdaParamContext) interface{}

func (v *BaseArcParserVisitor) VisitLambdaParamList(ctx *LambdaParamListContext) interface{}

func (v *BaseArcParserVisitor) VisitLeftHandSide(ctx *LeftHandSideContext) interface{}

func (v *BaseArcParserVisitor) VisitLiteral(ctx *LiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitMethodDecl(ctx *MethodDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitMutatingDecl(ctx *MutatingDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitNamespaceDecl(ctx *NamespaceDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitParameter(ctx *ParameterContext) interface{}

func (v *BaseArcParserVisitor) VisitParameterList(ctx *ParameterListContext) interface{}

func (v *BaseArcParserVisitor) VisitPointerType(ctx *PointerTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitPostfixOp(ctx *PostfixOpContext) interface{}

func (v *BaseArcParserVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{}

func (v *BaseArcParserVisitor) VisitQualifiedType(ctx *QualifiedTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitRangeExpression(ctx *RangeExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitReferenceType(ctx *ReferenceTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitReturnType(ctx *ReturnTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitShiftExpression(ctx *ShiftExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitSizeofExpression(ctx *SizeofExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitStatement(ctx *StatementContext) interface{}

func (v *BaseArcParserVisitor) VisitStructDecl(ctx *StructDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitStructField(ctx *StructFieldContext) interface{}

func (v *BaseArcParserVisitor) VisitStructLiteral(ctx *StructLiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitStructMember(ctx *StructMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitSwitchCase(ctx *SwitchCaseContext) interface{}

func (v *BaseArcParserVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitThrowStmt(ctx *ThrowStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitTopLevelDecl(ctx *TopLevelDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitTryStmt(ctx *TryStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitTupleExpression(ctx *TupleExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitTuplePattern(ctx *TuplePatternContext) interface{}

func (v *BaseArcParserVisitor) VisitTupleType(ctx *TupleTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitType(ctx *TypeContext) interface{}

func (v *BaseArcParserVisitor) VisitTypeList(ctx *TypeListContext) interface{}

func (v *BaseArcParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitVariableDecl(ctx *VariableDeclContext) interface{}

type BitAndExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewBitAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitAndExpressionContext

func NewEmptyBitAndExpressionContext() *BitAndExpressionContext

func (s *BitAndExpressionContext) AMP(i int) antlr.TerminalNode

func (s *BitAndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *BitAndExpressionContext) AllAMP() []antlr.TerminalNode

func (s *BitAndExpressionContext) AllEqualityExpression() []IEqualityExpressionContext

func (s *BitAndExpressionContext) EqualityExpression(i int) IEqualityExpressionContext

func (s *BitAndExpressionContext) GetParser() antlr.Parser

func (s *BitAndExpressionContext) GetRuleContext() antlr.RuleContext

func (*BitAndExpressionContext) IsBitAndExpressionContext()

func (s *BitAndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type BitOrExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewBitOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitOrExpressionContext

func NewEmptyBitOrExpressionContext() *BitOrExpressionContext

func (s *BitOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *BitOrExpressionContext) AllBIT_OR() []antlr.TerminalNode

func (s *BitOrExpressionContext) AllBitXorExpression() []IBitXorExpressionContext

func (s *BitOrExpressionContext) BIT_OR(i int) antlr.TerminalNode

func (s *BitOrExpressionContext) BitXorExpression(i int) IBitXorExpressionContext

func (s *BitOrExpressionContext) GetParser() antlr.Parser

func (s *BitOrExpressionContext) GetRuleContext() antlr.RuleContext

func (*BitOrExpressionContext) IsBitOrExpressionContext()

func (s *BitOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type BitXorExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewBitXorExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitXorExpressionContext

func NewEmptyBitXorExpressionContext() *BitXorExpressionContext

func (s *BitXorExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *BitXorExpressionContext) AllBIT_XOR() []antlr.TerminalNode

func (s *BitXorExpressionContext) AllBitAndExpression() []IBitAndExpressionContext

func (s *BitXorExpressionContext) BIT_XOR(i int) antlr.TerminalNode

func (s *BitXorExpressionContext) BitAndExpression(i int) IBitAndExpressionContext

func (s *BitXorExpressionContext) GetParser() antlr.Parser

func (s *BitXorExpressionContext) GetRuleContext() antlr.RuleContext

func (*BitXorExpressionContext) IsBitXorExpressionContext()

func (s *BitXorExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type BlockContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext

func NewEmptyBlockContext() *BlockContext

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *BlockContext) AllStatement() []IStatementContext

func (s *BlockContext) GetParser() antlr.Parser

func (s *BlockContext) GetRuleContext() antlr.RuleContext

func (*BlockContext) IsBlockContext()

func (s *BlockContext) LBRACE() antlr.TerminalNode

func (s *BlockContext) RBRACE() antlr.TerminalNode

func (s *BlockContext) Statement(i int) IStatementContext

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext

func NewEmptyBreakStmtContext() *BreakStmtContext

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *BreakStmtContext) BREAK() antlr.TerminalNode

func (s *BreakStmtContext) GetParser() antlr.Parser

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext

func (*BreakStmtContext) IsBreakStmtContext()

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type CCallingConventionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewCCallingConventionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CCallingConventionContext

func NewEmptyCCallingConventionContext() *CCallingConventionContext

func (s *CCallingConventionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *CCallingConventionContext) CDECL() antlr.TerminalNode

func (s *CCallingConventionContext) FASTCALL() antlr.TerminalNode

func (s *CCallingConventionContext) GetParser() antlr.Parser

func (s *CCallingConventionContext) GetRuleContext() antlr.RuleContext

func (*CCallingConventionContext) IsCCallingConventionContext()

func (s *CCallingConventionContext) STDCALL() antlr.TerminalNode

func (s *CCallingConventionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ClassDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewClassDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclContext

func NewEmptyClassDeclContext() *ClassDeclContext

func (s *ClassDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ClassDeclContext) AllClassMember() []IClassMemberContext

func (s *ClassDeclContext) CLASS() antlr.TerminalNode

func (s *ClassDeclContext) ClassMember(i int) IClassMemberContext

func (s *ClassDeclContext) GenericParams() IGenericParamsContext

func (s *ClassDeclContext) GetParser() antlr.Parser

func (s *ClassDeclContext) GetRuleContext() antlr.RuleContext

func (s *ClassDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ClassDeclContext) IsClassDeclContext()

func (s *ClassDeclContext) LBRACE() antlr.TerminalNode

func (s *ClassDeclContext) RBRACE() antlr.TerminalNode

func (s *ClassDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ClassFieldContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewClassFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassFieldContext

func NewEmptyClassFieldContext() *ClassFieldContext

func (s *ClassFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ClassFieldContext) COLON() antlr.TerminalNode

func (s *ClassFieldContext) GetParser() antlr.Parser

func (s *ClassFieldContext) GetRuleContext() antlr.RuleContext

func (s *ClassFieldContext) IDENTIFIER() antlr.TerminalNode

func (*ClassFieldContext) IsClassFieldContext()

func (s *ClassFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ClassFieldContext) Type_() ITypeContext

type ClassMemberContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewClassMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassMemberContext

func NewEmptyClassMemberContext() *ClassMemberContext

func (s *ClassMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ClassMemberContext) ClassField() IClassFieldContext

func (s *ClassMemberContext) DeinitDecl() IDeinitDeclContext

func (s *ClassMemberContext) FunctionDecl() IFunctionDeclContext

func (s *ClassMemberContext) GetParser() antlr.Parser

func (s *ClassMemberContext) GetRuleContext() antlr.RuleContext

func (s *ClassMemberContext) InitDecl() IInitDeclContext

func (*ClassMemberContext) IsClassMemberContext()

func (s *ClassMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type CompilationUnitContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext

func NewEmptyCompilationUnitContext() *CompilationUnitContext

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *CompilationUnitContext) AllImportDecl() []IImportDeclContext

func (s *CompilationUnitContext) AllNamespaceDecl() []INamespaceDeclContext

func (s *CompilationUnitContext) AllTopLevelDecl() []ITopLevelDeclContext

func (s *CompilationUnitContext) EOF() antlr.TerminalNode

func (s *CompilationUnitContext) GetParser() antlr.Parser

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext

func (s *CompilationUnitContext) ImportDecl(i int) IImportDeclContext

func (*CompilationUnitContext) IsCompilationUnitContext()

func (s *CompilationUnitContext) NamespaceDecl(i int) INamespaceDeclContext

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *CompilationUnitContext) TopLevelDecl(i int) ITopLevelDeclContext

type ComputeContextContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewComputeContextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComputeContextContext

func NewEmptyComputeContextContext() *ComputeContextContext

func (s *ComputeContextContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ComputeContextContext) GetParser() antlr.Parser

func (s *ComputeContextContext) GetRuleContext() antlr.RuleContext

func (s *ComputeContextContext) IDENTIFIER() antlr.TerminalNode

func (*ComputeContextContext) IsComputeContextContext()

func (s *ComputeContextContext) QualifiedIdentifier() IQualifiedIdentifierContext

func (s *ComputeContextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ComputeExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewComputeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComputeExpressionContext

func NewEmptyComputeExpressionContext() *ComputeExpressionContext

func (s *ComputeExpressionContext) ASYNC() antlr.TerminalNode

func (s *ComputeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ComputeExpressionContext) AllLPAREN() []antlr.TerminalNode

func (s *ComputeExpressionContext) AllRPAREN() []antlr.TerminalNode

func (s *ComputeExpressionContext) ArgumentList() IArgumentListContext

func (s *ComputeExpressionContext) Block() IBlockContext

func (s *ComputeExpressionContext) ComputeContext() IComputeContextContext

func (s *ComputeExpressionContext) FUNC() antlr.TerminalNode

func (s *ComputeExpressionContext) GenericParams() IGenericParamsContext

func (s *ComputeExpressionContext) GetParser() antlr.Parser

func (s *ComputeExpressionContext) GetRuleContext() antlr.RuleContext

func (*ComputeExpressionContext) IsComputeExpressionContext()

func (s *ComputeExpressionContext) LPAREN(i int) antlr.TerminalNode

func (s *ComputeExpressionContext) ParameterList() IParameterListContext

func (s *ComputeExpressionContext) RPAREN(i int) antlr.TerminalNode

func (s *ComputeExpressionContext) ReturnType() IReturnTypeContext

func (s *ComputeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ComputeMarkerContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewComputeMarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComputeMarkerContext

func NewEmptyComputeMarkerContext() *ComputeMarkerContext

func (s *ComputeMarkerContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ComputeMarkerContext) COMPUTE() antlr.TerminalNode

func (s *ComputeMarkerContext) GT() antlr.TerminalNode

func (s *ComputeMarkerContext) GetParser() antlr.Parser

func (s *ComputeMarkerContext) GetRuleContext() antlr.RuleContext

func (*ComputeMarkerContext) IsComputeMarkerContext()

func (s *ComputeMarkerContext) LT() antlr.TerminalNode

func (s *ComputeMarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ConstDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewConstDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDeclContext

func NewEmptyConstDeclContext() *ConstDeclContext

func (s *ConstDeclContext) ASSIGN() antlr.TerminalNode

func (s *ConstDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ConstDeclContext) COLON() antlr.TerminalNode

func (s *ConstDeclContext) CONST() antlr.TerminalNode

func (s *ConstDeclContext) Expression() IExpressionContext

func (s *ConstDeclContext) GetParser() antlr.Parser

func (s *ConstDeclContext) GetRuleContext() antlr.RuleContext

func (s *ConstDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ConstDeclContext) IsConstDeclContext()

func (s *ConstDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ConstDeclContext) Type_() ITypeContext

type ContinueStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext

func NewEmptyContinueStmtContext() *ContinueStmtContext

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode

func (s *ContinueStmtContext) GetParser() antlr.Parser

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext

func (*ContinueStmtContext) IsContinueStmtContext()

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type CppCallingConventionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewCppCallingConventionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CppCallingConventionContext

func NewEmptyCppCallingConventionContext() *CppCallingConventionContext

func (s *CppCallingConventionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *CppCallingConventionContext) CDECL() antlr.TerminalNode

func (s *CppCallingConventionContext) FASTCALL() antlr.TerminalNode

func (s *CppCallingConventionContext) GetParser() antlr.Parser

func (s *CppCallingConventionContext) GetRuleContext() antlr.RuleContext

func (*CppCallingConventionContext) IsCppCallingConventionContext()

func (s *CppCallingConventionContext) STDCALL() antlr.TerminalNode

func (s *CppCallingConventionContext) THISCALL() antlr.TerminalNode

func (s *CppCallingConventionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *CppCallingConventionContext) VECTORCALL() antlr.TerminalNode

type DefaultCaseContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewDefaultCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultCaseContext

func NewEmptyDefaultCaseContext() *DefaultCaseContext

func (s *DefaultCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *DefaultCaseContext) AllStatement() []IStatementContext

func (s *DefaultCaseContext) COLON() antlr.TerminalNode

func (s *DefaultCaseContext) DEFAULT() antlr.TerminalNode

func (s *DefaultCaseContext) GetParser() antlr.Parser

func (s *DefaultCaseContext) GetRuleContext() antlr.RuleContext

func (*DefaultCaseContext) IsDefaultCaseContext()

func (s *DefaultCaseContext) Statement(i int) IStatementContext

func (s *DefaultCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type DeferStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewDeferStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeferStmtContext

func NewEmptyDeferStmtContext() *DeferStmtContext

func (s *DeferStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *DeferStmtContext) AssignmentStmt() IAssignmentStmtContext

func (s *DeferStmtContext) DEFER() antlr.TerminalNode

func (s *DeferStmtContext) Expression() IExpressionContext

func (s *DeferStmtContext) GetParser() antlr.Parser

func (s *DeferStmtContext) GetRuleContext() antlr.RuleContext

func (*DeferStmtContext) IsDeferStmtContext()

func (s *DeferStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type DeinitDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewDeinitDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeinitDeclContext

func NewEmptyDeinitDeclContext() *DeinitDeclContext

func (s *DeinitDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *DeinitDeclContext) Block() IBlockContext

func (s *DeinitDeclContext) COLON() antlr.TerminalNode

func (s *DeinitDeclContext) DEINIT() antlr.TerminalNode

func (s *DeinitDeclContext) GetParser() antlr.Parser

func (s *DeinitDeclContext) GetRuleContext() antlr.RuleContext

func (s *DeinitDeclContext) IDENTIFIER() antlr.TerminalNode

func (*DeinitDeclContext) IsDeinitDeclContext()

func (s *DeinitDeclContext) LPAREN() antlr.TerminalNode

func (s *DeinitDeclContext) RPAREN() antlr.TerminalNode

func (s *DeinitDeclContext) SELF() antlr.TerminalNode

func (s *DeinitDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *DeinitDeclContext) Type_() ITypeContext

type EnumDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyEnumDeclContext() *EnumDeclContext

func NewEnumDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDeclContext

func (s *EnumDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *EnumDeclContext) AllEnumMember() []IEnumMemberContext

func (s *EnumDeclContext) COLON() antlr.TerminalNode

func (s *EnumDeclContext) ENUM() antlr.TerminalNode

func (s *EnumDeclContext) EnumMember(i int) IEnumMemberContext

func (s *EnumDeclContext) GetParser() antlr.Parser

func (s *EnumDeclContext) GetRuleContext() antlr.RuleContext

func (s *EnumDeclContext) IDENTIFIER() antlr.TerminalNode

func (*EnumDeclContext) IsEnumDeclContext()

func (s *EnumDeclContext) LBRACE() antlr.TerminalNode

func (s *EnumDeclContext) PrimitiveType() IPrimitiveTypeContext

func (s *EnumDeclContext) RBRACE() antlr.TerminalNode

func (s *EnumDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type EnumMemberContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyEnumMemberContext() *EnumMemberContext

func NewEnumMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumMemberContext

func (s *EnumMemberContext) ASSIGN() antlr.TerminalNode

func (s *EnumMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *EnumMemberContext) Expression() IExpressionContext

func (s *EnumMemberContext) GetParser() antlr.Parser

func (s *EnumMemberContext) GetRuleContext() antlr.RuleContext

func (s *EnumMemberContext) IDENTIFIER() antlr.TerminalNode

func (*EnumMemberContext) IsEnumMemberContext()

func (s *EnumMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type EqualityExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *EqualityExpressionContext) AllEQ() []antlr.TerminalNode

func (s *EqualityExpressionContext) AllNE() []antlr.TerminalNode

func (s *EqualityExpressionContext) AllRelationalExpression() []IRelationalExpressionContext

func (s *EqualityExpressionContext) EQ(i int) antlr.TerminalNode

func (s *EqualityExpressionContext) GetParser() antlr.Parser

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext

func (*EqualityExpressionContext) IsEqualityExpressionContext()

func (s *EqualityExpressionContext) NE(i int) antlr.TerminalNode

func (s *EqualityExpressionContext) RelationalExpression(i int) IRelationalExpressionContext

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExceptClauseContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExceptClauseContext() *ExceptClauseContext

func NewExceptClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExceptClauseContext

func (s *ExceptClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExceptClauseContext) Block() IBlockContext

func (s *ExceptClauseContext) EXCEPT() antlr.TerminalNode

func (s *ExceptClauseContext) GetParser() antlr.Parser

func (s *ExceptClauseContext) GetRuleContext() antlr.RuleContext

func (s *ExceptClauseContext) IDENTIFIER() antlr.TerminalNode

func (*ExceptClauseContext) IsExceptClauseContext()

func (s *ExceptClauseContext) QualifiedIdentifier() IQualifiedIdentifierContext

func (s *ExceptClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExpressionContext() *ExpressionContext

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExpressionContext) GetParser() antlr.Parser

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext

func (*ExpressionContext) IsExpressionContext()

func (s *ExpressionContext) LogicalOrExpression() ILogicalOrExpressionContext

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExpressionStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExpressionStmtContext() *ExpressionStmtContext

func NewExpressionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStmtContext

func (s *ExpressionStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExpressionStmtContext) Expression() IExpressionContext

func (s *ExpressionStmtContext) GetParser() antlr.Parser

func (s *ExpressionStmtContext) GetRuleContext() antlr.RuleContext

func (*ExpressionStmtContext) IsExpressionStmtContext()

func (s *ExpressionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCConstDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCConstDeclContext() *ExternCConstDeclContext

func NewExternCConstDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCConstDeclContext

func (s *ExternCConstDeclContext) ASSIGN() antlr.TerminalNode

func (s *ExternCConstDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCConstDeclContext) COLON() antlr.TerminalNode

func (s *ExternCConstDeclContext) CONST() antlr.TerminalNode

func (s *ExternCConstDeclContext) Expression() IExpressionContext

func (s *ExternCConstDeclContext) GetParser() antlr.Parser

func (s *ExternCConstDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternCConstDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternCConstDeclContext) IsExternCConstDeclContext()

func (s *ExternCConstDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternCConstDeclContext) Type_() ITypeContext

type ExternCDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCDeclContext() *ExternCDeclContext

func NewExternCDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCDeclContext

func (s *ExternCDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCDeclContext) AllExternCMember() []IExternCMemberContext

func (s *ExternCDeclContext) C_LANG() antlr.TerminalNode

func (s *ExternCDeclContext) EXTERN() antlr.TerminalNode

func (s *ExternCDeclContext) ExternCMember(i int) IExternCMemberContext

func (s *ExternCDeclContext) GetParser() antlr.Parser

func (s *ExternCDeclContext) GetRuleContext() antlr.RuleContext

func (*ExternCDeclContext) IsExternCDeclContext()

func (s *ExternCDeclContext) LBRACE() antlr.TerminalNode

func (s *ExternCDeclContext) RBRACE() antlr.TerminalNode

func (s *ExternCDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCFunctionDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCFunctionDeclContext() *ExternCFunctionDeclContext

func NewExternCFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCFunctionDeclContext

func (s *ExternCFunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCFunctionDeclContext) CCallingConvention() ICCallingConventionContext

func (s *ExternCFunctionDeclContext) ExternCParameterList() IExternCParameterListContext

func (s *ExternCFunctionDeclContext) FUNC() antlr.TerminalNode

func (s *ExternCFunctionDeclContext) GetParser() antlr.Parser

func (s *ExternCFunctionDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternCFunctionDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternCFunctionDeclContext) IsExternCFunctionDeclContext()

func (s *ExternCFunctionDeclContext) LPAREN() antlr.TerminalNode

func (s *ExternCFunctionDeclContext) RPAREN() antlr.TerminalNode

func (s *ExternCFunctionDeclContext) STRING_LITERAL() antlr.TerminalNode

func (s *ExternCFunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternCFunctionDeclContext) Type_() ITypeContext

type ExternCMemberContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCMemberContext() *ExternCMemberContext

func NewExternCMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCMemberContext

func (s *ExternCMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCMemberContext) ExternCConstDecl() IExternCConstDeclContext

func (s *ExternCMemberContext) ExternCFunctionDecl() IExternCFunctionDeclContext

func (s *ExternCMemberContext) ExternCOpaqueStructDecl() IExternCOpaqueStructDeclContext

func (s *ExternCMemberContext) ExternCTypeAlias() IExternCTypeAliasContext

func (s *ExternCMemberContext) GetParser() antlr.Parser

func (s *ExternCMemberContext) GetRuleContext() antlr.RuleContext

func (*ExternCMemberContext) IsExternCMemberContext()

func (s *ExternCMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCOpaqueStructDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCOpaqueStructDeclContext() *ExternCOpaqueStructDeclContext

func NewExternCOpaqueStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCOpaqueStructDeclContext

func (s *ExternCOpaqueStructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCOpaqueStructDeclContext) GetParser() antlr.Parser

func (s *ExternCOpaqueStructDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternCOpaqueStructDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternCOpaqueStructDeclContext) IsExternCOpaqueStructDeclContext()

func (s *ExternCOpaqueStructDeclContext) LBRACE() antlr.TerminalNode

func (s *ExternCOpaqueStructDeclContext) OPAQUE() antlr.TerminalNode

func (s *ExternCOpaqueStructDeclContext) RBRACE() antlr.TerminalNode

func (s *ExternCOpaqueStructDeclContext) STRUCT() antlr.TerminalNode

func (s *ExternCOpaqueStructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCParameterContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCParameterContext() *ExternCParameterContext

func NewExternCParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCParameterContext

func (s *ExternCParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCParameterContext) COLON() antlr.TerminalNode

func (s *ExternCParameterContext) GetParser() antlr.Parser

func (s *ExternCParameterContext) GetRuleContext() antlr.RuleContext

func (s *ExternCParameterContext) IDENTIFIER() antlr.TerminalNode

func (*ExternCParameterContext) IsExternCParameterContext()

func (s *ExternCParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternCParameterContext) Type_() ITypeContext

type ExternCParameterListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCParameterListContext() *ExternCParameterListContext

func NewExternCParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCParameterListContext

func (s *ExternCParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCParameterListContext) AllCOMMA() []antlr.TerminalNode

func (s *ExternCParameterListContext) AllExternCParameter() []IExternCParameterContext

func (s *ExternCParameterListContext) COMMA(i int) antlr.TerminalNode

func (s *ExternCParameterListContext) ELLIPSIS() antlr.TerminalNode

func (s *ExternCParameterListContext) ExternCParameter(i int) IExternCParameterContext

func (s *ExternCParameterListContext) GetParser() antlr.Parser

func (s *ExternCParameterListContext) GetRuleContext() antlr.RuleContext

func (*ExternCParameterListContext) IsExternCParameterListContext()

func (s *ExternCParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCTypeAliasContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCTypeAliasContext() *ExternCTypeAliasContext

func NewExternCTypeAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCTypeAliasContext

func (s *ExternCTypeAliasContext) ASSIGN() antlr.TerminalNode

func (s *ExternCTypeAliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCTypeAliasContext) GetParser() antlr.Parser

func (s *ExternCTypeAliasContext) GetRuleContext() antlr.RuleContext

func (s *ExternCTypeAliasContext) IDENTIFIER() antlr.TerminalNode

func (*ExternCTypeAliasContext) IsExternCTypeAliasContext()

func (s *ExternCTypeAliasContext) TYPE() antlr.TerminalNode

func (s *ExternCTypeAliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternCTypeAliasContext) Type_() ITypeContext

type ExternCppClassDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppClassDeclContext() *ExternCppClassDeclContext

func NewExternCppClassDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppClassDeclContext

func (s *ExternCppClassDeclContext) ABSTRACT() antlr.TerminalNode

func (s *ExternCppClassDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppClassDeclContext) AllExternCppClassMember() []IExternCppClassMemberContext

func (s *ExternCppClassDeclContext) CLASS() antlr.TerminalNode

func (s *ExternCppClassDeclContext) ExternCppClassMember(i int) IExternCppClassMemberContext

func (s *ExternCppClassDeclContext) GetParser() antlr.Parser

func (s *ExternCppClassDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternCppClassDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternCppClassDeclContext) IsExternCppClassDeclContext()

func (s *ExternCppClassDeclContext) LBRACE() antlr.TerminalNode

func (s *ExternCppClassDeclContext) RBRACE() antlr.TerminalNode

func (s *ExternCppClassDeclContext) STRING_LITERAL() antlr.TerminalNode

func (s *ExternCppClassDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCppClassMemberContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppClassMemberContext() *ExternCppClassMemberContext

func NewExternCppClassMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppClassMemberContext

func (s *ExternCppClassMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppClassMemberContext) ExternCppConstructorDecl() IExternCppConstructorDeclContext

func (s *ExternCppClassMemberContext) ExternCppDestructorDecl() IExternCppDestructorDeclContext

func (s *ExternCppClassMemberContext) ExternCppMethodDecl() IExternCppMethodDeclContext

func (s *ExternCppClassMemberContext) GetParser() antlr.Parser

func (s *ExternCppClassMemberContext) GetRuleContext() antlr.RuleContext

func (*ExternCppClassMemberContext) IsExternCppClassMemberContext()

func (s *ExternCppClassMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCppConstDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppConstDeclContext() *ExternCppConstDeclContext

func NewExternCppConstDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppConstDeclContext

func (s *ExternCppConstDeclContext) ASSIGN() antlr.TerminalNode

func (s *ExternCppConstDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppConstDeclContext) COLON() antlr.TerminalNode

func (s *ExternCppConstDeclContext) CONST() antlr.TerminalNode

func (s *ExternCppConstDeclContext) Expression() IExpressionContext

func (s *ExternCppConstDeclContext) GetParser() antlr.Parser

func (s *ExternCppConstDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternCppConstDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternCppConstDeclContext) IsExternCppConstDeclContext()

func (s *ExternCppConstDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternCppConstDeclContext) Type_() ITypeContext

type ExternCppConstructorDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppConstructorDeclContext() *ExternCppConstructorDeclContext

func NewExternCppConstructorDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppConstructorDeclContext

func (s *ExternCppConstructorDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppConstructorDeclContext) ExternCppParameterList() IExternCppParameterListContext

func (s *ExternCppConstructorDeclContext) GetParser() antlr.Parser

func (s *ExternCppConstructorDeclContext) GetRuleContext() antlr.RuleContext

func (*ExternCppConstructorDeclContext) IsExternCppConstructorDeclContext()

func (s *ExternCppConstructorDeclContext) LPAREN() antlr.TerminalNode

func (s *ExternCppConstructorDeclContext) NEW() antlr.TerminalNode

func (s *ExternCppConstructorDeclContext) RPAREN() antlr.TerminalNode

func (s *ExternCppConstructorDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternCppConstructorDeclContext) Type_() ITypeContext

type ExternCppDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppDeclContext() *ExternCppDeclContext

func NewExternCppDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppDeclContext

func (s *ExternCppDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppDeclContext) AllExternCppMember() []IExternCppMemberContext

func (s *ExternCppDeclContext) CPP_LANG() antlr.TerminalNode

func (s *ExternCppDeclContext) EXTERN() antlr.TerminalNode

func (s *ExternCppDeclContext) ExternCppMember(i int) IExternCppMemberContext

func (s *ExternCppDeclContext) GetParser() antlr.Parser

func (s *ExternCppDeclContext) GetRuleContext() antlr.RuleContext

func (*ExternCppDeclContext) IsExternCppDeclContext()

func (s *ExternCppDeclContext) LBRACE() antlr.TerminalNode

func (s *ExternCppDeclContext) RBRACE() antlr.TerminalNode

func (s *ExternCppDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCppDestructorDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppDestructorDeclContext() *ExternCppDestructorDeclContext

func NewExternCppDestructorDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppDestructorDeclContext

func (s *ExternCppDestructorDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppDestructorDeclContext) DELETE() antlr.TerminalNode

func (s *ExternCppDestructorDeclContext) ExternCppSelfParam() IExternCppSelfParamContext

func (s *ExternCppDestructorDeclContext) GetParser() antlr.Parser

func (s *ExternCppDestructorDeclContext) GetRuleContext() antlr.RuleContext

func (*ExternCppDestructorDeclContext) IsExternCppDestructorDeclContext()

func (s *ExternCppDestructorDeclContext) LPAREN() antlr.TerminalNode

func (s *ExternCppDestructorDeclContext) RPAREN() antlr.TerminalNode

func (s *ExternCppDestructorDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternCppDestructorDeclContext) Type_() ITypeContext

type ExternCppFunctionDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppFunctionDeclContext() *ExternCppFunctionDeclContext

func NewExternCppFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppFunctionDeclContext

func (s *ExternCppFunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppFunctionDeclContext) CppCallingConvention() ICppCallingConventionContext

func (s *ExternCppFunctionDeclContext) ExternCppParameterList() IExternCppParameterListContext

func (s *ExternCppFunctionDeclContext) FUNC() antlr.TerminalNode

func (s *ExternCppFunctionDeclContext) GetParser() antlr.Parser

func (s *ExternCppFunctionDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternCppFunctionDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternCppFunctionDeclContext) IsExternCppFunctionDeclContext()

func (s *ExternCppFunctionDeclContext) LPAREN() antlr.TerminalNode

func (s *ExternCppFunctionDeclContext) RPAREN() antlr.TerminalNode

func (s *ExternCppFunctionDeclContext) STRING_LITERAL() antlr.TerminalNode

func (s *ExternCppFunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternCppFunctionDeclContext) Type_() ITypeContext

type ExternCppMemberContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppMemberContext() *ExternCppMemberContext

func NewExternCppMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppMemberContext

func (s *ExternCppMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppMemberContext) ExternCppClassDecl() IExternCppClassDeclContext

func (s *ExternCppMemberContext) ExternCppConstDecl() IExternCppConstDeclContext

func (s *ExternCppMemberContext) ExternCppFunctionDecl() IExternCppFunctionDeclContext

func (s *ExternCppMemberContext) ExternCppNamespaceDecl() IExternCppNamespaceDeclContext

func (s *ExternCppMemberContext) ExternCppOpaqueClassDecl() IExternCppOpaqueClassDeclContext

func (s *ExternCppMemberContext) ExternCppTypeAlias() IExternCppTypeAliasContext

func (s *ExternCppMemberContext) GetParser() antlr.Parser

func (s *ExternCppMemberContext) GetRuleContext() antlr.RuleContext

func (*ExternCppMemberContext) IsExternCppMemberContext()

func (s *ExternCppMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCppMethodDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppMethodDeclContext() *ExternCppMethodDeclContext

func NewExternCppMethodDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppMethodDeclContext

func (s *ExternCppMethodDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppMethodDeclContext) CONST() antlr.TerminalNode

func (s *ExternCppMethodDeclContext) CppCallingConvention() ICppCallingConventionContext

func (s *ExternCppMethodDeclContext) ExternCppMethodParams() IExternCppMethodParamsContext

func (s *ExternCppMethodDeclContext) FUNC() antlr.TerminalNode

func (s *ExternCppMethodDeclContext) GetParser() antlr.Parser

func (s *ExternCppMethodDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternCppMethodDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternCppMethodDeclContext) IsExternCppMethodDeclContext()

func (s *ExternCppMethodDeclContext) LPAREN() antlr.TerminalNode

func (s *ExternCppMethodDeclContext) RPAREN() antlr.TerminalNode

func (s *ExternCppMethodDeclContext) STATIC() antlr.TerminalNode

func (s *ExternCppMethodDeclContext) STRING_LITERAL() antlr.TerminalNode

func (s *ExternCppMethodDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternCppMethodDeclContext) Type_() ITypeContext

func (s *ExternCppMethodDeclContext) VIRTUAL() antlr.TerminalNode

type ExternCppMethodParamsContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppMethodParamsContext() *ExternCppMethodParamsContext

func NewExternCppMethodParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppMethodParamsContext

func (s *ExternCppMethodParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppMethodParamsContext) AllCOMMA() []antlr.TerminalNode

func (s *ExternCppMethodParamsContext) AllExternCppParameter() []IExternCppParameterContext

func (s *ExternCppMethodParamsContext) COMMA(i int) antlr.TerminalNode

func (s *ExternCppMethodParamsContext) ExternCppParameter(i int) IExternCppParameterContext

func (s *ExternCppMethodParamsContext) ExternCppParameterList() IExternCppParameterListContext

func (s *ExternCppMethodParamsContext) ExternCppSelfParam() IExternCppSelfParamContext

func (s *ExternCppMethodParamsContext) GetParser() antlr.Parser

func (s *ExternCppMethodParamsContext) GetRuleContext() antlr.RuleContext

func (*ExternCppMethodParamsContext) IsExternCppMethodParamsContext()

func (s *ExternCppMethodParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCppNamespaceDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppNamespaceDeclContext() *ExternCppNamespaceDeclContext

func NewExternCppNamespaceDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppNamespaceDeclContext

func (s *ExternCppNamespaceDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppNamespaceDeclContext) AllExternCppMember() []IExternCppMemberContext

func (s *ExternCppNamespaceDeclContext) ExternCppMember(i int) IExternCppMemberContext

func (s *ExternCppNamespaceDeclContext) ExternNamespacePath() IExternNamespacePathContext

func (s *ExternCppNamespaceDeclContext) GetParser() antlr.Parser

func (s *ExternCppNamespaceDeclContext) GetRuleContext() antlr.RuleContext

func (*ExternCppNamespaceDeclContext) IsExternCppNamespaceDeclContext()

func (s *ExternCppNamespaceDeclContext) LBRACE() antlr.TerminalNode

func (s *ExternCppNamespaceDeclContext) NAMESPACE() antlr.TerminalNode

func (s *ExternCppNamespaceDeclContext) RBRACE() antlr.TerminalNode

func (s *ExternCppNamespaceDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCppOpaqueClassDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppOpaqueClassDeclContext() *ExternCppOpaqueClassDeclContext

func NewExternCppOpaqueClassDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppOpaqueClassDeclContext

func (s *ExternCppOpaqueClassDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppOpaqueClassDeclContext) CLASS() antlr.TerminalNode

func (s *ExternCppOpaqueClassDeclContext) GetParser() antlr.Parser

func (s *ExternCppOpaqueClassDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternCppOpaqueClassDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternCppOpaqueClassDeclContext) IsExternCppOpaqueClassDeclContext()

func (s *ExternCppOpaqueClassDeclContext) LBRACE() antlr.TerminalNode

func (s *ExternCppOpaqueClassDeclContext) OPAQUE() antlr.TerminalNode

func (s *ExternCppOpaqueClassDeclContext) RBRACE() antlr.TerminalNode

func (s *ExternCppOpaqueClassDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCppParamTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppParamTypeContext() *ExternCppParamTypeContext

func NewExternCppParamTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppParamTypeContext

func (s *ExternCppParamTypeContext) AMP() antlr.TerminalNode

func (s *ExternCppParamTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppParamTypeContext) CONST() antlr.TerminalNode

func (s *ExternCppParamTypeContext) GetParser() antlr.Parser

func (s *ExternCppParamTypeContext) GetRuleContext() antlr.RuleContext

func (*ExternCppParamTypeContext) IsExternCppParamTypeContext()

func (s *ExternCppParamTypeContext) STAR() antlr.TerminalNode

func (s *ExternCppParamTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternCppParamTypeContext) Type_() ITypeContext

type ExternCppParameterContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppParameterContext() *ExternCppParameterContext

func NewExternCppParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppParameterContext

func (s *ExternCppParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppParameterContext) COLON() antlr.TerminalNode

func (s *ExternCppParameterContext) ExternCppParamType() IExternCppParamTypeContext

func (s *ExternCppParameterContext) GetParser() antlr.Parser

func (s *ExternCppParameterContext) GetRuleContext() antlr.RuleContext

func (s *ExternCppParameterContext) IDENTIFIER() antlr.TerminalNode

func (*ExternCppParameterContext) IsExternCppParameterContext()

func (s *ExternCppParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCppParameterListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppParameterListContext() *ExternCppParameterListContext

func NewExternCppParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppParameterListContext

func (s *ExternCppParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppParameterListContext) AllCOMMA() []antlr.TerminalNode

func (s *ExternCppParameterListContext) AllExternCppParameter() []IExternCppParameterContext

func (s *ExternCppParameterListContext) COMMA(i int) antlr.TerminalNode

func (s *ExternCppParameterListContext) ELLIPSIS() antlr.TerminalNode

func (s *ExternCppParameterListContext) ExternCppParameter(i int) IExternCppParameterContext

func (s *ExternCppParameterListContext) GetParser() antlr.Parser

func (s *ExternCppParameterListContext) GetRuleContext() antlr.RuleContext

func (*ExternCppParameterListContext) IsExternCppParameterListContext()

func (s *ExternCppParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCppSelfParamContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppSelfParamContext() *ExternCppSelfParamContext

func NewExternCppSelfParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppSelfParamContext

func (s *ExternCppSelfParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppSelfParamContext) CONST() antlr.TerminalNode

func (s *ExternCppSelfParamContext) GetParser() antlr.Parser

func (s *ExternCppSelfParamContext) GetRuleContext() antlr.RuleContext

func (s *ExternCppSelfParamContext) IDENTIFIER() antlr.TerminalNode

func (*ExternCppSelfParamContext) IsExternCppSelfParamContext()

func (s *ExternCppSelfParamContext) SELF() antlr.TerminalNode

func (s *ExternCppSelfParamContext) STAR() antlr.TerminalNode

func (s *ExternCppSelfParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternCppTypeAliasContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternCppTypeAliasContext() *ExternCppTypeAliasContext

func NewExternCppTypeAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternCppTypeAliasContext

func (s *ExternCppTypeAliasContext) ASSIGN() antlr.TerminalNode

func (s *ExternCppTypeAliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternCppTypeAliasContext) GetParser() antlr.Parser

func (s *ExternCppTypeAliasContext) GetRuleContext() antlr.RuleContext

func (s *ExternCppTypeAliasContext) IDENTIFIER() antlr.TerminalNode

func (*ExternCppTypeAliasContext) IsExternCppTypeAliasContext()

func (s *ExternCppTypeAliasContext) TYPE() antlr.TerminalNode

func (s *ExternCppTypeAliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternCppTypeAliasContext) Type_() ITypeContext

type ExternNamespacePathContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternNamespacePathContext() *ExternNamespacePathContext

func NewExternNamespacePathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternNamespacePathContext

func (s *ExternNamespacePathContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternNamespacePathContext) AllDOT() []antlr.TerminalNode

func (s *ExternNamespacePathContext) AllIDENTIFIER() []antlr.TerminalNode

func (s *ExternNamespacePathContext) DOT(i int) antlr.TerminalNode

func (s *ExternNamespacePathContext) GetParser() antlr.Parser

func (s *ExternNamespacePathContext) GetRuleContext() antlr.RuleContext

func (s *ExternNamespacePathContext) IDENTIFIER(i int) antlr.TerminalNode

func (*ExternNamespacePathContext) IsExternNamespacePathContext()

func (s *ExternNamespacePathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type FieldInitContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyFieldInitContext() *FieldInitContext

func NewFieldInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldInitContext

func (s *FieldInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *FieldInitContext) COLON() antlr.TerminalNode

func (s *FieldInitContext) Expression() IExpressionContext

func (s *FieldInitContext) GetParser() antlr.Parser

func (s *FieldInitContext) GetRuleContext() antlr.RuleContext

func (s *FieldInitContext) IDENTIFIER() antlr.TerminalNode

func (*FieldInitContext) IsFieldInitContext()

func (s *FieldInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type FinallyClauseContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyFinallyClauseContext() *FinallyClauseContext

func NewFinallyClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FinallyClauseContext

func (s *FinallyClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *FinallyClauseContext) Block() IBlockContext

func (s *FinallyClauseContext) FINALLY() antlr.TerminalNode

func (s *FinallyClauseContext) GetParser() antlr.Parser

func (s *FinallyClauseContext) GetRuleContext() antlr.RuleContext

func (*FinallyClauseContext) IsFinallyClauseContext()

func (s *FinallyClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ForStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyForStmtContext() *ForStmtContext

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ForStmtContext) AllAssignmentStmt() []IAssignmentStmtContext

func (s *ForStmtContext) AllExpression() []IExpressionContext

func (s *ForStmtContext) AllIDENTIFIER() []antlr.TerminalNode

func (s *ForStmtContext) AllSEMICOLON() []antlr.TerminalNode

func (s *ForStmtContext) AssignmentStmt(i int) IAssignmentStmtContext

func (s *ForStmtContext) Block() IBlockContext

func (s *ForStmtContext) COMMA() antlr.TerminalNode

func (s *ForStmtContext) Expression(i int) IExpressionContext

func (s *ForStmtContext) FOR() antlr.TerminalNode

func (s *ForStmtContext) GetParser() antlr.Parser

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext

func (s *ForStmtContext) IDENTIFIER(i int) antlr.TerminalNode

func (s *ForStmtContext) IN() antlr.TerminalNode

func (*ForStmtContext) IsForStmtContext()

func (s *ForStmtContext) SEMICOLON(i int) antlr.TerminalNode

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ForStmtContext) VariableDecl() IVariableDeclContext

type FunctionDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyFunctionDeclContext() *FunctionDeclContext

func NewFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclContext

func (s *FunctionDeclContext) ASYNC() antlr.TerminalNode

func (s *FunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *FunctionDeclContext) Block() IBlockContext

func (s *FunctionDeclContext) CONTAINER() antlr.TerminalNode

func (s *FunctionDeclContext) FUNC() antlr.TerminalNode

func (s *FunctionDeclContext) GenericParams() IGenericParamsContext

func (s *FunctionDeclContext) GetParser() antlr.Parser

func (s *FunctionDeclContext) GetRuleContext() antlr.RuleContext

func (s *FunctionDeclContext) IDENTIFIER() antlr.TerminalNode

func (*FunctionDeclContext) IsFunctionDeclContext()

func (s *FunctionDeclContext) LPAREN() antlr.TerminalNode

func (s *FunctionDeclContext) PROCESS() antlr.TerminalNode

func (s *FunctionDeclContext) ParameterList() IParameterListContext

func (s *FunctionDeclContext) RPAREN() antlr.TerminalNode

func (s *FunctionDeclContext) ReturnType() IReturnTypeContext

func (s *FunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type FunctionTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyFunctionTypeContext() *FunctionTypeContext

func NewFunctionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionTypeContext

func (s *FunctionTypeContext) ASYNC() antlr.TerminalNode

func (s *FunctionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *FunctionTypeContext) FUNC() antlr.TerminalNode

func (s *FunctionTypeContext) GenericParams() IGenericParamsContext

func (s *FunctionTypeContext) GetParser() antlr.Parser

func (s *FunctionTypeContext) GetRuleContext() antlr.RuleContext

func (*FunctionTypeContext) IsFunctionTypeContext()

func (s *FunctionTypeContext) LPAREN() antlr.TerminalNode

func (s *FunctionTypeContext) RPAREN() antlr.TerminalNode

func (s *FunctionTypeContext) ReturnType() IReturnTypeContext

func (s *FunctionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *FunctionTypeContext) TypeList() ITypeListContext

type GenericArgContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyGenericArgContext() *GenericArgContext

func NewGenericArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericArgContext

func (s *GenericArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *GenericArgContext) Expression() IExpressionContext

func (s *GenericArgContext) GetParser() antlr.Parser

func (s *GenericArgContext) GetRuleContext() antlr.RuleContext

func (*GenericArgContext) IsGenericArgContext()

func (s *GenericArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *GenericArgContext) Type_() ITypeContext

type GenericArgListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyGenericArgListContext() *GenericArgListContext

func NewGenericArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericArgListContext

func (s *GenericArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *GenericArgListContext) AllCOMMA() []antlr.TerminalNode

func (s *GenericArgListContext) AllGenericArg() []IGenericArgContext

func (s *GenericArgListContext) COMMA(i int) antlr.TerminalNode

func (s *GenericArgListContext) GenericArg(i int) IGenericArgContext

func (s *GenericArgListContext) GetParser() antlr.Parser

func (s *GenericArgListContext) GetRuleContext() antlr.RuleContext

func (*GenericArgListContext) IsGenericArgListContext()

func (s *GenericArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type GenericArgsContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyGenericArgsContext() *GenericArgsContext

func NewGenericArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericArgsContext

func (s *GenericArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *GenericArgsContext) GT() antlr.TerminalNode

func (s *GenericArgsContext) GenericArgList() IGenericArgListContext

func (s *GenericArgsContext) GetParser() antlr.Parser

func (s *GenericArgsContext) GetRuleContext() antlr.RuleContext

func (*GenericArgsContext) IsGenericArgsContext()

func (s *GenericArgsContext) LT() antlr.TerminalNode

func (s *GenericArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type GenericParamContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyGenericParamContext() *GenericParamContext

func NewGenericParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericParamContext

func (s *GenericParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *GenericParamContext) AllDOT() []antlr.TerminalNode

func (s *GenericParamContext) AllIDENTIFIER() []antlr.TerminalNode

func (s *GenericParamContext) DOT(i int) antlr.TerminalNode

func (s *GenericParamContext) GetParser() antlr.Parser

func (s *GenericParamContext) GetRuleContext() antlr.RuleContext

func (s *GenericParamContext) IDENTIFIER(i int) antlr.TerminalNode

func (*GenericParamContext) IsGenericParamContext()

func (s *GenericParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type GenericParamListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyGenericParamListContext() *GenericParamListContext

func NewGenericParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericParamListContext

func (s *GenericParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *GenericParamListContext) AllCOMMA() []antlr.TerminalNode

func (s *GenericParamListContext) AllGenericParam() []IGenericParamContext

func (s *GenericParamListContext) COMMA(i int) antlr.TerminalNode

func (s *GenericParamListContext) GenericParam(i int) IGenericParamContext

func (s *GenericParamListContext) GetParser() antlr.Parser

func (s *GenericParamListContext) GetRuleContext() antlr.RuleContext

func (*GenericParamListContext) IsGenericParamListContext()

func (s *GenericParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type GenericParamsContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyGenericParamsContext() *GenericParamsContext

func NewGenericParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericParamsContext

func (s *GenericParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *GenericParamsContext) GT() antlr.TerminalNode

func (s *GenericParamsContext) GenericParamList() IGenericParamListContext

func (s *GenericParamsContext) GetParser() antlr.Parser

func (s *GenericParamsContext) GetRuleContext() antlr.RuleContext

func (*GenericParamsContext) IsGenericParamsContext()

func (s *GenericParamsContext) LT() antlr.TerminalNode

func (s *GenericParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMultiplicativeExpression() []IMultiplicativeExpressionContext
	MultiplicativeExpression(i int) IMultiplicativeExpressionContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}
    IAdditiveExpressionContext is an interface to support dynamic dispatch.

type IAlignofExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALIGNOF() antlr.TerminalNode
	LT() antlr.TerminalNode
	Type_() ITypeContext
	GT() antlr.TerminalNode

	// IsAlignofExpressionContext differentiates from other interfaces.
	IsAlignofExpressionContext()
}
    IAlignofExpressionContext is an interface to support dynamic dispatch.

type IAnonymousFuncExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	GenericParams() IGenericParamsContext
	ParameterList() IParameterListContext
	ReturnType() IReturnTypeContext
	ASYNC() antlr.TerminalNode
	PROCESS() antlr.TerminalNode
	CONTAINER() antlr.TerminalNode

	// IsAnonymousFuncExpressionContext differentiates from other interfaces.
	IsAnonymousFuncExpressionContext()
}
    IAnonymousFuncExpressionContext is an interface to support dynamic dispatch.

type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	LambdaExpression() ILambdaExpressionContext
	AnonymousFuncExpression() IAnonymousFuncExpressionContext

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}
    IArgumentContext is an interface to support dynamic dispatch.

type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArgument() []IArgumentContext
	Argument(i int) IArgumentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}
    IArgumentListContext is an interface to support dynamic dispatch.

type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	Expression() IExpressionContext
	RBRACKET() antlr.TerminalNode
	Type_() ITypeContext

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}
    IArrayTypeContext is an interface to support dynamic dispatch.

type IAssignmentOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	PLUS_ASSIGN() antlr.TerminalNode
	MINUS_ASSIGN() antlr.TerminalNode
	STAR_ASSIGN() antlr.TerminalNode
	SLASH_ASSIGN() antlr.TerminalNode
	PERCENT_ASSIGN() antlr.TerminalNode
	BIT_OR_ASSIGN() antlr.TerminalNode
	BIT_AND_ASSIGN() antlr.TerminalNode
	BIT_XOR_ASSIGN() antlr.TerminalNode
	LT() antlr.TerminalNode
	LE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GE() antlr.TerminalNode

	// IsAssignmentOpContext differentiates from other interfaces.
	IsAssignmentOpContext()
}
    IAssignmentOpContext is an interface to support dynamic dispatch.

type IAssignmentStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LeftHandSide() ILeftHandSideContext
	AssignmentOp() IAssignmentOpContext
	Expression() IExpressionContext

	// IsAssignmentStmtContext differentiates from other interfaces.
	IsAssignmentStmtContext()
}
    IAssignmentStmtContext is an interface to support dynamic dispatch.

type IBitAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEqualityExpression() []IEqualityExpressionContext
	EqualityExpression(i int) IEqualityExpressionContext
	AllAMP() []antlr.TerminalNode
	AMP(i int) antlr.TerminalNode

	// IsBitAndExpressionContext differentiates from other interfaces.
	IsBitAndExpressionContext()
}
    IBitAndExpressionContext is an interface to support dynamic dispatch.

type IBitOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBitXorExpression() []IBitXorExpressionContext
	BitXorExpression(i int) IBitXorExpressionContext
	AllBIT_OR() []antlr.TerminalNode
	BIT_OR(i int) antlr.TerminalNode

	// IsBitOrExpressionContext differentiates from other interfaces.
	IsBitOrExpressionContext()
}
    IBitOrExpressionContext is an interface to support dynamic dispatch.

type IBitXorExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBitAndExpression() []IBitAndExpressionContext
	BitAndExpression(i int) IBitAndExpressionContext
	AllBIT_XOR() []antlr.TerminalNode
	BIT_XOR(i int) antlr.TerminalNode

	// IsBitXorExpressionContext differentiates from other interfaces.
	IsBitXorExpressionContext()
}
    IBitXorExpressionContext is an interface to support dynamic dispatch.

type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}
    IBlockContext is an interface to support dynamic dispatch.

type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}
    IBreakStmtContext is an interface to support dynamic dispatch.

type ICCallingConventionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STDCALL() antlr.TerminalNode
	CDECL() antlr.TerminalNode
	FASTCALL() antlr.TerminalNode

	// IsCCallingConventionContext differentiates from other interfaces.
	IsCCallingConventionContext()
}
    ICCallingConventionContext is an interface to support dynamic dispatch.

type IClassDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLASS() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	GenericParams() IGenericParamsContext
	AllClassMember() []IClassMemberContext
	ClassMember(i int) IClassMemberContext

	// IsClassDeclContext differentiates from other interfaces.
	IsClassDeclContext()
}
    IClassDeclContext is an interface to support dynamic dispatch.

type IClassFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsClassFieldContext differentiates from other interfaces.
	IsClassFieldContext()
}
    IClassFieldContext is an interface to support dynamic dispatch.

type IClassMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ClassField() IClassFieldContext
	FunctionDecl() IFunctionDeclContext
	InitDecl() IInitDeclContext
	DeinitDecl() IDeinitDeclContext

	// IsClassMemberContext differentiates from other interfaces.
	IsClassMemberContext()
}
    IClassMemberContext is an interface to support dynamic dispatch.

type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllImportDecl() []IImportDeclContext
	ImportDecl(i int) IImportDeclContext
	AllNamespaceDecl() []INamespaceDeclContext
	NamespaceDecl(i int) INamespaceDeclContext
	AllTopLevelDecl() []ITopLevelDeclContext
	TopLevelDecl(i int) ITopLevelDeclContext

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}
    ICompilationUnitContext is an interface to support dynamic dispatch.

type IComputeContextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QualifiedIdentifier() IQualifiedIdentifierContext
	IDENTIFIER() antlr.TerminalNode

	// IsComputeContextContext differentiates from other interfaces.
	IsComputeContextContext()
}
    IComputeContextContext is an interface to support dynamic dispatch.

type IComputeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ComputeContext() IComputeContextContext
	FUNC() antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	Block() IBlockContext
	ASYNC() antlr.TerminalNode
	GenericParams() IGenericParamsContext
	ParameterList() IParameterListContext
	ReturnType() IReturnTypeContext
	ArgumentList() IArgumentListContext

	// IsComputeExpressionContext differentiates from other interfaces.
	IsComputeExpressionContext()
}
    IComputeExpressionContext is an interface to support dynamic dispatch.

type IComputeMarkerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LT() antlr.TerminalNode
	COMPUTE() antlr.TerminalNode
	GT() antlr.TerminalNode

	// IsComputeMarkerContext differentiates from other interfaces.
	IsComputeMarkerContext()
}
    IComputeMarkerContext is an interface to support dynamic dispatch.

type IConstDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsConstDeclContext differentiates from other interfaces.
	IsConstDeclContext()
}
    IConstDeclContext is an interface to support dynamic dispatch.

type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}
    IContinueStmtContext is an interface to support dynamic dispatch.

type ICppCallingConventionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STDCALL() antlr.TerminalNode
	CDECL() antlr.TerminalNode
	FASTCALL() antlr.TerminalNode
	VECTORCALL() antlr.TerminalNode
	THISCALL() antlr.TerminalNode

	// IsCppCallingConventionContext differentiates from other interfaces.
	IsCppCallingConventionContext()
}
    ICppCallingConventionContext is an interface to support dynamic dispatch.

type IDefaultCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsDefaultCaseContext differentiates from other interfaces.
	IsDefaultCaseContext()
}
    IDefaultCaseContext is an interface to support dynamic dispatch.

type IDeferStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFER() antlr.TerminalNode
	AssignmentStmt() IAssignmentStmtContext
	Expression() IExpressionContext

	// IsDeferStmtContext differentiates from other interfaces.
	IsDeferStmtContext()
}
    IDeferStmtContext is an interface to support dynamic dispatch.

type IDeinitDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEINIT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	SELF() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext

	// IsDeinitDeclContext differentiates from other interfaces.
	IsDeinitDeclContext()
}
    IDeinitDeclContext is an interface to support dynamic dispatch.

type IEnumDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	COLON() antlr.TerminalNode
	PrimitiveType() IPrimitiveTypeContext
	AllEnumMember() []IEnumMemberContext
	EnumMember(i int) IEnumMemberContext

	// IsEnumDeclContext differentiates from other interfaces.
	IsEnumDeclContext()
}
    IEnumDeclContext is an interface to support dynamic dispatch.

type IEnumMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsEnumMemberContext differentiates from other interfaces.
	IsEnumMemberContext()
}
    IEnumMemberContext is an interface to support dynamic dispatch.

type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRelationalExpression() []IRelationalExpressionContext
	RelationalExpression(i int) IRelationalExpressionContext
	AllEQ() []antlr.TerminalNode
	EQ(i int) antlr.TerminalNode
	AllNE() []antlr.TerminalNode
	NE(i int) antlr.TerminalNode

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}
    IEqualityExpressionContext is an interface to support dynamic dispatch.

type IExceptClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXCEPT() antlr.TerminalNode
	QualifiedIdentifier() IQualifiedIdentifierContext
	Block() IBlockContext
	IDENTIFIER() antlr.TerminalNode

	// IsExceptClauseContext differentiates from other interfaces.
	IsExceptClauseContext()
}
    IExceptClauseContext is an interface to support dynamic dispatch.

type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOrExpression() ILogicalOrExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}
    IExpressionContext is an interface to support dynamic dispatch.

type IExpressionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsExpressionStmtContext differentiates from other interfaces.
	IsExpressionStmtContext()
}
    IExpressionStmtContext is an interface to support dynamic dispatch.

type IExternCConstDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsExternCConstDeclContext differentiates from other interfaces.
	IsExternCConstDeclContext()
}
    IExternCConstDeclContext is an interface to support dynamic dispatch.

type IExternCDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTERN() antlr.TerminalNode
	C_LANG() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllExternCMember() []IExternCMemberContext
	ExternCMember(i int) IExternCMemberContext

	// IsExternCDeclContext differentiates from other interfaces.
	IsExternCDeclContext()
}
    IExternCDeclContext is an interface to support dynamic dispatch.

type IExternCFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	CCallingConvention() ICCallingConventionContext
	STRING_LITERAL() antlr.TerminalNode
	ExternCParameterList() IExternCParameterListContext
	Type_() ITypeContext

	// IsExternCFunctionDeclContext differentiates from other interfaces.
	IsExternCFunctionDeclContext()
}
    IExternCFunctionDeclContext is an interface to support dynamic dispatch.

type IExternCMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExternCFunctionDecl() IExternCFunctionDeclContext
	ExternCConstDecl() IExternCConstDeclContext
	ExternCTypeAlias() IExternCTypeAliasContext
	ExternCOpaqueStructDecl() IExternCOpaqueStructDeclContext

	// IsExternCMemberContext differentiates from other interfaces.
	IsExternCMemberContext()
}
    IExternCMemberContext is an interface to support dynamic dispatch.

type IExternCOpaqueStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPAQUE() antlr.TerminalNode
	STRUCT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode

	// IsExternCOpaqueStructDeclContext differentiates from other interfaces.
	IsExternCOpaqueStructDeclContext()
}
    IExternCOpaqueStructDeclContext is an interface to support dynamic dispatch.

type IExternCParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode

	// IsExternCParameterContext differentiates from other interfaces.
	IsExternCParameterContext()
}
    IExternCParameterContext is an interface to support dynamic dispatch.

type IExternCParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExternCParameter() []IExternCParameterContext
	ExternCParameter(i int) IExternCParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	ELLIPSIS() antlr.TerminalNode

	// IsExternCParameterListContext differentiates from other interfaces.
	IsExternCParameterListContext()
}
    IExternCParameterListContext is an interface to support dynamic dispatch.

type IExternCTypeAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Type_() ITypeContext

	// IsExternCTypeAliasContext differentiates from other interfaces.
	IsExternCTypeAliasContext()
}
    IExternCTypeAliasContext is an interface to support dynamic dispatch.

type IExternCppClassDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLASS() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	ABSTRACT() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	AllExternCppClassMember() []IExternCppClassMemberContext
	ExternCppClassMember(i int) IExternCppClassMemberContext

	// IsExternCppClassDeclContext differentiates from other interfaces.
	IsExternCppClassDeclContext()
}
    IExternCppClassDeclContext is an interface to support dynamic dispatch.

type IExternCppClassMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExternCppConstructorDecl() IExternCppConstructorDeclContext
	ExternCppDestructorDecl() IExternCppDestructorDeclContext
	ExternCppMethodDecl() IExternCppMethodDeclContext

	// IsExternCppClassMemberContext differentiates from other interfaces.
	IsExternCppClassMemberContext()
}
    IExternCppClassMemberContext is an interface to support dynamic dispatch.

type IExternCppConstDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsExternCppConstDeclContext differentiates from other interfaces.
	IsExternCppConstDeclContext()
}
    IExternCppConstDeclContext is an interface to support dynamic dispatch.

type IExternCppConstructorDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEW() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Type_() ITypeContext
	ExternCppParameterList() IExternCppParameterListContext

	// IsExternCppConstructorDeclContext differentiates from other interfaces.
	IsExternCppConstructorDeclContext()
}
    IExternCppConstructorDeclContext is an interface to support dynamic
    dispatch.

type IExternCppDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTERN() antlr.TerminalNode
	CPP_LANG() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllExternCppMember() []IExternCppMemberContext
	ExternCppMember(i int) IExternCppMemberContext

	// IsExternCppDeclContext differentiates from other interfaces.
	IsExternCppDeclContext()
}
    IExternCppDeclContext is an interface to support dynamic dispatch.

type IExternCppDestructorDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DELETE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	ExternCppSelfParam() IExternCppSelfParamContext
	RPAREN() antlr.TerminalNode
	Type_() ITypeContext

	// IsExternCppDestructorDeclContext differentiates from other interfaces.
	IsExternCppDestructorDeclContext()
}
    IExternCppDestructorDeclContext is an interface to support dynamic dispatch.

type IExternCppFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	CppCallingConvention() ICppCallingConventionContext
	STRING_LITERAL() antlr.TerminalNode
	ExternCppParameterList() IExternCppParameterListContext
	Type_() ITypeContext

	// IsExternCppFunctionDeclContext differentiates from other interfaces.
	IsExternCppFunctionDeclContext()
}
    IExternCppFunctionDeclContext is an interface to support dynamic dispatch.

type IExternCppMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExternCppNamespaceDecl() IExternCppNamespaceDeclContext
	ExternCppFunctionDecl() IExternCppFunctionDeclContext
	ExternCppConstDecl() IExternCppConstDeclContext
	ExternCppTypeAlias() IExternCppTypeAliasContext
	ExternCppOpaqueClassDecl() IExternCppOpaqueClassDeclContext
	ExternCppClassDecl() IExternCppClassDeclContext

	// IsExternCppMemberContext differentiates from other interfaces.
	IsExternCppMemberContext()
}
    IExternCppMemberContext is an interface to support dynamic dispatch.

type IExternCppMethodDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	ExternCppMethodParams() IExternCppMethodParamsContext
	RPAREN() antlr.TerminalNode
	CppCallingConvention() ICppCallingConventionContext
	VIRTUAL() antlr.TerminalNode
	STATIC() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	CONST() antlr.TerminalNode
	Type_() ITypeContext

	// IsExternCppMethodDeclContext differentiates from other interfaces.
	IsExternCppMethodDeclContext()
}
    IExternCppMethodDeclContext is an interface to support dynamic dispatch.

type IExternCppMethodParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExternCppSelfParam() IExternCppSelfParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllExternCppParameter() []IExternCppParameterContext
	ExternCppParameter(i int) IExternCppParameterContext
	ExternCppParameterList() IExternCppParameterListContext

	// IsExternCppMethodParamsContext differentiates from other interfaces.
	IsExternCppMethodParamsContext()
}
    IExternCppMethodParamsContext is an interface to support dynamic dispatch.

type IExternCppNamespaceDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAMESPACE() antlr.TerminalNode
	ExternNamespacePath() IExternNamespacePathContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllExternCppMember() []IExternCppMemberContext
	ExternCppMember(i int) IExternCppMemberContext

	// IsExternCppNamespaceDeclContext differentiates from other interfaces.
	IsExternCppNamespaceDeclContext()
}
    IExternCppNamespaceDeclContext is an interface to support dynamic dispatch.

type IExternCppOpaqueClassDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPAQUE() antlr.TerminalNode
	CLASS() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode

	// IsExternCppOpaqueClassDeclContext differentiates from other interfaces.
	IsExternCppOpaqueClassDeclContext()
}
    IExternCppOpaqueClassDeclContext is an interface to support dynamic
    dispatch.

type IExternCppParamTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	Type_() ITypeContext
	CONST() antlr.TerminalNode
	AMP() antlr.TerminalNode

	// IsExternCppParamTypeContext differentiates from other interfaces.
	IsExternCppParamTypeContext()
}
    IExternCppParamTypeContext is an interface to support dynamic dispatch.

type IExternCppParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExternCppParamType() IExternCppParamTypeContext
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode

	// IsExternCppParameterContext differentiates from other interfaces.
	IsExternCppParameterContext()
}
    IExternCppParameterContext is an interface to support dynamic dispatch.

type IExternCppParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExternCppParameter() []IExternCppParameterContext
	ExternCppParameter(i int) IExternCppParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	ELLIPSIS() antlr.TerminalNode

	// IsExternCppParameterListContext differentiates from other interfaces.
	IsExternCppParameterListContext()
}
    IExternCppParameterListContext is an interface to support dynamic dispatch.

type IExternCppSelfParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELF() antlr.TerminalNode
	STAR() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	CONST() antlr.TerminalNode

	// IsExternCppSelfParamContext differentiates from other interfaces.
	IsExternCppSelfParamContext()
}
    IExternCppSelfParamContext is an interface to support dynamic dispatch.

type IExternCppTypeAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Type_() ITypeContext

	// IsExternCppTypeAliasContext differentiates from other interfaces.
	IsExternCppTypeAliasContext()
}
    IExternCppTypeAliasContext is an interface to support dynamic dispatch.

type IExternNamespacePathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsExternNamespacePathContext differentiates from other interfaces.
	IsExternNamespacePathContext()
}
    IExternNamespacePathContext is an interface to support dynamic dispatch.

type IFieldInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Expression() IExpressionContext

	// IsFieldInitContext differentiates from other interfaces.
	IsFieldInitContext()
}
    IFieldInitContext is an interface to support dynamic dispatch.

type IFinallyClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FINALLY() antlr.TerminalNode
	Block() IBlockContext

	// IsFinallyClauseContext differentiates from other interfaces.
	IsFinallyClauseContext()
}
    IFinallyClauseContext is an interface to support dynamic dispatch.

type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	Block() IBlockContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	VariableDecl() IVariableDeclContext
	AllAssignmentStmt() []IAssignmentStmtContext
	AssignmentStmt(i int) IAssignmentStmtContext
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	IN() antlr.TerminalNode
	COMMA() antlr.TerminalNode

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}
    IForStmtContext is an interface to support dynamic dispatch.

type IFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	GenericParams() IGenericParamsContext
	ParameterList() IParameterListContext
	ReturnType() IReturnTypeContext
	ASYNC() antlr.TerminalNode
	PROCESS() antlr.TerminalNode
	CONTAINER() antlr.TerminalNode

	// IsFunctionDeclContext differentiates from other interfaces.
	IsFunctionDeclContext()
}
    IFunctionDeclContext is an interface to support dynamic dispatch.

type IFunctionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ASYNC() antlr.TerminalNode
	GenericParams() IGenericParamsContext
	TypeList() ITypeListContext
	ReturnType() IReturnTypeContext

	// IsFunctionTypeContext differentiates from other interfaces.
	IsFunctionTypeContext()
}
    IFunctionTypeContext is an interface to support dynamic dispatch.

type IGenericArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Expression() IExpressionContext

	// IsGenericArgContext differentiates from other interfaces.
	IsGenericArgContext()
}
    IGenericArgContext is an interface to support dynamic dispatch.

type IGenericArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllGenericArg() []IGenericArgContext
	GenericArg(i int) IGenericArgContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsGenericArgListContext differentiates from other interfaces.
	IsGenericArgListContext()
}
    IGenericArgListContext is an interface to support dynamic dispatch.

type IGenericArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LT() antlr.TerminalNode
	GenericArgList() IGenericArgListContext
	GT() antlr.TerminalNode

	// IsGenericArgsContext differentiates from other interfaces.
	IsGenericArgsContext()
}
    IGenericArgsContext is an interface to support dynamic dispatch.

type IGenericParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsGenericParamContext differentiates from other interfaces.
	IsGenericParamContext()
}
    IGenericParamContext is an interface to support dynamic dispatch.

type IGenericParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllGenericParam() []IGenericParamContext
	GenericParam(i int) IGenericParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsGenericParamListContext differentiates from other interfaces.
	IsGenericParamListContext()
}
    IGenericParamListContext is an interface to support dynamic dispatch.

type IGenericParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LT() antlr.TerminalNode
	GenericParamList() IGenericParamListContext
	GT() antlr.TerminalNode

	// IsGenericParamsContext differentiates from other interfaces.
	IsGenericParamsContext()
}
    IGenericParamsContext is an interface to support dynamic dispatch.

type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIF() []antlr.TerminalNode
	IF(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllELSE() []antlr.TerminalNode
	ELSE(i int) antlr.TerminalNode

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}
    IIfStmtContext is an interface to support dynamic dispatch.

type IImportDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllImportSpec() []IImportSpecContext
	ImportSpec(i int) IImportSpecContext

	// IsImportDeclContext differentiates from other interfaces.
	IsImportDeclContext()
}
    IImportDeclContext is an interface to support dynamic dispatch.

type IImportSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode

	// IsImportSpecContext differentiates from other interfaces.
	IsImportSpecContext()
}
    IImportSpecContext is an interface to support dynamic dispatch.

type IInitDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INIT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	SELF() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	ELLIPSIS() antlr.TerminalNode

	// IsInitDeclContext differentiates from other interfaces.
	IsInitDeclContext()
}
    IInitDeclContext is an interface to support dynamic dispatch.

type IInitializerEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COLON() antlr.TerminalNode

	// IsInitializerEntryContext differentiates from other interfaces.
	IsInitializerEntryContext()
}
    IInitializerEntryContext is an interface to support dynamic dispatch.

type IInitializerListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllInitializerEntry() []IInitializerEntryContext
	InitializerEntry(i int) IInitializerEntryContext

	// IsInitializerListContext differentiates from other interfaces.
	IsInitializerListContext()
}
    IInitializerListContext is an interface to support dynamic dispatch.

type ILambdaExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	FAT_ARROW() antlr.TerminalNode
	Block() IBlockContext
	LambdaParamList() ILambdaParamListContext
	ASYNC() antlr.TerminalNode
	PROCESS() antlr.TerminalNode
	CONTAINER() antlr.TerminalNode
	Expression() IExpressionContext

	// IsLambdaExpressionContext differentiates from other interfaces.
	IsLambdaExpressionContext()
}
    ILambdaExpressionContext is an interface to support dynamic dispatch.

type ILambdaParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsLambdaParamContext differentiates from other interfaces.
	IsLambdaParamContext()
}
    ILambdaParamContext is an interface to support dynamic dispatch.

type ILambdaParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLambdaParam() []ILambdaParamContext
	LambdaParam(i int) ILambdaParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsLambdaParamListContext differentiates from other interfaces.
	IsLambdaParamListContext()
}
    ILambdaParamListContext is an interface to support dynamic dispatch.

type ILeftHandSideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	PostfixExpression() IPostfixExpressionContext
	DOT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACKET() antlr.TerminalNode
	Expression() IExpressionContext
	RBRACKET() antlr.TerminalNode

	// IsLeftHandSideContext differentiates from other interfaces.
	IsLeftHandSideContext()
}
    ILeftHandSideContext is an interface to support dynamic dispatch.

type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER_LITERAL() antlr.TerminalNode
	FLOAT_LITERAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	CHAR_LITERAL() antlr.TerminalNode
	BOOLEAN_LITERAL() antlr.TerminalNode
	NULL() antlr.TerminalNode
	InitializerList() IInitializerListContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}
    ILiteralContext is an interface to support dynamic dispatch.

type ILogicalAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBitOrExpression() []IBitOrExpressionContext
	BitOrExpression(i int) IBitOrExpressionContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsLogicalAndExpressionContext differentiates from other interfaces.
	IsLogicalAndExpressionContext()
}
    ILogicalAndExpressionContext is an interface to support dynamic dispatch.

type ILogicalOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLogicalAndExpression() []ILogicalAndExpressionContext
	LogicalAndExpression(i int) ILogicalAndExpressionContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsLogicalOrExpressionContext differentiates from other interfaces.
	IsLogicalOrExpressionContext()
}
    ILogicalOrExpressionContext is an interface to support dynamic dispatch.

type IMethodDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	SELF() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	ASYNC() antlr.TerminalNode
	GenericParams() IGenericParamsContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	ReturnType() IReturnTypeContext

	// IsMethodDeclContext differentiates from other interfaces.
	IsMethodDeclContext()
}
    IMethodDeclContext is an interface to support dynamic dispatch.

type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnaryExpression() []IUnaryExpressionContext
	UnaryExpression(i int) IUnaryExpressionContext
	AllSTAR() []antlr.TerminalNode
	STAR(i int) antlr.TerminalNode
	AllSLASH() []antlr.TerminalNode
	SLASH(i int) antlr.TerminalNode
	AllPERCENT() []antlr.TerminalNode
	PERCENT(i int) antlr.TerminalNode

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}
    IMultiplicativeExpressionContext is an interface to support dynamic
    dispatch.

type IMutatingDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUTATING() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	SELF() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	ReturnType() IReturnTypeContext

	// IsMutatingDeclContext differentiates from other interfaces.
	IsMutatingDeclContext()
}
    IMutatingDeclContext is an interface to support dynamic dispatch.

type INamespaceDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAMESPACE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsNamespaceDeclContext differentiates from other interfaces.
	IsNamespaceDeclContext()
}
    INamespaceDeclContext is an interface to support dynamic dispatch.

type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	SELF() antlr.TerminalNode

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}
    IParameterContext is an interface to support dynamic dispatch.

type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	ELLIPSIS() antlr.TerminalNode

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}
    IParameterListContext is an interface to support dynamic dispatch.

type IPointerTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	Type_() ITypeContext
	CONST() antlr.TerminalNode

	// IsPointerTypeContext differentiates from other interfaces.
	IsPointerTypeContext()
}
    IPointerTypeContext is an interface to support dynamic dispatch.

type IPostfixExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpression() IPrimaryExpressionContext
	AllPostfixOp() []IPostfixOpContext
	PostfixOp(i int) IPostfixOpContext

	// IsPostfixExpressionContext differentiates from other interfaces.
	IsPostfixExpressionContext()
}
    IPostfixExpressionContext is an interface to support dynamic dispatch.

type IPostfixOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgumentList() IArgumentListContext
	LBRACKET() antlr.TerminalNode
	Expression() IExpressionContext
	RBRACKET() antlr.TerminalNode
	INCREMENT() antlr.TerminalNode
	DECREMENT() antlr.TerminalNode

	// IsPostfixOpContext differentiates from other interfaces.
	IsPostfixOpContext()
}
    IPostfixOpContext is an interface to support dynamic dispatch.

type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ComputeExpression() IComputeExpressionContext
	Literal() ILiteralContext
	StructLiteral() IStructLiteralContext
	SizeofExpression() ISizeofExpressionContext
	AlignofExpression() IAlignofExpressionContext
	LambdaExpression() ILambdaExpressionContext
	AnonymousFuncExpression() IAnonymousFuncExpressionContext
	TupleExpression() ITupleExpressionContext
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	QualifiedIdentifier() IQualifiedIdentifierContext
	GenericArgs() IGenericArgsContext
	ArgumentList() IArgumentListContext
	IDENTIFIER() antlr.TerminalNode

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}
    IPrimaryExpressionContext is an interface to support dynamic dispatch.

type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT8() antlr.TerminalNode
	INT16() antlr.TerminalNode
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT8() antlr.TerminalNode
	UINT16() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	USIZE() antlr.TerminalNode
	ISIZE() antlr.TerminalNode
	FLOAT32() antlr.TerminalNode
	FLOAT64() antlr.TerminalNode
	FLOAT16() antlr.TerminalNode
	BFLOAT16() antlr.TerminalNode
	BYTE() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	STRING() antlr.TerminalNode
	VOID() antlr.TerminalNode

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}
    IPrimitiveTypeContext is an interface to support dynamic dispatch.

type IQualifiedIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsQualifiedIdentifierContext differentiates from other interfaces.
	IsQualifiedIdentifierContext()
}
    IQualifiedIdentifierContext is an interface to support dynamic dispatch.

type IQualifiedTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	GenericArgs() IGenericArgsContext

	// IsQualifiedTypeContext differentiates from other interfaces.
	IsQualifiedTypeContext()
}
    IQualifiedTypeContext is an interface to support dynamic dispatch.

type IRangeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAdditiveExpression() []IAdditiveExpressionContext
	AdditiveExpression(i int) IAdditiveExpressionContext
	RANGE() antlr.TerminalNode

	// IsRangeExpressionContext differentiates from other interfaces.
	IsRangeExpressionContext()
}
    IRangeExpressionContext is an interface to support dynamic dispatch.

type IReferenceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AMP() antlr.TerminalNode
	Type_() ITypeContext
	CONST() antlr.TerminalNode

	// IsReferenceTypeContext differentiates from other interfaces.
	IsReferenceTypeContext()
}
    IReferenceTypeContext is an interface to support dynamic dispatch.

type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllShiftExpression() []IShiftExpressionContext
	ShiftExpression(i int) IShiftExpressionContext
	AllLT() []antlr.TerminalNode
	LT(i int) antlr.TerminalNode
	AllLE() []antlr.TerminalNode
	LE(i int) antlr.TerminalNode
	AllGT() []antlr.TerminalNode
	GT(i int) antlr.TerminalNode
	AllGE() []antlr.TerminalNode
	GE(i int) antlr.TerminalNode

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}
    IRelationalExpressionContext is an interface to support dynamic dispatch.

type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	TupleExpression() ITupleExpressionContext
	Expression() IExpressionContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}
    IReturnStmtContext is an interface to support dynamic dispatch.

type IReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	LPAREN() antlr.TerminalNode
	TypeList() ITypeListContext
	RPAREN() antlr.TerminalNode

	// IsReturnTypeContext differentiates from other interfaces.
	IsReturnTypeContext()
}
    IReturnTypeContext is an interface to support dynamic dispatch.

type IShiftExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRangeExpression() []IRangeExpressionContext
	RangeExpression(i int) IRangeExpressionContext
	AllLT() []antlr.TerminalNode
	LT(i int) antlr.TerminalNode
	AllGT() []antlr.TerminalNode
	GT(i int) antlr.TerminalNode

	// IsShiftExpressionContext differentiates from other interfaces.
	IsShiftExpressionContext()
}
    IShiftExpressionContext is an interface to support dynamic dispatch.

type ISizeofExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SIZEOF() antlr.TerminalNode
	LT() antlr.TerminalNode
	Type_() ITypeContext
	GT() antlr.TerminalNode

	// IsSizeofExpressionContext differentiates from other interfaces.
	IsSizeofExpressionContext()
}
    ISizeofExpressionContext is an interface to support dynamic dispatch.

type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	ReturnStmt() IReturnStmtContext
	BreakStmt() IBreakStmtContext
	ContinueStmt() IContinueStmtContext
	IfStmt() IIfStmtContext
	ForStmt() IForStmtContext
	SwitchStmt() ISwitchStmtContext
	TryStmt() ITryStmtContext
	ThrowStmt() IThrowStmtContext
	DeferStmt() IDeferStmtContext
	VariableDecl() IVariableDeclContext
	ConstDecl() IConstDeclContext
	AssignmentStmt() IAssignmentStmtContext
	ExpressionStmt() IExpressionStmtContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}
    IStatementContext is an interface to support dynamic dispatch.

type IStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	ComputeMarker() IComputeMarkerContext
	GenericParams() IGenericParamsContext
	AllStructMember() []IStructMemberContext
	StructMember(i int) IStructMemberContext

	// IsStructDeclContext differentiates from other interfaces.
	IsStructDeclContext()
}
    IStructDeclContext is an interface to support dynamic dispatch.

type IStructFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsStructFieldContext differentiates from other interfaces.
	IsStructFieldContext()
}
    IStructFieldContext is an interface to support dynamic dispatch.

type IStructLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	QualifiedIdentifier() IQualifiedIdentifierContext
	GenericArgs() IGenericArgsContext
	AllFieldInit() []IFieldInitContext
	FieldInit(i int) IFieldInitContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsStructLiteralContext differentiates from other interfaces.
	IsStructLiteralContext()
}
    IStructLiteralContext is an interface to support dynamic dispatch.

type IStructMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StructField() IStructFieldContext
	FunctionDecl() IFunctionDeclContext
	MutatingDecl() IMutatingDeclContext
	InitDecl() IInitDeclContext

	// IsStructMemberContext differentiates from other interfaces.
	IsStructMemberContext()
}
    IStructMemberContext is an interface to support dynamic dispatch.

type ISwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COLON() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsSwitchCaseContext differentiates from other interfaces.
	IsSwitchCaseContext()
}
    ISwitchCaseContext is an interface to support dynamic dispatch.

type ISwitchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expression() IExpressionContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllSwitchCase() []ISwitchCaseContext
	SwitchCase(i int) ISwitchCaseContext
	DefaultCase() IDefaultCaseContext

	// IsSwitchStmtContext differentiates from other interfaces.
	IsSwitchStmtContext()
}
    ISwitchStmtContext is an interface to support dynamic dispatch.

type IThrowStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	THROW() antlr.TerminalNode
	Expression() IExpressionContext

	// IsThrowStmtContext differentiates from other interfaces.
	IsThrowStmtContext()
}
    IThrowStmtContext is an interface to support dynamic dispatch.

type ITopLevelDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionDecl() IFunctionDeclContext
	StructDecl() IStructDeclContext
	ClassDecl() IClassDeclContext
	EnumDecl() IEnumDeclContext
	MethodDecl() IMethodDeclContext
	MutatingDecl() IMutatingDeclContext
	DeinitDecl() IDeinitDeclContext
	VariableDecl() IVariableDeclContext
	ConstDecl() IConstDeclContext
	ExternCDecl() IExternCDeclContext
	ExternCppDecl() IExternCppDeclContext

	// IsTopLevelDeclContext differentiates from other interfaces.
	IsTopLevelDeclContext()
}
    ITopLevelDeclContext is an interface to support dynamic dispatch.

type ITryStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRY() antlr.TerminalNode
	Block() IBlockContext
	AllExceptClause() []IExceptClauseContext
	ExceptClause(i int) IExceptClauseContext
	FinallyClause() IFinallyClauseContext

	// IsTryStmtContext differentiates from other interfaces.
	IsTryStmtContext()
}
    ITryStmtContext is an interface to support dynamic dispatch.

type ITupleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsTupleExpressionContext differentiates from other interfaces.
	IsTupleExpressionContext()
}
    ITupleExpressionContext is an interface to support dynamic dispatch.

type ITuplePatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsTuplePatternContext differentiates from other interfaces.
	IsTuplePatternContext()
}
    ITuplePatternContext is an interface to support dynamic dispatch.

type ITupleTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	TypeList() ITypeListContext
	RPAREN() antlr.TerminalNode

	// IsTupleTypeContext differentiates from other interfaces.
	IsTupleTypeContext()
}
    ITupleTypeContext is an interface to support dynamic dispatch.

type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimitiveType() IPrimitiveTypeContext
	PointerType() IPointerTypeContext
	ReferenceType() IReferenceTypeContext
	QualifiedType() IQualifiedTypeContext
	FunctionType() IFunctionTypeContext
	ArrayType() IArrayTypeContext
	IDENTIFIER() antlr.TerminalNode
	GenericArgs() IGenericArgsContext
	UNDERSCORE() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}
    ITypeContext is an interface to support dynamic dispatch.

type ITypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsTypeListContext differentiates from other interfaces.
	IsTypeListContext()
}
    ITypeListContext is an interface to support dynamic dispatch.

type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpression() IUnaryExpressionContext
	MINUS() antlr.TerminalNode
	NOT() antlr.TerminalNode
	BIT_NOT() antlr.TerminalNode
	STAR() antlr.TerminalNode
	AMP() antlr.TerminalNode
	AWAIT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	INCREMENT() antlr.TerminalNode
	DECREMENT() antlr.TerminalNode
	PostfixExpression() IPostfixExpressionContext

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}
    IUnaryExpressionContext is an interface to support dynamic dispatch.

type IVariableDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LET() antlr.TerminalNode
	TuplePattern() ITuplePatternContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	TupleType() ITupleTypeContext
	IDENTIFIER() antlr.TerminalNode
	Type_() ITypeContext

	// IsVariableDeclContext differentiates from other interfaces.
	IsVariableDeclContext()
}
    IVariableDeclContext is an interface to support dynamic dispatch.

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyIfStmtContext() *IfStmtContext

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *IfStmtContext) AllBlock() []IBlockContext

func (s *IfStmtContext) AllELSE() []antlr.TerminalNode

func (s *IfStmtContext) AllExpression() []IExpressionContext

func (s *IfStmtContext) AllIF() []antlr.TerminalNode

func (s *IfStmtContext) Block(i int) IBlockContext

func (s *IfStmtContext) ELSE(i int) antlr.TerminalNode

func (s *IfStmtContext) Expression(i int) IExpressionContext

func (s *IfStmtContext) GetParser() antlr.Parser

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext

func (s *IfStmtContext) IF(i int) antlr.TerminalNode

func (*IfStmtContext) IsIfStmtContext()

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ImportDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyImportDeclContext() *ImportDeclContext

func NewImportDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclContext

func (s *ImportDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ImportDeclContext) AllImportSpec() []IImportSpecContext

func (s *ImportDeclContext) GetParser() antlr.Parser

func (s *ImportDeclContext) GetRuleContext() antlr.RuleContext

func (s *ImportDeclContext) IMPORT() antlr.TerminalNode

func (s *ImportDeclContext) ImportSpec(i int) IImportSpecContext

func (*ImportDeclContext) IsImportDeclContext()

func (s *ImportDeclContext) LPAREN() antlr.TerminalNode

func (s *ImportDeclContext) RPAREN() antlr.TerminalNode

func (s *ImportDeclContext) STRING_LITERAL() antlr.TerminalNode

func (s *ImportDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ImportSpecContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyImportSpecContext() *ImportSpecContext

func NewImportSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportSpecContext

func (s *ImportSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ImportSpecContext) GetParser() antlr.Parser

func (s *ImportSpecContext) GetRuleContext() antlr.RuleContext

func (*ImportSpecContext) IsImportSpecContext()

func (s *ImportSpecContext) STRING_LITERAL() antlr.TerminalNode

func (s *ImportSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type InitDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyInitDeclContext() *InitDeclContext

func NewInitDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitDeclContext

func (s *InitDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *InitDeclContext) AllCOMMA() []antlr.TerminalNode

func (s *InitDeclContext) AllParameter() []IParameterContext

func (s *InitDeclContext) Block() IBlockContext

func (s *InitDeclContext) COLON() antlr.TerminalNode

func (s *InitDeclContext) COMMA(i int) antlr.TerminalNode

func (s *InitDeclContext) ELLIPSIS() antlr.TerminalNode

func (s *InitDeclContext) GetParser() antlr.Parser

func (s *InitDeclContext) GetRuleContext() antlr.RuleContext

func (s *InitDeclContext) IDENTIFIER() antlr.TerminalNode

func (s *InitDeclContext) INIT() antlr.TerminalNode

func (*InitDeclContext) IsInitDeclContext()

func (s *InitDeclContext) LPAREN() antlr.TerminalNode

func (s *InitDeclContext) Parameter(i int) IParameterContext

func (s *InitDeclContext) RPAREN() antlr.TerminalNode

func (s *InitDeclContext) SELF() antlr.TerminalNode

func (s *InitDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *InitDeclContext) Type_() ITypeContext

type InitializerEntryContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyInitializerEntryContext() *InitializerEntryContext

func NewInitializerEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitializerEntryContext

func (s *InitializerEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *InitializerEntryContext) AllExpression() []IExpressionContext

func (s *InitializerEntryContext) COLON() antlr.TerminalNode

func (s *InitializerEntryContext) Expression(i int) IExpressionContext

func (s *InitializerEntryContext) GetParser() antlr.Parser

func (s *InitializerEntryContext) GetRuleContext() antlr.RuleContext

func (*InitializerEntryContext) IsInitializerEntryContext()

func (s *InitializerEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type InitializerListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyInitializerListContext() *InitializerListContext

func NewInitializerListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitializerListContext

func (s *InitializerListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *InitializerListContext) AllCOMMA() []antlr.TerminalNode

func (s *InitializerListContext) AllExpression() []IExpressionContext

func (s *InitializerListContext) AllInitializerEntry() []IInitializerEntryContext

func (s *InitializerListContext) COMMA(i int) antlr.TerminalNode

func (s *InitializerListContext) Expression(i int) IExpressionContext

func (s *InitializerListContext) GetParser() antlr.Parser

func (s *InitializerListContext) GetRuleContext() antlr.RuleContext

func (s *InitializerListContext) InitializerEntry(i int) IInitializerEntryContext

func (*InitializerListContext) IsInitializerListContext()

func (s *InitializerListContext) LBRACE() antlr.TerminalNode

func (s *InitializerListContext) RBRACE() antlr.TerminalNode

func (s *InitializerListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type LambdaExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyLambdaExpressionContext() *LambdaExpressionContext

func NewLambdaExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaExpressionContext

func (s *LambdaExpressionContext) ASYNC() antlr.TerminalNode

func (s *LambdaExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *LambdaExpressionContext) Block() IBlockContext

func (s *LambdaExpressionContext) CONTAINER() antlr.TerminalNode

func (s *LambdaExpressionContext) Expression() IExpressionContext

func (s *LambdaExpressionContext) FAT_ARROW() antlr.TerminalNode

func (s *LambdaExpressionContext) GetParser() antlr.Parser

func (s *LambdaExpressionContext) GetRuleContext() antlr.RuleContext

func (*LambdaExpressionContext) IsLambdaExpressionContext()

func (s *LambdaExpressionContext) LPAREN() antlr.TerminalNode

func (s *LambdaExpressionContext) LambdaParamList() ILambdaParamListContext

func (s *LambdaExpressionContext) PROCESS() antlr.TerminalNode

func (s *LambdaExpressionContext) RPAREN() antlr.TerminalNode

func (s *LambdaExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type LambdaParamContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyLambdaParamContext() *LambdaParamContext

func NewLambdaParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaParamContext

func (s *LambdaParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *LambdaParamContext) COLON() antlr.TerminalNode

func (s *LambdaParamContext) GetParser() antlr.Parser

func (s *LambdaParamContext) GetRuleContext() antlr.RuleContext

func (s *LambdaParamContext) IDENTIFIER() antlr.TerminalNode

func (*LambdaParamContext) IsLambdaParamContext()

func (s *LambdaParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *LambdaParamContext) Type_() ITypeContext

type LambdaParamListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyLambdaParamListContext() *LambdaParamListContext

func NewLambdaParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaParamListContext

func (s *LambdaParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *LambdaParamListContext) AllCOMMA() []antlr.TerminalNode

func (s *LambdaParamListContext) AllLambdaParam() []ILambdaParamContext

func (s *LambdaParamListContext) COMMA(i int) antlr.TerminalNode

func (s *LambdaParamListContext) GetParser() antlr.Parser

func (s *LambdaParamListContext) GetRuleContext() antlr.RuleContext

func (*LambdaParamListContext) IsLambdaParamListContext()

func (s *LambdaParamListContext) LambdaParam(i int) ILambdaParamContext

func (s *LambdaParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type LeftHandSideContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyLeftHandSideContext() *LeftHandSideContext

func NewLeftHandSideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftHandSideContext

func (s *LeftHandSideContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *LeftHandSideContext) DOT() antlr.TerminalNode

func (s *LeftHandSideContext) Expression() IExpressionContext

func (s *LeftHandSideContext) GetParser() antlr.Parser

func (s *LeftHandSideContext) GetRuleContext() antlr.RuleContext

func (s *LeftHandSideContext) IDENTIFIER() antlr.TerminalNode

func (*LeftHandSideContext) IsLeftHandSideContext()

func (s *LeftHandSideContext) LBRACKET() antlr.TerminalNode

func (s *LeftHandSideContext) PostfixExpression() IPostfixExpressionContext

func (s *LeftHandSideContext) RBRACKET() antlr.TerminalNode

func (s *LeftHandSideContext) STAR() antlr.TerminalNode

func (s *LeftHandSideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type LiteralContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyLiteralContext() *LiteralContext

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *LiteralContext) BOOLEAN_LITERAL() antlr.TerminalNode

func (s *LiteralContext) CHAR_LITERAL() antlr.TerminalNode

func (s *LiteralContext) FLOAT_LITERAL() antlr.TerminalNode

func (s *LiteralContext) GetParser() antlr.Parser

func (s *LiteralContext) GetRuleContext() antlr.RuleContext

func (s *LiteralContext) INTEGER_LITERAL() antlr.TerminalNode

func (s *LiteralContext) InitializerList() IInitializerListContext

func (*LiteralContext) IsLiteralContext()

func (s *LiteralContext) NULL() antlr.TerminalNode

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type LogicalAndExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyLogicalAndExpressionContext() *LogicalAndExpressionContext

func NewLogicalAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndExpressionContext

func (s *LogicalAndExpressionContext) AND(i int) antlr.TerminalNode

func (s *LogicalAndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *LogicalAndExpressionContext) AllAND() []antlr.TerminalNode

func (s *LogicalAndExpressionContext) AllBitOrExpression() []IBitOrExpressionContext

func (s *LogicalAndExpressionContext) BitOrExpression(i int) IBitOrExpressionContext

func (s *LogicalAndExpressionContext) GetParser() antlr.Parser

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext

func (*LogicalAndExpressionContext) IsLogicalAndExpressionContext()

func (s *LogicalAndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type LogicalOrExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyLogicalOrExpressionContext() *LogicalOrExpressionContext

func NewLogicalOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrExpressionContext

func (s *LogicalOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *LogicalOrExpressionContext) AllLogicalAndExpression() []ILogicalAndExpressionContext

func (s *LogicalOrExpressionContext) AllOR() []antlr.TerminalNode

func (s *LogicalOrExpressionContext) GetParser() antlr.Parser

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext

func (*LogicalOrExpressionContext) IsLogicalOrExpressionContext()

func (s *LogicalOrExpressionContext) LogicalAndExpression(i int) ILogicalAndExpressionContext

func (s *LogicalOrExpressionContext) OR(i int) antlr.TerminalNode

func (s *LogicalOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type MethodDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyMethodDeclContext() *MethodDeclContext

func NewMethodDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodDeclContext

func (s *MethodDeclContext) ASYNC() antlr.TerminalNode

func (s *MethodDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *MethodDeclContext) AllCOMMA() []antlr.TerminalNode

func (s *MethodDeclContext) AllIDENTIFIER() []antlr.TerminalNode

func (s *MethodDeclContext) AllParameter() []IParameterContext

func (s *MethodDeclContext) Block() IBlockContext

func (s *MethodDeclContext) COLON() antlr.TerminalNode

func (s *MethodDeclContext) COMMA(i int) antlr.TerminalNode

func (s *MethodDeclContext) FUNC() antlr.TerminalNode

func (s *MethodDeclContext) GenericParams() IGenericParamsContext

func (s *MethodDeclContext) GetParser() antlr.Parser

func (s *MethodDeclContext) GetRuleContext() antlr.RuleContext

func (s *MethodDeclContext) IDENTIFIER(i int) antlr.TerminalNode

func (*MethodDeclContext) IsMethodDeclContext()

func (s *MethodDeclContext) LPAREN() antlr.TerminalNode

func (s *MethodDeclContext) Parameter(i int) IParameterContext

func (s *MethodDeclContext) RPAREN() antlr.TerminalNode

func (s *MethodDeclContext) ReturnType() IReturnTypeContext

func (s *MethodDeclContext) SELF() antlr.TerminalNode

func (s *MethodDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *MethodDeclContext) Type_() ITypeContext

type MultiplicativeExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *MultiplicativeExpressionContext) AllPERCENT() []antlr.TerminalNode

func (s *MultiplicativeExpressionContext) AllSLASH() []antlr.TerminalNode

func (s *MultiplicativeExpressionContext) AllSTAR() []antlr.TerminalNode

func (s *MultiplicativeExpressionContext) AllUnaryExpression() []IUnaryExpressionContext

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext()

func (s *MultiplicativeExpressionContext) PERCENT(i int) antlr.TerminalNode

func (s *MultiplicativeExpressionContext) SLASH(i int) antlr.TerminalNode

func (s *MultiplicativeExpressionContext) STAR(i int) antlr.TerminalNode

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *MultiplicativeExpressionContext) UnaryExpression(i int) IUnaryExpressionContext

type MutatingDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyMutatingDeclContext() *MutatingDeclContext

func NewMutatingDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutatingDeclContext

func (s *MutatingDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *MutatingDeclContext) AllCOMMA() []antlr.TerminalNode

func (s *MutatingDeclContext) AllIDENTIFIER() []antlr.TerminalNode

func (s *MutatingDeclContext) AllParameter() []IParameterContext

func (s *MutatingDeclContext) Block() IBlockContext

func (s *MutatingDeclContext) COLON() antlr.TerminalNode

func (s *MutatingDeclContext) COMMA(i int) antlr.TerminalNode

func (s *MutatingDeclContext) GetParser() antlr.Parser

func (s *MutatingDeclContext) GetRuleContext() antlr.RuleContext

func (s *MutatingDeclContext) IDENTIFIER(i int) antlr.TerminalNode

func (*MutatingDeclContext) IsMutatingDeclContext()

func (s *MutatingDeclContext) LPAREN() antlr.TerminalNode

func (s *MutatingDeclContext) MUTATING() antlr.TerminalNode

func (s *MutatingDeclContext) Parameter(i int) IParameterContext

func (s *MutatingDeclContext) RPAREN() antlr.TerminalNode

func (s *MutatingDeclContext) ReturnType() IReturnTypeContext

func (s *MutatingDeclContext) SELF() antlr.TerminalNode

func (s *MutatingDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *MutatingDeclContext) Type_() ITypeContext

type NamespaceDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyNamespaceDeclContext() *NamespaceDeclContext

func NewNamespaceDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceDeclContext

func (s *NamespaceDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *NamespaceDeclContext) GetParser() antlr.Parser

func (s *NamespaceDeclContext) GetRuleContext() antlr.RuleContext

func (s *NamespaceDeclContext) IDENTIFIER() antlr.TerminalNode

func (*NamespaceDeclContext) IsNamespaceDeclContext()

func (s *NamespaceDeclContext) NAMESPACE() antlr.TerminalNode

func (s *NamespaceDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ParameterContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyParameterContext() *ParameterContext

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ParameterContext) COLON() antlr.TerminalNode

func (s *ParameterContext) GetParser() antlr.Parser

func (s *ParameterContext) GetRuleContext() antlr.RuleContext

func (s *ParameterContext) IDENTIFIER() antlr.TerminalNode

func (*ParameterContext) IsParameterContext()

func (s *ParameterContext) SELF() antlr.TerminalNode

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ParameterContext) Type_() ITypeContext

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyParameterListContext() *ParameterListContext

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ParameterListContext) AllCOMMA() []antlr.TerminalNode

func (s *ParameterListContext) AllParameter() []IParameterContext

func (s *ParameterListContext) COMMA(i int) antlr.TerminalNode

func (s *ParameterListContext) ELLIPSIS() antlr.TerminalNode

func (s *ParameterListContext) GetParser() antlr.Parser

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext

func (*ParameterListContext) IsParameterListContext()

func (s *ParameterListContext) Parameter(i int) IParameterContext

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type PointerTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyPointerTypeContext() *PointerTypeContext

func NewPointerTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerTypeContext

func (s *PointerTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *PointerTypeContext) CONST() antlr.TerminalNode

func (s *PointerTypeContext) GetParser() antlr.Parser

func (s *PointerTypeContext) GetRuleContext() antlr.RuleContext

func (*PointerTypeContext) IsPointerTypeContext()

func (s *PointerTypeContext) STAR() antlr.TerminalNode

func (s *PointerTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *PointerTypeContext) Type_() ITypeContext

type PostfixExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyPostfixExpressionContext() *PostfixExpressionContext

func NewPostfixExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixExpressionContext

func (s *PostfixExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *PostfixExpressionContext) AllPostfixOp() []IPostfixOpContext

func (s *PostfixExpressionContext) GetParser() antlr.Parser

func (s *PostfixExpressionContext) GetRuleContext() antlr.RuleContext

func (*PostfixExpressionContext) IsPostfixExpressionContext()

func (s *PostfixExpressionContext) PostfixOp(i int) IPostfixOpContext

func (s *PostfixExpressionContext) PrimaryExpression() IPrimaryExpressionContext

func (s *PostfixExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type PostfixOpContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyPostfixOpContext() *PostfixOpContext

func NewPostfixOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixOpContext

func (s *PostfixOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *PostfixOpContext) ArgumentList() IArgumentListContext

func (s *PostfixOpContext) DECREMENT() antlr.TerminalNode

func (s *PostfixOpContext) DOT() antlr.TerminalNode

func (s *PostfixOpContext) Expression() IExpressionContext

func (s *PostfixOpContext) GetParser() antlr.Parser

func (s *PostfixOpContext) GetRuleContext() antlr.RuleContext

func (s *PostfixOpContext) IDENTIFIER() antlr.TerminalNode

func (s *PostfixOpContext) INCREMENT() antlr.TerminalNode

func (*PostfixOpContext) IsPostfixOpContext()

func (s *PostfixOpContext) LBRACKET() antlr.TerminalNode

func (s *PostfixOpContext) LPAREN() antlr.TerminalNode

func (s *PostfixOpContext) RBRACKET() antlr.TerminalNode

func (s *PostfixOpContext) RPAREN() antlr.TerminalNode

func (s *PostfixOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *PrimaryExpressionContext) AlignofExpression() IAlignofExpressionContext

func (s *PrimaryExpressionContext) AnonymousFuncExpression() IAnonymousFuncExpressionContext

func (s *PrimaryExpressionContext) ArgumentList() IArgumentListContext

func (s *PrimaryExpressionContext) ComputeExpression() IComputeExpressionContext

func (s *PrimaryExpressionContext) Expression() IExpressionContext

func (s *PrimaryExpressionContext) GenericArgs() IGenericArgsContext

func (s *PrimaryExpressionContext) GetParser() antlr.Parser

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext

func (s *PrimaryExpressionContext) IDENTIFIER() antlr.TerminalNode

func (*PrimaryExpressionContext) IsPrimaryExpressionContext()

func (s *PrimaryExpressionContext) LPAREN() antlr.TerminalNode

func (s *PrimaryExpressionContext) LambdaExpression() ILambdaExpressionContext

func (s *PrimaryExpressionContext) Literal() ILiteralContext

func (s *PrimaryExpressionContext) QualifiedIdentifier() IQualifiedIdentifierContext

func (s *PrimaryExpressionContext) RPAREN() antlr.TerminalNode

func (s *PrimaryExpressionContext) SizeofExpression() ISizeofExpressionContext

func (s *PrimaryExpressionContext) StructLiteral() IStructLiteralContext

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *PrimaryExpressionContext) TupleExpression() ITupleExpressionContext

type PrimitiveTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *PrimitiveTypeContext) BFLOAT16() antlr.TerminalNode

func (s *PrimitiveTypeContext) BOOL() antlr.TerminalNode

func (s *PrimitiveTypeContext) BYTE() antlr.TerminalNode

func (s *PrimitiveTypeContext) CHAR() antlr.TerminalNode

func (s *PrimitiveTypeContext) FLOAT16() antlr.TerminalNode

func (s *PrimitiveTypeContext) FLOAT32() antlr.TerminalNode

func (s *PrimitiveTypeContext) FLOAT64() antlr.TerminalNode

func (s *PrimitiveTypeContext) GetParser() antlr.Parser

func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext

func (s *PrimitiveTypeContext) INT16() antlr.TerminalNode

func (s *PrimitiveTypeContext) INT32() antlr.TerminalNode

func (s *PrimitiveTypeContext) INT64() antlr.TerminalNode

func (s *PrimitiveTypeContext) INT8() antlr.TerminalNode

func (s *PrimitiveTypeContext) ISIZE() antlr.TerminalNode

func (*PrimitiveTypeContext) IsPrimitiveTypeContext()

func (s *PrimitiveTypeContext) STRING() antlr.TerminalNode

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *PrimitiveTypeContext) UINT16() antlr.TerminalNode

func (s *PrimitiveTypeContext) UINT32() antlr.TerminalNode

func (s *PrimitiveTypeContext) UINT64() antlr.TerminalNode

func (s *PrimitiveTypeContext) UINT8() antlr.TerminalNode

func (s *PrimitiveTypeContext) USIZE() antlr.TerminalNode

func (s *PrimitiveTypeContext) VOID() antlr.TerminalNode

type QualifiedIdentifierContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyQualifiedIdentifierContext() *QualifiedIdentifierContext

func NewQualifiedIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedIdentifierContext

func (s *QualifiedIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *QualifiedIdentifierContext) AllDOT() []antlr.TerminalNode

func (s *QualifiedIdentifierContext) AllIDENTIFIER() []antlr.TerminalNode

func (s *QualifiedIdentifierContext) DOT(i int) antlr.TerminalNode

func (s *QualifiedIdentifierContext) GetParser() antlr.Parser

func (s *QualifiedIdentifierContext) GetRuleContext() antlr.RuleContext

func (s *QualifiedIdentifierContext) IDENTIFIER(i int) antlr.TerminalNode

func (*QualifiedIdentifierContext) IsQualifiedIdentifierContext()

func (s *QualifiedIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type QualifiedTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyQualifiedTypeContext() *QualifiedTypeContext

func NewQualifiedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedTypeContext

func (s *QualifiedTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *QualifiedTypeContext) AllDOT() []antlr.TerminalNode

func (s *QualifiedTypeContext) AllIDENTIFIER() []antlr.TerminalNode

func (s *QualifiedTypeContext) DOT(i int) antlr.TerminalNode

func (s *QualifiedTypeContext) GenericArgs() IGenericArgsContext

func (s *QualifiedTypeContext) GetParser() antlr.Parser

func (s *QualifiedTypeContext) GetRuleContext() antlr.RuleContext

func (s *QualifiedTypeContext) IDENTIFIER(i int) antlr.TerminalNode

func (*QualifiedTypeContext) IsQualifiedTypeContext()

func (s *QualifiedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type RangeExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyRangeExpressionContext() *RangeExpressionContext

func NewRangeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeExpressionContext

func (s *RangeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *RangeExpressionContext) AdditiveExpression(i int) IAdditiveExpressionContext

func (s *RangeExpressionContext) AllAdditiveExpression() []IAdditiveExpressionContext

func (s *RangeExpressionContext) GetParser() antlr.Parser

func (s *RangeExpressionContext) GetRuleContext() antlr.RuleContext

func (*RangeExpressionContext) IsRangeExpressionContext()

func (s *RangeExpressionContext) RANGE() antlr.TerminalNode

func (s *RangeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ReferenceTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyReferenceTypeContext() *ReferenceTypeContext

func NewReferenceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceTypeContext

func (s *ReferenceTypeContext) AMP() antlr.TerminalNode

func (s *ReferenceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ReferenceTypeContext) CONST() antlr.TerminalNode

func (s *ReferenceTypeContext) GetParser() antlr.Parser

func (s *ReferenceTypeContext) GetRuleContext() antlr.RuleContext

func (*ReferenceTypeContext) IsReferenceTypeContext()

func (s *ReferenceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ReferenceTypeContext) Type_() ITypeContext

type RelationalExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *RelationalExpressionContext) AllGE() []antlr.TerminalNode

func (s *RelationalExpressionContext) AllGT() []antlr.TerminalNode

func (s *RelationalExpressionContext) AllLE() []antlr.TerminalNode

func (s *RelationalExpressionContext) AllLT() []antlr.TerminalNode

func (s *RelationalExpressionContext) AllShiftExpression() []IShiftExpressionContext

func (s *RelationalExpressionContext) GE(i int) antlr.TerminalNode

func (s *RelationalExpressionContext) GT(i int) antlr.TerminalNode

func (s *RelationalExpressionContext) GetParser() antlr.Parser

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext

func (*RelationalExpressionContext) IsRelationalExpressionContext()

func (s *RelationalExpressionContext) LE(i int) antlr.TerminalNode

func (s *RelationalExpressionContext) LT(i int) antlr.TerminalNode

func (s *RelationalExpressionContext) ShiftExpression(i int) IShiftExpressionContext

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyReturnStmtContext() *ReturnStmtContext

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ReturnStmtContext) Expression() IExpressionContext

func (s *ReturnStmtContext) GetParser() antlr.Parser

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext

func (*ReturnStmtContext) IsReturnStmtContext()

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ReturnStmtContext) TupleExpression() ITupleExpressionContext

type ReturnTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyReturnTypeContext() *ReturnTypeContext

func NewReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTypeContext

func (s *ReturnTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ReturnTypeContext) GetParser() antlr.Parser

func (s *ReturnTypeContext) GetRuleContext() antlr.RuleContext

func (*ReturnTypeContext) IsReturnTypeContext()

func (s *ReturnTypeContext) LPAREN() antlr.TerminalNode

func (s *ReturnTypeContext) RPAREN() antlr.TerminalNode

func (s *ReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ReturnTypeContext) TypeList() ITypeListContext

func (s *ReturnTypeContext) Type_() ITypeContext

type ShiftExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyShiftExpressionContext() *ShiftExpressionContext

func NewShiftExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftExpressionContext

func (s *ShiftExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ShiftExpressionContext) AllGT() []antlr.TerminalNode

func (s *ShiftExpressionContext) AllLT() []antlr.TerminalNode

func (s *ShiftExpressionContext) AllRangeExpression() []IRangeExpressionContext

func (s *ShiftExpressionContext) GT(i int) antlr.TerminalNode

func (s *ShiftExpressionContext) GetParser() antlr.Parser

func (s *ShiftExpressionContext) GetRuleContext() antlr.RuleContext

func (*ShiftExpressionContext) IsShiftExpressionContext()

func (s *ShiftExpressionContext) LT(i int) antlr.TerminalNode

func (s *ShiftExpressionContext) RangeExpression(i int) IRangeExpressionContext

func (s *ShiftExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type SizeofExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptySizeofExpressionContext() *SizeofExpressionContext

func NewSizeofExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SizeofExpressionContext

func (s *SizeofExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *SizeofExpressionContext) GT() antlr.TerminalNode

func (s *SizeofExpressionContext) GetParser() antlr.Parser

func (s *SizeofExpressionContext) GetRuleContext() antlr.RuleContext

func (*SizeofExpressionContext) IsSizeofExpressionContext()

func (s *SizeofExpressionContext) LT() antlr.TerminalNode

func (s *SizeofExpressionContext) SIZEOF() antlr.TerminalNode

func (s *SizeofExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *SizeofExpressionContext) Type_() ITypeContext

type StatementContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyStatementContext() *StatementContext

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StatementContext) AssignmentStmt() IAssignmentStmtContext

func (s *StatementContext) Block() IBlockContext

func (s *StatementContext) BreakStmt() IBreakStmtContext

func (s *StatementContext) ConstDecl() IConstDeclContext

func (s *StatementContext) ContinueStmt() IContinueStmtContext

func (s *StatementContext) DeferStmt() IDeferStmtContext

func (s *StatementContext) ExpressionStmt() IExpressionStmtContext

func (s *StatementContext) ForStmt() IForStmtContext

func (s *StatementContext) GetParser() antlr.Parser

func (s *StatementContext) GetRuleContext() antlr.RuleContext

func (s *StatementContext) IfStmt() IIfStmtContext

func (*StatementContext) IsStatementContext()

func (s *StatementContext) ReturnStmt() IReturnStmtContext

func (s *StatementContext) SwitchStmt() ISwitchStmtContext

func (s *StatementContext) ThrowStmt() IThrowStmtContext

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *StatementContext) TryStmt() ITryStmtContext

func (s *StatementContext) VariableDecl() IVariableDeclContext

type StructDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyStructDeclContext() *StructDeclContext

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StructDeclContext) AllStructMember() []IStructMemberContext

func (s *StructDeclContext) ComputeMarker() IComputeMarkerContext

func (s *StructDeclContext) GenericParams() IGenericParamsContext

func (s *StructDeclContext) GetParser() antlr.Parser

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext

func (s *StructDeclContext) IDENTIFIER() antlr.TerminalNode

func (*StructDeclContext) IsStructDeclContext()

func (s *StructDeclContext) LBRACE() antlr.TerminalNode

func (s *StructDeclContext) RBRACE() antlr.TerminalNode

func (s *StructDeclContext) STRUCT() antlr.TerminalNode

func (s *StructDeclContext) StructMember(i int) IStructMemberContext

func (s *StructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type StructFieldContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyStructFieldContext() *StructFieldContext

func NewStructFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldContext

func (s *StructFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StructFieldContext) COLON() antlr.TerminalNode

func (s *StructFieldContext) GetParser() antlr.Parser

func (s *StructFieldContext) GetRuleContext() antlr.RuleContext

func (s *StructFieldContext) IDENTIFIER() antlr.TerminalNode

func (*StructFieldContext) IsStructFieldContext()

func (s *StructFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *StructFieldContext) Type_() ITypeContext

type StructLiteralContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyStructLiteralContext() *StructLiteralContext

func NewStructLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructLiteralContext

func (s *StructLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StructLiteralContext) AllCOMMA() []antlr.TerminalNode

func (s *StructLiteralContext) AllFieldInit() []IFieldInitContext

func (s *StructLiteralContext) COMMA(i int) antlr.TerminalNode

func (s *StructLiteralContext) FieldInit(i int) IFieldInitContext

func (s *StructLiteralContext) GenericArgs() IGenericArgsContext

func (s *StructLiteralContext) GetParser() antlr.Parser

func (s *StructLiteralContext) GetRuleContext() antlr.RuleContext

func (s *StructLiteralContext) IDENTIFIER() antlr.TerminalNode

func (*StructLiteralContext) IsStructLiteralContext()

func (s *StructLiteralContext) LBRACE() antlr.TerminalNode

func (s *StructLiteralContext) QualifiedIdentifier() IQualifiedIdentifierContext

func (s *StructLiteralContext) RBRACE() antlr.TerminalNode

func (s *StructLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type StructMemberContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyStructMemberContext() *StructMemberContext

func NewStructMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructMemberContext

func (s *StructMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StructMemberContext) FunctionDecl() IFunctionDeclContext

func (s *StructMemberContext) GetParser() antlr.Parser

func (s *StructMemberContext) GetRuleContext() antlr.RuleContext

func (s *StructMemberContext) InitDecl() IInitDeclContext

func (*StructMemberContext) IsStructMemberContext()

func (s *StructMemberContext) MutatingDecl() IMutatingDeclContext

func (s *StructMemberContext) StructField() IStructFieldContext

func (s *StructMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type SwitchCaseContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptySwitchCaseContext() *SwitchCaseContext

func NewSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchCaseContext

func (s *SwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *SwitchCaseContext) AllCOMMA() []antlr.TerminalNode

func (s *SwitchCaseContext) AllExpression() []IExpressionContext

func (s *SwitchCaseContext) AllStatement() []IStatementContext

func (s *SwitchCaseContext) CASE() antlr.TerminalNode

func (s *SwitchCaseContext) COLON() antlr.TerminalNode

func (s *SwitchCaseContext) COMMA(i int) antlr.TerminalNode

func (s *SwitchCaseContext) Expression(i int) IExpressionContext

func (s *SwitchCaseContext) GetParser() antlr.Parser

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext

func (*SwitchCaseContext) IsSwitchCaseContext()

func (s *SwitchCaseContext) Statement(i int) IStatementContext

func (s *SwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type SwitchStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptySwitchStmtContext() *SwitchStmtContext

func NewSwitchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStmtContext

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *SwitchStmtContext) AllSwitchCase() []ISwitchCaseContext

func (s *SwitchStmtContext) DefaultCase() IDefaultCaseContext

func (s *SwitchStmtContext) Expression() IExpressionContext

func (s *SwitchStmtContext) GetParser() antlr.Parser

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext

func (*SwitchStmtContext) IsSwitchStmtContext()

func (s *SwitchStmtContext) LBRACE() antlr.TerminalNode

func (s *SwitchStmtContext) RBRACE() antlr.TerminalNode

func (s *SwitchStmtContext) SWITCH() antlr.TerminalNode

func (s *SwitchStmtContext) SwitchCase(i int) ISwitchCaseContext

func (s *SwitchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ThrowStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyThrowStmtContext() *ThrowStmtContext

func NewThrowStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThrowStmtContext

func (s *ThrowStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ThrowStmtContext) Expression() IExpressionContext

func (s *ThrowStmtContext) GetParser() antlr.Parser

func (s *ThrowStmtContext) GetRuleContext() antlr.RuleContext

func (*ThrowStmtContext) IsThrowStmtContext()

func (s *ThrowStmtContext) THROW() antlr.TerminalNode

func (s *ThrowStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type TopLevelDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyTopLevelDeclContext() *TopLevelDeclContext

func NewTopLevelDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDeclContext

func (s *TopLevelDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *TopLevelDeclContext) ClassDecl() IClassDeclContext

func (s *TopLevelDeclContext) ConstDecl() IConstDeclContext

func (s *TopLevelDeclContext) DeinitDecl() IDeinitDeclContext

func (s *TopLevelDeclContext) EnumDecl() IEnumDeclContext

func (s *TopLevelDeclContext) ExternCDecl() IExternCDeclContext

func (s *TopLevelDeclContext) ExternCppDecl() IExternCppDeclContext

func (s *TopLevelDeclContext) FunctionDecl() IFunctionDeclContext

func (s *TopLevelDeclContext) GetParser() antlr.Parser

func (s *TopLevelDeclContext) GetRuleContext() antlr.RuleContext

func (*TopLevelDeclContext) IsTopLevelDeclContext()

func (s *TopLevelDeclContext) MethodDecl() IMethodDeclContext

func (s *TopLevelDeclContext) MutatingDecl() IMutatingDeclContext

func (s *TopLevelDeclContext) StructDecl() IStructDeclContext

func (s *TopLevelDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *TopLevelDeclContext) VariableDecl() IVariableDeclContext

type TryStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyTryStmtContext() *TryStmtContext

func NewTryStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TryStmtContext

func (s *TryStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *TryStmtContext) AllExceptClause() []IExceptClauseContext

func (s *TryStmtContext) Block() IBlockContext

func (s *TryStmtContext) ExceptClause(i int) IExceptClauseContext

func (s *TryStmtContext) FinallyClause() IFinallyClauseContext

func (s *TryStmtContext) GetParser() antlr.Parser

func (s *TryStmtContext) GetRuleContext() antlr.RuleContext

func (*TryStmtContext) IsTryStmtContext()

func (s *TryStmtContext) TRY() antlr.TerminalNode

func (s *TryStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type TupleExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyTupleExpressionContext() *TupleExpressionContext

func NewTupleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TupleExpressionContext

func (s *TupleExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *TupleExpressionContext) AllCOMMA() []antlr.TerminalNode

func (s *TupleExpressionContext) AllExpression() []IExpressionContext

func (s *TupleExpressionContext) COMMA(i int) antlr.TerminalNode

func (s *TupleExpressionContext) Expression(i int) IExpressionContext

func (s *TupleExpressionContext) GetParser() antlr.Parser

func (s *TupleExpressionContext) GetRuleContext() antlr.RuleContext

func (*TupleExpressionContext) IsTupleExpressionContext()

func (s *TupleExpressionContext) LPAREN() antlr.TerminalNode

func (s *TupleExpressionContext) RPAREN() antlr.TerminalNode

func (s *TupleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type TuplePatternContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyTuplePatternContext() *TuplePatternContext

func NewTuplePatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TuplePatternContext

func (s *TuplePatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *TuplePatternContext) AllCOMMA() []antlr.TerminalNode

func (s *TuplePatternContext) AllIDENTIFIER() []antlr.TerminalNode

func (s *TuplePatternContext) COMMA(i int) antlr.TerminalNode

func (s *TuplePatternContext) GetParser() antlr.Parser

func (s *TuplePatternContext) GetRuleContext() antlr.RuleContext

func (s *TuplePatternContext) IDENTIFIER(i int) antlr.TerminalNode

func (*TuplePatternContext) IsTuplePatternContext()

func (s *TuplePatternContext) LPAREN() antlr.TerminalNode

func (s *TuplePatternContext) RPAREN() antlr.TerminalNode

func (s *TuplePatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type TupleTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyTupleTypeContext() *TupleTypeContext

func NewTupleTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TupleTypeContext

func (s *TupleTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *TupleTypeContext) GetParser() antlr.Parser

func (s *TupleTypeContext) GetRuleContext() antlr.RuleContext

func (*TupleTypeContext) IsTupleTypeContext()

func (s *TupleTypeContext) LPAREN() antlr.TerminalNode

func (s *TupleTypeContext) RPAREN() antlr.TerminalNode

func (s *TupleTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *TupleTypeContext) TypeList() ITypeListContext

type TypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyTypeContext() *TypeContext

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *TypeContext) ArrayType() IArrayTypeContext

func (s *TypeContext) FunctionType() IFunctionTypeContext

func (s *TypeContext) GenericArgs() IGenericArgsContext

func (s *TypeContext) GetParser() antlr.Parser

func (s *TypeContext) GetRuleContext() antlr.RuleContext

func (s *TypeContext) IDENTIFIER() antlr.TerminalNode

func (*TypeContext) IsTypeContext()

func (s *TypeContext) PointerType() IPointerTypeContext

func (s *TypeContext) PrimitiveType() IPrimitiveTypeContext

func (s *TypeContext) QualifiedType() IQualifiedTypeContext

func (s *TypeContext) ReferenceType() IReferenceTypeContext

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *TypeContext) UNDERSCORE() antlr.TerminalNode

type TypeListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyTypeListContext() *TypeListContext

func NewTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeListContext

func (s *TypeListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *TypeListContext) AllCOMMA() []antlr.TerminalNode

func (s *TypeListContext) AllType_() []ITypeContext

func (s *TypeListContext) COMMA(i int) antlr.TerminalNode

func (s *TypeListContext) GetParser() antlr.Parser

func (s *TypeListContext) GetRuleContext() antlr.RuleContext

func (*TypeListContext) IsTypeListContext()

func (s *TypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *TypeListContext) Type_(i int) ITypeContext

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext

func (s *UnaryExpressionContext) AMP() antlr.TerminalNode

func (s *UnaryExpressionContext) AWAIT() antlr.TerminalNode

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *UnaryExpressionContext) BIT_NOT() antlr.TerminalNode

func (s *UnaryExpressionContext) DECREMENT() antlr.TerminalNode

func (s *UnaryExpressionContext) Expression() IExpressionContext

func (s *UnaryExpressionContext) GetParser() antlr.Parser

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext

func (s *UnaryExpressionContext) INCREMENT() antlr.TerminalNode

func (*UnaryExpressionContext) IsUnaryExpressionContext()

func (s *UnaryExpressionContext) LPAREN() antlr.TerminalNode

func (s *UnaryExpressionContext) MINUS() antlr.TerminalNode

func (s *UnaryExpressionContext) NOT() antlr.TerminalNode

func (s *UnaryExpressionContext) PostfixExpression() IPostfixExpressionContext

func (s *UnaryExpressionContext) RPAREN() antlr.TerminalNode

func (s *UnaryExpressionContext) STAR() antlr.TerminalNode

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *UnaryExpressionContext) UnaryExpression() IUnaryExpressionContext

type VariableDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyVariableDeclContext() *VariableDeclContext

func NewVariableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclContext

func (s *VariableDeclContext) ASSIGN() antlr.TerminalNode

func (s *VariableDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *VariableDeclContext) COLON() antlr.TerminalNode

func (s *VariableDeclContext) Expression() IExpressionContext

func (s *VariableDeclContext) GetParser() antlr.Parser

func (s *VariableDeclContext) GetRuleContext() antlr.RuleContext

func (s *VariableDeclContext) IDENTIFIER() antlr.TerminalNode

func (*VariableDeclContext) IsVariableDeclContext()

func (s *VariableDeclContext) LET() antlr.TerminalNode

func (s *VariableDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *VariableDeclContext) TuplePattern() ITuplePatternContext

func (s *VariableDeclContext) TupleType() ITupleTypeContext

func (s *VariableDeclContext) Type_() ITypeContext

