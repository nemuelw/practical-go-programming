package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
)

func main() {
	fmt.Printf("Name: %s\n", faker.Name())
	fmt.Printf("Email: %s\n", faker.Email())
	fmt.Printf("Username: %s\n", faker.Username())
	fmt.Printf("Password: %s\n", faker.Password())
}
