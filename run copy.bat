@echo off
chcp 65001

REM Step 1: 运行 test.go，生成模型回复到 model_output.json
echo 正在生成 ANTLR 解析器...
  antlr4 -Dlanguage=Go -o promptParser/ ./grammar/PromptDSLLexer.g4
  antlr4 -Dlanguage=Go -o promptParser/ -lib promptParser/grammar ./grammar/PromptDSLParser.g4
if errorlevel 1 (
  echo 运行 test.go 失败
  exit /b 1
)

REM Step 2: 执行 generated_prompt/generated_prompt.go，处理模型回复
go run generated_prompt/workflow.go
if errorlevel 1 (
  echo 执行 workflow.go 失败
  exit /b 1
)

echo 所有流程完成！
pause