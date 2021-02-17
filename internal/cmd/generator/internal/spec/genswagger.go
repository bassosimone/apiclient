package spec

import (
	"reflect"
	"strings"
	"sync"

	"github.com/bassosimone/apiclient/internal/cmd/internal/openapi"
)

const (
	tagForJSON = "json"
	tagForPath = "path"
)

func (d *Descriptor) genSwaggerURLPath() string {
	up := d.URLPath
	if up.InSwagger != "" {
		return up.InSwagger
	}
	if up.IsTemplate {
		panic("we should always use InSwapper and IsTemplate together")
	}
	return up.Value
}

func (d *Descriptor) genSwaggerSchema(cur reflect.Type) *openapi.Schema {
	switch cur.Kind() {
	case reflect.String:
		return &openapi.Schema{Type: "string"}
	case reflect.Bool:
		return &openapi.Schema{Type: "boolean"}
	case reflect.Int64:
		return &openapi.Schema{Type: "integer"}
	case reflect.Slice:
		return &openapi.Schema{Type: "array", Items: d.genSwaggerSchema(cur.Elem())}
	case reflect.Map:
		return &openapi.Schema{Type: "object"}
	case reflect.Ptr:
		return d.genSwaggerSchema(cur.Elem())
	case reflect.Struct:
		if cur.String() == "time.Time" {
			// Implementation note: we don't want to dive into time.Time but
			// rather we want to pretend it's a string. The JSON parser for
			// time.Time can indeed reconstruct a time.Time from a string, and
			// it's much easier for us to let it do the parsing.
			return &openapi.Schema{Type: "string"}
		}
		sinfo := &openapi.Schema{Type: "object"}
		var once sync.Once
		initmap := func() {
			sinfo.Properties = make(map[string]*openapi.Schema)
		}
		for idx := 0; idx < cur.NumField(); idx++ {
			field := cur.Field(idx)
			if field.Tag.Get(tagForPath) != "" {
				continue
			}
			if field.Tag.Get(tagForQuery) != "" {
				continue
			}
			v := field.Name
			if j := field.Tag.Get(tagForJSON); j != "" {
				j = strings.Replace(j, ",omitempty", "", 1)
				if j == "-" {
					continue
				}
				v = j
			}
			once.Do(initmap)
			sinfo.Properties[v] = d.genSwaggerSchema(field.Type)
		}
		return sinfo
	case reflect.Interface:
		return &openapi.Schema{Type: "object"}
	default:
		panic("unsupported type")
	}
}

func (d *Descriptor) swaggerParamForType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "boolean"
	case reflect.Int64:
		return "integer"
	default:
		panic("unsupported type")
	}
}

func (d *Descriptor) genSwaggerParams(cur reflect.Type) []*openapi.Parameter {
	if cur.Kind() != reflect.Ptr {
		panic("not a pointer")
	}
	cur = cur.Elem()
	if cur.Kind() != reflect.Struct {
		panic("not a pointer to struct")
	}
	var out []*openapi.Parameter
	for idx := 0; idx < cur.NumField(); idx++ {
		f := cur.Field(idx)
		if q := f.Tag.Get(tagForQuery); q != "" {
			out = append(out, &openapi.Parameter{
				Name:     q,
				In:       "query",
				Required: f.Tag.Get(tagForRequired) == "true",
				Type:     d.swaggerParamForType(f.Type),
			})
			continue
		}
		if p := f.Tag.Get(tagForPath); p != "" {
			out = append(out, &openapi.Parameter{
				Name:     p,
				In:       "path",
				Required: true,
				Type:     d.swaggerParamForType(f.Type),
			})
			continue
		}
	}
	return out
}

// GenSwaggerPath generates the OpenAPI 2.0 path for the Descriptor.
func (d *Descriptor) GenSwaggerPath() (string, *openapi.Path) {
	pathStr, pathInfo := d.genSwaggerURLPath(), &openapi.Path{}
	rtinfo := &openapi.RoundTrip{Produces: []string{"application/json"}}
	switch d.Method {
	case "GET":
		pathInfo.Get = rtinfo
	case "POST":
		rtinfo.Consumes = append(rtinfo.Consumes, "application/json")
		pathInfo.Post = rtinfo
	}
	rtinfo.Parameters = d.genSwaggerParams(reflect.TypeOf(d.Request))
	if d.Method != "GET" {
		rtinfo.Parameters = append(rtinfo.Parameters, &openapi.Parameter{
			Name:     "body",
			In:       "body",
			Required: true,
			Schema:   d.genSwaggerSchema(reflect.TypeOf(d.Request)),
		})
	}
	rtinfo.Responses = &openapi.Responses{Successful: openapi.Body{
		Description: "all good",
		Schema:      d.genSwaggerSchema(reflect.TypeOf(d.Response)),
	}}
	return pathStr, pathInfo
}
