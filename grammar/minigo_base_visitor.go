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

func (v *BaseMinigoVisitor) VisitTypedVarDecl(ctx *TypedVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitUntypedVarDecl(ctx *UntypedVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSingleVarDeclsNoExpsDecl(ctx *SingleVarDeclsNoExpsDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitMultiTypeDeclaration(ctx *MultiTypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitEmptyTypeDeclaration(ctx *EmptyTypeDeclarationContext) interface{} {
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

func (v *BaseMinigoVisitor) VisitNestedType(ctx *NestedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitIdentifierDeclType(ctx *IdentifierDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSliceType(ctx *SliceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
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

func (v *BaseMinigoVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitOperationPrecedence1(ctx *OperationPrecedence1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitExpressionPrimaryExpression(ctx *ExpressionPrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitOperationPrecedence2(ctx *OperationPrecedence2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitCaretExpression(ctx *CaretExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitBooleanOperation(ctx *BooleanOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSubIndex(ctx *SubIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitCapCall(ctx *CapCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitOperandExpression(ctx *OperandExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitAppendCall(ctx *AppendCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitLenCall(ctx *LenCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitMemberAccessor(ctx *MemberAccessorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitLiteralOperand(ctx *LiteralOperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitIdentifierOperand(ctx *IdentifierOperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitExpressionOperand(ctx *ExpressionOperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitRuneLiteral(ctx *RuneLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitRawStringLiteral(ctx *RawStringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitInterpretedStringLiteral(ctx *InterpretedStringLiteralContext) interface{} {
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

func (v *BaseMinigoVisitor) VisitExpressionSimpleStatement(ctx *ExpressionSimpleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitExpressionPostInc(ctx *ExpressionPostIncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitExpressionPostDec(ctx *ExpressionPostDecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitAssignmentSimpleStatement(ctx *AssignmentSimpleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitWalrusDeclaration(ctx *WalrusDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitNormalAssignment(ctx *NormalAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitInPlaceAssignment(ctx *InPlaceAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitIfSingleExpression(ctx *IfSingleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitIfElseIf(ctx *IfElseIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitIfElseBlock(ctx *IfElseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitIfSimpleNoElse(ctx *IfSimpleNoElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitIfSimpleElseIf(ctx *IfSimpleElseIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitIfSimpleElseBlock(ctx *IfSimpleElseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitInfiniteFor(ctx *InfiniteForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitWhileFor(ctx *WhileForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitThreePartFor(ctx *ThreePartForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSimpleStatementSwitchExpression(ctx *SimpleStatementSwitchExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitNormalSwitch(ctx *NormalSwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitNormalSwitchExpression(ctx *NormalSwitchExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSimpleStatementSwitch(ctx *SimpleStatementSwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitExpressionCaseClause(ctx *ExpressionCaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSwitchCaseBranch(ctx *SwitchCaseBranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinigoVisitor) VisitSwitchDefaultBranch(ctx *SwitchDefaultBranchContext) interface{} {
	return v.VisitChildren(ctx)
}
