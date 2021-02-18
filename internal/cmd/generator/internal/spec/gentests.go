package spec

import (
	"fmt"
	"reflect"
	"strings"
)

func (d *Descriptor) genTestNewRequest(sb *strings.Builder) {
	fmt.Fprintf(sb, "\treq := &%s{}\n", d.requestTypeNameAsStruct())
	fmt.Fprint(sb, "\tff := &fakeFill{}\n")
	fmt.Fprint(sb, "\tff.fill(req)\n")
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
		return // does not make sense when login isn't required
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
	fmt.Fprint(sb, "\tclnt := &mockableHTTPClient{Err: errMocked}\n")
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
	fmt.Fprint(sb, "\tif !errors.Is(err, errMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestMarshalErr(sb *strings.Builder) {
	if d.Method != "POST" {
		return // does not make sense when we don't send a request body
	}
	fmt.Fprintf(sb, "func Test%sMarshalErr(t *testing.T) {\n", d.Name)
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	fmt.Fprintf(sb, "\t\tBaseURL: \"https://ps1.ooni.io\",\n")
	fmt.Fprintf(sb, "\t\tmarshal: func(v interface{}) ([]byte, error) {\n")
	fmt.Fprintf(sb, "\t\t\treturn nil, errMocked\n")
	fmt.Fprintf(sb, "\t\t},\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, errMocked) {\n")
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
	fmt.Fprint(sb, "\t\t\treturn nil, errMocked\n")
	fmt.Fprint(sb, "\t\t},\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, errMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestWith400(sb *strings.Builder) {
	fmt.Fprintf(sb, "func Test%sWith400(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &mockableHTTPClient{Resp: &http.Response{StatusCode: 400}}\n")
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
	fmt.Fprint(sb, "\tclnt := &mockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &mockableBodyWithFailure{},\n")
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
	fmt.Fprint(sb, "\tif !errors.Is(err, errMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestWithUnmarshalFailure(sb *strings.Builder) {
	fmt.Fprintf(sb, "func Test%sWithUnmarshalFailure(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &mockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &mockableEmptyBody{},\n")
	fmt.Fprint(sb, "\t}}\n")
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin == true {
		fmt.Fprint(sb, "\t\tAuthorizer:      newStaticAuthorizer(\"fakeToken\"),\n")
	}
	fmt.Fprint(sb, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmt.Fprint(sb, "\t\tHTTPClient: clnt,\n")
	fmt.Fprintf(sb, "\t\tunmarshal: func(b []byte, v interface{}) error {\n")
	fmt.Fprintf(sb, "\t\t\treturn errMocked\n")
	fmt.Fprintf(sb, "\t\t},\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, errMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestRoundTrip(sb *strings.Builder) {
	fmt.Fprintf(sb, "func Test%sRoundTrip(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &mockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &mockableEmptyBody{},\n")
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
	switch d.responseTypeKind() {
	case reflect.Map:
		// fallthrough
	case reflect.Struct:
		return // test not applicable
	}
	fmt.Fprintf(sb, "func Test%sResponseLiteralNull(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &mockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &mockableLiteralNull{},\n")
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
	fmt.Fprint(sb, "\tclnt := &mockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &mockableLiteralNull{},\n")
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
		return // nothing to test
	}
	fmt.Fprintf(sb, "func Test%sTemplateParseErr(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &mockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &mockableLiteralNull{},\n")
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
	fmt.Fprint(sb, "\tif !errors.Is(err, errMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestTemplateExecuteErr(sb *strings.Builder) {
	if !d.URLPath.IsTemplate {
		return // nothing to test
	}
	fmt.Fprintf(sb, "func Test%sTemplateExecuteErr(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &mockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &mockableLiteralNull{},\n")
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
	fmt.Fprint(sb, "\tif !errors.Is(err, errMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestWithFailingAuthorizer(sb *strings.Builder) {
	if !d.RequiresLogin {
		return // nothing to test
	}
	fmt.Fprintf(sb, "func Test%sWithFailingAuthorizer(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\tclnt := &mockableHTTPClient{Resp: &http.Response{\n")
	fmt.Fprint(sb, "\t\tStatusCode: 200,\n")
	fmt.Fprint(sb, "\t\tBody: &mockableEmptyBody{},\n")
	fmt.Fprint(sb, "\t}}\n")
	fmt.Fprintf(sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin == true {
		fmt.Fprint(sb, "\t\tAuthorizer:      &failingAuthorizer{},\n")
	}
	fmt.Fprint(sb, "\t\tBaseURL:    \"https://ps1.ooni.io\",\n")
	fmt.Fprint(sb, "\t\tHTTPClient: clnt,\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tresp, err := api.call(ctx, req)\n")
	fmt.Fprint(sb, "\tif !errors.Is(err, errMocked) {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(\"not the error we expected\", err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tif resp != nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "}\n\n")
}

// TODO(bassosimone): we should add a panic for every switch for
// the type of a request or a response for robustness.

func (d *Descriptor) genHandlerForPublicAPI(sb *strings.Builder) {
	if d.Private {
		return // we only test public APIs here
	}
	fmt.Fprintf(sb, "type handle%s struct{\n", d.Name)
	if d.Method == "POST" {
		switch d.requestTypeKind() {
		case reflect.Struct:
			fmt.Fprintf(sb, "\treq %s\n", d.requestTypeName())
		default:
			panic("not supporting this case")
		}
	}
	fmt.Fprint(sb, "\tcount int32\n")
	fmt.Fprint(sb, "\tmethod string\n")
	fmt.Fprint(sb, "\tmu sync.Mutex\n")
	fmt.Fprintf(sb, "\tresp %s\n", d.responseTypeName())
	fmt.Fprint(sb, "\turl *url.URL\n")
	fmt.Fprint(sb, "}\n\n")
	fmt.Fprintf(sb, "func (h *handle%s) ServeHTTP(w http.ResponseWriter, r *http.Request) {\n", d.Name)
	fmt.Fprint(sb, "\tdefer h.mu.Unlock()\n")
	fmt.Fprint(sb, "\th.mu.Lock()\n")
	fmt.Fprint(sb, "\tif h.count > 0 {\n")
	fmt.Fprint(sb, "\t\tw.WriteHeader(400)\n")
	fmt.Fprint(sb, "\t\treturn\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\th.count++\n")
	fmt.Fprint(sb, "\th.method = r.Method\n")
	fmt.Fprint(sb, "\th.url = r.URL\n")
	if d.Method == "POST" {
		fmt.Fprint(sb, "\treqbody, err := ioutil.ReadAll(r.Body)\n")
		fmt.Fprint(sb, "\tif err != nil {\n")
		fmt.Fprint(sb, "\t\tw.WriteHeader(400)\n")
		fmt.Fprint(sb, "\t\treturn\n")
		fmt.Fprint(sb, "\t}\n")
		switch d.requestTypeKind() {
		case reflect.Struct:
			fmt.Fprintf(sb, "\tvar in %s\n", d.requestTypeNameAsStruct())
		default:
			panic("not supporting this case")
		}
		fmt.Fprint(sb, "\tif err := json.Unmarshal(reqbody, &in); err != nil {\n")
		fmt.Fprint(sb, "\t\tw.WriteHeader(400)\n")
		fmt.Fprint(sb, "\t\treturn\n")
		fmt.Fprint(sb, "\t}\n")
		fmt.Fprint(sb, "\tif reflect.ValueOf(in).IsZero() {\n")
		fmt.Fprint(sb, "\t\tw.WriteHeader(400)\n")
		fmt.Fprint(sb, "\t\treturn\n")
		fmt.Fprint(sb, "\t}\n")
		fmt.Fprint(sb, "\th.req = &in\n")
	}
	switch d.responseTypeKind() {
	case reflect.Struct:
		fmt.Fprintf(sb, "\tvar out %s\n", d.responseTypeNameAsStruct())
	case reflect.Map:
		fmt.Fprintf(sb, "\tvar out %s\n", d.responseTypeName())
	}
	fmt.Fprint(sb, "\tff := fakeFill{}\n")
	fmt.Fprint(sb, "\tff.fill(&out)\n")
	switch d.responseTypeKind() {
	case reflect.Struct:
		fmt.Fprint(sb, "\th.resp = &out\n")
	case reflect.Map:
		fmt.Fprint(sb, "\th.resp = out\n")
	}
	fmt.Fprint(sb, "\tdata, err := json.Marshal(out)\n")
	fmt.Fprint(sb, "\tif err != nil {\n")
	fmt.Fprint(sb, "\t\tw.WriteHeader(400)\n")
	fmt.Fprint(sb, "\t\treturn\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\tw.Write(data)\n")
	fmt.Fprint(sb, "}\n\n")
}

func (d *Descriptor) genTestClientWithHandlerForPublicAPI(sb *strings.Builder) {
	if d.Private {
		return // we only test public APIs here
	}
	fmt.Fprintf(sb, "func TestClientWithHandlerFor%s(t *testing.T) {\n", d.Name)
	fmt.Fprint(sb, "\t// setup\n")
	fmt.Fprintf(sb, "\thandler := &handle%s{}\n", d.Name)
	fmt.Fprint(sb, "\tsrvr := httptest.NewServer(handler)\n")
	fmt.Fprint(sb, "\tdefer srvr.Close()\n")
	d.genTestNewRequest(sb)
	fmt.Fprint(sb, "\tclnt := &Client{\n")
	fmt.Fprint(sb, "\t\tBaseURL: srvr.URL,\n")
	fmt.Fprint(sb, "\t}\n")
	if d.RequiresLogin == true {
		fmt.Fprint(sb, "\tkvstore := &memkvstore{}\n")
		fmt.Fprint(sb, "\t// hand-craft a state that does not require relogin\n")
		fmt.Fprint(sb, "\tlm := &loginManager{kvstore: kvstore, state: loginState{\n")
		fmt.Fprint(sb, "\t\tClientID: \"077c3985-b228-4df3-af22-bc3377c7a376\",\n")
		fmt.Fprint(sb, "\t\tExpire: time.Now().Add(3600*time.Second),\n")
		fmt.Fprint(sb, "\t\tPassword: \"077c3985-b228-4df3-af22-bc3377c7a376\",\n")
		fmt.Fprint(sb, "\t\tToken: \"077c3985-b228-4df3-af22-bc3377c7a376\",\n")
		fmt.Fprint(sb, "\t}}\n")
		fmt.Fprint(sb, "\tlm.writeback() // memory never fails\n")
		fmt.Fprint(sb, "\tclnt.KVStore = kvstore\n")
	}
	fmt.Fprint(sb, "\t// issue request\n")
	fmt.Fprint(sb, "\tctx := context.Background()\n")
	fmt.Fprintf(sb, "\tresp, err := clnt.%s(ctx, req)\n", d.Name)
	fmt.Fprint(sb, "\tif err != nil {\n")
	fmt.Fprintf(sb, "\t\tt.Fatal(err)\n")
	fmt.Fprint(sb, "\t}\n")
	fmt.Fprint(sb, "\t// check for data round trip\n")
	fmt.Fprint(sb, "\tif resp == nil {\n")
	fmt.Fprint(sb, "\t\tt.Fatal(\"expected non-nil resp\")\n")
	fmt.Fprint(sb, "\t}\n")
	switch d.responseTypeKind() {
	case reflect.Struct:
		fmt.Fprint(sb, "\tif reflect.ValueOf(*resp).IsZero() {\n")
		fmt.Fprint(sb, "\t\tt.Fatal(\"server returned a zero structure\")\n")
		fmt.Fprint(sb, "\t}\n")
	case reflect.Map:
		// nothing
	}
	if d.Method == "POST" {
		fmt.Fprint(sb, "\tif diff := cmp.Diff(req, handler.req); diff != \"\"{\n")
		fmt.Fprint(sb, "\t\tt.Fatal(diff)\n")
		fmt.Fprint(sb, "\t}\n")
	}
	fmt.Fprint(sb, "\tif diff := cmp.Diff(resp, handler.resp); diff != \"\"{\n")
	fmt.Fprint(sb, "\t\tt.Fatal(diff)\n")
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
	d.genTestWithFailingAuthorizer(&sb)
	d.genHandlerForPublicAPI(&sb)
	d.genTestClientWithHandlerForPublicAPI(&sb)
	return sb.String()
}
