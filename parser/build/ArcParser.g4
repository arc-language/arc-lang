parser grammar ArcParser;

options { tokenVocab = ArcLexer; }

// ═══════════════════════════════════════════════
//  Top Level
// ═══════════════════════════════════════════════

compilationUnit
    : namespaceDecl semi topLevelDecl* EOF
    ;

namespaceDecl
    : NAMESPACE qualifiedName
    ;

topLevelDecl
    : importDecl        semi
    | constDecl         semi
    | topLevelVarDecl   semi
    | topLevelLetDecl   semi
    | funcDecl
    | deinitDecl
    | attribute* interfaceDecl
    | enumDecl
    | typeAliasDecl     semi
    | externDecl
    ;

semi
    : SEMI*
    ;

// ═══════════════════════════════════════════════
//  Imports
// ═══════════════════════════════════════════════

importDecl
    : IMPORT importSpec
    | IMPORT LPAREN importSpec+ RPAREN
    ;

importSpec
    : importAlias? STRING_LIT semi
    ;

importAlias
    : IDENTIFIER        // named import
    | UNDERSCORE        // blank import
    | DOT               // dot import
    ;

// ═══════════════════════════════════════════════
//  Constants
// ═══════════════════════════════════════════════

constDecl
    : CONST constSpec
    | CONST LPAREN constSpec+ RPAREN
    ;

constSpec
    : IDENTIFIER (COLON typeRef)? ASSIGN expression semi
    ;

// ═══════════════════════════════════════════════
//  Variables (top-level)
// ═══════════════════════════════════════════════

topLevelVarDecl
    : VAR IDENTIFIER (COLON typeRef)? ASSIGN expression
    | VAR IDENTIFIER COLON typeRef ASSIGN NULL
    ;

topLevelLetDecl
    : LET IDENTIFIER (COLON typeRef)? ASSIGN expression
    ;

// ═══════════════════════════════════════════════
//  Functions
// ═══════════════════════════════════════════════

funcDecl
    : funcModifier* FUNC IDENTIFIER genericParams? LPAREN paramList? RPAREN returnType? block
    ;

funcModifier
    : ASYNC
    | GPU
    ;

deinitDecl
    : DEINIT LPAREN selfParam RPAREN block
    ;

paramList
    : param (COMMA param)* COMMA?
    ;

param
    : selfParam
    | IDENTIFIER COLON paramType
    | ELLIPSIS                          // variadic (arc-level)
    ;

selfParam
    : SELF IDENTIFIER COLON paramType                    // self by value
    | SELF AMP MUT IDENTIFIER COLON paramType            // self by mutable ref
    ;

paramType
    : AMP MUT typeRef                   // &mut T
    | typeRef
    ;

returnType
    : tupleType
    | typeRef
    ;

tupleType
    : LPAREN typeRef (COMMA typeRef)+ RPAREN
    ;

genericParams
    : LBRACKET IDENTIFIER (COMMA IDENTIFIER)* RBRACKET
    ;

genericArgs
    : LBRACKET typeRef (COMMA typeRef)* RBRACKET
    ;

// ═══════════════════════════════════════════════
//  Interfaces
// ═══════════════════════════════════════════════

interfaceDecl
    : INTERFACE IDENTIFIER genericParams? LBRACE interfaceField* RBRACE
    ;

interfaceField
    : IDENTIFIER COLON typeRef semi
    ;

// ═══════════════════════════════════════════════
//  Enums
// ═══════════════════════════════════════════════

enumDecl
    : ENUM IDENTIFIER (COLON primitiveType)? LBRACE enumMember+ RBRACE
    ;

enumMember
    : IDENTIFIER (ASSIGN expression)? semi
    ;

// ═══════════════════════════════════════════════
//  Type Alias / Opaque
// ═══════════════════════════════════════════════

