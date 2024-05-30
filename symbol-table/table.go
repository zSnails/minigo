package symboltable

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/zSnails/minigo/linked-list/double"
)

type SymbolTable struct {
	Symbols *double.LinkedList[*Symbol]
	Scope   int8
}

func (t SymbolTable) String() string {
	sb := strings.Builder{}
	fmt.Fprint(&sb, "[")
	current := t.Symbols.Head
	for current != nil {
		fmt.Fprintf(&sb, "%v", current.Value)
		if current.Next != nil {
			fmt.Fprint(&sb, ", ")
		}
		current = current.Next
	}
	fmt.Fprint(&sb, "]")
	return sb.String()
}

type SymbolType int

const (
	TypeSymbol SymbolType = 1 << iota
	PrimitiveTypeSymbol
	VariableSymbol
	FunctionSymbol
	SliceSymbol
	ArraySymbol
	StructSymbol
	ConstantSymbol
)

func (s SymbolType) String() string {
	sb := strings.Builder{}
	var bit SymbolType
	cp := s
	for bit = 0; bit < 32; bit++ {
		curr := cp >> bit
		if (curr & 1) == 0 {
			continue
		}
		var k SymbolType = 1 << bit
		fmt.Fprintf(&sb, "%s ", symbolTypeTable[k])
	}
	return sb.String()
}

var symbolTypeTable = map[SymbolType]string{
	TypeSymbol:          "Type",
	PrimitiveTypeSymbol: "Primitive",
	VariableSymbol:      "Variable",
	FunctionSymbol:      "Function",
	SliceSymbol:         "Slice",
	ArraySymbol:         "Array",
	StructSymbol:        "Struct",
}

type Symbol struct {
	Scope int8
	Name  string
	Value value.Value
}

func (s *SymbolTable) EnterScope() {
	s.Scope++
}

func (t *SymbolTable) AddSymbol(name string, value value.Value) error {
	_, added := t.Symbols.FindFirst(func(s *Symbol) bool {
		return (s.Scope == GLOBAL_SCOPE || s.Scope == t.Scope) && s.Name == name
	})

	if added {
		return fmt.Errorf("symbol '%s' already defined in the current scope", name)
	}

	t.Symbols.Add(&Symbol{
		Scope: t.Scope,
		Name:  name,
		Value: value,
	})

	return nil
}

func (t *SymbolTable) GetSymbol(name string) (*Symbol, bool) {
	val, found := t.Symbols.FindFirst(func(s *Symbol) bool {
		return s.Name == name
	})

	if !found {
		return nil, found
	}

	return val, found
}

func (t *SymbolTable) ExitScope() {
	t.Symbols.RemoveIf(func(s *Symbol) bool {
		return s.Scope == t.Scope
	})
	t.Scope--
}

const GLOBAL_SCOPE = -1

func (t *SymbolTable) addPrimitives() {
	t.AddSymbol("true", constant.True)
	t.AddSymbol("false", constant.False)

	printf := ir.NewFunc("printf", types.I32, ir.NewParam("", types.I8Ptr))
	printf.Sig.Variadic = true
	_ = t.AddSymbol("printf", printf)
}

func NewSymbolTable() *SymbolTable {
	t := &SymbolTable{
		Symbols: double.NewLinkedList[*Symbol](),
		Scope:   0,
	}
	t.addPrimitives()
	return t
}
