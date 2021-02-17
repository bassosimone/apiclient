package apimodel

import (
	"fmt"
	"strings"
)

// GenNewAPI generates the code that creates a new API instance.
func (d *Descriptor) GenNewAPI() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "func new%sAPI(c *Client) *%s {\n", d.Name, d.apiStructName())
	fmt.Fprintf(&sb, "\tapi := &%s{\n", d.apiStructName())
	if d.RequiresLogin {
		fmt.Fprintf(&sb, "\t\tAuthorizer: c,\n")
	}
	fmt.Fprintf(&sb, "\t\tBaseURL: c.baseURL(),\n")
	fmt.Fprintf(&sb, "\t\tHTTPClient: c.HTTPClient,\n")
	fmt.Fprintf(&sb, "\t\tUserAgent: c.UserAgent,\n")
	fmt.Fprintf(&sb, "\t}\n")
	fmt.Fprintf(&sb, "\treturn api\n")
	fmt.Fprint(&sb, "}\n\n")
	return sb.String()
}