typeAliasDecl
    : TYPE IDENTIFIER ASSIGN OPAQUE
    | TYPE IDENTIFIER ASSIGN typeRef
    ;

// ═══════════════════════════════════════════════
//  Attributes
// ═══════════════════════════════════════════════

attribute
    : AT IDENTIFIER (LPAREN expression RPAREN)?         // @packed, @align(16)
    ;

// ═══════════════════════════════════════════════
//  Types
// ═══════════════════════════════════════════════

typeRef
    : functionType
    | baseType
    ;

functionType
    : ASYNC? FUNC LPAREN typeList? RPAREN typeRef?      // async func(T,U) V
    ;

baseType
    : primitiveType
    | VOID
    | BOOL
    | STRING
    | BYTE
    | CHAR
    | qualifiedName genericArgs?                        // Point, Box[int32], net.Socket
    | VECTOR LBRACKET typeRef RBRACKET                  // vector[T]
    | MAP LBRACKET typeRef RBRACKET typeRef             // map[K]V
    | LBRACKET RBRACKET typeRef                         // []T   (slice)
    | LBRACKET expression RBRACKET typeRef              // [N]T  (fixed array)
    ;

primitiveType
    : INT8 | INT16 | INT32 | INT64
    | UINT8 | UINT16 | UINT32 | UINT64
    | USIZE | ISIZE
    | FLOAT32 | FLOAT64
    ;

typeList
    : typeRef (COMMA typeRef)*
    ;

// ═══════════════════════════════════════════════
//  Extern Blocks
// ═══════════════════════════════════════════════

externDecl
    : EXTERN IDENTIFIER LBRACE externMember* RBRACE     // extern c { ... } / extern cpp { ... }
    ;

externMember
    : externFuncDecl        semi
    | externTypeAlias       semi
    | externNamespace
    | externClass
    ;

// ── Extern Functions ─────────────────────────

externFuncDecl
    : callingConvention? FUNC IDENTIFIER externSymbol? LPAREN externParamList? RPAREN externReturnType?
    ;

callingConvention
    : CDECL | STDCALL | THISCALL | VECTORCALL | FASTCALL
    ;

externSymbol
    : STRING_LIT                                        // "printf", "?Clamp@Math@@..."
    ;

externParamList
    : externParam (COMMA externParam)* (COMMA ELLIPSIS)?
    | ELLIPSIS
    ;

externParam
    : externType
    ;

externReturnType
    : CONST? externType                                 // const for C++ const methods
    ;

externType
    : STAR STAR externType                              // **T
    | STAR CONST? externType                            // *T, *const T
    | AMP CONST? externType                             // &T, &const T (C++ refs)
    | primitiveType
    | VOID
    | BOOL
    | STRING
    | BYTE
    | CHAR
    | USIZE | ISIZE
    | qualifiedName
    | LBRACKET expression RBRACKET externType           // [N]T
    ;

// ── Extern Namespaces ────────────────────────

externNamespace
    : NAMESPACE qualifiedName LBRACE externMember* RBRACE
    ;

// ── Extern Classes ───────────────────────────

externClass
    : ABSTRACT? CLASS IDENTIFIER externSymbol? LBRACE externClassMember* RBRACE
    ;

externClassMember
    : externVirtualMethod   semi
    | externStaticMethod    semi
    | externConstructor     semi
    | externDestructor      semi
    ;

externVirtualMethod
    : callingConvention? VIRTUAL FUNC IDENTIFIER
      LPAREN externMethodParamList? RPAREN externReturnType?
    ;

externStaticMethod
    : STATIC FUNC IDENTIFIER externSymbol?
      LPAREN externParamList? RPAREN externReturnType?
    ;

externConstructor
    : NEW LPAREN externParamList? RPAREN externType     // new(...) *ClassName
    ;

externDestructor
    : DELETE LPAREN externMethodParam RPAREN VOID?      // delete(self *T) void
    ;

