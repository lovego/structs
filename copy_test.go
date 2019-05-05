package structs

import "fmt"

func ExampleShallowCopy() {
	type T struct{ A, B, c int }
	src := T{99, 88, 97}
	copy := ShallowCopy(src).(T)

	fmt.Println(copy, &src == &copy)
	// Output: {99 88 0} false
}
