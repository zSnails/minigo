package llvm

import (
	"fmt"
	"reflect"

	"strconv"

	"github.com/antlr4-go/antlr/v4"

	// "github.com/JimWen/gods-generic/queues/arrayqueue"
	_ "github.com/llir/llvm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/zSnails/minigo/grammar"
	"github.com/zSnails/minigo/stack"
	symboltable "github.com/zSnails/minigo/symbol-table"
)

const GLOBAL_SCOPE = 0

type LlvmBackend struct {
	listener          antlr.ErrorListener
	module            *ir.Module
	moduleSymbolTable *llTable
	loopStack         *stack.Stack[*ir.Block]
	blockStack        *stack.Stack[*ir.Block]
	funcStack         *stack.Stack[*Func]
}

// VisitNegativeExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitNegativeExpression(ctx *grammar.NegativeExpressionContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()
	expr := l.Visit(ctx.Expression()).(value.Value)
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
		panic("unreachable")
	}
}

// VisitPositiveExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitPositiveExpression(ctx *grammar.PositiveExpressionContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()
	expr := l.Visit(ctx.Expression()).(value.Value)
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
	fn := l.Visit(ctx.FuncFrontDecl()).(*Func)
	isdef = false

	l.moduleSymbolTable.AddSymbol(fn.Name(), fn)

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

var basicInt *ir.Global
var basicFloat *ir.Global
var basicBool *ir.Global

func (l *LlvmBackend) addBuiltIns() {
	puts = l.module.NewFunc("puts", types.I32, ir.NewParam("", types.I8Ptr))
	putchar = l.module.NewFunc("putchar", types.I8, ir.NewParam("", types.I8))
	printf = l.module.NewFunc("printf", types.I64, ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I64))
	strcat = l.module.NewFunc("strcat", types.I8Ptr, ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I8Ptr))
	strcmp = l.module.NewFunc("strcmp", types.I64, ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I8Ptr))

	l.moduleSymbolTable.AddSymbol("printf", printf)
	l.moduleSymbolTable.AddSymbol("putchar", putchar)
	l.moduleSymbolTable.AddSymbol("puts", puts)
	l.moduleSymbolTable.AddSymbol("strcat", strcat)
	l.moduleSymbolTable.AddSymbol("strcmp", strcmp)

	basicInt = l.module.NewGlobalDef("", constant.NewCharArrayFromString("%lld\n\x00"))
	basicFloat = l.module.NewGlobalDef("", constant.NewCharArrayFromString("%lf\n\x00"))
	basicBool = l.module.NewGlobalDef("", constant.NewCharArrayFromString("%b\n\x00"))

}

func New(listener antlr.ErrorListener) *LlvmBackend {
	b := &LlvmBackend{
		listener:          listener,
		module:            ir.NewModule(),
		moduleSymbolTable: NewTable(),
		loopStack:         stack.NewStack[*ir.Block](100),
		blockStack:        stack.NewStack[*ir.Block](100),
		funcStack:         stack.NewStack[*Func](20),
	}

	b.module.TargetTriple = "x86_64-pc-linux-gnu"
	b.addBuiltIns()
	return b
}

// Visit implements grammar.MinigoVisitor.
func (l *LlvmBackend) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(l)
}

// VisitAppendCall implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitAppendCall(ctx *grammar.AppendCallContext) interface{} {
	panic("unimplemented")
}

// VisitAppendExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitAppendExpression(ctx *grammar.AppendExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitArguments implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitArguments(ctx *grammar.ArgumentsContext) interface{} {
	panic("unimplemented")
}

// VisitArrayDeclType implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitArrayDeclType(ctx *grammar.ArrayDeclTypeContext) interface{} {
	panic("unimplemented")
}

// VisitArrayType implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitArrayType(ctx *grammar.ArrayTypeContext) interface{} {
	panic("unimplemented")
}

// VisitAssignmentSimpleStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitAssignmentSimpleStatement(ctx *grammar.AssignmentSimpleStatementContext) interface{} {
	// fn, err := l.funcStack.Peek()
	// if err != nil {
	// 	panic(err)
	// }

	// // %4 = alloca i32, align 4
	// // store i32 %1, ptr %4, align 4
	// // store i32 5, ptr %4, align 4

	// current := fn.NewBlock("")
	// l.
	// current.NewAlloca()

	// return nil
	return l.Visit(ctx.AssignmentStatement())
}

// VisitBlock implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitBlock(ctx *grammar.BlockContext) interface{} {
	return l.Visit(ctx.StatementList())
}

// VisitBlockStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitBlockStatement(ctx *grammar.BlockStatementContext) interface{} {
	panic("unimplemented")
}

// VisitBooleanOperation implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitBooleanOperation(ctx *grammar.BooleanOperationContext) interface{} {
	panic("unimplemented")
}

