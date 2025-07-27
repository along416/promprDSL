lexer grammar PromptDSLLexer;


// 保留关键字
PROMPT  : 'prompt';
PARAMS  : 'params';
SYSTEM  : 'system';
USER    : 'user';
NOTE    : 'note';
INPUT   : 'input';
OUTPUT  : 'output';
FORMAT  : 'format';
TYPE    : 'type';
STRUCT  : 'struct';
BEFORE  : 'before';
SCHEMA  : 'schema';
PARSE   : 'parse';
JSONFIX : 'jsonfix';
MARKDOWN : 'markdown';
IF : 'if';
ELSE : 'else';
OUTPUTSPEC : 'outputspec';


LBRACE: '{' ;
RBRACE: '}' ;
LPAREN: '(';
RPAREN: ')';
COLON: ':';
EQUALS: '==';
NOTEQUALS: '!=';
AT: '@';
DOT: '.';
LBRACK: '[';
RBRACK: ']';
COMMA: ',';
MINUS : '-';
// INTEGER: 'int';
// FLOAT: 'float';
BOOLEAN: 'bool';
ARRAY_STRING: '[]string';
ARRAY_INTEGER: '[]int';
ARRAY_FLOAT: '[]float';
ARRAY_BOOLEAN: '[]bool';
MD   : 'md';
JSON : 'json';


// 基础 Tokens
ID      : [a-zA-Z_][a-zA-Z_0-9]* ;
STRING  : '"' ( ~["\\] | '\\' . )* '"' ;
NUMBER  : '-'? [0-9]+ ('.' [0-9]+)? ;
BOOL    : 'true' | 'false' ;
PIPE    : '|';
SEMI    : ';';


HINT: 'hint';
FIX: 'fix' WS* '{' -> pushMode(CODE_BLOCK);
AFTER: 'after' WS* '{' -> pushMode(CODE_BLOCK);
// 默认模式跳过空白和注释
WS      : [ \t\r\n]+ -> skip ;
LINE_COMMENT  : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT : '/*' .*? '*/' -> skip ;

// JS_MODE 里的规则：匹配任意字符直到遇到结束标签（通过 popMode 实现）
mode CODE_BLOCK;

CODE_LBRACE: '{' -> type(LBRACE), pushMode(CODE_BLOCK);
CODE_RBRACE: '}' -> type(RBRACE), popMode;
CODE_STRING: '"' CODE_STRING_CHAR* '"';
fragment CODE_STRING_CHAR: ~["\n\r\\] | '\\\\' .;
CODE_TEXT: ~[{}"]+;

// 基础类型
STRING_TYPE : 'string';
FLOAT_TYPE  : 'float';
INT_TYPE    : 'int';