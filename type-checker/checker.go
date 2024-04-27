package checker

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"log"

	"github.com/antlr4-go/antlr/v4"
	"github.com/zSnails/minigo/grammar"
	"github.com/zSnails/minigo/stack"
	symboltable "github.com/zSnails/minigo/symbol-table"
)

type TypeChecker struct {
	listener antlr.ErrorListener
	filename string
	// currentFunction *symboltable.Symbol
	symbolStack *stack.Stack[*symboltable.Symbol]
	SymbolTable *symboltable.SymbolTable
}

// VisitArrayDeclType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitArrayDeclType(ctx *grammar.ArrayDeclTypeContext) interface{} {
	_symbol := t.Visit(ctx.DeclType()).(*symboltable.Symbol)
	size, _ := strconv.ParseUint(ctx.INTLITERAL().GetText(), 10, 64)
	return &symboltable.Symbol{
		SymbolType: symboltable.TypeSymbol | symboltable.ArraySymbol,
		Token:      _symbol.Token,
		Size:       size,
		Scope:      t.SymbolTable.Scope,
		Name:       "<array>",
		Type:       _symbol,
		Members:    nil,
	}
}

// VisitSliceDeclType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSliceDeclType(ctx *grammar.SliceDeclTypeContext) interface{} {
	_symbol, ok := t.Visit(ctx.DeclType()).(*symboltable.Symbol)
	if !ok {
		return nil // unrecoverable
	}
	return &symboltable.Symbol{
		SymbolType: symboltable.TypeSymbol | symboltable.SliceSymbol,
		Token:      _symbol.Token,
		Size:       0,
		Scope:      t.SymbolTable.Scope,
		Name:       "<slice>",
		Type:       _symbol,
		Members:    nil,
	}
}

// VisitStructDeclType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitStructDeclType(ctx *grammar.StructDeclTypeContext) interface{} {
	self, _ := t.symbolStack.Peek()
	out := t.SymbolTable.NewStructType(ctx.GetStart(), "<struct>")
	if members := ctx.StructMemDecls(); members != nil {
		decls := members.AllSingleVarDeclNoExps()
		for _, decl := range decls {
			idents := decl.IdentifierList().AllIDENTIFIER()
			declType := decl.DeclType()
			_type, ok := t.Visit(declType).(*symboltable.Symbol)
			fmt.Printf("_type: %v\n", _type)
			if !ok {
				continue
			}
			for _, ident := range idents {
				variable := t.SymbolTable.NewVariable(ident.GetSymbol(), ident.GetText(), _type)
				variable = makeSlice(variable, _type)
				out.Members = append(out.Members, variable)
			}
		}
	}
	return out
}

// VisitArrayType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitArrayType(ctx *grammar.ArrayTypeContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitIdentifierDeclType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIdentifierDeclType(ctx *grammar.IdentifierDeclTypeContext) interface{} {
	name := ctx.IDENTIFIER()
	symbol, found := t.SymbolTable.GetSymbol(name.GetText())
	if !found {
		t.MakeError(name.GetSymbol(), fmt.Errorf("unknown symbol '%s'", name.GetText()))
		return nil
	}
	return symbol
}

// VisitNestedType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitNestedType(ctx *grammar.NestedTypeContext) interface{} {
	return t.Visit(ctx.DeclType())
}

// VisitSliceType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSliceType(ctx *grammar.SliceTypeContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitStructType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitStructType(ctx *grammar.StructTypeContext) interface{} {
	return t.VisitChildren(ctx)
}

var _ grammar.MinigoVisitor = &TypeChecker{}

// VisitOperationPrecedence1 implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitOperationPrecedence1(ctx *grammar.OperationPrecedence1Context) interface{} {
	leftType, leftOk := t.Visit(ctx.GetLeft()).(*symboltable.Symbol)
	rightType, rightOk := t.Visit(ctx.GetRight()).(*symboltable.Symbol)

	if leftType == nil || rightType == nil {
		return nil // unrecoverable
	}
	leftType = getType(leftType)
	rightType = getType(rightType)

	if (leftOk && rightOk) && (!leftType.SameType(rightType)) {
		t.MakeError(ctx.GetStart(), fmt.Errorf("mismatched types: '%s' and '%s'", leftType, rightType))
	}

	if leftType == nil {
		return rightType
	}
	return leftType
}

// VisitOperationPrecedence2 implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitOperationPrecedence2(ctx *grammar.OperationPrecedence2Context) interface{} {
	leftType, leftOk := t.Visit(ctx.GetLeft()).(*symboltable.Symbol)
	rightType, rightOk := t.Visit(ctx.GetRight()).(*symboltable.Symbol)

	if leftType == nil || rightType == nil {
		return nil // unrecoverable
	}
	leftType = getType(leftType)
	rightType = getType(rightType)

	if (leftOk && rightOk) && (leftType != rightType) {
		t.MakeError(ctx.GetStart(), fmt.Errorf("mismatched types: '%s' and '%s'", leftType, rightType))
	}

	if leftType == nil {
		return rightType
	}
	return leftType
}

