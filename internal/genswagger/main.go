// This script generates swagger.json
package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/bassosimone/apiclient/internal/apimodel"
	"github.com/bassosimone/apiclient/internal/fatalx"
	"github.com/bassosimone/apiclient/internal/fmtx"
	"github.com/bassosimone/apiclient/internal/openapi"
	"github.com/bassosimone/apiclient/internal/osx"
	"github.com/bassosimone/apiclient/internal/reflectx"
)

func genparamtype(t reflect.Type) string {
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

func genparams(req *reflectx.TypeValueInfo) []openapi.Parameter {
	fields, err := req.AllFields()
	if err != nil {
		return nil
	}
	var out []openapi.Parameter
	for _, field := range fields {
		if q := field.Self.Tag.Get("query"); q != "" {
			out = append(out, openapi.Parameter{
				Name:     q,
				In:       "query",
				Required: field.Self.Tag.Get("required") == "true",
				Type:     genparamtype(field.Self.Type),
			})
			continue
		}
		if p := field.Self.Tag.Get("path"); p != "" {
			out = append(out, openapi.Parameter{
				Name:     p,
				In:       "path",
				Required: true,
				Type:     genparamtype(field.Self.Type),
			})
			continue
		}
	}
	return out
}

func genschemainfo(cur reflect.Type) *openapi.Schema {
	switch cur.Kind() {
	case reflect.String:
		return &openapi.Schema{Type: "string"}
	case reflect.Bool:
		return &openapi.Schema{Type: "boolean"}
	case reflect.Int64:
		return &openapi.Schema{Type: "integer"}
	case reflect.Slice:
		return &openapi.Schema{Type: "array", Items: genschemainfo(cur.Elem())}
	case reflect.Map:
		return &openapi.Schema{Type: "object"}
	case reflect.Ptr:
		return genschemainfo(cur.Elem())
	case reflect.Struct:
		sinfo := &openapi.Schema{Type: "object"}
		var once sync.Once
		initmap := func() {
			sinfo.Properties = make(map[string]*openapi.Schema)
		}
		for idx := 0; idx < cur.NumField(); idx++ {
			field := cur.Field(idx)
			if field.Tag.Get("path") != "" {
				continue
			}
			if field.Tag.Get("query") != "" {
				continue
			}
			v := field.Name
			if j := field.Tag.Get("json"); j != "" {
				j = strings.Replace(j, ",omitempty", "", 1)
				if j == "-" {
					continue
				}
				v = j
			}
			once.Do(initmap)
			sinfo.Properties[v] = genschemainfo(field.Type)
		}
		return sinfo
	case reflect.Interface:
		return &openapi.Schema{Type: "object"}
	default:
		panic("unsupported type")
	}
}

func genpath(up *apimodel.URLPath) string {
	if up.InSwagger != "" {
		return up.InSwagger
	}
	if up.IsTemplate {
		panic("we should always use InSwapper and IsTemplate together")
	}
	return up.Value
}

func genversion() string {
	return time.Now().UTC().Format("0.20060102.1150405")
}

func main() {
	swagger := openapi.Swagger{
		Swagger: "2.0",
		Info: openapi.API{
			Title:   "OONI API specification",
			Version: genversion(),
		},
		Host:     "api.ooni.io",
		BasePath: "/",
		Schemes:  []string{"https"},
		Paths:    make(map[string]*openapi.Path),
	}
	for _, descr := range apimodel.Descriptors {
		pinfo := &openapi.Path{}
		swagger.Paths[genpath(&descr.URLPath)] = pinfo
		rtinfo := &openapi.RoundTrip{
			Produces: []string{"application/json"},
		}
		switch descr.Method {
		case "GET":
			pinfo.Get = rtinfo
		case "POST":
			rtinfo.Consumes = append(rtinfo.Consumes, "application/json")
			pinfo.Post = rtinfo
		}
		req := reflectx.Must(reflectx.NewTypeValueInfo(descr.Request))
		resp := reflectx.Must(reflectx.NewTypeValueInfo(descr.Response))
		rtinfo.Parameters = genparams(req)
		if descr.Method != "GET" {
			rtinfo.Parameters = append(rtinfo.Parameters, openapi.Parameter{
				Name:   "body",
				In:     "body",
				Schema: genschemainfo(req.TypeInfo()),
			})
		}
		rtinfo.Responses = &openapi.Responses{Successful: openapi.Body{
			Description: "all good",
			Schema:      genschemainfo(resp.TypeInfo()),
		}}
	}
	data, err := json.MarshalIndent(swagger, "", "    ")
	fatalx.OnError(err, "json.Marshal failed")
	filep := osx.MustCreate("swagger.json")
	defer filep.Close()
	fmtx.Fprintf(filep, "%s\n", string(data))
}
