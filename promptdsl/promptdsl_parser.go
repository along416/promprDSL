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
		"", "'{'", "'}'", "':'", "'('", "','", "')'", "'='", "'.'", "'@'", "'['",
		"']'", "'-'", "'[]'", "'md'", "'json'", "'+'", "", "'string'", "'float'",
		"'int'", "'prompt'", "'params'", "'system'", "'user'", "'note'", "'input'",
		"'output'", "'format'", "'type'", "'struct'", "'before'", "'schema'",
		"'after'", "'parse'", "'jsonfix'", "", "", "", "", "'|'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "PLUS",
		"JAVASCRIPT_BLOCK", "STRING_TYPE", "FLOAT_TYPE", "INT_TYPE", "PROMPT",
		"PARAMS", "SYSTEM", "USER", "NOTE", "INPUT", "OUTPUT", "FORMAT", "TYPE",
		"STRUCT", "BEFORE", "SCHEMA", "AFTER", "PARSE", "JSONFIX", "ID", "STRING",
		"NUMBER", "BOOL", "PIPE", "SEMI", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"promptFile", "promptDef", "promptBlock", "inputSection", "outputSection",
		"outputEntry", "systemSection", "userSection", "noteSection", "beforeSection",
		"beforeContent", "beforeEntry", "afterSection", "afterContent", "afterEntry",
		"dslCallExpr", "expr", "fieldDef", "textLine", "paramPath", "structDef",
		"annotation", "annotationArgs", "annotationValue", "arrayLiteral", "textBlock",
		"type", "value", "formatType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 306, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 4, 0, 60, 8, 0, 11, 0, 12, 0,
		61, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 70, 8, 1, 11, 1, 12, 1, 71,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 83, 8, 2, 1,
		3, 1, 3, 1, 3, 4, 3, 88, 8, 3, 11, 3, 12, 3, 89, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 4, 4, 97, 8, 4, 11, 4, 12, 4, 98, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 107, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 114, 8, 5, 11,
		5, 12, 5, 115, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 124, 8, 5, 3,
		5, 126, 8, 5, 1, 6, 1, 6, 1, 6, 4, 6, 131, 8, 6, 11, 6, 12, 6, 132, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 4, 7, 140, 8, 7, 11, 7, 12, 7, 141, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 4, 8, 149, 8, 8, 11, 8, 12, 8, 150, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 163, 8, 10, 10,
		10, 12, 10, 166, 9, 10, 1, 10, 3, 10, 169, 8, 10, 1, 11, 1, 11, 1, 11,
		3, 11, 174, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 5, 13, 184, 8, 13, 10, 13, 12, 13, 187, 9, 13, 1, 13, 3, 13, 190, 8,
		13, 1, 14, 1, 14, 1, 14, 3, 14, 195, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 5, 15, 202, 8, 15, 10, 15, 12, 15, 205, 9, 15, 3, 15, 207, 8, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 215, 8, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 3, 17, 222, 8, 17, 1, 17, 5, 17, 225, 8, 17, 10,
		17, 12, 17, 228, 9, 17, 1, 17, 3, 17, 231, 8, 17, 1, 18, 1, 18, 1, 18,
		3, 18, 236, 8, 18, 1, 19, 1, 19, 1, 19, 5, 19, 241, 8, 19, 10, 19, 12,
		19, 244, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 4, 20, 250, 8, 20, 11, 20,
		12, 20, 251, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 260, 8, 21,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 5, 22, 267, 8, 22, 10, 22, 12, 22, 270,
		9, 22, 1, 23, 1, 23, 3, 23, 274, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5,
		24, 280, 8, 24, 10, 24, 12, 24, 283, 9, 24, 3, 24, 285, 8, 24, 1, 24, 1,
		24, 1, 25, 4, 25, 290, 8, 25, 11, 25, 12, 25, 291, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 3, 26, 300, 8, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		28, 0, 0, 29, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 0, 5, 4, 0, 26, 27,
		31, 31, 33, 33, 36, 36, 2, 0, 32, 32, 34, 36, 4, 0, 3, 3, 12, 12, 36, 38,
		42, 42, 1, 0, 37, 39, 1, 0, 14, 15, 324, 0, 59, 1, 0, 0, 0, 2, 65, 1, 0,
		0, 0, 4, 82, 1, 0, 0, 0, 6, 84, 1, 0, 0, 0, 8, 93, 1, 0, 0, 0, 10, 125,
		1, 0, 0, 0, 12, 127, 1, 0, 0, 0, 14, 136, 1, 0, 0, 0, 16, 145, 1, 0, 0,
		0, 18, 154, 1, 0, 0, 0, 20, 168, 1, 0, 0, 0, 22, 173, 1, 0, 0, 0, 24, 175,
		1, 0, 0, 0, 26, 189, 1, 0, 0, 0, 28, 194, 1, 0, 0, 0, 30, 196, 1, 0, 0,
		0, 32, 214, 1, 0, 0, 0, 34, 216, 1, 0, 0, 0, 36, 235, 1, 0, 0, 0, 38, 237,
		1, 0, 0, 0, 40, 245, 1, 0, 0, 0, 42, 255, 1, 0, 0, 0, 44, 263, 1, 0, 0,
		0, 46, 273, 1, 0, 0, 0, 48, 275, 1, 0, 0, 0, 50, 289, 1, 0, 0, 0, 52, 299,
		1, 0, 0, 0, 54, 301, 1, 0, 0, 0, 56, 303, 1, 0, 0, 0, 58, 60, 3, 2, 1,
		0, 59, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62,
		1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 5, 0, 0, 1, 64, 1, 1, 0, 0, 0,
		65, 66, 5, 21, 0, 0, 66, 67, 5, 36, 0, 0, 67, 69, 5, 1, 0, 0, 68, 70, 3,
		4, 2, 0, 69, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71,
		72, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 74, 5, 2, 0, 0, 74, 3, 1, 0, 0,
		0, 75, 83, 3, 6, 3, 0, 76, 83, 3, 8, 4, 0, 77, 83, 3, 12, 6, 0, 78, 83,
		3, 14, 7, 0, 79, 83, 3, 16, 8, 0, 80, 83, 3, 18, 9, 0, 81, 83, 3, 24, 12,
		0, 82, 75, 1, 0, 0, 0, 82, 76, 1, 0, 0, 0, 82, 77, 1, 0, 0, 0, 82, 78,
		1, 0, 0, 0, 82, 79, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 81, 1, 0, 0, 0,
		83, 5, 1, 0, 0, 0, 84, 85, 5, 26, 0, 0, 85, 87, 5, 1, 0, 0, 86, 88, 3,
		34, 17, 0, 87, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0,
		89, 90, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 5, 2, 0, 0, 92, 7, 1, 0,
		0, 0, 93, 94, 5, 27, 0, 0, 94, 96, 5, 1, 0, 0, 95, 97, 3, 10, 5, 0, 96,
		95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0,
		0, 99, 100, 1, 0, 0, 0, 100, 101, 5, 2, 0, 0, 101, 9, 1, 0, 0, 0, 102,
		103, 5, 28, 0, 0, 103, 104, 5, 3, 0, 0, 104, 106, 3, 56, 28, 0, 105, 107,
		5, 41, 0, 0, 106, 105, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 126, 1, 0,
		0, 0, 108, 109, 5, 29, 0, 0, 109, 110, 5, 36, 0, 0, 110, 111, 5, 30, 0,
		0, 111, 113, 5, 1, 0, 0, 112, 114, 3, 34, 17, 0, 113, 112, 1, 0, 0, 0,
		114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116,
		117, 1, 0, 0, 0, 117, 118, 5, 2, 0, 0, 118, 126, 1, 0, 0, 0, 119, 120,
		5, 32, 0, 0, 120, 121, 5, 3, 0, 0, 121, 123, 3, 52, 26, 0, 122, 124, 5,
		41, 0, 0, 123, 122, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0, 0,
		0, 125, 102, 1, 0, 0, 0, 125, 108, 1, 0, 0, 0, 125, 119, 1, 0, 0, 0, 126,
		11, 1, 0, 0, 0, 127, 128, 5, 23, 0, 0, 128, 130, 5, 1, 0, 0, 129, 131,
		3, 36, 18, 0, 130, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 130, 1,
		0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 5, 2, 0,
		0, 135, 13, 1, 0, 0, 0, 136, 137, 5, 24, 0, 0, 137, 139, 5, 1, 0, 0, 138,
		140, 3, 36, 18, 0, 139, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 139,
		1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 5, 2,
		0, 0, 144, 15, 1, 0, 0, 0, 145, 146, 5, 25, 0, 0, 146, 148, 5, 1, 0, 0,
		147, 149, 3, 36, 18, 0, 148, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150,
		148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153,
		5, 2, 0, 0, 153, 17, 1, 0, 0, 0, 154, 155, 5, 31, 0, 0, 155, 156, 5, 1,
		0, 0, 156, 157, 3, 20, 10, 0, 157, 158, 5, 2, 0, 0, 158, 19, 1, 0, 0, 0,
		159, 164, 3, 22, 11, 0, 160, 161, 5, 16, 0, 0, 161, 163, 3, 22, 11, 0,
		162, 160, 1, 0, 0, 0, 163, 166, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164,
		165, 1, 0, 0, 0, 165, 169, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 167, 169,
		5, 17, 0, 0, 168, 159, 1, 0, 0, 0, 168, 167, 1, 0, 0, 0, 169, 21, 1, 0,
		0, 0, 170, 174, 3, 30, 15, 0, 171, 174, 5, 37, 0, 0, 172, 174, 5, 17, 0,
		0, 173, 170, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 172, 1, 0, 0, 0, 174,
		23, 1, 0, 0, 0, 175, 176, 5, 33, 0, 0, 176, 177, 5, 1, 0, 0, 177, 178,
		3, 26, 13, 0, 178, 179, 5, 2, 0, 0, 179, 25, 1, 0, 0, 0, 180, 185, 3, 28,
		14, 0, 181, 182, 5, 16, 0, 0, 182, 184, 3, 28, 14, 0, 183, 181, 1, 0, 0,
		0, 184, 187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186,
		190, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 190, 5, 17, 0, 0, 189, 180,
		1, 0, 0, 0, 189, 188, 1, 0, 0, 0, 190, 27, 1, 0, 0, 0, 191, 195, 3, 30,
		15, 0, 192, 195, 5, 37, 0, 0, 193, 195, 5, 17, 0, 0, 194, 191, 1, 0, 0,
		0, 194, 192, 1, 0, 0, 0, 194, 193, 1, 0, 0, 0, 195, 29, 1, 0, 0, 0, 196,
		197, 3, 38, 19, 0, 197, 206, 5, 4, 0, 0, 198, 203, 3, 32, 16, 0, 199, 200,
		5, 5, 0, 0, 200, 202, 3, 32, 16, 0, 201, 199, 1, 0, 0, 0, 202, 205, 1,
		0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 207, 1, 0, 0,
		0, 205, 203, 1, 0, 0, 0, 206, 198, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207,
		208, 1, 0, 0, 0, 208, 209, 5, 6, 0, 0, 209, 31, 1, 0, 0, 0, 210, 215, 5,
		37, 0, 0, 211, 215, 5, 38, 0, 0, 212, 215, 5, 39, 0, 0, 213, 215, 3, 38,
		19, 0, 214, 210, 1, 0, 0, 0, 214, 211, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0,
		214, 213, 1, 0, 0, 0, 215, 33, 1, 0, 0, 0, 216, 217, 5, 36, 0, 0, 217,
		218, 5, 3, 0, 0, 218, 221, 3, 52, 26, 0, 219, 220, 5, 7, 0, 0, 220, 222,
		3, 54, 27, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 226, 1,
		0, 0, 0, 223, 225, 3, 42, 21, 0, 224, 223, 1, 0, 0, 0, 225, 228, 1, 0,
		0, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 230, 1, 0, 0, 0,
		228, 226, 1, 0, 0, 0, 229, 231, 5, 41, 0, 0, 230, 229, 1, 0, 0, 0, 230,
		231, 1, 0, 0, 0, 231, 35, 1, 0, 0, 0, 232, 236, 5, 37, 0, 0, 233, 236,
		5, 43, 0, 0, 234, 236, 3, 38, 19, 0, 235, 232, 1, 0, 0, 0, 235, 233, 1,
		0, 0, 0, 235, 234, 1, 0, 0, 0, 236, 37, 1, 0, 0, 0, 237, 242, 7, 0, 0,
		0, 238, 239, 5, 8, 0, 0, 239, 241, 7, 1, 0, 0, 240, 238, 1, 0, 0, 0, 241,
		244, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 39, 1,
		0, 0, 0, 244, 242, 1, 0, 0, 0, 245, 246, 5, 36, 0, 0, 246, 247, 5, 30,
		0, 0, 247, 249, 5, 1, 0, 0, 248, 250, 3, 34, 17, 0, 249, 248, 1, 0, 0,
		0, 250, 251, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252,
		253, 1, 0, 0, 0, 253, 254, 5, 2, 0, 0, 254, 41, 1, 0, 0, 0, 255, 256, 5,
		9, 0, 0, 256, 257, 5, 36, 0, 0, 257, 259, 5, 4, 0, 0, 258, 260, 3, 44,
		22, 0, 259, 258, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0,
		261, 262, 5, 6, 0, 0, 262, 43, 1, 0, 0, 0, 263, 268, 3, 46, 23, 0, 264,
		265, 5, 5, 0, 0, 265, 267, 3, 46, 23, 0, 266, 264, 1, 0, 0, 0, 267, 270,
		1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 45, 1, 0,
		0, 0, 270, 268, 1, 0, 0, 0, 271, 274, 5, 37, 0, 0, 272, 274, 3, 48, 24,
		0, 273, 271, 1, 0, 0, 0, 273, 272, 1, 0, 0, 0, 274, 47, 1, 0, 0, 0, 275,
		284, 5, 10, 0, 0, 276, 281, 5, 37, 0, 0, 277, 278, 5, 5, 0, 0, 278, 280,
		5, 37, 0, 0, 279, 277, 1, 0, 0, 0, 280, 283, 1, 0, 0, 0, 281, 279, 1, 0,
		0, 0, 281, 282, 1, 0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0,
		284, 276, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286,
		287, 5, 11, 0, 0, 287, 49, 1, 0, 0, 0, 288, 290, 7, 2, 0, 0, 289, 288,
		1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 291, 292, 1, 0,
		0, 0, 292, 51, 1, 0, 0, 0, 293, 300, 5, 18, 0, 0, 294, 300, 5, 19, 0, 0,
		295, 300, 5, 20, 0, 0, 296, 297, 5, 13, 0, 0, 297, 300, 3, 52, 26, 0, 298,
		300, 5, 36, 0, 0, 299, 293, 1, 0, 0, 0, 299, 294, 1, 0, 0, 0, 299, 295,
		1, 0, 0, 0, 299, 296, 1, 0, 0, 0, 299, 298, 1, 0, 0, 0, 300, 53, 1, 0,
		0, 0, 301, 302, 7, 3, 0, 0, 302, 55, 1, 0, 0, 0, 303, 304, 7, 4, 0, 0,
		304, 57, 1, 0, 0, 0, 34, 61, 71, 82, 89, 98, 106, 115, 123, 125, 132, 141,
		150, 164, 168, 173, 185, 189, 194, 203, 206, 214, 221, 226, 230, 235, 242,
		251, 259, 268, 273, 281, 284, 291, 299,
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
	PromptDSLParserEOF              = antlr.TokenEOF
	PromptDSLParserT__0             = 1
	PromptDSLParserT__1             = 2
	PromptDSLParserT__2             = 3
	PromptDSLParserT__3             = 4
	PromptDSLParserT__4             = 5
	PromptDSLParserT__5             = 6
	PromptDSLParserT__6             = 7
	PromptDSLParserT__7             = 8
	PromptDSLParserT__8             = 9
	PromptDSLParserT__9             = 10
	PromptDSLParserT__10            = 11
	PromptDSLParserT__11            = 12
	PromptDSLParserT__12            = 13
	PromptDSLParserT__13            = 14
	PromptDSLParserT__14            = 15
	PromptDSLParserPLUS             = 16
	PromptDSLParserJAVASCRIPT_BLOCK = 17
	PromptDSLParserSTRING_TYPE      = 18
	PromptDSLParserFLOAT_TYPE       = 19
	PromptDSLParserINT_TYPE         = 20
	PromptDSLParserPROMPT           = 21
	PromptDSLParserPARAMS           = 22
	PromptDSLParserSYSTEM           = 23
	PromptDSLParserUSER             = 24
	PromptDSLParserNOTE             = 25
	PromptDSLParserINPUT            = 26
	PromptDSLParserOUTPUT           = 27
	PromptDSLParserFORMAT           = 28
	PromptDSLParserTYPE             = 29
	PromptDSLParserSTRUCT           = 30
	PromptDSLParserBEFORE           = 31
	PromptDSLParserSCHEMA           = 32
	PromptDSLParserAFTER            = 33
	PromptDSLParserPARSE            = 34
	PromptDSLParserJSONFIX          = 35
	PromptDSLParserID               = 36
	PromptDSLParserSTRING           = 37
	PromptDSLParserNUMBER           = 38
	PromptDSLParserBOOL             = 39
	PromptDSLParserPIPE             = 40
	PromptDSLParserSEMI             = 41
	PromptDSLParserWS               = 42
	PromptDSLParserLINE_COMMENT     = 43
	PromptDSLParserBLOCK_COMMENT    = 44
)

