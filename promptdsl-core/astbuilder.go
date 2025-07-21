// promptdsl-core/astbuilder.go
package promptdslcore

import (
	"fmt"
	"promptdslcore/parser"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func BuildAST(parseTree *parser.PromptFileContext, tokens *antlr.CommonTokenStream) *RootNode {
	result := &RootNode{
		Vars:             make(map[string]interface{}), // 初始化空map
		SysNodes:         []Node{},
		UserNodes:        []Node{},
		ModuleDefs:       map[string][]Node{}, // 初始化 map
		InFields:         []FieldDef{},
		OutFields:        []FieldDef{},
		BeforeCode:       "",
		AfterCode:        []string{},
		FixCode:          []string{},
		BeforeNodes:      []Node{},
		outputspectNodes: OutputSpecNode{},
	}
	fmt.Println("Building AST...")
	def := parseTree.PromptDef(0)
	for _, block := range def.AllPromptBlock() {
		fmt.Println("  Processing block:")
		child := block.GetChild(0)
		switch b := child.(type) {
		case *parser.SystemSectionContext:
			if len(b.AllID()) > 0 {
				// ID 模式，生成 ModuleRefNode
				for _, id := range b.AllID() {
					modName := id.GetText()
					result.SysNodes = append(result.SysNodes, &ModuleRefNode{Name: modName})
				}
				fmt.Println("result.SysNodes:", result.SysNodes)
			} else {
				for _, content := range b.AllSysContent() {
					result.SysNodes = append(result.SysNodes, BuildSysNodesC(content))
				}
			}
		case *parser.ModuleDefContext:
			// 收集模块内容
			modName := b.ID().GetText()
			var contentNodes []Node
			for _, mc := range b.AllModuleContent() {
				mcCtx := mc.(*parser.ModuleContentContext)
				for i := 0; i < mcCtx.GetChildCount(); i++ {
					child := mcCtx.GetChild(i)
					switch sub := child.(type) {
					case *parser.TextLineContext:
						contentNodes = append(contentNodes, &StringNode{Val: cleanQuotes(sub.GetText())})
					case *parser.ParamPathContext:
						contentNodes = append(contentNodes, &ParamNode{Path: cleanQuotes(sub.GetText())})
					case *parser.IfStatementContext:
						contentNodes = append(contentNodes, buildIfNode(sub))
					case *parser.ExprContext:
						contentNodes = append(contentNodes, &StringNode{Val: cleanQuotes(sub.GetText())}) //表达式
					default:
						continue
					}
				}
			}
			result.ModuleDefs[modName] = contentNodes
			// fmt.Println("contentNodes:",contentNodes)
		case *parser.UserSectionContext:
			for _, uc := range b.AllUserContent() {
				node := buildUserNode(uc)
				result.UserNodes = append(result.UserNodes, node)
				if outSpec, ok := node.(*OutputSpecNode); ok {
					result.outputspectNodes = *outSpec
				}
			}
		case *parser.InputSectionContext:
			// 解析输入字段，放到 result.InDef
			var fields []FieldDef
			for _, param := range b.AllFieldDef() {
				name := param.ID().GetText()
				typ := param.Type_().GetText()

				// 解析注解
				var annotations []string
				for _, ann := range param.AllAnnotation() {
					if ann.AnnotationArgs() != nil {
						for _, v := range ann.AnnotationArgs().AllAnnotationValue() {
							if s := v.STRING(); s != nil {
								annotations = append(annotations, strings.Trim(s.GetText(), "\""))
							} else if arr := v.ArrayLiteral(); arr != nil {
								for _, s := range arr.AllSTRING() {
									annotations = append(annotations, strings.Trim(s.GetText(), "\""))
								}
							}
						}
					}
				}

				fields = append(fields, FieldDef{
					Name:        name,
					Type:        typ,
					JsonName:    name,
					Annotations: annotations,
				})
			}
			inNode := &InputNode{Fields: fields}
			result.InFields = inNode.Fields

		case *parser.OutputSectionContext:

			// 解析输出字段，放到 result.OutDef
			// 检查是哪种 output 类型
			//所以这里其实是构造结构体，而不是处理outputspec，还是都处理？
			fmt.Println("OutputSection")
			defaultAnnotations := b.AllDefaultAnnotation()
			// 先构建 defaultAnnotation map，方便查找
			defaultAnnoMap := map[string][]string{}
			for _, defAnn := range defaultAnnotations {
				name := defAnn.ID().GetText()
				var vals []string
				if defAnn.AnnotationArgs() != nil {
					for _, v := range defAnn.AnnotationArgs().AllAnnotationValue() {
						if s := v.STRING(); s != nil {
							raw := s.GetText()
							unquoted, err := strconv.Unquote(raw)
							if err != nil {
								unquoted = raw
							}
							vals = append(vals, unquoted)
						} else if arr := v.ArrayLiteral(); arr != nil {
							var parts []string
							for _, s := range arr.AllSTRING() {
								raw := s.GetText()
								unquoted, err := strconv.Unquote(raw)
								if err != nil {
									unquoted = raw
								}
								parts = append(parts, unquoted)
							}
							vals = append(vals, strings.Join(parts, ","))
						}
					}
				}
				defaultAnnoMap[name] = append(defaultAnnoMap[name], vals...)
			}
			if structCtx := b.OutputStruct(); structCtx != nil {
				var fields []FieldDef
				for _, field := range structCtx.AllFieldDef() {
					name := field.ID().GetText()
					typ := field.Type_().GetText()

					// 解析注解
					var annotations []string
					jsonName := name // 默认就是字段名

					for _, ann := range field.AllAnnotation() {
						annName := ann.ID().GetText()
						if ann.AnnotationArgs() != nil {
							for _, v := range ann.AnnotationArgs().AllAnnotationValue() {
								var val string

								if s := v.STRING(); s != nil {
									raw := s.GetText()
									unquoted, err := strconv.Unquote(raw)
									if err != nil {
										unquoted = raw
									}
									val = unquoted
								} else if arr := v.ArrayLiteral(); arr != nil {
									// 拼接数组内容
									var parts []string
									for _, s := range arr.AllSTRING() {
										raw := s.GetText()
										unquoted, err := strconv.Unquote(raw)
										if err != nil {
											unquoted = raw
										}
										parts = append(parts, unquoted)
									}
									val = strings.Join(parts, ",") // 或保留原结构
								}

								if annName == "jsonname" {
									jsonName = val
								} else {
									annotations = append(annotations, val)
								}
							}
						}
						// 再加 defaultAnnotation 中对应注解的参数（如果有）
						if defVals, ok := defaultAnnoMap[annName]; ok {
							annotations = append(annotations, defVals...)
						}
					}

					fields = append(fields, FieldDef{
						Name:        name,
						Type:        typ,
						JsonName:    jsonName,
						Annotations: annotations,
					})
				}

				outNode := &OutputNode{Fields: fields}
				result.OutFields = outNode.Fields
				// result.SysNodes = append(result.SysNodes, outNode)

			} else if mdCtx := b.OutputMarkdown(); mdCtx != nil {
				text := mdCtx.MARKDOWN().GetText()
				mdNode := &MarkdownNode{Content: cleanQuotes(text)}
				result.SysNodes = append(result.SysNodes, mdNode)
			}

		// case *parser.BeforeSectionContext:
		// 	for _, bc := range b.AllBeforeContent() {
		// 		node := buildNodeFromBeforeContent(bc) // 你自定义的函数，返回 Node 接口实现
		// 		result.BeforeNodes = append(result.BeforeNodes, node)
		// 	}
		case *parser.AfterSectionContext:
			result.AfterCode = extractRawText(b, tokens)

		case *parser.FixSectionContext:
			result.FixCode = extractRawText(b, tokens)
		}

	}

	return result
}

// // 假设 parser.BeforeContentContext 是 before 代码块中子内容的上下文类型
// func buildNodeFromBeforeContent(ctx parser.IBeforeContentContext) Node {
// 	// 这里根据 ctx 具体类型做判断，生成对应的 AST Node

// 	switch c := ctx.(type) {
// 	case *parser.TextLineContext:
// 		// 比如文本节点
// 		return &StringNode{Val: cleanQuotes(c.GetText())}
// 	case *parser.ParamPathContext:
// 		// 路径节点
// 		return &ParamNode{Path: cleanQuotes(c.GetText())}
// 	case *parser.IfStatementContext:
// 		// 条件节点，调用你已有的构造函数
// 		return buildIfNode(c)
// 	case *parser.ExprContext:
// 		// 表达式节点
// 		return &StringNode{Val: cleanQuotes(c.GetText())}
// 	// 你根据实际情况补充更多类型
// 	default:
// 		// 不支持的类型直接返回文本节点，避免panic
// 		return &StringNode{Val: cleanQuotes(ctx.GetText())}
// 	}
// }

func BuildSysNodes(root antlr.Tree) []Node {
	moduleContentMap := make(map[string][]Node)

	// Step 1: 收集所有 moduleDef
	var collectModules func(antlr.Tree)
	collectModules = func(node antlr.Tree) {
		switch ctx := node.(type) {
		case *parser.ModuleDefContext:
			moduleName := ctx.ID().GetText()
			var nodes []Node
			for _, mc := range ctx.AllModuleContent() {
				for i := 0; i < mc.(antlr.Tree).GetChildCount(); i++ {
					child := mc.(antlr.Tree).GetChild(i)
					switch t := child.(type) {
					case *parser.TextLineContext:
						nodes = append(nodes, &StringNode{Val: cleanQuotes(t.GetText())})
					}
				}
			}
			moduleContentMap[moduleName] = nodes
		}
		for i := 0; i < node.GetChildCount(); i++ {
			collectModules(node.GetChild(i))
		}
	}
	collectModules(root)

	// Step 2: 提取 sysSection 中的模块名，展开为 AST 节点
	var result []Node
	var expandSys func(antlr.Tree)
	expandSys = func(node antlr.Tree) {
		switch ctx := node.(type) {
		case *parser.SystemSectionContext:
			for _, id := range ctx.AllID() {
				modName := id.GetText()
				if nodes, ok := moduleContentMap[modName]; ok {
					result = append(result, nodes...) // 展开合并
				} else {
					// 未定义模块也可以加提示节点
					result = append(result, &StringNode{Val: fmt.Sprintf("[Missing module: %s]", modName)})
				}
			}
		}
		for i := 0; i < node.GetChildCount(); i++ {
			expandSys(node.GetChild(i))
		}
	}
	expandSys(root)

	return result
}
func BuildSysNodesC(ctx parser.ISysContentContext) Node {
	sysCtx := ctx.(*parser.SysContentContext)
	for i := 0; i < sysCtx.GetChildCount(); i++ {
		child := sysCtx.GetChild(i)
		switch sub := child.(type) {
		case *parser.TextLineContext:
			return &StringNode{Val: cleanQuotes(sub.GetText())}
		case *parser.IfStatementContext:
			return buildIfNode(sub)
		case *parser.ExprContext:
			return &StringNode{Val: cleanQuotes(sub.GetText())} // 临时
		default:
			continue
		}
	}
	return nil
}

// 构建一个 userContent 的 Node
func buildUserNode(ctx parser.IUserContentContext) Node {
	// fmt.Println("😊buildUserNode:")
	userCtx := ctx.(*parser.UserContentContext)
	// 优先判断 ARRAY_OUTPUTSPEC（形如 []outputspec）
	if userCtx.ARRAY_OUTPUTSPEC() != nil {
		text := userCtx.ARRAY_OUTPUTSPEC().GetText()
		fmt.Println("😊ARRAY_OUTPUTSPEC:", text)

		// 直接去掉前缀 []，拿到实际类型名
		rawType := strings.TrimPrefix(text, "[]")
		return &OutputSpecNode{
			IsArray: true,
			RawTyp:  rawType,
		}
	}

	// 普通 OUTPUTSPEC（非数组形式）
	if userCtx.OUTPUTSPEC() != nil {
		text := userCtx.OUTPUTSPEC().GetText()
		fmt.Println("😊OUTPUTSPEC:", text)

		return &OutputSpecNode{
			IsArray: false,
			RawTyp:  text,
		}
	}

	for i := 0; i < userCtx.GetChildCount(); i++ {
		child := userCtx.GetChild(i)
		switch sub := child.(type) {
		case *parser.ParamPathContext:
			// fmt.Println("😊param path:", sub.GetText())
			return &ParamNode{Path: cleanQuotes(sub.GetText())}
		case *parser.TextLineContext:
			fmt.Println("😊stringtext:", sub.GetText())
			return &StringNode{Val: cleanQuotes(sub.GetText())}
		case *parser.IfStatementContext:
			fmt.Println("😊IfStatementContext:", sub.GetText())
			return buildIfNode(sub)
		case *parser.ExprContext:
			return &StringNode{Val: cleanQuotes(sub.GetText())} // 临时
		default:
			continue
		}
	}
	return nil
}

func buildIfNode(ctx *parser.IfStatementContext) *IfNode {
	condList := ctx.Condition()
	// 断言接口为具体类型
	cctx := condList.(*parser.ConditionContext)
	condExpr := buildExprFromCondition(cctx)

	var thenNodes []Node
	// for _, uc := range ctx.AllUserContent() {
	// 	thenNodes = append(thenNodes, buildUserNode(uc))
	// }
	// usercontentls := ctx.AllUserContent()

	Thencontent := ctx.AllThencontent()
	for _, uc := range Thencontent {
		thenNodes = append(thenNodes, buildUserNode(uc.UserContent()))
	}
	// thenNodes = append(thenNodes, buildUserNode(thenNode))
	// fmt.Println("😐thenNodes", thenNodes)
	// for i, n := range thenNodes {
	// 	fmt.Printf("thenNode[%d]: %#v\n", i, n)
	// }
	var elseNodes []Node
	// 遍历 else 分支
	elsecontent := ctx.AllElsecontent()
	for _, uc := range elsecontent {
		elseNodes = append(elseNodes, buildUserNode(uc.UserContent()))
	}
	// for i := 0; i < ctx.GetChildCount(); i++ {
	// 	child := ctx.GetChild(i)
	// 	if ruleCtx, ok := child.(parser.IUserContentContext); ok {
	// 		elseNodes = append(elseNodes, buildUserNode(ruleCtx))
	// 	}

	// }

	return &IfNode{
		Condition: *condExpr,
		Then:      thenNodes,
		Else:      elseNodes,
	}
}

func buildExprFromCondition(ctx *parser.ConditionContext) *Expr {
	if ctx.GetOp() != nil {
		lhs := buildExpr(ctx.GetLhs())
		rhs := buildExpr(ctx.GetRhs())
		var op ExprOp
		switch ctx.GetOp().GetText() {
		case "==":
			op = ExprOp_Equal
		case "!=":
			op = ExprOp_NotEqual
		default:
			op = ExprOp_None
		}
		return &Expr{
			Op:       op,
			Operant0: []*Expr{lhs},
			Operant1: []*Expr{rhs},
		}
	}
	if single := ctx.GetSingle(); single != nil {
		return buildExpr(single)
	}
	return &Expr{Op: ExprOp_None}
}

func buildExpr(exprCtx parser.IExprContext) *Expr {
	switch expr := exprCtx.(type) {
	case *parser.ExprContext:
		if param := expr.ParamPath(); param != nil {
			paramName := getParamName(param)
			return &Expr{Leaf: &paramName}
		}
		if str := expr.STRING(); str != nil {
			s := strings.Trim(str.GetText(), "\"")
			return &Expr{Leaf: &s}
		}
		if num := expr.NUMBER(); num != nil {
			n := num.GetText()
			return &Expr{Leaf: &n}
		}
		if b := expr.BOOL(); b != nil {
			val := b.GetText()
			return &Expr{Leaf: &val}
		}
	}
	return &Expr{Op: ExprOp_None}
}

func cleanQuotes(s string) string {
	return strings.Trim(s, "\"")
}

func getParamName(p parser.IParamPathContext) string {
	parts := []string{}
	for _, id := range p.AllID() {
		name := id.GetText()
		if name != "in" {
			parts = append(parts, name)
		}
	}
	return strings.Join(parts, ".")
}