externMethodParamList
    : externMethodParam (COMMA externParam)* (COMMA ELLIPSIS)?
    ;

externMethodParam
    : SELF externType                                   // self *ClassName, self *const ClassName
    ;

// ── Extern Type Alias (inside extern block) ──

externTypeAlias
    : TYPE IDENTIFIER ASSIGN externFunctionPtrType
    ;

externFunctionPtrType
    : FUNC LPAREN externParamList? RPAREN externReturnType?
    ;

// ═══════════════════════════════════════════════
//  Statements
// ═══════════════════════════════════════════════

block
    : LBRACE statement* RBRACE
    ;

statement
    : letStatement          semi
    | varStatement          semi
    | constDecl             semi
    | returnStatement       semi
    | breakStatement        semi
    | continueStatement     semi
    | deferStatement        semi
    | ifStatement
    | forStatement
    | switchStatement
    | assignmentStatement   semi
    | expressionStatement   semi
    ;

letStatement
    : LET LPAREN IDENTIFIER (COMMA IDENTIFIER)+ RPAREN ASSIGN expression    // destructuring
    | LET IDENTIFIER (COLON typeRef)? ASSIGN expression
    ;

varStatement
    : VAR IDENTIFIER (COLON typeRef)? ASSIGN expression
    | VAR IDENTIFIER COLON typeRef ASSIGN NULL
    ;

returnStatement
    : RETURN expression?
    | RETURN LPAREN expression (COMMA expression)+ RPAREN     // return (a, b)
    ;

breakStatement
    : BREAK
    ;

continueStatement
    : CONTINUE
    ;

deferStatement
    : DEFER expression                                  // defer delete(x), defer close(fd)
    ;

ifStatement
    : IF expression block (ELSE IF expression block)* (ELSE block)?
    ;

forStatement
    : FOR forHeader block
    ;

forHeader
    : forInit SEMI expression SEMI forPost              // C-style
    | forIterator                                       // for-in
    | expression                                        // while-style
    |                                                   // infinite
    ;

forInit
    : LET IDENTIFIER (COLON typeRef)? ASSIGN expression
    | expression
    ;

forPost
    : expression
    | assignmentTarget assignOp expression
    | expression (INC | DEC)
    ;

forIterator
    : IDENTIFIER (COMMA IDENTIFIER)? IN expression      // for item in items / for k, v in map
    ;

switchStatement
    : SWITCH expression LBRACE switchCase* switchDefault? RBRACE
    ;

switchCase
    : CASE expressionList COLON statement*
    ;

switchDefault
    : DEFAULT COLON statement*
    ;

expressionList
    : expression (COMMA expression)*
    ;

assignmentStatement
    : assignmentTarget assignOp expression
    | expression (INC | DEC)
    ;

assignmentTarget
    : expression DOT IDENTIFIER                         // a.b
    | expression LBRACKET expression RBRACKET           // a[i]
    | IDENTIFIER                                        // x
    ;

assignOp
    : ASSIGN
    | ADD_ASSIGN | SUB_ASSIGN | MUL_ASSIGN | DIV_ASSIGN | MOD_ASSIGN
    | AND_ASSIGN | OR_ASSIGN | XOR_ASSIGN
    | SHL_ASSIGN | SHR_ASSIGN
    ;

expressionStatement
    : expression
    ;

// ═══════════════════════════════════════════════
//  Expressions (precedence by alternative order)
// ═══════════════════════════════════════════════

