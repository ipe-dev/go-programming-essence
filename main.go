package main

import (
	"fmt"
)

func server(ch chan string) {
	defer close(ch)
	// チャネルにデータを送信
	ch <- "one"
	ch <- "two"
	ch <- "three"
}

func main() {
	var s string

	// channelを作成
	ch := make(chan string)

	go server(ch)

	s = <-ch
	fmt.Println(s) // one

	s = <-ch
	fmt.Println(s) // two

	s = <-ch
	fmt.Println(s) // three
}
