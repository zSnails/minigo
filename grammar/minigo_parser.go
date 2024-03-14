// Code generated from Minigo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package grammar // Minigo
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

type MinigoParser struct {
	*antlr.BaseParser
}

var MinigoParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func minigoParserInit() {
	staticData := &MinigoParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'('", "')'", "'default'", "'case'", "':'", "'switch'", "'for'",
		"'if'", "'else'", "'/='", "'%='", "'&^='", "'<<='", "'>>='", "'^='",
		"'*='", "'|='", "'-='", "'&='", "'+='", "'++'", "'--'", "'continue'",
		"':='", "'break'", "'return'", "'println'", "'print'", "'cap'", "'len'",
		"'append'", "'.'", "", "", "", "", "", "", "", "'!'", "'||'", "'&&'",
		"'>='", "'<='", "'>'", "'<'", "'!='", "'=='", "'^'", "'|'", "'-'", "'+'",
		"'&^'", "'&'", "'>>'", "'<<'", "'%'", "'/'", "'*'", "'{'", "'}'", "'struct'",
		"", "'['", "']'", "','", "'func'", "'type'", "';'", "'var'", "'='",
		"'package'",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "LEFTPARENTHESIS", "RIGHTPARENTHESIS", "DEFAULT", "CASE",
		"COLON", "SWITCH", "FOR", "IF", "ELSE", "IDIV", "IMOD", "IANDXOR", "ILEFTSHIFT",
		"IRIGHTSHIFT", "IXOR", "IMUL", "IOR", "ISUB", "IAND", "IADD", "POSTINC",
		"POSTDEC", "CONTINUE", "WALRUS", "BREAK", "RETURN", "PRINTLN", "PRINT",
		"CAP", "LEN", "APPEND", "DOT", "WHITESPACE", "NEWLINE", "INTERPRETEDSTRINGLITERAL",
		"RAWSTRINGLITERAL", "ESCAPEDSEQUENCES", "RUNELITERAL", "FLOATLITERAL",
		"NOT", "OR", "AND", "GREATERTHANEQUAL", "LESSTHANEQUAL", "GREATERTHAN",
		"LESSTHAN", "NEGATION", "COMPARISON", "CARET", "PIPE", "MINUS", "PLUS",
		"AMPERSANDCARET", "AMPERSAND", "RIGHTSHIFT", "LEFTSHIFT", "MOD", "DIV",
		"TIMES", "LEFTCURLYBRACE", "RIGHTCURLYBRACE", "STRUCT", "INTLITERAL",
		"LEFTBRACKET", "RIGHTBRACKET", "COMMA", "FUNC", "TYPE", "SEMICOLON",
		"VAR", "EQUALS", "PACKAGE", "IDENTIFIER",
	}
	staticData.RuleNames = []string{
		"root", "topDeclarationList", "variableDecl", "innerVarDecls", "singleVarDecl",
		"singleVarDeclNoExps", "typeDecl", "innerTypeDecls", "singleTypeDecl",
		"funcDecl", "funcFrontDecl", "funcArgsDecls", "declType", "sliceDeclType",
		"arrayDeclType", "structDeclType", "structMemDecls", "identifierList",
		"expression", "expressionList", "primaryExpression", "operand", "literal",
		"index", "arguments", "selector", "appendExpression", "lengthExpression",
		"capExpression", "statementList", "block", "statement", "simpleStatement",
		"assignmentStatement", "ifStatement", "loop", "switch", "expressionCaseClauseList",
		"expressionCaseClause", "expressionSwitchCase",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 74, 500, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 5, 1, 89, 8, 1, 10, 1, 12, 1, 92, 9, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 108, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 115, 8, 3, 10, 3, 12,
		3, 118, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 130, 8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 149, 8, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 156, 8, 7, 10, 7, 12, 7, 159, 9, 7, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10,
		172, 8, 10, 1, 10, 1, 10, 5, 10, 176, 8, 10, 10, 10, 12, 10, 179, 9, 10,
		1, 11, 1, 11, 1, 11, 5, 11, 184, 8, 11, 10, 11, 12, 11, 187, 9, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 197, 8, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 5, 15, 211, 8, 15, 10, 15, 12, 15, 214, 9, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 223, 8, 16, 10, 16, 12, 16, 226,
		9, 16, 1, 17, 1, 17, 1, 17, 5, 17, 231, 8, 17, 10, 17, 12, 17, 234, 9,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		3, 18, 246, 8, 18, 1, 18, 1, 18, 1, 18, 5, 18, 251, 8, 18, 10, 18, 12,
		18, 254, 9, 18, 1, 19, 1, 19, 1, 19, 5, 19, 259, 8, 19, 10, 19, 12, 19,
		262, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 269, 8, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 277, 8, 20, 10, 20, 12, 20, 280,
		9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 288, 8, 21, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 5, 24, 298, 8, 24,
		10, 24, 12, 24, 301, 9, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 5, 29, 326, 8, 29, 10, 29, 12,
		29, 329, 9, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 3, 31,
		338, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 345, 8, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 3, 31, 351, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 375, 8, 31, 1,
		32, 1, 32, 3, 32, 379, 8, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32,
		386, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3,
		33, 396, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 3, 34, 430, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 453, 8, 35, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 3, 36, 481, 8, 36, 1, 37, 1, 37, 1, 37, 5, 37,
		486, 8, 37, 10, 37, 12, 37, 489, 9, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		39, 1, 39, 1, 39, 3, 39, 498, 8, 39, 1, 39, 0, 2, 36, 40, 40, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		0, 4, 1, 0, 42, 60, 3, 0, 36, 37, 39, 40, 64, 64, 1, 0, 22, 23, 1, 0, 11,
		21, 526, 0, 80, 1, 0, 0, 0, 2, 90, 1, 0, 0, 0, 4, 107, 1, 0, 0, 0, 6, 109,
		1, 0, 0, 0, 8, 129, 1, 0, 0, 0, 10, 131, 1, 0, 0, 0, 12, 148, 1, 0, 0,
		0, 14, 150, 1, 0, 0, 0, 16, 160, 1, 0, 0, 0, 18, 163, 1, 0, 0, 0, 20, 167,
		1, 0, 0, 0, 22, 180, 1, 0, 0, 0, 24, 196, 1, 0, 0, 0, 26, 198, 1, 0, 0,
		0, 28, 202, 1, 0, 0, 0, 30, 207, 1, 0, 0, 0, 32, 217, 1, 0, 0, 0, 34, 227,
		1, 0, 0, 0, 36, 245, 1, 0, 0, 0, 38, 255, 1, 0, 0, 0, 40, 268, 1, 0, 0,
		0, 42, 287, 1, 0, 0, 0, 44, 289, 1, 0, 0, 0, 46, 291, 1, 0, 0, 0, 48, 295,
		1, 0, 0, 0, 50, 304, 1, 0, 0, 0, 52, 307, 1, 0, 0, 0, 54, 314, 1, 0, 0,
		0, 56, 319, 1, 0, 0, 0, 58, 327, 1, 0, 0, 0, 60, 330, 1, 0, 0, 0, 62, 374,
		1, 0, 0, 0, 64, 385, 1, 0, 0, 0, 66, 395, 1, 0, 0, 0, 68, 429, 1, 0, 0,
		0, 70, 452, 1, 0, 0, 0, 72, 480, 1, 0, 0, 0, 74, 487, 1, 0, 0, 0, 76, 490,
		1, 0, 0, 0, 78, 497, 1, 0, 0, 0, 80, 81, 5, 73, 0, 0, 81, 82, 5, 74, 0,
		0, 82, 83, 5, 70, 0, 0, 83, 84, 3, 2, 1, 0, 84, 1, 1, 0, 0, 0, 85, 89,
		3, 4, 2, 0, 86, 89, 3, 12, 6, 0, 87, 89, 3, 18, 9, 0, 88, 85, 1, 0, 0,
		0, 88, 86, 1, 0, 0, 0, 88, 87, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88,
		1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 3, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0,
		93, 94, 5, 71, 0, 0, 94, 95, 3, 8, 4, 0, 95, 96, 5, 70, 0, 0, 96, 108,
		1, 0, 0, 0, 97, 98, 5, 71, 0, 0, 98, 99, 5, 2, 0, 0, 99, 100, 3, 6, 3,
		0, 100, 101, 5, 3, 0, 0, 101, 102, 5, 70, 0, 0, 102, 108, 1, 0, 0, 0, 103,
		104, 5, 71, 0, 0, 104, 105, 5, 2, 0, 0, 105, 106, 5, 3, 0, 0, 106, 108,
		5, 70, 0, 0, 107, 93, 1, 0, 0, 0, 107, 97, 1, 0, 0, 0, 107, 103, 1, 0,
		0, 0, 108, 5, 1, 0, 0, 0, 109, 110, 3, 8, 4, 0, 110, 116, 5, 70, 0, 0,
		111, 112, 3, 8, 4, 0, 112, 113, 5, 70, 0, 0, 113, 115, 1, 0, 0, 0, 114,
		111, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117,
		1, 0, 0, 0, 117, 7, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 120, 3, 34,
		17, 0, 120, 121, 3, 24, 12, 0, 121, 122, 5, 72, 0, 0, 122, 123, 3, 38,
		19, 0, 123, 130, 1, 0, 0, 0, 124, 125, 3, 34, 17, 0, 125, 126, 5, 72, 0,
		0, 126, 127, 3, 38, 19, 0, 127, 130, 1, 0, 0, 0, 128, 130, 3, 10, 5, 0,
		129, 119, 1, 0, 0, 0, 129, 124, 1, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130,
		9, 1, 0, 0, 0, 131, 132, 3, 34, 17, 0, 132, 133, 3, 24, 12, 0, 133, 11,
		1, 0, 0, 0, 134, 135, 5, 69, 0, 0, 135, 136, 3, 16, 8, 0, 136, 137, 5,
		70, 0, 0, 137, 149, 1, 0, 0, 0, 138, 139, 5, 69, 0, 0, 139, 140, 5, 2,
		0, 0, 140, 141, 3, 14, 7, 0, 141, 142, 5, 3, 0, 0, 142, 143, 5, 70, 0,
		0, 143, 149, 1, 0, 0, 0, 144, 145, 5, 69, 0, 0, 145, 146, 5, 2, 0, 0, 146,
		147, 5, 3, 0, 0, 147, 149, 5, 70, 0, 0, 148, 134, 1, 0, 0, 0, 148, 138,
		1, 0, 0, 0, 148, 144, 1, 0, 0, 0, 149, 13, 1, 0, 0, 0, 150, 151, 3, 16,
		8, 0, 151, 157, 5, 70, 0, 0, 152, 153, 3, 16, 8, 0, 153, 154, 5, 70, 0,
		0, 154, 156, 1, 0, 0, 0, 155, 152, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157,
		155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 15, 1, 0, 0, 0, 159, 157, 1,
		0, 0, 0, 160, 161, 5, 74, 0, 0, 161, 162, 3, 24, 12, 0, 162, 17, 1, 0,
		0, 0, 163, 164, 3, 20, 10, 0, 164, 165, 3, 60, 30, 0, 165, 166, 5, 70,
		0, 0, 166, 19, 1, 0, 0, 0, 167, 168, 5, 68, 0, 0, 168, 169, 5, 74, 0, 0,
		169, 171, 5, 2, 0, 0, 170, 172, 3, 22, 11, 0, 171, 170, 1, 0, 0, 0, 171,
		172, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 177, 5, 3, 0, 0, 174, 176,
		3, 24, 12, 0, 175, 174, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1,
		0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 21, 1, 0, 0, 0, 179, 177, 1, 0, 0,
		0, 180, 185, 3, 10, 5, 0, 181, 182, 5, 67, 0, 0, 182, 184, 3, 10, 5, 0,
		183, 181, 1, 0, 0, 0, 184, 187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185,
		186, 1, 0, 0, 0, 186, 23, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 189, 5,
		2, 0, 0, 189, 190, 3, 24, 12, 0, 190, 191, 5, 3, 0, 0, 191, 197, 1, 0,
		0, 0, 192, 197, 5, 74, 0, 0, 193, 197, 3, 26, 13, 0, 194, 197, 3, 28, 14,
		0, 195, 197, 3, 30, 15, 0, 196, 188, 1, 0, 0, 0, 196, 192, 1, 0, 0, 0,
		196, 193, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 195, 1, 0, 0, 0, 197,
		25, 1, 0, 0, 0, 198, 199, 5, 65, 0, 0, 199, 200, 5, 66, 0, 0, 200, 201,
		3, 24, 12, 0, 201, 27, 1, 0, 0, 0, 202, 203, 5, 65, 0, 0, 203, 204, 5,
		64, 0, 0, 204, 205, 5, 66, 0, 0, 205, 206, 3, 24, 12, 0, 206, 29, 1, 0,
		0, 0, 207, 208, 5, 63, 0, 0, 208, 212, 5, 61, 0, 0, 209, 211, 3, 32, 16,
		0, 210, 209, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212,
		213, 1, 0, 0, 0, 213, 215, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 215, 216,
		5, 62, 0, 0, 216, 31, 1, 0, 0, 0, 217, 218, 3, 10, 5, 0, 218, 224, 5, 70,
		0, 0, 219, 220, 3, 10, 5, 0, 220, 221, 5, 70, 0, 0, 221, 223, 1, 0, 0,
		0, 222, 219, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224,
		225, 1, 0, 0, 0, 225, 33, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 232, 5,
		74, 0, 0, 228, 229, 5, 67, 0, 0, 229, 231, 5, 74, 0, 0, 230, 228, 1, 0,
		0, 0, 231, 234, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0,
		233, 35, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 235, 236, 6, 18, -1, 0, 236,
		246, 3, 40, 20, 0, 237, 238, 5, 53, 0, 0, 238, 246, 3, 36, 18, 4, 239,
		240, 5, 52, 0, 0, 240, 246, 3, 36, 18, 3, 241, 242, 5, 41, 0, 0, 242, 246,
		3, 36, 18, 2, 243, 244, 5, 50, 0, 0, 244, 246, 3, 36, 18, 1, 245, 235,
		1, 0, 0, 0, 245, 237, 1, 0, 0, 0, 245, 239, 1, 0, 0, 0, 245, 241, 1, 0,
		0, 0, 245, 243, 1, 0, 0, 0, 246, 252, 1, 0, 0, 0, 247, 248, 10, 5, 0, 0,
		248, 249, 7, 0, 0, 0, 249, 251, 3, 36, 18, 6, 250, 247, 1, 0, 0, 0, 251,
		254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 37, 1,
		0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 260, 3, 36, 18, 0, 256, 257, 5, 67,
		0, 0, 257, 259, 3, 36, 18, 0, 258, 256, 1, 0, 0, 0, 259, 262, 1, 0, 0,
		0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 39, 1, 0, 0, 0, 262,
		260, 1, 0, 0, 0, 263, 264, 6, 20, -1, 0, 264, 269, 3, 42, 21, 0, 265, 269,
		3, 52, 26, 0, 266, 269, 3, 54, 27, 0, 267, 269, 3, 56, 28, 0, 268, 263,
		1, 0, 0, 0, 268, 265, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 268, 267, 1, 0,
		0, 0, 269, 278, 1, 0, 0, 0, 270, 271, 10, 6, 0, 0, 271, 277, 3, 50, 25,
		0, 272, 273, 10, 5, 0, 0, 273, 277, 3, 46, 23, 0, 274, 275, 10, 4, 0, 0,
		275, 277, 3, 48, 24, 0, 276, 270, 1, 0, 0, 0, 276, 272, 1, 0, 0, 0, 276,
		274, 1, 0, 0, 0, 277, 280, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278, 279,
		1, 0, 0, 0, 279, 41, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 281, 288, 3, 44,
		22, 0, 282, 288, 5, 74, 0, 0, 283, 284, 5, 2, 0, 0, 284, 285, 3, 36, 18,
		0, 285, 286, 5, 3, 0, 0, 286, 288, 1, 0, 0, 0, 287, 281, 1, 0, 0, 0, 287,
		282, 1, 0, 0, 0, 287, 283, 1, 0, 0, 0, 288, 43, 1, 0, 0, 0, 289, 290, 7,
		1, 0, 0, 290, 45, 1, 0, 0, 0, 291, 292, 5, 65, 0, 0, 292, 293, 3, 36, 18,
		0, 293, 294, 5, 66, 0, 0, 294, 47, 1, 0, 0, 0, 295, 299, 5, 2, 0, 0, 296,
		298, 3, 38, 19, 0, 297, 296, 1, 0, 0, 0, 298, 301, 1, 0, 0, 0, 299, 297,
		1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 302, 1, 0, 0, 0, 301, 299, 1, 0,
		0, 0, 302, 303, 5, 3, 0, 0, 303, 49, 1, 0, 0, 0, 304, 305, 5, 33, 0, 0,
		305, 306, 5, 74, 0, 0, 306, 51, 1, 0, 0, 0, 307, 308, 5, 32, 0, 0, 308,
		309, 5, 2, 0, 0, 309, 310, 3, 36, 18, 0, 310, 311, 5, 67, 0, 0, 311, 312,
		3, 36, 18, 0, 312, 313, 5, 3, 0, 0, 313, 53, 1, 0, 0, 0, 314, 315, 5, 31,
		0, 0, 315, 316, 5, 2, 0, 0, 316, 317, 3, 36, 18, 0, 317, 318, 5, 3, 0,
		0, 318, 55, 1, 0, 0, 0, 319, 320, 5, 30, 0, 0, 320, 321, 5, 2, 0, 0, 321,
		322, 3, 36, 18, 0, 322, 323, 5, 3, 0, 0, 323, 57, 1, 0, 0, 0, 324, 326,
		3, 62, 31, 0, 325, 324, 1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327, 325, 1,
		0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 59, 1, 0, 0, 0, 329, 327, 1, 0, 0,
		0, 330, 331, 5, 61, 0, 0, 331, 332, 3, 58, 29, 0, 332, 333, 5, 62, 0, 0,
		333, 61, 1, 0, 0, 0, 334, 335, 5, 29, 0, 0, 335, 337, 5, 2, 0, 0, 336,
		338, 3, 38, 19, 0, 337, 336, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 338, 339,
		1, 0, 0, 0, 339, 340, 5, 3, 0, 0, 340, 375, 5, 70, 0, 0, 341, 342, 5, 28,
		0, 0, 342, 344, 5, 2, 0, 0, 343, 345, 3, 38, 19, 0, 344, 343, 1, 0, 0,
		0, 344, 345, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 347, 5, 3, 0, 0, 347,
		375, 5, 70, 0, 0, 348, 350, 5, 27, 0, 0, 349, 351, 3, 36, 18, 0, 350, 349,
		1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 375, 5, 70,
		0, 0, 353, 354, 5, 26, 0, 0, 354, 375, 5, 70, 0, 0, 355, 356, 5, 24, 0,
		0, 356, 375, 5, 70, 0, 0, 357, 358, 3, 64, 32, 0, 358, 359, 5, 70, 0, 0,
		359, 375, 1, 0, 0, 0, 360, 361, 3, 60, 30, 0, 361, 362, 5, 70, 0, 0, 362,
		375, 1, 0, 0, 0, 363, 364, 3, 72, 36, 0, 364, 365, 5, 70, 0, 0, 365, 375,
		1, 0, 0, 0, 366, 367, 3, 68, 34, 0, 367, 368, 5, 70, 0, 0, 368, 375, 1,
		0, 0, 0, 369, 370, 3, 70, 35, 0, 370, 371, 5, 70, 0, 0, 371, 375, 1, 0,
		0, 0, 372, 375, 3, 12, 6, 0, 373, 375, 3, 4, 2, 0, 374, 334, 1, 0, 0, 0,
		374, 341, 1, 0, 0, 0, 374, 348, 1, 0, 0, 0, 374, 353, 1, 0, 0, 0, 374,
		355, 1, 0, 0, 0, 374, 357, 1, 0, 0, 0, 374, 360, 1, 0, 0, 0, 374, 363,
		1, 0, 0, 0, 374, 366, 1, 0, 0, 0, 374, 369, 1, 0, 0, 0, 374, 372, 1, 0,
		0, 0, 374, 373, 1, 0, 0, 0, 375, 63, 1, 0, 0, 0, 376, 378, 3, 36, 18, 0,
		377, 379, 7, 2, 0, 0, 378, 377, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379,
		386, 1, 0, 0, 0, 380, 386, 3, 66, 33, 0, 381, 382, 3, 38, 19, 0, 382, 383,
		5, 25, 0, 0, 383, 384, 3, 38, 19, 0, 384, 386, 1, 0, 0, 0, 385, 376, 1,
		0, 0, 0, 385, 380, 1, 0, 0, 0, 385, 381, 1, 0, 0, 0, 386, 65, 1, 0, 0,
		0, 387, 388, 3, 38, 19, 0, 388, 389, 5, 72, 0, 0, 389, 390, 3, 38, 19,
		0, 390, 396, 1, 0, 0, 0, 391, 392, 3, 36, 18, 0, 392, 393, 7, 3, 0, 0,
		393, 394, 3, 36, 18, 0, 394, 396, 1, 0, 0, 0, 395, 387, 1, 0, 0, 0, 395,
		391, 1, 0, 0, 0, 396, 67, 1, 0, 0, 0, 397, 398, 5, 9, 0, 0, 398, 399, 3,
		36, 18, 0, 399, 400, 3, 60, 30, 0, 400, 430, 1, 0, 0, 0, 401, 402, 5, 9,
		0, 0, 402, 403, 3, 36, 18, 0, 403, 404, 3, 60, 30, 0, 404, 405, 5, 10,
		0, 0, 405, 406, 3, 68, 34, 0, 406, 430, 1, 0, 0, 0, 407, 408, 5, 9, 0,
		0, 408, 409, 3, 64, 32, 0, 409, 410, 5, 70, 0, 0, 410, 411, 3, 36, 18,
		0, 411, 412, 3, 60, 30, 0, 412, 430, 1, 0, 0, 0, 413, 414, 5, 9, 0, 0,
		414, 415, 3, 64, 32, 0, 415, 416, 5, 70, 0, 0, 416, 417, 3, 36, 18, 0,
		417, 418, 3, 60, 30, 0, 418, 419, 5, 10, 0, 0, 419, 420, 3, 68, 34, 0,
		420, 430, 1, 0, 0, 0, 421, 422, 5, 9, 0, 0, 422, 423, 3, 64, 32, 0, 423,
		424, 5, 70, 0, 0, 424, 425, 3, 36, 18, 0, 425, 426, 3, 60, 30, 0, 426,
		427, 5, 10, 0, 0, 427, 428, 3, 60, 30, 0, 428, 430, 1, 0, 0, 0, 429, 397,
		1, 0, 0, 0, 429, 401, 1, 0, 0, 0, 429, 407, 1, 0, 0, 0, 429, 413, 1, 0,
		0, 0, 429, 421, 1, 0, 0, 0, 430, 69, 1, 0, 0, 0, 431, 432, 5, 8, 0, 0,
		432, 453, 3, 60, 30, 0, 433, 434, 5, 8, 0, 0, 434, 435, 3, 36, 18, 0, 435,
		436, 3, 60, 30, 0, 436, 453, 1, 0, 0, 0, 437, 438, 5, 8, 0, 0, 438, 439,
		3, 64, 32, 0, 439, 440, 5, 70, 0, 0, 440, 441, 3, 36, 18, 0, 441, 442,
		5, 70, 0, 0, 442, 443, 3, 64, 32, 0, 443, 444, 3, 60, 30, 0, 444, 453,
		1, 0, 0, 0, 445, 446, 5, 8, 0, 0, 446, 447, 3, 64, 32, 0, 447, 448, 5,
		70, 0, 0, 448, 449, 5, 70, 0, 0, 449, 450, 3, 64, 32, 0, 450, 451, 3, 60,
		30, 0, 451, 453, 1, 0, 0, 0, 452, 431, 1, 0, 0, 0, 452, 433, 1, 0, 0, 0,
		452, 437, 1, 0, 0, 0, 452, 445, 1, 0, 0, 0, 453, 71, 1, 0, 0, 0, 454, 455,
		5, 7, 0, 0, 455, 456, 3, 64, 32, 0, 456, 457, 5, 70, 0, 0, 457, 458, 3,
		36, 18, 0, 458, 459, 5, 61, 0, 0, 459, 460, 3, 74, 37, 0, 460, 461, 5,
		62, 0, 0, 461, 481, 1, 0, 0, 0, 462, 463, 5, 7, 0, 0, 463, 464, 3, 36,
		18, 0, 464, 465, 5, 61, 0, 0, 465, 466, 3, 74, 37, 0, 466, 467, 5, 62,
		0, 0, 467, 481, 1, 0, 0, 0, 468, 469, 5, 7, 0, 0, 469, 470, 3, 64, 32,
		0, 470, 471, 5, 70, 0, 0, 471, 472, 5, 61, 0, 0, 472, 473, 3, 74, 37, 0,
		473, 474, 5, 62, 0, 0, 474, 481, 1, 0, 0, 0, 475, 476, 5, 7, 0, 0, 476,
		477, 5, 61, 0, 0, 477, 478, 3, 74, 37, 0, 478, 479, 5, 62, 0, 0, 479, 481,
		1, 0, 0, 0, 480, 454, 1, 0, 0, 0, 480, 462, 1, 0, 0, 0, 480, 468, 1, 0,
		0, 0, 480, 475, 1, 0, 0, 0, 481, 73, 1, 0, 0, 0, 482, 483, 3, 76, 38, 0,
		483, 484, 3, 74, 37, 0, 484, 486, 1, 0, 0, 0, 485, 482, 1, 0, 0, 0, 486,
		489, 1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 488, 75, 1,
		0, 0, 0, 489, 487, 1, 0, 0, 0, 490, 491, 3, 78, 39, 0, 491, 492, 5, 6,
		0, 0, 492, 493, 3, 58, 29, 0, 493, 77, 1, 0, 0, 0, 494, 495, 5, 5, 0, 0,
		495, 498, 3, 38, 19, 0, 496, 498, 5, 4, 0, 0, 497, 494, 1, 0, 0, 0, 497,
		496, 1, 0, 0, 0, 498, 79, 1, 0, 0, 0, 35, 88, 90, 107, 116, 129, 148, 157,
		171, 177, 185, 196, 212, 224, 232, 245, 252, 260, 268, 276, 278, 287, 299,
		327, 337, 344, 350, 374, 378, 385, 395, 429, 452, 480, 487, 497,
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

// MinigoParserInit initializes any static state used to implement MinigoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMinigoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MinigoParserInit() {
	staticData := &MinigoParserStaticData
	staticData.once.Do(minigoParserInit)
}

// NewMinigoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMinigoParser(input antlr.TokenStream) *MinigoParser {
	MinigoParserInit()
	this := new(MinigoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MinigoParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Minigo.g4"

	return this
}

// MinigoParser tokens.
const (
	MinigoParserEOF                      = antlr.TokenEOF
	MinigoParserCOMMENT                  = 1
	MinigoParserLEFTPARENTHESIS          = 2
	MinigoParserRIGHTPARENTHESIS         = 3
	MinigoParserDEFAULT                  = 4
	MinigoParserCASE                     = 5
	MinigoParserCOLON                    = 6
	MinigoParserSWITCH                   = 7
	MinigoParserFOR                      = 8
	MinigoParserIF                       = 9
	MinigoParserELSE                     = 10
	MinigoParserIDIV                     = 11
	MinigoParserIMOD                     = 12
	MinigoParserIANDXOR                  = 13
	MinigoParserILEFTSHIFT               = 14
	MinigoParserIRIGHTSHIFT              = 15
	MinigoParserIXOR                     = 16
	MinigoParserIMUL                     = 17
	MinigoParserIOR                      = 18
	MinigoParserISUB                     = 19
	MinigoParserIAND                     = 20
	MinigoParserIADD                     = 21
	MinigoParserPOSTINC                  = 22
	MinigoParserPOSTDEC                  = 23
	MinigoParserCONTINUE                 = 24
	MinigoParserWALRUS                   = 25
	MinigoParserBREAK                    = 26
	MinigoParserRETURN                   = 27
	MinigoParserPRINTLN                  = 28
	MinigoParserPRINT                    = 29
	MinigoParserCAP                      = 30
	MinigoParserLEN                      = 31
	MinigoParserAPPEND                   = 32
	MinigoParserDOT                      = 33
	MinigoParserWHITESPACE               = 34
	MinigoParserNEWLINE                  = 35
	MinigoParserINTERPRETEDSTRINGLITERAL = 36
	MinigoParserRAWSTRINGLITERAL         = 37
	MinigoParserESCAPEDSEQUENCES         = 38
	MinigoParserRUNELITERAL              = 39
	MinigoParserFLOATLITERAL             = 40
	MinigoParserNOT                      = 41
	MinigoParserOR                       = 42
	MinigoParserAND                      = 43
	MinigoParserGREATERTHANEQUAL         = 44
	MinigoParserLESSTHANEQUAL            = 45
	MinigoParserGREATERTHAN              = 46
	MinigoParserLESSTHAN                 = 47
	MinigoParserNEGATION                 = 48
	MinigoParserCOMPARISON               = 49
	MinigoParserCARET                    = 50
	MinigoParserPIPE                     = 51
	MinigoParserMINUS                    = 52
	MinigoParserPLUS                     = 53
	MinigoParserAMPERSANDCARET           = 54
	MinigoParserAMPERSAND                = 55
	MinigoParserRIGHTSHIFT               = 56
	MinigoParserLEFTSHIFT                = 57
	MinigoParserMOD                      = 58
	MinigoParserDIV                      = 59
	MinigoParserTIMES                    = 60
	MinigoParserLEFTCURLYBRACE           = 61
	MinigoParserRIGHTCURLYBRACE          = 62
	MinigoParserSTRUCT                   = 63
	MinigoParserINTLITERAL               = 64
	MinigoParserLEFTBRACKET              = 65
	MinigoParserRIGHTBRACKET             = 66
	MinigoParserCOMMA                    = 67
	MinigoParserFUNC                     = 68
	MinigoParserTYPE                     = 69
	MinigoParserSEMICOLON                = 70
	MinigoParserVAR                      = 71
	MinigoParserEQUALS                   = 72
	MinigoParserPACKAGE                  = 73
	MinigoParserIDENTIFIER               = 74
)

