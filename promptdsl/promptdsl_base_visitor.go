// Code generated from PromptDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promptdsl // PromptDSL
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

func (v *BasePromptDSLVisitor) VisitInputSection(ctx *InputSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitOutputSection(ctx *OutputSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitOutputEntry(ctx *OutputEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitSystemSection(ctx *SystemSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitUserSection(ctx *UserSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitNoteSection(ctx *NoteSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitBeforeSection(ctx *BeforeSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitBeforeContent(ctx *BeforeContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitBeforeEntry(ctx *BeforeEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitAfterSection(ctx *AfterSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitAfterContent(ctx *AfterContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitAfterEntry(ctx *AfterEntryContext) interface{} {
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
