// promptdsl-core/runner.go
package promptdslcore

import (
	"fmt"
	"os"

	"promptdslcore/parser"
	// "strings"

	"github.com/antlr4-go/antlr/v4"
)

// 生成单prompt代码，返回调用名
func RunPromptDSL(input string, filename string) (*final, error) {

	// 1. 解析输入 DSL 文本，生成 Parse Tree
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewPromptDSLLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPromptDSLParser(tokenStream)
	tree := p.PromptFile()
	fmt.Println("🌳 start...")
	fmt.Println(tree.ToStringTree(nil, p))
	fmt.Println("🌳 ...end")
	// 2. 构建 AST Node
	ctx := tree.(*parser.PromptFileContext)

	rootNode := ConvertASTtoPrompt(ctx, tokenStream)
	// fmt.Printf("%v\n", rootNode)
	fmt.Printf("📦 OutFields: %+v\n", rootNode.OutFields)
	fmt.Printf("📦 aftercode: %+v\n", rootNode.AfterCode)

	// 3. 构造 Eval 上下文
	str := &PromptEvalContext{
		Vars:       make(map[string]interface{}),
		InFields:   rootNode.InFields,
		OutFields:  rootNode.OutFields,
		ModuleDefs: rootNode.ModuleDefs,
	}
	// fmt.Println("😅")
	//3.5before
	// 3. 先执行 before 节点，填充 Vars
	// for _, node := range rootNode.BeforeNodes {
	// 	err := node.Eval(str)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// }

	// 4. 执行 AST，得到 prompt 字符串
	outputParts, err := rootNode.Eval(str)
	if err != nil {
		return nil, err
	}
	//生成sys+user+after+fix
	// code:=Generateprompthandle(rootNode, getCurrentPackageName())
	code:=Generateprompthandle(rootNode, "generated", outputParts,filename)
	outputFile := "../generated_code/generated/" + filename + ".go"
	err = os.WriteFile(outputFile, []byte(code), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "写入文件失败: %v\n", err)
		os.Exit(1)
	}
	return outputParts, nil
}
