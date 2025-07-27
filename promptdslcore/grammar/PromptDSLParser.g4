parser grammar PromptDSLParser;

options { tokenVocab=PromptDSLLexer; }

// 解析规则
promptFile  : promptDef+ EOF ;

promptDef   : PROMPT ID LBRACE promptBlock+ RBRACE ;

promptBlock 
    : inputSection
    | outputSection
    | systemSection
    | userSection
    | noteSection 
    | afterSection
    | fixSection
    | moduleDef
    ;

inputSection
    : INPUT LBRACE fieldDef+ RBRACE
    ;

outputSection
    : defaultAnnotation* OUTPUT (outputStruct | outputMarkdown)
    ;

outputStruct : LBRACE fieldDef+ RBRACE ;
outputMarkdown : COLON MARKDOWN ;

beforeSection
    : BEFORE LBRACE beforeContent* RBRACE
    ;

beforeContent
    : varDef
    | expr
    | ifStatement
    | textLine
    ;

varDef
    : ID EQUAL expr
    ;

systemSection
    : SYSTEM LBRACE ID+ RBRACE
    | SYSTEM LBRACE sysContent+ RBRACE
    ;

sysContent
    : ifStatement
    | paramPath
    | ARRAY_OUTPUTSPEC
    | OUTPUTSPEC
    | expr
    | textLine
    ;

moduleDef : ID LBRACE moduleContent* RBRACE ;
moduleContent
    : ifStatement
    | paramPath
    | ARRAY_OUTPUTSPEC
    | OUTPUTSPEC
    | expr
    | textLine
    ;

userSection
    : USER LBRACE ID+ RBRACE
    | USER LBRACE userContent+ RBRACE
    ;

userContent
    : ifStatement
    | paramPath
    | ARRAY_OUTPUTSPEC
    | OUTPUTSPEC
    | expr
    | textLine
    ;

thencontent : userContent ;
elsecontent : userContent ;

ifStatement
    : IF LPAREN condition RPAREN LBRACE thencontent* RBRACE (ELSE LBRACE elsecontent* RBRACE)?
    ;

condition
    : lhs=expr op=(EQEQ | NOTEQ) rhs=expr
    | single=expr
    ;

// noteSection改成用token
noteSection
    : NOTE LBRACE textLine+ RBRACE
    ;

dslCallExpr
    : paramPath LPAREN (expr (COMMA expr)*)? RPAREN
    ;

expr
    : paramPath
    | STRING
    | NUMBER
    | BOOL
    ;

fieldDef
    : ID COLON type ( EQUAL value )? annotation* SEMI?
    ;

textLine
    : STRING
    | LINE_COMMENT
    | paramPath
    ;

paramPath
    : (ID | INPUT | OUTPUT | AFTER | BEFORE) (DOT ID)*
    ;

structDef
    : ID STRUCT LBRACE fieldDef+ RBRACE 
    ;

annotation
    : AT ID (LPAREN annotationArgs? RPAREN)?
    ;
annotationArgs
    : annotationValue (COMMA annotationValue)* 
    ;

annotationValue
    : STRING
    | arrayLiteral
    ;

arrayLiteral
    : LBRACE (STRING (COMMA STRING)*)? RBRACE 
    ;

defaultAnnotation
    : AT ID (LPAREN annotationArgs? RPAREN)?
    ;

// Fix section - uses CODE_BLOCK mode
fixSection
    : FIX  codeBlockContent RBRACE
    ;

// After section - uses CODE_BLOCK mode
afterSection
    : AFTER  codeBlockContent RBRACE
    ;

// Handle nested braces and strings in code blocks
codeBlockContent
    : (CODE_TEXT | CODE_STRING | LBRACE codeBlockContent RBRACE)*
    ;

type
    : STRUCT LBRACE fieldDef* RBRACE
    | FLOAT_TYPE
    | INT_TYPE
    | LBRACK RBRACK type  // 这里你之前用的是 '[]' type，建议改写对应你Lexer中的表示数组的token，如果没有，需要补充
    | STRING_TYPE   
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
