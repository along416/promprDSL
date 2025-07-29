// Code generated from ./promptdslcore/grammar/PromptDSLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromptDSLParser
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

var PromptDSLParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func promptdslparserParserInit() {
	staticData := &PromptDSLParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'string'", "'float'", "'int'", "'prompt'", "'params'", "'system'",
		"'user'", "'note'", "'input'", "'output'", "'format'", "'type'", "'struct'",
		"'before'", "'schema'", "'parse'", "'jsonfix'", "'markdown'", "'if'",
		"'else'", "'outputspec'", "'for'", "", "", "", "'{'", "", "'('", "')'",
		"':'", "'='", "','", "'.'", "'=='", "'!='", "'@'", "'md'", "'json'",
		"'['", "']'", "'++'", "'--'", "'-'", "'*'", "'/'", "'%'", "'+='", "'-='",
		"'*='", "'/='", "'%='", "'<'", "'<='", "'>'", "'>='", "", "", "", "",
		"'|'", "';'", "'+'",
	}
	staticData.SymbolicNames = []string{
		"", "STRING_TYPE", "FLOAT_TYPE", "INT_TYPE", "PROMPT", "PARAMS", "SYSTEM",
		"USER", "NOTE", "INPUT", "OUTPUT", "FORMAT", "TYPE", "STRUCT", "BEFORE",
		"SCHEMA", "PARSE", "JSONFIX", "MARKDOWN", "IF", "ELSE", "OUTPUTSPEC",
		"FOR", "FIX", "AFTER", "ARRAY_OUTPUTSPEC", "LBRACE", "RBRACE", "LPAREN",
		"RPAREN", "COLON", "EQUAL", "COMMA", "DOT", "EQEQ", "NOTEQ", "AT", "MD",
		"JSON", "LBRACK", "RBRACK", "INCREMENT", "DECREMENT", "MINUS", "STAR",
		"SLASH", "MOD", "PLUSEQ", "MINUSEQ", "MULTEQ", "DIVEQ", "MODEQ", "LT",
		"LTE", "GT", "GTE", "ID", "STRING", "NUMBER", "BOOL", "PIPE", "SEMI",
		"PLUS", "WS", "LINE_COMMENT", "BLOCK_COMMENT", "CODE_STRING", "CODE_TEXT",
	}
	staticData.RuleNames = []string{
		"promptFile", "promptDef", "promptBlock", "inputSection", "outputSection",
		"outputStruct", "outputMarkdown", "beforeSection", "beforeContent",
		"varDef", "systemSection", "sysContent", "userSection", "userContent",
		"moduleDef", "moduleContent", "thencontent", "elsecontent", "forcontent",
		"ifStatement", "condition", "forStatement", "assignExpr", "updateExpr",
		"dslCallExpr", "expr", "fieldDef", "textLine", "paramPath", "structDef",
		"annotation", "annotationArgs", "annotationValue", "arrayLiteral", "defaultAnnotation",
		"fixSection", "afterSection", "codeBlockContent", "type", "value", "formatType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 67, 463, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 4, 0, 84,
		8, 0, 11, 0, 12, 0, 85, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 94, 8,
		1, 11, 1, 12, 1, 95, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 107, 8, 2, 1, 3, 1, 3, 1, 3, 4, 3, 112, 8, 3, 11, 3, 12, 3, 113,
		1, 3, 1, 3, 1, 4, 5, 4, 119, 8, 4, 10, 4, 12, 4, 122, 9, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 127, 8, 4, 1, 5, 1, 5, 4, 5, 131, 8, 5, 11, 5, 12, 5, 132,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 143, 8, 7, 10, 7,
		12, 7, 146, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 154, 8, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 4, 10, 163, 8, 10, 11, 10,
		12, 10, 164, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 171, 8, 10, 11, 10, 12,
		10, 172, 1, 10, 1, 10, 3, 10, 177, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 3, 11, 186, 8, 11, 1, 12, 1, 12, 1, 12, 4, 12, 191, 8,
		12, 11, 12, 12, 12, 192, 1, 12, 1, 12, 1, 12, 1, 12, 4, 12, 199, 8, 12,
		11, 12, 12, 12, 200, 1, 12, 1, 12, 3, 12, 205, 8, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 214, 8, 13, 1, 14, 1, 14, 1, 14,
		5, 14, 219, 8, 14, 10, 14, 12, 14, 222, 9, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 233, 8, 15, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5,
		19, 247, 8, 19, 10, 19, 12, 19, 250, 9, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		5, 19, 256, 8, 19, 10, 19, 12, 19, 259, 9, 19, 1, 19, 3, 19, 262, 8, 19,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 269, 8, 20, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 281, 8, 21,
		10, 21, 12, 21, 284, 9, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 3, 23, 305, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5,
		24, 312, 8, 24, 10, 24, 12, 24, 315, 9, 24, 3, 24, 317, 8, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25,
		330, 8, 25, 1, 25, 1, 25, 1, 25, 5, 25, 335, 8, 25, 10, 25, 12, 25, 338,
		9, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 345, 8, 26, 1, 26, 5,
		26, 348, 8, 26, 10, 26, 12, 26, 351, 9, 26, 1, 26, 3, 26, 354, 8, 26, 1,
		27, 1, 27, 1, 27, 3, 27, 359, 8, 27, 1, 28, 1, 28, 1, 28, 5, 28, 364, 8,
		28, 10, 28, 12, 28, 367, 9, 28, 1, 29, 1, 29, 1, 29, 1, 29, 4, 29, 373,
		8, 29, 11, 29, 12, 29, 374, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 3,
		30, 383, 8, 30, 1, 30, 3, 30, 386, 8, 30, 1, 31, 1, 31, 1, 31, 5, 31, 391,
		8, 31, 10, 31, 12, 31, 394, 9, 31, 1, 32, 1, 32, 3, 32, 398, 8, 32, 1,
		33, 1, 33, 1, 33, 1, 33, 5, 33, 404, 8, 33, 10, 33, 12, 33, 407, 9, 33,
		3, 33, 409, 8, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 417,
		8, 34, 1, 34, 3, 34, 420, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 436,
		8, 37, 10, 37, 12, 37, 439, 9, 37, 1, 38, 1, 38, 1, 38, 5, 38, 444, 8,
		38, 10, 38, 12, 38, 447, 9, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 3, 38, 457, 8, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 0,
		1, 50, 41, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 80, 0, 5, 2, 0, 34, 35, 52, 55, 2, 0, 43, 46, 62, 62,
		4, 0, 9, 10, 14, 14, 24, 24, 56, 56, 1, 0, 57, 59, 1, 0, 37, 38, 504, 0,
		83, 1, 0, 0, 0, 2, 89, 1, 0, 0, 0, 4, 106, 1, 0, 0, 0, 6, 108, 1, 0, 0,
		0, 8, 120, 1, 0, 0, 0, 10, 128, 1, 0, 0, 0, 12, 136, 1, 0, 0, 0, 14, 139,
		1, 0, 0, 0, 16, 153, 1, 0, 0, 0, 18, 155, 1, 0, 0, 0, 20, 176, 1, 0, 0,
		0, 22, 185, 1, 0, 0, 0, 24, 204, 1, 0, 0, 0, 26, 213, 1, 0, 0, 0, 28, 215,
		1, 0, 0, 0, 30, 232, 1, 0, 0, 0, 32, 234, 1, 0, 0, 0, 34, 236, 1, 0, 0,
		0, 36, 238, 1, 0, 0, 0, 38, 240, 1, 0, 0, 0, 40, 268, 1, 0, 0, 0, 42, 270,
		1, 0, 0, 0, 44, 287, 1, 0, 0, 0, 46, 291, 1, 0, 0, 0, 48, 306, 1, 0, 0,
		0, 50, 329, 1, 0, 0, 0, 52, 339, 1, 0, 0, 0, 54, 358, 1, 0, 0, 0, 56, 360,
		1, 0, 0, 0, 58, 368, 1, 0, 0, 0, 60, 378, 1, 0, 0, 0, 62, 387, 1, 0, 0,
		0, 64, 397, 1, 0, 0, 0, 66, 399, 1, 0, 0, 0, 68, 412, 1, 0, 0, 0, 70, 421,
		1, 0, 0, 0, 72, 425, 1, 0, 0, 0, 74, 437, 1, 0, 0, 0, 76, 456, 1, 0, 0,
		0, 78, 458, 1, 0, 0, 0, 80, 460, 1, 0, 0, 0, 82, 84, 3, 2, 1, 0, 83, 82,
		1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0,
		86, 87, 1, 0, 0, 0, 87, 88, 5, 0, 0, 1, 88, 1, 1, 0, 0, 0, 89, 90, 5, 4,
		0, 0, 90, 91, 5, 56, 0, 0, 91, 93, 5, 26, 0, 0, 92, 94, 3, 4, 2, 0, 93,
		92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0,
		0, 96, 97, 1, 0, 0, 0, 97, 98, 5, 27, 0, 0, 98, 3, 1, 0, 0, 0, 99, 107,
		3, 6, 3, 0, 100, 107, 3, 8, 4, 0, 101, 107, 3, 20, 10, 0, 102, 107, 3,
		24, 12, 0, 103, 107, 3, 72, 36, 0, 104, 107, 3, 70, 35, 0, 105, 107, 3,
		28, 14, 0, 106, 99, 1, 0, 0, 0, 106, 100, 1, 0, 0, 0, 106, 101, 1, 0, 0,
		0, 106, 102, 1, 0, 0, 0, 106, 103, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106,
		105, 1, 0, 0, 0, 107, 5, 1, 0, 0, 0, 108, 109, 5, 9, 0, 0, 109, 111, 5,
		26, 0, 0, 110, 112, 3, 52, 26, 0, 111, 110, 1, 0, 0, 0, 112, 113, 1, 0,
		0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0,
		115, 116, 5, 27, 0, 0, 116, 7, 1, 0, 0, 0, 117, 119, 3, 68, 34, 0, 118,
		117, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121,
		1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 126, 5, 10,
		0, 0, 124, 127, 3, 10, 5, 0, 125, 127, 3, 12, 6, 0, 126, 124, 1, 0, 0,
		0, 126, 125, 1, 0, 0, 0, 127, 9, 1, 0, 0, 0, 128, 130, 5, 26, 0, 0, 129,
		131, 3, 52, 26, 0, 130, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 130,
		1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 5, 27,
		0, 0, 135, 11, 1, 0, 0, 0, 136, 137, 5, 30, 0, 0, 137, 138, 5, 18, 0, 0,
		138, 13, 1, 0, 0, 0, 139, 140, 5, 14, 0, 0, 140, 144, 5, 26, 0, 0, 141,
		143, 3, 16, 8, 0, 142, 141, 1, 0, 0, 0, 143, 146, 1, 0, 0, 0, 144, 142,
		1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 147, 1, 0, 0, 0, 146, 144, 1, 0,
		0, 0, 147, 148, 5, 27, 0, 0, 148, 15, 1, 0, 0, 0, 149, 154, 3, 18, 9, 0,
		150, 154, 3, 50, 25, 0, 151, 154, 3, 38, 19, 0, 152, 154, 3, 54, 27, 0,
		153, 149, 1, 0, 0, 0, 153, 150, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153,
		152, 1, 0, 0, 0, 154, 17, 1, 0, 0, 0, 155, 156, 5, 56, 0, 0, 156, 157,
		5, 31, 0, 0, 157, 158, 3, 50, 25, 0, 158, 19, 1, 0, 0, 0, 159, 160, 5,
		6, 0, 0, 160, 162, 5, 26, 0, 0, 161, 163, 5, 56, 0, 0, 162, 161, 1, 0,
		0, 0, 163, 164, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0,
		165, 166, 1, 0, 0, 0, 166, 177, 5, 27, 0, 0, 167, 168, 5, 6, 0, 0, 168,
		170, 5, 26, 0, 0, 169, 171, 3, 22, 11, 0, 170, 169, 1, 0, 0, 0, 171, 172,
		1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174, 1, 0,
		0, 0, 174, 175, 5, 27, 0, 0, 175, 177, 1, 0, 0, 0, 176, 159, 1, 0, 0, 0,
		176, 167, 1, 0, 0, 0, 177, 21, 1, 0, 0, 0, 178, 186, 3, 38, 19, 0, 179,
		186, 3, 56, 28, 0, 180, 186, 3, 42, 21, 0, 181, 186, 5, 25, 0, 0, 182,
		186, 5, 21, 0, 0, 183, 186, 3, 50, 25, 0, 184, 186, 3, 54, 27, 0, 185,
		178, 1, 0, 0, 0, 185, 179, 1, 0, 0, 0, 185, 180, 1, 0, 0, 0, 185, 181,
		1, 0, 0, 0, 185, 182, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 184, 1, 0,
		0, 0, 186, 23, 1, 0, 0, 0, 187, 188, 5, 7, 0, 0, 188, 190, 5, 26, 0, 0,
		189, 191, 5, 56, 0, 0, 190, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192,
		190, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 205,
		5, 27, 0, 0, 195, 196, 5, 7, 0, 0, 196, 198, 5, 26, 0, 0, 197, 199, 3,
		26, 13, 0, 198, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 198, 1, 0,
		0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 5, 27, 0, 0,
		203, 205, 1, 0, 0, 0, 204, 187, 1, 0, 0, 0, 204, 195, 1, 0, 0, 0, 205,
		25, 1, 0, 0, 0, 206, 214, 3, 38, 19, 0, 207, 214, 3, 56, 28, 0, 208, 214,
		3, 42, 21, 0, 209, 214, 5, 25, 0, 0, 210, 214, 5, 21, 0, 0, 211, 214, 3,
		50, 25, 0, 212, 214, 3, 54, 27, 0, 213, 206, 1, 0, 0, 0, 213, 207, 1, 0,
		0, 0, 213, 208, 1, 0, 0, 0, 213, 209, 1, 0, 0, 0, 213, 210, 1, 0, 0, 0,
		213, 211, 1, 0, 0, 0, 213, 212, 1, 0, 0, 0, 214, 27, 1, 0, 0, 0, 215, 216,
		5, 56, 0, 0, 216, 220, 5, 26, 0, 0, 217, 219, 3, 30, 15, 0, 218, 217, 1,
		0, 0, 0, 219, 222, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0,
		0, 221, 223, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 223, 224, 5, 27, 0, 0, 224,
		29, 1, 0, 0, 0, 225, 233, 3, 38, 19, 0, 226, 233, 3, 56, 28, 0, 227, 233,
		3, 42, 21, 0, 228, 233, 5, 25, 0, 0, 229, 233, 5, 21, 0, 0, 230, 233, 3,
		50, 25, 0, 231, 233, 3, 54, 27, 0, 232, 225, 1, 0, 0, 0, 232, 226, 1, 0,
		0, 0, 232, 227, 1, 0, 0, 0, 232, 228, 1, 0, 0, 0, 232, 229, 1, 0, 0, 0,
		232, 230, 1, 0, 0, 0, 232, 231, 1, 0, 0, 0, 233, 31, 1, 0, 0, 0, 234, 235,
		3, 26, 13, 0, 235, 33, 1, 0, 0, 0, 236, 237, 3, 26, 13, 0, 237, 35, 1,
		0, 0, 0, 238, 239, 3, 26, 13, 0, 239, 37, 1, 0, 0, 0, 240, 241, 5, 19,
		0, 0, 241, 242, 5, 28, 0, 0, 242, 243, 3, 40, 20, 0, 243, 244, 5, 29, 0,
		0, 244, 248, 5, 26, 0, 0, 245, 247, 3, 32, 16, 0, 246, 245, 1, 0, 0, 0,
		247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249,
		251, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 261, 5, 27, 0, 0, 252, 253,
		5, 20, 0, 0, 253, 257, 5, 26, 0, 0, 254, 256, 3, 34, 17, 0, 255, 254, 1,
		0, 0, 0, 256, 259, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 257, 258, 1, 0, 0,
		0, 258, 260, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 260, 262, 5, 27, 0, 0, 261,
		252, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 39, 1, 0, 0, 0, 263, 264, 3,
		50, 25, 0, 264, 265, 7, 0, 0, 0, 265, 266, 3, 50, 25, 0, 266, 269, 1, 0,
		0, 0, 267, 269, 3, 50, 25, 0, 268, 263, 1, 0, 0, 0, 268, 267, 1, 0, 0,
		0, 269, 41, 1, 0, 0, 0, 270, 271, 5, 22, 0, 0, 271, 272, 5, 28, 0, 0, 272,
		273, 3, 44, 22, 0, 273, 274, 5, 61, 0, 0, 274, 275, 3, 40, 20, 0, 275,
		276, 5, 61, 0, 0, 276, 277, 3, 46, 23, 0, 277, 278, 5, 29, 0, 0, 278, 282,
		5, 26, 0, 0, 279, 281, 3, 36, 18, 0, 280, 279, 1, 0, 0, 0, 281, 284, 1,
		0, 0, 0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 285, 1, 0, 0,
		0, 284, 282, 1, 0, 0, 0, 285, 286, 5, 27, 0, 0, 286, 43, 1, 0, 0, 0, 287,
		288, 3, 56, 28, 0, 288, 289, 5, 31, 0, 0, 289, 290, 3, 50, 25, 0, 290,
		45, 1, 0, 0, 0, 291, 304, 3, 56, 28, 0, 292, 305, 5, 41, 0, 0, 293, 305,
		5, 42, 0, 0, 294, 295, 5, 47, 0, 0, 295, 305, 3, 50, 25, 0, 296, 297, 5,
		48, 0, 0, 297, 305, 3, 50, 25, 0, 298, 299, 5, 49, 0, 0, 299, 305, 3, 50,
		25, 0, 300, 301, 5, 50, 0, 0, 301, 305, 3, 50, 25, 0, 302, 303, 5, 51,
		0, 0, 303, 305, 3, 50, 25, 0, 304, 292, 1, 0, 0, 0, 304, 293, 1, 0, 0,
		0, 304, 294, 1, 0, 0, 0, 304, 296, 1, 0, 0, 0, 304, 298, 1, 0, 0, 0, 304,
		300, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 305, 47, 1, 0, 0, 0, 306, 307, 3,
		56, 28, 0, 307, 316, 5, 28, 0, 0, 308, 313, 3, 50, 25, 0, 309, 310, 5,
		32, 0, 0, 310, 312, 3, 50, 25, 0, 311, 309, 1, 0, 0, 0, 312, 315, 1, 0,
		0, 0, 313, 311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 317, 1, 0, 0, 0,
		315, 313, 1, 0, 0, 0, 316, 308, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317,
		318, 1, 0, 0, 0, 318, 319, 5, 29, 0, 0, 319, 49, 1, 0, 0, 0, 320, 321,
		6, 25, -1, 0, 321, 330, 3, 56, 28, 0, 322, 330, 5, 57, 0, 0, 323, 330,
		5, 58, 0, 0, 324, 330, 5, 59, 0, 0, 325, 326, 5, 28, 0, 0, 326, 327, 3,
		50, 25, 0, 327, 328, 5, 29, 0, 0, 328, 330, 1, 0, 0, 0, 329, 320, 1, 0,
		0, 0, 329, 322, 1, 0, 0, 0, 329, 323, 1, 0, 0, 0, 329, 324, 1, 0, 0, 0,
		329, 325, 1, 0, 0, 0, 330, 336, 1, 0, 0, 0, 331, 332, 10, 6, 0, 0, 332,
		333, 7, 1, 0, 0, 333, 335, 3, 50, 25, 7, 334, 331, 1, 0, 0, 0, 335, 338,
		1, 0, 0, 0, 336, 334, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 51, 1, 0,
		0, 0, 338, 336, 1, 0, 0, 0, 339, 340, 5, 56, 0, 0, 340, 341, 5, 30, 0,
		0, 341, 344, 3, 76, 38, 0, 342, 343, 5, 31, 0, 0, 343, 345, 3, 78, 39,
		0, 344, 342, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 349, 1, 0, 0, 0, 346,
		348, 3, 60, 30, 0, 347, 346, 1, 0, 0, 0, 348, 351, 1, 0, 0, 0, 349, 347,
		1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 353, 1, 0, 0, 0, 351, 349, 1, 0,
		0, 0, 352, 354, 5, 61, 0, 0, 353, 352, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0,
		354, 53, 1, 0, 0, 0, 355, 359, 5, 57, 0, 0, 356, 359, 5, 64, 0, 0, 357,
		359, 3, 56, 28, 0, 358, 355, 1, 0, 0, 0, 358, 356, 1, 0, 0, 0, 358, 357,
		1, 0, 0, 0, 359, 55, 1, 0, 0, 0, 360, 365, 7, 2, 0, 0, 361, 362, 5, 33,
		0, 0, 362, 364, 5, 56, 0, 0, 363, 361, 1, 0, 0, 0, 364, 367, 1, 0, 0, 0,
		365, 363, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 57, 1, 0, 0, 0, 367, 365,
		1, 0, 0, 0, 368, 369, 5, 56, 0, 0, 369, 370, 5, 13, 0, 0, 370, 372, 5,
		26, 0, 0, 371, 373, 3, 52, 26, 0, 372, 371, 1, 0, 0, 0, 373, 374, 1, 0,
		0, 0, 374, 372, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0,
		376, 377, 5, 27, 0, 0, 377, 59, 1, 0, 0, 0, 378, 379, 5, 36, 0, 0, 379,
		385, 5, 56, 0, 0, 380, 382, 5, 28, 0, 0, 381, 383, 3, 62, 31, 0, 382, 381,
		1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 386, 5, 29,
		0, 0, 385, 380, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 61, 1, 0, 0, 0,
		387, 392, 3, 64, 32, 0, 388, 389, 5, 32, 0, 0, 389, 391, 3, 64, 32, 0,
		390, 388, 1, 0, 0, 0, 391, 394, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 392,
		393, 1, 0, 0, 0, 393, 63, 1, 0, 0, 0, 394, 392, 1, 0, 0, 0, 395, 398, 5,
		57, 0, 0, 396, 398, 3, 66, 33, 0, 397, 395, 1, 0, 0, 0, 397, 396, 1, 0,
		0, 0, 398, 65, 1, 0, 0, 0, 399, 408, 5, 26, 0, 0, 400, 405, 5, 57, 0, 0,
		401, 402, 5, 32, 0, 0, 402, 404, 5, 57, 0, 0, 403, 401, 1, 0, 0, 0, 404,
		407, 1, 0, 0, 0, 405, 403, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 409,
		1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 408, 400, 1, 0, 0, 0, 408, 409, 1, 0,
		0, 0, 409, 410, 1, 0, 0, 0, 410, 411, 5, 27, 0, 0, 411, 67, 1, 0, 0, 0,
		412, 413, 5, 36, 0, 0, 413, 419, 5, 56, 0, 0, 414, 416, 5, 28, 0, 0, 415,
		417, 3, 62, 31, 0, 416, 415, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 418,
		1, 0, 0, 0, 418, 420, 5, 29, 0, 0, 419, 414, 1, 0, 0, 0, 419, 420, 1, 0,
		0, 0, 420, 69, 1, 0, 0, 0, 421, 422, 5, 23, 0, 0, 422, 423, 3, 74, 37,
		0, 423, 424, 5, 27, 0, 0, 424, 71, 1, 0, 0, 0, 425, 426, 5, 24, 0, 0, 426,
		427, 3, 74, 37, 0, 427, 428, 5, 27, 0, 0, 428, 73, 1, 0, 0, 0, 429, 436,
		5, 67, 0, 0, 430, 436, 5, 66, 0, 0, 431, 432, 5, 26, 0, 0, 432, 433, 3,
		74, 37, 0, 433, 434, 5, 27, 0, 0, 434, 436, 1, 0, 0, 0, 435, 429, 1, 0,
		0, 0, 435, 430, 1, 0, 0, 0, 435, 431, 1, 0, 0, 0, 436, 439, 1, 0, 0, 0,
		437, 435, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 75, 1, 0, 0, 0, 439, 437,
		1, 0, 0, 0, 440, 441, 5, 13, 0, 0, 441, 445, 5, 26, 0, 0, 442, 444, 3,
		52, 26, 0, 443, 442, 1, 0, 0, 0, 444, 447, 1, 0, 0, 0, 445, 443, 1, 0,
		0, 0, 445, 446, 1, 0, 0, 0, 446, 448, 1, 0, 0, 0, 447, 445, 1, 0, 0, 0,
		448, 457, 5, 27, 0, 0, 449, 457, 5, 2, 0, 0, 450, 457, 5, 3, 0, 0, 451,
		452, 5, 39, 0, 0, 452, 453, 5, 40, 0, 0, 453, 457, 3, 76, 38, 0, 454, 457,
		5, 1, 0, 0, 455, 457, 5, 56, 0, 0, 456, 440, 1, 0, 0, 0, 456, 449, 1, 0,
		0, 0, 456, 450, 1, 0, 0, 0, 456, 451, 1, 0, 0, 0, 456, 454, 1, 0, 0, 0,
		456, 455, 1, 0, 0, 0, 457, 77, 1, 0, 0, 0, 458, 459, 7, 3, 0, 0, 459, 79,
		1, 0, 0, 0, 460, 461, 7, 4, 0, 0, 461, 81, 1, 0, 0, 0, 47, 85, 95, 106,
		113, 120, 126, 132, 144, 153, 164, 172, 176, 185, 192, 200, 204, 213, 220,
		232, 248, 257, 261, 268, 282, 304, 313, 316, 329, 336, 344, 349, 353, 358,
		365, 374, 382, 385, 392, 397, 405, 408, 416, 419, 435, 437, 445, 456,
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
	staticData := &PromptDSLParserParserStaticData
	staticData.once.Do(promptdslparserParserInit)
}

// NewPromptDSLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPromptDSLParser(input antlr.TokenStream) *PromptDSLParser {
	PromptDSLParserInit()
	this := new(PromptDSLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PromptDSLParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PromptDSLParser.g4"

	return this
}

// PromptDSLParser tokens.
const (
	PromptDSLParserEOF              = antlr.TokenEOF
	PromptDSLParserSTRING_TYPE      = 1
	PromptDSLParserFLOAT_TYPE       = 2
	PromptDSLParserINT_TYPE         = 3
	PromptDSLParserPROMPT           = 4
	PromptDSLParserPARAMS           = 5
	PromptDSLParserSYSTEM           = 6
	PromptDSLParserUSER             = 7
	PromptDSLParserNOTE             = 8
	PromptDSLParserINPUT            = 9
	PromptDSLParserOUTPUT           = 10
	PromptDSLParserFORMAT           = 11
	PromptDSLParserTYPE             = 12
	PromptDSLParserSTRUCT           = 13
	PromptDSLParserBEFORE           = 14
	PromptDSLParserSCHEMA           = 15
	PromptDSLParserPARSE            = 16
	PromptDSLParserJSONFIX          = 17
	PromptDSLParserMARKDOWN         = 18
	PromptDSLParserIF               = 19
	PromptDSLParserELSE             = 20
	PromptDSLParserOUTPUTSPEC       = 21
	PromptDSLParserFOR              = 22
	PromptDSLParserFIX              = 23
	PromptDSLParserAFTER            = 24
	PromptDSLParserARRAY_OUTPUTSPEC = 25
	PromptDSLParserLBRACE           = 26
	PromptDSLParserRBRACE           = 27
	PromptDSLParserLPAREN           = 28
	PromptDSLParserRPAREN           = 29
	PromptDSLParserCOLON            = 30
	PromptDSLParserEQUAL            = 31
	PromptDSLParserCOMMA            = 32
	PromptDSLParserDOT              = 33
	PromptDSLParserEQEQ             = 34
	PromptDSLParserNOTEQ            = 35
	PromptDSLParserAT               = 36
	PromptDSLParserMD               = 37
	PromptDSLParserJSON             = 38
	PromptDSLParserLBRACK           = 39
	PromptDSLParserRBRACK           = 40
	PromptDSLParserINCREMENT        = 41
	PromptDSLParserDECREMENT        = 42
	PromptDSLParserMINUS            = 43
	PromptDSLParserSTAR             = 44
	PromptDSLParserSLASH            = 45
	PromptDSLParserMOD              = 46
	PromptDSLParserPLUSEQ           = 47
	PromptDSLParserMINUSEQ          = 48
	PromptDSLParserMULTEQ           = 49
	PromptDSLParserDIVEQ            = 50
	PromptDSLParserMODEQ            = 51
	PromptDSLParserLT               = 52
	PromptDSLParserLTE              = 53
	PromptDSLParserGT               = 54
	PromptDSLParserGTE              = 55
	PromptDSLParserID               = 56
	PromptDSLParserSTRING           = 57
	PromptDSLParserNUMBER           = 58
	PromptDSLParserBOOL             = 59
	PromptDSLParserPIPE             = 60
	PromptDSLParserSEMI             = 61
	PromptDSLParserPLUS             = 62
	PromptDSLParserWS               = 63
	PromptDSLParserLINE_COMMENT     = 64
	PromptDSLParserBLOCK_COMMENT    = 65
	PromptDSLParserCODE_STRING      = 66
	PromptDSLParserCODE_TEXT        = 67
)

