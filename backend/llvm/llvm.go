package llvm

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"strconv"

	"github.com/antlr4-go/antlr/v4"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/zSnails/minigo/backend/llvm/internal"
	"github.com/zSnails/minigo/grammar"
	"github.com/zSnails/minigo/stack"
	"github.com/zSnails/minigo/typetable"
)

const GLOBAL_SCOPE = 0

type LlvmBackend struct {
	listener    antlr.ErrorListener
	module      *ir.Module
	symbolTable *internal.LlvmSymbolTable
	typeTable   *typetable.TypeTable
	loopStack   *stack.Stack[*ir.Block]
	blockStack  *stack.Stack[*ir.Block]
	funcStack   *stack.Stack[*Func]
}

// VisitThreePartForNoExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitThreePartForNoExpression(ctx *grammar.ThreePartForNoExpressionContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	l.Visit(ctx.GetFirst())

	_for := fn.NewBlock("")
	end := ir.NewBlock("")

	blk.NewBr(_for)

	_ = l.blockStack.Push(_for)
	_ = l.loopStack.Push(end)

	l.Visit(ctx.Block())
	l.Visit(ctx.GetLast())

	_, _ = l.loopStack.Pop()

	got, _ := l.blockStack.Peek()
	if got.Term == nil {
		got.NewBr(_for)
	}

	end.Parent = fn.Func
	fn.Blocks = append(fn.Blocks, end)

	_ = l.blockStack.Push(end)

	return nil
}

// VisitNegativeExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitNegativeExpression(ctx *grammar.NegativeExpressionContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()
	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}
	switch expr.(type) {
	case *constant.Float:
		{
			alloca := fn.body.NewAlloca(expr.Type())
			blk.NewStore(expr, alloca)
			sub := blk.NewFSub(zerof, expr)
			blk.NewStore(sub, alloca)
			return alloca
		}
	case *constant.Int:
		{
			alloca := fn.body.NewAlloca(expr.Type())
			blk.NewStore(expr, alloca)
			sub := blk.NewSub(zero, expr)
			blk.NewStore(sub, alloca)
			return alloca
		}
	default:
		l.reportError(ctx.GetStart(), "negation expressions are only implemented for constant values")
	}

	return nil
}

// VisitPositiveExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitPositiveExpression(ctx *grammar.PositiveExpressionContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()
	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}
	switch expr.(type) {
	case *constant.Float:
		{
			alloca := fn.body.NewAlloca(expr.Type())
			blk.NewStore(expr, alloca)
			sub := blk.NewFAdd(zerof, expr)
			blk.NewStore(sub, alloca)
			return alloca
		}
	case *constant.Int:
		{
			alloca := fn.body.NewAlloca(expr.Type())
			blk.NewStore(expr, alloca)
			sub := blk.NewAdd(zero, expr)
			blk.NewStore(sub, alloca)
			return alloca
		}
	default:
		panic("unreachable")
	}
}

// VisitFuncDef implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitFuncDef(ctx *grammar.FuncDefContext) interface{} {
	isdef = true
	fn, ok := l.Visit(ctx.FuncFrontDecl()).(*Func)
	if !ok {
		return nil
	}
	isdef = false

	l.symbolTable.AddSymbol(fn.Name(), fn)

	return fn
}

type Func struct {
	*ir.Func
	body *ir.Block
}

func (l *LlvmBackend) String() string {
	return l.module.String()
}

func (l *LlvmBackend) GetModule() *ir.Module {
	return l.module
}

var putchar *ir.Func
var puts *ir.Func
var printf *ir.Func
var strcat *ir.Func
var strcmp *ir.Func

var basicInt *ir.Global
var basicFloat *ir.Global
var basicBool *ir.Global
var basicChar *ir.Global

func (l *LlvmBackend) addBuiltIns() {
	puts = l.module.NewFunc("puts", types.I32, ir.NewParam("", types.I8Ptr))
	putchar = l.module.NewFunc("putchar", types.I8, ir.NewParam("", types.I8))
	printf = l.module.NewFunc("printf", types.I64, ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I64))
	strcat = l.module.NewFunc("strcat", types.I8Ptr, ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I8Ptr))
	strcmp = l.module.NewFunc("strcmp", types.I64, ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I8Ptr))

	l.symbolTable.AddSymbol("printf", printf)
	l.symbolTable.AddSymbol("putchar", putchar)
	l.symbolTable.AddSymbol("puts", puts)
	l.symbolTable.AddSymbol("strcat", strcat)
	l.symbolTable.AddSymbol("strcmp", strcmp)

	basicInt = l.module.NewGlobalDef("", constant.NewCharArrayFromString("%lld\n\x00"))
	basicFloat = l.module.NewGlobalDef("", constant.NewCharArrayFromString("%lf\n\x00"))
	basicBool = l.module.NewGlobalDef("", constant.NewCharArrayFromString("%b\n\x00"))
	basicChar = l.module.NewGlobalDef("", constant.NewCharArrayFromString("%c\n\x00"))

}

func NewLlvmBackend(listener antlr.ErrorListener) *LlvmBackend {
	b := &LlvmBackend{
		listener:    listener,
		module:      ir.NewModule(),
		typeTable:   typetable.NewTypeTable(),
		symbolTable: internal.NewTable(),
		loopStack:   stack.NewStack[*ir.Block](100),
		blockStack:  stack.NewStack[*ir.Block](100),
		funcStack:   stack.NewStack[*Func](20),
	}

	b.addBuiltIns()
	return b
}

// Visit implements grammar.MinigoVisitor.
func (l *LlvmBackend) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(l)
}

// VisitAppendCall implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitAppendCall(ctx *grammar.AppendCallContext) interface{} {
	l.reportError(ctx.GetStart(), "append calls are not yet implemented")
	return nil
}

// VisitAppendExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitAppendExpression(ctx *grammar.AppendExpressionContext) interface{} {
	l.reportError(ctx.GetStart(), "append expressions are not yet implemented")
	return nil
}

// VisitArguments implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitArguments(ctx *grammar.ArgumentsContext) interface{} {
	l.reportError(ctx.GetStart(), "arguments are not yet implemented")
	return nil
}

// VisitArrayDeclType implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitArrayDeclType(ctx *grammar.ArrayDeclTypeContext) interface{} {
	_type, ok := l.Visit(ctx.DeclType()).(types.Type)
	if !ok {
		return nil
	}

	got, err := strconv.ParseUint(ctx.INTLITERAL().GetText(), 0, 64)
	if err != nil {
		panic(err)
	}

	return types.NewArray(got, _type)
}

// VisitArrayType implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitArrayType(ctx *grammar.ArrayTypeContext) interface{} {
	return l.VisitChildren(ctx)
}

// VisitAssignmentSimpleStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitAssignmentSimpleStatement(ctx *grammar.AssignmentSimpleStatementContext) interface{} {
	return l.Visit(ctx.AssignmentStatement())
}

// VisitBlock implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitBlock(ctx *grammar.BlockContext) interface{} {
	return l.Visit(ctx.StatementList())
}

// VisitBlockStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitBlockStatement(ctx *grammar.BlockStatementContext) interface{} {
	l.reportError(ctx.GetStart(), "blocks are not yet implemented")
	return nil
}

func getPointerValue(block *ir.Block, val value.Value) value.Value {
	if l, ok := val.Type().(*types.PointerType); ok {
		return block.NewLoad(l.ElemType, val)
	}
	return val
}

