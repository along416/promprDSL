package main

import (
	"fmt"
	"promptdsl_test/promptdsl" // 替换为你实际的包路径

	"github.com/antlr4-go/antlr/v4"
	// "github.com/dop251/goja"
)

type MyVisitor struct {
	*promptdsl.BasePromptDSLVisitor
}
func (v *MyVisitor) VisitPromptBlock(ctx *promptdsl.PromptBlockContext) interface{} {
	fmt.Println("📦 Entering PromptBlock")
	return v.VisitChildren(ctx)
}
func (v *MyVisitor) VisitAfterSection(ctx *promptdsl.AfterSectionContext) interface{} {
	fmt.Println("🚪 Entering VisitAfterSection")
	afterCode := ctx.GetText()
	fmt.Println("📜 After code:", afterCode)
	return nil
}
// func (v *MyVisitor) VisitAfterSection(ctx *promptdsl.AfterSectionContext) interface{} {
// 	// 提取 `after` 部分的 JavaScript 代码
// 	fmt.Println("🚪 Entering VisitAfterSection")
// 	afterCode := ctx.GetText()
// 	fmt.Println("Extracted after code:", afterCode)

// 	// 在 Goja 中执行 `after` 部分的代码
// 	vm := goja.New()
// 	// 假设这是模型的原始回复
// 	response := "   This is the original response.   "
// 	err := vm.Set("response", response)
// 	if err != nil {
// 		fmt.Println("Error setting response:", err)
// 		return nil
// 	}

// 	// 执行 JavaScript 代码
// 	value, err := vm.RunString(afterCode)
// 	if err != nil {
// 		fmt.Println("Error executing JavaScript:", err)
// 		return nil
// 	}

// 	// 输出处理后的结果
// 	fmt.Println("✅ Processed response:", value)
// 	fmt.Println("🚪 Exiting VisitAfterSection")
// 	return value
// }

// 重写某规则访问方法，示例
func (v *MyVisitor) VisitPromptDef(ctx *promptdsl.PromptDefContext) interface{} {
	fmt.Println("解释执行 PromptDef:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func main() {

	input := `prompt SplitSolutionSteps {
    input {
            question: string;
        }

    output {
        format: json;
        type step struct {
            Conditions: []string;
            KnowledgePoint: string;
            ProcessResult: string;
        }
        schema: []step;
    }
        
    system {
        "你是一个擅长拆分解题步骤的数学老师，当前任务是将解题步骤进行合理拆分"
    }

    user {
        "请根据以下输入题目及其解答内容，将完整的解答过程拆分为多个“短链”，每个“短链”包含以下三个要素："
        "条件，知识点，结果"
        "请将输出内容严格按照以下格式返回："
        output.schema
        "特别提醒：本题可能涉及 extra_hint，请根据步骤合理提取对应知识点。"
        "以下是输入内容，请据此拆解："
        in.question
    }

    note {
        "每一步的条件必须完整、明确，不能模糊；"
        "知识点应尽量具体、准确，体现数学推理方法；"
        "结论应清晰表达出推导结果；"
        "上一步结果可作为下一步条件。"
    }

	after {
		"// 假设变量 response 是大模型的原始回复\n" +
		"// 将回复转成大写并去掉前后空白\n" +
		"let processed = response.trim().toUpperCase();\n" +
		"processed;"
	}
}
`

	// 创建输入流
	inputStream := antlr.NewInputStream(input)
	lexer := promptdsl.NewPromptDSLLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)

	tokenStream.Fill()
	tokens := tokenStream.GetAllTokens() // 使用 GetAllTokens() 而不是 GetTokens()
	for _, token := range tokens {
		if token.GetTokenType() != antlr.TokenEOF {
			fmt.Printf("Token: %s, text: %q\n",
				lexer.SymbolicNames[token.GetTokenType()],
				token.GetText())
		}
	}

	// 解析，调用顶层规则
	parser := promptdsl.NewPromptDSLParser(tokenStream)
	tree := parser.PromptFile()
	// 打印语法树文本结构
	fmt.Println(tree.ToStringTree(nil, parser))

	// 创建访问者，访问解析树
	// visitor := &MyVisitor{}
	visitor := &MyVisitor{
		BasePromptDSLVisitor: &promptdsl.BasePromptDSLVisitor{},
	}
	fmt.Println("🔍 Visiting Syntax Tree...")
	visitor.Visit(tree)
	// fmt.Println("🌳 Syntax Tree:")
	// fmt.Println(tree.ToStringTree(nil, parser))
}
