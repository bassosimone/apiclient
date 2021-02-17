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
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestWithMissingAuthorizer(sb *strings.Builder) {
	if d.RequiresLogin == false {
		return
	}
	fmt.Fprintf(sb, "func Test%sWithMissingAuthorizer(t *testing.T) {\n", d.Name)
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	fmt.Fprintf(sb, "\t\tBaseURL: \"https://ps1.ooni.io\",\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, errMissingAuthorizer) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestWithHTTPErr(sb *strings.Builder) {
	fmt.Fprintf(sb, "func Test%sWithHTTPErr(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &MockableHTTPClient{Err: ErrMocked}\n")
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin == true {
		fmt.Fprint(sb, "\t\tAuthorizer:      newStaticAuthorizer(\"fakeToken\"),\n")
	}
	fmt.Fprint(sb, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmt.Fprint(sb, "\t\tHTTPClient: clnt,\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, ErrMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestMarshalErr(sb *strings.Builder) {
	if d.Method != "POST" {
		return
	}
	fmt.Fprintf(sb, "func Test%sMarshalErr(t *testing.T) {\n", d.Name)
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	fmt.Fprintf(sb, "\t\tBaseURL: \"https://ps1.ooni.io\",\n")
	fmt.Fprintf(sb, "\t\tmarshal: func(v interface{}) ([]byte, error) {\n")
	fmt.Fprintf(sb, "\t\t\treturn nil, ErrMocked\n")
	fmt.Fprintf(sb, "\t\t},\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, ErrMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestWithNewRequestErr(sb *strings.Builder) {
	fmt.Fprintf(sb, "func Test%sWithNewRequestErr(t *testing.T) {\n", d.Name)
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin == true {
		fmt.Fprint(sb, "\t\tAuthorizer:      newStaticAuthorizer(\"fakeToken\"),\n")
	}
	fmt.Fprint(sb, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmt.Fprint(sb, "\t\tNewRequest: func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {\n")
	fmt.Fprint(sb, "\t\t\treturn nil, ErrMocked\n")
	fmt.Fprint(sb, "\t\t},\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, ErrMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestWith400(sb *strings.Builder) {
	fmt.Fprintf(sb, "func Test%sWith400(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &MockableHTTPClient{Resp: &http.Response{StatusCode: 400}}\n")
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin == true {
		fmt.Fprint(sb, "\t\tAuthorizer:      newStaticAuthorizer(\"fakeToken\"),\n")
	}
	fmt.Fprint(sb, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmt.Fprint(sb, "\t\tHTTPClient: clnt,\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, ErrHTTPFailure) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestWithResponseBodyReadErr(sb *strings.Builder) {
	fmt.Fprintf(sb, "func Test%sWithResponseBodyReadErr(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &MockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &MockableBodyWithFailure{},\n")
	fmt.Fprint(sb, "\t}}\n")
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin == true {
		fmt.Fprint(sb, "\t\tAuthorizer:      newStaticAuthorizer(\"fakeToken\"),\n")
	}
	fmt.Fprint(sb, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmt.Fprint(sb, "\t\tHTTPClient: clnt,\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, ErrMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestWithUnmarshalFailure(sb *strings.Builder) {
	fmt.Fprintf(sb, "func Test%sWithUnmarshalFailure(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &MockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &MockableEmptyBody{},\n")
	fmt.Fprint(sb, "\t}}\n")
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin == true {
		fmt.Fprint(sb, "\t\tAuthorizer:      newStaticAuthorizer(\"fakeToken\"),\n")
	}
	fmt.Fprint(sb, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmt.Fprint(sb, "\t\tHTTPClient: clnt,\n")
	fmt.Fprintf(sb, "\t\tunmarshal: func(b []byte, v interface{}) error {\n")
	fmt.Fprintf(sb, "\t\t\treturn ErrMocked\n")
	fmt.Fprintf(sb, "\t\t},\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, ErrMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestRoundTrip(sb *strings.Builder) {
	fmt.Fprintf(sb, "func Test%sRoundTrip(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &MockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &MockableEmptyBody{},\n")
	fmt.Fprint(sb, "\t}}\n")
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin == true {
		fmt.Fprint(sb, "\t\tAuthorizer:      newStaticAuthorizer(\"fakeToken\"),\n")
	}
	fmt.Fprint(sb, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmt.Fprint(sb, "\t\tHTTPClient: clnt,\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif err != nil{\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp == nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected non-nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestResponseLiteralNull(sb *strings.Builder) {
	// TODO(bassosimone): factor out this snippter
	switch reflect.TypeOf(d.Response).Kind() {
	case reflect.Map:
	case reflect.Ptr:
		return // test not applicable
	default:
		panic("unsupported type")
	}
	fmt.Fprintf(sb, "func Test%sResponseLiteralNull(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &MockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &MockableLiteralNull{},\n")
	fmt.Fprint(sb, "\t}}\n")
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin == true {
		fmt.Fprint(sb, "\t\tAuthorizer:      newStaticAuthorizer(\"fakeToken\"),\n")
	}
	fmt.Fprint(sb, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmt.Fprint(sb, "\t\tHTTPClient: clnt,\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, ErrJSONLiteralNull) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestMandatoryFields(sb *strings.Builder) {
	fields := d.getStructFieldsWithTag(d.Request, tagForRequired)
	if len(fields) < 1 {
		return // nothing to test
	}
	fmt.Fprintf(sb, "func Test%sMandatoryFields(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &MockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &MockableLiteralNull{},\n")
	fmt.Fprint(sb, "\t}}\n")
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin == true {
		fmt.Fprint(sb, "\t\tAuthorizer:      newStaticAuthorizer(\"fakeToken\"),\n")
	}
	fmt.Fprint(sb, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmt.Fprint(sb, "\t\tHTTPClient: clnt,\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	fmt.Fprintf(sb, "\treq := &%s{} // deliberately empty\n", d.requestTypeNameAsStruct())
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, ErrEmptyField) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestTemplateParseErr(sb *strings.Builder) {
	if !d.URLPath.IsTemplate {
		return
	}
	fmt.Fprintf(sb, "func Test%sTemplateParseErr(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &MockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &MockableLiteralNull{},\n")
	fmt.Fprint(sb, "\t}}\n")
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin == true {
		fmt.Fprint(sb, "\t\tAuthorizer:      newStaticAuthorizer(\"fakeToken\"),\n")
	}
	fmt.Fprint(sb, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmt.Fprint(sb, "\t\tHTTPClient: clnt,\n")
	fmt.Fprint(sb, "\t\tnewTemplate: func(name string) textTemplate {\n")
	fmt.Fprint(sb, "\t\t\treturn &templateParseError{}\n")
	fmt.Fprint(sb, "\t\t},\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, ErrMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestTemplateExecuteErr(sb *strings.Builder) {
	if !d.URLPath.IsTemplate {
		return
	}
	fmt.Fprintf(sb, "func Test%sTemplateExecuteErr(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &MockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &MockableLiteralNull{},\n")
	fmt.Fprint(sb, "\t}}\n")
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin == true {
		fmt.Fprint(sb, "\t\tAuthorizer:      newStaticAuthorizer(\"fakeToken\"),\n")
	}
	fmt.Fprint(sb, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmt.Fprint(sb, "\t\tHTTPClient: clnt,\n")
	fmt.Fprint(sb, "\t\tnewTemplate: func(name string) textTemplate {\n")
	fmt.Fprint(sb, "\t\t\treturn &templateExecuteError{}\n")
	fmt.Fprint(sb, "\t\t},\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, ErrMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
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
	d.genTestWithMissingAuthorizer(&sb)
	d.genTestWithHTTPErr(&sb)
	d.genTestMarshalErr(&sb)
	d.genTestWithNewRequestErr(&sb)
	d.genTestWith400(&sb)
	d.genTestWithResponseBodyReadErr(&sb)
	d.genTestWithUnmarshalFailure(&sb)
	d.genTestRoundTrip(&sb)
	d.genTestResponseLiteralNull(&sb)
	d.genTestMandatoryFields(&sb)
	d.genTestTemplateParseErr(&sb)
	d.genTestTemplateExecuteErr(&sb)
	return sb.String()
}
