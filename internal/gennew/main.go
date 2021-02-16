// Command gennew generates new.go.
package main

import (
	"strings"
	"time"

	"github.com/bassosimone/apiclient/internal/apimodel"
	"github.com/bassosimone/apiclient/internal/fmtx"
	"github.com/bassosimone/apiclient/internal/osx"
	"github.com/bassosimone/apiclient/internal/reflectx"
)

func genapi(filep osx.File, desc *apimodel.Descriptor) {
	typename := reflectx.Must(reflectx.NewTypeValueInfo(desc.Request)).TypeName()
	apiname := strings.Replace(typename, "Request", "API", 1)
	fmtx.Fprintf(filep, "// New%s creates a new instance of %s.\n", apiname, apiname)
	fmtx.Fprintf(filep, "func New%s(clnt *Client) (*%s, error) {\n", apiname, apiname)
	fmtx.Fprintf(filep, "\tapi := &%s{\n", apiname)
	if desc.RequiresLogin {
		fmtx.Fprintf(filep, "\t\tAuthorizer: clnt,\n")
	}
	fmtx.Fprintf(filep, "\t\tBaseURL: clnt.BaseURL,\n")
	fmtx.Fprintf(filep, "\t\tHTTPClient: clnt.HTTPClient,\n")
	fmtx.Fprintf(filep, "\t\tUserAgent: clnt.UserAgent,\n")
	fmtx.Fprintf(filep, "\t}\n")
	fmtx.Fprintf(filep, "\treturn api, nil\n")
	fmtx.Fprint(filep, "}\n\n")
}

func main() {
	filep := osx.MustCreate("new.go")
	defer filep.Close()

	fmtx.Fprint(filep, "// Code generated by go generate; DO NOT EDIT.\n")
	fmtx.Fprintf(filep, "// %v\n\n", time.Now())
	fmtx.Fprint(filep, "package apiclient\n\n")

	fmtx.Fprint(filep, "//go:generate go run ./internal/gennew/...\n\n")

	for _, descr := range apimodel.Descriptors {
		genapi(filep, &descr)
	}
}