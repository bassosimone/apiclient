package apiclient

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestSwaggerIsValidJSON(t *testing.T) {
	swagger := Swagger()
	if swagger == "" {
		t.Fatal("swagger must be non-empty")
	}
	var v map[string]interface{}
	if err := json.Unmarshal([]byte(swagger), &v); err != nil {
		t.Fatal(err)
	}
}

func TestClientDefaultHTTPClient(t *testing.T) {
	clnt := &Client{}
	if c := clnt.httpClient(); c != http.DefaultClient {
		t.Fatal("not the default http client")
	}
}

func TestClientCustomHTTPClient(t *testing.T) {
	cache := &cacheClient{
		HTTPClient: http.DefaultClient,
		KVStore:    &memkvstore{},
	}
	clnt := &Client{HTTPClient: cache}
	if cx := clnt.httpClient(); cx != cache {
		t.Fatal("not the custom http client")
	}
}

func TestBaseURLWorksAsIntended(t *testing.T) {
	clnt := &Client{}
	if clnt.baseURL() != defaultBaseURL {
		t.Fatal("unexpected default baseURL")
	}
}
