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

<<<<<<< HEAD
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
=======
// EnterVariableDecl is called when production variableDecl is entered.
func (s *BaseMinigoListener) EnterVariableDecl(ctx *VariableDeclContext) {}

// ExitVariableDecl is called when production variableDecl is exited.
func (s *BaseMinigoListener) ExitVariableDecl(ctx *VariableDeclContext) {}
>>>>>>> a9dd69d (Initial Commit)

// EnterInnerVarDecls is called when production innerVarDecls is entered.
func (s *BaseMinigoListener) EnterInnerVarDecls(ctx *InnerVarDeclsContext) {}

// ExitInnerVarDecls is called when production innerVarDecls is exited.
func (s *BaseMinigoListener) ExitInnerVarDecls(ctx *InnerVarDeclsContext) {}

// EnterSingleVarDecl is called when production singleVarDecl is entered.
func (s *BaseMinigoListener) EnterSingleVarDecl(ctx *SingleVarDeclContext) {}

// ExitSingleVarDecl is called when production singleVarDecl is exited.
func (s *BaseMinigoListener) ExitSingleVarDecl(ctx *SingleVarDeclContext) {}

// EnterSingleVarDeclNoExps is called when production singleVarDeclNoExps is entered.
func (s *BaseMinigoListener) EnterSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) {}

// ExitSingleVarDeclNoExps is called when production singleVarDeclNoExps is exited.
func (s *BaseMinigoListener) ExitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseMinigoListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseMinigoListener) ExitTypeDecl(ctx *TypeDeclContext) {}

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

// EnterFuncFrontDecl is called when production funcFrontDecl is entered.
func (s *BaseMinigoListener) EnterFuncFrontDecl(ctx *FuncFrontDeclContext) {}

// ExitFuncFrontDecl is called when production funcFrontDecl is exited.
func (s *BaseMinigoListener) ExitFuncFrontDecl(ctx *FuncFrontDeclContext) {}

// EnterFuncArgsDecls is called when production funcArgsDecls is entered.
func (s *BaseMinigoListener) EnterFuncArgsDecls(ctx *FuncArgsDeclsContext) {}

// ExitFuncArgsDecls is called when production funcArgsDecls is exited.
func (s *BaseMinigoListener) ExitFuncArgsDecls(ctx *FuncArgsDeclsContext) {}

// EnterDeclType is called when production declType is entered.
func (s *BaseMinigoListener) EnterDeclType(ctx *DeclTypeContext) {}

// ExitDeclType is called when production declType is exited.
func (s *BaseMinigoListener) ExitDeclType(ctx *DeclTypeContext) {}

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

// EnterExpression is called when production expression is entered.
func (s *BaseMinigoListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMinigoListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseMinigoListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseMinigoListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseMinigoListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseMinigoListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseMinigoListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseMinigoListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseMinigoListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseMinigoListener) ExitLiteral(ctx *LiteralContext) {}

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

<<<<<<< HEAD
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
=======
// EnterStatement is called when production statement is entered.
func (s *BaseMinigoListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseMinigoListener) ExitStatement(ctx *StatementContext) {}
>>>>>>> a9dd69d (Initial Commit)

// EnterSimpleStatement is called when production simpleStatement is entered.
func (s *BaseMinigoListener) EnterSimpleStatement(ctx *SimpleStatementContext) {}

// ExitSimpleStatement is called when production simpleStatement is exited.
func (s *BaseMinigoListener) ExitSimpleStatement(ctx *SimpleStatementContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseMinigoListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseMinigoListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseMinigoListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseMinigoListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaseMinigoListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaseMinigoListener) ExitLoop(ctx *LoopContext) {}

// EnterSwitch is called when production switch is entered.
func (s *BaseMinigoListener) EnterSwitch(ctx *SwitchContext) {}

// ExitSwitch is called when production switch is exited.
func (s *BaseMinigoListener) ExitSwitch(ctx *SwitchContext) {}

// EnterExpressionCaseClauseList is called when production expressionCaseClauseList is entered.
func (s *BaseMinigoListener) EnterExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) {}

// ExitExpressionCaseClauseList is called when production expressionCaseClauseList is exited.
func (s *BaseMinigoListener) ExitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) {}

// EnterExpressionCaseClause is called when production expressionCaseClause is entered.
func (s *BaseMinigoListener) EnterExpressionCaseClause(ctx *ExpressionCaseClauseContext) {}

// ExitExpressionCaseClause is called when production expressionCaseClause is exited.
func (s *BaseMinigoListener) ExitExpressionCaseClause(ctx *ExpressionCaseClauseContext) {}

// EnterExpressionSwitchCase is called when production expressionSwitchCase is entered.
func (s *BaseMinigoListener) EnterExpressionSwitchCase(ctx *ExpressionSwitchCaseContext) {}

// ExitExpressionSwitchCase is called when production expressionSwitchCase is exited.
func (s *BaseMinigoListener) ExitExpressionSwitchCase(ctx *ExpressionSwitchCaseContext) {}