// PromptDSLParser rules.
const (
	PromptDSLParserRULE_promptFile        = 0
	PromptDSLParserRULE_promptDef         = 1
	PromptDSLParserRULE_promptBlock       = 2
	PromptDSLParserRULE_inputSection      = 3
	PromptDSLParserRULE_outputSection     = 4
	PromptDSLParserRULE_outputStruct      = 5
	PromptDSLParserRULE_outputMarkdown    = 6
	PromptDSLParserRULE_beforeSection     = 7
	PromptDSLParserRULE_beforeContent     = 8
	PromptDSLParserRULE_varDef            = 9
	PromptDSLParserRULE_systemSection     = 10
	PromptDSLParserRULE_sysContent        = 11
	PromptDSLParserRULE_userSection       = 12
	PromptDSLParserRULE_userContent       = 13
	PromptDSLParserRULE_moduleDef         = 14
	PromptDSLParserRULE_moduleContent     = 15
	PromptDSLParserRULE_thencontent       = 16
	PromptDSLParserRULE_elsecontent       = 17
	PromptDSLParserRULE_forcontent        = 18
	PromptDSLParserRULE_ifStatement       = 19
	PromptDSLParserRULE_condition         = 20
	PromptDSLParserRULE_forStatement      = 21
	PromptDSLParserRULE_assignExpr        = 22
	PromptDSLParserRULE_updateExpr        = 23
	PromptDSLParserRULE_dslCallExpr       = 24
	PromptDSLParserRULE_expr              = 25
	PromptDSLParserRULE_fieldDef          = 26
	PromptDSLParserRULE_textLine          = 27
	PromptDSLParserRULE_paramPath         = 28
	PromptDSLParserRULE_structDef         = 29
	PromptDSLParserRULE_annotation        = 30
	PromptDSLParserRULE_annotationArgs    = 31
	PromptDSLParserRULE_annotationValue   = 32
	PromptDSLParserRULE_arrayLiteral      = 33
	PromptDSLParserRULE_defaultAnnotation = 34
	PromptDSLParserRULE_fixSection        = 35
	PromptDSLParserRULE_afterSection      = 36
	PromptDSLParserRULE_codeBlockContent  = 37
	PromptDSLParserRULE_type              = 38
	PromptDSLParserRULE_value             = 39
	PromptDSLParserRULE_formatType        = 40
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
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterPromptFile(s)
	}
}

func (s *PromptFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitPromptFile(s)
	}
}

func (p *PromptDSLParser) PromptFile() (localctx IPromptFileContext) {
	localctx = NewPromptFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PromptDSLParserRULE_promptFile)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserPROMPT {
		{
			p.SetState(82)
			p.PromptDef()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
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
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
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

func (s *PromptDefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *PromptDefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
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
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterPromptDef(s)
	}
}

func (s *PromptDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitPromptDef(s)
	}
}

func (p *PromptDSLParser) PromptDef() (localctx IPromptDefContext) {
	localctx = NewPromptDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PromptDSLParserRULE_promptDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(PromptDSLParserPROMPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72057662782572224) != 0) {
		{
			p.SetState(92)
			p.PromptBlock()
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(97)
		p.Match(PromptDSLParserRBRACE)
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
	InputSection() IInputSectionContext
	OutputSection() IOutputSectionContext
	SystemSection() ISystemSectionContext
	UserSection() IUserSectionContext
	AfterSection() IAfterSectionContext
	FixSection() IFixSectionContext
	ModuleDef() IModuleDefContext

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

func (s *PromptBlockContext) InputSection() IInputSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputSectionContext)
}

func (s *PromptBlockContext) OutputSection() IOutputSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutputSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutputSectionContext)
}

func (s *PromptBlockContext) SystemSection() ISystemSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISystemSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISystemSectionContext)
}

func (s *PromptBlockContext) UserSection() IUserSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserSectionContext)
}

