// Package reflectx extends the reflect package.
package reflectx

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/bassosimone/apiclient/internal/fatalx"
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

// Must fails if we cannot construct a TypeValueInfo
func Must(si *TypeValueInfo, err error) *TypeValueInfo {
	fatalx.OnError(err, "NewTypeValueInfo failed")
	return si
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

// AsInitialization generates an initialization for the named variable
func (si TypeValueInfo) AsInitialization(name string) string {
	switch si.typeInfo.Kind() {
	case reflect.Struct:
		return fmt.Sprintf("var %s %s", name, si.typeInfo.Name())
	case reflect.Map:
		return fmt.Sprintf("%s := %s{}", name, si.typeInfo.Name())
	default:
		panic("AsInitialization: unsupported type")
	}
}

// AsReturnType generates a declaration for si as a return type.
func (si TypeValueInfo) AsReturnType() string {
	switch si.typeInfo.Kind() {
	case reflect.Struct:
		return fmt.Sprintf("*%s", si.typeInfo.Name())
	case reflect.Map:
		return fmt.Sprintf("%s", si.typeInfo.Name())
	default:
		panic("AsReturnType: unsupported type")
	}
}

// AsReturnValue generates a declaration for si as a return value.
func (si TypeValueInfo) AsReturnValue(name string) string {
	switch si.typeInfo.Kind() {
	case reflect.Struct:
		return fmt.Sprintf("&%s", name)
	case reflect.Map:
		return fmt.Sprintf("%s", name)
	default:
		panic("AsReturnValue: unsupported type")
	}
}
