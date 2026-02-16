// ArcParser.g4
parser grammar ArcParser;

options {
    tokenVocab = ArcLexer;
}

// =============================================================================
// Top-Level Structure
// =============================================================================
compilationUnit
    : (importDecl | namespaceDecl | topLevelDecl)* EOF
    ;

importDecl
    : IMPORT (IDENTIFIER)? (STRING_LITERAL | LPAREN importSpec* RPAREN)
    ;

importSpec : STRING_LITERAL;

namespaceDecl : NAMESPACE IDENTIFIER;

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
attribute : AT IDENTIFIER (LPAREN expression RPAREN)?;

// =============================================================================
// Extern C
// =============================================================================
externCDecl : EXTERN C_LANG LBRACE externCMember* RBRACE;

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

externCConstDecl   : CONST IDENTIFIER COLON externType ASSIGN expression;
externCTypeAlias   : TYPE IDENTIFIER ASSIGN externType;

externCStructDecl  : STRUCT IDENTIFIER LBRACE externCStructField* RBRACE;
externCStructField : IDENTIFIER COLON externType;

// =============================================================================
// Extern C++
// =============================================================================
externCppDecl : EXTERN CPP_LANG LBRACE externCppMember* RBRACE;

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

externNamespacePath : IDENTIFIER (DOT IDENTIFIER)*;

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
    | AMP  CONST? externType
    | CONST? externType
    ;

externCppConstDecl  : CONST IDENTIFIER COLON externType ASSIGN expression;
externCppTypeAlias  : TYPE IDENTIFIER ASSIGN externType;

externCppClassDecl
    : ABSTRACT? CLASS IDENTIFIER (STRING_LITERAL)?
      LBRACE externCppClassMember* RBRACE
    ;

externCppClassMember
    : externCppConstructorDecl
    | externCppDestructorDecl
    | externCppMethodDecl
    ;

// new(...) return-type  — marks the extern C++ constructor
externCppConstructorDecl
    : NEW LPAREN externCppParameterList? RPAREN externType
    ;

// delete(self * const c)  — marks the extern C++ destructor
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

externCppSelfParam : SELF STAR CONST? IDENTIFIER;

// =============================================================================
// Extern Types
// =============================================================================
externType
    : externPointerType
    | externPrimitiveType
    | externFunctionType
    | IDENTIFIER (DOT IDENTIFIER)*
    ;

externPointerType
    : STAR CONST? externType
    | AMP  CONST? externType
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

externTypeList : externType (COMMA externType)*;

// =============================================================================
// Generics
// =============================================================================
genericParams   : LT genericParamList GT;
genericParamList: genericParam (COMMA genericParam)*;
genericParam    : IDENTIFIER (DOT IDENTIFIER)*;

genericArgs    : LT genericArgList GT;
genericArgList : genericArg (COMMA genericArg)*;
genericArg     : type;

// =============================================================================
// Type System
//
//   []byte         — slice view, ptr + length, no allocation
//   [4096]byte     — fixed-size array (used with new / stack)
//   &var int32     — mutable reference
//   vector[byte]   — owned dynamic array, heap allocated
// =============================================================================
type
    : LBRACKET RBRACKET type                       // []byte   — slice view
    | LBRACKET expression RBRACKET type            // [4096]byte — fixed-size array
    | AMP VAR type                                 // &var int32 — mutable reference
    | primitiveType
    | collectionType
    | qualifiedType
    | functionType
    | RAWPTR
    | IDENTIFIER genericArgs?
    | UNDERSCORE
    ;

collectionType  : IDENTIFIER LBRACKET type RBRACKET type?;
qualifiedType   : IDENTIFIER (DOT IDENTIFIER)+ genericArgs?;
functionType    : executionStrategy? FUNC genericParams? LPAREN typeList? RPAREN returnType?;

primitiveType
    : INT8 | INT16 | INT32 | INT64
    | UINT8 | UINT16 | UINT32 | UINT64
    | USIZE | ISIZE
    | FLOAT32 | FLOAT64
    | BYTE | BOOL | CHAR | STRING | VOID
    ;

typeList        : type (COMMA type)*;
returnType      : type | LPAREN typeList RPAREN;
qualifiedIdentifier : IDENTIFIER (DOT IDENTIFIER)+;

