package checker

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/zSnails/minigo/grammar"
	symboltable "github.com/zSnails/minigo/symbol-table"
)

type TypeChecker struct {
	Result      int
	filename    string
	SymbolTable *symboltable.SymbolTable
	errors      []error
}

// VisitAppendCall implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitAppendCall(ctx *grammar.AppendCallContext) interface{} {
	panic("unimplemented")
}

// VisitCapCall implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitCapCall(ctx *grammar.CapCallContext) interface{} {
	panic("unimplemented")
}

// VisitFunctionCall implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFunctionCall(ctx *grammar.FunctionCallContext) interface{} {
    panic("unimplemented")
}

// VisitLenCall implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitLenCall(ctx *grammar.LenCallContext) interface{} {
	panic("unimplemented")
}

// VisitMemberAccessor implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitMemberAccessor(ctx *grammar.MemberAccessorContext) interface{} {
	panic("unimplemented")
}

// VisitOperandExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitOperandExpression(ctx *grammar.OperandExpressionContext) interface{} {
	if operand := ctx.Operand(); operand != nil {
		return t.Visit(operand)
	}
    return t.VisitChildren(ctx)
}

// VisitSubIndex implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSubIndex(ctx *grammar.SubIndexContext) interface{} {
	panic("unimplemented")
}

var _ grammar.MinigoVisitor = &TypeChecker{}

// VisitInPlaceAssignment implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitInPlaceAssignment(ctx *grammar.InPlaceAssignmentContext) interface{} {
	lhs := ctx.GetLeft()
	rhs := ctx.GetRight()
	symbol, found := t.SymbolTable.Symbols.FindFirst(func(s *symboltable.Symbol) bool {
		return s.Name == lhs.GetText()
	})

	if !found {
		t.errors = append(t.errors, t.MakeError(ctx.GetStart(), fmt.Errorf("undefined: %s", lhs.GetText())))
		return nil
	}
	right, ok := t.Visit(rhs).(*symboltable.Symbol)
	if !ok {
		panic("unimplemented")
	}

	if symbol.Type != right {
		t.errors = append(t.errors, t.MakeError(ctx.GetStart(), fmt.Errorf("cannot use %s as %s value in assignment", right.Name, symbol.Type.Name)))
	}

	return nil
}

// VisitNormalAssignment implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitNormalAssignment(ctx *grammar.NormalAssignmentContext) interface{} {
	rhs := ctx.GetRight().AllExpression()
	for idx, ident := range ctx.GetLeft().AllExpression() {
		expression := rhs[idx]
		symbol, found := t.SymbolTable.Symbols.FindFirst(func(s *symboltable.Symbol) bool {
			return s.Name == ident.GetText()
		})

		if !found {
			t.errors = append(t.errors, t.MakeError(ctx.GetStart(), fmt.Errorf("undefined: %s", ident.GetText())))
			return nil
		}

		right, ok := t.Visit(expression).(*symboltable.Symbol)
		if !ok {
			panic("unimplemented")
		}

		if symbol.Type != right {
			t.errors = append(t.errors, t.MakeError(ctx.GetStart(), fmt.Errorf("cannot use %s as %s value in assignment", right.Name, symbol.Type.Name)))
		}
	}

	return nil
}

// VisitCaretExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitCaretExpression(ctx *grammar.CaretExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitExpressionPrimaryExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionPrimaryExpression(ctx *grammar.ExpressionPrimaryExpressionContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitMinusExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitMinusExpression(ctx *grammar.MinusExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitNotExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitNotExpression(ctx *grammar.NotExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitOperationExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitOperationExpression(ctx *grammar.OperationExpressionContext) interface{} {
	// TODO: check for possible edge cases on this piece of shit
	leftType, leftOk := t.Visit(ctx.GetLeft()).(*symboltable.Symbol)
	rightType, rightOk := t.Visit(ctx.GetRight()).(*symboltable.Symbol)
	if (leftType != rightType) && (leftOk && rightOk) {
		t.errors = append(t.errors, t.MakeError(ctx.GetStart(), fmt.Errorf("mismatched types: %s and %s", leftType.Name, rightType.Name)))
	}
	if leftType == nil {
		return rightOk
	}
	return leftType
}

// VisitPlusExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitPlusExpression(ctx *grammar.PlusExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitFloatLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFloatLiteral(ctx *grammar.FloatLiteralContext) interface{} {
	return symboltable.Float64
}

// VisitIntLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIntLiteral(ctx *grammar.IntLiteralContext) interface{} {
	return symboltable.Int
}

// VisitInterpretedStringLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitInterpretedStringLiteral(ctx *grammar.InterpretedStringLiteralContext) interface{} {
	return symboltable.String
}

// VisitRawStringLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitRawStringLiteral(ctx *grammar.RawStringLiteralContext) interface{} {
	return symboltable.String
}

// VisitRuneLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitRuneLiteral(ctx *grammar.RuneLiteralContext) interface{} {
	return symboltable.Rune
}

// VisitRoot implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitRoot(ctx *grammar.RootContext) interface{} {
	return t.VisitChildren(ctx)
}

// Visit implements grammar.MinigoVisitor.
func (t *TypeChecker) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(t)
}

// VisitEmptyTypeDeclaration implements grammar.MinigoVisitor.
// Subtle: this method shadows the method (BaseMinigoVisitor).VisitEmptyTypeDeclaration of TypeChecker.BaseMinigoVisitor.
func (t *TypeChecker) VisitEmptyTypeDeclaration(ctx *grammar.EmptyTypeDeclarationContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitMultiTypeDeclaration implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitMultiTypeDeclaration(ctx *grammar.MultiTypeDeclarationContext) interface{} {
	return t.VisitChildren(ctx)
}

func (t *TypeChecker) MakeError(token antlr.Token, err error) error {
	return fmt.Errorf("%s:%d:%d: %w", t.filename, token.GetLine(), token.GetTokenSource().GetCharPositionInLine(), err)
}

// VisitSingleVarDeclsNoExpsDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSingleVarDeclsNoExpsDecl(ctx *grammar.SingleVarDeclsNoExpsDeclContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitTypeDeclaration implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitTypeDeclaration(ctx *grammar.TypeDeclarationContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitTypedVarDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitTypedVarDecl(ctx *grammar.TypedVarDeclContext) interface{} {
	panic("unimplemented")
}

// VisitUntypedVarDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitUntypedVarDecl(ctx *grammar.UntypedVarDeclContext) interface{} {
	identList := ctx.IdentifierList().AllIDENTIFIER()
	exprList := ctx.ExpressionList().AllExpression()
	if idenLen, exprLen := len(identList), len(exprList); idenLen != exprLen {
		t.errors = append(t.errors, t.MakeError(ctx.GetStart(), fmt.Errorf("assignment mismatch: %d variables but %d vaules", idenLen, exprLen)))
		return nil // Unrecoverable
	}
	for idx, ident := range identList {
		expr := exprList[idx]
		_type, ok := t.Visit(expr).(*symboltable.Symbol)
		if !ok {
			t.errors = append(t.errors, t.MakeError(ctx.GetStart(), fmt.Errorf("invalid type")))
		}
		symbol := t.SymbolTable.NewVariable(ident.GetText(), _type)
		err := t.SymbolTable.AddSymbol(symbol)
		if err != nil {
			t.errors = append(t.errors, t.MakeError(ident.GetSymbol(), err))
		}
	}

	return nil
}

// VisitAppendExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitAppendExpression(ctx *grammar.AppendExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitArguments implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitArguments(ctx *grammar.ArgumentsContext) interface{} {
	panic("unimplemented")
}

// VisitArrayDeclType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitArrayDeclType(ctx *grammar.ArrayDeclTypeContext) interface{} {
	panic("unimplemented")
}

// VisitBlock implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitBlock(ctx *grammar.BlockContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitBlockStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitBlockStatement(ctx *grammar.BlockStatementContext) interface{} {
	panic("unimplemented")
}

// VisitBreakStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitBreakStatement(ctx *grammar.BreakStatementContext) interface{} {
	panic("unimplemented")
}

// VisitCapExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitCapExpression(ctx *grammar.CapExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitContinueStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitContinueStatement(ctx *grammar.ContinueStatementContext) interface{} {
	panic("unimplemented")
}

// VisitDeclType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitDeclType(ctx *grammar.DeclTypeContext) interface{} {
	// Get the name
	var name string
	if ident := ctx.IDENTIFIER(); ident != nil {
		name = ident.GetText()
	} else if ident := ctx.SliceDeclType().DeclType().IDENTIFIER(); ident != nil {
		name = ident.GetText()
	} else if ident := ctx.ArrayDeclType().DeclType().IDENTIFIER(); ident != nil {
		name = ident.GetText()
	}

	_type, found := t.SymbolTable.Symbols.FindFirst(func(s *symboltable.Symbol) bool {
		return s.Name == name
	})
	if !found {
		t.errors = append(t.errors, t.MakeError(ctx.GetStart(), fmt.Errorf("typename %s not found", name)))
		return nil
	}
	return _type
}