// VisitExpressionPostDec implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionPostDec(ctx *grammar.ExpressionPostDecContext) interface{} {
	expr := ctx.Expression()
	symbol := t.Visit(expr).(*symboltable.Symbol)
	if getType(symbol) != symboltable.Int {
		t.MakeError(expr.GetStart(), fmt.Errorf("cannot use expression of type '%s' as int in post decrement statement", getType(symbol)))
	}
	return nil
}

// VisitExpressionPostInc implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionPostInc(ctx *grammar.ExpressionPostIncContext) interface{} {
	expr := ctx.Expression()
	symbol := t.Visit(expr).(*symboltable.Symbol)
	if getType(symbol) != symboltable.Int {
		t.MakeError(expr.GetStart(), fmt.Errorf("cannot use expression of type '%s' as int in post increment statement", getType(symbol)))
	}
	return nil
}

// VisitSwitchCaseBranch implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSwitchCaseBranch(ctx *grammar.SwitchCaseBranchContext) interface{} {
	// switch {
	// case "hola":
	// 	println("adios")
	// }
	// if the current type is nil that means it's a switch statement without a statement
	currentType, _ := t.symbolStack.Peek()
	if currentType == nil {
		currentType = symboltable.Bool
	}

	types := t.Visit(ctx.ExpressionList()).([]*symboltable.Symbol)
	for _, _type := range types {
		if getType(_type) != getType(currentType) {
			t.MakeError(ctx.GetStart(), fmt.Errorf("cannot use symbol of type '%s' as '%s' type in case statement", getType(_type), getType(currentType)))
		}
	}
	return nil
}

// VisitSwitchDefaultBranch implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSwitchDefaultBranch(ctx *grammar.SwitchDefaultBranchContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitNormalSwitch implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitNormalSwitch(ctx *grammar.NormalSwitchContext) interface{} {
	t.symbolStack.Push(nil)
	return t.VisitChildren(ctx)
}

// VisitNormalSwitchExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitNormalSwitchExpression(ctx *grammar.NormalSwitchExpressionContext) interface{} {
	expr := ctx.Expression()
	symbol, ok := t.Visit(expr).(*symboltable.Symbol)
	if !ok {
		return nil // unreachable
	}
	symbolType := getType(symbol)
	err := t.symbolStack.Push(symbolType)
	if err != nil {
		panic("unreachable")
	}

	return t.Visit(ctx.ExpressionCaseClauseList())
}

// VisitSimpleStatementSwitch implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSimpleStatementSwitch(ctx *grammar.SimpleStatementSwitchContext) interface{} {
	panic("unimplemented")
}

// VisitSimpleStatementSwitchExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSimpleStatementSwitchExpression(ctx *grammar.SimpleStatementSwitchExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitIfElseBlock implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIfElseBlock(ctx *grammar.IfElseBlockContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitIfElseIf implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIfElseIf(ctx *grammar.IfElseIfContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitIfSimpleElseBlock implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIfSimpleElseBlock(ctx *grammar.IfSimpleElseBlockContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitIfSimpleElseIf implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIfSimpleElseIf(ctx *grammar.IfSimpleElseIfContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitIfSimpleNoElse implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIfSimpleNoElse(ctx *grammar.IfSimpleNoElseContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitIfSingleExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIfSingleExpression(ctx *grammar.IfSingleExpressionContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitInfiniteFor implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitInfiniteFor(ctx *grammar.InfiniteForContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitThreePartFor implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitThreePartFor(ctx *grammar.ThreePartForContext) interface{} {
	first := ctx.GetFirst()
	last := ctx.GetLast()
	t.Visit(first)
	t.Visit(last)
	expr := ctx.Expression()
	_type := t.Visit(expr).(*symboltable.Symbol)
	if getType(_type) != symboltable.Bool {
		t.MakeError(expr.GetStart(), fmt.Errorf("cannot use expression of type '%s' as boolean expression", getType(_type)))
	}
	return t.Visit(ctx.Block())
}

// VisitWhileFor implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitWhileFor(ctx *grammar.WhileForContext) interface{} {
	expressionType, ok := t.Visit(ctx.Expression()).(*symboltable.Symbol)
	if !ok {
		return nil // unrecoverable
	}

	_type := getType(expressionType)

	if _type != symboltable.Bool {
		t.MakeError(ctx.GetStart(), fmt.Errorf("cannot use expression of type '%s' in for condition", _type))
	}
	return t.Visit(ctx.Block())
}

// VisitAssignmentSimpleStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitAssignmentSimpleStatement(ctx *grammar.AssignmentSimpleStatementContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitExpressionSimpleStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionSimpleStatement(ctx *grammar.ExpressionSimpleStatementContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitWalrusDeclaration implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitWalrusDeclaration(ctx *grammar.WalrusDeclarationContext) interface{} {
	rhs := ctx.GetRight().AllExpression()
	lhs := ctx.GetLeft().AllExpression()

	idenLen := len(lhs)
	exprLen := len(rhs)
	if idenLen != exprLen {
		t.MakeError(ctx.GetStart(), fmt.Errorf("assignment mismatch: %d variables but %d vaules", idenLen, exprLen))
		return nil
	}
	for idx, ident := range lhs {
		expression := rhs[idx]

		right, ok := t.Visit(expression).(*symboltable.Symbol)
		if !ok {
			// panic(t.MakeError(expression.GetStart(), fmt.Errorf("Something went terribly wrong")))
			continue // unrecoverable
		}

		symbol := t.SymbolTable.NewVariable(ident.GetStart(), ident.GetText(), right)
		err := t.SymbolTable.AddSymbol(symbol)
		if err != nil {
			t.MakeError(ident.GetStart(), err)
		}

		sType := getType(symbol)
		rType := getType(right)

		if sType != rType {
			t.MakeError(ctx.GetStart(), fmt.Errorf("cannot use '%s' as '%s' value in assignment", sType, rType))
		}
	}
	return nil
}

// VisitExpressionOperand implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionOperand(ctx *grammar.ExpressionOperandContext) interface{} {
	return t.Visit(ctx.GetChild(1).(antlr.ParseTree))
}

// VisitIdentifierOperand implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIdentifierOperand(ctx *grammar.IdentifierOperandContext) interface{} {
	if ident := ctx.IDENTIFIER(); ident != nil {
		symbol, found := t.SymbolTable.GetSymbol(ident.GetText())
		if !found {
			t.MakeError(ctx.GetStart(), fmt.Errorf("unknown symbol: %s", ident.GetText()))
			return nil
		}
		return symbol // XXX: I removed the .Type call here, so now I gotta figure out what got affected by this
	}
	return t.VisitChildren(ctx)
}

// VisitLiteralOperand implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitLiteralOperand(ctx *grammar.LiteralOperandContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitAppendCall implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitAppendCall(ctx *grammar.AppendCallContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitCapCall implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitCapCall(ctx *grammar.CapCallContext) interface{} {
	return t.VisitChildren(ctx)
}

func (t *TypeChecker) FindGlobalSymbol(name string) (*symboltable.Symbol, bool) {
	return t.SymbolTable.Symbols.FindFirst(func(s *symboltable.Symbol) bool {
		return s.Scope == 0 && s.Name == name
	})
}

// VisitFunctionCall implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFunctionCall(ctx *grammar.FunctionCallContext) interface{} {
	fn, ok := t.Visit(ctx.PrimaryExpression()).(*symboltable.Symbol)
	if !ok {
		return nil
	}
	exprCtx, ok := ctx.Arguments().ExpressionList(0).(*grammar.ExpressionListContext)
	if !ok {
		return nil //unrecoverable
	}
	args, ok := t.VisitExpressionList(exprCtx).([]*symboltable.Symbol)
	if !ok {
		return nil // unrecoverable
	}
	expectedMembers := len(fn.Members)
	gotArgs := len(args)

	if gotArgs != expectedMembers {
		t.MakeError(ctx.GetStop(), fmt.Errorf("expected %d arguments but got %d instead", expectedMembers, gotArgs))
	}

	if expectedMembers <= gotArgs {
		for idx, member := range fn.Members {
			arg := getType(args[idx])
			if member.Type != arg {
				t.MakeError(ctx.GetStart(), fmt.Errorf("expected type '%s' but got '%s' instead", member.Type, arg))
			}
		}
	}
	return fn.Type
}

// VisitLenCall implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitLenCall(ctx *grammar.LenCallContext) interface{} {
	return t.VisitLengthExpression(ctx.LengthExpression().(*grammar.LengthExpressionContext))
	//return t.VisitChildren(ctx)
}

// VisitMemberAccessor implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitMemberAccessor(ctx *grammar.MemberAccessorContext) interface{} {
	op, ok := t.Visit(ctx.PrimaryExpression()).(*symboltable.Symbol)
	if !ok {
		return nil // unrecoverable
		// panic(t.MakeError(ctx.GetStart(), errors.New("something went terribly wrong")))
	}
	err := t.symbolStack.Push(op)
	if err != nil {
		panic("unreachable")
	}
	return t.VisitChildren(ctx)
}

// VisitOperandExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitOperandExpression(ctx *grammar.OperandExpressionContext) interface{} {
	return t.VisitChildren(ctx)
}

func getInnerMostType(in *symboltable.Symbol) *symboltable.Symbol {
	if in.Type != nil {
		if in.Is(symboltable.SliceSymbol | symboltable.ArraySymbol) {
			return in
		}
		return getInnerMostType(in.Type)
	}
	return in
}

