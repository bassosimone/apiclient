package spec

import (
	"fmt"
	"strings"
)

// GenAPIAdapterType generates the type adapter for login.
func (d *Descriptor) GenAPIAdapterType() string {
	if !d.RequiresLogin {
		return "" // nothing to generate
	}
	var sb strings.Builder

	// generate the login adapter
	fmt.Fprintf(&sb, "type %sLoginAdapter struct {\n", d.apiStructName())
	fmt.Fprint(&sb, "\tmu sync.Mutex\n")
	fmt.Fprintf(&sb, "\t\treq %s\n", d.requestTypeName())
	fmt.Fprintf(&sb, "\t\tresp %s\n", d.responseTypeName())
	fmt.Fprint(&sb, "}\n\n")

	// generate the adapter call function
	fmt.Fprintf(
		&sb, "func (la *%sLoginAdapter) call(ctx context.Context, clnt *Client, token string) error {\n",
		d.apiStructName())
	fmt.Fprint(&sb, "\tdefer la.mu.Unlock()\n")
	fmt.Fprint(&sb, "\tla.mu.Lock()\n")
	fmt.Fprint(&sb, "\tif la.resp != nil {\n")
	fmt.Fprint(&sb, "\t\treturn errors.New(\"apiclient: call already succeeded\")\n")
	fmt.Fprint(&sb, "\t}\n")
	fmt.Fprintf(&sb, "\tapi := new%sAPI(clnt, token)\n", d.Name)
	fmt.Fprint(&sb, "\tresp, err := api.call(ctx, la.req)\n")
	fmt.Fprint(&sb, "\tif err != nil {\n")
	fmt.Fprint(&sb, "\t\treturn err // we may need to try with another token\n")
	fmt.Fprint(&sb, "\t}\n")
	fmt.Fprint(&sb, "\tla.resp = resp\n")
	fmt.Fprint(&sb, "\treturn nil\n")
	fmt.Fprint(&sb, "}\n\n")

	return sb.String()
}
