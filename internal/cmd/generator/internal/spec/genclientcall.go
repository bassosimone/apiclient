package spec

import (
	"fmt"
	"strings"
)

// GenClientCall generates the code that calls the API from the client.
func (d *Descriptor) GenClientCall() string {
	var sb strings.Builder
	if d.Private {
		return "" // we don't generate a client call for private APIs
	}
	fmt.Fprintf(&sb, "// %s calls the %s API.\n", d.Name, d.Name)
	fmt.Fprintf(
		&sb, "func (c *Client) %s(ctx context.Context, req %s) (%s, error) {\n",
		d.Name, d.requestTypeName(), d.responseTypeName())
	if d.RequiresLogin {
		fmt.Fprint(&sb, "\ttoken, err := c.maybeRefreshToken(ctx)\n")
		fmt.Fprint(&sb, "\tif err != nil {\n")
		fmt.Fprint(&sb, "\t\treturn nil, err\n")
		fmt.Fprint(&sb, "\t}\n")
		fmt.Fprintf(&sb, "\tapi := new%sAPI(c, token)\n", d.Name)
		fmt.Fprint(&sb, "\treturn api.call(ctx, req)\n")
	} else {
		fmt.Fprintf(&sb, "\treturn new%sAPI(c).call(ctx, req)\n", d.Name)
	}
	fmt.Fprintf(&sb, "}\n\n")
	return sb.String()
}
