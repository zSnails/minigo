// Code generated from Minigo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package grammar // Minigo
import "github.com/antlr4-go/antlr/v4"

// BaseMinigoListener is a complete listener for a parse tree produced by MinigoParser.
type BaseMinigoListener struct{}

var _ MinigoListener = &BaseMinigoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMinigoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMinigoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMinigoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMinigoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseMinigoListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseMinigoListener) ExitRoot(ctx *RootContext) {}

// EnterTopDeclarationList is called when production topDeclarationList is entered.
func (s *BaseMinigoListener) EnterTopDeclarationList(ctx *TopDeclarationListContext) {}

// ExitTopDeclarationList is called when production topDeclarationList is exited.
func (s *BaseMinigoListener) ExitTopDeclarationList(ctx *TopDeclarationListContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseMinigoListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseMinigoListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterMultiVariableDeclaration is called when production multiVariableDeclaration is entered.
func (s *BaseMinigoListener) EnterMultiVariableDeclaration(ctx *MultiVariableDeclarationContext) {}

// ExitMultiVariableDeclaration is called when production multiVariableDeclaration is exited.
func (s *BaseMinigoListener) ExitMultiVariableDeclaration(ctx *MultiVariableDeclarationContext) {}

// EnterEmptyVariableDeclaration is called when production emptyVariableDeclaration is entered.
func (s *BaseMinigoListener) EnterEmptyVariableDeclaration(ctx *EmptyVariableDeclarationContext) {}

// ExitEmptyVariableDeclaration is called when production emptyVariableDeclaration is exited.
func (s *BaseMinigoListener) ExitEmptyVariableDeclaration(ctx *EmptyVariableDeclarationContext) {}

// EnterInnerVarDecls is called when production innerVarDecls is entered.
func (s *BaseMinigoListener) EnterInnerVarDecls(ctx *InnerVarDeclsContext) {}

// ExitInnerVarDecls is called when production innerVarDecls is exited.
func (s *BaseMinigoListener) ExitInnerVarDecls(ctx *InnerVarDeclsContext) {}

// EnterTypedVarDecl is called when production typedVarDecl is entered.
func (s *BaseMinigoListener) EnterTypedVarDecl(ctx *TypedVarDeclContext) {}

// ExitTypedVarDecl is called when production typedVarDecl is exited.
func (s *BaseMinigoListener) ExitTypedVarDecl(ctx *TypedVarDeclContext) {}

// EnterUntypedVarDecl is called when production untypedVarDecl is entered.
func (s *BaseMinigoListener) EnterUntypedVarDecl(ctx *UntypedVarDeclContext) {}

// ExitUntypedVarDecl is called when production untypedVarDecl is exited.
func (s *BaseMinigoListener) ExitUntypedVarDecl(ctx *UntypedVarDeclContext) {}

// EnterSingleVarDeclsNoExpsDecl is called when production singleVarDeclsNoExpsDecl is entered.
func (s *BaseMinigoListener) EnterSingleVarDeclsNoExpsDecl(ctx *SingleVarDeclsNoExpsDeclContext) {}

// ExitSingleVarDeclsNoExpsDecl is called when production singleVarDeclsNoExpsDecl is exited.
func (s *BaseMinigoListener) ExitSingleVarDeclsNoExpsDecl(ctx *SingleVarDeclsNoExpsDeclContext) {}

// EnterSingleVarDeclNoExps is called when production singleVarDeclNoExps is entered.
func (s *BaseMinigoListener) EnterSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) {}

// ExitSingleVarDeclNoExps is called when production singleVarDeclNoExps is exited.
func (s *BaseMinigoListener) ExitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseMinigoListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseMinigoListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterMultiTypeDeclaration is called when production multiTypeDeclaration is entered.
func (s *BaseMinigoListener) EnterMultiTypeDeclaration(ctx *MultiTypeDeclarationContext) {}

// ExitMultiTypeDeclaration is called when production multiTypeDeclaration is exited.
func (s *BaseMinigoListener) ExitMultiTypeDeclaration(ctx *MultiTypeDeclarationContext) {}

// EnterEmptyTypeDeclaration is called when production emptyTypeDeclaration is entered.
func (s *BaseMinigoListener) EnterEmptyTypeDeclaration(ctx *EmptyTypeDeclarationContext) {}

