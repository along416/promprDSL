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
AFTER   : 'after';
PARSE   : 'parse';
JSONFIX : 'jsonfix';
MARKDOWN : 'markdown';
IF : 'if';
ELSE : 'else';
OUTPUTSPEC : 'outputspec';

// 基础 Tokens
ID      : [a-zA-Z_][a-zA-Z_0-9]* ;
STRING  : '"' ( ~["\\] | '\\' . )* '"' ;
NUMBER  : '-'? [0-9]+ ('.' [0-9]+)? ;
BOOL    : 'true' | 'false' ;
PIPE    : '|';
SEMI    : ';';

// 特殊标签，切换模式
AFTER_START : '<after>' -> pushMode(JS_MODE);
AFTER_END   : '</after>' -> popMode;

FIX_START   : '<fix>' -> pushMode(JS_MODE);
FIX_END     : '</fix>' -> popMode;

// 默认模式跳过空白和注释
WS      : [ \t\r\n]+ -> skip ;
LINE_COMMENT  : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT : '/*' .*? '*/' -> skip ;

// JS_MODE 里的规则：匹配任意字符直到遇到结束标签（通过 popMode 实现）
mode JS_MODE;

JS_CODE
    : .+?    // 非贪婪匹配任意字符
    ;
    
// 基础类型
STRING_TYPE : 'string';
FLOAT_TYPE  : 'float';
INT_TYPE    : 'int';