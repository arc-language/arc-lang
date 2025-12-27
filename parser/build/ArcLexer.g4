lexer grammar ArcLexer;

// Keywords
IMPORT: 'import';
NAMESPACE: 'namespace';
LET: 'let';
CONST: 'const';
FUNC: 'func';
ASYNC: 'async';
AWAIT: 'await';
STRUCT: 'struct';
CLASS: 'class';
MUTATING: 'mutating';
DEINIT: 'deinit';
RETURN: 'return';
IF: 'if';
ELSE: 'else';
FOR: 'for';
IN: 'in';
BREAK: 'break';
CONTINUE: 'continue';
DEFER: 'defer';
EXTERN: 'extern';
SELF: 'self';
NULL: 'null';

// Type Keywords
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
VECTOR: 'vector';
MAP: 'map';

// Built-in Functions/Intrinsics
ALLOCA: 'alloca';
CAST: 'cast';
SYSCALL: 'syscall';
SIZEOF: 'sizeof';
ALIGNOF: 'alignof';
MEMSET: 'memset';
MEMCPY: 'memcpy';
MEMMOVE: 'memmove';
STRLEN: 'strlen';
MEMCHR: 'memchr';
VA_START: 'va_start';
VA_ARG: 'va_arg';
VA_END: 'va_end';
RAISE: 'raise';
MEMCMP: 'memcmp';
BIT_CAST: 'bit_cast';

// Operators (multi-character first!)
ARROW: '->';
RANGE: '..';      // Must come before DOT and ELLIPSIS
ELLIPSIS: '...';  // Must come after RANGE
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
INCREMENT: '++';
DECREMENT: '--';

// Single character operators
PLUS: '+';
MINUS: '-';
STAR: '*';
SLASH: '/';
PERCENT: '%';
LT: '<';
GT: '>';
NOT: '!';
AMP: '&';
AT: '@';
ASSIGN: '=';

// Delimiters
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';
COMMA: ',';
COLON: ':';
SEMICOLON: ';';
DOT: '.';         // Single dot comes after RANGE and ELLIPSIS

// Literals
BOOLEAN_LITERAL: 'true' | 'false';

INTEGER_LITERAL
    : DECIMAL_LITERAL
    | HEX_LITERAL
    | OCTAL_LITERAL
    | BINARY_LITERAL
    ;

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

STRING_LITERAL
    : '"' (~["\\\r\n] | ESCAPE_SEQUENCE)* '"'
    ;

CHAR_LITERAL
    : '\'' (~['\\\r\n] | ESCAPE_SEQUENCE) '\''
    ;

fragment ESCAPE_SEQUENCE
    : '\\' ['"\\nrt0]  // Added \0 for null character
    | '\\' 'x' HEX_DIGIT HEX_DIGIT
    | '\\' 'u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
    | '\\' 'U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
    ;

fragment HEX_DIGIT: [0-9a-fA-F];

// Identifiers
IDENTIFIER
    : [a-zA-Z_] [a-zA-Z0-9_]*
    ;

// Whitespace and Comments
WS: [ \t\r\n]+ -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;