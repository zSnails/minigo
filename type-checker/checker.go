package checker

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/zSnails/minigo/grammar"
	"github.com/zSnails/minigo/stack"
	symboltable "github.com/zSnails/minigo/symbol-table"
)

type TypeChecker struct {
	listener    antlr.ErrorListener
	symbolStack *stack.Stack[*symboltable.Symbol]
	typeStack   *stack.Stack[types.Type]
	typeTable   *symboltable.TypeTable
	nodeStack   *stack.Stack[antlr.TerminalNode]
	symbolTable *symboltable.SymbolTable
}

// VisitNegativeExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitNegativeExpression(ctx *grammar.NegativeExpressionContext) interface{} {
	expr, ok := t.Visit(ctx.Expression()).(types.Type)
	if !ok {
		return nil
	}

	expr = depointerize(expr)
	if !(expr.Equal(types.I64) || expr.Equal(types.Double)) {
		t.makeError(ctx.Expression().GetStart(), fmt.Errorf("cannot use expression of type '%s' as numeric expression", expr.Name()))
		return nil
	}
	return expr
}

// VisitPositiveExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitPositiveExpression(ctx *grammar.PositiveExpressionContext) interface{} {
	expr, ok := t.Visit(ctx.Expression()).(types.Type)
	if !ok {
		return nil
	}

	if !(expr.Equal(types.I64) || expr.Equal(types.Double)) {
		t.makeError(ctx.Expression().GetStart(), fmt.Errorf("cannot use expression of type '%s' as numeric expression", expr.Name()))
		return nil
	}
	return expr
}

// VisitFuncDef implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFuncDef(ctx *grammar.FuncDefContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitNumericIntLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitNumericIntLiteral(ctx *grammar.NumericIntLiteralContext) interface{} {
	number := ctx.INTLITERAL()
	_, err := strconv.ParseInt(number.GetText(), 10, 64)
	if err != nil {
		t.makeError(number.GetSymbol(), errors.New("value exceeds 64 bit limit"))
	}
	return types.I64
}

// VisitNumerixHexLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitNumerixHexLiteral(ctx *grammar.NumerixHexLiteralContext) interface{} {
	number := ctx.HEXINTLITERAL()
	_, err := strconv.ParseInt(number.GetText(), 0, 64)
	if err != nil {
		t.makeError(number.GetSymbol(), errors.New("value exceed 64 bit limit"))
	}
	// return symboltable.Int
	return types.I64
}

// VisitArrayDeclType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitArrayDeclType(ctx *grammar.ArrayDeclTypeContext) interface{} {
	_symbol, ok := t.Visit(ctx.DeclType()).(types.Type)
	if !ok {
		return nil
	}
	size, _ := strconv.ParseUint(ctx.INTLITERAL().GetText(), 10, 64)
	return types.NewArray(size, _symbol)
}

// VisitSliceDeclType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSliceDeclType(ctx *grammar.SliceDeclTypeContext) interface{} {
	_symbol, ok := t.Visit(ctx.DeclType()).(types.Type)
	if !ok {
		return nil // unrecoverable
	}

	return types.NewArray(0, _symbol)
}

var structMemberTable = map[string]map[string]int{}

// VisitStructDeclType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitStructDeclType(ctx *grammar.StructDeclTypeContext) interface{} {
	node, err := t.nodeStack.Pop()
	if err != nil {
		panic("unreachable")
	}
	nodeName := node.GetText()

	var mems []types.Type = nil
	if members := ctx.StructMemDecls(); members != nil {
		decls := members.AllSingleVarDeclNoExps()

		current := 0
		for _, decl := range decls {
			idents := decl.IdentifierList().AllIDENTIFIER()
			declType := decl.DeclType()
			_type, ok := t.Visit(declType).(types.Type)
			if !ok {
				continue
			}
			for _, ident := range idents {
				name := ident.GetText()
				if _, ok := structMemberTable[nodeName]; !ok {
					structMemberTable[nodeName] = map[string]int{}
				}
				structMemberTable[nodeName][name] = current
				mems = append(mems, _type)
				current++
			}
		}
	}

	_struct := types.NewStruct(mems...)
	_struct.SetName(nodeName)
	return _struct
}

