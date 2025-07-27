// Generated from d:/work/promptDSL/promptdsl-core/PromptDSL.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link PromptDSLParser}.
 */
public interface PromptDSLListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#promptFile}.
	 * @param ctx the parse tree
	 */
	void enterPromptFile(PromptDSLParser.PromptFileContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#promptFile}.
	 * @param ctx the parse tree
	 */
	void exitPromptFile(PromptDSLParser.PromptFileContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#promptDef}.
	 * @param ctx the parse tree
	 */
	void enterPromptDef(PromptDSLParser.PromptDefContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#promptDef}.
	 * @param ctx the parse tree
	 */
	void exitPromptDef(PromptDSLParser.PromptDefContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#promptBlock}.
	 * @param ctx the parse tree
	 */
	void enterPromptBlock(PromptDSLParser.PromptBlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#promptBlock}.
	 * @param ctx the parse tree
	 */
	void exitPromptBlock(PromptDSLParser.PromptBlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#sysSection}.
	 * @param ctx the parse tree
	 */
	void enterSysSection(PromptDSLParser.SysSectionContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#sysSection}.
	 * @param ctx the parse tree
	 */
	void exitSysSection(PromptDSLParser.SysSectionContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#moduleDef}.
	 * @param ctx the parse tree
	 */
	void enterModuleDef(PromptDSLParser.ModuleDefContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#moduleDef}.
	 * @param ctx the parse tree
	 */
	void exitModuleDef(PromptDSLParser.ModuleDefContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#moduleContent}.
	 * @param ctx the parse tree
	 */
	void enterModuleContent(PromptDSLParser.ModuleContentContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#moduleContent}.
	 * @param ctx the parse tree
	 */
	void exitModuleContent(PromptDSLParser.ModuleContentContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#inputSection}.
	 * @param ctx the parse tree
	 */
	void enterInputSection(PromptDSLParser.InputSectionContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#inputSection}.
	 * @param ctx the parse tree
	 */
	void exitInputSection(PromptDSLParser.InputSectionContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#outputSection}.
	 * @param ctx the parse tree
	 */
	void enterOutputSection(PromptDSLParser.OutputSectionContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#outputSection}.
	 * @param ctx the parse tree
	 */
	void exitOutputSection(PromptDSLParser.OutputSectionContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#outputStruct}.
	 * @param ctx the parse tree
	 */
	void enterOutputStruct(PromptDSLParser.OutputStructContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#outputStruct}.
	 * @param ctx the parse tree
	 */
	void exitOutputStruct(PromptDSLParser.OutputStructContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#outputMarkdown}.
	 * @param ctx the parse tree
	 */
	void enterOutputMarkdown(PromptDSLParser.OutputMarkdownContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#outputMarkdown}.
	 * @param ctx the parse tree
	 */
	void exitOutputMarkdown(PromptDSLParser.OutputMarkdownContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#systemSection}.
	 * @param ctx the parse tree
	 */
	void enterSystemSection(PromptDSLParser.SystemSectionContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#systemSection}.
	 * @param ctx the parse tree
	 */
	void exitSystemSection(PromptDSLParser.SystemSectionContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#userSection}.
	 * @param ctx the parse tree
	 */
	void enterUserSection(PromptDSLParser.UserSectionContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#userSection}.
	 * @param ctx the parse tree
	 */
	void exitUserSection(PromptDSLParser.UserSectionContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#userContent}.
	 * @param ctx the parse tree
	 */
	void enterUserContent(PromptDSLParser.UserContentContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#userContent}.
	 * @param ctx the parse tree
	 */
	void exitUserContent(PromptDSLParser.UserContentContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatement(PromptDSLParser.IfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatement(PromptDSLParser.IfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#condition}.
	 * @param ctx the parse tree
	 */
	void enterCondition(PromptDSLParser.ConditionContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#condition}.
	 * @param ctx the parse tree
	 */
	void exitCondition(PromptDSLParser.ConditionContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#noteSection}.
	 * @param ctx the parse tree
	 */
	void enterNoteSection(PromptDSLParser.NoteSectionContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#noteSection}.
	 * @param ctx the parse tree
	 */
	void exitNoteSection(PromptDSLParser.NoteSectionContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#dslCallExpr}.
	 * @param ctx the parse tree
	 */
	void enterDslCallExpr(PromptDSLParser.DslCallExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#dslCallExpr}.
	 * @param ctx the parse tree
	 */
	void exitDslCallExpr(PromptDSLParser.DslCallExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(PromptDSLParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(PromptDSLParser.ExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#fieldDef}.
	 * @param ctx the parse tree
	 */
	void enterFieldDef(PromptDSLParser.FieldDefContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#fieldDef}.
	 * @param ctx the parse tree
	 */
	void exitFieldDef(PromptDSLParser.FieldDefContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#textLine}.
	 * @param ctx the parse tree
	 */
	void enterTextLine(PromptDSLParser.TextLineContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#textLine}.
	 * @param ctx the parse tree
	 */
	void exitTextLine(PromptDSLParser.TextLineContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#paramPath}.
	 * @param ctx the parse tree
	 */
	void enterParamPath(PromptDSLParser.ParamPathContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#paramPath}.
	 * @param ctx the parse tree
	 */
	void exitParamPath(PromptDSLParser.ParamPathContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#structDef}.
	 * @param ctx the parse tree
	 */
	void enterStructDef(PromptDSLParser.StructDefContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#structDef}.
	 * @param ctx the parse tree
	 */
	void exitStructDef(PromptDSLParser.StructDefContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#annotation}.
	 * @param ctx the parse tree
	 */
	void enterAnnotation(PromptDSLParser.AnnotationContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#annotation}.
	 * @param ctx the parse tree
	 */
	void exitAnnotation(PromptDSLParser.AnnotationContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#annotationArgs}.
	 * @param ctx the parse tree
	 */
	void enterAnnotationArgs(PromptDSLParser.AnnotationArgsContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#annotationArgs}.
	 * @param ctx the parse tree
	 */
	void exitAnnotationArgs(PromptDSLParser.AnnotationArgsContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#annotationValue}.
	 * @param ctx the parse tree
	 */
	void enterAnnotationValue(PromptDSLParser.AnnotationValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#annotationValue}.
	 * @param ctx the parse tree
	 */
	void exitAnnotationValue(PromptDSLParser.AnnotationValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#arrayLiteral}.
	 * @param ctx the parse tree
	 */
	void enterArrayLiteral(PromptDSLParser.ArrayLiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#arrayLiteral}.
	 * @param ctx the parse tree
	 */
	void exitArrayLiteral(PromptDSLParser.ArrayLiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#afterSection}.
	 * @param ctx the parse tree
	 */
	void enterAfterSection(PromptDSLParser.AfterSectionContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#afterSection}.
	 * @param ctx the parse tree
	 */
	void exitAfterSection(PromptDSLParser.AfterSectionContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#fixSection}.
	 * @param ctx the parse tree
	 */
	void enterFixSection(PromptDSLParser.FixSectionContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#fixSection}.
	 * @param ctx the parse tree
	 */
	void exitFixSection(PromptDSLParser.FixSectionContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#textBlock}.
	 * @param ctx the parse tree
	 */
	void enterTextBlock(PromptDSLParser.TextBlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#textBlock}.
	 * @param ctx the parse tree
	 */
	void exitTextBlock(PromptDSLParser.TextBlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#type}.
	 * @param ctx the parse tree
	 */
	void enterType(PromptDSLParser.TypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#type}.
	 * @param ctx the parse tree
	 */
	void exitType(PromptDSLParser.TypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#value}.
	 * @param ctx the parse tree
	 */
	void enterValue(PromptDSLParser.ValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#value}.
	 * @param ctx the parse tree
	 */
	void exitValue(PromptDSLParser.ValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link PromptDSLParser#formatType}.
	 * @param ctx the parse tree
	 */
	void enterFormatType(PromptDSLParser.FormatTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link PromptDSLParser#formatType}.
	 * @param ctx the parse tree
	 */
	void exitFormatType(PromptDSLParser.FormatTypeContext ctx);
}