// From this function I have to return the type of the value that's getting indexed
func (t *TypeChecker) VisitSubIndex(ctx *grammar.SubIndexContext) interface{} {
	accessee, ok := t.Visit(ctx.PrimaryExpression()).(*symboltable.Symbol)
	if !ok {
		return nil
	}

	if !getInnerMostType(accessee).Is(symboltable.ArraySymbol | symboltable.SliceSymbol) {
		t.MakeError(ctx.GetStart(), fmt.Errorf("symbol '%s' is not a slice or array type", accessee.Name))
		return nil // unrecoverable
	}

	index, ok := t.Visit(ctx.Index()).(*symboltable.Symbol)
	if !ok {
		return nil // unreachable
	}

	if getType(index) != symboltable.Int {
		t.MakeError(accessee.Token, fmt.Errorf("cannot use value of type '%s' as index selector", getType(index)))
		return nil // unrecoverable
	}
	return accessee.Type
}

// VisitInPlaceAssignment implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitInPlaceAssignment(ctx *grammar.InPlaceAssignmentContext) interface{} {
	lhs := ctx.GetLeft()
	rhs := ctx.GetRight()
	symbol, ok := t.Visit(lhs).(*symboltable.Symbol)
	if !ok {
		return nil // unrecoverable
	}
	right, ok := t.Visit(rhs).(*symboltable.Symbol)
	if !ok {
		//t.MakeError(ctx.GetStart(), errors.New("something went terribly wrong")))
		return nil // unrecoverable
	}

	symbolType := getType(symbol)
	rightType := getType(right)

	if symbolType != rightType {
		t.MakeError(ctx.GetStart(), fmt.Errorf("cannot use '%s' as '%s' value in assignment", rightType, symbolType))
	}

	return nil
}

// VisitNormalAssignment implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitNormalAssignment(ctx *grammar.NormalAssignmentContext) interface{} {
	rhs := ctx.GetRight().AllExpression()
	lhs := ctx.GetLeft().AllExpression()
	lhsLen := len(lhs)
	rhsLen := len(rhs)
	if rhsLen != lhsLen {
		t.MakeError(ctx.GetStart(), fmt.Errorf("assignment mismatch: %d variables but %d values", lhsLen, rhsLen))
		return nil //
	}
	for idx, ident := range lhs {
		expression := rhs[idx]
		symbol, ok := t.Visit(ident).(*symboltable.Symbol)
		if !ok {
			return nil // unrecoverable
		}
		t.symbolStack.Push((symbol))
		right, ok := t.Visit(expression).(*symboltable.Symbol)
		if !ok {
			return nil // unrecoverable
		}
		t.symbolStack.Pop()

		symbolType := getType(symbol)
		rightType := getType(right)
		if symbolType != rightType {
			t.MakeError(ctx.GetStart(), fmt.Errorf("cannot use '%s' as '%s' value in assignment", rightType, symbolType))
			return nil // unrecoverable

		}
	}

	return nil
}

// VisitCaretExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitCaretExpression(ctx *grammar.CaretExpressionContext) interface{} {
	expr := ctx.Expression()
	symbol, ok := t.Visit(expr).(*symboltable.Symbol)
	if !ok {
		return nil // unrecoverable
	}
	if getType(symbol) != symboltable.Int {
		t.MakeError(expr.GetStart(), fmt.Errorf("cannot use expression of type '%s'", getType(symbol)))
	}
	return nil
}

// VisitExpressionPrimaryExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionPrimaryExpression(ctx *grammar.ExpressionPrimaryExpressionContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitNotExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitNotExpression(ctx *grammar.NotExpressionContext) interface{} {
	expr := ctx.Expression()
	symbol, ok := t.Visit(expr).(*symboltable.Symbol)
	if !ok {
		return nil // unrecoverable
	}
	symbolType := getType(symbol)
	if symbolType != symboltable.Bool {
		t.MakeError(expr.GetStart(), fmt.Errorf("expression of type '%s' is not boolean", symbolType))
	}

	return nil // unrecoverable
}

func secondToLastType(in *symboltable.Symbol) *symboltable.Symbol {
	if in != nil {
		if in.Is(symboltable.TypeSymbol) {
			return secondToLastType(in.Type)
		}

		return in.Type
	}

	return in
}

func getType(in *symboltable.Symbol) *symboltable.Symbol {

	// if in.Type != nil && !in.Is(symboltable.SliceSymbol|symboltable.SliceSymbol) {
	// 	return getType(in.Type)
	// }

	if in.Type != nil {
		return in.Type
	}

	return in

	// if in != nil && (in.SymbolType&symboltable.TypeSymbol != 0) {
	// 	if in.Type != nil {
	// 		return in.Type
	// 	}
	// 	return in
	// }

	// // return getType(in.Type)
	// return getType(in.Type)
}

// VisitComparison implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitComparison(ctx *grammar.ComparisonContext) interface{} {
	leftType, leftOk := t.Visit(ctx.GetLeft()).(*symboltable.Symbol)
	rightType, rightOk := t.Visit(ctx.GetRight()).(*symboltable.Symbol)

	leftType = getType(leftType)
	rightType = getType(rightType)

	if !(rightOk && leftOk) {
		return nil // unrecoverable
	}

	if (leftOk && rightOk) && (leftType != rightType) {
		t.MakeError(ctx.GetStart(), fmt.Errorf("mismatched types: '%s' and '%s'", leftType, rightType))
	}

	return symboltable.Bool
}

