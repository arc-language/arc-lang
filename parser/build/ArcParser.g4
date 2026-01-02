// ArcParser.g4
parser grammar ArcParser;

options {
    tokenVocab=ArcLexer;
}

// =============================================================================
// Compilation Unit
// =============================================================================

compilationUnit
    : (importDecl | namespaceDecl | topLevelDecl)* EOF
    ;

// =============================================================================
// Import Declaration
// =============================================================================

importDecl
    : IMPORT STRING_LITERAL
    | IMPORT LPAREN importSpec* RPAREN
    ;

importSpec
    : STRING_LITERAL
    ;

// =============================================================================
// Namespace Declaration
// =============================================================================

namespaceDecl
    : NAMESPACE IDENTIFIER
    ;

// =============================================================================
// Top Level Declarations
// =============================================================================

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

// =============================================================================
// Extern Declaration
// =============================================================================

externDecl
    : EXTERN IDENTIFIER LBRACE externMember* RBRACE
    ;

externMember
    : externFunctionDecl
    ;

externFunctionDecl
    : FUNC IDENTIFIER (STRING_LITERAL)? LPAREN externParameterList? RPAREN type?
    ;

externParameterList
    : type (COMMA type)* (COMMA ELLIPSIS)?
    | ELLIPSIS
    ;

// =============================================================================
// Generics
// =============================================================================

genericParams
    : LT genericParamList GT
    ;

genericParamList
    : IDENTIFIER (COMMA IDENTIFIER)*
    ;

genericArgs
    : LT genericArgList GT
    ;

genericArgList
    : genericArg (COMMA genericArg)*
    ;

genericArg
    : type
    | expression
    ;

// =============================================================================
// Function Declaration
// =============================================================================

functionDecl
    : ASYNC? FUNC IDENTIFIER genericParams? LPAREN parameterList? RPAREN returnType? block
    ;

returnType
    : type
    | LPAREN typeList RPAREN
    ;

typeList
    : type (COMMA type)*
    ;

parameterList
    : parameter (COMMA parameter)* (COMMA ELLIPSIS)?
    | ELLIPSIS
    ;

parameter
    : SELF? IDENTIFIER COLON type
    ;

// =============================================================================
// Struct Declaration
// =============================================================================

structDecl
    : STRUCT IDENTIFIER genericParams? LBRACE structMember* RBRACE
    ;

structMember
    : structField
    | functionDecl
    | mutatingDecl
    ;

structField
    : IDENTIFIER COLON type
    ;

// =============================================================================
// Class Declaration
// =============================================================================

classDecl
    : CLASS IDENTIFIER genericParams? LBRACE classMember* RBRACE
    ;

classMember
    : classField
    | functionDecl
    | deinitDecl
    ;

classField
    : IDENTIFIER COLON type
    ;

// =============================================================================
// Enum Declaration
// =============================================================================

enumDecl
    : ENUM IDENTIFIER (COLON primitiveType)? LBRACE enumMember* RBRACE
    ;

enumMember
    : IDENTIFIER (ASSIGN expression)?
    ;

// =============================================================================
// Method Declarations
// =============================================================================

methodDecl
    : ASYNC? FUNC IDENTIFIER genericParams? LPAREN SELF IDENTIFIER COLON type (COMMA parameter)* RPAREN returnType? block
    ;

mutatingDecl
    : MUTATING IDENTIFIER LPAREN SELF IDENTIFIER COLON type (COMMA parameter)* RPAREN returnType? block
    ;

deinitDecl
    : DEINIT LPAREN SELF IDENTIFIER COLON type RPAREN block
    ;

// =============================================================================
// Variable/Constant Declarations
// =============================================================================

variableDecl
    : LET tuplePattern (COLON tupleType)? ASSIGN expression
    | LET IDENTIFIER (COLON type)? ASSIGN expression
    ;

constDecl
    : CONST IDENTIFIER (COLON type)? ASSIGN expression
    ;

