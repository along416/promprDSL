// promptdsl-core/runner.go
package promptdslcore

import (
	"fmt"
	"go/format"
	"log"
	"os"
	"promptdslcore/parser"
	// "strings"

	"github.com/antlr4-go/antlr/v4"
)

func RunPromptDSL(input string) (*final, error) {

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
		Input: map[string]any{
			"process": "第一小题求学生的接受能力最强其实就是要求分段函数的最大值，方法是分别求出各段的最大值取其最大即可。\n第二小题比较5分钟和20分钟学生的接受能力何时强，方法是把x=5代入第一段函数中，而x=20要代入到第二段函数中，比较大小即可。不同的自变量代入相应的解析式才能符合要求",
			"question":"通过研究学生的学习行为，心理学家发现，学生接受能力依赖于老师引入概念和描述问题所用的时间，讲座开始时，学生的兴趣增长，中间有一段不太长的时间，学生的兴趣保持理想的状态，随后学生注意力开始分散．分析结果和实验表明，用f（x）表示学生掌握和接受概念的能力（f（x）值越大，表示接受能力越强），x表示提出和讲授概念的时间（单位：分），可以有以下的公式：\n$$f(x) = \\begin{cases} -0.1x^2 + 2.6x + 43 & (0 < x \\le 10) \\\\ 59 & (10 < x \\le 16) \\\\ -3x + 107 & (16 < x \\le 30) \\end{cases}$$\n（1）开讲多少分钟后，学生的接受能力最强？能维持多少时间？\n（2）开讲5分钟与开讲20分钟比较，学生的接受能力何时强一些？",
		},
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
		//生成sys+user+after+fix代码
	// code:=Generateprompthandle(rootNode, getCurrentPackageName())
	code:=Generateprompthandle(rootNode, "main",outputParts)

	outputFile := "generated_prompt/generated_prompt.go"
	err = os.WriteFile(outputFile, []byte(code), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "写入文件失败: %v\n", err)
		os.Exit(1)
	}

	// pkgName := getCurrentPackageName()
	code = GenerateAfterAndFixGoCode(rootNode, "main")
	// FIXcode := GenerateFIX(rootNode, "main")
	// outputParts.Fix=FIXcode
	// 2. 使用 go/format 美化生成代码
	formattedCode, err := format.Source([]byte(code))
	if err != nil {
		log.Printf("⚠️ Go代码格式化失败，使用未格式化代码: %v", err)
		formattedCode = []byte(code)
	}

	// 3. 写入文件
	err = os.WriteFile("code_gen/code_gen.go", formattedCode, 0644)
	if err != nil {
		log.Fatalf("写入 code_gen/code_gen.go 失败: %v", err)
	}

	return outputParts, nil
}


