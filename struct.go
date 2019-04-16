package structs

import (
	"fmt"
	"reflect"
	"strings"
)

var stringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

// Println print a struct without zero value fields.
func Println(structs ...interface{}) {
	var result = make([]interface{}, len(structs))
	for i, m := range structs {
		result[i] = Sprint(m)
	}
	fmt.Println(result...)
}

func Sprint(strct interface{}) string {
	if stringer, ok := strct.(fmt.Stringer); ok {
		rv := reflect.ValueOf(strct)
		if !(rv.Kind() == reflect.Ptr && rv.IsNil() && rv.Type().Elem().Implements(stringerType)) {
			return stringer.String()
		}
	}
	structV := reflect.ValueOf(strct)
	if structV.Kind() == reflect.Ptr {
		if structV.IsNil() {
			return fmt.Sprint(strct)
		} else {
			structV = structV.Elem()
		}
	}
	if structV.Kind() != reflect.Struct {
		return fmt.Sprint(strct)
	}

	slice := []string{}
	typ := structV.Type()
	for i := 0; i < structV.NumField(); i++ {
		if name := typ.Field(i).Name; name[0] >= 'A' && name[0] <= 'Z' {
			if value := structV.Field(i); !reflect.DeepEqual(
				value.Interface(), reflect.Zero(value.Type()).Interface(),
			) {
				slice = append(slice, fmt.Sprintf("%v:%v", name, Sprint(value.Interface())))
			}
		}
	}
	return fmt.Sprintf("{%s}", strings.Join(slice, " "))
}
