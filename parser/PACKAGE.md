

CONSTANTS

const (
	ArcLexerIMPORT          = 1
	ArcLexerNAMESPACE       = 2
	ArcLexerLET             = 3
	ArcLexerCONST           = 4
	ArcLexerFUNC            = 5
	ArcLexerASYNC           = 6
	ArcLexerAWAIT           = 7
	ArcLexerSTRUCT          = 8
	ArcLexerCLASS           = 9
	ArcLexerMUTATING        = 10
	ArcLexerDEINIT          = 11
	ArcLexerRETURN          = 12
	ArcLexerIF              = 13
	ArcLexerELSE            = 14
	ArcLexerFOR             = 15
	ArcLexerIN              = 16
	ArcLexerBREAK           = 17
	ArcLexerCONTINUE        = 18
	ArcLexerDEFER           = 19
	ArcLexerEXTERN          = 20
	ArcLexerSELF            = 21
	ArcLexerNULL            = 22
	ArcLexerINT8            = 23
	ArcLexerINT16           = 24
	ArcLexerINT32           = 25
	ArcLexerINT64           = 26
	ArcLexerUINT8           = 27
	ArcLexerUINT16          = 28
	ArcLexerUINT32          = 29
	ArcLexerUINT64          = 30
	ArcLexerUSIZE           = 31
	ArcLexerISIZE           = 32
	ArcLexerFLOAT32         = 33
	ArcLexerFLOAT64         = 34
	ArcLexerBYTE            = 35
	ArcLexerBOOL            = 36
	ArcLexerCHAR            = 37
	ArcLexerSTRING          = 38
	ArcLexerVOID            = 39
	ArcLexerVECTOR          = 40
	ArcLexerMAP             = 41
	ArcLexerALLOCA          = 42
	ArcLexerCAST            = 43
	ArcLexerSYSCALL         = 44
	ArcLexerSIZEOF          = 45
	ArcLexerALIGNOF         = 46
	ArcLexerMEMSET          = 47
	ArcLexerMEMCPY          = 48
	ArcLexerMEMMOVE         = 49
	ArcLexerSTRLEN          = 50
	ArcLexerMEMCHR          = 51
	ArcLexerVA_START        = 52
	ArcLexerVA_ARG          = 53
	ArcLexerVA_END          = 54
	ArcLexerRAISE           = 55
	ArcLexerMEMCMP          = 56
	ArcLexerBIT_CAST        = 57
	ArcLexerARROW           = 58
	ArcLexerRANGE           = 59
	ArcLexerELLIPSIS        = 60
	ArcLexerEQ              = 61
	ArcLexerNE              = 62
	ArcLexerLE              = 63
	ArcLexerGE              = 64
	ArcLexerAND             = 65
	ArcLexerOR              = 66
	ArcLexerPLUS_ASSIGN     = 67
	ArcLexerMINUS_ASSIGN    = 68
	ArcLexerSTAR_ASSIGN     = 69
	ArcLexerSLASH_ASSIGN    = 70
	ArcLexerPERCENT_ASSIGN  = 71
	ArcLexerINCREMENT       = 72
	ArcLexerDECREMENT       = 73
	ArcLexerPLUS            = 74
	ArcLexerMINUS           = 75
	ArcLexerSTAR            = 76
	ArcLexerSLASH           = 77
	ArcLexerPERCENT         = 78
	ArcLexerLT              = 79
	ArcLexerGT              = 80
	ArcLexerNOT             = 81
	ArcLexerAMP             = 82
	ArcLexerAT              = 83
	ArcLexerASSIGN          = 84
	ArcLexerLPAREN          = 85
	ArcLexerRPAREN          = 86
	ArcLexerLBRACE          = 87
	ArcLexerRBRACE          = 88
	ArcLexerLBRACKET        = 89
	ArcLexerRBRACKET        = 90
	ArcLexerCOMMA           = 91
	ArcLexerCOLON           = 92
	ArcLexerSEMICOLON       = 93
	ArcLexerDOT             = 94
	ArcLexerBOOLEAN_LITERAL = 95
	ArcLexerINTEGER_LITERAL = 96
	ArcLexerFLOAT_LITERAL   = 97
	ArcLexerSTRING_LITERAL  = 98
	ArcLexerCHAR_LITERAL    = 99
	ArcLexerIDENTIFIER      = 100
	ArcLexerWS              = 101
	ArcLexerLINE_COMMENT    = 102
	ArcLexerBLOCK_COMMENT   = 103
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
	ArcParserSTRUCT          = 8
	ArcParserCLASS           = 9
	ArcParserMUTATING        = 10
	ArcParserDEINIT          = 11
	ArcParserRETURN          = 12
	ArcParserIF              = 13
	ArcParserELSE            = 14
	ArcParserFOR             = 15
	ArcParserIN              = 16
	ArcParserBREAK           = 17
	ArcParserCONTINUE        = 18
	ArcParserDEFER           = 19
	ArcParserEXTERN          = 20
	ArcParserSELF            = 21
	ArcParserNULL            = 22
	ArcParserINT8            = 23
	ArcParserINT16           = 24
	ArcParserINT32           = 25
	ArcParserINT64           = 26
	ArcParserUINT8           = 27
	ArcParserUINT16          = 28
	ArcParserUINT32          = 29
	ArcParserUINT64          = 30
	ArcParserUSIZE           = 31
	ArcParserISIZE           = 32
	ArcParserFLOAT32         = 33
	ArcParserFLOAT64         = 34
	ArcParserBYTE            = 35
	ArcParserBOOL            = 36
	ArcParserCHAR            = 37
	ArcParserSTRING          = 38
	ArcParserVOID            = 39
	ArcParserVECTOR          = 40
	ArcParserMAP             = 41
	ArcParserALLOCA          = 42
	ArcParserCAST            = 43
	ArcParserSYSCALL         = 44
	ArcParserSIZEOF          = 45
	ArcParserALIGNOF         = 46
	ArcParserMEMSET          = 47
	ArcParserMEMCPY          = 48
	ArcParserMEMMOVE         = 49
	ArcParserSTRLEN          = 50
	ArcParserMEMCHR          = 51
	ArcParserVA_START        = 52
	ArcParserVA_ARG          = 53
	ArcParserVA_END          = 54
	ArcParserRAISE           = 55
	ArcParserMEMCMP          = 56
	ArcParserBIT_CAST        = 57
	ArcParserARROW           = 58
	ArcParserRANGE           = 59
	ArcParserELLIPSIS        = 60
	ArcParserEQ              = 61
	ArcParserNE              = 62
	ArcParserLE              = 63
	ArcParserGE              = 64
	ArcParserAND             = 65
	ArcParserOR              = 66
	ArcParserPLUS_ASSIGN     = 67
	ArcParserMINUS_ASSIGN    = 68
	ArcParserSTAR_ASSIGN     = 69
	ArcParserSLASH_ASSIGN    = 70
	ArcParserPERCENT_ASSIGN  = 71
	ArcParserINCREMENT       = 72
	ArcParserDECREMENT       = 73
	ArcParserPLUS            = 74
	ArcParserMINUS           = 75
	ArcParserSTAR            = 76
	ArcParserSLASH           = 77
	ArcParserPERCENT         = 78
	ArcParserLT              = 79
	ArcParserGT              = 80
	ArcParserNOT             = 81
	ArcParserAMP             = 82
	ArcParserAT              = 83
	ArcParserASSIGN          = 84
	ArcParserLPAREN          = 85
	ArcParserRPAREN          = 86
	ArcParserLBRACE          = 87
	ArcParserRBRACE          = 88
	ArcParserLBRACKET        = 89
	ArcParserRBRACKET        = 90
	ArcParserCOMMA           = 91
	ArcParserCOLON           = 92
	ArcParserSEMICOLON       = 93
	ArcParserDOT             = 94
	ArcParserBOOLEAN_LITERAL = 95
	ArcParserINTEGER_LITERAL = 96
	ArcParserFLOAT_LITERAL   = 97
	ArcParserSTRING_LITERAL  = 98
	ArcParserCHAR_LITERAL    = 99
	ArcParserIDENTIFIER      = 100
	ArcParserWS              = 101
	ArcParserLINE_COMMENT    = 102
	ArcParserBLOCK_COMMENT   = 103
)
    ArcParser tokens.

