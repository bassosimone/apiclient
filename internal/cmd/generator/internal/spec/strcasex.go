package spec

import (
	"github.com/iancoleman/strcase"
)

func init() {
	// make sure strcase behaves the way we want
	strcase.ConfigureAcronym("URLs", "urls")
}

func toLowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
