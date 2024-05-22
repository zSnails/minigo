package symboltable

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir/types"
	"github.com/zSnails/minigo/linked-list/double"
)

type SymbolTable struct {
	Symbols *double.LinkedList[*Symbol]
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
	SymbolType SymbolType
	Token      antlr.Token
	Size       uint64 //  this field will only be used if the symbol is an array
	Scope      uint8
	Variadic   bool
	Name       string
	Type       *Symbol
	LlvmType   types.Type
	Members    []*Symbol
}

func (s *Symbol) Is(against SymbolType) bool {
	return (s.SymbolType&against != 0)
}

// func (s *Symbol) Equals(other *Symbol) bool {
// 	return (s.SymbolType&TypeSymbol != 0) && s.Name == other.Name
// }

func (s *Symbol) SameType(other *Symbol) bool {
	return s.Name == other.Name && s.SymbolType == other.SymbolType
}

// func (s *Symbol) SameType(other *Symbol) bool {
// 	return s.SymbolType == other.SymbolType
// }

// func (s *Symbol) Equals(other *Symbol) bool {
// 	return other != nil && s.Type == other.Type && s.IsSlice == other.IsSlice && s.IsArray == other.IsArray && s.Name == other.Name
// }

func (s *Symbol) String() string {
	if s == nil {
		return "no value"
	}
	if (s.SymbolType)&ArraySymbol != 0 {
		if s.Type != nil {
			return fmt.Sprintf("[%d]%s", s.Size, s.Type.String())
		}
		return fmt.Sprintf("[%d]%s", s.Size, s.Name)
	} else if s.SymbolType&SliceSymbol != 0 {
		if s.Type != nil {
			return fmt.Sprintf("[]%s", s.Type.String())
		}
		return fmt.Sprintf("[]%s", s.Name)
	}

	return s.Name
}

func (s Symbol) Repr() string {
	if s.SymbolType&FunctionSymbol != 0 {
		return fmt.Sprintf("func %s%s -> %s", s.Name, s.Members, s.Type)
	} else if s.SymbolType&SliceSymbol != 0 {
		return fmt.Sprintf("[]%s", s.Type)
	} else if s.SymbolType&ArraySymbol != 0 {
		return fmt.Sprintf("[%d]%s", s.Size, s.Name)
	} else if s.SymbolType&VariableSymbol != 0 {
		return fmt.Sprintf("Variable(%d, %s, %s)", s.Scope, s.Name, s.Type)
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
		case FunctionSymbol:
			return fmt.Errorf("function '%s' already defined in the current scope", symbol.Name)
		case TypeSymbol:
			return fmt.Errorf("type '%s' already defined in the current scope", symbol.Name)
		default:
			return fmt.Errorf("symbol '%s' already defined in the current scope", symbol.Name)
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

func NewConstant(name string, _type *Symbol, llvmtype types.Type) *Symbol {
	return &Symbol{
		SymbolType: ConstantSymbol,
		Token:      nil,
		Scope:      0,
		Name:       name,
		Type:       _type,
		LlvmType:   llvmtype,
		Members:    nil,
	}
}

func NewPrimitive(name string, llvmtype types.Type) *Symbol {
	return &Symbol{
		SymbolType: TypeSymbol | PrimitiveTypeSymbol,
		Scope:      0, // This is the global scope
		Name:       name,
		LlvmType:   llvmtype,
		Type:       nil,
		Members:    nil,
	}
}

func (t *SymbolTable) NewVariable(token antlr.Token, name string, _type *Symbol) *Symbol {
	return &Symbol{
		SymbolType: VariableSymbol,
		Token:      token,
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
		Size:       0,
		Scope:      t.Scope,
		Name:       name,
		Type:       _type,
		Members:    nil,
	}
}
func (t *SymbolTable) NewArrayType(token antlr.Token, name string, _type *Symbol, size uint64) *Symbol {
	return &Symbol{
		SymbolType: TypeSymbol | ArraySymbol,
		Size:       size,
		Scope:      t.Scope,
		Name:       name,
		Type:       _type,
		Members:    nil,
	}
}

func (t *SymbolTable) NewSliceType(token antlr.Token, name string, _type *Symbol) *Symbol {
	return &Symbol{
		SymbolType: TypeSymbol | SliceSymbol,
		Token:      token,
		Size:       0,
		Scope:      t.Scope,
		Name:       name,
		Type:       _type,
		Members:    nil,
	}
}

var (
	Bool         = NewPrimitive("bool", types.I1)
	ConstantBool = NewConstant("bool", Bool, types.I1)
	Int          = NewPrimitive("int", types.I64)
	ConstantInt  = NewConstant("int", Int, types.I64)
	// Int8       = NewPrimitive("int8")
	// Int16      = NewPrimitive("int16")
	// Int32      = NewPrimitive("int32")
	// Int64      = NewPrimitive("int64")
	// Uint       = NewPrimitive("uint")
	// Uint8      = NewPrimitive("uint8")
	// Uint16     = NewPrimitive("uint16")
	// Uint32     = NewPrimitive("uint32")
	// Uint64     = NewPrimitive("uint64")
	// Float32    = NewPrimitive("float32")
	// Float64    = NewPrimitive("float64")
	Float         = NewPrimitive("float", types.Double)
	ConstantFloat = NewConstant("float", Float, types.Double)
	// Complex64  = NewPrimitive("complex64")
	// Complex128 = NewPrimitive("complex128")
	String         = NewPrimitive("string", types.I8Ptr)
	ConstantString = NewConstant("string", String, types.I8Ptr)
	Rune           = NewPrimitive("rune", types.I8)
	ConstantRune   = NewConstant("rune", Rune, types.I8)

	True  *Symbol
	False *Symbol
)

func (t *SymbolTable) addPrimitives() {
	_ = t.AddSymbol(Int)
	_ = t.AddSymbol(Float)
	// _ = t.AddSymbol(Int8)
	// _ = t.AddSymbol(Int16)
	// _ = t.AddSymbol(Int32)
	// _ = t.AddSymbol(Int64)
	// _ = t.AddSymbol(Uint)
	// _ = t.AddSymbol(Uint8)
	// _ = t.AddSymbol(Uint16)
	// _ = t.AddSymbol(Uint32)
	// _ = t.AddSymbol(Uint64)
	// _ = t.AddSymbol(Float32)
	// _ = t.AddSymbol(Float64)
	// _ = t.AddSymbol(Complex64)
	// _ = t.AddSymbol(Complex128)
	_ = t.AddSymbol(String)
	_ = t.AddSymbol(Rune)
	_ = t.AddSymbol(Bool)
	True = t.NewVariable(nil, "true", Bool)
	False = t.NewVariable(nil, "false", Bool)
	_ = t.AddSymbol(True)
	_ = t.AddSymbol(False)

	printf := t.NewFunction(nil, "printf", Int, t.NewVariable(nil, "", String))
	printf.Variadic = true
	_ = t.AddSymbol(printf)

	_ = t.AddSymbol(t.NewFunction(nil, "strcat", String, t.NewVariable(nil, "", String), t.NewVariable(nil, "", String)))
}

func NewSymbolTable() *SymbolTable {
	t := &SymbolTable{
		Symbols: double.NewLinkedList[*Symbol](),
		Scope:   0,
	}
	t.addPrimitives()
	return t
}