// MinigoParser rules.
const (
	MinigoParserRULE_root                     = 0
	MinigoParserRULE_topDeclarationList       = 1
	MinigoParserRULE_variableDecl             = 2
	MinigoParserRULE_innerVarDecls            = 3
	MinigoParserRULE_singleVarDecl            = 4
	MinigoParserRULE_singleVarDeclNoExps      = 5
	MinigoParserRULE_typeDecl                 = 6
	MinigoParserRULE_innerTypeDecls           = 7
	MinigoParserRULE_singleTypeDecl           = 8
	MinigoParserRULE_funcDecl                 = 9
	MinigoParserRULE_funcFrontDecl            = 10
	MinigoParserRULE_funcArgsDecls            = 11
	MinigoParserRULE_declType                 = 12
	MinigoParserRULE_sliceDeclType            = 13
	MinigoParserRULE_arrayDeclType            = 14
	MinigoParserRULE_structDeclType           = 15
	MinigoParserRULE_structMemDecls           = 16
	MinigoParserRULE_identifierList           = 17
	MinigoParserRULE_expression               = 18
	MinigoParserRULE_expressionList           = 19
	MinigoParserRULE_primaryExpression        = 20
	MinigoParserRULE_operand                  = 21
	MinigoParserRULE_literal                  = 22
	MinigoParserRULE_index                    = 23
	MinigoParserRULE_arguments                = 24
	MinigoParserRULE_selector                 = 25
	MinigoParserRULE_appendExpression         = 26
	MinigoParserRULE_lengthExpression         = 27
	MinigoParserRULE_capExpression            = 28
	MinigoParserRULE_statementList            = 29
	MinigoParserRULE_block                    = 30
	MinigoParserRULE_statement                = 31
	MinigoParserRULE_simpleStatement          = 32
	MinigoParserRULE_assignmentStatement      = 33
	MinigoParserRULE_ifStatement              = 34
	MinigoParserRULE_loop                     = 35
	MinigoParserRULE_switch                   = 36
	MinigoParserRULE_expressionCaseClauseList = 37
	MinigoParserRULE_expressionCaseClause     = 38
	MinigoParserRULE_expressionSwitchCase     = 39
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	TopDeclarationList() ITopDeclarationListContext

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(MinigoParserPACKAGE, 0)
}

func (s *RootContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MinigoParserIDENTIFIER, 0)
}

func (s *RootContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *RootContext) TopDeclarationList() ITopDeclarationListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopDeclarationListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopDeclarationListContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MinigoParserRULE_root)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(MinigoParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(MinigoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(MinigoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.TopDeclarationList()
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

// ITopDeclarationListContext is an interface to support dynamic dispatch.
type ITopDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariableDecl() []IVariableDeclContext
	VariableDecl(i int) IVariableDeclContext
	AllTypeDecl() []ITypeDeclContext
	TypeDecl(i int) ITypeDeclContext
	AllFuncDecl() []IFuncDeclContext
	FuncDecl(i int) IFuncDeclContext

	// IsTopDeclarationListContext differentiates from other interfaces.
	IsTopDeclarationListContext()
}

type TopDeclarationListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopDeclarationListContext() *TopDeclarationListContext {
	var p = new(TopDeclarationListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_topDeclarationList
	return p
}

func InitEmptyTopDeclarationListContext(p *TopDeclarationListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_topDeclarationList
}

func (*TopDeclarationListContext) IsTopDeclarationListContext() {}

func NewTopDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopDeclarationListContext {
	var p = new(TopDeclarationListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_topDeclarationList

	return p
}

func (s *TopDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *TopDeclarationListContext) AllVariableDecl() []IVariableDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclContext); ok {
			tst[i] = t.(IVariableDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListContext) VariableDecl(i int) IVariableDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
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

	return t.(IVariableDeclContext)
}

func (s *TopDeclarationListContext) AllTypeDecl() []ITypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclContext); ok {
			tst[i] = t.(ITypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListContext) TypeDecl(i int) ITypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
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

	return t.(ITypeDeclContext)
}

func (s *TopDeclarationListContext) AllFuncDecl() []IFuncDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDeclContext); ok {
			len++
		}
	}

	tst := make([]IFuncDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDeclContext); ok {
			tst[i] = t.(IFuncDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListContext) FuncDecl(i int) IFuncDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclContext); ok {
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

	return t.(IFuncDeclContext)
}

func (s *TopDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopDeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterTopDeclarationList(s)
	}
}

func (s *TopDeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitTopDeclarationList(s)
	}
}

func (s *TopDeclarationListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitTopDeclarationList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) TopDeclarationList() (localctx ITopDeclarationListContext) {
	localctx = NewTopDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MinigoParserRULE_topDeclarationList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-68)) & ^0x3f) == 0 && ((int64(1)<<(_la-68))&11) != 0 {
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MinigoParserVAR:
			{
				p.SetState(85)
				p.VariableDecl()
			}

		case MinigoParserTYPE:
			{
				p.SetState(86)
				p.TypeDecl()
			}

		case MinigoParserFUNC:
			{
				p.SetState(87)
				p.FuncDecl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(92)
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

// IVariableDeclContext is an interface to support dynamic dispatch.
type IVariableDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVariableDeclContext differentiates from other interfaces.
	IsVariableDeclContext()
}

type VariableDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclContext() *VariableDeclContext {
	var p = new(VariableDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_variableDecl
	return p
}

func InitEmptyVariableDeclContext(p *VariableDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_variableDecl
}

func (*VariableDeclContext) IsVariableDeclContext() {}

func NewVariableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclContext {
	var p = new(VariableDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_variableDecl

	return p
}

func (s *VariableDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclContext) CopyAll(ctx *VariableDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VariableDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultiVariableDeclarationContext struct {
	VariableDeclContext
}

func NewMultiVariableDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiVariableDeclarationContext {
	var p = new(MultiVariableDeclarationContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *MultiVariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiVariableDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(MinigoParserVAR, 0)
}

func (s *MultiVariableDeclarationContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *MultiVariableDeclarationContext) InnerVarDecls() IInnerVarDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerVarDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerVarDeclsContext)
}

func (s *MultiVariableDeclarationContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *MultiVariableDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *MultiVariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterMultiVariableDeclaration(s)
	}
}

func (s *MultiVariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitMultiVariableDeclaration(s)
	}
}

func (s *MultiVariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitMultiVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmptyVariableDeclarationContext struct {
	VariableDeclContext
}

func NewEmptyVariableDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyVariableDeclarationContext {
	var p = new(EmptyVariableDeclarationContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *EmptyVariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyVariableDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(MinigoParserVAR, 0)
}

func (s *EmptyVariableDeclarationContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *EmptyVariableDeclarationContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *EmptyVariableDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *EmptyVariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterEmptyVariableDeclaration(s)
	}
}

func (s *EmptyVariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitEmptyVariableDeclaration(s)
	}
}

func (s *EmptyVariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitEmptyVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableDeclarationContext struct {
	VariableDeclContext
}

func NewVariableDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(MinigoParserVAR, 0)
}

func (s *VariableDeclarationContext) SingleVarDecl() ISingleVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclContext)
}

