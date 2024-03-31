package qbe

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/zSnails/minigo/grammar"
	"github.com/zSnails/minigo/slices"
)

type QBEBackendVisitor struct {
	strings.Builder
}

// VisitBlockStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitBlockStatement(ctx *grammar.BlockStatementContext) interface{} {
	panic("unimplemented")
}

// VisitBreakStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitBreakStatement(ctx *grammar.BreakStatementContext) interface{} {
	panic("unimplemented")
}

// VisitContinueStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitContinueStatement(ctx *grammar.ContinueStatementContext) interface{} {
	panic("unimplemented")
}

// VisitIfStatementStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitIfStatementStatement(ctx *grammar.IfStatementStatementContext) interface{} {
	panic("unimplemented")
}

// VisitLoopStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitLoopStatement(ctx *grammar.LoopStatementContext) interface{} {
	panic("unimplemented")
}

// VisitPrintStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitPrintStatement(ctx *grammar.PrintStatementContext) interface{} {
	expressionList := ctx.ExpressionList()
	q.VisitExpressionList(expressionList.(*grammar.ExpressionListContext))
	return nil
}

// VisitPrintlnStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitPrintlnStatement(ctx *grammar.PrintlnStatementContext) interface{} {
	panic("unimplemented")
}

// VisitReturnStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitReturnStatement(ctx *grammar.ReturnStatementContext) interface{} {
	panic("unimplemented")
}

// VisitSimpleStatementStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitSimpleStatementStatement(ctx *grammar.SimpleStatementStatementContext) interface{} {
	panic("unimplemented")
}

// VisitSwitchStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitSwitchStatement(ctx *grammar.SwitchStatementContext) interface{} {
	panic("unimplemented")
}

// VisitTypeDeclStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitTypeDeclStatement(ctx *grammar.TypeDeclStatementContext) interface{} {
	panic("unimplemented")
}

// VisitVariableDeclStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitVariableDeclStatement(ctx *grammar.VariableDeclStatementContext) interface{} {
	panic("unimplemented")
}

func NewQBEBackendVisitor() *QBEBackendVisitor {
	return &QBEBackendVisitor{
		Builder: strings.Builder{},
	}
}

// VisitEmptyVariableDeclaration implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitEmptyVariableDeclaration(ctx *grammar.EmptyVariableDeclarationContext) interface{} {
	panic("unimplemented")
}

// VisitMultiVariableDeclaration implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitMultiVariableDeclaration(ctx *grammar.MultiVariableDeclarationContext) interface{} {
	panic("unimplemented")
}

// VisitVariableDeclaration implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitVariableDeclaration(ctx *grammar.VariableDeclarationContext) interface{} {
	panic("unimplemented")
}

// Visit implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) Visit(tree antlr.ParseTree) interface{} {
	panic("unimplemented: this function will not be implemented and shall not be called")
}

// VisitAppendExpression implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitAppendExpression(ctx *grammar.AppendExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitArguments implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitArguments(ctx *grammar.ArgumentsContext) interface{} {
	panic("unimplemented")
}

// VisitArrayDeclType implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitArrayDeclType(ctx *grammar.ArrayDeclTypeContext) interface{} {
	panic("unimplemented")
}

// VisitAssignmentStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitAssignmentStatement(ctx *grammar.AssignmentStatementContext) interface{} {
	panic("unimplemented")
}

// VisitBlock implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitBlock(ctx *grammar.BlockContext) interface{} {
	slist := ctx.StatementList()
	if slist != nil {
		q.VisitStatementList(slist.(*grammar.StatementListContext))
	}
	return nil
}

// VisitCapExpression implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitCapExpression(ctx *grammar.CapExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitChildren implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	panic("unimplemented")
}

// VisitDeclType implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitDeclType(ctx *grammar.DeclTypeContext) interface{} {
	panic("unimplemented")
}

// VisitErrorNode implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("unimplemented")
}

// VisitExpression implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitExpression(ctx *grammar.ExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitExpressionCaseClause implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitExpressionCaseClause(ctx *grammar.ExpressionCaseClauseContext) interface{} {
	panic("unimplemented")
}

// VisitExpressionCaseClauseList implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitExpressionCaseClauseList(ctx *grammar.ExpressionCaseClauseListContext) interface{} {
	panic("unimplemented")
}

// VisitExpressionList implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitExpressionList(ctx *grammar.ExpressionListContext) interface{} {
	panic("unimplemented")
}

// VisitExpressionSwitchCase implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitExpressionSwitchCase(ctx *grammar.ExpressionSwitchCaseContext) interface{} {
	panic("unimplemented")
}

// b = byte, h = half word
// datatypes b, h w, l, s, d = word, long, single, double, single and double represent single and double precision floats
var typeLookup = map[string]string{
	"byte":    "b",
	"int8":    "b",
	"int16":   "h",
	"int":     "w",
	"int32":   "w",
	"int64":   "l",
	"float32": "s",
	"float64": "d",
	"string":  ":string",
}

// VisitFuncArgsDecls implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitFuncArgsDecls(ctx *grammar.FuncArgsDeclsContext) interface{} {
	// TODO: implement custom struct types
	for idx, argument := range ctx.AllSingleVarDeclNoExps() {
		if idx > 0 {
			fmt.Fprint(q, ", ")
		}
		idents := slices.Map(argument.IdentifierList().AllIDENTIFIER(), func(node antlr.TerminalNode) string {
			return node.GetText()
		})
		for jdx, ident := range idents {
			if jdx > 0 {
				fmt.Fprint(q, ", ")
			}
			fmt.Fprintf(q, "%s %%%s", typeLookup[argument.DeclType().IDENTIFIER().GetText()], ident)
		}
		// fmt.Fprintf(q, "%s %%%s", typeLookup[argument.DeclType().IDENTIFIER().GetText()], argument.IdentifierList().IDENTIFIER)
	}
	return nil
}

