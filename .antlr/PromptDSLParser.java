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
		T__17=18, T__18=19, T__19=20, PROMPT=21, PARAMS=22, SYSTEM=23, USER=24, 
		NOTE=25, IN=26, OUTPUT=27, FORMAT=28, TYPE=29, STRUCT=30, SCHEMA=31, ID=32, 
		STRING=33, NUMBER=34, BOOL=35, PIPE=36, SEMI=37, WS=38, LINE_COMMENT=39, 
		BLOCK_COMMENT=40;
	public static final int
		RULE_promptFile = 0, RULE_promptDef = 1, RULE_promptBlock = 2, RULE_paramBody = 3, 
		RULE_paramSection = 4, RULE_outputBody = 5, RULE_fieldDef = 6, RULE_structDef = 7, 
		RULE_annotation = 8, RULE_annotationArgs = 9, RULE_annotationValue = 10, 
		RULE_arrayLiteral = 11, RULE_userBlock = 12, RULE_classificationBlock = 13, 
		RULE_summarizationBlock = 14, RULE_compilationBlock = 15, RULE_kvPair = 16, 
		RULE_textBlock = 17, RULE_type = 18, RULE_value = 19, RULE_formatType = 20;
	private static String[] makeRuleNames() {
		return new String[] {
			"promptFile", "promptDef", "promptBlock", "paramBody", "paramSection", 
			"outputBody", "fieldDef", "structDef", "annotation", "annotationArgs", 
			"annotationValue", "arrayLiteral", "userBlock", "classificationBlock", 
			"summarizationBlock", "compilationBlock", "kvPair", "textBlock", "type", 
			"value", "formatType"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'{'", "'}'", "':'", "'['", "']'", "'='", "'@'", "'('", "')'", 
			"','", "'classification'", "'summarization'", "'compilation'", "'extra_hint'", 
			"'-'", "'string'", "'float'", "'int'", "'md'", "'json'", "'prompt'", 
			"'params'", "'system'", "'user'", "'note'", "'in'", "'output'", "'format'", 
			"'type'", "'struct'", "'schema'", null, null, null, null, "'|'", "';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, "PROMPT", "PARAMS", 
			"SYSTEM", "USER", "NOTE", "IN", "OUTPUT", "FORMAT", "TYPE", "STRUCT", 
			"SCHEMA", "ID", "STRING", "NUMBER", "BOOL", "PIPE", "SEMI", "WS", "LINE_COMMENT", 
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
			setState(43); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(42);
				promptDef();
				}
				}
				setState(45); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==PROMPT );
			setState(47);
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
			setState(49);
			match(PROMPT);
			setState(50);
			match(ID);
			setState(51);
			match(T__0);
			setState(53); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(52);
				promptBlock();
				}
				}
				setState(55); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 62914560L) != 0) );
			setState(57);
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
		public TerminalNode PARAMS() { return getToken(PromptDSLParser.PARAMS, 0); }
		public ParamBodyContext paramBody() {
			return getRuleContext(ParamBodyContext.class,0);
		}
		public TerminalNode SYSTEM() { return getToken(PromptDSLParser.SYSTEM, 0); }
		public TextBlockContext textBlock() {
			return getRuleContext(TextBlockContext.class,0);
		}
		public TerminalNode USER() { return getToken(PromptDSLParser.USER, 0); }
		public UserBlockContext userBlock() {
			return getRuleContext(UserBlockContext.class,0);
		}
		public TerminalNode NOTE() { return getToken(PromptDSLParser.NOTE, 0); }
		public PromptBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_promptBlock; }
	}

	public final PromptBlockContext promptBlock() throws RecognitionException {
		PromptBlockContext _localctx = new PromptBlockContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_promptBlock);
		try {
			setState(79);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PARAMS:
				enterOuterAlt(_localctx, 1);
				{
				setState(59);
				match(PARAMS);
				setState(60);
				match(T__0);
				setState(61);
				paramBody();
				setState(62);
				match(T__1);
				}
				break;
			case SYSTEM:
				enterOuterAlt(_localctx, 2);
				{
				setState(64);
				match(SYSTEM);
				setState(65);
				match(T__0);
				setState(66);
				textBlock();
				setState(67);
				match(T__1);
				}
				break;
			case USER:
				enterOuterAlt(_localctx, 3);
				{
				setState(69);
				match(USER);
				setState(70);
				match(T__0);
				setState(71);
				userBlock();
				setState(72);
				match(T__1);
				}
				break;
			case NOTE:
				enterOuterAlt(_localctx, 4);
				{
				setState(74);
				match(NOTE);
				setState(75);
				match(T__0);
				setState(76);
				textBlock();
				setState(77);
				match(T__1);
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
	public static class ParamBodyContext extends ParserRuleContext {
		public List<ParamSectionContext> paramSection() {
			return getRuleContexts(ParamSectionContext.class);
		}
		public ParamSectionContext paramSection(int i) {
			return getRuleContext(ParamSectionContext.class,i);
		}
		public ParamBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramBody; }
	}

	public final ParamBodyContext paramBody() throws RecognitionException {
		ParamBodyContext _localctx = new ParamBodyContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_paramBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(82); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(81);
				paramSection();
				}
				}
				setState(84); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 4496293888L) != 0) );
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
	public static class ParamSectionContext extends ParserRuleContext {
		public ParamSectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramSection; }
	 
		public ParamSectionContext() { }
		public void copyFrom(ParamSectionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SingleFieldContext extends ParamSectionContext {
		public FieldDefContext fieldDef() {
			return getRuleContext(FieldDefContext.class,0);
		}
		public SingleFieldContext(ParamSectionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class InputSectionContext extends ParamSectionContext {
		public TerminalNode IN() { return getToken(PromptDSLParser.IN, 0); }
		public List<FieldDefContext> fieldDef() {
			return getRuleContexts(FieldDefContext.class);
		}
		public FieldDefContext fieldDef(int i) {
			return getRuleContext(FieldDefContext.class,i);
		}
		public InputSectionContext(ParamSectionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OutputSectionContext extends ParamSectionContext {
		public TerminalNode OUTPUT() { return getToken(PromptDSLParser.OUTPUT, 0); }
		public OutputBodyContext outputBody() {
			return getRuleContext(OutputBodyContext.class,0);
		}
		public OutputSectionContext(ParamSectionContext ctx) { copyFrom(ctx); }
	}

	public final ParamSectionContext paramSection() throws RecognitionException {
		ParamSectionContext _localctx = new ParamSectionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_paramSection);
		int _la;
		try {
			setState(103);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IN:
				_localctx = new InputSectionContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(86);
				match(IN);
				setState(87);
				match(T__2);
				setState(88);
				match(T__0);
				setState(90); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(89);
					fieldDef();
					}
					}
					setState(92); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==ID );
				setState(94);
				match(T__1);
				}
				break;
			case OUTPUT:
				_localctx = new OutputSectionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(96);
				match(OUTPUT);
				setState(97);
				match(T__2);
				setState(98);
				match(T__0);
				setState(99);
				outputBody();
				setState(100);
				match(T__1);
				}
				break;
			case ID:
				_localctx = new SingleFieldContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(102);
				fieldDef();
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
	public static class OutputBodyContext extends ParserRuleContext {
		public TerminalNode FORMAT() { return getToken(PromptDSLParser.FORMAT, 0); }
		public FormatTypeContext formatType() {
			return getRuleContext(FormatTypeContext.class,0);
		}
		public TerminalNode TYPE() { return getToken(PromptDSLParser.TYPE, 0); }
		public StructDefContext structDef() {
			return getRuleContext(StructDefContext.class,0);
		}
		public TerminalNode SCHEMA() { return getToken(PromptDSLParser.SCHEMA, 0); }
		public TerminalNode ID() { return getToken(PromptDSLParser.ID, 0); }
		public OutputBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_outputBody; }
	}

	public final OutputBodyContext outputBody() throws RecognitionException {
		OutputBodyContext _localctx = new OutputBodyContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_outputBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FORMAT) {
				{
				setState(105);
				match(FORMAT);
				setState(106);
				match(T__2);
				setState(107);
				formatType();
				}
			}

			setState(112);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==TYPE) {
				{
				setState(110);
				match(TYPE);
				setState(111);
				structDef();
				}
			}

			setState(119);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SCHEMA) {
				{
				setState(114);
				match(SCHEMA);
				setState(115);
				match(T__2);
				setState(116);
				match(T__3);
				setState(117);
				match(ID);
				setState(118);
				match(T__4);
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
		enterRule(_localctx, 12, RULE_fieldDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(121);
			match(ID);
			setState(122);
			match(T__2);
			setState(123);
			type();
			setState(126);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__5) {
				{
				setState(124);
				match(T__5);
				setState(125);
				value();
				}
			}

			setState(131);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__6) {
				{
				{
				setState(128);
				annotation();
				}
				}
				setState(133);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(135);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMI) {
				{
				setState(134);
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
		enterRule(_localctx, 14, RULE_structDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(137);
			match(ID);
			setState(138);
			match(STRUCT);
			setState(139);
			match(T__0);
			setState(141); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(140);
				fieldDef();
				}
				}
				setState(143); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			setState(145);
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
		enterRule(_localctx, 16, RULE_annotation);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
			match(T__6);
			setState(148);
			match(ID);
			setState(149);
			match(T__7);
			setState(151);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3 || _la==STRING) {
				{
				setState(150);
				annotationArgs();
				}
			}

			setState(153);
			match(T__8);
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
		enterRule(_localctx, 18, RULE_annotationArgs);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(155);
			annotationValue();
			setState(160);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__9) {
				{
				{
				setState(156);
				match(T__9);
				setState(157);
				annotationValue();
				}
				}
				setState(162);
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
		enterRule(_localctx, 20, RULE_annotationValue);
		try {
			setState(165);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(163);
				match(STRING);
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 2);
				{
				setState(164);
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
		enterRule(_localctx, 22, RULE_arrayLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(167);
			match(T__3);
			setState(176);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==STRING) {
				{
				setState(168);
				match(STRING);
				setState(173);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__9) {
					{
					{
					setState(169);
					match(T__9);
					setState(170);
					match(STRING);
					}
					}
					setState(175);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(178);
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
	public static class UserBlockContext extends ParserRuleContext {
		public List<TextBlockContext> textBlock() {
			return getRuleContexts(TextBlockContext.class);
		}
		public TextBlockContext textBlock(int i) {
			return getRuleContext(TextBlockContext.class,i);
		}
		public List<ClassificationBlockContext> classificationBlock() {
			return getRuleContexts(ClassificationBlockContext.class);
		}
		public ClassificationBlockContext classificationBlock(int i) {
			return getRuleContext(ClassificationBlockContext.class,i);
		}
		public List<SummarizationBlockContext> summarizationBlock() {
			return getRuleContexts(SummarizationBlockContext.class);
		}
		public SummarizationBlockContext summarizationBlock(int i) {
			return getRuleContext(SummarizationBlockContext.class,i);
		}
		public List<CompilationBlockContext> compilationBlock() {
			return getRuleContexts(CompilationBlockContext.class);
		}
		public CompilationBlockContext compilationBlock(int i) {
			return getRuleContext(CompilationBlockContext.class,i);
		}
		public UserBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_userBlock; }
	}

	public final UserBlockContext userBlock() throws RecognitionException {
		UserBlockContext _localctx = new UserBlockContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_userBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(184); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				setState(184);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case T__2:
				case T__14:
				case ID:
				case STRING:
				case NUMBER:
				case WS:
					{
					setState(180);
					textBlock();
					}
					break;
				case T__10:
					{
					setState(181);
					classificationBlock();
					}
					break;
				case T__11:
					{
					setState(182);
					summarizationBlock();
					}
					break;
				case T__12:
					{
					setState(183);
					compilationBlock();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(186); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 304942725128L) != 0) );
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
	public static class ClassificationBlockContext extends ParserRuleContext {
		public List<KvPairContext> kvPair() {
			return getRuleContexts(KvPairContext.class);
		}
		public KvPairContext kvPair(int i) {
			return getRuleContext(KvPairContext.class,i);
		}
		public ClassificationBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classificationBlock; }
	}

	public final ClassificationBlockContext classificationBlock() throws RecognitionException {
		ClassificationBlockContext _localctx = new ClassificationBlockContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_classificationBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(188);
			match(T__10);
			setState(189);
			match(T__0);
			setState(193);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ID) {
				{
				{
				setState(190);
				kvPair();
				}
				}
				setState(195);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(196);
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
	public static class SummarizationBlockContext extends ParserRuleContext {
		public List<KvPairContext> kvPair() {
			return getRuleContexts(KvPairContext.class);
		}
		public KvPairContext kvPair(int i) {
			return getRuleContext(KvPairContext.class,i);
		}
		public SummarizationBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_summarizationBlock; }
	}

	public final SummarizationBlockContext summarizationBlock() throws RecognitionException {
		SummarizationBlockContext _localctx = new SummarizationBlockContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_summarizationBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(198);
			match(T__11);
			setState(199);
			match(T__0);
			setState(203);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ID) {
				{
				{
				setState(200);
				kvPair();
				}
				}
				setState(205);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(206);
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
	public static class CompilationBlockContext extends ParserRuleContext {
		public List<KvPairContext> kvPair() {
			return getRuleContexts(KvPairContext.class);
		}
		public KvPairContext kvPair(int i) {
			return getRuleContext(KvPairContext.class,i);
		}
		public CompilationBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_compilationBlock; }
	}

	public final CompilationBlockContext compilationBlock() throws RecognitionException {
		CompilationBlockContext _localctx = new CompilationBlockContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_compilationBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(208);
			match(T__12);
			setState(209);
			match(T__0);
			setState(213);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ID) {
				{
				{
				setState(210);
				kvPair();
				}
				}
				setState(215);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(216);
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
	public static class KvPairContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(PromptDSLParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(PromptDSLParser.ID, i);
		}
		public TerminalNode STRING() { return getToken(PromptDSLParser.STRING, 0); }
		public KvPairContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_kvPair; }
	}

	public final KvPairContext kvPair() throws RecognitionException {
		KvPairContext _localctx = new KvPairContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_kvPair);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(218);
			match(ID);
			setState(219);
			match(T__2);
			setState(220);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 12884918272L) != 0)) ) {
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
		enterRule(_localctx, 34, RULE_textBlock);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(223); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(222);
					_la = _input.LA(1);
					if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 304942710792L) != 0)) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(225); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
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
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode ID() { return getToken(PromptDSLParser.ID, 0); }
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_type);
		try {
			setState(235);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__15:
				enterOuterAlt(_localctx, 1);
				{
				setState(227);
				match(T__15);
				}
				break;
			case T__16:
				enterOuterAlt(_localctx, 2);
				{
				setState(228);
				match(T__16);
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 3);
				{
				setState(229);
				match(T__17);
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 4);
				{
				setState(230);
				match(T__3);
				setState(231);
				type();
				setState(232);
				match(T__4);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 5);
				{
				setState(234);
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
		enterRule(_localctx, 38, RULE_value);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(237);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 60129542144L) != 0)) ) {
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
		enterRule(_localctx, 40, RULE_formatType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			_la = _input.LA(1);
			if ( !(_la==T__18 || _la==T__19) ) {
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
		"\u0004\u0001(\u00f2\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0001\u0000\u0004\u0000"+
		",\b\u0000\u000b\u0000\f\u0000-\u0001\u0000\u0001\u0000\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0004\u00016\b\u0001\u000b\u0001\f\u0001"+
		"7\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002P\b\u0002"+
		"\u0001\u0003\u0004\u0003S\b\u0003\u000b\u0003\f\u0003T\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0004\u0004[\b\u0004\u000b\u0004\f\u0004"+
		"\\\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004h\b\u0004\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0003\u0005m\b\u0005\u0001\u0005\u0001\u0005"+
		"\u0003\u0005q\b\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0003\u0005x\b\u0005\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0003\u0006\u007f\b\u0006\u0001\u0006\u0005\u0006"+
		"\u0082\b\u0006\n\u0006\f\u0006\u0085\t\u0006\u0001\u0006\u0003\u0006\u0088"+
		"\b\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0004\u0007\u008e"+
		"\b\u0007\u000b\u0007\f\u0007\u008f\u0001\u0007\u0001\u0007\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0003\b\u0098\b\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001"+
		"\t\u0005\t\u009f\b\t\n\t\f\t\u00a2\t\t\u0001\n\u0001\n\u0003\n\u00a6\b"+
		"\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0005\u000b\u00ac\b"+
		"\u000b\n\u000b\f\u000b\u00af\t\u000b\u0003\u000b\u00b1\b\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0004\f\u00b9\b\f\u000b\f"+
		"\f\f\u00ba\u0001\r\u0001\r\u0001\r\u0005\r\u00c0\b\r\n\r\f\r\u00c3\t\r"+
		"\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0005\u000e\u00ca"+
		"\b\u000e\n\u000e\f\u000e\u00cd\t\u000e\u0001\u000e\u0001\u000e\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0005\u000f\u00d4\b\u000f\n\u000f\f\u000f\u00d7"+
		"\t\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0011\u0004\u0011\u00e0\b\u0011\u000b\u0011\f\u0011\u00e1"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0003\u0012\u00ec\b\u0012\u0001\u0013\u0001\u0013"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0000\u0000\u0015\u0000\u0002\u0004"+
		"\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \""+
		"$&(\u0000\u0004\u0002\u0000\u000e\u000e !\u0004\u0000\u0003\u0003\u000f"+
		"\u000f \"&&\u0001\u0000!#\u0001\u0000\u0013\u0014\u00fd\u0000+\u0001\u0000"+
		"\u0000\u0000\u00021\u0001\u0000\u0000\u0000\u0004O\u0001\u0000\u0000\u0000"+
		"\u0006R\u0001\u0000\u0000\u0000\bg\u0001\u0000\u0000\u0000\nl\u0001\u0000"+
		"\u0000\u0000\fy\u0001\u0000\u0000\u0000\u000e\u0089\u0001\u0000\u0000"+
		"\u0000\u0010\u0093\u0001\u0000\u0000\u0000\u0012\u009b\u0001\u0000\u0000"+
		"\u0000\u0014\u00a5\u0001\u0000\u0000\u0000\u0016\u00a7\u0001\u0000\u0000"+
		"\u0000\u0018\u00b8\u0001\u0000\u0000\u0000\u001a\u00bc\u0001\u0000\u0000"+
		"\u0000\u001c\u00c6\u0001\u0000\u0000\u0000\u001e\u00d0\u0001\u0000\u0000"+
		"\u0000 \u00da\u0001\u0000\u0000\u0000\"\u00df\u0001\u0000\u0000\u0000"+
		"$\u00eb\u0001\u0000\u0000\u0000&\u00ed\u0001\u0000\u0000\u0000(\u00ef"+
		"\u0001\u0000\u0000\u0000*,\u0003\u0002\u0001\u0000+*\u0001\u0000\u0000"+
		"\u0000,-\u0001\u0000\u0000\u0000-+\u0001\u0000\u0000\u0000-.\u0001\u0000"+
		"\u0000\u0000./\u0001\u0000\u0000\u0000/0\u0005\u0000\u0000\u00010\u0001"+
		"\u0001\u0000\u0000\u000012\u0005\u0015\u0000\u000023\u0005 \u0000\u0000"+
		"35\u0005\u0001\u0000\u000046\u0003\u0004\u0002\u000054\u0001\u0000\u0000"+
		"\u000067\u0001\u0000\u0000\u000075\u0001\u0000\u0000\u000078\u0001\u0000"+
		"\u0000\u000089\u0001\u0000\u0000\u00009:\u0005\u0002\u0000\u0000:\u0003"+
		"\u0001\u0000\u0000\u0000;<\u0005\u0016\u0000\u0000<=\u0005\u0001\u0000"+
		"\u0000=>\u0003\u0006\u0003\u0000>?\u0005\u0002\u0000\u0000?P\u0001\u0000"+
		"\u0000\u0000@A\u0005\u0017\u0000\u0000AB\u0005\u0001\u0000\u0000BC\u0003"+
		"\"\u0011\u0000CD\u0005\u0002\u0000\u0000DP\u0001\u0000\u0000\u0000EF\u0005"+
		"\u0018\u0000\u0000FG\u0005\u0001\u0000\u0000GH\u0003\u0018\f\u0000HI\u0005"+
		"\u0002\u0000\u0000IP\u0001\u0000\u0000\u0000JK\u0005\u0019\u0000\u0000"+
		"KL\u0005\u0001\u0000\u0000LM\u0003\"\u0011\u0000MN\u0005\u0002\u0000\u0000"+
		"NP\u0001\u0000\u0000\u0000O;\u0001\u0000\u0000\u0000O@\u0001\u0000\u0000"+
		"\u0000OE\u0001\u0000\u0000\u0000OJ\u0001\u0000\u0000\u0000P\u0005\u0001"+
		"\u0000\u0000\u0000QS\u0003\b\u0004\u0000RQ\u0001\u0000\u0000\u0000ST\u0001"+
		"\u0000\u0000\u0000TR\u0001\u0000\u0000\u0000TU\u0001\u0000\u0000\u0000"+
		"U\u0007\u0001\u0000\u0000\u0000VW\u0005\u001a\u0000\u0000WX\u0005\u0003"+
		"\u0000\u0000XZ\u0005\u0001\u0000\u0000Y[\u0003\f\u0006\u0000ZY\u0001\u0000"+
		"\u0000\u0000[\\\u0001\u0000\u0000\u0000\\Z\u0001\u0000\u0000\u0000\\]"+
		"\u0001\u0000\u0000\u0000]^\u0001\u0000\u0000\u0000^_\u0005\u0002\u0000"+
		"\u0000_h\u0001\u0000\u0000\u0000`a\u0005\u001b\u0000\u0000ab\u0005\u0003"+
		"\u0000\u0000bc\u0005\u0001\u0000\u0000cd\u0003\n\u0005\u0000de\u0005\u0002"+
		"\u0000\u0000eh\u0001\u0000\u0000\u0000fh\u0003\f\u0006\u0000gV\u0001\u0000"+
		"\u0000\u0000g`\u0001\u0000\u0000\u0000gf\u0001\u0000\u0000\u0000h\t\u0001"+
		"\u0000\u0000\u0000ij\u0005\u001c\u0000\u0000jk\u0005\u0003\u0000\u0000"+
		"km\u0003(\u0014\u0000li\u0001\u0000\u0000\u0000lm\u0001\u0000\u0000\u0000"+
		"mp\u0001\u0000\u0000\u0000no\u0005\u001d\u0000\u0000oq\u0003\u000e\u0007"+
		"\u0000pn\u0001\u0000\u0000\u0000pq\u0001\u0000\u0000\u0000qw\u0001\u0000"+
		"\u0000\u0000rs\u0005\u001f\u0000\u0000st\u0005\u0003\u0000\u0000tu\u0005"+
		"\u0004\u0000\u0000uv\u0005 \u0000\u0000vx\u0005\u0005\u0000\u0000wr\u0001"+
		"\u0000\u0000\u0000wx\u0001\u0000\u0000\u0000x\u000b\u0001\u0000\u0000"+
		"\u0000yz\u0005 \u0000\u0000z{\u0005\u0003\u0000\u0000{~\u0003$\u0012\u0000"+
		"|}\u0005\u0006\u0000\u0000}\u007f\u0003&\u0013\u0000~|\u0001\u0000\u0000"+
		"\u0000~\u007f\u0001\u0000\u0000\u0000\u007f\u0083\u0001\u0000\u0000\u0000"+
		"\u0080\u0082\u0003\u0010\b\u0000\u0081\u0080\u0001\u0000\u0000\u0000\u0082"+
		"\u0085\u0001\u0000\u0000\u0000\u0083\u0081\u0001\u0000\u0000\u0000\u0083"+
		"\u0084\u0001\u0000\u0000\u0000\u0084\u0087\u0001\u0000\u0000\u0000\u0085"+
		"\u0083\u0001\u0000\u0000\u0000\u0086\u0088\u0005%\u0000\u0000\u0087\u0086"+
		"\u0001\u0000\u0000\u0000\u0087\u0088\u0001\u0000\u0000\u0000\u0088\r\u0001"+
		"\u0000\u0000\u0000\u0089\u008a\u0005 \u0000\u0000\u008a\u008b\u0005\u001e"+
		"\u0000\u0000\u008b\u008d\u0005\u0001\u0000\u0000\u008c\u008e\u0003\f\u0006"+
		"\u0000\u008d\u008c\u0001\u0000\u0000\u0000\u008e\u008f\u0001\u0000\u0000"+
		"\u0000\u008f\u008d\u0001\u0000\u0000\u0000\u008f\u0090\u0001\u0000\u0000"+
		"\u0000\u0090\u0091\u0001\u0000\u0000\u0000\u0091\u0092\u0005\u0002\u0000"+
		"\u0000\u0092\u000f\u0001\u0000\u0000\u0000\u0093\u0094\u0005\u0007\u0000"+
		"\u0000\u0094\u0095\u0005 \u0000\u0000\u0095\u0097\u0005\b\u0000\u0000"+
		"\u0096\u0098\u0003\u0012\t\u0000\u0097\u0096\u0001\u0000\u0000\u0000\u0097"+
		"\u0098\u0001\u0000\u0000\u0000\u0098\u0099\u0001\u0000\u0000\u0000\u0099"+
		"\u009a\u0005\t\u0000\u0000\u009a\u0011\u0001\u0000\u0000\u0000\u009b\u00a0"+
		"\u0003\u0014\n\u0000\u009c\u009d\u0005\n\u0000\u0000\u009d\u009f\u0003"+
		"\u0014\n\u0000\u009e\u009c\u0001\u0000\u0000\u0000\u009f\u00a2\u0001\u0000"+
		"\u0000\u0000\u00a0\u009e\u0001\u0000\u0000\u0000\u00a0\u00a1\u0001\u0000"+
		"\u0000\u0000\u00a1\u0013\u0001\u0000\u0000\u0000\u00a2\u00a0\u0001\u0000"+
		"\u0000\u0000\u00a3\u00a6\u0005!\u0000\u0000\u00a4\u00a6\u0003\u0016\u000b"+
		"\u0000\u00a5\u00a3\u0001\u0000\u0000\u0000\u00a5\u00a4\u0001\u0000\u0000"+
		"\u0000\u00a6\u0015\u0001\u0000\u0000\u0000\u00a7\u00b0\u0005\u0004\u0000"+
		"\u0000\u00a8\u00ad\u0005!\u0000\u0000\u00a9\u00aa\u0005\n\u0000\u0000"+
		"\u00aa\u00ac\u0005!\u0000\u0000\u00ab\u00a9\u0001\u0000\u0000\u0000\u00ac"+
		"\u00af\u0001\u0000\u0000\u0000\u00ad\u00ab\u0001\u0000\u0000\u0000\u00ad"+
		"\u00ae\u0001\u0000\u0000\u0000\u00ae\u00b1\u0001\u0000\u0000\u0000\u00af"+
		"\u00ad\u0001\u0000\u0000\u0000\u00b0\u00a8\u0001\u0000\u0000\u0000\u00b0"+
		"\u00b1\u0001\u0000\u0000\u0000\u00b1\u00b2\u0001\u0000\u0000\u0000\u00b2"+
		"\u00b3\u0005\u0005\u0000\u0000\u00b3\u0017\u0001\u0000\u0000\u0000\u00b4"+
		"\u00b9\u0003\"\u0011\u0000\u00b5\u00b9\u0003\u001a\r\u0000\u00b6\u00b9"+
		"\u0003\u001c\u000e\u0000\u00b7\u00b9\u0003\u001e\u000f\u0000\u00b8\u00b4"+
		"\u0001\u0000\u0000\u0000\u00b8\u00b5\u0001\u0000\u0000\u0000\u00b8\u00b6"+
		"\u0001\u0000\u0000\u0000\u00b8\u00b7\u0001\u0000\u0000\u0000\u00b9\u00ba"+
		"\u0001\u0000\u0000\u0000\u00ba\u00b8\u0001\u0000\u0000\u0000\u00ba\u00bb"+
		"\u0001\u0000\u0000\u0000\u00bb\u0019\u0001\u0000\u0000\u0000\u00bc\u00bd"+
		"\u0005\u000b\u0000\u0000\u00bd\u00c1\u0005\u0001\u0000\u0000\u00be\u00c0"+
		"\u0003 \u0010\u0000\u00bf\u00be\u0001\u0000\u0000\u0000\u00c0\u00c3\u0001"+
		"\u0000\u0000\u0000\u00c1\u00bf\u0001\u0000\u0000\u0000\u00c1\u00c2\u0001"+
		"\u0000\u0000\u0000\u00c2\u00c4\u0001\u0000\u0000\u0000\u00c3\u00c1\u0001"+
		"\u0000\u0000\u0000\u00c4\u00c5\u0005\u0002\u0000\u0000\u00c5\u001b\u0001"+
		"\u0000\u0000\u0000\u00c6\u00c7\u0005\f\u0000\u0000\u00c7\u00cb\u0005\u0001"+
		"\u0000\u0000\u00c8\u00ca\u0003 \u0010\u0000\u00c9\u00c8\u0001\u0000\u0000"+
		"\u0000\u00ca\u00cd\u0001\u0000\u0000\u0000\u00cb\u00c9\u0001\u0000\u0000"+
		"\u0000\u00cb\u00cc\u0001\u0000\u0000\u0000\u00cc\u00ce\u0001\u0000\u0000"+
		"\u0000\u00cd\u00cb\u0001\u0000\u0000\u0000\u00ce\u00cf\u0005\u0002\u0000"+
		"\u0000\u00cf\u001d\u0001\u0000\u0000\u0000\u00d0\u00d1\u0005\r\u0000\u0000"+
		"\u00d1\u00d5\u0005\u0001\u0000\u0000\u00d2\u00d4\u0003 \u0010\u0000\u00d3"+
		"\u00d2\u0001\u0000\u0000\u0000\u00d4\u00d7\u0001\u0000\u0000\u0000\u00d5"+
		"\u00d3\u0001\u0000\u0000\u0000\u00d5\u00d6\u0001\u0000\u0000\u0000\u00d6"+
		"\u00d8\u0001\u0000\u0000\u0000\u00d7\u00d5\u0001\u0000\u0000\u0000\u00d8"+
		"\u00d9\u0005\u0002\u0000\u0000\u00d9\u001f\u0001\u0000\u0000\u0000\u00da"+
		"\u00db\u0005 \u0000\u0000\u00db\u00dc\u0005\u0003\u0000\u0000\u00dc\u00dd"+
		"\u0007\u0000\u0000\u0000\u00dd!\u0001\u0000\u0000\u0000\u00de\u00e0\u0007"+
		"\u0001\u0000\u0000\u00df\u00de\u0001\u0000\u0000\u0000\u00e0\u00e1\u0001"+
		"\u0000\u0000\u0000\u00e1\u00df\u0001\u0000\u0000\u0000\u00e1\u00e2\u0001"+
		"\u0000\u0000\u0000\u00e2#\u0001\u0000\u0000\u0000\u00e3\u00ec\u0005\u0010"+
		"\u0000\u0000\u00e4\u00ec\u0005\u0011\u0000\u0000\u00e5\u00ec\u0005\u0012"+
		"\u0000\u0000\u00e6\u00e7\u0005\u0004\u0000\u0000\u00e7\u00e8\u0003$\u0012"+
		"\u0000\u00e8\u00e9\u0005\u0005\u0000\u0000\u00e9\u00ec\u0001\u0000\u0000"+
		"\u0000\u00ea\u00ec\u0005 \u0000\u0000\u00eb\u00e3\u0001\u0000\u0000\u0000"+
		"\u00eb\u00e4\u0001\u0000\u0000\u0000\u00eb\u00e5\u0001\u0000\u0000\u0000"+
		"\u00eb\u00e6\u0001\u0000\u0000\u0000\u00eb\u00ea\u0001\u0000\u0000\u0000"+
		"\u00ec%\u0001\u0000\u0000\u0000\u00ed\u00ee\u0007\u0002\u0000\u0000\u00ee"+
		"\'\u0001\u0000\u0000\u0000\u00ef\u00f0\u0007\u0003\u0000\u0000\u00f0)"+
		"\u0001\u0000\u0000\u0000\u0019-7OT\\glpw~\u0083\u0087\u008f\u0097\u00a0"+
		"\u00a5\u00ad\u00b0\u00b8\u00ba\u00c1\u00cb\u00d5\u00e1\u00eb";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}