tuplePattern
    : LPAREN IDENTIFIER (COMMA IDENTIFIER)+ RPAREN
    ;

tupleType
    : LPAREN typeList RPAREN
    ;

// =============================================================================
// Type System
// =============================================================================

type
    : primitiveType
    | pointerType
    | referenceType
    | qualifiedType            // Matches 'pkg.Type' (Strict 1-dot limit)
    | functionType             // Matches 'func(A) B'
    | IDENTIFIER genericArgs?  // Matches 'Type' or 'Type<T>'
    | UNDERSCORE
    ;

// STRICT RULE: Only allow One Dot for types (Namespace.Type)
// Go-style rule: Users must import "std/io" as 'io', then use 'io.File'
// Nested usage like 'std.io.File' is forbidden.
qualifiedType
    : IDENTIFIER DOT IDENTIFIER genericArgs?
    ;

// Definition for a function TYPE (e.g. for variables/fields)
functionType
    : ASYNC? FUNC genericParams? LPAREN typeList? RPAREN returnType?
    ;

// Used in expressions (not types), allows infinite chaining: a.b.c
qualifiedIdentifier
    : IDENTIFIER (DOT IDENTIFIER)+
    ;

primitiveType
    : INT8 | INT16 | INT32 | INT64
    | UINT8 | UINT16 | UINT32 | UINT64
    | USIZE | ISIZE
    | FLOAT32 | FLOAT64
    | BYTE | BOOL | CHAR
    | STRING
    | VOID
    ;

pointerType
    : STAR type
    ;

referenceType
    : AMP type
    ;

// =============================================================================
// Statements
// =============================================================================

block
    : LBRACE statement* RBRACE
    ;

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

assignmentStmt
    : leftHandSide assignmentOp expression
    ;

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
    | LT LE  // Decomposed <<=
    | GT GE  // Decomposed >>=
    ;

leftHandSide
    : STAR postfixExpression
    | postfixExpression DOT IDENTIFIER
    | postfixExpression LBRACKET expression RBRACKET
    | IDENTIFIER
    ;

expressionStmt
    : expression
    ;

returnStmt
    : RETURN tupleExpression
    | RETURN expression?
    ;

// =============================================================================
// Control Flow - If
// =============================================================================

ifStmt
    : IF expression block (ELSE IF expression block)* (ELSE block)?
    ;

// =============================================================================
// Control Flow - For
// =============================================================================

forStmt
    : FOR block
    | FOR expression block
    | FOR (variableDecl | assignmentStmt)? SEMICOLON expression? SEMICOLON (assignmentStmt | expression)? block
    | FOR IDENTIFIER IN expression block
    | FOR IDENTIFIER COMMA IDENTIFIER IN expression block
    ;

// =============================================================================
// Control Flow - Switch
// =============================================================================

switchStmt
    : SWITCH expression LBRACE switchCase* defaultCase? RBRACE
    ;

switchCase
    : CASE expression COLON statement*
    ;

defaultCase
    : DEFAULT COLON statement*
    ;

// =============================================================================
// Control Flow - Try/Except/Finally
// =============================================================================

tryStmt
    : TRY block exceptClause+ finallyClause?
    | TRY block finallyClause
    ;

exceptClause
    : EXCEPT qualifiedIdentifier block
    | EXCEPT IDENTIFIER block
    ;

finallyClause
    : FINALLY block
    ;

// =============================================================================
// Control Flow - Throw
// =============================================================================

throwStmt
    : THROW expression
    ;

// =============================================================================
// Control Flow - Break/Continue/Defer
// =============================================================================

breakStmt
    : BREAK
    ;

continueStmt
    : CONTINUE
    ;

deferStmt
    : DEFER (assignmentStmt | expression)
    ;

// =============================================================================
// Expressions
// =============================================================================

expression
    : logicalOrExpression
    ;

