package structs

import (
	"fmt"
	"reflect"
	"regexp"
)

func ExampleTraverse() {
	value := getTestValue()
	Traverse(reflect.ValueOf(value), true, nil, func(v reflect.Value, f reflect.StructField) bool {
		fmt.Println(f.Name)
		switch v.Kind() {
		case reflect.Int:
			v.SetInt(8)
		case reflect.String:
			v.SetString(f.Name)
		case reflect.Bool:
			v.SetBool(true)
			return false
		}
		return true
	})
	fmt.Println(regexp.MustCompile("0x[0-9a-f]+").ReplaceAllLiteralString(
		fmt.Sprintf("%+v", value),
		fmt.Sprintf("%+v", reflect.ValueOf(value).Elem().FieldByName("T5")),
	))
	// Output:
	// T1
	// T2
	// T3
	// T4
	// T5
	// Stop
	// &{T1:T1 T2:{T2:T2} t3:{T3:T3} T4:8 T5:&{T5:T5} t6:<nil> Stop:true AfterStop: notExported:0}
}

func ExampleTraverseType() {
	TraverseType(reflect.TypeOf(getTestValue()), nil, func(f reflect.StructField) {
		fmt.Println(f.Name, f.Index)
	})
	// Output:
	// T1 [0]
	// T2 [1 0]
	// T3 [2 0]
	// T4 [3]
	// T5 [4 0]
	// T6 [5 0]
	// Stop [6]
	// AfterStop [7]
}

func getTestValue() interface{} {
	type T2 struct {
		T2 string
	}
	type t3 struct {
		T3 string
	}
	type T4 int
	type T5 struct {
		T5 string
	}
	type t6 struct {
		T6 string
	}

	return &struct {
		T1 string
		T2
		t3
		T4
		*T5
		*t6
		Stop        bool
		AfterStop   string
		notExported int
	}{}
}
