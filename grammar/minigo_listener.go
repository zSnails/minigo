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

<<<<<<< HEAD
	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterMultiVariableDeclaration is called when entering the multiVariableDeclaration production.
	EnterMultiVariableDeclaration(c *MultiVariableDeclarationContext)

	// EnterEmptyVariableDeclaration is called when entering the emptyVariableDeclaration production.
	EnterEmptyVariableDeclaration(c *EmptyVariableDeclarationContext)
=======
	// EnterVariableDecl is called when entering the variableDecl production.
	EnterVariableDecl(c *VariableDeclContext)
>>>>>>> a9dd69d (Initial Commit)

	// EnterInnerVarDecls is called when entering the innerVarDecls production.
	EnterInnerVarDecls(c *InnerVarDeclsContext)

	// EnterSingleVarDecl is called when entering the singleVarDecl production.
	EnterSingleVarDecl(c *SingleVarDeclContext)

	// EnterSingleVarDeclNoExps is called when entering the singleVarDeclNoExps production.
	EnterSingleVarDeclNoExps(c *SingleVarDeclNoExpsContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterInnerTypeDecls is called when entering the innerTypeDecls production.
	EnterInnerTypeDecls(c *InnerTypeDeclsContext)

	// EnterSingleTypeDecl is called when entering the singleTypeDecl production.
	EnterSingleTypeDecl(c *SingleTypeDeclContext)

	// EnterFuncDecl is called when entering the funcDecl production.
	EnterFuncDecl(c *FuncDeclContext)

	// EnterFuncFrontDecl is called when entering the funcFrontDecl production.
	EnterFuncFrontDecl(c *FuncFrontDeclContext)

	// EnterFuncArgsDecls is called when entering the funcArgsDecls production.
	EnterFuncArgsDecls(c *FuncArgsDeclsContext)

	// EnterDeclType is called when entering the declType production.
	EnterDeclType(c *DeclTypeContext)

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

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

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

<<<<<<< HEAD
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
=======
	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)
>>>>>>> a9dd69d (Initial Commit)

	// EnterSimpleStatement is called when entering the simpleStatement production.
	EnterSimpleStatement(c *SimpleStatementContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterSwitch is called when entering the switch production.
	EnterSwitch(c *SwitchContext)

	// EnterExpressionCaseClauseList is called when entering the expressionCaseClauseList production.
	EnterExpressionCaseClauseList(c *ExpressionCaseClauseListContext)

	// EnterExpressionCaseClause is called when entering the expressionCaseClause production.
	EnterExpressionCaseClause(c *ExpressionCaseClauseContext)

	// EnterExpressionSwitchCase is called when entering the expressionSwitchCase production.
	EnterExpressionSwitchCase(c *ExpressionSwitchCaseContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitTopDeclarationList is called when exiting the topDeclarationList production.
	ExitTopDeclarationList(c *TopDeclarationListContext)

<<<<<<< HEAD
	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitMultiVariableDeclaration is called when exiting the multiVariableDeclaration production.
	ExitMultiVariableDeclaration(c *MultiVariableDeclarationContext)

	// ExitEmptyVariableDeclaration is called when exiting the emptyVariableDeclaration production.
	ExitEmptyVariableDeclaration(c *EmptyVariableDeclarationContext)
=======
	// ExitVariableDecl is called when exiting the variableDecl production.
	ExitVariableDecl(c *VariableDeclContext)
>>>>>>> a9dd69d (Initial Commit)

	// ExitInnerVarDecls is called when exiting the innerVarDecls production.
	ExitInnerVarDecls(c *InnerVarDeclsContext)

	// ExitSingleVarDecl is called when exiting the singleVarDecl production.
	ExitSingleVarDecl(c *SingleVarDeclContext)

	// ExitSingleVarDeclNoExps is called when exiting the singleVarDeclNoExps production.
	ExitSingleVarDeclNoExps(c *SingleVarDeclNoExpsContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitInnerTypeDecls is called when exiting the innerTypeDecls production.
	ExitInnerTypeDecls(c *InnerTypeDeclsContext)

	// ExitSingleTypeDecl is called when exiting the singleTypeDecl production.
	ExitSingleTypeDecl(c *SingleTypeDeclContext)

	// ExitFuncDecl is called when exiting the funcDecl production.
	ExitFuncDecl(c *FuncDeclContext)

	// ExitFuncFrontDecl is called when exiting the funcFrontDecl production.
	ExitFuncFrontDecl(c *FuncFrontDeclContext)

	// ExitFuncArgsDecls is called when exiting the funcArgsDecls production.
	ExitFuncArgsDecls(c *FuncArgsDeclsContext)

	// ExitDeclType is called when exiting the declType production.
	ExitDeclType(c *DeclTypeContext)

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

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

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

<<<<<<< HEAD
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
=======
	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)
>>>>>>> a9dd69d (Initial Commit)

	// ExitSimpleStatement is called when exiting the simpleStatement production.
	ExitSimpleStatement(c *SimpleStatementContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitSwitch is called when exiting the switch production.
	ExitSwitch(c *SwitchContext)

	// ExitExpressionCaseClauseList is called when exiting the expressionCaseClauseList production.
	ExitExpressionCaseClauseList(c *ExpressionCaseClauseListContext)

	// ExitExpressionCaseClause is called when exiting the expressionCaseClause production.
	ExitExpressionCaseClause(c *ExpressionCaseClauseContext)

	// ExitExpressionSwitchCase is called when exiting the expressionSwitchCase production.
	ExitExpressionSwitchCase(c *ExpressionSwitchCaseContext)
}
