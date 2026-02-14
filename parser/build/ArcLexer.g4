lexer grammar ArcLexer;

// =============================================================================
// Keywords
// =============================================================================
IMPORT: 'import';
NAMESPACE: 'namespace';
LET: 'let';
VAR: 'var'; // Added for mutable references
CONST: 'const';
FUNC: 'func';
ASYNC: 'async';
AWAIT: 'await';
PROCESS: 'process';
GPU: 'gpu';
STRUCT: 'struct';
CLASS: 'class';
DEINIT: 'deinit';
RETURN: 'return';
IF: 'if';
ELSE: 'else';
FOR: 'for';
IN: 'in';
BREAK: 'break';
CONTINUE: 'continue';
DEFER: 'defer';
SELF: 'self';
NULL: 'null';
SWITCH: 'switch';
CASE: 'case';
DEFAULT: 'default';
ENUM: 'enum';
RAWPTR: 'rawptr';

// Extern triggers mode switch
EXTERN: 'extern' -> pushMode(EXTERN_LANG_MODE);

// Extern-specific keywords (used inside extern blocks)
VIRTUAL: 'virtual';
STATIC: 'static';
ABSTRACT: 'abstract';
NEW: 'new';
DELETE: 'delete';
TYPE: 'type';

// Calling conventions
STDCALL: 'stdcall';
CDECL: 'cdecl';
FASTCALL: 'fastcall';
VECTORCALL: 'vectorcall';
THISCALL: 'thiscall';

// =============================================================================
// Types
// =============================================================================
INT8: 'int8';
INT16: 'int16';
INT32: 'int32';
INT64: 'int64';
UINT8: 'uint8';
UINT16: 'uint16';
UINT32: 'uint32';
UINT64: 'uint64';
USIZE: 'usize';
ISIZE: 'isize';
FLOAT32: 'float32';
FLOAT64: 'float64';
BYTE: 'byte';
BOOL: 'bool';
CHAR: 'char';
STRING: 'string';
VOID: 'void';

// =============================================================================
// Operators
// =============================================================================
EQ: '==';
NE: '!=';
LE: '<=';
GE: '>=';
AND: '&&';
OR: '||';
PLUS_ASSIGN: '+=';
MINUS_ASSIGN: '-=';
STAR_ASSIGN: '*=';
SLASH_ASSIGN: '/=';
PERCENT_ASSIGN: '%=';
BIT_OR_ASSIGN: '|=';
BIT_AND_ASSIGN: '&=';
BIT_XOR_ASSIGN: '^=';
SHL_ASSIGN: '<<=';
SHR_ASSIGN: '>>=';
INCREMENT: '++';
DECREMENT: '--';
FAT_ARROW: '=>';
RANGE: '..';
ELLIPSIS: '...';
PLUS: '+';
MINUS: '-';
STAR: '*';
SLASH: '/';
PERCENT: '%';
LT: '<';
GT: '>';
NOT: '!';
AMP: '&';
BIT_OR: '|';
BIT_XOR: '^';
BIT_NOT: '~';
AT: '@';
ASSIGN: '=';

// =============================================================================
// Delimiters
// =============================================================================
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';
COMMA: ',';
COLON: ':';
SEMICOLON: ';';
DOT: '.';
UNDERSCORE: '_';

// =============================================================================
// Literals
// =============================================================================
BOOLEAN_LITERAL: 'true' | 'false';

INTEGER_LITERAL: DECIMAL_LITERAL | HEX_LITERAL | OCTAL_LITERAL | BINARY_LITERAL;

fragment DECIMAL_LITERAL: [0-9] [0-9_]*;
fragment HEX_LITERAL: '0' [xX] [0-9a-fA-F] [0-9a-fA-F_]*;
fragment OCTAL_LITERAL: '0' [oO] [0-7] [0-7_]*;
fragment BINARY_LITERAL: '0' [bB] [01] [01_]*;

FLOAT_LITERAL
    : DECIMAL_LITERAL '.' DECIMAL_LITERAL EXPONENT?
    | DECIMAL_LITERAL EXPONENT
    | '.' DECIMAL_LITERAL EXPONENT?
    ;

fragment EXPONENT: [eE] [+-]? DECIMAL_LITERAL;

STRING_LITERAL: '"' (~["\\\r\n] | ESCAPE_SEQUENCE)* '"';

CHAR_LITERAL: '\'' (~['\\\r\n] | ESCAPE_SEQUENCE) '\'';

fragment ESCAPE_SEQUENCE
    : '\\' ['"\\nrt0]
    | '\\' 'x' HEX_DIGIT HEX_DIGIT
    | '\\' 'u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
    | '\\' 'U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
    ;

fragment HEX_DIGIT: [0-9a-fA-F];

IDENTIFIER: [a-zA-Z_] [a-zA-Z0-9_]*;

WS: [ \t\r\n]+ -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;

// =============================================================================
// Extern Language Mode
// =============================================================================
mode EXTERN_LANG_MODE;

EXTERN_WS: [ \t\r\n]+ -> skip;
C_LANG: 'c' -> popMode;
CPP_LANG: 'cpp' -> popMode;