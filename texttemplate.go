package apiclient

import (
	"io"
	"text/template"
)

type textTemplate interface {
	Parse(text string) (textTemplate, error)
	Execute(wr io.Writer, data interface{}) error
}

type stdlibTextTemplate struct {
	*template.Template
}

func (t *stdlibTextTemplate) Parse(text string) (textTemplate, error) {
	out, err := t.Template.Parse(text)
	if err != nil {
		return nil, err
	}
	return &stdlibTextTemplate{out}, nil
}

func newStdlibTextTemplate(name string) textTemplate {
	return &stdlibTextTemplate{template.New(name)}
}
