package spec

import (
	"fmt"
	"reflect"
	"strings"
)

// GenNewResponse generates the code that creates a response
// given the result of a specific HTTP round trip.
func (d *Descriptor) GenNewResponse() string {
	var sb strings.Builder
	fmt.Fprintf(&sb,
		"func (api *%s) newResponse(resp *http.Response, err error) (%s, error) {\n",
		d.apiStructName(), d.responseTypeName())

	fmt.Fprint(&sb, "\tif err != nil {\n")
	fmt.Fprint(&sb, "\t\treturn nil, err\n")
	fmt.Fprint(&sb, "\t}\n")
	fmt.Fprint(&sb, "\tif resp.StatusCode == 401 {\n")
	fmt.Fprint(&sb, "\t\treturn nil, ErrUnauthorized\n")
	fmt.Fprint(&sb, "\t}\n")
	fmt.Fprint(&sb, "\tif resp.StatusCode != 200 {\n")
	fmt.Fprint(&sb, "\t\treturn nil, newHTTPFailure(resp.StatusCode)\n")
	fmt.Fprint(&sb, "\t}\n")
	fmt.Fprint(&sb, "\tdefer resp.Body.Close()\n")
	fmt.Fprint(&sb, "\treader := io.LimitReader(resp.Body, 4<<20)\n")
	fmt.Fprint(&sb, "\tdata, err := ioutil.ReadAll(reader)\n")
	fmt.Fprint(&sb, "\tif err != nil {\n")
	fmt.Fprint(&sb, "\t\treturn nil, err\n")
	fmt.Fprint(&sb, "\t}\n")

	switch d.responseTypeKind() {
	case reflect.Map:
		fmt.Fprintf(&sb, "\tout := %s{}\n", d.responseTypeName())
	case reflect.Struct:
		fmt.Fprintf(&sb, "\tout := &%s{}\n", d.responseTypeNameAsStruct())
	}

	switch d.responseTypeKind() {
	case reflect.Map:
		fmt.Fprint(&sb, "\tif err := api.jsonCodec.Decode(data, &out); err != nil {\n")
	case reflect.Struct:
		fmt.Fprint(&sb, "\tif err := api.jsonCodec.Decode(data, out); err != nil {\n")
	}

	fmt.Fprint(&sb, "\t\treturn nil, err\n")
	fmt.Fprint(&sb, "\t}\n")

	switch d.responseTypeKind() {
	case reflect.Map:
		// For rationale, see https://play.golang.org/p/m9-MsTaQ5wt and
		// https://play.golang.org/p/6h-v-PShMk9.
		fmt.Fprint(&sb, "\tif out == nil {\n")
		fmt.Fprint(&sb, "\t\treturn nil, ErrJSONLiteralNull\n")
		fmt.Fprint(&sb, "\t}\n")
	case reflect.Struct:
		// nothing
	}
	fmt.Fprintf(&sb, "\treturn out, nil\n")
	fmt.Fprintf(&sb, "}\n\n")
	return sb.String()
}
