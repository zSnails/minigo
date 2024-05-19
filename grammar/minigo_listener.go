// Code generated from Minigo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package grammar // Minigo
import "github.com/antlr4-go/antlr/v4"

// MinigoListener is a complete listener for a parse tree produced by MinigoParser.
type MinigoListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterTopDeclarationList is called when entering the topDeclarationList production.
	EnterTopDeclarationList(c *TopDeclarationListContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterMultiVariableDeclaration is called when entering the multiVariableDeclaration production.
	EnterMultiVariableDeclaration(c *MultiVariableDeclarationContext)

	// EnterEmptyVariableDeclaration is called when entering the emptyVariableDeclaration production.
	EnterEmptyVariableDeclaration(c *EmptyVariableDeclarationContext)

	// EnterInnerVarDecls is called when entering the innerVarDecls production.
	EnterInnerVarDecls(c *InnerVarDeclsContext)

	// EnterTypedVarDecl is called when entering the typedVarDecl production.
	EnterTypedVarDecl(c *TypedVarDeclContext)

	// EnterUntypedVarDecl is called when entering the untypedVarDecl production.
	EnterUntypedVarDecl(c *UntypedVarDeclContext)

	// EnterSingleVarDeclsNoExpsDecl is called when entering the singleVarDeclsNoExpsDecl production.
	EnterSingleVarDeclsNoExpsDecl(c *SingleVarDeclsNoExpsDeclContext)

	// EnterSingleVarDeclNoExps is called when entering the singleVarDeclNoExps production.
	EnterSingleVarDeclNoExps(c *SingleVarDeclNoExpsContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterMultiTypeDeclaration is called when entering the multiTypeDeclaration production.
	EnterMultiTypeDeclaration(c *MultiTypeDeclarationContext)

	// EnterEmptyTypeDeclaration is called when entering the emptyTypeDeclaration production.
	EnterEmptyTypeDeclaration(c *EmptyTypeDeclarationContext)

	// EnterInnerTypeDecls is called when entering the innerTypeDecls production.
	EnterInnerTypeDecls(c *InnerTypeDeclsContext)

	// EnterSingleTypeDecl is called when entering the singleTypeDecl production.
	EnterSingleTypeDecl(c *SingleTypeDeclContext)

	// EnterFuncDecl is called when entering the funcDecl production.
	EnterFuncDecl(c *FuncDeclContext)

	// EnterFuncDef is called when entering the funcDef production.
	EnterFuncDef(c *FuncDefContext)

	// EnterFuncFrontDecl is called when entering the funcFrontDecl production.
	EnterFuncFrontDecl(c *FuncFrontDeclContext)

	// EnterFuncArgsDecls is called when entering the funcArgsDecls production.
	EnterFuncArgsDecls(c *FuncArgsDeclsContext)

	// EnterNestedType is called when entering the nestedType production.
	EnterNestedType(c *NestedTypeContext)

	// EnterIdentifierDeclType is called when entering the identifierDeclType production.
	EnterIdentifierDeclType(c *IdentifierDeclTypeContext)

	// EnterSliceType is called when entering the sliceType production.
	EnterSliceType(c *SliceTypeContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterStructType is called when entering the structType production.
	EnterStructType(c *StructTypeContext)

	// EnterSliceDeclType is called when entering the sliceDeclType production.
	EnterSliceDeclType(c *SliceDeclTypeContext)

	// EnterArrayDeclType is called when entering the arrayDeclType production.
	EnterArrayDeclType(c *ArrayDeclTypeContext)

	// EnterStructDeclType is called when entering the structDeclType production.
	EnterStructDeclType(c *StructDeclTypeContext)

	// EnterStructMemDecls is called when entering the structMemDecls production.
	EnterStructMemDecls(c *StructMemDeclsContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterOperationPrecedence1 is called when entering the operationPrecedence1 production.
	EnterOperationPrecedence1(c *OperationPrecedence1Context)

	// EnterExpressionPrimaryExpression is called when entering the expressionPrimaryExpression production.
	EnterExpressionPrimaryExpression(c *ExpressionPrimaryExpressionContext)

	// EnterOperationPrecedence2 is called when entering the operationPrecedence2 production.
	EnterOperationPrecedence2(c *OperationPrecedence2Context)

	// EnterNotExpression is called when entering the notExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterCaretExpression is called when entering the caretExpression production.
	EnterCaretExpression(c *CaretExpressionContext)

	// EnterBooleanOperation is called when entering the booleanOperation production.
	EnterBooleanOperation(c *BooleanOperationContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterSubIndex is called when entering the subIndex production.
	EnterSubIndex(c *SubIndexContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterCapCall is called when entering the capCall production.
	EnterCapCall(c *CapCallContext)

	// EnterOperandExpression is called when entering the operandExpression production.
	EnterOperandExpression(c *OperandExpressionContext)

	// EnterAppendCall is called when entering the appendCall production.
	EnterAppendCall(c *AppendCallContext)

	// EnterLenCall is called when entering the lenCall production.
	EnterLenCall(c *LenCallContext)

	// EnterMemberAccessor is called when entering the memberAccessor production.
	EnterMemberAccessor(c *MemberAccessorContext)

	// EnterLiteralOperand is called when entering the literalOperand production.
	EnterLiteralOperand(c *LiteralOperandContext)

	// EnterIdentifierOperand is called when entering the identifierOperand production.
	EnterIdentifierOperand(c *IdentifierOperandContext)

	// EnterExpressionOperand is called when entering the expressionOperand production.
	EnterExpressionOperand(c *ExpressionOperandContext)

	// EnterIntLiteral is called when entering the intLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterRuneLiteral is called when entering the runeLiteral production.
	EnterRuneLiteral(c *RuneLiteralContext)

	// EnterRawStringLiteral is called when entering the rawStringLiteral production.
	EnterRawStringLiteral(c *RawStringLiteralContext)

	// EnterInterpretedStringLiteral is called when entering the interpretedStringLiteral production.
	EnterInterpretedStringLiteral(c *InterpretedStringLiteralContext)

	// EnterNumericIntLiteral is called when entering the numericIntLiteral production.
	EnterNumericIntLiteral(c *NumericIntLiteralContext)

	// EnterNumerixHexLiteral is called when entering the numerixHexLiteral production.
	EnterNumerixHexLiteral(c *NumerixHexLiteralContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterAppendExpression is called when entering the appendExpression production.
	EnterAppendExpression(c *AppendExpressionContext)

	// EnterLengthExpression is called when entering the lengthExpression production.
	EnterLengthExpression(c *LengthExpressionContext)

	// EnterCapExpression is called when entering the capExpression production.
	EnterCapExpression(c *CapExpressionContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterPrintlnStatement is called when entering the printlnStatement production.
	EnterPrintlnStatement(c *PrintlnStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterSimpleStatementStatement is called when entering the simpleStatementStatement production.
	EnterSimpleStatementStatement(c *SimpleStatementStatementContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterIfStatementStatement is called when entering the ifStatementStatement production.
	EnterIfStatementStatement(c *IfStatementStatementContext)

	// EnterLoopStatement is called when entering the loopStatement production.
	EnterLoopStatement(c *LoopStatementContext)

	// EnterTypeDeclStatement is called when entering the typeDeclStatement production.
	EnterTypeDeclStatement(c *TypeDeclStatementContext)

	// EnterVariableDeclStatement is called when entering the variableDeclStatement production.
	EnterVariableDeclStatement(c *VariableDeclStatementContext)

	// EnterExpressionSimpleStatement is called when entering the expressionSimpleStatement production.
	EnterExpressionSimpleStatement(c *ExpressionSimpleStatementContext)

	// EnterExpressionPostInc is called when entering the expressionPostInc production.
	EnterExpressionPostInc(c *ExpressionPostIncContext)

	// EnterExpressionPostDec is called when entering the expressionPostDec production.
	EnterExpressionPostDec(c *ExpressionPostDecContext)

	// EnterAssignmentSimpleStatement is called when entering the assignmentSimpleStatement production.
	EnterAssignmentSimpleStatement(c *AssignmentSimpleStatementContext)

	// EnterWalrusDeclaration is called when entering the walrusDeclaration production.
	EnterWalrusDeclaration(c *WalrusDeclarationContext)

	// EnterNormalAssignment is called when entering the normalAssignment production.
	EnterNormalAssignment(c *NormalAssignmentContext)

	// EnterInPlaceAssignment is called when entering the inPlaceAssignment production.
	EnterInPlaceAssignment(c *InPlaceAssignmentContext)

	// EnterIfSingleExpression is called when entering the ifSingleExpression production.
	EnterIfSingleExpression(c *IfSingleExpressionContext)

	// EnterIfElseIf is called when entering the ifElseIf production.
	EnterIfElseIf(c *IfElseIfContext)

	// EnterIfElseBlock is called when entering the ifElseBlock production.
	EnterIfElseBlock(c *IfElseBlockContext)

	// EnterIfSimpleNoElse is called when entering the ifSimpleNoElse production.
	EnterIfSimpleNoElse(c *IfSimpleNoElseContext)

	// EnterIfSimpleElseIf is called when entering the ifSimpleElseIf production.
	EnterIfSimpleElseIf(c *IfSimpleElseIfContext)

	// EnterIfSimpleElseBlock is called when entering the ifSimpleElseBlock production.
	EnterIfSimpleElseBlock(c *IfSimpleElseBlockContext)

	// EnterInfiniteFor is called when entering the infiniteFor production.
	EnterInfiniteFor(c *InfiniteForContext)

	// EnterWhileFor is called when entering the whileFor production.
	EnterWhileFor(c *WhileForContext)

	// EnterThreePartFor is called when entering the threePartFor production.
	EnterThreePartFor(c *ThreePartForContext)

	// EnterSimpleStatementSwitchExpression is called when entering the simpleStatementSwitchExpression production.
	EnterSimpleStatementSwitchExpression(c *SimpleStatementSwitchExpressionContext)

	// EnterNormalSwitch is called when entering the normalSwitch production.
	EnterNormalSwitch(c *NormalSwitchContext)

	// EnterNormalSwitchExpression is called when entering the normalSwitchExpression production.
	EnterNormalSwitchExpression(c *NormalSwitchExpressionContext)

	// EnterSimpleStatementSwitch is called when entering the simpleStatementSwitch production.
	EnterSimpleStatementSwitch(c *SimpleStatementSwitchContext)

	// EnterExpressionCaseClauseList is called when entering the expressionCaseClauseList production.
	EnterExpressionCaseClauseList(c *ExpressionCaseClauseListContext)

	// EnterExpressionCaseClause is called when entering the expressionCaseClause production.
	EnterExpressionCaseClause(c *ExpressionCaseClauseContext)

	// EnterSwitchCaseBranch is called when entering the switchCaseBranch production.
	EnterSwitchCaseBranch(c *SwitchCaseBranchContext)

	// EnterSwitchDefaultBranch is called when entering the switchDefaultBranch production.
	EnterSwitchDefaultBranch(c *SwitchDefaultBranchContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitTopDeclarationList is called when exiting the topDeclarationList production.
	ExitTopDeclarationList(c *TopDeclarationListContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitMultiVariableDeclaration is called when exiting the multiVariableDeclaration production.
	ExitMultiVariableDeclaration(c *MultiVariableDeclarationContext)

	// ExitEmptyVariableDeclaration is called when exiting the emptyVariableDeclaration production.
	ExitEmptyVariableDeclaration(c *EmptyVariableDeclarationContext)

	// ExitInnerVarDecls is called when exiting the innerVarDecls production.
	ExitInnerVarDecls(c *InnerVarDeclsContext)

	// ExitTypedVarDecl is called when exiting the typedVarDecl production.
	ExitTypedVarDecl(c *TypedVarDeclContext)

	// ExitUntypedVarDecl is called when exiting the untypedVarDecl production.
	ExitUntypedVarDecl(c *UntypedVarDeclContext)

	// ExitSingleVarDeclsNoExpsDecl is called when exiting the singleVarDeclsNoExpsDecl production.
	ExitSingleVarDeclsNoExpsDecl(c *SingleVarDeclsNoExpsDeclContext)

	// ExitSingleVarDeclNoExps is called when exiting the singleVarDeclNoExps production.
	ExitSingleVarDeclNoExps(c *SingleVarDeclNoExpsContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitMultiTypeDeclaration is called when exiting the multiTypeDeclaration production.
	ExitMultiTypeDeclaration(c *MultiTypeDeclarationContext)

	// ExitEmptyTypeDeclaration is called when exiting the emptyTypeDeclaration production.
	ExitEmptyTypeDeclaration(c *EmptyTypeDeclarationContext)

	// ExitInnerTypeDecls is called when exiting the innerTypeDecls production.
	ExitInnerTypeDecls(c *InnerTypeDeclsContext)

	// ExitSingleTypeDecl is called when exiting the singleTypeDecl production.
	ExitSingleTypeDecl(c *SingleTypeDeclContext)

	// ExitFuncDecl is called when exiting the funcDecl production.
	ExitFuncDecl(c *FuncDeclContext)

	// ExitFuncDef is called when exiting the funcDef production.
	ExitFuncDef(c *FuncDefContext)

	// ExitFuncFrontDecl is called when exiting the funcFrontDecl production.
	ExitFuncFrontDecl(c *FuncFrontDeclContext)

	// ExitFuncArgsDecls is called when exiting the funcArgsDecls production.
	ExitFuncArgsDecls(c *FuncArgsDeclsContext)

	// ExitNestedType is called when exiting the nestedType production.
	ExitNestedType(c *NestedTypeContext)

	// ExitIdentifierDeclType is called when exiting the identifierDeclType production.
	ExitIdentifierDeclType(c *IdentifierDeclTypeContext)

	// ExitSliceType is called when exiting the sliceType production.
	ExitSliceType(c *SliceTypeContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitStructType is called when exiting the structType production.
	ExitStructType(c *StructTypeContext)

	// ExitSliceDeclType is called when exiting the sliceDeclType production.
	ExitSliceDeclType(c *SliceDeclTypeContext)

	// ExitArrayDeclType is called when exiting the arrayDeclType production.
	ExitArrayDeclType(c *ArrayDeclTypeContext)

	// ExitStructDeclType is called when exiting the structDeclType production.
	ExitStructDeclType(c *StructDeclTypeContext)

	// ExitStructMemDecls is called when exiting the structMemDecls production.
	ExitStructMemDecls(c *StructMemDeclsContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitOperationPrecedence1 is called when exiting the operationPrecedence1 production.
	ExitOperationPrecedence1(c *OperationPrecedence1Context)

	// ExitExpressionPrimaryExpression is called when exiting the expressionPrimaryExpression production.
	ExitExpressionPrimaryExpression(c *ExpressionPrimaryExpressionContext)

	// ExitOperationPrecedence2 is called when exiting the operationPrecedence2 production.
	ExitOperationPrecedence2(c *OperationPrecedence2Context)

	// ExitNotExpression is called when exiting the notExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitCaretExpression is called when exiting the caretExpression production.
	ExitCaretExpression(c *CaretExpressionContext)

	// ExitBooleanOperation is called when exiting the booleanOperation production.
	ExitBooleanOperation(c *BooleanOperationContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitSubIndex is called when exiting the subIndex production.
	ExitSubIndex(c *SubIndexContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitCapCall is called when exiting the capCall production.
	ExitCapCall(c *CapCallContext)

	// ExitOperandExpression is called when exiting the operandExpression production.
	ExitOperandExpression(c *OperandExpressionContext)

	// ExitAppendCall is called when exiting the appendCall production.
	ExitAppendCall(c *AppendCallContext)

	// ExitLenCall is called when exiting the lenCall production.
	ExitLenCall(c *LenCallContext)

	// ExitMemberAccessor is called when exiting the memberAccessor production.
	ExitMemberAccessor(c *MemberAccessorContext)

	// ExitLiteralOperand is called when exiting the literalOperand production.
	ExitLiteralOperand(c *LiteralOperandContext)

	// ExitIdentifierOperand is called when exiting the identifierOperand production.
	ExitIdentifierOperand(c *IdentifierOperandContext)

	// ExitExpressionOperand is called when exiting the expressionOperand production.
	ExitExpressionOperand(c *ExpressionOperandContext)

	// ExitIntLiteral is called when exiting the intLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitRuneLiteral is called when exiting the runeLiteral production.
	ExitRuneLiteral(c *RuneLiteralContext)

	// ExitRawStringLiteral is called when exiting the rawStringLiteral production.
	ExitRawStringLiteral(c *RawStringLiteralContext)

	// ExitInterpretedStringLiteral is called when exiting the interpretedStringLiteral production.
	ExitInterpretedStringLiteral(c *InterpretedStringLiteralContext)

	// ExitNumericIntLiteral is called when exiting the numericIntLiteral production.
	ExitNumericIntLiteral(c *NumericIntLiteralContext)

	// ExitNumerixHexLiteral is called when exiting the numerixHexLiteral production.
	ExitNumerixHexLiteral(c *NumerixHexLiteralContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitAppendExpression is called when exiting the appendExpression production.
	ExitAppendExpression(c *AppendExpressionContext)

	// ExitLengthExpression is called when exiting the lengthExpression production.
	ExitLengthExpression(c *LengthExpressionContext)

	// ExitCapExpression is called when exiting the capExpression production.
	ExitCapExpression(c *CapExpressionContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitPrintlnStatement is called when exiting the printlnStatement production.
	ExitPrintlnStatement(c *PrintlnStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitSimpleStatementStatement is called when exiting the simpleStatementStatement production.
	ExitSimpleStatementStatement(c *SimpleStatementStatementContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitIfStatementStatement is called when exiting the ifStatementStatement production.
	ExitIfStatementStatement(c *IfStatementStatementContext)

	// ExitLoopStatement is called when exiting the loopStatement production.
	ExitLoopStatement(c *LoopStatementContext)

	// ExitTypeDeclStatement is called when exiting the typeDeclStatement production.
	ExitTypeDeclStatement(c *TypeDeclStatementContext)

	// ExitVariableDeclStatement is called when exiting the variableDeclStatement production.
	ExitVariableDeclStatement(c *VariableDeclStatementContext)

	// ExitExpressionSimpleStatement is called when exiting the expressionSimpleStatement production.
	ExitExpressionSimpleStatement(c *ExpressionSimpleStatementContext)

	// ExitExpressionPostInc is called when exiting the expressionPostInc production.
	ExitExpressionPostInc(c *ExpressionPostIncContext)

	// ExitExpressionPostDec is called when exiting the expressionPostDec production.
	ExitExpressionPostDec(c *ExpressionPostDecContext)

	// ExitAssignmentSimpleStatement is called when exiting the assignmentSimpleStatement production.
	ExitAssignmentSimpleStatement(c *AssignmentSimpleStatementContext)

	// ExitWalrusDeclaration is called when exiting the walrusDeclaration production.
	ExitWalrusDeclaration(c *WalrusDeclarationContext)

	// ExitNormalAssignment is called when exiting the normalAssignment production.
	ExitNormalAssignment(c *NormalAssignmentContext)

	// ExitInPlaceAssignment is called when exiting the inPlaceAssignment production.
	ExitInPlaceAssignment(c *InPlaceAssignmentContext)

	// ExitIfSingleExpression is called when exiting the ifSingleExpression production.
	ExitIfSingleExpression(c *IfSingleExpressionContext)

	// ExitIfElseIf is called when exiting the ifElseIf production.
	ExitIfElseIf(c *IfElseIfContext)

	// ExitIfElseBlock is called when exiting the ifElseBlock production.
	ExitIfElseBlock(c *IfElseBlockContext)

	// ExitIfSimpleNoElse is called when exiting the ifSimpleNoElse production.
	ExitIfSimpleNoElse(c *IfSimpleNoElseContext)

	// ExitIfSimpleElseIf is called when exiting the ifSimpleElseIf production.
	ExitIfSimpleElseIf(c *IfSimpleElseIfContext)

	// ExitIfSimpleElseBlock is called when exiting the ifSimpleElseBlock production.
	ExitIfSimpleElseBlock(c *IfSimpleElseBlockContext)

	// ExitInfiniteFor is called when exiting the infiniteFor production.
	ExitInfiniteFor(c *InfiniteForContext)

	// ExitWhileFor is called when exiting the whileFor production.
	ExitWhileFor(c *WhileForContext)

	// ExitThreePartFor is called when exiting the threePartFor production.
	ExitThreePartFor(c *ThreePartForContext)

	// ExitSimpleStatementSwitchExpression is called when exiting the simpleStatementSwitchExpression production.
	ExitSimpleStatementSwitchExpression(c *SimpleStatementSwitchExpressionContext)

	// ExitNormalSwitch is called when exiting the normalSwitch production.
	ExitNormalSwitch(c *NormalSwitchContext)

	// ExitNormalSwitchExpression is called when exiting the normalSwitchExpression production.
	ExitNormalSwitchExpression(c *NormalSwitchExpressionContext)

	// ExitSimpleStatementSwitch is called when exiting the simpleStatementSwitch production.
	ExitSimpleStatementSwitch(c *SimpleStatementSwitchContext)

	// ExitExpressionCaseClauseList is called when exiting the expressionCaseClauseList production.
	ExitExpressionCaseClauseList(c *ExpressionCaseClauseListContext)

	// ExitExpressionCaseClause is called when exiting the expressionCaseClause production.
	ExitExpressionCaseClause(c *ExpressionCaseClauseContext)

	// ExitSwitchCaseBranch is called when exiting the switchCaseBranch production.
	ExitSwitchCaseBranch(c *SwitchCaseBranchContext)

	// ExitSwitchDefaultBranch is called when exiting the switchDefaultBranch production.
	ExitSwitchDefaultBranch(c *SwitchDefaultBranchContext)
}
