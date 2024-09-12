package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findLongestWord(text string) string {
	longestWord := ""
	maxLength := 0
	// Welcome to programming: {"Welcome", "to", "programming"}

	words := strings.Fields(text)

	for _, word := range words {
		if len(word) > maxLength {
			longestWord = word
			maxLength = len(word)
		}
	}

	return longestWord
}

func main() {
	var text string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a sentence or paragraph:")
	text, _ = reader.ReadString('\n')

	longestWord := findLongestWord(text)
	fmt.Printf("The longest word is: %s\n", longestWord)
}
