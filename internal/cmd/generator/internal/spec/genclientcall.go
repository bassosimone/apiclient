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
		fmt.Fprintf(&sb, "\tla := &%sLoginAdapter{req: req}\n", d.apiStructName())
		fmt.Fprint(&sb, "\tif err := c.doWithLoginAdapter(ctx, la); err != nil {\n")
		fmt.Fprint(&sb, "\t\treturn nil, err\n")
		fmt.Fprint(&sb, "\t}\n")
		fmt.Fprint(&sb, "\treturn la.resp, nil\n")
	} else {
		fmt.Fprintf(&sb, "\treturn new%sAPI(c).call(ctx, req)\n", d.Name)
	}
	fmt.Fprintf(&sb, "}\n\n")
	return sb.String()
}
