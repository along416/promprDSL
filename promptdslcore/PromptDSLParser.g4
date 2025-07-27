parser grammar PromptDSLParser;

options { tokenVocab=PromptDSLLexer; }

// ===============================
// ========== Parser Rules ==========
// ===============================

promptFile  : promptDef+ EOF ;

promptDef   : PROMPT ID LBRACE promptBlock+ RBRACE ;

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
sysSection: SYSTEM LBRACE ID+ RBRACE ;

// 自定义
moduleDef: ID LBRACE moduleContent* RBRACE ;
moduleContent
    : textLine
    | paramPath
    ;

// 输入
inputSection
    : INPUT LBRACE fieldDef+ RBRACE
    ;

// 输出
outputSection
    : OUTPUT (outputStruct | outputMarkdown)
    ;
outputStruct: LBRACE fieldDef+ RBRACE ;
outputMarkdown: COLON MARKDOWN ;

// system
systemSection
    : SYSTEM LBRACE textLine+ RBRACE
    ;

// user
userSection
    : USER LBRACE userContent+ RBRACE
    ;

userContent
    : ifStatement
    | textLine
    | OUTPUTSPEC
    | expr
    ;

// 条件
ifStatement
    : IF LPAREN condition RPAREN LBRACE userContent* RBRACE (ELSE LBRACE userContent* RBRACE)?
    ;
condition
    : expr (EQUALS | NOTEQUALS) expr
    | expr
    ;

// note
noteSection
    : NOTE LBRACE textLine+ RBRACE
    ;

// Fix section - uses CODE_BLOCK mode
fixSection
    : FIX codeBlockContent RBRACE
    ;

// After section - uses CODE_BLOCK mode
afterSection
    : AFTER codeBlockContent RBRACE
    ;

// Handle nested braces and strings in code blocks
codeBlockContent
    : (CODE_TEXT | CODE_STRING | LBRACE codeBlockContent RBRACE)*
    ;

// 表达式支持
expr
    : STRING
    | NUMBER
    | BOOL
    | paramPath
    ;

dslCallExpr
    : paramPath LPAREN (expr (COMMA expr)*)? RPAREN
    ;

// 字段定义
fieldDef
    : ID COLON type (EQUALS value)? annotation* SEMI?
    ;

// 单行文本
textLine
    : STRING
    | LINE_COMMENT
    | paramPath
    ;

// 参数路径形式，如 input.xxx
paramPath
    : (ID | INPUT | OUTPUT | AFTER | BEFORE) (DOT ID)*
    ;

// 结构体定义
structDef
    : ID STRUCT LBRACE fieldDef+ RBRACE
    ;

// 注解定义
annotation
    : AT ID LPAREN annotationArgs? RPAREN
    ;

annotationArgs
    : annotationValue (COMMA annotationValue)*
    ;

annotationValue
    : STRING
    | arrayLiteral
    ;

arrayLiteral
    : LBRACK (STRING (COMMA STRING)*)? RBRACK
    ;

// 模糊文本块
textBlock
    : (STRING | MINUS | COLON | ID | NUMBER | WS)+
    ;

// 类型支持
type
    : STRING_TYPE
    | FLOAT_TYPE
    | INT_TYPE
    | LBRACK RBRACK type
    | 'struct' LBRACE fieldDef* RBRACE
    | ID
    ;

value
    : STRING
    | NUMBER
    | BOOL
    ;

formatType
    : MD
    | JSON
    ;