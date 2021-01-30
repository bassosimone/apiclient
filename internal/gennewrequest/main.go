// This script generates newrequest.go
package main

import (
	"reflect"
	"strings"
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

func gettags(in interface{}, tagName string) []*reflectx.FieldInfo {
	sinfo, err := reflectx.NewTypeValueInfo(in)
	fatalx.OnError(err, "reflectx.NewStructInfo failed")
	finfo, err := sinfo.AllFieldsWithTag(tagName)
	fatalx.OnError(err, "sinfo.AllFieldsWithTag failed")
	return finfo
}

func genbeginfunc(filep osx.File, desc *apimodel.Descriptor) {
	typename := gettype(desc.Request)
	fmtx.Fprintf(filep, "func new%s", typename)
	fmtx.Fprint(filep, "(ctx context.Context, ")
	fmtx.Fprint(filep, "baseURL string, ")
	fmtx.Fprintf(filep, "req *%s)", typename)
	fmtx.Fprint(filep, " (*http.Request, error) {\n")
}

func genurlpath(filep osx.File, desc *apimodel.Descriptor) {
	if !strings.Contains(desc.URLPath, "{{ ") {
		fmtx.Fprintf(filep, "\tURL.Path = \"%s\"\n", desc.URLPath)
		return
	}
	fmtx.Fprintf(filep, "\ttmpl, err := template.New(\"urlpath\").Parse(\"%s\")\n", desc.URLPath)
	fmtx.Fprint(filep, "\tif err != nil {\n")
	fmtx.Fprint(filep, "\t\treturn nil, err\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tvar urlpath strings.Builder\n")
	fmtx.Fprint(filep, "\terr = tmpl.Execute(&urlpath, req)\n")
	fmtx.Fprint(filep, "\tif err != nil {\n")
	fmtx.Fprint(filep, "\t\treturn nil, err\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tURL.Path = urlpath.String()\n")
}

func genqueryforstring(filep osx.File, field *reflectx.FieldInfo) {
	name := field.Self.Name
	query := field.Self.Tag.Get("query")
	if field.Self.Tag.Get("mandatory") == "true" {
		fmtx.Fprintf(filep, "\tif req.%s == \"\" {\n", name)
		fmtx.Fprint(filep, "\t\treturn nil, errors.New(")
		fmtx.Fprintf(filep, "\"apiclient: empty %s field\")\n", name)
		fmtx.Fprintf(filep, "\t}\n")
		fmtx.Fprintf(filep, "\tquery.Add(\"%s\", ", query)
		fmtx.Fprintf(filep, "req.%s)\n", name)
		return
	}
	fmtx.Fprintf(filep, "\tif req.%s != \"\" {\n", name)
	fmtx.Fprintf(filep, "\t\tquery.Add(\"%s\", ", query)
	fmtx.Fprintf(filep, "req.%s)\n", name)
	fmtx.Fprintf(filep, "\t}\n")
}

func genqueryforbool(filep osx.File, field *reflectx.FieldInfo) {
	name := field.Self.Name
	query := field.Self.Tag.Get("query")
	// mandatory does not make much sense for a boolean field
	fmtx.Fprintf(filep, "\tif req.%s {\n", name)
	fmtx.Fprintf(filep, "\t\tquery.Add(\"%s\", \"true\")\n", query)
	fmtx.Fprintf(filep, "\t}\n")
}

func genqueryforint64(filep osx.File, field *reflectx.FieldInfo) {
	name := field.Self.Name
	query := field.Self.Tag.Get("query")
	// mandatory does not make much sense for an integer field
	fmtx.Fprintf(filep, "\tif req.%s != 0 {\n", name)
	fmtx.Fprintf(filep, "\t\tquery.Add(\"%s\", fmt.Sprintf(\"%%d\", req.%s))\n", query, name)
	fmtx.Fprintf(filep, "\t}\n")
}

func genquery(filep osx.File, desc *apimodel.Descriptor) {
	if desc.Method == "POST" {
		return
	}
	fields := gettags(desc.Request, "query")
	if fields == nil {
		return
	}
	fmtx.Fprint(filep, "\tquery := url.Values{}\n")
	for _, field := range fields {
		switch field.Self.Type.Kind() {
		case reflect.String:
			genqueryforstring(filep, field)
		case reflect.Bool:
			genqueryforbool(filep, field)
		case reflect.Int64:
			genqueryforint64(filep, field)
		default:
			panic("query: unsupported field type")
		}
	}
	fmtx.Fprint(filep, "\tURL.RawQuery = query.Encode()\n")
}

func gencreaterequest(filep osx.File, desc *apimodel.Descriptor) {
	if desc.Method == "POST" {
		fmtx.Fprint(filep, "\tbody, err := json.Marshal(req)\n")
		fmtx.Fprint(filep, "\tif err != nil {\n")
		fmtx.Fprint(filep, "\t\treturn nil, err\n")
		fmtx.Fprint(filep, "\t}\n")
		fmtx.Fprint(filep, "\tout, err := http.NewRequestWithContext(")
		fmtx.Fprintf(filep, "ctx, \"%s\", URL.String(), ", desc.Method)
		fmtx.Fprint(filep, "bytes.NewReader(body))\n")
		fmtx.Fprint(filep, "\tif err != nil {\n")
		fmtx.Fprint(filep, "\t\treturn nil, err\n")
		fmtx.Fprint(filep, "\t}\n")
		fmtx.Fprint(filep, "\tout.Header.Set(\"Content-Type\", \"application/json\")\n")
		fmtx.Fprint(filep, "\treturn out, nil\n")
		return
	}
	fmtx.Fprint(filep, "\treturn http.NewRequestWithContext(")
	fmtx.Fprintf(filep, "ctx, \"%s\", URL.String(), ", desc.Method)
	fmtx.Fprint(filep, "nil)\n")
}

func genmakeurl(filep osx.File, desc *apimodel.Descriptor) {
	fmtx.Fprint(filep, "\tURL, err := url.Parse(baseURL)\n")
	fmtx.Fprint(filep, "\tif err != nil {\n")
	fmtx.Fprint(filep, "\t\treturn nil, err\n")
	fmtx.Fprint(filep, "\t}\n")
}

func genendfunc(filep osx.File) {
	fmtx.Fprintf(filep, "}\n\n")
}

func genapi(filep osx.File, desc *apimodel.Descriptor) {
	genbeginfunc(filep, desc)
	genmakeurl(filep, desc)
	genurlpath(filep, desc)
	genquery(filep, desc)
	gencreaterequest(filep, desc)
	genendfunc(filep)
}

func main() {
	filep := osx.MustCreate("newrequest.go")
	defer filep.Close()

	fmtx.Fprint(filep, "// Code generated by go generate; DO NOT EDIT.\n")
	fmtx.Fprintf(filep, "// %v\n\n", time.Now())
	fmtx.Fprint(filep, "package apiclient\n\n")
	fmtx.Fprint(filep, "import (\n")
	fmtx.Fprint(filep, "\t\"bytes\"\n")
	fmtx.Fprint(filep, "\t\"context\"\n")
	fmtx.Fprint(filep, "\t\"encoding/json\"\n")
	fmtx.Fprint(filep, "\t\"errors\"\n")
	fmtx.Fprint(filep, "\t\"fmt\"\n")
	fmtx.Fprint(filep, "\t\"text/template\"\n")
	fmtx.Fprint(filep, "\t\"net/http\"\n")
	fmtx.Fprint(filep, "\t\"net/url\"\n")
	fmtx.Fprint(filep, "\t\"strings\"\n")
	fmtx.Fprint(filep, ")\n\n")

	fmtx.Fprint(filep, "//go:generate go run ./internal/gennewrequest/...\n\n")

	for _, descr := range apimodel.Descriptors {
		genapi(filep, &descr)
	}
}