expression
    : primary                                                           # PrimaryExpr

    // ── Postfix ──
    | expression DOT IDENTIFIER                                         # MemberAccess
    | expression LBRACKET expression RBRACKET                           # IndexExpr
    | expression LBRACKET expression RANGE expression RBRACKET          # SliceExpr
    | expression LPAREN argumentList? RPAREN                            # CallExpr
    | expression INC                                                    # PostIncrement
    | expression DEC                                                    # PostDecrement

    // ── Unary prefix ──
    | MINUS expression                                                  # UnaryMinus
    | BANG expression                                                   # LogicalNot
    | TILDE expression                                                  # BitwiseNot
    | AMP expression                                                    # AddressOf
    | AWAIT expression                                                  # AwaitExpr

    // ── Binary (tightest to loosest) ──
    | expression op=(STAR | SLASH | PERCENT) expression                 # MulExpr
    | expression op=(PLUS | MINUS) expression                           # AddExpr
    | expression op=(LSHIFT | RSHIFT) expression                        # ShiftExpr
    | expression op=(LT | GT | LE | GE) expression                     # RelationalExpr
    | expression op=(EQ | NEQ) expression                               # EqualityExpr
    | expression AMP expression                                         # BitwiseAndExpr
    | expression CARET expression                                       # BitwiseXorExpr
    | expression PIPE expression                                        # BitwiseOrExpr
    | expression AND expression                                         # LogicalAndExpr
    | expression OR expression                                          # LogicalOrExpr

    // ── Range (for `0..10`) ──
    | expression RANGE expression                                       # RangeExpr
    ;

primary
    : INT_LIT                                                           # IntLiteral
    | HEX_LIT                                                           # HexLiteral
    | FLOAT_LIT                                                         # FloatLiteral
    | STRING_LIT                                                        # StringLiteral
    | CHAR_LIT                                                          # CharLiteral
    | TRUE                                                              # TrueLiteral
    | FALSE                                                             # FalseLiteral
    | NULL                                                              # NullLiteral

    | IDENTIFIER                                                        # IdentExpr
    | qualifiedName                                                     # QualifiedExpr

    // ── Type in expression position (casts, sizeof, etc.) ──
    | primitiveType                                                     # PrimitiveTypeExpr

    // ── Parenthesized / tuple ──
    | LPAREN expression RPAREN                                          # ParenExpr
    | LPAREN expression (COMMA expression)+ RPAREN                      # TupleLiteral

    // ── New ──
    | NEW typeRef initializerBlock                                      # NewExpr
    | NEW LBRACKET expression RBRACKET typeRef                          # NewArrayExpr

    // ── Delete ──
    | DELETE LPAREN expression RPAREN                                   # DeleteExpr

    // ── Collection / interface initializer literals ──
    | qualifiedName genericArgs? initializerBlock                       # TypedInitExpr
    | VECTOR LBRACKET typeRef RBRACKET initializerBlock                 # VectorLiteral
    | MAP LBRACKET typeRef RBRACKET typeRef initializerBlock            # MapLiteral

    // ── Lambda ──
    | ASYNC? LPAREN lambdaParamList? RPAREN ARROW block                 # LambdaExpr

    // ── Process ──
    | PROCESS FUNC LPAREN paramList? RPAREN returnType? block
      LPAREN argumentList? RPAREN                                       # ProcessExpr
    ;

// ── Initializer block { field: val, ... } or { val, ... } or { k: v, ... } ──

initializerBlock
    : LBRACE RBRACE
    | LBRACE fieldInit (COMMA fieldInit)* COMMA? RBRACE
    | LBRACE expression (COMMA expression)* COMMA? RBRACE
    | LBRACE mapEntry (COMMA mapEntry)* COMMA? RBRACE
    ;

fieldInit
    : IDENTIFIER COLON expression
    ;

mapEntry
    : expression COLON expression
    ;

// ── Arguments ────────────────────────────────

argumentList
    : argument (COMMA argument)*
    ;

argument
    : expression
    ;

// ── Lambda parameters ────────────────────────

lambdaParamList
    : lambdaParam (COMMA lambdaParam)*
    ;

lambdaParam
    : IDENTIFIER COLON typeRef
    ;

// ═══════════════════════════════════════════════
//  Shared
// ═══════════════════════════════════════════════

qualifiedName
    : IDENTIFIER (DOT IDENTIFIER)+
    ;