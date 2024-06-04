package internal

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/zSnails/minigo/linked-list/single"
)

type SymbolTable struct {
	Symbols *single.LinkedList[*Symbol]
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
	_ = t.AddSymbol("true", constant.True)
	_ = t.AddSymbol("false", constant.False)

	printf := ir.NewFunc("printf", types.I32, ir.NewParam("", types.I8Ptr))
	printf.Sig.Variadic = true
	putchar := ir.NewFunc("putchar", types.Void, ir.NewParam("", types.I64))

	_ = t.AddSymbol("putchar", putchar)
	_ = t.AddSymbol("printf", printf)
}

func NewSymbolTable() *SymbolTable {
	t := &SymbolTable{
		Symbols: single.NewLinkedList[*Symbol](),
		Scope:   0,
	}
	t.addPrimitives()
	return t
}
