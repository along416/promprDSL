package promptdsl

import (
	"encoding/json"
	"fmt"
	// "log"
	"regexp"
	"strings"

	// "github.com/dop251/goja"
	"github.com/antlr4-go/antlr/v4"
	"github.com/dop251/goja"
)

type MyVisitor struct {
	*BasePromptDSLVisitor
	SystemText string
	UserText   string
	InputVars  map[string]string
	AfterJS    string
}

func NewMyVisitor() *MyVisitor {
	return &MyVisitor{
		InputVars: make(map[string]string),
	}
}

// 转义单斜杠
func escapeInvalidJSONBackslashes(s string) string {
	var buf strings.Builder
	n := len(s)

	for i := 0; i < n; i++ {
		if s[i] == '\\' {
			if i+1 < n {
				next := s[i+1]
				if strings.ContainsRune(`\"ntbrfu/`, rune(next)) {
					// 是合法的转义序列，保留原样
					buf.WriteByte(s[i])
					buf.WriteByte(next)
					i++ // skip next
				} else {
					// 非法反斜杠，转义为两个反斜杠
					buf.WriteString(`\\`)
				}
			} else {
				// 末尾只有一个 \
				buf.WriteString(`\\`)
			}
		} else {
			buf.WriteByte(s[i])
		}
	}
	return buf.String()
}
func ExtractJSON(text string, ret any) error {
	// fmt.Printf("源文本:%v", text)
	re := regexp.MustCompile("(?s)```json\\s*(.*?)\\s*```")
	matches := re.FindStringSubmatch(text)
	if len(matches) < 2 {
		return fmt.Errorf("no JSON data found")
	}
	// log.Printf("\nmatches%v\n",matches)
	jsonStr := escapeInvalidJSONBackslashes(matches[1])
	// var result interface{}
	// 防止反斜杠导致json.Unmarshal失败
	if err := json.Unmarshal([]byte(jsonStr), ret); err != nil {
		return fmt.Errorf("invalid JSON format: %v", err)
	}
	return nil
}

// 辅助函数：合并 textLine 列表为字符串
func extractTextLines(lines []ITextLineContext) string {
	var sb strings.Builder
	for _, line := range lines {
		sb.WriteString(line.GetText())
		sb.WriteString("\n")
	}
	return sb.String()
}

// 提取 After 中 JS 代码
func extractAfterCode(section IAfterSectionContext) string {
	// var sb strings.Builder
	// entries := section.AfterContent().AllAfterEntry()
	// for _, entry := range entries {
	// 	code := entry.GetText()
	// 	sb.WriteString(strings.Trim(code, "`"))
	// 	sb.WriteString("\n") // 你可以根据需要加入换行
	// }
	return ""
}

// var _ PromptDSLVisitor = (*MyVisitor)(nil)
// 拼接最终的 Prompt
func (v *MyVisitor) GeneratePromptf(inputContent string) string {
	// 拼接 system、user、input 等部分
	finalPrompt := fmt.Sprintf(`system:
%s

user:
%s

input:
%s
`, v.SystemText, v.UserText, inputContent)

	// 你还可以在最后加上 AfterJS 的后处理部分
	if v.AfterJS != "" {
		finalPrompt += fmt.Sprintf("\nafter:\n%s\n", v.AfterJS)
	}

	return finalPrompt
}
func (v *MyVisitor) GeneratePrompt() string {
	return v.GeneratePromptf("")
}
func (v *MyVisitor) Visit(tree antlr.ParseTree) interface{} {
	// fmt.Println("🔍 Visiting Syntax Tree...")
	return tree.Accept(v)
}

//	func (v *MyVisitor) VisitPromptBlock(ctx *PromptBlockContext) interface{} {
//		// fmt.Println("📦 Entering PromptBlock")
//		return v.VisitChildren(ctx)
//	}
func (v *MyVisitor) VisitPromptBlock(ctx PromptBlockContext) interface{} {
	switch {
	case ctx.InputSection() != nil:
		v.Visit(ctx.InputSection())
	case ctx.OutputSection() != nil:
		// 可忽略或保留结构定义
	case ctx.SystemSection() != nil:
		v.SystemText = extractTextLines(ctx.SystemSection().AllTextLine())
	case ctx.UserSection() != nil:
		// v.UserText = extractTextLines(ctx.UserSection().AllTextLine())
	case ctx.NoteSection() != nil:
		// 可以拼到 user prompt 结尾作为注意事项
	case ctx.AfterSection() != nil:
		v.AfterJS = extractAfterCode(ctx.AfterSection())
	}
	return nil
}

