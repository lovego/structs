package structs

import "reflect"

// ShallowCopy shallow copy a struct
func ShallowCopy(src interface{}) interface{} {
	return ShallowCopyV(reflect.ValueOf(src)).Interface()
}

// ShallowCopyV shallow copy a struct value
func ShallowCopyV(src reflect.Value) reflect.Value {
	typ := src.Type()
	dst := reflect.New(typ).Elem()
	for i := 0; i < typ.NumField(); i++ {
		if name := typ.Field(i).Name; name[0] >= 'A' && name[0] <= 'Z' {
			dst.Field(i).Set(src.Field(i))
		}
	}
	return dst
}
