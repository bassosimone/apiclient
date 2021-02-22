package apiclient

import (
	"errors"
	"testing"
)

func TestMemKVStoreGetFailure(t *testing.T) {
	value, err := (&memkvstore{}).Get("antani")
	if !errors.Is(err, errMemkvstoreNotFound) {
		t.Fatal("unexpected error", err)
	}
	if value != nil {
		t.Fatal("expected nil value here")
	}
}
