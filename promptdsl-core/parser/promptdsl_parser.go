// Code generated from PromptDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromptDSL
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
		"", "'{'", "'}'", "'sys'", "':'", "'('", "')'", "'=='", "'!='", "','",
		"'='", "'.'", "'@'", "'['", "']'", "'<after>'", "'</after>'", "'<fix>'",
		"'</fix>'", "'-'", "'[]'", "'md'", "'json'", "'+'", "'string'", "'float'",
		"'int'", "'prompt'", "'params'", "'system'", "'user'", "'note'", "'input'",
		"'output'", "'format'", "'type'", "'struct'", "'before'", "'schema'",
		"'after'", "'parse'", "'jsonfix'", "'markdown'", "'if'", "'else'", "'outputspec'",
		"", "", "", "", "'|'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "PLUS", "STRING_TYPE", "FLOAT_TYPE", "INT_TYPE",
		"PROMPT", "PARAMS", "SYSTEM", "USER", "NOTE", "INPUT", "OUTPUT", "FORMAT",
		"TYPE", "STRUCT", "BEFORE", "SCHEMA", "AFTER", "PARSE", "JSONFIX", "MARKDOWN",
		"IF", "ELSE", "OUTPUTSPEC", "ID", "STRING", "NUMBER", "BOOL", "PIPE",
		"SEMI", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"promptFile", "promptDef", "promptBlock", "sysSection", "moduleDef",
		"moduleContent", "inputSection", "outputSection", "outputStruct", "outputMarkdown",
		"systemSection", "userSection", "userContent", "ifStatement", "condition",
		"noteSection", "dslCallExpr", "expr", "fieldDef", "textLine", "paramPath",
		"structDef", "annotation", "annotationArgs", "annotationValue", "arrayLiteral",
		"afterSection", "fixSection", "textBlock", "type", "value", "formatType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 340, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 1, 0, 4, 0, 66, 8, 0, 11, 0, 12, 0, 67, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 4, 1, 76, 8, 1, 11, 1, 12, 1, 77, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 91, 8, 2, 1, 3, 1, 3, 1, 3, 4,
		3, 96, 8, 3, 11, 3, 12, 3, 97, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 105,
		8, 4, 10, 4, 12, 4, 108, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5, 114, 8, 5,
		1, 6, 1, 6, 1, 6, 4, 6, 119, 8, 6, 11, 6, 12, 6, 120, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 3, 7, 128, 8, 7, 1, 8, 1, 8, 4, 8, 132, 8, 8, 11, 8, 12, 8,
		133, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 4, 10, 144, 8,
		10, 11, 10, 12, 10, 145, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 4, 11, 153,
		8, 11, 11, 11, 12, 11, 154, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 3,
		12, 163, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 171, 8,
		13, 10, 13, 12, 13, 174, 9, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 180,
		8, 13, 10, 13, 12, 13, 183, 9, 13, 1, 13, 3, 13, 186, 8, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 3, 14, 193, 8, 14, 1, 15, 1, 15, 1, 15, 4, 15,
		198, 8, 15, 11, 15, 12, 15, 199, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 5, 16, 209, 8, 16, 10, 16, 12, 16, 212, 9, 16, 3, 16, 214, 8,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 222, 8, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 229, 8, 18, 1, 18, 5, 18, 232, 8, 18,
		10, 18, 12, 18, 235, 9, 18, 1, 18, 3, 18, 238, 8, 18, 1, 19, 1, 19, 1,
		19, 3, 19, 243, 8, 19, 1, 20, 1, 20, 1, 20, 5, 20, 248, 8, 20, 10, 20,
		12, 20, 251, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 4, 21, 257, 8, 21, 11,
		21, 12, 21, 258, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 267,
		8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 5, 23, 274, 8, 23, 10, 23, 12,
		23, 277, 9, 23, 1, 24, 1, 24, 3, 24, 281, 8, 24, 1, 25, 1, 25, 1, 25, 1,
		25, 5, 25, 287, 8, 25, 10, 25, 12, 25, 290, 9, 25, 3, 25, 292, 8, 25, 1,
		25, 1, 25, 1, 26, 1, 26, 5, 26, 298, 8, 26, 10, 26, 12, 26, 301, 9, 26,
		1, 26, 1, 26, 1, 27, 1, 27, 5, 27, 307, 8, 27, 10, 27, 12, 27, 310, 9,
		27, 1, 27, 1, 27, 1, 28, 4, 28, 315, 8, 28, 11, 28, 12, 28, 316, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 327, 8, 29, 10,
		29, 12, 29, 330, 9, 29, 1, 29, 1, 29, 3, 29, 334, 8, 29, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 31, 2, 299, 308, 0, 32, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 0, 5, 1, 0, 7, 8, 4, 0, 32, 33, 37, 37, 39, 39, 46,
		46, 4, 0, 4, 4, 19, 19, 46, 48, 52, 52, 1, 0, 47, 49, 1, 0, 21, 22, 359,
		0, 65, 1, 0, 0, 0, 2, 71, 1, 0, 0, 0, 4, 90, 1, 0, 0, 0, 6, 92, 1, 0, 0,
		0, 8, 101, 1, 0, 0, 0, 10, 113, 1, 0, 0, 0, 12, 115, 1, 0, 0, 0, 14, 124,
		1, 0, 0, 0, 16, 129, 1, 0, 0, 0, 18, 137, 1, 0, 0, 0, 20, 140, 1, 0, 0,
		0, 22, 149, 1, 0, 0, 0, 24, 162, 1, 0, 0, 0, 26, 164, 1, 0, 0, 0, 28, 192,
		1, 0, 0, 0, 30, 194, 1, 0, 0, 0, 32, 203, 1, 0, 0, 0, 34, 221, 1, 0, 0,
		0, 36, 223, 1, 0, 0, 0, 38, 242, 1, 0, 0, 0, 40, 244, 1, 0, 0, 0, 42, 252,
		1, 0, 0, 0, 44, 262, 1, 0, 0, 0, 46, 270, 1, 0, 0, 0, 48, 280, 1, 0, 0,
		0, 50, 282, 1, 0, 0, 0, 52, 295, 1, 0, 0, 0, 54, 304, 1, 0, 0, 0, 56, 314,
		1, 0, 0, 0, 58, 333, 1, 0, 0, 0, 60, 335, 1, 0, 0, 0, 62, 337, 1, 0, 0,
		0, 64, 66, 3, 2, 1, 0, 65, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 65,
		1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 5, 0, 0, 1,
		70, 1, 1, 0, 0, 0, 71, 72, 5, 27, 0, 0, 72, 73, 5, 46, 0, 0, 73, 75, 5,
		1, 0, 0, 74, 76, 3, 4, 2, 0, 75, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77,
		75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 5, 2, 0,
		0, 80, 3, 1, 0, 0, 0, 81, 91, 3, 12, 6, 0, 82, 91, 3, 14, 7, 0, 83, 91,
		3, 6, 3, 0, 84, 91, 3, 20, 10, 0, 85, 91, 3, 22, 11, 0, 86, 91, 3, 30,
		15, 0, 87, 91, 3, 52, 26, 0, 88, 91, 3, 54, 27, 0, 89, 91, 3, 8, 4, 0,
		90, 81, 1, 0, 0, 0, 90, 82, 1, 0, 0, 0, 90, 83, 1, 0, 0, 0, 90, 84, 1,
		0, 0, 0, 90, 85, 1, 0, 0, 0, 90, 86, 1, 0, 0, 0, 90, 87, 1, 0, 0, 0, 90,
		88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 5, 1, 0, 0, 0, 92, 93, 5, 3, 0,
		0, 93, 95, 5, 1, 0, 0, 94, 96, 5, 46, 0, 0, 95, 94, 1, 0, 0, 0, 96, 97,
		1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0,
		99, 100, 5, 2, 0, 0, 100, 7, 1, 0, 0, 0, 101, 102, 5, 46, 0, 0, 102, 106,
		5, 1, 0, 0, 103, 105, 3, 10, 5, 0, 104, 103, 1, 0, 0, 0, 105, 108, 1, 0,
		0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 109, 1, 0, 0, 0,
		108, 106, 1, 0, 0, 0, 109, 110, 5, 2, 0, 0, 110, 9, 1, 0, 0, 0, 111, 114,
		3, 38, 19, 0, 112, 114, 3, 40, 20, 0, 113, 111, 1, 0, 0, 0, 113, 112, 1,
		0, 0, 0, 114, 11, 1, 0, 0, 0, 115, 116, 5, 32, 0, 0, 116, 118, 5, 1, 0,
		0, 117, 119, 3, 36, 18, 0, 118, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0,
		120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122,
		123, 5, 2, 0, 0, 123, 13, 1, 0, 0, 0, 124, 127, 5, 33, 0, 0, 125, 128,
		3, 16, 8, 0, 126, 128, 3, 18, 9, 0, 127, 125, 1, 0, 0, 0, 127, 126, 1,
		0, 0, 0, 128, 15, 1, 0, 0, 0, 129, 131, 5, 1, 0, 0, 130, 132, 3, 36, 18,
		0, 131, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133,
		134, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 5, 2, 0, 0, 136, 17, 1,
		0, 0, 0, 137, 138, 5, 4, 0, 0, 138, 139, 5, 42, 0, 0, 139, 19, 1, 0, 0,
		0, 140, 141, 5, 29, 0, 0, 141, 143, 5, 1, 0, 0, 142, 144, 3, 38, 19, 0,
		143, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145,
		146, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 5, 2, 0, 0, 148, 21, 1,
		0, 0, 0, 149, 150, 5, 30, 0, 0, 150, 152, 5, 1, 0, 0, 151, 153, 3, 24,
		12, 0, 152, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0,
		154, 155, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 5, 2, 0, 0, 157,
		23, 1, 0, 0, 0, 158, 163, 3, 26, 13, 0, 159, 163, 3, 38, 19, 0, 160, 163,
		5, 45, 0, 0, 161, 163, 3, 34, 17, 0, 162, 158, 1, 0, 0, 0, 162, 159, 1,
		0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 161, 1, 0, 0, 0, 163, 25, 1, 0, 0,
		0, 164, 165, 5, 43, 0, 0, 165, 166, 5, 5, 0, 0, 166, 167, 3, 28, 14, 0,
		167, 168, 5, 6, 0, 0, 168, 172, 5, 1, 0, 0, 169, 171, 3, 24, 12, 0, 170,
		169, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173,
		1, 0, 0, 0, 173, 175, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175, 185, 5, 2,
		0, 0, 176, 177, 5, 44, 0, 0, 177, 181, 5, 1, 0, 0, 178, 180, 3, 24, 12,
		0, 179, 178, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181,
		182, 1, 0, 0, 0, 182, 184, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 186,
		5, 2, 0, 0, 185, 176, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 27, 1, 0,
		0, 0, 187, 188, 3, 34, 17, 0, 188, 189, 7, 0, 0, 0, 189, 190, 3, 34, 17,
		0, 190, 193, 1, 0, 0, 0, 191, 193, 3, 34, 17, 0, 192, 187, 1, 0, 0, 0,
		192, 191, 1, 0, 0, 0, 193, 29, 1, 0, 0, 0, 194, 195, 5, 31, 0, 0, 195,
		197, 5, 1, 0, 0, 196, 198, 3, 38, 19, 0, 197, 196, 1, 0, 0, 0, 198, 199,
		1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 1, 0,
		0, 0, 201, 202, 5, 2, 0, 0, 202, 31, 1, 0, 0, 0, 203, 204, 3, 40, 20, 0,
		204, 213, 5, 5, 0, 0, 205, 210, 3, 34, 17, 0, 206, 207, 5, 9, 0, 0, 207,
		209, 3, 34, 17, 0, 208, 206, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208,
		1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212, 210, 1, 0,
		0, 0, 213, 205, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0,
		215, 216, 5, 6, 0, 0, 216, 33, 1, 0, 0, 0, 217, 222, 5, 47, 0, 0, 218,
		222, 5, 48, 0, 0, 219, 222, 5, 49, 0, 0, 220, 222, 3, 40, 20, 0, 221, 217,
		1, 0, 0, 0, 221, 218, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 220, 1, 0,
		0, 0, 222, 35, 1, 0, 0, 0, 223, 224, 5, 46, 0, 0, 224, 225, 5, 4, 0, 0,
		225, 228, 3, 58, 29, 0, 226, 227, 5, 10, 0, 0, 227, 229, 3, 60, 30, 0,
		228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 233, 1, 0, 0, 0, 230,
		232, 3, 44, 22, 0, 231, 230, 1, 0, 0, 0, 232, 235, 1, 0, 0, 0, 233, 231,
		1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 237, 1, 0, 0, 0, 235, 233, 1, 0,
		0, 0, 236, 238, 5, 51, 0, 0, 237, 236, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0,
		238, 37, 1, 0, 0, 0, 239, 243, 5, 47, 0, 0, 240, 243, 5, 53, 0, 0, 241,
		243, 3, 40, 20, 0, 242, 239, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 241,
		1, 0, 0, 0, 243, 39, 1, 0, 0, 0, 244, 249, 7, 1, 0, 0, 245, 246, 5, 11,
		0, 0, 246, 248, 5, 46, 0, 0, 247, 245, 1, 0, 0, 0, 248, 251, 1, 0, 0, 0,
		249, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 41, 1, 0, 0, 0, 251, 249,
		1, 0, 0, 0, 252, 253, 5, 46, 0, 0, 253, 254, 5, 36, 0, 0, 254, 256, 5,
		1, 0, 0, 255, 257, 3, 36, 18, 0, 256, 255, 1, 0, 0, 0, 257, 258, 1, 0,
		0, 0, 258, 256, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0,
		260, 261, 5, 2, 0, 0, 261, 43, 1, 0, 0, 0, 262, 263, 5, 12, 0, 0, 263,
		264, 5, 46, 0, 0, 264, 266, 5, 5, 0, 0, 265, 267, 3, 46, 23, 0, 266, 265,
		1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 269, 5, 6,
		0, 0, 269, 45, 1, 0, 0, 0, 270, 275, 3, 48, 24, 0, 271, 272, 5, 9, 0, 0,
		272, 274, 3, 48, 24, 0, 273, 271, 1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275,
		273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 47, 1, 0, 0, 0, 277, 275, 1,
		0, 0, 0, 278, 281, 5, 47, 0, 0, 279, 281, 3, 50, 25, 0, 280, 278, 1, 0,
		0, 0, 280, 279, 1, 0, 0, 0, 281, 49, 1, 0, 0, 0, 282, 291, 5, 13, 0, 0,
		283, 288, 5, 47, 0, 0, 284, 285, 5, 9, 0, 0, 285, 287, 5, 47, 0, 0, 286,
		284, 1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289,
		1, 0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 291, 283, 1, 0,
		0, 0, 291, 292, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 294, 5, 14, 0, 0,
		294, 51, 1, 0, 0, 0, 295, 299, 5, 15, 0, 0, 296, 298, 9, 0, 0, 0, 297,
		296, 1, 0, 0, 0, 298, 301, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 299, 297,
		1, 0, 0, 0, 300, 302, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 302, 303, 5, 16,
		0, 0, 303, 53, 1, 0, 0, 0, 304, 308, 5, 17, 0, 0, 305, 307, 9, 0, 0, 0,
		306, 305, 1, 0, 0, 0, 307, 310, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 308,
		306, 1, 0, 0, 0, 309, 311, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 311, 312,
		5, 18, 0, 0, 312, 55, 1, 0, 0, 0, 313, 315, 7, 2, 0, 0, 314, 313, 1, 0,
		0, 0, 315, 316, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0,
		317, 57, 1, 0, 0, 0, 318, 334, 5, 24, 0, 0, 319, 334, 5, 25, 0, 0, 320,
		334, 5, 26, 0, 0, 321, 322, 5, 20, 0, 0, 322, 334, 3, 58, 29, 0, 323, 324,
		5, 36, 0, 0, 324, 328, 5, 1, 0, 0, 325, 327, 3, 36, 18, 0, 326, 325, 1,
		0, 0, 0, 327, 330, 1, 0, 0, 0, 328, 326, 1, 0, 0, 0, 328, 329, 1, 0, 0,
		0, 329, 331, 1, 0, 0, 0, 330, 328, 1, 0, 0, 0, 331, 334, 5, 2, 0, 0, 332,
		334, 5, 46, 0, 0, 333, 318, 1, 0, 0, 0, 333, 319, 1, 0, 0, 0, 333, 320,
		1, 0, 0, 0, 333, 321, 1, 0, 0, 0, 333, 323, 1, 0, 0, 0, 333, 332, 1, 0,
		0, 0, 334, 59, 1, 0, 0, 0, 335, 336, 7, 3, 0, 0, 336, 61, 1, 0, 0, 0, 337,
		338, 7, 4, 0, 0, 338, 63, 1, 0, 0, 0, 36, 67, 77, 90, 97, 106, 113, 120,
		127, 133, 145, 154, 162, 172, 181, 185, 192, 199, 210, 213, 221, 228, 233,
		237, 242, 249, 258, 266, 275, 280, 288, 291, 299, 308, 316, 328, 333,
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
	PromptDSLParserT__20         = 21
	PromptDSLParserT__21         = 22
	PromptDSLParserPLUS          = 23
	PromptDSLParserSTRING_TYPE   = 24
	PromptDSLParserFLOAT_TYPE    = 25
	PromptDSLParserINT_TYPE      = 26
	PromptDSLParserPROMPT        = 27
	PromptDSLParserPARAMS        = 28
	PromptDSLParserSYSTEM        = 29
	PromptDSLParserUSER          = 30
	PromptDSLParserNOTE          = 31
	PromptDSLParserINPUT         = 32
	PromptDSLParserOUTPUT        = 33
	PromptDSLParserFORMAT        = 34
	PromptDSLParserTYPE          = 35
	PromptDSLParserSTRUCT        = 36
	PromptDSLParserBEFORE        = 37
	PromptDSLParserSCHEMA        = 38
	PromptDSLParserAFTER         = 39
	PromptDSLParserPARSE         = 40
	PromptDSLParserJSONFIX       = 41
	PromptDSLParserMARKDOWN      = 42
	PromptDSLParserIF            = 43
	PromptDSLParserELSE          = 44
	PromptDSLParserOUTPUTSPEC    = 45
	PromptDSLParserID            = 46
	PromptDSLParserSTRING        = 47
	PromptDSLParserNUMBER        = 48
	PromptDSLParserBOOL          = 49
	PromptDSLParserPIPE          = 50
	PromptDSLParserSEMI          = 51
	PromptDSLParserWS            = 52
	PromptDSLParserLINE_COMMENT  = 53
	PromptDSLParserBLOCK_COMMENT = 54
)

