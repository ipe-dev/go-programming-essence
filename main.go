package main

import (
	"time"
)
func sendMessage(msg string) {
	println(msg)
}

func main() {
	message := "hi"
	go func() {
		sendMessage(message)
	}()
	// race condition
	// goroutineと呼び出し元の間に競合がある
	message = "ho"

	time.Sleep(time.Second)
	println(message)
	time.Sleep(time.Second)
}
