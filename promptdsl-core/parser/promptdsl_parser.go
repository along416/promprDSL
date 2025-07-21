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
		"", "'{'", "'}'", "':'", "'='", "'('", "')'", "'=='", "'!='", "','",
		"'.'", "'@'", "'['", "']'", "'-'", "'[]'", "'md'", "'json'", "", "'+'",
		"", "", "'string'", "'float'", "'int'", "'prompt'", "'params'", "'system'",
		"'user'", "'note'", "'input'", "'output'", "'format'", "'type'", "'struct'",
		"'before'", "'schema'", "'after'", "'parse'", "'jsonfix'", "'markdown'",
		"'if'", "'else'", "'outputspec'", "", "", "", "", "'|'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "ARRAY_OUTPUTSPEC", "PLUS", "AFTER_BLOCK", "FIX_BLOCK", "STRING_TYPE",
		"FLOAT_TYPE", "INT_TYPE", "PROMPT", "PARAMS", "SYSTEM", "USER", "NOTE",
		"INPUT", "OUTPUT", "FORMAT", "TYPE", "STRUCT", "BEFORE", "SCHEMA", "AFTER",
		"PARSE", "JSONFIX", "MARKDOWN", "IF", "ELSE", "OUTPUTSPEC", "ID", "STRING",
		"NUMBER", "BOOL", "PIPE", "SEMI", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"promptFile", "promptDef", "promptBlock", "inputSection", "outputSection",
		"outputStruct", "outputMarkdown", "beforeSection", "beforeContent",
		"varDef", "systemSection", "sysContent", "moduleDef", "moduleContent",
		"userSection", "userContent", "thencontent", "elsecontent", "ifStatement",
		"condition", "noteSection", "dslCallExpr", "expr", "fieldDef", "textLine",
		"paramPath", "structDef", "annotation", "annotationArgs", "annotationValue",
		"arrayLiteral", "defaultAnnotation", "afterSection", "fixSection", "textBlock",
		"type", "value", "formatType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 52, 385, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 1, 0, 4, 0, 78, 8, 0, 11, 0, 12, 0, 79, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 4, 1, 88, 8, 1, 11, 1, 12, 1, 89, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 102, 8, 2, 1, 3, 1, 3, 1,
		3, 4, 3, 107, 8, 3, 11, 3, 12, 3, 108, 1, 3, 1, 3, 1, 4, 5, 4, 114, 8,
		4, 10, 4, 12, 4, 117, 9, 4, 1, 4, 1, 4, 1, 4, 3, 4, 122, 8, 4, 1, 5, 1,
		5, 4, 5, 126, 8, 5, 11, 5, 12, 5, 127, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 5, 7, 138, 8, 7, 10, 7, 12, 7, 141, 9, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 3, 8, 149, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 4, 10, 158, 8, 10, 11, 10, 12, 10, 159, 1, 10, 1, 10, 1, 10,
		1, 10, 4, 10, 166, 8, 10, 11, 10, 12, 10, 167, 1, 10, 1, 10, 3, 10, 172,
		8, 10, 1, 11, 1, 11, 1, 11, 3, 11, 177, 8, 11, 1, 12, 1, 12, 1, 12, 5,
		12, 182, 8, 12, 10, 12, 12, 12, 185, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		3, 13, 191, 8, 13, 1, 14, 1, 14, 1, 14, 4, 14, 196, 8, 14, 11, 14, 12,
		14, 197, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15,
		208, 8, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 5, 18, 220, 8, 18, 10, 18, 12, 18, 223, 9, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 5, 18, 229, 8, 18, 10, 18, 12, 18, 232, 9, 18, 1, 18, 3,
		18, 235, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 242, 8, 19, 1,
		20, 1, 20, 1, 20, 4, 20, 247, 8, 20, 11, 20, 12, 20, 248, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 258, 8, 21, 10, 21, 12, 21, 261,
		9, 21, 3, 21, 263, 8, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3,
		22, 271, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 278, 8, 23, 1,
		23, 5, 23, 281, 8, 23, 10, 23, 12, 23, 284, 9, 23, 1, 23, 3, 23, 287, 8,
		23, 1, 24, 1, 24, 1, 24, 3, 24, 292, 8, 24, 1, 25, 1, 25, 1, 25, 5, 25,
		297, 8, 25, 10, 25, 12, 25, 300, 9, 25, 1, 26, 1, 26, 1, 26, 1, 26, 4,
		26, 306, 8, 26, 11, 26, 12, 26, 307, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27,
		1, 27, 3, 27, 316, 8, 27, 1, 27, 3, 27, 319, 8, 27, 1, 28, 1, 28, 1, 28,
		5, 28, 324, 8, 28, 10, 28, 12, 28, 327, 9, 28, 1, 29, 1, 29, 3, 29, 331,
		8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 337, 8, 30, 10, 30, 12, 30, 340,
		9, 30, 3, 30, 342, 8, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 3,
		31, 350, 8, 31, 1, 31, 3, 31, 353, 8, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		34, 4, 34, 360, 8, 34, 11, 34, 12, 34, 361, 1, 35, 1, 35, 1, 35, 5, 35,
		367, 8, 35, 10, 35, 12, 35, 370, 9, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 3, 35, 379, 8, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37,
		0, 0, 38, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 0, 5, 1, 0, 7, 8, 4, 0, 30, 31, 35, 35, 37, 37, 44, 44, 4,
		0, 3, 3, 14, 14, 44, 46, 50, 50, 1, 0, 45, 47, 1, 0, 16, 17, 408, 0, 77,
		1, 0, 0, 0, 2, 83, 1, 0, 0, 0, 4, 101, 1, 0, 0, 0, 6, 103, 1, 0, 0, 0,
		8, 115, 1, 0, 0, 0, 10, 123, 1, 0, 0, 0, 12, 131, 1, 0, 0, 0, 14, 134,
		1, 0, 0, 0, 16, 148, 1, 0, 0, 0, 18, 150, 1, 0, 0, 0, 20, 171, 1, 0, 0,
		0, 22, 176, 1, 0, 0, 0, 24, 178, 1, 0, 0, 0, 26, 190, 1, 0, 0, 0, 28, 192,
		1, 0, 0, 0, 30, 207, 1, 0, 0, 0, 32, 209, 1, 0, 0, 0, 34, 211, 1, 0, 0,
		0, 36, 213, 1, 0, 0, 0, 38, 241, 1, 0, 0, 0, 40, 243, 1, 0, 0, 0, 42, 252,
		1, 0, 0, 0, 44, 270, 1, 0, 0, 0, 46, 272, 1, 0, 0, 0, 48, 291, 1, 0, 0,
		0, 50, 293, 1, 0, 0, 0, 52, 301, 1, 0, 0, 0, 54, 311, 1, 0, 0, 0, 56, 320,
		1, 0, 0, 0, 58, 330, 1, 0, 0, 0, 60, 332, 1, 0, 0, 0, 62, 345, 1, 0, 0,
		0, 64, 354, 1, 0, 0, 0, 66, 356, 1, 0, 0, 0, 68, 359, 1, 0, 0, 0, 70, 378,
		1, 0, 0, 0, 72, 380, 1, 0, 0, 0, 74, 382, 1, 0, 0, 0, 76, 78, 3, 2, 1,
		0, 77, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80,
		1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 82, 5, 0, 0, 1, 82, 1, 1, 0, 0, 0,
		83, 84, 5, 25, 0, 0, 84, 85, 5, 44, 0, 0, 85, 87, 5, 1, 0, 0, 86, 88, 3,
		4, 2, 0, 87, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89,
		90, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 5, 2, 0, 0, 92, 3, 1, 0, 0,
		0, 93, 102, 3, 6, 3, 0, 94, 102, 3, 8, 4, 0, 95, 102, 3, 20, 10, 0, 96,
		102, 3, 28, 14, 0, 97, 102, 3, 40, 20, 0, 98, 102, 3, 64, 32, 0, 99, 102,
		3, 66, 33, 0, 100, 102, 3, 24, 12, 0, 101, 93, 1, 0, 0, 0, 101, 94, 1,
		0, 0, 0, 101, 95, 1, 0, 0, 0, 101, 96, 1, 0, 0, 0, 101, 97, 1, 0, 0, 0,
		101, 98, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 100, 1, 0, 0, 0, 102, 5,
		1, 0, 0, 0, 103, 104, 5, 30, 0, 0, 104, 106, 5, 1, 0, 0, 105, 107, 3, 46,
		23, 0, 106, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0,
		108, 109, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 5, 2, 0, 0, 111,
		7, 1, 0, 0, 0, 112, 114, 3, 62, 31, 0, 113, 112, 1, 0, 0, 0, 114, 117,
		1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 118, 1, 0,
		0, 0, 117, 115, 1, 0, 0, 0, 118, 121, 5, 31, 0, 0, 119, 122, 3, 10, 5,
		0, 120, 122, 3, 12, 6, 0, 121, 119, 1, 0, 0, 0, 121, 120, 1, 0, 0, 0, 122,
		9, 1, 0, 0, 0, 123, 125, 5, 1, 0, 0, 124, 126, 3, 46, 23, 0, 125, 124,
		1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0,
		0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 5, 2, 0, 0, 130, 11, 1, 0, 0, 0,
		131, 132, 5, 3, 0, 0, 132, 133, 5, 40, 0, 0, 133, 13, 1, 0, 0, 0, 134,
		135, 5, 35, 0, 0, 135, 139, 5, 1, 0, 0, 136, 138, 3, 16, 8, 0, 137, 136,
		1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 140, 142, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 143, 5, 2, 0, 0,
		143, 15, 1, 0, 0, 0, 144, 149, 3, 18, 9, 0, 145, 149, 3, 44, 22, 0, 146,
		149, 3, 36, 18, 0, 147, 149, 3, 48, 24, 0, 148, 144, 1, 0, 0, 0, 148, 145,
		1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 147, 1, 0, 0, 0, 149, 17, 1, 0,
		0, 0, 150, 151, 5, 44, 0, 0, 151, 152, 5, 4, 0, 0, 152, 153, 3, 44, 22,
		0, 153, 19, 1, 0, 0, 0, 154, 155, 5, 27, 0, 0, 155, 157, 5, 1, 0, 0, 156,
		158, 5, 44, 0, 0, 157, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 157,
		1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 172, 5, 2,
		0, 0, 162, 163, 5, 27, 0, 0, 163, 165, 5, 1, 0, 0, 164, 166, 3, 22, 11,
		0, 165, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167,
		168, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 5, 2, 0, 0, 170, 172,
		1, 0, 0, 0, 171, 154, 1, 0, 0, 0, 171, 162, 1, 0, 0, 0, 172, 21, 1, 0,
		0, 0, 173, 177, 3, 36, 18, 0, 174, 177, 3, 44, 22, 0, 175, 177, 3, 48,
		24, 0, 176, 173, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176, 175, 1, 0, 0, 0,
		177, 23, 1, 0, 0, 0, 178, 179, 5, 44, 0, 0, 179, 183, 5, 1, 0, 0, 180,
		182, 3, 26, 13, 0, 181, 180, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 181,
		1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 186, 1, 0, 0, 0, 185, 183, 1, 0,
		0, 0, 186, 187, 5, 2, 0, 0, 187, 25, 1, 0, 0, 0, 188, 191, 3, 50, 25, 0,
		189, 191, 3, 48, 24, 0, 190, 188, 1, 0, 0, 0, 190, 189, 1, 0, 0, 0, 191,
		27, 1, 0, 0, 0, 192, 193, 5, 28, 0, 0, 193, 195, 5, 1, 0, 0, 194, 196,
		3, 30, 15, 0, 195, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 195, 1,
		0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 200, 5, 2, 0,
		0, 200, 29, 1, 0, 0, 0, 201, 208, 3, 36, 18, 0, 202, 208, 3, 50, 25, 0,
		203, 208, 5, 18, 0, 0, 204, 208, 5, 43, 0, 0, 205, 208, 3, 44, 22, 0, 206,
		208, 3, 48, 24, 0, 207, 201, 1, 0, 0, 0, 207, 202, 1, 0, 0, 0, 207, 203,
		1, 0, 0, 0, 207, 204, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 206, 1, 0,
		0, 0, 208, 31, 1, 0, 0, 0, 209, 210, 3, 30, 15, 0, 210, 33, 1, 0, 0, 0,
		211, 212, 3, 30, 15, 0, 212, 35, 1, 0, 0, 0, 213, 214, 5, 41, 0, 0, 214,
		215, 5, 5, 0, 0, 215, 216, 3, 38, 19, 0, 216, 217, 5, 6, 0, 0, 217, 221,
		5, 1, 0, 0, 218, 220, 3, 32, 16, 0, 219, 218, 1, 0, 0, 0, 220, 223, 1,
		0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 224, 1, 0, 0,
		0, 223, 221, 1, 0, 0, 0, 224, 234, 5, 2, 0, 0, 225, 226, 5, 42, 0, 0, 226,
		230, 5, 1, 0, 0, 227, 229, 3, 34, 17, 0, 228, 227, 1, 0, 0, 0, 229, 232,
		1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 233, 1, 0,
		0, 0, 232, 230, 1, 0, 0, 0, 233, 235, 5, 2, 0, 0, 234, 225, 1, 0, 0, 0,
		234, 235, 1, 0, 0, 0, 235, 37, 1, 0, 0, 0, 236, 237, 3, 44, 22, 0, 237,
		238, 7, 0, 0, 0, 238, 239, 3, 44, 22, 0, 239, 242, 1, 0, 0, 0, 240, 242,
		3, 44, 22, 0, 241, 236, 1, 0, 0, 0, 241, 240, 1, 0, 0, 0, 242, 39, 1, 0,
		0, 0, 243, 244, 5, 29, 0, 0, 244, 246, 5, 1, 0, 0, 245, 247, 3, 48, 24,
		0, 246, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248,
		249, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 251, 5, 2, 0, 0, 251, 41, 1,
		0, 0, 0, 252, 253, 3, 50, 25, 0, 253, 262, 5, 5, 0, 0, 254, 259, 3, 44,
		22, 0, 255, 256, 5, 9, 0, 0, 256, 258, 3, 44, 22, 0, 257, 255, 1, 0, 0,
		0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260,
		263, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 254, 1, 0, 0, 0, 262, 263,
		1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 265, 5, 6, 0, 0, 265, 43, 1, 0,
		0, 0, 266, 271, 3, 50, 25, 0, 267, 271, 5, 45, 0, 0, 268, 271, 5, 46, 0,
		0, 269, 271, 5, 47, 0, 0, 270, 266, 1, 0, 0, 0, 270, 267, 1, 0, 0, 0, 270,
		268, 1, 0, 0, 0, 270, 269, 1, 0, 0, 0, 271, 45, 1, 0, 0, 0, 272, 273, 5,
		44, 0, 0, 273, 274, 5, 3, 0, 0, 274, 277, 3, 70, 35, 0, 275, 276, 5, 4,
		0, 0, 276, 278, 3, 72, 36, 0, 277, 275, 1, 0, 0, 0, 277, 278, 1, 0, 0,
		0, 278, 282, 1, 0, 0, 0, 279, 281, 3, 54, 27, 0, 280, 279, 1, 0, 0, 0,
		281, 284, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283,
		286, 1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 285, 287, 5, 49, 0, 0, 286, 285,
		1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 47, 1, 0, 0, 0, 288, 292, 5, 45,
		0, 0, 289, 292, 5, 51, 0, 0, 290, 292, 3, 50, 25, 0, 291, 288, 1, 0, 0,
		0, 291, 289, 1, 0, 0, 0, 291, 290, 1, 0, 0, 0, 292, 49, 1, 0, 0, 0, 293,
		298, 7, 1, 0, 0, 294, 295, 5, 10, 0, 0, 295, 297, 5, 44, 0, 0, 296, 294,
		1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0,
		0, 0, 299, 51, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 301, 302, 5, 44, 0, 0,
		302, 303, 5, 34, 0, 0, 303, 305, 5, 1, 0, 0, 304, 306, 3, 46, 23, 0, 305,
		304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 307, 308,
		1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 310, 5, 2, 0, 0, 310, 53, 1, 0,
		0, 0, 311, 312, 5, 11, 0, 0, 312, 318, 5, 44, 0, 0, 313, 315, 5, 5, 0,
		0, 314, 316, 3, 56, 28, 0, 315, 314, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0,
		316, 317, 1, 0, 0, 0, 317, 319, 5, 6, 0, 0, 318, 313, 1, 0, 0, 0, 318,
		319, 1, 0, 0, 0, 319, 55, 1, 0, 0, 0, 320, 325, 3, 58, 29, 0, 321, 322,
		5, 9, 0, 0, 322, 324, 3, 58, 29, 0, 323, 321, 1, 0, 0, 0, 324, 327, 1,
		0, 0, 0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 57, 1, 0, 0,
		0, 327, 325, 1, 0, 0, 0, 328, 331, 5, 45, 0, 0, 329, 331, 3, 60, 30, 0,
		330, 328, 1, 0, 0, 0, 330, 329, 1, 0, 0, 0, 331, 59, 1, 0, 0, 0, 332, 341,
		5, 12, 0, 0, 333, 338, 5, 45, 0, 0, 334, 335, 5, 9, 0, 0, 335, 337, 5,
		45, 0, 0, 336, 334, 1, 0, 0, 0, 337, 340, 1, 0, 0, 0, 338, 336, 1, 0, 0,
		0, 338, 339, 1, 0, 0, 0, 339, 342, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0, 341,
		333, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 344,
		5, 13, 0, 0, 344, 61, 1, 0, 0, 0, 345, 346, 5, 11, 0, 0, 346, 352, 5, 44,
		0, 0, 347, 349, 5, 5, 0, 0, 348, 350, 3, 56, 28, 0, 349, 348, 1, 0, 0,
		0, 349, 350, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 353, 5, 6, 0, 0, 352,
		347, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 63, 1, 0, 0, 0, 354, 355, 5,
		20, 0, 0, 355, 65, 1, 0, 0, 0, 356, 357, 5, 21, 0, 0, 357, 67, 1, 0, 0,
		0, 358, 360, 7, 2, 0, 0, 359, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361,
		359, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 69, 1, 0, 0, 0, 363, 364, 5,
		34, 0, 0, 364, 368, 5, 1, 0, 0, 365, 367, 3, 46, 23, 0, 366, 365, 1, 0,
		0, 0, 367, 370, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0,
		369, 371, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 371, 379, 5, 2, 0, 0, 372,
		379, 5, 23, 0, 0, 373, 379, 5, 24, 0, 0, 374, 375, 5, 15, 0, 0, 375, 379,
		3, 70, 35, 0, 376, 379, 5, 22, 0, 0, 377, 379, 5, 44, 0, 0, 378, 363, 1,
		0, 0, 0, 378, 372, 1, 0, 0, 0, 378, 373, 1, 0, 0, 0, 378, 374, 1, 0, 0,
		0, 378, 376, 1, 0, 0, 0, 378, 377, 1, 0, 0, 0, 379, 71, 1, 0, 0, 0, 380,
		381, 7, 3, 0, 0, 381, 73, 1, 0, 0, 0, 382, 383, 7, 4, 0, 0, 383, 75, 1,
		0, 0, 0, 42, 79, 89, 101, 108, 115, 121, 127, 139, 148, 159, 167, 171,
		176, 183, 190, 197, 207, 221, 230, 234, 241, 248, 259, 262, 270, 277, 282,
		286, 291, 298, 307, 315, 318, 325, 330, 338, 341, 349, 352, 361, 368, 378,
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
	PromptDSLParserT__15            = 16
	PromptDSLParserT__16            = 17
	PromptDSLParserARRAY_OUTPUTSPEC = 18
	PromptDSLParserPLUS             = 19
	PromptDSLParserAFTER_BLOCK      = 20
	PromptDSLParserFIX_BLOCK        = 21
	PromptDSLParserSTRING_TYPE      = 22
	PromptDSLParserFLOAT_TYPE       = 23
	PromptDSLParserINT_TYPE         = 24
	PromptDSLParserPROMPT           = 25
	PromptDSLParserPARAMS           = 26
	PromptDSLParserSYSTEM           = 27
	PromptDSLParserUSER             = 28
	PromptDSLParserNOTE             = 29
	PromptDSLParserINPUT            = 30
	PromptDSLParserOUTPUT           = 31
	PromptDSLParserFORMAT           = 32
	PromptDSLParserTYPE             = 33
	PromptDSLParserSTRUCT           = 34
	PromptDSLParserBEFORE           = 35
	PromptDSLParserSCHEMA           = 36
	PromptDSLParserAFTER            = 37
	PromptDSLParserPARSE            = 38
	PromptDSLParserJSONFIX          = 39
	PromptDSLParserMARKDOWN         = 40
	PromptDSLParserIF               = 41
	PromptDSLParserELSE             = 42
	PromptDSLParserOUTPUTSPEC       = 43
	PromptDSLParserID               = 44
	PromptDSLParserSTRING           = 45
	PromptDSLParserNUMBER           = 46
	PromptDSLParserBOOL             = 47
	PromptDSLParserPIPE             = 48
	PromptDSLParserSEMI             = 49
	PromptDSLParserWS               = 50
	PromptDSLParserLINE_COMMENT     = 51
	PromptDSLParserBLOCK_COMMENT    = 52
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
	PromptDSLParserRULE_moduleDef         = 12
	PromptDSLParserRULE_moduleContent     = 13
	PromptDSLParserRULE_userSection       = 14
	PromptDSLParserRULE_userContent       = 15
	PromptDSLParserRULE_thencontent       = 16
	PromptDSLParserRULE_elsecontent       = 17
	PromptDSLParserRULE_ifStatement       = 18
	PromptDSLParserRULE_condition         = 19
	PromptDSLParserRULE_noteSection       = 20
	PromptDSLParserRULE_dslCallExpr       = 21
	PromptDSLParserRULE_expr              = 22
	PromptDSLParserRULE_fieldDef          = 23
	PromptDSLParserRULE_textLine          = 24
	PromptDSLParserRULE_paramPath         = 25
	PromptDSLParserRULE_structDef         = 26
	PromptDSLParserRULE_annotation        = 27
	PromptDSLParserRULE_annotationArgs    = 28
	PromptDSLParserRULE_annotationValue   = 29
	PromptDSLParserRULE_arrayLiteral      = 30
	PromptDSLParserRULE_defaultAnnotation = 31
	PromptDSLParserRULE_afterSection      = 32
	PromptDSLParserRULE_fixSection        = 33
	PromptDSLParserRULE_textBlock         = 34
	PromptDSLParserRULE_type              = 35
	PromptDSLParserRULE_value             = 36
	PromptDSLParserRULE_formatType        = 37
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
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserPROMPT {
		{
			p.SetState(76)
			p.PromptDef()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(81)
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
		p.SetState(83)
		p.Match(PromptDSLParserPROMPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(PromptDSLParserID)
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

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17596349941760) != 0) {
		{
			p.SetState(86)
			p.PromptBlock()
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
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINPUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.InputSection()
		}

	case PromptDSLParserT__10, PromptDSLParserOUTPUT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.OutputSection()
		}

	case PromptDSLParserSYSTEM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
			p.SystemSection()
		}

	case PromptDSLParserUSER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(96)
			p.UserSection()
		}

	case PromptDSLParserNOTE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(97)
			p.NoteSection()
		}

	case PromptDSLParserAFTER_BLOCK:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(98)
			p.AfterSection()
		}

	case PromptDSLParserFIX_BLOCK:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(99)
			p.FixSection()
		}

	case PromptDSLParserID:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(100)
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
		p.SetState(103)
		p.Match(PromptDSLParserINPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
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

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(105)
			p.FieldDef()
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(110)
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
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserT__10 {
		{
			p.SetState(112)
			p.DefaultAnnotation()
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(118)
		p.Match(PromptDSLParserOUTPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserT__0:
		{
			p.SetState(119)
			p.OutputStruct()
		}

	case PromptDSLParserT__2:
		{
			p.SetState(120)
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
	p.EnterRule(localctx, 10, PromptDSLParserRULE_outputStruct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(124)
			p.FieldDef()
		}

		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(129)
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
	p.EnterRule(localctx, 12, PromptDSLParserRULE_outputMarkdown)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(PromptDSLParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
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
	p.EnterRule(localctx, 14, PromptDSLParserRULE_beforeSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(PromptDSLParserBEFORE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2518056647524352) != 0 {
		{
			p.SetState(136)
			p.BeforeContent()
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(142)
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
	p.EnterRule(localctx, 16, PromptDSLParserRULE_beforeContent)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.VarDef()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
			p.IfStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(147)
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
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterVarDef(s)
	}
}

func (s *VarDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitVarDef(s)
	}
}

func (s *VarDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitVarDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) VarDef() (localctx IVarDefContext) {
	localctx = NewVarDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PromptDSLParserRULE_varDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Match(PromptDSLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Expr()
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

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.Match(PromptDSLParserSYSTEM)
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
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PromptDSLParserID {
			{
				p.SetState(156)
				p.Match(PromptDSLParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(161)
			p.Match(PromptDSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(PromptDSLParserSYSTEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Match(PromptDSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2518056647524352) != 0) {
			{
				p.SetState(164)
				p.SysContent()
			}

			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(169)
			p.Match(PromptDSLParserT__1)
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
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterSysContent(s)
	}
}

func (s *SysContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitSysContent(s)
	}
}

func (s *SysContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitSysContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) SysContent() (localctx ISysContentContext) {
	localctx = NewSysContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PromptDSLParserRULE_sysContent)
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.IfStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.Expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(175)
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
	p.EnterRule(localctx, 24, PromptDSLParserRULE_moduleDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2304751391735808) != 0 {
		{
			p.SetState(180)
			p.ModuleContent()
		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(186)
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
	ParamPath() IParamPathContext
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
	p.EnterRule(localctx, 26, PromptDSLParserRULE_moduleContent)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.ParamPath()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
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
	p.EnterRule(localctx, 28, PromptDSLParserRULE_userSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(PromptDSLParserUSER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2526852740808704) != 0) {
		{
			p.SetState(194)
			p.UserContent()
		}

		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(199)
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
	ParamPath() IParamPathContext
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
	p.EnterRule(localctx, 30, PromptDSLParserRULE_userContent)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.IfStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.ParamPath()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(203)
			p.Match(PromptDSLParserARRAY_OUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(204)
			p.Match(PromptDSLParserOUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(205)
			p.Expr()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(206)
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
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterThencontent(s)
	}
}

func (s *ThencontentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitThencontent(s)
	}
}

func (s *ThencontentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitThencontent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) Thencontent() (localctx IThencontentContext) {
	localctx = NewThencontentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PromptDSLParserRULE_thencontent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
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
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterElsecontent(s)
	}
}

func (s *ElsecontentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitElsecontent(s)
	}
}

func (s *ElsecontentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitElsecontent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) Elsecontent() (localctx IElsecontentContext) {
	localctx = NewElsecontentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PromptDSLParserRULE_elsecontent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
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
	Condition() IConditionContext
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
	p.EnterRule(localctx, 36, PromptDSLParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(PromptDSLParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(PromptDSLParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Condition()
	}
	{
		p.SetState(216)
		p.Match(PromptDSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2526852740808704) != 0 {
		{
			p.SetState(218)
			p.Thencontent()
		}

		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(224)
		p.Match(PromptDSLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserELSE {
		{
			p.SetState(225)
			p.Match(PromptDSLParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.Match(PromptDSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2526852740808704) != 0 {
			{
				p.SetState(227)
				p.Elsecontent()
			}

			p.SetState(232)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(233)
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
	p.EnterRule(localctx, 38, PromptDSLParserRULE_condition)
	var _la int

	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(236)

			var _x = p.Expr()

			localctx.(*ConditionContext).lhs = _x
		}
		{
			p.SetState(237)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ConditionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == PromptDSLParserT__6 || _la == PromptDSLParserT__7) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ConditionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(238)

			var _x = p.Expr()

			localctx.(*ConditionContext).rhs = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(240)

			var _x = p.Expr()

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
	p.EnterRule(localctx, 40, PromptDSLParserRULE_noteSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(PromptDSLParserNOTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2304751391735808) != 0) {
		{
			p.SetState(245)
			p.TextLine()
		}

		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(250)
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
	p.EnterRule(localctx, 42, PromptDSLParserRULE_dslCallExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.ParamPath()
	}
	{
		p.SetState(253)
		p.Match(PromptDSLParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&264057810583552) != 0 {
		{
			p.SetState(254)
			p.Expr()
		}
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserT__8 {
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
				p.Expr()
			}

			p.SetState(261)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(264)
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
	ParamPath() IParamPathContext
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOL() antlr.TerminalNode

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
	p.EnterRule(localctx, 44, PromptDSLParserRULE_expr)
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.ParamPath()
		}

	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(268)
			p.Match(PromptDSLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserBOOL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(269)
			p.Match(PromptDSLParserBOOL)
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
	p.EnterRule(localctx, 46, PromptDSLParserRULE_fieldDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Match(PromptDSLParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.Type_()
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserT__3 {
		{
			p.SetState(275)
			p.Match(PromptDSLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Value()
		}

	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserT__10 {
		{
			p.SetState(279)
			p.Annotation()
		}

		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSEMI {
		{
			p.SetState(285)
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
	p.EnterRule(localctx, 48, PromptDSLParserRULE_textLine)
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(288)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserLINE_COMMENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(289)
			p.Match(PromptDSLParserLINE_COMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(290)
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
	p.EnterRule(localctx, 50, PromptDSLParserRULE_paramPath)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17767205961728) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserT__9 {
		{
			p.SetState(294)
			p.Match(PromptDSLParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)
			p.Match(PromptDSLParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(300)
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
	p.EnterRule(localctx, 52, PromptDSLParserRULE_structDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(302)
		p.Match(PromptDSLParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.Match(PromptDSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(304)
			p.FieldDef()
		}

		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(309)
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
	p.EnterRule(localctx, 54, PromptDSLParserRULE_annotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(PromptDSLParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserT__4 {
		{
			p.SetState(313)
			p.Match(PromptDSLParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PromptDSLParserT__11 || _la == PromptDSLParserSTRING {
			{
				p.SetState(314)
				p.AnnotationArgs()
			}

		}
		{
			p.SetState(317)
			p.Match(PromptDSLParserT__5)
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
	p.EnterRule(localctx, 56, PromptDSLParserRULE_annotationArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.AnnotationValue()
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserT__8 {
		{
			p.SetState(321)
			p.Match(PromptDSLParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.AnnotationValue()
		}

		p.SetState(327)
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
	p.EnterRule(localctx, 58, PromptDSLParserRULE_annotationValue)
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(328)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserT__11:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
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
	p.EnterRule(localctx, 60, PromptDSLParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(PromptDSLParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSTRING {
		{
			p.SetState(333)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserT__8 {
			{
				p.SetState(334)
				p.Match(PromptDSLParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(335)
				p.Match(PromptDSLParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(340)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(343)
		p.Match(PromptDSLParserT__12)
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
	ID() antlr.TerminalNode
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

func (s *DefaultAnnotationContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
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
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.EnterDefaultAnnotation(s)
	}
}

func (s *DefaultAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLListener); ok {
		listenerT.ExitDefaultAnnotation(s)
	}
}

func (s *DefaultAnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromptDSLVisitor:
		return t.VisitDefaultAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromptDSLParser) DefaultAnnotation() (localctx IDefaultAnnotationContext) {
	localctx = NewDefaultAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PromptDSLParserRULE_defaultAnnotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Match(PromptDSLParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(346)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserT__4 {
		{
			p.SetState(347)
			p.Match(PromptDSLParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PromptDSLParserT__11 || _la == PromptDSLParserSTRING {
			{
				p.SetState(348)
				p.AnnotationArgs()
			}

		}
		{
			p.SetState(351)
			p.Match(PromptDSLParserT__5)
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

// IAfterSectionContext is an interface to support dynamic dispatch.
type IAfterSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AFTER_BLOCK() antlr.TerminalNode

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

func (s *AfterSectionContext) AFTER_BLOCK() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserAFTER_BLOCK, 0)
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
	p.EnterRule(localctx, 64, PromptDSLParserRULE_afterSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(PromptDSLParserAFTER_BLOCK)
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

	// Getter signatures
	FIX_BLOCK() antlr.TerminalNode

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

func (s *FixSectionContext) FIX_BLOCK() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserFIX_BLOCK, 0)
}

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
	p.EnterRule(localctx, 66, PromptDSLParserRULE_fixSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(356)
		p.Match(PromptDSLParserFIX_BLOCK)
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
	p.EnterRule(localctx, 68, PromptDSLParserRULE_textBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1249045209169928) != 0) {
		{
			p.SetState(358)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1249045209169928) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(361)
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
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext
	FLOAT_TYPE() antlr.TerminalNode
	INT_TYPE() antlr.TerminalNode
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
	p.EnterRule(localctx, 70, PromptDSLParserRULE_type)
	var _la int

	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRUCT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(363)
			p.Match(PromptDSLParserSTRUCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)
			p.Match(PromptDSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserID {
			{
				p.SetState(365)
				p.FieldDef()
			}

			p.SetState(370)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(371)
			p.Match(PromptDSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserFLOAT_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.Match(PromptDSLParserFLOAT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserINT_TYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
			p.Match(PromptDSLParserINT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserT__14:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(374)
			p.Match(PromptDSLParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.Type_()
		}

	case PromptDSLParserSTRING_TYPE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(376)
			p.Match(PromptDSLParserSTRING_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserID:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(377)
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
	p.EnterRule(localctx, 72, PromptDSLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&246290604621824) != 0) {
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
	p.EnterRule(localctx, 74, PromptDSLParserRULE_formatType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromptDSLParserT__15 || _la == PromptDSLParserT__16) {
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
