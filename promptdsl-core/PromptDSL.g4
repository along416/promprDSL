grammar PromptDSL;

// @lexer::header {
// import "strings"
// }

//@lexer::members {
//func (l *PromptDSLLexer) isKeyword(text string) bool {
//    keywords := []string{"prompt", "params", "system", "user", "note"}
//    trimmed := strings.TrimSpace(text)
//    for _, kw := range keywords {
//        if trimmed == kw {
//            return true
//        }
//    }
//    return false
//}
//}

// 解析规则
promptFile  : promptDef+ EOF ;

promptDef   : PROMPT ID '{' promptBlock+ '}' ;

promptBlock 
    : inputSection
    | outputSection
    | sysSection
    | systemSection
    | userSection
    | noteSection 
    | afterSection
    | fixSection
    | moduleDef //自定义模块
    ;
//sys    
sysSection: 'sys' '{' ID+ '}';

moduleDef: ID '{' moduleContent* '}';
moduleContent
    : textLine
    | paramPath
    ;

//input
inputSection
    : 'input'  '{' fieldDef+ '}'
    ;

//output
outputSection
    : 'output' (outputStruct | outputMarkdown)
    ;
outputStruct:'{' fieldDef+ '}';
outputMarkdown: ':' MARKDOWN;


//sys
systemSection
    : 'system' '{' textLine+ '}'
    ;
//user
userSection
    : 'user' '{' userContent+ '}'
    ;

userContent
    : ifStatement
    | textLine
    | OUTPUTSPEC
    | expr
    ;

ifStatement
    : 'if' '(' condition ')' '{' userContent* '}' ('else' '{' userContent* '}')?
    ;

condition
    : lhs=expr op=('==' | '!=') rhs=expr
    | single=expr
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
    : STRING
    | NUMBER
    | BOOL
    | paramPath
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
    : (ID | INPUT | OUTPUT | AFTER | BEFORE) ('.' ID )*
    ;
// 结构体定义
structDef
    : ID STRUCT '{' fieldDef+ '}' 
    ;

// 注解
annotation
    : '@' ID '(' annotationArgs? ')' 
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

afterSection
    : '<after>' .*? '</after>'
    ;

fixSection
    : '<fix>' .*? '</fix>'
    ;



// 纯文本段
textBlock
    : (STRING | '-' | ':' | ID | NUMBER | WS)+ 
    ;

STRING_TYPE : 'string';
FLOAT_TYPE  : 'float';
INT_TYPE    : 'int';
// 类型定义
type
    : STRING_TYPE
    | FLOAT_TYPE
    | INT_TYPE
    | '[]' type
    | 'struct' '{' fieldDef* '}'    
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
WS      : [ \t\r\n]+ -> skip ;
LINE_COMMENT  : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT : '/*' .*? '*/' -> skip ;