// VisitEmptyVariableDeclaration implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitEmptyVariableDeclaration(ctx *grammar.EmptyVariableDeclarationContext) interface{} {
	panic("unimplemented")
}

// VisitErrorNode implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}

// VisitExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpression(ctx *grammar.ExpressionContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitExpressionCaseClause implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionCaseClause(ctx *grammar.ExpressionCaseClauseContext) interface{} {
	panic("unimplemented")
}

// VisitExpressionCaseClauseList implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionCaseClauseList(ctx *grammar.ExpressionCaseClauseListContext) interface{} {
	panic("unimplemented")
}

// VisitExpressionList implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionList(ctx *grammar.ExpressionListContext) interface{} {
	panic("unimplemented")
}

// VisitExpressionSwitchCase implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionSwitchCase(ctx *grammar.ExpressionSwitchCaseContext) interface{} {
	panic("unimplemented")
}

// VisitFuncArgsDecls implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFuncArgsDecls(ctx *grammar.FuncArgsDeclsContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitFuncDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFuncDecl(ctx *grammar.FuncDeclContext) interface{} {
	defer t.SymbolTable.ExitScope() // This will exit the scope created in the FuncFrontDecl
	return t.VisitChildren(ctx)
}

// VisitFuncFrontDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFuncFrontDecl(ctx *grammar.FuncFrontDeclContext) interface{} {
	name := ctx.IDENTIFIER()
	var _type *symboltable.Symbol = nil
	if declType := ctx.DeclType(); declType != nil {
		var ok bool
		_type, ok = t.Visit(ctx.DeclType()).(*symboltable.Symbol)
		if !ok {
			_type = nil
		}
	}
	symbol := t.SymbolTable.NewFunction(name.GetText(), _type)
	err := t.SymbolTable.AddSymbol(symbol)
	if err != nil {
		t.errors = append(t.errors, t.MakeError(ctx.GetStart(), err))
	}
	t.SymbolTable.EnterScope() // This will enter the function scope
	if funcArgs := ctx.FuncArgsDecls(); funcArgs != nil {
		t.Visit(funcArgs)
	}
	return nil
}

// VisitIdentifierList implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIdentifierList(ctx *grammar.IdentifierListContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitIfStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIfStatement(ctx *grammar.IfStatementContext) interface{} {
	panic("unimplemented")
}

// VisitIfStatementStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIfStatementStatement(ctx *grammar.IfStatementStatementContext) interface{} {
	panic("unimplemented")
}

// VisitIndex implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIndex(ctx *grammar.IndexContext) interface{} {
	panic("unimplemented")
}

// VisitInnerTypeDecls implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitInnerTypeDecls(ctx *grammar.InnerTypeDeclsContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitInnerVarDecls implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitInnerVarDecls(ctx *grammar.InnerVarDeclsContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitLengthExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitLengthExpression(ctx *grammar.LengthExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitLoop implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitLoop(ctx *grammar.LoopContext) interface{} {
	panic("unimplemented")
}

// VisitLoopStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitLoopStatement(ctx *grammar.LoopStatementContext) interface{} {
	panic("unimplemented")
}

// VisitMultiVariableDeclaration implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitMultiVariableDeclaration(ctx *grammar.MultiVariableDeclarationContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitOperand implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitOperand(ctx *grammar.OperandContext) interface{} {
	if ident := ctx.IDENTIFIER(); ident != nil {
		symbol, found := t.SymbolTable.Symbols.FindFirst(func(s *symboltable.Symbol) bool {
			return ident.GetText() == s.Name
		})
		if !found {
			t.errors = append(t.errors, t.MakeError(ctx.GetStart(), fmt.Errorf("undefined: %s", ident.GetText())))
			return nil
		}
		return symbol.Type
	}
	return t.VisitChildren(ctx)
}

// VisitPrimaryExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitPrimaryExpression(ctx *grammar.PrimaryExpressionContext) interface{} {
	if operand := ctx.Operand(); operand != nil {
		return t.Visit(operand)
	}
	return t.VisitChildren(ctx)
}

// VisitPrintStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitPrintStatement(ctx *grammar.PrintStatementContext) interface{} {
	panic("unimplemented")
}

// VisitPrintlnStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitPrintlnStatement(ctx *grammar.PrintlnStatementContext) interface{} {
	panic("unimplemented")
}

// VisitReturnStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitReturnStatement(ctx *grammar.ReturnStatementContext) interface{} {
	panic("unimplemented")
}

// VisitSelector implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSelector(ctx *grammar.SelectorContext) interface{} {
	panic("unimplemented")
}

// VisitSimpleStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSimpleStatement(ctx *grammar.SimpleStatementContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitSimpleStatementStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSimpleStatementStatement(ctx *grammar.SimpleStatementStatementContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitSingleTypeDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSingleTypeDecl(ctx *grammar.SingleTypeDeclContext) interface{} {
	name := ctx.IDENTIFIER()
	declType := ctx.DeclType()
	if typeName := declType.IDENTIFIER(); typeName != nil {
		_type, err := t.SymbolTable.NewType(name.GetText(), typeName.GetText())
		if err != nil {
			t.errors = append(t.errors, t.MakeError(typeName.GetSymbol(), err))
			return nil // Can't continue from here, it will just not work
		}
		err = t.SymbolTable.AddSymbol(_type)
		if err != nil {
			t.errors = append(t.errors, t.MakeError(name.GetSymbol(), err))
		}
	} else if typeName := declType.SliceDeclType(); typeName != nil {
		// TODO: add slice to the symbol table
	} else if typeName := declType.ArrayDeclType(); typeName != nil {
		// TODO: add array to the symbol table
	} else if typeName := declType.StructDeclType(); typeName != nil {
		// TODO: add struct to the symbol table
	}
	return nil
}

// VisitSingleVarDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSingleVarDecl(ctx *grammar.SingleVarDeclContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitSingleVarDeclNoExps implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSingleVarDeclNoExps(ctx *grammar.SingleVarDeclNoExpsContext) interface{} {
	_type, ok := t.Visit(ctx.DeclType()).(*symboltable.Symbol)
	if !ok {
		return nil // Unrecoverable error
	}
	for _, ident := range ctx.IdentifierList().AllIDENTIFIER() {
		symbol := t.SymbolTable.NewVariable(ident.GetText(), _type)
		err := t.SymbolTable.AddSymbol(symbol)
		if err != nil {
			t.errors = append(t.errors, t.MakeError(ident.GetSymbol(), err))
		}
	}
	return nil
}

// VisitSliceDeclType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSliceDeclType(ctx *grammar.SliceDeclTypeContext) interface{} {
	panic("unimplemented")
}

// VisitStatementList implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitStatementList(ctx *grammar.StatementListContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitStructDeclType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitStructDeclType(ctx *grammar.StructDeclTypeContext) interface{} {
	panic("unimplemented")
}

// VisitStructMemDecls implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitStructMemDecls(ctx *grammar.StructMemDeclsContext) interface{} {
	panic("unimplemented")
}

// VisitSwitch implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSwitch(ctx *grammar.SwitchContext) interface{} {
	panic("unimplemented")
}

// VisitSwitchStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSwitchStatement(ctx *grammar.SwitchStatementContext) interface{} {
	panic("unimplemented")
}

// VisitTerminal implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitTerminal(node antlr.TerminalNode) interface{} {
	// symbol, found := t.SymbolTable.Symbols.FindFirst(func(s *symboltable.Symbol) bool {
	//     return node.GetText() == s.Name
	// })
	// if found {
	//     fmt.Printf("symbol: %v\n", symbol)
	//     return symbol
	// } else {
	//     return nil
	// }
	return nil
}

// VisitTopDeclarationList implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitTopDeclarationList(ctx *grammar.TopDeclarationListContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitTypeDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitTypeDecl(ctx *grammar.TypeDeclContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitTypeDeclStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitTypeDeclStatement(ctx *grammar.TypeDeclStatementContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitVariableDeclStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitVariableDeclStatement(ctx *grammar.VariableDeclStatementContext) interface{} {
	// TODO: use the symbol table and add the symbols declared here to it
	return t.VisitChildren(ctx)
}

// VisitVariableDeclaration implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitVariableDeclaration(ctx *grammar.VariableDeclarationContext) interface{} {
	// TODO: use the symbol table and add the symbols declared here to it
	return t.VisitChildren(ctx)
}

func (v *TypeChecker) VisitChildren(node antlr.RuleNode) interface{} {
	var result any = nil
	for _, child := range node.GetChildren() {
		cpt := child.(antlr.ParseTree)
		childResult := cpt.Accept(v)
		result = childResult
	}
	return result
}

func (t *TypeChecker) HasErrors() bool {
	return len(t.errors) > 0
}

func (t TypeChecker) String() string {
	sb := strings.Builder{}
	for _, err := range t.errors {
		fmt.Fprintf(&sb, "%s\n", err)
	}
	return sb.String()
}

func NewTypeChecker(filename string) *TypeChecker {
	return &TypeChecker{
		filename:    filename,
		Result:      0,
		SymbolTable: symboltable.NewSymbolTable(),
		errors:      []error{},
	}
}
