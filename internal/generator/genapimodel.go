package main

import (
	"time"

	"github.com/bassosimone/apiclient/internal/generator/internal/spec"
)

func genAPIModel() {
	filep := mustCreateFile("spec.go")
	defer filep.Close()

	fprint(filep, "// Code generated by go generate; DO NOT EDIT.\n")
	fprintf(filep, "// %v\n\n", time.Now())
	fprint(filep, "package apiclient\n\n")
	fprint(filep, "import (\n")
	fprint(filep, "\t\"bytes\"\n")
	fprint(filep, "\t\"context\"\n")
	fprint(filep, "\t\"encoding/json\"\n")
	fprint(filep, "\t\"io/ioutil\"\n")
	fprint(filep, "\t\"io\"\n")
	fprint(filep, "\t\"net/http\"\n")
	fprint(filep, "\t\"net/url\"\n")
	fprint(filep, "\t\"strings\"\n")
	fprint(filep, "\n")
	fprint(filep, "\t\"github.com/bassosimone/apiclient/internal/imodel\"\n")
	fprint(filep, "\t\"github.com/bassosimone/apiclient/model\"\n")
	fprint(filep, ")\n\n")

	fprint(filep, "//go:generate go run ./internal/generator/...\n\n")

	for _, desc := range spec.Descriptors {
		fprintf(filep, desc.GenClientCall())
		fprintf(filep, desc.GenAPIType())
		fprintf(filep, desc.GenNewAPI())
		fprintf(filep, desc.GenAPICall())
		fprintf(filep, desc.GenNewRequest())
		fprintf(filep, desc.GenNewResponse())
	}
}