logicalOrExpression
    : logicalAndExpression (OR logicalAndExpression)*
    ;

logicalAndExpression
    : bitOrExpression (AND bitOrExpression)*
    ;

bitOrExpression
    : bitXorExpression (BIT_OR bitXorExpression)*
    ;

bitXorExpression
    : bitAndExpression (BIT_XOR bitAndExpression)*
    ;

bitAndExpression
    : equalityExpression (AMP equalityExpression)*
    ;

equalityExpression
    : relationalExpression ((EQ | NE) relationalExpression)*
    ;

relationalExpression
    : shiftExpression ((LT | LE | GT | GE) shiftExpression)*
    ;

shiftExpression
    : rangeExpression ((LT LT | GT GT) rangeExpression)*
    ;

rangeExpression
    : additiveExpression (RANGE additiveExpression)?
    ;

additiveExpression
    : multiplicativeExpression ((PLUS | MINUS) multiplicativeExpression)*
    ;

multiplicativeExpression
    : unaryExpression ((STAR | SLASH | PERCENT) unaryExpression)*
    ;

unaryExpression
    : (MINUS | NOT | BIT_NOT | STAR | AMP) unaryExpression
    | AWAIT (LPAREN expression RPAREN)? unaryExpression
    | INCREMENT unaryExpression
    | DECREMENT unaryExpression
    | postfixExpression
    ;

postfixExpression
    : primaryExpression postfixOp*
    ;

postfixOp
    : DOT IDENTIFIER
    | DOT IDENTIFIER LPAREN argumentList? RPAREN
    | LPAREN argumentList? RPAREN
    | LBRACKET expression RBRACKET
    | INCREMENT
    | DECREMENT
    ;

// =============================================================================
// Primary Expressions
// =============================================================================

primaryExpression
    : literal
    | structLiteral
    | sizeofExpression
    | alignofExpression
    | lambdaExpression
    | anonymousFuncExpression
    | tupleExpression
    | LPAREN expression RPAREN
    | qualifiedIdentifier genericArgs? (LPAREN argumentList? RPAREN)? // Allows func calls from namespace
    | IDENTIFIER genericArgs? (LPAREN argumentList? RPAREN)?
    | IDENTIFIER genericArgs?
    | qualifiedIdentifier genericArgs?
    ;

// =============================================================================
// Sizeof/Alignof
// =============================================================================

sizeofExpression
    : SIZEOF LT type GT
    ;

alignofExpression
    : ALIGNOF LT type GT
    ;

// =============================================================================
// Literals
// =============================================================================

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

initializerEntry
    : expression COLON expression
    ;

// =============================================================================
// Struct Literal
// =============================================================================

structLiteral
    : (IDENTIFIER | qualifiedIdentifier) genericArgs? LBRACE (fieldInit (COMMA fieldInit)*)? RBRACE
    ;

fieldInit
    : IDENTIFIER COLON expression
    ;

// =============================================================================
// Arguments
// =============================================================================

argumentList
    : argument (COMMA argument)*
    ;

argument
    : expression
    | lambdaExpression
    | anonymousFuncExpression
    ;

// =============================================================================
// Lambda & Anonymous Function Expressions
// =============================================================================

lambdaExpression
    : ASYNC? LPAREN lambdaParamList? RPAREN FAT_ARROW block
    | ASYNC? LPAREN lambdaParamList? RPAREN FAT_ARROW expression
    ;

anonymousFuncExpression
    : (THREAD | PROCESS | CONTAINER | ASYNC)? FUNC genericParams? LPAREN parameterList? RPAREN returnType? block
    ;

lambdaParamList
    : lambdaParam (COMMA lambdaParam)*
    ;

lambdaParam
    : IDENTIFIER COLON type
    | IDENTIFIER
    ;

// =============================================================================
// Tuple Expression
// =============================================================================

tupleExpression
    : LPAREN expression COMMA expression (COMMA expression)* RPAREN
    ;