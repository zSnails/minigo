package llvm

import (
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
	"github.com/zSnails/minigo/linked-list/single"
)

type llTable struct {
	currentScope int
	Symbols      *single.LinkedList[*llSymbol]
}

func NewTable() *llTable {
	table := &llTable{
		currentScope: 0,
		Symbols:      single.NewLinkedList[*llSymbol](),
	}
	table.AddSymbol(constant.True)
	table.AddSymbol(constant.False)
	return table
}

type llSymbol struct {
	Scope  int
	Symbol value.Value
}

func (l *llTable) EnterScope() {
	l.currentScope++
}

func (l *llTable) ExitScope() {
	l.Symbols.RemoveIf(func(ls *llSymbol) bool {
		return ls.Scope == l.currentScope
	})
	l.currentScope--
}

func (l *llTable) AddSymbol(val value.Value) {
	v := &llSymbol{
		Scope:  l.currentScope,
		Symbol: val,
	}
	l.Symbols.Add(v)
}

func (l *llTable) GetSymbol(name string) (*llSymbol, bool) {
	symbol, found := l.Symbols.FindFirst(func(ls *llSymbol) bool {
		switch ls := ls.Symbol.(type) {
		case value.Named:
			return ls.Name() == name
		case constant.Constant:
			return ls.Ident() == name
		}
		return false
	})
	if !found {
		return nil, found
	}

	return symbol, found
}