// VisitArrayType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitArrayType(ctx *grammar.ArrayTypeContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitIdentifierDeclType implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIdentifierDeclType(ctx *grammar.IdentifierDeclTypeContext) interface{} {
	name := ctx.IDENTIFIER()
	symbol, found := t.typeTable.GetSymbol(name.GetText())
	if !found {
		t.makeError(name.GetSymbol(), fmt.Errorf("unknown symbol '%s'", name.GetText()))
		return nil
	}
	return symbol.Type
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
	leftType, leftOk := t.Visit(ctx.GetLeft()).(types.Type)
	rightType, rightOk := t.Visit(ctx.GetRight()).(types.Type)

	if leftType == nil || rightType == nil {
		return nil // unrecoverable
	}

	leftType = depointerize(leftType)
	rightType = depointerize(rightType)

	if (leftOk && rightOk) && (!leftType.Equal(rightType)) {
		t.makeError(ctx.GetStart(), fmt.Errorf("mismatched types: '%s' and '%s'", leftType, rightType))
		return nil
	}

	if ctx.MOD() != nil && leftType.Equal(types.Double) {
		t.makeError(ctx.GetStart(), errors.New("modulo operation is not defined for floating point values"))
		return nil
	}

	return leftType
}

// VisitOperationPrecedence2 implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitOperationPrecedence2(ctx *grammar.OperationPrecedence2Context) interface{} {
	leftType, leftOk := t.Visit(ctx.GetLeft()).(types.Type)
	rightType, rightOk := t.Visit(ctx.GetRight()).(types.Type)

	// if leftType == nil || rightType == nil {
	// 	return nil // unrecoverable
	// }

	leftType = depointerize(leftType)
	rightType = depointerize(rightType)

	if (leftOk && rightOk) && (!leftType.Equal(rightType)) {
		t.makeError(ctx.GetStart(), fmt.Errorf("mismatched types: '%s' and '%s'", leftType, rightType))
	}

	if ctx.PLUS() == nil && ctx.MINUS() == nil {
		if (leftOk && rightOk) && !leftType.Equal(types.I64) {
			t.makeError(ctx.GetStart(), fmt.Errorf("bitwise operations only support integer types"))
		}
	}

	return leftType
}

// VisitExpressionPostDec implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionPostDec(ctx *grammar.ExpressionPostDecContext) interface{} {
	expr := ctx.Expression()
	symbol, ok := t.Visit(expr).(types.Type)
	if !ok {
		return nil
	}

	symbol = depointerize(symbol)
	if !symbol.Equal(types.I64) {
		t.makeError(expr.GetStart(), fmt.Errorf("cannot use expression of type '%s' as int in post decrement statement", symbol))
	}
	return nil
}

// VisitExpressionPostInc implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitExpressionPostInc(ctx *grammar.ExpressionPostIncContext) interface{} {
	expr := ctx.Expression()
	symbol, ok := t.Visit(expr).(types.Type)
	if !ok {
		return nil
	}

	symbol = depointerize(symbol)

	if !symbol.Equal(types.I64) {
		t.makeError(expr.GetStart(), fmt.Errorf("cannot use expression of type '%s' as int in post increment statement", symbol))
	}
	return nil
}

// VisitSwitchCaseBranch implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSwitchCaseBranch(ctx *grammar.SwitchCaseBranchContext) interface{} {
	currentType, _ := t.typeStack.Peek()

	types, ok := t.Visit(ctx.ExpressionList()).([]types.Type)
	if !ok {
		return nil
	}
	for _, _type := range types {
		_type = depointerize(_type)
		if !_type.Equal(currentType) {
			t.makeError(ctx.GetStart(), fmt.Errorf("cannot use symbol of type '%s' as '%s' type in case statement", _type, currentType))
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
	_ = t.typeStack.Push(types.I1)
	return t.VisitChildren(ctx)
}

// VisitNormalSwitchExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitNormalSwitchExpression(ctx *grammar.NormalSwitchExpressionContext) interface{} {
	expr := ctx.Expression()
	_type, ok := t.Visit(expr).(types.Type)
	if !ok {
		return nil // unreachable
	}
	err := t.typeStack.Push(_type)
	if err != nil {
		panic("unreachable")
	}

	return t.Visit(ctx.ExpressionCaseClauseList())
}

// VisitSimpleStatementSwitch implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSimpleStatementSwitch(ctx *grammar.SimpleStatementSwitchContext) interface{} {
	_ = t.typeStack.Push(types.I1)
	return t.VisitChildren(ctx)
}

// VisitSimpleStatementSwitchExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSimpleStatementSwitchExpression(ctx *grammar.SimpleStatementSwitchExpressionContext) interface{} {
	simpleStatement := ctx.SimpleStatement()
	t.Visit(simpleStatement)

	expr, ok := t.Visit(ctx.Expression()).(*symboltable.Symbol)
	if !ok {
		return nil
	}

	_ = t.symbolStack.Push(expr)
	clauseList := ctx.ExpressionCaseClauseList()
	return t.Visit(clauseList)
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
	t.Visit(ctx.SimpleStatement())
	_type, ok := t.Visit(ctx.Expression()).(types.Type)
	if !ok {
		return nil
	}

	_type = depointerize(_type)

	if !_type.Equal(types.I1) {
		t.makeError(ctx.Expression().GetStart(), fmt.Errorf("mismatched types '%s' and 'bool'", _type))
		return nil
	}

	t.Visit(ctx.Block())
	return t.Visit(ctx.IfStatement())
}

// VisitIfSimpleNoElse implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIfSimpleNoElse(ctx *grammar.IfSimpleNoElseContext) interface{} {
	t.Visit(ctx.SimpleStatement())
	_type, ok := t.Visit(ctx.Expression()).(types.Type)
	if !ok {
		return nil
	}

	_type = depointerize(_type)
	if !_type.Equal(types.I1) {
		t.makeError(ctx.Expression().GetStart(), fmt.Errorf("mismatched types '%s' and 'bool'", _type))
		return nil
	}
	return t.Visit(ctx.Block())
}

// VisitIfSingleExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIfSingleExpression(ctx *grammar.IfSingleExpressionContext) interface{} {
	_type, ok := t.Visit(ctx.Expression()).(types.Type)
	if !ok {
		return nil
	}

	_type = depointerize(_type)
	if !_type.Equal(types.I1) {
		t.makeError(ctx.Expression().GetStart(), fmt.Errorf("mismatched types '%s' and 'bool'", _type))
		return nil
	}
	return t.Visit(ctx.Block())
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
	_type, ok := t.Visit(expr).(types.Type)
	if !ok {
		return nil
	}

	_type = depointerize(_type)

	if !_type.Equal(types.I1) {
		t.makeError(expr.GetStart(), fmt.Errorf("cannot use expression of type '%s' as boolean expression", _type))
	}
	return t.Visit(ctx.Block())
}

// VisitWhileFor implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitWhileFor(ctx *grammar.WhileForContext) interface{} {
	_type, ok := t.Visit(ctx.Expression()).(types.Type)
	if !ok {
		return nil // unrecoverable
	}

	_type = depointerize(_type)

	if !_type.Equal(types.I1) {
		t.makeError(ctx.GetStart(), fmt.Errorf("cannot use expression of type '%s' in for condition", _type))
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
		t.makeError(ctx.GetStart(), fmt.Errorf("assignment mismatch: %d variables but %d vaules", idenLen, exprLen))
		return nil
	}
	for idx, ident := range lhs {
		expression := rhs[idx]

		right, ok := t.Visit(expression).(types.Type)
		if !ok {
			continue
		}

		name := ident.GetText()
		symbol := ir.NewGlobal(name, right)
		err := t.symbolTable.AddSymbol(name, symbol)
		if err != nil {
			t.makeError(ident.GetStart(), err)
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

		symbol, found := t.symbolTable.GetSymbol(ident.GetText())
		if !found {
			t.makeError(ctx.GetStart(), fmt.Errorf("unknown symbol: %s", ident.GetText()))
			return nil
		}
		switch symbol.Value.(type) {
		case *ir.Func:
			return symbol.Value
		default:
			return symbol.Value.Type()
		}
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
	return t.symbolTable.Symbols.FindFirst(func(s *symboltable.Symbol) bool {
		return s.Scope == 0 && s.Name == name
	})
}

// VisitFunctionCall implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFunctionCall(ctx *grammar.FunctionCallContext) interface{} {
	fn, ok := t.Visit(ctx.PrimaryExpression()).(*ir.Func)
	if !ok {
		return nil
	}

	if fn.Name() == "main" {
		t.makeError(ctx.GetStart(), fmt.Errorf("the main function cannot be called"))
		return nil
	}

	expectedMembers := len(fn.Params)
	exprCtx, ok := ctx.Arguments().ExpressionList().(*grammar.ExpressionListContext)
	if ok {
		args, ok := t.VisitExpressionList(exprCtx).([]types.Type)
		if !ok {
			return nil // unrecoverable
		}
		gotArgs := len(args)

		if (!fn.Sig.Variadic && gotArgs != expectedMembers) || (fn.Sig.Variadic && gotArgs < expectedMembers) {
			t.makeError(ctx.GetStop(), fmt.Errorf("expected %d arguments but got %d instead", expectedMembers, gotArgs))
			return nil
		}

		if expectedMembers <= gotArgs {
			for idx, member := range fn.Params {
				arg := args[idx]
				mem := depointerize(member.Typ)
				ar := depointerize(arg)
				if !mem.Equal(ar) {
					t.makeError(ctx.GetStart(), fmt.Errorf("expected type '%s' but got '%s' instead", member.Typ, arg))
				}
			}
		}
	} else {
		if expectedMembers > 0 {
			t.makeError(ctx.GetStart(), fmt.Errorf("expected at least %d arguments but got zero", expectedMembers))
			return nil
		}
	}

	return fn.Sig.RetType
}