// ExitEmptyTypeDeclaration is called when production emptyTypeDeclaration is exited.
func (s *BaseMinigoListener) ExitEmptyTypeDeclaration(ctx *EmptyTypeDeclarationContext) {}

// EnterInnerTypeDecls is called when production innerTypeDecls is entered.
func (s *BaseMinigoListener) EnterInnerTypeDecls(ctx *InnerTypeDeclsContext) {}

// ExitInnerTypeDecls is called when production innerTypeDecls is exited.
func (s *BaseMinigoListener) ExitInnerTypeDecls(ctx *InnerTypeDeclsContext) {}

// EnterSingleTypeDecl is called when production singleTypeDecl is entered.
func (s *BaseMinigoListener) EnterSingleTypeDecl(ctx *SingleTypeDeclContext) {}

// ExitSingleTypeDecl is called when production singleTypeDecl is exited.
func (s *BaseMinigoListener) ExitSingleTypeDecl(ctx *SingleTypeDeclContext) {}

// EnterFuncDecl is called when production funcDecl is entered.
func (s *BaseMinigoListener) EnterFuncDecl(ctx *FuncDeclContext) {}

// ExitFuncDecl is called when production funcDecl is exited.
func (s *BaseMinigoListener) ExitFuncDecl(ctx *FuncDeclContext) {}

// EnterFuncDef is called when production funcDef is entered.
func (s *BaseMinigoListener) EnterFuncDef(ctx *FuncDefContext) {}

// ExitFuncDef is called when production funcDef is exited.
func (s *BaseMinigoListener) ExitFuncDef(ctx *FuncDefContext) {}

// EnterFuncFrontDecl is called when production funcFrontDecl is entered.
func (s *BaseMinigoListener) EnterFuncFrontDecl(ctx *FuncFrontDeclContext) {}

// ExitFuncFrontDecl is called when production funcFrontDecl is exited.
func (s *BaseMinigoListener) ExitFuncFrontDecl(ctx *FuncFrontDeclContext) {}

// EnterFuncArgsDecls is called when production funcArgsDecls is entered.
func (s *BaseMinigoListener) EnterFuncArgsDecls(ctx *FuncArgsDeclsContext) {}

// ExitFuncArgsDecls is called when production funcArgsDecls is exited.
func (s *BaseMinigoListener) ExitFuncArgsDecls(ctx *FuncArgsDeclsContext) {}

// EnterNestedType is called when production nestedType is entered.
func (s *BaseMinigoListener) EnterNestedType(ctx *NestedTypeContext) {}

// ExitNestedType is called when production nestedType is exited.
func (s *BaseMinigoListener) ExitNestedType(ctx *NestedTypeContext) {}

// EnterIdentifierDeclType is called when production identifierDeclType is entered.
func (s *BaseMinigoListener) EnterIdentifierDeclType(ctx *IdentifierDeclTypeContext) {}

// ExitIdentifierDeclType is called when production identifierDeclType is exited.
func (s *BaseMinigoListener) ExitIdentifierDeclType(ctx *IdentifierDeclTypeContext) {}

// EnterSliceType is called when production sliceType is entered.
func (s *BaseMinigoListener) EnterSliceType(ctx *SliceTypeContext) {}

