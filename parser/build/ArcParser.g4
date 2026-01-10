parser grammar ArcParser;

options {
    tokenVocab=ArcLexer;
}

compilationUnit
    : (importDecl | namespaceDecl | topLevelDecl)* EOF
    ;

// =============================================================================
// Declarations
// =============================================================================

importDecl: IMPORT (STRING_LITERAL | LPAREN importSpec* RPAREN);
importSpec: STRING_LITERAL;

namespaceDecl: NAMESPACE IDENTIFIER;

topLevelDecl
    : functionDecl
    | structDecl
    | classDecl
    | enumDecl
    | methodDecl
    | mutatingDecl
    | deinitDecl
    | variableDecl
    | constDecl
    | externCDecl
    | externCppDecl
    ;

// =============================================================================
// Extern C Declarations
// =============================================================================

externCDecl: EXTERN C_LANG LBRACE externCMember* RBRACE;

externCMember
    : externCFunctionDecl
    | externCConstDecl
    | externCTypeAlias
    | externCOpaqueStructDecl
    ;

externCFunctionDecl
    : cCallingConvention? FUNC IDENTIFIER (STRING_LITERAL)?
      LPAREN externCParameterList? RPAREN type?
    ;

cCallingConvention
    : STDCALL
    | CDECL
    | FASTCALL
    ;

externCParameterList
    : externCParameter (COMMA externCParameter)* (COMMA ELLIPSIS)?
    | ELLIPSIS
    ;

externCParameter
    : type
    | IDENTIFIER COLON type
    ;

externCConstDecl
    : CONST IDENTIFIER COLON type ASSIGN expression
    ;

externCTypeAlias
    : TYPE IDENTIFIER ASSIGN type
    ;

externCOpaqueStructDecl
    : OPAQUE STRUCT IDENTIFIER LBRACE RBRACE
    ;

// =============================================================================
// Extern C++ Declarations
// =============================================================================

externCppDecl: EXTERN CPP_LANG LBRACE externCppMember* RBRACE;

externCppMember
    : externCppNamespaceDecl
    | externCppFunctionDecl
    | externCppConstDecl
    | externCppTypeAlias
    | externCppOpaqueClassDecl
    | externCppClassDecl
    ;

externCppNamespaceDecl
    : NAMESPACE externNamespacePath LBRACE externCppMember* RBRACE
    ;

externNamespacePath
    : IDENTIFIER (DOT IDENTIFIER)*
    ;

externCppFunctionDecl
    : cppCallingConvention? FUNC IDENTIFIER (STRING_LITERAL)?
      LPAREN externCppParameterList? RPAREN type?
    ;

cppCallingConvention
    : STDCALL
    | CDECL
    | FASTCALL
    | VECTORCALL
    | THISCALL
    ;

externCppParameterList
    : externCppParameter (COMMA externCppParameter)* (COMMA ELLIPSIS)?
    | ELLIPSIS
    ;

externCppParameter
    : externCppParamType
    | IDENTIFIER COLON externCppParamType
    ;

externCppParamType
    : STAR CONST? type
    | AMP CONST? type
    | CONST? type
    ;

externCppConstDecl
    : CONST IDENTIFIER COLON type ASSIGN expression
    ;

externCppTypeAlias
    : TYPE IDENTIFIER ASSIGN type
    ;

externCppOpaqueClassDecl
    : OPAQUE CLASS IDENTIFIER LBRACE RBRACE
    ;

externCppClassDecl
    : ABSTRACT? CLASS IDENTIFIER (STRING_LITERAL)? LBRACE externCppClassMember* RBRACE
    ;

externCppClassMember
    : externCppConstructorDecl
    | externCppDestructorDecl
    | externCppMethodDecl
    ;

externCppConstructorDecl
    : NEW LPAREN externCppParameterList? RPAREN type
    ;

externCppDestructorDecl
    : DELETE LPAREN externCppSelfParam RPAREN type?
    ;

externCppMethodDecl
    : cppCallingConvention? VIRTUAL? STATIC? FUNC IDENTIFIER (STRING_LITERAL)?
      LPAREN externCppMethodParams RPAREN CONST? type?
    ;

externCppMethodParams
    : externCppSelfParam (COMMA externCppParameter)*
    | externCppParameterList?
    ;

externCppSelfParam
    : SELF STAR CONST? IDENTIFIER
    ;

// =============================================================================
// Generics
// =============================================================================

genericParams: LT genericParamList GT;
genericParamList: genericParam (COMMA genericParam)*;
genericParam: IDENTIFIER (DOT IDENTIFIER)*;

genericArgs: LT genericArgList GT;
genericArgList: genericArg (COMMA genericArg)*;
genericArg: type | expression;

// =============================================================================
// Functions
// =============================================================================

