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
    : NAMESPACE (IDENTIFIER | SYSCALL)
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
    : LT typeList GT
    ;

typeList
    : type (COMMA type)*
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
    | arrayType
    | qualifiedType
    | IDENTIFIER genericArgs?
    | UNDERSCORE
    ;

qualifiedType
    : (IDENTIFIER | SYSCALL) (DOT IDENTIFIER)+ genericArgs?
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

arrayType
    : ARRAY LT type COMMA arraySize GT
    ;

arraySize
    : INTEGER_LITERAL
    | IDENTIFIER
    | UNDERSCORE
    ;

// =============================================================================
// Statements
// =============================================================================

block
    : LBRACE statement* RBRACE
    ;

statement
    : variableDecl
    | constDecl
    | assignmentStmt
    | expressionStmt
    | returnStmt
    | ifStmt
    | forStmt
    | switchStmt
    | tryStmt
    | throwStmt
    | breakStmt
    | continueStmt
    | deferStmt
    | block
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
    : rangeExpression ((LT LT | GT GT) rangeExpression)* // Decomposed << and >>
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
    : (MINUS | NOT | BIT_NOT | STAR | AMP | AWAIT) unaryExpression
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
    | castExpression
    | allocaExpression
    | syscallExpression
    | intrinsicExpression
    | lambdaExpression
    | tupleExpression
    | LPAREN expression RPAREN
    | qualifiedIdentifier genericArgs?
    | IDENTIFIER genericArgs?
    ;

qualifiedIdentifier
    : (IDENTIFIER | SYSCALL) (DOT IDENTIFIER)+
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

// Generic initializer list - used for arrays, vectors, maps, structs
// Compiler determines actual type from context
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
    ;

// =============================================================================
// Lambda Expression
// =============================================================================

lambdaExpression
    : ASYNC? LPAREN lambdaParamList? RPAREN FAT_ARROW block
    | ASYNC? LPAREN lambdaParamList? RPAREN FAT_ARROW expression
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

// =============================================================================
// Cast Expression
// =============================================================================

castExpression
    : CAST LT type GT LPAREN expression RPAREN
    ;

// =============================================================================
// Alloca Expression
// =============================================================================

allocaExpression
    : ALLOCA LPAREN type (COMMA expression)? RPAREN
    ;

// =============================================================================
// Syscall Expression
// =============================================================================

syscallExpression
    : SYSCALL LPAREN expression (COMMA expression)* RPAREN
    ;

// =============================================================================
// Intrinsic Expressions
// =============================================================================

intrinsicExpression
    : SIZEOF LT type GT
    | ALIGNOF LT type GT
    | MEMSET LPAREN expression COMMA expression COMMA expression RPAREN
    | MEMCPY LPAREN expression COMMA expression COMMA expression RPAREN
    | MEMMOVE LPAREN expression COMMA expression COMMA expression RPAREN
    | STRLEN LPAREN expression RPAREN
    | MEMCHR LPAREN expression COMMA expression COMMA expression RPAREN
    | VA_START LPAREN IDENTIFIER RPAREN
    | VA_ARG LT type GT LPAREN expression RPAREN
    | VA_END LPAREN expression RPAREN
    | RAISE LPAREN expression RPAREN
    | MEMCMP LPAREN expression COMMA expression COMMA expression RPAREN
    | BIT_CAST LT type GT LPAREN expression RPAREN
    | SLICE LPAREN expression COMMA expression RPAREN
    ;