const (
	ArcParserRULE_compilationUnit          = 0
	ArcParserRULE_importDecl               = 1
	ArcParserRULE_importSpec               = 2
	ArcParserRULE_namespaceDecl            = 3
	ArcParserRULE_topLevelDecl             = 4
	ArcParserRULE_externDecl               = 5
	ArcParserRULE_externMember             = 6
	ArcParserRULE_externFunctionDecl       = 7
	ArcParserRULE_externParameterList      = 8
	ArcParserRULE_functionDecl             = 9
	ArcParserRULE_parameterList            = 10
	ArcParserRULE_parameter                = 11
	ArcParserRULE_structDecl               = 12
	ArcParserRULE_structMember             = 13
	ArcParserRULE_structField              = 14
	ArcParserRULE_classDecl                = 15
	ArcParserRULE_classMember              = 16
	ArcParserRULE_classField               = 17
	ArcParserRULE_methodDecl               = 18
	ArcParserRULE_mutatingDecl             = 19
	ArcParserRULE_deinitDecl               = 20
	ArcParserRULE_variableDecl             = 21
	ArcParserRULE_constDecl                = 22
	ArcParserRULE_type                     = 23
	ArcParserRULE_primitiveType            = 24
	ArcParserRULE_pointerType              = 25
	ArcParserRULE_referenceType            = 26
	ArcParserRULE_vectorType               = 27
	ArcParserRULE_mapType                  = 28
	ArcParserRULE_block                    = 29
	ArcParserRULE_statement                = 30
	ArcParserRULE_assignmentStmt           = 31
	ArcParserRULE_leftHandSide             = 32
	ArcParserRULE_expressionStmt           = 33
	ArcParserRULE_returnStmt               = 34
	ArcParserRULE_ifStmt                   = 35
	ArcParserRULE_forStmt                  = 36
	ArcParserRULE_breakStmt                = 37
	ArcParserRULE_continueStmt             = 38
	ArcParserRULE_deferStmt                = 39
	ArcParserRULE_expression               = 40
	ArcParserRULE_logicalOrExpression      = 41
	ArcParserRULE_logicalAndExpression     = 42
	ArcParserRULE_equalityExpression       = 43
	ArcParserRULE_relationalExpression     = 44
	ArcParserRULE_rangeExpression          = 45
	ArcParserRULE_additiveExpression       = 46
	ArcParserRULE_multiplicativeExpression = 47
	ArcParserRULE_unaryExpression          = 48
	ArcParserRULE_postfixExpression        = 49
	ArcParserRULE_postfixOp                = 50
	ArcParserRULE_primaryExpression        = 51
	ArcParserRULE_literal                  = 52
	ArcParserRULE_vectorLiteral            = 53
	ArcParserRULE_mapLiteral               = 54
	ArcParserRULE_mapEntry                 = 55
	ArcParserRULE_structLiteral            = 56
	ArcParserRULE_fieldInit                = 57
	ArcParserRULE_argumentList             = 58
	ArcParserRULE_castExpression           = 59
	ArcParserRULE_allocaExpression         = 60
	ArcParserRULE_syscallExpression        = 61
	ArcParserRULE_intrinsicExpression      = 62
)
    ArcParser rules.


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
func InitEmptyAllocaExpressionContext(p *AllocaExpressionContext)
func InitEmptyArgumentListContext(p *ArgumentListContext)
func InitEmptyAssignmentStmtContext(p *AssignmentStmtContext)
func InitEmptyBlockContext(p *BlockContext)
func InitEmptyBreakStmtContext(p *BreakStmtContext)
func InitEmptyCastExpressionContext(p *CastExpressionContext)
func InitEmptyClassDeclContext(p *ClassDeclContext)
func InitEmptyClassFieldContext(p *ClassFieldContext)
func InitEmptyClassMemberContext(p *ClassMemberContext)
func InitEmptyCompilationUnitContext(p *CompilationUnitContext)
func InitEmptyConstDeclContext(p *ConstDeclContext)
func InitEmptyContinueStmtContext(p *ContinueStmtContext)
func InitEmptyDeferStmtContext(p *DeferStmtContext)
func InitEmptyDeinitDeclContext(p *DeinitDeclContext)
func InitEmptyEqualityExpressionContext(p *EqualityExpressionContext)
func InitEmptyExpressionContext(p *ExpressionContext)
func InitEmptyExpressionStmtContext(p *ExpressionStmtContext)
func InitEmptyExternDeclContext(p *ExternDeclContext)
func InitEmptyExternFunctionDeclContext(p *ExternFunctionDeclContext)
func InitEmptyExternMemberContext(p *ExternMemberContext)
func InitEmptyExternParameterListContext(p *ExternParameterListContext)
func InitEmptyFieldInitContext(p *FieldInitContext)
func InitEmptyForStmtContext(p *ForStmtContext)
func InitEmptyFunctionDeclContext(p *FunctionDeclContext)
func InitEmptyIfStmtContext(p *IfStmtContext)
func InitEmptyImportDeclContext(p *ImportDeclContext)
func InitEmptyImportSpecContext(p *ImportSpecContext)
func InitEmptyIntrinsicExpressionContext(p *IntrinsicExpressionContext)
func InitEmptyLeftHandSideContext(p *LeftHandSideContext)
func InitEmptyLiteralContext(p *LiteralContext)
func InitEmptyLogicalAndExpressionContext(p *LogicalAndExpressionContext)
func InitEmptyLogicalOrExpressionContext(p *LogicalOrExpressionContext)
func InitEmptyMapEntryContext(p *MapEntryContext)
func InitEmptyMapLiteralContext(p *MapLiteralContext)
func InitEmptyMapTypeContext(p *MapTypeContext)
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
func InitEmptyRangeExpressionContext(p *RangeExpressionContext)
func InitEmptyReferenceTypeContext(p *ReferenceTypeContext)
func InitEmptyRelationalExpressionContext(p *RelationalExpressionContext)
func InitEmptyReturnStmtContext(p *ReturnStmtContext)
func InitEmptyStatementContext(p *StatementContext)
func InitEmptyStructDeclContext(p *StructDeclContext)
func InitEmptyStructFieldContext(p *StructFieldContext)
func InitEmptyStructLiteralContext(p *StructLiteralContext)
func InitEmptyStructMemberContext(p *StructMemberContext)
func InitEmptySyscallExpressionContext(p *SyscallExpressionContext)
func InitEmptyTopLevelDeclContext(p *TopLevelDeclContext)
func InitEmptyTypeContext(p *TypeContext)
func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext)
func InitEmptyVariableDeclContext(p *VariableDeclContext)
func InitEmptyVectorLiteralContext(p *VectorLiteralContext)
func InitEmptyVectorTypeContext(p *VectorTypeContext)

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

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener)

func (s *AdditiveExpressionContext) GetParser() antlr.Parser

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext

func (*AdditiveExpressionContext) IsAdditiveExpressionContext()

func (s *AdditiveExpressionContext) MINUS(i int) antlr.TerminalNode

func (s *AdditiveExpressionContext) MultiplicativeExpression(i int) IMultiplicativeExpressionContext

func (s *AdditiveExpressionContext) PLUS(i int) antlr.TerminalNode

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type AllocaExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewAllocaExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllocaExpressionContext

func NewEmptyAllocaExpressionContext() *AllocaExpressionContext

func (s *AllocaExpressionContext) ALLOCA() antlr.TerminalNode

func (s *AllocaExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *AllocaExpressionContext) COMMA() antlr.TerminalNode

func (s *AllocaExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *AllocaExpressionContext) ExitRule(listener antlr.ParseTreeListener)

func (s *AllocaExpressionContext) Expression() IExpressionContext

func (s *AllocaExpressionContext) GetParser() antlr.Parser

func (s *AllocaExpressionContext) GetRuleContext() antlr.RuleContext

func (*AllocaExpressionContext) IsAllocaExpressionContext()

func (s *AllocaExpressionContext) LPAREN() antlr.TerminalNode

func (s *AllocaExpressionContext) RPAREN() antlr.TerminalNode

func (s *AllocaExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *AllocaExpressionContext) Type_() ITypeContext

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

func (p *ArcParser) AllocaExpression() (localctx IAllocaExpressionContext)

func (p *ArcParser) ArgumentList() (localctx IArgumentListContext)

func (p *ArcParser) AssignmentStmt() (localctx IAssignmentStmtContext)

func (p *ArcParser) Block() (localctx IBlockContext)

func (p *ArcParser) BreakStmt() (localctx IBreakStmtContext)

func (p *ArcParser) CastExpression() (localctx ICastExpressionContext)

func (p *ArcParser) ClassDecl() (localctx IClassDeclContext)

func (p *ArcParser) ClassField() (localctx IClassFieldContext)

func (p *ArcParser) ClassMember() (localctx IClassMemberContext)

func (p *ArcParser) CompilationUnit() (localctx ICompilationUnitContext)

func (p *ArcParser) ConstDecl() (localctx IConstDeclContext)

func (p *ArcParser) ContinueStmt() (localctx IContinueStmtContext)

func (p *ArcParser) DeferStmt() (localctx IDeferStmtContext)

func (p *ArcParser) DeinitDecl() (localctx IDeinitDeclContext)

func (p *ArcParser) EqualityExpression() (localctx IEqualityExpressionContext)

func (p *ArcParser) Expression() (localctx IExpressionContext)

func (p *ArcParser) ExpressionStmt() (localctx IExpressionStmtContext)

func (p *ArcParser) ExternDecl() (localctx IExternDeclContext)

func (p *ArcParser) ExternFunctionDecl() (localctx IExternFunctionDeclContext)

func (p *ArcParser) ExternMember() (localctx IExternMemberContext)

func (p *ArcParser) ExternParameterList() (localctx IExternParameterListContext)

func (p *ArcParser) FieldInit() (localctx IFieldInitContext)

func (p *ArcParser) ForStmt() (localctx IForStmtContext)

func (p *ArcParser) FunctionDecl() (localctx IFunctionDeclContext)

func (p *ArcParser) IfStmt() (localctx IIfStmtContext)

func (p *ArcParser) ImportDecl() (localctx IImportDeclContext)

func (p *ArcParser) ImportSpec() (localctx IImportSpecContext)

func (p *ArcParser) IntrinsicExpression() (localctx IIntrinsicExpressionContext)

func (p *ArcParser) LeftHandSide() (localctx ILeftHandSideContext)

func (p *ArcParser) Literal() (localctx ILiteralContext)

func (p *ArcParser) LogicalAndExpression() (localctx ILogicalAndExpressionContext)

func (p *ArcParser) LogicalOrExpression() (localctx ILogicalOrExpressionContext)

func (p *ArcParser) MapEntry() (localctx IMapEntryContext)

func (p *ArcParser) MapLiteral() (localctx IMapLiteralContext)

func (p *ArcParser) MapType() (localctx IMapTypeContext)

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

func (p *ArcParser) RangeExpression() (localctx IRangeExpressionContext)

func (p *ArcParser) ReferenceType() (localctx IReferenceTypeContext)

func (p *ArcParser) RelationalExpression() (localctx IRelationalExpressionContext)