func (s *VariableDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) VariableDecl() (localctx IVariableDeclContext) {
	localctx = NewVariableDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MinigoParserRULE_variableDecl)
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVariableDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(MinigoParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.SingleVarDecl()
		}
		{
			p.SetState(95)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewMultiVariableDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.Match(MinigoParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(MinigoParserLEFTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.InnerVarDecls()
		}
		{
			p.SetState(100)
			p.Match(MinigoParserRIGHTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewEmptyVariableDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)
			p.Match(MinigoParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Match(MinigoParserLEFTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(MinigoParserRIGHTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(MinigoParserSEMICOLON)
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

// IInnerVarDeclsContext is an interface to support dynamic dispatch.
type IInnerVarDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleVarDecl() []ISingleVarDeclContext
	SingleVarDecl(i int) ISingleVarDeclContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsInnerVarDeclsContext differentiates from other interfaces.
	IsInnerVarDeclsContext()
}

type InnerVarDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerVarDeclsContext() *InnerVarDeclsContext {
	var p = new(InnerVarDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_innerVarDecls
	return p
}

func InitEmptyInnerVarDeclsContext(p *InnerVarDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_innerVarDecls
}

func (*InnerVarDeclsContext) IsInnerVarDeclsContext() {}

func NewInnerVarDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerVarDeclsContext {
	var p = new(InnerVarDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_innerVarDecls

	return p
}

func (s *InnerVarDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerVarDeclsContext) AllSingleVarDecl() []ISingleVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclContext); ok {
			tst[i] = t.(ISingleVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *InnerVarDeclsContext) SingleVarDecl(i int) ISingleVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
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

	return t.(ISingleVarDeclContext)
}

func (s *InnerVarDeclsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(MinigoParserSEMICOLON)
}

func (s *InnerVarDeclsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, i)
}

func (s *InnerVarDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerVarDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerVarDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterInnerVarDecls(s)
	}
}

func (s *InnerVarDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitInnerVarDecls(s)
	}
}

func (s *InnerVarDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitInnerVarDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) InnerVarDecls() (localctx IInnerVarDeclsContext) {
	localctx = NewInnerVarDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MinigoParserRULE_innerVarDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.SingleVarDecl()
	}
	{
		p.SetState(110)
		p.Match(MinigoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinigoParserIDENTIFIER {
		{
			p.SetState(111)
			p.SingleVarDecl()
		}
		{
			p.SetState(112)
			p.Match(MinigoParserSEMICOLON)
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

// ISingleVarDeclContext is an interface to support dynamic dispatch.
type ISingleVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierList() IIdentifierListContext
	DeclType() IDeclTypeContext
	EQUALS() antlr.TerminalNode
	ExpressionList() IExpressionListContext
	SingleVarDeclNoExps() ISingleVarDeclNoExpsContext

	// IsSingleVarDeclContext differentiates from other interfaces.
	IsSingleVarDeclContext()
}

type SingleVarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleVarDeclContext() *SingleVarDeclContext {
	var p = new(SingleVarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_singleVarDecl
	return p
}

func InitEmptySingleVarDeclContext(p *SingleVarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_singleVarDecl
}

func (*SingleVarDeclContext) IsSingleVarDeclContext() {}

func NewSingleVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclContext {
	var p = new(SingleVarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_singleVarDecl

	return p
}

func (s *SingleVarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleVarDeclContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *SingleVarDeclContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleVarDeclContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(MinigoParserEQUALS, 0)
}

func (s *SingleVarDeclContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *SingleVarDeclContext) SingleVarDeclNoExps() ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *SingleVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleVarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSingleVarDecl(s)
	}
}

func (s *SingleVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSingleVarDecl(s)
	}
}

func (s *SingleVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSingleVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) SingleVarDecl() (localctx ISingleVarDeclContext) {
	localctx = NewSingleVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MinigoParserRULE_singleVarDecl)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.IdentifierList()
		}
		{
			p.SetState(120)
			p.DeclType()
		}
		{
			p.SetState(121)
			p.Match(MinigoParserEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.ExpressionList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.IdentifierList()
		}
		{
			p.SetState(125)
			p.Match(MinigoParserEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.ExpressionList()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
			p.SingleVarDeclNoExps()
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

// ISingleVarDeclNoExpsContext is an interface to support dynamic dispatch.
type ISingleVarDeclNoExpsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierList() IIdentifierListContext
	DeclType() IDeclTypeContext

	// IsSingleVarDeclNoExpsContext differentiates from other interfaces.
	IsSingleVarDeclNoExpsContext()
}

type SingleVarDeclNoExpsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleVarDeclNoExpsContext() *SingleVarDeclNoExpsContext {
	var p = new(SingleVarDeclNoExpsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_singleVarDeclNoExps
	return p
}

func InitEmptySingleVarDeclNoExpsContext(p *SingleVarDeclNoExpsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_singleVarDeclNoExps
}

func (*SingleVarDeclNoExpsContext) IsSingleVarDeclNoExpsContext() {}

func NewSingleVarDeclNoExpsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclNoExpsContext {
	var p = new(SingleVarDeclNoExpsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_singleVarDeclNoExps

	return p
}

func (s *SingleVarDeclNoExpsContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleVarDeclNoExpsContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *SingleVarDeclNoExpsContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleVarDeclNoExpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclNoExpsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleVarDeclNoExpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSingleVarDeclNoExps(s)
	}
}

func (s *SingleVarDeclNoExpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSingleVarDeclNoExps(s)
	}
}

func (s *SingleVarDeclNoExpsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSingleVarDeclNoExps(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) SingleVarDeclNoExps() (localctx ISingleVarDeclNoExpsContext) {
	localctx = NewSingleVarDeclNoExpsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MinigoParserRULE_singleVarDeclNoExps)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.IdentifierList()
	}
	{
		p.SetState(132)
		p.DeclType()
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

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	SingleTypeDecl() ISingleTypeDeclContext
	SEMICOLON() antlr.TerminalNode
	LEFTPARENTHESIS() antlr.TerminalNode
	InnerTypeDecls() IInnerTypeDeclsContext
	RIGHTPARENTHESIS() antlr.TerminalNode

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_typeDecl
	return p
}

func InitEmptyTypeDeclContext(p *TypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_typeDecl
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(MinigoParserTYPE, 0)
}

func (s *TypeDeclContext) SingleTypeDecl() ISingleTypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleTypeDeclContext)
}

func (s *TypeDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *TypeDeclContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *TypeDeclContext) InnerTypeDecls() IInnerTypeDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerTypeDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerTypeDeclsContext)
}

func (s *TypeDeclContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (s *TypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MinigoParserRULE_typeDecl)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Match(MinigoParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.SingleTypeDecl()
		}
		{
			p.SetState(136)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.Match(MinigoParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(MinigoParserLEFTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.InnerTypeDecls()
		}
		{
			p.SetState(141)
			p.Match(MinigoParserRIGHTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.Match(MinigoParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(MinigoParserLEFTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(MinigoParserRIGHTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(MinigoParserSEMICOLON)
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

// IInnerTypeDeclsContext is an interface to support dynamic dispatch.
type IInnerTypeDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleTypeDecl() []ISingleTypeDeclContext
	SingleTypeDecl(i int) ISingleTypeDeclContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsInnerTypeDeclsContext differentiates from other interfaces.
	IsInnerTypeDeclsContext()
}

type InnerTypeDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerTypeDeclsContext() *InnerTypeDeclsContext {
	var p = new(InnerTypeDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_innerTypeDecls
	return p
}

func InitEmptyInnerTypeDeclsContext(p *InnerTypeDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_innerTypeDecls
}

func (*InnerTypeDeclsContext) IsInnerTypeDeclsContext() {}

func NewInnerTypeDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerTypeDeclsContext {
	var p = new(InnerTypeDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_innerTypeDecls

	return p
}

func (s *InnerTypeDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerTypeDeclsContext) AllSingleTypeDecl() []ISingleTypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ISingleTypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleTypeDeclContext); ok {
			tst[i] = t.(ISingleTypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *InnerTypeDeclsContext) SingleTypeDecl(i int) ISingleTypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
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

	return t.(ISingleTypeDeclContext)
}

func (s *InnerTypeDeclsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(MinigoParserSEMICOLON)
}

func (s *InnerTypeDeclsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, i)
}

func (s *InnerTypeDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerTypeDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerTypeDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterInnerTypeDecls(s)
	}
}

func (s *InnerTypeDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitInnerTypeDecls(s)
	}
}

func (s *InnerTypeDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitInnerTypeDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) InnerTypeDecls() (localctx IInnerTypeDeclsContext) {
	localctx = NewInnerTypeDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MinigoParserRULE_innerTypeDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.SingleTypeDecl()
	}
	{
		p.SetState(151)
		p.Match(MinigoParserSEMICOLON)
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

	for _la == MinigoParserIDENTIFIER {
		{
			p.SetState(152)
			p.SingleTypeDecl()
		}
		{
			p.SetState(153)
			p.Match(MinigoParserSEMICOLON)
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

// ISingleTypeDeclContext is an interface to support dynamic dispatch.
type ISingleTypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	DeclType() IDeclTypeContext

	// IsSingleTypeDeclContext differentiates from other interfaces.
	IsSingleTypeDeclContext()
}

type SingleTypeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleTypeDeclContext() *SingleTypeDeclContext {
	var p = new(SingleTypeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_singleTypeDecl
	return p
}

func InitEmptySingleTypeDeclContext(p *SingleTypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_singleTypeDecl
}

func (*SingleTypeDeclContext) IsSingleTypeDeclContext() {}

func NewSingleTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleTypeDeclContext {
	var p = new(SingleTypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_singleTypeDecl

	return p
}

func (s *SingleTypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleTypeDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MinigoParserIDENTIFIER, 0)
}

func (s *SingleTypeDeclContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleTypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleTypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleTypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSingleTypeDecl(s)
	}
}

func (s *SingleTypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSingleTypeDecl(s)
	}
}

func (s *SingleTypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSingleTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) SingleTypeDecl() (localctx ISingleTypeDeclContext) {
	localctx = NewSingleTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MinigoParserRULE_singleTypeDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(MinigoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.DeclType()
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

// IFuncDeclContext is an interface to support dynamic dispatch.
type IFuncDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncFrontDecl() IFuncFrontDeclContext
	Block() IBlockContext
	SEMICOLON() antlr.TerminalNode

	// IsFuncDeclContext differentiates from other interfaces.
	IsFuncDeclContext()
}

type FuncDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclContext() *FuncDeclContext {
	var p = new(FuncDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_funcDecl
	return p
}

func InitEmptyFuncDeclContext(p *FuncDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_funcDecl
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) FuncFrontDecl() IFuncFrontDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncFrontDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncFrontDeclContext)
}

func (s *FuncDeclContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterFuncDecl(s)
	}
}

func (s *FuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitFuncDecl(s)
	}
}

func (s *FuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MinigoParserRULE_funcDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.FuncFrontDecl()
	}
	{
		p.SetState(164)
		p.Block()
	}
	{
		p.SetState(165)
		p.Match(MinigoParserSEMICOLON)
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

// IFuncFrontDeclContext is an interface to support dynamic dispatch.
type IFuncFrontDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LEFTPARENTHESIS() antlr.TerminalNode
	RIGHTPARENTHESIS() antlr.TerminalNode
	FuncArgsDecls() IFuncArgsDeclsContext
	AllDeclType() []IDeclTypeContext
	DeclType(i int) IDeclTypeContext

	// IsFuncFrontDeclContext differentiates from other interfaces.
	IsFuncFrontDeclContext()
}

type FuncFrontDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncFrontDeclContext() *FuncFrontDeclContext {
	var p = new(FuncFrontDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_funcFrontDecl
	return p
}

func InitEmptyFuncFrontDeclContext(p *FuncFrontDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_funcFrontDecl
}

func (*FuncFrontDeclContext) IsFuncFrontDeclContext() {}

func NewFuncFrontDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncFrontDeclContext {
	var p = new(FuncFrontDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_funcFrontDecl

	return p
}

func (s *FuncFrontDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncFrontDeclContext) FUNC() antlr.TerminalNode {
	return s.GetToken(MinigoParserFUNC, 0)
}

func (s *FuncFrontDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MinigoParserIDENTIFIER, 0)
}

func (s *FuncFrontDeclContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *FuncFrontDeclContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *FuncFrontDeclContext) FuncArgsDecls() IFuncArgsDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncArgsDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncArgsDeclsContext)
}

func (s *FuncFrontDeclContext) AllDeclType() []IDeclTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclTypeContext); ok {
			len++
		}
	}

	tst := make([]IDeclTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclTypeContext); ok {
			tst[i] = t.(IDeclTypeContext)
			i++
		}
	}

	return tst
}

func (s *FuncFrontDeclContext) DeclType(i int) IDeclTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
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

	return t.(IDeclTypeContext)
}