// PromptDSLParser rules.
const (
	PromptDSLParserRULE_promptFile      = 0
	PromptDSLParserRULE_promptDef       = 1
	PromptDSLParserRULE_promptBlock     = 2
	PromptDSLParserRULE_inputSection    = 3
	PromptDSLParserRULE_outputSection   = 4
	PromptDSLParserRULE_outputEntry     = 5
	PromptDSLParserRULE_systemSection   = 6
	PromptDSLParserRULE_userSection     = 7
	PromptDSLParserRULE_noteSection     = 8
	PromptDSLParserRULE_beforeSection   = 9
	PromptDSLParserRULE_beforeContent   = 10
	PromptDSLParserRULE_beforeEntry     = 11
	PromptDSLParserRULE_afterSection    = 12
	PromptDSLParserRULE_afterContent    = 13
	PromptDSLParserRULE_afterEntry      = 14
	PromptDSLParserRULE_dslCallExpr     = 15
	PromptDSLParserRULE_expr            = 16
	PromptDSLParserRULE_fieldDef        = 17
	PromptDSLParserRULE_textLine        = 18
	PromptDSLParserRULE_paramPath       = 19
	PromptDSLParserRULE_structDef       = 20
	PromptDSLParserRULE_annotation      = 21
	PromptDSLParserRULE_annotationArgs  = 22
	PromptDSLParserRULE_annotationValue = 23
	PromptDSLParserRULE_arrayLiteral    = 24
	PromptDSLParserRULE_textBlock       = 25
	PromptDSLParserRULE_type            = 26
	PromptDSLParserRULE_value           = 27
	PromptDSLParserRULE_formatType      = 28
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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserPROMPT {
		{
			p.SetState(58)
			p.PromptDef()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(63)
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
		p.SetState(65)
		p.Match(PromptDSLParserPROMPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&10997465088) != 0) {
		{
			p.SetState(68)
			p.PromptBlock()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(73)
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
	SystemSection() ISystemSectionContext
	UserSection() IUserSectionContext
	NoteSection() INoteSectionContext
	BeforeSection() IBeforeSectionContext
	AfterSection() IAfterSectionContext

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

func (s *PromptBlockContext) BeforeSection() IBeforeSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBeforeSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBeforeSectionContext)
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
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINPUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.InputSection()
		}

	case PromptDSLParserOUTPUT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.OutputSection()
		}

	case PromptDSLParserSYSTEM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(77)
			p.SystemSection()
		}

	case PromptDSLParserUSER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(78)
			p.UserSection()
		}

	case PromptDSLParserNOTE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(79)
			p.NoteSection()
		}

	case PromptDSLParserBEFORE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(80)
			p.BeforeSection()
		}

	case PromptDSLParserAFTER:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(81)
			p.AfterSection()
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
	p.EnterRule(localctx, 6, PromptDSLParserRULE_inputSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(PromptDSLParserINPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(86)
			p.FieldDef()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
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
	AllOutputEntry() []IOutputEntryContext
	OutputEntry(i int) IOutputEntryContext

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

func (s *OutputSectionContext) AllOutputEntry() []IOutputEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOutputEntryContext); ok {
			len++
		}
	}

	tst := make([]IOutputEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOutputEntryContext); ok {
			tst[i] = t.(IOutputEntryContext)
			i++
		}
	}

	return tst
}

