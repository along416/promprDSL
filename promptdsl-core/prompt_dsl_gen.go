package promptdslcore

import (
	"encoding/json"
	"fmt"
	"strings"
)

type FieldDef struct {
	Name        string
	Type        string
	JsonName    string
	Hint        string
	Annotations []string
}

type PromptEvalContext struct {
	InFields   []FieldDef
	OutFields  []FieldDef
	Input      any
	ModuleDefs map[string][]Node
}

type Node interface {
	Eval(_ *PromptEvalContext) ([]string, error)
	// ConvertToCode() string
}

// 引用结点

type ExprOp int

const (
	ExprOp_And ExprOp = iota
	ExprOp_Or
	ExprOp_Not
	ExprOp_Equal
	ExprOp_NotEqual
	ExprOp_GreaterThan
	ExprOp_LessThan
	ExprOp_GreaterThanOrEqual
	ExprOp_LessThanOrEqual
	ExprOp_Add
	ExprOp_Sub
	ExprOp_Mul
	ExprOp_Div
	ExprOp_Mod

	ExprOp_None // for leaf nodes
)

type Expr struct {
	Leaf *string

	Op       ExprOp
	Operant0 []*Expr
	Operant1 []*Expr
}

func (e *Expr) EvalToBool(ctx *PromptEvalContext) (bool, error) {
	// 如果是叶子节点，不能直接判断真假，必须是比较表达式
	if e.Leaf != nil {
		return false, fmt.Errorf("cannot evaluate leaf as boolean directly: %s", *e.Leaf)
	}

	// 支持的布尔比较运算
	switch e.Op {
	case ExprOp_Equal, ExprOp_NotEqual:
		leftExpr := e.Operant0[0]
		rightExpr := e.Operant1[0]

		leftVals, err := leftExpr.Eval(ctx)
		if err != nil {
			return false, err
		}
		rightVals, err := rightExpr.Eval(ctx)
		if err != nil {
			return false, err
		}

		left := firstOrEmpty(leftVals)
		right := firstOrEmpty(rightVals)
		fmt.Println("😶left:", left)
		fmt.Println("😶right:", right)
		// fmt.Println("🔍 Comparing:", left, e.Op, right)

		switch e.Op {
		case ExprOp_Equal:
			return left == right, nil
		case ExprOp_NotEqual:
			return left != right, nil
		}
	}

	return false, fmt.Errorf("unsupported boolean operator: %v", e.Op)
}
func (e *Expr) Eval(ctx *PromptEvalContext) ([]string, error) {
	// 如果是叶子节点
	if e.Leaf != nil {
		leafVal := *e.Leaf

		// 假设叶子是参数路径格式，比如 "input.question" 或 "output.answer"
		parts := strings.SplitN(leafVal, ".", 2)
		if len(parts) == 2 {
			section, key := parts[0], parts[1]
			var val any
			var ok bool

			switch section {
			case "input":
				// 假设 ctx.Input 是 map[string]any
				if m, okCast := ctx.Input.(map[string]any); okCast {
					val, ok = m[key]
				}
			case "output":
				// 你也可以在 ctx 中保存输出数据结构，这里示例用 OutFields 不一样，按实际需求调整
				// 这里仅示例直接返回 key
				val, ok = key, true
			default:
				return nil, fmt.Errorf("unknown section in leaf: %s", section)
			}

			if !ok {
				return []string{""}, nil // 找不到值返回空字符串，或者你也可以返回错误
			}
			return []string{fmt.Sprintf("%v", val)}, nil
		}

		// 如果不是带点的路径，就直接返回叶子内容
		return []string{leafVal}, nil
	}

	// 非叶子节点暂时不支持Eval
	return nil, fmt.Errorf("Eval not supported on non-leaf expressions yet")
}
func firstOrEmpty(list []string) string {
	if len(list) == 0 {
		return ""
	}
	return list[0]
}
func (e *Expr) EvalToInt(_ *PromptEvalContext) (int, error) {
	return 0, nil
}

func (e *Expr) EvalToFloat(_ *PromptEvalContext) (float32, error) {
	return 0.0, nil
}

func (e *Expr) EvalToString(_ *PromptEvalContext) (string, error) {
	return "", nil
}

type StringNode struct {
	Val string
}

func (node *StringNode) Eval(_ *PromptEvalContext) ([]string, error) {
	return []string{node.Val}, nil
}

type OutputSpecNode struct {
	Format string
}

func (node *OutputSpecNode) Eval(_ *PromptEvalContext) ([]string, error) {
	// TODO
	return []string{}, nil
}

type ModuleRefNode struct {
	Name string
}

//	type EvalContext struct {
//		ModuleDefs map[string][]Node
//		// 你可以根据需要添加其他字段，比如变量、环境等
//	}
//
// LookupModule 从上下文中查找模块定义,平铺moudlenode列表方便遍历
func (ctx *PromptEvalContext) LookupModule(name string) []Node {
	return ctx.ModuleDefs[name]
}
func (m *ModuleRefNode) Eval(ctx *PromptEvalContext) ([]string, error) {

	nodes := ctx.LookupModule(m.Name)
	if nodes == nil {
		return []string{fmt.Sprintf("[Missing module: %s]", m.Name)}, nil
		// return nil,nil
	}

	var result []string
	for _, node := range nodes {
		str, err := node.Eval(ctx)
		if err != nil {
			return nil, err
		}
		result = append(result, str...)
	}
	return result, nil
}

