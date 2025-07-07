// Code generated from PromptDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promptdsl // PromptDSL
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PromptDSLParser struct {
	*antlr.BaseParser
}

var PromptDSLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func promptdslParserInit() {
	staticData := &PromptDSLParserStaticData
	staticData.LiteralNames = []string{
		"", "'{'", "'}'", "':'", "'['", "']'", "'='", "'@'", "'('", "')'", "','",
		"'classification'", "'summarization'", "'compilation'", "'extra_hint'",
		"'-'", "'string'", "'float'", "'int'", "'md'", "'json'", "'prompt'",
		"'params'", "'system'", "'user'", "'note'", "'in'", "'output'", "'format'",
		"'type'", "'struct'", "'schema'", "", "", "", "", "'|'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "PROMPT", "PARAMS", "SYSTEM", "USER", "NOTE", "IN",
		"OUTPUT", "FORMAT", "TYPE", "STRUCT", "SCHEMA", "ID", "STRING", "NUMBER",
		"BOOL", "PIPE", "SEMI", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"promptFile", "promptDef", "promptBlock", "paramBody", "paramSection",
		"outputBody", "fieldDef", "structDef", "annotation", "annotationArgs",
		"annotationValue", "arrayLiteral", "userBlock", "classificationBlock",
		"summarizationBlock", "compilationBlock", "kvPair", "textBlock", "type",
		"value", "formatType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 40, 242, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 4, 0, 44, 8, 0, 11, 0, 12, 0, 45, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		4, 1, 54, 8, 1, 11, 1, 12, 1, 55, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 80, 8, 2, 1, 3, 4, 3, 83, 8, 3, 11, 3, 12, 3,
		84, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4, 91, 8, 4, 11, 4, 12, 4, 92, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 104, 8, 4, 1, 5, 1,
		5, 1, 5, 3, 5, 109, 8, 5, 1, 5, 1, 5, 3, 5, 113, 8, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 120, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 127, 8,
		6, 1, 6, 5, 6, 130, 8, 6, 10, 6, 12, 6, 133, 9, 6, 1, 6, 3, 6, 136, 8,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 4, 7, 142, 8, 7, 11, 7, 12, 7, 143, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 152, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 5, 9, 159, 8, 9, 10, 9, 12, 9, 162, 9, 9, 1, 10, 1, 10, 3, 10, 166,
		8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 172, 8, 11, 10, 11, 12, 11, 175,
		9, 11, 3, 11, 177, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 4,
		12, 185, 8, 12, 11, 12, 12, 12, 186, 1, 13, 1, 13, 1, 13, 5, 13, 192, 8,
		13, 10, 13, 12, 13, 195, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 5, 14,
		202, 8, 14, 10, 14, 12, 14, 205, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 5, 15, 212, 8, 15, 10, 15, 12, 15, 215, 9, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 17, 4, 17, 224, 8, 17, 11, 17, 12, 17, 225, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 236, 8, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 0, 0, 21, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 0, 4, 2, 0, 14, 14,
		32, 33, 4, 0, 3, 3, 15, 15, 32, 34, 38, 38, 1, 0, 33, 35, 1, 0, 19, 20,
		253, 0, 43, 1, 0, 0, 0, 2, 49, 1, 0, 0, 0, 4, 79, 1, 0, 0, 0, 6, 82, 1,
		0, 0, 0, 8, 103, 1, 0, 0, 0, 10, 108, 1, 0, 0, 0, 12, 121, 1, 0, 0, 0,
		14, 137, 1, 0, 0, 0, 16, 147, 1, 0, 0, 0, 18, 155, 1, 0, 0, 0, 20, 165,
		1, 0, 0, 0, 22, 167, 1, 0, 0, 0, 24, 184, 1, 0, 0, 0, 26, 188, 1, 0, 0,
		0, 28, 198, 1, 0, 0, 0, 30, 208, 1, 0, 0, 0, 32, 218, 1, 0, 0, 0, 34, 223,
		1, 0, 0, 0, 36, 235, 1, 0, 0, 0, 38, 237, 1, 0, 0, 0, 40, 239, 1, 0, 0,
		0, 42, 44, 3, 2, 1, 0, 43, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 43,
		1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 48, 5, 0, 0, 1,
		48, 1, 1, 0, 0, 0, 49, 50, 5, 21, 0, 0, 50, 51, 5, 32, 0, 0, 51, 53, 5,
		1, 0, 0, 52, 54, 3, 4, 2, 0, 53, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55,
		53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 5, 2, 0,
		0, 58, 3, 1, 0, 0, 0, 59, 60, 5, 22, 0, 0, 60, 61, 5, 1, 0, 0, 61, 62,
		3, 6, 3, 0, 62, 63, 5, 2, 0, 0, 63, 80, 1, 0, 0, 0, 64, 65, 5, 23, 0, 0,
		65, 66, 5, 1, 0, 0, 66, 67, 3, 34, 17, 0, 67, 68, 5, 2, 0, 0, 68, 80, 1,
		0, 0, 0, 69, 70, 5, 24, 0, 0, 70, 71, 5, 1, 0, 0, 71, 72, 3, 24, 12, 0,
		72, 73, 5, 2, 0, 0, 73, 80, 1, 0, 0, 0, 74, 75, 5, 25, 0, 0, 75, 76, 5,
		1, 0, 0, 76, 77, 3, 34, 17, 0, 77, 78, 5, 2, 0, 0, 78, 80, 1, 0, 0, 0,
		79, 59, 1, 0, 0, 0, 79, 64, 1, 0, 0, 0, 79, 69, 1, 0, 0, 0, 79, 74, 1,
		0, 0, 0, 80, 5, 1, 0, 0, 0, 81, 83, 3, 8, 4, 0, 82, 81, 1, 0, 0, 0, 83,
		84, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 7, 1, 0, 0,
		0, 86, 87, 5, 26, 0, 0, 87, 88, 5, 3, 0, 0, 88, 90, 5, 1, 0, 0, 89, 91,
		3, 12, 6, 0, 90, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0,
		92, 93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 5, 2, 0, 0, 95, 104, 1,
		0, 0, 0, 96, 97, 5, 27, 0, 0, 97, 98, 5, 3, 0, 0, 98, 99, 5, 1, 0, 0, 99,
		100, 3, 10, 5, 0, 100, 101, 5, 2, 0, 0, 101, 104, 1, 0, 0, 0, 102, 104,
		3, 12, 6, 0, 103, 86, 1, 0, 0, 0, 103, 96, 1, 0, 0, 0, 103, 102, 1, 0,
		0, 0, 104, 9, 1, 0, 0, 0, 105, 106, 5, 28, 0, 0, 106, 107, 5, 3, 0, 0,
		107, 109, 3, 40, 20, 0, 108, 105, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109,
		112, 1, 0, 0, 0, 110, 111, 5, 29, 0, 0, 111, 113, 3, 14, 7, 0, 112, 110,
		1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 119, 1, 0, 0, 0, 114, 115, 5, 31,
		0, 0, 115, 116, 5, 3, 0, 0, 116, 117, 5, 4, 0, 0, 117, 118, 5, 32, 0, 0,
		118, 120, 5, 5, 0, 0, 119, 114, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120,
		11, 1, 0, 0, 0, 121, 122, 5, 32, 0, 0, 122, 123, 5, 3, 0, 0, 123, 126,
		3, 36, 18, 0, 124, 125, 5, 6, 0, 0, 125, 127, 3, 38, 19, 0, 126, 124, 1,
		0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 131, 1, 0, 0, 0, 128, 130, 3, 16, 8,
		0, 129, 128, 1, 0, 0, 0, 130, 133, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131,
		132, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 136,
		5, 37, 0, 0, 135, 134, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 13, 1, 0,
		0, 0, 137, 138, 5, 32, 0, 0, 138, 139, 5, 30, 0, 0, 139, 141, 5, 1, 0,
		0, 140, 142, 3, 12, 6, 0, 141, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143,
		141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146,
		5, 2, 0, 0, 146, 15, 1, 0, 0, 0, 147, 148, 5, 7, 0, 0, 148, 149, 5, 32,
		0, 0, 149, 151, 5, 8, 0, 0, 150, 152, 3, 18, 9, 0, 151, 150, 1, 0, 0, 0,
		151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 5, 9, 0, 0, 154,
		17, 1, 0, 0, 0, 155, 160, 3, 20, 10, 0, 156, 157, 5, 10, 0, 0, 157, 159,
		3, 20, 10, 0, 158, 156, 1, 0, 0, 0, 159, 162, 1, 0, 0, 0, 160, 158, 1,
		0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 19, 1, 0, 0, 0, 162, 160, 1, 0, 0,
		0, 163, 166, 5, 33, 0, 0, 164, 166, 3, 22, 11, 0, 165, 163, 1, 0, 0, 0,
		165, 164, 1, 0, 0, 0, 166, 21, 1, 0, 0, 0, 167, 176, 5, 4, 0, 0, 168, 173,
		5, 33, 0, 0, 169, 170, 5, 10, 0, 0, 170, 172, 5, 33, 0, 0, 171, 169, 1,
		0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0,
		0, 174, 177, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 168, 1, 0, 0, 0, 176,
		177, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 5, 5, 0, 0, 179, 23, 1,
		0, 0, 0, 180, 185, 3, 34, 17, 0, 181, 185, 3, 26, 13, 0, 182, 185, 3, 28,
		14, 0, 183, 185, 3, 30, 15, 0, 184, 180, 1, 0, 0, 0, 184, 181, 1, 0, 0,
		0, 184, 182, 1, 0, 0, 0, 184, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186,
		184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 25, 1, 0, 0, 0, 188, 189, 5,
		11, 0, 0, 189, 193, 5, 1, 0, 0, 190, 192, 3, 32, 16, 0, 191, 190, 1, 0,
		0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0,
		194, 196, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 197, 5, 2, 0, 0, 197,
		27, 1, 0, 0, 0, 198, 199, 5, 12, 0, 0, 199, 203, 5, 1, 0, 0, 200, 202,
		3, 32, 16, 0, 201, 200, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1,
		0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 206, 1, 0, 0, 0, 205, 203, 1, 0, 0,
		0, 206, 207, 5, 2, 0, 0, 207, 29, 1, 0, 0, 0, 208, 209, 5, 13, 0, 0, 209,
		213, 5, 1, 0, 0, 210, 212, 3, 32, 16, 0, 211, 210, 1, 0, 0, 0, 212, 215,
		1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 216, 1, 0,
		0, 0, 215, 213, 1, 0, 0, 0, 216, 217, 5, 2, 0, 0, 217, 31, 1, 0, 0, 0,
		218, 219, 5, 32, 0, 0, 219, 220, 5, 3, 0, 0, 220, 221, 7, 0, 0, 0, 221,
		33, 1, 0, 0, 0, 222, 224, 7, 1, 0, 0, 223, 222, 1, 0, 0, 0, 224, 225, 1,
		0, 0, 0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 35, 1, 0, 0,
		0, 227, 236, 5, 16, 0, 0, 228, 236, 5, 17, 0, 0, 229, 236, 5, 18, 0, 0,
		230, 231, 5, 4, 0, 0, 231, 232, 3, 36, 18, 0, 232, 233, 5, 5, 0, 0, 233,
		236, 1, 0, 0, 0, 234, 236, 5, 32, 0, 0, 235, 227, 1, 0, 0, 0, 235, 228,
		1, 0, 0, 0, 235, 229, 1, 0, 0, 0, 235, 230, 1, 0, 0, 0, 235, 234, 1, 0,
		0, 0, 236, 37, 1, 0, 0, 0, 237, 238, 7, 2, 0, 0, 238, 39, 1, 0, 0, 0, 239,
		240, 7, 3, 0, 0, 240, 41, 1, 0, 0, 0, 25, 45, 55, 79, 84, 92, 103, 108,
		112, 119, 126, 131, 135, 143, 151, 160, 165, 173, 176, 184, 186, 193, 203,
		213, 225, 235,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PromptDSLParserInit initializes any static state used to implement PromptDSLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPromptDSLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PromptDSLParserInit() {
	staticData := &PromptDSLParserStaticData
	staticData.once.Do(promptdslParserInit)
}

// NewPromptDSLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPromptDSLParser(input antlr.TokenStream) *PromptDSLParser {
	PromptDSLParserInit()
	this := new(PromptDSLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PromptDSLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PromptDSL.g4"

	return this
}

// PromptDSLParser tokens.
const (
	PromptDSLParserEOF           = antlr.TokenEOF
	PromptDSLParserT__0          = 1
	PromptDSLParserT__1          = 2
	PromptDSLParserT__2          = 3
	PromptDSLParserT__3          = 4
	PromptDSLParserT__4          = 5
	PromptDSLParserT__5          = 6
	PromptDSLParserT__6          = 7
	PromptDSLParserT__7          = 8
	PromptDSLParserT__8          = 9
	PromptDSLParserT__9          = 10
	PromptDSLParserT__10         = 11
	PromptDSLParserT__11         = 12
	PromptDSLParserT__12         = 13
	PromptDSLParserT__13         = 14
	PromptDSLParserT__14         = 15
	PromptDSLParserT__15         = 16
	PromptDSLParserT__16         = 17
	PromptDSLParserT__17         = 18
	PromptDSLParserT__18         = 19
	PromptDSLParserT__19         = 20
	PromptDSLParserPROMPT        = 21
	PromptDSLParserPARAMS        = 22
	PromptDSLParserSYSTEM        = 23
	PromptDSLParserUSER          = 24
	PromptDSLParserNOTE          = 25
	PromptDSLParserIN            = 26
	PromptDSLParserOUTPUT        = 27
	PromptDSLParserFORMAT        = 28
	PromptDSLParserTYPE          = 29
	PromptDSLParserSTRUCT        = 30
	PromptDSLParserSCHEMA        = 31
	PromptDSLParserID            = 32
	PromptDSLParserSTRING        = 33
	PromptDSLParserNUMBER        = 34
	PromptDSLParserBOOL          = 35
	PromptDSLParserPIPE          = 36
	PromptDSLParserSEMI          = 37
	PromptDSLParserWS            = 38
	PromptDSLParserLINE_COMMENT  = 39
	PromptDSLParserBLOCK_COMMENT = 40
)

// PromptDSLParser rules.
const (
	PromptDSLParserRULE_promptFile          = 0
	PromptDSLParserRULE_promptDef           = 1
	PromptDSLParserRULE_promptBlock         = 2
	PromptDSLParserRULE_paramBody           = 3
	PromptDSLParserRULE_paramSection        = 4
	PromptDSLParserRULE_outputBody          = 5
	PromptDSLParserRULE_fieldDef            = 6
	PromptDSLParserRULE_structDef           = 7
	PromptDSLParserRULE_annotation          = 8
	PromptDSLParserRULE_annotationArgs      = 9
	PromptDSLParserRULE_annotationValue     = 10
	PromptDSLParserRULE_arrayLiteral        = 11
	PromptDSLParserRULE_userBlock           = 12
	PromptDSLParserRULE_classificationBlock = 13
	PromptDSLParserRULE_summarizationBlock  = 14
	PromptDSLParserRULE_compilationBlock    = 15
	PromptDSLParserRULE_kvPair              = 16
	PromptDSLParserRULE_textBlock           = 17
	PromptDSLParserRULE_type                = 18
	PromptDSLParserRULE_value               = 19
	PromptDSLParserRULE_formatType          = 20
)

// IPromptFileContext is an interface to support dynamic dispatch.
type IPromptFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllPromptDef() []IPromptDefContext
	PromptDef(i int) IPromptDefContext

	// IsPromptFileContext differentiates from other interfaces.
	IsPromptFileContext()
}

type PromptFileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPromptFileContext() *PromptFileContext {
	var p = new(PromptFileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_promptFile
	return p
}

func InitEmptyPromptFileContext(p *PromptFileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_promptFile
}

func (*PromptFileContext) IsPromptFileContext() {}

func NewPromptFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PromptFileContext {
	var p = new(PromptFileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_promptFile

	return p
}

func (s *PromptFileContext) GetParser() antlr.Parser { return s.parser }

func (s *PromptFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserEOF, 0)
}

func (s *PromptFileContext) AllPromptDef() []IPromptDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPromptDefContext); ok {
			len++
		}
	}

	tst := make([]IPromptDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPromptDefContext); ok {
			tst[i] = t.(IPromptDefContext)
			i++
		}
	}

	return tst
}