// VisitCapCall implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitCapCall(ctx *grammar.CapCallContext) interface{} {
	panic("unimplemented")
}

// VisitCapExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitCapExpression(ctx *grammar.CapExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitCaretExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitCaretExpression(ctx *grammar.CaretExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitChildren implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitChildren(node antlr.RuleNode) interface{} {
	children := node.GetChildren()

	var result any // Always remember that this will just return the top level
	for _, child := range children {
		result = l.Visit(child.(antlr.ParseTree))
	}

	return result
}

// VisitComparison implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitComparison(ctx *grammar.ComparisonContext) interface{} {
	// LESSTHAN | GREATERTHAN | LESSTHANEQUAL | GREATERTHANEQUAL | COMPARISON | NEGATION
	fn, _ := l.blockStack.Peek()
	blk, _ := l.blockStack.Peek()

	left := l.Visit(ctx.GetLeft()).(value.Value)
	right := l.Visit(ctx.GetRight()).(value.Value)

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
	}

	panic("unreachable")
}

// VisitContinueStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitContinueStatement(ctx *grammar.ContinueStatementContext) interface{} {
	panic("unimplemented")
}

// VisitEmptyTypeDeclaration implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitEmptyTypeDeclaration(ctx *grammar.EmptyTypeDeclarationContext) interface{} {
	panic("unimplemented")
}

// VisitEmptyVariableDeclaration implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitEmptyVariableDeclaration(ctx *grammar.EmptyVariableDeclarationContext) interface{} {
	panic("unimplemented")
}

// VisitErrorNode implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("unimplemented")
}

// VisitExpressionCaseClause implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionCaseClause(ctx *grammar.ExpressionCaseClauseContext) interface{} {
	panic("unimplemented")
}

// VisitExpressionCaseClauseList implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionCaseClauseList(ctx *grammar.ExpressionCaseClauseListContext) interface{} {
	panic("unimplemented")
}

// VisitExpressionList implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionList(ctx *grammar.ExpressionListContext) interface{} {
	panic("unimplemented")
}

// VisitExpressionOperand implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionOperand(ctx *grammar.ExpressionOperandContext) interface{} {
	return l.Visit(ctx.Expression())
}

// VisitExpressionPostDec implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionPostDec(ctx *grammar.ExpressionPostDecContext) interface{} {
	blk, _ := l.blockStack.Peek()
	expr := l.Visit(ctx.Expression()).(value.Value)

	var temp value.Value
	if types.IsPointer(expr.Type()) {
		temp = blk.NewLoad(types.I64, expr)
	}

	result := blk.NewSub(temp, constant.NewInt(types.I64, 1))
	blk.NewStore(result, expr)
	return nil
}

// VisitExpressionPostInc implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitExpressionPostInc(ctx *grammar.ExpressionPostIncContext) interface{} {
	blk, _ := l.blockStack.Peek()
	expr := l.Visit(ctx.Expression()).(value.Value)

	var sex value.Value
	if types.IsPointer(expr.Type()) {
		sex = blk.NewLoad(types.I64, expr)
	}

	result := blk.NewAdd(sex, constant.NewInt(types.I64, 1))
	blk.NewStore(result, expr)
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

// VisitFuncDecl implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitFuncDecl(ctx *grammar.FuncDeclContext) interface{} {

	fn := l.Visit(ctx.FuncFrontDecl()).(*Func)
	l.moduleSymbolTable.AddSymbol(fn.Name(), fn)

	_ = l.funcStack.Push(fn)
	// l.blockStack = stack.NewStack[*ir.Block](100)
	l.blockStack.Push(fn.body)
	l.moduleSymbolTable.EnterScope()
	l.Visit(ctx.Block())

	if name := fn.Func.Name(); name != "main" {
		fn.Func.Linkage = enum.LinkagePrivate
	} else if fn.Func.Sig.RetType == types.Void && name != "main" {
		blk, _ := l.blockStack.Peek()
		blk.NewRet(nil)
	} else {
		fn.Func.Sig.RetType = types.I64
		l := len(fn.Blocks)
		if l > 0 {
			fn.Blocks[l-1].NewRet(zero)
		}
	}

	l.moduleSymbolTable.ExitScope()
	_, _ = l.blockStack.Pop()
	_, _ = l.funcStack.Pop()

	return nil
}

// OpaquePointerType is a custom type representing an opaque pointer.
type OpaquePointerType struct{}

var _ types.Type = &OpaquePointerType{}

// LLString implements types.Type.
func (t OpaquePointerType) LLString() string {
	return "ptr"
}

// Name implements types.Type.
func (t OpaquePointerType) Name() string {
	return "ptr"
}

// SetName implements types.Type.
func (t OpaquePointerType) SetName(name string) {
}