// =============================================================================
// Declarations
// =============================================================================
functionDecl
    : executionStrategy? FUNC IDENTIFIER genericParams?
      LPAREN parameterList? RPAREN returnType? block
    ;

// self is explicit and typed — no magic this
methodDecl
    : executionStrategy? FUNC IDENTIFIER genericParams?
      LPAREN SELF IDENTIFIER COLON type (COMMA parameter)* RPAREN returnType? block
    ;

deinitDecl
    : DEINIT LPAREN SELF IDENTIFIER COLON type RPAREN block
    ;

structDecl
    : attribute* STRUCT IDENTIFIER genericParams? LBRACE structMember* RBRACE
    ;

structMember : structField | functionDecl;
structField  : IDENTIFIER COLON type;

classDecl
    : CLASS IDENTIFIER genericParams? LBRACE classMember* RBRACE
    ;

classMember : classField | functionDecl | deinitDecl;
classField  : IDENTIFIER COLON type;

enumDecl
    : ENUM IDENTIFIER (COLON primitiveType)? LBRACE enumMember* RBRACE
    ;

enumMember : IDENTIFIER (ASSIGN expression)?;

// =============================================================================
// Variable Declarations
//
//   let p = Point{...}    — stack, immutable binding
//   var p = Point{...}    — heap, ref-counted, mutable
//
// The declaration keyword is the memory model. The type is the shape.
// These are orthogonal concerns kept separate.
// =============================================================================
variableDecl
    : LET tuplePattern (COLON tupleType)? ASSIGN expression
    | LET IDENTIFIER (COLON type)? ASSIGN expression
    | LET IDENTIFIER (COLON type)? ASSIGN NULL
    | VAR IDENTIFIER (COLON type)? ASSIGN expression
    | VAR IDENTIFIER (COLON type)? ASSIGN NULL
    ;

constDecl    : CONST IDENTIFIER (COLON type)? ASSIGN expression;
tuplePattern : LPAREN IDENTIFIER (COMMA IDENTIFIER)+ RPAREN;
tupleType    : LPAREN typeList RPAREN;

// =============================================================================
// Parameters
// &var in the type carries all mutability info — no var on the parameter name
// =============================================================================
parameterList
    : parameter (COMMA parameter)* (COMMA ELLIPSIS)?
    | ELLIPSIS
    ;

parameter : SELF? IDENTIFIER COLON type;

// =============================================================================
// Statements
// =============================================================================
block : LBRACE statement* RBRACE;

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

assignmentStmt : unaryExpression assignmentOp expression;

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

expressionStmt : expression;
returnStmt     : RETURN tupleExpression | RETURN expression?;
deferStmt      : DEFER (assignmentStmt | expression);
breakStmt      : BREAK;
continueStmt   : CONTINUE;

ifStmt
    : IF expression block (ELSE IF expression block)* (ELSE block)?
    ;

forStmt
    : FOR block
    | FOR expression block
    | FOR (variableDecl | assignmentStmt)? SEMICOLON expression?
          SEMICOLON (assignmentStmt | expression)? block
    | FOR IDENTIFIER IN expression block
    | FOR IDENTIFIER COMMA IDENTIFIER IN expression block
    ;

switchStmt : SWITCH expression LBRACE switchCase* defaultCase? RBRACE;
switchCase  : CASE expression (COMMA expression)* COLON statement*;
defaultCase : DEFAULT COLON statement*;

// =============================================================================
// Expressions
// =============================================================================
expression           : logicalOrExpression;
logicalOrExpression  : logicalAndExpression (OR  logicalAndExpression)*;
logicalAndExpression : bitOrExpression      (AND bitOrExpression)*;
bitOrExpression      : bitXorExpression     (BIT_OR  bitXorExpression)*;
bitXorExpression     : bitAndExpression     (BIT_XOR bitAndExpression)*;
bitAndExpression     : equalityExpression   (AMP equalityExpression)*;
equalityExpression   : relationalExpression ((EQ | NE) relationalExpression)*;
relationalExpression : shiftExpression      ((LT | LE | GT | GE) shiftExpression)*;
shiftExpression      : rangeExpression      ((LT LT | GT GT) rangeExpression)*;
rangeExpression      : additiveExpression   (RANGE additiveExpression)?;
additiveExpression   : multiplicativeExpression ((PLUS | MINUS) multiplicativeExpression)*;
multiplicativeExpression : unaryExpression  ((STAR | SLASH | PERCENT) unaryExpression)*;

