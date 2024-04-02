package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/zSnails/minigo/grammar"
	"github.com/zSnails/minigo/reporter"
)

const (
	ParsingError = iota + 1
	InternalError
)

var (
	filename string
	isJson   bool
)

func init() {
	flag.StringVar(&filename, "file", "", "The file to parse")
	flag.BoolVar(&isJson, "json-output", false, "Whether or not to show json output")
	flag.Parse()
}

func main() {
	if filename == "" {
		fmt.Fprintln(os.Stderr, "`file` can't be empty")
		os.Exit(InternalError)
	}

	fileStream, err := antlr.NewFileStream(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(InternalError)
	}
	lexer := grammar.NewMinigoLexer(fileStream)
	cts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := grammar.NewMinigoParser(cts)

	var r reporter.Reporter
	parser.RemoveErrorListeners()

	if isJson {
		r = reporter.NewJsonReporter(fileStream.GetSourceName())
	} else {
		r = reporter.NewReporter(fileStream.GetSourceName())
	}

	parser.AddErrorListener(r.(antlr.ErrorListener))
	ctx := parser.Root()
	if r.HasErrors() {
		fmt.Fprintf(os.Stderr, "%s", r.String())
		os.Exit(ParsingError)
	}
	s := ctx.ToStringTree(nil, parser)
	fmt.Printf("s: %v\n", s)
}