// VisitBooleanOperation implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitBooleanOperation(ctx *grammar.BooleanOperationContext) interface{} {
	blk, _ := l.blockStack.Peek()

	left, leftOk := l.Visit(ctx.GetLeft()).(value.Value)
	right, rightOk := l.Visit(ctx.GetRight()).(value.Value)
	if !(leftOk && rightOk) {
		return nil
	}

	left = getPointerValue(blk, left)
	right = getPointerValue(blk, right)

	switch {
	case ctx.AND() != nil:
		return blk.NewAnd(left, right)
	case ctx.OR() != nil:
		return blk.NewOr(left, right)
	default:
		panic("unreachable")
	}
}

// VisitCapCall implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitCapCall(ctx *grammar.CapCallContext) interface{} {
	l.reportError(ctx.GetStart(), "cap calls are not yet implemented")
	return nil
}

// VisitCapExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitCapExpression(ctx *grammar.CapExpressionContext) interface{} {
	l.reportError(ctx.GetStart(), "cap expressions are not yet implemented")
	return nil
}

// VisitCaretExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitCaretExpression(ctx *grammar.CaretExpressionContext) interface{} {
	l.reportError(ctx.GetStart(), "caret expressions are not yet implemented")
	return nil
}

// VisitChildren implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitChildren(node antlr.RuleNode) interface{} {
	children := node.GetChildren()

	var result any // Always remember that this will just return the top level
	for _, child := range children {
		c, ok := child.(antlr.ParseTree)
		if !ok {
			return nil
		}
		r := l.Visit(c)
		if r != nil {
			result = r
		}
	}

	return result
}

// VisitComparison implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitComparison(ctx *grammar.ComparisonContext) interface{} {
	fn, _ := l.blockStack.Peek()
	blk, _ := l.blockStack.Peek()

	left, leftOk := l.Visit(ctx.GetLeft()).(value.Value)
	right, rightOk := l.Visit(ctx.GetRight()).(value.Value)

	if !(leftOk && rightOk) {
		return nil
	}

	// TODO: Figure out how to do string comparison

	if types.IsPointer(left.Type()) {
		left = fn.NewLoad(right.Type(), left)
	}

	if types.IsPointer(right.Type()) {
		right = fn.NewLoad(left.Type(), right)
	}

	switch {
	case ctx.COMPARISON() != nil:
		if left.Type() == types.Double {
			return blk.NewFCmp(enum.FPredOEQ, left, right)
		}
		return blk.NewICmp(enum.IPredEQ, left, right)
	case ctx.GREATERTHAN() != nil:
		if left.Type() == types.Double {
			return blk.NewFCmp(enum.FPredOGT, left, right)
		}
		return blk.NewICmp(enum.IPredSGT, left, right)
	case ctx.LESSTHAN() != nil:
		if left.Type() == types.Double {
			return blk.NewFCmp(enum.FPredOLT, left, right)
		}
		return blk.NewICmp(enum.IPredSLT, left, right)
	case ctx.GREATERTHANEQUAL() != nil:
		if left.Type() == types.Double {
			return blk.NewFCmp(enum.FPredOGE, left, right)
		}
		return blk.NewICmp(enum.IPredSGE, left, right)
	case ctx.LESSTHANEQUAL() != nil:
		if left.Type() == types.Double {
			return blk.NewFCmp(enum.FPredOLE, left, right)
		}
		return blk.NewICmp(enum.IPredSLE, left, right)
	case ctx.NEGATION() != nil:
		if left.Type() == types.Double {
			return blk.NewFCmp(enum.FPredONE, left, right)
		}
		return blk.NewICmp(enum.IPredNE, left, right)
	default:
		l.reportError(ctx.GetStart(), "operation not defined")
	}

	panic("unreachable")
}

// VisitContinueStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitContinueStatement(ctx *grammar.ContinueStatementContext) interface{} {
	l.reportError(ctx.GetStart(), "continue statements are not yet implemented")
	return nil
}

// VisitEmptyTypeDeclaration implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitEmptyTypeDeclaration(ctx *grammar.EmptyTypeDeclarationContext) interface{} {
	return nil
}

// VisitEmptyVariableDeclaration implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitEmptyVariableDeclaration(ctx *grammar.EmptyVariableDeclarationContext) interface{} {
	return nil
}

// VisitErrorNode implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("unimplemented")
}

// VisitExpressionCaseClause implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionCaseClause(ctx *grammar.ExpressionCaseClauseContext) interface{} {
	l.reportError(ctx.GetStart(), "case clauses are not yet implemented")
	return nil
}

// VisitExpressionCaseClauseList implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionCaseClauseList(ctx *grammar.ExpressionCaseClauseListContext) interface{} {
	l.reportError(ctx.GetStart(), "case clause lists are not yet implemented")
	return nil
}

// VisitExpressionList implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionList(ctx *grammar.ExpressionListContext) interface{} {
	l.reportError(ctx.GetStart(), "expression lists are not yet implemented")
	return nil
}

// VisitExpressionOperand implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionOperand(ctx *grammar.ExpressionOperandContext) interface{} {
	return l.Visit(ctx.Expression())
}

// VisitExpressionPostDec implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionPostDec(ctx *grammar.ExpressionPostDecContext) interface{} {
	blk, _ := l.blockStack.Peek()
	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}

	var temp value.Value
	if types.IsPointer(expr.Type()) {
		temp = blk.NewLoad(types.I64, expr)
	} else {
		temp = expr
	}

	result := blk.NewSub(temp, constant.NewInt(types.I64, 1))
	blk.NewStore(result, expr)
	return nil
}

// VisitExpressionPostInc implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionPostInc(ctx *grammar.ExpressionPostIncContext) interface{} {
	blk, _ := l.blockStack.Peek()
	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}

	var sex value.Value
	if types.IsPointer(expr.Type()) {
		sex = blk.NewLoad(types.I64, expr)
		result := blk.NewAdd(sex, constant.NewInt(types.I64, 1))
		blk.NewStore(result, expr)
	} else {
		sex = expr
		blk.NewAdd(sex, constant.NewInt(types.I64, 1))
	}

	return nil
}

// VisitExpressionPrimaryExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionPrimaryExpression(ctx *grammar.ExpressionPrimaryExpressionContext) interface{} {
	return l.VisitChildren(ctx)
}

// VisitExpressionSimpleStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionSimpleStatement(ctx *grammar.ExpressionSimpleStatementContext) interface{} {
	return l.Visit(ctx.Expression())
}

// VisitFloatLiteral implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitFloatLiteral(ctx *grammar.FloatLiteralContext) interface{} {
	val, err := strconv.ParseFloat(ctx.FLOATLITERAL().GetText(), 64)
	if err != nil {
		panic(err)
	}

	return constant.NewFloat(types.Double, val)
}

// VisitFuncArgsDecls implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitFuncArgsDecls(ctx *grammar.FuncArgsDeclsContext) interface{} {
	panic("unimplemented")
}

const MAINFUNCNAME = "main"

