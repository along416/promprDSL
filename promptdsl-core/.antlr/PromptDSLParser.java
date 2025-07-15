// Generated from d:/work/promptDSL/promptdsl-core/PromptDSL.g4 by ANTLR 4.13.1
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
		T__17=18, T__18=19, T__19=20, T__20=21, PLUS=22, STRING_TYPE=23, FLOAT_TYPE=24, 
		INT_TYPE=25, PROMPT=26, PARAMS=27, SYSTEM=28, USER=29, NOTE=30, INPUT=31, 
		OUTPUT=32, FORMAT=33, TYPE=34, STRUCT=35, BEFORE=36, SCHEMA=37, AFTER=38, 
		PARSE=39, JSONFIX=40, MARKDOWN=41, IF=42, ELSE=43, OUTPUTSPEC=44, ID=45, 
		STRING=46, NUMBER=47, BOOL=48, PIPE=49, SEMI=50, WS=51, LINE_COMMENT=52, 
		BLOCK_COMMENT=53;
	public static final int
		RULE_promptFile = 0, RULE_promptDef = 1, RULE_promptBlock = 2, RULE_inputSection = 3, 
		RULE_outputSection = 4, RULE_outputStruct = 5, RULE_outputMarkdown = 6, 
		RULE_systemSection = 7, RULE_sysContent = 8, RULE_moduleDef = 9, RULE_moduleContent = 10, 
		RULE_userSection = 11, RULE_userContent = 12, RULE_thencontent = 13, RULE_elsecontent = 14, 
		RULE_ifStatement = 15, RULE_condition = 16, RULE_noteSection = 17, RULE_dslCallExpr = 18, 
		RULE_expr = 19, RULE_fieldDef = 20, RULE_textLine = 21, RULE_paramPath = 22, 
		RULE_structDef = 23, RULE_annotation = 24, RULE_annotationArgs = 25, RULE_annotationValue = 26, 
		RULE_arrayLiteral = 27, RULE_afterSection = 28, RULE_fixSection = 29, 
		RULE_textBlock = 30, RULE_type = 31, RULE_value = 32, RULE_formatType = 33;
	private static String[] makeRuleNames() {
		return new String[] {
			"promptFile", "promptDef", "promptBlock", "inputSection", "outputSection", 
			"outputStruct", "outputMarkdown", "systemSection", "sysContent", "moduleDef", 
			"moduleContent", "userSection", "userContent", "thencontent", "elsecontent", 
			"ifStatement", "condition", "noteSection", "dslCallExpr", "expr", "fieldDef", 
			"textLine", "paramPath", "structDef", "annotation", "annotationArgs", 
			"annotationValue", "arrayLiteral", "afterSection", "fixSection", "textBlock", 
			"type", "value", "formatType"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'{'", "'}'", "':'", "'('", "')'", "'=='", "'!='", "','", "'='", 
			"'.'", "'@'", "'['", "']'", "'<after>'", "'</after>'", "'<fix>'", "'</fix>'", 
			"'-'", "'[]'", "'md'", "'json'", "'+'", "'string'", "'float'", "'int'", 
			"'prompt'", "'params'", "'system'", "'user'", "'note'", "'input'", "'output'", 
			"'format'", "'type'", "'struct'", "'before'", "'schema'", "'after'", 
			"'parse'", "'jsonfix'", "'markdown'", "'if'", "'else'", "'outputspec'", 
			null, null, null, null, "'|'", "';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, "PLUS", "STRING_TYPE", 
			"FLOAT_TYPE", "INT_TYPE", "PROMPT", "PARAMS", "SYSTEM", "USER", "NOTE", 
			"INPUT", "OUTPUT", "FORMAT", "TYPE", "STRUCT", "BEFORE", "SCHEMA", "AFTER", 
			"PARSE", "JSONFIX", "MARKDOWN", "IF", "ELSE", "OUTPUTSPEC", "ID", "STRING", 
			"NUMBER", "BOOL", "PIPE", "SEMI", "WS", "LINE_COMMENT", "BLOCK_COMMENT"
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
			setState(69); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(68);
				promptDef();
				}
				}
				setState(71); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==PROMPT );
			setState(73);
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
			setState(75);
			match(PROMPT);
			setState(76);
			match(ID);
			setState(77);
			match(T__0);
			setState(79); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(78);
				promptBlock();
				}
				}
				setState(81); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 35192693669888L) != 0) );
			setState(83);
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
			setState(93);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INPUT:
				enterOuterAlt(_localctx, 1);
				{
				setState(85);
				inputSection();
				}
				break;
			case OUTPUT:
				enterOuterAlt(_localctx, 2);
				{
				setState(86);
				outputSection();
				}
				break;
			case SYSTEM:
				enterOuterAlt(_localctx, 3);
				{
				setState(87);
				systemSection();
				}
				break;
			case USER:
				enterOuterAlt(_localctx, 4);
				{
				setState(88);
				userSection();
				}
				break;
			case NOTE:
				enterOuterAlt(_localctx, 5);
				{
				setState(89);
				noteSection();
				}
				break;
			case T__13:
				enterOuterAlt(_localctx, 6);
				{
				setState(90);
				afterSection();
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 7);
				{
				setState(91);
				fixSection();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 8);
				{
				setState(92);
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
		enterRule(_localctx, 6, RULE_inputSection);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
			match(INPUT);
			setState(96);
			match(T__0);
			setState(98); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(97);
				fieldDef();
				}
				}
				setState(100); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			setState(102);
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
		enterRule(_localctx, 8, RULE_outputSection);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(104);
			match(OUTPUT);
			setState(107);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				{
				setState(105);
				outputStruct();
				}
				break;
			case T__2:
				{
				setState(106);
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
		enterRule(_localctx, 10, RULE_outputStruct);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(109);
			match(T__0);
			setState(111); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(110);
				fieldDef();
				}
				}
				setState(113); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			setState(115);
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
		enterRule(_localctx, 12, RULE_outputMarkdown);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			match(T__2);
			setState(118);
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
		public List<TerminalNode> ID() { return getTokens(PromptDSLParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(PromptDSLParser.ID, i);
		}
		public List<SysContentContext> sysContent() {
			return getRuleContexts(SysContentContext.class);
		}
		public SysContentContext sysContent(int i) {
			return getRuleContext(SysContentContext.class,i);
		}
		public SystemSectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_systemSection; }
	}

	public final SystemSectionContext systemSection() throws RecognitionException {
		SystemSectionContext _localctx = new SystemSectionContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_systemSection);
		int _la;
		try {
			setState(137);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(120);
				match(SYSTEM);
				setState(121);
				match(T__0);
				setState(123); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(122);
					match(ID);
					}
					}
					setState(125); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==ID );
				setState(127);
				match(T__1);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(128);
				match(SYSTEM);
				setState(129);
				match(T__0);
				setState(131); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(130);
					sysContent();
					}
					}
					setState(133); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 5036113295048704L) != 0) );
				setState(135);
				match(T__1);
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
	public static class SysContentContext extends ParserRuleContext {
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TextLineContext textLine() {
			return getRuleContext(TextLineContext.class,0);
		}
		public SysContentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sysContent; }
	}

	public final SysContentContext sysContent() throws RecognitionException {
		SysContentContext _localctx = new SysContentContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_sysContent);
		try {
			setState(142);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(139);
				ifStatement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(140);
				expr();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(141);
				textLine();
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
		enterRule(_localctx, 18, RULE_moduleDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			match(ID);
			setState(145);
			match(T__0);
			setState(149);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4609502783471616L) != 0)) {
				{
				{
				setState(146);
				moduleContent();
				}
				}
				setState(151);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(152);
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
		public ParamPathContext paramPath() {
			return getRuleContext(ParamPathContext.class,0);
		}
		public TextLineContext textLine() {
			return getRuleContext(TextLineContext.class,0);
		}
		public ModuleContentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_moduleContent; }
	}

	public final ModuleContentContext moduleContent() throws RecognitionException {
		ModuleContentContext _localctx = new ModuleContentContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_moduleContent);
		try {
			setState(156);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(154);
				paramPath();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(155);
				textLine();
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
			setState(158);
			match(USER);
			setState(159);
			match(T__0);
			setState(161); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(160);
				userContent();
				}
				}
				setState(163); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 5053705481093120L) != 0) );
			setState(165);
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
		public ParamPathContext paramPath() {
			return getRuleContext(ParamPathContext.class,0);
		}
		public TerminalNode OUTPUTSPEC() { return getToken(PromptDSLParser.OUTPUTSPEC, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TextLineContext textLine() {
			return getRuleContext(TextLineContext.class,0);
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
			setState(172);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(167);
				ifStatement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(168);
				paramPath();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(169);
				match(OUTPUTSPEC);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(170);
				expr();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(171);
				textLine();
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
	public static class ThencontentContext extends ParserRuleContext {
		public UserContentContext userContent() {
			return getRuleContext(UserContentContext.class,0);
		}
		public ThencontentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_thencontent; }
	}

	public final ThencontentContext thencontent() throws RecognitionException {
		ThencontentContext _localctx = new ThencontentContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_thencontent);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(174);
			userContent();
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
	public static class ElsecontentContext extends ParserRuleContext {
		public UserContentContext userContent() {
			return getRuleContext(UserContentContext.class,0);
		}
		public ElsecontentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elsecontent; }
	}

	public final ElsecontentContext elsecontent() throws RecognitionException {
		ElsecontentContext _localctx = new ElsecontentContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_elsecontent);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(176);
			userContent();
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
		public List<ThencontentContext> thencontent() {
			return getRuleContexts(ThencontentContext.class);
		}
		public ThencontentContext thencontent(int i) {
			return getRuleContext(ThencontentContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(PromptDSLParser.ELSE, 0); }
		public List<ElsecontentContext> elsecontent() {
			return getRuleContexts(ElsecontentContext.class);
		}
		public ElsecontentContext elsecontent(int i) {
			return getRuleContext(ElsecontentContext.class,i);
		}
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_ifStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
			match(IF);
			setState(179);
			match(T__3);
			setState(180);
			condition();
			setState(181);
			match(T__4);
			setState(182);
			match(T__0);
			setState(186);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 5053705481093120L) != 0)) {
				{
				{
				setState(183);
				thencontent();
				}
				}
				setState(188);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(189);
			match(T__1);
			setState(199);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(190);
				match(ELSE);
				setState(191);
				match(T__0);
				setState(195);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 5053705481093120L) != 0)) {
					{
					{
					setState(192);
					elsecontent();
					}
					}
					setState(197);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(198);
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
		public ExprContext lhs;
		public Token op;
		public ExprContext rhs;
		public ExprContext single;
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
		enterRule(_localctx, 32, RULE_condition);
		int _la;
		try {
			setState(206);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(201);
				((ConditionContext)_localctx).lhs = expr();
				setState(202);
				((ConditionContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==T__5 || _la==T__6) ) {
					((ConditionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(203);
				((ConditionContext)_localctx).rhs = expr();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(205);
				((ConditionContext)_localctx).single = expr();
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
		enterRule(_localctx, 34, RULE_noteSection);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(208);
			match(NOTE);
			setState(209);
			match(T__0);
			setState(211); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(210);
				textLine();
				}
				}
				setState(213); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 4609502783471616L) != 0) );
			setState(215);
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
		enterRule(_localctx, 36, RULE_dslCallExpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(217);
			paramPath();
			setState(218);
			match(T__3);
			setState(227);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 528115621167104L) != 0)) {
				{
				setState(219);
				expr();
				setState(224);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__7) {
					{
					{
					setState(220);
					match(T__7);
					setState(221);
					expr();
					}
					}
					setState(226);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(229);
			match(T__4);
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
		public ParamPathContext paramPath() {
			return getRuleContext(ParamPathContext.class,0);
		}
		public TerminalNode STRING() { return getToken(PromptDSLParser.STRING, 0); }
		public TerminalNode NUMBER() { return getToken(PromptDSLParser.NUMBER, 0); }
		public TerminalNode BOOL() { return getToken(PromptDSLParser.BOOL, 0); }
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	}

	public final ExprContext expr() throws RecognitionException {
		ExprContext _localctx = new ExprContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_expr);
		try {
			setState(235);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INPUT:
			case OUTPUT:
			case BEFORE:
			case AFTER:
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				setState(231);
				paramPath();
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 2);
				{
				setState(232);
				match(STRING);
				}
				break;
			case NUMBER:
				enterOuterAlt(_localctx, 3);
				{
				setState(233);
				match(NUMBER);
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(234);
				match(BOOL);
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
		enterRule(_localctx, 40, RULE_fieldDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(237);
			match(ID);
			setState(238);
			match(T__2);
			setState(239);
			type();
			setState(242);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__8) {
				{
				setState(240);
				match(T__8);
				setState(241);
				value();
				}
			}

			setState(247);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__10) {
				{
				{
				setState(244);
				annotation();
				}
				}
				setState(249);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(251);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(250);
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
		enterRule(_localctx, 42, RULE_textLine);
		try {
			setState(256);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(253);
				match(STRING);
				}
				break;
			case LINE_COMMENT:
				enterOuterAlt(_localctx, 2);
				{
				setState(254);
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
				setState(255);
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
		enterRule(_localctx, 44, RULE_paramPath);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(258);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 35534411923456L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(263);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__9) {
				{
				{
				setState(259);
				match(T__9);
				setState(260);
				match(ID);
				}
				}
				setState(265);
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
		enterRule(_localctx, 46, RULE_structDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(266);
			match(ID);
			setState(267);
			match(STRUCT);
			setState(268);
			match(T__0);
			setState(270); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(269);
				fieldDef();
				}
				}
				setState(272); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			setState(274);
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
		enterRule(_localctx, 48, RULE_annotation);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(276);
			match(T__10);
			setState(277);
			match(ID);
			setState(278);
			match(T__3);
			setState(280);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11 || _la==STRING) {
				{
				setState(279);
				annotationArgs();
				}
			}

			setState(282);
			match(T__4);
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
		enterRule(_localctx, 50, RULE_annotationArgs);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(284);
			annotationValue();
			setState(289);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__7) {
				{
				{
				setState(285);
				match(T__7);
				setState(286);
				annotationValue();
				}
				}
				setState(291);
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
		enterRule(_localctx, 52, RULE_annotationValue);
		try {
			setState(294);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(292);
				match(STRING);
				}
				break;
			case T__11:
				enterOuterAlt(_localctx, 2);
				{
				setState(293);
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
		enterRule(_localctx, 54, RULE_arrayLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(296);
			match(T__11);
			setState(305);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING) {
				{
				setState(297);
				match(STRING);
				setState(302);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__7) {
					{
					{
					setState(298);
					match(T__7);
					setState(299);
					match(STRING);
					}
					}
					setState(304);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(307);
			match(T__12);
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
		enterRule(_localctx, 56, RULE_afterSection);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(309);
			match(T__13);
			setState(313);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,33,_ctx);
			while ( _alt!=1 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1+1 ) {
					{
					{
					setState(310);
					matchWildcard();
					}
					} 
				}
				setState(315);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,33,_ctx);
			}
			setState(316);
			match(T__14);
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
		enterRule(_localctx, 58, RULE_fixSection);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(318);
			match(T__15);
			setState(322);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,34,_ctx);
			while ( _alt!=1 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1+1 ) {
					{
					{
					setState(319);
					matchWildcard();
					}
					} 
				}
				setState(324);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,34,_ctx);
			}
			setState(325);
			match(T__16);
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
		enterRule(_localctx, 60, RULE_textBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(328); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(327);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 2498090418569224L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				}
				setState(330); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 2498090418569224L) != 0) );
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
		enterRule(_localctx, 62, RULE_type);
		int _la;
		try {
			setState(347);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING_TYPE:
				enterOuterAlt(_localctx, 1);
				{
				setState(332);
				match(STRING_TYPE);
				}
				break;
			case FLOAT_TYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(333);
				match(FLOAT_TYPE);
				}
				break;
			case INT_TYPE:
				enterOuterAlt(_localctx, 3);
				{
				setState(334);
				match(INT_TYPE);
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 4);
				{
				setState(335);
				match(T__18);
				setState(336);
				type();
				}
				break;
			case STRUCT:
				enterOuterAlt(_localctx, 5);
				{
				setState(337);
				match(STRUCT);
				setState(338);
				match(T__0);
				setState(342);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==ID) {
					{
					{
					setState(339);
					fieldDef();
					}
					}
					setState(344);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(345);
				match(T__1);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 6);
				{
				setState(346);
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
		enterRule(_localctx, 64, RULE_value);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(349);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 492581209243648L) != 0)) ) {
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
		enterRule(_localctx, 66, RULE_formatType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(351);
			_la = _input.LA(1);
			if ( !(_la==T__19 || _la==T__20) ) {
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
		"\u0004\u00015\u0162\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0001\u0000\u0004"+
		"\u0000F\b\u0000\u000b\u0000\f\u0000G\u0001\u0000\u0001\u0000\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0004\u0001P\b\u0001\u000b\u0001"+
		"\f\u0001Q\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002"+
		"^\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0004\u0003c\b\u0003\u000b"+
		"\u0003\f\u0003d\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0003\u0004l\b\u0004\u0001\u0005\u0001\u0005\u0004\u0005p\b\u0005"+
		"\u000b\u0005\f\u0005q\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0004\u0007|\b\u0007"+
		"\u000b\u0007\f\u0007}\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0004\u0007\u0084\b\u0007\u000b\u0007\f\u0007\u0085\u0001\u0007\u0001"+
		"\u0007\u0003\u0007\u008a\b\u0007\u0001\b\u0001\b\u0001\b\u0003\b\u008f"+
		"\b\b\u0001\t\u0001\t\u0001\t\u0005\t\u0094\b\t\n\t\f\t\u0097\t\t\u0001"+
		"\t\u0001\t\u0001\n\u0001\n\u0003\n\u009d\b\n\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0004\u000b\u00a2\b\u000b\u000b\u000b\f\u000b\u00a3\u0001\u000b"+
		"\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003\f\u00ad\b\f"+
		"\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0005\u000f\u00b9\b\u000f\n"+
		"\u000f\f\u000f\u00bc\t\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0005\u000f\u00c2\b\u000f\n\u000f\f\u000f\u00c5\t\u000f\u0001\u000f"+
		"\u0003\u000f\u00c8\b\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0003\u0010\u00cf\b\u0010\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0004\u0011\u00d4\b\u0011\u000b\u0011\f\u0011\u00d5\u0001\u0011\u0001"+
		"\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0005"+
		"\u0012\u00df\b\u0012\n\u0012\f\u0012\u00e2\t\u0012\u0003\u0012\u00e4\b"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0003\u0013\u00ec\b\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0003\u0014\u00f3\b\u0014\u0001\u0014\u0005\u0014\u00f6"+
		"\b\u0014\n\u0014\f\u0014\u00f9\t\u0014\u0001\u0014\u0003\u0014\u00fc\b"+
		"\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u0101\b\u0015\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0005\u0016\u0106\b\u0016\n\u0016\f\u0016"+
		"\u0109\t\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0004\u0017"+
		"\u010f\b\u0017\u000b\u0017\f\u0017\u0110\u0001\u0017\u0001\u0017\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0003\u0018\u0119\b\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0005\u0019\u0120"+
		"\b\u0019\n\u0019\f\u0019\u0123\t\u0019\u0001\u001a\u0001\u001a\u0003\u001a"+
		"\u0127\b\u001a\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0005\u001b"+
		"\u012d\b\u001b\n\u001b\f\u001b\u0130\t\u001b\u0003\u001b\u0132\b\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0005\u001c\u0138\b\u001c"+
		"\n\u001c\f\u001c\u013b\t\u001c\u0001\u001c\u0001\u001c\u0001\u001d\u0001"+
		"\u001d\u0005\u001d\u0141\b\u001d\n\u001d\f\u001d\u0144\t\u001d\u0001\u001d"+
		"\u0001\u001d\u0001\u001e\u0004\u001e\u0149\b\u001e\u000b\u001e\f\u001e"+
		"\u014a\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0005\u001f\u0155\b\u001f\n\u001f\f\u001f"+
		"\u0158\t\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u015c\b\u001f\u0001"+
		" \u0001 \u0001!\u0001!\u0001!\u0002\u0139\u0142\u0000\"\u0000\u0002\u0004"+
		"\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \""+
		"$&(*,.02468:<>@B\u0000\u0005\u0001\u0000\u0006\u0007\u0004\u0000\u001f"+
		" $$&&--\u0004\u0000\u0003\u0003\u0012\u0012-/33\u0001\u0000.0\u0001\u0000"+
		"\u0014\u0015\u0176\u0000E\u0001\u0000\u0000\u0000\u0002K\u0001\u0000\u0000"+
		"\u0000\u0004]\u0001\u0000\u0000\u0000\u0006_\u0001\u0000\u0000\u0000\b"+
		"h\u0001\u0000\u0000\u0000\nm\u0001\u0000\u0000\u0000\fu\u0001\u0000\u0000"+
		"\u0000\u000e\u0089\u0001\u0000\u0000\u0000\u0010\u008e\u0001\u0000\u0000"+
		"\u0000\u0012\u0090\u0001\u0000\u0000\u0000\u0014\u009c\u0001\u0000\u0000"+
		"\u0000\u0016\u009e\u0001\u0000\u0000\u0000\u0018\u00ac\u0001\u0000\u0000"+
		"\u0000\u001a\u00ae\u0001\u0000\u0000\u0000\u001c\u00b0\u0001\u0000\u0000"+
		"\u0000\u001e\u00b2\u0001\u0000\u0000\u0000 \u00ce\u0001\u0000\u0000\u0000"+
		"\"\u00d0\u0001\u0000\u0000\u0000$\u00d9\u0001\u0000\u0000\u0000&\u00eb"+
		"\u0001\u0000\u0000\u0000(\u00ed\u0001\u0000\u0000\u0000*\u0100\u0001\u0000"+
		"\u0000\u0000,\u0102\u0001\u0000\u0000\u0000.\u010a\u0001\u0000\u0000\u0000"+
		"0\u0114\u0001\u0000\u0000\u00002\u011c\u0001\u0000\u0000\u00004\u0126"+
		"\u0001\u0000\u0000\u00006\u0128\u0001\u0000\u0000\u00008\u0135\u0001\u0000"+
		"\u0000\u0000:\u013e\u0001\u0000\u0000\u0000<\u0148\u0001\u0000\u0000\u0000"+
		">\u015b\u0001\u0000\u0000\u0000@\u015d\u0001\u0000\u0000\u0000B\u015f"+
		"\u0001\u0000\u0000\u0000DF\u0003\u0002\u0001\u0000ED\u0001\u0000\u0000"+
		"\u0000FG\u0001\u0000\u0000\u0000GE\u0001\u0000\u0000\u0000GH\u0001\u0000"+
		"\u0000\u0000HI\u0001\u0000\u0000\u0000IJ\u0005\u0000\u0000\u0001J\u0001"+
		"\u0001\u0000\u0000\u0000KL\u0005\u001a\u0000\u0000LM\u0005-\u0000\u0000"+
		"MO\u0005\u0001\u0000\u0000NP\u0003\u0004\u0002\u0000ON\u0001\u0000\u0000"+
		"\u0000PQ\u0001\u0000\u0000\u0000QO\u0001\u0000\u0000\u0000QR\u0001\u0000"+
		"\u0000\u0000RS\u0001\u0000\u0000\u0000ST\u0005\u0002\u0000\u0000T\u0003"+
		"\u0001\u0000\u0000\u0000U^\u0003\u0006\u0003\u0000V^\u0003\b\u0004\u0000"+
		"W^\u0003\u000e\u0007\u0000X^\u0003\u0016\u000b\u0000Y^\u0003\"\u0011\u0000"+
		"Z^\u00038\u001c\u0000[^\u0003:\u001d\u0000\\^\u0003\u0012\t\u0000]U\u0001"+
		"\u0000\u0000\u0000]V\u0001\u0000\u0000\u0000]W\u0001\u0000\u0000\u0000"+
		"]X\u0001\u0000\u0000\u0000]Y\u0001\u0000\u0000\u0000]Z\u0001\u0000\u0000"+
		"\u0000][\u0001\u0000\u0000\u0000]\\\u0001\u0000\u0000\u0000^\u0005\u0001"+
		"\u0000\u0000\u0000_`\u0005\u001f\u0000\u0000`b\u0005\u0001\u0000\u0000"+
		"ac\u0003(\u0014\u0000ba\u0001\u0000\u0000\u0000cd\u0001\u0000\u0000\u0000"+
		"db\u0001\u0000\u0000\u0000de\u0001\u0000\u0000\u0000ef\u0001\u0000\u0000"+
		"\u0000fg\u0005\u0002\u0000\u0000g\u0007\u0001\u0000\u0000\u0000hk\u0005"+
		" \u0000\u0000il\u0003\n\u0005\u0000jl\u0003\f\u0006\u0000ki\u0001\u0000"+
		"\u0000\u0000kj\u0001\u0000\u0000\u0000l\t\u0001\u0000\u0000\u0000mo\u0005"+
		"\u0001\u0000\u0000np\u0003(\u0014\u0000on\u0001\u0000\u0000\u0000pq\u0001"+
		"\u0000\u0000\u0000qo\u0001\u0000\u0000\u0000qr\u0001\u0000\u0000\u0000"+
		"rs\u0001\u0000\u0000\u0000st\u0005\u0002\u0000\u0000t\u000b\u0001\u0000"+
		"\u0000\u0000uv\u0005\u0003\u0000\u0000vw\u0005)\u0000\u0000w\r\u0001\u0000"+
		"\u0000\u0000xy\u0005\u001c\u0000\u0000y{\u0005\u0001\u0000\u0000z|\u0005"+
		"-\u0000\u0000{z\u0001\u0000\u0000\u0000|}\u0001\u0000\u0000\u0000}{\u0001"+
		"\u0000\u0000\u0000}~\u0001\u0000\u0000\u0000~\u007f\u0001\u0000\u0000"+
		"\u0000\u007f\u008a\u0005\u0002\u0000\u0000\u0080\u0081\u0005\u001c\u0000"+
		"\u0000\u0081\u0083\u0005\u0001\u0000\u0000\u0082\u0084\u0003\u0010\b\u0000"+
		"\u0083\u0082\u0001\u0000\u0000\u0000\u0084\u0085\u0001\u0000\u0000\u0000"+
		"\u0085\u0083\u0001\u0000\u0000\u0000\u0085\u0086\u0001\u0000\u0000\u0000"+
		"\u0086\u0087\u0001\u0000\u0000\u0000\u0087\u0088\u0005\u0002\u0000\u0000"+
		"\u0088\u008a\u0001\u0000\u0000\u0000\u0089x\u0001\u0000\u0000\u0000\u0089"+
		"\u0080\u0001\u0000\u0000\u0000\u008a\u000f\u0001\u0000\u0000\u0000\u008b"+
		"\u008f\u0003\u001e\u000f\u0000\u008c\u008f\u0003&\u0013\u0000\u008d\u008f"+
		"\u0003*\u0015\u0000\u008e\u008b\u0001\u0000\u0000\u0000\u008e\u008c\u0001"+
		"\u0000\u0000\u0000\u008e\u008d\u0001\u0000\u0000\u0000\u008f\u0011\u0001"+
		"\u0000\u0000\u0000\u0090\u0091\u0005-\u0000\u0000\u0091\u0095\u0005\u0001"+
		"\u0000\u0000\u0092\u0094\u0003\u0014\n\u0000\u0093\u0092\u0001\u0000\u0000"+
		"\u0000\u0094\u0097\u0001\u0000\u0000\u0000\u0095\u0093\u0001\u0000\u0000"+
		"\u0000\u0095\u0096\u0001\u0000\u0000\u0000\u0096\u0098\u0001\u0000\u0000"+
		"\u0000\u0097\u0095\u0001\u0000\u0000\u0000\u0098\u0099\u0005\u0002\u0000"+
		"\u0000\u0099\u0013\u0001\u0000\u0000\u0000\u009a\u009d\u0003,\u0016\u0000"+
		"\u009b\u009d\u0003*\u0015\u0000\u009c\u009a\u0001\u0000\u0000\u0000\u009c"+
		"\u009b\u0001\u0000\u0000\u0000\u009d\u0015\u0001\u0000\u0000\u0000\u009e"+
		"\u009f\u0005\u001d\u0000\u0000\u009f\u00a1\u0005\u0001\u0000\u0000\u00a0"+
		"\u00a2\u0003\u0018\f\u0000\u00a1\u00a0\u0001\u0000\u0000\u0000\u00a2\u00a3"+
		"\u0001\u0000\u0000\u0000\u00a3\u00a1\u0001\u0000\u0000\u0000\u00a3\u00a4"+
		"\u0001\u0000\u0000\u0000\u00a4\u00a5\u0001\u0000\u0000\u0000\u00a5\u00a6"+
		"\u0005\u0002\u0000\u0000\u00a6\u0017\u0001\u0000\u0000\u0000\u00a7\u00ad"+
		"\u0003\u001e\u000f\u0000\u00a8\u00ad\u0003,\u0016\u0000\u00a9\u00ad\u0005"+
		",\u0000\u0000\u00aa\u00ad\u0003&\u0013\u0000\u00ab\u00ad\u0003*\u0015"+
		"\u0000\u00ac\u00a7\u0001\u0000\u0000\u0000\u00ac\u00a8\u0001\u0000\u0000"+
		"\u0000\u00ac\u00a9\u0001\u0000\u0000\u0000\u00ac\u00aa\u0001\u0000\u0000"+
		"\u0000\u00ac\u00ab\u0001\u0000\u0000\u0000\u00ad\u0019\u0001\u0000\u0000"+
		"\u0000\u00ae\u00af\u0003\u0018\f\u0000\u00af\u001b\u0001\u0000\u0000\u0000"+
		"\u00b0\u00b1\u0003\u0018\f\u0000\u00b1\u001d\u0001\u0000\u0000\u0000\u00b2"+
		"\u00b3\u0005*\u0000\u0000\u00b3\u00b4\u0005\u0004\u0000\u0000\u00b4\u00b5"+
		"\u0003 \u0010\u0000\u00b5\u00b6\u0005\u0005\u0000\u0000\u00b6\u00ba\u0005"+
		"\u0001\u0000\u0000\u00b7\u00b9\u0003\u001a\r\u0000\u00b8\u00b7\u0001\u0000"+
		"\u0000\u0000\u00b9\u00bc\u0001\u0000\u0000\u0000\u00ba\u00b8\u0001\u0000"+
		"\u0000\u0000\u00ba\u00bb\u0001\u0000\u0000\u0000\u00bb\u00bd\u0001\u0000"+
		"\u0000\u0000\u00bc\u00ba\u0001\u0000\u0000\u0000\u00bd\u00c7\u0005\u0002"+
		"\u0000\u0000\u00be\u00bf\u0005+\u0000\u0000\u00bf\u00c3\u0005\u0001\u0000"+
		"\u0000\u00c0\u00c2\u0003\u001c\u000e\u0000\u00c1\u00c0\u0001\u0000\u0000"+
		"\u0000\u00c2\u00c5\u0001\u0000\u0000\u0000\u00c3\u00c1\u0001\u0000\u0000"+
		"\u0000\u00c3\u00c4\u0001\u0000\u0000\u0000\u00c4\u00c6\u0001\u0000\u0000"+
		"\u0000\u00c5\u00c3\u0001\u0000\u0000\u0000\u00c6\u00c8\u0005\u0002\u0000"+
		"\u0000\u00c7\u00be\u0001\u0000\u0000\u0000\u00c7\u00c8\u0001\u0000\u0000"+
		"\u0000\u00c8\u001f\u0001\u0000\u0000\u0000\u00c9\u00ca\u0003&\u0013\u0000"+
		"\u00ca\u00cb\u0007\u0000\u0000\u0000\u00cb\u00cc\u0003&\u0013\u0000\u00cc"+
		"\u00cf\u0001\u0000\u0000\u0000\u00cd\u00cf\u0003&\u0013\u0000\u00ce\u00c9"+
		"\u0001\u0000\u0000\u0000\u00ce\u00cd\u0001\u0000\u0000\u0000\u00cf!\u0001"+
		"\u0000\u0000\u0000\u00d0\u00d1\u0005\u001e\u0000\u0000\u00d1\u00d3\u0005"+
		"\u0001\u0000\u0000\u00d2\u00d4\u0003*\u0015\u0000\u00d3\u00d2\u0001\u0000"+
		"\u0000\u0000\u00d4\u00d5\u0001\u0000\u0000\u0000\u00d5\u00d3\u0001\u0000"+
		"\u0000\u0000\u00d5\u00d6\u0001\u0000\u0000\u0000\u00d6\u00d7\u0001\u0000"+
		"\u0000\u0000\u00d7\u00d8\u0005\u0002\u0000\u0000\u00d8#\u0001\u0000\u0000"+
		"\u0000\u00d9\u00da\u0003,\u0016\u0000\u00da\u00e3\u0005\u0004\u0000\u0000"+
		"\u00db\u00e0\u0003&\u0013\u0000\u00dc\u00dd\u0005\b\u0000\u0000\u00dd"+
		"\u00df\u0003&\u0013\u0000\u00de\u00dc\u0001\u0000\u0000\u0000\u00df\u00e2"+
		"\u0001\u0000\u0000\u0000\u00e0\u00de\u0001\u0000\u0000\u0000\u00e0\u00e1"+
		"\u0001\u0000\u0000\u0000\u00e1\u00e4\u0001\u0000\u0000\u0000\u00e2\u00e0"+
		"\u0001\u0000\u0000\u0000\u00e3\u00db\u0001\u0000\u0000\u0000\u00e3\u00e4"+
		"\u0001\u0000\u0000\u0000\u00e4\u00e5\u0001\u0000\u0000\u0000\u00e5\u00e6"+
		"\u0005\u0005\u0000\u0000\u00e6%\u0001\u0000\u0000\u0000\u00e7\u00ec\u0003"+
		",\u0016\u0000\u00e8\u00ec\u0005.\u0000\u0000\u00e9\u00ec\u0005/\u0000"+
		"\u0000\u00ea\u00ec\u00050\u0000\u0000\u00eb\u00e7\u0001\u0000\u0000\u0000"+
		"\u00eb\u00e8\u0001\u0000\u0000\u0000\u00eb\u00e9\u0001\u0000\u0000\u0000"+
		"\u00eb\u00ea\u0001\u0000\u0000\u0000\u00ec\'\u0001\u0000\u0000\u0000\u00ed"+
		"\u00ee\u0005-\u0000\u0000\u00ee\u00ef\u0005\u0003\u0000\u0000\u00ef\u00f2"+
		"\u0003>\u001f\u0000\u00f0\u00f1\u0005\t\u0000\u0000\u00f1\u00f3\u0003"+
		"@ \u0000\u00f2\u00f0\u0001\u0000\u0000\u0000\u00f2\u00f3\u0001\u0000\u0000"+
		"\u0000\u00f3\u00f7\u0001\u0000\u0000\u0000\u00f4\u00f6\u00030\u0018\u0000"+
		"\u00f5\u00f4\u0001\u0000\u0000\u0000\u00f6\u00f9\u0001\u0000\u0000\u0000"+
		"\u00f7\u00f5\u0001\u0000\u0000\u0000\u00f7\u00f8\u0001\u0000\u0000\u0000"+
		"\u00f8\u00fb\u0001\u0000\u0000\u0000\u00f9\u00f7\u0001\u0000\u0000\u0000"+
		"\u00fa\u00fc\u00052\u0000\u0000\u00fb\u00fa\u0001\u0000\u0000\u0000\u00fb"+
		"\u00fc\u0001\u0000\u0000\u0000\u00fc)\u0001\u0000\u0000\u0000\u00fd\u0101"+
		"\u0005.\u0000\u0000\u00fe\u0101\u00054\u0000\u0000\u00ff\u0101\u0003,"+
		"\u0016\u0000\u0100\u00fd\u0001\u0000\u0000\u0000\u0100\u00fe\u0001\u0000"+
		"\u0000\u0000\u0100\u00ff\u0001\u0000\u0000\u0000\u0101+\u0001\u0000\u0000"+
		"\u0000\u0102\u0107\u0007\u0001\u0000\u0000\u0103\u0104\u0005\n\u0000\u0000"+
		"\u0104\u0106\u0005-\u0000\u0000\u0105\u0103\u0001\u0000\u0000\u0000\u0106"+
		"\u0109\u0001\u0000\u0000\u0000\u0107\u0105\u0001\u0000\u0000\u0000\u0107"+
		"\u0108\u0001\u0000\u0000\u0000\u0108-\u0001\u0000\u0000\u0000\u0109\u0107"+
		"\u0001\u0000\u0000\u0000\u010a\u010b\u0005-\u0000\u0000\u010b\u010c\u0005"+
		"#\u0000\u0000\u010c\u010e\u0005\u0001\u0000\u0000\u010d\u010f\u0003(\u0014"+
		"\u0000\u010e\u010d\u0001\u0000\u0000\u0000\u010f\u0110\u0001\u0000\u0000"+
		"\u0000\u0110\u010e\u0001\u0000\u0000\u0000\u0110\u0111\u0001\u0000\u0000"+
		"\u0000\u0111\u0112\u0001\u0000\u0000\u0000\u0112\u0113\u0005\u0002\u0000"+
		"\u0000\u0113/\u0001\u0000\u0000\u0000\u0114\u0115\u0005\u000b\u0000\u0000"+
		"\u0115\u0116\u0005-\u0000\u0000\u0116\u0118\u0005\u0004\u0000\u0000\u0117"+
		"\u0119\u00032\u0019\u0000\u0118\u0117\u0001\u0000\u0000\u0000\u0118\u0119"+
		"\u0001\u0000\u0000\u0000\u0119\u011a\u0001\u0000\u0000\u0000\u011a\u011b"+
		"\u0005\u0005\u0000\u0000\u011b1\u0001\u0000\u0000\u0000\u011c\u0121\u0003"+
		"4\u001a\u0000\u011d\u011e\u0005\b\u0000\u0000\u011e\u0120\u00034\u001a"+
		"\u0000\u011f\u011d\u0001\u0000\u0000\u0000\u0120\u0123\u0001\u0000\u0000"+
		"\u0000\u0121\u011f\u0001\u0000\u0000\u0000\u0121\u0122\u0001\u0000\u0000"+
		"\u0000\u01223\u0001\u0000\u0000\u0000\u0123\u0121\u0001\u0000\u0000\u0000"+
		"\u0124\u0127\u0005.\u0000\u0000\u0125\u0127\u00036\u001b\u0000\u0126\u0124"+
		"\u0001\u0000\u0000\u0000\u0126\u0125\u0001\u0000\u0000\u0000\u01275\u0001"+
		"\u0000\u0000\u0000\u0128\u0131\u0005\f\u0000\u0000\u0129\u012e\u0005."+
		"\u0000\u0000\u012a\u012b\u0005\b\u0000\u0000\u012b\u012d\u0005.\u0000"+
		"\u0000\u012c\u012a\u0001\u0000\u0000\u0000\u012d\u0130\u0001\u0000\u0000"+
		"\u0000\u012e\u012c\u0001\u0000\u0000\u0000\u012e\u012f\u0001\u0000\u0000"+
		"\u0000\u012f\u0132\u0001\u0000\u0000\u0000\u0130\u012e\u0001\u0000\u0000"+
		"\u0000\u0131\u0129\u0001\u0000\u0000\u0000\u0131\u0132\u0001\u0000\u0000"+
		"\u0000\u0132\u0133\u0001\u0000\u0000\u0000\u0133\u0134\u0005\r\u0000\u0000"+
		"\u01347\u0001\u0000\u0000\u0000\u0135\u0139\u0005\u000e\u0000\u0000\u0136"+
		"\u0138\t\u0000\u0000\u0000\u0137\u0136\u0001\u0000\u0000\u0000\u0138\u013b"+
		"\u0001\u0000\u0000\u0000\u0139\u013a\u0001\u0000\u0000\u0000\u0139\u0137"+
		"\u0001\u0000\u0000\u0000\u013a\u013c\u0001\u0000\u0000\u0000\u013b\u0139"+
		"\u0001\u0000\u0000\u0000\u013c\u013d\u0005\u000f\u0000\u0000\u013d9\u0001"+
		"\u0000\u0000\u0000\u013e\u0142\u0005\u0010\u0000\u0000\u013f\u0141\t\u0000"+
		"\u0000\u0000\u0140\u013f\u0001\u0000\u0000\u0000\u0141\u0144\u0001\u0000"+
		"\u0000\u0000\u0142\u0143\u0001\u0000\u0000\u0000\u0142\u0140\u0001\u0000"+
		"\u0000\u0000\u0143\u0145\u0001\u0000\u0000\u0000\u0144\u0142\u0001\u0000"+
		"\u0000\u0000\u0145\u0146\u0005\u0011\u0000\u0000\u0146;\u0001\u0000\u0000"+
		"\u0000\u0147\u0149\u0007\u0002\u0000\u0000\u0148\u0147\u0001\u0000\u0000"+
		"\u0000\u0149\u014a\u0001\u0000\u0000\u0000\u014a\u0148\u0001\u0000\u0000"+
		"\u0000\u014a\u014b\u0001\u0000\u0000\u0000\u014b=\u0001\u0000\u0000\u0000"+
		"\u014c\u015c\u0005\u0017\u0000\u0000\u014d\u015c\u0005\u0018\u0000\u0000"+
		"\u014e\u015c\u0005\u0019\u0000\u0000\u014f\u0150\u0005\u0013\u0000\u0000"+
		"\u0150\u015c\u0003>\u001f\u0000\u0151\u0152\u0005#\u0000\u0000\u0152\u0156"+
		"\u0005\u0001\u0000\u0000\u0153\u0155\u0003(\u0014\u0000\u0154\u0153\u0001"+
		"\u0000\u0000\u0000\u0155\u0158\u0001\u0000\u0000\u0000\u0156\u0154\u0001"+
		"\u0000\u0000\u0000\u0156\u0157\u0001\u0000\u0000\u0000\u0157\u0159\u0001"+
		"\u0000\u0000\u0000\u0158\u0156\u0001\u0000\u0000\u0000\u0159\u015c\u0005"+
		"\u0002\u0000\u0000\u015a\u015c\u0005-\u0000\u0000\u015b\u014c\u0001\u0000"+
		"\u0000\u0000\u015b\u014d\u0001\u0000\u0000\u0000\u015b\u014e\u0001\u0000"+
		"\u0000\u0000\u015b\u014f\u0001\u0000\u0000\u0000\u015b\u0151\u0001\u0000"+
		"\u0000\u0000\u015b\u015a\u0001\u0000\u0000\u0000\u015c?\u0001\u0000\u0000"+
		"\u0000\u015d\u015e\u0007\u0003\u0000\u0000\u015eA\u0001\u0000\u0000\u0000"+
		"\u015f\u0160\u0007\u0004\u0000\u0000\u0160C\u0001\u0000\u0000\u0000&G"+
		"Q]dkq}\u0085\u0089\u008e\u0095\u009c\u00a3\u00ac\u00ba\u00c3\u00c7\u00ce"+
		"\u00d5\u00e0\u00e3\u00eb\u00f2\u00f7\u00fb\u0100\u0107\u0110\u0118\u0121"+
		"\u0126\u012e\u0131\u0139\u0142\u014a\u0156\u015b";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}