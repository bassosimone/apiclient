package spec

import (
	"fmt"
	"strings"
)

// GenNewAPI generates the code that creates a new API instance.
func (d *Descriptor) GenNewAPI() string {
	var sb strings.Builder
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
	return sb.String()
}