// VisitFuncDecl implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitFuncDecl(ctx *grammar.FuncDeclContext) interface{} {

	fn, ok := l.Visit(ctx.FuncFrontDecl()).(*Func)
	if !ok {
		return nil
	}

	l.symbolTable.AddSymbol(fn.Name(), fn)

	_ = l.funcStack.Push(fn)
	l.symbolTable.EnterScope()
	l.Visit(ctx.Block())

	name := fn.Func.Name()
	if name != MAINFUNCNAME && !unicode.IsUpper(rune(name[0])) {
		fn.Func.Linkage = enum.LinkagePrivate
	}
	if fn.Func.Sig.RetType == types.Void && name != MAINFUNCNAME {
		blk, _ := l.blockStack.Peek()
		blk.NewRet(nil)
	} else if name == MAINFUNCNAME {
		fn.Func.Sig.RetType = types.I64
		l := len(fn.Blocks)
		if l > 0 {
			fn.Blocks[l-1].NewRet(zero)
		}
	}

	l.symbolTable.ExitScope()
	_, _ = l.blockStack.Pop()
	_, _ = l.funcStack.Pop()

	return nil
}

var isdef = false

// VisitFuncFrontDecl implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitFuncFrontDecl(ctx *grammar.FuncFrontDeclContext) interface{} {
	funcName := ctx.IDENTIFIER().GetText()

	var params []*ir.Param
	if fargs := ctx.FuncArgsDecls(); fargs != nil {
		for _, param := range fargs.AllSingleVarDeclNoExps() {
			for _, ident := range param.IdentifierList().AllIDENTIFIER() {
				_type, ok := l.Visit(param.DeclType()).(types.Type)
				if !ok {
					return nil
				}
				params = append(params, ir.NewParam(ident.GetText(), _type))
			}
		}
	}

	var rt types.Type = types.Void
	if typ := ctx.DeclType(); typ != nil {
		var ok bool
		rt, ok = l.Visit(typ).(types.Type)
		if !ok {
			return nil
		}
	}

	fn := l.module.NewFunc(funcName, rt, params...)

	var entry *ir.Block
	if !isdef {
		entry = fn.NewBlock("")
		for _, param := range params {
			alloca := entry.NewAlloca(param.Type())
			entry.NewStore(param, alloca)
			// l.moduleSymbolTable.Replace(param.Name(), alloca)
			l.symbolTable.AddSymbol(param.Name(), alloca)
		}

		_ = l.blockStack.Push(entry)
	}

	return &Func{
		Func: fn,
		body: entry,
	}
}

// VisitFunctionCall implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitFunctionCall(ctx *grammar.FunctionCallContext) interface{} {
	blk, _ := l.blockStack.Peek()

	callee, ok := l.Visit(ctx.PrimaryExpression()).(value.Named)
	if !ok {
		return nil
	}
	var args []value.Value

	if exprL := ctx.Arguments().ExpressionList(); exprL != nil {
		for _, expr := range exprL.AllExpression() {
			argument, ok := l.Visit(expr).(value.Value)
			if !ok {
				return nil
			}
			switch argument := argument.(type) {
			case *ir.Global:
				ptr := blk.NewGetElementPtr(types.I8, argument, zero)
				args = append(args, ptr)
			case *ir.InstGetElementPtr:
				switch argument.Type().(type) {
				case *types.PointerType:
					switch argument.ElemType {
					case types.I8:
						args = append(args, argument)
					default:
						load := blk.NewLoad(argument.ElemType, argument)
						args = append(args, load)
					}
				default:
					load := blk.NewLoad(argument.ElemType, argument) // BUG: this might cause yet another bug
					args = append(args, load)
				}
			case *ir.InstAlloca: // BUG: there's a bug with this piece of shit, I'll check it out tomorrow
				load := blk.NewLoad(argument.ElemType, argument)
				args = append(args, load)
			case constant.Constant, *ir.Param, ir.Instruction: // BUG: this might cause a bug
				args = append(args, argument)
			default:
				l.reportError(ctx.GetStart(), fmt.Sprintf("function call not defined for argument of type %s", argument))
			}
		}
	}

	return blk.NewCall(callee, args...)
}

// VisitIdentifierDeclType implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIdentifierDeclType(ctx *grammar.IdentifierDeclTypeContext) interface{} {
	typeName := ctx.IDENTIFIER().GetText()

	got, found := l.typeTable.GetSymbol(typeName)
	if !found {
		return nil
	}

	return got.Type
}

// VisitIdentifierList implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIdentifierList(ctx *grammar.IdentifierListContext) interface{} {
	l.reportError(ctx.GetStart(), "identifier lists are not yet implemented")
	return nil
}

// VisitIdentifierOperand implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIdentifierOperand(ctx *grammar.IdentifierOperandContext) interface{} {
	name := ctx.IDENTIFIER().GetText()

	switch name {
	case "int":
		return zero
	case "float":
		return zerof
	case "rune":
		return constant.NewInt(types.I8, 0)
	case "string":
		return l.module.NewGlobalDef("", constant.NewCharArrayFromString(""))
	}

	symbol, found := l.symbolTable.GetSymbol(name)
	if !found {
		return nil
	}

	return symbol.Symbol
}

// VisitIfElseBlock implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfElseBlock(ctx *grammar.IfElseBlockContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}

	if types.IsPointer(expr.Type()) {
		expr = blk.NewLoad(types.I1, expr)
	}

	True := fn.NewBlock("")

	_ = l.blockStack.Push(True)
	l.Visit(ctx.GetFirstBlock())
	ifBlock, _ := l.blockStack.Pop()

	False := fn.NewBlock("")
	_ = l.blockStack.Push(False)

	l.Visit(ctx.GetLastBlock())
	elseBlock, _ := l.blockStack.Pop()
	Done := fn.NewBlock("")

	if ifBlock != nil && ifBlock.Term == nil {
		ifBlock.NewBr(Done)
	}

	if elseBlock != nil && elseBlock.Term == nil {
		elseBlock.NewBr(Done)
	}

	blk.NewCondBr(expr, True, False)

	if True.Term == nil {
		True.NewBr(Done)
	}

	if False.Term == nil {
		False.NewBr(Done)
	}

	_ = l.blockStack.Push(Done)

	return nil
}

// VisitIfElseIf implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfElseIf(ctx *grammar.IfElseIfContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}
	if types.IsPointer(expr.Type()) {
		expr = blk.NewLoad(types.I1, expr)
	}

	True := fn.NewBlock("")

	_ = l.blockStack.Push(True)
	l.Visit(ctx.Block())
	ifBlock, _ := l.blockStack.Pop()

	False := fn.NewBlock("")
	_ = l.blockStack.Push(False)
	l.Visit(ctx.IfStatement())
	toConnect, _ := l.blockStack.Pop()

	Done := fn.NewBlock("")
	if toConnect != nil && toConnect.Term == nil {
		toConnect.NewBr(Done)
	}

	if ifBlock != nil && ifBlock.Term == nil {
		ifBlock.NewBr(Done)
	}

	blk.NewCondBr(expr, True, False)

	if True.Term == nil {
		True.NewBr(Done)
	}
	if False.Term == nil {
		False.NewBr(Done)
	}

	_ = l.blockStack.Push(Done)

	return nil
}

// VisitIfSimpleElseBlock implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfSimpleElseBlock(ctx *grammar.IfSimpleElseBlockContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	l.Visit(ctx.SimpleStatement()) // This will probably never return anything
	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}

	if types.IsPointer(expr.Type()) {
		expr = blk.NewLoad(types.I1, expr)
	}

	True := fn.NewBlock("")

	_ = l.blockStack.Push(True)
	l.Visit(ctx.GetFirstBlock())
	ifBlock, _ := l.blockStack.Pop()

	False := fn.NewBlock("")
	_ = l.blockStack.Push(False)
	l.Visit(ctx.GetLastBlock())
	elseBlock, _ := l.blockStack.Pop()

	Done := fn.NewBlock("")
	if ifBlock != nil && ifBlock.Term == nil {
		ifBlock.NewBr(Done)
	}
	if elseBlock != nil && elseBlock.Term == nil {
		elseBlock.NewBr(Done)
	}
	blk.NewCondBr(expr, True, False)
	if True.Term == nil {
		True.NewBr(Done)
	}

	if False.Term == nil {
		False.NewBr(Done)
	}
	_ = l.blockStack.Push(Done)

	return nil
}

