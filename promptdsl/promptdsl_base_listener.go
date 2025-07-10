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

// EnterInputSection is called when production inputSection is entered.
func (s *BasePromptDSLListener) EnterInputSection(ctx *InputSectionContext) {}

// ExitInputSection is called when production inputSection is exited.
func (s *BasePromptDSLListener) ExitInputSection(ctx *InputSectionContext) {}

// EnterOutputSection is called when production outputSection is entered.
func (s *BasePromptDSLListener) EnterOutputSection(ctx *OutputSectionContext) {}

// ExitOutputSection is called when production outputSection is exited.
func (s *BasePromptDSLListener) ExitOutputSection(ctx *OutputSectionContext) {}

// EnterOutputEntry is called when production outputEntry is entered.
func (s *BasePromptDSLListener) EnterOutputEntry(ctx *OutputEntryContext) {}

// ExitOutputEntry is called when production outputEntry is exited.
func (s *BasePromptDSLListener) ExitOutputEntry(ctx *OutputEntryContext) {}

// EnterSystemSection is called when production systemSection is entered.
func (s *BasePromptDSLListener) EnterSystemSection(ctx *SystemSectionContext) {}

// ExitSystemSection is called when production systemSection is exited.
func (s *BasePromptDSLListener) ExitSystemSection(ctx *SystemSectionContext) {}

// EnterUserSection is called when production userSection is entered.
func (s *BasePromptDSLListener) EnterUserSection(ctx *UserSectionContext) {}

// ExitUserSection is called when production userSection is exited.
func (s *BasePromptDSLListener) ExitUserSection(ctx *UserSectionContext) {}

// EnterNoteSection is called when production noteSection is entered.
func (s *BasePromptDSLListener) EnterNoteSection(ctx *NoteSectionContext) {}

// ExitNoteSection is called when production noteSection is exited.
func (s *BasePromptDSLListener) ExitNoteSection(ctx *NoteSectionContext) {}

// EnterBeforeSection is called when production beforeSection is entered.
func (s *BasePromptDSLListener) EnterBeforeSection(ctx *BeforeSectionContext) {}

// ExitBeforeSection is called when production beforeSection is exited.
func (s *BasePromptDSLListener) ExitBeforeSection(ctx *BeforeSectionContext) {}

// EnterBeforeContent is called when production beforeContent is entered.
func (s *BasePromptDSLListener) EnterBeforeContent(ctx *BeforeContentContext) {}

// ExitBeforeContent is called when production beforeContent is exited.
func (s *BasePromptDSLListener) ExitBeforeContent(ctx *BeforeContentContext) {}

// EnterBeforeEntry is called when production beforeEntry is entered.
func (s *BasePromptDSLListener) EnterBeforeEntry(ctx *BeforeEntryContext) {}

// ExitBeforeEntry is called when production beforeEntry is exited.
func (s *BasePromptDSLListener) ExitBeforeEntry(ctx *BeforeEntryContext) {}

// EnterAfterSection is called when production afterSection is entered.
func (s *BasePromptDSLListener) EnterAfterSection(ctx *AfterSectionContext) {}

// ExitAfterSection is called when production afterSection is exited.
func (s *BasePromptDSLListener) ExitAfterSection(ctx *AfterSectionContext) {}

// EnterAfterContent is called when production afterContent is entered.
func (s *BasePromptDSLListener) EnterAfterContent(ctx *AfterContentContext) {}

// ExitAfterContent is called when production afterContent is exited.
func (s *BasePromptDSLListener) ExitAfterContent(ctx *AfterContentContext) {}

// EnterAfterEntry is called when production afterEntry is entered.
func (s *BasePromptDSLListener) EnterAfterEntry(ctx *AfterEntryContext) {}

// ExitAfterEntry is called when production afterEntry is exited.
func (s *BasePromptDSLListener) ExitAfterEntry(ctx *AfterEntryContext) {}

// EnterDslCallExpr is called when production dslCallExpr is entered.
func (s *BasePromptDSLListener) EnterDslCallExpr(ctx *DslCallExprContext) {}

// ExitDslCallExpr is called when production dslCallExpr is exited.
func (s *BasePromptDSLListener) ExitDslCallExpr(ctx *DslCallExprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePromptDSLListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePromptDSLListener) ExitExpr(ctx *ExprContext) {}

// EnterFieldDef is called when production fieldDef is entered.
func (s *BasePromptDSLListener) EnterFieldDef(ctx *FieldDefContext) {}

// ExitFieldDef is called when production fieldDef is exited.
func (s *BasePromptDSLListener) ExitFieldDef(ctx *FieldDefContext) {}

// EnterTextLine is called when production textLine is entered.
func (s *BasePromptDSLListener) EnterTextLine(ctx *TextLineContext) {}

// ExitTextLine is called when production textLine is exited.
func (s *BasePromptDSLListener) ExitTextLine(ctx *TextLineContext) {}

// EnterParamPath is called when production paramPath is entered.
func (s *BasePromptDSLListener) EnterParamPath(ctx *ParamPathContext) {}

// ExitParamPath is called when production paramPath is exited.
func (s *BasePromptDSLListener) ExitParamPath(ctx *ParamPathContext) {}

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
