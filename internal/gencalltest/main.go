// Command gencalltest generates call_test.go files
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

func getapiame(in interface{}) string {
	name := reflectx.Must(reflectx.NewTypeValueInfo(in)).TypeName()
	name = strings.Replace(name, "Request", "", 1)
	name = strings.Replace(name, "Response", "", 1)
	return name
}

func genTestInvalidURL(filep osx.File, desc *apimodel.Descriptor) {
	apiname := getapiame(desc.Response)
	fmtx.Fprintf(filep, "func Test%sInvalidURL(t *testing.T) {\n", apiname)
	fmtx.Fprintf(filep, "\tapi := &%sAPI{\n", apiname)
	fmtx.Fprintf(filep, "\t\tBaseURL: \"\\t\", // invalid\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tctx := context.Background()\n")
	fmtx.Fprintf(filep, "\treq := &%sRequest{}\n", apiname)
	fmtx.Fprint(filep, "\tresp, err := api.Call(ctx, req)\n")
	fmtx.Fprint(filep, "\tif err == nil || !strings.HasSuffix(err.Error(), \"invalid control character in URL\") {\n")
	fmtx.Fprintf(filep, "\t\tt.Fatalf(\"not the error we expected: %%+v\", err)\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tif resp != nil {\n")
	fmtx.Fprint(filep, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "}\n\n")
}

func genTestWithEmptyToken(filep osx.File, desc *apimodel.Descriptor) {
	if desc.RequiresLogin == false {
		return
	}
	apiname := getapiame(desc.Response)
	fmtx.Fprintf(filep, "func Test%sWithEmptyToken(t *testing.T) {\n", apiname)
	fmtx.Fprintf(filep, "\tapi := &%sAPI{\n", apiname)
	fmtx.Fprintf(filep, "\t\tBaseURL: \"https://ps1.ooni.io\",\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tctx := context.Background()\n")
	fmtx.Fprintf(filep, "\treq := &%sRequest{}\n", apiname)
	fmtx.Fprint(filep, "\tresp, err := api.Call(ctx, req)\n")
	fmtx.Fprint(filep, "\tif !errors.Is(err, ErrEmptyToken) {\n")
	fmtx.Fprintf(filep, "\t\tt.Fatalf(\"not the error we expected: %%+v\", err)\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tif resp != nil {\n")
	fmtx.Fprint(filep, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "}\n\n")
}

