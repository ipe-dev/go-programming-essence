package main

import "sync"

func main() {

	// main関数が終了するとgoroutineも強制終了してしまう
	// それを防ぐためにsync.WaitGroupを使ってgoroutineの終了を待つ
	var wg sync.WaitGroup
	// リファレンスカウンタを+1
	wg.Add(1)
	go func() {
		// リファレンスカウンタを-1
		defer wg.Done()
		// 重い処理
		print("重い処理2")
	}()

	// 別の重い処理
	print("重い処理1")
	// リファレンスカウンタが0になるまで待つ
	wg.Wait()
}