// PromptDSLParser rules.
const (
	PromptDSLParserRULE_promptFile      = 0
	PromptDSLParserRULE_promptDef       = 1
	PromptDSLParserRULE_promptBlock     = 2
	PromptDSLParserRULE_sysSection      = 3
	PromptDSLParserRULE_moduleDef       = 4
	PromptDSLParserRULE_moduleContent   = 5
	PromptDSLParserRULE_inputSection    = 6
	PromptDSLParserRULE_outputSection   = 7
	PromptDSLParserRULE_outputStruct    = 8
	PromptDSLParserRULE_outputMarkdown  = 9
	PromptDSLParserRULE_systemSection   = 10
	PromptDSLParserRULE_userSection     = 11
	PromptDSLParserRULE_userContent     = 12
	PromptDSLParserRULE_ifStatement     = 13
	PromptDSLParserRULE_condition       = 14
	PromptDSLParserRULE_noteSection     = 15
	PromptDSLParserRULE_dslCallExpr     = 16
	PromptDSLParserRULE_expr            = 17
	PromptDSLParserRULE_fieldDef        = 18
	PromptDSLParserRULE_textLine        = 19
	PromptDSLParserRULE_paramPath       = 20
	PromptDSLParserRULE_structDef       = 21
	PromptDSLParserRULE_annotation      = 22
	PromptDSLParserRULE_annotationArgs  = 23
	PromptDSLParserRULE_annotationValue = 24
	PromptDSLParserRULE_arrayLiteral    = 25
	PromptDSLParserRULE_afterSection    = 26
	PromptDSLParserRULE_fixSection      = 27
	PromptDSLParserRULE_textBlock       = 28
	PromptDSLParserRULE_type            = 29
	PromptDSLParserRULE_value           = 30
	PromptDSLParserRULE_formatType      = 31
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserPROMPT {
		{
			p.SetState(64)
			p.PromptDef()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(69)
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
		p.SetState(71)
		p.Match(PromptDSLParserPROMPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70385387339784) != 0) {
		{
			p.SetState(74)
			p.PromptBlock()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(79)
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
	InputSection() IInputSectionContext
	OutputSection() IOutputSectionContext
	SysSection() ISysSectionContext
	SystemSection() ISystemSectionContext
	UserSection() IUserSectionContext
	NoteSection() INoteSectionContext
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

func (s *PromptBlockContext) SysSection() ISysSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISysSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISysSectionContext)
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

func (s *PromptBlockContext) NoteSection() INoteSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoteSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoteSectionContext)
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
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINPUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)
			p.InputSection()
		}

	case PromptDSLParserOUTPUT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.OutputSection()
		}

	case PromptDSLParserT__2:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(83)
			p.SysSection()
		}

	case PromptDSLParserSYSTEM:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(84)
			p.SystemSection()
		}

	case PromptDSLParserUSER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(85)
			p.UserSection()
		}

	case PromptDSLParserNOTE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(86)
			p.NoteSection()
		}

	case PromptDSLParserT__14:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(87)
			p.AfterSection()
		}

	case PromptDSLParserT__16:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(88)
			p.FixSection()
		}

	case PromptDSLParserID:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(89)
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

