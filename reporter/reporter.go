package reporter

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type MinigoReporter struct {
	antlr.DefaultErrorListener
	errors   []error
	fileName string
}

func (m *MinigoReporter) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	m.errors = append(m.errors, fmt.Errorf("%s:%d:%d: %s", m.fileName, line, column, msg))
}

func (m *MinigoReporter) GetErrors() []error {
	return m.errors
}

func (m *MinigoReporter) HasErrors() bool {
	return len(m.errors) > 0
}

func NewReporter(filename string) *MinigoReporter {
	return &MinigoReporter{
		DefaultErrorListener: antlr.DefaultErrorListener{},
		errors:               []error{},
		fileName:             filename,
	}
}