// VisitLenCall implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitLenCall(ctx *grammar.LenCallContext) interface{} {
	return t.VisitLengthExpression(ctx.LengthExpression().(*grammar.LengthExpressionContext))
	//return t.VisitChildren(ctx)
}

// VisitMemberAccessor implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitMemberAccessor(ctx *grammar.MemberAccessorContext) interface{} {
	op, ok := t.Visit(ctx.PrimaryExpression()).(types.Type)
	if !ok {
		return nil // unrecoverable
	}

	err := t.typeStack.Push(op)
	if err != nil {
		panic("unreachable")
	}
	return t.Visit(ctx.Selector())
}

// VisitOperandExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitOperandExpression(ctx *grammar.OperandExpressionContext) interface{} {
	return t.VisitChildren(ctx)
}

// 	if in != nil && in.Type != nil {
// 		if in.Is(symboltable.SliceSymbol | symboltable.ArraySymbol) {
// 			return in
// 		}
// 	}
// 	return in
// }

// From this function I have to return the type of the value that's getting indexed
func (t *TypeChecker) VisitSubIndex(ctx *grammar.SubIndexContext) interface{} {
	accessee, ok := t.Visit(ctx.PrimaryExpression()).(types.Type)
	if !ok {
		return nil
	}

	accessee = depointerize(accessee)
	if !types.IsArray(accessee) {
		t.makeError(ctx.GetStart(), fmt.Errorf("symbol '%s' is not a slice or array type", accessee))
		return nil // unrecoverable
	}

	index, ok := t.Visit(ctx.Index()).(types.Type)
	if !ok {
		return nil // unreachable
	}

	index = depointerize(index)

	if !index.Equal(types.I64) {
		t.makeError(ctx.GetStart(), fmt.Errorf("cannot use value of type '%s' as index selector", index))
		return nil // unrecoverable
	}

	acc, ok := accessee.(*types.ArrayType)
	if !ok {
		return nil
	}
	return acc.ElemType
}

// VisitInPlaceAssignment implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitInPlaceAssignment(ctx *grammar.InPlaceAssignmentContext) interface{} {
	lhs := ctx.GetLeft()
	rhs := ctx.GetRight()

	symbol, ok := t.Visit(lhs).(types.Type)
	if !ok {
		return nil // unrecoverable
	}

	right, ok := t.Visit(rhs).(types.Type)
	if !ok {
		//t.MakeError(ctx.GetStart(), errors.New("something went terribly wrong")))
		return nil // unrecoverable
	}

	symbol = depointerize(symbol)
	right = depointerize(right)

	if !symbol.Equal(right) {
		t.makeError(ctx.GetStart(), fmt.Errorf("cannot use '%s' as '%s' value in assignment", right, symbol))
	}

	// if symbolType.Is(symboltable.ArraySymbol | symboltable.SliceSymbol) {
	// 		t.makeError(ctx.GetStart(), fmt.Errorf("cannot use '%s' as '%s' value in assignment", rightType, symbolType))
	// 	}
	// } else {
	// 	if symbolType != rightType {
	// 		t.makeError(ctx.GetStart(), fmt.Errorf("cannot use '%s' as '%s' value in assignment", rightType, symbolType))
	// 	}
	// }

	return nil
}