func (s *PromptFileContext) PromptDef(i int) IPromptDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPromptDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPromptDefContext)
}

func (s *PromptFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PromptFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PromptFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterPromptFile(s)
	}
}

func (s *PromptFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitPromptFile(s)
	}
}

func (s *PromptFileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitPromptFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) PromptFile() (localctx IPromptFileContext) {
	localctx = NewPromptFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PromptDSLParserRULE_promptFile)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserPROMPT {
		{
			p.SetState(42)
			p.PromptDef()
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
		p.Match(PromptDSLParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPromptDefContext is an interface to support dynamic dispatch.
type IPromptDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROMPT() antlr.TerminalNode
	ID() antlr.TerminalNode
	AllPromptBlock() []IPromptBlockContext
	PromptBlock(i int) IPromptBlockContext

	// IsPromptDefContext differentiates from other interfaces.
	IsPromptDefContext()
}

type PromptDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPromptDefContext() *PromptDefContext {
	var p = new(PromptDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_promptDef
	return p
}

func InitEmptyPromptDefContext(p *PromptDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_promptDef
}

func (*PromptDefContext) IsPromptDefContext() {}

func NewPromptDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PromptDefContext {
	var p = new(PromptDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_promptDef

	return p
}

func (s *PromptDefContext) GetParser() antlr.Parser { return s.parser }

func (s *PromptDefContext) PROMPT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserPROMPT, 0)
}

func (s *PromptDefContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *PromptDefContext) AllPromptBlock() []IPromptBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPromptBlockContext); ok {
			len++
		}
	}

	tst := make([]IPromptBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPromptBlockContext); ok {
			tst[i] = t.(IPromptBlockContext)
			i++
		}
	}

	return tst
}

