package structs

import "reflect"

// Traverse traverses a reflect.Value
// convertNilPtr:  convert nil pointer to non-nil or not.
func Traverse(val reflect.Value, convertNilPtr bool, fn func(
	val reflect.Value, field reflect.StructField,
) bool) bool {
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
				if convertNilPtr && fieldVal.CanSet() {
					fieldVal.Set(reflect.New(fieldTyp))
				} else {
					continue
				}
			}
			fieldVal = fieldVal.Elem()
		}

		if field.Anonymous && fieldTyp.Kind() == reflect.Struct {
			if !Traverse(fieldVal, convertNilPtr, fn) {
				return false // stop traverse
			}
		} else if field.Name[0] >= 'A' && field.Name[0] <= 'Z' {
			if !fn(fieldVal, field) {
				return false // stop traverse
			}
		}
	}
	return true // go on traverse
}

// TraverseType traverses a reflect.Type
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
