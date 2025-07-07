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

func (v *BasePromptDSLVisitor) VisitParamBody(ctx *ParamBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitInputSection(ctx *InputSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitOutputSection(ctx *OutputSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitSingleField(ctx *SingleFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitOutputBody(ctx *OutputBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitFieldDef(ctx *FieldDefContext) interface{} {
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

func (v *BasePromptDSLVisitor) VisitUserBlock(ctx *UserBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitClassificationBlock(ctx *ClassificationBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitSummarizationBlock(ctx *SummarizationBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitCompilationBlock(ctx *CompilationBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromptDSLVisitor) VisitKvPair(ctx *KvPairContext) interface{} {
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