func (s *PromptBlockContext) AfterSection() IAfterSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAfterSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAfterSectionContext)
}

func (s *PromptBlockContext) FixSection() IFixSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFixSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFixSectionContext)
}

func (s *PromptBlockContext) ModuleDef() IModuleDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModuleDefContext)
}

func (s *PromptBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PromptBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PromptBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterPromptBlock(s)
	}
}

func (s *PromptBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitPromptBlock(s)
	}
}

func (p *PromptDSLParser) PromptBlock() (localctx IPromptBlockContext) {
	localctx = NewPromptBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PromptDSLParserRULE_promptBlock)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINPUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.InputSection()
		}

	case PromptDSLParserOUTPUT, PromptDSLParserAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.OutputSection()
		}

	case PromptDSLParserSYSTEM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(101)
			p.SystemSection()
		}

	case PromptDSLParserUSER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(102)
			p.UserSection()
		}

	case PromptDSLParserAFTER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(103)
			p.AfterSection()
		}

	case PromptDSLParserFIX:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(104)
			p.FixSection()
		}

	case PromptDSLParserID:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(105)
			p.ModuleDef()
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

// IInputSectionContext is an interface to support dynamic dispatch.
type IInputSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INPUT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext

	// IsInputSectionContext differentiates from other interfaces.
	IsInputSectionContext()
}

type InputSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputSectionContext() *InputSectionContext {
	var p = new(InputSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_inputSection
	return p
}

func InitEmptyInputSectionContext(p *InputSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_inputSection
}

func (*InputSectionContext) IsInputSectionContext() {}

func NewInputSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputSectionContext {
	var p = new(InputSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_inputSection

	return p
}

func (s *InputSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputSectionContext) INPUT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserINPUT, 0)
}

func (s *InputSectionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *InputSectionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
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

func (s *InputSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterInputSection(s)
	}
}

func (s *InputSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitInputSection(s)
	}
}

func (p *PromptDSLParser) InputSection() (localctx IInputSectionContext) {
	localctx = NewInputSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PromptDSLParserRULE_inputSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(PromptDSLParserINPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(110)
			p.FieldDef()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
		p.Match(PromptDSLParserRBRACE)
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

// IOutputSectionContext is an interface to support dynamic dispatch.
type IOutputSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OUTPUT() antlr.TerminalNode
	OutputStruct() IOutputStructContext
	OutputMarkdown() IOutputMarkdownContext
	AllDefaultAnnotation() []IDefaultAnnotationContext
	DefaultAnnotation(i int) IDefaultAnnotationContext

	// IsOutputSectionContext differentiates from other interfaces.
	IsOutputSectionContext()
}

type OutputSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputSectionContext() *OutputSectionContext {
	var p = new(OutputSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputSection
	return p
}

func InitEmptyOutputSectionContext(p *OutputSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputSection
}

func (*OutputSectionContext) IsOutputSectionContext() {}

func NewOutputSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputSectionContext {
	var p = new(OutputSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_outputSection

	return p
}

func (s *OutputSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputSectionContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserOUTPUT, 0)
}

func (s *OutputSectionContext) OutputStruct() IOutputStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutputStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutputStructContext)
}

func (s *OutputSectionContext) OutputMarkdown() IOutputMarkdownContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutputMarkdownContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutputMarkdownContext)
}

func (s *OutputSectionContext) AllDefaultAnnotation() []IDefaultAnnotationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefaultAnnotationContext); ok {
			len++
		}
	}

	tst := make([]IDefaultAnnotationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefaultAnnotationContext); ok {
			tst[i] = t.(IDefaultAnnotationContext)
			i++
		}
	}

	return tst
}

func (s *OutputSectionContext) DefaultAnnotation(i int) IDefaultAnnotationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultAnnotationContext); ok {
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

	return t.(IDefaultAnnotationContext)
}

func (s *OutputSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterOutputSection(s)
	}
}

func (s *OutputSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitOutputSection(s)
	}
}

func (p *PromptDSLParser) OutputSection() (localctx IOutputSectionContext) {
	localctx = NewOutputSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PromptDSLParserRULE_outputSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserAT {
		{
			p.SetState(117)
			p.DefaultAnnotation()
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(123)
		p.Match(PromptDSLParserOUTPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserLBRACE:
		{
			p.SetState(124)
			p.OutputStruct()
		}

	case PromptDSLParserCOLON:
		{
			p.SetState(125)
			p.OutputMarkdown()
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

// IOutputStructContext is an interface to support dynamic dispatch.
type IOutputStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext

	// IsOutputStructContext differentiates from other interfaces.
	IsOutputStructContext()
}

type OutputStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputStructContext() *OutputStructContext {
	var p = new(OutputStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputStruct
	return p
}

func InitEmptyOutputStructContext(p *OutputStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputStruct
}

func (*OutputStructContext) IsOutputStructContext() {}

func NewOutputStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputStructContext {
	var p = new(OutputStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_outputStruct

	return p
}

func (s *OutputStructContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputStructContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *OutputStructContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *OutputStructContext) AllFieldDef() []IFieldDefContext {
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

func (s *OutputStructContext) FieldDef(i int) IFieldDefContext {
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

func (s *OutputStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterOutputStruct(s)
	}
}

func (s *OutputStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitOutputStruct(s)
	}
}

func (p *PromptDSLParser) OutputStruct() (localctx IOutputStructContext) {
	localctx = NewOutputStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PromptDSLParserRULE_outputStruct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(129)
			p.FieldDef()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(134)
		p.Match(PromptDSLParserRBRACE)
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

// IOutputMarkdownContext is an interface to support dynamic dispatch.
type IOutputMarkdownContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	MARKDOWN() antlr.TerminalNode

	// IsOutputMarkdownContext differentiates from other interfaces.
	IsOutputMarkdownContext()
}

type OutputMarkdownContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputMarkdownContext() *OutputMarkdownContext {
	var p = new(OutputMarkdownContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputMarkdown
	return p
}

func InitEmptyOutputMarkdownContext(p *OutputMarkdownContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputMarkdown
}

func (*OutputMarkdownContext) IsOutputMarkdownContext() {}

func NewOutputMarkdownContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputMarkdownContext {
	var p = new(OutputMarkdownContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_outputMarkdown

	return p
}

func (s *OutputMarkdownContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputMarkdownContext) COLON() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOLON, 0)
}

func (s *OutputMarkdownContext) MARKDOWN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMARKDOWN, 0)
}

func (s *OutputMarkdownContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputMarkdownContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputMarkdownContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterOutputMarkdown(s)
	}
}

func (s *OutputMarkdownContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitOutputMarkdown(s)
	}
}

func (p *PromptDSLParser) OutputMarkdown() (localctx IOutputMarkdownContext) {
	localctx = NewOutputMarkdownContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PromptDSLParserRULE_outputMarkdown)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(PromptDSLParserMARKDOWN)
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

// IBeforeSectionContext is an interface to support dynamic dispatch.
type IBeforeSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BEFORE() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllBeforeContent() []IBeforeContentContext
	BeforeContent(i int) IBeforeContentContext

	// IsBeforeSectionContext differentiates from other interfaces.
	IsBeforeSectionContext()
}

type BeforeSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeforeSectionContext() *BeforeSectionContext {
	var p = new(BeforeSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_beforeSection
	return p
}

func InitEmptyBeforeSectionContext(p *BeforeSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_beforeSection
}

func (*BeforeSectionContext) IsBeforeSectionContext() {}

func NewBeforeSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeforeSectionContext {
	var p = new(BeforeSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_beforeSection

	return p
}

func (s *BeforeSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *BeforeSectionContext) BEFORE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserBEFORE, 0)
}

func (s *BeforeSectionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *BeforeSectionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *BeforeSectionContext) AllBeforeContent() []IBeforeContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBeforeContentContext); ok {
			len++
		}
	}

	tst := make([]IBeforeContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBeforeContentContext); ok {
			tst[i] = t.(IBeforeContentContext)
			i++
		}
	}

	return tst
}

func (s *BeforeSectionContext) BeforeContent(i int) IBeforeContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBeforeContentContext); ok {
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

	return t.(IBeforeContentContext)
}

func (s *BeforeSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeforeSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeforeSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterBeforeSection(s)
	}
}

func (s *BeforeSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitBeforeSection(s)
	}
}

func (p *PromptDSLParser) BeforeSection() (localctx IBeforeSectionContext) {
	localctx = NewBeforeSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PromptDSLParserRULE_beforeSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(PromptDSLParserBEFORE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&38139859344852003) != 0 {
		{
			p.SetState(141)
			p.BeforeContent()
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(147)
		p.Match(PromptDSLParserRBRACE)
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

// IBeforeContentContext is an interface to support dynamic dispatch.
type IBeforeContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDef() IVarDefContext
	Expr() IExprContext
	IfStatement() IIfStatementContext
	TextLine() ITextLineContext

	// IsBeforeContentContext differentiates from other interfaces.
	IsBeforeContentContext()
}

type BeforeContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeforeContentContext() *BeforeContentContext {
	var p = new(BeforeContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_beforeContent
	return p
}

func InitEmptyBeforeContentContext(p *BeforeContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_beforeContent
}

func (*BeforeContentContext) IsBeforeContentContext() {}

func NewBeforeContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeforeContentContext {
	var p = new(BeforeContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_beforeContent

	return p
}

func (s *BeforeContentContext) GetParser() antlr.Parser { return s.parser }

func (s *BeforeContentContext) VarDef() IVarDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDefContext)
}

func (s *BeforeContentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BeforeContentContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *BeforeContentContext) TextLine() ITextLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextLineContext)
}

func (s *BeforeContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeforeContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeforeContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterBeforeContent(s)
	}
}

func (s *BeforeContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitBeforeContent(s)
	}
}

func (p *PromptDSLParser) BeforeContent() (localctx IBeforeContentContext) {
	localctx = NewBeforeContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PromptDSLParserRULE_beforeContent)
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.VarDef()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.expr(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.IfStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(152)
			p.TextLine()
		}

	case antlr.ATNInvalidAltNumber:
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

// IVarDefContext is an interface to support dynamic dispatch.
type IVarDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	Expr() IExprContext

	// IsVarDefContext differentiates from other interfaces.
	IsVarDefContext()
}

type VarDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDefContext() *VarDefContext {
	var p = new(VarDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_varDef
	return p
}

func InitEmptyVarDefContext(p *VarDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_varDef
}

func (*VarDefContext) IsVarDefContext() {}

func NewVarDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDefContext {
	var p = new(VarDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_varDef

	return p
}

func (s *VarDefContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDefContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *VarDefContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserEQUAL, 0)
}

func (s *VarDefContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterVarDef(s)
	}
}

func (s *VarDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitVarDef(s)
	}
}

