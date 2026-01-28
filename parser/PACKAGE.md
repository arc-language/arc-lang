

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
	ArcLexerSTRUCT          = 9
	ArcLexerCLASS           = 10
	ArcLexerMUTATING        = 11
	ArcLexerINIT            = 12
	ArcLexerDEINIT          = 13
	ArcLexerRETURN          = 14
	ArcLexerIF              = 15
	ArcLexerELSE            = 16
	ArcLexerFOR             = 17
	ArcLexerIN              = 18
	ArcLexerBREAK           = 19
	ArcLexerCONTINUE        = 20
	ArcLexerDEFER           = 21
	ArcLexerSELF            = 22
	ArcLexerNULL            = 23
	ArcLexerSWITCH          = 24
	ArcLexerCASE            = 25
	ArcLexerDEFAULT         = 26
	ArcLexerENUM            = 27
	ArcLexerCOMPUTE         = 28
	ArcLexerPROPERTY        = 29
	ArcLexerPROTOCOL        = 30
	ArcLexerOPTIONAL        = 31
	ArcLexerEXTERN          = 32
	ArcLexerOPAQUE          = 33
	ArcLexerVIRTUAL         = 34
	ArcLexerSTATIC          = 35
	ArcLexerABSTRACT        = 36
	ArcLexerNEW             = 37
	ArcLexerDELETE          = 38
	ArcLexerTYPE            = 39
	ArcLexerSTDCALL         = 40
	ArcLexerCDECL           = 41
	ArcLexerFASTCALL        = 42
	ArcLexerVECTORCALL      = 43
	ArcLexerTHISCALL        = 44
	ArcLexerINT8            = 45
	ArcLexerINT16           = 46
	ArcLexerINT32           = 47
	ArcLexerINT64           = 48
	ArcLexerUINT8           = 49
	ArcLexerUINT16          = 50
	ArcLexerUINT32          = 51
	ArcLexerUINT64          = 52
	ArcLexerUSIZE           = 53
	ArcLexerISIZE           = 54
	ArcLexerFLOAT32         = 55
	ArcLexerFLOAT64         = 56
	ArcLexerBYTE            = 57
	ArcLexerBOOL            = 58
	ArcLexerCHAR            = 59
	ArcLexerSTRING          = 60
	ArcLexerVOID            = 61
	ArcLexerFLOAT16         = 62
	ArcLexerBFLOAT16        = 63
	ArcLexerSIZEOF          = 64
	ArcLexerALIGNOF         = 65
	ArcLexerARROW           = 66
	ArcLexerRANGE           = 67
	ArcLexerELLIPSIS        = 68
	ArcLexerEQ              = 69
	ArcLexerNE              = 70
	ArcLexerLE              = 71
	ArcLexerGE              = 72
	ArcLexerAND             = 73
	ArcLexerOR              = 74
	ArcLexerPLUS_ASSIGN     = 75
	ArcLexerMINUS_ASSIGN    = 76
	ArcLexerSTAR_ASSIGN     = 77
	ArcLexerSLASH_ASSIGN    = 78
	ArcLexerPERCENT_ASSIGN  = 79
	ArcLexerBIT_OR_ASSIGN   = 80
	ArcLexerBIT_AND_ASSIGN  = 81
	ArcLexerBIT_XOR_ASSIGN  = 82
	ArcLexerINCREMENT       = 83
	ArcLexerDECREMENT       = 84
	ArcLexerFAT_ARROW       = 85
	ArcLexerPLUS            = 86
	ArcLexerMINUS           = 87
	ArcLexerSTAR            = 88
	ArcLexerSLASH           = 89
	ArcLexerPERCENT         = 90
	ArcLexerLT              = 91
	ArcLexerGT              = 92
	ArcLexerNOT             = 93
	ArcLexerAMP             = 94
	ArcLexerBIT_OR          = 95
	ArcLexerBIT_XOR         = 96
	ArcLexerBIT_NOT         = 97
	ArcLexerAT              = 98
	ArcLexerASSIGN          = 99
	ArcLexerLPAREN          = 100
	ArcLexerRPAREN          = 101
	ArcLexerLBRACE          = 102
	ArcLexerRBRACE          = 103
	ArcLexerLBRACKET        = 104
	ArcLexerRBRACKET        = 105
	ArcLexerCOMMA           = 106
	ArcLexerCOLON           = 107
	ArcLexerSEMICOLON       = 108
	ArcLexerDOT             = 109
	ArcLexerUNDERSCORE      = 110
	ArcLexerBOOLEAN_LITERAL = 111
	ArcLexerINTEGER_LITERAL = 112
	ArcLexerFLOAT_LITERAL   = 113
	ArcLexerSTRING_LITERAL  = 114
	ArcLexerCHAR_LITERAL    = 115
	ArcLexerIDENTIFIER      = 116
	ArcLexerWS              = 117
	ArcLexerLINE_COMMENT    = 118
	ArcLexerBLOCK_COMMENT   = 119
	ArcLexerEXTERN_WS       = 120
	ArcLexerC_LANG          = 121
	ArcLexerCPP_LANG        = 122
	ArcLexerOBJC_LANG       = 123
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
	ArcParserSTRUCT          = 9
	ArcParserCLASS           = 10
	ArcParserMUTATING        = 11
	ArcParserINIT            = 12
	ArcParserDEINIT          = 13
	ArcParserRETURN          = 14
	ArcParserIF              = 15
	ArcParserELSE            = 16
	ArcParserFOR             = 17
	ArcParserIN              = 18
	ArcParserBREAK           = 19
	ArcParserCONTINUE        = 20
	ArcParserDEFER           = 21
	ArcParserSELF            = 22
	ArcParserNULL            = 23
	ArcParserSWITCH          = 24
	ArcParserCASE            = 25
	ArcParserDEFAULT         = 26
	ArcParserENUM            = 27
	ArcParserCOMPUTE         = 28
	ArcParserPROPERTY        = 29
	ArcParserPROTOCOL        = 30
	ArcParserOPTIONAL        = 31
	ArcParserEXTERN          = 32
	ArcParserOPAQUE          = 33
	ArcParserVIRTUAL         = 34
	ArcParserSTATIC          = 35
	ArcParserABSTRACT        = 36
	ArcParserNEW             = 37
	ArcParserDELETE          = 38
	ArcParserTYPE            = 39
	ArcParserSTDCALL         = 40
	ArcParserCDECL           = 41
	ArcParserFASTCALL        = 42
	ArcParserVECTORCALL      = 43
	ArcParserTHISCALL        = 44
	ArcParserINT8            = 45
	ArcParserINT16           = 46
	ArcParserINT32           = 47
	ArcParserINT64           = 48
	ArcParserUINT8           = 49
	ArcParserUINT16          = 50
	ArcParserUINT32          = 51
	ArcParserUINT64          = 52
	ArcParserUSIZE           = 53
	ArcParserISIZE           = 54
	ArcParserFLOAT32         = 55
	ArcParserFLOAT64         = 56
	ArcParserBYTE            = 57
	ArcParserBOOL            = 58
	ArcParserCHAR            = 59
	ArcParserSTRING          = 60
	ArcParserVOID            = 61
	ArcParserFLOAT16         = 62
	ArcParserBFLOAT16        = 63
	ArcParserSIZEOF          = 64
	ArcParserALIGNOF         = 65
	ArcParserARROW           = 66
	ArcParserRANGE           = 67
	ArcParserELLIPSIS        = 68
	ArcParserEQ              = 69
	ArcParserNE              = 70
	ArcParserLE              = 71
	ArcParserGE              = 72
	ArcParserAND             = 73
	ArcParserOR              = 74
	ArcParserPLUS_ASSIGN     = 75
	ArcParserMINUS_ASSIGN    = 76
	ArcParserSTAR_ASSIGN     = 77
	ArcParserSLASH_ASSIGN    = 78
	ArcParserPERCENT_ASSIGN  = 79
	ArcParserBIT_OR_ASSIGN   = 80
	ArcParserBIT_AND_ASSIGN  = 81
	ArcParserBIT_XOR_ASSIGN  = 82
	ArcParserINCREMENT       = 83
	ArcParserDECREMENT       = 84
	ArcParserFAT_ARROW       = 85
	ArcParserPLUS            = 86
	ArcParserMINUS           = 87
	ArcParserSTAR            = 88
	ArcParserSLASH           = 89
	ArcParserPERCENT         = 90
	ArcParserLT              = 91
	ArcParserGT              = 92
	ArcParserNOT             = 93
	ArcParserAMP             = 94
	ArcParserBIT_OR          = 95
	ArcParserBIT_XOR         = 96
	ArcParserBIT_NOT         = 97
	ArcParserAT              = 98
	ArcParserASSIGN          = 99
	ArcParserLPAREN          = 100
	ArcParserRPAREN          = 101
	ArcParserLBRACE          = 102
	ArcParserRBRACE          = 103
	ArcParserLBRACKET        = 104
	ArcParserRBRACKET        = 105
	ArcParserCOMMA           = 106
	ArcParserCOLON           = 107
	ArcParserSEMICOLON       = 108
	ArcParserDOT             = 109
	ArcParserUNDERSCORE      = 110
	ArcParserBOOLEAN_LITERAL = 111
	ArcParserINTEGER_LITERAL = 112
	ArcParserFLOAT_LITERAL   = 113
	ArcParserSTRING_LITERAL  = 114
	ArcParserCHAR_LITERAL    = 115
	ArcParserIDENTIFIER      = 116
	ArcParserWS              = 117
	ArcParserLINE_COMMENT    = 118
	ArcParserBLOCK_COMMENT   = 119
	ArcParserEXTERN_WS       = 120
	ArcParserC_LANG          = 121
	ArcParserCPP_LANG        = 122
	ArcParserOBJC_LANG       = 123
)
    ArcParser tokens.