// VisitBooleanOperation implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitBooleanOperation(ctx *grammar.BooleanOperationContext) interface{} {
	leftType, leftOk := t.Visit(ctx.GetLeft()).(*symboltable.Symbol)
	rightType, rightOk := t.Visit(ctx.GetRight()).(*symboltable.Symbol)
	leftType = getType(leftType)
	rightType = getType(rightType)

	if !(leftOk && rightOk) {
		return nil // BUG: this could cause a potential bug unrecoverable
	}

	if (leftOk && rightOk) && (leftType != rightType) {
		t.MakeError(ctx.GetStart(), fmt.Errorf("mismatched types: '%s' and '%s'", leftType, rightType))
	}

	return symboltable.Bool
}

// VisitOperationExpression implements grammar.MinigoVisitor.

// VisitFloatLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFloatLiteral(ctx *grammar.FloatLiteralContext) interface{} {
	return symboltable.Float
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
	funcs := ctx.TopDeclarationList().AllFuncDecl()

	types := ctx.TopDeclarationList().AllTypeDecl()
	for _, _type := range types {
		t.Visit(_type)
	}

	// I did this piece of shit to register functions in the global scope, for
	// cross reference
	for _, fn := range funcs {
		fun := fn.FuncFrontDecl()
		fun.IDENTIFIER()
		var rt *symboltable.Symbol
		if returnType := fun.DeclType(); returnType != nil {
			rt = t.Visit(returnType).(*symboltable.Symbol)
		}

		var members []*symboltable.Symbol = nil
		if arguments := fun.FuncArgsDecls(); arguments != nil {
			for _, variable := range arguments.AllSingleVarDeclNoExps() {
				varType, ok := t.Visit(variable.DeclType()).(*symboltable.Symbol)
				if !ok {
					continue // unrecoverable
				}
				for _, identifier := range variable.IdentifierList().AllIDENTIFIER() {
					symbol := t.SymbolTable.NewVariable(identifier.GetSymbol(), identifier.GetText(), getType(varType))
					symbol = makeSlice(symbol, varType)
					members = append(members, symbol)
				}
			}
		}

		symbol := t.SymbolTable.NewFunction(fun.IDENTIFIER().GetSymbol(), fun.IDENTIFIER().GetText(), rt, members...)
		symbol = makeSlice(symbol, rt)
		err := t.SymbolTable.AddSymbol(symbol)
		if err != nil {
			t.MakeError(ctx.GetStart(), err)
			return nil
		}
	}

	variables := ctx.TopDeclarationList().AllVariableDecl()
	for _, variable := range variables {
		t.Visit(variable)
	}

	funcs = ctx.TopDeclarationList().AllFuncDecl()
	for _, _func := range funcs {
		t.Visit(_func)
	}

	// t.SymbolTable.EnterScope()
	// defer t.SymbolTable.ExitScope()
	// return t.VisitChildren(ctx)
	return nil
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

func (t *TypeChecker) MakeError(token antlr.Token, err error) {
	t.listener.SyntaxError(nil, token, token.GetLine(), token.GetColumn(), err.Error(), nil)
	//return fmt.Errorf("%s:%d:%d: %w", t.filename, token.GetLine(), token.GetColumn(), err)
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
	identifiers := ctx.IdentifierList().AllIDENTIFIER()
	_type := t.Visit(ctx.DeclType()).(*symboltable.Symbol)
	expressions := ctx.ExpressionList().AllExpression()
	if idenLen, exprLen := len(identifiers), len(expressions); idenLen != exprLen {
		t.MakeError(ctx.GetStart(), fmt.Errorf("assignment mismatch: %d variables but %d values", idenLen, exprLen))
		return nil
	}

	for idx, identifier := range identifiers {
		expression := expressions[idx]
		expressionType, ok := t.Visit(expression).(*symboltable.Symbol)
		if !ok {
			t.MakeError(ctx.GetStart(), fmt.Errorf("invalid type")) // This should be unreachable
		}
		expressionType = getType(expressionType)
		symbolType := getType(_type)

		if expressionType != symbolType {
			t.MakeError(expression.GetStart(), fmt.Errorf("cannot use expression of type '%s' as '%s' value in assignment", symbolType, expressionType))
			continue
		}

		symbol := t.SymbolTable.NewVariable(identifier.GetSymbol(), identifier.GetText(), getType(expressionType))
		// BUG: this might cause a bug, check later
		// symbol = makeSlice(symbol, expressionType)
		err := t.SymbolTable.AddSymbol(symbol)
		if err != nil {
			t.MakeError(identifier.GetSymbol(), err)
		}
	}

	return nil
}

// VisitUntypedVarDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitUntypedVarDecl(ctx *grammar.UntypedVarDeclContext) interface{} {
	identifiers := ctx.IdentifierList().AllIDENTIFIER()
	expressions := ctx.ExpressionList().AllExpression()
	if idenLen, exprLen := len(identifiers), len(expressions); idenLen != exprLen {
		t.MakeError(ctx.GetStart(), fmt.Errorf("assignment mismatch: %d variables but %d vaules", idenLen, exprLen))
		return nil // Unrecoverable
	}
	for idx, ident := range identifiers {
		expr := expressions[idx]
		_type, ok := t.Visit(expr).(*symboltable.Symbol)
		if !ok {
			t.MakeError(ctx.GetStart(), fmt.Errorf("invalid type"))
		}
		symbol := t.SymbolTable.NewVariable(ident.GetSymbol(), ident.GetText(), _type)
		err := t.SymbolTable.AddSymbol(symbol)
		if err != nil {
			t.MakeError(ident.GetSymbol(), err)
		}
	}

	return nil
}

// VisitAppendExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitAppendExpression(ctx *grammar.AppendExpressionContext) interface{} {
	// TODO: rework thie piece of shit to use the new system
	sliceExpr := ctx.GetSlice()
	slice := t.Visit(sliceExpr).(*symboltable.Symbol)
	valueExpr := ctx.GetValue()
	value := t.Visit(valueExpr).(*symboltable.Symbol)
	sliceOrArray := symboltable.SliceSymbol

	leftSideExpression, _ := t.symbolStack.Peek()
	if !leftSideExpression.Type.Is(sliceOrArray) {
		t.MakeError(ctx.GetStart(), fmt.Errorf("cannot use expression of type '%s' as slice expression in append statement", leftSideExpression.Type))
		return nil
	}
	if !slice.Type.Is(sliceOrArray) {
		t.MakeError(sliceExpr.GetStart(), fmt.Errorf("cannot use expression of type '%s' as slice expression in append statement", slice.Type))
		return nil
	}

	sliceType := getType(slice)
	valueType := getType(value)
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("value: %v\n", value)
	fmt.Printf("sliceType: %v\n", sliceType)
	fmt.Printf("valueType: %v\n", valueType)

	if !sliceType.Type.SameType(valueType) {
		t.MakeError(valueExpr.GetStart(), fmt.Errorf("cannot use expression of type '%s' as '%s' expression in append statement", valueType, slice))
	}

	return slice
}

// VisitArguments implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitArguments(ctx *grammar.ArgumentsContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitBlock implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitBlock(ctx *grammar.BlockContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitBlockStatement implements grammar.MinigoVisitor.o
func (t *TypeChecker) VisitBlockStatement(ctx *grammar.BlockStatementContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitBreakStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitBreakStatement(ctx *grammar.BreakStatementContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitCapExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitCapExpression(ctx *grammar.CapExpressionContext) interface{} {
	expr := ctx.Expression()
	_type := t.Visit(expr).(*symboltable.Symbol)
	if _type.SymbolType&(symboltable.ArraySymbol|symboltable.SliceSymbol) == 0 {
		t.MakeError(expr.GetStart(), fmt.Errorf("expression in cap call is not a slice or array type"))
		return nil
	}
	return symboltable.Int
}

// VisitContinueStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitContinueStatement(ctx *grammar.ContinueStatementContext) interface{} {
	return t.VisitChildren(ctx)
}

type declTypePayload struct {
	symbol  *symboltable.Symbol
	IsSlice bool
	IsArray bool
}

// VisitDeclType implements grammar.MinigoVisitor.

// VisitEmptyVariableDeclaration implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitEmptyVariableDeclaration(ctx *grammar.EmptyVariableDeclarationContext) interface{} {
	return t.VisitChildren(ctx)
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
	t.Visit(ctx.ExpressionSwitchCase())
	return t.Visit(ctx.StatementList())
}

// VisitExpressionCaseClauseList implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionCaseClauseList(ctx *grammar.ExpressionCaseClauseListContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitExpressionList implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionList(ctx *grammar.ExpressionListContext) interface{} {
	exprs := ctx.AllExpression()
	out := []*symboltable.Symbol{}
	for _, expr := range exprs {
		_type, ok := t.Visit(expr).(*symboltable.Symbol)
		if !ok {
			// panic("TODO: fix this piece of shit, this should never happen")
			continue
		}
		out = append(out, _type)
	}
	return out
}

// VisitExpressionSwitchCase implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionSwitchCase(ctx *grammar.ExpressionSwitchCaseContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitFuncArgsDecls implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFuncArgsDecls(ctx *grammar.FuncArgsDeclsContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitFuncDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFuncDecl(ctx *grammar.FuncDeclContext) interface{} {
	defer t.SymbolTable.ExitScope() // This will exit the scope created in the FuncFrontDecl
	val, found := t.SymbolTable.GetSymbol(ctx.FuncFrontDecl().IDENTIFIER().GetText())
	if !found {
		panic("TODO: proper error message here")
	}
	// t.currentFunction = val
	t.symbolStack.Push(val)
	res := t.VisitChildren(ctx)
	_, _ = t.symbolStack.Pop()
	// t.currentFunction = nil
	return res
}

// VisitFuncFrontDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFuncFrontDecl(ctx *grammar.FuncFrontDeclContext) interface{} {
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

// VisitIfStatementStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIfStatementStatement(ctx *grammar.IfStatementStatementContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitIndex implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIndex(ctx *grammar.IndexContext) interface{} {
	// return t.VisitChildren(ctx)
	return t.Visit(ctx.GetChild(1).(*grammar.ExpressionPrimaryExpressionContext))
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
	_type, ok := t.Visit(ctx.Expression()).(*symboltable.Symbol)
	if !ok {
		return nil // unrecoverable
	}

	if !_type.Is(symboltable.SliceSymbol | symboltable.ArraySymbol) {
		t.MakeError(_type.Token, fmt.Errorf("cannot use symbol of type '%s' in len call", getType(_type)))
		return nil
	}

	return symboltable.Int
}

// VisitLoop implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitLoop(ctx *grammar.LoopContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitLoopStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitLoopStatement(ctx *grammar.LoopStatementContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitMultiVariableDeclaration implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitMultiVariableDeclaration(ctx *grammar.MultiVariableDeclarationContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitOperand implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitOperand(ctx *grammar.OperandContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitPrintStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitPrintStatement(ctx *grammar.PrintStatementContext) interface{} {
	log.Println("TODO: Implement this piece of shit (VisitPrintStatement)")
	return t.VisitChildren(ctx)
}

// VisitPrintlnStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitPrintlnStatement(ctx *grammar.PrintlnStatementContext) interface{} {
	log.Println("TODO: Implement this piece of shit (VisitPrintlnStatement)")
	return t.VisitChildren(ctx)
}

// VisitReturnStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitReturnStatement(ctx *grammar.ReturnStatementContext) interface{} {
	expr := ctx.Expression()
	val, ok := t.Visit(expr).(*symboltable.Symbol)
	if !ok {
		return nil // unrecoverable
		// panic(t.MakeError(expr.GetStart(), fmt.Errorf("something went terribly wrong")))
	}

	valType := getType(val)
	// if !valType.Equals(t.currentFunction.Type) {
	currentFunction, _ := t.symbolStack.Peek()
	if valType != getType(currentFunction) {
		t.MakeError(expr.GetStart(), fmt.Errorf("cannot return value of type '%s' in function with return type of '%s'", valType, getType(currentFunction)))
	}
	return nil
}

