package structs

import "reflect"

func TraverseExportedFields(typ reflect.Type, fn func(field reflect.StructField)) {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Anonymous {
			fieldTyp := field.Type
			if fieldTyp.Kind() == reflect.Ptr {
				fieldTyp = fieldTyp.Elem()
			}
			if fieldTyp.Kind() == reflect.Struct {
				TraverseExportedFields(fieldTyp, fn)
				continue
			}
		}
		if field.Name[0] >= 'A' && field.Name[0] <= 'Z' {
			fn(field)
		}
	}
}
