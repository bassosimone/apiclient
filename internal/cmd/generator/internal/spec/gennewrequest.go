package spec

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	tagForQuery    = "query"
	tagForRequired = "required"
)

func (d *Descriptor) genNewRequestQueryElemString(sb *strings.Builder, f *reflect.StructField) {
	name := f.Name
	query := f.Tag.Get(tagForQuery)
	if f.Tag.Get(tagForRequired) == "true" {
		fmt.Fprintf(sb, "\tif req.%s == \"\" {\n", name)
		fmt.Fprintf(sb, "\t\treturn nil, newErrEmptyField(\"%s\")\n", name)
		fmt.Fprintf(sb, "\t}\n")
		fmt.Fprintf(sb, "\tq.Add(\"%s\", req.%s)\n", query, name)
		return
	}
	fmt.Fprintf(sb, "\tif req.%s != \"\" {\n", name)
	fmt.Fprintf(sb, "\t\tq.Add(\"%s\", req.%s)\n", query, name)
	fmt.Fprintf(sb, "\t}\n")
}

func (d *Descriptor) genNewRequestQueryElemBool(sb *strings.Builder, f *reflect.StructField) {
	// required does not make much sense for a boolean field
	name := f.Name
	query := f.Tag.Get(tagForQuery)
	fmt.Fprintf(sb, "\tif req.%s {\n", name)
	fmt.Fprintf(sb, "\t\tq.Add(\"%s\", \"true\")\n", query)
	fmt.Fprintf(sb, "\t}\n")
}

func (d *Descriptor) genNewRequestQueryElemInt64(sb *strings.Builder, f *reflect.StructField) {
	// required does not make much sense for an integer field
	name := f.Name
	query := f.Tag.Get(tagForQuery)
	fmt.Fprintf(sb, "\tif req.%s != 0 {\n", name)
	fmt.Fprintf(sb, "\t\tq.Add(\"%s\", newQueryFieldInt64(req.%s))\n", query, name)
	fmt.Fprintf(sb, "\t}\n")
}

func (d *Descriptor) genNewRequestQuery(sb *strings.Builder) {
	if d.Method != "GET" {
		return // we only generate query for GET
	}
	fields := d.getStructFieldsWithTag(d.Request, tagForQuery)
	if len(fields) <= 0 {
		return
	}
	fmt.Fprint(sb, "\tq := url.Values{}\n")
	for idx, f := range fields {
		switch f.Type.Kind() {
		case reflect.String:
			d.genNewRequestQueryElemString(sb, f)
		case reflect.Bool:
			d.genNewRequestQueryElemBool(sb, f)
		case reflect.Int64:
			d.genNewRequestQueryElemInt64(sb, f)
		default:
			panic(fmt.Sprintf("unexpected query type at index %d", idx))
		}
	}
	fmt.Fprint(sb, "\tURL.RawQuery = q.Encode()\n")
}

func (d *Descriptor) genNewRequestCallNewRequest(sb *strings.Builder) {
	if d.Method == "POST" {
		fmt.Fprint(sb, "\tbody, err := api.jsonCodec.Encode(req)\n")
		fmt.Fprint(sb, "\tif err != nil {\n")
		fmt.Fprint(sb, "\t\treturn nil, err\n")
		fmt.Fprint(sb, "\t}\n")
		fmt.Fprint(sb, "\tout, err := api.requestMaker.NewRequest(")
		fmt.Fprintf(sb, "ctx, \"%s\", URL.String(), ", d.Method)
		fmt.Fprint(sb, "bytes.NewReader(body))\n")
		fmt.Fprint(sb, "\tif err != nil {\n")
		fmt.Fprint(sb, "\t\treturn nil, err\n")
		fmt.Fprint(sb, "\t}\n")
		fmt.Fprint(sb, "\tout.Header.Set(\"Content-Type\", \"application/json\")\n")
		fmt.Fprint(sb, "\treturn out, nil\n")
		return
	}
	fmt.Fprint(sb, "\treturn api.requestMaker.NewRequest(")
	fmt.Fprintf(sb, "ctx, \"%s\", URL.String(), ", d.Method)
	fmt.Fprint(sb, "nil)\n")
}

func (d *Descriptor) genNewRequest(sb *strings.Builder) {

	fmt.Fprintf(
		sb, "func (api *%s) newRequest(ctx context.Context, req %s) %s {\n",
		d.apiStructName(), d.requestTypeName(), "(*http.Request, error)")
	fmt.Fprint(sb, "\tURL, err := url.Parse(api.baseURL)\n")
	fmt.Fprint(sb, "\tif err != nil {\n")
	fmt.Fprint(sb, "\t\treturn nil, err\n")
	fmt.Fprint(sb, "\t}\n")

	switch d.URLPath.IsTemplate {
	case false:
		fmt.Fprintf(sb, "\tURL.Path = \"%s\"\n", d.URLPath.Value)
	case true:
		fmt.Fprintf(
			sb, "\tup, err := api.templateExecutor.Execute(\"%s\", req)\n", d.URLPath.Value)
		fmt.Fprint(sb, "\tif err != nil {\n")
		fmt.Fprint(sb, "\t\treturn nil, err\n")
		fmt.Fprint(sb, "\t}\n")
		fmt.Fprint(sb, "\tURL.Path = up\n")
	}

	d.genNewRequestQuery(sb)
	d.genNewRequestCallNewRequest(sb)

	fmt.Fprintf(sb, "}\n\n")
}

// GenNewRequest generates the code that creates a http.Request
// given a specific API call.
func (d *Descriptor) GenNewRequest() string {
	var sb strings.Builder
	d.genNewRequest(&sb)
	return sb.String()
}
