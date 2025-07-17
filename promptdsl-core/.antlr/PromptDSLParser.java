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
		PLUS=18, AFTER_BLOCK=19, FIX_BLOCK=20, STRING_TYPE=21, FLOAT_TYPE=22, 
		INT_TYPE=23, PROMPT=24, PARAMS=25, SYSTEM=26, USER=27, NOTE=28, INPUT=29, 
		OUTPUT=30, FORMAT=31, TYPE=32, STRUCT=33, BEFORE=34, SCHEMA=35, AFTER=36, 
		PARSE=37, JSONFIX=38, MARKDOWN=39, IF=40, ELSE=41, OUTPUTSPEC=42, ID=43, 
		STRING=44, NUMBER=45, BOOL=46, PIPE=47, SEMI=48, WS=49, LINE_COMMENT=50, 
		BLOCK_COMMENT=51;
	public static final int
		RULE_promptFile = 0, RULE_promptDef = 1, RULE_promptBlock = 2, RULE_inputSection = 3, 
		RULE_outputSection = 4, RULE_outputStruct = 5, RULE_outputMarkdown = 6, 
		RULE_beforeSection = 7, RULE_beforeContent = 8, RULE_varDef = 9, RULE_systemSection = 10, 
		RULE_sysContent = 11, RULE_moduleDef = 12, RULE_moduleContent = 13, RULE_userSection = 14, 
		RULE_userContent = 15, RULE_thencontent = 16, RULE_elsecontent = 17, RULE_ifStatement = 18, 
		RULE_condition = 19, RULE_noteSection = 20, RULE_dslCallExpr = 21, RULE_expr = 22, 
		RULE_fieldDef = 23, RULE_textLine = 24, RULE_paramPath = 25, RULE_structDef = 26, 
		RULE_annotation = 27, RULE_annotationArgs = 28, RULE_annotationValue = 29, 
		RULE_arrayLiteral = 30, RULE_afterSection = 31, RULE_fixSection = 32, 
		RULE_textBlock = 33, RULE_type = 34, RULE_value = 35, RULE_formatType = 36;
	private static String[] makeRuleNames() {
		return new String[] {
			"promptFile", "promptDef", "promptBlock", "inputSection", "outputSection", 
			"outputStruct", "outputMarkdown", "beforeSection", "beforeContent", "varDef", 
			"systemSection", "sysContent", "moduleDef", "moduleContent", "userSection", 
			"userContent", "thencontent", "elsecontent", "ifStatement", "condition", 
			"noteSection", "dslCallExpr", "expr", "fieldDef", "textLine", "paramPath", 
			"structDef", "annotation", "annotationArgs", "annotationValue", "arrayLiteral", 
			"afterSection", "fixSection", "textBlock", "type", "value", "formatType"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'{'", "'}'", "':'", "'='", "'('", "')'", "'=='", "'!='", "','", 
			"'.'", "'@'", "'['", "']'", "'-'", "'[]'", "'md'", "'json'", "'+'", null, 
			null, "'string'", "'float'", "'int'", "'prompt'", "'params'", "'system'", 
			"'user'", "'note'", "'input'", "'output'", "'format'", "'type'", "'struct'", 
			"'before'", "'schema'", "'after'", "'parse'", "'jsonfix'", "'markdown'", 
			"'if'", "'else'", "'outputspec'", null, null, null, null, "'|'", "';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, "PLUS", "AFTER_BLOCK", "FIX_BLOCK", 
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
			setState(75); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(74);
				promptDef();
				}
				}
				setState(77); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==PROMPT );
			setState(79);
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
			setState(81);
			match(PROMPT);
			setState(82);
			match(ID);
			setState(83);
			match(T__0);
			setState(85); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(84);
				promptBlock();
				}
				}
				setState(87); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 8798174969856L) != 0) );
			setState(89);
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
			setState(99);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INPUT:
				enterOuterAlt(_localctx, 1);
				{
				setState(91);
				inputSection();
				}
				break;
			case OUTPUT:
				enterOuterAlt(_localctx, 2);
				{
				setState(92);
				outputSection();
				}
				break;
			case SYSTEM:
				enterOuterAlt(_localctx, 3);
				{
				setState(93);
				systemSection();
				}
				break;
			case USER:
				enterOuterAlt(_localctx, 4);
				{
				setState(94);
				userSection();
				}
				break;
			case NOTE:
				enterOuterAlt(_localctx, 5);
				{
				setState(95);
				noteSection();
				}
				break;
			case AFTER_BLOCK:
				enterOuterAlt(_localctx, 6);
				{
				setState(96);
				afterSection();
				}
				break;
			case FIX_BLOCK:
				enterOuterAlt(_localctx, 7);
				{
				setState(97);
				fixSection();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 8);
				{
				setState(98);
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
			setState(101);
			match(INPUT);
			setState(102);
			match(T__0);
			setState(104); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(103);
				fieldDef();
				}
				}
				setState(106); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			setState(108);
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
			setState(110);
			match(OUTPUT);
			setState(113);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				{
				setState(111);
				outputStruct();
				}
				break;
			case T__2:
				{
				setState(112);
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
			setState(115);
			match(T__0);
			setState(117); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(116);
				fieldDef();
				}
				}
				setState(119); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			setState(121);
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
			setState(123);
			match(T__2);
			setState(124);
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
	public static class BeforeSectionContext extends ParserRuleContext {
		public TerminalNode BEFORE() { return getToken(PromptDSLParser.BEFORE, 0); }
		public List<BeforeContentContext> beforeContent() {
			return getRuleContexts(BeforeContentContext.class);
		}
		public BeforeContentContext beforeContent(int i) {
			return getRuleContext(BeforeContentContext.class,i);
		}
		public BeforeSectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_beforeSection; }
	}

	public final BeforeSectionContext beforeSection() throws RecognitionException {
		BeforeSectionContext _localctx = new BeforeSectionContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_beforeSection);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(126);
			match(BEFORE);
			setState(127);
			match(T__0);
			setState(131);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1263426370273280L) != 0)) {
				{
				{
				setState(128);
				beforeContent();
				}
				}
				setState(133);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(134);
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
	public static class BeforeContentContext extends ParserRuleContext {
		public VarDefContext varDef() {
			return getRuleContext(VarDefContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public TextLineContext textLine() {
			return getRuleContext(TextLineContext.class,0);
		}
		public BeforeContentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_beforeContent; }
	}

	public final BeforeContentContext beforeContent() throws RecognitionException {
		BeforeContentContext _localctx = new BeforeContentContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_beforeContent);
		try {
			setState(140);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(136);
				varDef();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(137);
				expr();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(138);
				ifStatement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(139);
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
	public static class VarDefContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(PromptDSLParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VarDefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varDef; }
	}

	public final VarDefContext varDef() throws RecognitionException {
		VarDefContext _localctx = new VarDefContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_varDef);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(142);
			match(ID);
			setState(143);
			match(T__3);
			setState(144);
			expr();
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
		enterRule(_localctx, 20, RULE_systemSection);
		int _la;
		try {
			setState(163);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(146);
				match(SYSTEM);
				setState(147);
				match(T__0);
				setState(149); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(148);
					match(ID);
					}
					}
					setState(151); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==ID );
				setState(153);
				match(T__1);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(154);
				match(SYSTEM);
				setState(155);
				match(T__0);
				setState(157); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(156);
					sysContent();
					}
					}
					setState(159); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 1263426370273280L) != 0) );
				setState(161);
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
		enterRule(_localctx, 22, RULE_sysContent);
		try {
			setState(168);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(165);
				ifStatement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(166);
				expr();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(167);
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
		enterRule(_localctx, 24, RULE_moduleDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(170);
			match(ID);
			setState(171);
			match(T__0);
			setState(175);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1156773742379008L) != 0)) {
				{
				{
				setState(172);
				moduleContent();
				}
				}
				setState(177);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(178);
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
		enterRule(_localctx, 26, RULE_moduleContent);
		try {
			setState(182);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(180);
				paramPath();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(181);
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
		enterRule(_localctx, 28, RULE_userSection);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			match(USER);
			setState(185);
			match(T__0);
			setState(187); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(186);
				userContent();
				}
				}
				setState(189); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 1263426370273280L) != 0) );
			setState(191);
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
		enterRule(_localctx, 30, RULE_userContent);
		try {
			setState(198);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(193);
				ifStatement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(194);
				paramPath();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(195);
				match(OUTPUTSPEC);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(196);
				expr();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(197);
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
		enterRule(_localctx, 32, RULE_thencontent);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(200);
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
		enterRule(_localctx, 34, RULE_elsecontent);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(202);
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
		enterRule(_localctx, 36, RULE_ifStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(204);
			match(IF);
			setState(205);
			match(T__4);
			setState(206);
			condition();
			setState(207);
			match(T__5);
			setState(208);
			match(T__0);
			setState(212);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1263426370273280L) != 0)) {
				{
				{
				setState(209);
				thencontent();
				}
				}
				setState(214);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(215);
			match(T__1);
			setState(225);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(216);
				match(ELSE);
				setState(217);
				match(T__0);
				setState(221);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1263426370273280L) != 0)) {
					{
					{
					setState(218);
					elsecontent();
					}
					}
					setState(223);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(224);
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
		enterRule(_localctx, 38, RULE_condition);
		int _la;
		try {
			setState(232);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(227);
				((ConditionContext)_localctx).lhs = expr();
				setState(228);
				((ConditionContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==T__6 || _la==T__7) ) {
					((ConditionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(229);
				((ConditionContext)_localctx).rhs = expr();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(231);
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
		enterRule(_localctx, 40, RULE_noteSection);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(234);
			match(NOTE);
			setState(235);
			match(T__0);
			setState(237); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(236);
				textLine();
				}
				}
				setState(239); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 1156773742379008L) != 0) );
			setState(241);
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
		enterRule(_localctx, 42, RULE_dslCallExpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(243);
			paramPath();
			setState(244);
			match(T__4);
			setState(253);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 136426951802880L) != 0)) {
				{
				setState(245);
				expr();
				setState(250);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__8) {
					{
					{
					setState(246);
					match(T__8);
					setState(247);
					expr();
					}
					}
					setState(252);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(255);
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
		enterRule(_localctx, 44, RULE_expr);
		try {
			setState(261);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INPUT:
			case OUTPUT:
			case BEFORE:
			case AFTER:
			case OUTPUTSPEC:
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				setState(257);
				paramPath();
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 2);
				{
				setState(258);
				match(STRING);
				}
				break;
			case NUMBER:
				enterOuterAlt(_localctx, 3);
				{
				setState(259);
				match(NUMBER);
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(260);
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
		enterRule(_localctx, 46, RULE_fieldDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(263);
			match(ID);
			setState(264);
			match(T__2);
			setState(265);
			type();
			setState(268);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(266);
				match(T__3);
				setState(267);
				value();
				}
			}

			setState(273);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__10) {
				{
				{
				setState(270);
				annotation();
				}
				}
				setState(275);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(277);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(276);
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
		enterRule(_localctx, 48, RULE_textLine);
		try {
			setState(282);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(279);
				match(STRING);
				}
				break;
			case LINE_COMMENT:
				enterOuterAlt(_localctx, 2);
				{
				setState(280);
				match(LINE_COMMENT);
				}
				break;
			case INPUT:
			case OUTPUT:
			case BEFORE:
			case AFTER:
			case OUTPUTSPEC:
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(281);
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
		public TerminalNode OUTPUTSPEC() { return getToken(PromptDSLParser.OUTPUTSPEC, 0); }
		public ParamPathContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramPath; }
	}

	public final ParamPathContext paramPath() throws RecognitionException {
		ParamPathContext _localctx = new ParamPathContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_paramPath);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(284);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 13281649491968L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(289);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__9) {
				{
				{
				setState(285);
				match(T__9);
				setState(286);
				match(ID);
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
		enterRule(_localctx, 52, RULE_structDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(292);
			match(ID);
			setState(293);
			match(STRUCT);
			setState(294);
			match(T__0);
			setState(296); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(295);
				fieldDef();
				}
				}
				setState(298); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			setState(300);
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
		enterRule(_localctx, 54, RULE_annotation);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(302);
			match(T__10);
			setState(303);
			match(ID);
			setState(304);
			match(T__4);
			setState(306);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11 || _la==STRING) {
				{
				setState(305);
				annotationArgs();
				}
			}

			setState(308);
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
		enterRule(_localctx, 56, RULE_annotationArgs);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(310);
			annotationValue();
			setState(315);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__8) {
				{
				{
				setState(311);
				match(T__8);
				setState(312);
				annotationValue();
				}
				}
				setState(317);
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
		enterRule(_localctx, 58, RULE_annotationValue);
		try {
			setState(320);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(318);
				match(STRING);
				}
				break;
			case T__11:
				enterOuterAlt(_localctx, 2);
				{
				setState(319);
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
		enterRule(_localctx, 60, RULE_arrayLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(322);
			match(T__11);
			setState(331);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING) {
				{
				setState(323);
				match(STRING);
				setState(328);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__8) {
					{
					{
					setState(324);
					match(T__8);
					setState(325);
					match(STRING);
					}
					}
					setState(330);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(333);
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
		public TerminalNode AFTER_BLOCK() { return getToken(PromptDSLParser.AFTER_BLOCK, 0); }
		public AfterSectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_afterSection; }
	}

	public final AfterSectionContext afterSection() throws RecognitionException {
		AfterSectionContext _localctx = new AfterSectionContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_afterSection);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(335);
			match(AFTER_BLOCK);
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
		public TerminalNode FIX_BLOCK() { return getToken(PromptDSLParser.FIX_BLOCK, 0); }
		public FixSectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fixSection; }
	}

	public final FixSectionContext fixSection() throws RecognitionException {
		FixSectionContext _localctx = new FixSectionContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_fixSection);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(337);
			match(FIX_BLOCK);
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
		enterRule(_localctx, 66, RULE_textBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(340); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(339);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 624522604593160L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				}
				setState(342); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 624522604593160L) != 0) );
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
		enterRule(_localctx, 68, RULE_type);
		int _la;
		try {
			setState(359);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING_TYPE:
				enterOuterAlt(_localctx, 1);
				{
				setState(344);
				match(STRING_TYPE);
				}
				break;
			case FLOAT_TYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(345);
				match(FLOAT_TYPE);
				}
				break;
			case INT_TYPE:
				enterOuterAlt(_localctx, 3);
				{
				setState(346);
				match(INT_TYPE);
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 4);
				{
				setState(347);
				match(T__14);
				setState(348);
				type();
				}
				break;
			case STRUCT:
				enterOuterAlt(_localctx, 5);
				{
				setState(349);
				match(STRUCT);
				setState(350);
				match(T__0);
				setState(354);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==ID) {
					{
					{
					setState(351);
					fieldDef();
					}
					}
					setState(356);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(357);
				match(T__1);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 6);
				{
				setState(358);
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
		enterRule(_localctx, 70, RULE_value);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(361);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 123145302310912L) != 0)) ) {
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
		enterRule(_localctx, 72, RULE_formatType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			_la = _input.LA(1);
			if ( !(_la==T__15 || _la==T__16) ) {
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
		"\u0004\u00013\u016e\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0001\u0000\u0004\u0000L\b\u0000\u000b\u0000\f"+
		"\u0000M\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0004\u0001V\b\u0001\u000b\u0001\f\u0001W\u0001\u0001\u0001\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0003\u0002d\b\u0002\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0004\u0003i\b\u0003\u000b\u0003\f\u0003j\u0001\u0003\u0001"+
		"\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004r\b\u0004\u0001"+
		"\u0005\u0001\u0005\u0004\u0005v\b\u0005\u000b\u0005\f\u0005w\u0001\u0005"+
		"\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0005\u0007\u0082\b\u0007\n\u0007\f\u0007\u0085\t\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b\u008d\b\b\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0004\n\u0096\b\n\u000b"+
		"\n\f\n\u0097\u0001\n\u0001\n\u0001\n\u0001\n\u0004\n\u009e\b\n\u000b\n"+
		"\f\n\u009f\u0001\n\u0001\n\u0003\n\u00a4\b\n\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0003\u000b\u00a9\b\u000b\u0001\f\u0001\f\u0001\f\u0005\f\u00ae"+
		"\b\f\n\f\f\f\u00b1\t\f\u0001\f\u0001\f\u0001\r\u0001\r\u0003\r\u00b7\b"+
		"\r\u0001\u000e\u0001\u000e\u0001\u000e\u0004\u000e\u00bc\b\u000e\u000b"+
		"\u000e\f\u000e\u00bd\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u00c7\b\u000f\u0001\u0010\u0001"+
		"\u0010\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0005\u0012\u00d3\b\u0012\n\u0012\f\u0012"+
		"\u00d6\t\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0005\u0012"+
		"\u00dc\b\u0012\n\u0012\f\u0012\u00df\t\u0012\u0001\u0012\u0003\u0012\u00e2"+
		"\b\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0003"+
		"\u0013\u00e9\b\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0004\u0014\u00ee"+
		"\b\u0014\u000b\u0014\f\u0014\u00ef\u0001\u0014\u0001\u0014\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0005\u0015\u00f9\b\u0015"+
		"\n\u0015\f\u0015\u00fc\t\u0015\u0003\u0015\u00fe\b\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u0106"+
		"\b\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0003"+
		"\u0017\u010d\b\u0017\u0001\u0017\u0005\u0017\u0110\b\u0017\n\u0017\f\u0017"+
		"\u0113\t\u0017\u0001\u0017\u0003\u0017\u0116\b\u0017\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0003\u0018\u011b\b\u0018\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0005\u0019\u0120\b\u0019\n\u0019\f\u0019\u0123\t\u0019\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0004\u001a\u0129\b\u001a\u000b\u001a"+
		"\f\u001a\u012a\u0001\u001a\u0001\u001a\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0003\u001b\u0133\b\u001b\u0001\u001b\u0001\u001b\u0001\u001c"+
		"\u0001\u001c\u0001\u001c\u0005\u001c\u013a\b\u001c\n\u001c\f\u001c\u013d"+
		"\t\u001c\u0001\u001d\u0001\u001d\u0003\u001d\u0141\b\u001d\u0001\u001e"+
		"\u0001\u001e\u0001\u001e\u0001\u001e\u0005\u001e\u0147\b\u001e\n\u001e"+
		"\f\u001e\u014a\t\u001e\u0003\u001e\u014c\b\u001e\u0001\u001e\u0001\u001e"+
		"\u0001\u001f\u0001\u001f\u0001 \u0001 \u0001!\u0004!\u0155\b!\u000b!\f"+
		"!\u0156\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0005\"\u0161\b\"\n\"\f\"\u0164\t\"\u0001\"\u0001\"\u0003\"\u0168\b"+
		"\"\u0001#\u0001#\u0001$\u0001$\u0001$\u0000\u0000%\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,."+
		"02468:<>@BDFH\u0000\u0005\u0001\u0000\u0007\b\u0004\u0000\u001d\u001e"+
		"\"\"$$*+\u0004\u0000\u0003\u0003\u000e\u000e+-11\u0001\u0000,.\u0001\u0000"+
		"\u0010\u0011\u0181\u0000K\u0001\u0000\u0000\u0000\u0002Q\u0001\u0000\u0000"+
		"\u0000\u0004c\u0001\u0000\u0000\u0000\u0006e\u0001\u0000\u0000\u0000\b"+
		"n\u0001\u0000\u0000\u0000\ns\u0001\u0000\u0000\u0000\f{\u0001\u0000\u0000"+
		"\u0000\u000e~\u0001\u0000\u0000\u0000\u0010\u008c\u0001\u0000\u0000\u0000"+
		"\u0012\u008e\u0001\u0000\u0000\u0000\u0014\u00a3\u0001\u0000\u0000\u0000"+
		"\u0016\u00a8\u0001\u0000\u0000\u0000\u0018\u00aa\u0001\u0000\u0000\u0000"+
		"\u001a\u00b6\u0001\u0000\u0000\u0000\u001c\u00b8\u0001\u0000\u0000\u0000"+
		"\u001e\u00c6\u0001\u0000\u0000\u0000 \u00c8\u0001\u0000\u0000\u0000\""+
		"\u00ca\u0001\u0000\u0000\u0000$\u00cc\u0001\u0000\u0000\u0000&\u00e8\u0001"+
		"\u0000\u0000\u0000(\u00ea\u0001\u0000\u0000\u0000*\u00f3\u0001\u0000\u0000"+
		"\u0000,\u0105\u0001\u0000\u0000\u0000.\u0107\u0001\u0000\u0000\u00000"+
		"\u011a\u0001\u0000\u0000\u00002\u011c\u0001\u0000\u0000\u00004\u0124\u0001"+
		"\u0000\u0000\u00006\u012e\u0001\u0000\u0000\u00008\u0136\u0001\u0000\u0000"+
		"\u0000:\u0140\u0001\u0000\u0000\u0000<\u0142\u0001\u0000\u0000\u0000>"+
		"\u014f\u0001\u0000\u0000\u0000@\u0151\u0001\u0000\u0000\u0000B\u0154\u0001"+
		"\u0000\u0000\u0000D\u0167\u0001\u0000\u0000\u0000F\u0169\u0001\u0000\u0000"+
		"\u0000H\u016b\u0001\u0000\u0000\u0000JL\u0003\u0002\u0001\u0000KJ\u0001"+
		"\u0000\u0000\u0000LM\u0001\u0000\u0000\u0000MK\u0001\u0000\u0000\u0000"+
		"MN\u0001\u0000\u0000\u0000NO\u0001\u0000\u0000\u0000OP\u0005\u0000\u0000"+
		"\u0001P\u0001\u0001\u0000\u0000\u0000QR\u0005\u0018\u0000\u0000RS\u0005"+
		"+\u0000\u0000SU\u0005\u0001\u0000\u0000TV\u0003\u0004\u0002\u0000UT\u0001"+
		"\u0000\u0000\u0000VW\u0001\u0000\u0000\u0000WU\u0001\u0000\u0000\u0000"+
		"WX\u0001\u0000\u0000\u0000XY\u0001\u0000\u0000\u0000YZ\u0005\u0002\u0000"+
		"\u0000Z\u0003\u0001\u0000\u0000\u0000[d\u0003\u0006\u0003\u0000\\d\u0003"+
		"\b\u0004\u0000]d\u0003\u0014\n\u0000^d\u0003\u001c\u000e\u0000_d\u0003"+
		"(\u0014\u0000`d\u0003>\u001f\u0000ad\u0003@ \u0000bd\u0003\u0018\f\u0000"+
		"c[\u0001\u0000\u0000\u0000c\\\u0001\u0000\u0000\u0000c]\u0001\u0000\u0000"+
		"\u0000c^\u0001\u0000\u0000\u0000c_\u0001\u0000\u0000\u0000c`\u0001\u0000"+
		"\u0000\u0000ca\u0001\u0000\u0000\u0000cb\u0001\u0000\u0000\u0000d\u0005"+
		"\u0001\u0000\u0000\u0000ef\u0005\u001d\u0000\u0000fh\u0005\u0001\u0000"+
		"\u0000gi\u0003.\u0017\u0000hg\u0001\u0000\u0000\u0000ij\u0001\u0000\u0000"+
		"\u0000jh\u0001\u0000\u0000\u0000jk\u0001\u0000\u0000\u0000kl\u0001\u0000"+
		"\u0000\u0000lm\u0005\u0002\u0000\u0000m\u0007\u0001\u0000\u0000\u0000"+
		"nq\u0005\u001e\u0000\u0000or\u0003\n\u0005\u0000pr\u0003\f\u0006\u0000"+
		"qo\u0001\u0000\u0000\u0000qp\u0001\u0000\u0000\u0000r\t\u0001\u0000\u0000"+
		"\u0000su\u0005\u0001\u0000\u0000tv\u0003.\u0017\u0000ut\u0001\u0000\u0000"+
		"\u0000vw\u0001\u0000\u0000\u0000wu\u0001\u0000\u0000\u0000wx\u0001\u0000"+
		"\u0000\u0000xy\u0001\u0000\u0000\u0000yz\u0005\u0002\u0000\u0000z\u000b"+
		"\u0001\u0000\u0000\u0000{|\u0005\u0003\u0000\u0000|}\u0005\'\u0000\u0000"+
		"}\r\u0001\u0000\u0000\u0000~\u007f\u0005\"\u0000\u0000\u007f\u0083\u0005"+
		"\u0001\u0000\u0000\u0080\u0082\u0003\u0010\b\u0000\u0081\u0080\u0001\u0000"+
		"\u0000\u0000\u0082\u0085\u0001\u0000\u0000\u0000\u0083\u0081\u0001\u0000"+
		"\u0000\u0000\u0083\u0084\u0001\u0000\u0000\u0000\u0084\u0086\u0001\u0000"+
		"\u0000\u0000\u0085\u0083\u0001\u0000\u0000\u0000\u0086\u0087\u0005\u0002"+
		"\u0000\u0000\u0087\u000f\u0001\u0000\u0000\u0000\u0088\u008d\u0003\u0012"+
		"\t\u0000\u0089\u008d\u0003,\u0016\u0000\u008a\u008d\u0003$\u0012\u0000"+
		"\u008b\u008d\u00030\u0018\u0000\u008c\u0088\u0001\u0000\u0000\u0000\u008c"+
		"\u0089\u0001\u0000\u0000\u0000\u008c\u008a\u0001\u0000\u0000\u0000\u008c"+
		"\u008b\u0001\u0000\u0000\u0000\u008d\u0011\u0001\u0000\u0000\u0000\u008e"+
		"\u008f\u0005+\u0000\u0000\u008f\u0090\u0005\u0004\u0000\u0000\u0090\u0091"+
		"\u0003,\u0016\u0000\u0091\u0013\u0001\u0000\u0000\u0000\u0092\u0093\u0005"+
		"\u001a\u0000\u0000\u0093\u0095\u0005\u0001\u0000\u0000\u0094\u0096\u0005"+
		"+\u0000\u0000\u0095\u0094\u0001\u0000\u0000\u0000\u0096\u0097\u0001\u0000"+
		"\u0000\u0000\u0097\u0095\u0001\u0000\u0000\u0000\u0097\u0098\u0001\u0000"+
		"\u0000\u0000\u0098\u0099\u0001\u0000\u0000\u0000\u0099\u00a4\u0005\u0002"+
		"\u0000\u0000\u009a\u009b\u0005\u001a\u0000\u0000\u009b\u009d\u0005\u0001"+
		"\u0000\u0000\u009c\u009e\u0003\u0016\u000b\u0000\u009d\u009c\u0001\u0000"+
		"\u0000\u0000\u009e\u009f\u0001\u0000\u0000\u0000\u009f\u009d\u0001\u0000"+
		"\u0000\u0000\u009f\u00a0\u0001\u0000\u0000\u0000\u00a0\u00a1\u0001\u0000"+
		"\u0000\u0000\u00a1\u00a2\u0005\u0002\u0000\u0000\u00a2\u00a4\u0001\u0000"+
		"\u0000\u0000\u00a3\u0092\u0001\u0000\u0000\u0000\u00a3\u009a\u0001\u0000"+
		"\u0000\u0000\u00a4\u0015\u0001\u0000\u0000\u0000\u00a5\u00a9\u0003$\u0012"+
		"\u0000\u00a6\u00a9\u0003,\u0016\u0000\u00a7\u00a9\u00030\u0018\u0000\u00a8"+
		"\u00a5\u0001\u0000\u0000\u0000\u00a8\u00a6\u0001\u0000\u0000\u0000\u00a8"+
		"\u00a7\u0001\u0000\u0000\u0000\u00a9\u0017\u0001\u0000\u0000\u0000\u00aa"+
		"\u00ab\u0005+\u0000\u0000\u00ab\u00af\u0005\u0001\u0000\u0000\u00ac\u00ae"+
		"\u0003\u001a\r\u0000\u00ad\u00ac\u0001\u0000\u0000\u0000\u00ae\u00b1\u0001"+
		"\u0000\u0000\u0000\u00af\u00ad\u0001\u0000\u0000\u0000\u00af\u00b0\u0001"+
		"\u0000\u0000\u0000\u00b0\u00b2\u0001\u0000\u0000\u0000\u00b1\u00af\u0001"+
		"\u0000\u0000\u0000\u00b2\u00b3\u0005\u0002\u0000\u0000\u00b3\u0019\u0001"+
		"\u0000\u0000\u0000\u00b4\u00b7\u00032\u0019\u0000\u00b5\u00b7\u00030\u0018"+
		"\u0000\u00b6\u00b4\u0001\u0000\u0000\u0000\u00b6\u00b5\u0001\u0000\u0000"+
		"\u0000\u00b7\u001b\u0001\u0000\u0000\u0000\u00b8\u00b9\u0005\u001b\u0000"+
		"\u0000\u00b9\u00bb\u0005\u0001\u0000\u0000\u00ba\u00bc\u0003\u001e\u000f"+
		"\u0000\u00bb\u00ba\u0001\u0000\u0000\u0000\u00bc\u00bd\u0001\u0000\u0000"+
		"\u0000\u00bd\u00bb\u0001\u0000\u0000\u0000\u00bd\u00be\u0001\u0000\u0000"+
		"\u0000\u00be\u00bf\u0001\u0000\u0000\u0000\u00bf\u00c0\u0005\u0002\u0000"+
		"\u0000\u00c0\u001d\u0001\u0000\u0000\u0000\u00c1\u00c7\u0003$\u0012\u0000"+
		"\u00c2\u00c7\u00032\u0019\u0000\u00c3\u00c7\u0005*\u0000\u0000\u00c4\u00c7"+
		"\u0003,\u0016\u0000\u00c5\u00c7\u00030\u0018\u0000\u00c6\u00c1\u0001\u0000"+
		"\u0000\u0000\u00c6\u00c2\u0001\u0000\u0000\u0000\u00c6\u00c3\u0001\u0000"+
		"\u0000\u0000\u00c6\u00c4\u0001\u0000\u0000\u0000\u00c6\u00c5\u0001\u0000"+
		"\u0000\u0000\u00c7\u001f\u0001\u0000\u0000\u0000\u00c8\u00c9\u0003\u001e"+
		"\u000f\u0000\u00c9!\u0001\u0000\u0000\u0000\u00ca\u00cb\u0003\u001e\u000f"+
		"\u0000\u00cb#\u0001\u0000\u0000\u0000\u00cc\u00cd\u0005(\u0000\u0000\u00cd"+
		"\u00ce\u0005\u0005\u0000\u0000\u00ce\u00cf\u0003&\u0013\u0000\u00cf\u00d0"+
		"\u0005\u0006\u0000\u0000\u00d0\u00d4\u0005\u0001\u0000\u0000\u00d1\u00d3"+
		"\u0003 \u0010\u0000\u00d2\u00d1\u0001\u0000\u0000\u0000\u00d3\u00d6\u0001"+
		"\u0000\u0000\u0000\u00d4\u00d2\u0001\u0000\u0000\u0000\u00d4\u00d5\u0001"+
		"\u0000\u0000\u0000\u00d5\u00d7\u0001\u0000\u0000\u0000\u00d6\u00d4\u0001"+
		"\u0000\u0000\u0000\u00d7\u00e1\u0005\u0002\u0000\u0000\u00d8\u00d9\u0005"+
		")\u0000\u0000\u00d9\u00dd\u0005\u0001\u0000\u0000\u00da\u00dc\u0003\""+
		"\u0011\u0000\u00db\u00da\u0001\u0000\u0000\u0000\u00dc\u00df\u0001\u0000"+
		"\u0000\u0000\u00dd\u00db\u0001\u0000\u0000\u0000\u00dd\u00de\u0001\u0000"+
		"\u0000\u0000\u00de\u00e0\u0001\u0000\u0000\u0000\u00df\u00dd\u0001\u0000"+
		"\u0000\u0000\u00e0\u00e2\u0005\u0002\u0000\u0000\u00e1\u00d8\u0001\u0000"+
		"\u0000\u0000\u00e1\u00e2\u0001\u0000\u0000\u0000\u00e2%\u0001\u0000\u0000"+
		"\u0000\u00e3\u00e4\u0003,\u0016\u0000\u00e4\u00e5\u0007\u0000\u0000\u0000"+
		"\u00e5\u00e6\u0003,\u0016\u0000\u00e6\u00e9\u0001\u0000\u0000\u0000\u00e7"+
		"\u00e9\u0003,\u0016\u0000\u00e8\u00e3\u0001\u0000\u0000\u0000\u00e8\u00e7"+
		"\u0001\u0000\u0000\u0000\u00e9\'\u0001\u0000\u0000\u0000\u00ea\u00eb\u0005"+
		"\u001c\u0000\u0000\u00eb\u00ed\u0005\u0001\u0000\u0000\u00ec\u00ee\u0003"+
		"0\u0018\u0000\u00ed\u00ec\u0001\u0000\u0000\u0000\u00ee\u00ef\u0001\u0000"+
		"\u0000\u0000\u00ef\u00ed\u0001\u0000\u0000\u0000\u00ef\u00f0\u0001\u0000"+
		"\u0000\u0000\u00f0\u00f1\u0001\u0000\u0000\u0000\u00f1\u00f2\u0005\u0002"+
		"\u0000\u0000\u00f2)\u0001\u0000\u0000\u0000\u00f3\u00f4\u00032\u0019\u0000"+
		"\u00f4\u00fd\u0005\u0005\u0000\u0000\u00f5\u00fa\u0003,\u0016\u0000\u00f6"+
		"\u00f7\u0005\t\u0000\u0000\u00f7\u00f9\u0003,\u0016\u0000\u00f8\u00f6"+
		"\u0001\u0000\u0000\u0000\u00f9\u00fc\u0001\u0000\u0000\u0000\u00fa\u00f8"+
		"\u0001\u0000\u0000\u0000\u00fa\u00fb\u0001\u0000\u0000\u0000\u00fb\u00fe"+
		"\u0001\u0000\u0000\u0000\u00fc\u00fa\u0001\u0000\u0000\u0000\u00fd\u00f5"+
		"\u0001\u0000\u0000\u0000\u00fd\u00fe\u0001\u0000\u0000\u0000\u00fe\u00ff"+
		"\u0001\u0000\u0000\u0000\u00ff\u0100\u0005\u0006\u0000\u0000\u0100+\u0001"+
		"\u0000\u0000\u0000\u0101\u0106\u00032\u0019\u0000\u0102\u0106\u0005,\u0000"+
		"\u0000\u0103\u0106\u0005-\u0000\u0000\u0104\u0106\u0005.\u0000\u0000\u0105"+
		"\u0101\u0001\u0000\u0000\u0000\u0105\u0102\u0001\u0000\u0000\u0000\u0105"+
		"\u0103\u0001\u0000\u0000\u0000\u0105\u0104\u0001\u0000\u0000\u0000\u0106"+
		"-\u0001\u0000\u0000\u0000\u0107\u0108\u0005+\u0000\u0000\u0108\u0109\u0005"+
		"\u0003\u0000\u0000\u0109\u010c\u0003D\"\u0000\u010a\u010b\u0005\u0004"+
		"\u0000\u0000\u010b\u010d\u0003F#\u0000\u010c\u010a\u0001\u0000\u0000\u0000"+
		"\u010c\u010d\u0001\u0000\u0000\u0000\u010d\u0111\u0001\u0000\u0000\u0000"+
		"\u010e\u0110\u00036\u001b\u0000\u010f\u010e\u0001\u0000\u0000\u0000\u0110"+
		"\u0113\u0001\u0000\u0000\u0000\u0111\u010f\u0001\u0000\u0000\u0000\u0111"+
		"\u0112\u0001\u0000\u0000\u0000\u0112\u0115\u0001\u0000\u0000\u0000\u0113"+
		"\u0111\u0001\u0000\u0000\u0000\u0114\u0116\u00050\u0000\u0000\u0115\u0114"+
		"\u0001\u0000\u0000\u0000\u0115\u0116\u0001\u0000\u0000\u0000\u0116/\u0001"+
		"\u0000\u0000\u0000\u0117\u011b\u0005,\u0000\u0000\u0118\u011b\u00052\u0000"+
		"\u0000\u0119\u011b\u00032\u0019\u0000\u011a\u0117\u0001\u0000\u0000\u0000"+
		"\u011a\u0118\u0001\u0000\u0000\u0000\u011a\u0119\u0001\u0000\u0000\u0000"+
		"\u011b1\u0001\u0000\u0000\u0000\u011c\u0121\u0007\u0001\u0000\u0000\u011d"+
		"\u011e\u0005\n\u0000\u0000\u011e\u0120\u0005+\u0000\u0000\u011f\u011d"+
		"\u0001\u0000\u0000\u0000\u0120\u0123\u0001\u0000\u0000\u0000\u0121\u011f"+
		"\u0001\u0000\u0000\u0000\u0121\u0122\u0001\u0000\u0000\u0000\u01223\u0001"+
		"\u0000\u0000\u0000\u0123\u0121\u0001\u0000\u0000\u0000\u0124\u0125\u0005"+
		"+\u0000\u0000\u0125\u0126\u0005!\u0000\u0000\u0126\u0128\u0005\u0001\u0000"+
		"\u0000\u0127\u0129\u0003.\u0017\u0000\u0128\u0127\u0001\u0000\u0000\u0000"+
		"\u0129\u012a\u0001\u0000\u0000\u0000\u012a\u0128\u0001\u0000\u0000\u0000"+
		"\u012a\u012b\u0001\u0000\u0000\u0000\u012b\u012c\u0001\u0000\u0000\u0000"+
		"\u012c\u012d\u0005\u0002\u0000\u0000\u012d5\u0001\u0000\u0000\u0000\u012e"+
		"\u012f\u0005\u000b\u0000\u0000\u012f\u0130\u0005+\u0000\u0000\u0130\u0132"+
		"\u0005\u0005\u0000\u0000\u0131\u0133\u00038\u001c\u0000\u0132\u0131\u0001"+
		"\u0000\u0000\u0000\u0132\u0133\u0001\u0000\u0000\u0000\u0133\u0134\u0001"+
		"\u0000\u0000\u0000\u0134\u0135\u0005\u0006\u0000\u0000\u01357\u0001\u0000"+
		"\u0000\u0000\u0136\u013b\u0003:\u001d\u0000\u0137\u0138\u0005\t\u0000"+
		"\u0000\u0138\u013a\u0003:\u001d\u0000\u0139\u0137\u0001\u0000\u0000\u0000"+
		"\u013a\u013d\u0001\u0000\u0000\u0000\u013b\u0139\u0001\u0000\u0000\u0000"+
		"\u013b\u013c\u0001\u0000\u0000\u0000\u013c9\u0001\u0000\u0000\u0000\u013d"+
		"\u013b\u0001\u0000\u0000\u0000\u013e\u0141\u0005,\u0000\u0000\u013f\u0141"+
		"\u0003<\u001e\u0000\u0140\u013e\u0001\u0000\u0000\u0000\u0140\u013f\u0001"+
		"\u0000\u0000\u0000\u0141;\u0001\u0000\u0000\u0000\u0142\u014b\u0005\f"+
		"\u0000\u0000\u0143\u0148\u0005,\u0000\u0000\u0144\u0145\u0005\t\u0000"+
		"\u0000\u0145\u0147\u0005,\u0000\u0000\u0146\u0144\u0001\u0000\u0000\u0000"+
		"\u0147\u014a\u0001\u0000\u0000\u0000\u0148\u0146\u0001\u0000\u0000\u0000"+
		"\u0148\u0149\u0001\u0000\u0000\u0000\u0149\u014c\u0001\u0000\u0000\u0000"+
		"\u014a\u0148\u0001\u0000\u0000\u0000\u014b\u0143\u0001\u0000\u0000\u0000"+
		"\u014b\u014c\u0001\u0000\u0000\u0000\u014c\u014d\u0001\u0000\u0000\u0000"+
		"\u014d\u014e\u0005\r\u0000\u0000\u014e=\u0001\u0000\u0000\u0000\u014f"+
		"\u0150\u0005\u0013\u0000\u0000\u0150?\u0001\u0000\u0000\u0000\u0151\u0152"+
		"\u0005\u0014\u0000\u0000\u0152A\u0001\u0000\u0000\u0000\u0153\u0155\u0007"+
		"\u0002\u0000\u0000\u0154\u0153\u0001\u0000\u0000\u0000\u0155\u0156\u0001"+
		"\u0000\u0000\u0000\u0156\u0154\u0001\u0000\u0000\u0000\u0156\u0157\u0001"+
		"\u0000\u0000\u0000\u0157C\u0001\u0000\u0000\u0000\u0158\u0168\u0005\u0015"+
		"\u0000\u0000\u0159\u0168\u0005\u0016\u0000\u0000\u015a\u0168\u0005\u0017"+
		"\u0000\u0000\u015b\u015c\u0005\u000f\u0000\u0000\u015c\u0168\u0003D\""+
		"\u0000\u015d\u015e\u0005!\u0000\u0000\u015e\u0162\u0005\u0001\u0000\u0000"+
		"\u015f\u0161\u0003.\u0017\u0000\u0160\u015f\u0001\u0000\u0000\u0000\u0161"+
		"\u0164\u0001\u0000\u0000\u0000\u0162\u0160\u0001\u0000\u0000\u0000\u0162"+
		"\u0163\u0001\u0000\u0000\u0000\u0163\u0165\u0001\u0000\u0000\u0000\u0164"+
		"\u0162\u0001\u0000\u0000\u0000\u0165\u0168\u0005\u0002\u0000\u0000\u0166"+
		"\u0168\u0005+\u0000\u0000\u0167\u0158\u0001\u0000\u0000\u0000\u0167\u0159"+
		"\u0001\u0000\u0000\u0000\u0167\u015a\u0001\u0000\u0000\u0000\u0167\u015b"+
		"\u0001\u0000\u0000\u0000\u0167\u015d\u0001\u0000\u0000\u0000\u0167\u0166"+
		"\u0001\u0000\u0000\u0000\u0168E\u0001\u0000\u0000\u0000\u0169\u016a\u0007"+
		"\u0003\u0000\u0000\u016aG\u0001\u0000\u0000\u0000\u016b\u016c\u0007\u0004"+
		"\u0000\u0000\u016cI\u0001\u0000\u0000\u0000&MWcjqw\u0083\u008c\u0097\u009f"+
		"\u00a3\u00a8\u00af\u00b6\u00bd\u00c6\u00d4\u00dd\u00e1\u00e8\u00ef\u00fa"+
		"\u00fd\u0105\u010c\u0111\u0115\u011a\u0121\u012a\u0132\u013b\u0140\u0148"+
		"\u014b\u0156\u0162\u0167";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}