// ISysSectionContext is an interface to support dynamic dispatch.
type ISysSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsSysSectionContext differentiates from other interfaces.
	IsSysSectionContext()
}

type SysSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySysSectionContext() *SysSectionContext {
	var p = new(SysSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_sysSection
	return p
}

func InitEmptySysSectionContext(p *SysSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_sysSection
}

func (*SysSectionContext) IsSysSectionContext() {}

func NewSysSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SysSectionContext {
	var p = new(SysSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_sysSection

	return p
}

func (s *SysSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SysSectionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserID)
}

func (s *SysSectionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, i)
}

func (s *SysSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SysSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SysSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterSysSection(s)
	}
}

func (s *SysSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitSysSection(s)
	}
}

func (s *SysSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitSysSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) SysSection() (localctx ISysSectionContext) {
	localctx = NewSysSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PromptDSLParserRULE_sysSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(PromptDSLParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(94)
			p.Match(PromptDSLParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
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

// IModuleDefContext is an interface to support dynamic dispatch.
type IModuleDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
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
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterModuleDef(s)
	}
}

func (s *ModuleDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitModuleDef(s)
	}
}

func (s *ModuleDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitModuleDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) ModuleDef() (localctx IModuleDefContext) {
	localctx = NewModuleDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PromptDSLParserRULE_moduleDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9219005566943232) != 0 {
		{
			p.SetState(103)
			p.ModuleContent()
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(109)
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

// IModuleContentContext is an interface to support dynamic dispatch.
type IModuleContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TextLine() ITextLineContext
	ParamPath() IParamPathContext

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

func (s *ModuleContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterModuleContent(s)
	}
}

func (s *ModuleContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitModuleContent(s)
	}
}

