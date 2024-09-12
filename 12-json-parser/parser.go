package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string
	Age int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Syntax: ./parser <json_file>")
		return
	}

	json_file := os.Args[1]
	contents, err := os.ReadFile(json_file)
	if err != nil {
		fmt.Println("Error reading from file")
		return
	}

	var user User
	err = json.Unmarshal(contents, &user)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Name: %s\n", user.Name)
		fmt.Printf("Age:  %d\n", user.Age)
	}
}
