package structs

import "reflect"

// Traverse traverses a reflect.Value
// convertNilPtr:  convert anonymous nil struct pointer to non-nil or not.
// nil pointer to anonymous unexported struct fields are not traversed by Traverse,
// but pointer to anonymous unexported struct fields are always traversed by TraverseType.
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
		fieldVal := val.FieldByIndex(field.Index)

		if field.Anonymous {
			fieldTyp := field.Type
			if fieldTyp.Kind() == reflect.Ptr && fieldTyp.Elem().Kind() == reflect.Struct {
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
			if fieldTyp.Kind() == reflect.Struct {
				if !Traverse(fieldVal, convertNilPtr, fn) {
					return false // stop traverse
				}
				continue
			}
		}
		if field.Name[0] >= 'A' && field.Name[0] <= 'Z' {
			if !fn(fieldVal, field) {
				return false // stop traverse
			}
		}
	}
	return true // go on traverse
}

// TraverseType traverses a reflect.Type
func TraverseType(typ reflect.Type, fn func(field reflect.StructField), index ...int) {
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
				TraverseType(fieldTyp, fn, append(index, field.Index...)...)
				continue
			}
		}
		if field.Name[0] >= 'A' && field.Name[0] <= 'Z' {
			field.Index = append(index, field.Index...)
			fn(field)
		}
	}
}