func (s *OutputSectionContext) OutputEntry(i int) IOutputEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutputEntryContext); ok {
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

	return t.(IOutputEntryContext)
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
	p.EnterRule(localctx, 8, PromptDSLParserRULE_outputSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(PromptDSLParserOUTPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5100273664) != 0) {
		{
			p.SetState(95)
			p.OutputEntry()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(100)
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

// IOutputEntryContext is an interface to support dynamic dispatch.
type IOutputEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FORMAT() antlr.TerminalNode
	FormatType() IFormatTypeContext
	SEMI() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	ID() antlr.TerminalNode
	STRUCT() antlr.TerminalNode
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext
	SCHEMA() antlr.TerminalNode
	Type_() ITypeContext

	// IsOutputEntryContext differentiates from other interfaces.
	IsOutputEntryContext()
}

type OutputEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputEntryContext() *OutputEntryContext {
	var p = new(OutputEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputEntry
	return p
}

func InitEmptyOutputEntryContext(p *OutputEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputEntry
}

func (*OutputEntryContext) IsOutputEntryContext() {}

func NewOutputEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputEntryContext {
	var p = new(OutputEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_outputEntry

	return p
}

func (s *OutputEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputEntryContext) FORMAT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserFORMAT, 0)
}

func (s *OutputEntryContext) FormatType() IFormatTypeContext {
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

func (s *OutputEntryContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSEMI, 0)
}

func (s *OutputEntryContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserTYPE, 0)
}

func (s *OutputEntryContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *OutputEntryContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRUCT, 0)
}

