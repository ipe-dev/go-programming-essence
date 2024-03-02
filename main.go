package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func f(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// 中断
			return
		default:
		}
		fmt.Println("goroutine:処理")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// goroutineを起動した側で処理の中断をするため、contextとcancelを生成
	ctx, cancel := context.WithCancel(context.Background())
	go f(ctx, &wg)
	time.Sleep(10 * time.Second)

	// contextで処理を中止
	cancel()
	wg.Wait()
}
