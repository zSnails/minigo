package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/zSnails/minigo/grammar"
)

func main() {
	fileStream, err := antlr.NewFileStream("tests/fuap.minigo")
	if err != nil {
		panic(err)
	}
	lexer := grammar.NewMinigoLexer(fileStream)
	cts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := grammar.NewMinigoParser(cts)
	ctx := parser.Root()
	s := ctx.ToStringTree(nil, parser)
	fmt.Printf("s: %v\n", s)
}
