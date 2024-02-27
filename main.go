//go:generate stringer -type Fruit
package main

import "fmt"

type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)

func main() {
	f := Apple
	fmt.Println(f)
}