type InputNode struct {
	Fields []FieldDef
}

func (n *InputNode) Eval(ctx *PromptEvalContext) ([]string, error) {
	// Eval逻辑根据需要实现
	return nil, nil
}

type OutputNode struct {
	Fields []FieldDef
}

func (n *OutputNode) Eval(ctx *PromptEvalContext) ([]string, error) {
	// Eval逻辑根据需要实现
	return nil, nil
}

type MarkdownNode struct {
	Content string
}

func (m *MarkdownNode) Eval(ctx *PromptEvalContext) ([]string, error) {
	return []string{m.Content}, nil
}

type IfNode struct {
	Condition Expr
	Then      []Node
	Else      []Node
}

func (node *IfNode) Eval(ctx *PromptEvalContext) ([]string, error) {

	condResult, err := node.Condition.EvalToBool(ctx)
	if err != nil {
		return nil, err
	}

	var result []string
	if condResult {
		// fmt.Println("IfNode: condition is true")
		for _, n := range node.Then {
			texts, err := n.Eval(ctx)
			if err != nil {
				return nil, err
			}
			result = append(result, texts...)
		}
		// fmt.Println("IfNode: then done",node)
	} else {
		// fmt.Println("IfNode: condition is false")
		for _, n := range node.Else {
			texts, err := n.Eval(ctx)
			if err != nil {
				return nil, err
			}
			result = append(result, texts...)
		}
	}

	return result, nil
}

type ParamNode struct {
	Path string
}

// 实现 Node 接口的 Eval 方法
func (p *ParamNode) Eval(ctx *PromptEvalContext) ([]string, error) {

	var section, key string
	parts := strings.SplitN(p.Path, ".", 2)
	if len(parts) == 2 {
		section, key = parts[0], parts[1]
	} else if len(parts) == 1 {
		section = parts[0]
		key = "" 
	} else {
		return nil, fmt.Errorf("invalid param path: %s", p.Path)
	}
	// section, key := parts[0], parts[1]

	switch section {
	case "input":
		// var input string
		// 从 ctx.Input 中获取真实值
		// fmt.Println("解释执行 ParamNode:", ctx.Input)
		inputMap, ok := ctx.Input.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("input is not a valid map[string]any")
		}

		if val, ok := inputMap[key]; ok {
			return []string{fmt.Sprintf("%v", val)}, nil
		}
	case "output":
		if key == "" {
			// 取全部 OutFields，组装 JSON 返回
			b, err := json.Marshal(ctx.OutFields)
			if err != nil {
				return nil, err
			}
			return []string{string(b)}, nil
		}
		for _, field := range ctx.OutFields {
			if field.Name == key {
				return []string{fmt.Sprintf("[Output: %s of type %s]", field.Name, field.Type)}, nil
			}
		}

	default:
		return nil, fmt.Errorf("unknown param section: %s", section)
	}

	return nil, fmt.Errorf("param not found: %s", p.Path)
}

type CasePair struct {
	Case Expr
	Body []Node
}

type SwitchNode struct {
	Switch  Expr
	Cases   []CasePair
	Default Node
}

func (node *SwitchNode) Eval(_ *PromptEvalContext) ([]string, error) {
	// TODO
	return []string{}, nil
}

type ForNode struct {
	Init string
	Cond Expr
	Post string
	Body []Node
}

func (node *ForNode) Eval(_ *PromptEvalContext) ([]string, error) {
	// TODO
	return []string{}, nil
}

type RootNode struct {
	SysNodes   []Node
	UserNodes  []Node
	ModuleDefs map[string][]Node // 初始化 map
	InFields   []FieldDef
	OutFields  []FieldDef
	// 其它部分
}

func (r *RootNode) Eval(ctx *PromptEvalContext) ([]string, error) {
	var result []string
	for _, node := range r.SysNodes {
		out, err := node.Eval(ctx)
		if err != nil {
			return nil, err
		}
		result = append(result, out...)
	}
	for _, node := range r.UserNodes {
		out, err := node.Eval(ctx)
		if err != nil {
			return nil, err
		}
		result = append(result, out...)
	}
	// 你可以加更多逻辑
	return result, nil
}

type ParsedPrompt struct {
	InDef  []FieldDef
	OutDef []FieldDef

	// Sys []Node
	// User []Node
	// Or:
	SysVars  map[string][]Node
	UserVars map[string][]Node

	BeforeCode string
	FixCode    string
	AfterCode  string
}

func (parsed *ParsedPrompt) GetSystemPrompt(inp any) (string, error) {
	return "", nil
}

func (parsed *ParsedPrompt) GetUserPrompt(inp any) (string, error) {
	return "", nil
}

// Or
func (parsed *ParsedPrompt) GenSystemPromptFn() (string, error) {
	return "", nil
}

func (parsed *ParsedPrompt) GenUserPromptFn() (string, error) {
	return "", nil
}