// VisitSelector implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSelector(ctx *grammar.SelectorContext) interface{} {
	// XXX: I need to store the current selected item
	symbol, err := t.symbolStack.Pop()
	if err != nil {
		panic("unreachable")
	}
	memberName := ctx.IDENTIFIER().GetText()
	symbolType := getType(symbol)
	idx := slices.IndexFunc(symbolType.Type.Members, func(s *symboltable.Symbol) bool {
		return s.Name == memberName
	})
	if idx == -1 {
		t.MakeError(ctx.IDENTIFIER().GetSymbol(), fmt.Errorf("symbol '%s' of type '%s' has no member called '%s'", symbol.Name, symbolType, memberName))
		return nil
	}

	return symbolType.Type.Members[idx]
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
	// TODO: fix this piece of shit
	name := ctx.IDENTIFIER()
	symbol := t.SymbolTable.NewAliasType(name.GetSymbol(), name.GetText(), nil)

	err := t.symbolStack.Push(symbol)
	if err != nil {
		panic("unreachable")
	}
	_type, ok := t.Visit(ctx.DeclType()).(*symboltable.Symbol)
	if !ok {
		return nil //unreachable
	}
	_, err = t.symbolStack.Pop()
	if err != nil {
		panic("unreachable")
	}
	// symbol = _type
	if _type.Is(symboltable.SliceSymbol | symboltable.ArraySymbol) {
		name := symbol.Name
		symbol.Name = name
		symbol.SymbolType = _type.SymbolType
		symbol.Members = _type.Members
		symbol.Size = _type.Size
		symbol.Scope = _type.Scope
		symbol.Type = _type.Type
	} else {
		symbol.Type = _type
	}

	// err = t.SymbolTable.AddSymbol(symbol)
	// if err != nil {
	// 	t.MakeError(name.GetSymbol(), err)
	// }

	return symbol
	// if typeName := declType.IDENTIFIER(); typeName != nil {
	// 	_type, found := t.SymbolTable.GetSymbol(typeName.GetText())
	// 	if !found {
	// 		t.MakeError(typeName.GetSymbol(), fmt.Errorf("unknown symbol '%s'", typeName.GetText()))
	// 		return nil // Can't continue from here, it will just not work
	// 	}
	// 	symbol := t.SymbolTable.NewAliasType(name.GetSymbol(), name.GetText(), _type)
	// 	err := t.SymbolTable.AddSymbol(symbol)
	// 	if err != nil {
	// 		t.MakeError(name.GetSymbol(), err)
	// 	}
	// } else if sliceDecl := declType.SliceDeclType(); sliceDecl != nil {
	// 	typeName := sliceDecl.DeclType().IDENTIFIER().GetText()
	// 	_type, found := t.SymbolTable.GetSymbol(typeName)
	// 	if !found {
	// 		t.MakeError(sliceDecl.GetStart(), fmt.Errorf("unknown symbol '%s'", typeName))
	// 		return nil // Can't continue from here, it will just not work
	// 	}
	// 	symbol := t.SymbolTable.NewSliceType(name.GetSymbol(), name.GetText(), _type)
	// 	err := t.SymbolTable.AddSymbol(symbol)
	// 	if err != nil {
	// 		t.MakeError(name.GetSymbol(), err)
	// 	}
	// } else if arrayDecl := declType.ArrayDeclType(); arrayDecl != nil {
	// 	originalSize := arrayDecl.INTLITERAL().GetText()
	// 	size, _ := strconv.ParseUint(originalSize, 10, 64)
	// 	typeName := arrayDecl.DeclType().IDENTIFIER().GetText()
	// 	_type, found := t.SymbolTable.GetSymbol(typeName)
	// 	if !found {
	// 		t.MakeError(arrayDecl.GetStart(), fmt.Errorf("unknown symbol '%s'", typeName))
	// 		return nil // Can't continue from here, it will just not work
	// 	}
	// 	symbol := t.SymbolTable.NewArrayType(name.GetSymbol(), name.GetText(), _type, size)
	// 	err := t.SymbolTable.AddSymbol(symbol)
	// 	if err != nil {
	// 		t.MakeError(name.GetSymbol(), err)
	// 	}
	// } else if structType := declType.StructDeclType(); structType != nil {
	// 	symbol := t.SymbolTable.NewStructType(name.GetSymbol(), name.GetText())
	// 	err := t.SymbolTable.AddSymbol(symbol)
	// 	if err != nil {
	// 		t.MakeError(ctx.GetStart(), err)
	// 		return nil
	// 	}
	// 	if members := structType.StructMemDecls(); members != nil {
	// 		decls := members.AllSingleVarDeclNoExps()
	// 		for _, decl := range decls {
	// 			idents := decl.IdentifierList().AllIDENTIFIER()
	// 			_type, ok := t.Visit(decl.DeclType()).(declTypePayload)
	// 			if !ok {
	// 				continue
	// 			}
	// 			for _, ident := range idents {
	// 				variable := t.SymbolTable.NewVariable(ident.GetSymbol(), ident.GetText(), _type.symbol)
	// 				variable = makeSlice(variable, _type)
	// 				symbol.Members = append(symbol.Members, variable)
	// 			}
	// 		}
	// 	}
	// 	return nil
	// }
}

// VisitSingleVarDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSingleVarDecl(ctx *grammar.SingleVarDeclContext) interface{} {
	return t.VisitChildren(ctx)
}

// Deprecated: this function was been deprecated
func makeSlice(symbol *symboltable.Symbol, _ *symboltable.Symbol) *symboltable.Symbol {
	return symbol
	// if _type == nil {
	// 	return symbol
	// }
	// if _type.Is(symboltable.SliceSymbol) {
	// 	symbol.SymbolType |= symboltable.SliceSymbol
	// }

	// if _type.Is(symboltable.ArraySymbol) {
	// 	symbol.SymbolType |= symboltable.ArraySymbol
	// }

	// return symbol
}

// VisitSingleVarDeclNoExps implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSingleVarDeclNoExps(ctx *grammar.SingleVarDeclNoExpsContext) interface{} {
	_type, ok := t.Visit(ctx.DeclType()).(*symboltable.Symbol)
	if !ok {
		return nil // Unrecoverable error
	}
	for _, ident := range ctx.IdentifierList().AllIDENTIFIER() {
		// BUG: this will fix case A, but will break case B (case A is when the
		// type is of the following form -> []string, case B is when the type)
		// is of the following form -> SliceAliasType
		symbol := t.SymbolTable.NewVariable(ident.GetSymbol(), ident.GetText(), _type)
		err := t.SymbolTable.AddSymbol(symbol)
		if err != nil {
			t.MakeError(ident.GetSymbol(), err)
		}
	}
	return nil
}

// VisitStatementList implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitStatementList(ctx *grammar.StatementListContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitStructMemDecls implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitStructMemDecls(ctx *grammar.StructMemDeclsContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitSwitchStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSwitchStatement(ctx *grammar.SwitchStatementContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitTerminal implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitTerminal(node antlr.TerminalNode) interface{} {
	// return node.GetText()
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
		cpt, ok := child.(antlr.ParseTree)
		if !ok {
			panic("unreachable")
		}
		result = cpt.Accept(v)
	}
	return result
}

func NewTypeChecker(filename string, listener antlr.ErrorListener) *TypeChecker {
	return &TypeChecker{
		filename: filename,
		listener: listener,
		// currentFunction: &symboltable.Symbol{},
		symbolStack: stack.NewStack[*symboltable.Symbol](100),
		SymbolTable: symboltable.NewSymbolTable(),
	}
}
