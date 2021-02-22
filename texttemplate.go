package apiclient

import (
	"strings"
	"text/template"
)

type stdlibTemplateExecutor struct{}

func (*stdlibTemplateExecutor) Execute(tmpl string, v interface{}) (string, error) {
	to, err := template.New("t").Parse(tmpl)
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	if err := to.Execute(&sb, v); err != nil {
		return "", err
	}
	return sb.String(), nil
}