// String returns the LLVM syntax representation of the type.
func (t *OpaquePointerType) String() string {
	return "ptr"
}

// Equal reports whether t and u are of equal type.
func (t *OpaquePointerType) Equal(u types.Type) bool {
	_, ok := u.(*OpaquePointerType)
	return ok
}

// BitSize returns the size in bits of the type.
func (t *OpaquePointerType) BitSize() int64 {
	// Opaque pointers do not have a specified bit size, but we can use the typical
	// size for pointers (e.g., 64 bits on most platforms).
	return 64
}

var opaquePtr types.Type = &OpaquePointerType{}

// VisitFuncFrontDecl implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitFuncFrontDecl(ctx *grammar.FuncFrontDeclContext) interface{} {
	symbol, found := l.symbolTable.GetSymbol(ctx.IDENTIFIER().GetText())
	if !found {
		panic("unreachable")
	}

	var _type types.Type

	if symbol.Type != nil {
		_type = symbol.Type.LlvmType
	} else {
		_type = types.Void
	}

	var params []*ir.Param
	for _, argument := range symbol.Members {
		param := ir.NewParam(argument.Name, argument.Type.LlvmType)
		l.moduleSymbolTable.AddSymbol(argument.Name, param)
		params = append(params, param)
	}

	fn := l.module.NewFunc(symbol.Name, _type, params...)
	entry := fn.NewBlock("")

	return &Func{
		Func: fn,
		body: entry,
		// idents: map[string]value.Value{},
	}
}

// VisitFunctionCall implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitFunctionCall(ctx *grammar.FunctionCallContext) interface{} {
	blk, _ := l.blockStack.Peek()

	callee := l.Visit(ctx.PrimaryExpression()).(value.Named)
	var args []value.Value

	for _, expr := range ctx.Arguments().ExpressionList().AllExpression() {
		argument := l.Visit(expr).(value.Value)
		switch argument := argument.(type) {
		case *ir.InstGetElementPtr, *ir.Global:
			ptr := blk.NewGetElementPtr(types.I8, argument, zero)
			args = append(args, ptr)
		case *ir.InstAlloca:
			load := blk.NewLoad(argument.ElemType, argument)
			args = append(args, load)
		case constant.Constant, ir.Instruction: // BUG: this might cause a bug
			args = append(args, argument)
		default:
			fmt.Printf("argument: %v\n", argument)
			fmt.Printf("reflect.TypeOf(argument).String(): %v\n", reflect.TypeOf(argument).String())
			panic("unimplemented")
		}
	}

	return blk.NewCall(callee, args...)
}

// VisitIdentifierDeclType implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIdentifierDeclType(ctx *grammar.IdentifierDeclTypeContext) interface{} {
	panic("unimplemented")
}

// VisitIdentifierList implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIdentifierList(ctx *grammar.IdentifierListContext) interface{} {
	panic("unimplemented")
}

// VisitIdentifierOperand implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIdentifierOperand(ctx *grammar.IdentifierOperandContext) interface{} {
	// Check if it's one of the built in types
	name := ctx.IDENTIFIER().GetText()

	switch name {
	case "int":
		return zero
	case "float":
		return constant.NewFloat(types.Double, 0.0)
	case "rune":
		return constant.NewInt(types.I8, 0)
	case "string":
		return l.module.NewGlobalDef("", constant.NewCharArrayFromString(""))
	}

	symbol, found := l.moduleSymbolTable.GetSymbol(name)
	if !found {
		panic("unreachable")
	}

	return symbol.Symbol
}

// VisitIfElseBlock implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfElseBlock(ctx *grammar.IfElseBlockContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	expr := l.Visit(ctx.Expression()).(value.Value)
	if types.IsPointer(expr.Type()) {
		expr = blk.NewLoad(types.I1, expr)
	}

	True := fn.NewBlock("")

	l.blockStack.Push(True)
	l.Visit(ctx.GetFirstBlock())
	ifBlock, _ := l.blockStack.Pop()

	False := fn.NewBlock("")
	l.blockStack.Push(False)

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

	l.blockStack.Push(Done)

	return nil
}

// VisitIfElseIf implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfElseIf(ctx *grammar.IfElseIfContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	expr := l.Visit(ctx.Expression()).(value.Value)
	if types.IsPointer(expr.Type()) {
		expr = blk.NewLoad(types.I1, expr)
	}

	True := fn.NewBlock("")

	l.blockStack.Push(True)
	l.Visit(ctx.Block())
	ifBlock, _ := l.blockStack.Pop()

	False := fn.NewBlock("")
	l.blockStack.Push(False)
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

	l.blockStack.Push(Done)

	return nil
}

