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

	// EnterSysSection is called when entering the sysSection production.
	EnterSysSection(c *SysSectionContext)

	// EnterModuleDef is called when entering the moduleDef production.
	EnterModuleDef(c *ModuleDefContext)

	// EnterModuleContent is called when entering the moduleContent production.
	EnterModuleContent(c *ModuleContentContext)

	// EnterInputSection is called when entering the inputSection production.
	EnterInputSection(c *InputSectionContext)

	// EnterOutputSection is called when entering the outputSection production.
	EnterOutputSection(c *OutputSectionContext)

	// EnterOutputStruct is called when entering the outputStruct production.
	EnterOutputStruct(c *OutputStructContext)

	// EnterOutputMarkdown is called when entering the outputMarkdown production.
	EnterOutputMarkdown(c *OutputMarkdownContext)

	// EnterSystemSection is called when entering the systemSection production.
	EnterSystemSection(c *SystemSectionContext)

	// EnterUserSection is called when entering the userSection production.
	EnterUserSection(c *UserSectionContext)

	// EnterUserContent is called when entering the userContent production.
	EnterUserContent(c *UserContentContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterNoteSection is called when entering the noteSection production.
	EnterNoteSection(c *NoteSectionContext)

	// EnterDslCallExpr is called when entering the dslCallExpr production.
	EnterDslCallExpr(c *DslCallExprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterFieldDef is called when entering the fieldDef production.
	EnterFieldDef(c *FieldDefContext)

	// EnterTextLine is called when entering the textLine production.
	EnterTextLine(c *TextLineContext)

	// EnterParamPath is called when entering the paramPath production.
	EnterParamPath(c *ParamPathContext)

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

	// EnterAfterSection is called when entering the afterSection production.
	EnterAfterSection(c *AfterSectionContext)

	// EnterFixSection is called when entering the fixSection production.
	EnterFixSection(c *FixSectionContext)

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

	// ExitSysSection is called when exiting the sysSection production.
	ExitSysSection(c *SysSectionContext)

	// ExitModuleDef is called when exiting the moduleDef production.
	ExitModuleDef(c *ModuleDefContext)

	// ExitModuleContent is called when exiting the moduleContent production.
	ExitModuleContent(c *ModuleContentContext)

	// ExitInputSection is called when exiting the inputSection production.
	ExitInputSection(c *InputSectionContext)

	// ExitOutputSection is called when exiting the outputSection production.
	ExitOutputSection(c *OutputSectionContext)

	// ExitOutputStruct is called when exiting the outputStruct production.
	ExitOutputStruct(c *OutputStructContext)

	// ExitOutputMarkdown is called when exiting the outputMarkdown production.
	ExitOutputMarkdown(c *OutputMarkdownContext)

	// ExitSystemSection is called when exiting the systemSection production.
	ExitSystemSection(c *SystemSectionContext)

	// ExitUserSection is called when exiting the userSection production.
	ExitUserSection(c *UserSectionContext)

	// ExitUserContent is called when exiting the userContent production.
	ExitUserContent(c *UserContentContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitNoteSection is called when exiting the noteSection production.
	ExitNoteSection(c *NoteSectionContext)

	// ExitDslCallExpr is called when exiting the dslCallExpr production.
	ExitDslCallExpr(c *DslCallExprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitFieldDef is called when exiting the fieldDef production.
	ExitFieldDef(c *FieldDefContext)

	// ExitTextLine is called when exiting the textLine production.
	ExitTextLine(c *TextLineContext)

	// ExitParamPath is called when exiting the paramPath production.
	ExitParamPath(c *ParamPathContext)

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

	// ExitAfterSection is called when exiting the afterSection production.
	ExitAfterSection(c *AfterSectionContext)

	// ExitFixSection is called when exiting the fixSection production.
	ExitFixSection(c *FixSectionContext)

	// ExitTextBlock is called when exiting the textBlock production.
	ExitTextBlock(c *TextBlockContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitFormatType is called when exiting the formatType production.
	ExitFormatType(c *FormatTypeContext)
}
