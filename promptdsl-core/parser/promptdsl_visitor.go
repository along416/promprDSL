// Code generated from PromptDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromptDSL
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

	// Visit a parse tree produced by PromptDSLParser#inputSection.
	VisitInputSection(ctx *InputSectionContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#outputSection.
	VisitOutputSection(ctx *OutputSectionContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#outputStruct.
	VisitOutputStruct(ctx *OutputStructContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#outputMarkdown.
	VisitOutputMarkdown(ctx *OutputMarkdownContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#beforeSection.
	VisitBeforeSection(ctx *BeforeSectionContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#beforeContent.
	VisitBeforeContent(ctx *BeforeContentContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#varDef.
	VisitVarDef(ctx *VarDefContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#systemSection.
	VisitSystemSection(ctx *SystemSectionContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#sysContent.
	VisitSysContent(ctx *SysContentContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#moduleDef.
	VisitModuleDef(ctx *ModuleDefContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#moduleContent.
	VisitModuleContent(ctx *ModuleContentContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#userSection.
	VisitUserSection(ctx *UserSectionContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#userContent.
	VisitUserContent(ctx *UserContentContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#thencontent.
	VisitThencontent(ctx *ThencontentContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#elsecontent.
	VisitElsecontent(ctx *ElsecontentContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#noteSection.
	VisitNoteSection(ctx *NoteSectionContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#dslCallExpr.
	VisitDslCallExpr(ctx *DslCallExprContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#fieldDef.
	VisitFieldDef(ctx *FieldDefContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#textLine.
	VisitTextLine(ctx *TextLineContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#paramPath.
	VisitParamPath(ctx *ParamPathContext) interface{}

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

	// Visit a parse tree produced by PromptDSLParser#defaultAnnotation.
	VisitDefaultAnnotation(ctx *DefaultAnnotationContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#afterSection.
	VisitAfterSection(ctx *AfterSectionContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#fixSection.
	VisitFixSection(ctx *FixSectionContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#textBlock.
	VisitTextBlock(ctx *TextBlockContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by PromptDSLParser#formatType.
	VisitFormatType(ctx *FormatTypeContext) interface{}
}