// VisitIfSimpleElseBlock implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfSimpleElseBlock(ctx *grammar.IfSimpleElseBlockContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	l.Visit(ctx.SimpleStatement()) // This will probably never return anything
	expr := l.Visit(ctx.Expression()).(value.Value)
	if types.IsPointer(expr.Type()) {
		expr = blk.NewLoad(types.I1, expr)
	}

	True := fn.NewBlock("")

	l.blockStack.Push(True)
	l.Visit(ctx.GetFirstBlock())
	ifBlock, _ := l.blockStack.Pop()

	False := fn.NewBlock("")
	l.blockStack.Push(False)
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
	l.blockStack.Push(Done)

	return nil
}

// VisitIfSimpleElseIf implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfSimpleElseIf(ctx *grammar.IfSimpleElseIfContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	l.Visit(ctx.SimpleStatement()) // This will probably never return anything
	expr := l.Visit(ctx.Expression()).(value.Value)
	if types.IsPointer(expr.Type()) {
		expr = blk.NewLoad(types.I1, expr)
	}

	True := fn.NewBlock("")
	l.blockStack.Push(True)
	l.Visit(ctx.Block())
	ifBlock, _ := l.blockStack.Pop()

	False := fn.NewBlock("")
	l.blockStack.Push(False)
	l.Visit(ctx.IfStatement())
	got, _ := l.blockStack.Peek()

	blk.NewCondBr(expr, True, False)
	if True.Term == nil {
		// True.NewBr(Done)
		True.NewBr(got)
	}

	if False.Term == nil {
		False.NewBr(got)
	}

	if ifBlock != nil && ifBlock.Term == nil {
		ifBlock.NewBr(got)
	}
	// l.blockStack.Push(Done)

	return nil
}

// VisitIfSimpleNoElse implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfSimpleNoElse(ctx *grammar.IfSimpleNoElseContext) interface{} {

	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	l.Visit(ctx.SimpleStatement()) // This will probably never return anything
	expr := l.Visit(ctx.Expression()).(value.Value)
	if types.IsPointer(expr.Type()) {
		expr = blk.NewLoad(types.I1, expr)
	}

	True := fn.NewBlock("")
	l.blockStack.Push(True)
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

	l.blockStack.Push(Done)

	return nil
}

// VisitIfSingleExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfSingleExpression(ctx *grammar.IfSingleExpressionContext) interface{} {
	fn, _ := l.funcStack.Peek()
	blk, _ := l.blockStack.Peek()

	expr := l.Visit(ctx.Expression()).(value.Value)
	if types.IsPointer(expr.Type()) {
		expr = blk.NewLoad(types.I1, expr)
	}

	then := fn.NewBlock("")
	l.blockStack.Push(then)
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

	l.blockStack.Push(end)

	return nil
}

// VisitIfStatementStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIfStatementStatement(ctx *grammar.IfStatementStatementContext) interface{} {
	l.moduleSymbolTable.EnterScope()
	defer l.moduleSymbolTable.ExitScope()
	return l.Visit(ctx.IfStatement())
}

// VisitInPlaceAssignment implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitInPlaceAssignment(ctx *grammar.InPlaceAssignmentContext) interface{} {
	panic("unimplemented")
}

// VisitIndex implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIndex(ctx *grammar.IndexContext) interface{} {
	panic("unimplemented")
}

// VisitBreakStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitBreakStatement(ctx *grammar.BreakStatementContext) interface{} {
	blk, _ := l.blockStack.Peek()
	loop, err := l.loopStack.Pop()
	if err != nil {
		panic(err)
	}
	blk.NewBr(loop)
	return nil
}

// VisitInnerTypeDecls implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitInnerTypeDecls(ctx *grammar.InnerTypeDeclsContext) interface{} {
	panic("unimplemented")
}

// VisitInnerVarDecls implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitInnerVarDecls(ctx *grammar.InnerVarDeclsContext) interface{} {
	panic("unimplemented")
}

// VisitIntLiteral implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitIntLiteral(ctx *grammar.IntLiteralContext) interface{} {
	return l.Visit(ctx.NumericLiteral())
}

func makeCstr(input string) string {
	return fmt.Sprintf("%s\x00", input)
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

// VisitLenCall implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitLenCall(ctx *grammar.LenCallContext) interface{} {
	panic("unimplemented")
}

// VisitLengthExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitLengthExpression(ctx *grammar.LengthExpressionContext) interface{} {
	panic("unimplemented")
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
	panic("unimplemented")
}

// VisitMultiTypeDeclaration implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitMultiTypeDeclaration(ctx *grammar.MultiTypeDeclarationContext) interface{} {
	panic("unimplemented")
}

// VisitMultiVariableDeclaration implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitMultiVariableDeclaration(ctx *grammar.MultiVariableDeclarationContext) interface{} {
	panic("unimplemented")
}

// VisitNestedType implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitNestedType(ctx *grammar.NestedTypeContext) interface{} {
	panic("unimplemented")
}

