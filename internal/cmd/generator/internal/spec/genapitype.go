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
	fmt.Fprint(&sb, "\tbaseURL     string\n")
	fmt.Fprint(&sb, "\thttpClient  HTTPClient\n")
	fmt.Fprint(&sb, "\tjsonCodec   JSONCodec\n")
	fmt.Fprint(&sb, "\trequestMaker RequestMaker\n")
	if d.RequiresLogin {
		fmt.Fprintf(&sb, "\ttoken string\n")
	}
	if d.URLPath.IsTemplate {
		fmt.Fprint(&sb, "\ttemplateExecutor TemplateExecutor\n")
	}
	fmt.Fprint(&sb, "\tuserAgent   string\n")
	fmt.Fprint(&sb, "}\n\n")

	// generate the newAPI factory
	if d.RequiresLogin {
		fmt.Fprintf(&sb, "func new%sAPI(c *Client, token string) *%s {\n", d.Name, d.apiStructName())
	} else {
		fmt.Fprintf(&sb, "func new%sAPI(c *Client) *%s {\n", d.Name, d.apiStructName())
	}
	fmt.Fprintf(&sb, "\tvar clnt HTTPClient = c.httpClient\n")
	if d.Cache {
		fmt.Fprintf(&sb, "\tclnt = &cacheClient{\n")
		fmt.Fprintf(&sb, "\t\tHTTPClient: clnt,\n")
		fmt.Fprintf(&sb, "\t\tKVStore: c.kvStore,\n")
		fmt.Fprintf(&sb, "\t}\n")
	}
	fmt.Fprintf(&sb, "\tapi := &%s{\n", d.apiStructName())
	fmt.Fprintf(&sb, "\t\tbaseURL: c.baseURL,\n")
	fmt.Fprintf(&sb, "\t\thttpClient: clnt,\n")
	fmt.Fprintf(&sb, "\t\tjsonCodec: c.jsonCodec,\n")
	fmt.Fprintf(&sb, "\t\trequestMaker: c.requestMaker,\n")
	if d.RequiresLogin {
		fmt.Fprintf(&sb, "\t\ttoken: token,\n")
	}
	fmt.Fprintf(&sb, "\t\tuserAgent: c.userAgent,\n")
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
		fmt.Fprint(&sb, "\tif api.token == \"\" {\n")
		fmt.Fprint(&sb, "\t\treturn nil, errMissingToken\n")
		fmt.Fprint(&sb, "\t}\n")
		fmt.Fprintf(&sb, "\tauthorization := newAuthorizationHeader(api.token)\n")
		fmt.Fprint(&sb, "\treq.Header.Add(\"Authorization\", authorization)\n")
	}
	fmt.Fprint(&sb, "\treq.Header.Add(\"User-Agent\", api.userAgent)\n")
	fmt.Fprint(&sb, "\treturn api.newResponse(api.httpClient.Do(req))\n")
	fmt.Fprint(&sb, "}\n\n")

	return sb.String()
}
