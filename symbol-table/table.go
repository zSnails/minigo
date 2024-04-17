package symboltable

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
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
	TypeSymbol SymbolType = 2 << iota
	VariableSymbol
	FunctionSymbol
	SliceSymbol
	ArraySymbol
	StructSymbol
)

type Symbol struct {
	SymbolType SymbolType
	Token      antlr.Token
	IsSlice    bool
	Size       uint64 //  this field will only be used if the symbol is an array
	Scope      uint8
	Name       string
	Type       *Symbol
	Members    []*Symbol
}

func (s Symbol) String() string {
	if s.SymbolType&FunctionSymbol != 0 {
		return fmt.Sprintf("func %s%s -> %s", s.Name, s.Members, s.Type)
	} else if s.SymbolType&VariableSymbol != 0 {
		if s.IsSlice {
			return fmt.Sprintf("Slice[%d, %s, %s]", s.Scope, s.Name, s.Type)
		}
		return fmt.Sprintf("Variable(%d, %s, %s)", s.Scope, s.Name, s.Type)
	} else if s.SymbolType&SliceSymbol != 0 {
		return fmt.Sprintf("Slice[%d, %s, %s]", s.Scope, s.Name, s.Type)
	} else if s.SymbolType&ArraySymbol != 0 {
		return fmt.Sprintf("Array(%d)[%d, %s, %s]", s.Size, s.Scope, s.Name, s.Type)
	} else if s.SymbolType&StructSymbol != 0 {
		return fmt.Sprintf("Struct(%d, %s, %d members (optimized out))", s.Scope, s.Name, len(s.Members))
	} else if s.SymbolType&TypeSymbol != 0 {
		if s.Type == nil {
			return fmt.Sprintf("Type(%d, %s)", s.Scope, s.Name)
		}
		return fmt.Sprintf("Type(%d, %s, %s)", s.Scope, s.Name, s.Type)
	}
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

func (t *SymbolTable) NewVariable(token antlr.Token, name string, _type *Symbol) *Symbol {
	return &Symbol{
		SymbolType: VariableSymbol,
		Token:      token,
		IsSlice:    false,
		Size:       0,
		Scope:      t.Scope,
		Name:       name,
		Type:       _type,
		Members:    []*Symbol{},
	}
}

func (t *SymbolTable) NewFunction(token antlr.Token, name string, _type *Symbol, members ...*Symbol) *Symbol {
	return &Symbol{
		SymbolType: FunctionSymbol,
		Token:      token,
		IsSlice:    false,
		Size:       0,
		Scope:      t.Scope,
		Name:       name,
		Type:       _type,
		Members:    members,
	}
}

func (t *SymbolTable) NewStructType(token antlr.Token, name string, members ...*Symbol) *Symbol {
	return &Symbol{
		SymbolType: TypeSymbol | StructSymbol,
		Token:      token,
		IsSlice:    false,
		Size:       0,
		Scope:      t.Scope,
		Name:       name,
		Type:       nil,
		Members:    members,
	}
}

func (t *SymbolTable) NewAliasType(token antlr.Token, name string, _type *Symbol) *Symbol {
	return &Symbol{
		SymbolType: TypeSymbol,
		Token:      token,
		IsSlice:    false,
		Size:       0,
		Scope:      t.Scope,
		Name:       name,
		Type:       _type,
		Members:    nil,
	}
}
func (t *SymbolTable) NewArrayType(token antlr.Token, name string, _type *Symbol, size uint64) *Symbol {
	return &Symbol{
		SymbolType: ArraySymbol,
		Size:       size,
		Scope:      t.Scope,
		Name:       name,
		Type:       _type,
		Members:    nil,
	}
}

func (t *SymbolTable) NewSliceType(token antlr.Token, name string, _type *Symbol) *Symbol {
	return &Symbol{
		SymbolType: SliceSymbol,
		Token:      token,
		IsSlice:    false,
		Size:       0,
		Scope:      t.Scope,
		Name:       name,
		Type:       _type,
		Members:    nil,
	}
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
