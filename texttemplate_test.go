package apiclient

import (
	"io"
	"strings"
	"testing"
)

type templateParseError struct{}

func (t *templateParseError) Parse(text string) (textTemplate, error) {
	return nil, errMocked
}

func (t *templateParseError) Execute(wr io.Writer, data interface{}) error {
	panic("should not be called")
}

type templateExecuteError struct{}

func (t *templateExecuteError) Parse(text string) (textTemplate, error) {
	return t, nil
}

func (t *templateExecuteError) Execute(wr io.Writer, data interface{}) error {
	return errMocked
}

func TestNewStdlibTextTemplateParseError(t *testing.T) {
	tmpl := newStdlibTextTemplate("antani")
	out, err := tmpl.Parse("{{ .Foo")
	if err == nil || !strings.HasSuffix(err.Error(), "unclosed action") {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if out != nil {
		t.Fatal("expected nil output here")
	}
}
