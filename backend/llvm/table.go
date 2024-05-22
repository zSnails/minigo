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
	table.AddSymbol("true", constant.True)
	table.AddSymbol("false", constant.False)
	return table
}

type llSymbol struct {
	Name   string
	Scope  int
	Symbol value.Value
}

func (l *llTable) EnterScope() {
	l.currentScope++
}

func (l *llTable) Replace(name string, symbol value.Value) {
    l.Symbols.ForEach(func(ls *llSymbol) {
        if ls.Name == name {
            ls.Symbol = symbol
        }
    })
}

func (l *llTable) ExitScope() {
	l.Symbols.RemoveIf(func(ls *llSymbol) bool {
		return ls.Scope == l.currentScope
	})
	l.currentScope--
}

func (l *llTable) AddSymbol(name string, val value.Value) {
	v := &llSymbol{
		Name:   name,
		Scope:  l.currentScope,
		Symbol: val,
	}
	l.Symbols.Add(v)
}

func (l *llTable) GetSymbol(name string) (*llSymbol, bool) {
	symbol, found := l.Symbols.FindFirst(func(ls *llSymbol) bool {
		return ls.Name == name
	})
	if !found {
		return nil, found
	}

	return symbol, found
}