// Hardcoded keywords for concurrency models
functionDecl: (ASYNC | PROCESS | CONTAINER)? FUNC IDENTIFIER genericParams? LPAREN parameterList? RPAREN returnType? block;

returnType: type | LPAREN typeList RPAREN;
typeList: type (COMMA type)*;
parameterList: parameter (COMMA parameter)* (COMMA ELLIPSIS)? | ELLIPSIS;
parameter: SELF? IDENTIFIER COLON type;

// =============================================================================
// Structs
// =============================================================================

structDecl: STRUCT IDENTIFIER computeMarker? genericParams? LBRACE structMember* RBRACE;
computeMarker: LT COMPUTE GT;
structMember: structField | functionDecl | mutatingDecl | initDecl;
structField: IDENTIFIER COLON type;

initDecl: INIT LPAREN SELF IDENTIFIER COLON type (COMMA parameter)* (COMMA ELLIPSIS)? RPAREN block;

// =============================================================================
// Classes
// =============================================================================

classDecl: CLASS IDENTIFIER genericParams? LBRACE classMember* RBRACE;
classMember: classField | functionDecl | initDecl | deinitDecl;
classField: IDENTIFIER COLON type;

// =============================================================================
// Enums
// =============================================================================

enumDecl: ENUM IDENTIFIER (COLON primitiveType)? LBRACE enumMember* RBRACE;
enumMember: IDENTIFIER (ASSIGN expression)?;

// =============================================================================
// Methods
// =============================================================================

methodDecl: ASYNC? FUNC IDENTIFIER genericParams? LPAREN SELF IDENTIFIER COLON type (COMMA parameter)* RPAREN returnType? block;
mutatingDecl: MUTATING IDENTIFIER LPAREN SELF IDENTIFIER COLON type (COMMA parameter)* RPAREN returnType? block;
deinitDecl: DEINIT LPAREN SELF IDENTIFIER COLON type RPAREN block;

// =============================================================================
// Variables
// =============================================================================

variableDecl
    : LET tuplePattern (COLON tupleType)? ASSIGN expression
    | LET IDENTIFIER (COLON type)? ASSIGN expression
    ;
constDecl: CONST IDENTIFIER (COLON type)? ASSIGN expression;
tuplePattern: LPAREN IDENTIFIER (COMMA IDENTIFIER)+ RPAREN;
tupleType: LPAREN typeList RPAREN;

// =============================================================================
// Type System
// =============================================================================

type
    : primitiveType
    | pointerType
    | referenceType
    | qualifiedType
    | functionType
    | arrayType
    | IDENTIFIER genericArgs?
    | UNDERSCORE
    ;

qualifiedType: IDENTIFIER (DOT IDENTIFIER)+ genericArgs?;

functionType: ASYNC? FUNC genericParams? LPAREN typeList? RPAREN returnType?;

arrayType: LBRACKET expression RBRACKET type;

qualifiedIdentifier: IDENTIFIER (DOT IDENTIFIER)+;

primitiveType
    : INT8 | INT16 | INT32 | INT64
    | UINT8 | UINT16 | UINT32 | UINT64
    | USIZE | ISIZE
    | FLOAT32 | FLOAT64
    | FLOAT16 | BFLOAT16
    | BYTE | BOOL | CHAR | STRING | VOID
    ;

pointerType: STAR CONST? type;
referenceType: AMP CONST? type;

// =============================================================================
// Statements
// =============================================================================

block: LBRACE statement* RBRACE;

statement
    : block
    | returnStmt
    | breakStmt
    | continueStmt
    | ifStmt
    | forStmt
    | switchStmt
    | tryStmt
    | throwStmt
    | deferStmt
    | variableDecl
    | constDecl
    | assignmentStmt
    | expressionStmt
    ;

assignmentStmt: leftHandSide assignmentOp expression;
assignmentOp
    : ASSIGN
    | PLUS_ASSIGN
    | MINUS_ASSIGN
    | STAR_ASSIGN
    | SLASH_ASSIGN
    | PERCENT_ASSIGN
    | BIT_OR_ASSIGN
    | BIT_AND_ASSIGN
    | BIT_XOR_ASSIGN
    | LT LE
    | GT GE
    ;
leftHandSide
    : STAR postfixExpression
    | postfixExpression DOT IDENTIFIER
    | postfixExpression LBRACKET expression RBRACKET
    | IDENTIFIER
    ;

expressionStmt: expression;
returnStmt: RETURN tupleExpression | RETURN expression?;
deferStmt: DEFER (assignmentStmt | expression);
breakStmt: BREAK;
continueStmt: CONTINUE;
throwStmt: THROW expression;

ifStmt: IF expression block (ELSE IF expression block)* (ELSE block)?;