var zero = constant.NewInt(types.I64, 0)
var zerof = constant.NewFloat(types.Double, 0.0)

func (l *LlvmBackend) GetSymbol(name string) (value.Value, bool) {
	symbol, found := l.moduleSymbolTable.GetSymbol(name)
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

		name := ident.GetText()
		symbol, found := l.GetSymbol(name)
		if !found {
			panic("unreachable")
		}

		expr := l.Visit(rhs.Expression(idx)).(value.Value)
		switch expr := expr.(type) {
		case *ir.InstAdd, *ir.InstCall, ir.Instruction:
			blk.NewStore(expr, symbol)
		case *ir.Global:
			ptr := blk.NewGetElementPtr(types.I8, expr, zero)
			l.moduleSymbolTable.Replace(name, ptr)
		default:
			fmt.Printf("symbol: %v\n", symbol)
			fmt.Printf("expr: %v\n", expr)
			fmt.Printf("reflect.TypeOf(expr).String(): %v\n", reflect.TypeOf(expr).String())
			panic("unimplemented")
		}
	}

	return nil
}

// VisitNormalSwitch implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitNormalSwitch(ctx *grammar.NormalSwitchContext) interface{} {
	panic("unimplemented")
}

// VisitNormalSwitchExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitNormalSwitchExpression(ctx *grammar.NormalSwitchExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitNotExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitNotExpression(ctx *grammar.NotExpressionContext) interface{} {
	panic("unimplemented")
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

	leftNode := l.Visit(ctx.GetLeft()).(value.Value)
	rightNode := l.Visit(ctx.GetRight()).(value.Value)

	if types.IsPointer(leftNode.Type()) {
		leftNode = blk.NewLoad(rightNode.Type(), leftNode)
	}

	if types.IsPointer(rightNode.Type()) {
		rightNode = blk.NewLoad(leftNode.Type(), rightNode)
	}
	// left=expression (MOD) right=expression #operationPrecedence1
	switch {
	case ctx.TIMES() != nil:
		if leftNode.Type() == types.Double {
			return blk.NewFMul(leftNode, rightNode)
		}
		return blk.NewMul(leftNode, rightNode)
	case ctx.DIV() != nil:
		if leftNode.Type() == types.Double {
			return blk.NewFDiv(leftNode, rightNode)
		}
		return blk.NewSDiv(leftNode, rightNode)
	case ctx.MOD() != nil:
		if leftNode.Type() == types.Double {
			// TODO: find a way of getting the modulo return blk.NewFDiv(leftNode, rightNode)
		}
		// TODO: find a way of getting the modulo return blk.NewSDiv(leftNode, rightNode)
	default:
		panic("unreachable")
	}

	return nil
}

// VisitOperationPrecedence2 implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitOperationPrecedence2(ctx *grammar.OperationPrecedence2Context) interface{} {
	blk, err := l.blockStack.Peek()
	if err != nil {
		panic(err)
	}

	leftNode := l.Visit(ctx.GetLeft()).(value.Value)
	rightNode := l.Visit(ctx.GetRight()).(value.Value)

	if ptr, ok := leftNode.Type().(*types.PointerType); ok {
		leftNode = blk.NewLoad(ptr.ElemType, leftNode)
	}

	if ptr, ok := rightNode.Type().(*types.PointerType); ok {
		rightNode = blk.NewLoad(ptr.ElemType, rightNode)
	}

	switch {
	// LEFTSHIFT | RIGHTSHIFT | AMPERSAND | AMPERSANDCARET | PLUS | MINUS | PIPE | CARET
	case ctx.PLUS() != nil:
		if leftNode.Type() == types.Double {
			return blk.NewFAdd(leftNode, rightNode)
		}
		return blk.NewAdd(leftNode, rightNode)
	case ctx.MINUS() != nil:
		if leftNode.Type() == types.Double {
			return blk.NewFSub(leftNode, rightNode)
		}
		return blk.NewSub(leftNode, rightNode)
	default:
		panic("unimplemented")
	}
	return nil
}

// VisitPrintStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitPrintStatement(ctx *grammar.PrintStatementContext) interface{} {
	panic("TODO: fix print statement")
	blk, _ := l.blockStack.Peek()

	for _, expr := range ctx.ExpressionList().AllExpression() {
		innerString := l.Visit(expr).(value.Named)
		blk.NewCall(puts, innerString)
	}

	return nil
}

var newline = constant.NewInt(types.I8, 0x0A)

// VisitPrintlnStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitPrintlnStatement(ctx *grammar.PrintlnStatementContext) interface{} {
	blk, _ := l.blockStack.Peek()

	for _, expr := range ctx.ExpressionList().AllExpression() {
		val := l.Visit(expr).(value.Value)
		fmt.Printf("reflect.TypeOf(val).String(): %v\n", reflect.TypeOf(val).String())
		switch v := val.(type) {
		case *ir.Global:
			blk.NewCall(puts, val)
		case *ir.InstGetElementPtr:
			blk.NewCall(puts, val)
		case *ir.InstAlloca:
			switch d := v.ElemType.(type) {
			case *types.FloatType:
				blk.NewCall(printf, basicFloat, v)
			case *types.PointerType:
				switch d.ElemType {
				case types.I8:
					load := blk.NewLoad(v.Type(), v)
					blk.NewCall(puts, load)
				}
            case *types.IntType:
                load := blk.NewLoad(v.ElemType, v)
                blk.NewCall(printf, basicBool, load)
			default:
				fmt.Printf("reflect.TypeOf(v.ElemType): %v\n", reflect.TypeOf(v.ElemType))
				fmt.Printf("v.ElemType: %v\n", v.ElemType)
				panic("unimplemented")
			}
		case *ir.InstFAdd, *ir.InstFSub, *ir.InstFMul, *ir.InstFDiv:
			blk.NewCall(printf, basicFloat, val)
		case *ir.InstAdd, *ir.InstSub, *ir.InstMul, *ir.InstSDiv, *ir.InstAnd, *ir.InstXor, *ir.InstOr:
			blk.NewCall(printf, basicInt, val)
		case *constant.Int:
			blk.NewCall(printf, basicInt, val)
		case *constant.Float:
			blk.NewCall(printf, basicFloat, val)
		default:
			fmt.Printf("val: %v\n", val)
			fmt.Printf("v: %v\n", v)
			fmt.Printf("reflect.TypeOf(val).String(): %v\n", reflect.TypeOf(val).String())
			panic("unimplemented")

		}
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

	if fn.Name() == "main" {
		l.blockStack.Push(lk)
		return nblk.NewRet(zero)
	}

	if expr := ctx.Expression(); expr != nil {
		expr := l.Visit(expr).(value.Value)
		switch expr := expr.(type) {
		case *ir.Global:
			ptr := nblk.NewGetElementPtr(expr.ContentType, expr, zero)
			return nblk.NewRet(ptr)
		case *ir.InstAlloca:
			load := nblk.NewLoad(expr.ElemType, expr)
			return nblk.NewRet(load)
		case *ir.InstGetElementPtr:
			return nblk.NewRet(expr)
		case constant.Constant:
			return nblk.NewRet(expr)
		default:
			fmt.Printf("reflect.TypeOf(expr): %v\n", reflect.TypeOf(expr))
		}
	}

	return nblk.NewRet(nil)
}

var typeMap = map[string]types.Type{
	"string": types.I8Ptr,
	"int":    types.I64,
	"float":  types.Double,
	"bool":   types.I1,
	"rune":   types.I8,
}

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
	panic("unimplemented")
}

// VisitSimpleStatementStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSimpleStatementStatement(ctx *grammar.SimpleStatementStatementContext) interface{} {
	return l.Visit(ctx.SimpleStatement())
}

// VisitSimpleStatementSwitch implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSimpleStatementSwitch(ctx *grammar.SimpleStatementSwitchContext) interface{} {
	panic("unimplemented")
}

// VisitSimpleStatementSwitchExpression implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSimpleStatementSwitchExpression(ctx *grammar.SimpleStatementSwitchExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitSingleTypeDecl implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSingleTypeDecl(ctx *grammar.SingleTypeDeclContext) interface{} {
	panic("unimplemented")
}

// VisitSingleVarDeclNoExps implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSingleVarDeclNoExps(ctx *grammar.SingleVarDeclNoExpsContext) interface{} {

	// blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	expr := l.Visit(ctx.DeclType()).(types.Type)

	for _, ident := range ctx.IdentifierList().AllIDENTIFIER() {
		name := ident.GetText()
		if fn.body != nil {
			alloca := fn.body.NewAlloca(expr)
			l.moduleSymbolTable.AddSymbol(name, alloca)
		} else {
			def := l.module.NewGlobalDef("", constant.NewZeroInitializer(expr))
			l.moduleSymbolTable.Replace(name, def)
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
	_type := l.Visit(ctx.DeclType()).(types.Type)
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
	panic("unimplemented")
}

// VisitStructMemDecls implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitStructMemDecls(ctx *grammar.StructMemDeclsContext) interface{} {
	panic("unimplemented")
}

// VisitStructType implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitStructType(ctx *grammar.StructTypeContext) interface{} {
	panic("unimplemented")
}

// VisitSubIndex implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSubIndex(ctx *grammar.SubIndexContext) interface{} {
	blk, _ := l.blockStack.Peek()

	expr := l.Visit(ctx.PrimaryExpression()).(value.Value)
	index := ctx.Index()
	idx := l.Visit(index).(value.Value)

	exprt := expr.Type().(*types.PointerType)

	if id, ok := idx.Type().(*types.PointerType); ok {
		idx = blk.NewLoad(id.ElemType, idx)
	}

	gep := blk.NewGetElementPtr(exprt.ElemType, expr, zero, idx)
	return gep
}

