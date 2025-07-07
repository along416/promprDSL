// Code generated from PromptDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promptdsl // PromptDSL
import "github.com/antlr4-go/antlr/v4"

// PromptDSLListener is a complete listener for a parse tree produced by PromptDSLParser.
type PromptDSLListener interface {
	antlr.ParseTreeListener

	// EnterPromptFile is called when entering the promptFile production.
	EnterPromptFile(c *PromptFileContext)

	// EnterPromptDef is called when entering the promptDef production.
	EnterPromptDef(c *PromptDefContext)

	// EnterPromptBlock is called when entering the promptBlock production.
	EnterPromptBlock(c *PromptBlockContext)

	// EnterParamBody is called when entering the paramBody production.
	EnterParamBody(c *ParamBodyContext)

	// EnterInputSection is called when entering the inputSection production.
	EnterInputSection(c *InputSectionContext)

	// EnterOutputSection is called when entering the outputSection production.
	EnterOutputSection(c *OutputSectionContext)

	// EnterSingleField is called when entering the singleField production.
	EnterSingleField(c *SingleFieldContext)

	// EnterOutputBody is called when entering the outputBody production.
	EnterOutputBody(c *OutputBodyContext)

	// EnterFieldDef is called when entering the fieldDef production.
	EnterFieldDef(c *FieldDefContext)

	// EnterStructDef is called when entering the structDef production.
	EnterStructDef(c *StructDefContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterAnnotationArgs is called when entering the annotationArgs production.
	EnterAnnotationArgs(c *AnnotationArgsContext)

	// EnterAnnotationValue is called when entering the annotationValue production.
	EnterAnnotationValue(c *AnnotationValueContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterUserBlock is called when entering the userBlock production.
	EnterUserBlock(c *UserBlockContext)

	// EnterClassificationBlock is called when entering the classificationBlock production.
	EnterClassificationBlock(c *ClassificationBlockContext)

	// EnterSummarizationBlock is called when entering the summarizationBlock production.
	EnterSummarizationBlock(c *SummarizationBlockContext)

	// EnterCompilationBlock is called when entering the compilationBlock production.
	EnterCompilationBlock(c *CompilationBlockContext)

	// EnterKvPair is called when entering the kvPair production.
	EnterKvPair(c *KvPairContext)

	// EnterTextBlock is called when entering the textBlock production.
	EnterTextBlock(c *TextBlockContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterFormatType is called when entering the formatType production.
	EnterFormatType(c *FormatTypeContext)

	// ExitPromptFile is called when exiting the promptFile production.
	ExitPromptFile(c *PromptFileContext)

	// ExitPromptDef is called when exiting the promptDef production.
	ExitPromptDef(c *PromptDefContext)

	// ExitPromptBlock is called when exiting the promptBlock production.
	ExitPromptBlock(c *PromptBlockContext)

	// ExitParamBody is called when exiting the paramBody production.
	ExitParamBody(c *ParamBodyContext)

	// ExitInputSection is called when exiting the inputSection production.
	ExitInputSection(c *InputSectionContext)

	// ExitOutputSection is called when exiting the outputSection production.
	ExitOutputSection(c *OutputSectionContext)

	// ExitSingleField is called when exiting the singleField production.
	ExitSingleField(c *SingleFieldContext)

	// ExitOutputBody is called when exiting the outputBody production.
	ExitOutputBody(c *OutputBodyContext)

	// ExitFieldDef is called when exiting the fieldDef production.
	ExitFieldDef(c *FieldDefContext)

	// ExitStructDef is called when exiting the structDef production.
	ExitStructDef(c *StructDefContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitAnnotationArgs is called when exiting the annotationArgs production.
	ExitAnnotationArgs(c *AnnotationArgsContext)

	// ExitAnnotationValue is called when exiting the annotationValue production.
	ExitAnnotationValue(c *AnnotationValueContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitUserBlock is called when exiting the userBlock production.
	ExitUserBlock(c *UserBlockContext)

	// ExitClassificationBlock is called when exiting the classificationBlock production.
	ExitClassificationBlock(c *ClassificationBlockContext)

	// ExitSummarizationBlock is called when exiting the summarizationBlock production.
	ExitSummarizationBlock(c *SummarizationBlockContext)

	// ExitCompilationBlock is called when exiting the compilationBlock production.
	ExitCompilationBlock(c *CompilationBlockContext)

	// ExitKvPair is called when exiting the kvPair production.
	ExitKvPair(c *KvPairContext)

	// ExitTextBlock is called when exiting the textBlock production.
	ExitTextBlock(c *TextBlockContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitFormatType is called when exiting the formatType production.
	ExitFormatType(c *FormatTypeContext)
}
