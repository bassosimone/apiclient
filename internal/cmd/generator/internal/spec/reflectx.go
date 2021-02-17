package spec

import (
	"fmt"
	"reflect"
)

// typeName returns v's package-qualified type name.
func (d *Descriptor) typeName(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// requestTypeName calls d.typeName(d.Request).
func (d *Descriptor) requestTypeName() string {
	return d.typeName(d.Request)
}

// responseTypeName calls d.typeName(d.Response).
func (d *Descriptor) responseTypeName() string {
	return d.typeName(d.Response)
}

// apiStructName returns the correct struct type name
// for the API we're currently processing (i.e., d).
func (d *Descriptor) apiStructName() string {
	return fmt.Sprintf("%sAPI", toLowerCamel(d.Name))
}

// getStructFields returns all the struct fields of in. This function
// assumes that in is a pointer to struct, and will otherwise panic.
func (d *Descriptor) getStructFields(in interface{}) []*reflect.StructField {
	t := reflect.TypeOf(in)
	if t.Kind() != reflect.Ptr {
		panic("not a pointer")
	}
	t = t.Elem()
	if t.Kind() != reflect.Struct {
		panic("not a struct")
	}
	var out []*reflect.StructField
	for idx := 0; idx < t.NumField(); idx++ {
		f := t.Field(idx)
		out = append(out, &f)
	}
	return out
}

// getStructFieldsWithTag returns all the struct fields of
// in that have the specified tag.
func (d *Descriptor) getStructFieldsWithTag(in interface{}, tag string) []*reflect.StructField {
	var out []*reflect.StructField
	for _, f := range d.getStructFields(in) {
		if f.Tag.Get(tag) != "" {
			out = append(out, f)
		}
	}
	return out
}

// requestOrResponseTypeKind returns the type kind of in, which should
// be a request or a response. This function assumes that in is either a
// pointer to struct or a map and will panic otherwise.
func (d *Descriptor) requestOrResponseTypeKind(in interface{}) reflect.Kind {
	t := reflect.TypeOf(in)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		if t.Kind() != reflect.Struct {
			panic("not a struct")
		}
		return reflect.Struct
	}
	if t.Kind() != reflect.Map {
		panic("not a map")
	}
	return reflect.Map
}

// requestTypeKind calls d.requestOrResponseTypeKind(d.Request).
func (d *Descriptor) requestTypeKind() reflect.Kind {
	return d.requestOrResponseTypeKind(d.Request)
}

// responseTypeKind calls d.requestOrResponseTypeKind(d.Response).
func (d *Descriptor) responseTypeKind() reflect.Kind {
	return d.requestOrResponseTypeKind(d.Response)
}

// typeNameAsStruct assumes that in is a pointer to struct and
// returns the type of the corresponding struct. The returned
// type is package qualified.
func (d *Descriptor) typeNameAsStruct(in interface{}) string {
	t := reflect.TypeOf(in)
	if t.Kind() != reflect.Ptr {
		panic("not a pointer")
	}
	t = t.Elem()
	if t.Kind() != reflect.Struct {
		panic("not a struct")
	}
	return t.String()
}

// requestTypeNameAsStruct calls d.typeNameAsStruct(d.Request)
func (d *Descriptor) requestTypeNameAsStruct() string {
	return d.typeNameAsStruct(d.Request)
}

// responseTypeNameAsStruct calls d.typeNameAsStruct(d.Response)
func (d *Descriptor) responseTypeNameAsStruct() string {
	return d.typeNameAsStruct(d.Response)
}
