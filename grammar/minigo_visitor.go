// Code generated from Minigo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package grammar // Minigo
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MinigoParser.
type MinigoVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MinigoParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by MinigoParser#topDeclarationList.
	VisitTopDeclarationList(ctx *TopDeclarationListContext) interface{}

	// Visit a parse tree produced by MinigoParser#variableDecl.
	VisitVariableDecl(ctx *VariableDeclContext) interface{}

	// Visit a parse tree produced by MinigoParser#innerVarDecls.
	VisitInnerVarDecls(ctx *InnerVarDeclsContext) interface{}

	// Visit a parse tree produced by MinigoParser#singleVarDecl.
	VisitSingleVarDecl(ctx *SingleVarDeclContext) interface{}

	// Visit a parse tree produced by MinigoParser#singleVarDeclNoExps.
	VisitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) interface{}

	// Visit a parse tree produced by MinigoParser#typeDecl.
	VisitTypeDecl(ctx *TypeDeclContext) interface{}

	// Visit a parse tree produced by MinigoParser#innerTypeDecls.
	VisitInnerTypeDecls(ctx *InnerTypeDeclsContext) interface{}

	// Visit a parse tree produced by MinigoParser#singleTypeDecl.
	VisitSingleTypeDecl(ctx *SingleTypeDeclContext) interface{}

	// Visit a parse tree produced by MinigoParser#funcDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by MinigoParser#funcFrontDecl.
	VisitFuncFrontDecl(ctx *FuncFrontDeclContext) interface{}

	// Visit a parse tree produced by MinigoParser#funcArgsDecls.
	VisitFuncArgsDecls(ctx *FuncArgsDeclsContext) interface{}

	// Visit a parse tree produced by MinigoParser#declType.
	VisitDeclType(ctx *DeclTypeContext) interface{}

	// Visit a parse tree produced by MinigoParser#sliceDeclType.
	VisitSliceDeclType(ctx *SliceDeclTypeContext) interface{}

	// Visit a parse tree produced by MinigoParser#arrayDeclType.
	VisitArrayDeclType(ctx *ArrayDeclTypeContext) interface{}

	// Visit a parse tree produced by MinigoParser#structDeclType.
	VisitStructDeclType(ctx *StructDeclTypeContext) interface{}

	// Visit a parse tree produced by MinigoParser#structMemDecls.
	VisitStructMemDecls(ctx *StructMemDeclsContext) interface{}

	// Visit a parse tree produced by MinigoParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by MinigoParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by MinigoParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by MinigoParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by MinigoParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by MinigoParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by MinigoParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by MinigoParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by MinigoParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by MinigoParser#appendExpression.
	VisitAppendExpression(ctx *AppendExpressionContext) interface{}

	// Visit a parse tree produced by MinigoParser#lengthExpression.
	VisitLengthExpression(ctx *LengthExpressionContext) interface{}

	// Visit a parse tree produced by MinigoParser#capExpression.
	VisitCapExpression(ctx *CapExpressionContext) interface{}

	// Visit a parse tree produced by MinigoParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by MinigoParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by MinigoParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#simpleStatement.
	VisitSimpleStatement(ctx *SimpleStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#assignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#loop.
	VisitLoop(ctx *LoopContext) interface{}

	// Visit a parse tree produced by MinigoParser#switch.
	VisitSwitch(ctx *SwitchContext) interface{}

	// Visit a parse tree produced by MinigoParser#expressionCaseClauseList.
	VisitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) interface{}

	// Visit a parse tree produced by MinigoParser#expressionCaseClause.
	VisitExpressionCaseClause(ctx *ExpressionCaseClauseContext) interface{}

	// Visit a parse tree produced by MinigoParser#expressionSwitchCase.
	VisitExpressionSwitchCase(ctx *ExpressionSwitchCaseContext) interface{}
}
