// Code generated from PromptDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromptDSL
import "github.com/antlr4-go/antlr/v4"

type BasePromptDSLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePromptDSLVisitor) VisitPromptFile(ctx *PromptFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitPromptDef(ctx *PromptDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitPromptBlock(ctx *PromptBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitSysSection(ctx *SysSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitModuleDef(ctx *ModuleDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitModuleContent(ctx *ModuleContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitInputSection(ctx *InputSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitOutputSection(ctx *OutputSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitOutputStruct(ctx *OutputStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitOutputMarkdown(ctx *OutputMarkdownContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitSystemSection(ctx *SystemSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitUserSection(ctx *UserSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitUserContent(ctx *UserContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitNoteSection(ctx *NoteSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitDslCallExpr(ctx *DslCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitFieldDef(ctx *FieldDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitTextLine(ctx *TextLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitParamPath(ctx *ParamPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitStructDef(ctx *StructDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitAnnotationArgs(ctx *AnnotationArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitAnnotationValue(ctx *AnnotationValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitAfterSection(ctx *AfterSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitFixSection(ctx *FixSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitTextBlock(ctx *TextBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitFormatType(ctx *FormatTypeContext) interface{} {
	return v.VisitChildren(ctx)
}