func (s *FuncFrontDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncFrontDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncFrontDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterFuncFrontDecl(s)
	}
}

func (s *FuncFrontDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitFuncFrontDecl(s)
	}
}

func (s *FuncFrontDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitFuncFrontDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) FuncFrontDecl() (localctx IFuncFrontDeclContext) {
	localctx = NewFuncFrontDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MinigoParserRULE_funcFrontDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(MinigoParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(MinigoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(MinigoParserLEFTPARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MinigoParserIDENTIFIER {
		{
			p.SetState(170)
			p.FuncArgsDecls()
		}

	}
	{
		p.SetState(173)
		p.Match(MinigoParserRIGHTPARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinigoParserLEFTPARENTHESIS || _la == MinigoParserSTRUCT || _la == MinigoParserLEFTBRACKET || _la == MinigoParserIDENTIFIER {
		{
			p.SetState(174)
			p.DeclType()
		}

		p.SetState(179)
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

// IFuncArgsDeclsContext is an interface to support dynamic dispatch.
type IFuncArgsDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext
	SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncArgsDeclsContext differentiates from other interfaces.
	IsFuncArgsDeclsContext()
}

type FuncArgsDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgsDeclsContext() *FuncArgsDeclsContext {
	var p = new(FuncArgsDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_funcArgsDecls
	return p
}

func InitEmptyFuncArgsDeclsContext(p *FuncArgsDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_funcArgsDecls
}

func (*FuncArgsDeclsContext) IsFuncArgsDeclsContext() {}

func NewFuncArgsDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgsDeclsContext {
	var p = new(FuncArgsDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_funcArgsDecls

	return p
}

func (s *FuncArgsDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgsDeclsContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclNoExpsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			tst[i] = t.(ISingleVarDeclNoExpsContext)
			i++
		}
	}

	return tst
}

func (s *FuncArgsDeclsContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
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

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *FuncArgsDeclsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MinigoParserCOMMA)
}

func (s *FuncArgsDeclsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MinigoParserCOMMA, i)
}

func (s *FuncArgsDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgsDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgsDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterFuncArgsDecls(s)
	}
}

func (s *FuncArgsDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitFuncArgsDecls(s)
	}
}

func (s *FuncArgsDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitFuncArgsDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) FuncArgsDecls() (localctx IFuncArgsDeclsContext) {
	localctx = NewFuncArgsDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MinigoParserRULE_funcArgsDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.SingleVarDeclNoExps()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinigoParserCOMMA {
		{
			p.SetState(181)
			p.Match(MinigoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.SingleVarDeclNoExps()
		}

		p.SetState(187)
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

// IDeclTypeContext is an interface to support dynamic dispatch.
type IDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFTPARENTHESIS() antlr.TerminalNode
	DeclType() IDeclTypeContext
	RIGHTPARENTHESIS() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	SliceDeclType() ISliceDeclTypeContext
	ArrayDeclType() IArrayDeclTypeContext
	StructDeclType() IStructDeclTypeContext

	// IsDeclTypeContext differentiates from other interfaces.
	IsDeclTypeContext()
}

type DeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclTypeContext() *DeclTypeContext {
	var p = new(DeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_declType
	return p
}

func InitEmptyDeclTypeContext(p *DeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_declType
}

func (*DeclTypeContext) IsDeclTypeContext() {}

func NewDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclTypeContext {
	var p = new(DeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_declType

	return p
}

func (s *DeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclTypeContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *DeclTypeContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *DeclTypeContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *DeclTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MinigoParserIDENTIFIER, 0)
}

func (s *DeclTypeContext) SliceDeclType() ISliceDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceDeclTypeContext)
}

func (s *DeclTypeContext) ArrayDeclType() IArrayDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayDeclTypeContext)
}

func (s *DeclTypeContext) StructDeclType() IStructDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclTypeContext)
}

func (s *DeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterDeclType(s)
	}
}

func (s *DeclTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitDeclType(s)
	}
}

func (s *DeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) DeclType() (localctx IDeclTypeContext) {
	localctx = NewDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MinigoParserRULE_declType)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(MinigoParserLEFTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.DeclType()
		}
		{
			p.SetState(190)
			p.Match(MinigoParserRIGHTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.Match(MinigoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(193)
			p.SliceDeclType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(194)
			p.ArrayDeclType()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(195)
			p.StructDeclType()
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

// ISliceDeclTypeContext is an interface to support dynamic dispatch.
type ISliceDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFTBRACKET() antlr.TerminalNode
	RIGHTBRACKET() antlr.TerminalNode
	DeclType() IDeclTypeContext

	// IsSliceDeclTypeContext differentiates from other interfaces.
	IsSliceDeclTypeContext()
}

type SliceDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceDeclTypeContext() *SliceDeclTypeContext {
	var p = new(SliceDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_sliceDeclType
	return p
}

func InitEmptySliceDeclTypeContext(p *SliceDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_sliceDeclType
}

func (*SliceDeclTypeContext) IsSliceDeclTypeContext() {}

func NewSliceDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceDeclTypeContext {
	var p = new(SliceDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_sliceDeclType

	return p
}

func (s *SliceDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceDeclTypeContext) LEFTBRACKET() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTBRACKET, 0)
}

func (s *SliceDeclTypeContext) RIGHTBRACKET() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTBRACKET, 0)
}

func (s *SliceDeclTypeContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SliceDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceDeclTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSliceDeclType(s)
	}
}

func (s *SliceDeclTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSliceDeclType(s)
	}
}

func (s *SliceDeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSliceDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) SliceDeclType() (localctx ISliceDeclTypeContext) {
	localctx = NewSliceDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MinigoParserRULE_sliceDeclType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(MinigoParserLEFTBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Match(MinigoParserRIGHTBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.DeclType()
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

// IArrayDeclTypeContext is an interface to support dynamic dispatch.
type IArrayDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFTBRACKET() antlr.TerminalNode
	INTLITERAL() antlr.TerminalNode
	RIGHTBRACKET() antlr.TerminalNode
	DeclType() IDeclTypeContext

	// IsArrayDeclTypeContext differentiates from other interfaces.
	IsArrayDeclTypeContext()
}

type ArrayDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayDeclTypeContext() *ArrayDeclTypeContext {
	var p = new(ArrayDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_arrayDeclType
	return p
}

func InitEmptyArrayDeclTypeContext(p *ArrayDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_arrayDeclType
}

func (*ArrayDeclTypeContext) IsArrayDeclTypeContext() {}

func NewArrayDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDeclTypeContext {
	var p = new(ArrayDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_arrayDeclType

	return p
}

func (s *ArrayDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDeclTypeContext) LEFTBRACKET() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTBRACKET, 0)
}

func (s *ArrayDeclTypeContext) INTLITERAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserINTLITERAL, 0)
}

func (s *ArrayDeclTypeContext) RIGHTBRACKET() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTBRACKET, 0)
}

func (s *ArrayDeclTypeContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *ArrayDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDeclTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterArrayDeclType(s)
	}
}

func (s *ArrayDeclTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitArrayDeclType(s)
	}
}

func (s *ArrayDeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitArrayDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) ArrayDeclType() (localctx IArrayDeclTypeContext) {
	localctx = NewArrayDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MinigoParserRULE_arrayDeclType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(MinigoParserLEFTBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.Match(MinigoParserINTLITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.Match(MinigoParserRIGHTBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.DeclType()
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

// IStructDeclTypeContext is an interface to support dynamic dispatch.
type IStructDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	LEFTCURLYBRACE() antlr.TerminalNode
	RIGHTCURLYBRACE() antlr.TerminalNode
	AllStructMemDecls() []IStructMemDeclsContext
	StructMemDecls(i int) IStructMemDeclsContext

	// IsStructDeclTypeContext differentiates from other interfaces.
	IsStructDeclTypeContext()
}

type StructDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclTypeContext() *StructDeclTypeContext {
	var p = new(StructDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_structDeclType
	return p
}

func InitEmptyStructDeclTypeContext(p *StructDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_structDeclType
}

func (*StructDeclTypeContext) IsStructDeclTypeContext() {}

func NewStructDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclTypeContext {
	var p = new(StructDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_structDeclType

	return p
}

func (s *StructDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclTypeContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(MinigoParserSTRUCT, 0)
}

func (s *StructDeclTypeContext) LEFTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTCURLYBRACE, 0)
}

func (s *StructDeclTypeContext) RIGHTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTCURLYBRACE, 0)
}

func (s *StructDeclTypeContext) AllStructMemDecls() []IStructMemDeclsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructMemDeclsContext); ok {
			len++
		}
	}

	tst := make([]IStructMemDeclsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructMemDeclsContext); ok {
			tst[i] = t.(IStructMemDeclsContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclTypeContext) StructMemDecls(i int) IStructMemDeclsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructMemDeclsContext); ok {
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

	return t.(IStructMemDeclsContext)
}

func (s *StructDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterStructDeclType(s)
	}
}

func (s *StructDeclTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitStructDeclType(s)
	}
}

func (s *StructDeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitStructDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) StructDeclType() (localctx IStructDeclTypeContext) {
	localctx = NewStructDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MinigoParserRULE_structDeclType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(MinigoParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(MinigoParserLEFTCURLYBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinigoParserIDENTIFIER {
		{
			p.SetState(209)
			p.StructMemDecls()
		}

		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(215)
		p.Match(MinigoParserRIGHTCURLYBRACE)
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

// IStructMemDeclsContext is an interface to support dynamic dispatch.
type IStructMemDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext
	SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsStructMemDeclsContext differentiates from other interfaces.
	IsStructMemDeclsContext()
}

type StructMemDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructMemDeclsContext() *StructMemDeclsContext {
	var p = new(StructMemDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_structMemDecls
	return p
}

func InitEmptyStructMemDeclsContext(p *StructMemDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_structMemDecls
}

func (*StructMemDeclsContext) IsStructMemDeclsContext() {}

func NewStructMemDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructMemDeclsContext {
	var p = new(StructMemDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_structMemDecls

	return p
}

func (s *StructMemDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *StructMemDeclsContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclNoExpsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			tst[i] = t.(ISingleVarDeclNoExpsContext)
			i++
		}
	}

	return tst
}

func (s *StructMemDeclsContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
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

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *StructMemDeclsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(MinigoParserSEMICOLON)
}

func (s *StructMemDeclsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, i)
}

func (s *StructMemDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructMemDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterStructMemDecls(s)
	}
}

func (s *StructMemDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitStructMemDecls(s)
	}
}