func (s *PromptDefContext) PromptBlock(i int) IPromptBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPromptBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPromptBlockContext)
}

func (s *PromptDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PromptDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PromptDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterPromptDef(s)
	}
}

func (s *PromptDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitPromptDef(s)
	}
}

func (s *PromptDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitPromptDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) PromptDef() (localctx IPromptDefContext) {
	localctx = NewPromptDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PromptDSLParserRULE_promptDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(PromptDSLParserPROMPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62914560) != 0) {
		{
			p.SetState(52)
			p.PromptBlock()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
		p.Match(PromptDSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPromptBlockContext is an interface to support dynamic dispatch.
type IPromptBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARAMS() antlr.TerminalNode
	ParamBody() IParamBodyContext
	SYSTEM() antlr.TerminalNode
	TextBlock() ITextBlockContext
	USER() antlr.TerminalNode
	UserBlock() IUserBlockContext
	NOTE() antlr.TerminalNode

	// IsPromptBlockContext differentiates from other interfaces.
	IsPromptBlockContext()
}

type PromptBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPromptBlockContext() *PromptBlockContext {
	var p = new(PromptBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_promptBlock
	return p
}

func InitEmptyPromptBlockContext(p *PromptBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_promptBlock
}

func (*PromptBlockContext) IsPromptBlockContext() {}

func NewPromptBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PromptBlockContext {
	var p = new(PromptBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_promptBlock

	return p
}

func (s *PromptBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *PromptBlockContext) PARAMS() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserPARAMS, 0)
}

func (s *PromptBlockContext) ParamBody() IParamBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamBodyContext)
}

