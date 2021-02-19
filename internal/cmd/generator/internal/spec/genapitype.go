package spec

import (
	"fmt"
	"strings"
)

// GenAPIType generates the type definition for the API.
func (d *Descriptor) GenAPIType() string {
	var sb strings.Builder

	// generate the struct itself
	fmt.Fprintf(&sb, "type %s struct {\n", d.apiStructName())
	fmt.Fprint(&sb, "\tBaseURL     string\n")
	fmt.Fprint(&sb, "\tHTTPClient  HTTPClient\n")
	fmt.Fprint(&sb, "\tNewRequest  func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)\n")
	if d.RequiresLogin {
		fmt.Fprintf(&sb, "\tToken string\n")
	}
	fmt.Fprint(&sb, "\tUserAgent   string\n")
	fmt.Fprint(&sb, "\tmarshal     func(v interface{}) ([]byte, error)\n")
	if d.URLPath.IsTemplate {
		fmt.Fprint(&sb, "\tnewTemplate func(s string) textTemplate\n")
	}
	fmt.Fprint(&sb, "\tunmarshal   func(b []byte, v interface{}) error\n")
	fmt.Fprint(&sb, "}\n\n")

	// generate the newAPI factory
	if d.RequiresLogin {
		fmt.Fprintf(&sb, "func new%sAPI(c *Client, token string) *%s {\n", d.Name, d.apiStructName())
	} else {
		fmt.Fprintf(&sb, "func new%sAPI(c *Client) *%s {\n", d.Name, d.apiStructName())
	}
	fmt.Fprintf(&sb, "\tvar clnt HTTPClient = c.httpClient()\n")
	if d.Cache {
		fmt.Fprintf(&sb, "\tclnt = &cacheClient{\n")
		fmt.Fprintf(&sb, "\t\tHTTPClient: clnt,\n")
		fmt.Fprintf(&sb, "\t\tKVStore: c.kvstore(),\n")
		fmt.Fprintf(&sb, "\t}\n")
	}
	fmt.Fprintf(&sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin {
		fmt.Fprintf(&sb, "\t\tToken: token,\n")
	}
	fmt.Fprintf(&sb, "\t\tBaseURL: c.baseURL(),\n")
	fmt.Fprintf(&sb, "\t\tHTTPClient: clnt,\n")
	fmt.Fprintf(&sb, "\t\tUserAgent: c.UserAgent,\n")
	fmt.Fprintf(&sb, "\t}\n")
	fmt.Fprintf(&sb, "\treturn api\n")
	fmt.Fprint(&sb, "}\n\n")

	// generate the API call method
	fmt.Fprintf(
		&sb, "func (api *%s) call(ctx context.Context, in %s) (%s, error) {\n",
		d.apiStructName(), d.requestTypeName(), d.responseTypeName())
	fmt.Fprint(&sb, "\treq, err := api.newRequest(ctx, in)\n")
	fmt.Fprint(&sb, "\tif err != nil {\n")
	fmt.Fprint(&sb, "\t\treturn nil, err\n")
	fmt.Fprint(&sb, "\t}\n")
	fmt.Fprint(&sb, "\treq.Header.Add(\"Accept\", \"application/json\")\n")
	if d.RequiresLogin {
		fmt.Fprint(&sb, "\tif api.Token == \"\" {\n")
		fmt.Fprint(&sb, "\t\treturn nil, errMissingToken\n")
		fmt.Fprint(&sb, "\t}\n")
		fmt.Fprintf(&sb, "\tauthorization := newAuthorizationHeader(api.Token)\n")
		fmt.Fprint(&sb, "\treq.Header.Add(\"Authorization\", authorization)\n")
	}
	fmt.Fprint(&sb, "\treq.Header.Add(\"User-Agent\", api.UserAgent)\n")
	fmt.Fprint(&sb, "\tvar httpClient HTTPClient = http.DefaultClient\n")
	fmt.Fprint(&sb, "\tif api.HTTPClient != nil {\n")
	fmt.Fprint(&sb, "\t\thttpClient = api.HTTPClient\n")
	fmt.Fprint(&sb, "\t}\n")
	fmt.Fprint(&sb, "\treturn api.newResponse(httpClient.Do(req))\n")
	fmt.Fprint(&sb, "}\n\n")

	return sb.String()
}