func genTestWithHTTPErr(filep osx.File, desc *apimodel.Descriptor) {
	req := reflectx.Must(reflectx.NewTypeValueInfo(desc.Request))
	apiname := getapiame(desc.Response)
	fmtx.Fprintf(filep, "func Test%sWithHTTPErr(t *testing.T) {\n", apiname)
	fmtx.Fprint(filep, "\tclnt := &MockableHTTPClient{Err: ErrMocked}\n")
	fmtx.Fprintf(filep, "\tapi := &%sAPI{\n", apiname)
	fmtx.Fprint(filep, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmtx.Fprint(filep, "\t\tHTTPClient: clnt,\n")
	if desc.RequiresLogin == true {
		fmtx.Fprint(filep, "\t\tToken:      \"fakeToken\",\n")
	}
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tctx := context.Background()\n")
	fields, err := req.AllFieldsWithTag("required")
	fatalx.OnError(err, "req.AllFieldsWithTag failed")
	if len(fields) > 0 {
		fmtx.Fprintf(filep, "\treq := &%sRequest{\n", apiname)
		for _, field := range fields {
			switch field.Self.Type.Kind() {
			case reflect.String:
				fmtx.Fprintf(filep, "\t\t%s: \"antani\",\n", field.Self.Name)
			case reflect.Bool:
				fmtx.Fprintf(filep, "\t\t%s: true,\n", field.Self.Name)
			case reflect.Int64:
				fmtx.Fprintf(filep, "\t\t%s: 123456789,\n", field.Self.Name)
			default:
				panic("genTestWithHTTPErr: unsupported field type")
			}
		}
		fmtx.Fprint(filep, "\t}\n")
	} else {
		fmtx.Fprintf(filep, "\treq := &%sRequest{}\n", apiname)
	}
	fmtx.Fprint(filep, "\tresp, err := api.Call(ctx, req)\n")
	fmtx.Fprint(filep, "\tif !errors.Is(err, ErrMocked) {\n")
	fmtx.Fprintf(filep, "\t\tt.Fatalf(\"not the error we expected: %%+v\", err)\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tif resp != nil {\n")
	fmtx.Fprint(filep, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "}\n\n")
}

func genTestMarshalErr(filep osx.File, desc *apimodel.Descriptor) {
	if desc.Method != "POST" {
		return
	}
	apiname := getapiame(desc.Response)
	fmtx.Fprintf(filep, "func Test%sMarshalErr(t *testing.T) {\n", apiname)
	fmtx.Fprintf(filep, "\tapi := &%sAPI{\n", apiname)
	fmtx.Fprintf(filep, "\t\tBaseURL: \"https://ps1.ooni.io\",\n")
	fmtx.Fprintf(filep, "\t\tmarshal: func(v interface{}) ([]byte, error) {\n")
	fmtx.Fprintf(filep, "\t\t\treturn nil, ErrMocked\n")
	fmtx.Fprintf(filep, "\t\t},\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tctx := context.Background()\n")
	fmtx.Fprintf(filep, "\treq := &%sRequest{}\n", apiname)
	fmtx.Fprint(filep, "\tresp, err := api.Call(ctx, req)\n")
	fmtx.Fprint(filep, "\tif !errors.Is(err, ErrMocked) {\n")
	fmtx.Fprintf(filep, "\t\tt.Fatalf(\"not the error we expected: %%+v\", err)\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tif resp != nil {\n")
	fmtx.Fprint(filep, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "}\n\n")
}

func genTestWithNewRequestErr(filep osx.File, desc *apimodel.Descriptor) {
	req := reflectx.Must(reflectx.NewTypeValueInfo(desc.Request))
	apiname := getapiame(desc.Response)
	fmtx.Fprintf(filep, "func Test%sWithNewRequestErr(t *testing.T) {\n", apiname)
	fmtx.Fprintf(filep, "\tapi := &%sAPI{\n", apiname)
	fmtx.Fprint(filep, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmtx.Fprint(filep, "\t\tNewRequest: func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {\n")
	fmtx.Fprint(filep, "\t\t\treturn nil, ErrMocked\n")
	fmtx.Fprint(filep, "\t\t},\n")
	if desc.RequiresLogin == true {
		fmtx.Fprint(filep, "\t\tToken:      \"fakeToken\",\n")
	}
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tctx := context.Background()\n")
	fields, err := req.AllFieldsWithTag("required")
	fatalx.OnError(err, "req.AllFieldsWithTag failed")
	if len(fields) > 0 {
		fmtx.Fprintf(filep, "\treq := &%sRequest{\n", apiname)
		for _, field := range fields {
			switch field.Self.Type.Kind() {
			case reflect.String:
				fmtx.Fprintf(filep, "\t\t%s: \"antani\",\n", field.Self.Name)
			case reflect.Bool:
				fmtx.Fprintf(filep, "\t\t%s: true,\n", field.Self.Name)
			case reflect.Int64:
				fmtx.Fprintf(filep, "\t\t%s: 123456789,\n", field.Self.Name)
			default:
				panic("genTestWithHTTPErr: unsupported field type")
			}
		}
		fmtx.Fprint(filep, "\t}\n")
	} else {
		fmtx.Fprintf(filep, "\treq := &%sRequest{}\n", apiname)
	}
	fmtx.Fprint(filep, "\tresp, err := api.Call(ctx, req)\n")
	fmtx.Fprint(filep, "\tif !errors.Is(err, ErrMocked) {\n")
	fmtx.Fprintf(filep, "\t\tt.Fatalf(\"not the error we expected: %%+v\", err)\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tif resp != nil {\n")
	fmtx.Fprint(filep, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "}\n\n")
}

func genTestWith400(filep osx.File, desc *apimodel.Descriptor) {
	req := reflectx.Must(reflectx.NewTypeValueInfo(desc.Request))
	apiname := getapiame(desc.Response)
	fmtx.Fprintf(filep, "func Test%sWith400(t *testing.T) {\n", apiname)
	fmtx.Fprint(filep, "\tclnt := &MockableHTTPClient{Resp: &http.Response{StatusCode: 400}}\n")
	fmtx.Fprintf(filep, "\tapi := &%sAPI{\n", apiname)
	fmtx.Fprint(filep, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmtx.Fprint(filep, "\t\tHTTPClient: clnt,\n")
	if desc.RequiresLogin == true {
		fmtx.Fprint(filep, "\t\tToken:      \"fakeToken\",\n")
	}
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tctx := context.Background()\n")
	fields, err := req.AllFieldsWithTag("required")
	fatalx.OnError(err, "req.AllFieldsWithTag failed")
	if len(fields) > 0 {
		fmtx.Fprintf(filep, "\treq := &%sRequest{\n", apiname)
		for _, field := range fields {
			switch field.Self.Type.Kind() {
			case reflect.String:
				fmtx.Fprintf(filep, "\t\t%s: \"antani\",\n", field.Self.Name)
			case reflect.Bool:
				fmtx.Fprintf(filep, "\t\t%s: true,\n", field.Self.Name)
			case reflect.Int64:
				fmtx.Fprintf(filep, "\t\t%s: 123456789,\n", field.Self.Name)
			default:
				panic("genTestWithHTTPErr: unsupported field type")
			}
		}
		fmtx.Fprint(filep, "\t}\n")
	} else {
		fmtx.Fprintf(filep, "\treq := &%sRequest{}\n", apiname)
	}
	fmtx.Fprint(filep, "\tresp, err := api.Call(ctx, req)\n")
	fmtx.Fprint(filep, "\tif !errors.Is(err, ErrHTTPFailure) {\n")
	fmtx.Fprintf(filep, "\t\tt.Fatalf(\"not the error we expected: %%+v\", err)\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\tif resp != nil {\n")
	fmtx.Fprint(filep, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "}\n\n")
}

func genapi(filep osx.File, desc *apimodel.Descriptor) {
	genTestInvalidURL(filep, desc)
	genTestWithEmptyToken(filep, desc)
	genTestWithHTTPErr(filep, desc)
	genTestMarshalErr(filep, desc)
	genTestWithNewRequestErr(filep, desc)
	genTestWith400(filep, desc)
}

func main() {
	filep := osx.MustCreate("call_test.go")
	defer filep.Close()

	fmtx.Fprint(filep, "// Code generated by go generate; DO NOT EDIT.\n")
	fmtx.Fprintf(filep, "// %v\n\n", time.Now())
	fmtx.Fprint(filep, "package apiclient\n\n")
	fmtx.Fprint(filep, "import (\n")
	fmtx.Fprint(filep, "\t\"context\"\n")
	fmtx.Fprint(filep, "\t\"errors\"\n")
	fmtx.Fprint(filep, "\t\"io\"\n")
	fmtx.Fprint(filep, "\t\"net/http\"\n")
	fmtx.Fprint(filep, "\t\"strings\"\n")
	fmtx.Fprint(filep, "\t\"testing\"\n")
	fmtx.Fprint(filep, ")\n\n")

	fmtx.Fprint(filep, "//go:generate go run ./internal/gencalltest/...\n\n")

	for _, descr := range apimodel.Descriptors {
		genapi(filep, &descr)
	}
}
