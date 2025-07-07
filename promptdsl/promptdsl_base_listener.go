// Code generated from PromptDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promptdsl // PromptDSL
import "github.com/antlr4-go/antlr/v4"

// BasePromptDSLListener is a complete listener for a parse tree produced by PromptDSLParser.
type BasePromptDSLListener struct{}

var _ PromptDSLListener = &BasePromptDSLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePromptDSLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePromptDSLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePromptDSLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePromptDSLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPromptFile is called when production promptFile is entered.
func (s *BasePromptDSLListener) EnterPromptFile(ctx *PromptFileContext) {}

// ExitPromptFile is called when production promptFile is exited.
func (s *BasePromptDSLListener) ExitPromptFile(ctx *PromptFileContext) {}

// EnterPromptDef is called when production promptDef is entered.
func (s *BasePromptDSLListener) EnterPromptDef(ctx *PromptDefContext) {}

// ExitPromptDef is called when production promptDef is exited.
func (s *BasePromptDSLListener) ExitPromptDef(ctx *PromptDefContext) {}

// EnterPromptBlock is called when production promptBlock is entered.
func (s *BasePromptDSLListener) EnterPromptBlock(ctx *PromptBlockContext) {}

// ExitPromptBlock is called when production promptBlock is exited.
func (s *BasePromptDSLListener) ExitPromptBlock(ctx *PromptBlockContext) {}

// EnterParamBody is called when production paramBody is entered.
func (s *BasePromptDSLListener) EnterParamBody(ctx *ParamBodyContext) {}

// ExitParamBody is called when production paramBody is exited.
func (s *BasePromptDSLListener) ExitParamBody(ctx *ParamBodyContext) {}

// EnterInputSection is called when production inputSection is entered.
func (s *BasePromptDSLListener) EnterInputSection(ctx *InputSectionContext) {}

// ExitInputSection is called when production inputSection is exited.
func (s *BasePromptDSLListener) ExitInputSection(ctx *InputSectionContext) {}

// EnterOutputSection is called when production outputSection is entered.
func (s *BasePromptDSLListener) EnterOutputSection(ctx *OutputSectionContext) {}

// ExitOutputSection is called when production outputSection is exited.
func (s *BasePromptDSLListener) ExitOutputSection(ctx *OutputSectionContext) {}

// EnterSingleField is called when production singleField is entered.
func (s *BasePromptDSLListener) EnterSingleField(ctx *SingleFieldContext) {}

// ExitSingleField is called when production singleField is exited.
func (s *BasePromptDSLListener) ExitSingleField(ctx *SingleFieldContext) {}

// EnterOutputBody is called when production outputBody is entered.
func (s *BasePromptDSLListener) EnterOutputBody(ctx *OutputBodyContext) {}

// ExitOutputBody is called when production outputBody is exited.
func (s *BasePromptDSLListener) ExitOutputBody(ctx *OutputBodyContext) {}

// EnterFieldDef is called when production fieldDef is entered.
func (s *BasePromptDSLListener) EnterFieldDef(ctx *FieldDefContext) {}

// ExitFieldDef is called when production fieldDef is exited.
func (s *BasePromptDSLListener) ExitFieldDef(ctx *FieldDefContext) {}

// EnterStructDef is called when production structDef is entered.
func (s *BasePromptDSLListener) EnterStructDef(ctx *StructDefContext) {}

// ExitStructDef is called when production structDef is exited.
func (s *BasePromptDSLListener) ExitStructDef(ctx *StructDefContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BasePromptDSLListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BasePromptDSLListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterAnnotationArgs is called when production annotationArgs is entered.
func (s *BasePromptDSLListener) EnterAnnotationArgs(ctx *AnnotationArgsContext) {}

// ExitAnnotationArgs is called when production annotationArgs is exited.
func (s *BasePromptDSLListener) ExitAnnotationArgs(ctx *AnnotationArgsContext) {}

// EnterAnnotationValue is called when production annotationValue is entered.
func (s *BasePromptDSLListener) EnterAnnotationValue(ctx *AnnotationValueContext) {}

// ExitAnnotationValue is called when production annotationValue is exited.
func (s *BasePromptDSLListener) ExitAnnotationValue(ctx *AnnotationValueContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BasePromptDSLListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BasePromptDSLListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterUserBlock is called when production userBlock is entered.
func (s *BasePromptDSLListener) EnterUserBlock(ctx *UserBlockContext) {}

// ExitUserBlock is called when production userBlock is exited.
func (s *BasePromptDSLListener) ExitUserBlock(ctx *UserBlockContext) {}

// EnterClassificationBlock is called when production classificationBlock is entered.
func (s *BasePromptDSLListener) EnterClassificationBlock(ctx *ClassificationBlockContext) {}

// ExitClassificationBlock is called when production classificationBlock is exited.
func (s *BasePromptDSLListener) ExitClassificationBlock(ctx *ClassificationBlockContext) {}

// EnterSummarizationBlock is called when production summarizationBlock is entered.
func (s *BasePromptDSLListener) EnterSummarizationBlock(ctx *SummarizationBlockContext) {}

// ExitSummarizationBlock is called when production summarizationBlock is exited.
func (s *BasePromptDSLListener) ExitSummarizationBlock(ctx *SummarizationBlockContext) {}

// EnterCompilationBlock is called when production compilationBlock is entered.
func (s *BasePromptDSLListener) EnterCompilationBlock(ctx *CompilationBlockContext) {}

// ExitCompilationBlock is called when production compilationBlock is exited.
func (s *BasePromptDSLListener) ExitCompilationBlock(ctx *CompilationBlockContext) {}

// EnterKvPair is called when production kvPair is entered.
func (s *BasePromptDSLListener) EnterKvPair(ctx *KvPairContext) {}

// ExitKvPair is called when production kvPair is exited.
func (s *BasePromptDSLListener) ExitKvPair(ctx *KvPairContext) {}

// EnterTextBlock is called when production textBlock is entered.
func (s *BasePromptDSLListener) EnterTextBlock(ctx *TextBlockContext) {}

// ExitTextBlock is called when production textBlock is exited.
func (s *BasePromptDSLListener) ExitTextBlock(ctx *TextBlockContext) {}

// EnterType is called when production type is entered.
func (s *BasePromptDSLListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BasePromptDSLListener) ExitType(ctx *TypeContext) {}

// EnterValue is called when production value is entered.
func (s *BasePromptDSLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasePromptDSLListener) ExitValue(ctx *ValueContext) {}

// EnterFormatType is called when production formatType is entered.
func (s *BasePromptDSLListener) EnterFormatType(ctx *FormatTypeContext) {}

// ExitFormatType is called when production formatType is exited.
func (s *BasePromptDSLListener) ExitFormatType(ctx *FormatTypeContext) {}