func (s *PromptBlockContext) SYSTEM() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSYSTEM, 0)
}

func (s *PromptBlockContext) TextBlock() ITextBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextBlockContext)
}

func (s *PromptBlockContext) USER() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserUSER, 0)
}

func (s *PromptBlockContext) UserBlock() IUserBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserBlockContext)
}

func (s *PromptBlockContext) NOTE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserNOTE, 0)
}

func (s *PromptBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PromptBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PromptBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterPromptBlock(s)
	}
}

func (s *PromptBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitPromptBlock(s)
	}
}

func (s *PromptBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitPromptBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) PromptBlock() (localctx IPromptBlockContext) {
	localctx = NewPromptBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PromptDSLParserRULE_promptBlock)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserPARAMS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Match(PromptDSLParserPARAMS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Match(PromptDSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.ParamBody()
		}
		{
			p.SetState(62)
			p.Match(PromptDSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserSYSTEM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Match(PromptDSLParserSYSTEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Match(PromptDSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.TextBlock()
		}
		{
			p.SetState(67)
			p.Match(PromptDSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserUSER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.Match(PromptDSLParserUSER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(PromptDSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.UserBlock()
		}
		{
			p.SetState(72)
			p.Match(PromptDSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserNOTE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(74)
			p.Match(PromptDSLParserNOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Match(PromptDSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.TextBlock()
		}
		{
			p.SetState(77)
			p.Match(PromptDSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamBodyContext is an interface to support dynamic dispatch.
type IParamBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParamSection() []IParamSectionContext
	ParamSection(i int) IParamSectionContext

	// IsParamBodyContext differentiates from other interfaces.
	IsParamBodyContext()
}

type ParamBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamBodyContext() *ParamBodyContext {
	var p = new(ParamBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_paramBody
	return p
}

func InitEmptyParamBodyContext(p *ParamBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_paramBody
}

func (*ParamBodyContext) IsParamBodyContext() {}

func NewParamBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamBodyContext {
	var p = new(ParamBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_paramBody

	return p
}

func (s *ParamBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamBodyContext) AllParamSection() []IParamSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamSectionContext); ok {
			len++
		}
	}

	tst := make([]IParamSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamSectionContext); ok {
			tst[i] = t.(IParamSectionContext)
			i++
		}
	}

	return tst
}

func (s *ParamBodyContext) ParamSection(i int) IParamSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamSectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamSectionContext)
}

func (s *ParamBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterParamBody(s)
	}
}

func (s *ParamBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitParamBody(s)
	}
}

func (s *ParamBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitParamBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) ParamBody() (localctx IParamBodyContext) {
	localctx = NewParamBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PromptDSLParserRULE_paramBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4496293888) != 0) {
		{
			p.SetState(81)
			p.ParamSection()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamSectionContext is an interface to support dynamic dispatch.
type IParamSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParamSectionContext differentiates from other interfaces.
	IsParamSectionContext()
}

type ParamSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamSectionContext() *ParamSectionContext {
	var p = new(ParamSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_paramSection
	return p
}

func InitEmptyParamSectionContext(p *ParamSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_paramSection
}

func (*ParamSectionContext) IsParamSectionContext() {}

func NewParamSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamSectionContext {
	var p = new(ParamSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_paramSection

	return p
}

func (s *ParamSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamSectionContext) CopyAll(ctx *ParamSectionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ParamSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleFieldContext struct {
	ParamSectionContext
}

func NewSingleFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleFieldContext {
	var p = new(SingleFieldContext)

	InitEmptyParamSectionContext(&p.ParamSectionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParamSectionContext))

	return p
}

func (s *SingleFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleFieldContext) FieldDef() IFieldDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldDefContext)
}

func (s *SingleFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterSingleField(s)
	}
}

func (s *SingleFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitSingleField(s)
	}
}

func (s *SingleFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitSingleField(s)

	default:
		return t.VisitChildren(s)
	}
}

type InputSectionContext struct {
	ParamSectionContext
}

func NewInputSectionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InputSectionContext {
	var p = new(InputSectionContext)

	InitEmptyParamSectionContext(&p.ParamSectionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParamSectionContext))

	return p
}

func (s *InputSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputSectionContext) IN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserIN, 0)
}

