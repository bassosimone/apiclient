package main

import "testing"

func TestMain(t *testing.T) {
	if compare("../../swagger.json") != 0 {
		t.Fatal("model mismatch (see above)")
	}
}
