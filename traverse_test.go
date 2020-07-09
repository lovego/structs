package structs

import (
	"fmt"
	"reflect"
)

func ExampleTraverse() {
	type TestT2 struct {
		T2Name string
	}
	type testT3 struct {
		T3Name string
	}
	type TestT4 int
	type testT5 string
	type TestT6 struct {
		T6Name string
	}

	type TestT struct {
		Name        string
		notExported int
		TestT2
		*testT3
		TestT4
		testT5
		TestT6
		Stop      bool
		AfterStop string
	}

	value := &TestT{}
	Traverse(reflect.ValueOf(value), func(v reflect.Value, f reflect.StructField) bool {
		if !v.CanSet() {
			fmt.Println(f)
			return false
		}
		switch v.Kind() {
		case reflect.Int:
			v.SetInt(8)
		case reflect.String:
			v.SetString(f.Name)
		case reflect.Bool:
			v.SetBool(true)
			return true
		}
		return false
	})
	fmt.Printf("%+v\n", *value)
	// Output:
	// {Name:Name notExported:0 TestT2:{T2Name:T2Name} testT3:<nil> TestT4:8 testT5: TestT6:{T6Name:T6Name} Stop:true AfterStop:}
}

func ExampleTraverseType() {
	type TestT2 struct {
		T2Name string
	}
	type testT3 struct {
		T3Name string
	}
	type TestT4 int
	type testT5 string

	type TestT struct {
		Name        string
		notExported int
		TestT2
		*testT3
		TestT4
		testT5
	}
	TraverseType(reflect.TypeOf(TestT{}), func(f reflect.StructField) {
		fmt.Println(f.Name)
	})
	// Output:
	// Name
	// T2Name
	// T3Name
	// TestT4
}
