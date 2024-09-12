package main

import (
	"fmt"
	"time"
)

func main() {
	var birthYear int
	fmt.Print("Enter your year of birth: ")
	fmt.Scanln(&birthYear)

	currentYear := time.Now().Year()
	age := currentYear - birthYear
	fmt.Printf("You are %d years old\n", age)
}
