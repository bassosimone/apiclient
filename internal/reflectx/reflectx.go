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

// TypeValueInfo contains info about a type
type TypeValueInfo struct {
	typeInfo  reflect.Type
	valueInfo *reflect.Value
}

// NewTypeValueInfo creates a new TypeValueInfo
func NewTypeValueInfo(in interface{}) (*TypeValueInfo, error) {
	valueInfo := reflect.ValueOf(in)
	if valueInfo.Kind() == reflect.Ptr {
		valueInfo = valueInfo.Elem()
		if valueInfo.IsZero() {
			return nil, ErrNilPointer
		}
	}
	typeInfo := valueInfo.Type()
	return &TypeValueInfo{typeInfo: typeInfo, valueInfo: &valueInfo}, nil
}

// TypeName returns the name of the struct type.
func (si TypeValueInfo) TypeName() string {
	return si.typeInfo.Name()
}

// FieldInfo contains information about a field.
type FieldInfo struct {
	Self  *reflect.StructField
	Value *reflect.Value
}

// AllFieldsWithTag returns all fields with a given tag name.
func (si TypeValueInfo) AllFieldsWithTag(tagName string) ([]*FieldInfo, error) {
	if si.typeInfo.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}
	var out []*FieldInfo
	for idx := 0; idx < si.typeInfo.NumField(); idx++ {
		fieldType := si.typeInfo.Field(idx)
		if tag := fieldType.Tag.Get(tagName); tag == "" {
			continue
		}
		fieldValue := si.valueInfo.Field(idx)
		out = append(out, &FieldInfo{Self: &fieldType, Value: &fieldValue})
	}
	return out, nil
}
