// promptdsl-core/runner.go
package promptdslcore

import (
	"promptdslcore/parser" // 你 ANTLR 生成文件所在包名

	"github.com/antlr4-go/antlr/v4"
)

func RunPromptDSL(input string) (string, error) {
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewPromptDSLLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPromptDSLParser(tokenStream)
	tree := p.PromptFile()

	// visitor := parser.NewMyVisitor()
	visitor := parser.NewMyVisitor()

	visitor.Visit(tree)

	// 返回最终 prompt
	return visitor.GeneratePrompt(), nil
}
