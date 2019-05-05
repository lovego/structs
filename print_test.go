package structs

import "time"

func ExamplePrintln() {
	type c struct {
		C int
	}
	type D struct {
		D int
	}
	var s = struct {
		A int
		B string
		C c
		D
		E c
	}{
		A: 1, B: "", C: c{C: 3}, D: D{D: 4},
	}
	Println(s)

	// Output: {A:1 C:{C:3} D:{D:4}}
}

func ExamplePrintln_nilStringer() {
	t := time.Date(2019, 1, 2, 10, 11, 12, 0, time.UTC)

	Println(struct{ Time *time.Time }{})
	Println(struct{ Time *time.Time }{&t})
	// Output:
	// {}
	// {Time:2019-01-02 10:11:12 +0000 UTC}
}