func depointerize(a types.Type) types.Type {
	if b, ok := a.(*types.PointerType); ok {
		if b.ElemType.Equal(types.I8) {
			return a
		}
		return depointerize(b.ElemType)
	}

	return a
}

// VisitNormalAssignment implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitNormalAssignment(ctx *grammar.NormalAssignmentContext) interface{} {
	rhs := ctx.GetRight().AllExpression()
	lhs := ctx.GetLeft().AllExpression()
	lhsLen := len(lhs)
	rhsLen := len(rhs)

	if rhsLen != lhsLen {
		t.makeError(ctx.GetStart(), fmt.Errorf("assignment mismatch: %d variables but %d values", lhsLen, rhsLen))
		return nil //
	}

	for idx, ident := range lhs {
		expression := rhs[idx]
		symbol, ok := t.Visit(ident).(types.Type)
		if !ok {
			return nil // unrecoverable
		}

		_ = t.typeStack.Push(symbol)
		right, ok := t.Visit(expression).(types.Type)
		if !ok {
			return nil // unrecoverable
		}
		_, _ = t.typeStack.Pop()

		symbol = depointerize(symbol)
		right = depointerize(right)

		if !symbol.Equal(right) {
			t.makeError(ctx.GetStart(), fmt.Errorf("cannot use '%s' as '%s' value in assignment", right, symbol))
			return nil // unrecoverable
		}

	}

	return nil
}

// VisitCaretExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitCaretExpression(ctx *grammar.CaretExpressionContext) interface{} {
	expr := ctx.Expression()
	_type, ok := t.Visit(expr).(types.Type)
	if !ok {
		return nil // unrecoverable
	}

	_type = depointerize(_type)
	if !_type.Equal(types.I64) {
		t.makeError(expr.GetStart(), fmt.Errorf("cannot use expression of type '%s' as 'int' value", _type))
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
	_type, ok := t.Visit(expr).(types.Type)
	if !ok {
		return nil // unrecoverable
	}

	_type = depointerize(_type)

	if !_type.Equal(types.I1) {
		t.makeError(expr.GetStart(), fmt.Errorf("expression of type '%s' is not boolean", _type))
	}

	return nil // unrecoverable
}

// VisitComparison implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitComparison(ctx *grammar.ComparisonContext) interface{} {
	leftType, leftOk := t.Visit(ctx.GetLeft()).(types.Type)
	rightType, rightOk := t.Visit(ctx.GetRight()).(types.Type)

	if !(rightOk && leftOk) {
		return nil // unrecoverable
	}

	leftType = depointerize(leftType)
	rightType = depointerize(rightType)

	if !leftType.Equal(rightType) {
		t.makeError(ctx.GetStart(), fmt.Errorf("mismatched types: '%s' and '%s'", leftType, rightType))
	}

	return types.I1
}

// VisitBooleanOperation implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitBooleanOperation(ctx *grammar.BooleanOperationContext) interface{} {
	leftType, leftOk := t.Visit(ctx.GetLeft()).(types.Type)
	rightType, rightOk := t.Visit(ctx.GetRight()).(types.Type)

	if !(leftOk && rightOk) {
		return nil
	}

	leftType = depointerize(leftType)
	rightType = depointerize(rightType)

	if (leftOk && rightOk) && (!leftType.Equal(rightType)) {
		t.makeError(ctx.GetStart(), fmt.Errorf("mismatched types: '%s' and '%s'", leftType, rightType))
	}

	return types.I1
}

// VisitOperationExpression implements grammar.MinigoVisitor.

// VisitFloatLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFloatLiteral(ctx *grammar.FloatLiteralContext) interface{} {
	return types.Double
}

// VisitIntLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIntLiteral(ctx *grammar.IntLiteralContext) interface{} {

	return t.VisitChildren(ctx)
}

// VisitInterpretedStringLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitInterpretedStringLiteral(ctx *grammar.InterpretedStringLiteralContext) interface{} {
	return types.I8Ptr
}

// VisitRawStringLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitRawStringLiteral(ctx *grammar.RawStringLiteralContext) interface{} {
	return types.I8Ptr
}

// VisitRuneLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitRuneLiteral(ctx *grammar.RuneLiteralContext) interface{} {
	return types.I8
}

