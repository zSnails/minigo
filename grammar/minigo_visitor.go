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

	// Visit a parse tree produced by MinigoParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by MinigoParser#multiVariableDeclaration.
	VisitMultiVariableDeclaration(ctx *MultiVariableDeclarationContext) interface{}

	// Visit a parse tree produced by MinigoParser#emptyVariableDeclaration.
	VisitEmptyVariableDeclaration(ctx *EmptyVariableDeclarationContext) interface{}

	// Visit a parse tree produced by MinigoParser#innerVarDecls.
	VisitInnerVarDecls(ctx *InnerVarDeclsContext) interface{}

	// Visit a parse tree produced by MinigoParser#typedVarDecl.
	VisitTypedVarDecl(ctx *TypedVarDeclContext) interface{}

	// Visit a parse tree produced by MinigoParser#untypedVarDecl.
	VisitUntypedVarDecl(ctx *UntypedVarDeclContext) interface{}

	// Visit a parse tree produced by MinigoParser#singleVarDeclsNoExpsDecl.
	VisitSingleVarDeclsNoExpsDecl(ctx *SingleVarDeclsNoExpsDeclContext) interface{}

	// Visit a parse tree produced by MinigoParser#singleVarDeclNoExps.
	VisitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) interface{}

	// Visit a parse tree produced by MinigoParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by MinigoParser#multiTypeDeclaration.
	VisitMultiTypeDeclaration(ctx *MultiTypeDeclarationContext) interface{}

	// Visit a parse tree produced by MinigoParser#emptyTypeDeclaration.
	VisitEmptyTypeDeclaration(ctx *EmptyTypeDeclarationContext) interface{}

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

	// Visit a parse tree produced by MinigoParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by MinigoParser#expressionPrimaryExpression.
	VisitExpressionPrimaryExpression(ctx *ExpressionPrimaryExpressionContext) interface{}

	// Visit a parse tree produced by MinigoParser#notExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by MinigoParser#caretExpression.
	VisitCaretExpression(ctx *CaretExpressionContext) interface{}

	// Visit a parse tree produced by MinigoParser#booleanOperation.
	VisitBooleanOperation(ctx *BooleanOperationContext) interface{}

	// Visit a parse tree produced by MinigoParser#operation.
	VisitOperation(ctx *OperationContext) interface{}

	// Visit a parse tree produced by MinigoParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by MinigoParser#subIndex.
	VisitSubIndex(ctx *SubIndexContext) interface{}

	// Visit a parse tree produced by MinigoParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by MinigoParser#capCall.
	VisitCapCall(ctx *CapCallContext) interface{}

	// Visit a parse tree produced by MinigoParser#operandExpression.
	VisitOperandExpression(ctx *OperandExpressionContext) interface{}

	// Visit a parse tree produced by MinigoParser#appendCall.
	VisitAppendCall(ctx *AppendCallContext) interface{}

	// Visit a parse tree produced by MinigoParser#lenCall.
	VisitLenCall(ctx *LenCallContext) interface{}

	// Visit a parse tree produced by MinigoParser#memberAccessor.
	VisitMemberAccessor(ctx *MemberAccessorContext) interface{}

	// Visit a parse tree produced by MinigoParser#literalOperand.
	VisitLiteralOperand(ctx *LiteralOperandContext) interface{}

	// Visit a parse tree produced by MinigoParser#identifierOperand.
	VisitIdentifierOperand(ctx *IdentifierOperandContext) interface{}

	// Visit a parse tree produced by MinigoParser#expressionOperand.
	VisitExpressionOperand(ctx *ExpressionOperandContext) interface{}

	// Visit a parse tree produced by MinigoParser#intLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by MinigoParser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by MinigoParser#runeLiteral.
	VisitRuneLiteral(ctx *RuneLiteralContext) interface{}

	// Visit a parse tree produced by MinigoParser#rawStringLiteral.
	VisitRawStringLiteral(ctx *RawStringLiteralContext) interface{}

	// Visit a parse tree produced by MinigoParser#interpretedStringLiteral.
	VisitInterpretedStringLiteral(ctx *InterpretedStringLiteralContext) interface{}

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

	// Visit a parse tree produced by MinigoParser#printStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#printlnStatement.
	VisitPrintlnStatement(ctx *PrintlnStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#simpleStatementStatement.
	VisitSimpleStatementStatement(ctx *SimpleStatementStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#ifStatementStatement.
	VisitIfStatementStatement(ctx *IfStatementStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#typeDeclStatement.
	VisitTypeDeclStatement(ctx *TypeDeclStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#variableDeclStatement.
	VisitVariableDeclStatement(ctx *VariableDeclStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#expressionSimpleStatement.
	VisitExpressionSimpleStatement(ctx *ExpressionSimpleStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#assignmentSimpleStatement.
	VisitAssignmentSimpleStatement(ctx *AssignmentSimpleStatementContext) interface{}

	// Visit a parse tree produced by MinigoParser#walrusDeclaration.
	VisitWalrusDeclaration(ctx *WalrusDeclarationContext) interface{}

	// Visit a parse tree produced by MinigoParser#normalAssignment.
	VisitNormalAssignment(ctx *NormalAssignmentContext) interface{}

	// Visit a parse tree produced by MinigoParser#inPlaceAssignment.
	VisitInPlaceAssignment(ctx *InPlaceAssignmentContext) interface{}

	// Visit a parse tree produced by MinigoParser#ifSingleExpression.
	VisitIfSingleExpression(ctx *IfSingleExpressionContext) interface{}

	// Visit a parse tree produced by MinigoParser#ifElseIf.
	VisitIfElseIf(ctx *IfElseIfContext) interface{}

	// Visit a parse tree produced by MinigoParser#ifElseBlock.
	VisitIfElseBlock(ctx *IfElseBlockContext) interface{}

	// Visit a parse tree produced by MinigoParser#ifSimpleNoElse.
	VisitIfSimpleNoElse(ctx *IfSimpleNoElseContext) interface{}

	// Visit a parse tree produced by MinigoParser#ifSimpleElseIf.
	VisitIfSimpleElseIf(ctx *IfSimpleElseIfContext) interface{}

	// Visit a parse tree produced by MinigoParser#ifSimpleElseBlock.
	VisitIfSimpleElseBlock(ctx *IfSimpleElseBlockContext) interface{}

	// Visit a parse tree produced by MinigoParser#infiniteFor.
	VisitInfiniteFor(ctx *InfiniteForContext) interface{}

	// Visit a parse tree produced by MinigoParser#whileFor.
	VisitWhileFor(ctx *WhileForContext) interface{}

	// Visit a parse tree produced by MinigoParser#threePartFor.
	VisitThreePartFor(ctx *ThreePartForContext) interface{}

	// Visit a parse tree produced by MinigoParser#simpleStatementSwitchExpression.
	VisitSimpleStatementSwitchExpression(ctx *SimpleStatementSwitchExpressionContext) interface{}

	// Visit a parse tree produced by MinigoParser#normalSwitch.
	VisitNormalSwitch(ctx *NormalSwitchContext) interface{}

	// Visit a parse tree produced by MinigoParser#normalSwitchExpression.
	VisitNormalSwitchExpression(ctx *NormalSwitchExpressionContext) interface{}

	// Visit a parse tree produced by MinigoParser#simpleStatementSwitch.
	VisitSimpleStatementSwitch(ctx *SimpleStatementSwitchContext) interface{}

	// Visit a parse tree produced by MinigoParser#expressionCaseClauseList.
	VisitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) interface{}

	// Visit a parse tree produced by MinigoParser#expressionCaseClause.
	VisitExpressionCaseClause(ctx *ExpressionCaseClauseContext) interface{}

	// Visit a parse tree produced by MinigoParser#switchCaseBranch.
	VisitSwitchCaseBranch(ctx *SwitchCaseBranchContext) interface{}

	// Visit a parse tree produced by MinigoParser#switchDefaultBranch.
	VisitSwitchDefaultBranch(ctx *SwitchDefaultBranchContext) interface{}
}
