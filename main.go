//go:generate stringer -type Fruit
package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("hoge.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	log.Println("app started")
}