// VisitIfSimpleElseIf implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfSimpleElseIf(ctx *grammar.IfSimpleElseIfContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	l.Visit(ctx.SimpleStatement()) // This will probably never return anything
	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}
	if types.IsPointer(expr.Type()) {
		expr = blk.NewLoad(types.I1, expr)
	}

	True := fn.NewBlock("")
	_ = l.blockStack.Push(True)
	l.Visit(ctx.Block())
	ifBlock, _ := l.blockStack.Pop()

	False := fn.NewBlock("")
	_ = l.blockStack.Push(False)
	l.Visit(ctx.IfStatement())
	got, _ := l.blockStack.Peek()

	blk.NewCondBr(expr, True, False)
	if True.Term == nil {
		True.NewBr(got)
	}

	if False.Term == nil {
		False.NewBr(got)
	}

	if ifBlock != nil && ifBlock.Term == nil {
		ifBlock.NewBr(got)
	}

	return nil
}

// VisitIfSimpleNoElse implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfSimpleNoElse(ctx *grammar.IfSimpleNoElseContext) interface{} {

	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	l.Visit(ctx.SimpleStatement()) // This will probably never return anything
	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}
	if types.IsPointer(expr.Type()) {
		expr = blk.NewLoad(types.I1, expr)
	}

	True := fn.NewBlock("")
	_ = l.blockStack.Push(True)
	l.Visit(ctx.Block())
	s, _ := l.blockStack.Pop()
	Done := fn.NewBlock("")
	if s != nil && s.Term == nil {
		s.NewBr(Done)
	}

	blk.NewCondBr(expr, True, Done)
	if True.Term == nil {
		True.NewBr(Done)
	}

	_ = l.blockStack.Push(Done)

	return nil
}

// VisitIfSingleExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfSingleExpression(ctx *grammar.IfSingleExpressionContext) interface{} {
	fn, _ := l.funcStack.Peek()
	blk, _ := l.blockStack.Peek()

	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}
	if types.IsPointer(expr.Type()) {
		expr = blk.NewLoad(types.I1, expr)
	}

	then := fn.NewBlock("")
	_ = l.blockStack.Push(then)
	l.Visit(ctx.Block())
	s, _ := l.blockStack.Pop()

	end := fn.NewBlock("")
	if s != nil && s.Term == nil {
		s.NewBr(end)
	}

	blk.NewCondBr(expr, then, end)
	if then.Term == nil {
		then.NewBr(end)
	}

	_ = l.blockStack.Push(end)

	return nil
}

// VisitIfStatementStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfStatementStatement(ctx *grammar.IfStatementStatementContext) interface{} {
	l.symbolTable.EnterScope()
	defer l.symbolTable.ExitScope()
	return l.Visit(ctx.IfStatement())
}

// VisitInPlaceAssignment implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitInPlaceAssignment(ctx *grammar.InPlaceAssignmentContext) interface{} {
	blk, _ := l.blockStack.Peek()
	left, leftOk := l.Visit(ctx.GetLeft()).(value.Value)
	right, rightOk := l.Visit(ctx.GetRight()).(value.Value)

	if !(leftOk && rightOk) {
		return nil
	}

	leftValue := getPointerValue(blk, left)
	rightValue := getPointerValue(blk, right)

	switch {
	case ctx.IDIV() != nil:
		{
			var div value.Value
			if leftValue.Type().Equal(types.Double) {
				div = blk.NewFDiv(leftValue, rightValue)
			} else {
				div = blk.NewSDiv(leftValue, rightValue)
			}
			return blk.NewStore(div, left)
		}
	case ctx.ISUB() != nil:
		{
			var sub value.Value
			if leftValue.Type().Equal(types.Double) {
				sub = blk.NewFSub(leftValue, rightValue)
			} else {
				sub = blk.NewSub(leftValue, rightValue)
			}
			return blk.NewStore(sub, left)
		}
	case ctx.IADD() != nil:
		{
			var add value.Value
			if leftValue.Type().Equal(types.Double) {
				add = blk.NewFAdd(leftValue, rightValue)
			} else {
				add = blk.NewAdd(leftValue, rightValue)
			}
			return blk.NewStore(add, left)
		}
	case ctx.IMUL() != nil:
		{
			var mul value.Value
			if leftValue.Type().Equal(types.Double) {
				mul = blk.NewFMul(leftValue, rightValue)
			} else {
				mul = blk.NewMul(leftValue, rightValue)
			}
			return blk.NewStore(mul, left)
		}
	case ctx.IRIGHTSHIFT() != nil:
		{
			shift := blk.NewLShr(leftValue, rightValue)
			if leftValue.Type().Equal(types.Double) {
				panic("unreachable")
			}

			return blk.NewStore(shift, left)
		}
		// case ctx.
	default:
		l.reportError(ctx.GetStart(), "operation not defined")
	}
	return nil
}

// VisitIndex implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIndex(ctx *grammar.IndexContext) interface{} {
	return l.Visit(ctx.Expression())
}

// VisitBreakStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitBreakStatement(ctx *grammar.BreakStatementContext) interface{} {
	blk, _ := l.blockStack.Peek()
	loop, err := l.loopStack.Peek()
	if err != nil {
		panic(err)
	}
	blk.NewBr(loop)
	return nil
}

// VisitInnerTypeDecls implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitInnerTypeDecls(ctx *grammar.InnerTypeDeclsContext) interface{} {
	l.reportError(ctx.GetStart(), "nested type declarations are not yet implemented")
	return nil
}

// VisitInnerVarDecls implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitInnerVarDecls(ctx *grammar.InnerVarDeclsContext) interface{} {
	l.reportError(ctx.GetStart(), "nested type declarations are not yet implemented")
	return nil
}

// VisitIntLiteral implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIntLiteral(ctx *grammar.IntLiteralContext) interface{} {
	return l.Visit(ctx.NumericLiteral())
}

func makeCstr(input string) string {
	return fmt.Sprintf("%s\x00", input)
}

// VisitLenCall implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitLenCall(ctx *grammar.LenCallContext) interface{} {
	return l.Visit(ctx.LengthExpression())
}

// VisitLengthExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitLengthExpression(ctx *grammar.LengthExpressionContext) interface{} {
	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}
	switch expr := expr.(type) {
	case *ir.InstAlloca:
		switch expr := expr.ElemType.(type) {
		case *types.ArrayType:
			return constant.NewInt(types.I64, int64(expr.Len))
		default:
			l.reportError(ctx.GetStart(), fmt.Sprintf("length expression is not defined for %s", expr))
		}
	default:
		l.reportError(ctx.GetStart(), fmt.Sprintf("length expression is not defined for %s", expr))
	}

	return nil
}

// VisitLiteralOperand implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitLiteralOperand(ctx *grammar.LiteralOperandContext) interface{} {
	return l.Visit(ctx.Literal())
}

// VisitLoopStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitLoopStatement(ctx *grammar.LoopStatementContext) interface{} {
	return l.Visit(ctx.Loop())
}