func (s *ModuleContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitModuleContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) ModuleContent() (localctx IModuleContentContext) {
	localctx = NewModuleContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PromptDSLParserRULE_moduleContent)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.TextLine()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.ParamPath()
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

// IInputSectionContext is an interface to support dynamic dispatch.
type IInputSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INPUT() antlr.TerminalNode
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

func (p *PromptDSLParser) InputSection() (localctx IInputSectionContext) {
	localctx = NewInputSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PromptDSLParserRULE_inputSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(PromptDSLParserINPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(117)
			p.FieldDef()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(122)
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

// IOutputSectionContext is an interface to support dynamic dispatch.
type IOutputSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OUTPUT() antlr.TerminalNode
	OutputStruct() IOutputStructContext
	OutputMarkdown() IOutputMarkdownContext

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

func (s *OutputSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *PromptDSLParser) OutputSection() (localctx IOutputSectionContext) {
	localctx = NewOutputSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PromptDSLParserRULE_outputSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(PromptDSLParserOUTPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserT__0:
		{
			p.SetState(125)
			p.OutputStruct()
		}

	case PromptDSLParserT__3:
		{
			p.SetState(126)
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
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterOutputStruct(s)
	}
}

func (s *OutputStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitOutputStruct(s)
	}
}

func (s *OutputStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitOutputStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) OutputStruct() (localctx IOutputStructContext) {
	localctx = NewOutputStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PromptDSLParserRULE_outputStruct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(130)
			p.FieldDef()
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(135)
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

// IOutputMarkdownContext is an interface to support dynamic dispatch.
type IOutputMarkdownContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
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
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterOutputMarkdown(s)
	}
}

func (s *OutputMarkdownContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitOutputMarkdown(s)
	}
}

