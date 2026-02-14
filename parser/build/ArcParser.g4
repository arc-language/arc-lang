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

importDecl: IMPORT (IDENTIFIER)? (STRING_LITERAL | LPAREN importSpec* RPAREN);
importSpec: STRING_LITERAL;

namespaceDecl: NAMESPACE IDENTIFIER;

topLevelDecl
    : functionDecl
    | structDecl
    | classDecl
    | enumDecl
    | methodDecl
    | deinitDecl
    | variableDecl
    | constDecl
    | externCDecl
    | externCppDecl
    ;

// =============================================================================
// Attributes
// =============================================================================

attribute: AT IDENTIFIER (LPAREN expression RPAREN)?;

// =============================================================================
// Extern C Declarations
// =============================================================================

externCDecl: EXTERN C_LANG LBRACE externCMember* RBRACE;

externCMember
    : externCFunctionDecl
    | externCConstDecl
    | externCTypeAlias
    | externCStructDecl
    ;

externCFunctionDecl
    : cCallingConvention? FUNC IDENTIFIER (STRING_LITERAL)?
      LPAREN externCParameterList? RPAREN externType?
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
    : externType
    | IDENTIFIER COLON externType
    ;

externCConstDecl
    : CONST IDENTIFIER COLON externType ASSIGN expression
    ;

externCTypeAlias
    : TYPE IDENTIFIER ASSIGN externType
    ;

externCStructDecl
    : STRUCT IDENTIFIER LBRACE externCStructField* RBRACE
    ;

externCStructField
    : IDENTIFIER COLON externType
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
      LPAREN externCppParameterList? RPAREN externType?
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
    : STAR CONST? externType
    | AMP CONST? externType
    | CONST? externType
    ;

externCppConstDecl
    : CONST IDENTIFIER COLON externType ASSIGN expression
    ;

externCppTypeAlias
    : TYPE IDENTIFIER ASSIGN externType
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
    : NEW LPAREN externCppParameterList? RPAREN externType
    ;

externCppDestructorDecl
    : DELETE LPAREN externCppSelfParam RPAREN externType?
    ;

externCppMethodDecl
    : cppCallingConvention? VIRTUAL? STATIC? FUNC IDENTIFIER (STRING_LITERAL)?
      LPAREN externCppMethodParams RPAREN CONST? externType?
    ;

externCppMethodParams
    : externCppSelfParam (COMMA externCppParameter)*
    | externCppParameterList?
    ;

externCppSelfParam
    : SELF STAR CONST? IDENTIFIER
    ;

// =============================================================================
// Extern Type System (stars live here only)
// =============================================================================

externType
    : externPointerType
    | externPrimitiveType
    | externFunctionType
    | IDENTIFIER (DOT IDENTIFIER)*
    ;

externPointerType
    : STAR CONST? externType
    | AMP CONST? externType
    ;

externPrimitiveType
    : INT8 | INT16 | INT32 | INT64
    | UINT8 | UINT16 | UINT32 | UINT64
    | USIZE | ISIZE
    | FLOAT32 | FLOAT64
    | BYTE | BOOL | CHAR | STRING | VOID
    ;

externFunctionType
    : FUNC LPAREN externTypeList? RPAREN externType?
    ;

externTypeList: externType (COMMA externType)*;

// =============================================================================
// Generics (structs and functions only)
// =============================================================================

genericParams: LT genericParamList GT;
genericParamList: genericParam (COMMA genericParam)*;
genericParam: IDENTIFIER (DOT IDENTIFIER)*;

genericArgs: LT genericArgList GT;
genericArgList: genericArg (COMMA genericArg)*;
genericArg: type;

// =============================================================================
// Collection Types (map, vector, set use [key]value syntax)
// =============================================================================

collectionType
    : IDENTIFIER LBRACKET type RBRACKET type?
    ;

// map[string]int32  -> IDENTIFIER[keyType]valueType
// vector[int32]     -> IDENTIFIER[elemType]
// set[string]       -> IDENTIFIER[elemType]

// =============================================================================
// Functions
// =============================================================================

functionDecl
    : executionStrategy? FUNC IDENTIFIER genericParams?
      LPAREN parameterList? RPAREN returnType? block
    ;

returnType: type | LPAREN typeList RPAREN;
typeList: type (COMMA type)*;
parameterList: parameter (COMMA parameter)* (COMMA ELLIPSIS)? | ELLIPSIS;
parameter: SELF? IDENTIFIER COLON type;

// =============================================================================
// Structs
// =============================================================================

structDecl: attribute* STRUCT IDENTIFIER genericParams? LBRACE structMember* RBRACE;
structMember: structField | functionDecl;
structField: IDENTIFIER COLON type;

// =============================================================================
// Classes
// =============================================================================

classDecl: CLASS IDENTIFIER genericParams? LBRACE classMember* RBRACE;
classMember: classField | functionDecl | deinitDecl;
classField: IDENTIFIER COLON type;

// =============================================================================
// Enums
// =============================================================================

enumDecl: ENUM IDENTIFIER (COLON primitiveType)? LBRACE enumMember* RBRACE;
enumMember: IDENTIFIER (ASSIGN expression)?;