// VisitMemberAccessor implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitMemberAccessor(ctx *grammar.MemberAccessorContext) interface{} {
	l.reportError(ctx.GetStart(), "member accessors are not yet implemented")
	return nil
}

// VisitMultiTypeDeclaration implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitMultiTypeDeclaration(ctx *grammar.MultiTypeDeclarationContext) interface{} {
	l.reportError(ctx.GetStart(), "multiple type declarations are not yet implemented")
	return nil
}

// VisitMultiVariableDeclaration implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitMultiVariableDeclaration(ctx *grammar.MultiVariableDeclarationContext) interface{} {
	l.reportError(ctx.GetStart(), "multiple type declarations are not yet implemented")
	return nil
}

// VisitNestedType implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitNestedType(ctx *grammar.NestedTypeContext) interface{} {
	return l.Visit(ctx.DeclType())
}

var zero = constant.NewInt(types.I64, 0)
var zerof = constant.NewFloat(types.Double, 0.0)

func (l *LlvmBackend) GetSymbol(name string) (value.Value, bool) {
	symbol, found := l.symbolTable.GetSymbol(name)
	if !found {
		return nil, false
	}

	return symbol.Symbol, true
}

// VisitNormalAssignment implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitNormalAssignment(ctx *grammar.NormalAssignmentContext) interface{} {
	blk, err := l.blockStack.Peek()
	if err != nil {
		panic(err)
	}

	rhs := ctx.GetRight()
	for idx, ident := range ctx.GetLeft().AllExpression() {

		symbol, ok := l.Visit(ident).(value.Value)
		if !ok {
			return nil
		}
		expr, ok := l.Visit(rhs.Expression(idx)).(value.Value)
		if !ok {
			return nil
		}
		switch expr := expr.(type) {
		case *ir.InstAlloca:
			load := blk.NewLoad(expr.ElemType, expr)
			blk.NewStore(load, symbol)
		case *ir.InstAdd, *ir.InstSDiv, *ir.InstSub,
			*ir.InstMul, *ir.InstFAdd, *ir.InstFSub,
			*ir.InstFMul, *ir.InstFDiv, *ir.InstSRem, *ir.InstFRem:
			if types.IsPointer(expr.Type()) {
				load := blk.NewLoad(expr.Type(), expr)
				blk.NewStore(load, symbol)
			} else {
				blk.NewStore(expr, symbol)
			}
		case *ir.InstCall, *ir.InstLShr, *ir.InstShl:
			blk.NewStore(expr, symbol)
		case *ir.Global:
			ptr := blk.NewGetElementPtr(types.I8, expr, zero)
			l.symbolTable.Replace(symbol.Ident(), ptr)
		case constant.Constant:
			blk.NewStore(expr, symbol)
		case *ir.InstGetElementPtr:
			switch exp := expr.Type().(type) {
			case *types.PointerType:
				switch e := exp.ElemType.(type) {
				case *types.ArrayType:
					load := blk.NewLoad(e.ElemType, expr)
					blk.NewStore(load, symbol)
				case *types.PointerType:
					// BUG: this is prolly the cause of the current bug
					load := blk.NewLoad(e, expr)
					blk.NewStore(load, symbol)
				case *types.IntType:
					load := blk.NewLoad(e, expr)
					blk.NewStore(load, symbol)
				case constant.Constant:
					blk.NewStore(expr, symbol)
				default:
					l.reportError(ctx.GetStart(), fmt.Sprintf("normal assignment is not defined for %s", exp.ElemType))
				}
			case *types.IntType:
				load := blk.NewLoad(expr.ElemType, expr)
				blk.NewStore(load, symbol)
			default:
				blk.NewStore(expr, symbol)
			}
		default:
			l.reportError(ctx.GetStart(), fmt.Sprintf("normal assignment is not defined for %s", reflect.TypeOf(expr)))
		}
	}

	return nil
}

// VisitNormalSwitch implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitNormalSwitch(ctx *grammar.NormalSwitchContext) interface{} {
	l.reportError(ctx.GetStart(), "switch statements are not yet implemented")
	return nil
}

// VisitNormalSwitchExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitNormalSwitchExpression(ctx *grammar.NormalSwitchExpressionContext) interface{} {
	l.reportError(ctx.GetStart(), "switch statements are not yet implemented")
	return nil
}

// VisitNotExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitNotExpression(ctx *grammar.NotExpressionContext) interface{} {
	blk, _ := l.blockStack.Peek()
	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}

	expr = getPointerValue(blk, expr)

	return blk.NewXor(expr, constant.True)
}

// VisitNumericIntLiteral implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitNumericIntLiteral(ctx *grammar.NumericIntLiteralContext) interface{} {
	val, err := strconv.ParseInt(ctx.INTLITERAL().GetText(), 10, 64)
	if err != nil {
		panic(err)
	}

	return constant.NewInt(types.I64, val)
}

// VisitNumerixHexLiteral implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitNumerixHexLiteral(ctx *grammar.NumerixHexLiteralContext) interface{} {
	val, err := strconv.ParseInt(ctx.HEXINTLITERAL().GetText(), 0, 64)
	if err != nil {
		panic(err)
	}

	return constant.NewInt(types.I64, val)
}

// VisitOperandExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitOperandExpression(ctx *grammar.OperandExpressionContext) interface{} {
	return l.Visit(ctx.Operand())
}

// VisitOperationPrecedence1 implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitOperationPrecedence1(ctx *grammar.OperationPrecedence1Context) interface{} {
	blk, err := l.blockStack.Peek()
	if err != nil {
		panic(err)
	}

	left, leftOk := l.Visit(ctx.GetLeft()).(value.Value)
	right, rightOk := l.Visit(ctx.GetRight()).(value.Value)

	if !(leftOk && rightOk) {
		return nil
	}

	left = getPointerValue(blk, left)
	right = getPointerValue(blk, right)

	switch {
	case ctx.TIMES() != nil:
		if left.Type() == types.Double {
			return blk.NewFMul(left, right)
		}
		return blk.NewMul(left, right)
	case ctx.DIV() != nil:
		if left.Type() == types.Double {
			return blk.NewFDiv(left, right)
		}
		return blk.NewSDiv(left, right)
	case ctx.MOD() != nil:
		return blk.NewSRem(left, right)
	default:
		l.reportError(ctx.GetStart(), "operation not defined")
		return nil
	}
}

func (l *LlvmBackend) reportError(token antlr.Token, message string) {
	line, column := token.GetLine(), token.GetColumn()
	l.listener.SyntaxError(nil, token, line, column, message, nil)
}

// VisitOperationPrecedence2 implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitOperationPrecedence2(ctx *grammar.OperationPrecedence2Context) interface{} {
	blk, err := l.blockStack.Peek()
	if err != nil {
		panic(err)
	}

	left, leftOk := l.Visit(ctx.GetLeft()).(value.Value)
	right, rightOk := l.Visit(ctx.GetRight()).(value.Value)
	if !(leftOk && rightOk) {
		return nil
	}

	left = getPointerValue(blk, left)
	right = getPointerValue(blk, right)

	switch {
	// LEFTSHIFT | RIGHTSHIFT | AMPERSAND | AMPERSANDCARET | PLUS | MINUS | PIPE | CARET
	case ctx.PLUS() != nil:
		if types.IsArray(left.Type()) || left.Type().Equal(types.I8Ptr) {
			return blk.NewCall(strcat, left, right)
		} else {
			if left.Type() == types.Double {
				return blk.NewFAdd(left, right)
			}
			return blk.NewAdd(left, right)
		}
	case ctx.MINUS() != nil:
		if left.Type() == types.Double {
			return blk.NewFSub(left, right)
		}
		return blk.NewSub(left, right)
	case ctx.AMPERSAND() != nil:
		return blk.NewAnd(left, right)
	case ctx.RIGHTSHIFT() != nil:
		return blk.NewLShr(left, right)
	case ctx.LEFTSHIFT() != nil:
		return blk.NewShl(left, right)
	case ctx.CARET() != nil:
		return blk.NewXor(left, right)
	case ctx.PIPE() != nil:
		return blk.NewOr(left, right)
	default:
		l.reportError(ctx.GetStart(), "operation not defined")
	}
	return nil
}

// VisitPrintStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitPrintStatement(ctx *grammar.PrintStatementContext) interface{} {
	l.reportError(ctx.GetStart(), "print statements are not yet implemented")
	return nil
	// panic("TODO: fix print statement")
	// blk, _ := l.blockStack.Peek()

	// for _, expr := range ctx.ExpressionList().AllExpression() {
	// 	innerString := l.Visit(expr).(value.Named)
	// 	blk.NewCall(puts, innerString)
	// }

	// return nil
}

var newline = constant.NewInt(types.I8, 0x0A)

// VisitPrintlnStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitPrintlnStatement(ctx *grammar.PrintlnStatementContext) interface{} {
	blk, _ := l.blockStack.Peek()

	if exprlist := ctx.ExpressionList(); exprlist != nil {
		for _, expr := range exprlist.AllExpression() {
			val, ok := l.Visit(expr).(value.Value)
			if !ok {
				return nil
			}

			switch v := val.(type) {
			case *ir.Global:
				blk.NewCall(puts, val)
			case *ir.InstGetElementPtr:
				if types.IsArray(v.ElemType) {
					load := blk.NewLoad(v.ElemType, val)
					blk.NewCall(printf, basicChar, load)
				} else {
					blk.NewCall(puts, val)
				}
			case *ir.InstAlloca:
				switch d := v.ElemType.(type) {
				case *types.FloatType:
					load := blk.NewLoad(d, v)
					blk.NewCall(printf, basicFloat, load)
				case *types.PointerType:
					switch d.ElemType {
					case types.I8:
						load := blk.NewLoad(v.Type(), v)
						blk.NewCall(puts, load)
					default:
						l.reportError(ctx.GetStart(), fmt.Sprintf("println statement is not defined for type %s", d.ElemType))
					}
				case *types.IntType:
					load := blk.NewLoad(v.ElemType, v)
					switch v.ElemType {
					case types.I64:
						blk.NewCall(printf, basicInt, load)
					case types.I1:
						blk.NewCall(printf, basicBool, load)
					default:
						l.reportError(ctx.GetStart(), fmt.Sprintf("println statement not defined for %s", v.ElemType))
					}
				default:
					l.reportError(ctx.GetStart(), fmt.Sprintf("println statement not defined for %s", reflect.TypeOf(d)))
				}
			case *ir.InstFAdd, *ir.InstFSub, *ir.InstFMul, *ir.InstFDiv:
				blk.NewCall(printf, basicFloat, val)
			case *ir.InstAdd, *ir.InstSub, *ir.InstMul, *ir.InstSDiv, *ir.InstAnd, *ir.InstXor, *ir.InstOr:
				blk.NewCall(printf, basicInt, val)
			case *constant.Int:
				blk.NewCall(printf, basicInt, val)
			case *constant.Float:
				blk.NewCall(printf, basicFloat, val)
			case *ir.InstCall:
				switch typ := v.Sig().RetType.(type) {
				case *types.FloatType:
					blk.NewCall(printf, basicFloat, v)
				case *types.IntType:
					blk.NewCall(printf, basicInt, v)
				default:
					l.reportError(ctx.GetStart(), fmt.Sprintf("println statement not defined for %s", reflect.TypeOf(typ)))
				}
			default:
				l.reportError(ctx.GetStart(), fmt.Sprintf("println statement not defined for %s", reflect.TypeOf(v)))
			}
		}
	} else {
		blk.NewCall(putchar, newline)
	}

	return nil
}

// VisitRawStringLiteral implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitRawStringLiteral(ctx *grammar.RawStringLiteralContext) interface{} {
	res := strings.Trim(ctx.RAWSTRINGLITERAL().GetText(), "`")
	return l.module.NewGlobalDef("", constant.NewCharArrayFromString(makeCstr(res)))
}

// VisitInterpretedStringLiteral implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitInterpretedStringLiteral(ctx *grammar.InterpretedStringLiteralContext) interface{} {
	unquoted, err := strconv.Unquote(ctx.GetText())
	if err != nil {
		panic(err)
	}
	car := constant.NewCharArrayFromString(makeCstr(unquoted))
	return l.module.NewGlobalDef("", car)
}

// VisitReturnStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitReturnStatement(ctx *grammar.ReturnStatementContext) interface{} {
	fn, _ := l.funcStack.Peek()
	blk, _ := l.blockStack.Peek()

	nblk := fn.NewBlock("")
	blk.NewBr(nblk)

	lk := fn.NewBlock("")
	lk.NewBr(lk)

	if fn.Name() == MAINFUNCNAME {
		_ = l.blockStack.Push(lk)
		return nblk.NewRet(zero)
	}

	if expr := ctx.Expression(); expr != nil {
		expr, ok := l.Visit(expr).(value.Value)
		if !ok {
			return nil
		}
		switch exp := expr.(type) {
		case *ir.Global:
			ptr := nblk.NewGetElementPtr(exp.ContentType, exp, zero)
			return nblk.NewRet(ptr)
		case *ir.InstAlloca:
			load := nblk.NewLoad(exp.ElemType, exp)
			return nblk.NewRet(load)
		case *ir.InstGetElementPtr:
			return nblk.NewRet(exp)
		case constant.Constant:
			return nblk.NewRet(exp)
		case *ir.InstAdd, *ir.InstMul, *ir.InstSDiv, *ir.InstFMul,
			*ir.InstOr, *ir.InstAnd, *ir.InstXor, *ir.InstCall,
			*ir.InstFDiv, *ir.InstFSub, *ir.InstFAdd:
			return nblk.NewRet(expr)
		default:
			l.reportError(ctx.GetStart(), fmt.Sprintf("return expression not defined for instruction %s", reflect.TypeOf(expr)))
		}
	}

	return nil
}

// var typeMap = map[string]types.Type{
// 	"string": types.I8Ptr,
// 	"int":    types.I64,
// 	"float":  types.Double,
// 	"bool":   types.I1,
// 	"rune":   types.I8,
// }

// VisitRoot implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitRoot(ctx *grammar.RootContext) interface{} {

	for _, _type := range ctx.TopDeclarationList().AllTypeDecl() {
		l.Visit(_type)
	}

	for _, decl := range ctx.TopDeclarationList().AllVariableDecl() {
		l.Visit(decl)
	}

	for _, fn := range ctx.TopDeclarationList().AllFuncDef() {
		l.Visit(fn)
	}

	for _, fn := range ctx.TopDeclarationList().AllFuncDecl() {
		l.Visit(fn)
	}

	return nil
}

// VisitRuneLiteral implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitRuneLiteral(ctx *grammar.RuneLiteralContext) interface{} {
	str, err := strconv.Unquote(ctx.RUNELITERAL().GetText())
	if err != nil {
		panic(err)
	}
	return constant.NewInt(types.I8, int64(str[0]))
}

