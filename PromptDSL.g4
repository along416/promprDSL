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

promptBlock : 
      PARAMS '{' paramBody '}' 
    | SYSTEM '{' textBlock '}' 
    | USER   '{' userBlock '}' 
    | NOTE   '{' textBlock '}' ;

// params 结构
paramBody   : paramSection+ ;

paramSection
    : IN ':' '{' fieldDef+ '}'                 # inputSection
    | OUTPUT ':' '{' outputBody '}'            # outputSection
    | fieldDef                                 # singleField
    ;

outputBody
    : (FORMAT ':' formatType)? 
      (TYPE structDef)? 
      (SCHEMA ':' '[' ID ']')? 
    ;

// 字段定义
fieldDef
    : ID ':' type ( '=' value )? annotation* SEMI? 
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

// 用户内容块
userBlock
    : (textBlock | classificationBlock | summarizationBlock | compilationBlock)+ 
    ;

classificationBlock
    : 'classification' '{' kvPair* '}' 
    ;

summarizationBlock
    : 'summarization' '{' kvPair* '}' 
    ;

compilationBlock
    : 'compilation' '{' kvPair* '}' 
    ;

kvPair
    : ID ':' (STRING | ID | 'extra_hint') 
    ;

// 纯文本段
textBlock
    : (STRING | '-' | ':' | ID | NUMBER | WS)+ 
    ;

// 类型定义
type
    : 'string'
    | 'float'
    | 'int'
    | '[' type ']' 
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
IN      : 'in';
OUTPUT  : 'output';
FORMAT  : 'format';
TYPE    : 'type';
STRUCT  : 'struct';
SCHEMA  : 'schema';

// 基础 Tokens
ID      : [a-zA-Z_][a-zA-Z_0-9]* ;
STRING  : '"' ( ~["\\] | '\\' . )* '"' ;
NUMBER  : '-'? [0-9]+ ('.' [0-9]+)? ;
BOOL    : 'true' | 'false' ;
PIPE    : '|';
SEMI    : ';';

// 特殊文本规则
// TEXT: ~[{}:";@[\]\r\n]+ {
//     switch l.GetText() {
//     case "prompt":
//         l.SetType(PromptDSLLexerPROMPT)
//     case "params":
//         l.SetType(PromptDSLLexerPARAMS)
//     case "system":
//         l.SetType(PromptDSLLexerSYSTEM)
//     case "user":
//         l.SetType(PromptDSLLexerUSER)
//     case "note":
//         l.SetType(PromptDSLLexerNOTE)
//     }
// };

// 空白和注释
WS      : [ \t\r\n]+ -> skip ;
LINE_COMMENT  : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT : '/*' .*? '*/' -> skip ;