func (s *OutputMarkdownContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitOutputMarkdown(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) OutputMarkdown() (localctx IOutputMarkdownContext) {
	localctx = NewOutputMarkdownContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PromptDSLParserRULE_outputMarkdown)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(PromptDSLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
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

// ISystemSectionContext is an interface to support dynamic dispatch.
type ISystemSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYSTEM() antlr.TerminalNode
	AllTextLine() []ITextLineContext
	TextLine(i int) ITextLineContext

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

func (s *SystemSectionContext) AllTextLine() []ITextLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITextLineContext); ok {
			len++
		}
	}

	tst := make([]ITextLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITextLineContext); ok {
			tst[i] = t.(ITextLineContext)
			i++
		}
	}

	return tst
}

func (s *SystemSectionContext) TextLine(i int) ITextLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextLineContext); ok {
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

	return t.(ITextLineContext)
}

func (s *SystemSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SystemSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SystemSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterSystemSection(s)
	}
}

func (s *SystemSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitSystemSection(s)
	}
}

func (s *SystemSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitSystemSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) SystemSection() (localctx ISystemSectionContext) {
	localctx = NewSystemSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PromptDSLParserRULE_systemSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(PromptDSLParserSYSTEM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9219005566943232) != 0) {
		{
			p.SetState(142)
			p.TextLine()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(147)
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

// IUserSectionContext is an interface to support dynamic dispatch.
type IUserSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	USER() antlr.TerminalNode
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
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterUserSection(s)
	}
}

