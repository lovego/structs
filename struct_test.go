package structs

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