// ExitSliceType is called when production sliceType is exited.
func (s *BaseMinigoListener) ExitSliceType(ctx *SliceTypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseMinigoListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseMinigoListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterStructType is called when production structType is entered.
func (s *BaseMinigoListener) EnterStructType(ctx *StructTypeContext) {}

// ExitStructType is called when production structType is exited.
func (s *BaseMinigoListener) ExitStructType(ctx *StructTypeContext) {}

// EnterSliceDeclType is called when production sliceDeclType is entered.
func (s *BaseMinigoListener) EnterSliceDeclType(ctx *SliceDeclTypeContext) {}

// ExitSliceDeclType is called when production sliceDeclType is exited.
func (s *BaseMinigoListener) ExitSliceDeclType(ctx *SliceDeclTypeContext) {}

// EnterArrayDeclType is called when production arrayDeclType is entered.
func (s *BaseMinigoListener) EnterArrayDeclType(ctx *ArrayDeclTypeContext) {}

// ExitArrayDeclType is called when production arrayDeclType is exited.
func (s *BaseMinigoListener) ExitArrayDeclType(ctx *ArrayDeclTypeContext) {}

// EnterStructDeclType is called when production structDeclType is entered.
func (s *BaseMinigoListener) EnterStructDeclType(ctx *StructDeclTypeContext) {}

// ExitStructDeclType is called when production structDeclType is exited.
func (s *BaseMinigoListener) ExitStructDeclType(ctx *StructDeclTypeContext) {}

// EnterStructMemDecls is called when production structMemDecls is entered.
func (s *BaseMinigoListener) EnterStructMemDecls(ctx *StructMemDeclsContext) {}

// ExitStructMemDecls is called when production structMemDecls is exited.
func (s *BaseMinigoListener) ExitStructMemDecls(ctx *StructMemDeclsContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseMinigoListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseMinigoListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterNegativeExpression is called when production negativeExpression is entered.
func (s *BaseMinigoListener) EnterNegativeExpression(ctx *NegativeExpressionContext) {}

// ExitNegativeExpression is called when production negativeExpression is exited.
func (s *BaseMinigoListener) ExitNegativeExpression(ctx *NegativeExpressionContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseMinigoListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseMinigoListener) ExitComparison(ctx *ComparisonContext) {}

// EnterOperationPrecedence1 is called when production operationPrecedence1 is entered.
func (s *BaseMinigoListener) EnterOperationPrecedence1(ctx *OperationPrecedence1Context) {}

// ExitOperationPrecedence1 is called when production operationPrecedence1 is exited.
func (s *BaseMinigoListener) ExitOperationPrecedence1(ctx *OperationPrecedence1Context) {}

// EnterExpressionPrimaryExpression is called when production expressionPrimaryExpression is entered.
func (s *BaseMinigoListener) EnterExpressionPrimaryExpression(ctx *ExpressionPrimaryExpressionContext) {
}

// ExitExpressionPrimaryExpression is called when production expressionPrimaryExpression is exited.
func (s *BaseMinigoListener) ExitExpressionPrimaryExpression(ctx *ExpressionPrimaryExpressionContext) {
}

// EnterOperationPrecedence2 is called when production operationPrecedence2 is entered.
func (s *BaseMinigoListener) EnterOperationPrecedence2(ctx *OperationPrecedence2Context) {}

// ExitOperationPrecedence2 is called when production operationPrecedence2 is exited.
func (s *BaseMinigoListener) ExitOperationPrecedence2(ctx *OperationPrecedence2Context) {}

// EnterPositiveExpression is called when production positiveExpression is entered.
func (s *BaseMinigoListener) EnterPositiveExpression(ctx *PositiveExpressionContext) {}

// ExitPositiveExpression is called when production positiveExpression is exited.
func (s *BaseMinigoListener) ExitPositiveExpression(ctx *PositiveExpressionContext) {}

// EnterNotExpression is called when production notExpression is entered.
func (s *BaseMinigoListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production notExpression is exited.
func (s *BaseMinigoListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterCaretExpression is called when production caretExpression is entered.
func (s *BaseMinigoListener) EnterCaretExpression(ctx *CaretExpressionContext) {}

// ExitCaretExpression is called when production caretExpression is exited.
func (s *BaseMinigoListener) ExitCaretExpression(ctx *CaretExpressionContext) {}

// EnterBooleanOperation is called when production booleanOperation is entered.
func (s *BaseMinigoListener) EnterBooleanOperation(ctx *BooleanOperationContext) {}

// ExitBooleanOperation is called when production booleanOperation is exited.
func (s *BaseMinigoListener) ExitBooleanOperation(ctx *BooleanOperationContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseMinigoListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseMinigoListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterSubIndex is called when production subIndex is entered.
func (s *BaseMinigoListener) EnterSubIndex(ctx *SubIndexContext) {}

// ExitSubIndex is called when production subIndex is exited.
func (s *BaseMinigoListener) ExitSubIndex(ctx *SubIndexContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseMinigoListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseMinigoListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterCapCall is called when production capCall is entered.
func (s *BaseMinigoListener) EnterCapCall(ctx *CapCallContext) {}

// ExitCapCall is called when production capCall is exited.
func (s *BaseMinigoListener) ExitCapCall(ctx *CapCallContext) {}

// EnterOperandExpression is called when production operandExpression is entered.
func (s *BaseMinigoListener) EnterOperandExpression(ctx *OperandExpressionContext) {}

// ExitOperandExpression is called when production operandExpression is exited.
func (s *BaseMinigoListener) ExitOperandExpression(ctx *OperandExpressionContext) {}

// EnterAppendCall is called when production appendCall is entered.
func (s *BaseMinigoListener) EnterAppendCall(ctx *AppendCallContext) {}

// ExitAppendCall is called when production appendCall is exited.
func (s *BaseMinigoListener) ExitAppendCall(ctx *AppendCallContext) {}

// EnterLenCall is called when production lenCall is entered.
func (s *BaseMinigoListener) EnterLenCall(ctx *LenCallContext) {}

// ExitLenCall is called when production lenCall is exited.
func (s *BaseMinigoListener) ExitLenCall(ctx *LenCallContext) {}

// EnterMemberAccessor is called when production memberAccessor is entered.
func (s *BaseMinigoListener) EnterMemberAccessor(ctx *MemberAccessorContext) {}

// ExitMemberAccessor is called when production memberAccessor is exited.
func (s *BaseMinigoListener) ExitMemberAccessor(ctx *MemberAccessorContext) {}

// EnterLiteralOperand is called when production literalOperand is entered.
func (s *BaseMinigoListener) EnterLiteralOperand(ctx *LiteralOperandContext) {}

// ExitLiteralOperand is called when production literalOperand is exited.
func (s *BaseMinigoListener) ExitLiteralOperand(ctx *LiteralOperandContext) {}

// EnterIdentifierOperand is called when production identifierOperand is entered.
func (s *BaseMinigoListener) EnterIdentifierOperand(ctx *IdentifierOperandContext) {}

// ExitIdentifierOperand is called when production identifierOperand is exited.
func (s *BaseMinigoListener) ExitIdentifierOperand(ctx *IdentifierOperandContext) {}

// EnterExpressionOperand is called when production expressionOperand is entered.
func (s *BaseMinigoListener) EnterExpressionOperand(ctx *ExpressionOperandContext) {}

// ExitExpressionOperand is called when production expressionOperand is exited.
func (s *BaseMinigoListener) ExitExpressionOperand(ctx *ExpressionOperandContext) {}

// EnterIntLiteral is called when production intLiteral is entered.
func (s *BaseMinigoListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production intLiteral is exited.
func (s *BaseMinigoListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseMinigoListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseMinigoListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterRuneLiteral is called when production runeLiteral is entered.
func (s *BaseMinigoListener) EnterRuneLiteral(ctx *RuneLiteralContext) {}

// ExitRuneLiteral is called when production runeLiteral is exited.
func (s *BaseMinigoListener) ExitRuneLiteral(ctx *RuneLiteralContext) {}

// EnterRawStringLiteral is called when production rawStringLiteral is entered.
func (s *BaseMinigoListener) EnterRawStringLiteral(ctx *RawStringLiteralContext) {}

// ExitRawStringLiteral is called when production rawStringLiteral is exited.
func (s *BaseMinigoListener) ExitRawStringLiteral(ctx *RawStringLiteralContext) {}

// EnterInterpretedStringLiteral is called when production interpretedStringLiteral is entered.
func (s *BaseMinigoListener) EnterInterpretedStringLiteral(ctx *InterpretedStringLiteralContext) {}

// ExitInterpretedStringLiteral is called when production interpretedStringLiteral is exited.
func (s *BaseMinigoListener) ExitInterpretedStringLiteral(ctx *InterpretedStringLiteralContext) {}

// EnterNumericIntLiteral is called when production numericIntLiteral is entered.
func (s *BaseMinigoListener) EnterNumericIntLiteral(ctx *NumericIntLiteralContext) {}

// ExitNumericIntLiteral is called when production numericIntLiteral is exited.
func (s *BaseMinigoListener) ExitNumericIntLiteral(ctx *NumericIntLiteralContext) {}

// EnterNumerixHexLiteral is called when production numerixHexLiteral is entered.
func (s *BaseMinigoListener) EnterNumerixHexLiteral(ctx *NumerixHexLiteralContext) {}

// ExitNumerixHexLiteral is called when production numerixHexLiteral is exited.
func (s *BaseMinigoListener) ExitNumerixHexLiteral(ctx *NumerixHexLiteralContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseMinigoListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseMinigoListener) ExitIndex(ctx *IndexContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseMinigoListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseMinigoListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseMinigoListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseMinigoListener) ExitSelector(ctx *SelectorContext) {}

// EnterAppendExpression is called when production appendExpression is entered.
func (s *BaseMinigoListener) EnterAppendExpression(ctx *AppendExpressionContext) {}

// ExitAppendExpression is called when production appendExpression is exited.
func (s *BaseMinigoListener) ExitAppendExpression(ctx *AppendExpressionContext) {}

// EnterLengthExpression is called when production lengthExpression is entered.
func (s *BaseMinigoListener) EnterLengthExpression(ctx *LengthExpressionContext) {}

// ExitLengthExpression is called when production lengthExpression is exited.
func (s *BaseMinigoListener) ExitLengthExpression(ctx *LengthExpressionContext) {}

// EnterCapExpression is called when production capExpression is entered.
func (s *BaseMinigoListener) EnterCapExpression(ctx *CapExpressionContext) {}

// ExitCapExpression is called when production capExpression is exited.
func (s *BaseMinigoListener) ExitCapExpression(ctx *CapExpressionContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseMinigoListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseMinigoListener) ExitStatementList(ctx *StatementListContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMinigoListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMinigoListener) ExitBlock(ctx *BlockContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseMinigoListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseMinigoListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterPrintlnStatement is called when production printlnStatement is entered.
func (s *BaseMinigoListener) EnterPrintlnStatement(ctx *PrintlnStatementContext) {}

// ExitPrintlnStatement is called when production printlnStatement is exited.
func (s *BaseMinigoListener) ExitPrintlnStatement(ctx *PrintlnStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseMinigoListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseMinigoListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseMinigoListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseMinigoListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseMinigoListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseMinigoListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterSimpleStatementStatement is called when production simpleStatementStatement is entered.
func (s *BaseMinigoListener) EnterSimpleStatementStatement(ctx *SimpleStatementStatementContext) {}

// ExitSimpleStatementStatement is called when production simpleStatementStatement is exited.
func (s *BaseMinigoListener) ExitSimpleStatementStatement(ctx *SimpleStatementStatementContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseMinigoListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseMinigoListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseMinigoListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseMinigoListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterIfStatementStatement is called when production ifStatementStatement is entered.
func (s *BaseMinigoListener) EnterIfStatementStatement(ctx *IfStatementStatementContext) {}

// ExitIfStatementStatement is called when production ifStatementStatement is exited.
func (s *BaseMinigoListener) ExitIfStatementStatement(ctx *IfStatementStatementContext) {}

// EnterLoopStatement is called when production loopStatement is entered.
func (s *BaseMinigoListener) EnterLoopStatement(ctx *LoopStatementContext) {}

// ExitLoopStatement is called when production loopStatement is exited.
func (s *BaseMinigoListener) ExitLoopStatement(ctx *LoopStatementContext) {}

// EnterTypeDeclStatement is called when production typeDeclStatement is entered.
func (s *BaseMinigoListener) EnterTypeDeclStatement(ctx *TypeDeclStatementContext) {}

// ExitTypeDeclStatement is called when production typeDeclStatement is exited.
func (s *BaseMinigoListener) ExitTypeDeclStatement(ctx *TypeDeclStatementContext) {}

// EnterVariableDeclStatement is called when production variableDeclStatement is entered.
func (s *BaseMinigoListener) EnterVariableDeclStatement(ctx *VariableDeclStatementContext) {}

// ExitVariableDeclStatement is called when production variableDeclStatement is exited.
func (s *BaseMinigoListener) ExitVariableDeclStatement(ctx *VariableDeclStatementContext) {}

// EnterExpressionSimpleStatement is called when production expressionSimpleStatement is entered.
func (s *BaseMinigoListener) EnterExpressionSimpleStatement(ctx *ExpressionSimpleStatementContext) {}

// ExitExpressionSimpleStatement is called when production expressionSimpleStatement is exited.
func (s *BaseMinigoListener) ExitExpressionSimpleStatement(ctx *ExpressionSimpleStatementContext) {}

// EnterExpressionPostInc is called when production expressionPostInc is entered.
func (s *BaseMinigoListener) EnterExpressionPostInc(ctx *ExpressionPostIncContext) {}

// ExitExpressionPostInc is called when production expressionPostInc is exited.
func (s *BaseMinigoListener) ExitExpressionPostInc(ctx *ExpressionPostIncContext) {}

// EnterExpressionPostDec is called when production expressionPostDec is entered.
func (s *BaseMinigoListener) EnterExpressionPostDec(ctx *ExpressionPostDecContext) {}

// ExitExpressionPostDec is called when production expressionPostDec is exited.
func (s *BaseMinigoListener) ExitExpressionPostDec(ctx *ExpressionPostDecContext) {}

// EnterAssignmentSimpleStatement is called when production assignmentSimpleStatement is entered.
func (s *BaseMinigoListener) EnterAssignmentSimpleStatement(ctx *AssignmentSimpleStatementContext) {}

// ExitAssignmentSimpleStatement is called when production assignmentSimpleStatement is exited.
func (s *BaseMinigoListener) ExitAssignmentSimpleStatement(ctx *AssignmentSimpleStatementContext) {}

// EnterWalrusDeclaration is called when production walrusDeclaration is entered.
func (s *BaseMinigoListener) EnterWalrusDeclaration(ctx *WalrusDeclarationContext) {}

// ExitWalrusDeclaration is called when production walrusDeclaration is exited.
func (s *BaseMinigoListener) ExitWalrusDeclaration(ctx *WalrusDeclarationContext) {}

// EnterNormalAssignment is called when production normalAssignment is entered.
func (s *BaseMinigoListener) EnterNormalAssignment(ctx *NormalAssignmentContext) {}

// ExitNormalAssignment is called when production normalAssignment is exited.
func (s *BaseMinigoListener) ExitNormalAssignment(ctx *NormalAssignmentContext) {}

// EnterInPlaceAssignment is called when production inPlaceAssignment is entered.
func (s *BaseMinigoListener) EnterInPlaceAssignment(ctx *InPlaceAssignmentContext) {}

// ExitInPlaceAssignment is called when production inPlaceAssignment is exited.
func (s *BaseMinigoListener) ExitInPlaceAssignment(ctx *InPlaceAssignmentContext) {}

// EnterIfSingleExpression is called when production ifSingleExpression is entered.
func (s *BaseMinigoListener) EnterIfSingleExpression(ctx *IfSingleExpressionContext) {}

// ExitIfSingleExpression is called when production ifSingleExpression is exited.
func (s *BaseMinigoListener) ExitIfSingleExpression(ctx *IfSingleExpressionContext) {}

// EnterIfElseIf is called when production ifElseIf is entered.
func (s *BaseMinigoListener) EnterIfElseIf(ctx *IfElseIfContext) {}

// ExitIfElseIf is called when production ifElseIf is exited.
func (s *BaseMinigoListener) ExitIfElseIf(ctx *IfElseIfContext) {}

// EnterIfElseBlock is called when production ifElseBlock is entered.
func (s *BaseMinigoListener) EnterIfElseBlock(ctx *IfElseBlockContext) {}

// ExitIfElseBlock is called when production ifElseBlock is exited.
func (s *BaseMinigoListener) ExitIfElseBlock(ctx *IfElseBlockContext) {}

// EnterIfSimpleNoElse is called when production ifSimpleNoElse is entered.
func (s *BaseMinigoListener) EnterIfSimpleNoElse(ctx *IfSimpleNoElseContext) {}

// ExitIfSimpleNoElse is called when production ifSimpleNoElse is exited.
func (s *BaseMinigoListener) ExitIfSimpleNoElse(ctx *IfSimpleNoElseContext) {}

// EnterIfSimpleElseIf is called when production ifSimpleElseIf is entered.
func (s *BaseMinigoListener) EnterIfSimpleElseIf(ctx *IfSimpleElseIfContext) {}

// ExitIfSimpleElseIf is called when production ifSimpleElseIf is exited.
func (s *BaseMinigoListener) ExitIfSimpleElseIf(ctx *IfSimpleElseIfContext) {}

// EnterIfSimpleElseBlock is called when production ifSimpleElseBlock is entered.
func (s *BaseMinigoListener) EnterIfSimpleElseBlock(ctx *IfSimpleElseBlockContext) {}

// ExitIfSimpleElseBlock is called when production ifSimpleElseBlock is exited.
func (s *BaseMinigoListener) ExitIfSimpleElseBlock(ctx *IfSimpleElseBlockContext) {}

// EnterInfiniteFor is called when production infiniteFor is entered.
func (s *BaseMinigoListener) EnterInfiniteFor(ctx *InfiniteForContext) {}

// ExitInfiniteFor is called when production infiniteFor is exited.
func (s *BaseMinigoListener) ExitInfiniteFor(ctx *InfiniteForContext) {}

// EnterWhileFor is called when production whileFor is entered.
func (s *BaseMinigoListener) EnterWhileFor(ctx *WhileForContext) {}

// ExitWhileFor is called when production whileFor is exited.
func (s *BaseMinigoListener) ExitWhileFor(ctx *WhileForContext) {}

// EnterThreePartFor is called when production threePartFor is entered.
func (s *BaseMinigoListener) EnterThreePartFor(ctx *ThreePartForContext) {}

// ExitThreePartFor is called when production threePartFor is exited.
func (s *BaseMinigoListener) ExitThreePartFor(ctx *ThreePartForContext) {}

// EnterThreePartForNoExpression is called when production threePartForNoExpression is entered.
func (s *BaseMinigoListener) EnterThreePartForNoExpression(ctx *ThreePartForNoExpressionContext) {}

// ExitThreePartForNoExpression is called when production threePartForNoExpression is exited.
func (s *BaseMinigoListener) ExitThreePartForNoExpression(ctx *ThreePartForNoExpressionContext) {}

// EnterSimpleStatementSwitchExpression is called when production simpleStatementSwitchExpression is entered.
func (s *BaseMinigoListener) EnterSimpleStatementSwitchExpression(ctx *SimpleStatementSwitchExpressionContext) {
}

// ExitSimpleStatementSwitchExpression is called when production simpleStatementSwitchExpression is exited.
func (s *BaseMinigoListener) ExitSimpleStatementSwitchExpression(ctx *SimpleStatementSwitchExpressionContext) {
}

// EnterNormalSwitch is called when production normalSwitch is entered.
func (s *BaseMinigoListener) EnterNormalSwitch(ctx *NormalSwitchContext) {}

// ExitNormalSwitch is called when production normalSwitch is exited.
func (s *BaseMinigoListener) ExitNormalSwitch(ctx *NormalSwitchContext) {}

// EnterNormalSwitchExpression is called when production normalSwitchExpression is entered.
func (s *BaseMinigoListener) EnterNormalSwitchExpression(ctx *NormalSwitchExpressionContext) {}

// ExitNormalSwitchExpression is called when production normalSwitchExpression is exited.
func (s *BaseMinigoListener) ExitNormalSwitchExpression(ctx *NormalSwitchExpressionContext) {}

// EnterSimpleStatementSwitch is called when production simpleStatementSwitch is entered.
func (s *BaseMinigoListener) EnterSimpleStatementSwitch(ctx *SimpleStatementSwitchContext) {}

// ExitSimpleStatementSwitch is called when production simpleStatementSwitch is exited.
func (s *BaseMinigoListener) ExitSimpleStatementSwitch(ctx *SimpleStatementSwitchContext) {}

// EnterExpressionCaseClauseList is called when production expressionCaseClauseList is entered.
func (s *BaseMinigoListener) EnterExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) {}

// ExitExpressionCaseClauseList is called when production expressionCaseClauseList is exited.
func (s *BaseMinigoListener) ExitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) {}

// EnterExpressionCaseClause is called when production expressionCaseClause is entered.
func (s *BaseMinigoListener) EnterExpressionCaseClause(ctx *ExpressionCaseClauseContext) {}

// ExitExpressionCaseClause is called when production expressionCaseClause is exited.
func (s *BaseMinigoListener) ExitExpressionCaseClause(ctx *ExpressionCaseClauseContext) {}

// EnterSwitchCaseBranch is called when production switchCaseBranch is entered.
func (s *BaseMinigoListener) EnterSwitchCaseBranch(ctx *SwitchCaseBranchContext) {}

// ExitSwitchCaseBranch is called when production switchCaseBranch is exited.
func (s *BaseMinigoListener) ExitSwitchCaseBranch(ctx *SwitchCaseBranchContext) {}

// EnterSwitchDefaultBranch is called when production switchDefaultBranch is entered.
func (s *BaseMinigoListener) EnterSwitchDefaultBranch(ctx *SwitchDefaultBranchContext) {}

// ExitSwitchDefaultBranch is called when production switchDefaultBranch is exited.
func (s *BaseMinigoListener) ExitSwitchDefaultBranch(ctx *SwitchDefaultBranchContext) {}
