package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://api.quotable.io/random"

type Quote struct {
	Content string `json:"content"`
	Author string `json:"author"`
}

func main() {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var quote Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s (%s)\n", quote.Content, quote.Author)
}
