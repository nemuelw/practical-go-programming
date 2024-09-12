package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func hash_content(content []byte) string {
	hasher := sha256.New()
	hasher.Write(content)
	result := hasher.Sum(nil)
	hash := hex.EncodeToString(result)
	return hash
}

// c0535e4be2b79ffd93291305436bf889314e4a3faec05ecffcbb7df31ad9e51a

func main() {
	// ./fichecker filename originalhash
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./fichecker <filename> <originalhash>")
		return
	}

	file_name := os.Args[1]
	original_hash := os.Args[2]
	contents, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	hash := hash_content(contents)
	if hash == original_hash {
		fmt.Println("The file is intact")
	} else {
		fmt.Println("The file is not intact and has been tampered with")
	}
}
