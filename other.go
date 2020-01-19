package goson

import (
	"errors"
	"fmt"
	"reflect"
)

func getAddr(value interface{}) (*string, error) {
	var val reflect.Value

	switch value.(type) {
	case reflect.Value:
		val = value.(reflect.Value)
	default:
		val = reflect.ValueOf(value)
	}

	kind := val.Kind()

	for kind == reflect.Struct || kind == reflect.Interface || kind == reflect.Ptr {
		if kind == reflect.Ptr {
			elem := val.Elem()
			if !elem.IsValid() {
				break
			}
		}

		switch kind {
		case reflect.Struct:
			if val.NumField() < 1 {
				return nil, errors.New("struct not value")
			}
			val = val.Field(0)
			kind = val.Kind()
		case reflect.Ptr:
			val = val.Elem()
			kind = val.Kind()
		default:
			val = val.Elem()
			kind = val.Kind()
		}
	}

	var addr string
	if val.CanAddr() {
		addr = fmt.Sprintf("%x", val.Addr())
	}

	return &addr, nil
}

func getFieldNameFromStructField(structField *reflect.StructField, _ *string) *string {
	name := structField.Name
	return &name
}

func getFieldNameFromMetaTag(structField *reflect.StructField, tag *string) *string {
	name := structField.Tag.Get(*tag)
	return &name
}