func (p *PromptDSLParser) VarDef() (localctx IVarDefContext) {
	localctx = NewVarDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PromptDSLParserRULE_varDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(PromptDSLParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.expr(0)
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

// ISystemSectionContext is an interface to support dynamic dispatch.
type ISystemSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYSTEM() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllSysContent() []ISysContentContext
	SysContent(i int) ISysContentContext

	// IsSystemSectionContext differentiates from other interfaces.
	IsSystemSectionContext()
}

type SystemSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySystemSectionContext() *SystemSectionContext {
	var p = new(SystemSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_systemSection
	return p
}

func InitEmptySystemSectionContext(p *SystemSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_systemSection
}

func (*SystemSectionContext) IsSystemSectionContext() {}

func NewSystemSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SystemSectionContext {
	var p = new(SystemSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_systemSection

	return p
}

func (s *SystemSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SystemSectionContext) SYSTEM() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSYSTEM, 0)
}

func (s *SystemSectionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *SystemSectionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *SystemSectionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserID)
}

func (s *SystemSectionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, i)
}

func (s *SystemSectionContext) AllSysContent() []ISysContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISysContentContext); ok {
			len++
		}
	}

	tst := make([]ISysContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISysContentContext); ok {
			tst[i] = t.(ISysContentContext)
			i++
		}
	}

	return tst
}

func (s *SystemSectionContext) SysContent(i int) ISysContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISysContentContext); ok {
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

	return t.(ISysContentContext)
}

func (s *SystemSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SystemSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SystemSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterSystemSection(s)
	}
}

func (s *SystemSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitSystemSection(s)
	}
}

func (p *PromptDSLParser) SystemSection() (localctx ISystemSectionContext) {
	localctx = NewSystemSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PromptDSLParserRULE_systemSection)
	var _la int

	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.Match(PromptDSLParserSYSTEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PromptDSLParserID {
			{
				p.SetState(161)
				p.Match(PromptDSLParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(166)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Match(PromptDSLParserSYSTEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&38139859344929827) != 0) {
			{
				p.SetState(169)
				p.SysContent()
			}

			p.SetState(172)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(174)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// ISysContentContext is an interface to support dynamic dispatch.
type ISysContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfStatement() IIfStatementContext
	ParamPath() IParamPathContext
	ForStatement() IForStatementContext
	ARRAY_OUTPUTSPEC() antlr.TerminalNode
	OUTPUTSPEC() antlr.TerminalNode
	Expr() IExprContext
	TextLine() ITextLineContext

	// IsSysContentContext differentiates from other interfaces.
	IsSysContentContext()
}

type SysContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySysContentContext() *SysContentContext {
	var p = new(SysContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_sysContent
	return p
}

func InitEmptySysContentContext(p *SysContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_sysContent
}

func (*SysContentContext) IsSysContentContext() {}

func NewSysContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SysContentContext {
	var p = new(SysContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_sysContent

	return p
}

func (s *SysContentContext) GetParser() antlr.Parser { return s.parser }

func (s *SysContentContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *SysContentContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *SysContentContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *SysContentContext) ARRAY_OUTPUTSPEC() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserARRAY_OUTPUTSPEC, 0)
}

func (s *SysContentContext) OUTPUTSPEC() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserOUTPUTSPEC, 0)
}

func (s *SysContentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SysContentContext) TextLine() ITextLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextLineContext)
}

func (s *SysContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SysContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SysContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterSysContent(s)
	}
}

func (s *SysContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitSysContent(s)
	}
}

func (p *PromptDSLParser) SysContent() (localctx ISysContentContext) {
	localctx = NewSysContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PromptDSLParserRULE_sysContent)
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(178)
			p.IfStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.ParamPath()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(180)
			p.ForStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(181)
			p.Match(PromptDSLParserARRAY_OUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(182)
			p.Match(PromptDSLParserOUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(183)
			p.expr(0)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(184)
			p.TextLine()
		}

	case antlr.ATNInvalidAltNumber:
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

// IUserSectionContext is an interface to support dynamic dispatch.
type IUserSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	USER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllUserContent() []IUserContentContext
	UserContent(i int) IUserContentContext

	// IsUserSectionContext differentiates from other interfaces.
	IsUserSectionContext()
}

type UserSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserSectionContext() *UserSectionContext {
	var p = new(UserSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_userSection
	return p
}

func InitEmptyUserSectionContext(p *UserSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_userSection
}

func (*UserSectionContext) IsUserSectionContext() {}

func NewUserSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserSectionContext {
	var p = new(UserSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_userSection

	return p
}

func (s *UserSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *UserSectionContext) USER() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserUSER, 0)
}

func (s *UserSectionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *UserSectionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *UserSectionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserID)
}

func (s *UserSectionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, i)
}

func (s *UserSectionContext) AllUserContent() []IUserContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUserContentContext); ok {
			len++
		}
	}

	tst := make([]IUserContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUserContentContext); ok {
			tst[i] = t.(IUserContentContext)
			i++
		}
	}

	return tst
}

func (s *UserSectionContext) UserContent(i int) IUserContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContentContext); ok {
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

	return t.(IUserContentContext)
}

func (s *UserSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterUserSection(s)
	}
}

func (s *UserSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitUserSection(s)
	}
}

func (p *PromptDSLParser) UserSection() (localctx IUserSectionContext) {
	localctx = NewUserSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PromptDSLParserRULE_userSection)
	var _la int

	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Match(PromptDSLParserUSER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PromptDSLParserID {
			{
				p.SetState(189)
				p.Match(PromptDSLParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(194)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.Match(PromptDSLParserUSER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&38139859344929827) != 0) {
			{
				p.SetState(197)
				p.UserContent()
			}

			p.SetState(200)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(202)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IUserContentContext is an interface to support dynamic dispatch.
type IUserContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfStatement() IIfStatementContext
	ParamPath() IParamPathContext
	ForStatement() IForStatementContext
	ARRAY_OUTPUTSPEC() antlr.TerminalNode
	OUTPUTSPEC() antlr.TerminalNode
	Expr() IExprContext
	TextLine() ITextLineContext

	// IsUserContentContext differentiates from other interfaces.
	IsUserContentContext()
}

type UserContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserContentContext() *UserContentContext {
	var p = new(UserContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_userContent
	return p
}

func InitEmptyUserContentContext(p *UserContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_userContent
}

func (*UserContentContext) IsUserContentContext() {}

func NewUserContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserContentContext {
	var p = new(UserContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_userContent

	return p
}

func (s *UserContentContext) GetParser() antlr.Parser { return s.parser }

func (s *UserContentContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *UserContentContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *UserContentContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *UserContentContext) ARRAY_OUTPUTSPEC() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserARRAY_OUTPUTSPEC, 0)
}

func (s *UserContentContext) OUTPUTSPEC() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserOUTPUTSPEC, 0)
}

func (s *UserContentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UserContentContext) TextLine() ITextLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextLineContext)
}

func (s *UserContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterUserContent(s)
	}
}

func (s *UserContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitUserContent(s)
	}
}

func (p *PromptDSLParser) UserContent() (localctx IUserContentContext) {
	localctx = NewUserContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PromptDSLParserRULE_userContent)
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.IfStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
			p.ParamPath()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(208)
			p.ForStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(209)
			p.Match(PromptDSLParserARRAY_OUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(210)
			p.Match(PromptDSLParserOUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(211)
			p.expr(0)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(212)
			p.TextLine()
		}

	case antlr.ATNInvalidAltNumber:
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

// IModuleDefContext is an interface to support dynamic dispatch.
type IModuleDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllModuleContent() []IModuleContentContext
	ModuleContent(i int) IModuleContentContext

	// IsModuleDefContext differentiates from other interfaces.
	IsModuleDefContext()
}

type ModuleDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleDefContext() *ModuleDefContext {
	var p = new(ModuleDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_moduleDef
	return p
}

func InitEmptyModuleDefContext(p *ModuleDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_moduleDef
}

func (*ModuleDefContext) IsModuleDefContext() {}

func NewModuleDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleDefContext {
	var p = new(ModuleDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_moduleDef

	return p
}

func (s *ModuleDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleDefContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *ModuleDefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *ModuleDefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *ModuleDefContext) AllModuleContent() []IModuleContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IModuleContentContext); ok {
			len++
		}
	}

	tst := make([]IModuleContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IModuleContentContext); ok {
			tst[i] = t.(IModuleContentContext)
			i++
		}
	}

	return tst
}

func (s *ModuleDefContext) ModuleContent(i int) IModuleContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleContentContext); ok {
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

	return t.(IModuleContentContext)
}

func (s *ModuleDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterModuleDef(s)
	}
}

func (s *ModuleDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitModuleDef(s)
	}
}

func (p *PromptDSLParser) ModuleDef() (localctx IModuleDefContext) {
	localctx = NewModuleDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PromptDSLParserRULE_moduleDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&38139859344929827) != 0 {
		{
			p.SetState(217)
			p.ModuleContent()
		}

		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(223)
		p.Match(PromptDSLParserRBRACE)
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

// IModuleContentContext is an interface to support dynamic dispatch.
type IModuleContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfStatement() IIfStatementContext
	ParamPath() IParamPathContext
	ForStatement() IForStatementContext
	ARRAY_OUTPUTSPEC() antlr.TerminalNode
	OUTPUTSPEC() antlr.TerminalNode
	Expr() IExprContext
	TextLine() ITextLineContext

	// IsModuleContentContext differentiates from other interfaces.
	IsModuleContentContext()
}

type ModuleContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContentContext() *ModuleContentContext {
	var p = new(ModuleContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_moduleContent
	return p
}

func InitEmptyModuleContentContext(p *ModuleContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_moduleContent
}

func (*ModuleContentContext) IsModuleContentContext() {}

func NewModuleContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContentContext {
	var p = new(ModuleContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_moduleContent

	return p
}

func (s *ModuleContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContentContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *ModuleContentContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *ModuleContentContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *ModuleContentContext) ARRAY_OUTPUTSPEC() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserARRAY_OUTPUTSPEC, 0)
}

func (s *ModuleContentContext) OUTPUTSPEC() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserOUTPUTSPEC, 0)
}

func (s *ModuleContentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ModuleContentContext) TextLine() ITextLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextLineContext)
}

func (s *ModuleContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterModuleContent(s)
	}
}

func (s *ModuleContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitModuleContent(s)
	}
}

func (p *PromptDSLParser) ModuleContent() (localctx IModuleContentContext) {
	localctx = NewModuleContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PromptDSLParserRULE_moduleContent)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(225)
			p.IfStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.ParamPath()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(227)
			p.ForStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(228)
			p.Match(PromptDSLParserARRAY_OUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(229)
			p.Match(PromptDSLParserOUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(230)
			p.expr(0)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(231)
			p.TextLine()
		}

	case antlr.ATNInvalidAltNumber:
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

// IThencontentContext is an interface to support dynamic dispatch.
type IThencontentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UserContent() IUserContentContext

	// IsThencontentContext differentiates from other interfaces.
	IsThencontentContext()
}

type ThencontentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThencontentContext() *ThencontentContext {
	var p = new(ThencontentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_thencontent
	return p
}

func InitEmptyThencontentContext(p *ThencontentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_thencontent
}

func (*ThencontentContext) IsThencontentContext() {}

func NewThencontentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThencontentContext {
	var p = new(ThencontentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_thencontent

	return p
}

func (s *ThencontentContext) GetParser() antlr.Parser { return s.parser }

func (s *ThencontentContext) UserContent() IUserContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContentContext)
}

func (s *ThencontentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThencontentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThencontentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterThencontent(s)
	}
}

func (s *ThencontentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitThencontent(s)
	}
}

func (p *PromptDSLParser) Thencontent() (localctx IThencontentContext) {
	localctx = NewThencontentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PromptDSLParserRULE_thencontent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.UserContent()
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

// IElsecontentContext is an interface to support dynamic dispatch.
type IElsecontentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UserContent() IUserContentContext

	// IsElsecontentContext differentiates from other interfaces.
	IsElsecontentContext()
}

type ElsecontentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElsecontentContext() *ElsecontentContext {
	var p = new(ElsecontentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_elsecontent
	return p
}

func InitEmptyElsecontentContext(p *ElsecontentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_elsecontent
}

func (*ElsecontentContext) IsElsecontentContext() {}

func NewElsecontentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElsecontentContext {
	var p = new(ElsecontentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_elsecontent

	return p
}

func (s *ElsecontentContext) GetParser() antlr.Parser { return s.parser }

func (s *ElsecontentContext) UserContent() IUserContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContentContext)
}

func (s *ElsecontentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElsecontentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElsecontentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterElsecontent(s)
	}
}

func (s *ElsecontentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitElsecontent(s)
	}
}

func (p *PromptDSLParser) Elsecontent() (localctx IElsecontentContext) {
	localctx = NewElsecontentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PromptDSLParserRULE_elsecontent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.UserContent()
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

// IForcontentContext is an interface to support dynamic dispatch.
type IForcontentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UserContent() IUserContentContext

	// IsForcontentContext differentiates from other interfaces.
	IsForcontentContext()
}

type ForcontentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForcontentContext() *ForcontentContext {
	var p = new(ForcontentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_forcontent
	return p
}

func InitEmptyForcontentContext(p *ForcontentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_forcontent
}

func (*ForcontentContext) IsForcontentContext() {}

func NewForcontentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForcontentContext {
	var p = new(ForcontentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_forcontent

	return p
}

func (s *ForcontentContext) GetParser() antlr.Parser { return s.parser }

func (s *ForcontentContext) UserContent() IUserContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContentContext)
}

func (s *ForcontentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForcontentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForcontentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterForcontent(s)
	}
}

func (s *ForcontentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitForcontent(s)
	}
}

func (p *PromptDSLParser) Forcontent() (localctx IForcontentContext) {
	localctx = NewForcontentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PromptDSLParserRULE_forcontent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.UserContent()
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Condition() IConditionContext
	RPAREN() antlr.TerminalNode
	AllLBRACE() []antlr.TerminalNode
	LBRACE(i int) antlr.TerminalNode
	AllRBRACE() []antlr.TerminalNode
	RBRACE(i int) antlr.TerminalNode
	AllThencontent() []IThencontentContext
	Thencontent(i int) IThencontentContext
	ELSE() antlr.TerminalNode
	AllElsecontent() []IElsecontentContext
	Elsecontent(i int) IElsecontentContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserIF, 0)
}

func (s *IfStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLPAREN, 0)
}

func (s *IfStatementContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *IfStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRPAREN, 0)
}

func (s *IfStatementContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserLBRACE)
}

func (s *IfStatementContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, i)
}

func (s *IfStatementContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserRBRACE)
}

func (s *IfStatementContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, i)
}

func (s *IfStatementContext) AllThencontent() []IThencontentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IThencontentContext); ok {
			len++
		}
	}

	tst := make([]IThencontentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IThencontentContext); ok {
			tst[i] = t.(IThencontentContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Thencontent(i int) IThencontentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThencontentContext); ok {
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

	return t.(IThencontentContext)
}

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserELSE, 0)
}

func (s *IfStatementContext) AllElsecontent() []IElsecontentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElsecontentContext); ok {
			len++
		}
	}

	tst := make([]IElsecontentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElsecontentContext); ok {
			tst[i] = t.(IElsecontentContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Elsecontent(i int) IElsecontentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElsecontentContext); ok {
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

	return t.(IElsecontentContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *PromptDSLParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PromptDSLParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(PromptDSLParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Match(PromptDSLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Condition()
	}
	{
		p.SetState(243)
		p.Match(PromptDSLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&38139859344929827) != 0 {
		{
			p.SetState(245)
			p.Thencontent()
		}

		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(251)
		p.Match(PromptDSLParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserELSE {
		{
			p.SetState(252)
			p.Match(PromptDSLParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&38139859344929827) != 0 {
			{
				p.SetState(254)
				p.Elsecontent()
			}

			p.SetState(259)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(260)
			p.Match(PromptDSLParserRBRACE)
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

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLhs returns the lhs rule contexts.
	GetLhs() IExprContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExprContext

	// GetSingle returns the single rule contexts.
	GetSingle() IExprContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(IExprContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExprContext)

	// SetSingle sets the single rule contexts.
	SetSingle(IExprContext)

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	EQEQ() antlr.TerminalNode
	NOTEQ() antlr.TerminalNode
	LT() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GTE() antlr.TerminalNode

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    IExprContext
	op     antlr.Token
	rhs    IExprContext
	single IExprContext
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) GetOp() antlr.Token { return s.op }

func (s *ConditionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ConditionContext) GetLhs() IExprContext { return s.lhs }

func (s *ConditionContext) GetRhs() IExprContext { return s.rhs }

func (s *ConditionContext) GetSingle() IExprContext { return s.single }

func (s *ConditionContext) SetLhs(v IExprContext) { s.lhs = v }

func (s *ConditionContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *ConditionContext) SetSingle(v IExprContext) { s.single = v }

func (s *ConditionContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ConditionContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ConditionContext) EQEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserEQEQ, 0)
}

func (s *ConditionContext) NOTEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserNOTEQ, 0)
}

func (s *ConditionContext) LT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLT, 0)
}

func (s *ConditionContext) LTE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLTE, 0)
}

func (s *ConditionContext) GT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserGT, 0)
}

func (s *ConditionContext) GTE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserGTE, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *PromptDSLParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PromptDSLParserRULE_condition)
	var _la int

	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)

			var _x = p.expr(0)

			localctx.(*ConditionContext).lhs = _x
		}
		{
			p.SetState(264)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ConditionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67554045950164992) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ConditionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(265)

			var _x = p.expr(0)

			localctx.(*ConditionContext).rhs = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)

			var _x = p.expr(0)

			localctx.(*ConditionContext).single = _x
		}

	case antlr.ATNInvalidAltNumber:
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

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInit returns the init rule contexts.
	GetInit() IAssignExprContext

	// GetUpdate returns the update rule contexts.
	GetUpdate() IUpdateExprContext

	// SetInit sets the init rule contexts.
	SetInit(IAssignExprContext)

	// SetUpdate sets the update rule contexts.
	SetUpdate(IUpdateExprContext)

	// Getter signatures
	FOR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode
	Condition() IConditionContext
	RPAREN() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AssignExpr() IAssignExprContext
	UpdateExpr() IUpdateExprContext
	AllForcontent() []IForcontentContext
	Forcontent(i int) IForcontentContext

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	init   IAssignExprContext
	update IUpdateExprContext
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) GetInit() IAssignExprContext { return s.init }

func (s *ForStatementContext) GetUpdate() IUpdateExprContext { return s.update }

func (s *ForStatementContext) SetInit(v IAssignExprContext) { s.init = v }

func (s *ForStatementContext) SetUpdate(v IUpdateExprContext) { s.update = v }

func (s *ForStatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserFOR, 0)
}

func (s *ForStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLPAREN, 0)
}

func (s *ForStatementContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserSEMI)
}

func (s *ForStatementContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSEMI, i)
}

func (s *ForStatementContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ForStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRPAREN, 0)
}

func (s *ForStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *ForStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *ForStatementContext) AssignExpr() IAssignExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignExprContext)
}

func (s *ForStatementContext) UpdateExpr() IUpdateExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdateExprContext)
}

func (s *ForStatementContext) AllForcontent() []IForcontentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForcontentContext); ok {
			len++
		}
	}

	tst := make([]IForcontentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForcontentContext); ok {
			tst[i] = t.(IForcontentContext)
			i++
		}
	}

	return tst
}

func (s *ForStatementContext) Forcontent(i int) IForcontentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForcontentContext); ok {
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

	return t.(IForcontentContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (p *PromptDSLParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PromptDSLParserRULE_forStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Match(PromptDSLParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Match(PromptDSLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)

		var _x = p.AssignExpr()

		localctx.(*ForStatementContext).init = _x
	}
	{
		p.SetState(273)
		p.Match(PromptDSLParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.Condition()
	}
	{
		p.SetState(275)
		p.Match(PromptDSLParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(276)

		var _x = p.UpdateExpr()

		localctx.(*ForStatementContext).update = _x
	}
	{
		p.SetState(277)
		p.Match(PromptDSLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&38139859344929827) != 0 {
		{
			p.SetState(279)
			p.Forcontent()
		}

		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(285)
		p.Match(PromptDSLParserRBRACE)
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

// IAssignExprContext is an interface to support dynamic dispatch.
type IAssignExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLhs returns the lhs rule contexts.
	GetLhs() IParamPathContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExprContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(IParamPathContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExprContext)

	// Getter signatures
	EQUAL() antlr.TerminalNode
	ParamPath() IParamPathContext
	Expr() IExprContext

	// IsAssignExprContext differentiates from other interfaces.
	IsAssignExprContext()
}

type AssignExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    IParamPathContext
	rhs    IExprContext
}

func NewEmptyAssignExprContext() *AssignExprContext {
	var p = new(AssignExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_assignExpr
	return p
}

func InitEmptyAssignExprContext(p *AssignExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_assignExpr
}

func (*AssignExprContext) IsAssignExprContext() {}

func NewAssignExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignExprContext {
	var p = new(AssignExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_assignExpr

	return p
}

func (s *AssignExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignExprContext) GetLhs() IParamPathContext { return s.lhs }

func (s *AssignExprContext) GetRhs() IExprContext { return s.rhs }

func (s *AssignExprContext) SetLhs(v IParamPathContext) { s.lhs = v }

func (s *AssignExprContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *AssignExprContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserEQUAL, 0)
}

func (s *AssignExprContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *AssignExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterAssignExpr(s)
	}
}

func (s *AssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitAssignExpr(s)
	}
}

func (p *PromptDSLParser) AssignExpr() (localctx IAssignExprContext) {
	localctx = NewAssignExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PromptDSLParserRULE_assignExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)

		var _x = p.ParamPath()

		localctx.(*AssignExprContext).lhs = _x
	}
	{
		p.SetState(288)
		p.Match(PromptDSLParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)

		var _x = p.expr(0)

		localctx.(*AssignExprContext).rhs = _x
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

// IUpdateExprContext is an interface to support dynamic dispatch.
type IUpdateExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParamPath() IParamPathContext
	INCREMENT() antlr.TerminalNode
	DECREMENT() antlr.TerminalNode
	PLUSEQ() antlr.TerminalNode
	Expr() IExprContext
	MINUSEQ() antlr.TerminalNode
	MULTEQ() antlr.TerminalNode
	DIVEQ() antlr.TerminalNode
	MODEQ() antlr.TerminalNode

	// IsUpdateExprContext differentiates from other interfaces.
	IsUpdateExprContext()
}

type UpdateExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateExprContext() *UpdateExprContext {
	var p = new(UpdateExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_updateExpr
	return p
}

func InitEmptyUpdateExprContext(p *UpdateExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_updateExpr
}

func (*UpdateExprContext) IsUpdateExprContext() {}

func NewUpdateExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateExprContext {
	var p = new(UpdateExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_updateExpr

	return p
}

func (s *UpdateExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateExprContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *UpdateExprContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserINCREMENT, 0)
}

func (s *UpdateExprContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDECREMENT, 0)
}

func (s *UpdateExprContext) PLUSEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserPLUSEQ, 0)
}

func (s *UpdateExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UpdateExprContext) MINUSEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMINUSEQ, 0)
}