func (v *MyVisitor) VisitAfterSection(ctx *AfterSectionContext) interface{} {
	fmt.Println("🚪 Entering VisitAfterSection")
	vm := goja.New()
	//模型回复
	response := `[{"Conditions": ["已知a+b=5"], "KnowledgePoint": "代数加法性质", "ProcessResult": "a=5-b"}]`
	vm.Set("response", response)
	//如果是json，则调用转义
	// format:="json"
	// if format=="json" {
	// 	var content Checkcontent
	// 	err:=ExtractJSON(response, &content)
	// 	if err != nil {
	// 		// log.Printf("task1: %s\nerror:解析失败❌%v\n", t.id, err)
	// 		// return lp.TaskResultFailedRetry
	// 	}
	// }
	//md
	//else

	//jscode
	// var jsCode string
	// afterContent := ""
	// fmt.Printf("📦 AfterContent text: %s\n", afterContent.GetText())
	// if jsBlock := afterContent.JAVASCRIPT_BLOCK(); jsBlock != nil {
	// 	// 去掉反引号（`）包裹
	// 	raw := jsBlock.GetText()
	// 	jsCode = strings.Trim(raw, "`")
	// } else {
	// 	// 多个 entry 拼接
	// 	for _, entry := range afterContent.AllAfterEntry() {
	// 		text := entry.GetText()
	// 		if strings.HasPrefix(text, "`") {
	// 			text = strings.Trim(text, "`")
	// 		} else {
	// 			text = strings.Trim(text, "\"") // 也可能是字符串
	// 		}
	// 		jsCode += text + "\n"
	// 	}
	// }

	// value, err := vm.RunString(jsCode)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("结果:", value.String())
	return v.VisitChildren(ctx)
}


func (v *MyVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	fmt.Println("📦 VisitChildren")
	n := node.GetChildCount()
	for i := 0; i < n; i++ {
		child := node.GetChild(i)
		if ptree, ok := child.(antlr.ParseTree); ok {
			ptree.Accept(v)
		}
	}
	return nil
}

// func (v *MyVisitor) VisitPromptDef(ctx *PromptDefContext) interface{} {
// 	// fmt.Println("🌳 Syntax Tree:myvisitor")
// 	// log.Println("解释执行 PromptDef:", ctx.GetText())
// 	return v.VisitChildren(ctx)
// }

func (v *MyVisitor) VisitPromptDef(ctx *PromptDefContext) interface{} {
	for _, block := range ctx.AllPromptBlock() {
		v.Visit(block)
	}
	return nil
}

func (v *MyVisitor) VisitPromptFile(ctx *PromptFileContext) interface{} {
	// fmt.Println("🌳 Syntax Tree:VisitPromptFile")
	return v.VisitChildren(ctx)
}
func (v *MyVisitor) VisitInputSection(ctx *InputSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitOutputSection(ctx *OutputSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

// func (v *MyVisitor) VisitOutputEntry(ctx *OutputEntryContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

func (v *MyVisitor) VisitSystemSection(ctx *SystemSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitUserSection(ctx *UserSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitNoteSection(ctx *NoteSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

//	func (v *MyVisitor) VisitAfterSection(ctx *AfterSectionContext) interface{} {
//		// fmt.Println("🌳 Syntax Tree:")
//		return v.VisitChildren(ctx)
//	}
func (v *MyVisitor) VisitFieldDef(ctx *FieldDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitTextLine(ctx *TextLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitParamPath(ctx *ParamPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitStructDef(ctx *StructDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitAnnotationArgs(ctx *AnnotationArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitAnnotationValue(ctx *AnnotationValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitTextBlock(ctx *TextBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitFormatType(ctx *FormatTypeContext) interface{} {
	return v.VisitChildren(ctx)
}