func (s *InputSectionContext) AllFieldDef() []IFieldDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDefContext); ok {
			len++
		}
	}

	tst := make([]IFieldDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDefContext); ok {
			tst[i] = t.(IFieldDefContext)
			i++
		}
	}

	return tst
}

func (s *InputSectionContext) FieldDef(i int) IFieldDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldDefContext)
}

func (s *InputSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterInputSection(s)
	}
}

func (s *InputSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitInputSection(s)
	}
}

func (s *InputSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitInputSection(s)

	default:
		return t.VisitChildren(s)
	}
}

type OutputSectionContext struct {
	ParamSectionContext
}

func NewOutputSectionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OutputSectionContext {
	var p = new(OutputSectionContext)

	InitEmptyParamSectionContext(&p.ParamSectionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParamSectionContext))

	return p
}

func (s *OutputSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputSectionContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserOUTPUT, 0)
}

func (s *OutputSectionContext) OutputBody() IOutputBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutputBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutputBodyContext)
}

func (s *OutputSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterOutputSection(s)
	}
}

func (s *OutputSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitOutputSection(s)
	}
}

func (s *OutputSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitOutputSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) ParamSection() (localctx IParamSectionContext) {
	localctx = NewParamSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PromptDSLParserRULE_paramSection)
	var _la int

	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserIN:
		localctx = NewInputSectionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(PromptDSLParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(PromptDSLParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Match(PromptDSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PromptDSLParserID {
			{
				p.SetState(89)
				p.FieldDef()
			}

			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(94)
			p.Match(PromptDSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserOUTPUT:
		localctx = NewOutputSectionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.Match(PromptDSLParserOUTPUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.Match(PromptDSLParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(PromptDSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.OutputBody()
		}
		{
			p.SetState(100)
			p.Match(PromptDSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserID:
		localctx = NewSingleFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(102)
			p.FieldDef()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOutputBodyContext is an interface to support dynamic dispatch.
type IOutputBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FORMAT() antlr.TerminalNode
	FormatType() IFormatTypeContext
	TYPE() antlr.TerminalNode
	StructDef() IStructDefContext
	SCHEMA() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsOutputBodyContext differentiates from other interfaces.
	IsOutputBodyContext()
}

type OutputBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputBodyContext() *OutputBodyContext {
	var p = new(OutputBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputBody
	return p
}

func InitEmptyOutputBodyContext(p *OutputBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputBody
}

func (*OutputBodyContext) IsOutputBodyContext() {}

func NewOutputBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputBodyContext {
	var p = new(OutputBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_outputBody

	return p
}

func (s *OutputBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputBodyContext) FORMAT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserFORMAT, 0)
}

func (s *OutputBodyContext) FormatType() IFormatTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormatTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormatTypeContext)
}

func (s *OutputBodyContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserTYPE, 0)
}

func (s *OutputBodyContext) StructDef() IStructDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDefContext)
}

func (s *OutputBodyContext) SCHEMA() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSCHEMA, 0)
}

func (s *OutputBodyContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *OutputBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterOutputBody(s)
	}
}

func (s *OutputBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitOutputBody(s)
	}
}

