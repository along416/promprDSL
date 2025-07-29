grammar PromptDSL;

// 解析规则
promptFile  : promptDef+ EOF ;

promptDef   : PROMPT ID '{' promptBlock+ '}' ;

promptBlock 
    : inputSection
    | outputSection
    | systemSection
    | userSection
    | noteSection 
    | afterSection
    | fixSection
    | moduleDef //自定义模块
    ;

//input
inputSection
    : 'input'  '{' fieldDef+ '}'
    ;

//output
outputSection

    : defaultAnnotation* 'output' (outputStruct | outputMarkdown)
    ;
outputStruct:'{' fieldDef+ '}';
outputMarkdown: ':' MARKDOWN;
//before
beforeSection
    : 'before' '{' beforeContent* '}'
    ;

beforeContent
    : varDef
    | expr
    | ifStatement
    | textLine
    ;

varDef
    : ID '=' expr
    ;
//sys    
systemSection
    : 'system' '{' ID+ '}'
    | 'system' '{' sysContent+ '}'
    ;
sysContent
    : ifStatement
    | paramPath
    | ARRAY_OUTPUTSPEC
    | OUTPUTSPEC
    | expr
    | textLine
    ;
moduleDef: ID '{' moduleContent* '}';
moduleContent
    : ifStatement
    | paramPath
    | ARRAY_OUTPUTSPEC
    | OUTPUTSPEC
    | expr
    | textLine
    ;
//user
userSection
    : 'user' '{' ID+ '}'
    | 'user' '{' userContent+ '}'
    ;

userContent
    : ifStatement
    | paramPath
    | forStatement
    | ARRAY_OUTPUTSPEC
    | OUTPUTSPEC
    | expr
    | textLine
    ;
ARRAY_OUTPUTSPEC 
    : '[' ']' OUTPUTSPEC
    ;
thencontent
    : userContent
    ;
elsecontent
    : userContent
    ;
ifStatement
    : 'if' '(' condition ')' '{' thencontent* '}' ('else' '{' elsecontent* '}')?
    ;

condition
    : lhs=expr op=('==' | '!=') rhs=expr
    | single=expr
    ;
// for
forStatement
    : 'for' '(' init=expr ';' condition ';' update=expr ')' '{' userContent* '}'
    ;


//note
noteSection
    : 'note' '{' textLine+ '}'
    ;


// after
PLUS : '+';

// dslCallExpr
//     : paramPath '(' (paramPath | STRING | NUMBER | BOOL)? ')' // 支持函数调用
//     ;
dslCallExpr
    : paramPath '(' (expr (',' expr)*)? ')'
    ;
expr
    : paramPath
    | STRING
    | NUMBER
    | BOOL
    // | paramPath '==' (paramPath | STRING | NUMBER | BOOL)
    ;


// 字段定义
fieldDef
    : ID ':' type ( '=' value )? annotation* SEMI? 
    ;
// 文本行支持    
textLine
    : STRING
    | LINE_COMMENT
    | paramPath
    ;
paramPath
    : (ID | INPUT | OUTPUT | AFTER | BEFORE ) ('.' ID )*
    ;
// 结构体定义
structDef
    : ID STRUCT '{' fieldDef+ '}' 
    ;

// 注解
annotation
    : '@' ID ('(' annotationArgs? ')')? 
    ;

annotationArgs
    : annotationValue (',' annotationValue)* 
    ;

annotationValue
    : STRING
    | arrayLiteral
    ;

arrayLiteral
    : '[' (STRING (',' STRING)*)? ']' 
    ;

defaultAnnotation
    : '@' ID ('(' annotationArgs? ')')? 
    ;
afterSection : AFTER_BLOCK ;
fixSection   : FIX_BLOCK ;

AFTER_BLOCK : '<after>' .*? '</after>' ;
FIX_BLOCK   : '<fix>'   .*? '</fix>' ;


// 纯文本段
textBlock
    : (STRING | '-' | ':' | ID | NUMBER | WS)+ 
    ;

STRING_TYPE : 'string';
FLOAT_TYPE  : 'float';
INT_TYPE    : 'int';
// 类型定义
type
    : 'struct' '{' fieldDef* '}' 
    | FLOAT_TYPE
    | INT_TYPE
    | '[]' type
    | STRING_TYPE   
    | ID 
    ;

value
    : STRING
    | NUMBER
    | BOOL
    ;

formatType
    : 'md'
    | 'json'
    ;

// 词法规则
// 关键字 Tokens (必须放在ID之前)
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


// 空白和注释
WS              : [ \t\r\n]+         -> channel(HIDDEN) ;
LINE_COMMENT    : '//' ~[\r\n]*      -> channel(HIDDEN) ;
BLOCK_COMMENT   : '/*' .*? '*/'      -> channel(HIDDEN) ;