func (p *ArcParser) ReturnStmt() (localctx IReturnStmtContext)

func (p *ArcParser) Statement() (localctx IStatementContext)

func (p *ArcParser) StructDecl() (localctx IStructDeclContext)

func (p *ArcParser) StructField() (localctx IStructFieldContext)

func (p *ArcParser) StructLiteral() (localctx IStructLiteralContext)

func (p *ArcParser) StructMember() (localctx IStructMemberContext)

func (p *ArcParser) SyscallExpression() (localctx ISyscallExpressionContext)

func (p *ArcParser) TopLevelDecl() (localctx ITopLevelDeclContext)

func (p *ArcParser) Type_() (localctx ITypeContext)

func (p *ArcParser) UnaryExpression() (localctx IUnaryExpressionContext)

func (p *ArcParser) VariableDecl() (localctx IVariableDeclContext)

func (p *ArcParser) VectorLiteral() (localctx IVectorLiteralContext)

func (p *ArcParser) VectorType() (localctx IVectorTypeContext)

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

	// Visit a parse tree produced by ArcParser#externDecl.
	VisitExternDecl(ctx *ExternDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externMember.
	VisitExternMember(ctx *ExternMemberContext) interface{}

	// Visit a parse tree produced by ArcParser#externFunctionDecl.
	VisitExternFunctionDecl(ctx *ExternFunctionDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#externParameterList.
	VisitExternParameterList(ctx *ExternParameterListContext) interface{}

	// Visit a parse tree produced by ArcParser#functionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

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

	// Visit a parse tree produced by ArcParser#classDecl.
	VisitClassDecl(ctx *ClassDeclContext) interface{}

	// Visit a parse tree produced by ArcParser#classMember.
	VisitClassMember(ctx *ClassMemberContext) interface{}

	// Visit a parse tree produced by ArcParser#classField.
	VisitClassField(ctx *ClassFieldContext) interface{}

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

	// Visit a parse tree produced by ArcParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by ArcParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#pointerType.
	VisitPointerType(ctx *PointerTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#referenceType.
	VisitReferenceType(ctx *ReferenceTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#vectorType.
	VisitVectorType(ctx *VectorTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by ArcParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ArcParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ArcParser#assignmentStmt.
	VisitAssignmentStmt(ctx *AssignmentStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#leftHandSide.
	VisitLeftHandSide(ctx *LeftHandSideContext) interface{}

	// Visit a parse tree produced by ArcParser#expressionStmt.
	VisitExpressionStmt(ctx *ExpressionStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#continueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#deferStmt.
	VisitDeferStmt(ctx *DeferStmtContext) interface{}

	// Visit a parse tree produced by ArcParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

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

	// Visit a parse tree produced by ArcParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#vectorLiteral.
	VisitVectorLiteral(ctx *VectorLiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#mapLiteral.
	VisitMapLiteral(ctx *MapLiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#mapEntry.
	VisitMapEntry(ctx *MapEntryContext) interface{}

	// Visit a parse tree produced by ArcParser#structLiteral.
	VisitStructLiteral(ctx *StructLiteralContext) interface{}

	// Visit a parse tree produced by ArcParser#fieldInit.
	VisitFieldInit(ctx *FieldInitContext) interface{}

	// Visit a parse tree produced by ArcParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by ArcParser#castExpression.
	VisitCastExpression(ctx *CastExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#allocaExpression.
	VisitAllocaExpression(ctx *AllocaExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#syscallExpression.
	VisitSyscallExpression(ctx *SyscallExpressionContext) interface{}

	// Visit a parse tree produced by ArcParser#intrinsicExpression.
	VisitIntrinsicExpression(ctx *IntrinsicExpressionContext) interface{}
}
    A complete Visitor for a parse tree produced by ArcParser.

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext

func NewEmptyArgumentListContext() *ArgumentListContext

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode

func (s *ArgumentListContext) AllExpression() []IExpressionContext

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ArgumentListContext) Expression(i int) IExpressionContext

func (s *ArgumentListContext) GetParser() antlr.Parser

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext

func (*ArgumentListContext) IsArgumentListContext()

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type AssignmentStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewAssignmentStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStmtContext

func NewEmptyAssignmentStmtContext() *AssignmentStmtContext

func (s *AssignmentStmtContext) ASSIGN() antlr.TerminalNode

func (s *AssignmentStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *AssignmentStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *AssignmentStmtContext) ExitRule(listener antlr.ParseTreeListener)

func (s *AssignmentStmtContext) Expression() IExpressionContext

func (s *AssignmentStmtContext) GetParser() antlr.Parser

func (s *AssignmentStmtContext) GetRuleContext() antlr.RuleContext

func (*AssignmentStmtContext) IsAssignmentStmtContext()

func (s *AssignmentStmtContext) LeftHandSide() ILeftHandSideContext

func (s *AssignmentStmtContext) MINUS_ASSIGN() antlr.TerminalNode

func (s *AssignmentStmtContext) PERCENT_ASSIGN() antlr.TerminalNode

func (s *AssignmentStmtContext) PLUS_ASSIGN() antlr.TerminalNode

func (s *AssignmentStmtContext) SLASH_ASSIGN() antlr.TerminalNode

func (s *AssignmentStmtContext) STAR_ASSIGN() antlr.TerminalNode

func (s *AssignmentStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type BaseArcParserListener struct{}
    BaseArcParserListener is a complete listener for a parse tree produced by
    ArcParser.

func (s *BaseArcParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext)
    EnterAdditiveExpression is called when production additiveExpression is
    entered.

func (s *BaseArcParserListener) EnterAllocaExpression(ctx *AllocaExpressionContext)
    EnterAllocaExpression is called when production allocaExpression is entered.

func (s *BaseArcParserListener) EnterArgumentList(ctx *ArgumentListContext)
    EnterArgumentList is called when production argumentList is entered.

func (s *BaseArcParserListener) EnterAssignmentStmt(ctx *AssignmentStmtContext)
    EnterAssignmentStmt is called when production assignmentStmt is entered.

func (s *BaseArcParserListener) EnterBlock(ctx *BlockContext)
    EnterBlock is called when production block is entered.

func (s *BaseArcParserListener) EnterBreakStmt(ctx *BreakStmtContext)
    EnterBreakStmt is called when production breakStmt is entered.

func (s *BaseArcParserListener) EnterCastExpression(ctx *CastExpressionContext)
    EnterCastExpression is called when production castExpression is entered.

func (s *BaseArcParserListener) EnterClassDecl(ctx *ClassDeclContext)
    EnterClassDecl is called when production classDecl is entered.

func (s *BaseArcParserListener) EnterClassField(ctx *ClassFieldContext)
    EnterClassField is called when production classField is entered.

func (s *BaseArcParserListener) EnterClassMember(ctx *ClassMemberContext)
    EnterClassMember is called when production classMember is entered.

func (s *BaseArcParserListener) EnterCompilationUnit(ctx *CompilationUnitContext)
    EnterCompilationUnit is called when production compilationUnit is entered.

func (s *BaseArcParserListener) EnterConstDecl(ctx *ConstDeclContext)
    EnterConstDecl is called when production constDecl is entered.

func (s *BaseArcParserListener) EnterContinueStmt(ctx *ContinueStmtContext)
    EnterContinueStmt is called when production continueStmt is entered.

func (s *BaseArcParserListener) EnterDeferStmt(ctx *DeferStmtContext)
    EnterDeferStmt is called when production deferStmt is entered.

func (s *BaseArcParserListener) EnterDeinitDecl(ctx *DeinitDeclContext)
    EnterDeinitDecl is called when production deinitDecl is entered.

func (s *BaseArcParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext)
    EnterEqualityExpression is called when production equalityExpression is
    entered.

func (s *BaseArcParserListener) EnterEveryRule(ctx antlr.ParserRuleContext)
    EnterEveryRule is called when any rule is entered.

func (s *BaseArcParserListener) EnterExpression(ctx *ExpressionContext)
    EnterExpression is called when production expression is entered.

func (s *BaseArcParserListener) EnterExpressionStmt(ctx *ExpressionStmtContext)
    EnterExpressionStmt is called when production expressionStmt is entered.

func (s *BaseArcParserListener) EnterExternDecl(ctx *ExternDeclContext)
    EnterExternDecl is called when production externDecl is entered.

func (s *BaseArcParserListener) EnterExternFunctionDecl(ctx *ExternFunctionDeclContext)
    EnterExternFunctionDecl is called when production externFunctionDecl is
    entered.

func (s *BaseArcParserListener) EnterExternMember(ctx *ExternMemberContext)
    EnterExternMember is called when production externMember is entered.

func (s *BaseArcParserListener) EnterExternParameterList(ctx *ExternParameterListContext)
    EnterExternParameterList is called when production externParameterList is
    entered.

func (s *BaseArcParserListener) EnterFieldInit(ctx *FieldInitContext)
    EnterFieldInit is called when production fieldInit is entered.

func (s *BaseArcParserListener) EnterForStmt(ctx *ForStmtContext)
    EnterForStmt is called when production forStmt is entered.

func (s *BaseArcParserListener) EnterFunctionDecl(ctx *FunctionDeclContext)
    EnterFunctionDecl is called when production functionDecl is entered.

func (s *BaseArcParserListener) EnterIfStmt(ctx *IfStmtContext)
    EnterIfStmt is called when production ifStmt is entered.

func (s *BaseArcParserListener) EnterImportDecl(ctx *ImportDeclContext)
    EnterImportDecl is called when production importDecl is entered.

func (s *BaseArcParserListener) EnterImportSpec(ctx *ImportSpecContext)
    EnterImportSpec is called when production importSpec is entered.

func (s *BaseArcParserListener) EnterIntrinsicExpression(ctx *IntrinsicExpressionContext)
    EnterIntrinsicExpression is called when production intrinsicExpression is
    entered.

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

func (s *BaseArcParserListener) EnterMapEntry(ctx *MapEntryContext)
    EnterMapEntry is called when production mapEntry is entered.

func (s *BaseArcParserListener) EnterMapLiteral(ctx *MapLiteralContext)
    EnterMapLiteral is called when production mapLiteral is entered.

func (s *BaseArcParserListener) EnterMapType(ctx *MapTypeContext)
    EnterMapType is called when production mapType is entered.

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

func (s *BaseArcParserListener) EnterRangeExpression(ctx *RangeExpressionContext)
    EnterRangeExpression is called when production rangeExpression is entered.

func (s *BaseArcParserListener) EnterReferenceType(ctx *ReferenceTypeContext)
    EnterReferenceType is called when production referenceType is entered.

func (s *BaseArcParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext)
    EnterRelationalExpression is called when production relationalExpression is
    entered.

func (s *BaseArcParserListener) EnterReturnStmt(ctx *ReturnStmtContext)
    EnterReturnStmt is called when production returnStmt is entered.

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

func (s *BaseArcParserListener) EnterSyscallExpression(ctx *SyscallExpressionContext)
    EnterSyscallExpression is called when production syscallExpression is
    entered.

func (s *BaseArcParserListener) EnterTopLevelDecl(ctx *TopLevelDeclContext)
    EnterTopLevelDecl is called when production topLevelDecl is entered.

func (s *BaseArcParserListener) EnterType(ctx *TypeContext)
    EnterType is called when production type is entered.

func (s *BaseArcParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext)
    EnterUnaryExpression is called when production unaryExpression is entered.

func (s *BaseArcParserListener) EnterVariableDecl(ctx *VariableDeclContext)
    EnterVariableDecl is called when production variableDecl is entered.

func (s *BaseArcParserListener) EnterVectorLiteral(ctx *VectorLiteralContext)
    EnterVectorLiteral is called when production vectorLiteral is entered.

func (s *BaseArcParserListener) EnterVectorType(ctx *VectorTypeContext)
    EnterVectorType is called when production vectorType is entered.

func (s *BaseArcParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext)
    ExitAdditiveExpression is called when production additiveExpression is
    exited.

func (s *BaseArcParserListener) ExitAllocaExpression(ctx *AllocaExpressionContext)
    ExitAllocaExpression is called when production allocaExpression is exited.

func (s *BaseArcParserListener) ExitArgumentList(ctx *ArgumentListContext)
    ExitArgumentList is called when production argumentList is exited.

func (s *BaseArcParserListener) ExitAssignmentStmt(ctx *AssignmentStmtContext)
    ExitAssignmentStmt is called when production assignmentStmt is exited.

func (s *BaseArcParserListener) ExitBlock(ctx *BlockContext)
    ExitBlock is called when production block is exited.

func (s *BaseArcParserListener) ExitBreakStmt(ctx *BreakStmtContext)
    ExitBreakStmt is called when production breakStmt is exited.

func (s *BaseArcParserListener) ExitCastExpression(ctx *CastExpressionContext)
    ExitCastExpression is called when production castExpression is exited.

func (s *BaseArcParserListener) ExitClassDecl(ctx *ClassDeclContext)
    ExitClassDecl is called when production classDecl is exited.

func (s *BaseArcParserListener) ExitClassField(ctx *ClassFieldContext)
    ExitClassField is called when production classField is exited.

func (s *BaseArcParserListener) ExitClassMember(ctx *ClassMemberContext)
    ExitClassMember is called when production classMember is exited.

func (s *BaseArcParserListener) ExitCompilationUnit(ctx *CompilationUnitContext)
    ExitCompilationUnit is called when production compilationUnit is exited.

func (s *BaseArcParserListener) ExitConstDecl(ctx *ConstDeclContext)
    ExitConstDecl is called when production constDecl is exited.

func (s *BaseArcParserListener) ExitContinueStmt(ctx *ContinueStmtContext)
    ExitContinueStmt is called when production continueStmt is exited.

func (s *BaseArcParserListener) ExitDeferStmt(ctx *DeferStmtContext)
    ExitDeferStmt is called when production deferStmt is exited.

func (s *BaseArcParserListener) ExitDeinitDecl(ctx *DeinitDeclContext)
    ExitDeinitDecl is called when production deinitDecl is exited.

func (s *BaseArcParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext)
    ExitEqualityExpression is called when production equalityExpression is
    exited.

func (s *BaseArcParserListener) ExitEveryRule(ctx antlr.ParserRuleContext)
    ExitEveryRule is called when any rule is exited.

func (s *BaseArcParserListener) ExitExpression(ctx *ExpressionContext)
    ExitExpression is called when production expression is exited.

func (s *BaseArcParserListener) ExitExpressionStmt(ctx *ExpressionStmtContext)
    ExitExpressionStmt is called when production expressionStmt is exited.

func (s *BaseArcParserListener) ExitExternDecl(ctx *ExternDeclContext)
    ExitExternDecl is called when production externDecl is exited.

func (s *BaseArcParserListener) ExitExternFunctionDecl(ctx *ExternFunctionDeclContext)
    ExitExternFunctionDecl is called when production externFunctionDecl is
    exited.

func (s *BaseArcParserListener) ExitExternMember(ctx *ExternMemberContext)
    ExitExternMember is called when production externMember is exited.

func (s *BaseArcParserListener) ExitExternParameterList(ctx *ExternParameterListContext)
    ExitExternParameterList is called when production externParameterList is
    exited.

func (s *BaseArcParserListener) ExitFieldInit(ctx *FieldInitContext)
    ExitFieldInit is called when production fieldInit is exited.

func (s *BaseArcParserListener) ExitForStmt(ctx *ForStmtContext)
    ExitForStmt is called when production forStmt is exited.

func (s *BaseArcParserListener) ExitFunctionDecl(ctx *FunctionDeclContext)
    ExitFunctionDecl is called when production functionDecl is exited.

func (s *BaseArcParserListener) ExitIfStmt(ctx *IfStmtContext)
    ExitIfStmt is called when production ifStmt is exited.

func (s *BaseArcParserListener) ExitImportDecl(ctx *ImportDeclContext)
    ExitImportDecl is called when production importDecl is exited.

func (s *BaseArcParserListener) ExitImportSpec(ctx *ImportSpecContext)
    ExitImportSpec is called when production importSpec is exited.

func (s *BaseArcParserListener) ExitIntrinsicExpression(ctx *IntrinsicExpressionContext)
    ExitIntrinsicExpression is called when production intrinsicExpression is
    exited.

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

func (s *BaseArcParserListener) ExitMapEntry(ctx *MapEntryContext)
    ExitMapEntry is called when production mapEntry is exited.

func (s *BaseArcParserListener) ExitMapLiteral(ctx *MapLiteralContext)
    ExitMapLiteral is called when production mapLiteral is exited.

func (s *BaseArcParserListener) ExitMapType(ctx *MapTypeContext)
    ExitMapType is called when production mapType is exited.

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

func (s *BaseArcParserListener) ExitRangeExpression(ctx *RangeExpressionContext)
    ExitRangeExpression is called when production rangeExpression is exited.

func (s *BaseArcParserListener) ExitReferenceType(ctx *ReferenceTypeContext)
    ExitReferenceType is called when production referenceType is exited.

func (s *BaseArcParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext)
    ExitRelationalExpression is called when production relationalExpression is
    exited.

func (s *BaseArcParserListener) ExitReturnStmt(ctx *ReturnStmtContext)
    ExitReturnStmt is called when production returnStmt is exited.

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

func (s *BaseArcParserListener) ExitSyscallExpression(ctx *SyscallExpressionContext)
    ExitSyscallExpression is called when production syscallExpression is exited.

func (s *BaseArcParserListener) ExitTopLevelDecl(ctx *TopLevelDeclContext)
    ExitTopLevelDecl is called when production topLevelDecl is exited.

func (s *BaseArcParserListener) ExitType(ctx *TypeContext)
    ExitType is called when production type is exited.

func (s *BaseArcParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext)
    ExitUnaryExpression is called when production unaryExpression is exited.

func (s *BaseArcParserListener) ExitVariableDecl(ctx *VariableDeclContext)
    ExitVariableDecl is called when production variableDecl is exited.

func (s *BaseArcParserListener) ExitVectorLiteral(ctx *VectorLiteralContext)
    ExitVectorLiteral is called when production vectorLiteral is exited.

func (s *BaseArcParserListener) ExitVectorType(ctx *VectorTypeContext)
    ExitVectorType is called when production vectorType is exited.

func (s *BaseArcParserListener) VisitErrorNode(node antlr.ErrorNode)
    VisitErrorNode is called when an error node is visited.

func (s *BaseArcParserListener) VisitTerminal(node antlr.TerminalNode)
    VisitTerminal is called when a terminal node is visited.

type BaseArcParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseArcParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitAllocaExpression(ctx *AllocaExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{}

func (v *BaseArcParserVisitor) VisitAssignmentStmt(ctx *AssignmentStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitBlock(ctx *BlockContext) interface{}

func (v *BaseArcParserVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitCastExpression(ctx *CastExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitClassDecl(ctx *ClassDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitClassField(ctx *ClassFieldContext) interface{}

func (v *BaseArcParserVisitor) VisitClassMember(ctx *ClassMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

func (v *BaseArcParserVisitor) VisitConstDecl(ctx *ConstDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitDeferStmt(ctx *DeferStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitDeinitDecl(ctx *DeinitDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitExpression(ctx *ExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitExpressionStmt(ctx *ExpressionStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitExternDecl(ctx *ExternDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternFunctionDecl(ctx *ExternFunctionDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitExternMember(ctx *ExternMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitExternParameterList(ctx *ExternParameterListContext) interface{}

func (v *BaseArcParserVisitor) VisitFieldInit(ctx *FieldInitContext) interface{}

func (v *BaseArcParserVisitor) VisitForStmt(ctx *ForStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitIfStmt(ctx *IfStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitImportSpec(ctx *ImportSpecContext) interface{}

func (v *BaseArcParserVisitor) VisitIntrinsicExpression(ctx *IntrinsicExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitLeftHandSide(ctx *LeftHandSideContext) interface{}

func (v *BaseArcParserVisitor) VisitLiteral(ctx *LiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitMapEntry(ctx *MapEntryContext) interface{}

func (v *BaseArcParserVisitor) VisitMapLiteral(ctx *MapLiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitMapType(ctx *MapTypeContext) interface{}

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

func (v *BaseArcParserVisitor) VisitRangeExpression(ctx *RangeExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitReferenceType(ctx *ReferenceTypeContext) interface{}

func (v *BaseArcParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{}

func (v *BaseArcParserVisitor) VisitStatement(ctx *StatementContext) interface{}

func (v *BaseArcParserVisitor) VisitStructDecl(ctx *StructDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitStructField(ctx *StructFieldContext) interface{}

func (v *BaseArcParserVisitor) VisitStructLiteral(ctx *StructLiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitStructMember(ctx *StructMemberContext) interface{}

func (v *BaseArcParserVisitor) VisitSyscallExpression(ctx *SyscallExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitTopLevelDecl(ctx *TopLevelDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitType(ctx *TypeContext) interface{}

func (v *BaseArcParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

func (v *BaseArcParserVisitor) VisitVariableDecl(ctx *VariableDeclContext) interface{}

func (v *BaseArcParserVisitor) VisitVectorLiteral(ctx *VectorLiteralContext) interface{}

func (v *BaseArcParserVisitor) VisitVectorType(ctx *VectorTypeContext) interface{}

type BlockContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext

func NewEmptyBlockContext() *BlockContext

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *BlockContext) AllStatement() []IStatementContext

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener)

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener)

func (s *BreakStmtContext) GetParser() antlr.Parser

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext

func (*BreakStmtContext) IsBreakStmtContext()

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type CastExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewCastExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CastExpressionContext

func NewEmptyCastExpressionContext() *CastExpressionContext

func (s *CastExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *CastExpressionContext) CAST() antlr.TerminalNode

func (s *CastExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *CastExpressionContext) ExitRule(listener antlr.ParseTreeListener)

func (s *CastExpressionContext) Expression() IExpressionContext

func (s *CastExpressionContext) GT() antlr.TerminalNode

func (s *CastExpressionContext) GetParser() antlr.Parser

func (s *CastExpressionContext) GetRuleContext() antlr.RuleContext

func (*CastExpressionContext) IsCastExpressionContext()

func (s *CastExpressionContext) LPAREN() antlr.TerminalNode

func (s *CastExpressionContext) LT() antlr.TerminalNode

func (s *CastExpressionContext) RPAREN() antlr.TerminalNode

func (s *CastExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *CastExpressionContext) Type_() ITypeContext

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

func (s *ClassDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ClassDeclContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *ClassFieldContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ClassFieldContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *ClassMemberContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ClassMemberContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ClassMemberContext) FunctionDecl() IFunctionDeclContext

func (s *ClassMemberContext) GetParser() antlr.Parser

func (s *ClassMemberContext) GetRuleContext() antlr.RuleContext

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

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener)

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener)

func (s *CompilationUnitContext) GetParser() antlr.Parser

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext

func (s *CompilationUnitContext) ImportDecl(i int) IImportDeclContext

func (*CompilationUnitContext) IsCompilationUnitContext()

func (s *CompilationUnitContext) NamespaceDecl(i int) INamespaceDeclContext

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *CompilationUnitContext) TopLevelDecl(i int) ITopLevelDeclContext

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

func (s *ConstDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ConstDeclContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ContinueStmtContext) GetParser() antlr.Parser

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext

func (*ContinueStmtContext) IsContinueStmtContext()

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type DeferStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewDeferStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeferStmtContext

func NewEmptyDeferStmtContext() *DeferStmtContext

func (s *DeferStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *DeferStmtContext) AssignmentStmt() IAssignmentStmtContext

func (s *DeferStmtContext) DEFER() antlr.TerminalNode

func (s *DeferStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *DeferStmtContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *DeinitDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *DeinitDeclContext) ExitRule(listener antlr.ParseTreeListener)

func (s *DeinitDeclContext) GetParser() antlr.Parser

func (s *DeinitDeclContext) GetRuleContext() antlr.RuleContext

func (s *DeinitDeclContext) IDENTIFIER() antlr.TerminalNode

func (*DeinitDeclContext) IsDeinitDeclContext()

func (s *DeinitDeclContext) LPAREN() antlr.TerminalNode

func (s *DeinitDeclContext) RPAREN() antlr.TerminalNode

func (s *DeinitDeclContext) SELF() antlr.TerminalNode

func (s *DeinitDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *DeinitDeclContext) Type_() ITypeContext

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

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener)

func (s *EqualityExpressionContext) GetParser() antlr.Parser

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext

func (*EqualityExpressionContext) IsEqualityExpressionContext()

func (s *EqualityExpressionContext) NE(i int) antlr.TerminalNode

func (s *EqualityExpressionContext) RelationalExpression(i int) IRelationalExpressionContext

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExpressionContext() *ExpressionContext

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *ExpressionStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ExpressionStmtContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ExpressionStmtContext) Expression() IExpressionContext

func (s *ExpressionStmtContext) GetParser() antlr.Parser

func (s *ExpressionStmtContext) GetRuleContext() antlr.RuleContext

func (*ExpressionStmtContext) IsExpressionStmtContext()

func (s *ExpressionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternDeclContext() *ExternDeclContext

func NewExternDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternDeclContext

func (s *ExternDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternDeclContext) AllExternMember() []IExternMemberContext

func (s *ExternDeclContext) EXTERN() antlr.TerminalNode

func (s *ExternDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ExternDeclContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ExternDeclContext) ExternMember(i int) IExternMemberContext

func (s *ExternDeclContext) GetParser() antlr.Parser

func (s *ExternDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternDeclContext) IsExternDeclContext()

func (s *ExternDeclContext) LBRACE() antlr.TerminalNode

func (s *ExternDeclContext) RBRACE() antlr.TerminalNode

func (s *ExternDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternFunctionDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternFunctionDeclContext() *ExternFunctionDeclContext

func NewExternFunctionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternFunctionDeclContext

func (s *ExternFunctionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternFunctionDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ExternFunctionDeclContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ExternFunctionDeclContext) ExternParameterList() IExternParameterListContext

func (s *ExternFunctionDeclContext) FUNC() antlr.TerminalNode

func (s *ExternFunctionDeclContext) GetParser() antlr.Parser

func (s *ExternFunctionDeclContext) GetRuleContext() antlr.RuleContext

func (s *ExternFunctionDeclContext) IDENTIFIER() antlr.TerminalNode

func (*ExternFunctionDeclContext) IsExternFunctionDeclContext()

func (s *ExternFunctionDeclContext) LPAREN() antlr.TerminalNode

func (s *ExternFunctionDeclContext) RPAREN() antlr.TerminalNode

func (s *ExternFunctionDeclContext) STRING_LITERAL() antlr.TerminalNode

func (s *ExternFunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternFunctionDeclContext) Type_() ITypeContext

type ExternMemberContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternMemberContext() *ExternMemberContext

func NewExternMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternMemberContext

func (s *ExternMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternMemberContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ExternMemberContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ExternMemberContext) ExternFunctionDecl() IExternFunctionDeclContext

func (s *ExternMemberContext) GetParser() antlr.Parser

func (s *ExternMemberContext) GetRuleContext() antlr.RuleContext

func (*ExternMemberContext) IsExternMemberContext()

func (s *ExternMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ExternParameterListContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyExternParameterListContext() *ExternParameterListContext

func NewExternParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternParameterListContext

func (s *ExternParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ExternParameterListContext) AllCOMMA() []antlr.TerminalNode

func (s *ExternParameterListContext) AllType_() []ITypeContext

func (s *ExternParameterListContext) COMMA(i int) antlr.TerminalNode

func (s *ExternParameterListContext) ELLIPSIS() antlr.TerminalNode

func (s *ExternParameterListContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ExternParameterListContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ExternParameterListContext) GetParser() antlr.Parser

func (s *ExternParameterListContext) GetRuleContext() antlr.RuleContext

func (*ExternParameterListContext) IsExternParameterListContext()

func (s *ExternParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *ExternParameterListContext) Type_(i int) ITypeContext

type FieldInitContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyFieldInitContext() *FieldInitContext

func NewFieldInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldInitContext

func (s *FieldInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *FieldInitContext) COLON() antlr.TerminalNode

func (s *FieldInitContext) EnterRule(listener antlr.ParseTreeListener)

func (s *FieldInitContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *FunctionDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *FunctionDeclContext) ExitRule(listener antlr.ParseTreeListener)

func (s *FunctionDeclContext) FUNC() antlr.TerminalNode

func (s *FunctionDeclContext) GetParser() antlr.Parser

func (s *FunctionDeclContext) GetRuleContext() antlr.RuleContext

func (s *FunctionDeclContext) IDENTIFIER() antlr.TerminalNode

func (*FunctionDeclContext) IsFunctionDeclContext()

func (s *FunctionDeclContext) LPAREN() antlr.TerminalNode

func (s *FunctionDeclContext) ParameterList() IParameterListContext

func (s *FunctionDeclContext) RPAREN() antlr.TerminalNode

func (s *FunctionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *FunctionDeclContext) Type_() ITypeContext

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

type IAllocaExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALLOCA() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Type_() ITypeContext
	RPAREN() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAllocaExpressionContext differentiates from other interfaces.
	IsAllocaExpressionContext()
}
    IAllocaExpressionContext is an interface to support dynamic dispatch.

type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}
    IArgumentListContext is an interface to support dynamic dispatch.

type IAssignmentStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LeftHandSide() ILeftHandSideContext
	Expression() IExpressionContext
	ASSIGN() antlr.TerminalNode
	PLUS_ASSIGN() antlr.TerminalNode
	MINUS_ASSIGN() antlr.TerminalNode
	STAR_ASSIGN() antlr.TerminalNode
	SLASH_ASSIGN() antlr.TerminalNode
	PERCENT_ASSIGN() antlr.TerminalNode

	// IsAssignmentStmtContext differentiates from other interfaces.
	IsAssignmentStmtContext()
}
    IAssignmentStmtContext is an interface to support dynamic dispatch.

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

type ICastExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CAST() antlr.TerminalNode
	LT() antlr.TerminalNode
	Type_() ITypeContext
	GT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsCastExpressionContext differentiates from other interfaces.
	IsCastExpressionContext()
}
    ICastExpressionContext is an interface to support dynamic dispatch.

type IClassDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLASS() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
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

type IExternDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTERN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllExternMember() []IExternMemberContext
	ExternMember(i int) IExternMemberContext

	// IsExternDeclContext differentiates from other interfaces.
	IsExternDeclContext()
}
    IExternDeclContext is an interface to support dynamic dispatch.

type IExternFunctionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	ExternParameterList() IExternParameterListContext
	Type_() ITypeContext

	// IsExternFunctionDeclContext differentiates from other interfaces.
	IsExternFunctionDeclContext()
}
    IExternFunctionDeclContext is an interface to support dynamic dispatch.

type IExternMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExternFunctionDecl() IExternFunctionDeclContext

	// IsExternMemberContext differentiates from other interfaces.
	IsExternMemberContext()
}
    IExternMemberContext is an interface to support dynamic dispatch.

type IExternParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	ELLIPSIS() antlr.TerminalNode

	// IsExternParameterListContext differentiates from other interfaces.
	IsExternParameterListContext()
}
    IExternParameterListContext is an interface to support dynamic dispatch.

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
	ASYNC() antlr.TerminalNode
	ParameterList() IParameterListContext
	Type_() ITypeContext

	// IsFunctionDeclContext differentiates from other interfaces.
	IsFunctionDeclContext()
}
    IFunctionDeclContext is an interface to support dynamic dispatch.

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

type IIntrinsicExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SIZEOF() antlr.TerminalNode
	LT() antlr.TerminalNode
	Type_() ITypeContext
	GT() antlr.TerminalNode
	ALIGNOF() antlr.TerminalNode
	MEMSET() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	MEMCPY() antlr.TerminalNode
	MEMMOVE() antlr.TerminalNode
	STRLEN() antlr.TerminalNode
	MEMCHR() antlr.TerminalNode
	VA_START() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	VA_ARG() antlr.TerminalNode
	VA_END() antlr.TerminalNode
	RAISE() antlr.TerminalNode
	MEMCMP() antlr.TerminalNode
	BIT_CAST() antlr.TerminalNode

	// IsIntrinsicExpressionContext differentiates from other interfaces.
	IsIntrinsicExpressionContext()
}
    IIntrinsicExpressionContext is an interface to support dynamic dispatch.

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
	VectorLiteral() IVectorLiteralContext
	MapLiteral() IMapLiteralContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}
    ILiteralContext is an interface to support dynamic dispatch.

type ILogicalAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEqualityExpression() []IEqualityExpressionContext
	EqualityExpression(i int) IEqualityExpressionContext
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

type IMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COLON() antlr.TerminalNode

	// IsMapEntryContext differentiates from other interfaces.
	IsMapEntryContext()
}
    IMapEntryContext is an interface to support dynamic dispatch.

type IMapLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllMapEntry() []IMapEntryContext
	MapEntry(i int) IMapEntryContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMapLiteralContext differentiates from other interfaces.
	IsMapLiteralContext()
}
    IMapLiteralContext is an interface to support dynamic dispatch.

type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAP() antlr.TerminalNode
	LT() antlr.TerminalNode
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	COMMA() antlr.TerminalNode
	GT() antlr.TerminalNode

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}
    IMapTypeContext is an interface to support dynamic dispatch.

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
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	ASYNC() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext

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
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext

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
	Literal() ILiteralContext
	StructLiteral() IStructLiteralContext
	CastExpression() ICastExpressionContext
	AllocaExpression() IAllocaExpressionContext
	SyscallExpression() ISyscallExpressionContext
	IntrinsicExpression() IIntrinsicExpressionContext
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
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
	BYTE() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	STRING() antlr.TerminalNode
	VOID() antlr.TerminalNode

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}
    IPrimitiveTypeContext is an interface to support dynamic dispatch.

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

	// IsReferenceTypeContext differentiates from other interfaces.
	IsReferenceTypeContext()
}
    IReferenceTypeContext is an interface to support dynamic dispatch.

type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRangeExpression() []IRangeExpressionContext
	RangeExpression(i int) IRangeExpressionContext
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
	Expression() IExpressionContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}
    IReturnStmtContext is an interface to support dynamic dispatch.

type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableDecl() IVariableDeclContext
	ConstDecl() IConstDeclContext
	AssignmentStmt() IAssignmentStmtContext
	ExpressionStmt() IExpressionStmtContext
	ReturnStmt() IReturnStmtContext
	IfStmt() IIfStmtContext
	ForStmt() IForStmtContext
	BreakStmt() IBreakStmtContext
	ContinueStmt() IContinueStmtContext
	DeferStmt() IDeferStmtContext
	Block() IBlockContext

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
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
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

	// IsStructMemberContext differentiates from other interfaces.
	IsStructMemberContext()
}
    IStructMemberContext is an interface to support dynamic dispatch.

type ISyscallExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYSCALL() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsSyscallExpressionContext differentiates from other interfaces.
	IsSyscallExpressionContext()
}
    ISyscallExpressionContext is an interface to support dynamic dispatch.

type ITopLevelDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionDecl() IFunctionDeclContext
	StructDecl() IStructDeclContext
	ClassDecl() IClassDeclContext
	MethodDecl() IMethodDeclContext
	MutatingDecl() IMutatingDeclContext
	DeinitDecl() IDeinitDeclContext
	VariableDecl() IVariableDeclContext
	ConstDecl() IConstDeclContext
	ExternDecl() IExternDeclContext

	// IsTopLevelDeclContext differentiates from other interfaces.
	IsTopLevelDeclContext()
}
    ITopLevelDeclContext is an interface to support dynamic dispatch.

type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimitiveType() IPrimitiveTypeContext
	PointerType() IPointerTypeContext
	ReferenceType() IReferenceTypeContext
	VectorType() IVectorTypeContext
	MapType() IMapTypeContext
	IDENTIFIER() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}
    ITypeContext is an interface to support dynamic dispatch.

type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpression() IUnaryExpressionContext
	MINUS() antlr.TerminalNode
	NOT() antlr.TerminalNode
	STAR() antlr.TerminalNode
	AMP() antlr.TerminalNode
	AWAIT() antlr.TerminalNode
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
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsVariableDeclContext differentiates from other interfaces.
	IsVariableDeclContext()
}
    IVariableDeclContext is an interface to support dynamic dispatch.

type IVectorLiteralContext interface {
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

	// IsVectorLiteralContext differentiates from other interfaces.
	IsVectorLiteralContext()
}
    IVectorLiteralContext is an interface to support dynamic dispatch.

type IVectorTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VECTOR() antlr.TerminalNode
	LT() antlr.TerminalNode
	Type_() ITypeContext
	GT() antlr.TerminalNode

	// IsVectorTypeContext differentiates from other interfaces.
	IsVectorTypeContext()
}
    IVectorTypeContext is an interface to support dynamic dispatch.

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

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *ImportDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ImportDeclContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *ImportSpecContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ImportSpecContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ImportSpecContext) GetParser() antlr.Parser

func (s *ImportSpecContext) GetRuleContext() antlr.RuleContext

func (*ImportSpecContext) IsImportSpecContext()

func (s *ImportSpecContext) STRING_LITERAL() antlr.TerminalNode

func (s *ImportSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type IntrinsicExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyIntrinsicExpressionContext() *IntrinsicExpressionContext

func NewIntrinsicExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntrinsicExpressionContext

func (s *IntrinsicExpressionContext) ALIGNOF() antlr.TerminalNode

func (s *IntrinsicExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *IntrinsicExpressionContext) AllCOMMA() []antlr.TerminalNode

func (s *IntrinsicExpressionContext) AllExpression() []IExpressionContext

func (s *IntrinsicExpressionContext) BIT_CAST() antlr.TerminalNode

func (s *IntrinsicExpressionContext) COMMA(i int) antlr.TerminalNode

func (s *IntrinsicExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *IntrinsicExpressionContext) ExitRule(listener antlr.ParseTreeListener)

func (s *IntrinsicExpressionContext) Expression(i int) IExpressionContext

func (s *IntrinsicExpressionContext) GT() antlr.TerminalNode

func (s *IntrinsicExpressionContext) GetParser() antlr.Parser

func (s *IntrinsicExpressionContext) GetRuleContext() antlr.RuleContext

func (s *IntrinsicExpressionContext) IDENTIFIER() antlr.TerminalNode

func (*IntrinsicExpressionContext) IsIntrinsicExpressionContext()

func (s *IntrinsicExpressionContext) LPAREN() antlr.TerminalNode

func (s *IntrinsicExpressionContext) LT() antlr.TerminalNode

func (s *IntrinsicExpressionContext) MEMCHR() antlr.TerminalNode

func (s *IntrinsicExpressionContext) MEMCMP() antlr.TerminalNode

func (s *IntrinsicExpressionContext) MEMCPY() antlr.TerminalNode

func (s *IntrinsicExpressionContext) MEMMOVE() antlr.TerminalNode

func (s *IntrinsicExpressionContext) MEMSET() antlr.TerminalNode

func (s *IntrinsicExpressionContext) RAISE() antlr.TerminalNode

func (s *IntrinsicExpressionContext) RPAREN() antlr.TerminalNode

func (s *IntrinsicExpressionContext) SIZEOF() antlr.TerminalNode

func (s *IntrinsicExpressionContext) STRLEN() antlr.TerminalNode

func (s *IntrinsicExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *IntrinsicExpressionContext) Type_() ITypeContext

func (s *IntrinsicExpressionContext) VA_ARG() antlr.TerminalNode

func (s *IntrinsicExpressionContext) VA_END() antlr.TerminalNode

func (s *IntrinsicExpressionContext) VA_START() antlr.TerminalNode

type LeftHandSideContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyLeftHandSideContext() *LeftHandSideContext

func NewLeftHandSideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftHandSideContext

func (s *LeftHandSideContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *LeftHandSideContext) DOT() antlr.TerminalNode

func (s *LeftHandSideContext) EnterRule(listener antlr.ParseTreeListener)

func (s *LeftHandSideContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener)

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener)

func (s *LiteralContext) FLOAT_LITERAL() antlr.TerminalNode

func (s *LiteralContext) GetParser() antlr.Parser

func (s *LiteralContext) GetRuleContext() antlr.RuleContext

func (s *LiteralContext) INTEGER_LITERAL() antlr.TerminalNode

func (*LiteralContext) IsLiteralContext()

func (s *LiteralContext) MapLiteral() IMapLiteralContext

func (s *LiteralContext) NULL() antlr.TerminalNode

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *LiteralContext) VectorLiteral() IVectorLiteralContext

type LogicalAndExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyLogicalAndExpressionContext() *LogicalAndExpressionContext

func NewLogicalAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndExpressionContext

func (s *LogicalAndExpressionContext) AND(i int) antlr.TerminalNode

func (s *LogicalAndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *LogicalAndExpressionContext) AllAND() []antlr.TerminalNode

func (s *LogicalAndExpressionContext) AllEqualityExpression() []IEqualityExpressionContext

func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *LogicalAndExpressionContext) EqualityExpression(i int) IEqualityExpressionContext

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener)

func (s *LogicalOrExpressionContext) GetParser() antlr.Parser

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext

func (*LogicalOrExpressionContext) IsLogicalOrExpressionContext()

func (s *LogicalOrExpressionContext) LogicalAndExpression(i int) ILogicalAndExpressionContext

func (s *LogicalOrExpressionContext) OR(i int) antlr.TerminalNode

func (s *LogicalOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type MapEntryContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyMapEntryContext() *MapEntryContext

func NewMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapEntryContext

func (s *MapEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *MapEntryContext) AllExpression() []IExpressionContext

func (s *MapEntryContext) COLON() antlr.TerminalNode

func (s *MapEntryContext) EnterRule(listener antlr.ParseTreeListener)

func (s *MapEntryContext) ExitRule(listener antlr.ParseTreeListener)

func (s *MapEntryContext) Expression(i int) IExpressionContext

func (s *MapEntryContext) GetParser() antlr.Parser

func (s *MapEntryContext) GetRuleContext() antlr.RuleContext

func (*MapEntryContext) IsMapEntryContext()

func (s *MapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type MapLiteralContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyMapLiteralContext() *MapLiteralContext

func NewMapLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapLiteralContext

func (s *MapLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *MapLiteralContext) AllCOMMA() []antlr.TerminalNode

func (s *MapLiteralContext) AllMapEntry() []IMapEntryContext

func (s *MapLiteralContext) COMMA(i int) antlr.TerminalNode

func (s *MapLiteralContext) EnterRule(listener antlr.ParseTreeListener)

func (s *MapLiteralContext) ExitRule(listener antlr.ParseTreeListener)

func (s *MapLiteralContext) GetParser() antlr.Parser

func (s *MapLiteralContext) GetRuleContext() antlr.RuleContext

func (*MapLiteralContext) IsMapLiteralContext()

func (s *MapLiteralContext) LBRACE() antlr.TerminalNode

func (s *MapLiteralContext) MapEntry(i int) IMapEntryContext

func (s *MapLiteralContext) RBRACE() antlr.TerminalNode

func (s *MapLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type MapTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyMapTypeContext() *MapTypeContext

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext

func (s *MapTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *MapTypeContext) AllType_() []ITypeContext

func (s *MapTypeContext) COMMA() antlr.TerminalNode

func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener)

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener)

func (s *MapTypeContext) GT() antlr.TerminalNode

func (s *MapTypeContext) GetParser() antlr.Parser

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext

func (*MapTypeContext) IsMapTypeContext()

func (s *MapTypeContext) LT() antlr.TerminalNode

func (s *MapTypeContext) MAP() antlr.TerminalNode

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *MapTypeContext) Type_(i int) ITypeContext

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

func (s *MethodDeclContext) AllType_() []ITypeContext

func (s *MethodDeclContext) Block() IBlockContext

func (s *MethodDeclContext) COLON() antlr.TerminalNode

func (s *MethodDeclContext) COMMA(i int) antlr.TerminalNode

func (s *MethodDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *MethodDeclContext) ExitRule(listener antlr.ParseTreeListener)

func (s *MethodDeclContext) FUNC() antlr.TerminalNode

func (s *MethodDeclContext) GetParser() antlr.Parser

func (s *MethodDeclContext) GetRuleContext() antlr.RuleContext

func (s *MethodDeclContext) IDENTIFIER(i int) antlr.TerminalNode

func (*MethodDeclContext) IsMethodDeclContext()

func (s *MethodDeclContext) LPAREN() antlr.TerminalNode

func (s *MethodDeclContext) Parameter(i int) IParameterContext

func (s *MethodDeclContext) RPAREN() antlr.TerminalNode

func (s *MethodDeclContext) SELF() antlr.TerminalNode

func (s *MethodDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *MethodDeclContext) Type_(i int) ITypeContext

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

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *MutatingDeclContext) AllType_() []ITypeContext

func (s *MutatingDeclContext) Block() IBlockContext

func (s *MutatingDeclContext) COLON() antlr.TerminalNode

func (s *MutatingDeclContext) COMMA(i int) antlr.TerminalNode

func (s *MutatingDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *MutatingDeclContext) ExitRule(listener antlr.ParseTreeListener)

func (s *MutatingDeclContext) GetParser() antlr.Parser

func (s *MutatingDeclContext) GetRuleContext() antlr.RuleContext

func (s *MutatingDeclContext) IDENTIFIER(i int) antlr.TerminalNode

func (*MutatingDeclContext) IsMutatingDeclContext()

func (s *MutatingDeclContext) LPAREN() antlr.TerminalNode

func (s *MutatingDeclContext) MUTATING() antlr.TerminalNode

func (s *MutatingDeclContext) Parameter(i int) IParameterContext

func (s *MutatingDeclContext) RPAREN() antlr.TerminalNode

func (s *MutatingDeclContext) SELF() antlr.TerminalNode

func (s *MutatingDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *MutatingDeclContext) Type_(i int) ITypeContext

type NamespaceDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyNamespaceDeclContext() *NamespaceDeclContext

func NewNamespaceDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceDeclContext

func (s *NamespaceDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *NamespaceDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *NamespaceDeclContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *PointerTypeContext) EnterRule(listener antlr.ParseTreeListener)

func (s *PointerTypeContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *PostfixExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *PostfixExpressionContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *PostfixOpContext) EnterRule(listener antlr.ParseTreeListener)

func (s *PostfixOpContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *PrimaryExpressionContext) AllocaExpression() IAllocaExpressionContext

func (s *PrimaryExpressionContext) CastExpression() ICastExpressionContext

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener)

func (s *PrimaryExpressionContext) Expression() IExpressionContext

func (s *PrimaryExpressionContext) GetParser() antlr.Parser

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext

func (s *PrimaryExpressionContext) IDENTIFIER() antlr.TerminalNode

func (s *PrimaryExpressionContext) IntrinsicExpression() IIntrinsicExpressionContext

func (*PrimaryExpressionContext) IsPrimaryExpressionContext()

func (s *PrimaryExpressionContext) LPAREN() antlr.TerminalNode

func (s *PrimaryExpressionContext) Literal() ILiteralContext

func (s *PrimaryExpressionContext) RPAREN() antlr.TerminalNode

func (s *PrimaryExpressionContext) StructLiteral() IStructLiteralContext

func (s *PrimaryExpressionContext) SyscallExpression() ISyscallExpressionContext

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type PrimitiveTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *PrimitiveTypeContext) BOOL() antlr.TerminalNode

func (s *PrimitiveTypeContext) BYTE() antlr.TerminalNode

func (s *PrimitiveTypeContext) CHAR() antlr.TerminalNode

func (s *PrimitiveTypeContext) EnterRule(listener antlr.ParseTreeListener)

func (s *PrimitiveTypeContext) ExitRule(listener antlr.ParseTreeListener)

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

type RangeExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyRangeExpressionContext() *RangeExpressionContext

func NewRangeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeExpressionContext

func (s *RangeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *RangeExpressionContext) AdditiveExpression(i int) IAdditiveExpressionContext

func (s *RangeExpressionContext) AllAdditiveExpression() []IAdditiveExpressionContext

func (s *RangeExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *RangeExpressionContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *ReferenceTypeContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ReferenceTypeContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *RelationalExpressionContext) AllRangeExpression() []IRangeExpressionContext

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener)