forStmt
    : FOR block
    | FOR expression block
    | FOR (variableDecl | assignmentStmt)? SEMICOLON expression? SEMICOLON (assignmentStmt | expression)? block
    | FOR IDENTIFIER IN expression block
    | FOR IDENTIFIER COMMA IDENTIFIER IN expression block
    ;

switchStmt: SWITCH expression LBRACE switchCase* defaultCase? RBRACE;
switchCase: CASE expression (COMMA expression)* COLON statement*;
defaultCase: DEFAULT COLON statement*;

tryStmt: TRY block exceptClause+ finallyClause? | TRY block finallyClause;
exceptClause: EXCEPT qualifiedIdentifier block | EXCEPT IDENTIFIER block;
finallyClause: FINALLY block;

// =============================================================================
// Expressions
// =============================================================================

expression: logicalOrExpression;
logicalOrExpression: logicalAndExpression (OR logicalAndExpression)*;
logicalAndExpression: bitOrExpression (AND bitOrExpression)*;
bitOrExpression: bitXorExpression (BIT_OR bitXorExpression)*;
bitXorExpression: bitAndExpression (BIT_XOR bitAndExpression)*;
bitAndExpression: equalityExpression (AMP equalityExpression)*;
equalityExpression: relationalExpression ((EQ | NE) relationalExpression)*;
relationalExpression: shiftExpression ((LT | LE | GT | GE) shiftExpression)*;
shiftExpression: rangeExpression ((LT LT | GT GT) rangeExpression)*;
rangeExpression: additiveExpression (RANGE additiveExpression)?;
additiveExpression: multiplicativeExpression ((PLUS | MINUS) multiplicativeExpression)*;
multiplicativeExpression: unaryExpression ((STAR | SLASH | PERCENT) unaryExpression)*;

unaryExpression
    : (MINUS | NOT | BIT_NOT | STAR | AMP) unaryExpression
    | AWAIT (LPAREN expression RPAREN)? unaryExpression
    | INCREMENT unaryExpression
    | DECREMENT unaryExpression
    | postfixExpression
    ;

postfixExpression: primaryExpression postfixOp*;
postfixOp
    : DOT IDENTIFIER
    | DOT IDENTIFIER LPAREN argumentList? RPAREN
    | LPAREN argumentList? RPAREN
    | LBRACKET expression RBRACKET
    | INCREMENT
    | DECREMENT
    ;

primaryExpression
    : computeExpression
    | literal
    | structLiteral
    | sizeofExpression
    | alignofExpression
    | lambdaExpression
    | anonymousFuncExpression
    | tupleExpression
    | LPAREN expression RPAREN
    | qualifiedIdentifier genericArgs? (LPAREN argumentList? RPAREN)?
    | IDENTIFIER genericArgs? (LPAREN argumentList? RPAREN)?
    | IDENTIFIER genericArgs?
    | qualifiedIdentifier genericArgs?
    ;

computeExpression
    : computeContext ASYNC? FUNC genericParams? LPAREN parameterList? RPAREN returnType? block LPAREN argumentList? RPAREN
    ;

computeContext
    : qualifiedIdentifier
    | IDENTIFIER
    ;

sizeofExpression: SIZEOF LT type GT;
alignofExpression: ALIGNOF LT type GT;

literal
    : INTEGER_LITERAL
    | FLOAT_LITERAL
    | STRING_LITERAL
    | CHAR_LITERAL
    | BOOLEAN_LITERAL
    | NULL
    | initializerList
    ;

initializerList
    : LBRACE RBRACE
    | LBRACE expression (COMMA expression)* RBRACE
    | LBRACE initializerEntry (COMMA initializerEntry)* RBRACE
    ;
initializerEntry: expression COLON expression;

structLiteral
    : (IDENTIFIER | qualifiedIdentifier) genericArgs? LBRACE (fieldInit (COMMA fieldInit)*)? RBRACE
    ;
fieldInit: IDENTIFIER COLON expression;

argumentList: argument (COMMA argument)*;
argument: expression | lambdaExpression | anonymousFuncExpression;

lambdaExpression
    : (ASYNC | PROCESS | CONTAINER)? LPAREN lambdaParamList? RPAREN FAT_ARROW block
    | (ASYNC | PROCESS | CONTAINER)? LPAREN lambdaParamList? RPAREN FAT_ARROW expression
    ;

// Hardcoded keywords here too
anonymousFuncExpression
    : (ASYNC | PROCESS | CONTAINER)? FUNC genericParams? LPAREN parameterList? RPAREN returnType? block
    ;

lambdaParamList: lambdaParam (COMMA lambdaParam)*;
lambdaParam: IDENTIFIER COLON type | IDENTIFIER;

tupleExpression: LPAREN expression COMMA expression (COMMA expression)* RPAREN;