// VisitRoot implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitRoot(ctx *grammar.RootContext) interface{} {
	funcs := ctx.TopDeclarationList().AllFuncDecl()

	_types := ctx.TopDeclarationList().AllTypeDecl()
	for _, _type := range _types {
		t.Visit(_type)
	}

	// I did this piece of shit to register functions in the global scope, for
	// cross reference
	for _, fn := range funcs {
		fun := fn.FuncFrontDecl()
		var rt types.Type = types.Void
		if returnType := fun.DeclType(); returnType != nil {
			var ok bool
			rt, ok = t.Visit(returnType).(types.Type)
			if !ok {
				return nil
			}
			rt = depointerize(rt)
		}

		if fun.IDENTIFIER().GetText() == "main" {
			if !rt.Equal(types.Void) {
				t.makeError(fun.IDENTIFIER().GetSymbol(), fmt.Errorf("the main function cannot have a return type"))
				return nil
			}
		}

		var members []*ir.Param = nil
		if arguments := fun.FuncArgsDecls(); arguments != nil {
			for _, variable := range arguments.AllSingleVarDeclNoExps() {
				varType, ok := t.Visit(variable.DeclType()).(types.Type)
				if !ok {
					continue // unrecoverable
				}
				for _, identifier := range variable.IdentifierList().AllIDENTIFIER() {
					symbol := ir.NewParam(identifier.GetText(), varType)
					members = append(members, symbol)
				}
			}
		}

		name := fun.IDENTIFIER().GetText()
		symbol := ir.NewFunc(name, rt, members...)
		err := t.symbolTable.AddSymbol(name, symbol)
		if err != nil {
			t.makeError(ctx.GetStart(), err)
			return nil
		}
	}

	variables := ctx.TopDeclarationList().AllVariableDecl()
	for _, variable := range variables {
		t.Visit(variable)
	}

	funcs2 := ctx.TopDeclarationList().AllFuncDef()
	for _, fn := range funcs2 {
		fun := fn.FuncFrontDecl()
		var rt types.Type = types.Void
		if returnType := fun.DeclType(); returnType != nil {
			var ok bool
			rt, ok = t.Visit(returnType).(types.Type)
			if !ok {
				return nil
			}
		}

		var members []*ir.Param = nil
		if arguments := fun.FuncArgsDecls(); arguments != nil {
			for _, variable := range arguments.AllSingleVarDeclNoExps() {
				varType, ok := t.Visit(variable.DeclType()).(types.Type)
				if !ok {
					continue // unrecoverable
				}
				for _, identifier := range variable.IdentifierList().AllIDENTIFIER() {
					symbol := ir.NewParam(identifier.GetText(), varType)
					members = append(members, symbol)
				}
			}
		}

		name := fun.IDENTIFIER().GetText()
		symbol := ir.NewFunc(name, rt, members...)
		err := t.symbolTable.AddSymbol(name, symbol)
		if err != nil {
			t.makeError(ctx.GetStart(), err)
			return nil
		}
	}

	funcs = ctx.TopDeclarationList().AllFuncDecl()
	for _, _func := range funcs {
		t.Visit(_func)
	}

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

func (t *TypeChecker) makeError(token antlr.Token, err error) {
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
	_type, ok := t.Visit(ctx.DeclType()).(types.Type)
	if !ok {
		return nil
	}
	expressions := ctx.ExpressionList().AllExpression()
	if idenLen, exprLen := len(identifiers), len(expressions); idenLen != exprLen {
		t.makeError(ctx.GetStart(), fmt.Errorf("assignment mismatch: %d variables but %d values", idenLen, exprLen))
		return nil
	}

	for idx, identifier := range identifiers {
		expression := expressions[idx]
		expressionType, ok := t.Visit(expression).(types.Type)
		if !ok {
			t.makeError(ctx.GetStart(), fmt.Errorf("invalid type")) // This should be unreachable
			continue
		}

		expressionType = depointerize(expressionType)
		_type = depointerize(_type)

		if !expressionType.Equal(_type) {
			t.makeError(expression.GetStart(), fmt.Errorf("cannot use '%s' as '%s' value in assignment", expressionType, _type))
			continue
		}

		name := identifier.GetText()
		symbol := ir.NewGlobal(name, expressionType)
		err := t.symbolTable.AddSymbol(name, symbol)
		if err != nil {
			t.makeError(identifier.GetSymbol(), err)
		}
	}

	return nil
}

// VisitUntypedVarDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitUntypedVarDecl(ctx *grammar.UntypedVarDeclContext) interface{} {
	identifiers := ctx.IdentifierList().AllIDENTIFIER()
	expressions := ctx.ExpressionList().AllExpression()
	if idenLen, exprLen := len(identifiers), len(expressions); idenLen != exprLen {
		t.makeError(ctx.GetStart(), fmt.Errorf("assignment mismatch: %d variables but %d vaules", idenLen, exprLen))
		return nil // Unrecoverable
	}
	for idx, ident := range identifiers {
		expr := expressions[idx]
		_type, ok := t.Visit(expr).(types.Type)
		if !ok {
			t.makeError(ctx.GetStart(), fmt.Errorf("invalid type"))
		}
		name := ident.GetText()
		symbol := ir.NewGlobal(name, _type)
		err := t.symbolTable.AddSymbol(name, symbol)
		if err != nil {
			t.makeError(ident.GetSymbol(), err)
		}
	}

	return nil
}

// VisitAppendExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitAppendExpression(ctx *grammar.AppendExpressionContext) interface{} {
	sliceExpr := ctx.GetSlice()
	slice, ok := t.Visit(sliceExpr).(types.Type)
	if !ok {
		return nil
	}
	valueExpr := ctx.GetValue()
	value, ok := t.Visit(valueExpr).(types.Type)
	if !ok {
		return nil
	}

	leftSideExpression, _ := t.typeStack.Peek()

	left := depointerize(leftSideExpression)

	slice = depointerize(slice)
	value = depointerize(value)

	// if !leftSideExpression.Type.Is(sliceOrArray) {
	if !types.IsArray(left) {
		t.makeError(ctx.GetStart(), fmt.Errorf("cannot use expression of type '%s' as slice expression in append statement", left))
		return nil
	}

	// if !slice.Type.Is(sliceOrArray) {
	if !types.IsArray(slice) {
		t.makeError(sliceExpr.GetStart(), fmt.Errorf("cannot use expression of type '%s' as slice expression in append statement", slice))
		return nil
	}

	// if !sliceType.Type.SameType(valueType) {
	{
		slice, ok := slice.(*types.ArrayType)
		if !ok {
			return nil
		}

		if slice.Len != 0 {
			t.makeError(valueExpr.GetStart(), fmt.Errorf("cannot use append expression on array value (can only be used with slices)"))
		}

		if !slice.ElemType.Equal(value) {
			t.makeError(valueExpr.GetStart(), fmt.Errorf("cannot use expression of type '%s' as '%s' expression in append statement", value, slice))
		}
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
	t.symbolTable.EnterScope()
	defer t.symbolTable.ExitScope()
	return t.VisitChildren(ctx)
}

// VisitBreakStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitBreakStatement(ctx *grammar.BreakStatementContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitCapExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitCapExpression(ctx *grammar.CapExpressionContext) interface{} {
	expr := ctx.Expression()
	symbol, ok := t.Visit(expr).(types.Type)
	if !ok {
		return nil
	}
	// if symbol.Type.SymbolType&(symboltable.ArraySymbol|symboltable.SliceSymbol) == 0 {
	symbol = depointerize(symbol)

	if !types.IsArray(symbol) {
		t.makeError(expr.GetStart(), fmt.Errorf("expression in cap call is not a slice or array type"))
		return nil
	}
	return types.I64
}

// VisitContinueStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitContinueStatement(ctx *grammar.ContinueStatementContext) interface{} {
	return t.VisitChildren(ctx)
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
	var out []types.Type = make([]types.Type, 0)
	for _, expr := range exprs {
		_type, ok := t.Visit(expr).(types.Type)
		if !ok {
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
	defer t.symbolTable.ExitScope() // This will exit the scope created in the FuncFrontDecl
	val, found := t.symbolTable.GetSymbol(ctx.FuncFrontDecl().IDENTIFIER().GetText())
	if !found {
		panic("unreachable")
	}
	// t.currentFunction = val
	_ = t.symbolStack.Push(val)
	res, ok := t.VisitChildren(ctx).(*ir.Func)
	if !ok {
		return nil
	}
	_, _ = t.symbolStack.Pop()
	// t.currentFunction = nil
	return res
}

// VisitFuncFrontDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFuncFrontDecl(ctx *grammar.FuncFrontDeclContext) interface{} {
	t.symbolTable.EnterScope() // This will enter the function scope
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
	t.symbolTable.EnterScope()
	defer t.symbolTable.ExitScope()
	return t.VisitChildren(ctx)
}

