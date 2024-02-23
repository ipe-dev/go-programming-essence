package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func () {
			defer wg.Done()
			// goroutineの中でiを出力するときにはすでにforループが終わってることがある
			// そうすると10しか出力しない
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
