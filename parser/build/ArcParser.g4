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
    | externDecl
    ;

// Extern
externDecl: EXTERN IDENTIFIER LBRACE externMember* RBRACE;
externMember: externFunctionDecl;
externFunctionDecl: FUNC IDENTIFIER (STRING_LITERAL)? LPAREN externParameterList? RPAREN type?;
externParameterList: type (COMMA type)* (COMMA ELLIPSIS)? | ELLIPSIS;

// Generics
genericParams: LT genericParamList GT;
genericParamList: IDENTIFIER (COMMA IDENTIFIER)*;
genericArgs: LT genericArgList GT;
genericArgList: genericArg (COMMA genericArg)*;
genericArg: type | expression;

// Functions
functionDecl: ASYNC? FUNC IDENTIFIER genericParams? LPAREN parameterList? RPAREN returnType? block;
returnType: type | LPAREN typeList RPAREN;
typeList: type (COMMA type)*;
parameterList: parameter (COMMA parameter)* (COMMA ELLIPSIS)? | ELLIPSIS;
parameter: SELF? IDENTIFIER COLON type;

// Structs
structDecl: STRUCT IDENTIFIER genericParams? LBRACE structMember* RBRACE;
structMember: structField | functionDecl | mutatingDecl;
structField: IDENTIFIER COLON type;

// Classes
classDecl: CLASS IDENTIFIER genericParams? LBRACE classMember* RBRACE;
classMember: classField | functionDecl | deinitDecl;
classField: IDENTIFIER COLON type;

// Enums
enumDecl: ENUM IDENTIFIER (COLON primitiveType)? LBRACE enumMember* RBRACE;
enumMember: IDENTIFIER (ASSIGN expression)?;

// Methods
methodDecl: ASYNC? FUNC IDENTIFIER genericParams? LPAREN SELF IDENTIFIER COLON type (COMMA parameter)* RPAREN returnType? block;
mutatingDecl: MUTATING IDENTIFIER LPAREN SELF IDENTIFIER COLON type (COMMA parameter)* RPAREN returnType? block;
deinitDecl: DEINIT LPAREN SELF IDENTIFIER COLON type RPAREN block;

// Variables
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
    | IDENTIFIER genericArgs?
    | UNDERSCORE
    ;

qualifiedType: IDENTIFIER DOT IDENTIFIER genericArgs?;

functionType: ASYNC? FUNC genericParams? LPAREN typeList? RPAREN returnType?;

qualifiedIdentifier: IDENTIFIER (DOT IDENTIFIER)+;

primitiveType
    : INT8 | INT16 | INT32 | INT64
    | UINT8 | UINT16 | UINT32 | UINT64
    | USIZE | ISIZE
    | FLOAT32 | FLOAT64
    | FLOAT16 | BFLOAT16
    | BYTE | BOOL | CHAR | STRING | VOID
    ;

pointerType: STAR type;
referenceType: AMP type;

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
assignmentOp: ASSIGN | PLUS_ASSIGN | MINUS_ASSIGN | STAR_ASSIGN | SLASH_ASSIGN | PERCENT_ASSIGN | BIT_OR_ASSIGN | BIT_AND_ASSIGN | BIT_XOR_ASSIGN | LT LE | GT GE;
leftHandSide: STAR postfixExpression | postfixExpression DOT IDENTIFIER | postfixExpression LBRACKET expression RBRACKET | IDENTIFIER;

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
    : literal
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

sizeofExpression: SIZEOF LT type GT;
alignofExpression: ALIGNOF LT type GT;

literal
    : INTEGER_LITERAL | FLOAT_LITERAL | STRING_LITERAL | CHAR_LITERAL
    | BOOLEAN_LITERAL | NULL | initializerList
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
    : ASYNC? LPAREN lambdaParamList? RPAREN FAT_ARROW block
    | ASYNC? LPAREN lambdaParamList? RPAREN FAT_ARROW expression
    ;

anonymousFuncExpression
    : (THREAD | PROCESS | CONTAINER | ASYNC)? FUNC genericParams? LPAREN parameterList? RPAREN returnType? block
    ;

lambdaParamList: lambdaParam (COMMA lambdaParam)*;
lambdaParam: IDENTIFIER COLON type | IDENTIFIER;

tupleExpression: LPAREN expression COMMA expression (COMMA expression)* RPAREN;