parser grammar PromptDSLParser;

options { tokenVocab=PromptDSLLexer; }

// ===============================
// ========== Parser Rules ==========
// ===============================

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
    | moduleDef // 自定义
    ;

// sys
sysSection: 'sys' '{' ID+ '}' ;

// 自定义
moduleDef: ID '{' moduleContent* '}' ;
moduleContent
    : textLine
    | paramPath
    ;

// 输入
inputSection
    : 'input'  '{' fieldDef+ '}'
    ;

// 输出
outputSection
    : 'output' (outputStruct | outputMarkdown)
    ;
outputStruct:'{' fieldDef+ '}';
outputMarkdown: ':' MARKDOWN;

// system
systemSection
    : 'system' '{' textLine+ '}'
    ;

// user
userSection
    : 'user' '{' userContent+ '}'
    ;

userContent
    : ifStatement
    | textLine
    | OUTPUTSPEC
    | expr
    ;

// 条件
ifStatement
    : 'if' '(' condition ')' '{' userContent* '}' ('else' '{' userContent* '}')?
    ;
condition
    : expr ('==' | '!=') expr
    | expr
    ;

// note
noteSection
    : 'note' '{' textLine+ '}'
    ;

// after 与 fix 特殊代码块包裹
afterSection
    : AFTER_START JS_CODE AFTER_END
    ;

fixSection
    : FIX_START JS_CODE FIX_END
    ;

// 表达式支持
expr
    : STRING
    | NUMBER
    | BOOL
    | paramPath
    ;

dslCallExpr
    : paramPath '(' (expr (',' expr)*)? ')'
    ;

// 字段定义
fieldDef
    : ID ':' type ( '=' value )? annotation* SEMI? 
    ;

// 单行文本
textLine
    : STRING
    | LINE_COMMENT
    | paramPath
    ;

// 参数路径形式，如 input.xxx
paramPath
    : (ID | INPUT | OUTPUT | AFTER | BEFORE) ('.' ID )*
    ;

// 结构体定义
structDef
    : ID STRUCT '{' fieldDef+ '}' 
    ;

// 注解定义
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

// 模糊文本块
textBlock
    : (STRING | '-' | ':' | ID | NUMBER | WS)+ 
    ;

// 类型支持
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


