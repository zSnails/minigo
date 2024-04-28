package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/antlr4-go/antlr/v4"
	"github.com/zSnails/minigo/grammar"
	"github.com/zSnails/minigo/reporter"
	checker "github.com/zSnails/minigo/type-checker"
)

const (
	CompilerError = iota + 1
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

	filename = path.Clean(filename)

	fileStream, err := antlr.NewFileStream(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(InternalError)
	}
	lexer := grammar.NewMinigoLexer(fileStream)
	lexer.RemoveErrorListeners()
	cts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := grammar.NewMinigoParser(cts)

	var r reporter.Reporter
	parser.RemoveErrorListeners()

	if isJson {
		r = reporter.NewJsonReporter(fileStream.GetSourceName())
	} else {
		r = reporter.NewReporter(fileStream.GetSourceName())
	}

	lexer.AddErrorListener(r.(antlr.ErrorListener))
	parser.AddErrorListener(r.(antlr.ErrorListener))
	ctx := parser.Root()
	if r.HasErrors() {
		fmt.Fprintf(os.Stderr, "%s", r.String())
		os.Exit(CompilerError)
	}

	typeChecker := checker.NewTypeChecker(fileStream.GetSourceName(), r.(antlr.ErrorListener))
	typeChecker.Visit(ctx)
	if r.HasErrors() {
		fmt.Fprintf(os.Stderr, "%s", r.String())
		os.Exit(CompilerError)
	}
}