// VisitSelector implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSelector(ctx *grammar.SelectorContext) interface{} {
	l.reportError(ctx.GetStart(), "member selectors are not yet implemented")
	return nil
}

// VisitSimpleStatementStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSimpleStatementStatement(ctx *grammar.SimpleStatementStatementContext) interface{} {
	return l.Visit(ctx.SimpleStatement())
}

// VisitSimpleStatementSwitch implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSimpleStatementSwitch(ctx *grammar.SimpleStatementSwitchContext) interface{} {
	l.reportError(ctx.GetStart(), "switch expressions are not yet implemented")
	return nil
}

// VisitSimpleStatementSwitchExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSimpleStatementSwitchExpression(ctx *grammar.SimpleStatementSwitchExpressionContext) interface{} {
	l.reportError(ctx.GetStart(), "switch expressions are not yet implemented")
	return nil
}

// VisitSingleTypeDecl implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSingleTypeDecl(ctx *grammar.SingleTypeDeclContext) interface{} {
	l.reportError(ctx.GetStart(), "type declarations are not yet implemented")
	return nil
}

// VisitSingleVarDeclNoExps implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSingleVarDeclNoExps(ctx *grammar.SingleVarDeclNoExpsContext) interface{} {

	// blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	expr, ok := l.Visit(ctx.DeclType()).(types.Type)
	if !ok {
		return nil
	}

	for _, ident := range ctx.IdentifierList().AllIDENTIFIER() {
		name := ident.GetText()
		if fn != nil && fn.body != nil {
			alloca := fn.body.NewAlloca(expr)
			l.symbolTable.AddSymbol(name, alloca)
		} else {
			def := l.module.NewGlobalDef("", constant.NewZeroInitializer(expr))
			l.symbolTable.Replace(name, def)
		}
	}

	return expr
}

// VisitSingleVarDeclsNoExpsDecl implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSingleVarDeclsNoExpsDecl(ctx *grammar.SingleVarDeclsNoExpsDeclContext) interface{} {
	return l.VisitChildren(ctx)
}

// VisitSliceDeclType implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSliceDeclType(ctx *grammar.SliceDeclTypeContext) interface{} {
	_type, ok := l.Visit(ctx.DeclType()).(types.Type)
	if !ok {
		return nil
	}
	return types.NewArray(0, _type)
}

// VisitSliceType implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSliceType(ctx *grammar.SliceTypeContext) interface{} {
	return l.VisitChildren(ctx)
}

// VisitStatementList implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitStatementList(ctx *grammar.StatementListContext) interface{} {
	return l.VisitChildren(ctx)
}

// VisitStructDeclType implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitStructDeclType(ctx *grammar.StructDeclTypeContext) interface{} {
	l.reportError(ctx.GetStart(), "struct type declarations are not yet implemented")
	return nil
}

// VisitStructMemDecls implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitStructMemDecls(ctx *grammar.StructMemDeclsContext) interface{} {
	l.reportError(ctx.GetStart(), "struct member declarations are not yet implemented")
	return nil
}

// VisitStructType implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitStructType(ctx *grammar.StructTypeContext) interface{} {
	l.reportError(ctx.GetStart(), "struct type declarations are not yet implemented")
	return nil
}

// VisitSubIndex implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSubIndex(ctx *grammar.SubIndexContext) interface{} {
	blk, _ := l.blockStack.Peek()

	expr, ok := l.Visit(ctx.PrimaryExpression()).(value.Value)
	if !ok {
		return nil
	}
	index := ctx.Index()
	idx, ok := l.Visit(index).(value.Value)
	if !ok {
		return nil
	}

	exprt, ok := expr.Type().(*types.PointerType)
	if !ok {
		return nil
	}

	idx = getPointerValue(blk, idx)

	gep := blk.NewGetElementPtr(exprt.ElemType, expr, zero, idx)
	return gep
}

// VisitSwitchCaseBranch implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSwitchCaseBranch(ctx *grammar.SwitchCaseBranchContext) interface{} {
	l.reportError(ctx.GetStart(), "switch case branches are not yet implemented")
	return nil
}

// VisitSwitchDefaultBranch implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSwitchDefaultBranch(ctx *grammar.SwitchDefaultBranchContext) interface{} {
	l.reportError(ctx.GetStart(), "switch default branches are not yet implemented")
	return nil
}

// VisitSwitchStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSwitchStatement(ctx *grammar.SwitchStatementContext) interface{} {
	l.reportError(ctx.GetStart(), "switch statements are not yet implemented")
	return nil
}

// VisitTerminal implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitTerminal(node antlr.TerminalNode) interface{} {
	// BUG: For some reason I'm returning the text of a terminal node, this might cause a bug later on
	return node.GetText()
}

// VisitThreePartFor implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitThreePartFor(ctx *grammar.ThreePartForContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	l.Visit(ctx.GetFirst())

	_for := fn.NewBlock("")
	end := ir.NewBlock("")

	blk.NewBr(_for)

	_ = l.blockStack.Push(_for)
	_ = l.loopStack.Push(end)

	l.Visit(ctx.Block())
	l.Visit(ctx.GetLast())

	_, _ = l.loopStack.Pop()

	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}

	got, _ := l.blockStack.Peek()
	if got.Term == nil {
		got.NewCondBr(expr, _for, end)
	}

	end.Parent = fn.Func
	fn.Blocks = append(fn.Blocks, end)

	_ = l.blockStack.Push(end)

	return nil
}

// VisitTopDeclarationList implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitTopDeclarationList(ctx *grammar.TopDeclarationListContext) interface{} {
	return l.VisitChildren(ctx)
}

// VisitTypeDeclStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitTypeDeclStatement(ctx *grammar.TypeDeclStatementContext) interface{} {
	l.reportError(ctx.GetStart(), "type decl statements are not yet implemented")
	return nil
}

// VisitTypeDeclaration implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitTypeDeclaration(ctx *grammar.TypeDeclarationContext) interface{} {
	return l.VisitChildren(ctx)
}

// VisitTypedVarDecl implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitTypedVarDecl(ctx *grammar.TypedVarDeclContext) interface{} {

	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	for idx, ident := range ctx.IdentifierList().AllIDENTIFIER() {
		name := ident.GetText()
		expr, ok := l.Visit(ctx.ExpressionList().Expression(idx)).(value.Value)
		if !ok {
			return nil
		}

		switch argument := expr.(type) {
		case *ir.InstAlloca:
			load := blk.NewLoad(argument.ElemType, expr)
			alloca := fn.body.NewAlloca(argument.ElemType)
			blk.NewStore(load, alloca)
			l.symbolTable.AddSymbol(name, alloca)
		case *ir.Global: // XXX: This might cause a bug somewhere else, I still have to run the checks
			switch arg := argument.ContentType.(type) {
			case *types.ArrayType:
				ptr := blk.NewGetElementPtr(arg.ElemType, argument, zero)
				l.symbolTable.AddSymbol(name, ptr)
			default:
				alloca := fn.body.NewAlloca(argument.Type())
				load := blk.NewLoad(argument.Type(), argument)
				blk.NewStore(load, alloca)
				l.symbolTable.AddSymbol(name, alloca)
			}
		case *ir.InstGetElementPtr:
			if l.symbolTable.GetScope() == GLOBAL_SCOPE {
				l.symbolTable.AddSymbol(name, argument)
			} else {
				ptr := blk.NewGetElementPtr(types.I8, argument, zero)
				l.symbolTable.AddSymbol(name, ptr)
			}
		case constant.Constant:
			if blk != nil {
				alloca := fn.body.NewAlloca(expr.Type())
				blk.NewStore(expr, alloca)
				l.symbolTable.AddSymbol(name, alloca)
			} else {
				def := l.module.NewGlobalDef("", argument)
				l.symbolTable.AddSymbol(name, def)
			}
		case ir.Instruction:
			alloca := fn.body.NewAlloca(expr.Type())
			blk.NewStore(expr, alloca)
			l.symbolTable.AddSymbol(name, alloca)
		default:
			l.reportError(ctx.GetStart(), fmt.Sprintf("typed var decl not defined not defined for %s", argument))
		}
	}

	return nil
}