func (s *StructMemDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitStructMemDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) StructMemDecls() (localctx IStructMemDeclsContext) {
	localctx = NewStructMemDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MinigoParserRULE_structMemDecls)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.SingleVarDeclNoExps()
	}
	{
		p.SetState(218)
		p.Match(MinigoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(219)
				p.SingleVarDeclNoExps()
			}
			{
				p.SetState(220)
				p.Match(MinigoParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_identifierList
	return p
}

func InitEmptyIdentifierListContext(p *IdentifierListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_identifierList
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(MinigoParserIDENTIFIER)
}

func (s *IdentifierListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(MinigoParserIDENTIFIER, i)
}

func (s *IdentifierListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MinigoParserCOMMA)
}

func (s *IdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MinigoParserCOMMA, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MinigoParserRULE_identifierList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(MinigoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinigoParserCOMMA {
		{
			p.SetState(228)
			p.Match(MinigoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Match(MinigoParserIDENTIFIER)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpression() IPrimaryExpressionContext
	PLUS() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	MINUS() antlr.TerminalNode
	NOT() antlr.TerminalNode
	CARET() antlr.TerminalNode
	TIMES() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	LEFTSHIFT() antlr.TerminalNode
	RIGHTSHIFT() antlr.TerminalNode
	AMPERSAND() antlr.TerminalNode
	AMPERSANDCARET() antlr.TerminalNode
	PIPE() antlr.TerminalNode
	COMPARISON() antlr.TerminalNode
	NEGATION() antlr.TerminalNode
	LESSTHAN() antlr.TerminalNode
	GREATERTHAN() antlr.TerminalNode
	LESSTHANEQUAL() antlr.TerminalNode
	GREATERTHANEQUAL() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MinigoParserPLUS, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MinigoParserMINUS, 0)
}

func (s *ExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(MinigoParserNOT, 0)
}

func (s *ExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(MinigoParserCARET, 0)
}

func (s *ExpressionContext) TIMES() antlr.TerminalNode {
	return s.GetToken(MinigoParserTIMES, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(MinigoParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(MinigoParserMOD, 0)
}

func (s *ExpressionContext) LEFTSHIFT() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTSHIFT, 0)
}

func (s *ExpressionContext) RIGHTSHIFT() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTSHIFT, 0)
}

func (s *ExpressionContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(MinigoParserAMPERSAND, 0)
}

func (s *ExpressionContext) AMPERSANDCARET() antlr.TerminalNode {
	return s.GetToken(MinigoParserAMPERSANDCARET, 0)
}

func (s *ExpressionContext) PIPE() antlr.TerminalNode {
	return s.GetToken(MinigoParserPIPE, 0)
}

func (s *ExpressionContext) COMPARISON() antlr.TerminalNode {
	return s.GetToken(MinigoParserCOMPARISON, 0)
}

func (s *ExpressionContext) NEGATION() antlr.TerminalNode {
	return s.GetToken(MinigoParserNEGATION, 0)
}

func (s *ExpressionContext) LESSTHAN() antlr.TerminalNode {
	return s.GetToken(MinigoParserLESSTHAN, 0)
}

func (s *ExpressionContext) GREATERTHAN() antlr.TerminalNode {
	return s.GetToken(MinigoParserGREATERTHAN, 0)
}

func (s *ExpressionContext) LESSTHANEQUAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserLESSTHANEQUAL, 0)
}

func (s *ExpressionContext) GREATERTHANEQUAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserGREATERTHANEQUAL, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(MinigoParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(MinigoParserOR, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *MinigoParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, MinigoParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinigoParserLEFTPARENTHESIS, MinigoParserCAP, MinigoParserLEN, MinigoParserAPPEND, MinigoParserINTERPRETEDSTRINGLITERAL, MinigoParserRAWSTRINGLITERAL, MinigoParserRUNELITERAL, MinigoParserFLOATLITERAL, MinigoParserINTLITERAL, MinigoParserIDENTIFIER:
		{
			p.SetState(236)
			p.primaryExpression(0)
		}

	case MinigoParserPLUS:
		{
			p.SetState(237)
			p.Match(MinigoParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.expression(4)
		}

	case MinigoParserMINUS:
		{
			p.SetState(239)
			p.Match(MinigoParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.expression(3)
		}

	case MinigoParserNOT:
		{
			p.SetState(241)
			p.Match(MinigoParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.expression(2)
		}

	case MinigoParserCARET:
		{
			p.SetState(243)
			p.Match(MinigoParserCARET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.expression(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, MinigoParserRULE_expression)
			p.SetState(247)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				goto errorExit
			}
			{
				p.SetState(248)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2305838611167182848) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(249)
				p.expression(6)
			}

		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
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

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MinigoParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MinigoParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MinigoParserRULE_expressionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.expression(0)
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinigoParserCOMMA {
		{
			p.SetState(256)
			p.Match(MinigoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)
			p.expression(0)
		}

		p.SetState(262)
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

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext
	AppendExpression() IAppendExpressionContext
	LengthExpression() ILengthExpressionContext
	CapExpression() ICapExpressionContext
	PrimaryExpression() IPrimaryExpressionContext
	Selector() ISelectorContext
	Index() IIndexContext
	Arguments() IArgumentsContext

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExpressionContext) AppendExpression() IAppendExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAppendExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAppendExpressionContext)
}

func (s *PrimaryExpressionContext) LengthExpression() ILengthExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthExpressionContext)
}

func (s *PrimaryExpressionContext) CapExpression() ICapExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICapExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICapExpressionContext)
}

func (s *PrimaryExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpressionContext) Selector() ISelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *PrimaryExpressionContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *PrimaryExpressionContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	return p.primaryExpression(0)
}

func (p *MinigoParser) primaryExpression(_p int) (localctx IPrimaryExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, MinigoParserRULE_primaryExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinigoParserLEFTPARENTHESIS, MinigoParserINTERPRETEDSTRINGLITERAL, MinigoParserRAWSTRINGLITERAL, MinigoParserRUNELITERAL, MinigoParserFLOATLITERAL, MinigoParserINTLITERAL, MinigoParserIDENTIFIER:
		{
			p.SetState(264)
			p.Operand()
		}

	case MinigoParserAPPEND:
		{
			p.SetState(265)
			p.AppendExpression()
		}

	case MinigoParserLEN:
		{
			p.SetState(266)
			p.LengthExpression()
		}

	case MinigoParserCAP:
		{
			p.SetState(267)
			p.CapExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(276)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimaryExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MinigoParserRULE_primaryExpression)
				p.SetState(270)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(271)
					p.Selector()
				}

			case 2:
				localctx = NewPrimaryExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MinigoParserRULE_primaryExpression)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(273)
					p.Index()
				}

			case 3:
				localctx = NewPrimaryExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MinigoParserRULE_primaryExpression)
				p.SetState(274)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(275)
					p.Arguments()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
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

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	IDENTIFIER() antlr.TerminalNode
	LEFTPARENTHESIS() antlr.TerminalNode
	Expression() IExpressionContext
	RIGHTPARENTHESIS() antlr.TerminalNode

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *OperandContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MinigoParserIDENTIFIER, 0)
}

func (s *OperandContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *OperandContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperandContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MinigoParserRULE_operand)
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinigoParserINTERPRETEDSTRINGLITERAL, MinigoParserRAWSTRINGLITERAL, MinigoParserRUNELITERAL, MinigoParserFLOATLITERAL, MinigoParserINTLITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Literal()
		}

	case MinigoParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(282)
			p.Match(MinigoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserLEFTPARENTHESIS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(283)
			p.Match(MinigoParserLEFTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.expression(0)
		}
		{
			p.SetState(285)
			p.Match(MinigoParserRIGHTPARENTHESIS)
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTLITERAL() antlr.TerminalNode
	FLOATLITERAL() antlr.TerminalNode
	RUNELITERAL() antlr.TerminalNode
	RAWSTRINGLITERAL() antlr.TerminalNode
	INTERPRETEDSTRINGLITERAL() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INTLITERAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserINTLITERAL, 0)
}

func (s *LiteralContext) FLOATLITERAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserFLOATLITERAL, 0)
}

func (s *LiteralContext) RUNELITERAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserRUNELITERAL, 0)
}

func (s *LiteralContext) RAWSTRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserRAWSTRINGLITERAL, 0)
}

func (s *LiteralContext) INTERPRETEDSTRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserINTERPRETEDSTRINGLITERAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MinigoParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-36)) & ^0x3f) == 0 && ((int64(1)<<(_la-36))&268435483) != 0) {
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

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFTBRACKET() antlr.TerminalNode
	Expression() IExpressionContext
	RIGHTBRACKET() antlr.TerminalNode

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) LEFTBRACKET() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTBRACKET, 0)
}

func (s *IndexContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexContext) RIGHTBRACKET() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTBRACKET, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MinigoParserRULE_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(MinigoParserLEFTBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.expression(0)
	}
	{
		p.SetState(293)
		p.Match(MinigoParserRIGHTBRACKET)
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

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFTPARENTHESIS() antlr.TerminalNode
	RIGHTPARENTHESIS() antlr.TerminalNode
	AllExpressionList() []IExpressionListContext
	ExpressionList(i int) IExpressionListContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *ArgumentsContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *ArgumentsContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
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

	return t.(IExpressionListContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MinigoParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Match(MinigoParserLEFTPARENTHESIS)
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
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14640760754274308) != 0) || _la == MinigoParserINTLITERAL || _la == MinigoParserIDENTIFIER {
		{
			p.SetState(296)
			p.ExpressionList()
		}

		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(302)
		p.Match(MinigoParserRIGHTPARENTHESIS)
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

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) DOT() antlr.TerminalNode {
	return s.GetToken(MinigoParserDOT, 0)
}

func (s *SelectorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MinigoParserIDENTIFIER, 0)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (s *SelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MinigoParserRULE_selector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(MinigoParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.Match(MinigoParserIDENTIFIER)
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

// IAppendExpressionContext is an interface to support dynamic dispatch.
type IAppendExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	APPEND() antlr.TerminalNode
	LEFTPARENTHESIS() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COMMA() antlr.TerminalNode
	RIGHTPARENTHESIS() antlr.TerminalNode

	// IsAppendExpressionContext differentiates from other interfaces.
	IsAppendExpressionContext()
}

type AppendExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAppendExpressionContext() *AppendExpressionContext {
	var p = new(AppendExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_appendExpression
	return p
}

func InitEmptyAppendExpressionContext(p *AppendExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_appendExpression
}

func (*AppendExpressionContext) IsAppendExpressionContext() {}

func NewAppendExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppendExpressionContext {
	var p = new(AppendExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_appendExpression

	return p
}

func (s *AppendExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AppendExpressionContext) APPEND() antlr.TerminalNode {
	return s.GetToken(MinigoParserAPPEND, 0)
}

func (s *AppendExpressionContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *AppendExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AppendExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *AppendExpressionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MinigoParserCOMMA, 0)
}

func (s *AppendExpressionContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *AppendExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AppendExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterAppendExpression(s)
	}
}

func (s *AppendExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitAppendExpression(s)
	}
}

func (s *AppendExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitAppendExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) AppendExpression() (localctx IAppendExpressionContext) {
	localctx = NewAppendExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MinigoParserRULE_appendExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(MinigoParserAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.Match(MinigoParserLEFTPARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)
		p.expression(0)
	}
	{
		p.SetState(310)
		p.Match(MinigoParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.expression(0)
	}
	{
		p.SetState(312)
		p.Match(MinigoParserRIGHTPARENTHESIS)
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

// ILengthExpressionContext is an interface to support dynamic dispatch.
type ILengthExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEN() antlr.TerminalNode
	LEFTPARENTHESIS() antlr.TerminalNode
	Expression() IExpressionContext
	RIGHTPARENTHESIS() antlr.TerminalNode

	// IsLengthExpressionContext differentiates from other interfaces.
	IsLengthExpressionContext()
}

type LengthExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthExpressionContext() *LengthExpressionContext {
	var p = new(LengthExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_lengthExpression
	return p
}

func InitEmptyLengthExpressionContext(p *LengthExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_lengthExpression
}

func (*LengthExpressionContext) IsLengthExpressionContext() {}

func NewLengthExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthExpressionContext {
	var p = new(LengthExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_lengthExpression

	return p
}

func (s *LengthExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthExpressionContext) LEN() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEN, 0)
}

func (s *LengthExpressionContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *LengthExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LengthExpressionContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *LengthExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterLengthExpression(s)
	}
}

func (s *LengthExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitLengthExpression(s)
	}
}

