// Code generated from PromptDSLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromptDSLParser
import "github.com/antlr4-go/antlr/v4"

type BasePromptDSLParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePromptDSLParserVisitor) VisitPromptFile(ctx *PromptFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitPromptDef(ctx *PromptDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitPromptBlock(ctx *PromptBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitSysSection(ctx *SysSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitModuleDef(ctx *ModuleDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitModuleContent(ctx *ModuleContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitInputSection(ctx *InputSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitOutputSection(ctx *OutputSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitOutputStruct(ctx *OutputStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitOutputMarkdown(ctx *OutputMarkdownContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitSystemSection(ctx *SystemSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitUserSection(ctx *UserSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitUserContent(ctx *UserContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitNoteSection(ctx *NoteSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitAfterSection(ctx *AfterSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitFixSection(ctx *FixSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitDslCallExpr(ctx *DslCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitFieldDef(ctx *FieldDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitTextLine(ctx *TextLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitParamPath(ctx *ParamPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitStructDef(ctx *StructDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitAnnotationArgs(ctx *AnnotationArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitAnnotationValue(ctx *AnnotationValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitTextBlock(ctx *TextBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLParserVisitor) VisitFormatType(ctx *FormatTypeContext) interface{} {
	return v.VisitChildren(ctx)
}
