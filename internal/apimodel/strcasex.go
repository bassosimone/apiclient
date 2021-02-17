package apimodel

import (
	"github.com/iancoleman/strcase"
)

func init() {
	strcase.ConfigureAcronym("URLs", "urls")
}

func toLowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
