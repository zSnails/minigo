package reporter

import (
	"encoding/json"

	"github.com/antlr4-go/antlr/v4"
)

type JsonReporter struct {
	antlr.DefaultErrorListener
	errors   []error
	fileName string
}

type jsonError struct {
	Line     int    `json:"line"`
	Column   int    `json:"column"`
	Message  string `json:"message"`
	FileName string `json:"fileName"`
}

func (j *jsonError) Error() string {
	data, _ := json.Marshal(j)
	return string(data)
}

func (j *JsonReporter) String() string {
	data, _ := json.Marshal(j.errors)
	return string(data)
}

func (m *JsonReporter) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	m.errors = append(m.errors, &jsonError{FileName: m.fileName, Line: line, Column: column, Message: msg})
}

func (m *JsonReporter) GetErrors() []error {
	return m.errors
}

func (m *JsonReporter) HasErrors() bool {
	return len(m.errors) > 0
}

func NewJsonReporter(filename string) *JsonReporter {
	return &JsonReporter{
		DefaultErrorListener: antlr.DefaultErrorListener{},
		errors:               []error{},
		fileName:             filename,
	}
}
