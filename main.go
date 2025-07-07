package main

import (
	"fmt"
	"promptdsl_test/promptdsl" // 替换为你实际的包路径

	"github.com/antlr4-go/antlr/v4"
)

type MyVisitor struct {
	*promptdsl.BasePromptDSLVisitor
}

// 重写某规则访问方法，示例
func (v *MyVisitor) VisitPromptDef(ctx *promptdsl.PromptDefContext) interface{} {
	fmt.Println("解释执行 PromptDef:", ctx.GetText())
	return v.VisitChildren(ctx)
}

func main() {
	input := `prompt SplitSolutionSteps {
    params {
        in: {
            question: string;
        }
        output: {
            format: md | json;
            type step struct {
                Conditions: []string;
                KnowledgePoint: string;
                ProcessResult: string;
            }
            schema: []step;
        }
        temperature: float = 0.3;
    }

    system {
        你是一个擅长拆分解题步骤的数学老师，当前任务是将解题步骤进行合理拆分
    }

    user {
        请根据以下输入题目及其解答内容，将完整的解答过程拆分为多个“短链”，每个“短链”包含以下三个要素：
        条件，知识点，结果
        请将输出内容严格按照以下格式返回：
        params.output.schema
        特别提醒：本题可能涉及 extra_hint，请根据步骤合理提取对应知识点。
        以下是输入内容，请据此拆解：
        params.in.question
    }

    note {
        1. 每一步的条件必须完整、明确，不能模糊；
        2. 知识点应尽量具体、准确，体现数学推理方法；
        3. 结论应清晰表达出推导结果；
        4. 上一步结果可作为下一步条件。
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
	visitor := &MyVisitor{}
	visitor.Visit(tree)
}
