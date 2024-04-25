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
		"", "", "", "'('", "')'", "'default'", "'case'", "':'", "'switch'",
		"'for'", "'if'", "'else'", "'/='", "'%='", "'&^='", "'<<='", "'>>='",
		"'^='", "'*='", "'|='", "'-='", "'&='", "'+='", "'++'", "'--'", "'continue'",
		"':='", "'break'", "'return'", "'println'", "'print'", "'cap'", "'len'",
		"'append'", "'.'", "", "", "", "", "", "", "", "'!'", "'||'", "'&&'",
		"'>='", "'<='", "'>'", "'<'", "'!='", "'=='", "'^'", "'|'", "'-'", "'+'",
		"'&^'", "'&'", "'>>'", "'<<'", "'%'", "'/'", "'*'", "'{'", "'}'", "'struct'",
		"", "'['", "']'", "','", "'func'", "'type'", "';'", "'var'", "'='",
		"'package'",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "MULTILINE_COMMENT", "LEFTPARENTHESIS", "RIGHTPARENTHESIS",
		"DEFAULT", "CASE", "COLON", "SWITCH", "FOR", "IF", "ELSE", "IDIV", "IMOD",
		"IANDXOR", "ILEFTSHIFT", "IRIGHTSHIFT", "IXOR", "IMUL", "IOR", "ISUB",
		"IAND", "IADD", "POSTINC", "POSTDEC", "CONTINUE", "WALRUS", "BREAK",
		"RETURN", "PRINTLN", "PRINT", "CAP", "LEN", "APPEND", "DOT", "WHITESPACE",
		"NEWLINE", "INTERPRETEDSTRINGLITERAL", "RAWSTRINGLITERAL", "ESCAPEDSEQUENCES",
		"RUNELITERAL", "FLOATLITERAL", "NOT", "OR", "AND", "GREATERTHANEQUAL",
		"LESSTHANEQUAL", "GREATERTHAN", "LESSTHAN", "NEGATION", "COMPARISON",
		"CARET", "PIPE", "MINUS", "PLUS", "AMPERSANDCARET", "AMPERSAND", "RIGHTSHIFT",
		"LEFTSHIFT", "MOD", "DIV", "TIMES", "LEFTCURLYBRACE", "RIGHTCURLYBRACE",
		"STRUCT", "INTLITERAL", "LEFTBRACKET", "RIGHTBRACKET", "COMMA", "FUNC",
		"TYPE", "SEMICOLON", "VAR", "EQUALS", "PACKAGE", "IDENTIFIER",
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
		4, 1, 75, 502, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		172, 8, 10, 1, 10, 1, 10, 3, 10, 176, 8, 10, 1, 11, 1, 11, 1, 11, 5, 11,
		181, 8, 11, 10, 11, 12, 11, 184, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 3, 12, 194, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15, 208, 8,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 217, 8, 16,
		10, 16, 12, 16, 220, 9, 16, 1, 17, 1, 17, 1, 17, 5, 17, 225, 8, 17, 10,
		17, 12, 17, 228, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18,
		236, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 5, 18, 247, 8, 18, 10, 18, 12, 18, 250, 9, 18, 1, 19, 1, 19, 1, 19,
		5, 19, 255, 8, 19, 10, 19, 12, 19, 258, 9, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 3, 20, 265, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		5, 20, 273, 8, 20, 10, 20, 12, 20, 276, 9, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 3, 21, 284, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		3, 22, 291, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 5, 24, 299,
		8, 24, 10, 24, 12, 24, 302, 9, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 5, 29, 327, 8, 29, 10,
		29, 12, 29, 330, 9, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31,
		3, 31, 339, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 346, 8, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 352, 8, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 376, 8,
		31, 1, 32, 1, 32, 3, 32, 380, 8, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		3, 32, 387, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 3, 33, 397, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3,
		34, 437, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 3, 35, 449, 8, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 455,
		8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 483, 8, 36, 1, 37, 1,
		37, 1, 37, 5, 37, 488, 8, 37, 10, 37, 12, 37, 491, 9, 37, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 3, 39, 500, 8, 39, 1, 39, 0, 2, 36,
		40, 40, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 0, 5, 1, 0, 43, 44, 1, 0, 45, 50, 1, 0, 51, 61, 1,
		0, 23, 24, 1, 0, 12, 22, 533, 0, 80, 1, 0, 0, 0, 2, 90, 1, 0, 0, 0, 4,
		107, 1, 0, 0, 0, 6, 109, 1, 0, 0, 0, 8, 129, 1, 0, 0, 0, 10, 131, 1, 0,
		0, 0, 12, 148, 1, 0, 0, 0, 14, 150, 1, 0, 0, 0, 16, 160, 1, 0, 0, 0, 18,
		163, 1, 0, 0, 0, 20, 167, 1, 0, 0, 0, 22, 177, 1, 0, 0, 0, 24, 193, 1,
		0, 0, 0, 26, 195, 1, 0, 0, 0, 28, 199, 1, 0, 0, 0, 30, 204, 1, 0, 0, 0,
		32, 211, 1, 0, 0, 0, 34, 221, 1, 0, 0, 0, 36, 235, 1, 0, 0, 0, 38, 251,
		1, 0, 0, 0, 40, 264, 1, 0, 0, 0, 42, 283, 1, 0, 0, 0, 44, 290, 1, 0, 0,
		0, 46, 292, 1, 0, 0, 0, 48, 296, 1, 0, 0, 0, 50, 305, 1, 0, 0, 0, 52, 308,
		1, 0, 0, 0, 54, 315, 1, 0, 0, 0, 56, 320, 1, 0, 0, 0, 58, 328, 1, 0, 0,
		0, 60, 331, 1, 0, 0, 0, 62, 375, 1, 0, 0, 0, 64, 386, 1, 0, 0, 0, 66, 396,
		1, 0, 0, 0, 68, 436, 1, 0, 0, 0, 70, 454, 1, 0, 0, 0, 72, 482, 1, 0, 0,
		0, 74, 489, 1, 0, 0, 0, 76, 492, 1, 0, 0, 0, 78, 499, 1, 0, 0, 0, 80, 81,
		5, 74, 0, 0, 81, 82, 5, 75, 0, 0, 82, 83, 5, 71, 0, 0, 83, 84, 3, 2, 1,
		0, 84, 1, 1, 0, 0, 0, 85, 89, 3, 4, 2, 0, 86, 89, 3, 12, 6, 0, 87, 89,
		3, 18, 9, 0, 88, 85, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 87, 1, 0, 0, 0,
		89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 3, 1, 0,
		0, 0, 92, 90, 1, 0, 0, 0, 93, 94, 5, 72, 0, 0, 94, 95, 3, 8, 4, 0, 95,
		96, 5, 71, 0, 0, 96, 108, 1, 0, 0, 0, 97, 98, 5, 72, 0, 0, 98, 99, 5, 3,
		0, 0, 99, 100, 3, 6, 3, 0, 100, 101, 5, 4, 0, 0, 101, 102, 5, 71, 0, 0,
		102, 108, 1, 0, 0, 0, 103, 104, 5, 72, 0, 0, 104, 105, 5, 3, 0, 0, 105,
		106, 5, 4, 0, 0, 106, 108, 5, 71, 0, 0, 107, 93, 1, 0, 0, 0, 107, 97, 1,
		0, 0, 0, 107, 103, 1, 0, 0, 0, 108, 5, 1, 0, 0, 0, 109, 110, 3, 8, 4, 0,
		110, 116, 5, 71, 0, 0, 111, 112, 3, 8, 4, 0, 112, 113, 5, 71, 0, 0, 113,
		115, 1, 0, 0, 0, 114, 111, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114,
		1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 7, 1, 0, 0, 0, 118, 116, 1, 0, 0,
		0, 119, 120, 3, 34, 17, 0, 120, 121, 3, 24, 12, 0, 121, 122, 5, 73, 0,
		0, 122, 123, 3, 38, 19, 0, 123, 130, 1, 0, 0, 0, 124, 125, 3, 34, 17, 0,
		125, 126, 5, 73, 0, 0, 126, 127, 3, 38, 19, 0, 127, 130, 1, 0, 0, 0, 128,
		130, 3, 10, 5, 0, 129, 119, 1, 0, 0, 0, 129, 124, 1, 0, 0, 0, 129, 128,
		1, 0, 0, 0, 130, 9, 1, 0, 0, 0, 131, 132, 3, 34, 17, 0, 132, 133, 3, 24,
		12, 0, 133, 11, 1, 0, 0, 0, 134, 135, 5, 70, 0, 0, 135, 136, 3, 16, 8,
		0, 136, 137, 5, 71, 0, 0, 137, 149, 1, 0, 0, 0, 138, 139, 5, 70, 0, 0,
		139, 140, 5, 3, 0, 0, 140, 141, 3, 14, 7, 0, 141, 142, 5, 4, 0, 0, 142,
		143, 5, 71, 0, 0, 143, 149, 1, 0, 0, 0, 144, 145, 5, 70, 0, 0, 145, 146,
		5, 3, 0, 0, 146, 147, 5, 4, 0, 0, 147, 149, 5, 71, 0, 0, 148, 134, 1, 0,
		0, 0, 148, 138, 1, 0, 0, 0, 148, 144, 1, 0, 0, 0, 149, 13, 1, 0, 0, 0,
		150, 151, 3, 16, 8, 0, 151, 157, 5, 71, 0, 0, 152, 153, 3, 16, 8, 0, 153,
		154, 5, 71, 0, 0, 154, 156, 1, 0, 0, 0, 155, 152, 1, 0, 0, 0, 156, 159,
		1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 15, 1, 0,
		0, 0, 159, 157, 1, 0, 0, 0, 160, 161, 5, 75, 0, 0, 161, 162, 3, 24, 12,
		0, 162, 17, 1, 0, 0, 0, 163, 164, 3, 20, 10, 0, 164, 165, 3, 60, 30, 0,
		165, 166, 5, 71, 0, 0, 166, 19, 1, 0, 0, 0, 167, 168, 5, 69, 0, 0, 168,
		169, 5, 75, 0, 0, 169, 171, 5, 3, 0, 0, 170, 172, 3, 22, 11, 0, 171, 170,
		1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 175, 5, 4,
		0, 0, 174, 176, 3, 24, 12, 0, 175, 174, 1, 0, 0, 0, 175, 176, 1, 0, 0,
		0, 176, 21, 1, 0, 0, 0, 177, 182, 3, 10, 5, 0, 178, 179, 5, 68, 0, 0, 179,
		181, 3, 10, 5, 0, 180, 178, 1, 0, 0, 0, 181, 184, 1, 0, 0, 0, 182, 180,
		1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 23, 1, 0, 0, 0, 184, 182, 1, 0,
		0, 0, 185, 186, 5, 3, 0, 0, 186, 187, 3, 24, 12, 0, 187, 188, 5, 4, 0,
		0, 188, 194, 1, 0, 0, 0, 189, 194, 5, 75, 0, 0, 190, 194, 3, 26, 13, 0,
		191, 194, 3, 28, 14, 0, 192, 194, 3, 30, 15, 0, 193, 185, 1, 0, 0, 0, 193,
		189, 1, 0, 0, 0, 193, 190, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 192,
		1, 0, 0, 0, 194, 25, 1, 0, 0, 0, 195, 196, 5, 66, 0, 0, 196, 197, 5, 67,
		0, 0, 197, 198, 3, 24, 12, 0, 198, 27, 1, 0, 0, 0, 199, 200, 5, 66, 0,
		0, 200, 201, 5, 65, 0, 0, 201, 202, 5, 67, 0, 0, 202, 203, 3, 24, 12, 0,
		203, 29, 1, 0, 0, 0, 204, 205, 5, 64, 0, 0, 205, 207, 5, 62, 0, 0, 206,
		208, 3, 32, 16, 0, 207, 206, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209,
		1, 0, 0, 0, 209, 210, 5, 63, 0, 0, 210, 31, 1, 0, 0, 0, 211, 212, 3, 10,
		5, 0, 212, 218, 5, 71, 0, 0, 213, 214, 3, 10, 5, 0, 214, 215, 5, 71, 0,
		0, 215, 217, 1, 0, 0, 0, 216, 213, 1, 0, 0, 0, 217, 220, 1, 0, 0, 0, 218,
		216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 33, 1, 0, 0, 0, 220, 218, 1,
		0, 0, 0, 221, 226, 5, 75, 0, 0, 222, 223, 5, 68, 0, 0, 223, 225, 5, 75,
		0, 0, 224, 222, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0,
		226, 227, 1, 0, 0, 0, 227, 35, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 229, 230,
		6, 18, -1, 0, 230, 236, 3, 40, 20, 0, 231, 232, 5, 42, 0, 0, 232, 236,
		3, 36, 18, 2, 233, 234, 5, 51, 0, 0, 234, 236, 3, 36, 18, 1, 235, 229,
		1, 0, 0, 0, 235, 231, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 236, 248, 1, 0,
		0, 0, 237, 238, 10, 6, 0, 0, 238, 239, 7, 0, 0, 0, 239, 247, 3, 36, 18,
		7, 240, 241, 10, 5, 0, 0, 241, 242, 7, 1, 0, 0, 242, 247, 3, 36, 18, 6,
		243, 244, 10, 4, 0, 0, 244, 245, 7, 2, 0, 0, 245, 247, 3, 36, 18, 5, 246,
		237, 1, 0, 0, 0, 246, 240, 1, 0, 0, 0, 246, 243, 1, 0, 0, 0, 247, 250,
		1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 37, 1, 0,
		0, 0, 250, 248, 1, 0, 0, 0, 251, 256, 3, 36, 18, 0, 252, 253, 5, 68, 0,
		0, 253, 255, 3, 36, 18, 0, 254, 252, 1, 0, 0, 0, 255, 258, 1, 0, 0, 0,
		256, 254, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 39, 1, 0, 0, 0, 258, 256,
		1, 0, 0, 0, 259, 260, 6, 20, -1, 0, 260, 265, 3, 42, 21, 0, 261, 265, 3,
		52, 26, 0, 262, 265, 3, 54, 27, 0, 263, 265, 3, 56, 28, 0, 264, 259, 1,
		0, 0, 0, 264, 261, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 264, 263, 1, 0, 0,
		0, 265, 274, 1, 0, 0, 0, 266, 267, 10, 6, 0, 0, 267, 273, 3, 50, 25, 0,
		268, 269, 10, 5, 0, 0, 269, 273, 3, 46, 23, 0, 270, 271, 10, 4, 0, 0, 271,
		273, 3, 48, 24, 0, 272, 266, 1, 0, 0, 0, 272, 268, 1, 0, 0, 0, 272, 270,
		1, 0, 0, 0, 273, 276, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 274, 275, 1, 0,
		0, 0, 275, 41, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0, 277, 284, 3, 44, 22, 0,
		278, 284, 5, 75, 0, 0, 279, 280, 5, 3, 0, 0, 280, 281, 3, 36, 18, 0, 281,
		282, 5, 4, 0, 0, 282, 284, 1, 0, 0, 0, 283, 277, 1, 0, 0, 0, 283, 278,
		1, 0, 0, 0, 283, 279, 1, 0, 0, 0, 284, 43, 1, 0, 0, 0, 285, 291, 5, 65,
		0, 0, 286, 291, 5, 41, 0, 0, 287, 291, 5, 40, 0, 0, 288, 291, 5, 38, 0,
		0, 289, 291, 5, 37, 0, 0, 290, 285, 1, 0, 0, 0, 290, 286, 1, 0, 0, 0, 290,
		287, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 289, 1, 0, 0, 0, 291, 45, 1,
		0, 0, 0, 292, 293, 5, 66, 0, 0, 293, 294, 3, 36, 18, 0, 294, 295, 5, 67,
		0, 0, 295, 47, 1, 0, 0, 0, 296, 300, 5, 3, 0, 0, 297, 299, 3, 38, 19, 0,
		298, 297, 1, 0, 0, 0, 299, 302, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 300,
		301, 1, 0, 0, 0, 301, 303, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 303, 304,
		5, 4, 0, 0, 304, 49, 1, 0, 0, 0, 305, 306, 5, 34, 0, 0, 306, 307, 5, 75,
		0, 0, 307, 51, 1, 0, 0, 0, 308, 309, 5, 33, 0, 0, 309, 310, 5, 3, 0, 0,
		310, 311, 3, 36, 18, 0, 311, 312, 5, 68, 0, 0, 312, 313, 3, 36, 18, 0,
		313, 314, 5, 4, 0, 0, 314, 53, 1, 0, 0, 0, 315, 316, 5, 32, 0, 0, 316,
		317, 5, 3, 0, 0, 317, 318, 3, 36, 18, 0, 318, 319, 5, 4, 0, 0, 319, 55,
		1, 0, 0, 0, 320, 321, 5, 31, 0, 0, 321, 322, 5, 3, 0, 0, 322, 323, 3, 36,
		18, 0, 323, 324, 5, 4, 0, 0, 324, 57, 1, 0, 0, 0, 325, 327, 3, 62, 31,
		0, 326, 325, 1, 0, 0, 0, 327, 330, 1, 0, 0, 0, 328, 326, 1, 0, 0, 0, 328,
		329, 1, 0, 0, 0, 329, 59, 1, 0, 0, 0, 330, 328, 1, 0, 0, 0, 331, 332, 5,
		62, 0, 0, 332, 333, 3, 58, 29, 0, 333, 334, 5, 63, 0, 0, 334, 61, 1, 0,
		0, 0, 335, 336, 5, 30, 0, 0, 336, 338, 5, 3, 0, 0, 337, 339, 3, 38, 19,
		0, 338, 337, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340,
		341, 5, 4, 0, 0, 341, 376, 5, 71, 0, 0, 342, 343, 5, 29, 0, 0, 343, 345,
		5, 3, 0, 0, 344, 346, 3, 38, 19, 0, 345, 344, 1, 0, 0, 0, 345, 346, 1,
		0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 348, 5, 4, 0, 0, 348, 376, 5, 71, 0,
		0, 349, 351, 5, 28, 0, 0, 350, 352, 3, 36, 18, 0, 351, 350, 1, 0, 0, 0,
		351, 352, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 376, 5, 71, 0, 0, 354,
		355, 5, 27, 0, 0, 355, 376, 5, 71, 0, 0, 356, 357, 5, 25, 0, 0, 357, 376,
		5, 71, 0, 0, 358, 359, 3, 64, 32, 0, 359, 360, 5, 71, 0, 0, 360, 376, 1,
		0, 0, 0, 361, 362, 3, 60, 30, 0, 362, 363, 5, 71, 0, 0, 363, 376, 1, 0,
		0, 0, 364, 365, 3, 72, 36, 0, 365, 366, 5, 71, 0, 0, 366, 376, 1, 0, 0,
		0, 367, 368, 3, 68, 34, 0, 368, 369, 5, 71, 0, 0, 369, 376, 1, 0, 0, 0,
		370, 371, 3, 70, 35, 0, 371, 372, 5, 71, 0, 0, 372, 376, 1, 0, 0, 0, 373,
		376, 3, 12, 6, 0, 374, 376, 3, 4, 2, 0, 375, 335, 1, 0, 0, 0, 375, 342,
		1, 0, 0, 0, 375, 349, 1, 0, 0, 0, 375, 354, 1, 0, 0, 0, 375, 356, 1, 0,
		0, 0, 375, 358, 1, 0, 0, 0, 375, 361, 1, 0, 0, 0, 375, 364, 1, 0, 0, 0,
		375, 367, 1, 0, 0, 0, 375, 370, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375,
		374, 1, 0, 0, 0, 376, 63, 1, 0, 0, 0, 377, 379, 3, 36, 18, 0, 378, 380,
		7, 3, 0, 0, 379, 378, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380, 387, 1, 0,
		0, 0, 381, 387, 3, 66, 33, 0, 382, 383, 3, 38, 19, 0, 383, 384, 5, 26,
		0, 0, 384, 385, 3, 38, 19, 0, 385, 387, 1, 0, 0, 0, 386, 377, 1, 0, 0,
		0, 386, 381, 1, 0, 0, 0, 386, 382, 1, 0, 0, 0, 387, 65, 1, 0, 0, 0, 388,
		389, 3, 38, 19, 0, 389, 390, 5, 73, 0, 0, 390, 391, 3, 38, 19, 0, 391,
		397, 1, 0, 0, 0, 392, 393, 3, 36, 18, 0, 393, 394, 7, 4, 0, 0, 394, 395,
		3, 36, 18, 0, 395, 397, 1, 0, 0, 0, 396, 388, 1, 0, 0, 0, 396, 392, 1,
		0, 0, 0, 397, 67, 1, 0, 0, 0, 398, 399, 5, 10, 0, 0, 399, 400, 3, 36, 18,
		0, 400, 401, 3, 60, 30, 0, 401, 437, 1, 0, 0, 0, 402, 403, 5, 10, 0, 0,
		403, 404, 3, 36, 18, 0, 404, 405, 3, 60, 30, 0, 405, 406, 5, 11, 0, 0,
		406, 407, 3, 68, 34, 0, 407, 437, 1, 0, 0, 0, 408, 409, 5, 10, 0, 0, 409,
		410, 3, 36, 18, 0, 410, 411, 3, 60, 30, 0, 411, 412, 5, 11, 0, 0, 412,
		413, 3, 60, 30, 0, 413, 437, 1, 0, 0, 0, 414, 415, 5, 10, 0, 0, 415, 416,
		3, 64, 32, 0, 416, 417, 5, 71, 0, 0, 417, 418, 3, 36, 18, 0, 418, 419,
		3, 60, 30, 0, 419, 437, 1, 0, 0, 0, 420, 421, 5, 10, 0, 0, 421, 422, 3,
		64, 32, 0, 422, 423, 5, 71, 0, 0, 423, 424, 3, 36, 18, 0, 424, 425, 3,
		60, 30, 0, 425, 426, 5, 11, 0, 0, 426, 427, 3, 68, 34, 0, 427, 437, 1,
		0, 0, 0, 428, 429, 5, 10, 0, 0, 429, 430, 3, 64, 32, 0, 430, 431, 5, 71,
		0, 0, 431, 432, 3, 36, 18, 0, 432, 433, 3, 60, 30, 0, 433, 434, 5, 11,
		0, 0, 434, 435, 3, 60, 30, 0, 435, 437, 1, 0, 0, 0, 436, 398, 1, 0, 0,
		0, 436, 402, 1, 0, 0, 0, 436, 408, 1, 0, 0, 0, 436, 414, 1, 0, 0, 0, 436,
		420, 1, 0, 0, 0, 436, 428, 1, 0, 0, 0, 437, 69, 1, 0, 0, 0, 438, 439, 5,
		9, 0, 0, 439, 455, 3, 60, 30, 0, 440, 441, 5, 9, 0, 0, 441, 442, 3, 36,
		18, 0, 442, 443, 3, 60, 30, 0, 443, 455, 1, 0, 0, 0, 444, 445, 5, 9, 0,
		0, 445, 446, 3, 64, 32, 0, 446, 448, 5, 71, 0, 0, 447, 449, 3, 36, 18,
		0, 448, 447, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 450, 1, 0, 0, 0, 450,
		451, 5, 71, 0, 0, 451, 452, 3, 64, 32, 0, 452, 453, 3, 60, 30, 0, 453,
		455, 1, 0, 0, 0, 454, 438, 1, 0, 0, 0, 454, 440, 1, 0, 0, 0, 454, 444,
		1, 0, 0, 0, 455, 71, 1, 0, 0, 0, 456, 457, 5, 8, 0, 0, 457, 458, 3, 64,
		32, 0, 458, 459, 5, 71, 0, 0, 459, 460, 3, 36, 18, 0, 460, 461, 5, 62,
		0, 0, 461, 462, 3, 74, 37, 0, 462, 463, 5, 63, 0, 0, 463, 483, 1, 0, 0,
		0, 464, 465, 5, 8, 0, 0, 465, 466, 5, 62, 0, 0, 466, 467, 3, 74, 37, 0,
		467, 468, 5, 63, 0, 0, 468, 483, 1, 0, 0, 0, 469, 470, 5, 8, 0, 0, 470,
		471, 3, 36, 18, 0, 471, 472, 5, 62, 0, 0, 472, 473, 3, 74, 37, 0, 473,
		474, 5, 63, 0, 0, 474, 483, 1, 0, 0, 0, 475, 476, 5, 8, 0, 0, 476, 477,
		3, 64, 32, 0, 477, 478, 5, 71, 0, 0, 478, 479, 5, 62, 0, 0, 479, 480, 3,
		74, 37, 0, 480, 481, 5, 63, 0, 0, 481, 483, 1, 0, 0, 0, 482, 456, 1, 0,
		0, 0, 482, 464, 1, 0, 0, 0, 482, 469, 1, 0, 0, 0, 482, 475, 1, 0, 0, 0,
		483, 73, 1, 0, 0, 0, 484, 485, 3, 76, 38, 0, 485, 486, 3, 74, 37, 0, 486,
		488, 1, 0, 0, 0, 487, 484, 1, 0, 0, 0, 488, 491, 1, 0, 0, 0, 489, 487,
		1, 0, 0, 0, 489, 490, 1, 0, 0, 0, 490, 75, 1, 0, 0, 0, 491, 489, 1, 0,
		0, 0, 492, 493, 3, 78, 39, 0, 493, 494, 5, 7, 0, 0, 494, 495, 3, 58, 29,
		0, 495, 77, 1, 0, 0, 0, 496, 497, 5, 6, 0, 0, 497, 500, 3, 38, 19, 0, 498,
		500, 5, 5, 0, 0, 499, 496, 1, 0, 0, 0, 499, 498, 1, 0, 0, 0, 500, 79, 1,
		0, 0, 0, 38, 88, 90, 107, 116, 129, 148, 157, 171, 175, 182, 193, 207,
		218, 226, 235, 246, 248, 256, 264, 272, 274, 283, 290, 300, 328, 338, 345,
		351, 375, 379, 386, 396, 436, 448, 454, 482, 489, 499,
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
	MinigoParserMULTILINE_COMMENT        = 2
	MinigoParserLEFTPARENTHESIS          = 3
	MinigoParserRIGHTPARENTHESIS         = 4
	MinigoParserDEFAULT                  = 5
	MinigoParserCASE                     = 6
	MinigoParserCOLON                    = 7
	MinigoParserSWITCH                   = 8
	MinigoParserFOR                      = 9
	MinigoParserIF                       = 10
	MinigoParserELSE                     = 11
	MinigoParserIDIV                     = 12
	MinigoParserIMOD                     = 13
	MinigoParserIANDXOR                  = 14
	MinigoParserILEFTSHIFT               = 15
	MinigoParserIRIGHTSHIFT              = 16
	MinigoParserIXOR                     = 17
	MinigoParserIMUL                     = 18
	MinigoParserIOR                      = 19
	MinigoParserISUB                     = 20
	MinigoParserIAND                     = 21
	MinigoParserIADD                     = 22
	MinigoParserPOSTINC                  = 23
	MinigoParserPOSTDEC                  = 24
	MinigoParserCONTINUE                 = 25
	MinigoParserWALRUS                   = 26
	MinigoParserBREAK                    = 27
	MinigoParserRETURN                   = 28
	MinigoParserPRINTLN                  = 29
	MinigoParserPRINT                    = 30
	MinigoParserCAP                      = 31
	MinigoParserLEN                      = 32
	MinigoParserAPPEND                   = 33
	MinigoParserDOT                      = 34
	MinigoParserWHITESPACE               = 35
	MinigoParserNEWLINE                  = 36
	MinigoParserINTERPRETEDSTRINGLITERAL = 37
	MinigoParserRAWSTRINGLITERAL         = 38
	MinigoParserESCAPEDSEQUENCES         = 39
	MinigoParserRUNELITERAL              = 40
	MinigoParserFLOATLITERAL             = 41
	MinigoParserNOT                      = 42
	MinigoParserOR                       = 43
	MinigoParserAND                      = 44
	MinigoParserGREATERTHANEQUAL         = 45
	MinigoParserLESSTHANEQUAL            = 46
	MinigoParserGREATERTHAN              = 47
	MinigoParserLESSTHAN                 = 48
	MinigoParserNEGATION                 = 49
	MinigoParserCOMPARISON               = 50
	MinigoParserCARET                    = 51
	MinigoParserPIPE                     = 52
	MinigoParserMINUS                    = 53
	MinigoParserPLUS                     = 54
	MinigoParserAMPERSANDCARET           = 55
	MinigoParserAMPERSAND                = 56
	MinigoParserRIGHTSHIFT               = 57
	MinigoParserLEFTSHIFT                = 58
	MinigoParserMOD                      = 59
	MinigoParserDIV                      = 60
	MinigoParserTIMES                    = 61
	MinigoParserLEFTCURLYBRACE           = 62
	MinigoParserRIGHTCURLYBRACE          = 63
	MinigoParserSTRUCT                   = 64
	MinigoParserINTLITERAL               = 65
	MinigoParserLEFTBRACKET              = 66
	MinigoParserRIGHTBRACKET             = 67
	MinigoParserCOMMA                    = 68
	MinigoParserFUNC                     = 69
	MinigoParserTYPE                     = 70
	MinigoParserSEMICOLON                = 71
	MinigoParserVAR                      = 72
	MinigoParserEQUALS                   = 73
	MinigoParserPACKAGE                  = 74
	MinigoParserIDENTIFIER               = 75
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

	for (int64((_la-69)) & ^0x3f) == 0 && ((int64(1)<<(_la-69))&11) != 0 {
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

func (s *SingleVarDeclContext) CopyAll(ctx *SingleVarDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SingleVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypedVarDeclContext struct {
	SingleVarDeclContext
}

func NewTypedVarDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedVarDeclContext {
	var p = new(TypedVarDeclContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *TypedVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedVarDeclContext) IdentifierList() IIdentifierListContext {
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

func (s *TypedVarDeclContext) DeclType() IDeclTypeContext {
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

func (s *TypedVarDeclContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(MinigoParserEQUALS, 0)
}

func (s *TypedVarDeclContext) ExpressionList() IExpressionListContext {
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

func (s *TypedVarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterTypedVarDecl(s)
	}
}

func (s *TypedVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitTypedVarDecl(s)
	}
}

func (s *TypedVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitTypedVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleVarDeclsNoExpsDeclContext struct {
	SingleVarDeclContext
}

func NewSingleVarDeclsNoExpsDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleVarDeclsNoExpsDeclContext {
	var p = new(SingleVarDeclsNoExpsDeclContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *SingleVarDeclsNoExpsDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclsNoExpsDeclContext) SingleVarDeclNoExps() ISingleVarDeclNoExpsContext {
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

func (s *SingleVarDeclsNoExpsDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSingleVarDeclsNoExpsDecl(s)
	}
}

func (s *SingleVarDeclsNoExpsDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSingleVarDeclsNoExpsDecl(s)
	}
}

func (s *SingleVarDeclsNoExpsDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSingleVarDeclsNoExpsDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type UntypedVarDeclContext struct {
	SingleVarDeclContext
}

func NewUntypedVarDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UntypedVarDeclContext {
	var p = new(UntypedVarDeclContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *UntypedVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntypedVarDeclContext) IdentifierList() IIdentifierListContext {
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

func (s *UntypedVarDeclContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(MinigoParserEQUALS, 0)
}

func (s *UntypedVarDeclContext) ExpressionList() IExpressionListContext {
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

func (s *UntypedVarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterUntypedVarDecl(s)
	}
}

func (s *UntypedVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitUntypedVarDecl(s)
	}
}

func (s *UntypedVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitUntypedVarDecl(s)

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
		localctx = NewTypedVarDeclContext(p, localctx)
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
		localctx = NewUntypedVarDeclContext(p, localctx)
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
		localctx = NewSingleVarDeclsNoExpsDeclContext(p, localctx)
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

func (s *TypeDeclContext) CopyAll(ctx *TypeDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeDeclarationContext struct {
	TypeDeclContext
}

func NewTypeDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(MinigoParserTYPE, 0)
}

func (s *TypeDeclarationContext) SingleTypeDecl() ISingleTypeDeclContext {
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

func (s *TypeDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiTypeDeclarationContext struct {
	TypeDeclContext
}

func NewMultiTypeDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiTypeDeclarationContext {
	var p = new(MultiTypeDeclarationContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *MultiTypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiTypeDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(MinigoParserTYPE, 0)
}

func (s *MultiTypeDeclarationContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *MultiTypeDeclarationContext) InnerTypeDecls() IInnerTypeDeclsContext {
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

func (s *MultiTypeDeclarationContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *MultiTypeDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *MultiTypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterMultiTypeDeclaration(s)
	}
}

func (s *MultiTypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitMultiTypeDeclaration(s)
	}
}

func (s *MultiTypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitMultiTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmptyTypeDeclarationContext struct {
	TypeDeclContext
}

func NewEmptyTypeDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyTypeDeclarationContext {
	var p = new(EmptyTypeDeclarationContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *EmptyTypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyTypeDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(MinigoParserTYPE, 0)
}

func (s *EmptyTypeDeclarationContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *EmptyTypeDeclarationContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *EmptyTypeDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *EmptyTypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterEmptyTypeDeclaration(s)
	}
}

func (s *EmptyTypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitEmptyTypeDeclaration(s)
	}
}

func (s *EmptyTypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitEmptyTypeDeclaration(s)

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
		localctx = NewTypeDeclarationContext(p, localctx)
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
		localctx = NewMultiTypeDeclarationContext(p, localctx)
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
		localctx = NewEmptyTypeDeclarationContext(p, localctx)
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
	DeclType() IDeclTypeContext

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

func (s *FuncFrontDeclContext) DeclType() IDeclTypeContext {
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
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MinigoParserLEFTPARENTHESIS || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&2053) != 0) {
		{
			p.SetState(174)
			p.DeclType()
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
		p.SetState(177)
		p.SingleVarDeclNoExps()
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinigoParserCOMMA {
		{
			p.SetState(178)
			p.Match(MinigoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.SingleVarDeclNoExps()
		}

		p.SetState(184)
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
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Match(MinigoParserLEFTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.DeclType()
		}
		{
			p.SetState(187)
			p.Match(MinigoParserRIGHTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Match(MinigoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(190)
			p.SliceDeclType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(191)
			p.ArrayDeclType()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(192)
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
		p.SetState(195)
		p.Match(MinigoParserLEFTBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.Match(MinigoParserRIGHTBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
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
		p.SetState(199)
		p.Match(MinigoParserLEFTBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.Match(MinigoParserINTLITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Match(MinigoParserRIGHTBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
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
	StructMemDecls() IStructMemDeclsContext

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

func (s *StructDeclTypeContext) StructMemDecls() IStructMemDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructMemDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
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
		p.SetState(204)
		p.Match(MinigoParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(MinigoParserLEFTCURLYBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MinigoParserIDENTIFIER {
		{
			p.SetState(206)
			p.StructMemDecls()
		}

	}
	{
		p.SetState(209)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.SingleVarDeclNoExps()
	}
	{
		p.SetState(212)
		p.Match(MinigoParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinigoParserIDENTIFIER {
		{
			p.SetState(213)
			p.SingleVarDeclNoExps()
		}
		{
			p.SetState(214)
			p.Match(MinigoParserSEMICOLON)
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
		p.SetState(221)
		p.Match(MinigoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinigoParserCOMMA {
		{
			p.SetState(222)
			p.Match(MinigoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Match(MinigoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(228)
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

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ComparisonContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonContext {
	var p = new(ComparisonContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ComparisonContext) GetLeft() IExpressionContext { return s.left }

func (s *ComparisonContext) GetRight() IExpressionContext { return s.right }

func (s *ComparisonContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *ComparisonContext) SetRight(v IExpressionContext) { s.right = v }

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) AllExpression() []IExpressionContext {
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

func (s *ComparisonContext) Expression(i int) IExpressionContext {
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

func (s *ComparisonContext) LESSTHAN() antlr.TerminalNode {
	return s.GetToken(MinigoParserLESSTHAN, 0)
}

func (s *ComparisonContext) GREATERTHAN() antlr.TerminalNode {
	return s.GetToken(MinigoParserGREATERTHAN, 0)
}

func (s *ComparisonContext) LESSTHANEQUAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserLESSTHANEQUAL, 0)
}

func (s *ComparisonContext) GREATERTHANEQUAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserGREATERTHANEQUAL, 0)
}

func (s *ComparisonContext) COMPARISON() antlr.TerminalNode {
	return s.GetToken(MinigoParserCOMPARISON, 0)
}

func (s *ComparisonContext) NEGATION() antlr.TerminalNode {
	return s.GetToken(MinigoParserNEGATION, 0)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionPrimaryExpressionContext struct {
	ExpressionContext
}

func NewExpressionPrimaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionPrimaryExpressionContext {
	var p = new(ExpressionPrimaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionPrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionPrimaryExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
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

func (s *ExpressionPrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterExpressionPrimaryExpression(s)
	}
}

func (s *ExpressionPrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitExpressionPrimaryExpression(s)
	}
}

func (s *ExpressionPrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitExpressionPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExpressionContext struct {
	ExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(MinigoParserNOT, 0)
}

func (s *NotExpressionContext) Expression() IExpressionContext {
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

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

func (s *NotExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitNotExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CaretExpressionContext struct {
	ExpressionContext
}

func NewCaretExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaretExpressionContext {
	var p = new(CaretExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CaretExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaretExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(MinigoParserCARET, 0)
}

func (s *CaretExpressionContext) Expression() IExpressionContext {
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

func (s *CaretExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterCaretExpression(s)
	}
}

func (s *CaretExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitCaretExpression(s)
	}
}

func (s *CaretExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitCaretExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanOperationContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewBooleanOperationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanOperationContext {
	var p = new(BooleanOperationContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BooleanOperationContext) GetLeft() IExpressionContext { return s.left }

func (s *BooleanOperationContext) GetRight() IExpressionContext { return s.right }

func (s *BooleanOperationContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *BooleanOperationContext) SetRight(v IExpressionContext) { s.right = v }

func (s *BooleanOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanOperationContext) AllExpression() []IExpressionContext {
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

func (s *BooleanOperationContext) Expression(i int) IExpressionContext {
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

func (s *BooleanOperationContext) AND() antlr.TerminalNode {
	return s.GetToken(MinigoParserAND, 0)
}

func (s *BooleanOperationContext) OR() antlr.TerminalNode {
	return s.GetToken(MinigoParserOR, 0)
}

func (s *BooleanOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterBooleanOperation(s)
	}
}

func (s *BooleanOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitBooleanOperation(s)
	}
}

func (s *BooleanOperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitBooleanOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

type OperationContext struct {
	ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewOperationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperationContext {
	var p = new(OperationContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OperationContext) GetLeft() IExpressionContext { return s.left }

func (s *OperationContext) GetRight() IExpressionContext { return s.right }

func (s *OperationContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *OperationContext) SetRight(v IExpressionContext) { s.right = v }

func (s *OperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationContext) AllExpression() []IExpressionContext {
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

func (s *OperationContext) Expression(i int) IExpressionContext {
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

func (s *OperationContext) TIMES() antlr.TerminalNode {
	return s.GetToken(MinigoParserTIMES, 0)
}

func (s *OperationContext) DIV() antlr.TerminalNode {
	return s.GetToken(MinigoParserDIV, 0)
}

func (s *OperationContext) MOD() antlr.TerminalNode {
	return s.GetToken(MinigoParserMOD, 0)
}

func (s *OperationContext) LEFTSHIFT() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTSHIFT, 0)
}

func (s *OperationContext) RIGHTSHIFT() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTSHIFT, 0)
}

func (s *OperationContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(MinigoParserAMPERSAND, 0)
}

func (s *OperationContext) AMPERSANDCARET() antlr.TerminalNode {
	return s.GetToken(MinigoParserAMPERSANDCARET, 0)
}

func (s *OperationContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MinigoParserPLUS, 0)
}

func (s *OperationContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MinigoParserMINUS, 0)
}

func (s *OperationContext) PIPE() antlr.TerminalNode {
	return s.GetToken(MinigoParserPIPE, 0)
}

func (s *OperationContext) CARET() antlr.TerminalNode {
	return s.GetToken(MinigoParserCARET, 0)
}

func (s *OperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterOperation(s)
	}
}

func (s *OperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitOperation(s)
	}
}

func (s *OperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitOperation(s)

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
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinigoParserLEFTPARENTHESIS, MinigoParserCAP, MinigoParserLEN, MinigoParserAPPEND, MinigoParserINTERPRETEDSTRINGLITERAL, MinigoParserRAWSTRINGLITERAL, MinigoParserRUNELITERAL, MinigoParserFLOATLITERAL, MinigoParserINTLITERAL, MinigoParserIDENTIFIER:
		localctx = NewExpressionPrimaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(230)
			p.primaryExpression(0)
		}

	case MinigoParserNOT:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(231)
			p.Match(MinigoParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.expression(2)
		}

	case MinigoParserCARET:
		localctx = NewCaretExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(233)
			p.Match(MinigoParserCARET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.expression(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(246)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBooleanOperationContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BooleanOperationContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, MinigoParserRULE_expression)
				p.SetState(237)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(238)
					_la = p.GetTokenStream().LA(1)

					if !(_la == MinigoParserOR || _la == MinigoParserAND) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(239)

					var _x = p.expression(7)

					localctx.(*BooleanOperationContext).right = _x
				}

			case 2:
				localctx = NewComparisonContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ComparisonContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, MinigoParserRULE_expression)
				p.SetState(240)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(241)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2216615441596416) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(242)

					var _x = p.expression(6)

					localctx.(*ComparisonContext).right = _x
				}

			case 3:
				localctx = NewOperationContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*OperationContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, MinigoParserRULE_expression)
				p.SetState(243)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(244)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4609434218613702656) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(245)

					var _x = p.expression(5)

					localctx.(*OperationContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
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
		p.SetState(251)
		p.expression(0)
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MinigoParserCOMMA {
		{
			p.SetState(252)
			p.Match(MinigoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.expression(0)
		}

		p.SetState(258)
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

func (s *PrimaryExpressionContext) CopyAll(ctx *PrimaryExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SubIndexContext struct {
	PrimaryExpressionContext
}

func NewSubIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubIndexContext {
	var p = new(SubIndexContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *SubIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubIndexContext) PrimaryExpression() IPrimaryExpressionContext {
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

func (s *SubIndexContext) Index() IIndexContext {
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

func (s *SubIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSubIndex(s)
	}
}

func (s *SubIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSubIndex(s)
	}
}

func (s *SubIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSubIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	PrimaryExpressionContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) PrimaryExpression() IPrimaryExpressionContext {
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

func (s *FunctionCallContext) Arguments() IArgumentsContext {
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

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type CapCallContext struct {
	PrimaryExpressionContext
}

func NewCapCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CapCallContext {
	var p = new(CapCallContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *CapCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapCallContext) CapExpression() ICapExpressionContext {
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

func (s *CapCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterCapCall(s)
	}
}

func (s *CapCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitCapCall(s)
	}
}

func (s *CapCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitCapCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type OperandExpressionContext struct {
	PrimaryExpressionContext
}

func NewOperandExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperandExpressionContext {
	var p = new(OperandExpressionContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *OperandExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandExpressionContext) Operand() IOperandContext {
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

func (s *OperandExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterOperandExpression(s)
	}
}

func (s *OperandExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitOperandExpression(s)
	}
}

func (s *OperandExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitOperandExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppendCallContext struct {
	PrimaryExpressionContext
}

func NewAppendCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendCallContext {
	var p = new(AppendCallContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *AppendCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendCallContext) AppendExpression() IAppendExpressionContext {
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

func (s *AppendCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterAppendCall(s)
	}
}

func (s *AppendCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitAppendCall(s)
	}
}

func (s *AppendCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitAppendCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type LenCallContext struct {
	PrimaryExpressionContext
}

func NewLenCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LenCallContext {
	var p = new(LenCallContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *LenCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LenCallContext) LengthExpression() ILengthExpressionContext {
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

func (s *LenCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterLenCall(s)
	}
}

func (s *LenCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitLenCall(s)
	}
}

func (s *LenCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitLenCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type MemberAccessorContext struct {
	PrimaryExpressionContext
}

func NewMemberAccessorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberAccessorContext {
	var p = new(MemberAccessorContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *MemberAccessorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberAccessorContext) PrimaryExpression() IPrimaryExpressionContext {
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

func (s *MemberAccessorContext) Selector() ISelectorContext {
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

func (s *MemberAccessorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterMemberAccessor(s)
	}
}

func (s *MemberAccessorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitMemberAccessor(s)
	}
}

func (s *MemberAccessorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitMemberAccessor(s)

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
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinigoParserLEFTPARENTHESIS, MinigoParserINTERPRETEDSTRINGLITERAL, MinigoParserRAWSTRINGLITERAL, MinigoParserRUNELITERAL, MinigoParserFLOATLITERAL, MinigoParserINTLITERAL, MinigoParserIDENTIFIER:
		localctx = NewOperandExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(260)
			p.Operand()
		}

	case MinigoParserAPPEND:
		localctx = NewAppendCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(261)
			p.AppendExpression()
		}

	case MinigoParserLEN:
		localctx = NewLenCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(262)
			p.LengthExpression()
		}

	case MinigoParserCAP:
		localctx = NewCapCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(263)
			p.CapExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(272)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMemberAccessorContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MinigoParserRULE_primaryExpression)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(267)
					p.Selector()
				}

			case 2:
				localctx = NewSubIndexContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MinigoParserRULE_primaryExpression)
				p.SetState(268)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(269)
					p.Index()
				}

			case 3:
				localctx = NewFunctionCallContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MinigoParserRULE_primaryExpression)
				p.SetState(270)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(271)
					p.Arguments()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
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

func (s *OperandContext) CopyAll(ctx *OperandContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LiteralOperandContext struct {
	OperandContext
}

func NewLiteralOperandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralOperandContext {
	var p = new(LiteralOperandContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *LiteralOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralOperandContext) Literal() ILiteralContext {
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

func (s *LiteralOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterLiteralOperand(s)
	}
}

func (s *LiteralOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitLiteralOperand(s)
	}
}

func (s *LiteralOperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitLiteralOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierOperandContext struct {
	OperandContext
}

func NewIdentifierOperandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierOperandContext {
	var p = new(IdentifierOperandContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *IdentifierOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierOperandContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MinigoParserIDENTIFIER, 0)
}

func (s *IdentifierOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterIdentifierOperand(s)
	}
}

func (s *IdentifierOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitIdentifierOperand(s)
	}
}

func (s *IdentifierOperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitIdentifierOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionOperandContext struct {
	OperandContext
}

func NewExpressionOperandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionOperandContext {
	var p = new(ExpressionOperandContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *ExpressionOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionOperandContext) LEFTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTPARENTHESIS, 0)
}

func (s *ExpressionOperandContext) Expression() IExpressionContext {
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

func (s *ExpressionOperandContext) RIGHTPARENTHESIS() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTPARENTHESIS, 0)
}

func (s *ExpressionOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterExpressionOperand(s)
	}
}

func (s *ExpressionOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitExpressionOperand(s)
	}
}

func (s *ExpressionOperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitExpressionOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MinigoParserRULE_operand)
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinigoParserINTERPRETEDSTRINGLITERAL, MinigoParserRAWSTRINGLITERAL, MinigoParserRUNELITERAL, MinigoParserFLOATLITERAL, MinigoParserINTLITERAL:
		localctx = NewLiteralOperandContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(277)
			p.Literal()
		}

	case MinigoParserIDENTIFIER:
		localctx = NewIdentifierOperandContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)
			p.Match(MinigoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserLEFTPARENTHESIS:
		localctx = NewExpressionOperandContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(279)
			p.Match(MinigoParserLEFTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.expression(0)
		}
		{
			p.SetState(281)
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

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InterpretedStringLiteralContext struct {
	LiteralContext
}

func NewInterpretedStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InterpretedStringLiteralContext {
	var p = new(InterpretedStringLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *InterpretedStringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterpretedStringLiteralContext) INTERPRETEDSTRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserINTERPRETEDSTRINGLITERAL, 0)
}

func (s *InterpretedStringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterInterpretedStringLiteral(s)
	}
}

func (s *InterpretedStringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitInterpretedStringLiteral(s)
	}
}

func (s *InterpretedStringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitInterpretedStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntLiteralContext struct {
	LiteralContext
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext {
	var p = new(IntLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) INTLITERAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserINTLITERAL, 0)
}

func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatLiteralContext struct {
	LiteralContext
}

func NewFloatLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) FLOATLITERAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserFLOATLITERAL, 0)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type RuneLiteralContext struct {
	LiteralContext
}

func NewRuneLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RuneLiteralContext {
	var p = new(RuneLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *RuneLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuneLiteralContext) RUNELITERAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserRUNELITERAL, 0)
}

func (s *RuneLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterRuneLiteral(s)
	}
}

func (s *RuneLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitRuneLiteral(s)
	}
}

func (s *RuneLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitRuneLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type RawStringLiteralContext struct {
	LiteralContext
}

func NewRawStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RawStringLiteralContext {
	var p = new(RawStringLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *RawStringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawStringLiteralContext) RAWSTRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(MinigoParserRAWSTRINGLITERAL, 0)
}

func (s *RawStringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterRawStringLiteral(s)
	}
}

func (s *RawStringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitRawStringLiteral(s)
	}
}

func (s *RawStringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitRawStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MinigoParserRULE_literal)
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinigoParserINTLITERAL:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.Match(MinigoParserINTLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserFLOATLITERAL:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.Match(MinigoParserFLOATLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserRUNELITERAL:
		localctx = NewRuneLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(287)
			p.Match(MinigoParserRUNELITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserRAWSTRINGLITERAL:
		localctx = NewRawStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(288)
			p.Match(MinigoParserRAWSTRINGLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserINTERPRETEDSTRINGLITERAL:
		localctx = NewInterpretedStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(289)
			p.Match(MinigoParserINTERPRETEDSTRINGLITERAL)
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
		p.SetState(292)
		p.Match(MinigoParserLEFTBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.expression(0)
	}
	{
		p.SetState(294)
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
		p.SetState(296)
		p.Match(MinigoParserLEFTPARENTHESIS)
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

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2259923744325640) != 0) || _la == MinigoParserINTLITERAL || _la == MinigoParserIDENTIFIER {
		{
			p.SetState(297)
			p.ExpressionList()
		}

		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(303)
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
		p.SetState(305)
		p.Match(MinigoParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
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
		p.SetState(308)
		p.Match(MinigoParserAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)
		p.Match(MinigoParserLEFTPARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
		p.expression(0)
	}
	{
		p.SetState(311)
		p.Match(MinigoParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
		p.expression(0)
	}
	{
		p.SetState(313)
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
		p.SetState(315)
		p.Match(MinigoParserLEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(316)
		p.Match(MinigoParserLEFTPARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(317)
		p.expression(0)
	}
	{
		p.SetState(318)
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
		p.SetState(320)
		p.Match(MinigoParserCAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.Match(MinigoParserLEFTPARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.expression(0)
	}
	{
		p.SetState(323)
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
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4613945944218535688) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&1185) != 0) {
		{
			p.SetState(325)
			p.Statement()
		}

		p.SetState(330)
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
		p.SetState(331)
		p.Match(MinigoParserLEFTCURLYBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(332)
		p.StatementList()
	}
	{
		p.SetState(333)
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

	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinigoParserPRINT:
		localctx = NewPrintStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(335)
			p.Match(MinigoParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Match(MinigoParserLEFTPARENTHESIS)
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

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2259923744325640) != 0) || _la == MinigoParserINTLITERAL || _la == MinigoParserIDENTIFIER {
			{
				p.SetState(337)
				p.ExpressionList()
			}

		}
		{
			p.SetState(340)
			p.Match(MinigoParserRIGHTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
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
			p.SetState(342)
			p.Match(MinigoParserPRINTLN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Match(MinigoParserLEFTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2259923744325640) != 0) || _la == MinigoParserINTLITERAL || _la == MinigoParserIDENTIFIER {
			{
				p.SetState(344)
				p.ExpressionList()
			}

		}
		{
			p.SetState(347)
			p.Match(MinigoParserRIGHTPARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)
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
			p.SetState(349)
			p.Match(MinigoParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2259923744325640) != 0) || _la == MinigoParserINTLITERAL || _la == MinigoParserIDENTIFIER {
			{
				p.SetState(350)
				p.expression(0)
			}

		}
		{
			p.SetState(353)
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
			p.SetState(354)
			p.Match(MinigoParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
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
			p.SetState(356)
			p.Match(MinigoParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MinigoParserLEFTPARENTHESIS, MinigoParserCAP, MinigoParserLEN, MinigoParserAPPEND, MinigoParserINTERPRETEDSTRINGLITERAL, MinigoParserRAWSTRINGLITERAL, MinigoParserRUNELITERAL, MinigoParserFLOATLITERAL, MinigoParserNOT, MinigoParserCARET, MinigoParserINTLITERAL, MinigoParserIDENTIFIER:
		localctx = NewSimpleStatementStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(358)
			p.SimpleStatement()
		}
		{
			p.SetState(359)
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
			p.SetState(361)
			p.Block()
		}
		{
			p.SetState(362)
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
			p.SetState(364)
			p.Switch_()
		}
		{
			p.SetState(365)
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
			p.SetState(367)
			p.IfStatement()
		}
		{
			p.SetState(368)
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
			p.SetState(370)
			p.Loop()
		}
		{
			p.SetState(371)
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
			p.SetState(373)
			p.TypeDecl()
		}

	case MinigoParserVAR:
		localctx = NewVariableDeclStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(374)
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

func (s *SimpleStatementContext) CopyAll(ctx *SimpleStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SimpleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionSimpleStatementContext struct {
	SimpleStatementContext
}

func NewExpressionSimpleStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionSimpleStatementContext {
	var p = new(ExpressionSimpleStatementContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *ExpressionSimpleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSimpleStatementContext) Expression() IExpressionContext {
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

func (s *ExpressionSimpleStatementContext) POSTINC() antlr.TerminalNode {
	return s.GetToken(MinigoParserPOSTINC, 0)
}

func (s *ExpressionSimpleStatementContext) POSTDEC() antlr.TerminalNode {
	return s.GetToken(MinigoParserPOSTDEC, 0)
}

func (s *ExpressionSimpleStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterExpressionSimpleStatement(s)
	}
}

func (s *ExpressionSimpleStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitExpressionSimpleStatement(s)
	}
}

func (s *ExpressionSimpleStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitExpressionSimpleStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentSimpleStatementContext struct {
	SimpleStatementContext
}

func NewAssignmentSimpleStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentSimpleStatementContext {
	var p = new(AssignmentSimpleStatementContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *AssignmentSimpleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentSimpleStatementContext) AssignmentStatement() IAssignmentStatementContext {
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

func (s *AssignmentSimpleStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterAssignmentSimpleStatement(s)
	}
}

func (s *AssignmentSimpleStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitAssignmentSimpleStatement(s)
	}
}

func (s *AssignmentSimpleStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitAssignmentSimpleStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type WalrusDeclarationContext struct {
	SimpleStatementContext
	left  IExpressionListContext
	right IExpressionListContext
}

func NewWalrusDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WalrusDeclarationContext {
	var p = new(WalrusDeclarationContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *WalrusDeclarationContext) GetLeft() IExpressionListContext { return s.left }

func (s *WalrusDeclarationContext) GetRight() IExpressionListContext { return s.right }

func (s *WalrusDeclarationContext) SetLeft(v IExpressionListContext) { s.left = v }

func (s *WalrusDeclarationContext) SetRight(v IExpressionListContext) { s.right = v }

func (s *WalrusDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WalrusDeclarationContext) WALRUS() antlr.TerminalNode {
	return s.GetToken(MinigoParserWALRUS, 0)
}

func (s *WalrusDeclarationContext) AllExpressionList() []IExpressionListContext {
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

func (s *WalrusDeclarationContext) ExpressionList(i int) IExpressionListContext {
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

func (s *WalrusDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterWalrusDeclaration(s)
	}
}

func (s *WalrusDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitWalrusDeclaration(s)
	}
}

func (s *WalrusDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitWalrusDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) SimpleStatement() (localctx ISimpleStatementContext) {
	localctx = NewSimpleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MinigoParserRULE_simpleStatement)
	var _la int

	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpressionSimpleStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(377)
			p.expression(0)
		}
		p.SetState(379)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MinigoParserPOSTINC || _la == MinigoParserPOSTDEC {
			{
				p.SetState(378)
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
		localctx = NewAssignmentSimpleStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(381)
			p.AssignmentStatement()
		}

	case 3:
		localctx = NewWalrusDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(382)

			var _x = p.ExpressionList()

			localctx.(*WalrusDeclarationContext).left = _x
		}
		{
			p.SetState(383)
			p.Match(MinigoParserWALRUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(384)

			var _x = p.ExpressionList()

			localctx.(*WalrusDeclarationContext).right = _x
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

func (s *AssignmentStatementContext) CopyAll(ctx *AssignmentStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NormalAssignmentContext struct {
	AssignmentStatementContext
	left  IExpressionListContext
	right IExpressionListContext
}

func NewNormalAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NormalAssignmentContext {
	var p = new(NormalAssignmentContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *NormalAssignmentContext) GetLeft() IExpressionListContext { return s.left }

func (s *NormalAssignmentContext) GetRight() IExpressionListContext { return s.right }

func (s *NormalAssignmentContext) SetLeft(v IExpressionListContext) { s.left = v }

func (s *NormalAssignmentContext) SetRight(v IExpressionListContext) { s.right = v }

func (s *NormalAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NormalAssignmentContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(MinigoParserEQUALS, 0)
}

func (s *NormalAssignmentContext) AllExpressionList() []IExpressionListContext {
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

func (s *NormalAssignmentContext) ExpressionList(i int) IExpressionListContext {
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

func (s *NormalAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterNormalAssignment(s)
	}
}

func (s *NormalAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitNormalAssignment(s)
	}
}

func (s *NormalAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitNormalAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type InPlaceAssignmentContext struct {
	AssignmentStatementContext
	left  IExpressionContext
	right IExpressionContext
}

func NewInPlaceAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InPlaceAssignmentContext {
	var p = new(InPlaceAssignmentContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *InPlaceAssignmentContext) GetLeft() IExpressionContext { return s.left }

func (s *InPlaceAssignmentContext) GetRight() IExpressionContext { return s.right }

func (s *InPlaceAssignmentContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *InPlaceAssignmentContext) SetRight(v IExpressionContext) { s.right = v }

func (s *InPlaceAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InPlaceAssignmentContext) AllExpression() []IExpressionContext {
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

func (s *InPlaceAssignmentContext) Expression(i int) IExpressionContext {
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

func (s *InPlaceAssignmentContext) IADD() antlr.TerminalNode {
	return s.GetToken(MinigoParserIADD, 0)
}

func (s *InPlaceAssignmentContext) IAND() antlr.TerminalNode {
	return s.GetToken(MinigoParserIAND, 0)
}

func (s *InPlaceAssignmentContext) ISUB() antlr.TerminalNode {
	return s.GetToken(MinigoParserISUB, 0)
}

func (s *InPlaceAssignmentContext) IOR() antlr.TerminalNode {
	return s.GetToken(MinigoParserIOR, 0)
}

func (s *InPlaceAssignmentContext) IMUL() antlr.TerminalNode {
	return s.GetToken(MinigoParserIMUL, 0)
}

func (s *InPlaceAssignmentContext) IXOR() antlr.TerminalNode {
	return s.GetToken(MinigoParserIXOR, 0)
}

func (s *InPlaceAssignmentContext) ILEFTSHIFT() antlr.TerminalNode {
	return s.GetToken(MinigoParserILEFTSHIFT, 0)
}

func (s *InPlaceAssignmentContext) IRIGHTSHIFT() antlr.TerminalNode {
	return s.GetToken(MinigoParserIRIGHTSHIFT, 0)
}

func (s *InPlaceAssignmentContext) IANDXOR() antlr.TerminalNode {
	return s.GetToken(MinigoParserIANDXOR, 0)
}

func (s *InPlaceAssignmentContext) IMOD() antlr.TerminalNode {
	return s.GetToken(MinigoParserIMOD, 0)
}

func (s *InPlaceAssignmentContext) IDIV() antlr.TerminalNode {
	return s.GetToken(MinigoParserIDIV, 0)
}

func (s *InPlaceAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterInPlaceAssignment(s)
	}
}

func (s *InPlaceAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitInPlaceAssignment(s)
	}
}

func (s *InPlaceAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitInPlaceAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MinigoParserRULE_assignmentStatement)
	var _la int

	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNormalAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(388)

			var _x = p.ExpressionList()

			localctx.(*NormalAssignmentContext).left = _x
		}
		{
			p.SetState(389)
			p.Match(MinigoParserEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)

			var _x = p.ExpressionList()

			localctx.(*NormalAssignmentContext).right = _x
		}

	case 2:
		localctx = NewInPlaceAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(392)

			var _x = p.expression(0)

			localctx.(*InPlaceAssignmentContext).left = _x
		}
		{
			p.SetState(393)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8384512) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(394)

			var _x = p.expression(0)

			localctx.(*InPlaceAssignmentContext).right = _x
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

func (s *IfStatementContext) CopyAll(ctx *IfStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfSingleExpressionContext struct {
	IfStatementContext
}

func NewIfSingleExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfSingleExpressionContext {
	var p = new(IfSingleExpressionContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfSingleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfSingleExpressionContext) IF() antlr.TerminalNode {
	return s.GetToken(MinigoParserIF, 0)
}

func (s *IfSingleExpressionContext) Expression() IExpressionContext {
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

func (s *IfSingleExpressionContext) Block() IBlockContext {
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

func (s *IfSingleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterIfSingleExpression(s)
	}
}

func (s *IfSingleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitIfSingleExpression(s)
	}
}

func (s *IfSingleExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitIfSingleExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfSimpleElseIfContext struct {
	IfStatementContext
}

func NewIfSimpleElseIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfSimpleElseIfContext {
	var p = new(IfSimpleElseIfContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfSimpleElseIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfSimpleElseIfContext) IF() antlr.TerminalNode {
	return s.GetToken(MinigoParserIF, 0)
}

func (s *IfSimpleElseIfContext) SimpleStatement() ISimpleStatementContext {
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

func (s *IfSimpleElseIfContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *IfSimpleElseIfContext) Expression() IExpressionContext {
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

func (s *IfSimpleElseIfContext) Block() IBlockContext {
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

func (s *IfSimpleElseIfContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MinigoParserELSE, 0)
}

func (s *IfSimpleElseIfContext) IfStatement() IIfStatementContext {
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

func (s *IfSimpleElseIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterIfSimpleElseIf(s)
	}
}

func (s *IfSimpleElseIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitIfSimpleElseIf(s)
	}
}

func (s *IfSimpleElseIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitIfSimpleElseIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfSimpleNoElseContext struct {
	IfStatementContext
}

func NewIfSimpleNoElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfSimpleNoElseContext {
	var p = new(IfSimpleNoElseContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfSimpleNoElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfSimpleNoElseContext) IF() antlr.TerminalNode {
	return s.GetToken(MinigoParserIF, 0)
}

func (s *IfSimpleNoElseContext) SimpleStatement() ISimpleStatementContext {
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

func (s *IfSimpleNoElseContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *IfSimpleNoElseContext) Expression() IExpressionContext {
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

func (s *IfSimpleNoElseContext) Block() IBlockContext {
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

func (s *IfSimpleNoElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterIfSimpleNoElse(s)
	}
}

func (s *IfSimpleNoElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitIfSimpleNoElse(s)
	}
}

func (s *IfSimpleNoElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitIfSimpleNoElse(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseBlockContext struct {
	IfStatementContext
}

func NewIfElseBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseBlockContext {
	var p = new(IfElseBlockContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseBlockContext) IF() antlr.TerminalNode {
	return s.GetToken(MinigoParserIF, 0)
}

func (s *IfElseBlockContext) Expression() IExpressionContext {
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

func (s *IfElseBlockContext) AllBlock() []IBlockContext {
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

func (s *IfElseBlockContext) Block(i int) IBlockContext {
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

func (s *IfElseBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MinigoParserELSE, 0)
}

func (s *IfElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterIfElseBlock(s)
	}
}

func (s *IfElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitIfElseBlock(s)
	}
}

func (s *IfElseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitIfElseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfSimpleElseBlockContext struct {
	IfStatementContext
}

func NewIfSimpleElseBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfSimpleElseBlockContext {
	var p = new(IfSimpleElseBlockContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfSimpleElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfSimpleElseBlockContext) IF() antlr.TerminalNode {
	return s.GetToken(MinigoParserIF, 0)
}

func (s *IfSimpleElseBlockContext) SimpleStatement() ISimpleStatementContext {
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

func (s *IfSimpleElseBlockContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *IfSimpleElseBlockContext) Expression() IExpressionContext {
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

func (s *IfSimpleElseBlockContext) AllBlock() []IBlockContext {
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

func (s *IfSimpleElseBlockContext) Block(i int) IBlockContext {
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

func (s *IfSimpleElseBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MinigoParserELSE, 0)
}

func (s *IfSimpleElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterIfSimpleElseBlock(s)
	}
}

func (s *IfSimpleElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitIfSimpleElseBlock(s)
	}
}

func (s *IfSimpleElseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitIfSimpleElseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseIfContext struct {
	IfStatementContext
}

func NewIfElseIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseIfContext {
	var p = new(IfElseIfContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfElseIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseIfContext) IF() antlr.TerminalNode {
	return s.GetToken(MinigoParserIF, 0)
}

func (s *IfElseIfContext) Expression() IExpressionContext {
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

func (s *IfElseIfContext) Block() IBlockContext {
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

func (s *IfElseIfContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MinigoParserELSE, 0)
}

func (s *IfElseIfContext) IfStatement() IIfStatementContext {
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

func (s *IfElseIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterIfElseIf(s)
	}
}

func (s *IfElseIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitIfElseIf(s)
	}
}

func (s *IfElseIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitIfElseIf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, MinigoParserRULE_ifStatement)
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfSingleExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)
			p.Match(MinigoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.expression(0)
		}
		{
			p.SetState(400)
			p.Block()
		}

	case 2:
		localctx = NewIfElseIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(402)
			p.Match(MinigoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.expression(0)
		}
		{
			p.SetState(404)
			p.Block()
		}
		{
			p.SetState(405)
			p.Match(MinigoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.IfStatement()
		}

	case 3:
		localctx = NewIfElseBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(408)
			p.Match(MinigoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.expression(0)
		}
		{
			p.SetState(410)
			p.Block()
		}
		{
			p.SetState(411)
			p.Match(MinigoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)
			p.Block()
		}

	case 4:
		localctx = NewIfSimpleNoElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(414)
			p.Match(MinigoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.SimpleStatement()
		}
		{
			p.SetState(416)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.expression(0)
		}
		{
			p.SetState(418)
			p.Block()
		}

	case 5:
		localctx = NewIfSimpleElseIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(420)
			p.Match(MinigoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.SimpleStatement()
		}
		{
			p.SetState(422)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.expression(0)
		}
		{
			p.SetState(424)
			p.Block()
		}
		{
			p.SetState(425)
			p.Match(MinigoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(426)
			p.IfStatement()
		}

	case 6:
		localctx = NewIfSimpleElseBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(428)
			p.Match(MinigoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)
			p.SimpleStatement()
		}
		{
			p.SetState(430)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.expression(0)
		}
		{
			p.SetState(432)
			p.Block()
		}
		{
			p.SetState(433)
			p.Match(MinigoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
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

func (s *LoopContext) CopyAll(ctx *LoopContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InfiniteForContext struct {
	LoopContext
}

func NewInfiniteForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InfiniteForContext {
	var p = new(InfiniteForContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *InfiniteForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfiniteForContext) FOR() antlr.TerminalNode {
	return s.GetToken(MinigoParserFOR, 0)
}

func (s *InfiniteForContext) Block() IBlockContext {
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

func (s *InfiniteForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterInfiniteFor(s)
	}
}

func (s *InfiniteForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitInfiniteFor(s)
	}
}

func (s *InfiniteForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitInfiniteFor(s)

	default:
		return t.VisitChildren(s)
	}
}

type WhileForContext struct {
	LoopContext
}

func NewWhileForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileForContext {
	var p = new(WhileForContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *WhileForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileForContext) FOR() antlr.TerminalNode {
	return s.GetToken(MinigoParserFOR, 0)
}

func (s *WhileForContext) Expression() IExpressionContext {
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

func (s *WhileForContext) Block() IBlockContext {
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

func (s *WhileForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterWhileFor(s)
	}
}

func (s *WhileForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitWhileFor(s)
	}
}

func (s *WhileForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitWhileFor(s)

	default:
		return t.VisitChildren(s)
	}
}

type ThreePartForContext struct {
	LoopContext
	first ISimpleStatementContext
	last  ISimpleStatementContext
}

func NewThreePartForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThreePartForContext {
	var p = new(ThreePartForContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *ThreePartForContext) GetFirst() ISimpleStatementContext { return s.first }

func (s *ThreePartForContext) GetLast() ISimpleStatementContext { return s.last }

func (s *ThreePartForContext) SetFirst(v ISimpleStatementContext) { s.first = v }

func (s *ThreePartForContext) SetLast(v ISimpleStatementContext) { s.last = v }

func (s *ThreePartForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThreePartForContext) FOR() antlr.TerminalNode {
	return s.GetToken(MinigoParserFOR, 0)
}

func (s *ThreePartForContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(MinigoParserSEMICOLON)
}

func (s *ThreePartForContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, i)
}

func (s *ThreePartForContext) Block() IBlockContext {
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

func (s *ThreePartForContext) AllSimpleStatement() []ISimpleStatementContext {
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

func (s *ThreePartForContext) SimpleStatement(i int) ISimpleStatementContext {
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

func (s *ThreePartForContext) Expression() IExpressionContext {
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

func (s *ThreePartForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterThreePartFor(s)
	}
}

func (s *ThreePartForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitThreePartFor(s)
	}
}

func (s *ThreePartForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitThreePartFor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, MinigoParserRULE_loop)
	var _la int

	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInfiniteForContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(438)
			p.Match(MinigoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(439)
			p.Block()
		}

	case 2:
		localctx = NewWhileForContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(440)
			p.Match(MinigoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(441)
			p.expression(0)
		}
		{
			p.SetState(442)
			p.Block()
		}

	case 3:
		localctx = NewThreePartForContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(444)
			p.Match(MinigoParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)

			var _x = p.SimpleStatement()

			localctx.(*ThreePartForContext).first = _x
		}
		{
			p.SetState(446)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(448)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2259923744325640) != 0) || _la == MinigoParserINTLITERAL || _la == MinigoParserIDENTIFIER {
			{
				p.SetState(447)
				p.expression(0)
			}

		}
		{
			p.SetState(450)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(451)

			var _x = p.SimpleStatement()

			localctx.(*ThreePartForContext).last = _x
		}
		{
			p.SetState(452)
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

func (s *SwitchContext) CopyAll(ctx *SwitchContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SimpleStatementSwitchExpressionContext struct {
	SwitchContext
}

func NewSimpleStatementSwitchExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleStatementSwitchExpressionContext {
	var p = new(SimpleStatementSwitchExpressionContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *SimpleStatementSwitchExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementSwitchExpressionContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(MinigoParserSWITCH, 0)
}

func (s *SimpleStatementSwitchExpressionContext) SimpleStatement() ISimpleStatementContext {
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

func (s *SimpleStatementSwitchExpressionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *SimpleStatementSwitchExpressionContext) Expression() IExpressionContext {
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

func (s *SimpleStatementSwitchExpressionContext) LEFTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTCURLYBRACE, 0)
}

func (s *SimpleStatementSwitchExpressionContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
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

func (s *SimpleStatementSwitchExpressionContext) RIGHTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTCURLYBRACE, 0)
}

func (s *SimpleStatementSwitchExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSimpleStatementSwitchExpression(s)
	}
}

func (s *SimpleStatementSwitchExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSimpleStatementSwitchExpression(s)
	}
}

func (s *SimpleStatementSwitchExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSimpleStatementSwitchExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NormalSwitchContext struct {
	SwitchContext
}

func NewNormalSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NormalSwitchContext {
	var p = new(NormalSwitchContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *NormalSwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NormalSwitchContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(MinigoParserSWITCH, 0)
}

func (s *NormalSwitchContext) LEFTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTCURLYBRACE, 0)
}

func (s *NormalSwitchContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
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

func (s *NormalSwitchContext) RIGHTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTCURLYBRACE, 0)
}

func (s *NormalSwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterNormalSwitch(s)
	}
}

func (s *NormalSwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitNormalSwitch(s)
	}
}

func (s *NormalSwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitNormalSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleStatementSwitchContext struct {
	SwitchContext
}

func NewSimpleStatementSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleStatementSwitchContext {
	var p = new(SimpleStatementSwitchContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *SimpleStatementSwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementSwitchContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(MinigoParserSWITCH, 0)
}

func (s *SimpleStatementSwitchContext) SimpleStatement() ISimpleStatementContext {
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

func (s *SimpleStatementSwitchContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MinigoParserSEMICOLON, 0)
}

func (s *SimpleStatementSwitchContext) LEFTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTCURLYBRACE, 0)
}

func (s *SimpleStatementSwitchContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
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

func (s *SimpleStatementSwitchContext) RIGHTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTCURLYBRACE, 0)
}

func (s *SimpleStatementSwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSimpleStatementSwitch(s)
	}
}

func (s *SimpleStatementSwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSimpleStatementSwitch(s)
	}
}

func (s *SimpleStatementSwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSimpleStatementSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

type NormalSwitchExpressionContext struct {
	SwitchContext
}

func NewNormalSwitchExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NormalSwitchExpressionContext {
	var p = new(NormalSwitchExpressionContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *NormalSwitchExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NormalSwitchExpressionContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(MinigoParserSWITCH, 0)
}

func (s *NormalSwitchExpressionContext) Expression() IExpressionContext {
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

func (s *NormalSwitchExpressionContext) LEFTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserLEFTCURLYBRACE, 0)
}

func (s *NormalSwitchExpressionContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
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

func (s *NormalSwitchExpressionContext) RIGHTCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(MinigoParserRIGHTCURLYBRACE, 0)
}

func (s *NormalSwitchExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterNormalSwitchExpression(s)
	}
}

func (s *NormalSwitchExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitNormalSwitchExpression(s)
	}
}

func (s *NormalSwitchExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitNormalSwitchExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) Switch_() (localctx ISwitchContext) {
	localctx = NewSwitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, MinigoParserRULE_switch)
	p.SetState(482)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimpleStatementSwitchExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(456)
			p.Match(MinigoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.SimpleStatement()
		}
		{
			p.SetState(458)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(459)
			p.expression(0)
		}
		{
			p.SetState(460)
			p.Match(MinigoParserLEFTCURLYBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(462)
			p.Match(MinigoParserRIGHTCURLYBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewNormalSwitchContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(464)
			p.Match(MinigoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(465)
			p.Match(MinigoParserLEFTCURLYBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(466)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(467)
			p.Match(MinigoParserRIGHTCURLYBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewNormalSwitchExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(469)
			p.Match(MinigoParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)
			p.expression(0)
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
		localctx = NewSimpleStatementSwitchContext(p, localctx)
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
			p.SimpleStatement()
		}
		{
			p.SetState(477)
			p.Match(MinigoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(478)
			p.Match(MinigoParserLEFTCURLYBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(480)
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
	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(484)
				p.ExpressionCaseClause()
			}
			{
				p.SetState(485)
				p.ExpressionCaseClauseList()
			}

		}
		p.SetState(491)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
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
		p.SetState(492)
		p.ExpressionSwitchCase()
	}
	{
		p.SetState(493)
		p.Match(MinigoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(494)
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

func (s *ExpressionSwitchCaseContext) CopyAll(ctx *ExpressionSwitchCaseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionSwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchCaseBranchContext struct {
	ExpressionSwitchCaseContext
}

func NewSwitchCaseBranchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchCaseBranchContext {
	var p = new(SwitchCaseBranchContext)

	InitEmptyExpressionSwitchCaseContext(&p.ExpressionSwitchCaseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSwitchCaseContext))

	return p
}

func (s *SwitchCaseBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseBranchContext) CASE() antlr.TerminalNode {
	return s.GetToken(MinigoParserCASE, 0)
}

func (s *SwitchCaseBranchContext) ExpressionList() IExpressionListContext {
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

func (s *SwitchCaseBranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSwitchCaseBranch(s)
	}
}

func (s *SwitchCaseBranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSwitchCaseBranch(s)
	}
}

func (s *SwitchCaseBranchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSwitchCaseBranch(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchDefaultBranchContext struct {
	ExpressionSwitchCaseContext
}

func NewSwitchDefaultBranchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchDefaultBranchContext {
	var p = new(SwitchDefaultBranchContext)

	InitEmptyExpressionSwitchCaseContext(&p.ExpressionSwitchCaseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSwitchCaseContext))

	return p
}

func (s *SwitchDefaultBranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchDefaultBranchContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(MinigoParserDEFAULT, 0)
}

func (s *SwitchDefaultBranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.EnterSwitchDefaultBranch(s)
	}
}

func (s *SwitchDefaultBranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinigoListener); ok {
		listenerT.ExitSwitchDefaultBranch(s)
	}
}

func (s *SwitchDefaultBranchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinigoVisitor:
		return t.VisitSwitchDefaultBranch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinigoParser) ExpressionSwitchCase() (localctx IExpressionSwitchCaseContext) {
	localctx = NewExpressionSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, MinigoParserRULE_expressionSwitchCase)
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MinigoParserCASE:
		localctx = NewSwitchCaseBranchContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(496)
			p.Match(MinigoParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(497)
			p.ExpressionList()
		}

	case MinigoParserDEFAULT:
		localctx = NewSwitchDefaultBranchContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(498)
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
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MinigoParser) PrimaryExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