func (s *RelationalExpressionContext) GE(i int) antlr.TerminalNode

func (s *RelationalExpressionContext) GT(i int) antlr.TerminalNode

func (s *RelationalExpressionContext) GetParser() antlr.Parser

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext

func (*RelationalExpressionContext) IsRelationalExpressionContext()

func (s *RelationalExpressionContext) LE(i int) antlr.TerminalNode

func (s *RelationalExpressionContext) LT(i int) antlr.TerminalNode

func (s *RelationalExpressionContext) RangeExpression(i int) IRangeExpressionContext

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyReturnStmtContext() *ReturnStmtContext

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener)

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener)

func (s *ReturnStmtContext) Expression() IExpressionContext

func (s *ReturnStmtContext) GetParser() antlr.Parser

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext

func (*ReturnStmtContext) IsReturnStmtContext()

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

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

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener)

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener)

func (s *StatementContext) ExpressionStmt() IExpressionStmtContext

func (s *StatementContext) ForStmt() IForStmtContext

func (s *StatementContext) GetParser() antlr.Parser

func (s *StatementContext) GetRuleContext() antlr.RuleContext

func (s *StatementContext) IfStmt() IIfStmtContext

func (*StatementContext) IsStatementContext()

func (s *StatementContext) ReturnStmt() IReturnStmtContext

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *StatementContext) VariableDecl() IVariableDeclContext

type StructDeclContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyStructDeclContext() *StructDeclContext

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StructDeclContext) AllStructMember() []IStructMemberContext

func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *StructFieldContext) EnterRule(listener antlr.ParseTreeListener)

func (s *StructFieldContext) ExitRule(listener antlr.ParseTreeListener)

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

func (s *StructLiteralContext) EnterRule(listener antlr.ParseTreeListener)

func (s *StructLiteralContext) ExitRule(listener antlr.ParseTreeListener)

func (s *StructLiteralContext) FieldInit(i int) IFieldInitContext

func (s *StructLiteralContext) GetParser() antlr.Parser

func (s *StructLiteralContext) GetRuleContext() antlr.RuleContext

func (s *StructLiteralContext) IDENTIFIER() antlr.TerminalNode

func (*StructLiteralContext) IsStructLiteralContext()

func (s *StructLiteralContext) LBRACE() antlr.TerminalNode

func (s *StructLiteralContext) RBRACE() antlr.TerminalNode

func (s *StructLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type StructMemberContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyStructMemberContext() *StructMemberContext

func NewStructMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructMemberContext

func (s *StructMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *StructMemberContext) EnterRule(listener antlr.ParseTreeListener)

func (s *StructMemberContext) ExitRule(listener antlr.ParseTreeListener)

func (s *StructMemberContext) FunctionDecl() IFunctionDeclContext

func (s *StructMemberContext) GetParser() antlr.Parser

func (s *StructMemberContext) GetRuleContext() antlr.RuleContext

func (*StructMemberContext) IsStructMemberContext()

func (s *StructMemberContext) MutatingDecl() IMutatingDeclContext

func (s *StructMemberContext) StructField() IStructFieldContext

func (s *StructMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type SyscallExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptySyscallExpressionContext() *SyscallExpressionContext

func NewSyscallExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyscallExpressionContext

func (s *SyscallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *SyscallExpressionContext) AllCOMMA() []antlr.TerminalNode

func (s *SyscallExpressionContext) AllExpression() []IExpressionContext

func (s *SyscallExpressionContext) COMMA(i int) antlr.TerminalNode

func (s *SyscallExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *SyscallExpressionContext) ExitRule(listener antlr.ParseTreeListener)

func (s *SyscallExpressionContext) Expression(i int) IExpressionContext

func (s *SyscallExpressionContext) GetParser() antlr.Parser

func (s *SyscallExpressionContext) GetRuleContext() antlr.RuleContext

func (*SyscallExpressionContext) IsSyscallExpressionContext()

func (s *SyscallExpressionContext) LPAREN() antlr.TerminalNode

func (s *SyscallExpressionContext) RPAREN() antlr.TerminalNode

func (s *SyscallExpressionContext) SYSCALL() antlr.TerminalNode

func (s *SyscallExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

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

func (s *TopLevelDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *TopLevelDeclContext) ExitRule(listener antlr.ParseTreeListener)

func (s *TopLevelDeclContext) ExternDecl() IExternDeclContext

func (s *TopLevelDeclContext) FunctionDecl() IFunctionDeclContext

func (s *TopLevelDeclContext) GetParser() antlr.Parser

func (s *TopLevelDeclContext) GetRuleContext() antlr.RuleContext

func (*TopLevelDeclContext) IsTopLevelDeclContext()

func (s *TopLevelDeclContext) MethodDecl() IMethodDeclContext

func (s *TopLevelDeclContext) MutatingDecl() IMutatingDeclContext

func (s *TopLevelDeclContext) StructDecl() IStructDeclContext

func (s *TopLevelDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *TopLevelDeclContext) VariableDecl() IVariableDeclContext

type TypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyTypeContext() *TypeContext

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener)

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener)

func (s *TypeContext) GetParser() antlr.Parser

func (s *TypeContext) GetRuleContext() antlr.RuleContext

func (s *TypeContext) IDENTIFIER() antlr.TerminalNode

func (*TypeContext) IsTypeContext()

func (s *TypeContext) MapType() IMapTypeContext

func (s *TypeContext) PointerType() IPointerTypeContext

func (s *TypeContext) PrimitiveType() IPrimitiveTypeContext

func (s *TypeContext) ReferenceType() IReferenceTypeContext

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *TypeContext) VectorType() IVectorTypeContext

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext

func (s *UnaryExpressionContext) AMP() antlr.TerminalNode

func (s *UnaryExpressionContext) AWAIT() antlr.TerminalNode

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *UnaryExpressionContext) DECREMENT() antlr.TerminalNode

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener)

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener)

func (s *UnaryExpressionContext) GetParser() antlr.Parser

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext

func (s *UnaryExpressionContext) INCREMENT() antlr.TerminalNode

func (*UnaryExpressionContext) IsUnaryExpressionContext()

func (s *UnaryExpressionContext) MINUS() antlr.TerminalNode

func (s *UnaryExpressionContext) NOT() antlr.TerminalNode

func (s *UnaryExpressionContext) PostfixExpression() IPostfixExpressionContext

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

func (s *VariableDeclContext) EnterRule(listener antlr.ParseTreeListener)

func (s *VariableDeclContext) ExitRule(listener antlr.ParseTreeListener)

func (s *VariableDeclContext) Expression() IExpressionContext

func (s *VariableDeclContext) GetParser() antlr.Parser

func (s *VariableDeclContext) GetRuleContext() antlr.RuleContext

func (s *VariableDeclContext) IDENTIFIER() antlr.TerminalNode

func (*VariableDeclContext) IsVariableDeclContext()

func (s *VariableDeclContext) LET() antlr.TerminalNode