// VisitFuncDecl implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitFuncDecl(ctx *grammar.FuncDeclContext) interface{} {
	_func := ctx.FuncFrontDecl()
	name := _func.IDENTIFIER()
	fmt.Fprintf(q, "function $%s (", name)
	arguments := _func.FuncArgsDecls()
	if arguments != nil {
		q.VisitFuncArgsDecls(arguments.(*grammar.FuncArgsDeclsContext))
	}
	fmt.Fprint(q, ") {\n@start\n")

	block := ctx.Block()
	if block != nil {
		q.VisitBlock(block.(*grammar.BlockContext))
	}
	fmt.Fprint(q, "\n}")
	// TODO: implement function block
	return nil
}

// VisitFuncFrontDecl implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitFuncFrontDecl(ctx *grammar.FuncFrontDeclContext) interface{} {
	panic("unimplemented")
}

// VisitIdentifierList implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitIdentifierList(ctx *grammar.IdentifierListContext) interface{} {
	panic("unimplemented")
}

// VisitIfStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitIfStatement(ctx *grammar.IfStatementContext) interface{} {
	panic("unimplemented")
}

// VisitIndex implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitIndex(ctx *grammar.IndexContext) interface{} {
	panic("unimplemented")
}

// VisitInnerTypeDecls implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitInnerTypeDecls(ctx *grammar.InnerTypeDeclsContext) interface{} {
	panic("unimplemented")
}

// VisitInnerVarDecls implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitInnerVarDecls(ctx *grammar.InnerVarDeclsContext) interface{} {
	panic("unimplemented")
}

// VisitLengthExpression implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitLengthExpression(ctx *grammar.LengthExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitLiteral implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitLiteral(ctx *grammar.LiteralContext) interface{} {
	panic("unimplemented")
}

// VisitLoop implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitLoop(ctx *grammar.LoopContext) interface{} {
	panic("unimplemented")
}

// VisitOperand implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitOperand(ctx *grammar.OperandContext) interface{} {
	panic("unimplemented")
}

// VisitPrimaryExpression implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitPrimaryExpression(ctx *grammar.PrimaryExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitRoot implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitRoot(ctx *grammar.RootContext) interface{} {
	declList := ctx.TopDeclarationList()
	types := declList.AllTypeDecl()
	variables := declList.AllVariableDecl()
	funcs := declList.AllFuncDecl()

	for _, _type := range types {
		q.VisitTypeDecl(_type.(*grammar.TypeDeclContext))
	}

	for _, variable := range variables {
		q.VisitVariableDecl(variable.(*grammar.VariableDeclContext))
	}

	for _, _func := range funcs {
		q.VisitFuncDecl(_func.(*grammar.FuncDeclContext))
	}

	// TODO: handle the return type of those programs
	return nil
}

// VisitSelector implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitSelector(ctx *grammar.SelectorContext) interface{} {
	panic("unimplemented")
}

// VisitSimpleStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitSimpleStatement(ctx *grammar.SimpleStatementContext) interface{} {
	panic("unimplemented")
}

// VisitSingleTypeDecl implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitSingleTypeDecl(ctx *grammar.SingleTypeDeclContext) interface{} {
	panic("unimplemented")
}

// VisitSingleVarDecl implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitSingleVarDecl(ctx *grammar.SingleVarDeclContext) interface{} {
	panic("unimplemented")
}

// VisitSingleVarDeclNoExps implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitSingleVarDeclNoExps(ctx *grammar.SingleVarDeclNoExpsContext) interface{} {
	panic("unimplemented")
}

// VisitSliceDeclType implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitSliceDeclType(ctx *grammar.SliceDeclTypeContext) interface{} {
	panic("unimplemented")
}

// VisitStatement implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitStatement(ctx *grammar.StatementContext) interface{} {
	ruleCtx := ctx.GetRuleContext()
	switch ruleCtx := ruleCtx.(type) {
	case *grammar.PrintStatementContext:
		q.VisitPrintStatement(ruleCtx)
	default:
		panic("sexo tilin")
	}

	return nil
}

// VisitStatementList implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitStatementList(ctx *grammar.StatementListContext) interface{} {
	statements := ctx.AllStatement()
	for _, statement := range statements {
		switch statement := statement.(type) {
		case *grammar.PrintStatementContext:
			q.VisitPrintStatement(statement)
		default:
			panic("sexo tilin")
		}
	}

	return nil
}

// VisitStructDeclType implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitStructDeclType(ctx *grammar.StructDeclTypeContext) interface{} {
	panic("unimplemented")
}

// VisitStructMemDecls implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitStructMemDecls(ctx *grammar.StructMemDeclsContext) interface{} {
	panic("unimplemented")
}

// VisitSwitch implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitSwitch(ctx *grammar.SwitchContext) interface{} {
	panic("unimplemented")
}

// VisitTerminal implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("unimplemented")
}

// VisitTopDeclarationList implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitTopDeclarationList(ctx *grammar.TopDeclarationListContext) interface{} {
	panic("unimplemented")
}

// VisitTypeDecl implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitTypeDecl(ctx *grammar.TypeDeclContext) interface{} {
	panic("unimplemented")
}

// VisitVariableDecl implements grammar.MinigoVisitor.
func (q *QBEBackendVisitor) VisitVariableDecl(ctx *grammar.VariableDeclContext) interface{} {
	panic("unimplemented")
}

// func test() {
// 	a(&QBEBackendVisitor{})
// }

// func a(grammar.MinigoVisitor) {}
