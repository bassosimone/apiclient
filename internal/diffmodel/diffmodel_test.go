package main

import "testing"

func TestWithProductionAPI(t *testing.T) {
	t.Log("using ", productionURL)
	if compare("../../swagger.json", productionURL) != 0 {
		t.Fatal("model mismatch (see above)")
	}
}

func TestWithTestingAPI(t *testing.T) {
	t.Log("using ", testingURL)
	if compare("../../swagger.json", testingURL) != 0 {
		t.Fatal("model mismatch (see above)")
	}
}