func (s *OutputBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitOutputBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) OutputBody() (localctx IOutputBodyContext) {
	localctx = NewOutputBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PromptDSLParserRULE_outputBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserFORMAT {
		{
			p.SetState(105)
			p.Match(PromptDSLParserFORMAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(PromptDSLParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.FormatType()
		}

	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserTYPE {
		{
			p.SetState(110)
			p.Match(PromptDSLParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.StructDef()
		}

	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSCHEMA {
		{
			p.SetState(114)
			p.Match(PromptDSLParserSCHEMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Match(PromptDSLParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Match(PromptDSLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(PromptDSLParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Match(PromptDSLParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldDefContext is an interface to support dynamic dispatch.
type IFieldDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Type_() ITypeContext
	Value() IValueContext
	AllAnnotation() []IAnnotationContext
	Annotation(i int) IAnnotationContext
	SEMI() antlr.TerminalNode

	// IsFieldDefContext differentiates from other interfaces.
	IsFieldDefContext()
}

type FieldDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDefContext() *FieldDefContext {
	var p = new(FieldDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_fieldDef
	return p
}

func InitEmptyFieldDefContext(p *FieldDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_fieldDef
}

func (*FieldDefContext) IsFieldDefContext() {}

func NewFieldDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDefContext {
	var p = new(FieldDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_fieldDef

	return p
}

func (s *FieldDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDefContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *FieldDefContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FieldDefContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *FieldDefContext) AllAnnotation() []IAnnotationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnnotationContext); ok {
			len++
		}
	}

	tst := make([]IAnnotationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnnotationContext); ok {
			tst[i] = t.(IAnnotationContext)
			i++
		}
	}

	return tst
}

func (s *FieldDefContext) Annotation(i int) IAnnotationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *FieldDefContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSEMI, 0)
}

func (s *FieldDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterFieldDef(s)
	}
}

func (s *FieldDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitFieldDef(s)
	}
}

func (s *FieldDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitFieldDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) FieldDef() (localctx IFieldDefContext) {
	localctx = NewFieldDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PromptDSLParserRULE_fieldDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(PromptDSLParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Type_()
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserT__5 {
		{
			p.SetState(124)
			p.Match(PromptDSLParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Value()
		}

	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserT__6 {
		{
			p.SetState(128)
			p.Annotation()
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSEMI {
		{
			p.SetState(134)
			p.Match(PromptDSLParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDefContext is an interface to support dynamic dispatch.
type IStructDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	STRUCT() antlr.TerminalNode
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext

	// IsStructDefContext differentiates from other interfaces.
	IsStructDefContext()
}

type StructDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDefContext() *StructDefContext {
	var p = new(StructDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_structDef
	return p
}

func InitEmptyStructDefContext(p *StructDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_structDef
}

func (*StructDefContext) IsStructDefContext() {}

func NewStructDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDefContext {
	var p = new(StructDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_structDef

	return p
}

func (s *StructDefContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDefContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *StructDefContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRUCT, 0)
}

func (s *StructDefContext) AllFieldDef() []IFieldDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDefContext); ok {
			len++
		}
	}

	tst := make([]IFieldDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDefContext); ok {
			tst[i] = t.(IFieldDefContext)
			i++
		}
	}

	return tst
}

func (s *StructDefContext) FieldDef(i int) IFieldDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldDefContext)
}

func (s *StructDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterStructDef(s)
	}
}

func (s *StructDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitStructDef(s)
	}
}

func (s *StructDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitStructDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) StructDef() (localctx IStructDefContext) {
	localctx = NewStructDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PromptDSLParserRULE_structDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(PromptDSLParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(140)
			p.FieldDef()
		}

		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(145)
		p.Match(PromptDSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AnnotationArgs() IAnnotationArgsContext

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_annotation
	return p
}

func InitEmptyAnnotationContext(p *AnnotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_annotation
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *AnnotationContext) AnnotationArgs() IAnnotationArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotationArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotationArgsContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (s *AnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PromptDSLParserRULE_annotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(PromptDSLParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(PromptDSLParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserT__3 || _la == PromptDSLParserSTRING {
		{
			p.SetState(150)
			p.AnnotationArgs()
		}

	}
	{
		p.SetState(153)
		p.Match(PromptDSLParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnnotationArgsContext is an interface to support dynamic dispatch.
type IAnnotationArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAnnotationValue() []IAnnotationValueContext
	AnnotationValue(i int) IAnnotationValueContext

	// IsAnnotationArgsContext differentiates from other interfaces.
	IsAnnotationArgsContext()
}

type AnnotationArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationArgsContext() *AnnotationArgsContext {
	var p = new(AnnotationArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_annotationArgs
	return p
}

func InitEmptyAnnotationArgsContext(p *AnnotationArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_annotationArgs
}

func (*AnnotationArgsContext) IsAnnotationArgsContext() {}

func NewAnnotationArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationArgsContext {
	var p = new(AnnotationArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_annotationArgs

	return p
}

func (s *AnnotationArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationArgsContext) AllAnnotationValue() []IAnnotationValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnnotationValueContext); ok {
			len++
		}
	}

	tst := make([]IAnnotationValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnnotationValueContext); ok {
			tst[i] = t.(IAnnotationValueContext)
			i++
		}
	}

	return tst
}

func (s *AnnotationArgsContext) AnnotationValue(i int) IAnnotationValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotationValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotationValueContext)
}

func (s *AnnotationArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterAnnotationArgs(s)
	}
}

func (s *AnnotationArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitAnnotationArgs(s)
	}
}

func (s *AnnotationArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitAnnotationArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) AnnotationArgs() (localctx IAnnotationArgsContext) {
	localctx = NewAnnotationArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PromptDSLParserRULE_annotationArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.AnnotationValue()
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserT__9 {
		{
			p.SetState(156)
			p.Match(PromptDSLParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.AnnotationValue()
		}

		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnnotationValueContext is an interface to support dynamic dispatch.
type IAnnotationValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	ArrayLiteral() IArrayLiteralContext

	// IsAnnotationValueContext differentiates from other interfaces.
	IsAnnotationValueContext()
}

type AnnotationValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationValueContext() *AnnotationValueContext {
	var p = new(AnnotationValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_annotationValue
	return p
}

func InitEmptyAnnotationValueContext(p *AnnotationValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_annotationValue
}

func (*AnnotationValueContext) IsAnnotationValueContext() {}

func NewAnnotationValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationValueContext {
	var p = new(AnnotationValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_annotationValue

	return p
}

func (s *AnnotationValueContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *AnnotationValueContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *AnnotationValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterAnnotationValue(s)
	}
}

func (s *AnnotationValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitAnnotationValue(s)
	}
}

func (s *AnnotationValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitAnnotationValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) AnnotationValue() (localctx IAnnotationValueContext) {
	localctx = NewAnnotationValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PromptDSLParserRULE_annotationValue)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.ArrayLiteral()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserSTRING)
}

func (s *ArrayLiteralContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, i)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PromptDSLParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(PromptDSLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSTRING {
		{
			p.SetState(168)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserT__9 {
			{
				p.SetState(169)
				p.Match(PromptDSLParserT__9)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(170)
				p.Match(PromptDSLParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(178)
		p.Match(PromptDSLParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUserBlockContext is an interface to support dynamic dispatch.
type IUserBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTextBlock() []ITextBlockContext
	TextBlock(i int) ITextBlockContext
	AllClassificationBlock() []IClassificationBlockContext
	ClassificationBlock(i int) IClassificationBlockContext
	AllSummarizationBlock() []ISummarizationBlockContext
	SummarizationBlock(i int) ISummarizationBlockContext
	AllCompilationBlock() []ICompilationBlockContext
	CompilationBlock(i int) ICompilationBlockContext

	// IsUserBlockContext differentiates from other interfaces.
	IsUserBlockContext()
}

type UserBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserBlockContext() *UserBlockContext {
	var p = new(UserBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_userBlock
	return p
}

func InitEmptyUserBlockContext(p *UserBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_userBlock
}

func (*UserBlockContext) IsUserBlockContext() {}

func NewUserBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserBlockContext {
	var p = new(UserBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_userBlock

	return p
}

func (s *UserBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *UserBlockContext) AllTextBlock() []ITextBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITextBlockContext); ok {
			len++
		}
	}

	tst := make([]ITextBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITextBlockContext); ok {
			tst[i] = t.(ITextBlockContext)
			i++
		}
	}

	return tst
}

func (s *UserBlockContext) TextBlock(i int) ITextBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextBlockContext)
}

func (s *UserBlockContext) AllClassificationBlock() []IClassificationBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IClassificationBlockContext); ok {
			len++
		}
	}

	tst := make([]IClassificationBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IClassificationBlockContext); ok {
			tst[i] = t.(IClassificationBlockContext)
			i++
		}
	}

	return tst
}

func (s *UserBlockContext) ClassificationBlock(i int) IClassificationBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassificationBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassificationBlockContext)
}

func (s *UserBlockContext) AllSummarizationBlock() []ISummarizationBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISummarizationBlockContext); ok {
			len++
		}
	}

	tst := make([]ISummarizationBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISummarizationBlockContext); ok {
			tst[i] = t.(ISummarizationBlockContext)
			i++
		}
	}

	return tst
}

