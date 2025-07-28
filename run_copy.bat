@echo off
chcp 65001


REM 生成词法分析器
antlr4 -Dlanguage=Go -package parser -o parser ./promptdslcore/grammar/PromptDSLLexer.g4

REM 生成语法分析器
antlr4 -Dlanguage=Go -package parser -o ./parser -lib  ./parser ./promptdslcore/grammar/PromptDSLParser.g4

echo 完成 ANTLR Go 代码生成。
@REM pause