func (s *OutputEntryContext) AllFieldDef() []IFieldDefContext {
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

func (s *OutputEntryContext) FieldDef(i int) IFieldDefContext {
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

func (s *OutputEntryContext) SCHEMA() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSCHEMA, 0)
}

func (s *OutputEntryContext) Type_() ITypeContext {
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

func (s *OutputEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterOutputEntry(s)
	}
}

func (s *OutputEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitOutputEntry(s)
	}
}

func (s *OutputEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitOutputEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) OutputEntry() (localctx IOutputEntryContext) {
	localctx = NewOutputEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PromptDSLParserRULE_outputEntry)
	var _la int

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserFORMAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Match(PromptDSLParserFORMAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(PromptDSLParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.FormatType()
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PromptDSLParserSEMI {
			{
				p.SetState(105)
				p.Match(PromptDSLParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case PromptDSLParserTYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.Match(PromptDSLParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(PromptDSLParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.Match(PromptDSLParserSTRUCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Match(PromptDSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PromptDSLParserID {
			{
				p.SetState(112)
				p.FieldDef()
			}

			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(117)
			p.Match(PromptDSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserSCHEMA:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(119)
			p.Match(PromptDSLParserSCHEMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(PromptDSLParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Type_()
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PromptDSLParserSEMI {
			{
				p.SetState(122)
				p.Match(PromptDSLParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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
	p.EnterRule(localctx, 12, PromptDSLParserRULE_systemSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(PromptDSLParserSYSTEM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(PromptDSLParserT__0)
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

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9013190197248) != 0) {
		{
			p.SetState(129)
			p.TextLine()
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
	AllTextLine() []ITextLineContext
	TextLine(i int) ITextLineContext

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

func (s *UserSectionContext) AllTextLine() []ITextLineContext {
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

func (s *UserSectionContext) TextLine(i int) ITextLineContext {
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
	p.EnterRule(localctx, 14, PromptDSLParserRULE_userSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(PromptDSLParserUSER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9013190197248) != 0) {
		{
			p.SetState(138)
			p.TextLine()
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
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
	p.EnterRule(localctx, 16, PromptDSLParserRULE_noteSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(PromptDSLParserNOTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9013190197248) != 0) {
		{
			p.SetState(147)
			p.TextLine()
		}

		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(152)
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

// IBeforeSectionContext is an interface to support dynamic dispatch.
type IBeforeSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BEFORE() antlr.TerminalNode
	BeforeContent() IBeforeContentContext

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

func (s *BeforeSectionContext) BeforeContent() IBeforeContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBeforeContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
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
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterBeforeSection(s)
	}
}

func (s *BeforeSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitBeforeSection(s)
	}
}

func (s *BeforeSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitBeforeSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) BeforeSection() (localctx IBeforeSectionContext) {
	localctx = NewBeforeSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PromptDSLParserRULE_beforeSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(PromptDSLParserBEFORE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.BeforeContent()
	}
	{
		p.SetState(157)
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

// IBeforeContentContext is an interface to support dynamic dispatch.
type IBeforeContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBeforeEntry() []IBeforeEntryContext
	BeforeEntry(i int) IBeforeEntryContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	JAVASCRIPT_BLOCK() antlr.TerminalNode

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

func (s *BeforeContentContext) AllBeforeEntry() []IBeforeEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBeforeEntryContext); ok {
			len++
		}
	}

	tst := make([]IBeforeEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBeforeEntryContext); ok {
			tst[i] = t.(IBeforeEntryContext)
			i++
		}
	}

	return tst
}

func (s *BeforeContentContext) BeforeEntry(i int) IBeforeEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBeforeEntryContext); ok {
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

	return t.(IBeforeEntryContext)
}

func (s *BeforeContentContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserPLUS)
}

func (s *BeforeContentContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserPLUS, i)
}

func (s *BeforeContentContext) JAVASCRIPT_BLOCK() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserJAVASCRIPT_BLOCK, 0)
}

func (s *BeforeContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeforeContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeforeContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterBeforeContent(s)
	}
}

func (s *BeforeContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitBeforeContent(s)
	}
}

func (s *BeforeContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitBeforeContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) BeforeContent() (localctx IBeforeContentContext) {
	localctx = NewBeforeContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PromptDSLParserRULE_beforeContent)
	var _la int

	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.BeforeEntry()
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserPLUS {
			{
				p.SetState(160)
				p.Match(PromptDSLParserPLUS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(161)
				p.BeforeEntry()
			}

			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Match(PromptDSLParserJAVASCRIPT_BLOCK)
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

// IBeforeEntryContext is an interface to support dynamic dispatch.
type IBeforeEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DslCallExpr() IDslCallExprContext
	STRING() antlr.TerminalNode
	JAVASCRIPT_BLOCK() antlr.TerminalNode

	// IsBeforeEntryContext differentiates from other interfaces.
	IsBeforeEntryContext()
}

type BeforeEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeforeEntryContext() *BeforeEntryContext {
	var p = new(BeforeEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_beforeEntry
	return p
}

func InitEmptyBeforeEntryContext(p *BeforeEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_beforeEntry
}

func (*BeforeEntryContext) IsBeforeEntryContext() {}

func NewBeforeEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeforeEntryContext {
	var p = new(BeforeEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_beforeEntry

	return p
}

func (s *BeforeEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *BeforeEntryContext) DslCallExpr() IDslCallExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDslCallExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDslCallExprContext)
}

func (s *BeforeEntryContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *BeforeEntryContext) JAVASCRIPT_BLOCK() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserJAVASCRIPT_BLOCK, 0)
}

func (s *BeforeEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeforeEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeforeEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterBeforeEntry(s)
	}
}

func (s *BeforeEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitBeforeEntry(s)
	}
}

func (s *BeforeEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitBeforeEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) BeforeEntry() (localctx IBeforeEntryContext) {
	localctx = NewBeforeEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PromptDSLParserRULE_beforeEntry)
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.DslCallExpr()
		}

	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserJAVASCRIPT_BLOCK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(172)
			p.Match(PromptDSLParserJAVASCRIPT_BLOCK)
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

// IAfterSectionContext is an interface to support dynamic dispatch.
type IAfterSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AFTER() antlr.TerminalNode
	AfterContent() IAfterContentContext

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

func (s *AfterSectionContext) AfterContent() IAfterContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAfterContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAfterContentContext)
}

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
	p.EnterRule(localctx, 24, PromptDSLParserRULE_afterSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(PromptDSLParserAFTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.AfterContent()
	}
	{
		p.SetState(178)
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

// IAfterContentContext is an interface to support dynamic dispatch.
type IAfterContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAfterEntry() []IAfterEntryContext
	AfterEntry(i int) IAfterEntryContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	JAVASCRIPT_BLOCK() antlr.TerminalNode

	// IsAfterContentContext differentiates from other interfaces.
	IsAfterContentContext()
}

type AfterContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAfterContentContext() *AfterContentContext {
	var p = new(AfterContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_afterContent
	return p
}

func InitEmptyAfterContentContext(p *AfterContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_afterContent
}

func (*AfterContentContext) IsAfterContentContext() {}

func NewAfterContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AfterContentContext {
	var p = new(AfterContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_afterContent

	return p
}

func (s *AfterContentContext) GetParser() antlr.Parser { return s.parser }

func (s *AfterContentContext) AllAfterEntry() []IAfterEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAfterEntryContext); ok {
			len++
		}
	}

	tst := make([]IAfterEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAfterEntryContext); ok {
			tst[i] = t.(IAfterEntryContext)
			i++
		}
	}

	return tst
}

func (s *AfterContentContext) AfterEntry(i int) IAfterEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAfterEntryContext); ok {
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

	return t.(IAfterEntryContext)
}

func (s *AfterContentContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserPLUS)
}

func (s *AfterContentContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserPLUS, i)
}

func (s *AfterContentContext) JAVASCRIPT_BLOCK() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserJAVASCRIPT_BLOCK, 0)
}

