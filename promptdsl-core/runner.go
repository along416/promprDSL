// promptdsl-core/runner.go
package promptdslcore

import (
	"fmt"
	"promptdslcore/parser"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func RunPromptDSL(input string) (string, error) {

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
	rootNode := BuildAST(ctx)
	fmt.Printf("%v\n", rootNode)

	// 3. 构造 Eval 上下文
	str := &PromptEvalContext{
		InFields:   rootNode.InFields,
		OutFields:  rootNode.OutFields,
		ModuleDefs: rootNode.ModuleDefs,
		Input: map[string]any{
			"question": "参数内容",
			"level":    "easy",
		},
	}

	// 4. 执行 AST，得到 prompt 字符串
	outputParts, err := rootNode.Eval(str)
	if err != nil {
		return "", err
	}

	return strings.Join(outputParts, "\n"), nil
}











// func extractParts(node antlr.ParseTree) (systemTexts, userTexts []string, inputs map[string]string) {
// 	inputs = make(map[string]string)
// 	// var sysBuilder, userBuilder strings.Builder
// 	// 内部递归函数
// 	var walk func(antlr.Tree)
// 	walk = func(n antlr.Tree) {
// 		switch ctx := n.(type) {
// 		case *parser.SystemSectionContext:
// 			// 把 systemSection 的所有 textLine 拼起来
// 			var lines []string
// 			for _, tl := range ctx.AllTextLine() {
// 				text := cleanQuotes(tl.GetText())
// 				lines = append(lines, text)
// 			}
// 			systemTexts = append(systemTexts, strings.Join(lines, "\n"))

// 		case *parser.UserSectionContext:
// 			// userSection 里也可能是多段 userContent，这里只简单拼接所有 textLine 文本
// 			var lines []string
// 			inputs := map[string]string{
// 				"question": "这是输入的题目内容",
// 				"level":    "easy",
// 			}

// 			for _, uc := range ctx.AllUserContent() {
// 				lines = append(lines, extractUserContentText(uc, inputs))
// 			}
// 			userTexts = append(userTexts, strings.Join(lines, "\n"))

// 		case *parser.InputSectionContext:
// 			// inputSection 包含多个 fieldDef，提取字段名和类型
// 			for _, fd := range ctx.AllFieldDef() {
// 				fieldName := fd.ID().GetText()
// 				fieldType := fd.Type_().GetText()
// 				inputs[fieldName] = fieldType
// 			}
// 		}

// 		// 递归遍历子节点
// 		for i := 0; i < n.GetChildCount(); i++ {
// 			child := n.GetChild(i)
// 			walk(child)
// 		}
// 	}
// 	walk(node)
// 	return
// }

// // userContent 有多种情况，这里演示对 textLine 的简单提取
// func extractUserContentText(ctx parser.IUserContentContext, inputs map[string]string) string {
// 	// 先断言成 *parser.UserContentContext
// 	userCtx, ok := ctx.(*parser.UserContentContext)
// 	if !ok {
// 		return ""
// 	}

// 	// 遍历 userContent 的所有子节点
// 	var texts []string
// 	for i := 0; i < userCtx.GetChildCount(); i++ {
// 		child := userCtx.GetChild(i)
// 		switch c := child.(type) {
// 		case antlr.TerminalNode:
// 			// 如果是终结符，比如字符串直接拿文本
// 			texts = append(texts, c.GetText())
// 		case antlr.ParserRuleContext:
// 			// 进一步判断是哪种具体类型，比如 IfStatementContext 或 TextLineContext
// 			switch subCtx := c.(type) {
// 			case *parser.TextLineContext:
// 				texts = append(texts, cleanQuotes(subCtx.GetText()))
// 			case *parser.IfStatementContext:
// 				cond := subCtx.Condition()

// 				var chosenBranch []parser.IUserContentContext
// 				if cond.GetLhs() != nil && cond.GetRhs() != nil {
// 					lhsExpr := cond.GetLhs()
// 					rhsExpr := cond.GetRhs()
// 					op := cond.GetOp().GetText()
// 					// 判断左边类型
// 					lhs := lhsExpr.(*parser.ExprContext)
// 					if lhs.ParamPath() != nil {
// 						lhsStr := lhs.ParamPath().GetText()
// 						paramName := strings.TrimPrefix(lhsStr, "in.")

// 						inputVal, exists := inputs[paramName]
// 						if !exists {
// 							inputVal = ""
// 						}

// 						rhsText := strings.Trim(rhsExpr.GetText(), "\"")

// 						var conditionTrue bool
// 						switch op {
// 						case "==":
// 							conditionTrue = inputVal == rhsText
// 						case "!=":
// 							conditionTrue = inputVal != rhsText
// 						}

// 						if conditionTrue {
// 							chosenBranch = subCtx.AllUserContent()
// 						} else {
// 							// 手动找到 else 分支的内容
// 							for i := 0; i < subCtx.GetChildCount(); i++ {
// 								child := subCtx.GetChild(i)
// 								if ruleCtx, ok := child.(antlr.ParserRuleContext); ok {
// 									if ruleCtx.GetText() == "else" && i+1 < subCtx.GetChildCount() {
// 										if block, ok := subCtx.GetChild(i + 1).(antlr.ParserRuleContext); ok {
// 											for j := 0; j < block.GetChildCount(); j++ {
// 												if uc, ok := block.GetChild(j).(parser.IUserContentContext); ok {
// 													chosenBranch = append(chosenBranch, uc)
// 												}
// 											}
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return strings.Join(texts, "\n")
// }

// // 去掉字符串两端的双引号，方便阅读
// func cleanQuotes(s string) string {
// 	return strings.Trim(s, "\"")
// }
// func extractSysCombined(root antlr.Tree) string {
// 	// 先遍历整棵树，收集所有 moduleDef 的文本，key是模块名
// 	moduleContentMap := make(map[string]string)

// 	var collectModules func(antlr.Tree)
// 	collectModules = func(node antlr.Tree) {
// 		switch ctx := node.(type) {
// 		case *parser.ModuleDefContext:
// 			moduleName := ctx.ID().GetText() // 模块名，比如 sys_role
// 			var texts []string
// 			for _, mc := range ctx.AllModuleContent() {
// 				// moduleContent 里可能是 textLine，也可能是 paramPath 等，这里只演示 textLine
// 				for i := 0; i < mc.(antlr.Tree).GetChildCount(); i++ {
// 					child := mc.(antlr.Tree).GetChild(i)
// 					switch t := child.(type) {
// 					case *parser.TextLineContext:
// 						texts = append(texts, cleanQuotes(t.GetText()))
// 					}
// 				}
// 			}
// 			moduleContentMap[moduleName] = strings.Join(texts, "\n")
// 		}
// 		for i := 0; i < node.GetChildCount(); i++ {
// 			collectModules(node.GetChild(i))
// 		}
// 	}
// 	collectModules(root)

// 	// 遍历 sysSection，按顺序取模块名，拼接对应文本
// 	var combined strings.Builder
// 	var combineSys func(antlr.Tree)
// 	combineSys = func(node antlr.Tree) {
// 		switch ctx := node.(type) {
// 		case *parser.SysSectionContext:
// 			for _, id := range ctx.AllID() {
// 				key := id.GetText()
// 				if val, ok := moduleContentMap[key]; ok {
// 					combined.WriteString(val)
// 				}
// 			}
// 		}
// 		for i := 0; i < node.GetChildCount(); i++ {
// 			combineSys(node.GetChild(i))
// 		}
// 	}
// 	combineSys(root)

// 	return combined.String()
// }
// func getParamName(p *parser.ParamPathContext) string {
// 	// 你可以拼接所有 ID，忽略 "." 和 "in"
// 	parts := []string{}
// 	for _, id := range p.AllID() {
// 		name := id.GetText()
// 		if name != "in" { // 忽略 in 前缀
// 			parts = append(parts, name)
// 		}
// 	}
// 	return strings.Join(parts, ".")
// }
