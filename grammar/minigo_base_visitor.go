// Code generated from Minigo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package grammar // Minigo
import "github.com/antlr4-go/antlr/v4"

type BaseMinigoVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMinigoVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitTopDeclarationList(ctx *TopDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitMultiVariableDeclaration(ctx *MultiVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitEmptyVariableDeclaration(ctx *EmptyVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitInnerVarDecls(ctx *InnerVarDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSingleVarDecl(ctx *SingleVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitTypeDecl(ctx *TypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitInnerTypeDecls(ctx *InnerTypeDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSingleTypeDecl(ctx *SingleTypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitFuncDecl(ctx *FuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitFuncFrontDecl(ctx *FuncFrontDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitFuncArgsDecls(ctx *FuncArgsDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitDeclType(ctx *DeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSliceDeclType(ctx *SliceDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitArrayDeclType(ctx *ArrayDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitStructDeclType(ctx *StructDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitStructMemDecls(ctx *StructMemDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitAppendExpression(ctx *AppendExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitLengthExpression(ctx *LengthExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitCapExpression(ctx *CapExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitPrintlnStatement(ctx *PrintlnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSimpleStatementStatement(ctx *SimpleStatementStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitIfStatementStatement(ctx *IfStatementStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitTypeDeclStatement(ctx *TypeDeclStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitVariableDeclStatement(ctx *VariableDeclStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSimpleStatement(ctx *SimpleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitLoop(ctx *LoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSwitch(ctx *SwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitExpressionCaseClause(ctx *ExpressionCaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitExpressionSwitchCase(ctx *ExpressionSwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}
