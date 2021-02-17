// Command gencall generates callapi.go.
package main

import (
	"strings"
	"time"

	"github.com/bassosimone/apiclient/internal/apimodel"
	"github.com/bassosimone/apiclient/internal/fmtx"
	"github.com/bassosimone/apiclient/internal/osx"
	"github.com/bassosimone/apiclient/internal/reflectx"
	"github.com/bassosimone/apiclient/internal/strcasex"
)

func getapiame(in interface{}) string {
	name := reflectx.Must(reflectx.NewTypeValueInfo(in)).TypeName()
	name = strings.Replace(name, "Request", "", 1)
	name = strings.Replace(name, "Response", "", 1)
	return strcasex.ToLowerCamel(name)
}

func genapitype(filep osx.File, desc *apimodel.Descriptor) {
	apiname := getapiame(desc.Response)
	fmtx.Fprintf(filep, "// %sAPI is the %s API. The zero-value structure\n", apiname, apiname)
	if desc.RequiresLogin {
		fmtx.Fprint(filep, "// is not valid because Authorizer is always required. We use\n")
		fmtx.Fprint(filep, "// suitable defaults for any other zero-initialized field.\n")
	} else {
		fmtx.Fprint(filep, "// works as intended using suitable default values.\n")
	}
	fmtx.Fprintf(filep, "type %sAPI struct {\n", apiname)
	if desc.RequiresLogin {
		fmtx.Fprint(filep, "\tAuthorizer authorizer\n")
	}
	fmtx.Fprint(filep, "\tBaseURL     string\n")
	fmtx.Fprint(filep, "\tHTTPClient  HTTPClient\n")
	fmtx.Fprint(filep, "\tNewRequest  func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)\n")
	fmtx.Fprint(filep, "\tUserAgent   string\n")
	fmtx.Fprint(filep, "\tmarshal     func(v interface{}) ([]byte, error)\n")
	if desc.URLPath.IsTemplate {
		fmtx.Fprint(filep, "\tnewTemplate func(s string) textTemplate\n")
	}
	fmtx.Fprint(filep, "\tunmarshal   func(b []byte, v interface{}) error\n")
	fmtx.Fprint(filep, "}\n\n")
}

func genbeginfunc(filep osx.File, desc *apimodel.Descriptor) {
	resp := reflectx.Must(reflectx.NewTypeValueInfo(desc.Response))
	req := reflectx.Must(reflectx.NewTypeValueInfo(desc.Request)).TypeName()
	apiname := getapiame(desc.Response)
	fmtx.Fprintf(filep, "// Call calls %s %s. Arguments MUST NOT be nil. The return\n", desc.Method, desc.URLPath.Value)
	fmtx.Fprint(filep, "// value is either a non-nil error or a non-nil result.\n")
	fmtx.Fprintf(filep, "func (api %sAPI) Call", apiname)
	fmtx.Fprintf(filep, "(ctx context.Context, in *%s)", req)
	fmtx.Fprintf(filep, " (%s, error) {\n", resp.AsReturnType())
}

func gencall(filep osx.File, desc *apimodel.Descriptor) {
	fmtx.Fprint(filep, "\treq, err := api.newRequest(ctx, api.BaseURL, in)\n")
	fmtx.Fprint(filep, "\tif err != nil {\n")
	fmtx.Fprint(filep, "\t\treturn nil, err\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\treq.Header.Add(\"Accept\", \"application/json\")\n")
	if desc.RequiresLogin {
		fmtx.Fprint(filep, "\tif api.Authorizer == nil {\n")
		fmtx.Fprint(filep, "\t\treturn nil, errMissingAuthorizer\n")
		fmtx.Fprint(filep, "\t}\n")
		fmtx.Fprint(filep, "\ttoken, err := api.Authorizer.maybeRefreshToken(ctx)\n")
		fmtx.Fprint(filep, "\tif err != nil {\n")
		fmtx.Fprint(filep, "\t\treturn nil, err\n")
		fmtx.Fprint(filep, "\t}\n")
		fmtx.Fprintf(filep, "\tauthorization := fmt.Sprintf(\"Bearer %%s\", token)\n")
		fmtx.Fprint(filep, "\treq.Header.Add(\"Authorization\", authorization)\n")
	}
	fmtx.Fprint(filep, "\treq.Header.Add(\"User-Agent\", api.UserAgent)\n")
	fmtx.Fprint(filep, "\tvar httpClient HTTPClient = http.DefaultClient\n")
	fmtx.Fprint(filep, "\tif api.HTTPClient != nil {\n")
	fmtx.Fprint(filep, "\t\thttpClient = api.HTTPClient\n")
	fmtx.Fprint(filep, "\t}\n")
	fmtx.Fprint(filep, "\treturn api.newResponse(httpClient.Do(req))\n")
}

func genendfunc(filep osx.File) {
	fmtx.Fprintf(filep, "}\n\n")
}

func genapi(filep osx.File, desc *apimodel.Descriptor) {
	genapitype(filep, desc)
	genbeginfunc(filep, desc)
	gencall(filep, desc)
	genendfunc(filep)
}

func main() {
	filep := osx.MustCreate("callapi.go")
	defer filep.Close()

	fmtx.Fprint(filep, "// Code generated by go generate; DO NOT EDIT.\n")
	fmtx.Fprintf(filep, "// %v\n\n", time.Now())
	fmtx.Fprint(filep, "package apiclient\n\n")
	fmtx.Fprint(filep, "import (\n")
	fmtx.Fprint(filep, "\t\"context\"\n")
	fmtx.Fprint(filep, "\t\"fmt\"\n")
	fmtx.Fprint(filep, "\t\"io\"\n")
	fmtx.Fprint(filep, "\t\"net/http\"\n")
	fmtx.Fprint(filep, ")\n\n")

	fmtx.Fprint(filep, "//go:generate go run ./internal/gencallapi/...\n\n")

	for _, descr := range apimodel.Descriptors {
		genapi(filep, &descr)
	}
}
