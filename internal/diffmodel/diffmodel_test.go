package main

import "testing"

func TestWithProductionAPI(t *testing.T) {
	t.Log("using ", productionURL)
	if compare(productionURL) != 0 {
		t.Fatal("model mismatch (see above)")
	}
}

func TestWithTestingAPI(t *testing.T) {
	t.Log("using ", testingURL)
	if compare(testingURL) != 0 {
		t.Fatal("model mismatch (see above)")
	}
}
