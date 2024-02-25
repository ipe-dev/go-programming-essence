package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"sync"
)

func downloadCSV(wg *sync.WaitGroup, urls []string, ch chan []byte) {
	defer wg.Done()
	defer close(ch) //終わったら閉じる (5)

	// HTTPサーバからダウンロード(あんまりCPUを使わない処理)
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Println("cannot download CSV: ", err)
			continue
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Println("cannot read content: ", err)
			continue
		}
		resp.Body.Close()
		ch <- b // main関数にコンテンツを送信(3)
	}
}

func main() {
	urls := []string{
		"http://my-server.com/data01.csv",
		"http://my-server.com/data02.csv",
		"http://my-server.com/data03.csv",
		"http://my-server.com/data04.csv",
		"http://my-server.com/data05.csv",
		"http://my-server.com/data06.csv",
		"http://my-server.com/data07.csv",
		"http://my-server.com/data08.csv",
	}

	// バイト列を転送するためのchannelを作成 (1)
	ch := make(chan []byte)

	var wg sync.WaitGroup
	wg.Add(1)
	go downloadCSV(&wg, urls, ch) // (2)

	// goroutineからコンテンツを受け取る(4)
	for b := range ch {
		r := csv.NewReader(bytes.NewReader(b))
		for {
			records, err := r.Read()
			if err != nil {
				log.Fatal(err)
			}
			// レコードの登録(あんまりCPUを使わない処理)
			insertRecords(records)
		}
	}
	wg.Wait()
}
