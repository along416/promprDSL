package promptdslcore

type FieldDef struct {
	Name        string
	Type        string
	JsonName    string
	Hint        string
	Annotations []string
}

type PromptEvalContext struct {
	InFields  []FieldDef
	OutFields []FieldDef
	Input     any
}

type Node interface {
	Eval(_ *PromptEvalContext) ([]string, error)
	ConvertToCode() string
}

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

func (e *Expr) EvalToBool(_ *PromptEvalContext) (bool, error) {
	return false, nil
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
		for _, n := range node.Then {
			texts, err := n.Eval(ctx)
			if err != nil {
				return nil, err
			}
			result = append(result, texts...)
		}
	} else {
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
	SysNodes  []Node
	UserNodes []Node
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