func (s *UserSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitUserSection(s)
	}
}

func (s *UserSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitUserSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) UserSection() (localctx IUserSectionContext) {
	localctx = NewUserSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PromptDSLParserRULE_userSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(PromptDSLParserUSER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&10107410962186240) != 0) {
		{
			p.SetState(151)
			p.UserContent()
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
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

// IUserContentContext is an interface to support dynamic dispatch.
type IUserContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfStatement() IIfStatementContext
	TextLine() ITextLineContext
	OUTPUTSPEC() antlr.TerminalNode
	Expr() IExprContext

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

func (s *UserContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterUserContent(s)
	}
}

func (s *UserContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitUserContent(s)
	}
}

func (s *UserContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitUserContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) UserContent() (localctx IUserContentContext) {
	localctx = NewUserContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PromptDSLParserRULE_userContent)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.IfStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.TextLine()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
			p.Match(PromptDSLParserOUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(161)
			p.Expr()
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Condition() IConditionContext
	AllUserContent() []IUserContentContext
	UserContent(i int) IUserContentContext
	ELSE() antlr.TerminalNode

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

func (s *IfStatementContext) AllUserContent() []IUserContentContext {
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

func (s *IfStatementContext) UserContent(i int) IUserContentContext {
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

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserELSE, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PromptDSLParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(PromptDSLParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Match(PromptDSLParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Condition()
	}
	{
		p.SetState(167)
		p.Match(PromptDSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&10107410962186240) != 0 {
		{
			p.SetState(169)
			p.UserContent()
		}

		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(175)
		p.Match(PromptDSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserELSE {
		{
			p.SetState(176)
			p.Match(PromptDSLParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Match(PromptDSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&10107410962186240) != 0 {
			{
				p.SetState(178)
				p.UserContent()
			}

			p.SetState(183)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(184)
			p.Match(PromptDSLParserT__1)
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

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PromptDSLParserRULE_condition)
	var _la int

	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Expr()
		}
		{
			p.SetState(188)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PromptDSLParserT__6 || _la == PromptDSLParserT__7) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(189)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.Expr()
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

// INoteSectionContext is an interface to support dynamic dispatch.
type INoteSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOTE() antlr.TerminalNode
	AllTextLine() []ITextLineContext
	TextLine(i int) ITextLineContext

	// IsNoteSectionContext differentiates from other interfaces.
	IsNoteSectionContext()
}

type NoteSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoteSectionContext() *NoteSectionContext {
	var p = new(NoteSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_noteSection
	return p
}

func InitEmptyNoteSectionContext(p *NoteSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_noteSection
}

func (*NoteSectionContext) IsNoteSectionContext() {}

func NewNoteSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoteSectionContext {
	var p = new(NoteSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_noteSection

	return p
}

func (s *NoteSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *NoteSectionContext) NOTE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserNOTE, 0)
}

func (s *NoteSectionContext) AllTextLine() []ITextLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITextLineContext); ok {
			len++
		}
	}

	tst := make([]ITextLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITextLineContext); ok {
			tst[i] = t.(ITextLineContext)
			i++
		}
	}

	return tst
}

func (s *NoteSectionContext) TextLine(i int) ITextLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextLineContext); ok {
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

	return t.(ITextLineContext)
}

func (s *NoteSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoteSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoteSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterNoteSection(s)
	}
}

func (s *NoteSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitNoteSection(s)
	}
}

func (s *NoteSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitNoteSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) NoteSection() (localctx INoteSectionContext) {
	localctx = NewNoteSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PromptDSLParserRULE_noteSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(PromptDSLParserNOTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9219005566943232) != 0) {
		{
			p.SetState(196)
			p.TextLine()
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(201)
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

// IDslCallExprContext is an interface to support dynamic dispatch.
type IDslCallExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParamPath() IParamPathContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext

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

func (s *DslCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DslCallExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DslCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterDslCallExpr(s)
	}
}

func (s *DslCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitDslCallExpr(s)
	}
}

func (s *DslCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitDslCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) DslCallExpr() (localctx IDslCallExprContext) {
	localctx = NewDslCallExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PromptDSLParserRULE_dslCallExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.ParamPath()
	}
	{
		p.SetState(204)
		p.Match(PromptDSLParserT__4)
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

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1056231242334208) != 0 {
		{
			p.SetState(205)
			p.Expr()
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserT__8 {
			{
				p.SetState(206)
				p.Match(PromptDSLParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(207)
				p.Expr()
			}

			p.SetState(212)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(215)
		p.Match(PromptDSLParserT__5)
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

	// Getter signatures
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	ParamPath() IParamPathContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserNUMBER, 0)
}

func (s *ExprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserBOOL, 0)
}

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

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PromptDSLParserRULE_expr)
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(218)
			p.Match(PromptDSLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(219)
			p.Match(PromptDSLParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(220)
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
	p.EnterRule(localctx, 36, PromptDSLParserRULE_fieldDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(PromptDSLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Type_()
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserT__9 {
		{
			p.SetState(226)
			p.Match(PromptDSLParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Value()
		}

	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserT__11 {
		{
			p.SetState(230)
			p.Annotation()
		}

		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSEMI {
		{
			p.SetState(236)
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
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterTextLine(s)
	}
}

func (s *TextLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitTextLine(s)
	}
}

func (s *TextLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitTextLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) TextLine() (localctx ITextLineContext) {
	localctx = NewTextLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PromptDSLParserRULE_textLine)
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(239)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserLINE_COMMENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(240)
			p.Match(PromptDSLParserLINE_COMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(241)
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

func (s *ParamPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterParamPath(s)
	}
}

func (s *ParamPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitParamPath(s)
	}
}

func (s *ParamPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitParamPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) ParamPath() (localctx IParamPathContext) {
	localctx = NewParamPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PromptDSLParserRULE_paramPath)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&71068823846912) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserT__10 {
		{
			p.SetState(245)
			p.Match(PromptDSLParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Match(PromptDSLParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(251)
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
	p.EnterRule(localctx, 42, PromptDSLParserRULE_structDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
		p.Match(PromptDSLParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(255)
			p.FieldDef()
		}

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(260)
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
	p.EnterRule(localctx, 44, PromptDSLParserRULE_annotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Match(PromptDSLParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.Match(PromptDSLParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserT__12 || _la == PromptDSLParserSTRING {
		{
			p.SetState(265)
			p.AnnotationArgs()
		}

	}
	{
		p.SetState(268)
		p.Match(PromptDSLParserT__5)
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
	p.EnterRule(localctx, 46, PromptDSLParserRULE_annotationArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.AnnotationValue()
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserT__8 {
		{
			p.SetState(271)
			p.Match(PromptDSLParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)
			p.AnnotationValue()
		}

		p.SetState(277)
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
	p.EnterRule(localctx, 48, PromptDSLParserRULE_annotationValue)
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(278)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(279)
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
	p.EnterRule(localctx, 50, PromptDSLParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(PromptDSLParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSTRING {
		{
			p.SetState(283)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserT__8 {
			{
				p.SetState(284)
				p.Match(PromptDSLParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(285)
				p.Match(PromptDSLParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(290)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(293)
		p.Match(PromptDSLParserT__13)
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
func (s *AfterSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AfterSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AfterSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterAfterSection(s)
	}
}

func (s *AfterSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitAfterSection(s)
	}
}

func (s *AfterSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitAfterSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) AfterSection() (localctx IAfterSectionContext) {
	localctx = NewAfterSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PromptDSLParserRULE_afterSection)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Match(PromptDSLParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(296)
			p.MatchWildcard()

		}
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(302)
		p.Match(PromptDSLParserT__15)
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

// IFixSectionContext is an interface to support dynamic dispatch.
type IFixSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
func (s *FixSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FixSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterFixSection(s)
	}
}

func (s *FixSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitFixSection(s)
	}
}

func (s *FixSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitFixSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) FixSection() (localctx IFixSectionContext) {
	localctx = NewFixSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PromptDSLParserRULE_fixSection)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(PromptDSLParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(305)
			p.MatchWildcard()

		}
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.Match(PromptDSLParserT__17)
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
	p.EnterRule(localctx, 56, PromptDSLParserRULE_textBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4996180837138448) != 0) {
		{
			p.SetState(313)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4996180837138448) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(316)
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
	STRING_TYPE() antlr.TerminalNode
	FLOAT_TYPE() antlr.TerminalNode
	INT_TYPE() antlr.TerminalNode
	Type_() ITypeContext
	STRUCT() antlr.TerminalNode
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext
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

func (s *TypeContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING_TYPE, 0)
}

func (s *TypeContext) FLOAT_TYPE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserFLOAT_TYPE, 0)
}

func (s *TypeContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserINT_TYPE, 0)
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

func (s *TypeContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRUCT, 0)
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
	p.EnterRule(localctx, 58, PromptDSLParserRULE_type)
	var _la int

	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)
			p.Match(PromptDSLParserSTRING_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserFLOAT_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(319)
			p.Match(PromptDSLParserFLOAT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserINT_TYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(320)
			p.Match(PromptDSLParserINT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserT__19:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(321)
			p.Match(PromptDSLParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Type_()
		}

	case PromptDSLParserSTRUCT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(323)
			p.Match(PromptDSLParserSTRUCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
			p.Match(PromptDSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserID {
			{
				p.SetState(325)
				p.FieldDef()
			}

			p.SetState(330)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(331)
			p.Match(PromptDSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserID:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(332)
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
	p.EnterRule(localctx, 60, PromptDSLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&985162418487296) != 0) {
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
	p.EnterRule(localctx, 62, PromptDSLParserRULE_formatType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromptDSLParserT__20 || _la == PromptDSLParserT__21) {
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
