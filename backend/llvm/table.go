package llvm

import (
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
	"github.com/zSnails/minigo/linked-list/single"
)

type LlvmSymbolTable struct {
	currentScope int
	Symbols      *single.LinkedList[*llSymbol]
}

func NewTable() *LlvmSymbolTable {
	table := &LlvmSymbolTable{
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

func (l *LlvmSymbolTable) EnterScope() {
	l.currentScope++
}

func (l *LlvmSymbolTable) Replace(name string, symbol value.Value) {
	l.Symbols.ForEach(func(ls *llSymbol) {
		if ls.Name == name {
			ls.Symbol = symbol
		}
	})
}

func (l *LlvmSymbolTable) ExitScope() {
	l.Symbols.RemoveIf(func(ls *llSymbol) bool {
		return ls.Scope == l.currentScope
	})
	l.currentScope--
}

func (l *LlvmSymbolTable) AddSymbol(name string, val value.Value) {
	v := &llSymbol{
		Name:   name,
		Scope:  l.currentScope,
		Symbol: val,
	}
	l.Symbols.Add(v)
}

func (l *LlvmSymbolTable) GetSymbol(name string) (*llSymbol, bool) {
	symbol, found := l.Symbols.FindFirst(func(ls *llSymbol) bool {
		return ls.Name == name
	})
	if !found {
		return nil, found
	}

	return symbol, found
}
