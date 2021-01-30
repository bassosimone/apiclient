// This script generates newresponse.go.
package main

import (
	"time"

	"github.com/bassosimone/apiclient/internal/apimodel"
	"github.com/bassosimone/apiclient/internal/fatalx"
	"github.com/bassosimone/apiclient/internal/fmtx"
	"github.com/bassosimone/apiclient/internal/osx"
	"github.com/bassosimone/apiclient/internal/reflectx"
)

func gettype(in interface{}) string {
	sinfo, err := reflectx.NewTypeValueInfo(in)
	fatalx.OnError(err, "reflectx.NewStructInfo failed")
	return sinfo.TypeName()
}

func genbeginfunc(filep osx.File, desc *apimodel.Descriptor) {
	typename := gettype(desc.Response)
	fmtx.Fprintf(filep, "func new%s", typename)
	fmtx.Fprint(filep, "(resp *http.Response, err error)")
	fmtx.Fprintf(filep, " (*%s, error) {\n", typename)
}

func genparse(filep osx.File, desc *apimodel.Descriptor) {
	typename := gettype(desc.Response)
	fmtx.Fprint(filep, "\tif err != nil {\n")
	fmtx.Fprint(filep, "\t\treturn nil, err\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tif resp.StatusCode != 200 {\n")
	fmtx.Fprint(filep, "\t\treturn nil, errors.New(\"apiclient: http request failed\")\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tdefer resp.Body.Close()\n")
	fmtx.Fprint(filep, "\treader := io.LimitReader(resp.Body, 4<<20)\n")
	fmtx.Fprint(filep, "\tdata, err := ioutil.ReadAll(reader)\n")
	fmtx.Fprint(filep, "\tif err != nil {\n")
	fmtx.Fprint(filep, "\t\treturn nil, err\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprintf(filep, "\tvar out %s\n", typename)
	fmtx.Fprint(filep, "\tif err := json.Unmarshal(data, &out); err != nil {\n")
	fmtx.Fprint(filep, "\t\treturn nil, err\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\treturn &out, nil\n")
}

func genendfunc(filep osx.File) {
	fmtx.Fprintf(filep, "}\n\n")
}

func genapi(filep osx.File, desc *apimodel.Descriptor) {
	genbeginfunc(filep, desc)
	genparse(filep, desc)
	genendfunc(filep)
}

func main() {
	filep := osx.MustCreate("newresponse.go")
	defer filep.Close()

	fmtx.Fprint(filep, "// Code generated by go generate; DO NOT EDIT.\n")
	fmtx.Fprintf(filep, "// %v\n\n", time.Now())
	fmtx.Fprint(filep, "package apiclient\n\n")
	fmtx.Fprint(filep, "import (\n")
	fmtx.Fprint(filep, "\t\"encoding/json\"\n")
	fmtx.Fprint(filep, "\t\"errors\"\n")
	fmtx.Fprint(filep, "\t\"io/ioutil\"\n")
	fmtx.Fprint(filep, "\t\"io\"\n")
	fmtx.Fprint(filep, "\t\"net/http\"\n")
	fmtx.Fprint(filep, ")\n\n")

	fmtx.Fprint(filep, "//go:generate go run ./internal/gennewresponse/...\n\n")

	for _, descr := range apimodel.Descriptors {
		genapi(filep, &descr)
	}
}