func (s *UpdateExprContext) MULTEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMULTEQ, 0)
}

func (s *UpdateExprContext) DIVEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDIVEQ, 0)
}

func (s *UpdateExprContext) MODEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMODEQ, 0)
}

func (s *UpdateExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterUpdateExpr(s)
	}
}

func (s *UpdateExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitUpdateExpr(s)
	}
}

func (p *PromptDSLParser) UpdateExpr() (localctx IUpdateExprContext) {
	localctx = NewUpdateExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PromptDSLParserRULE_updateExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.ParamPath()
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINCREMENT:
		{
			p.SetState(292)
			p.Match(PromptDSLParserINCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserDECREMENT:
		{
			p.SetState(293)
			p.Match(PromptDSLParserDECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserPLUSEQ:
		{
			p.SetState(294)
			p.Match(PromptDSLParserPLUSEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)
			p.expr(0)
		}

	case PromptDSLParserMINUSEQ:
		{
			p.SetState(296)
			p.Match(PromptDSLParserMINUSEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)
			p.expr(0)
		}

	case PromptDSLParserMULTEQ:
		{
			p.SetState(298)
			p.Match(PromptDSLParserMULTEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.expr(0)
		}

	case PromptDSLParserDIVEQ:
		{
			p.SetState(300)
			p.Match(PromptDSLParserDIVEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.expr(0)
		}

	case PromptDSLParserMODEQ:
		{
			p.SetState(302)
			p.Match(PromptDSLParserMODEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.expr(0)
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

// IDslCallExprContext is an interface to support dynamic dispatch.
type IDslCallExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParamPath() IParamPathContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsDslCallExprContext differentiates from other interfaces.
	IsDslCallExprContext()
}

type DslCallExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDslCallExprContext() *DslCallExprContext {
	var p = new(DslCallExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_dslCallExpr
	return p
}

func InitEmptyDslCallExprContext(p *DslCallExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_dslCallExpr
}

func (*DslCallExprContext) IsDslCallExprContext() {}

func NewDslCallExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DslCallExprContext {
	var p = new(DslCallExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_dslCallExpr

	return p
}

func (s *DslCallExprContext) GetParser() antlr.Parser { return s.parser }

func (s *DslCallExprContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *DslCallExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLPAREN, 0)
}

func (s *DslCallExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRPAREN, 0)
}

func (s *DslCallExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *DslCallExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *DslCallExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserCOMMA)
}

func (s *DslCallExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOMMA, i)
}

func (s *DslCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DslCallExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DslCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterDslCallExpr(s)
	}
}

func (s *DslCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitDslCallExpr(s)
	}
}

func (p *PromptDSLParser) DslCallExpr() (localctx IDslCallExprContext) {
	localctx = NewDslCallExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PromptDSLParserRULE_dslCallExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.ParamPath()
	}
	{
		p.SetState(307)
		p.Match(PromptDSLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1080863910854149632) != 0 {
		{
			p.SetState(308)
			p.expr(0)
		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserCOMMA {
			{
				p.SetState(309)
				p.Match(PromptDSLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(310)
				p.expr(0)
			}

			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(318)
		p.Match(PromptDSLParserRPAREN)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ParamPath() IParamPathContext
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	RPAREN() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	STAR() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	MOD() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserNUMBER, 0)
}

func (s *ExprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserBOOL, 0)
}

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLPAREN, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRPAREN, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserPLUS, 0)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMINUS, 0)
}

func (s *ExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTAR, 0)
}

func (s *ExprContext) SLASH() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSLASH, 0)
}

func (s *ExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMOD, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *PromptDSLParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *PromptDSLParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, PromptDSLParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		{
			p.SetState(321)
			p.ParamPath()
		}

	case PromptDSLParserSTRING:
		{
			p.SetState(322)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserNUMBER:
		{
			p.SetState(323)
			p.Match(PromptDSLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserBOOL:
		{
			p.SetState(324)
			p.Match(PromptDSLParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserLPAREN:
		{
			p.SetState(325)
			p.Match(PromptDSLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)
			p.expr(0)
		}
		{
			p.SetState(327)
			p.Match(PromptDSLParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, PromptDSLParserRULE_expr)
			p.SetState(331)

			if !(p.Precpred(p.GetParserRuleContext(), 6)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				goto errorExit
			}
			{
				p.SetState(332)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4611817959822721024) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(333)
				p.expr(7)
			}

		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
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
	p.UnrollRecursionContexts(_parentctx)
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
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	EQUAL() antlr.TerminalNode
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

func (s *FieldDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOLON, 0)
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

func (s *FieldDefContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserEQUAL, 0)
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
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterFieldDef(s)
	}
}

func (s *FieldDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitFieldDef(s)
	}
}

func (p *PromptDSLParser) FieldDef() (localctx IFieldDefContext) {
	localctx = NewFieldDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PromptDSLParserRULE_fieldDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(341)
		p.Type_()
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserEQUAL {
		{
			p.SetState(342)
			p.Match(PromptDSLParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Value()
		}

	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserAT {
		{
			p.SetState(346)
			p.Annotation()
		}

		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSEMI {
		{
			p.SetState(352)
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

// ITextLineContext is an interface to support dynamic dispatch.
type ITextLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	LINE_COMMENT() antlr.TerminalNode
	ParamPath() IParamPathContext

	// IsTextLineContext differentiates from other interfaces.
	IsTextLineContext()
}

type TextLineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextLineContext() *TextLineContext {
	var p = new(TextLineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_textLine
	return p
}

func InitEmptyTextLineContext(p *TextLineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_textLine
}

func (*TextLineContext) IsTextLineContext() {}

func NewTextLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextLineContext {
	var p = new(TextLineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_textLine

	return p
}

func (s *TextLineContext) GetParser() antlr.Parser { return s.parser }

func (s *TextLineContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *TextLineContext) LINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLINE_COMMENT, 0)
}

func (s *TextLineContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *TextLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterTextLine(s)
	}
}

func (s *TextLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitTextLine(s)
	}
}

func (p *PromptDSLParser) TextLine() (localctx ITextLineContext) {
	localctx = NewTextLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PromptDSLParserRULE_textLine)
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(355)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserLINE_COMMENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(356)
			p.Match(PromptDSLParserLINE_COMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(357)
			p.ParamPath()
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

// IParamPathContext is an interface to support dynamic dispatch.
type IParamPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	INPUT() antlr.TerminalNode
	OUTPUT() antlr.TerminalNode
	AFTER() antlr.TerminalNode
	BEFORE() antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsParamPathContext differentiates from other interfaces.
	IsParamPathContext()
}

type ParamPathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamPathContext() *ParamPathContext {
	var p = new(ParamPathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_paramPath
	return p
}

func InitEmptyParamPathContext(p *ParamPathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_paramPath
}

func (*ParamPathContext) IsParamPathContext() {}

func NewParamPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamPathContext {
	var p = new(ParamPathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_paramPath

	return p
}

func (s *ParamPathContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamPathContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserID)
}

func (s *ParamPathContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, i)
}

func (s *ParamPathContext) INPUT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserINPUT, 0)
}

func (s *ParamPathContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserOUTPUT, 0)
}

func (s *ParamPathContext) AFTER() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserAFTER, 0)
}

func (s *ParamPathContext) BEFORE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserBEFORE, 0)
}

func (s *ParamPathContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserDOT)
}

func (s *ParamPathContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDOT, i)
}

func (s *ParamPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterParamPath(s)
	}
}

func (s *ParamPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitParamPath(s)
	}
}

func (p *PromptDSLParser) ParamPath() (localctx IParamPathContext) {
	localctx = NewParamPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PromptDSLParserRULE_paramPath)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72057594054723072) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(361)
				p.Match(PromptDSLParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(362)
				p.Match(PromptDSLParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(367)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
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

// IStructDefContext is an interface to support dynamic dispatch.
type IStructDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
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

func (s *StructDefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *StructDefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
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
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterStructDef(s)
	}
}

func (s *StructDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitStructDef(s)
	}
}

func (p *PromptDSLParser) StructDef() (localctx IStructDefContext) {
	localctx = NewStructDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PromptDSLParserRULE_structDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
		p.Match(PromptDSLParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(370)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(371)
			p.FieldDef()
		}

		p.SetState(374)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(376)
		p.Match(PromptDSLParserRBRACE)
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
	AT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
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

func (s *AnnotationContext) AT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserAT, 0)
}

func (s *AnnotationContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *AnnotationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLPAREN, 0)
}

func (s *AnnotationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRPAREN, 0)
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
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (p *PromptDSLParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PromptDSLParserRULE_annotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(PromptDSLParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserLPAREN {
		{
			p.SetState(380)
			p.Match(PromptDSLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(382)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PromptDSLParserLBRACE || _la == PromptDSLParserSTRING {
			{
				p.SetState(381)
				p.AnnotationArgs()
			}

		}
		{
			p.SetState(384)
			p.Match(PromptDSLParserRPAREN)
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

// IAnnotationArgsContext is an interface to support dynamic dispatch.
type IAnnotationArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAnnotationValue() []IAnnotationValueContext
	AnnotationValue(i int) IAnnotationValueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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

func (s *AnnotationArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserCOMMA)
}

func (s *AnnotationArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOMMA, i)
}

func (s *AnnotationArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterAnnotationArgs(s)
	}
}

func (s *AnnotationArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitAnnotationArgs(s)
	}
}

func (p *PromptDSLParser) AnnotationArgs() (localctx IAnnotationArgsContext) {
	localctx = NewAnnotationArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PromptDSLParserRULE_annotationArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.AnnotationValue()
	}
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserCOMMA {
		{
			p.SetState(388)
			p.Match(PromptDSLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(389)
			p.AnnotationValue()
		}

		p.SetState(394)
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
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterAnnotationValue(s)
	}
}

func (s *AnnotationValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitAnnotationValue(s)
	}
}

func (p *PromptDSLParser) AnnotationValue() (localctx IAnnotationValueContext) {
	localctx = NewAnnotationValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PromptDSLParserRULE_annotationValue)
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(395)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(396)
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
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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

func (s *ArrayLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *ArrayLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *ArrayLiteralContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserSTRING)
}