// VisitSwitchCaseBranch implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSwitchCaseBranch(ctx *grammar.SwitchCaseBranchContext) interface{} {
	panic("unimplemented")
}

// VisitSwitchDefaultBranch implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSwitchDefaultBranch(ctx *grammar.SwitchDefaultBranchContext) interface{} {
	panic("unimplemented")
}

// VisitSwitchStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitSwitchStatement(ctx *grammar.SwitchStatementContext) interface{} {
	panic("unimplemented")
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

	l.blockStack.Push(_for)
	l.loopStack.Push(end)

	l.Visit(ctx.Block())
	l.Visit(ctx.GetLast())

	l.loopStack.Pop()

	expr := l.Visit(ctx.Expression()).(value.Value)

	got, _ := l.blockStack.Peek()
	if got.Term == nil {
		got.NewCondBr(expr, _for, end)
	}

	end.Parent = fn.Func
	fn.Blocks = append(fn.Blocks, end)

	l.blockStack.Push(end)

	return nil
}

// VisitTopDeclarationList implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitTopDeclarationList(ctx *grammar.TopDeclarationListContext) interface{} {
	return l.VisitChildren(ctx)
}

// VisitTypeDeclStatement implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitTypeDeclStatement(ctx *grammar.TypeDeclStatementContext) interface{} {
	panic("unimplemented")
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
		expr := l.Visit(ctx.ExpressionList().Expression(idx)).(value.Value)

		switch argument := expr.(type) {
		case *ir.InstAlloca:
			load := blk.NewLoad(argument.ElemType, expr)
			alloca := fn.body.NewAlloca(argument.ElemType)
			blk.NewStore(load, alloca)
			l.moduleSymbolTable.AddSymbol(name, alloca)
		case *ir.Global: // XXX: This might cause a bug somewhere else, I still have to run the checks
			switch arg := argument.ContentType.(type) {
			case *types.ArrayType:
				ptr := blk.NewGetElementPtr(arg.ElemType, argument, zero)
				l.moduleSymbolTable.AddSymbol(name, ptr)
			default:
				alloca := fn.body.NewAlloca(argument.Type())
				load := blk.NewLoad(argument.Type(), argument)
				blk.NewStore(load, alloca)
				l.moduleSymbolTable.AddSymbol(name, alloca)
			}
		case *ir.InstGetElementPtr:
			if l.moduleSymbolTable.currentScope == GLOBAL_SCOPE {
				l.moduleSymbolTable.AddSymbol(name, argument)
			} else {
				ptr := blk.NewGetElementPtr(types.I8, argument, zero)
				l.moduleSymbolTable.AddSymbol(name, ptr)
			}
		case constant.Constant:
			if blk != nil {
				alloca := fn.body.NewAlloca(expr.Type())
				blk.NewStore(expr, alloca)
				l.moduleSymbolTable.AddSymbol(name, alloca)
			} else {
				def := l.module.NewGlobalDef("", argument)
				l.moduleSymbolTable.AddSymbol(name, def)
			}
		case ir.Instruction:
			alloca := fn.body.NewAlloca(expr.Type())
			blk.NewStore(expr, alloca)
			l.moduleSymbolTable.AddSymbol(name, alloca)
		default:
			fmt.Printf("argument: %v\n", argument)
			fmt.Printf("reflect.TypeOf(argument).String(): %v\n", reflect.TypeOf(argument).String())
			panic("unimplemented")
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
		expr := l.Visit(ctx.ExpressionList().Expression(idx)).(value.Value)

		switch argument := expr.(type) {
		case *ir.InstAlloca:
			load := blk.NewLoad(argument.ElemType, expr)
			alloca := fn.body.NewAlloca(argument.ElemType)
			blk.NewStore(load, alloca)
			l.moduleSymbolTable.AddSymbol(name, alloca)
		case *ir.Global: // XXX: This might cause a bug somewhere else, I still have to run the checks
			switch arg := argument.ContentType.(type) {
			case *types.ArrayType:
				ptr := blk.NewGetElementPtr(arg.ElemType, argument, zero)
				l.moduleSymbolTable.AddSymbol(name, ptr)
			default:
				alloca := fn.body.NewAlloca(argument.Type())
				load := blk.NewLoad(argument.Type(), argument)
				blk.NewStore(load, alloca)
				l.moduleSymbolTable.AddSymbol(name, alloca)
			}
		case *ir.InstGetElementPtr:
			if l.moduleSymbolTable.currentScope == GLOBAL_SCOPE {
				l.moduleSymbolTable.AddSymbol(name, argument)
			} else {
				ptr := blk.NewGetElementPtr(types.I8, argument, zero)
				l.moduleSymbolTable.AddSymbol(name, ptr)
			}
		case constant.Constant:
			if blk != nil {
				alloca := fn.body.NewAlloca(expr.Type())
				blk.NewStore(expr, alloca)
				l.moduleSymbolTable.AddSymbol(name, alloca)
			} else {
				def := l.module.NewGlobalDef("", argument)
				l.moduleSymbolTable.AddSymbol(name, def)
			}
		case ir.Instruction:
			alloca := fn.body.NewAlloca(expr.Type())
			blk.NewStore(expr, alloca)
			l.moduleSymbolTable.AddSymbol(name, alloca)
		default:
			fmt.Printf("argument: %v\n", argument)
			fmt.Printf("reflect.TypeOf(argument).String(): %v\n", reflect.TypeOf(argument).String())
			panic("unimplemented")
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

		expr := l.Visit(ctx.GetRight().Expression(idx)).(value.Value)

		fmt.Printf("reflect.TypeOf(expr): %v\n", reflect.TypeOf(expr))
		switch argument := expr.(type) {
		case *ir.InstAlloca:
			load := blk.NewLoad(argument.ElemType, expr)
			alloca := fn.body.NewAlloca(argument.ElemType)
			blk.NewStore(load, alloca)
			l.moduleSymbolTable.AddSymbol(name, alloca)
		case *ir.Global: // XXX: This might cause a bug somewhere else, I still have to run the checks
			switch arg := argument.ContentType.(type) {
			case *types.ArrayType:
				ptr := blk.NewGetElementPtr(arg.ElemType, argument, zero)
				l.moduleSymbolTable.AddSymbol(name, ptr)
			default:
				alloca := fn.body.NewAlloca(argument.Type())
				load := blk.NewLoad(argument.Type(), argument)
				blk.NewStore(load, alloca)
				l.moduleSymbolTable.AddSymbol(name, alloca)
			}
		case *ir.InstGetElementPtr:
			if l.moduleSymbolTable.currentScope == GLOBAL_SCOPE {
				l.moduleSymbolTable.AddSymbol(name, argument)
			} else {
				switch e := argument.ElemType.(type) {
				case *types.ArrayType:
					ptr := blk.NewGetElementPtr(argument.ElemType, argument, zero)
					alloca := fn.body.NewAlloca(e.ElemType)
					load := blk.NewLoad(e.ElemType, ptr)
					blk.NewStore(load, alloca)
					l.moduleSymbolTable.AddSymbol(name, alloca)
				case *types.IntType:
					alloca := fn.body.NewAlloca(expr.Type())
					blk.NewStore(expr, alloca)
					l.moduleSymbolTable.AddSymbol(name, alloca)
				default:
					fmt.Printf("reflect.TypeOf(argument.ElemType): %v\n", reflect.TypeOf(argument.ElemType))
					panic("unimplemented")
				}
			}
		case constant.Constant:
			alloca := fn.body.NewAlloca(expr.Type())
			blk.NewStore(expr, alloca)
			l.moduleSymbolTable.AddSymbol(name, alloca)
		case ir.Instruction:
			alloca := fn.body.NewAlloca(expr.Type())
			blk.NewStore(expr, alloca)
			l.moduleSymbolTable.AddSymbol(name, alloca)
		default:
			fmt.Printf("expr: %v\n", argument)
			fmt.Printf("reflect.TypeOf(expr).String(): %v\n", reflect.TypeOf(argument).String())
			panic("unimplemented")
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

	l.blockStack.Push(_for)
	l.loopStack.Push(end)

	l.Visit(ctx.Block())

	l.loopStack.Pop()

	expr := l.Visit(ctx.Expression()).(value.Value)

	got, _ := l.blockStack.Peek()
	if got.Term == nil {
		got.NewCondBr(expr, _for, end)
	}

	end.Parent = fn.Func
	fn.Blocks = append(fn.Blocks, end)

	l.blockStack.Push(end)

	return nil
}

// VisitInfiniteFor implements grammar.MinigoVisitor.
func (l *LlvmBackend) VisitInfiniteFor(ctx *grammar.InfiniteForContext) interface{} {
	blk, _ := l.blockStack.Peek()
	fn, _ := l.funcStack.Peek()

	_for := fn.NewBlock("")
	end := ir.NewBlock("")

	blk.NewBr(_for)

	l.blockStack.Push(_for)
	l.loopStack.Push(end)

	l.Visit(ctx.Block())

	l.loopStack.Pop()

	got, _ := l.blockStack.Peek()
	if got.Term == nil {
		got.NewBr(_for)
	}

	end.Parent = fn.Func
	fn.Blocks = append(fn.Blocks, end)

	l.blockStack.Push(end)

	return nil
}

var _ grammar.MinigoVisitor = &LlvmBackend{}
