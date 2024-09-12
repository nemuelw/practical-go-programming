package main

import (
	"fmt"
	"math/rand"
	"time"
)

const pool = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func main() {
	password_length := 8
	var password string

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < password_length; i++ {
		character := pool[rand.Intn(len(pool))]
		password += string(character)
	}

	fmt.Println("Generated password is: ", password)
}
