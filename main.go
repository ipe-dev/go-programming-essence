package main

import "fmt"

type F struct {
	Name string
	Age  int
}

func main() {
	f := &F{
		Name: "John",
		Age:  20,
	}

	// &{John 20}
	fmt.Printf("%v\n", f)

	// &{Name:John Age:20}
	fmt.Printf("%+v\n", f)

	// &main.F{Name:"John", Age:20}
	fmt.Printf("%#v\n", f)

	// *main.F
	fmt.Printf("%T\n", f)

	// main.F
	fmt.Printf("%T\n", *f)
}