func (s *LengthExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitLengthExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) LengthExpression() (localctx ILengthExpressionContext) {
	localctx = NewLengthExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MinigoParserRULE_lengthExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(MinigoParserLEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
		p.Match(MinigoParserLEFTPARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(316)
		p.expression(0)
	}
	{
		p.SetState(317)
		p.Match(MinigoParserRIGHTPARENTHESIS)
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

// ICapExpressionContext is an interface to support dynamic dispatch.
type ICapExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CAP() antlr.TerminalNode
	LEFTPARENTHESIS() antlr.TerminalNode
	Expression() IExpressionContext
	RIGHTPARENTHESIS() antlr.TerminalNode

	// IsCapExpressionContext differentiates from other interfaces.
	IsCapExpressionContext()
}

type CapExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCapExpressionContext() *CapExpressionContext {
	var p = new(CapExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_capExpression
	return p
}

func InitEmptyCapExpressionContext(p *CapExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_capExpression
}

func (*CapExpressionContext) IsCapExpressionContext() {}

func NewCapExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CapExpressionContext {
	var p = new(CapExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_capExpression

	return p
}

func (s *CapExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CapExpressionContext) CAP() antlr.TerminalNode {
	return s.GetToken(MinigoParserCAP, 0)
}

func (s *CapExpressionContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *CapExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CapExpressionContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *CapExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CapExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterCapExpression(s)
	}
}

func (s *CapExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitCapExpression(s)
	}
}

func (s *CapExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitCapExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) CapExpression() (localctx ICapExpressionContext) {
	localctx = NewCapExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, MinigoParserRULE_capExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)
		p.Match(MinigoParserCAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(320)
		p.Match(MinigoParserLEFTPARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.expression(0)
	}
	{
		p.SetState(322)
		p.Match(MinigoParserRIGHTPARENTHESIS)
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

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_statementList
	return p
}

func InitEmptyStatementListContext(p *StatementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_statementList
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MinigoParserRULE_statementList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2320483770991379332) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&1185) != 0) {
		{
			p.SetState(324)
			p.Statement()
		}

		p.SetState(329)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFTCURLYBRACE() antlr.TerminalNode
	StatementList() IStatementListContext
	RIGHTCURLYBRACE() antlr.TerminalNode

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LEFTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTCURLYBRACE, 0)
}

