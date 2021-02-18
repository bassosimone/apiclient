package main

import (
	"time"

	"github.com/bassosimone/apiclient/internal/cmd/generator/internal/spec"
)

func genTests(filep file) {
	fprint(filep, "// Code generated by go generate; DO NOT EDIT.\n")
	fprintf(filep, "// %v\n\n", time.Now())
	fprint(filep, "package apiclient\n\n")
	fprint(filep, "import (\n")
	fprint(filep, "\t\"context\"\n")
	fprint(filep, "\t\"encoding/json\"\n")
	fprint(filep, "\t\"errors\"\n")
	fprint(filep, "\t\"io\"\n")
	fprint(filep, "\t\"io/ioutil\"\n")
	fprint(filep, "\t\"net/http/httptest\"\n")
	fprint(filep, "\t\"net/http\"\n")
	fprint(filep, "\t\"net/url\"\n")
	fprint(filep, "\t\"strings\"\n")
	fprint(filep, "\t\"sync\"\n")
	fprint(filep, "\t\"testing\"\n")
	fprint(filep, "\t\"time\"\n")
	fprint(filep, "\n")
	fprint(filep, "\t\"github.com/bassosimone/apiclient/internal/imodel\"\n")
	fprint(filep, "\t\"github.com/bassosimone/apiclient/model\"\n")
	fprint(filep, "\t\"github.com/google/go-cmp/cmp\"\n")
	fprint(filep, ")\n\n")
	fprint(filep, "//go:generate go run ./internal/cmd/generator\n\n")
	for _, desc := range spec.Descriptors {
		fprintf(filep, desc.GenTests())
	}
}
