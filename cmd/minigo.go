package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/zSnails/minigo/backend/qbe"
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
	root := parser.Root()
	var visitor qbe.QBEBackendVisitor
	visitor.VisitRoot(root.(*grammar.RootContext))
    fmt.Printf("%s\n", visitor.Builder.String())
}
