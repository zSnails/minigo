package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/teris-io/cli"
	"github.com/zSnails/minigo/backend/llvm"
	"github.com/zSnails/minigo/grammar"
	"github.com/zSnails/minigo/reporter"
	checker "github.com/zSnails/minigo/type-checker"
)

const (
	CompilerError = iota + 1
	InternalError
	ExternalError
)

var tmp = os.TempDir()

func getFileName(input string) string {
	return strings.TrimSuffix(input, filepath.Ext(input))
}

func main() {
	app := cli.New("minigo compiler")
	json := cli.NewOption("json", "wheter or not to report errors as json").WithType(cli.TypeBool)
	app.WithCommand(cli.NewCommand("build", "build source code into an executable").
		WithArg(cli.NewArg("file", "the source file to be built")).
		WithOption(json).
		WithAction(build))
	app.WithCommand(cli.NewCommand("run", "build and run source code into an executable").
		WithOption(json).
		WithArg(cli.NewArg("file", "the source file to be built and run")).
		WithAction(func(args []string, options map[string]string) int {
			code := build(args, options)
			if code != 0 {
				return code
			}
			filename := path.Clean(args[0])
			executable := getFileName(path.Base(filename))
			return run(executable)
		}))

	code := app.Run(os.Args, os.Stdout)
	os.Exit(code)
}

func build(args []string, options map[string]string) int {
	filename := path.Clean(args[0])

	fileStream, err := antlr.NewFileStream(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return InternalError
	}
	lexer := grammar.NewMinigoLexer(fileStream)
	lexer.RemoveErrorListeners()
	cts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := grammar.NewMinigoParser(cts)

	var r reporter.Reporter
	parser.RemoveErrorListeners()

	_, json := options["json"]

	if json {
		r = reporter.NewJsonReporter(fileStream.GetSourceName())
	} else {
		r = reporter.NewReporter(fileStream.GetSourceName())
	}
	listener, ok := r.(antlr.ErrorListener)
	if !ok {
		return InternalError
	}

	lexer.AddErrorListener(listener)
	parser.AddErrorListener(listener)
	ctx := parser.Root()
	if r.HasErrors() {
		fmt.Fprintln(os.Stderr, r.String())
		return CompilerError
	}

	typeChecker := checker.NewTypeChecker(fileStream.GetSourceName(), listener)
	typeChecker.Visit(ctx)
	if r.HasErrors() {
		fmt.Fprintln(os.Stderr, r.String())
		return CompilerError
	}

	backend := llvm.NewLlvmBackend(listener)
	backend.Visit(ctx)
	if r.HasErrors() {
		fmt.Fprintln(os.Stderr, r.String())
		return CompilerError
	}
	backend.GetModule().SourceFilename = filename

	base := filepath.Base(filename)
	out, err := os.CreateTemp(tmp, base+"*.ll")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return CompilerError
	}

	defer func() {
		err := os.Remove(out.Name())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(CompilerError)
		}
	}()

	defer out.Close()

	cmd := exec.Command("clang", out.Name(), "-o", getFileName(base)+EXTENSION)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	_, err = backend.GetModule().WriteTo(out)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return CompilerError
	}

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return CompilerError
	}

	return 0
}