func (s *AfterContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AfterContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AfterContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterAfterContent(s)
	}
}

func (s *AfterContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitAfterContent(s)
	}
}

func (s *AfterContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitAfterContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) AfterContent() (localctx IAfterContentContext) {
	localctx = NewAfterContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PromptDSLParserRULE_afterContent)
	var _la int

	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.AfterEntry()
		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserPLUS {
			{
				p.SetState(181)
				p.Match(PromptDSLParserPLUS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(182)
				p.AfterEntry()
			}

			p.SetState(187)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Match(PromptDSLParserJAVASCRIPT_BLOCK)
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

// IAfterEntryContext is an interface to support dynamic dispatch.
type IAfterEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DslCallExpr() IDslCallExprContext
	STRING() antlr.TerminalNode
	JAVASCRIPT_BLOCK() antlr.TerminalNode

	// IsAfterEntryContext differentiates from other interfaces.
	IsAfterEntryContext()
}

type AfterEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAfterEntryContext() *AfterEntryContext {
	var p = new(AfterEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_afterEntry
	return p
}

func InitEmptyAfterEntryContext(p *AfterEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_afterEntry
}

func (*AfterEntryContext) IsAfterEntryContext() {}

func NewAfterEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AfterEntryContext {
	var p = new(AfterEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_afterEntry

	return p
}

func (s *AfterEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *AfterEntryContext) DslCallExpr() IDslCallExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDslCallExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDslCallExprContext)
}

func (s *AfterEntryContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *AfterEntryContext) JAVASCRIPT_BLOCK() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserJAVASCRIPT_BLOCK, 0)
}