func (s *BlockContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockContext) RIGHTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTCURLYBRACE, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MinigoParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(MinigoParserLEFTCURLYBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.StatementList()
	}
	{
		p.SetState(332)
		p.Match(MinigoParserRIGHTCURLYBRACE)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintStatementContext struct {
	StatementContext
}

func NewPrintStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStatementContext {
	var p = new(PrintStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) PRINT() antlr.TerminalNode {
	return s.GetToken(MinigoParserPRINT, 0)
}

func (s *PrintStatementContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *PrintStatementContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *PrintStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *PrintStatementContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (s *PrintStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitPrintStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableDeclStatementContext struct {
	StatementContext
}

func NewVariableDeclStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclStatementContext {
	var p = new(VariableDeclStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *VariableDeclStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclStatementContext) VariableDecl() IVariableDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *VariableDeclStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterVariableDeclStatement(s)
	}
}

func (s *VariableDeclStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitVariableDeclStatement(s)
	}
}

func (s *VariableDeclStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitVariableDeclStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStatementStatementContext struct {
	StatementContext
}

func NewIfStatementStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementStatementContext {
	var p = new(IfStatementStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IfStatementStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementStatementContext) IfStatement() IIfStatementContext {
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

func (s *IfStatementStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *IfStatementStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterIfStatementStatement(s)
	}
}

func (s *IfStatementStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitIfStatementStatement(s)
	}
}

func (s *IfStatementStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitIfStatementStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleStatementStatementContext struct {
	StatementContext
}

func NewSimpleStatementStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleStatementStatementContext {
	var p = new(SimpleStatementStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *SimpleStatementStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementStatementContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *SimpleStatementStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *SimpleStatementStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSimpleStatementStatement(s)
	}
}

func (s *SimpleStatementStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSimpleStatementStatement(s)
	}
}

func (s *SimpleStatementStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSimpleStatementStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlockStatementContext struct {
	StatementContext
}

func NewBlockStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStatementContext {
	var p = new(BlockStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeDeclStatementContext struct {
	StatementContext
}

func NewTypeDeclStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDeclStatementContext {
	var p = new(TypeDeclStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *TypeDeclStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclStatementContext) TypeDecl() ITypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *TypeDeclStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterTypeDeclStatement(s)
	}
}

func (s *TypeDeclStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitTypeDeclStatement(s)
	}
}

func (s *TypeDeclStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitTypeDeclStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStatementContext struct {
	StatementContext
}

func NewBreakStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStatementContext {
	var p = new(BreakStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(MinigoParserBREAK, 0)
}

func (s *BreakStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueStatementContext struct {
	StatementContext
}

func NewContinueStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(MinigoParserCONTINUE, 0)
}

func (s *ContinueStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type LoopStatementContext struct {
	StatementContext
}

func NewLoopStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LoopStatementContext {
	var p = new(LoopStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *LoopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStatementContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *LoopStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *LoopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterLoopStatement(s)
	}
}

func (s *LoopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitLoopStatement(s)
	}
}

func (s *LoopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitLoopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStatementContext struct {
	StatementContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(MinigoParserRETURN, 0)
}

func (s *ReturnStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintlnStatementContext struct {
	StatementContext
}

func NewPrintlnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintlnStatementContext {
	var p = new(PrintlnStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintlnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnStatementContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(MinigoParserPRINTLN, 0)
}

func (s *PrintlnStatementContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *PrintlnStatementContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *PrintlnStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *PrintlnStatementContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *PrintlnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterPrintlnStatement(s)
	}
}

func (s *PrintlnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitPrintlnStatement(s)
	}
}

func (s *PrintlnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitPrintlnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchStatementContext struct {
	StatementContext
}

func NewSwitchStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) Switch_() ISwitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchContext)
}

func (s *SwitchStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSwitchStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MinigoParserRULE_statement)
	var _la int

	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinigoParserPRINT:
		localctx = NewPrintStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(334)
			p.Match(MinigoParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(MinigoParserLEFTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14640760754274308) != 0) || _la == MinigoParserINTLITERAL || _la == MinigoParserIDENTIFIER {
			{
				p.SetState(336)
				p.ExpressionList()
			}

		}
		{
			p.SetState(339)
			p.Match(MinigoParserRIGHTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserPRINTLN:
		localctx = NewPrintlnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(341)
			p.Match(MinigoParserPRINTLN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Match(MinigoParserLEFTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14640760754274308) != 0) || _la == MinigoParserINTLITERAL || _la == MinigoParserIDENTIFIER {
			{
				p.SetState(343)
				p.ExpressionList()
			}

		}
		{
			p.SetState(346)
			p.Match(MinigoParserRIGHTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserRETURN:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(348)
			p.Match(MinigoParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14640760754274308) != 0) || _la == MinigoParserINTLITERAL || _la == MinigoParserIDENTIFIER {
			{
				p.SetState(349)
				p.expression(0)
			}

		}
		{
			p.SetState(352)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserBREAK:
		localctx = NewBreakStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(353)
			p.Match(MinigoParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserCONTINUE:
		localctx = NewContinueStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(355)
			p.Match(MinigoParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserLEFTPARENTHESIS, MinigoParserCAP, MinigoParserLEN, MinigoParserAPPEND, MinigoParserINTERPRETEDSTRINGLITERAL, MinigoParserRAWSTRINGLITERAL, MinigoParserRUNELITERAL, MinigoParserFLOATLITERAL, MinigoParserNOT, MinigoParserCARET, MinigoParserMINUS, MinigoParserPLUS, MinigoParserINTLITERAL, MinigoParserIDENTIFIER:
		localctx = NewSimpleStatementStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(357)
			p.SimpleStatement()
		}
		{
			p.SetState(358)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserLEFTCURLYBRACE:
		localctx = NewBlockStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(360)
			p.Block()
		}
		{
			p.SetState(361)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserSWITCH:
		localctx = NewSwitchStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(363)
			p.Switch_()
		}
		{
			p.SetState(364)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserIF:
		localctx = NewIfStatementStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(366)
			p.IfStatement()
		}
		{
			p.SetState(367)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserFOR:
		localctx = NewLoopStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(369)
			p.Loop()
		}
		{
			p.SetState(370)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserTYPE:
		localctx = NewTypeDeclStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(372)
			p.TypeDecl()
		}

	case MinigoParserVAR:
		localctx = NewVariableDeclStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(373)
			p.VariableDecl()
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

// ISimpleStatementContext is an interface to support dynamic dispatch.
type ISimpleStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	POSTINC() antlr.TerminalNode
	POSTDEC() antlr.TerminalNode
	AssignmentStatement() IAssignmentStatementContext
	AllExpressionList() []IExpressionListContext
	ExpressionList(i int) IExpressionListContext
	WALRUS() antlr.TerminalNode

	// IsSimpleStatementContext differentiates from other interfaces.
	IsSimpleStatementContext()
}

type SimpleStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStatementContext() *SimpleStatementContext {
	var p = new(SimpleStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_simpleStatement
	return p
}

func InitEmptySimpleStatementContext(p *SimpleStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_simpleStatement
}

func (*SimpleStatementContext) IsSimpleStatementContext() {}

func NewSimpleStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStatementContext {
	var p = new(SimpleStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_simpleStatement

	return p
}

func (s *SimpleStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SimpleStatementContext) POSTINC() antlr.TerminalNode {
	return s.GetToken(MinigoParserPOSTINC, 0)
}

func (s *SimpleStatementContext) POSTDEC() antlr.TerminalNode {
	return s.GetToken(MinigoParserPOSTDEC, 0)
}

func (s *SimpleStatementContext) AssignmentStatement() IAssignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *SimpleStatementContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *SimpleStatementContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
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

	return t.(IExpressionListContext)
}

func (s *SimpleStatementContext) WALRUS() antlr.TerminalNode {
	return s.GetToken(MinigoParserWALRUS, 0)
}

func (s *SimpleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSimpleStatement(s)
	}
}

func (s *SimpleStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSimpleStatement(s)
	}
}

func (s *SimpleStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSimpleStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) SimpleStatement() (localctx ISimpleStatementContext) {
	localctx = NewSimpleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MinigoParserRULE_simpleStatement)
	var _la int

	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(376)
			p.expression(0)
		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MinigoParserPOSTINC || _la == MinigoParserPOSTDEC {
			{
				p.SetState(377)
				_la = p.GetTokenStream().LA(1)

				if !(_la == MinigoParserPOSTINC || _la == MinigoParserPOSTDEC) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(380)
			p.AssignmentStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(381)
			p.ExpressionList()
		}
		{
			p.SetState(382)
			p.Match(MinigoParserWALRUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.ExpressionList()
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

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpressionList() []IExpressionListContext
	ExpressionList(i int) IExpressionListContext
	EQUALS() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	IADD() antlr.TerminalNode
	IAND() antlr.TerminalNode
	ISUB() antlr.TerminalNode
	IOR() antlr.TerminalNode
	IMUL() antlr.TerminalNode
	IXOR() antlr.TerminalNode
	ILEFTSHIFT() antlr.TerminalNode
	IRIGHTSHIFT() antlr.TerminalNode
	IANDXOR() antlr.TerminalNode
	IMOD() antlr.TerminalNode
	IDIV() antlr.TerminalNode

	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
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

	return t.(IExpressionListContext)
}

func (s *AssignmentStatementContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(MinigoParserEQUALS, 0)
}

func (s *AssignmentStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *AssignmentStatementContext) IADD() antlr.TerminalNode {
	return s.GetToken(MinigoParserIADD, 0)
}

func (s *AssignmentStatementContext) IAND() antlr.TerminalNode {
	return s.GetToken(MinigoParserIAND, 0)
}

func (s *AssignmentStatementContext) ISUB() antlr.TerminalNode {
	return s.GetToken(MinigoParserISUB, 0)
}

func (s *AssignmentStatementContext) IOR() antlr.TerminalNode {
	return s.GetToken(MinigoParserIOR, 0)
}

func (s *AssignmentStatementContext) IMUL() antlr.TerminalNode {
	return s.GetToken(MinigoParserIMUL, 0)
}

func (s *AssignmentStatementContext) IXOR() antlr.TerminalNode {
	return s.GetToken(MinigoParserIXOR, 0)
}

func (s *AssignmentStatementContext) ILEFTSHIFT() antlr.TerminalNode {
	return s.GetToken(MinigoParserILEFTSHIFT, 0)
}

func (s *AssignmentStatementContext) IRIGHTSHIFT() antlr.TerminalNode {
	return s.GetToken(MinigoParserIRIGHTSHIFT, 0)
}

func (s *AssignmentStatementContext) IANDXOR() antlr.TerminalNode {
	return s.GetToken(MinigoParserIANDXOR, 0)
}

func (s *AssignmentStatementContext) IMOD() antlr.TerminalNode {
	return s.GetToken(MinigoParserIMOD, 0)
}

func (s *AssignmentStatementContext) IDIV() antlr.TerminalNode {
	return s.GetToken(MinigoParserIDIV, 0)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitAssignmentStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MinigoParserRULE_assignmentStatement)
	var _la int

	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(387)
			p.ExpressionList()
		}
		{
			p.SetState(388)
			p.Match(MinigoParserEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(389)
			p.ExpressionList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(391)
			p.expression(0)
		}
		{
			p.SetState(392)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4192256) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(393)
			p.expression(0)
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
	Expression() IExpressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode
	IfStatement() IIfStatementContext
	SimpleStatement() ISimpleStatementContext
	SEMICOLON() antlr.TerminalNode

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
	p.RuleIndex = MinigoParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(MinigoParserIF, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MinigoParserELSE, 0)
}

func (s *IfStatementContext) IfStatement() IIfStatementContext {
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

func (s *IfStatementContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IfStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, MinigoParserRULE_ifStatement)
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(397)
			p.Match(MinigoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.expression(0)
		}
		{
			p.SetState(399)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(401)
			p.Match(MinigoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.expression(0)
		}
		{
			p.SetState(403)
			p.Block()
		}
		{
			p.SetState(404)
			p.Match(MinigoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.IfStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(407)
			p.Match(MinigoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
			p.SimpleStatement()
		}
		{
			p.SetState(409)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.expression(0)
		}
		{
			p.SetState(411)
			p.Block()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(413)
			p.Match(MinigoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.SimpleStatement()
		}
		{
			p.SetState(415)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.expression(0)
		}
		{
			p.SetState(417)
			p.Block()
		}
		{
			p.SetState(418)
			p.Match(MinigoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.IfStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(421)
			p.Match(MinigoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.SimpleStatement()
		}
		{
			p.SetState(423)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.expression(0)
		}
		{
			p.SetState(425)
			p.Block()
		}
		{
			p.SetState(426)
			p.Match(MinigoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.Block()
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

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	Block() IBlockContext
	Expression() IExpressionContext
	AllSimpleStatement() []ISimpleStatementContext
	SimpleStatement(i int) ISimpleStatementContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_loop
	return p
}

func InitEmptyLoopContext(p *LoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_loop
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(MinigoParserFOR, 0)
}

func (s *LoopContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopContext) AllSimpleStatement() []ISimpleStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStatementContext); ok {
			tst[i] = t.(ISimpleStatementContext)
			i++
		}
	}

	return tst
}

func (s *LoopContext) SimpleStatement(i int) ISimpleStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
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

	return t.(ISimpleStatementContext)
}

func (s *LoopContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(MinigoParserSEMICOLON)
}

func (s *LoopContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, i)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (s *LoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, MinigoParserRULE_loop)
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(431)
			p.Match(MinigoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(432)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(433)
			p.Match(MinigoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
			p.expression(0)
		}
		{
			p.SetState(435)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(437)
			p.Match(MinigoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)
			p.SimpleStatement()
		}
		{
			p.SetState(439)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(440)
			p.expression(0)
		}
		{
			p.SetState(441)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(442)
			p.SimpleStatement()
		}
		{
			p.SetState(443)
			p.Block()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(445)
			p.Match(MinigoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)
			p.SimpleStatement()
		}
		{
			p.SetState(447)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(449)
			p.SimpleStatement()
		}
		{
			p.SetState(450)
			p.Block()
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

// ISwitchContext is an interface to support dynamic dispatch.
type ISwitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	SimpleStatement() ISimpleStatementContext
	SEMICOLON() antlr.TerminalNode
	Expression() IExpressionContext
	LEFTCURLYBRACE() antlr.TerminalNode
	ExpressionCaseClauseList() IExpressionCaseClauseListContext
	RIGHTCURLYBRACE() antlr.TerminalNode

	// IsSwitchContext differentiates from other interfaces.
	IsSwitchContext()
}

type SwitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchContext() *SwitchContext {
	var p = new(SwitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_switch
	return p
}

func InitEmptySwitchContext(p *SwitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_switch
}

func (*SwitchContext) IsSwitchContext() {}

func NewSwitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchContext {
	var p = new(SwitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_switch

	return p
}

func (s *SwitchContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(MinigoParserSWITCH, 0)
}

func (s *SwitchContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *SwitchContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *SwitchContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchContext) LEFTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTCURLYBRACE, 0)
}

func (s *SwitchContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SwitchContext) RIGHTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTCURLYBRACE, 0)
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSwitch(s)
	}
}

func (s *SwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSwitch(s)
	}
}

func (s *SwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Switch_() (localctx ISwitchContext) {
	localctx = NewSwitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, MinigoParserRULE_switch)
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(454)
			p.Match(MinigoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(455)
			p.SimpleStatement()
		}
		{
			p.SetState(456)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.expression(0)
		}
		{
			p.SetState(458)
			p.Match(MinigoParserLEFTCURLYBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(459)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(460)
			p.Match(MinigoParserRIGHTCURLYBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(462)
			p.Match(MinigoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)
			p.expression(0)
		}
		{
			p.SetState(464)
			p.Match(MinigoParserLEFTCURLYBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(465)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(466)
			p.Match(MinigoParserRIGHTCURLYBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(468)
			p.Match(MinigoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)
			p.SimpleStatement()
		}
		{
			p.SetState(470)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(471)
			p.Match(MinigoParserLEFTCURLYBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(472)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(473)
			p.Match(MinigoParserRIGHTCURLYBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(475)
			p.Match(MinigoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(476)
			p.Match(MinigoParserLEFTCURLYBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(477)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(478)
			p.Match(MinigoParserRIGHTCURLYBRACE)
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

// IExpressionCaseClauseListContext is an interface to support dynamic dispatch.
type IExpressionCaseClauseListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpressionCaseClause() []IExpressionCaseClauseContext
	ExpressionCaseClause(i int) IExpressionCaseClauseContext
	AllExpressionCaseClauseList() []IExpressionCaseClauseListContext
	ExpressionCaseClauseList(i int) IExpressionCaseClauseListContext

	// IsExpressionCaseClauseListContext differentiates from other interfaces.
	IsExpressionCaseClauseListContext()
}

type ExpressionCaseClauseListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionCaseClauseListContext() *ExpressionCaseClauseListContext {
	var p = new(ExpressionCaseClauseListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_expressionCaseClauseList
	return p
}

func InitEmptyExpressionCaseClauseListContext(p *ExpressionCaseClauseListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_expressionCaseClauseList
}

func (*ExpressionCaseClauseListContext) IsExpressionCaseClauseListContext() {}

func NewExpressionCaseClauseListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseListContext {
	var p = new(ExpressionCaseClauseListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_expressionCaseClauseList

	return p
}

func (s *ExpressionCaseClauseListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionCaseClauseListContext) AllExpressionCaseClause() []IExpressionCaseClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionCaseClauseContext); ok {
			len++
		}
	}

	tst := make([]IExpressionCaseClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionCaseClauseContext); ok {
			tst[i] = t.(IExpressionCaseClauseContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionCaseClauseListContext) ExpressionCaseClause(i int) IExpressionCaseClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseContext); ok {
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

	return t.(IExpressionCaseClauseContext)
}

func (s *ExpressionCaseClauseListContext) AllExpressionCaseClauseList() []IExpressionCaseClauseListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionCaseClauseListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionCaseClauseListContext); ok {
			tst[i] = t.(IExpressionCaseClauseListContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionCaseClauseListContext) ExpressionCaseClauseList(i int) IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
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

	return t.(IExpressionCaseClauseListContext)
}

func (s *ExpressionCaseClauseListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionCaseClauseListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterExpressionCaseClauseList(s)
	}
}

func (s *ExpressionCaseClauseListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitExpressionCaseClauseList(s)
	}
}

func (s *ExpressionCaseClauseListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitExpressionCaseClauseList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) ExpressionCaseClauseList() (localctx IExpressionCaseClauseListContext) {
	localctx = NewExpressionCaseClauseListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, MinigoParserRULE_expressionCaseClauseList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(487)
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
				p.SetState(482)
				p.ExpressionCaseClause()
			}
			{
				p.SetState(483)
				p.ExpressionCaseClauseList()
			}

		}
		p.SetState(489)
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

// IExpressionCaseClauseContext is an interface to support dynamic dispatch.
type IExpressionCaseClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionSwitchCase() IExpressionSwitchCaseContext
	COLON() antlr.TerminalNode
	StatementList() IStatementListContext

	// IsExpressionCaseClauseContext differentiates from other interfaces.
	IsExpressionCaseClauseContext()
}

type ExpressionCaseClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionCaseClauseContext() *ExpressionCaseClauseContext {
	var p = new(ExpressionCaseClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_expressionCaseClause
	return p
}

func InitEmptyExpressionCaseClauseContext(p *ExpressionCaseClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_expressionCaseClause
}

func (*ExpressionCaseClauseContext) IsExpressionCaseClauseContext() {}

func NewExpressionCaseClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseContext {
	var p = new(ExpressionCaseClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_expressionCaseClause

	return p
}

func (s *ExpressionCaseClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionCaseClauseContext) ExpressionSwitchCase() IExpressionSwitchCaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSwitchCaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSwitchCaseContext)
}

func (s *ExpressionCaseClauseContext) COLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserCOLON, 0)
}

func (s *ExpressionCaseClauseContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *ExpressionCaseClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionCaseClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterExpressionCaseClause(s)
	}
}

func (s *ExpressionCaseClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitExpressionCaseClause(s)
	}
}

func (s *ExpressionCaseClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitExpressionCaseClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) ExpressionCaseClause() (localctx IExpressionCaseClauseContext) {
	localctx = NewExpressionCaseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, MinigoParserRULE_expressionCaseClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)
		p.ExpressionSwitchCase()
	}
	{
		p.SetState(491)
		p.Match(MinigoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(492)
		p.StatementList()
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

// IExpressionSwitchCaseContext is an interface to support dynamic dispatch.
type IExpressionSwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	ExpressionList() IExpressionListContext
	DEFAULT() antlr.TerminalNode

	// IsExpressionSwitchCaseContext differentiates from other interfaces.
	IsExpressionSwitchCaseContext()
}

type ExpressionSwitchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionSwitchCaseContext() *ExpressionSwitchCaseContext {
	var p = new(ExpressionSwitchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_expressionSwitchCase
	return p
}

func InitEmptyExpressionSwitchCaseContext(p *ExpressionSwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MinigoParserRULE_expressionSwitchCase
}

func (*ExpressionSwitchCaseContext) IsExpressionSwitchCaseContext() {}

func NewExpressionSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSwitchCaseContext {
	var p = new(ExpressionSwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinigoParserRULE_expressionSwitchCase

	return p
}

func (s *ExpressionSwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionSwitchCaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(MinigoParserCASE, 0)
}

func (s *ExpressionSwitchCaseContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ExpressionSwitchCaseContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(MinigoParserDEFAULT, 0)
}

func (s *ExpressionSwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionSwitchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterExpressionSwitchCase(s)
	}
}

func (s *ExpressionSwitchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitExpressionSwitchCase(s)
	}
}

func (s *ExpressionSwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitExpressionSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) ExpressionSwitchCase() (localctx IExpressionSwitchCaseContext) {
	localctx = NewExpressionSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, MinigoParserRULE_expressionSwitchCase)
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinigoParserCASE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(494)
			p.Match(MinigoParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(495)
			p.ExpressionList()
		}

	case MinigoParserDEFAULT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(496)
			p.Match(MinigoParserDEFAULT)
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

func (p *MinigoParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 20:
		var t *PrimaryExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExpressionContext)
		}
		return p.PrimaryExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MinigoParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MinigoParser) PrimaryExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
