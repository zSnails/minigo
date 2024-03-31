package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/zSnails/minigo/grammar"
	"github.com/zSnails/minigo/reporter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	fileStream, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		panic(err)
	}
	lexer := grammar.NewMinigoLexer(fileStream)
	cts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := grammar.NewMinigoParser(cts)
	r := reporter.NewReporter(fileStream.GetSourceName())
	parser.RemoveErrorListeners()
	parser.AddErrorListener(r)
	ctx := parser.Root()
	if r.HasErrors() {
		errs := r.GetErrors()
		for _, err := range errs {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
		return
	}
	s := ctx.ToStringTree(nil, parser)
	fmt.Printf("s: %v\n", s)
}
