package main

import (
	"encoding/json"
	"log"
	"os"
)

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Data struct {
	Species     string     `json:"species"`
	Description string     `json:"description"`
	Dimensions  Dimensions `json:"dimensions"`
}

func main() {
	var data Data
	f, err := os.Open("input.json")
	if err != nil {
		log.Fatal(err)
	}
	// ネットワークやファイルのjsonを扱うときはデコーダを使う
	err = json.NewDecoder(f).Decode(data)
	if err != nil {
		log.Fatal(err)
	}
}