func (s *AfterEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AfterEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AfterEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterAfterEntry(s)
	}
}

func (s *AfterEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitAfterEntry(s)
	}
}

func (s *AfterEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitAfterEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) AfterEntry() (localctx IAfterEntryContext) {
	localctx = NewAfterEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PromptDSLParserRULE_afterEntry)
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)
			p.DslCallExpr()
		}

	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserJAVASCRIPT_BLOCK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(193)
			p.Match(PromptDSLParserJAVASCRIPT_BLOCK)
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
	p.EnterRule(localctx, 30, PromptDSLParserRULE_dslCallExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.ParamPath()
	}
	{
		p.SetState(197)
		p.Match(PromptDSLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1041730895872) != 0 {
		{
			p.SetState(198)
			p.Expr()
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserT__4 {
			{
				p.SetState(199)
				p.Match(PromptDSLParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(200)
				p.Expr()
			}

			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(208)
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
	p.EnterRule(localctx, 32, PromptDSLParserRULE_expr)
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
			p.Match(PromptDSLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(212)
			p.Match(PromptDSLParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(213)
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
	p.EnterRule(localctx, 34, PromptDSLParserRULE_fieldDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.Match(PromptDSLParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.Type_()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserT__6 {
		{
			p.SetState(219)
			p.Match(PromptDSLParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)
			p.Value()
		}

	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserT__8 {
		{
			p.SetState(223)
			p.Annotation()
		}

		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSEMI {
		{
			p.SetState(229)
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
	p.EnterRule(localctx, 36, PromptDSLParserRULE_textLine)
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserLINE_COMMENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.Match(PromptDSLParserLINE_COMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(234)
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
	AllSCHEMA() []antlr.TerminalNode
	SCHEMA(i int) antlr.TerminalNode
	AllPARSE() []antlr.TerminalNode
	PARSE(i int) antlr.TerminalNode
	AllJSONFIX() []antlr.TerminalNode
	JSONFIX(i int) antlr.TerminalNode

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

func (s *ParamPathContext) AllSCHEMA() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserSCHEMA)
}

func (s *ParamPathContext) SCHEMA(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSCHEMA, i)
}

func (s *ParamPathContext) AllPARSE() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserPARSE)
}

func (s *ParamPathContext) PARSE(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserPARSE, i)
}

func (s *ParamPathContext) AllJSONFIX() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserJSONFIX)
}

func (s *ParamPathContext) JSONFIX(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserJSONFIX, i)
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
	p.EnterRule(localctx, 38, PromptDSLParserRULE_paramPath)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&79658221568) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserT__7 {
		{
			p.SetState(238)
			p.Match(PromptDSLParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&124554051584) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(244)
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
	p.EnterRule(localctx, 40, PromptDSLParserRULE_structDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Match(PromptDSLParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(248)
			p.FieldDef()
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(253)
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
	p.EnterRule(localctx, 42, PromptDSLParserRULE_annotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(PromptDSLParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Match(PromptDSLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserT__9 || _la == PromptDSLParserSTRING {
		{
			p.SetState(258)
			p.AnnotationArgs()
		}

	}
	{
		p.SetState(261)
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
	p.EnterRule(localctx, 44, PromptDSLParserRULE_annotationArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.AnnotationValue()
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserT__4 {
		{
			p.SetState(264)
			p.Match(PromptDSLParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.AnnotationValue()
		}

		p.SetState(270)
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
	p.EnterRule(localctx, 46, PromptDSLParserRULE_annotationValue)
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)
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
	p.EnterRule(localctx, 48, PromptDSLParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.Match(PromptDSLParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSTRING {
		{
			p.SetState(276)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserT__4 {
			{
				p.SetState(277)
				p.Match(PromptDSLParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(278)
				p.Match(PromptDSLParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(283)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(286)
		p.Match(PromptDSLParserT__10)
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
	p.EnterRule(localctx, 50, PromptDSLParserRULE_textBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4879082852360) != 0) {
		{
			p.SetState(288)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4879082852360) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(291)
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
	p.EnterRule(localctx, 52, PromptDSLParserRULE_type)
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.Match(PromptDSLParserSTRING_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserFLOAT_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
			p.Match(PromptDSLParserFLOAT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserINT_TYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(295)
			p.Match(PromptDSLParserINT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserT__12:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(296)
			p.Match(PromptDSLParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)
			p.Type_()
		}

	case PromptDSLParserID:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(298)
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
	p.EnterRule(localctx, 54, PromptDSLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&962072674304) != 0) {
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
	p.EnterRule(localctx, 56, PromptDSLParserRULE_formatType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromptDSLParserT__13 || _la == PromptDSLParserT__14) {
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
