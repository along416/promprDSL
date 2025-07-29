@echo off
chcp 65001


REM 生成词法分析器
antlr4 -Dlanguage=Go -package parser -o ./promptdslcore/parser ./promptdslcore/grammar/PromptDSLLexer.g4

REM 生成语法分析器
antlr4 -Dlanguage=Go -package parser -o ./promptdslcore/parser -lib  ./promptdslcore/parser ./promptdslcore/grammar/PromptDSLParser.g4

echo 解析器生成完成。
@REM pausec