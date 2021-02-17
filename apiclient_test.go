package apiclient_test

import (
	"encoding/json"
	"testing"

	"github.com/bassosimone/apiclient"
)

func TestSwaggerIsValidJSON(t *testing.T) {
	swagger := apiclient.Swagger()
	if swagger == "" {
		t.Fatal("swagger must be non-empty")
	}
	var v map[string]interface{}
	if err := json.Unmarshal([]byte(swagger), &v); err != nil {
		t.Fatal(err)
	}
}
