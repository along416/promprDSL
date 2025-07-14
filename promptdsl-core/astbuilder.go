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
		SysNodes:  []Node{},
		UserNodes: []Node{},
	}

	for _, def := range parseTree.AllPromptDef() {
		for _, block := range def.AllPromptBlock() {
			child := block.GetChild(0)
			switch b := child.(type) {
			case *parser.SysSectionContext:
				// 收集 sys 使用的模块名
				result.SysNodes = append(result.UserNodes, BuildSysNodes(b))
			case *parser.ModuleDefContext:
				// 收集模块内容

			case *parser.UserSectionContext:
				for _, uc := range b.AllUserContent() {
					result.UserNodes = append(result.UserNodes, buildUserNode(uc))
				}
			case *parser.SystemSectionContext:
				// system 文字内容处理

			case *parser.InputSectionContext:
				// 解析输入字段，放到 result.InDef

			case *parser.OutputSectionContext:
				// 解析输出字段，放到 result.OutDef

			case *parser.AfterSectionContext:
				// result.AfterCode = extractRawText(b)

			case *parser.FixSectionContext:
				// result.FixCode = extractRawText(b)
			}
		}
	}

	return result
}

// sysnode
//
//	func buildUSysNode(ctx parser.ISysSectionContext) Node {
//		sysCtx := ctx.(*parser.SysSectionContext)
//		for i := 0; i < sysCtx.GetChildCount(); i++ {
//			child := sysCtx.GetChild(i)
//			switch sub := child.(type) {
//			case *parser.TextLineContext:
//				return &StringNode{Val: cleanQuotes(sub.GetText())}
//			case *parser.IfStatementContext:
//				return buildIfNode(sub)
//			case *parser.ExprContext:
//				return &StringNode{Val: cleanQuotes(sub.GetText())} // 临时
//			default:
//				continue
//			}
//		}
//		return nil
//	}
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
		case *parser.SysSectionContext:
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

// 构建一个 userContent 的 Node
func buildUserNode(ctx parser.IUserContentContext) Node {
	userCtx := ctx.(*parser.UserContentContext)
	for i := 0; i < userCtx.GetChildCount(); i++ {
		child := userCtx.GetChild(i)
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

func buildIfNode(ctx *parser.IfStatementContext) *IfNode {
	condList := ctx.Condition()
	// 断言接口为具体类型
	cctx := condList.(*parser.ConditionContext)
	condExpr := buildExprFromCondition(cctx)

	var thenNodes []Node
	for _, uc := range ctx.AllUserContent() {
		thenNodes = append(thenNodes, buildUserNode(uc))
	}

	var elseNodes []Node
	// 遍历 else 分支
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		if ruleCtx, ok := child.(parser.IUserContentContext); ok {
			elseNodes = append(elseNodes, buildUserNode(ruleCtx))
		}
	}

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
