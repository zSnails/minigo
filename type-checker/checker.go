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

// Builtin symbols
// const (
// 	Int symboltable.Type = iota + 1
// 	Int8
// 	Int16
// 	Int32
// 	Int64
// 	Uint
// 	Uint8
// 	Uint16
// 	Uint32
// 	Uint64

// 	Float32
// 	Float64

// 	Complex64
// 	Complex128

// 	String
// 	Rune
// )

type Primitive struct {
	Name    string
	Members []*symboltable.Type
}

// GetMembers implements symboltable.Type.
func (p *Primitive) GetMembers() []*symboltable.Symbol {
	return p.Members
}

// GetName implements symboltable.Type.
func (p *Primitive) GetName() string {
	return p.Name
}

// HasMembers implements symboltable.Type.
func (p *Primitive) HasMembers() bool {
	return len(p.Members) > 0
}

// IsPrimitive implements symboltable.Type.
func (p *Primitive) IsPrimitive() bool {
	return true
}

func NewPrimitive(name string) symboltable.Type {
	return &Primitive{
		Name:    name,
		Members: nil,
	}
}

// Dynamic symbol table

var TypeTable = map[string]symboltable.Type{
	"int":        NewPrimitive("int"),
	"int8":       NewPrimitive("int8"),
	"int16":      NewPrimitive("int16"),
	"int32":      NewPrimitive("int32"),
	"int64":      NewPrimitive("int64"),
	"uint":       NewPrimitive("uint"),
	"uint8":      NewPrimitive("uint8"),
	"uint16":     NewPrimitive("uint16"),
	"uint32":     NewPrimitive("uint32"),
	"uint64":     NewPrimitive("uint64"),
	"float32":    NewPrimitive("float32"),
	"float64":    NewPrimitive("float64"),
	"complex64":  NewPrimitive("complex64"),
	"complex128": NewPrimitive("complex128"),
	"string":     NewPrimitive("string"),
	"rune":       NewPrimitive("rune"),
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
	a := ctx.InnerTypeDecls()
	for _, decl := range a.AllSingleTypeDecl() {
		name := decl.IDENTIFIER()
		typeName := decl.DeclType().IDENTIFIER()
		err := t.SymbolTable.AddSymbol(name.GetText(), TypeTable[typeName.GetText()])
		if err != nil {
			t.errors = append(t.errors, t.MakeError(name.GetSymbol(), err))
		}
	}
	return nil
}

func (t *TypeChecker) MakeError(token antlr.Token, err error) error {
	return fmt.Errorf("%s:%d:%d: %w", t.filename, token.GetLine(), token.GetTokenSource().GetCharPositionInLine(), err)
}

// VisitSingleVarDeclsNoExpsDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSingleVarDeclsNoExpsDecl(ctx *grammar.SingleVarDeclsNoExpsDeclContext) interface{} {
	panic("unimplemented")
}

// VisitTypeDeclaration implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitTypeDeclaration(ctx *grammar.TypeDeclarationContext) interface{} {
	name := ctx.SingleTypeDecl().IDENTIFIER()
	// TODO: this doesn't work for struct types or slice/array types
	declType := ctx.SingleTypeDecl().DeclType()
	if typeName := declType.IDENTIFIER(); typeName != nil {
		err := t.SymbolTable.AddSymbol(name.GetText(), TypeTable[typeName.GetText()])
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

// VisitTypedVarDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitTypedVarDecl(ctx *grammar.TypedVarDeclContext) interface{} {
	panic("unimplemented")
}

// VisitUntypedVarDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitUntypedVarDecl(ctx *grammar.UntypedVarDeclContext) interface{} {
	panic("unimplemented")
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

// VisitAssignmentStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitAssignmentStatement(ctx *grammar.AssignmentStatementContext) interface{} {
	panic("unimplemented")
}

// VisitBlock implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitBlock(ctx *grammar.BlockContext) interface{} {
	t.SymbolTable.EnterScope()
	res := t.VisitChildren(ctx)
	t.SymbolTable.ExitScope()
	return res
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
	panic("unimplemented")
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
	panic("unimplemented")
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
	panic("unimplemented")
}

// VisitFuncDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFuncDecl(ctx *grammar.FuncDeclContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitFuncFrontDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitFuncFrontDecl(ctx *grammar.FuncFrontDeclContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitIdentifierList implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitIdentifierList(ctx *grammar.IdentifierListContext) interface{} {
	panic("unimplemented")
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
	panic("unimplemented")
}

// VisitInnerVarDecls implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitInnerVarDecls(ctx *grammar.InnerVarDeclsContext) interface{} {
	panic("unimplemented")
}

// VisitLengthExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitLengthExpression(ctx *grammar.LengthExpressionContext) interface{} {
	panic("unimplemented")
}

// VisitLiteral implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitLiteral(ctx *grammar.LiteralContext) interface{} {
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
	panic("unimplemented")
}

// VisitOperand implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitOperand(ctx *grammar.OperandContext) interface{} {
	panic("unimplemented")
}

// VisitPrimaryExpression implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitPrimaryExpression(ctx *grammar.PrimaryExpressionContext) interface{} {
	panic("unimplemented")
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
	panic("unimplemented")
}

// VisitSimpleStatementStatement implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSimpleStatementStatement(ctx *grammar.SimpleStatementStatementContext) interface{} {
	panic("unimplemented")
}

// VisitSingleTypeDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSingleTypeDecl(ctx *grammar.SingleTypeDeclContext) interface{} {
	panic("unimplemented")
}

// VisitSingleVarDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSingleVarDecl(ctx *grammar.SingleVarDeclContext) interface{} {
	panic("unimplemented")
}

// VisitSingleVarDeclNoExps implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitSingleVarDeclNoExps(ctx *grammar.SingleVarDeclNoExpsContext) interface{} {
	panic("unimplemented")
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
	return nil
}

// VisitTopDeclarationList implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitTopDeclarationList(ctx *grammar.TopDeclarationListContext) interface{} {
	return t.VisitChildren(ctx)
}

// VisitTypeDecl implements grammar.MinigoVisitor.
func (t *TypeChecker) VisitTypeDecl(ctx *grammar.TypeDeclContext) interface{} {
	panic("unimplemented")
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
