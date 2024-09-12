package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://v2.jokeapi.dev/joke/Any"

type Joke struct {
    Setup string
	Delivery string
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

	var joke Joke
	err = json.Unmarshal(body, &joke)
	if err != nil {
			fmt.Println(err)
			return
	}
	if joke == (Joke{}) {
		fmt.Println("No joke returned")
	} else {
		fmt.Println(joke.Setup)
		fmt.Println(joke.Delivery)
	}
	
}
