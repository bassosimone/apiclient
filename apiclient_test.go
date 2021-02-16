package apiclient

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSwagger(t *testing.T) {
	swagger := Swagger()
	if swagger == "" {
		t.Fatal("swagger must be non-empty")
	}
	var v map[string]interface{}
	if err := json.Unmarshal([]byte(swagger), &v); err != nil {
		t.Fatal(err)
	}
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
