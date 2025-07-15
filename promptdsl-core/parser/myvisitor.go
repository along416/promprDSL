package parser // PromptDSL

import (
	"encoding/json"
	"fmt"
	// "log"
	"regexp"
	"strings"

	// "github.com/dop251/goja"
	"github.com/antlr4-go/antlr/v4"
	// "github.com/dop251/goja"
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
func (v *MyVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	// 默认啥都不做，直接返回
	return nil
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
	raw := section.GetText()
	return strings.TrimSpace(
		strings.TrimSuffix(strings.TrimPrefix(raw, "<after>"), "</after>"),
	)
}

func (v *MyVisitor) GeneratePromptf(inputContent string) string {
	finalPrompt := fmt.Sprintf(`system:
%s

user:
%s

input:
%s
`, v.SystemText, v.UserText, inputContent)
	if v.AfterJS != "" {
		finalPrompt += fmt.Sprintf("\nafter:\n%s\n", v.AfterJS)
	}
	return finalPrompt
}

// var _ PromptDSLVisitor = (*MyVisitor)(nil)
// 拼接最终的 Prompt

func (v *MyVisitor) GeneratePrompt() string {
	return v.GeneratePromptf("")
}

func (v *MyVisitor) Visit(tree antlr.ParseTree) interface{} {
	fmt.Println("🔍 Visiting Syntax Tree...")
	return tree.Accept(v)
}

//	func (v *MyVisitor) VisitPromptBlock(ctx *PromptBlockContext) interface{} {
//		// fmt.Println("📦 Entering PromptBlock")
//		return v.VisitChildren(ctx)
//	}
func (v *MyVisitor) VisitPromptBlock(ctx PromptBlockContext) interface{} {

	fmt.Println("📦 Entering PromptBlock")
	switch {
	case ctx.InputSection() != nil:
		v.Visit(ctx.InputSection())
		fmt.Printf("\n📦 Entering InputSection：\n%v\n",ctx.InputSection())
	case ctx.OutputSection() != nil:
		// 可忽略或保留结构定义
	case ctx.SystemSection() != nil:
		// v.SystemText = extractTextLines(ctx.SystemSection().AllTextLine())
	case ctx.UserSection() != nil:
		// v.UserText = extractTextLines(ctx.UserSection().AllTextLine())
	case ctx.NoteSection() != nil:
		// 可以拼到 user prompt 结尾作为注意事项
	case ctx.AfterSection() != nil:
		v.AfterJS = extractAfterCode(ctx.AfterSection())
	}
	return nil
}

// func (v *MyVisitor) VisitPromptFile(ctx *PromptFileContext) interface{} {
// 	var children []Node
// 	for _, block := range ctx.AllPromptBlock() {
// 		node := v.Visit(block)
// 		if node != nil {
// 			children = append(children, node.(Node))
// 		}
// 	}
// 	return &RootNode{Children: children}
// }

func (v *MyVisitor) VisitPromptDef(ctx *PromptDefContext) interface{} {
	return v.VisitChildren(ctx)
}

// func (v *MyVisitor) VisitPromptBlock(ctx *PromptBlockContext) interface{} {
// 	return v.VisitChildren(ctx)
// }


func (v *MyVisitor) VisitModuleDef(ctx *ModuleDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitModuleContent(ctx *ModuleContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitInputSection(ctx *InputSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitOutputSection(ctx *OutputSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitOutputStruct(ctx *OutputStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitOutputMarkdown(ctx *OutputMarkdownContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitSystemSection(ctx *SystemSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitUserSection(ctx *UserSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitUserContent(ctx *UserContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitNoteSection(ctx *NoteSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitDslCallExpr(ctx *DslCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

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

func (v *MyVisitor) VisitAfterSection(ctx *AfterSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *MyVisitor) VisitFixSection(ctx *FixSectionContext) interface{} {
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

