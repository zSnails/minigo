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

type SymbolType int

const (
	TypeSymbol SymbolType = iota + 1
	VariableSymbol
	FunctionSymbol
	SliceSymbol
	ArraySymbol
	StructSymbol
)

type Symbol struct {
	SymbolType SymbolType
	Scope      uint8
	Name       string
	Type       *Symbol
	Members    []*Symbol
}

func (s Symbol) String() string {
	return fmt.Sprintf("Symbol(%d, %s, %s)", s.Scope, s.Name, s.Type)
}

func (s *SymbolTable) EnterScope() {
	s.Scope++
}

func (t *SymbolTable) AddSymbol(symbol *Symbol) error {
	_, added := t.Symbols.FindFirst(func(s *Symbol) bool {
		return s.Name == symbol.Name && s.Scope == t.Scope
	})
	if added {
		switch symbol.SymbolType {
		case FunctionSymbol: // TODO:
			return fmt.Errorf("function %s already defined in the current scope", symbol.Name)
		case TypeSymbol:
			return fmt.Errorf("type %s already defined in the current scope", symbol.Name)
		default:
			return fmt.Errorf("symbol %s already defined in the current scope", symbol.Name)
		}
	}
	t.Symbols.Add(symbol)
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

func NewPrimitive(name string) *Symbol {
	return &Symbol{
		SymbolType: TypeSymbol,
		Scope:      0, // This is the global scope
		Name:       name,
		Type:       nil,
		Members:    nil,
	}
}

func (t *SymbolTable) NewVariable(name string, _type *Symbol) *Symbol {
	return &Symbol{
		SymbolType: VariableSymbol,
		Scope:      t.Scope,
		Name:       name,
		Type:       _type,
		Members:    []*Symbol{},
	}
}

func (t *SymbolTable) NewFunction(name string, _type *Symbol, members ...*Symbol) *Symbol {
	return &Symbol{
		SymbolType: FunctionSymbol,
		Scope:      t.Scope,
		Name:       name,
		Type:       _type,
		Members:    []*Symbol{},
	}
}

func (t *SymbolTable) NewType(name string, typename string) (*Symbol, error) {
	_type, found := t.Symbols.FindFirst(func(s *Symbol) bool {
		return s.Name == typename
	})
	if !found {
		return nil, fmt.Errorf("type %s does not exist", typename)
	}
	return &Symbol{
		SymbolType: TypeSymbol,
		Scope:      0, // Types are always global, so the 0 scope is used
		Name:       name,
		Type:       _type,
		Members:    nil, // TODO: this should be defined for struct types, I'll probably just create another helper method for that
	}, nil
}

var (
	Int        = NewPrimitive("int")
	Int8       = NewPrimitive("int8")
	Int16      = NewPrimitive("int16")
	Int32      = NewPrimitive("int32")
	Int64      = NewPrimitive("int64")
	Uint       = NewPrimitive("uint")
	Uint8      = NewPrimitive("uint8")
	Uint16     = NewPrimitive("uint16")
	Uint32     = NewPrimitive("uint32")
	Uint64     = NewPrimitive("uint64")
	Float32    = NewPrimitive("float32")
	Float64    = NewPrimitive("float64")
	Complex64  = NewPrimitive("complex64")
	Complex128 = NewPrimitive("complex128")
	String     = NewPrimitive("string")
	Rune       = NewPrimitive("rune")
)

func (t *SymbolTable) addPrimitives() {
	_ = t.AddSymbol(Int)
	_ = t.AddSymbol(Int8)
	_ = t.AddSymbol(Int16)
	_ = t.AddSymbol(Int32)
	_ = t.AddSymbol(Int64)
	_ = t.AddSymbol(Uint)
	_ = t.AddSymbol(Uint8)
	_ = t.AddSymbol(Uint16)
	_ = t.AddSymbol(Uint32)
	_ = t.AddSymbol(Uint64)
	_ = t.AddSymbol(Float32)
	_ = t.AddSymbol(Float64)
	_ = t.AddSymbol(Complex64)
	_ = t.AddSymbol(Complex128)
	_ = t.AddSymbol(String)
	_ = t.AddSymbol(Rune)
}

func NewSymbolTable() *SymbolTable {
	t := &SymbolTable{
		Symbols: single.NewLinkedList[*Symbol](),
		Scope:   0,
	}
	t.addPrimitives()
	return t
}
