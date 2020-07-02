package structs

import (
	"fmt"
	"reflect"
)

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