const (
	ArcParserRULE_compilationUnit           = 0
	ArcParserRULE_importDecl                = 1
	ArcParserRULE_importSpec                = 2
	ArcParserRULE_namespaceDecl             = 3
	ArcParserRULE_topLevelDecl              = 4
	ArcParserRULE_attribute                 = 5
	ArcParserRULE_externCDecl               = 6
	ArcParserRULE_externCMember             = 7
	ArcParserRULE_externCFunctionDecl       = 8
	ArcParserRULE_cCallingConvention        = 9
	ArcParserRULE_externCParameterList      = 10
	ArcParserRULE_externCParameter          = 11
	ArcParserRULE_externCConstDecl          = 12
	ArcParserRULE_externCTypeAlias          = 13
	ArcParserRULE_externCOpaqueStructDecl   = 14
	ArcParserRULE_externCppDecl             = 15
	ArcParserRULE_externCppMember           = 16
	ArcParserRULE_externCppNamespaceDecl    = 17
	ArcParserRULE_externNamespacePath       = 18
	ArcParserRULE_externCppFunctionDecl     = 19
	ArcParserRULE_cppCallingConvention      = 20
	ArcParserRULE_externCppParameterList    = 21
	ArcParserRULE_externCppParameter        = 22
	ArcParserRULE_externCppParamType        = 23
	ArcParserRULE_externCppConstDecl        = 24
	ArcParserRULE_externCppTypeAlias        = 25
	ArcParserRULE_externCppOpaqueClassDecl  = 26
	ArcParserRULE_externCppClassDecl        = 27
	ArcParserRULE_externCppClassMember      = 28
	ArcParserRULE_externCppConstructorDecl  = 29
	ArcParserRULE_externCppDestructorDecl   = 30
	ArcParserRULE_externCppMethodDecl       = 31
	ArcParserRULE_externCppMethodParams     = 32
	ArcParserRULE_externCppSelfParam        = 33
	ArcParserRULE_externObjCDecl            = 34
	ArcParserRULE_externObjCMember          = 35
	ArcParserRULE_externObjCStructDecl      = 36
	ArcParserRULE_externObjCClassDecl       = 37
	ArcParserRULE_externObjCProtocolDecl    = 38
	ArcParserRULE_externObjCClassMember     = 39
	ArcParserRULE_externObjCProtocolMember  = 40
	ArcParserRULE_externObjCNewDecl         = 41
	ArcParserRULE_externObjCPropertyDecl    = 42
	ArcParserRULE_propertyAttributes        = 43
	ArcParserRULE_propertyAttribute         = 44
	ArcParserRULE_externObjCMethodDecl      = 45
	ArcParserRULE_externObjCMethodParams    = 46
	ArcParserRULE_externObjCSelfParam       = 47
	ArcParserRULE_externObjCOpaqueClassDecl = 48
	ArcParserRULE_externObjCFunctionDecl    = 49
	ArcParserRULE_genericParams             = 50
	ArcParserRULE_genericParamList          = 51
	ArcParserRULE_genericParam              = 52
	ArcParserRULE_genericArgs               = 53
	ArcParserRULE_genericArgList            = 54
	ArcParserRULE_genericArg                = 55
	ArcParserRULE_functionDecl              = 56
	ArcParserRULE_returnType                = 57
	ArcParserRULE_typeList                  = 58
	ArcParserRULE_parameterList             = 59
	ArcParserRULE_parameter                 = 60
	ArcParserRULE_structDecl                = 61
	ArcParserRULE_structMember              = 62
	ArcParserRULE_structField               = 63
	ArcParserRULE_initDecl                  = 64
	ArcParserRULE_computeDecl               = 65
	ArcParserRULE_classDecl                 = 66
	ArcParserRULE_classMember               = 67
	ArcParserRULE_classField                = 68
	ArcParserRULE_enumDecl                  = 69
	ArcParserRULE_enumMember                = 70
	ArcParserRULE_methodDecl                = 71
	ArcParserRULE_mutatingDecl              = 72
	ArcParserRULE_deinitDecl                = 73
	ArcParserRULE_variableDecl              = 74
	ArcParserRULE_constDecl                 = 75
	ArcParserRULE_tuplePattern              = 76
	ArcParserRULE_tupleType                 = 77
	ArcParserRULE_type                      = 78
	ArcParserRULE_qualifiedType             = 79
	ArcParserRULE_functionType              = 80
	ArcParserRULE_arrayType                 = 81
	ArcParserRULE_qualifiedIdentifier       = 82
	ArcParserRULE_primitiveType             = 83
	ArcParserRULE_pointerType               = 84
	ArcParserRULE_referenceType             = 85
	ArcParserRULE_block                     = 86
	ArcParserRULE_statement                 = 87
	ArcParserRULE_assignmentStmt            = 88
	ArcParserRULE_assignmentOp              = 89
	ArcParserRULE_expressionStmt            = 90
	ArcParserRULE_returnStmt                = 91
	ArcParserRULE_deferStmt                 = 92
	ArcParserRULE_breakStmt                 = 93
	ArcParserRULE_continueStmt              = 94
	ArcParserRULE_ifStmt                    = 95
	ArcParserRULE_forStmt                   = 96
	ArcParserRULE_switchStmt                = 97
	ArcParserRULE_switchCase                = 98
	ArcParserRULE_defaultCase               = 99
	ArcParserRULE_expression                = 100
	ArcParserRULE_logicalOrExpression       = 101
	ArcParserRULE_logicalAndExpression      = 102
	ArcParserRULE_bitOrExpression           = 103
	ArcParserRULE_bitXorExpression          = 104
	ArcParserRULE_bitAndExpression          = 105
	ArcParserRULE_equalityExpression        = 106
	ArcParserRULE_relationalExpression      = 107
	ArcParserRULE_shiftExpression           = 108
	ArcParserRULE_rangeExpression           = 109
	ArcParserRULE_additiveExpression        = 110
	ArcParserRULE_multiplicativeExpression  = 111
	ArcParserRULE_unaryExpression           = 112
	ArcParserRULE_postfixExpression         = 113
	ArcParserRULE_postfixOp                 = 114
	ArcParserRULE_primaryExpression         = 115
	ArcParserRULE_builtinExpression         = 116
	ArcParserRULE_sizeofExpression          = 117
	ArcParserRULE_alignofExpression         = 118
	ArcParserRULE_literal                   = 119
	ArcParserRULE_initializerList           = 120
	ArcParserRULE_initializerEntry          = 121
	ArcParserRULE_structLiteral             = 122
	ArcParserRULE_fieldInit                 = 123
	ArcParserRULE_argumentList              = 124
	ArcParserRULE_argument                  = 125
	ArcParserRULE_lambdaExpression          = 126
	ArcParserRULE_anonymousFuncExpression   = 127
	ArcParserRULE_executionStrategy         = 128
	ArcParserRULE_contextIdentifier         = 129
	ArcParserRULE_lambdaParamList           = 130
	ArcParserRULE_lambdaParam               = 131
	ArcParserRULE_tupleExpression           = 132
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
func InitEmptyAttributeContext(p *AttributeContext)
func InitEmptyBitAndExpressionContext(p *BitAndExpressionContext)
func InitEmptyBitOrExpressionContext(p *BitOrExpressionContext)
func InitEmptyBitXorExpressionContext(p *BitXorExpressionContext)
func InitEmptyBlockContext(p *BlockContext)
func InitEmptyBreakStmtContext(p *BreakStmtContext)
func InitEmptyBuiltinExpressionContext(p *BuiltinExpressionContext)
func InitEmptyCCallingConventionContext(p *CCallingConventionContext)
func InitEmptyClassDeclContext(p *ClassDeclContext)
func InitEmptyClassFieldContext(p *ClassFieldContext)
func InitEmptyClassMemberContext(p *ClassMemberContext)
func InitEmptyCompilationUnitContext(p *CompilationUnitContext)
func InitEmptyComputeDeclContext(p *ComputeDeclContext)
func InitEmptyConstDeclContext(p *ConstDeclContext)
func InitEmptyContextIdentifierContext(p *ContextIdentifierContext)
func InitEmptyContinueStmtContext(p *ContinueStmtContext)
func InitEmptyCppCallingConventionContext(p *CppCallingConventionContext)
func InitEmptyDefaultCaseContext(p *DefaultCaseContext)
func InitEmptyDeferStmtContext(p *DeferStmtContext)
func InitEmptyDeinitDeclContext(p *DeinitDeclContext)
func InitEmptyEnumDeclContext(p *EnumDeclContext)
func InitEmptyEnumMemberContext(p *EnumMemberContext)
func InitEmptyEqualityExpressionContext(p *EqualityExpressionContext)
func InitEmptyExecutionStrategyContext(p *ExecutionStrategyContext)
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
func InitEmptyExternObjCClassDeclContext(p *ExternObjCClassDeclContext)
func InitEmptyExternObjCClassMemberContext(p *ExternObjCClassMemberContext)
func InitEmptyExternObjCDeclContext(p *ExternObjCDeclContext)
func InitEmptyExternObjCFunctionDeclContext(p *ExternObjCFunctionDeclContext)
func InitEmptyExternObjCMemberContext(p *ExternObjCMemberContext)
func InitEmptyExternObjCMethodDeclContext(p *ExternObjCMethodDeclContext)
func InitEmptyExternObjCMethodParamsContext(p *ExternObjCMethodParamsContext)
func InitEmptyExternObjCNewDeclContext(p *ExternObjCNewDeclContext)
func InitEmptyExternObjCOpaqueClassDeclContext(p *ExternObjCOpaqueClassDeclContext)
func InitEmptyExternObjCPropertyDeclContext(p *ExternObjCPropertyDeclContext)
func InitEmptyExternObjCProtocolDeclContext(p *ExternObjCProtocolDeclContext)
func InitEmptyExternObjCProtocolMemberContext(p *ExternObjCProtocolMemberContext)
func InitEmptyExternObjCSelfParamContext(p *ExternObjCSelfParamContext)
func InitEmptyExternObjCStructDeclContext(p *ExternObjCStructDeclContext)
func InitEmptyFieldInitContext(p *FieldInitContext)
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
func InitEmptyPropertyAttributeContext(p *PropertyAttributeContext)
func InitEmptyPropertyAttributesContext(p *PropertyAttributesContext)
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
func InitEmptyTopLevelDeclContext(p *TopLevelDeclContext)
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

func (s *AlignofExpressionContext) GetParser() antlr.Parser

func (s *AlignofExpressionContext) GetRuleContext() antlr.RuleContext

func (*AlignofExpressionContext) IsAlignofExpressionContext()

func (s *AlignofExpressionContext) LPAREN() antlr.TerminalNode

func (s *AlignofExpressionContext) RPAREN() antlr.TerminalNode

func (s *AlignofExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *AlignofExpressionContext) Type_() ITypeContext

type AnonymousFuncExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewAnonymousFuncExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnonymousFuncExpressionContext

func NewEmptyAnonymousFuncExpressionContext() *AnonymousFuncExpressionContext

func (s *AnonymousFuncExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *AnonymousFuncExpressionContext) Block() IBlockContext

func (s *AnonymousFuncExpressionContext) ExecutionStrategy() IExecutionStrategyContext

func (s *AnonymousFuncExpressionContext) FUNC() antlr.TerminalNode

func (s *AnonymousFuncExpressionContext) GenericParams() IGenericParamsContext

func (s *AnonymousFuncExpressionContext) GetParser() antlr.Parser

func (s *AnonymousFuncExpressionContext) GetRuleContext() antlr.RuleContext

func (*AnonymousFuncExpressionContext) IsAnonymousFuncExpressionContext()

func (s *AnonymousFuncExpressionContext) LPAREN() antlr.TerminalNode

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

func (p *ArcParser) Attribute() (localctx IAttributeContext)

func (p *ArcParser) BitAndExpression() (localctx IBitAndExpressionContext)

func (p *ArcParser) BitOrExpression() (localctx IBitOrExpressionContext)

func (p *ArcParser) BitXorExpression() (localctx IBitXorExpressionContext)

func (p *ArcParser) Block() (localctx IBlockContext)

func (p *ArcParser) BreakStmt() (localctx IBreakStmtContext)

func (p *ArcParser) BuiltinExpression() (localctx IBuiltinExpressionContext)

func (p *ArcParser) CCallingConvention() (localctx ICCallingConventionContext)

func (p *ArcParser) ClassDecl() (localctx IClassDeclContext)

func (p *ArcParser) ClassField() (localctx IClassFieldContext)

func (p *ArcParser) ClassMember() (localctx IClassMemberContext)

func (p *ArcParser) CompilationUnit() (localctx ICompilationUnitContext)

func (p *ArcParser) ComputeDecl() (localctx IComputeDeclContext)

func (p *ArcParser) ConstDecl() (localctx IConstDeclContext)

func (p *ArcParser) ContextIdentifier() (localctx IContextIdentifierContext)

func (p *ArcParser) ContinueStmt() (localctx IContinueStmtContext)

func (p *ArcParser) CppCallingConvention() (localctx ICppCallingConventionContext)

func (p *ArcParser) DefaultCase() (localctx IDefaultCaseContext)

func (p *ArcParser) DeferStmt() (localctx IDeferStmtContext)

func (p *ArcParser) DeinitDecl() (localctx IDeinitDeclContext)

func (p *ArcParser) EnumDecl() (localctx IEnumDeclContext)

func (p *ArcParser) EnumMember() (localctx IEnumMemberContext)

func (p *ArcParser) EqualityExpression() (localctx IEqualityExpressionContext)

func (p *ArcParser) ExecutionStrategy() (localctx IExecutionStrategyContext)

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

func (p *ArcParser) ExternObjCClassDecl() (localctx IExternObjCClassDeclContext)

func (p *ArcParser) ExternObjCClassMember() (localctx IExternObjCClassMemberContext)

func (p *ArcParser) ExternObjCDecl() (localctx IExternObjCDeclContext)

func (p *ArcParser) ExternObjCFunctionDecl() (localctx IExternObjCFunctionDeclContext)

func (p *ArcParser) ExternObjCMember() (localctx IExternObjCMemberContext)

func (p *ArcParser) ExternObjCMethodDecl() (localctx IExternObjCMethodDeclContext)

func (p *ArcParser) ExternObjCMethodParams() (localctx IExternObjCMethodParamsContext)

func (p *ArcParser) ExternObjCNewDecl() (localctx IExternObjCNewDeclContext)

func (p *ArcParser) ExternObjCOpaqueClassDecl() (localctx IExternObjCOpaqueClassDeclContext)

func (p *ArcParser) ExternObjCPropertyDecl() (localctx IExternObjCPropertyDeclContext)

func (p *ArcParser) ExternObjCProtocolDecl() (localctx IExternObjCProtocolDeclContext)

func (p *ArcParser) ExternObjCProtocolMember() (localctx IExternObjCProtocolMemberContext)

func (p *ArcParser) ExternObjCSelfParam() (localctx IExternObjCSelfParamContext)

func (p *ArcParser) ExternObjCStructDecl() (localctx IExternObjCStructDeclContext)

func (p *ArcParser) FieldInit() (localctx IFieldInitContext)

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

func (p *ArcParser) PropertyAttribute() (localctx IPropertyAttributeContext)

func (p *ArcParser) PropertyAttributes() (localctx IPropertyAttributesContext)

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

func (p *ArcParser) TopLevelDecl() (localctx ITopLevelDeclContext)

func (p *ArcParser) TupleExpression() (localctx ITupleExpressionContext)

func (p *ArcParser) TuplePattern() (localctx ITuplePatternContext)

func (p *ArcParser) TupleType() (localctx ITupleTypeContext)

func (p *ArcParser) TypeList() (localctx ITypeListContext)

func (p *ArcParser) Type_() (localctx ITypeContext)

func (p *ArcParser) UnaryExpression() (localctx IUnaryExpressionContext)

func (p *ArcParser) VariableDecl() (localctx IVariableDeclContext)

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

	// Visit a parse tree produced by ArcParser#externObjCDecl.
	VisitExternObjCDecl(ctx *ExternObjCDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externObjCMember.
	VisitExternObjCMember(ctx *ExternObjCMemberContext) interface{}

	// Visit a parse tree produced by ArcParser#externObjCStructDecl.
	VisitExternObjCStructDecl(ctx *ExternObjCStructDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externObjCClassDecl.
	VisitExternObjCClassDecl(ctx *ExternObjCClassDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externObjCProtocolDecl.
	VisitExternObjCProtocolDecl(ctx *ExternObjCProtocolDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externObjCClassMember.
	VisitExternObjCClassMember(ctx *ExternObjCClassMemberContext) interface{}

	// Visit a parse tree produced by ArcParser#externObjCProtocolMember.
	VisitExternObjCProtocolMember(ctx *ExternObjCProtocolMemberContext) interface{}

	// Visit a parse tree produced by ArcParser#externObjCNewDecl.
	VisitExternObjCNewDecl(ctx *ExternObjCNewDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externObjCPropertyDecl.
	VisitExternObjCPropertyDecl(ctx *ExternObjCPropertyDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#propertyAttributes.
	VisitPropertyAttributes(ctx *PropertyAttributesContext) interface{}

	// Visit a parse tree produced by ArcParser#propertyAttribute.
	VisitPropertyAttribute(ctx *PropertyAttributeContext) interface{}

	// Visit a parse tree produced by ArcParser#externObjCMethodDecl.
	VisitExternObjCMethodDecl(ctx *ExternObjCMethodDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externObjCMethodParams.
	VisitExternObjCMethodParams(ctx *ExternObjCMethodParamsContext) interface{}

	// Visit a parse tree produced by ArcParser#externObjCSelfParam.
	VisitExternObjCSelfParam(ctx *ExternObjCSelfParamContext) interface{}

	// Visit a parse tree produced by ArcParser#externObjCOpaqueClassDecl.
	VisitExternObjCOpaqueClassDecl(ctx *ExternObjCOpaqueClassDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externObjCFunctionDecl.
	VisitExternObjCFunctionDecl(ctx *ExternObjCFunctionDeclContext) interface{}

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

	// Visit a parse tree produced by ArcParser#structMember.
	VisitStructMember(ctx *StructMemberContext) interface{}

	// Visit a parse tree produced by ArcParser#structField.
	VisitStructField(ctx *StructFieldContext) interface{}

	// Visit a parse tree produced by ArcParser#initDecl.
	VisitInitDecl(ctx *InitDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#computeDecl.
	VisitComputeDecl(ctx *ComputeDeclContext) interface{}

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

	// Visit a parse tree produced by ArcParser#builtinExpression.
	VisitBuiltinExpression(ctx *BuiltinExpressionContext) interface{}

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

	// Visit a parse tree produced by ArcParser#executionStrategy.
	VisitExecutionStrategy(ctx *ExecutionStrategyContext) interface{}

	// Visit a parse tree produced by ArcParser#contextIdentifier.
	VisitContextIdentifier(ctx *ContextIdentifierContext) interface{}

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

func (s *AssignmentStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *AssignmentStmtContext) UnaryExpression() IUnaryExpressionContext

type AttributeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext

func NewEmptyAttributeContext() *AttributeContext

func (s *AttributeContext) AT() antlr.TerminalNode

func (s *AttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *AttributeContext) Expression() IExpressionContext

func (s *AttributeContext) GetParser() antlr.Parser

func (s *AttributeContext) GetRuleContext() antlr.RuleContext

func (s *AttributeContext) IDENTIFIER() antlr.TerminalNode

func (*AttributeContext) IsAttributeContext()

func (s *AttributeContext) LPAREN() antlr.TerminalNode

func (s *AttributeContext) RPAREN() antlr.TerminalNode

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

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

func (v *BaseArcParserVisitor) VisitAttribute(ctx *AttributeContext) interface{}

func (v *BaseArcParserVisitor) VisitBitAndExpression(ctx *BitAndExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitBitOrExpression(ctx *BitOrExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitBitXorExpression(ctx *BitXorExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitBlock(ctx *BlockContext) interface{}

func (v *BaseArcParserVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitBuiltinExpression(ctx *BuiltinExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitCCallingConvention(ctx *CCallingConventionContext) interface{}

func (v *BaseArcParserVisitor) VisitClassDecl(ctx *ClassDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitClassField(ctx *ClassFieldContext) interface{}

func (v *BaseArcParserVisitor) VisitClassMember(ctx *ClassMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

func (v *BaseArcParserVisitor) VisitComputeDecl(ctx *ComputeDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitConstDecl(ctx *ConstDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitContextIdentifier(ctx *ContextIdentifierContext) interface{}

func (v *BaseArcParserVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitCppCallingConvention(ctx *CppCallingConventionContext) interface{}

func (v *BaseArcParserVisitor) VisitDefaultCase(ctx *DefaultCaseContext) interface{}

func (v *BaseArcParserVisitor) VisitDeferStmt(ctx *DeferStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitDeinitDecl(ctx *DeinitDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitEnumDecl(ctx *EnumDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitEnumMember(ctx *EnumMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitExecutionStrategy(ctx *ExecutionStrategyContext) interface{}

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

func (v *BaseArcParserVisitor) VisitExternObjCClassDecl(ctx *ExternObjCClassDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternObjCClassMember(ctx *ExternObjCClassMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitExternObjCDecl(ctx *ExternObjCDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternObjCFunctionDecl(ctx *ExternObjCFunctionDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternObjCMember(ctx *ExternObjCMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitExternObjCMethodDecl(ctx *ExternObjCMethodDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternObjCMethodParams(ctx *ExternObjCMethodParamsContext) interface{}

func (v *BaseArcParserVisitor) VisitExternObjCNewDecl(ctx *ExternObjCNewDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternObjCOpaqueClassDecl(ctx *ExternObjCOpaqueClassDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternObjCPropertyDecl(ctx *ExternObjCPropertyDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternObjCProtocolDecl(ctx *ExternObjCProtocolDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternObjCProtocolMember(ctx *ExternObjCProtocolMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitExternObjCSelfParam(ctx *ExternObjCSelfParamContext) interface{}

func (v *BaseArcParserVisitor) VisitExternObjCStructDecl(ctx *ExternObjCStructDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitFieldInit(ctx *FieldInitContext) interface{}

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

func (v *BaseArcParserVisitor) VisitPropertyAttribute(ctx *PropertyAttributeContext) interface{}

func (v *BaseArcParserVisitor) VisitPropertyAttributes(ctx *PropertyAttributesContext) interface{}

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

func (v *BaseArcParserVisitor) VisitTopLevelDecl(ctx *TopLevelDeclContext) interface{}

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

type BuiltinExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewBuiltinExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltinExpressionContext

func NewEmptyBuiltinExpressionContext() *BuiltinExpressionContext

func (s *BuiltinExpressionContext) AT() antlr.TerminalNode

func (s *BuiltinExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *BuiltinExpressionContext) ArgumentList() IArgumentListContext

func (s *BuiltinExpressionContext) GetParser() antlr.Parser

func (s *BuiltinExpressionContext) GetRuleContext() antlr.RuleContext

func (s *BuiltinExpressionContext) IDENTIFIER() antlr.TerminalNode

func (*BuiltinExpressionContext) IsBuiltinExpressionContext()

func (s *BuiltinExpressionContext) LPAREN() antlr.TerminalNode

func (s *BuiltinExpressionContext) RPAREN() antlr.TerminalNode

func (s *BuiltinExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

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

type ComputeDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewComputeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComputeDeclContext

func NewEmptyComputeDeclContext() *ComputeDeclContext

func (s *ComputeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ComputeDeclContext) AllStructMember() []IStructMemberContext

func (s *ComputeDeclContext) COMPUTE() antlr.TerminalNode

func (s *ComputeDeclContext) GetParser() antlr.Parser

func (s *ComputeDeclContext) GetRuleContext() antlr.RuleContext

func (s *ComputeDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ComputeDeclContext) IsComputeDeclContext()

func (s *ComputeDeclContext) LBRACE() antlr.TerminalNode

func (s *ComputeDeclContext) RBRACE() antlr.TerminalNode

func (s *ComputeDeclContext) StructMember(i int) IStructMemberContext

func (s *ComputeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

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

type ContextIdentifierContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewContextIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContextIdentifierContext

func NewEmptyContextIdentifierContext() *ContextIdentifierContext

func (s *ContextIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ContextIdentifierContext) AllDOT() []antlr.TerminalNode

func (s *ContextIdentifierContext) AllIDENTIFIER() []antlr.TerminalNode

func (s *ContextIdentifierContext) DOT(i int) antlr.TerminalNode

func (s *ContextIdentifierContext) GetParser() antlr.Parser

func (s *ContextIdentifierContext) GetRuleContext() antlr.RuleContext

func (s *ContextIdentifierContext) IDENTIFIER(i int) antlr.TerminalNode

func (*ContextIdentifierContext) IsContextIdentifierContext()

func (s *ContextIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

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

type ExecutionStrategyContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExecutionStrategyContext() *ExecutionStrategyContext

func NewExecutionStrategyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecutionStrategyContext

func (s *ExecutionStrategyContext) ASYNC() antlr.TerminalNode

func (s *ExecutionStrategyContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExecutionStrategyContext) COMPUTE() antlr.TerminalNode

func (s *ExecutionStrategyContext) ContextIdentifier() IContextIdentifierContext

func (s *ExecutionStrategyContext) GetParser() antlr.Parser

func (s *ExecutionStrategyContext) GetRuleContext() antlr.RuleContext

func (*ExecutionStrategyContext) IsExecutionStrategyContext()

func (s *ExecutionStrategyContext) PROCESS() antlr.TerminalNode

func (s *ExecutionStrategyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

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

type ExternObjCClassDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCClassDeclContext() *ExternObjCClassDeclContext

func NewExternObjCClassDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCClassDeclContext

func (s *ExternObjCClassDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCClassDeclContext) AllExternObjCClassMember() []IExternObjCClassMemberContext

func (s *ExternObjCClassDeclContext) CLASS() antlr.TerminalNode

func (s *ExternObjCClassDeclContext) COLON() antlr.TerminalNode

func (s *ExternObjCClassDeclContext) ExternObjCClassMember(i int) IExternObjCClassMemberContext

func (s *ExternObjCClassDeclContext) GetParser() antlr.Parser

func (s *ExternObjCClassDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternObjCClassDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternObjCClassDeclContext) IsExternObjCClassDeclContext()

func (s *ExternObjCClassDeclContext) LBRACE() antlr.TerminalNode

func (s *ExternObjCClassDeclContext) RBRACE() antlr.TerminalNode

func (s *ExternObjCClassDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternObjCClassDeclContext) TypeList() ITypeListContext

type ExternObjCClassMemberContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCClassMemberContext() *ExternObjCClassMemberContext

func NewExternObjCClassMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCClassMemberContext

func (s *ExternObjCClassMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCClassMemberContext) ExternObjCMethodDecl() IExternObjCMethodDeclContext

func (s *ExternObjCClassMemberContext) ExternObjCNewDecl() IExternObjCNewDeclContext

func (s *ExternObjCClassMemberContext) ExternObjCPropertyDecl() IExternObjCPropertyDeclContext

func (s *ExternObjCClassMemberContext) GetParser() antlr.Parser

func (s *ExternObjCClassMemberContext) GetRuleContext() antlr.RuleContext

func (*ExternObjCClassMemberContext) IsExternObjCClassMemberContext()

func (s *ExternObjCClassMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternObjCDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCDeclContext() *ExternObjCDeclContext

func NewExternObjCDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCDeclContext

func (s *ExternObjCDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCDeclContext) AllExternObjCMember() []IExternObjCMemberContext

func (s *ExternObjCDeclContext) EXTERN() antlr.TerminalNode

func (s *ExternObjCDeclContext) ExternObjCMember(i int) IExternObjCMemberContext

func (s *ExternObjCDeclContext) GetParser() antlr.Parser

func (s *ExternObjCDeclContext) GetRuleContext() antlr.RuleContext

func (*ExternObjCDeclContext) IsExternObjCDeclContext()

func (s *ExternObjCDeclContext) LBRACE() antlr.TerminalNode

func (s *ExternObjCDeclContext) OBJC_LANG() antlr.TerminalNode

func (s *ExternObjCDeclContext) RBRACE() antlr.TerminalNode

func (s *ExternObjCDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternObjCFunctionDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCFunctionDeclContext() *ExternObjCFunctionDeclContext

func NewExternObjCFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCFunctionDeclContext

func (s *ExternObjCFunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCFunctionDeclContext) ExternCParameterList() IExternCParameterListContext

func (s *ExternObjCFunctionDeclContext) FUNC() antlr.TerminalNode

func (s *ExternObjCFunctionDeclContext) GetParser() antlr.Parser

func (s *ExternObjCFunctionDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternObjCFunctionDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternObjCFunctionDeclContext) IsExternObjCFunctionDeclContext()

func (s *ExternObjCFunctionDeclContext) LPAREN() antlr.TerminalNode

func (s *ExternObjCFunctionDeclContext) RPAREN() antlr.TerminalNode

func (s *ExternObjCFunctionDeclContext) STRING_LITERAL() antlr.TerminalNode

func (s *ExternObjCFunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternObjCFunctionDeclContext) Type_() ITypeContext

type ExternObjCMemberContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCMemberContext() *ExternObjCMemberContext

func NewExternObjCMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCMemberContext

func (s *ExternObjCMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCMemberContext) ExternCConstDecl() IExternCConstDeclContext

func (s *ExternObjCMemberContext) ExternCOpaqueStructDecl() IExternCOpaqueStructDeclContext

func (s *ExternObjCMemberContext) ExternCTypeAlias() IExternCTypeAliasContext

func (s *ExternObjCMemberContext) ExternObjCClassDecl() IExternObjCClassDeclContext

func (s *ExternObjCMemberContext) ExternObjCFunctionDecl() IExternObjCFunctionDeclContext

func (s *ExternObjCMemberContext) ExternObjCOpaqueClassDecl() IExternObjCOpaqueClassDeclContext

func (s *ExternObjCMemberContext) ExternObjCProtocolDecl() IExternObjCProtocolDeclContext

func (s *ExternObjCMemberContext) ExternObjCStructDecl() IExternObjCStructDeclContext

func (s *ExternObjCMemberContext) GetParser() antlr.Parser

func (s *ExternObjCMemberContext) GetRuleContext() antlr.RuleContext

func (*ExternObjCMemberContext) IsExternObjCMemberContext()

func (s *ExternObjCMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternObjCMethodDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCMethodDeclContext() *ExternObjCMethodDeclContext

func NewExternObjCMethodDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCMethodDeclContext

func (s *ExternObjCMethodDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCMethodDeclContext) ExternObjCMethodParams() IExternObjCMethodParamsContext

func (s *ExternObjCMethodDeclContext) FUNC() antlr.TerminalNode

func (s *ExternObjCMethodDeclContext) GetParser() antlr.Parser

func (s *ExternObjCMethodDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternObjCMethodDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternObjCMethodDeclContext) IsExternObjCMethodDeclContext()

func (s *ExternObjCMethodDeclContext) LPAREN() antlr.TerminalNode

func (s *ExternObjCMethodDeclContext) RPAREN() antlr.TerminalNode

func (s *ExternObjCMethodDeclContext) STATIC() antlr.TerminalNode

func (s *ExternObjCMethodDeclContext) STRING_LITERAL() antlr.TerminalNode

func (s *ExternObjCMethodDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternObjCMethodDeclContext) Type_() ITypeContext

type ExternObjCMethodParamsContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCMethodParamsContext() *ExternObjCMethodParamsContext

func NewExternObjCMethodParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCMethodParamsContext

func (s *ExternObjCMethodParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCMethodParamsContext) AllCOMMA() []antlr.TerminalNode

func (s *ExternObjCMethodParamsContext) AllExternCParameter() []IExternCParameterContext

func (s *ExternObjCMethodParamsContext) COMMA(i int) antlr.TerminalNode

func (s *ExternObjCMethodParamsContext) ExternCParameter(i int) IExternCParameterContext

func (s *ExternObjCMethodParamsContext) ExternCParameterList() IExternCParameterListContext

func (s *ExternObjCMethodParamsContext) ExternObjCSelfParam() IExternObjCSelfParamContext

func (s *ExternObjCMethodParamsContext) GetParser() antlr.Parser

func (s *ExternObjCMethodParamsContext) GetRuleContext() antlr.RuleContext

func (*ExternObjCMethodParamsContext) IsExternObjCMethodParamsContext()

func (s *ExternObjCMethodParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternObjCNewDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCNewDeclContext() *ExternObjCNewDeclContext

func NewExternObjCNewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCNewDeclContext

func (s *ExternObjCNewDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCNewDeclContext) ExternCParameterList() IExternCParameterListContext

func (s *ExternObjCNewDeclContext) GetParser() antlr.Parser

func (s *ExternObjCNewDeclContext) GetRuleContext() antlr.RuleContext

func (*ExternObjCNewDeclContext) IsExternObjCNewDeclContext()

func (s *ExternObjCNewDeclContext) LPAREN() antlr.TerminalNode

func (s *ExternObjCNewDeclContext) NEW() antlr.TerminalNode

func (s *ExternObjCNewDeclContext) RPAREN() antlr.TerminalNode

func (s *ExternObjCNewDeclContext) STRING_LITERAL() antlr.TerminalNode

func (s *ExternObjCNewDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternObjCNewDeclContext) Type_() ITypeContext

type ExternObjCOpaqueClassDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCOpaqueClassDeclContext() *ExternObjCOpaqueClassDeclContext

func NewExternObjCOpaqueClassDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCOpaqueClassDeclContext

func (s *ExternObjCOpaqueClassDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCOpaqueClassDeclContext) CLASS() antlr.TerminalNode

func (s *ExternObjCOpaqueClassDeclContext) GetParser() antlr.Parser

func (s *ExternObjCOpaqueClassDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternObjCOpaqueClassDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternObjCOpaqueClassDeclContext) IsExternObjCOpaqueClassDeclContext()

func (s *ExternObjCOpaqueClassDeclContext) LBRACE() antlr.TerminalNode

func (s *ExternObjCOpaqueClassDeclContext) OPAQUE() antlr.TerminalNode

func (s *ExternObjCOpaqueClassDeclContext) RBRACE() antlr.TerminalNode

func (s *ExternObjCOpaqueClassDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternObjCPropertyDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCPropertyDeclContext() *ExternObjCPropertyDeclContext

func NewExternObjCPropertyDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCPropertyDeclContext

func (s *ExternObjCPropertyDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCPropertyDeclContext) COLON() antlr.TerminalNode

func (s *ExternObjCPropertyDeclContext) GetParser() antlr.Parser

func (s *ExternObjCPropertyDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternObjCPropertyDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternObjCPropertyDeclContext) IsExternObjCPropertyDeclContext()

func (s *ExternObjCPropertyDeclContext) LPAREN() antlr.TerminalNode

func (s *ExternObjCPropertyDeclContext) PROPERTY() antlr.TerminalNode

func (s *ExternObjCPropertyDeclContext) PropertyAttributes() IPropertyAttributesContext

func (s *ExternObjCPropertyDeclContext) RPAREN() antlr.TerminalNode

func (s *ExternObjCPropertyDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternObjCPropertyDeclContext) Type_() ITypeContext

type ExternObjCProtocolDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCProtocolDeclContext() *ExternObjCProtocolDeclContext

func NewExternObjCProtocolDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCProtocolDeclContext

func (s *ExternObjCProtocolDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCProtocolDeclContext) AllExternObjCProtocolMember() []IExternObjCProtocolMemberContext

func (s *ExternObjCProtocolDeclContext) COLON() antlr.TerminalNode

func (s *ExternObjCProtocolDeclContext) ExternObjCProtocolMember(i int) IExternObjCProtocolMemberContext

func (s *ExternObjCProtocolDeclContext) GetParser() antlr.Parser

func (s *ExternObjCProtocolDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternObjCProtocolDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternObjCProtocolDeclContext) IsExternObjCProtocolDeclContext()

func (s *ExternObjCProtocolDeclContext) LBRACE() antlr.TerminalNode

func (s *ExternObjCProtocolDeclContext) PROTOCOL() antlr.TerminalNode

func (s *ExternObjCProtocolDeclContext) RBRACE() antlr.TerminalNode

func (s *ExternObjCProtocolDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternObjCProtocolDeclContext) TypeList() ITypeListContext

type ExternObjCProtocolMemberContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCProtocolMemberContext() *ExternObjCProtocolMemberContext

func NewExternObjCProtocolMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCProtocolMemberContext

func (s *ExternObjCProtocolMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCProtocolMemberContext) ExternObjCMethodDecl() IExternObjCMethodDeclContext

func (s *ExternObjCProtocolMemberContext) ExternObjCPropertyDecl() IExternObjCPropertyDeclContext

func (s *ExternObjCProtocolMemberContext) GetParser() antlr.Parser

func (s *ExternObjCProtocolMemberContext) GetRuleContext() antlr.RuleContext

func (*ExternObjCProtocolMemberContext) IsExternObjCProtocolMemberContext()

func (s *ExternObjCProtocolMemberContext) OPTIONAL() antlr.TerminalNode

func (s *ExternObjCProtocolMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternObjCSelfParamContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCSelfParamContext() *ExternObjCSelfParamContext

func NewExternObjCSelfParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCSelfParamContext

func (s *ExternObjCSelfParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCSelfParamContext) COLON() antlr.TerminalNode

func (s *ExternObjCSelfParamContext) GetParser() antlr.Parser

func (s *ExternObjCSelfParamContext) GetRuleContext() antlr.RuleContext

func (s *ExternObjCSelfParamContext) IDENTIFIER() antlr.TerminalNode

func (*ExternObjCSelfParamContext) IsExternObjCSelfParamContext()

func (s *ExternObjCSelfParamContext) SELF() antlr.TerminalNode

func (s *ExternObjCSelfParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternObjCSelfParamContext) Type_() ITypeContext

type ExternObjCStructDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternObjCStructDeclContext() *ExternObjCStructDeclContext

func NewExternObjCStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternObjCStructDeclContext

func (s *ExternObjCStructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternObjCStructDeclContext) GetParser() antlr.Parser

func (s *ExternObjCStructDeclContext) GetRuleContext() antlr.RuleContext

func (*ExternObjCStructDeclContext) IsExternObjCStructDeclContext()

func (s *ExternObjCStructDeclContext) StructDecl() IStructDeclContext

func (s *ExternObjCStructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

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

func (s *FunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *FunctionDeclContext) Block() IBlockContext

func (s *FunctionDeclContext) ExecutionStrategy() IExecutionStrategyContext

func (s *FunctionDeclContext) FUNC() antlr.TerminalNode

func (s *FunctionDeclContext) GenericParams() IGenericParamsContext

func (s *FunctionDeclContext) GetParser() antlr.Parser

func (s *FunctionDeclContext) GetRuleContext() antlr.RuleContext

func (s *FunctionDeclContext) IDENTIFIER() antlr.TerminalNode

func (*FunctionDeclContext) IsFunctionDeclContext()

func (s *FunctionDeclContext) LPAREN() antlr.TerminalNode

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

func (s *FunctionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *FunctionTypeContext) ExecutionStrategy() IExecutionStrategyContext

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
	LPAREN() antlr.TerminalNode
	Type_() ITypeContext
	RPAREN() antlr.TerminalNode

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
	ExecutionStrategy() IExecutionStrategyContext
	GenericParams() IGenericParamsContext
	ParameterList() IParameterListContext
	ReturnType() IReturnTypeContext

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
	UnaryExpression() IUnaryExpressionContext
	AssignmentOp() IAssignmentOpContext
	Expression() IExpressionContext

	// IsAssignmentStmtContext differentiates from other interfaces.
	IsAssignmentStmtContext()
}
    IAssignmentStmtContext is an interface to support dynamic dispatch.

type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}
    IAttributeContext is an interface to support dynamic dispatch.

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

type IBuiltinExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgumentList() IArgumentListContext

	// IsBuiltinExpressionContext differentiates from other interfaces.
	IsBuiltinExpressionContext()
}
    IBuiltinExpressionContext is an interface to support dynamic dispatch.

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

type IComputeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMPUTE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStructMember() []IStructMemberContext
	StructMember(i int) IStructMemberContext

	// IsComputeDeclContext differentiates from other interfaces.
	IsComputeDeclContext()
}
    IComputeDeclContext is an interface to support dynamic dispatch.

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

type IContextIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsContextIdentifierContext differentiates from other interfaces.
	IsContextIdentifierContext()
}
    IContextIdentifierContext is an interface to support dynamic dispatch.

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

type IExecutionStrategyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ContextIdentifier() IContextIdentifierContext
	ASYNC() antlr.TerminalNode
	PROCESS() antlr.TerminalNode
	COMPUTE() antlr.TerminalNode

	// IsExecutionStrategyContext differentiates from other interfaces.
	IsExecutionStrategyContext()
}
    IExecutionStrategyContext is an interface to support dynamic dispatch.

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

type IExternObjCClassDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLASS() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	COLON() antlr.TerminalNode
	TypeList() ITypeListContext
	AllExternObjCClassMember() []IExternObjCClassMemberContext
	ExternObjCClassMember(i int) IExternObjCClassMemberContext

	// IsExternObjCClassDeclContext differentiates from other interfaces.
	IsExternObjCClassDeclContext()
}
    IExternObjCClassDeclContext is an interface to support dynamic dispatch.

type IExternObjCClassMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExternObjCMethodDecl() IExternObjCMethodDeclContext
	ExternObjCPropertyDecl() IExternObjCPropertyDeclContext
	ExternObjCNewDecl() IExternObjCNewDeclContext

	// IsExternObjCClassMemberContext differentiates from other interfaces.
	IsExternObjCClassMemberContext()
}
    IExternObjCClassMemberContext is an interface to support dynamic dispatch.

type IExternObjCDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTERN() antlr.TerminalNode
	OBJC_LANG() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllExternObjCMember() []IExternObjCMemberContext
	ExternObjCMember(i int) IExternObjCMemberContext

	// IsExternObjCDeclContext differentiates from other interfaces.
	IsExternObjCDeclContext()
}
    IExternObjCDeclContext is an interface to support dynamic dispatch.

type IExternObjCFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	ExternCParameterList() IExternCParameterListContext
	Type_() ITypeContext

	// IsExternObjCFunctionDeclContext differentiates from other interfaces.
	IsExternObjCFunctionDeclContext()
}
    IExternObjCFunctionDeclContext is an interface to support dynamic dispatch.

type IExternObjCMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExternObjCClassDecl() IExternObjCClassDeclContext
	ExternObjCProtocolDecl() IExternObjCProtocolDeclContext
	ExternObjCOpaqueClassDecl() IExternObjCOpaqueClassDeclContext
	ExternCOpaqueStructDecl() IExternCOpaqueStructDeclContext
	ExternObjCStructDecl() IExternObjCStructDeclContext
	ExternObjCFunctionDecl() IExternObjCFunctionDeclContext
	ExternCTypeAlias() IExternCTypeAliasContext
	ExternCConstDecl() IExternCConstDeclContext

	// IsExternObjCMemberContext differentiates from other interfaces.
	IsExternObjCMemberContext()
}
    IExternObjCMemberContext is an interface to support dynamic dispatch.

type IExternObjCMethodDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	STATIC() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	ExternObjCMethodParams() IExternObjCMethodParamsContext
	Type_() ITypeContext

	// IsExternObjCMethodDeclContext differentiates from other interfaces.
	IsExternObjCMethodDeclContext()
}
    IExternObjCMethodDeclContext is an interface to support dynamic dispatch.

type IExternObjCMethodParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExternObjCSelfParam() IExternObjCSelfParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllExternCParameter() []IExternCParameterContext
	ExternCParameter(i int) IExternCParameterContext
	ExternCParameterList() IExternCParameterListContext

	// IsExternObjCMethodParamsContext differentiates from other interfaces.
	IsExternObjCMethodParamsContext()
}
    IExternObjCMethodParamsContext is an interface to support dynamic dispatch.

type IExternObjCNewDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEW() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Type_() ITypeContext
	STRING_LITERAL() antlr.TerminalNode
	ExternCParameterList() IExternCParameterListContext

	// IsExternObjCNewDeclContext differentiates from other interfaces.
	IsExternObjCNewDeclContext()
}
    IExternObjCNewDeclContext is an interface to support dynamic dispatch.

type IExternObjCOpaqueClassDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPAQUE() antlr.TerminalNode
	CLASS() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode

	// IsExternObjCOpaqueClassDeclContext differentiates from other interfaces.
	IsExternObjCOpaqueClassDeclContext()
}
    IExternObjCOpaqueClassDeclContext is an interface to support dynamic
    dispatch.

type IExternObjCPropertyDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROPERTY() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	LPAREN() antlr.TerminalNode
	PropertyAttributes() IPropertyAttributesContext
	RPAREN() antlr.TerminalNode

	// IsExternObjCPropertyDeclContext differentiates from other interfaces.
	IsExternObjCPropertyDeclContext()
}
    IExternObjCPropertyDeclContext is an interface to support dynamic dispatch.

type IExternObjCProtocolDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROTOCOL() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	COLON() antlr.TerminalNode
	TypeList() ITypeListContext
	AllExternObjCProtocolMember() []IExternObjCProtocolMemberContext
	ExternObjCProtocolMember(i int) IExternObjCProtocolMemberContext

	// IsExternObjCProtocolDeclContext differentiates from other interfaces.
	IsExternObjCProtocolDeclContext()
}
    IExternObjCProtocolDeclContext is an interface to support dynamic dispatch.

type IExternObjCProtocolMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExternObjCMethodDecl() IExternObjCMethodDeclContext
	OPTIONAL() antlr.TerminalNode
	ExternObjCPropertyDecl() IExternObjCPropertyDeclContext

	// IsExternObjCProtocolMemberContext differentiates from other interfaces.
	IsExternObjCProtocolMemberContext()
}
    IExternObjCProtocolMemberContext is an interface to support dynamic
    dispatch.

type IExternObjCSelfParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELF() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsExternObjCSelfParamContext differentiates from other interfaces.
	IsExternObjCSelfParamContext()
}
    IExternObjCSelfParamContext is an interface to support dynamic dispatch.

type IExternObjCStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StructDecl() IStructDeclContext

	// IsExternObjCStructDeclContext differentiates from other interfaces.
	IsExternObjCStructDeclContext()
}
    IExternObjCStructDeclContext is an interface to support dynamic dispatch.

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
	ExecutionStrategy() IExecutionStrategyContext
	GenericParams() IGenericParamsContext
	ParameterList() IParameterListContext
	ReturnType() IReturnTypeContext

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
	ExecutionStrategy() IExecutionStrategyContext
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
	IDENTIFIER() antlr.TerminalNode
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
	ExecutionStrategy() IExecutionStrategyContext
	LambdaParamList() ILambdaParamListContext
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
	ExecutionStrategy() IExecutionStrategyContext
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
	BuiltinExpression() IBuiltinExpressionContext
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

type IPropertyAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	COLON() antlr.TerminalNode

	// IsPropertyAttributeContext differentiates from other interfaces.
	IsPropertyAttributeContext()
}
    IPropertyAttributeContext is an interface to support dynamic dispatch.

type IPropertyAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPropertyAttribute() []IPropertyAttributeContext
	PropertyAttribute(i int) IPropertyAttributeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPropertyAttributesContext differentiates from other interfaces.
	IsPropertyAttributesContext()
}
    IPropertyAttributesContext is an interface to support dynamic dispatch.

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
	LPAREN() antlr.TerminalNode
	Type_() ITypeContext
	RPAREN() antlr.TerminalNode

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
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext
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

type ITopLevelDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionDecl() IFunctionDeclContext
	StructDecl() IStructDeclContext
	ClassDecl() IClassDeclContext
	EnumDecl() IEnumDeclContext
	ComputeDecl() IComputeDeclContext
	MethodDecl() IMethodDeclContext
	MutatingDecl() IMutatingDeclContext
	DeinitDecl() IDeinitDeclContext
	VariableDecl() IVariableDeclContext
	ConstDecl() IConstDeclContext
	ExternCDecl() IExternCDeclContext
	ExternCppDecl() IExternCppDeclContext
	ExternObjCDecl() IExternObjCDeclContext

	// IsTopLevelDeclContext differentiates from other interfaces.
	IsTopLevelDeclContext()
}
    ITopLevelDeclContext is an interface to support dynamic dispatch.

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

func (s *ImportDeclContext) IDENTIFIER() antlr.TerminalNode

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

func (s *LambdaExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *LambdaExpressionContext) Block() IBlockContext

func (s *LambdaExpressionContext) ExecutionStrategy() IExecutionStrategyContext

func (s *LambdaExpressionContext) Expression() IExpressionContext

func (s *LambdaExpressionContext) FAT_ARROW() antlr.TerminalNode

func (s *LambdaExpressionContext) GetParser() antlr.Parser

func (s *LambdaExpressionContext) GetRuleContext() antlr.RuleContext

func (*LambdaExpressionContext) IsLambdaExpressionContext()

func (s *LambdaExpressionContext) LPAREN() antlr.TerminalNode

func (s *LambdaExpressionContext) LambdaParamList() ILambdaParamListContext

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

func (s *MethodDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *MethodDeclContext) AllCOMMA() []antlr.TerminalNode

func (s *MethodDeclContext) AllIDENTIFIER() []antlr.TerminalNode

func (s *MethodDeclContext) AllParameter() []IParameterContext

func (s *MethodDeclContext) Block() IBlockContext

func (s *MethodDeclContext) COLON() antlr.TerminalNode

func (s *MethodDeclContext) COMMA(i int) antlr.TerminalNode

func (s *MethodDeclContext) ExecutionStrategy() IExecutionStrategyContext

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

func (s *PrimaryExpressionContext) BuiltinExpression() IBuiltinExpressionContext

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

type PropertyAttributeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyPropertyAttributeContext() *PropertyAttributeContext

func NewPropertyAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyAttributeContext

func (s *PropertyAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *PropertyAttributeContext) AllIDENTIFIER() []antlr.TerminalNode

func (s *PropertyAttributeContext) COLON() antlr.TerminalNode

func (s *PropertyAttributeContext) GetParser() antlr.Parser

func (s *PropertyAttributeContext) GetRuleContext() antlr.RuleContext

func (s *PropertyAttributeContext) IDENTIFIER(i int) antlr.TerminalNode

func (*PropertyAttributeContext) IsPropertyAttributeContext()

func (s *PropertyAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type PropertyAttributesContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyPropertyAttributesContext() *PropertyAttributesContext

func NewPropertyAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyAttributesContext

func (s *PropertyAttributesContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *PropertyAttributesContext) AllCOMMA() []antlr.TerminalNode

func (s *PropertyAttributesContext) AllPropertyAttribute() []IPropertyAttributeContext

func (s *PropertyAttributesContext) COMMA(i int) antlr.TerminalNode

func (s *PropertyAttributesContext) GetParser() antlr.Parser

func (s *PropertyAttributesContext) GetRuleContext() antlr.RuleContext

func (*PropertyAttributesContext) IsPropertyAttributesContext()

func (s *PropertyAttributesContext) PropertyAttribute(i int) IPropertyAttributeContext

func (s *PropertyAttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

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

func (s *SizeofExpressionContext) GetParser() antlr.Parser

func (s *SizeofExpressionContext) GetRuleContext() antlr.RuleContext

func (*SizeofExpressionContext) IsSizeofExpressionContext()

func (s *SizeofExpressionContext) LPAREN() antlr.TerminalNode

func (s *SizeofExpressionContext) RPAREN() antlr.TerminalNode

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

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *StatementContext) VariableDecl() IVariableDeclContext

type StructDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyStructDeclContext() *StructDeclContext

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StructDeclContext) AllAttribute() []IAttributeContext

func (s *StructDeclContext) AllStructMember() []IStructMemberContext

func (s *StructDeclContext) Attribute(i int) IAttributeContext

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

type TopLevelDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyTopLevelDeclContext() *TopLevelDeclContext

func NewTopLevelDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDeclContext

func (s *TopLevelDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *TopLevelDeclContext) ClassDecl() IClassDeclContext

func (s *TopLevelDeclContext) ComputeDecl() IComputeDeclContext

func (s *TopLevelDeclContext) ConstDecl() IConstDeclContext

func (s *TopLevelDeclContext) DeinitDecl() IDeinitDeclContext

func (s *TopLevelDeclContext) EnumDecl() IEnumDeclContext

func (s *TopLevelDeclContext) ExternCDecl() IExternCDeclContext

func (s *TopLevelDeclContext) ExternCppDecl() IExternCppDeclContext

func (s *TopLevelDeclContext) ExternObjCDecl() IExternObjCDeclContext

func (s *TopLevelDeclContext) FunctionDecl() IFunctionDeclContext

func (s *TopLevelDeclContext) GetParser() antlr.Parser

func (s *TopLevelDeclContext) GetRuleContext() antlr.RuleContext

func (*TopLevelDeclContext) IsTopLevelDeclContext()

func (s *TopLevelDeclContext) MethodDecl() IMethodDeclContext

func (s *TopLevelDeclContext) MutatingDecl() IMutatingDeclContext

func (s *TopLevelDeclContext) StructDecl() IStructDeclContext

func (s *TopLevelDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *TopLevelDeclContext) VariableDecl() IVariableDeclContext

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

