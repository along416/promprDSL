// promptdsl-core/astbuilder.go
package promptdslcore

import (
	"fmt"
	"promptdslcore/parser"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func BuildAST(parseTree *parser.PromptFileContext) *RootNode {
	result := &RootNode{
		SysNodes:   []Node{},
		UserNodes:  []Node{},
		ModuleDefs: map[string][]Node{}, // 初始化 map
		InFields:   []FieldDef{},
		OutFields:  []FieldDef{},
	}
	def := parseTree.PromptDef(0)
	for _, block := range def.AllPromptBlock() {
		child := block.GetChild(0)
		switch b := child.(type) {
		case *parser.SystemSectionContext:
			// fmt.Println("AllID:%v",b.AllID())
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
				result.UserNodes = append(result.UserNodes, buildUserNode(uc))
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
					annotations = append(annotations, ann.ID().GetText())
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
			if structCtx := b.OutputStruct(); structCtx != nil {
				var fields []FieldDef
				for _, field := range structCtx.AllFieldDef() {
					name := field.ID().GetText()
					typ := field.Type_().GetText()

					var annotations []string
					for _, ann := range field.AllAnnotation() {
						annotations = append(annotations, ann.ID().GetText())
					}

					fields = append(fields, FieldDef{
						Name:        name,
						Type:        typ,
						JsonName:    name,
						Annotations: annotations,
					})
				}

				outNode := &OutputNode{Fields: fields}
				result.OutFields = outNode.Fields
				result.SysNodes = append(result.SysNodes, outNode)

			} else if mdCtx := b.OutputMarkdown(); mdCtx != nil {
				text := mdCtx.MARKDOWN().GetText()
				mdNode := &MarkdownNode{Content: cleanQuotes(text)}
				result.SysNodes = append(result.SysNodes, mdNode)
			}

		case *parser.AfterSectionContext:
			// result.AfterCode = extractRawText(b)

		case *parser.FixSectionContext:
			// result.FixCode = extractRawText(b)
		}
	}

	return result
}

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
	elsecontent := ctx.AllThencontent()
	for _, uc := range elsecontent {
		thenNodes = append(elseNodes, buildUserNode(uc.UserContent()))
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
