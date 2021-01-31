// This script generates swagger.json
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"

	"github.com/bassosimone/apiclient/internal/apimodel"
	"github.com/bassosimone/apiclient/internal/fatalx"
	"github.com/bassosimone/apiclient/internal/reflectx"
)

type apiInfo struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

type serverInfo struct {
	URL string `json:"url"`
}

type schemaInfo struct {
	Type       string                 `json:"type"`
	Properties map[string]*schemaInfo `json:"properties,omitempty"`
	Items      *schemaInfo            `json:"items,omitempty"`
}

type parameterInfo struct {
	Name     string     `json:"name"`
	In       string     `json:"in"`
	Required bool       `json:"required"`
	Schema   schemaInfo `json:"schema"`
}

type jsonInfo struct {
	Schema *schemaInfo `json:"schema"`
}

type contentInfo struct {
	JSON *jsonInfo `json:"application/json,omitempty"`
}

type bodyInfo struct {
	Description string       `json:"description,omitempty"`
	Content     *contentInfo `json:"content,omitempty"`
}

type responsesInfo struct {
	Successful bodyInfo `json:"200"`
}

type roundTripInfo struct {
	Parameters  []parameterInfo `json:"parameters,omitempty"`
	RequestBody *bodyInfo       `json:"requestBody,omitempty"`
	Responses   *responsesInfo  `json:"responses,omitempty"`
}

type pathInfo struct {
	Get  *roundTripInfo `json:"get,omitempty"`
	Post *roundTripInfo `json:"post,omitempty"`
}

type swagger struct {
	OpenAPI string               `json:"openapi"`
	Info    apiInfo              `json:"info"`
	Servers []serverInfo         `json:"servers"`
	Paths   map[string]*pathInfo `json:"paths"`
}

func gettype(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Bool:
		return "boolean"
	case reflect.Int64:
		return "integer"
	case reflect.String:
		return "string"
	case reflect.Slice:
		return "array"
	default:
		return "object"
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
				Schema: schemaInfo{
					Type: gettype(field.Self.Type),
				},
			})
			continue
		}
		if p := field.Self.Tag.Get("path"); p != "" {
			out = append(out, parameterInfo{
				Name:     p,
				In:       "path",
				Required: true,
				Schema:   schemaInfo{Type: "string"},
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
				if j == "-" {
					continue
				}
				v = strings.Replace(j, ",omitempty", "", 1)
			}
			once.Do(initmap)
			sinfo.Properties[v] = genschemainfo(field.Type)
		}
		return sinfo
	case reflect.Interface:
		return &schemaInfo{Type: "object"}
	default:
		log.Fatalf("unsupported type: %+v", cur)
		return nil
	}
}

func genrequestbody(req *reflectx.TypeValueInfo) *bodyInfo {
	sinfo := genschemainfo(req.TypeInfo())
	return &bodyInfo{Content: &contentInfo{JSON: &jsonInfo{Schema: sinfo}}}
	/*
		fields, err := req.AllFields()
		if err != nil {
			return nil
		}
		sinfo := &schemaInfo{Type: "object", Properties: make(map[string]propertyInfo)}
		out := &bodyInfo{Content: &contentInfo{JSON: &jsonInfo{Schema: sinfo}}}
		for _, field := range fields {
			if q := field.Self.Tag.Get("query"); q != "" {
				continue
			}
			if p := field.Self.Tag.Get("path"); p != "" {
				continue
			}
			v := field.Self.Name
			if j := field.Self.Tag.Get("json"); j != "" {
				if j == "-" {
					continue
				}
				v = strings.Replace(j, ",omitempty", "", 1)
			}
			pt := gettype(field.Self.Type)
			sinfo.Properties[v] = propertyInfo{
				Type:  pt,
				Items: genitems(field),
			}
		}
		return out
	*/
}

func genresponsebody(req *reflectx.TypeValueInfo) *contentInfo {
	sinfo := genschemainfo(req.TypeInfo())
	return &contentInfo{JSON: &jsonInfo{Schema: sinfo}}
	/*
		fields, err := req.AllFields()
		if err != nil {
			// TODO(bassosimone): we don't support map[string]interface{} responses;
			// the catch is that we should always generate a schema dynamically from
			// the object that we're currently serializing.
			return nil
		}
		sinfo := &schemaInfo{Type: "object", Properties: make(map[string]propertyInfo)}
		out := &contentInfo{JSON: &jsonInfo{Schema: sinfo}}
		for _, field := range fields {
			v := field.Self.Name
			if j := field.Self.Tag.Get("json"); j != "" {
				if j == "-" {
					continue
				}
				v = strings.Replace(j, ",omitempty", "", 1)
			}
			pt := gettype(field.Self.Type)
			sinfo.Properties[v] = propertyInfo{
				Type:  pt,
				Items: genitems(field),
			}
		}
		return out
	*/
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

func main() {
	swagger := swagger{
		OpenAPI: "3.0.0",
		Info: apiInfo{
			Title:   "OONI API specification",
			Version: "2021.01.30",
		},
		Servers: []serverInfo{{
			URL: "https://api.ooni.io/",
		}},
		Paths: make(map[string]*pathInfo),
	}
	for _, descr := range apimodel.Descriptors {
		pinfo := &pathInfo{}
		swagger.Paths[genpath(&descr.URLPath)] = pinfo
		rtinfo := &roundTripInfo{}
		switch descr.Method {
		case "GET":
			pinfo.Get = rtinfo
		case "POST":
			pinfo.Post = rtinfo
		}
		req := reflectx.Must(reflectx.NewTypeValueInfo(descr.Request))
		resp := reflectx.Must(reflectx.NewTypeValueInfo(descr.Response))
		rtinfo.Parameters = genparams(req)
		if descr.Method != "GET" {
			rtinfo.RequestBody = genrequestbody(req)
		}
		rtinfo.Responses = &responsesInfo{Successful: bodyInfo{
			Description: "all good",
			Content:     genresponsebody(resp),
		}}
	}
	data, err := json.MarshalIndent(swagger, "", "    ")
	fatalx.OnError(err, "json.Marshal failed")
	fmt.Printf("%s\n", string(data))
}
