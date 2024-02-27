package main

import "fmt"

type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)
func (i Fruit) String() string {
	switch i {
	case Apple:
			return "Apple"
	case Orange:
		return "Orange"
	case Banana:
		return "Banana"
	}
	return fmt.Sprintf("Fruit(%d)", i)
}

func main() {
	var f Fruit
	f = Apple
	fmt.Println(f)
}
