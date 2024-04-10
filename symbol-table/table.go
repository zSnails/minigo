package symboltable

import (
	"fmt"
	"strings"

	"github.com/zSnails/minigo/linked-list/single"
)

type SymbolTable struct {
	Symbols *single.LinkedList[*Symbol]
	Scope   uint8
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

type Type interface {
	GetName() string
	IsPrimitive() bool
	HasMembers() bool
	GetMembers() []*Symbol
}

type Symbol struct {
	Scope   uint8
	Name    string
	IsSlice bool
	IsArray bool
	Type    Type // TODO: find a way of doing this for complex types such a structs
}

func (s Symbol) String() string {
	return fmt.Sprintf("Symbol(%d, %s, %t, %t, %d)", s.Scope, s.Name, s.IsSlice, s.IsArray, s.Type)
}

func (s *SymbolTable) EnterScope() {
	s.Scope++
}

func (t *SymbolTable) AddSymbol(name string, _type Type) error {
	_, added := t.Symbols.FindFirst(func(s *Symbol) bool {
		return s.Name == name && s.Scope == t.Scope
	})
	if added {
		return fmt.Errorf("symbol %s already defined in the current scope", name)
	}
	t.Symbols.Add(&Symbol{Name: name, Type: _type, Scope: t.Scope})
	return nil
}

func (t *SymbolTable) GetSymbol(name string) (*Symbol, bool) {
	return t.Symbols.FindFirst(func(s *Symbol) bool {
		return s.Name == name
	})
}

func (t *SymbolTable) ExitScope() {
	t.Symbols.RemoveIf(func(s *Symbol) bool {
		return s.Scope == t.Scope
	})
	t.Scope--
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		Symbols: single.NewLinkedList[*Symbol](),
		Scope:   0,
	}
}