func (s *UserBlockContext) SummarizationBlock(i int) ISummarizationBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISummarizationBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISummarizationBlockContext)
}

func (s *UserBlockContext) AllCompilationBlock() []ICompilationBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICompilationBlockContext); ok {
			len++
		}
	}

	tst := make([]ICompilationBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICompilationBlockContext); ok {
			tst[i] = t.(ICompilationBlockContext)
			i++
		}
	}

	return tst
}

func (s *UserBlockContext) CompilationBlock(i int) ICompilationBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompilationBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompilationBlockContext)
}

func (s *UserBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterUserBlock(s)
	}
}

func (s *UserBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitUserBlock(s)
	}
}

func (s *UserBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitUserBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) UserBlock() (localctx IUserBlockContext) {
	localctx = NewUserBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PromptDSLParserRULE_userBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&304942725128) != 0) {
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromptDSLParserT__2, PromptDSLParserT__14, PromptDSLParserID, PromptDSLParserSTRING, PromptDSLParserNUMBER, PromptDSLParserWS:
			{
				p.SetState(180)
				p.TextBlock()
			}

		case PromptDSLParserT__10:
			{
				p.SetState(181)
				p.ClassificationBlock()
			}

		case PromptDSLParserT__11:
			{
				p.SetState(182)
				p.SummarizationBlock()
			}

		case PromptDSLParserT__12:
			{
				p.SetState(183)
				p.CompilationBlock()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClassificationBlockContext is an interface to support dynamic dispatch.
type IClassificationBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllKvPair() []IKvPairContext
	KvPair(i int) IKvPairContext

	// IsClassificationBlockContext differentiates from other interfaces.
	IsClassificationBlockContext()
}

type ClassificationBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassificationBlockContext() *ClassificationBlockContext {
	var p = new(ClassificationBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_classificationBlock
	return p
}

func InitEmptyClassificationBlockContext(p *ClassificationBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_classificationBlock
}

func (*ClassificationBlockContext) IsClassificationBlockContext() {}

func NewClassificationBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassificationBlockContext {
	var p = new(ClassificationBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_classificationBlock

	return p
}

func (s *ClassificationBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassificationBlockContext) AllKvPair() []IKvPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKvPairContext); ok {
			len++
		}
	}

	tst := make([]IKvPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKvPairContext); ok {
			tst[i] = t.(IKvPairContext)
			i++
		}
	}

	return tst
}

func (s *ClassificationBlockContext) KvPair(i int) IKvPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKvPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKvPairContext)
}

func (s *ClassificationBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassificationBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassificationBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterClassificationBlock(s)
	}
}

func (s *ClassificationBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitClassificationBlock(s)
	}
}

func (s *ClassificationBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitClassificationBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) ClassificationBlock() (localctx IClassificationBlockContext) {
	localctx = NewClassificationBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PromptDSLParserRULE_classificationBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(PromptDSLParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserID {
		{
			p.SetState(190)
			p.KvPair()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(196)
		p.Match(PromptDSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISummarizationBlockContext is an interface to support dynamic dispatch.
type ISummarizationBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllKvPair() []IKvPairContext
	KvPair(i int) IKvPairContext

	// IsSummarizationBlockContext differentiates from other interfaces.
	IsSummarizationBlockContext()
}

type SummarizationBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySummarizationBlockContext() *SummarizationBlockContext {
	var p = new(SummarizationBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_summarizationBlock
	return p
}

func InitEmptySummarizationBlockContext(p *SummarizationBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_summarizationBlock
}

func (*SummarizationBlockContext) IsSummarizationBlockContext() {}

func NewSummarizationBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SummarizationBlockContext {
	var p = new(SummarizationBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_summarizationBlock

	return p
}

func (s *SummarizationBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *SummarizationBlockContext) AllKvPair() []IKvPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKvPairContext); ok {
			len++
		}
	}

	tst := make([]IKvPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKvPairContext); ok {
			tst[i] = t.(IKvPairContext)
			i++
		}
	}

	return tst
}

func (s *SummarizationBlockContext) KvPair(i int) IKvPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKvPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKvPairContext)
}

