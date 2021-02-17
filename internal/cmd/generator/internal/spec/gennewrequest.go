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

func (d *Descriptor) getStructFields(in interface{}) []*reflect.StructField {
	value := reflect.ValueOf(in)
	if value.Type().Kind() == reflect.Ptr {
		if value.IsNil() {
			panic("null pointer")
		}
		value = value.Elem()
	}
	if value.Type().Kind() != reflect.Struct {
		panic("expected a struct")
	}
	var out []*reflect.StructField
	for idx := 0; idx < value.NumField(); idx++ {
		f := value.Type().Field(idx)
		out = append(out, &f)
	}
	return out
}

func (d *Descriptor) getStructFieldsWithTag(in interface{}, tag string) []*reflect.StructField {
	var out []*reflect.StructField
	for _, f := range d.getStructFields(in) {
		if f.Tag.Get(tag) != "" {
			out = append(out, f)
		}
	}
	return out
}

func (d *Descriptor) genNewRequestQueryElemString(sb *strings.Builder, f *reflect.StructField) {
	name := f.Name
	query := f.Tag.Get(tagForQuery)
	if f.Tag.Get(tagForRequired) == "true" {
		fmt.Fprintf(sb, "\tif req.%s == \"\" {\n", name)
		fmt.Fprintf(sb, "\t\treturn \"\", newErrEmptyField(\"%s\")\n", name)
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

func (d *Descriptor) genNewRequestQuery(sb *strings.Builder) bool {
	if d.Method != "GET" {
		return false
	}
	fields := d.getStructFieldsWithTag(d.Request, tagForQuery)
	if len(fields) <= 0 {
		return false
	}
	fmt.Fprintf(sb, "func (api *%s) newRequestQuery(req %s) (string, error) {\n",
		d.apiStructName(), d.requestTypeName())
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
	fmt.Fprint(sb, "\treturn q.Encode(), nil\n")
	fmt.Fprintf(sb, "}\n\n")
	return true
}

func (d *Descriptor) genNewRequestURLPath(sb *strings.Builder) {
	if !d.URLPath.IsTemplate {
		return
	}
	fmt.Fprintf(
		sb, "func (api *%s) newRequestURLPath(req %s) (string, error) {\n",
		d.apiStructName(), d.requestTypeName())
	fmt.Fprintf(sb, "\tnewTemplate := newStdlibTextTemplate\n")
	fmt.Fprintf(sb, "\tif api.newTemplate != nil {\n")
	fmt.Fprintf(sb, "\t\tnewTemplate = api.newTemplate\n")
	fmt.Fprintf(sb, "\t}\n")
	fmt.Fprintf(sb, "\ttmpl, err := newTemplate(\"urlpath\").Parse(\"%s\")\n", d.URLPath.Value)
	fmt.Fprint(sb, "\tif err != nil {\n")
	fmt.Fprint(sb, "\t\treturn \"\", err\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tvar urlpath strings.Builder\n")
	fmt.Fprint(sb, "\terr = tmpl.Execute(&urlpath, req)\n")
	fmt.Fprint(sb, "\tif err != nil {\n")
	fmt.Fprint(sb, "\t\treturn \"\", err\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\treturn urlpath.String(), nil\n")
	fmt.Fprintf(sb, "}\n\n")
}

func (d *Descriptor) genNewRequestCallNewRequest(sb *strings.Builder) {
	emit := func() {
		fmt.Fprint(sb, "\tnewRequest := http.NewRequestWithContext\n")
		fmt.Fprint(sb, "\tif api.NewRequest != nil {\n")
		fmt.Fprint(sb, "\t\tnewRequest = api.NewRequest\n")
		fmt.Fprint(sb, "\t}\n")
	}
	if d.Method == "POST" {
		fmt.Fprint(sb, "\tmarshal := json.Marshal\n")
		fmt.Fprint(sb, "\tif api.marshal != nil {\n")
		fmt.Fprint(sb, "\t\tmarshal = api.marshal\n")
		fmt.Fprint(sb, "\t}\n")
		fmt.Fprint(sb, "\tbody, err := marshal(req)\n")
		fmt.Fprint(sb, "\tif err != nil {\n")
		fmt.Fprint(sb, "\t\treturn nil, err\n")
		fmt.Fprint(sb, "\t}\n")
		emit()
		fmt.Fprint(sb, "\tout, err := newRequest(")
		fmt.Fprintf(sb, "ctx, \"%s\", URL.String(), ", d.Method)
		fmt.Fprint(sb, "bytes.NewReader(body))\n")
		fmt.Fprint(sb, "\tif err != nil {\n")
		fmt.Fprint(sb, "\t\treturn nil, err\n")
		fmt.Fprint(sb, "\t}\n")
		fmt.Fprint(sb, "\tout.Header.Set(\"Content-Type\", \"application/json\")\n")
		fmt.Fprint(sb, "\treturn out, nil\n")
		return
	}
	emit()
	fmt.Fprint(sb, "\treturn newRequest(")
	fmt.Fprintf(sb, "ctx, \"%s\", URL.String(), ", d.Method)
	fmt.Fprint(sb, "nil)\n")
}

func (d *Descriptor) genNewRequest(sb *strings.Builder, hasQuery bool) {

	fmt.Fprintf(
		sb, "func (api *%s) newRequest(ctx context.Context, req %s) %s {\n",
		d.apiStructName(), d.requestTypeName(), "(*http.Request, error)")
	fmt.Fprint(sb, "\tURL, err := url.Parse(api.BaseURL)\n")
	fmt.Fprint(sb, "\tif err != nil {\n")
	fmt.Fprint(sb, "\t\treturn nil, err\n")
	fmt.Fprint(sb, "\t}\n")

	switch d.URLPath.IsTemplate {
	case false:
		fmt.Fprintf(sb, "\tURL.Path = \"%s\"\n", d.URLPath.Value)
	case true:
		fmt.Fprint(sb, "\tURL.Path, err = api.newRequestURLPath(req)\n")
		fmt.Fprint(sb, "\tif err != nil {\n")
		fmt.Fprint(sb, "\t\treturn nil, err\n")
		fmt.Fprint(sb, "\t}\n")
	}

	if hasQuery {
		fmt.Fprint(sb, "\tURL.RawQuery, err = api.newRequestQuery(req)\n")
		fmt.Fprint(sb, "\tif err != nil {\n")
		fmt.Fprint(sb, "\t\treturn nil, err\n")
		fmt.Fprint(sb, "\t}\n")
	}

	d.genNewRequestCallNewRequest(sb)

	fmt.Fprintf(sb, "}\n\n")
}

// GenNewRequest generates the code that creates a http.Request
// given a specific API call.
func (d *Descriptor) GenNewRequest() string {
	var sb strings.Builder
	hasQuery := d.genNewRequestQuery(&sb)
	d.genNewRequestURLPath(&sb)
	d.genNewRequest(&sb, hasQuery)
	return sb.String()
}