// =============================================================================
// Unary Expressions
//   & in unary position  = take address / pass mutable ref (&val, &n)
//   & in infix position  = bitwise AND (a & b) — disambiguated by grammar position
// =============================================================================
unaryExpression
    : (MINUS | NOT | BIT_NOT) unaryExpression
    | AMP unaryExpression                       // &val — address-of / mutable ref at callsite
    | AWAIT (LPAREN expression RPAREN)? unaryExpression
    | INCREMENT unaryExpression
    | DECREMENT unaryExpression
    | postfixExpression
    ;

postfixExpression : primaryExpression postfixOp*;

postfixOp
    : DOT IDENTIFIER
    | DOT IDENTIFIER LPAREN argumentList? RPAREN
    | LPAREN argumentList? RPAREN
    | LBRACKET expression RBRACKET                         // buf[0]   — single element
    | LBRACKET expression RANGE expression RBRACKET       // buf[0..4] — slice
    | INCREMENT
    | DECREMENT
    ;

primaryExpression
    : newExpression
    | deleteExpression
    | builtinExpression
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
// new — Manual Heap Allocation
//
//   new Node{}               — allocate struct, zero-initialised
//   new Node{value: 10}      — allocate struct with field initialisation
//   new [4096]byte           — allocate fixed-size array
//   new [256]uint32          — allocate fixed-size array of any type
//
// NOT ref-counted. You own it. Always pair with defer delete(...).
// =============================================================================
newExpression
    : NEW (IDENTIFIER | qualifiedIdentifier) genericArgs?
          LBRACE (fieldInit (COMMA fieldInit)*)? RBRACE   // new Node{...}
    | NEW LBRACKET expression RBRACKET type               // new [4096]byte
    ;

// =============================================================================
// delete — Manual Heap Deallocation
//
//   delete(node)   — free a value allocated with new
//   delete(buf)    — free an array allocated with new
//
// Always called via defer: defer delete(node)
// Never triggers deinit — that is only for var.
// =============================================================================
deleteExpression
    : DELETE LPAREN expression RPAREN
    ;

// =============================================================================
// Cast Expression
//   rawptr(x)    — cast value to raw pointer
//   rawptr(&val) — get raw pointer to address of val
//   int32(x)     — numeric cast
// =============================================================================
castExpression  : castTargetType LPAREN expression RPAREN;
castTargetType  : primitiveType | RAWPTR;

// =============================================================================
// Builtin Expression  (@identifier or @identifier(...))
// For compiler-level directives that are NOT intrinsic functions.
// Intrinsics (memptr, sizeof, len, syscall, etc.) resolve as plain identifiers.
// =============================================================================
builtinExpression
    : AT IDENTIFIER LPAREN argumentList? RPAREN
    | AT IDENTIFIER
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
    | LBRACE expression  (COMMA expression)*  RBRACE
    | LBRACE initializerEntry (COMMA initializerEntry)* RBRACE
    ;

initializerEntry : expression COLON expression;

structLiteral
    : (IDENTIFIER | qualifiedIdentifier) genericArgs?
      LBRACE (fieldInit (COMMA fieldInit)*)? RBRACE
    ;

fieldInit : IDENTIFIER COLON expression;

// =============================================================================
// Function Arguments & Lambdas
// =============================================================================
argumentList : argument (COMMA argument)*;
argument     : expression | lambdaExpression | anonymousFuncExpression;

lambdaExpression
    : executionStrategy? LPAREN lambdaParamList? RPAREN FAT_ARROW block
    | executionStrategy? LPAREN lambdaParamList? RPAREN FAT_ARROW expression
    ;

anonymousFuncExpression
    : executionStrategy? FUNC genericParams? LPAREN parameterList? RPAREN returnType? block
    ;

executionStrategy : GPU | ASYNC | PROCESS;

lambdaParamList : lambdaParam (COMMA lambdaParam)*;
lambdaParam     : IDENTIFIER COLON type | IDENTIFIER;

// =============================================================================
// Tuple Expression
// =============================================================================
tupleExpression
    : LPAREN expression COMMA expression (COMMA expression)* RPAREN
    ;