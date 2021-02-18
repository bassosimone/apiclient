package apiclient

import "testing"

func TestBaseURLWorksAsIntended(t *testing.T) {
	clnt := &Client{}
	if clnt.baseURL() != defaultBaseURL {
		t.Fatal("unexpected default baseURL")
	}
}
