// Package reflectx extends the reflect package.
package reflectx

import (
	"errors"
	"reflect"
)

// This package returns the following errors.
var (
	ErrNilPointer = errors.New("reflectx: nil pointer")
	ErrNotStruct  = errors.New("reflectx: not a struct")
)

// StructInfo contains info about a struct
type StructInfo struct {
	typeInfo  reflect.Type
	valueInfo *reflect.Value
}

// NewStructInfo creates a new StructInfo
func NewStructInfo(in interface{}) (*StructInfo, error) {
	valueInfo := reflect.ValueOf(in)
	if valueInfo.Kind() == reflect.Ptr {
		valueInfo = valueInfo.Elem()
		if valueInfo.IsZero() {
			return nil, ErrNilPointer
		}
	}
	typeInfo := valueInfo.Type()
	if typeInfo.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}
	return &StructInfo{typeInfo: typeInfo, valueInfo: &valueInfo}, nil
}

// TypeName returns the name of the struct type.
func (si StructInfo) TypeName() string {
	return si.typeInfo.Name()
}

// FieldInfo contains information about a field.
type FieldInfo struct {
	Self  *reflect.StructField
	Value *reflect.Value
}

// AllFieldsWithTag returns all fields with a given tag name.
func (si StructInfo) AllFieldsWithTag(tagName string) []*FieldInfo {
	var out []*FieldInfo
	for idx := 0; idx < si.typeInfo.NumField(); idx++ {
		fieldType := si.typeInfo.Field(idx)
		if tag := fieldType.Tag.Get(tagName); tag == "" {
			continue
		}
		fieldValue := si.valueInfo.Field(idx)
		out = append(out, &FieldInfo{Self: &fieldType, Value: &fieldValue})
	}
	return out
}