// VisitIndex implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIndex(ctx *grammar.IndexContext) interface{} {
	// return t.VisitChildren(ctx)
	idx := ctx.GetChild(1)
	if idx == nil {
		return nil
	}
	i, ok := idx.(antlr.ParseTree)
	if !ok {
		return nil
	}
	return t.Visit(i)
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
	symbol, ok := t.Visit(ctx.Expression()).(types.Type)
	if !ok {
		return nil // unrecoverable
	}

	symbol = depointerize(symbol)
	if !types.IsArray(symbol) {
		t.makeError(ctx.LEFTPARENTHESIS().GetSymbol(), fmt.Errorf("cannot use symbol of type '%s' in len call", symbol))
		return nil
	}

	return types.I64
}

// VisitLoopStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitLoopStatement(ctx *grammar.LoopStatementContext) interface{} {
	t.symbolTable.EnterScope()
	defer t.symbolTable.ExitScope()
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
	return t.VisitChildren(ctx)
}

// VisitPrintlnStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitPrintlnStatement(ctx *grammar.PrintlnStatementContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitReturnStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitReturnStatement(ctx *grammar.ReturnStatementContext) interface{} {
	cf, _ := t.symbolStack.Peek()

	currentFunction, ok := cf.Value.(*ir.Func)
	if !ok {
		return nil
	}

	expr := ctx.Expression()
	if expr == nil {
		if !currentFunction.Sig.RetType.Equal(types.Void) {
			t.makeError(ctx.RETURN().GetSymbol(), fmt.Errorf("empty return statement on function with %s return type", currentFunction.Sig.RetType))
			return nil
		}
		return nil
	}

	val, ok := t.Visit(expr).(types.Type)
	if !ok {
		return nil // unrecoverable
	}

	val = depointerize(val)

	if !val.Equal(currentFunction.Sig.RetType) {
		t.makeError(expr.GetStart(), fmt.Errorf("cannot return value of type '%s' in function with return type of '%s'", val, currentFunction.Sig.RetType))
	}

	return nil
}

// VisitSelector implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSelector(ctx *grammar.SelectorContext) interface{} {
	symbol, err := t.typeStack.Pop()
	if err != nil {
		panic("unreachable")
	}
	_struct, ok := depointerize(symbol).(*types.StructType)
	if !ok {
		return nil
	}
	stru := structMemberTable[_struct.Name()]
	memberName := ctx.IDENTIFIER().GetText()
	idx, ok := stru[memberName]
	if !ok {
		t.makeError(ctx.IDENTIFIER().GetSymbol(), fmt.Errorf("type '%s' has no member called '%s'", _struct, memberName))
		return nil
	}

	return _struct.Fields[idx]
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
	_ = t.nodeStack.Push(name)

	_type, ok := t.Visit(ctx.DeclType()).(types.Type)
	if !ok {
		return nil //unreachable
	}

	_ = t.typeTable.AddSymbol(name.GetText(), _type)

	return _type
}

// VisitSingleVarDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSingleVarDecl(ctx *grammar.SingleVarDeclContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitSingleVarDeclNoExps implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSingleVarDeclNoExps(ctx *grammar.SingleVarDeclNoExpsContext) interface{} {
	_type, ok := t.Visit(ctx.DeclType()).(types.Type)
	if !ok {
		return nil
	}
	for _, ident := range ctx.IdentifierList().AllIDENTIFIER() {
		name := ident.GetText()
		symbol := ir.NewGlobal(name, _type)
		err := t.symbolTable.AddSymbol(name, symbol)
		if err != nil {
			t.makeError(ident.GetSymbol(), err)
			return nil
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
	t.symbolTable.EnterScope()
	defer t.symbolTable.ExitScope()
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
	return t.VisitChildren(ctx)
}

// VisitVariableDeclaration implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitVariableDeclaration(ctx *grammar.VariableDeclarationContext) interface{} {
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
		listener:    listener,
		symbolStack: stack.NewStack[*symboltable.Symbol](100),
		typeStack:   stack.NewStack[types.Type](100),
		typeTable:   symboltable.NewTypeTable(),
		nodeStack:   stack.NewStack[antlr.TerminalNode](20),
		symbolTable: symboltable.NewSymbolTable(),
	}
}
