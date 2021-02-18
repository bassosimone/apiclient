package spec

import (
	"fmt"
	"strings"
)

// GenAPIType generates the type definition for the API.
func (d *Descriptor) GenAPIType() string {
	var sb strings.Builder
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
	return sb.String()
}
