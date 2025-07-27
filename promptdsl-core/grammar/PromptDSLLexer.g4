lexer grammar PromptDSLLexer;

AFTER_BLOCK : '<after>' .*? '</after>' ;
FIX_BLOCK   : '<fix>'   .*? '</fix>' ;

// 基础类型
STRING_TYPE : 'string';
FLOAT_TYPE  : 'float';
INT_TYPE    : 'int';

// 关键字 Tokens（必须放在 ID 之前）
PROMPT   : 'prompt';
PARAMS   : 'params';
SYSTEM   : 'system';
USER     : 'user';
NOTE     : 'note';
INPUT    : 'input';
OUTPUT   : 'output';
FORMAT   : 'format';
TYPE     : 'type';
STRUCT   : 'struct';
BEFORE   : 'before';
SCHEMA   : 'schema';
PARSE    : 'parse';
JSONFIX  : 'jsonfix';
MARKDOWN : 'markdown';
IF       : 'if';
ELSE     : 'else';
OUTPUTSPEC : 'outputspec';

FIX: 'fix' -> pushMode(CODE_BLOCK);
AFTER: 'after' -> pushMode(CODE_BLOCK);
ARRAY_OUTPUTSPEC : '[' ']' OUTPUTSPEC ;

// 补充符号Token定义
LBRACE : '{' ;
RBRACE : '}' -> popMode ;
LPAREN : '(' ;
RPAREN : ')' ;
COLON  : ':' ;
EQUAL  : '=' ;
COMMA  : ',' ;
DOT    : '.' ;
EQEQ : '==' ;
NOTEQ : '!=' ;
AT       : '@' ;
MD       : 'md' ;
JSON     : 'json' ;

// 基础 Tokens
ID     : [a-zA-Z_][a-zA-Z_0-9]* ;
STRING : '"' ( ~["\\] | '\\' . )* '"' ;
NUMBER : '-'? [0-9]+ ('.' [0-9]+)? ;
BOOL   : 'true' | 'false' ;
PIPE   : '|';
SEMI   : ';';
PLUS   : '+';

// Mode for code blocks in FIX and AFTER sections
mode CODE_BLOCK;

CODE_LBRACE: '{' -> type(LBRACE), pushMode(CODE_BLOCK);
CODE_RBRACE: '}' -> type(RBRACE), popMode;
CODE_STRING: '"' CODE_STRING_CHAR* '"';
fragment CODE_STRING_CHAR: ~["\n\r\\] | '\\\\' .;
CODE_TEXT: ~[{}"]+;

// 空白和注释
WS            : [ \t\r\n]+      -> channel(HIDDEN) ;
LINE_COMMENT  : '//' ~[\r\n]*   -> channel(HIDDEN) ;
BLOCK_COMMENT : '/*' .*? '*/'   -> channel(HIDDEN) ;
