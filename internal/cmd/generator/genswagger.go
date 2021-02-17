package main

import (
	"encoding/json"
	"time"

	"github.com/bassosimone/apiclient/internal/cmd/generator/internal/spec"
	"github.com/bassosimone/apiclient/internal/cmd/internal/fatalx"
	"github.com/bassosimone/apiclient/internal/openapi"
)

func genSwaggerVersion() string {
	return time.Now().UTC().Format("0.20060102.1150405")
}

func genSwagger() {
	swagger := openapi.Swagger{
		Swagger: "2.0",
		Info: openapi.API{
			Title:   "OONI API specification",
			Version: genSwaggerVersion(),
		},
		Host:     "api.ooni.io",
		BasePath: "/",
		Schemes:  []string{"https"},
		Paths:    make(map[string]*openapi.Path),
	}
	for _, desc := range spec.Descriptors {
		pathStr, pathInfo := desc.GenSwaggerPath()
		swagger.Paths[pathStr] = pathInfo
	}
	data, err := json.MarshalIndent(swagger, "", "    ")
	fatalx.OnError(err, "json.Marshal failed")
	filep := mustCreateFile("swagger.go")
	defer filep.Close()
	fprint(filep, "// Code generated by go generate; DO NOT EDIT.\n")
	fprintf(filep, "// %v\n\n", time.Now())
	fprint(filep, "package apiclient\n\n")
	fprint(filep, "//go:generate go run ./internal/cmd/generator\n\n")
	fprintf(filep, "var swagger = `%s`\n", string(data))
}
