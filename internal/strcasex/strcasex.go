package strcasex

import (
	"github.com/iancoleman/strcase"
)

func init() {
	strcase.ConfigureAcronym("URLs", "urls")
}

// ToLowerCamel converts a string to camelCase with the
// first letter being lowercase.
func ToLowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