func (s *VariableDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *VariableDeclContext) Type_() ITypeContext

type VectorLiteralContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyVectorLiteralContext() *VectorLiteralContext

func NewVectorLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorLiteralContext

func (s *VectorLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *VectorLiteralContext) AllCOMMA() []antlr.TerminalNode

func (s *VectorLiteralContext) AllExpression() []IExpressionContext

func (s *VectorLiteralContext) COMMA(i int) antlr.TerminalNode

func (s *VectorLiteralContext) EnterRule(listener antlr.ParseTreeListener)

func (s *VectorLiteralContext) ExitRule(listener antlr.ParseTreeListener)

func (s *VectorLiteralContext) Expression(i int) IExpressionContext

func (s *VectorLiteralContext) GetParser() antlr.Parser

func (s *VectorLiteralContext) GetRuleContext() antlr.RuleContext

func (*VectorLiteralContext) IsVectorLiteralContext()

func (s *VectorLiteralContext) LBRACE() antlr.TerminalNode

func (s *VectorLiteralContext) RBRACE() antlr.TerminalNode

func (s *VectorLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

type VectorTypeContext struct {
	antlr.BaseParserRuleContext
	// Has unexported fields.
}

func NewEmptyVectorTypeContext() *VectorTypeContext

func NewVectorTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorTypeContext

func (s *VectorTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{}

func (s *VectorTypeContext) EnterRule(listener antlr.ParseTreeListener)

func (s *VectorTypeContext) ExitRule(listener antlr.ParseTreeListener)

func (s *VectorTypeContext) GT() antlr.TerminalNode

func (s *VectorTypeContext) GetParser() antlr.Parser

func (s *VectorTypeContext) GetRuleContext() antlr.RuleContext

func (*VectorTypeContext) IsVectorTypeContext()

func (s *VectorTypeContext) LT() antlr.TerminalNode

func (s *VectorTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string

func (s *VectorTypeContext) Type_() ITypeContext

func (s *VectorTypeContext) VECTOR() antlr.TerminalNode