func (s *ArrayLiteralContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, i)
}

func (s *ArrayLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserCOMMA)
}

func (s *ArrayLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOMMA, i)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (p *PromptDSLParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PromptDSLParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSTRING {
		{
			p.SetState(400)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserCOMMA {
			{
				p.SetState(401)
				p.Match(PromptDSLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(402)
				p.Match(PromptDSLParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(407)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(410)
		p.Match(PromptDSLParserRBRACE)
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

// IDefaultAnnotationContext is an interface to support dynamic dispatch.
type IDefaultAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AnnotationArgs() IAnnotationArgsContext

	// IsDefaultAnnotationContext differentiates from other interfaces.
	IsDefaultAnnotationContext()
}

type DefaultAnnotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultAnnotationContext() *DefaultAnnotationContext {
	var p = new(DefaultAnnotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_defaultAnnotation
	return p
}

func InitEmptyDefaultAnnotationContext(p *DefaultAnnotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_defaultAnnotation
}

func (*DefaultAnnotationContext) IsDefaultAnnotationContext() {}

func NewDefaultAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultAnnotationContext {
	var p = new(DefaultAnnotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_defaultAnnotation

	return p
}

func (s *DefaultAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultAnnotationContext) AT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserAT, 0)
}

func (s *DefaultAnnotationContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *DefaultAnnotationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLPAREN, 0)
}

func (s *DefaultAnnotationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRPAREN, 0)
}

func (s *DefaultAnnotationContext) AnnotationArgs() IAnnotationArgsContext {
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

func (s *DefaultAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterDefaultAnnotation(s)
	}
}

func (s *DefaultAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitDefaultAnnotation(s)
	}
}

func (p *PromptDSLParser) DefaultAnnotation() (localctx IDefaultAnnotationContext) {
	localctx = NewDefaultAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, PromptDSLParserRULE_defaultAnnotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(412)
		p.Match(PromptDSLParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(413)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserLPAREN {
		{
			p.SetState(414)
			p.Match(PromptDSLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PromptDSLParserLBRACE || _la == PromptDSLParserSTRING {
			{
				p.SetState(415)
				p.AnnotationArgs()
			}

		}
		{
			p.SetState(418)
			p.Match(PromptDSLParserRPAREN)
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

// IFixSectionContext is an interface to support dynamic dispatch.
type IFixSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FIX() antlr.TerminalNode
	CodeBlockContent() ICodeBlockContentContext
	RBRACE() antlr.TerminalNode

	// IsFixSectionContext differentiates from other interfaces.
	IsFixSectionContext()
}

type FixSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFixSectionContext() *FixSectionContext {
	var p = new(FixSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_fixSection
	return p
}

func InitEmptyFixSectionContext(p *FixSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_fixSection
}

func (*FixSectionContext) IsFixSectionContext() {}

func NewFixSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FixSectionContext {
	var p = new(FixSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_fixSection

	return p
}

func (s *FixSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *FixSectionContext) FIX() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserFIX, 0)
}

func (s *FixSectionContext) CodeBlockContent() ICodeBlockContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContentContext)
}

func (s *FixSectionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *FixSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FixSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterFixSection(s)
	}
}

func (s *FixSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitFixSection(s)
	}
}

func (p *PromptDSLParser) FixSection() (localctx IFixSectionContext) {
	localctx = NewFixSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, PromptDSLParserRULE_fixSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.Match(PromptDSLParserFIX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(422)
		p.CodeBlockContent()
	}
	{
		p.SetState(423)
		p.Match(PromptDSLParserRBRACE)
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

// IAfterSectionContext is an interface to support dynamic dispatch.
type IAfterSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AFTER() antlr.TerminalNode
	CodeBlockContent() ICodeBlockContentContext
	RBRACE() antlr.TerminalNode

	// IsAfterSectionContext differentiates from other interfaces.
	IsAfterSectionContext()
}

type AfterSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAfterSectionContext() *AfterSectionContext {
	var p = new(AfterSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_afterSection
	return p
}

func InitEmptyAfterSectionContext(p *AfterSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_afterSection
}

func (*AfterSectionContext) IsAfterSectionContext() {}

func NewAfterSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AfterSectionContext {
	var p = new(AfterSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_afterSection

	return p
}

func (s *AfterSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *AfterSectionContext) AFTER() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserAFTER, 0)
}

func (s *AfterSectionContext) CodeBlockContent() ICodeBlockContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContentContext)
}

func (s *AfterSectionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *AfterSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AfterSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AfterSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterAfterSection(s)
	}
}

func (s *AfterSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitAfterSection(s)
	}
}

func (p *PromptDSLParser) AfterSection() (localctx IAfterSectionContext) {
	localctx = NewAfterSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, PromptDSLParserRULE_afterSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)
		p.Match(PromptDSLParserAFTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(426)
		p.CodeBlockContent()
	}
	{
		p.SetState(427)
		p.Match(PromptDSLParserRBRACE)
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

// ICodeBlockContentContext is an interface to support dynamic dispatch.
type ICodeBlockContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCODE_TEXT() []antlr.TerminalNode
	CODE_TEXT(i int) antlr.TerminalNode
	AllCODE_STRING() []antlr.TerminalNode
	CODE_STRING(i int) antlr.TerminalNode
	AllLBRACE() []antlr.TerminalNode
	LBRACE(i int) antlr.TerminalNode
	AllCodeBlockContent() []ICodeBlockContentContext
	CodeBlockContent(i int) ICodeBlockContentContext
	AllRBRACE() []antlr.TerminalNode
	RBRACE(i int) antlr.TerminalNode

	// IsCodeBlockContentContext differentiates from other interfaces.
	IsCodeBlockContentContext()
}

type CodeBlockContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeBlockContentContext() *CodeBlockContentContext {
	var p = new(CodeBlockContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_codeBlockContent
	return p
}

func InitEmptyCodeBlockContentContext(p *CodeBlockContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_codeBlockContent
}

func (*CodeBlockContentContext) IsCodeBlockContentContext() {}

func NewCodeBlockContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeBlockContentContext {
	var p = new(CodeBlockContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_codeBlockContent

	return p
}

func (s *CodeBlockContentContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeBlockContentContext) AllCODE_TEXT() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserCODE_TEXT)
}

func (s *CodeBlockContentContext) CODE_TEXT(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCODE_TEXT, i)
}

func (s *CodeBlockContentContext) AllCODE_STRING() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserCODE_STRING)
}

func (s *CodeBlockContentContext) CODE_STRING(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCODE_STRING, i)
}

func (s *CodeBlockContentContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserLBRACE)
}

func (s *CodeBlockContentContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, i)
}

func (s *CodeBlockContentContext) AllCodeBlockContent() []ICodeBlockContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICodeBlockContentContext); ok {
			len++
		}
	}

	tst := make([]ICodeBlockContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICodeBlockContentContext); ok {
			tst[i] = t.(ICodeBlockContentContext)
			i++
		}
	}

	return tst
}

func (s *CodeBlockContentContext) CodeBlockContent(i int) ICodeBlockContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContentContext); ok {
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

	return t.(ICodeBlockContentContext)
}

func (s *CodeBlockContentContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserRBRACE)
}

func (s *CodeBlockContentContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, i)
}

func (s *CodeBlockContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeBlockContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeBlockContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterCodeBlockContent(s)
	}
}

func (s *CodeBlockContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitCodeBlockContent(s)
	}
}

func (p *PromptDSLParser) CodeBlockContent() (localctx ICodeBlockContentContext) {
	localctx = NewCodeBlockContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, PromptDSLParserRULE_codeBlockContent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&3298534883329) != 0 {
		p.SetState(435)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromptDSLParserCODE_TEXT:
			{
				p.SetState(429)
				p.Match(PromptDSLParserCODE_TEXT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case PromptDSLParserCODE_STRING:
			{
				p.SetState(430)
				p.Match(PromptDSLParserCODE_STRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case PromptDSLParserLBRACE:
			{
				p.SetState(431)
				p.Match(PromptDSLParserLBRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(432)
				p.CodeBlockContent()
			}
			{
				p.SetState(433)
				p.Match(PromptDSLParserRBRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(439)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext
	FLOAT_TYPE() antlr.TerminalNode
	INT_TYPE() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	Type_() ITypeContext
	STRING_TYPE() antlr.TerminalNode
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

func (s *TypeContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRUCT, 0)
}

func (s *TypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *TypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *TypeContext) AllFieldDef() []IFieldDefContext {
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

func (s *TypeContext) FieldDef(i int) IFieldDefContext {
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

func (s *TypeContext) FLOAT_TYPE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserFLOAT_TYPE, 0)
}

func (s *TypeContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserINT_TYPE, 0)
}

func (s *TypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACK, 0)
}

func (s *TypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACK, 0)
}

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

func (s *TypeContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING_TYPE, 0)
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
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *PromptDSLParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, PromptDSLParserRULE_type)
	var _la int

	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRUCT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(440)
			p.Match(PromptDSLParserSTRUCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(441)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(445)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserID {
			{
				p.SetState(442)
				p.FieldDef()
			}

			p.SetState(447)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(448)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserFLOAT_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(449)
			p.Match(PromptDSLParserFLOAT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserINT_TYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(450)
			p.Match(PromptDSLParserINT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserLBRACK:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(451)
			p.Match(PromptDSLParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(452)
			p.Match(PromptDSLParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)
			p.Type_()
		}

	case PromptDSLParserSTRING_TYPE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(454)
			p.Match(PromptDSLParserSTRING_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserID:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(455)
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
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *PromptDSLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, PromptDSLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(458)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1008806316530991104) != 0) {
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

	// Getter signatures
	MD() antlr.TerminalNode
	JSON() antlr.TerminalNode

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

func (s *FormatTypeContext) MD() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMD, 0)
}

func (s *FormatTypeContext) JSON() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserJSON, 0)
}

func (s *FormatTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormatTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormatTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterFormatType(s)
	}
}

func (s *FormatTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitFormatType(s)
	}
}

func (p *PromptDSLParser) FormatType() (localctx IFormatTypeContext) {
	localctx = NewFormatTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, PromptDSLParserRULE_formatType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromptDSLParserMD || _la == PromptDSLParserJSON) {
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

func (p *PromptDSLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 25:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PromptDSLParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
