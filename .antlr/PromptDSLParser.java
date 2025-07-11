// Generated from d:/work/promptDSL/PromptDSL.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class PromptDSLParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, PLUS=23, STRING_TYPE=24, 
		FLOAT_TYPE=25, INT_TYPE=26, PROMPT=27, PARAMS=28, SYSTEM=29, USER=30, 
		NOTE=31, INPUT=32, OUTPUT=33, FORMAT=34, TYPE=35, STRUCT=36, BEFORE=37, 
		SCHEMA=38, AFTER=39, PARSE=40, JSONFIX=41, MARKDOWN=42, IF=43, ELSE=44, 
		OUTPUTSPEC=45, ID=46, STRING=47, NUMBER=48, BOOL=49, PIPE=50, SEMI=51, 
		WS=52, LINE_COMMENT=53, BLOCK_COMMENT=54;
	public static final int
		RULE_promptFile = 0, RULE_promptDef = 1, RULE_promptBlock = 2, RULE_sysSection = 3, 
		RULE_moduleDef = 4, RULE_moduleContent = 5, RULE_inputSection = 6, RULE_outputSection = 7, 
		RULE_outputStruct = 8, RULE_outputMarkdown = 9, RULE_systemSection = 10, 
		RULE_userSection = 11, RULE_userContent = 12, RULE_ifStatement = 13, RULE_condition = 14, 
		RULE_noteSection = 15, RULE_dslCallExpr = 16, RULE_expr = 17, RULE_fieldDef = 18, 
		RULE_textLine = 19, RULE_paramPath = 20, RULE_structDef = 21, RULE_annotation = 22, 
		RULE_annotationArgs = 23, RULE_annotationValue = 24, RULE_arrayLiteral = 25, 
		RULE_afterSection = 26, RULE_fixSection = 27, RULE_textBlock = 28, RULE_type = 29, 
		RULE_value = 30, RULE_formatType = 31;
	private static String[] makeRuleNames() {
		return new String[] {
			"promptFile", "promptDef", "promptBlock", "sysSection", "moduleDef", 
			"moduleContent", "inputSection", "outputSection", "outputStruct", "outputMarkdown", 
			"systemSection", "userSection", "userContent", "ifStatement", "condition", 
			"noteSection", "dslCallExpr", "expr", "fieldDef", "textLine", "paramPath", 
			"structDef", "annotation", "annotationArgs", "annotationValue", "arrayLiteral", 
			"afterSection", "fixSection", "textBlock", "type", "value", "formatType"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'{'", "'}'", "'sys'", "':'", "'('", "')'", "'=='", "'!='", "','", 
			"'='", "'.'", "'@'", "'['", "']'", "'<after>'", "'</after>'", "'<fix>'", 
			"'</fix>'", "'-'", "'[]'", "'md'", "'json'", "'+'", "'string'", "'float'", 
			"'int'", "'prompt'", "'params'", "'system'", "'user'", "'note'", "'input'", 
			"'output'", "'format'", "'type'", "'struct'", "'before'", "'schema'", 
			"'after'", "'parse'", "'jsonfix'", "'markdown'", "'if'", "'else'", "'outputspec'", 
			null, null, null, null, "'|'", "';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, "PLUS", 
			"STRING_TYPE", "FLOAT_TYPE", "INT_TYPE", "PROMPT", "PARAMS", "SYSTEM", 
			"USER", "NOTE", "INPUT", "OUTPUT", "FORMAT", "TYPE", "STRUCT", "BEFORE", 
			"SCHEMA", "AFTER", "PARSE", "JSONFIX", "MARKDOWN", "IF", "ELSE", "OUTPUTSPEC", 
			"ID", "STRING", "NUMBER", "BOOL", "PIPE", "SEMI", "WS", "LINE_COMMENT", 
			"BLOCK_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "PromptDSL.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public PromptDSLParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PromptFileContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(PromptDSLParser.EOF, 0); }
		public List<PromptDefContext> promptDef() {
			return getRuleContexts(PromptDefContext.class);
		}
		public PromptDefContext promptDef(int i) {
			return getRuleContext(PromptDefContext.class,i);
		}
		public PromptFileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_promptFile; }
	}

	public final PromptFileContext promptFile() throws RecognitionException {
		PromptFileContext _localctx = new PromptFileContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_promptFile);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(65); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(64);
				promptDef();
				}
				}
				setState(67); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==PROMPT );
			setState(69);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PromptDefContext extends ParserRuleContext {
		public TerminalNode PROMPT() { return getToken(PromptDSLParser.PROMPT, 0); }
		public TerminalNode ID() { return getToken(PromptDSLParser.ID, 0); }
		public List<PromptBlockContext> promptBlock() {
			return getRuleContexts(PromptBlockContext.class);
		}
		public PromptBlockContext promptBlock(int i) {
			return getRuleContext(PromptBlockContext.class,i);
		}
		public PromptDefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_promptDef; }
	}

	public final PromptDefContext promptDef() throws RecognitionException {
		PromptDefContext _localctx = new PromptDefContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_promptDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			match(PROMPT);
			setState(72);
			match(ID);
			setState(73);
			match(T__0);
			setState(75); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(74);
				promptBlock();
				}
				}
				setState(77); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 70385387339784L) != 0) );
			setState(79);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PromptBlockContext extends ParserRuleContext {
		public InputSectionContext inputSection() {
			return getRuleContext(InputSectionContext.class,0);
		}
		public OutputSectionContext outputSection() {
			return getRuleContext(OutputSectionContext.class,0);
		}
		public SysSectionContext sysSection() {
			return getRuleContext(SysSectionContext.class,0);
		}
		public SystemSectionContext systemSection() {
			return getRuleContext(SystemSectionContext.class,0);
		}
		public UserSectionContext userSection() {
			return getRuleContext(UserSectionContext.class,0);
		}
		public NoteSectionContext noteSection() {
			return getRuleContext(NoteSectionContext.class,0);
		}
		public AfterSectionContext afterSection() {
			return getRuleContext(AfterSectionContext.class,0);
		}
		public FixSectionContext fixSection() {
			return getRuleContext(FixSectionContext.class,0);
		}
		public ModuleDefContext moduleDef() {
			return getRuleContext(ModuleDefContext.class,0);
		}
		public PromptBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_promptBlock; }
	}

	public final PromptBlockContext promptBlock() throws RecognitionException {
		PromptBlockContext _localctx = new PromptBlockContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_promptBlock);
		try {
			setState(90);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INPUT:
				enterOuterAlt(_localctx, 1);
				{
				setState(81);
				inputSection();
				}
				break;
			case OUTPUT:
				enterOuterAlt(_localctx, 2);
				{
				setState(82);
				outputSection();
				}
				break;
			case T__2:
				enterOuterAlt(_localctx, 3);
				{
				setState(83);
				sysSection();
				}
				break;
			case SYSTEM:
				enterOuterAlt(_localctx, 4);
				{
				setState(84);
				systemSection();
				}
				break;
			case USER:
				enterOuterAlt(_localctx, 5);
				{
				setState(85);
				userSection();
				}
				break;
			case NOTE:
				enterOuterAlt(_localctx, 6);
				{
				setState(86);
				noteSection();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 7);
				{
				setState(87);
				afterSection();
				}
				break;
			case T__16:
				enterOuterAlt(_localctx, 8);
				{
				setState(88);
				fixSection();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 9);
				{
				setState(89);
				moduleDef();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SysSectionContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(PromptDSLParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(PromptDSLParser.ID, i);
		}
		public SysSectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sysSection; }
	}

	public final SysSectionContext sysSection() throws RecognitionException {
		SysSectionContext _localctx = new SysSectionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_sysSection);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(92);
			match(T__2);
			setState(93);
			match(T__0);
			setState(95); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(94);
				match(ID);
				}
				}
				setState(97); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			setState(99);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ModuleDefContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(PromptDSLParser.ID, 0); }
		public List<ModuleContentContext> moduleContent() {
			return getRuleContexts(ModuleContentContext.class);
		}
		public ModuleContentContext moduleContent(int i) {
			return getRuleContext(ModuleContentContext.class,i);
		}
		public ModuleDefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_moduleDef; }
	}

	public final ModuleDefContext moduleDef() throws RecognitionException {
		ModuleDefContext _localctx = new ModuleDefContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_moduleDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(101);
			match(ID);
			setState(102);
			match(T__0);
			setState(106);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 9219005566943232L) != 0)) {
				{
				{
				setState(103);
				moduleContent();
				}
				}
				setState(108);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(109);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ModuleContentContext extends ParserRuleContext {
		public TextLineContext textLine() {
			return getRuleContext(TextLineContext.class,0);
		}
		public ParamPathContext paramPath() {
			return getRuleContext(ParamPathContext.class,0);
		}
		public ModuleContentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_moduleContent; }
	}

	public final ModuleContentContext moduleContent() throws RecognitionException {
		ModuleContentContext _localctx = new ModuleContentContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_moduleContent);
		try {
			setState(113);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(111);
				textLine();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(112);
				paramPath();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InputSectionContext extends ParserRuleContext {
		public TerminalNode INPUT() { return getToken(PromptDSLParser.INPUT, 0); }
		public List<FieldDefContext> fieldDef() {
			return getRuleContexts(FieldDefContext.class);
		}
		public FieldDefContext fieldDef(int i) {
			return getRuleContext(FieldDefContext.class,i);
		}
		public InputSectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inputSection; }
	}

	public final InputSectionContext inputSection() throws RecognitionException {
		InputSectionContext _localctx = new InputSectionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_inputSection);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(115);
			match(INPUT);
			setState(116);
			match(T__0);
			setState(118); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(117);
				fieldDef();
				}
				}
				setState(120); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			setState(122);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OutputSectionContext extends ParserRuleContext {
		public TerminalNode OUTPUT() { return getToken(PromptDSLParser.OUTPUT, 0); }
		public OutputStructContext outputStruct() {
			return getRuleContext(OutputStructContext.class,0);
		}
		public OutputMarkdownContext outputMarkdown() {
			return getRuleContext(OutputMarkdownContext.class,0);
		}
		public OutputSectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_outputSection; }
	}

	public final OutputSectionContext outputSection() throws RecognitionException {
		OutputSectionContext _localctx = new OutputSectionContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_outputSection);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(124);
			match(OUTPUT);
			setState(127);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				{
				setState(125);
				outputStruct();
				}
				break;
			case T__3:
				{
				setState(126);
				outputMarkdown();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OutputStructContext extends ParserRuleContext {
		public List<FieldDefContext> fieldDef() {
			return getRuleContexts(FieldDefContext.class);
		}
		public FieldDefContext fieldDef(int i) {
			return getRuleContext(FieldDefContext.class,i);
		}
		public OutputStructContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_outputStruct; }
	}

	public final OutputStructContext outputStruct() throws RecognitionException {
		OutputStructContext _localctx = new OutputStructContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_outputStruct);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(129);
			match(T__0);
			setState(131); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(130);
				fieldDef();
				}
				}
				setState(133); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			setState(135);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OutputMarkdownContext extends ParserRuleContext {
		public TerminalNode MARKDOWN() { return getToken(PromptDSLParser.MARKDOWN, 0); }
		public OutputMarkdownContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_outputMarkdown; }
	}

	public final OutputMarkdownContext outputMarkdown() throws RecognitionException {
		OutputMarkdownContext _localctx = new OutputMarkdownContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_outputMarkdown);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(137);
			match(T__3);
			setState(138);
			match(MARKDOWN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SystemSectionContext extends ParserRuleContext {
		public TerminalNode SYSTEM() { return getToken(PromptDSLParser.SYSTEM, 0); }
		public List<TextLineContext> textLine() {
			return getRuleContexts(TextLineContext.class);
		}
		public TextLineContext textLine(int i) {
			return getRuleContext(TextLineContext.class,i);
		}
		public SystemSectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_systemSection; }
	}

	public final SystemSectionContext systemSection() throws RecognitionException {
		SystemSectionContext _localctx = new SystemSectionContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_systemSection);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(140);
			match(SYSTEM);
			setState(141);
			match(T__0);
			setState(143); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(142);
				textLine();
				}
				}
				setState(145); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 9219005566943232L) != 0) );
			setState(147);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UserSectionContext extends ParserRuleContext {
		public TerminalNode USER() { return getToken(PromptDSLParser.USER, 0); }
		public List<UserContentContext> userContent() {
			return getRuleContexts(UserContentContext.class);
		}
		public UserContentContext userContent(int i) {
			return getRuleContext(UserContentContext.class,i);
		}
		public UserSectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_userSection; }
	}

	public final UserSectionContext userSection() throws RecognitionException {
		UserSectionContext _localctx = new UserSectionContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_userSection);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(149);
			match(USER);
			setState(150);
			match(T__0);
			setState(152); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(151);
				userContent();
				}
				}
				setState(154); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 10107410962186240L) != 0) );
			setState(156);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UserContentContext extends ParserRuleContext {
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public TextLineContext textLine() {
			return getRuleContext(TextLineContext.class,0);
		}
		public TerminalNode OUTPUTSPEC() { return getToken(PromptDSLParser.OUTPUTSPEC, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public UserContentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_userContent; }
	}

	public final UserContentContext userContent() throws RecognitionException {
		UserContentContext _localctx = new UserContentContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_userContent);
		try {
			setState(162);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(158);
				ifStatement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(159);
				textLine();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(160);
				match(OUTPUTSPEC);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(161);
				expr();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(PromptDSLParser.IF, 0); }
		public ConditionContext condition() {
			return getRuleContext(ConditionContext.class,0);
		}
		public List<UserContentContext> userContent() {
			return getRuleContexts(UserContentContext.class);
		}
		public UserContentContext userContent(int i) {
			return getRuleContext(UserContentContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(PromptDSLParser.ELSE, 0); }
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_ifStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(164);
			match(IF);
			setState(165);
			match(T__4);
			setState(166);
			condition();
			setState(167);
			match(T__5);
			setState(168);
			match(T__0);
			setState(172);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 10107410962186240L) != 0)) {
				{
				{
				setState(169);
				userContent();
				}
				}
				setState(174);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(175);
			match(T__1);
			setState(185);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(176);
				match(ELSE);
				setState(177);
				match(T__0);
				setState(181);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 10107410962186240L) != 0)) {
					{
					{
					setState(178);
					userContent();
					}
					}
					setState(183);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(184);
				match(T__1);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condition; }
	}

	public final ConditionContext condition() throws RecognitionException {
		ConditionContext _localctx = new ConditionContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_condition);
		int _la;
		try {
			setState(192);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(187);
				expr();
				setState(188);
				_la = _input.LA(1);
				if ( !(_la==T__6 || _la==T__7) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(189);
				expr();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(191);
				expr();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class NoteSectionContext extends ParserRuleContext {
		public TerminalNode NOTE() { return getToken(PromptDSLParser.NOTE, 0); }
		public List<TextLineContext> textLine() {
			return getRuleContexts(TextLineContext.class);
		}
		public TextLineContext textLine(int i) {
			return getRuleContext(TextLineContext.class,i);
		}
		public NoteSectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_noteSection; }
	}

	public final NoteSectionContext noteSection() throws RecognitionException {
		NoteSectionContext _localctx = new NoteSectionContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_noteSection);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(194);
			match(NOTE);
			setState(195);
			match(T__0);
			setState(197); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(196);
				textLine();
				}
				}
				setState(199); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 9219005566943232L) != 0) );
			setState(201);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DslCallExprContext extends ParserRuleContext {
		public ParamPathContext paramPath() {
			return getRuleContext(ParamPathContext.class,0);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public DslCallExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dslCallExpr; }
	}

	public final DslCallExprContext dslCallExpr() throws RecognitionException {
		DslCallExprContext _localctx = new DslCallExprContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_dslCallExpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(203);
			paramPath();
			setState(204);
			match(T__4);
			setState(213);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1056231242334208L) != 0)) {
				{
				setState(205);
				expr();
				setState(210);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__8) {
					{
					{
					setState(206);
					match(T__8);
					setState(207);
					expr();
					}
					}
					setState(212);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(215);
			match(T__5);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(PromptDSLParser.STRING, 0); }
		public TerminalNode NUMBER() { return getToken(PromptDSLParser.NUMBER, 0); }
		public TerminalNode BOOL() { return getToken(PromptDSLParser.BOOL, 0); }
		public ParamPathContext paramPath() {
			return getRuleContext(ParamPathContext.class,0);
		}
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	}

	public final ExprContext expr() throws RecognitionException {
		ExprContext _localctx = new ExprContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_expr);
		try {
			setState(221);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(217);
				match(STRING);
				}
				break;
			case NUMBER:
				enterOuterAlt(_localctx, 2);
				{
				setState(218);
				match(NUMBER);
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 3);
				{
				setState(219);
				match(BOOL);
				}
				break;
			case INPUT:
			case OUTPUT:
			case BEFORE:
			case AFTER:
			case ID:
				enterOuterAlt(_localctx, 4);
				{
				setState(220);
				paramPath();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FieldDefContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(PromptDSLParser.ID, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public List<AnnotationContext> annotation() {
			return getRuleContexts(AnnotationContext.class);
		}
		public AnnotationContext annotation(int i) {
			return getRuleContext(AnnotationContext.class,i);
		}
		public TerminalNode SEMI() { return getToken(PromptDSLParser.SEMI, 0); }
		public FieldDefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fieldDef; }
	}

	public final FieldDefContext fieldDef() throws RecognitionException {
		FieldDefContext _localctx = new FieldDefContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_fieldDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(223);
			match(ID);
			setState(224);
			match(T__3);
			setState(225);
			type();
			setState(228);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__9) {
				{
				setState(226);
				match(T__9);
				setState(227);
				value();
				}
			}

			setState(233);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__11) {
				{
				{
				setState(230);
				annotation();
				}
				}
				setState(235);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(237);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(236);
				match(SEMI);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TextLineContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(PromptDSLParser.STRING, 0); }
		public TerminalNode LINE_COMMENT() { return getToken(PromptDSLParser.LINE_COMMENT, 0); }
		public ParamPathContext paramPath() {
			return getRuleContext(ParamPathContext.class,0);
		}
		public TextLineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_textLine; }
	}

	public final TextLineContext textLine() throws RecognitionException {
		TextLineContext _localctx = new TextLineContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_textLine);
		try {
			setState(242);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(239);
				match(STRING);
				}
				break;
			case LINE_COMMENT:
				enterOuterAlt(_localctx, 2);
				{
				setState(240);
				match(LINE_COMMENT);
				}
				break;
			case INPUT:
			case OUTPUT:
			case BEFORE:
			case AFTER:
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(241);
				paramPath();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParamPathContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(PromptDSLParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(PromptDSLParser.ID, i);
		}
		public TerminalNode INPUT() { return getToken(PromptDSLParser.INPUT, 0); }
		public TerminalNode OUTPUT() { return getToken(PromptDSLParser.OUTPUT, 0); }
		public TerminalNode AFTER() { return getToken(PromptDSLParser.AFTER, 0); }
		public TerminalNode BEFORE() { return getToken(PromptDSLParser.BEFORE, 0); }
		public ParamPathContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramPath; }
	}

	public final ParamPathContext paramPath() throws RecognitionException {
		ParamPathContext _localctx = new ParamPathContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_paramPath);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(244);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 71068823846912L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(249);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__10) {
				{
				{
				setState(245);
				match(T__10);
				setState(246);
				match(ID);
				}
				}
				setState(251);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructDefContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(PromptDSLParser.ID, 0); }
		public TerminalNode STRUCT() { return getToken(PromptDSLParser.STRUCT, 0); }
		public List<FieldDefContext> fieldDef() {
			return getRuleContexts(FieldDefContext.class);
		}
		public FieldDefContext fieldDef(int i) {
			return getRuleContext(FieldDefContext.class,i);
		}
		public StructDefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structDef; }
	}

	public final StructDefContext structDef() throws RecognitionException {
		StructDefContext _localctx = new StructDefContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_structDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(252);
			match(ID);
			setState(253);
			match(STRUCT);
			setState(254);
			match(T__0);
			setState(256); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(255);
				fieldDef();
				}
				}
				setState(258); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			setState(260);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AnnotationContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(PromptDSLParser.ID, 0); }
		public AnnotationArgsContext annotationArgs() {
			return getRuleContext(AnnotationArgsContext.class,0);
		}
		public AnnotationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_annotation; }
	}

	public final AnnotationContext annotation() throws RecognitionException {
		AnnotationContext _localctx = new AnnotationContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_annotation);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(262);
			match(T__11);
			setState(263);
			match(ID);
			setState(264);
			match(T__4);
			setState(266);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12 || _la==STRING) {
				{
				setState(265);
				annotationArgs();
				}
			}

			setState(268);
			match(T__5);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AnnotationArgsContext extends ParserRuleContext {
		public List<AnnotationValueContext> annotationValue() {
			return getRuleContexts(AnnotationValueContext.class);
		}
		public AnnotationValueContext annotationValue(int i) {
			return getRuleContext(AnnotationValueContext.class,i);
		}
		public AnnotationArgsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_annotationArgs; }
	}

	public final AnnotationArgsContext annotationArgs() throws RecognitionException {
		AnnotationArgsContext _localctx = new AnnotationArgsContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_annotationArgs);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(270);
			annotationValue();
			setState(275);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__8) {
				{
				{
				setState(271);
				match(T__8);
				setState(272);
				annotationValue();
				}
				}
				setState(277);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AnnotationValueContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(PromptDSLParser.STRING, 0); }
		public ArrayLiteralContext arrayLiteral() {
			return getRuleContext(ArrayLiteralContext.class,0);
		}
		public AnnotationValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_annotationValue; }
	}

	public final AnnotationValueContext annotationValue() throws RecognitionException {
		AnnotationValueContext _localctx = new AnnotationValueContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_annotationValue);
		try {
			setState(280);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(278);
				match(STRING);
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 2);
				{
				setState(279);
				arrayLiteral();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayLiteralContext extends ParserRuleContext {
		public List<TerminalNode> STRING() { return getTokens(PromptDSLParser.STRING); }
		public TerminalNode STRING(int i) {
			return getToken(PromptDSLParser.STRING, i);
		}
		public ArrayLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayLiteral; }
	}

	public final ArrayLiteralContext arrayLiteral() throws RecognitionException {
		ArrayLiteralContext _localctx = new ArrayLiteralContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_arrayLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(282);
			match(T__12);
			setState(291);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING) {
				{
				setState(283);
				match(STRING);
				setState(288);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__8) {
					{
					{
					setState(284);
					match(T__8);
					setState(285);
					match(STRING);
					}
					}
					setState(290);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(293);
			match(T__13);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AfterSectionContext extends ParserRuleContext {
		public AfterSectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_afterSection; }
	}

	public final AfterSectionContext afterSection() throws RecognitionException {
		AfterSectionContext _localctx = new AfterSectionContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_afterSection);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(295);
			match(T__14);
			setState(299);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,31,_ctx);
			while ( _alt!=1 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1+1 ) {
					{
					{
					setState(296);
					matchWildcard();
					}
					} 
				}
				setState(301);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,31,_ctx);
			}
			setState(302);
			match(T__15);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FixSectionContext extends ParserRuleContext {
		public FixSectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fixSection; }
	}

	public final FixSectionContext fixSection() throws RecognitionException {
		FixSectionContext _localctx = new FixSectionContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_fixSection);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(304);
			match(T__16);
			setState(308);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,32,_ctx);
			while ( _alt!=1 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1+1 ) {
					{
					{
					setState(305);
					matchWildcard();
					}
					} 
				}
				setState(310);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,32,_ctx);
			}
			setState(311);
			match(T__17);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TextBlockContext extends ParserRuleContext {
		public List<TerminalNode> STRING() { return getTokens(PromptDSLParser.STRING); }
		public TerminalNode STRING(int i) {
			return getToken(PromptDSLParser.STRING, i);
		}
		public List<TerminalNode> ID() { return getTokens(PromptDSLParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(PromptDSLParser.ID, i);
		}
		public List<TerminalNode> NUMBER() { return getTokens(PromptDSLParser.NUMBER); }
		public TerminalNode NUMBER(int i) {
			return getToken(PromptDSLParser.NUMBER, i);
		}
		public List<TerminalNode> WS() { return getTokens(PromptDSLParser.WS); }
		public TerminalNode WS(int i) {
			return getToken(PromptDSLParser.WS, i);
		}
		public TextBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_textBlock; }
	}

	public final TextBlockContext textBlock() throws RecognitionException {
		TextBlockContext _localctx = new TextBlockContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_textBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(314); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(313);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 4996180837138448L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				}
				setState(316); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 4996180837138448L) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeContext extends ParserRuleContext {
		public TerminalNode STRING_TYPE() { return getToken(PromptDSLParser.STRING_TYPE, 0); }
		public TerminalNode FLOAT_TYPE() { return getToken(PromptDSLParser.FLOAT_TYPE, 0); }
		public TerminalNode INT_TYPE() { return getToken(PromptDSLParser.INT_TYPE, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode STRUCT() { return getToken(PromptDSLParser.STRUCT, 0); }
		public List<FieldDefContext> fieldDef() {
			return getRuleContexts(FieldDefContext.class);
		}
		public FieldDefContext fieldDef(int i) {
			return getRuleContext(FieldDefContext.class,i);
		}
		public TerminalNode ID() { return getToken(PromptDSLParser.ID, 0); }
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_type);
		int _la;
		try {
			setState(333);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING_TYPE:
				enterOuterAlt(_localctx, 1);
				{
				setState(318);
				match(STRING_TYPE);
				}
				break;
			case FLOAT_TYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(319);
				match(FLOAT_TYPE);
				}
				break;
			case INT_TYPE:
				enterOuterAlt(_localctx, 3);
				{
				setState(320);
				match(INT_TYPE);
				}
				break;
			case T__19:
				enterOuterAlt(_localctx, 4);
				{
				setState(321);
				match(T__19);
				setState(322);
				type();
				}
				break;
			case STRUCT:
				enterOuterAlt(_localctx, 5);
				{
				setState(323);
				match(STRUCT);
				setState(324);
				match(T__0);
				setState(328);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==ID) {
					{
					{
					setState(325);
					fieldDef();
					}
					}
					setState(330);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(331);
				match(T__1);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 6);
				{
				setState(332);
				match(ID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ValueContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(PromptDSLParser.STRING, 0); }
		public TerminalNode NUMBER() { return getToken(PromptDSLParser.NUMBER, 0); }
		public TerminalNode BOOL() { return getToken(PromptDSLParser.BOOL, 0); }
		public ValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value; }
	}

	public final ValueContext value() throws RecognitionException {
		ValueContext _localctx = new ValueContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_value);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(335);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 985162418487296L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FormatTypeContext extends ParserRuleContext {
		public FormatTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_formatType; }
	}

	public final FormatTypeContext formatType() throws RecognitionException {
		FormatTypeContext _localctx = new FormatTypeContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_formatType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(337);
			_la = _input.LA(1);
			if ( !(_la==T__20 || _la==T__21) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u00016\u0154\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0001\u0000\u0004\u0000B\b\u0000\u000b\u0000"+
		"\f\u0000C\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0004\u0001L\b\u0001\u000b\u0001\f\u0001M\u0001\u0001\u0001"+
		"\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002[\b\u0002\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0004\u0003`\b\u0003\u000b\u0003\f\u0003"+
		"a\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0005\u0004"+
		"i\b\u0004\n\u0004\f\u0004l\t\u0004\u0001\u0004\u0001\u0004\u0001\u0005"+
		"\u0001\u0005\u0003\u0005r\b\u0005\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0004\u0006w\b\u0006\u000b\u0006\f\u0006x\u0001\u0006\u0001\u0006\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0003\u0007\u0080\b\u0007\u0001\b\u0001"+
		"\b\u0004\b\u0084\b\b\u000b\b\f\b\u0085\u0001\b\u0001\b\u0001\t\u0001\t"+
		"\u0001\t\u0001\n\u0001\n\u0001\n\u0004\n\u0090\b\n\u000b\n\f\n\u0091\u0001"+
		"\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0004\u000b\u0099\b\u000b"+
		"\u000b\u000b\f\u000b\u009a\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0003\f\u00a3\b\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0005\r\u00ab\b\r\n\r\f\r\u00ae\t\r\u0001\r\u0001\r\u0001\r\u0001\r"+
		"\u0005\r\u00b4\b\r\n\r\f\r\u00b7\t\r\u0001\r\u0003\r\u00ba\b\r\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u00c1\b\u000e"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0004\u000f\u00c6\b\u000f\u000b\u000f"+
		"\f\u000f\u00c7\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0005\u0010\u00d1\b\u0010\n\u0010\f\u0010\u00d4"+
		"\t\u0010\u0003\u0010\u00d6\b\u0010\u0001\u0010\u0001\u0010\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0003\u0011\u00de\b\u0011\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u00e5\b\u0012"+
		"\u0001\u0012\u0005\u0012\u00e8\b\u0012\n\u0012\f\u0012\u00eb\t\u0012\u0001"+
		"\u0012\u0003\u0012\u00ee\b\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0003"+
		"\u0013\u00f3\b\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0005\u0014\u00f8"+
		"\b\u0014\n\u0014\f\u0014\u00fb\t\u0014\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0004\u0015\u0101\b\u0015\u000b\u0015\f\u0015\u0102\u0001"+
		"\u0015\u0001\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0003"+
		"\u0016\u010b\b\u0016\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0005\u0017\u0112\b\u0017\n\u0017\f\u0017\u0115\t\u0017\u0001\u0018"+
		"\u0001\u0018\u0003\u0018\u0119\b\u0018\u0001\u0019\u0001\u0019\u0001\u0019"+
		"\u0001\u0019\u0005\u0019\u011f\b\u0019\n\u0019\f\u0019\u0122\t\u0019\u0003"+
		"\u0019\u0124\b\u0019\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0005"+
		"\u001a\u012a\b\u001a\n\u001a\f\u001a\u012d\t\u001a\u0001\u001a\u0001\u001a"+
		"\u0001\u001b\u0001\u001b\u0005\u001b\u0133\b\u001b\n\u001b\f\u001b\u0136"+
		"\t\u001b\u0001\u001b\u0001\u001b\u0001\u001c\u0004\u001c\u013b\b\u001c"+
		"\u000b\u001c\f\u001c\u013c\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d"+
		"\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0005\u001d\u0147\b\u001d"+
		"\n\u001d\f\u001d\u014a\t\u001d\u0001\u001d\u0001\u001d\u0003\u001d\u014e"+
		"\b\u001d\u0001\u001e\u0001\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0002"+
		"\u012b\u0134\u0000 \u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014"+
		"\u0016\u0018\u001a\u001c\u001e \"$&(*,.02468:<>\u0000\u0005\u0001\u0000"+
		"\u0007\b\u0004\u0000 !%%\'\'..\u0004\u0000\u0004\u0004\u0013\u0013.04"+
		"4\u0001\u0000/1\u0001\u0000\u0015\u0016\u0167\u0000A\u0001\u0000\u0000"+
		"\u0000\u0002G\u0001\u0000\u0000\u0000\u0004Z\u0001\u0000\u0000\u0000\u0006"+
		"\\\u0001\u0000\u0000\u0000\be\u0001\u0000\u0000\u0000\nq\u0001\u0000\u0000"+
		"\u0000\fs\u0001\u0000\u0000\u0000\u000e|\u0001\u0000\u0000\u0000\u0010"+
		"\u0081\u0001\u0000\u0000\u0000\u0012\u0089\u0001\u0000\u0000\u0000\u0014"+
		"\u008c\u0001\u0000\u0000\u0000\u0016\u0095\u0001\u0000\u0000\u0000\u0018"+
		"\u00a2\u0001\u0000\u0000\u0000\u001a\u00a4\u0001\u0000\u0000\u0000\u001c"+
		"\u00c0\u0001\u0000\u0000\u0000\u001e\u00c2\u0001\u0000\u0000\u0000 \u00cb"+
		"\u0001\u0000\u0000\u0000\"\u00dd\u0001\u0000\u0000\u0000$\u00df\u0001"+
		"\u0000\u0000\u0000&\u00f2\u0001\u0000\u0000\u0000(\u00f4\u0001\u0000\u0000"+
		"\u0000*\u00fc\u0001\u0000\u0000\u0000,\u0106\u0001\u0000\u0000\u0000."+
		"\u010e\u0001\u0000\u0000\u00000\u0118\u0001\u0000\u0000\u00002\u011a\u0001"+
		"\u0000\u0000\u00004\u0127\u0001\u0000\u0000\u00006\u0130\u0001\u0000\u0000"+
		"\u00008\u013a\u0001\u0000\u0000\u0000:\u014d\u0001\u0000\u0000\u0000<"+
		"\u014f\u0001\u0000\u0000\u0000>\u0151\u0001\u0000\u0000\u0000@B\u0003"+
		"\u0002\u0001\u0000A@\u0001\u0000\u0000\u0000BC\u0001\u0000\u0000\u0000"+
		"CA\u0001\u0000\u0000\u0000CD\u0001\u0000\u0000\u0000DE\u0001\u0000\u0000"+
		"\u0000EF\u0005\u0000\u0000\u0001F\u0001\u0001\u0000\u0000\u0000GH\u0005"+
		"\u001b\u0000\u0000HI\u0005.\u0000\u0000IK\u0005\u0001\u0000\u0000JL\u0003"+
		"\u0004\u0002\u0000KJ\u0001\u0000\u0000\u0000LM\u0001\u0000\u0000\u0000"+
		"MK\u0001\u0000\u0000\u0000MN\u0001\u0000\u0000\u0000NO\u0001\u0000\u0000"+
		"\u0000OP\u0005\u0002\u0000\u0000P\u0003\u0001\u0000\u0000\u0000Q[\u0003"+
		"\f\u0006\u0000R[\u0003\u000e\u0007\u0000S[\u0003\u0006\u0003\u0000T[\u0003"+
		"\u0014\n\u0000U[\u0003\u0016\u000b\u0000V[\u0003\u001e\u000f\u0000W[\u0003"+
		"4\u001a\u0000X[\u00036\u001b\u0000Y[\u0003\b\u0004\u0000ZQ\u0001\u0000"+
		"\u0000\u0000ZR\u0001\u0000\u0000\u0000ZS\u0001\u0000\u0000\u0000ZT\u0001"+
		"\u0000\u0000\u0000ZU\u0001\u0000\u0000\u0000ZV\u0001\u0000\u0000\u0000"+
		"ZW\u0001\u0000\u0000\u0000ZX\u0001\u0000\u0000\u0000ZY\u0001\u0000\u0000"+
		"\u0000[\u0005\u0001\u0000\u0000\u0000\\]\u0005\u0003\u0000\u0000]_\u0005"+
		"\u0001\u0000\u0000^`\u0005.\u0000\u0000_^\u0001\u0000\u0000\u0000`a\u0001"+
		"\u0000\u0000\u0000a_\u0001\u0000\u0000\u0000ab\u0001\u0000\u0000\u0000"+
		"bc\u0001\u0000\u0000\u0000cd\u0005\u0002\u0000\u0000d\u0007\u0001\u0000"+
		"\u0000\u0000ef\u0005.\u0000\u0000fj\u0005\u0001\u0000\u0000gi\u0003\n"+
		"\u0005\u0000hg\u0001\u0000\u0000\u0000il\u0001\u0000\u0000\u0000jh\u0001"+
		"\u0000\u0000\u0000jk\u0001\u0000\u0000\u0000km\u0001\u0000\u0000\u0000"+
		"lj\u0001\u0000\u0000\u0000mn\u0005\u0002\u0000\u0000n\t\u0001\u0000\u0000"+
		"\u0000or\u0003&\u0013\u0000pr\u0003(\u0014\u0000qo\u0001\u0000\u0000\u0000"+
		"qp\u0001\u0000\u0000\u0000r\u000b\u0001\u0000\u0000\u0000st\u0005 \u0000"+
		"\u0000tv\u0005\u0001\u0000\u0000uw\u0003$\u0012\u0000vu\u0001\u0000\u0000"+
		"\u0000wx\u0001\u0000\u0000\u0000xv\u0001\u0000\u0000\u0000xy\u0001\u0000"+
		"\u0000\u0000yz\u0001\u0000\u0000\u0000z{\u0005\u0002\u0000\u0000{\r\u0001"+
		"\u0000\u0000\u0000|\u007f\u0005!\u0000\u0000}\u0080\u0003\u0010\b\u0000"+
		"~\u0080\u0003\u0012\t\u0000\u007f}\u0001\u0000\u0000\u0000\u007f~\u0001"+
		"\u0000\u0000\u0000\u0080\u000f\u0001\u0000\u0000\u0000\u0081\u0083\u0005"+
		"\u0001\u0000\u0000\u0082\u0084\u0003$\u0012\u0000\u0083\u0082\u0001\u0000"+
		"\u0000\u0000\u0084\u0085\u0001\u0000\u0000\u0000\u0085\u0083\u0001\u0000"+
		"\u0000\u0000\u0085\u0086\u0001\u0000\u0000\u0000\u0086\u0087\u0001\u0000"+
		"\u0000\u0000\u0087\u0088\u0005\u0002\u0000\u0000\u0088\u0011\u0001\u0000"+
		"\u0000\u0000\u0089\u008a\u0005\u0004\u0000\u0000\u008a\u008b\u0005*\u0000"+
		"\u0000\u008b\u0013\u0001\u0000\u0000\u0000\u008c\u008d\u0005\u001d\u0000"+
		"\u0000\u008d\u008f\u0005\u0001\u0000\u0000\u008e\u0090\u0003&\u0013\u0000"+
		"\u008f\u008e\u0001\u0000\u0000\u0000\u0090\u0091\u0001\u0000\u0000\u0000"+
		"\u0091\u008f\u0001\u0000\u0000\u0000\u0091\u0092\u0001\u0000\u0000\u0000"+
		"\u0092\u0093\u0001\u0000\u0000\u0000\u0093\u0094\u0005\u0002\u0000\u0000"+
		"\u0094\u0015\u0001\u0000\u0000\u0000\u0095\u0096\u0005\u001e\u0000\u0000"+
		"\u0096\u0098\u0005\u0001\u0000\u0000\u0097\u0099\u0003\u0018\f\u0000\u0098"+
		"\u0097\u0001\u0000\u0000\u0000\u0099\u009a\u0001\u0000\u0000\u0000\u009a"+
		"\u0098\u0001\u0000\u0000\u0000\u009a\u009b\u0001\u0000\u0000\u0000\u009b"+
		"\u009c\u0001\u0000\u0000\u0000\u009c\u009d\u0005\u0002\u0000\u0000\u009d"+
		"\u0017\u0001\u0000\u0000\u0000\u009e\u00a3\u0003\u001a\r\u0000\u009f\u00a3"+
		"\u0003&\u0013\u0000\u00a0\u00a3\u0005-\u0000\u0000\u00a1\u00a3\u0003\""+
		"\u0011\u0000\u00a2\u009e\u0001\u0000\u0000\u0000\u00a2\u009f\u0001\u0000"+
		"\u0000\u0000\u00a2\u00a0\u0001\u0000\u0000\u0000\u00a2\u00a1\u0001\u0000"+
		"\u0000\u0000\u00a3\u0019\u0001\u0000\u0000\u0000\u00a4\u00a5\u0005+\u0000"+
		"\u0000\u00a5\u00a6\u0005\u0005\u0000\u0000\u00a6\u00a7\u0003\u001c\u000e"+
		"\u0000\u00a7\u00a8\u0005\u0006\u0000\u0000\u00a8\u00ac\u0005\u0001\u0000"+
		"\u0000\u00a9\u00ab\u0003\u0018\f\u0000\u00aa\u00a9\u0001\u0000\u0000\u0000"+
		"\u00ab\u00ae\u0001\u0000\u0000\u0000\u00ac\u00aa\u0001\u0000\u0000\u0000"+
		"\u00ac\u00ad\u0001\u0000\u0000\u0000\u00ad\u00af\u0001\u0000\u0000\u0000"+
		"\u00ae\u00ac\u0001\u0000\u0000\u0000\u00af\u00b9\u0005\u0002\u0000\u0000"+
		"\u00b0\u00b1\u0005,\u0000\u0000\u00b1\u00b5\u0005\u0001\u0000\u0000\u00b2"+
		"\u00b4\u0003\u0018\f\u0000\u00b3\u00b2\u0001\u0000\u0000\u0000\u00b4\u00b7"+
		"\u0001\u0000\u0000\u0000\u00b5\u00b3\u0001\u0000\u0000\u0000\u00b5\u00b6"+
		"\u0001\u0000\u0000\u0000\u00b6\u00b8\u0001\u0000\u0000\u0000\u00b7\u00b5"+
		"\u0001\u0000\u0000\u0000\u00b8\u00ba\u0005\u0002\u0000\u0000\u00b9\u00b0"+
		"\u0001\u0000\u0000\u0000\u00b9\u00ba\u0001\u0000\u0000\u0000\u00ba\u001b"+
		"\u0001\u0000\u0000\u0000\u00bb\u00bc\u0003\"\u0011\u0000\u00bc\u00bd\u0007"+
		"\u0000\u0000\u0000\u00bd\u00be\u0003\"\u0011\u0000\u00be\u00c1\u0001\u0000"+
		"\u0000\u0000\u00bf\u00c1\u0003\"\u0011\u0000\u00c0\u00bb\u0001\u0000\u0000"+
		"\u0000\u00c0\u00bf\u0001\u0000\u0000\u0000\u00c1\u001d\u0001\u0000\u0000"+
		"\u0000\u00c2\u00c3\u0005\u001f\u0000\u0000\u00c3\u00c5\u0005\u0001\u0000"+
		"\u0000\u00c4\u00c6\u0003&\u0013\u0000\u00c5\u00c4\u0001\u0000\u0000\u0000"+
		"\u00c6\u00c7\u0001\u0000\u0000\u0000\u00c7\u00c5\u0001\u0000\u0000\u0000"+
		"\u00c7\u00c8\u0001\u0000\u0000\u0000\u00c8\u00c9\u0001\u0000\u0000\u0000"+
		"\u00c9\u00ca\u0005\u0002\u0000\u0000\u00ca\u001f\u0001\u0000\u0000\u0000"+
		"\u00cb\u00cc\u0003(\u0014\u0000\u00cc\u00d5\u0005\u0005\u0000\u0000\u00cd"+
		"\u00d2\u0003\"\u0011\u0000\u00ce\u00cf\u0005\t\u0000\u0000\u00cf\u00d1"+
		"\u0003\"\u0011\u0000\u00d0\u00ce\u0001\u0000\u0000\u0000\u00d1\u00d4\u0001"+
		"\u0000\u0000\u0000\u00d2\u00d0\u0001\u0000\u0000\u0000\u00d2\u00d3\u0001"+
		"\u0000\u0000\u0000\u00d3\u00d6\u0001\u0000\u0000\u0000\u00d4\u00d2\u0001"+
		"\u0000\u0000\u0000\u00d5\u00cd\u0001\u0000\u0000\u0000\u00d5\u00d6\u0001"+
		"\u0000\u0000\u0000\u00d6\u00d7\u0001\u0000\u0000\u0000\u00d7\u00d8\u0005"+
		"\u0006\u0000\u0000\u00d8!\u0001\u0000\u0000\u0000\u00d9\u00de\u0005/\u0000"+
		"\u0000\u00da\u00de\u00050\u0000\u0000\u00db\u00de\u00051\u0000\u0000\u00dc"+
		"\u00de\u0003(\u0014\u0000\u00dd\u00d9\u0001\u0000\u0000\u0000\u00dd\u00da"+
		"\u0001\u0000\u0000\u0000\u00dd\u00db\u0001\u0000\u0000\u0000\u00dd\u00dc"+
		"\u0001\u0000\u0000\u0000\u00de#\u0001\u0000\u0000\u0000\u00df\u00e0\u0005"+
		".\u0000\u0000\u00e0\u00e1\u0005\u0004\u0000\u0000\u00e1\u00e4\u0003:\u001d"+
		"\u0000\u00e2\u00e3\u0005\n\u0000\u0000\u00e3\u00e5\u0003<\u001e\u0000"+
		"\u00e4\u00e2\u0001\u0000\u0000\u0000\u00e4\u00e5\u0001\u0000\u0000\u0000"+
		"\u00e5\u00e9\u0001\u0000\u0000\u0000\u00e6\u00e8\u0003,\u0016\u0000\u00e7"+
		"\u00e6\u0001\u0000\u0000\u0000\u00e8\u00eb\u0001\u0000\u0000\u0000\u00e9"+
		"\u00e7\u0001\u0000\u0000\u0000\u00e9\u00ea\u0001\u0000\u0000\u0000\u00ea"+
		"\u00ed\u0001\u0000\u0000\u0000\u00eb\u00e9\u0001\u0000\u0000\u0000\u00ec"+
		"\u00ee\u00053\u0000\u0000\u00ed\u00ec\u0001\u0000\u0000\u0000\u00ed\u00ee"+
		"\u0001\u0000\u0000\u0000\u00ee%\u0001\u0000\u0000\u0000\u00ef\u00f3\u0005"+
		"/\u0000\u0000\u00f0\u00f3\u00055\u0000\u0000\u00f1\u00f3\u0003(\u0014"+
		"\u0000\u00f2\u00ef\u0001\u0000\u0000\u0000\u00f2\u00f0\u0001\u0000\u0000"+
		"\u0000\u00f2\u00f1\u0001\u0000\u0000\u0000\u00f3\'\u0001\u0000\u0000\u0000"+
		"\u00f4\u00f9\u0007\u0001\u0000\u0000\u00f5\u00f6\u0005\u000b\u0000\u0000"+
		"\u00f6\u00f8\u0005.\u0000\u0000\u00f7\u00f5\u0001\u0000\u0000\u0000\u00f8"+
		"\u00fb\u0001\u0000\u0000\u0000\u00f9\u00f7\u0001\u0000\u0000\u0000\u00f9"+
		"\u00fa\u0001\u0000\u0000\u0000\u00fa)\u0001\u0000\u0000\u0000\u00fb\u00f9"+
		"\u0001\u0000\u0000\u0000\u00fc\u00fd\u0005.\u0000\u0000\u00fd\u00fe\u0005"+
		"$\u0000\u0000\u00fe\u0100\u0005\u0001\u0000\u0000\u00ff\u0101\u0003$\u0012"+
		"\u0000\u0100\u00ff\u0001\u0000\u0000\u0000\u0101\u0102\u0001\u0000\u0000"+
		"\u0000\u0102\u0100\u0001\u0000\u0000\u0000\u0102\u0103\u0001\u0000\u0000"+
		"\u0000\u0103\u0104\u0001\u0000\u0000\u0000\u0104\u0105\u0005\u0002\u0000"+
		"\u0000\u0105+\u0001\u0000\u0000\u0000\u0106\u0107\u0005\f\u0000\u0000"+
		"\u0107\u0108\u0005.\u0000\u0000\u0108\u010a\u0005\u0005\u0000\u0000\u0109"+
		"\u010b\u0003.\u0017\u0000\u010a\u0109\u0001\u0000\u0000\u0000\u010a\u010b"+
		"\u0001\u0000\u0000\u0000\u010b\u010c\u0001\u0000\u0000\u0000\u010c\u010d"+
		"\u0005\u0006\u0000\u0000\u010d-\u0001\u0000\u0000\u0000\u010e\u0113\u0003"+
		"0\u0018\u0000\u010f\u0110\u0005\t\u0000\u0000\u0110\u0112\u00030\u0018"+
		"\u0000\u0111\u010f\u0001\u0000\u0000\u0000\u0112\u0115\u0001\u0000\u0000"+
		"\u0000\u0113\u0111\u0001\u0000\u0000\u0000\u0113\u0114\u0001\u0000\u0000"+
		"\u0000\u0114/\u0001\u0000\u0000\u0000\u0115\u0113\u0001\u0000\u0000\u0000"+
		"\u0116\u0119\u0005/\u0000\u0000\u0117\u0119\u00032\u0019\u0000\u0118\u0116"+
		"\u0001\u0000\u0000\u0000\u0118\u0117\u0001\u0000\u0000\u0000\u01191\u0001"+
		"\u0000\u0000\u0000\u011a\u0123\u0005\r\u0000\u0000\u011b\u0120\u0005/"+
		"\u0000\u0000\u011c\u011d\u0005\t\u0000\u0000\u011d\u011f\u0005/\u0000"+
		"\u0000\u011e\u011c\u0001\u0000\u0000\u0000\u011f\u0122\u0001\u0000\u0000"+
		"\u0000\u0120\u011e\u0001\u0000\u0000\u0000\u0120\u0121\u0001\u0000\u0000"+
		"\u0000\u0121\u0124\u0001\u0000\u0000\u0000\u0122\u0120\u0001\u0000\u0000"+
		"\u0000\u0123\u011b\u0001\u0000\u0000\u0000\u0123\u0124\u0001\u0000\u0000"+
		"\u0000\u0124\u0125\u0001\u0000\u0000\u0000\u0125\u0126\u0005\u000e\u0000"+
		"\u0000\u01263\u0001\u0000\u0000\u0000\u0127\u012b\u0005\u000f\u0000\u0000"+
		"\u0128\u012a\t\u0000\u0000\u0000\u0129\u0128\u0001\u0000\u0000\u0000\u012a"+
		"\u012d\u0001\u0000\u0000\u0000\u012b\u012c\u0001\u0000\u0000\u0000\u012b"+
		"\u0129\u0001\u0000\u0000\u0000\u012c\u012e\u0001\u0000\u0000\u0000\u012d"+
		"\u012b\u0001\u0000\u0000\u0000\u012e\u012f\u0005\u0010\u0000\u0000\u012f"+
		"5\u0001\u0000\u0000\u0000\u0130\u0134\u0005\u0011\u0000\u0000\u0131\u0133"+
		"\t\u0000\u0000\u0000\u0132\u0131\u0001\u0000\u0000\u0000\u0133\u0136\u0001"+
		"\u0000\u0000\u0000\u0134\u0135\u0001\u0000\u0000\u0000\u0134\u0132\u0001"+
		"\u0000\u0000\u0000\u0135\u0137\u0001\u0000\u0000\u0000\u0136\u0134\u0001"+
		"\u0000\u0000\u0000\u0137\u0138\u0005\u0012\u0000\u0000\u01387\u0001\u0000"+
		"\u0000\u0000\u0139\u013b\u0007\u0002\u0000\u0000\u013a\u0139\u0001\u0000"+
		"\u0000\u0000\u013b\u013c\u0001\u0000\u0000\u0000\u013c\u013a\u0001\u0000"+
		"\u0000\u0000\u013c\u013d\u0001\u0000\u0000\u0000\u013d9\u0001\u0000\u0000"+
		"\u0000\u013e\u014e\u0005\u0018\u0000\u0000\u013f\u014e\u0005\u0019\u0000"+
		"\u0000\u0140\u014e\u0005\u001a\u0000\u0000\u0141\u0142\u0005\u0014\u0000"+
		"\u0000\u0142\u014e\u0003:\u001d\u0000\u0143\u0144\u0005$\u0000\u0000\u0144"+
		"\u0148\u0005\u0001\u0000\u0000\u0145\u0147\u0003$\u0012\u0000\u0146\u0145"+
		"\u0001\u0000\u0000\u0000\u0147\u014a\u0001\u0000\u0000\u0000\u0148\u0146"+
		"\u0001\u0000\u0000\u0000\u0148\u0149\u0001\u0000\u0000\u0000\u0149\u014b"+
		"\u0001\u0000\u0000\u0000\u014a\u0148\u0001\u0000\u0000\u0000\u014b\u014e"+
		"\u0005\u0002\u0000\u0000\u014c\u014e\u0005.\u0000\u0000\u014d\u013e\u0001"+
		"\u0000\u0000\u0000\u014d\u013f\u0001\u0000\u0000\u0000\u014d\u0140\u0001"+
		"\u0000\u0000\u0000\u014d\u0141\u0001\u0000\u0000\u0000\u014d\u0143\u0001"+
		"\u0000\u0000\u0000\u014d\u014c\u0001\u0000\u0000\u0000\u014e;\u0001\u0000"+
		"\u0000\u0000\u014f\u0150\u0007\u0003\u0000\u0000\u0150=\u0001\u0000\u0000"+
		"\u0000\u0151\u0152\u0007\u0004\u0000\u0000\u0152?\u0001\u0000\u0000\u0000"+
		"$CMZajqx\u007f\u0085\u0091\u009a\u00a2\u00ac\u00b5\u00b9\u00c0\u00c7\u00d2"+
		"\u00d5\u00dd\u00e4\u00e9\u00ed\u00f2\u00f9\u0102\u010a\u0113\u0118\u0120"+
		"\u0123\u012b\u0134\u013c\u0148\u014d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}