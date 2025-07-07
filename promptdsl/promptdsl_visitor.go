// Code generated from PromptDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promptdsl // PromptDSL
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PromptDSLParser.
type PromptDSLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PromptDSLParser#promptFile.
	VisitPromptFile(ctx *PromptFileContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#promptDef.
	VisitPromptDef(ctx *PromptDefContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#promptBlock.
	VisitPromptBlock(ctx *PromptBlockContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#paramBody.
	VisitParamBody(ctx *ParamBodyContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#inputSection.
	VisitInputSection(ctx *InputSectionContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#outputSection.
	VisitOutputSection(ctx *OutputSectionContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#singleField.
	VisitSingleField(ctx *SingleFieldContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#outputBody.
	VisitOutputBody(ctx *OutputBodyContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#fieldDef.
	VisitFieldDef(ctx *FieldDefContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#structDef.
	VisitStructDef(ctx *StructDefContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#annotationArgs.
	VisitAnnotationArgs(ctx *AnnotationArgsContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#annotationValue.
	VisitAnnotationValue(ctx *AnnotationValueContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#userBlock.
	VisitUserBlock(ctx *UserBlockContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#classificationBlock.
	VisitClassificationBlock(ctx *ClassificationBlockContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#summarizationBlock.
	VisitSummarizationBlock(ctx *SummarizationBlockContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#compilationBlock.
	VisitCompilationBlock(ctx *CompilationBlockContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#kvPair.
	VisitKvPair(ctx *KvPairContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#textBlock.
	VisitTextBlock(ctx *TextBlockContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#formatType.
	VisitFormatType(ctx *FormatTypeContext) interface{}
}
