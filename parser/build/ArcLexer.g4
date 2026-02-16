lexer grammar ArcLexer;

// ─────────────────────────────────────────────
//  NOTE: This grammar expects a post-lexer semicolon insertion pass
//  (similar to Go). After lexing, a token filter inserts synthetic SEMI
//  tokens at newlines when the preceding token is one of:
//    IDENTIFIER, INT_LIT, FLOAT_LIT, STRING_LIT, CHAR_LIT,
//    TRUE, FALSE, NULL, RETURN, BREAK, CONTINUE,
//    RPAREN, RBRACKET, RBRACE, INC, DEC,
//    or any primitive type keyword.
//  The parser consumes SEMI as an optional statement terminator.
// ─────────────────────────────────────────────

// ── Keywords: structure ──────────────────────

NAMESPACE   : 'namespace';
IMPORT      : 'import';
FUNC        : 'func';
ASYNC       : 'async';
GPU         : 'gpu';
INTERFACE   : 'interface';
ENUM        : 'enum';
CONST       : 'const';
LET         : 'let';
VAR         : 'var';
NEW         : 'new';
DELETE      : 'delete';
DEFER       : 'defer';
DEINIT      : 'deinit';
RETURN      : 'return';
IF          : 'if';
ELSE        : 'else';
FOR         : 'for';
IN          : 'in';
SWITCH      : 'switch';
CASE        : 'case';
DEFAULT     : 'default';
BREAK       : 'break';
CONTINUE    : 'continue';
PROCESS     : 'process';
AWAIT       : 'await';
EXTERN      : 'extern';
TYPE        : 'type';
OPAQUE      : 'opaque';
SELF        : 'self';
MUT         : 'mut';
VOID        : 'void';

// ── Keywords: literals ───────────────────────

NULL        : 'null';
TRUE        : 'true';
FALSE       : 'false';

// ── Keywords: extern ─────────────────────────

CLASS       : 'class';
VIRTUAL     : 'virtual';
STATIC      : 'static';
ABSTRACT    : 'abstract';

// ── Keywords: calling conventions ────────────

CDECL       : 'cdecl';
STDCALL     : 'stdcall';
THISCALL    : 'thiscall';
VECTORCALL  : 'vectorcall';
FASTCALL    : 'fastcall';

// ── Keywords: primitive types ────────────────

INT8        : 'int8';
INT16       : 'int16';
INT32       : 'int32';
INT64       : 'int64';
UINT8       : 'uint8';
UINT16      : 'uint16';
UINT32      : 'uint32';
UINT64      : 'uint64';
USIZE       : 'usize';
ISIZE       : 'isize';
FLOAT32     : 'float32';
FLOAT64     : 'float64';
BYTE        : 'byte';
BOOL        : 'bool';
CHAR        : 'char';
STRING      : 'string';

// ── Keywords: collection types ───────────────

VECTOR      : 'vector';
MAP         : 'map';

// ── Operators: multi-char (order matters) ────

ARROW       : '=>';
ELLIPSIS    : '...';
RANGE       : '..';
LSHIFT      : '<<';
RSHIFT      : '>>';
LE          : '<=';
GE          : '>=';
EQ          : '==';
NEQ         : '!=';
AND         : '&&';
OR          : '||';
INC         : '++';
DEC         : '--';
ADD_ASSIGN  : '+=';
SUB_ASSIGN  : '-=';
MUL_ASSIGN  : '*=';
DIV_ASSIGN  : '/=';
MOD_ASSIGN  : '%=';
AND_ASSIGN  : '&=';
OR_ASSIGN   : '|=';
XOR_ASSIGN  : '^=';
SHL_ASSIGN  : '<<=';
SHR_ASSIGN  : '>>=';

// ── Operators: single-char ───────────────────

LPAREN      : '(';
RPAREN      : ')';
LBRACKET    : '[';
RBRACKET    : ']';
LBRACE      : '{';
RBRACE      : '}';
DOT         : '.';
COMMA       : ',';
COLON       : ':';
SEMI        : ';';
ASSIGN      : '=';
PLUS        : '+';
MINUS       : '-';
STAR        : '*';
SLASH       : '/';
PERCENT     : '%';
AMP         : '&';
PIPE        : '|';
CARET       : '^';
TILDE       : '~';
BANG        : '!';
LT          : '<';
GT          : '>';
AT          : '@';
UNDERSCORE  : '_';

// ── Literals ─────────────────────────────────

HEX_LIT
    : '0' [xX] [0-9a-fA-F] [0-9a-fA-F_]*
    ;

FLOAT_LIT
    : [0-9] [0-9_]* '.' [0-9] [0-9_]* ([eE] [+\-]? [0-9]+)?
    | [0-9] [0-9_]* [eE] [+\-]? [0-9]+
    ;

INT_LIT
    : [0-9] [0-9_]*
    ;

CHAR_LIT
    : '\'' ( '\\' [nrt\\'0] | ~['\\\r\n] ) '\''
    ;

STRING_LIT
    : '"' ( '\\' [nrt\\"0] | ~["\\\r\n] )* '"'
    ;

// ── Identifiers ──────────────────────────────

IDENTIFIER
    : [a-zA-Z_] [a-zA-Z0-9_]*
    ;

// ── Whitespace and comments ──────────────────

NL
    : [\r\n]+ -> channel(HIDDEN)
    ;

WS
    : [ \t]+ -> skip
    ;

LINE_COMMENT
    : '//' ~[\r\n]* -> channel(HIDDEN)
    ;

BLOCK_COMMENT
    : '/*' .*? '*/' -> channel(HIDDEN)
    ;