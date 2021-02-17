package apimodel

import (
	"fmt"
	"reflect"
	"strings"
)

func (d *Descriptor) typeName(v interface{}) string {
	return reflect.ValueOf(v).Type().String()
}

func (d *Descriptor) requestTypeName() string {
	return d.typeName(d.Request)
}

func (d *Descriptor) responseTypeName() string {
	return d.typeName(d.Response)
}

// GenAPICall generates the Call method of the API.
func (d *Descriptor) GenAPICall() string {
	var sb strings.Builder
	fmt.Fprintf(
		&sb, "func (api *%s) call(ctx context.Context, in %s) (%s, error) {\n",
		d.apiStructName(), d.requestTypeName(), d.responseTypeName())

	fmt.Fprint(&sb, "\treq, err := api.newRequest(ctx, in)\n")
	fmt.Fprint(&sb, "\tif err != nil {\n")
	fmt.Fprint(&sb, "\t\treturn nil, err\n")
	fmt.Fprint(&sb, "\t}\n")
	fmt.Fprint(&sb, "\treq.Header.Add(\"Accept\", \"application/json\")\n")

	if d.RequiresLogin {
		fmt.Fprint(&sb, "\tif api.Authorizer == nil {\n")
		fmt.Fprint(&sb, "\t\treturn nil, errMissingAuthorizer\n")
		fmt.Fprint(&sb, "\t}\n")
		fmt.Fprint(&sb, "\ttoken, err := api.Authorizer.maybeRefreshToken(ctx)\n")
		fmt.Fprint(&sb, "\tif err != nil {\n")
		fmt.Fprint(&sb, "\t\treturn nil, err\n")
		fmt.Fprint(&sb, "\t}\n")
		fmt.Fprintf(&sb, "\tauthorization := newAuthorizationHeader(token)\n")
		fmt.Fprint(&sb, "\treq.Header.Add(\"Authorization\", authorization)\n")
	}

	fmt.Fprint(&sb, "\treq.Header.Add(\"User-Agent\", api.UserAgent)\n")
	fmt.Fprint(&sb, "\tvar httpClient HTTPClient = http.DefaultClient\n")
	fmt.Fprint(&sb, "\tif api.HTTPClient != nil {\n")
	fmt.Fprint(&sb, "\t\thttpClient = api.HTTPClient\n")
	fmt.Fprint(&sb, "\t}\n")
	fmt.Fprint(&sb, "\treturn api.newResponse(httpClient.Do(req))\n")

	fmt.Fprintf(&sb, "}\n\n")
	return sb.String()
}
