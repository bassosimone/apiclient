package apiclient

import (
	"encoding/json"
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
