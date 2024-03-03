package main

import (
	_ "embed"
	"go-programming-essence/server"
	"log"
)

func main() {
	svr := server.New("localhost", 8888)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
