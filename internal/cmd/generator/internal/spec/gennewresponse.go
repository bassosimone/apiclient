package spec

import (
	"fmt"
	"reflect"
	"strings"
)

func (d *Descriptor) responseTypeKind() reflect.Kind {
	value := reflect.ValueOf(d.Response)
	if value.Type().Kind() == reflect.Ptr {
		if value.IsNil() {
			panic("null pointer")
		}
		value = value.Elem()
		if value.Type().Kind() != reflect.Struct {
			panic("not a struct")
		}
		return reflect.Struct
	}
	if value.Type().Kind() != reflect.Map {
		panic("not a map")
	}
	return reflect.Map
}

func (d *Descriptor) responseTypeNameAsStruct() string {
	value := reflect.ValueOf(d.Response)
	if value.Type().Kind() != reflect.Ptr {
		panic("not a pointer")
	}
	if value.IsNil() {
		panic("null pointer")
	}
	value = value.Elem()
	if value.Type().Kind() != reflect.Struct {
		panic("not a struct")
	}
	return value.Type().String()
}

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

	fmt.Fprint(&sb, "\tunmarshal := json.Unmarshal\n")
	fmt.Fprint(&sb, "\tif api.unmarshal != nil {\n")
	fmt.Fprint(&sb, "\t\tunmarshal = api.unmarshal\n")
	fmt.Fprint(&sb, "\t}\n")

	switch d.responseTypeKind() {
	case reflect.Map:
		fmt.Fprint(&sb, "\tif err := unmarshal(data, &out); err != nil {\n")
	case reflect.Struct:
		fmt.Fprint(&sb, "\tif err := unmarshal(data, out); err != nil {\n")
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
