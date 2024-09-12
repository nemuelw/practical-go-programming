package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {
	http.HandleFunc("/", helloWorldHandler) // localhost:1234/
	fmt.Println("Server running on port 1234 ...")
	http.ListenAndServe("localhost:1234", nil)
}
