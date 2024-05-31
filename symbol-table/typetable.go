package symboltable

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/zSnails/minigo/linked-list/double"
)

type TypeTable struct {
	Symbols *double.LinkedList[*Type]
	Scope   int8
}

type Type struct {
	Type types.Type
	Name string
}

func (t *TypeTable) AddSymbol(name string, value types.Type) error {
	_, added := t.Symbols.FindFirst(func(s *Type) bool {
		return name == s.Name
	})

	if added {
		return fmt.Errorf("type '%s' already defined in the current scope", value.Name())
	}

	t.Symbols.Add(&Type{
		Type: value,
		Name: name,
	})

	return nil
}

func (t *TypeTable) GetSymbol(name string) (*Type, bool) {
	val, found := t.Symbols.FindFirst(func(s *Type) bool {
		return s.Name == name
	})

	if !found {
		return nil, found
	}

	return val, found
}

func (t *TypeTable) addBuiltins() {
	// types.I1.SetName("bool")
	// types.I64.SetName("int")
	// types.Double.SetName("float")
	// types.I8Ptr.SetName("string")
	// types.I8.SetName("rune")

	t.AddSymbol("int", types.I64)
	t.AddSymbol("bool", types.I1)
	t.AddSymbol("float", types.Double)
	t.AddSymbol("string", types.I8Ptr)
	t.AddSymbol("rune", types.I8)
}

func NewTypeTable() *TypeTable {
	t := &TypeTable{
		Symbols: double.NewLinkedList[*Type](),
		Scope:   0,
	}

	t.addBuiltins()

	return t
}