// =============================================================================
// Methods
// =============================================================================

methodDecl
    : executionStrategy? FUNC IDENTIFIER genericParams?
      LPAREN SELF IDENTIFIER COLON type (COMMA parameter)* RPAREN returnType? block
    ;

deinitDecl: DEINIT LPAREN SELF IDENTIFIER COLON type RPAREN block;

// =============================================================================
// Variables
// =============================================================================

variableDecl
    : LET tuplePattern (COLON tupleType)? ASSIGN expression
    | LET IDENTIFIER (COLON type)? ASSIGN expression
    | LET IDENTIFIER (COLON type)? ASSIGN NULL
    ;

constDecl: CONST IDENTIFIER (COLON type)? ASSIGN expression;
tuplePattern: LPAREN IDENTIFIER (COMMA IDENTIFIER)+ RPAREN;
tupleType: LPAREN typeList RPAREN;

// =============================================================================
// Type System (no stars, no arrays, no opaque)
// =============================================================================

type
    : primitiveType
    | collectionType
    | qualifiedType
    | functionType
    | RAWPTR
    | IDENTIFIER genericArgs?
    | UNDERSCORE
    ;

qualifiedType: IDENTIFIER (DOT IDENTIFIER)+ genericArgs?;

functionType: executionStrategy? FUNC genericParams? LPAREN typeList? RPAREN returnType?;

qualifiedIdentifier: IDENTIFIER (DOT IDENTIFIER)+;

primitiveType
    : INT8 | INT16 | INT32 | INT64
    | UINT8 | UINT16 | UINT32 | UINT64
    | USIZE | ISIZE
    | FLOAT32 | FLOAT64
    | BYTE | BOOL | CHAR | STRING | VOID
    ;

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
    | deferStmt
    | variableDecl
    | constDecl
    | assignmentStmt
    | expressionStmt
    ;

assignmentStmt: unaryExpression assignmentOp expression;

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
    | SHL_ASSIGN
    | SHR_ASSIGN
    ;

expressionStmt: expression;
returnStmt: RETURN tupleExpression | RETURN expression?;
deferStmt: DEFER (assignmentStmt | expression);
breakStmt: BREAK;
continueStmt: CONTINUE;

ifStmt: IF expression block (ELSE IF expression block)* (ELSE block)?;

forStmt
    : FOR block
    | FOR expression block
    | FOR (variableDecl | assignmentStmt)? SEMICOLON expression?
      SEMICOLON (assignmentStmt | expression)? block
    | FOR IDENTIFIER IN expression block
    | FOR IDENTIFIER COMMA IDENTIFIER IN expression block
    ;

switchStmt: SWITCH expression LBRACE switchCase* defaultCase? RBRACE;
switchCase: CASE expression (COMMA expression)* COLON statement*;
defaultCase: DEFAULT COLON statement*;

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
    : (MINUS | NOT | BIT_NOT | AMP) unaryExpression
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
    : builtinExpression
    | castExpression
    | literal
    | structLiteral
    | lambdaExpression
    | anonymousFuncExpression
    | tupleExpression
    | LPAREN expression RPAREN
    | qualifiedIdentifier genericArgs? (LPAREN argumentList? RPAREN)?
    | IDENTIFIER genericArgs? (LPAREN argumentList? RPAREN)?
    | IDENTIFIER genericArgs?
    | qualifiedIdentifier genericArgs?
    ;

// =============================================================================
// Cast Expression
// type(value) style: int32(3.14), float64(x), rawptr(-1)
// =============================================================================

castExpression
    : castTargetType LPAREN expression RPAREN
    ;

castTargetType
    : primitiveType
    | RAWPTR
    ;

// =============================================================================
// Compiler Builtins
// sizeof, alignof, memset etc are intrinsics not grammar keywords
// =============================================================================

builtinExpression
    : AT IDENTIFIER LPAREN argumentList? RPAREN
    | AT IDENTIFIER
    ;

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
    : (IDENTIFIER | qualifiedIdentifier) genericArgs?
      LBRACE (fieldInit (COMMA fieldInit)*)? RBRACE
    ;

fieldInit: IDENTIFIER COLON expression;

argumentList: argument (COMMA argument)*;
argument: expression | lambdaExpression | anonymousFuncExpression;

lambdaExpression
    : executionStrategy? LPAREN lambdaParamList? RPAREN FAT_ARROW block
    | executionStrategy? LPAREN lambdaParamList? RPAREN FAT_ARROW expression
    ;

anonymousFuncExpression
    : executionStrategy? FUNC genericParams? LPAREN parameterList? RPAREN returnType? block
    ;

// =============================================================================
// Execution Strategy
// =============================================================================

executionStrategy
    : GPU
    | ASYNC
    | PROCESS
    ;

lambdaParamList: lambdaParam (COMMA lambdaParam)*;
lambdaParam: IDENTIFIER COLON type | IDENTIFIER;

tupleExpression
    : LPAREN expression COMMA expression (COMMA expression)* RPAREN
    ;