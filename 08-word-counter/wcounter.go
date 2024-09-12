package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var text string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a sentence or paragraph:")
	text, _ = reader.ReadString('\n')

	words := strings.Fields(text)
	count := len(words)
	fmt.Printf("The number of words is: %d\n", count)
}
