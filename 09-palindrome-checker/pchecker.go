package main

import (
	"fmt"
)

func reverseWord(word string) string {
	reversed := ""
	// word := "racecar" word[6]

	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}

	return reversed
}

func main() {
	var word string
	fmt.Print("Enter the word to check: ")
	fmt.Scanln(&word)

	reversed := reverseWord(word)
	if reversed == word {
		fmt.Printf("%s is a palindrome\n", word)
	} else {
		fmt.Printf("%s is not a palindrome\n", word)
	}
}