// VisitUntypedVarDecl implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitUntypedVarDecl(ctx *grammar.UntypedVarDeclContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	for idx, ident := range ctx.IdentifierList().AllIDENTIFIER() {
		name := ident.GetText()
		expr, ok := l.Visit(ctx.ExpressionList().Expression(idx)).(value.Value)
		if !ok {
			return nil
		}

		switch argument := expr.(type) {
		case *ir.InstAlloca:
			load := blk.NewLoad(argument.ElemType, expr)
			alloca := fn.body.NewAlloca(argument.ElemType)
			blk.NewStore(load, alloca)
			l.symbolTable.AddSymbol(name, alloca)
		case *ir.Global: // XXX: This might cause a bug somewhere else, I still have to run the checks
			switch arg := argument.ContentType.(type) {
			case *types.ArrayType:
				ptr := blk.NewGetElementPtr(arg.ElemType, argument, zero)
				l.symbolTable.AddSymbol(name, ptr)
			default:
				alloca := fn.body.NewAlloca(argument.Type())
				load := blk.NewLoad(argument.Type(), argument)
				blk.NewStore(load, alloca)
				l.symbolTable.AddSymbol(name, alloca)
			}
		case *ir.InstGetElementPtr:
			if l.symbolTable.GetScope() == GLOBAL_SCOPE {
				l.symbolTable.AddSymbol(name, argument)
			} else {
				ptr := blk.NewGetElementPtr(types.I8, argument, zero)
				l.symbolTable.AddSymbol(name, ptr)
			}
		case constant.Constant:
			if blk != nil {
				alloca := fn.body.NewAlloca(expr.Type())
				blk.NewStore(expr, alloca)
				l.symbolTable.AddSymbol(name, alloca)
			} else {
				def := l.module.NewGlobalDef("", argument)
				l.symbolTable.AddSymbol(name, def)
			}
		case ir.Instruction:
			alloca := fn.body.NewAlloca(expr.Type())
			blk.NewStore(expr, alloca)
			l.symbolTable.AddSymbol(name, alloca)
		default:
			l.reportError(ctx.GetStart(), fmt.Sprintf("untyped var decl not defined not defined for %s", argument))
		}
	}

	return nil
}

// VisitVariableDeclStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitVariableDeclStatement(ctx *grammar.VariableDeclStatementContext) interface{} {
	return l.Visit(ctx.VariableDecl())
}

// VisitVariableDeclaration implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitVariableDeclaration(ctx *grammar.VariableDeclarationContext) interface{} {
	// No need to visit all cildren, when only this one child matters
	return l.Visit(ctx.SingleVarDecl())
}

// VisitWalrusDeclaration implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitWalrusDeclaration(ctx *grammar.WalrusDeclarationContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	for idx, ident := range ctx.GetLeft().AllExpression() {
		name := ident.GetText()

		expr, ok := l.Visit(ctx.GetRight().Expression(idx)).(value.Value)
		if !ok {
			return nil
		}

		switch argument := expr.(type) {
		case *ir.InstAlloca:
			load := blk.NewLoad(argument.ElemType, expr)
			alloca := fn.body.NewAlloca(argument.ElemType)
			blk.NewStore(load, alloca)
			l.symbolTable.AddSymbol(name, alloca)
		case *ir.Global: // XXX: This might cause a bug somewhere else, I still have to run the checks
			switch arg := argument.ContentType.(type) {
			case *types.ArrayType:
				ptr := blk.NewGetElementPtr(arg.ElemType, argument, zero)
				l.symbolTable.AddSymbol(name, ptr)
			default:
				alloca := fn.body.NewAlloca(argument.Type())
				load := blk.NewLoad(argument.Type(), argument)
				blk.NewStore(load, alloca)
				l.symbolTable.AddSymbol(name, alloca)
			}
		case *ir.InstGetElementPtr:
			if l.symbolTable.GetScope() == GLOBAL_SCOPE {
				l.symbolTable.AddSymbol(name, argument)
			} else {
				switch e := argument.ElemType.(type) {
				case *types.ArrayType:
					ptr := blk.NewGetElementPtr(argument.ElemType, argument, zero)
					alloca := fn.body.NewAlloca(e.ElemType)
					load := blk.NewLoad(e.ElemType, ptr)
					blk.NewStore(load, alloca)
					l.symbolTable.AddSymbol(name, alloca)
				case *types.IntType:
					alloca := fn.body.NewAlloca(expr.Type())
					blk.NewStore(expr, alloca)
					l.symbolTable.AddSymbol(name, alloca)
				default:
					l.reportError(ctx.GetStart(), fmt.Sprintf("walrus declaration not defined not defined for %s", argument.ElemType))
				}
			}
		case constant.Constant:
			alloca := fn.body.NewAlloca(expr.Type())
			blk.NewStore(expr, alloca)
			l.symbolTable.AddSymbol(name, alloca)
		case ir.Instruction:
			alloca := fn.body.NewAlloca(expr.Type())
			blk.NewStore(expr, alloca)
			l.symbolTable.AddSymbol(name, alloca)
		default:
			l.reportError(ctx.GetStart(), fmt.Sprintf("walrus declaration not defined not defined for %s", argument))
		}
	}

	return nil
}

// VisitWhileFor implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitWhileFor(ctx *grammar.WhileForContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	_for := fn.NewBlock("")
	end := ir.NewBlock("")

	blk.NewBr(_for)

	_ = l.blockStack.Push(_for)
	_ = l.loopStack.Push(end)

	l.Visit(ctx.Block())

	_, _ = l.loopStack.Pop()

	expr, ok := l.Visit(ctx.Expression()).(value.Value)
	if !ok {
		return nil
	}

	got, _ := l.blockStack.Peek()
	if got.Term == nil {
		got.NewCondBr(expr, _for, end)
	}

	end.Parent = fn.Func
	fn.Blocks = append(fn.Blocks, end)

	_ = l.blockStack.Push(end)

	return nil
}

// VisitInfiniteFor implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitInfiniteFor(ctx *grammar.InfiniteForContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	_for := fn.NewBlock("")
	end := ir.NewBlock("")

	blk.NewBr(_for)

	_ = l.blockStack.Push(_for)
	_ = l.loopStack.Push(end)

	l.Visit(ctx.Block())

	_, _ = l.loopStack.Pop()

	got, _ := l.blockStack.Peek()
	if got.Term == nil {
		got.NewBr(_for)
	}

	end.Parent = fn.Func
	fn.Blocks = append(fn.Blocks, end)

	_ = l.blockStack.Push(end)

	return nil
}

var _ grammar.MinigoVisitor = &LlvmBackend{}
