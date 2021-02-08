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
	"github.com/bassosimone/apiclient/internal/osx"
	"github.com/bassosimone/apiclient/internal/reflectx"
)

type schemaInfo struct {
	Properties map[string]*schemaInfo `json:"properties,omitempty"`
	Items      *schemaInfo            `json:"items,omitempty"`
	Type       string                 `json:"type"`
}

type parameterInfo struct {
	In       string      `json:"in"`
	Name     string      `json:"name"`
	Required bool        `json:"required,omitempty"`
	Schema   *schemaInfo `json:"schema,omitempty"`
	Type     string      `json:"type,omitempty"`
}

type bodyInfo struct {
	Description string      `json:"description,omitempty"`
	Schema      *schemaInfo `json:"schema"`
}

type responsesInfo struct {
	Successful bodyInfo `json:"200"`
}

type roundTripInfo struct {
	Consumes   []string        `json:"consumes,omitempty"`
	Produces   []string        `json:"produces,omitempty"`
	Parameters []parameterInfo `json:"parameters,omitempty"`
	Responses  *responsesInfo  `json:"responses,omitempty"`
}

type pathInfo struct {
	Get  *roundTripInfo `json:"get,omitempty"`
	Post *roundTripInfo `json:"post,omitempty"`
}

type apiInfo struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

type swagger struct {
	Swagger  string               `json:"swagger"`
	Info     apiInfo              `json:"info"`
	Host     string               `json:"host"`
	BasePath string               `json:"basePath"`
	Schemes  []string             `json:"schemes"`
	Paths    map[string]*pathInfo `json:"paths"`
}

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

func genparams(req *reflectx.TypeValueInfo) []parameterInfo {
	fields, err := req.AllFields()
	if err != nil {
		return nil
	}
	var out []parameterInfo
	for _, field := range fields {
		if q := field.Self.Tag.Get("query"); q != "" {
			out = append(out, parameterInfo{
				Name:     q,
				In:       "query",
				Required: field.Self.Tag.Get("required") == "true",
				Type:     genparamtype(field.Self.Type),
			})
			continue
		}
		if p := field.Self.Tag.Get("path"); p != "" {
			out = append(out, parameterInfo{
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

func genschemainfo(cur reflect.Type) *schemaInfo {
	switch cur.Kind() {
	case reflect.String:
		return &schemaInfo{Type: "string"}
	case reflect.Bool:
		return &schemaInfo{Type: "boolean"}
	case reflect.Int64:
		return &schemaInfo{Type: "integer"}
	case reflect.Slice:
		return &schemaInfo{Type: "array", Items: genschemainfo(cur.Elem())}
	case reflect.Map:
		return &schemaInfo{Type: "object"}
	case reflect.Ptr:
		return genschemainfo(cur.Elem())
	case reflect.Struct:
		sinfo := &schemaInfo{Type: "object"}
		var once sync.Once
		initmap := func() {
			sinfo.Properties = make(map[string]*schemaInfo)
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
		return &schemaInfo{Type: "object"}
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
	swagger := swagger{
		Swagger: "2.0",
		Info: apiInfo{
			Title:   "OONI API specification",
			Version: genversion(),
		},
		Host:     "api.ooni.io",
		BasePath: "/",
		Schemes:  []string{"https"},
		Paths:    make(map[string]*pathInfo),
	}
	for _, descr := range apimodel.Descriptors {
		pinfo := &pathInfo{}
		swagger.Paths[genpath(&descr.URLPath)] = pinfo
		rtinfo := &roundTripInfo{
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
			rtinfo.Parameters = append(rtinfo.Parameters, parameterInfo{
				Name:   "body",
				In:     "body",
				Schema: genschemainfo(req.TypeInfo()),
			})
		}
		rtinfo.Responses = &responsesInfo{Successful: bodyInfo{
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
