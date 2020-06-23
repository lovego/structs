package structs

import "reflect"

func TraverseExportedFields(typ reflect.Type, fn func(field reflect.StructField)) {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			TraverseExportedFields(field.Type, fn)
		} else if field.Name[0] >= 'A' && field.Name[0] <= 'Z' {
			fn(field)
		}
	}
}