func (s *SummarizationBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SummarizationBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SummarizationBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterSummarizationBlock(s)
	}
}

func (s *SummarizationBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitSummarizationBlock(s)
	}
}

func (s *SummarizationBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitSummarizationBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) SummarizationBlock() (localctx ISummarizationBlockContext) {
	localctx = NewSummarizationBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PromptDSLParserRULE_summarizationBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(PromptDSLParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserID {
		{
			p.SetState(200)
			p.KvPair()
		}

		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(206)
		p.Match(PromptDSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompilationBlockContext is an interface to support dynamic dispatch.
type ICompilationBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllKvPair() []IKvPairContext
	KvPair(i int) IKvPairContext

	// IsCompilationBlockContext differentiates from other interfaces.
	IsCompilationBlockContext()
}

type CompilationBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationBlockContext() *CompilationBlockContext {
	var p = new(CompilationBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_compilationBlock
	return p
}

func InitEmptyCompilationBlockContext(p *CompilationBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_compilationBlock
}

func (*CompilationBlockContext) IsCompilationBlockContext() {}

func NewCompilationBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationBlockContext {
	var p = new(CompilationBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_compilationBlock

	return p
}

func (s *CompilationBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationBlockContext) AllKvPair() []IKvPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKvPairContext); ok {
			len++
		}
	}

	tst := make([]IKvPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKvPairContext); ok {
			tst[i] = t.(IKvPairContext)
			i++
		}
	}

	return tst
}

func (s *CompilationBlockContext) KvPair(i int) IKvPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKvPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKvPairContext)
}

func (s *CompilationBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterCompilationBlock(s)
	}
}

func (s *CompilationBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitCompilationBlock(s)
	}
}

func (s *CompilationBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitCompilationBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) CompilationBlock() (localctx ICompilationBlockContext) {
	localctx = NewCompilationBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PromptDSLParserRULE_compilationBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(PromptDSLParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserID {
		{
			p.SetState(210)
			p.KvPair()
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
		p.Match(PromptDSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKvPairContext is an interface to support dynamic dispatch.
type IKvPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsKvPairContext differentiates from other interfaces.
	IsKvPairContext()
}

type KvPairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKvPairContext() *KvPairContext {
	var p = new(KvPairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_kvPair
	return p
}

func InitEmptyKvPairContext(p *KvPairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_kvPair
}

func (*KvPairContext) IsKvPairContext() {}

func NewKvPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KvPairContext {
	var p = new(KvPairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_kvPair

	return p
}

func (s *KvPairContext) GetParser() antlr.Parser { return s.parser }

func (s *KvPairContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserID)
}

func (s *KvPairContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, i)
}

func (s *KvPairContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *KvPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KvPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KvPairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterKvPair(s)
	}
}

func (s *KvPairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitKvPair(s)
	}
}

func (s *KvPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitKvPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) KvPair() (localctx IKvPairContext) {
	localctx = NewKvPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PromptDSLParserRULE_kvPair)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.Match(PromptDSLParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&12884918272) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITextBlockContext is an interface to support dynamic dispatch.
type ITextBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsTextBlockContext differentiates from other interfaces.
	IsTextBlockContext()
}

type TextBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextBlockContext() *TextBlockContext {
	var p = new(TextBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_textBlock
	return p
}

func InitEmptyTextBlockContext(p *TextBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_textBlock
}

func (*TextBlockContext) IsTextBlockContext() {}

func NewTextBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextBlockContext {
	var p = new(TextBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_textBlock

	return p
}

func (s *TextBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *TextBlockContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserSTRING)
}

func (s *TextBlockContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, i)
}

func (s *TextBlockContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserID)
}

func (s *TextBlockContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, i)
}

func (s *TextBlockContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserNUMBER)
}

func (s *TextBlockContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserNUMBER, i)
}

func (s *TextBlockContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserWS)
}

func (s *TextBlockContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserWS, i)
}

func (s *TextBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterTextBlock(s)
	}
}

func (s *TextBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitTextBlock(s)
	}
}

func (s *TextBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitTextBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) TextBlock() (localctx ITextBlockContext) {
	localctx = NewTextBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PromptDSLParserRULE_textBlock)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(222)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&304942710792) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ID() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PromptDSLParserRULE_type)
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.Match(PromptDSLParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.Match(PromptDSLParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserT__17:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(229)
			p.Match(PromptDSLParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserT__3:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(230)
			p.Match(PromptDSLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Type_()
		}
		{
			p.SetState(232)
			p.Match(PromptDSLParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserID:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(234)
			p.Match(PromptDSLParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOL() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserNUMBER, 0)
}

func (s *ValueContext) BOOL() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserBOOL, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PromptDSLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&60129542144) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFormatTypeContext is an interface to support dynamic dispatch.
type IFormatTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFormatTypeContext differentiates from other interfaces.
	IsFormatTypeContext()
}

type FormatTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormatTypeContext() *FormatTypeContext {
	var p = new(FormatTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_formatType
	return p
}

func InitEmptyFormatTypeContext(p *FormatTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_formatType
}

func (*FormatTypeContext) IsFormatTypeContext() {}

func NewFormatTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormatTypeContext {
	var p = new(FormatTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_formatType

	return p
}

func (s *FormatTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *FormatTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormatTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormatTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterFormatType(s)
	}
}

func (s *FormatTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitFormatType(s)
	}
}

func (s *FormatTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitFormatType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) FormatType() (localctx IFormatTypeContext) {
	localctx = NewFormatTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PromptDSLParserRULE_formatType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromptDSLParserT__18 || _la == PromptDSLParserT__19) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
