package structs

import "reflect"

func Traverse(val reflect.Value, fn func(
	val reflect.Value, field reflect.StructField,
)) {
	typ := val.Type()
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldTyp := field.Type
		fieldVal := val.FieldByIndex(field.Index)
		if fieldTyp.Kind() == reflect.Ptr {
			fieldTyp = fieldTyp.Elem()
			if fieldVal.IsNil() {
				fieldVal.Set(reflect.New(fieldTyp))
			}
			fieldVal = fieldVal.Elem()
		}

		if field.Anonymous && fieldTyp.Kind() == reflect.Struct {
			Traverse(fieldVal, fn)
		} else if field.Name[0] >= 'A' && field.Name[0] <= 'Z' {
			fn(fieldVal, field)
		}
	}
}

func TraverseType(typ reflect.Type, fn func(field reflect.StructField)) {
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
				TraverseType(fieldTyp, fn)
				continue
			}
		}
		if field.Name[0] >= 'A' && field.Name[0] <= 'Z' {
			fn(field)
		}
	}
}
