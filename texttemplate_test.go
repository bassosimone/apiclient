package apiclient

import "io"

type templateParseError struct{}

func (t *templateParseError) Parse(text string) (textTemplate, error) {
	return nil, ErrMocked
}

func (t *templateParseError) Execute(wr io.Writer, data interface{}) error {
	panic("should not be called")
}

type templateExecuteError struct{}

func (t *templateExecuteError) Parse(text string) (textTemplate, error) {
	return t, nil
}

func (t *templateExecuteError) Execute(wr io.Writer, data interface{}) error {
	return ErrMocked
}
