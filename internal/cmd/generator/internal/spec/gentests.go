package spec

import (
	"fmt"
	"reflect"
	"strings"
)

func (d *Descriptor) genTestNewRequest(sb *strings.Builder) {
	fields := d.getStructFieldsWithTag(d.Request, tagForRequired)
	if len(fields) > 0 {
		fmt.Fprintf(sb, "\treq := &%s{\n", d.requestTypeNameAsStruct())
		for idx, field := range fields {
			switch field.Type.Kind() {
			case reflect.String:
				fmt.Fprintf(sb, "\t\t%s: \"antani\",\n", field.Name)
			case reflect.Bool:
				fmt.Fprintf(sb, "\t\t%s: true,\n", field.Name)
			case reflect.Int64:
				fmt.Fprintf(sb, "\t\t%s: 123456789,\n", field.Name)
			default:
				panic(fmt.Sprintf("genTestNewRequest: unsupported field type: %d", idx))
			}
		}
		fmt.Fprint(sb, "\t}\n")
	} else {
		fmt.Fprintf(sb, "\treq := &%s{}\n", d.requestTypeNameAsStruct())
	}
}

func (d *Descriptor) genTestInvalidURL(sb *strings.Builder) {
	fmt.Fprintf(sb, "func Test%sInvalidURL(t *testing.T) {\n", d.Name)
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	fmt.Fprintf(sb, "\t\tBaseURL: \"\\t\", // invalid\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif err == nil || !strings.HasSuffix(err.Error(), \"invalid control character in URL\") {\n")
	fmt.Fprintf(sb, "\t\tt.Fatalf(\"not the error we expected: %%+v\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

// GenTests generates tests for generated code.
func (d *Descriptor) GenTests() string {
	var sb strings.Builder
	d.genTestInvalidURL(&sb)
	return sb.String()
}
