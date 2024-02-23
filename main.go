package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		// iを別の変数に入れる
		v := i
		wg.Add(1)
		go func () {
			defer wg.Done()
			// vはメモリ領域を別で確保されているから正しく出力できるってことかな？
			fmt.Println(v)
		}()
	}
	wg.Wait()
}
