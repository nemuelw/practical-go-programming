package main

import "fmt"

type Person struct {
	Name string
	Age int
	Email string
}

func displayInformation(p Person) {
	fmt.Println("Profile information:")
	fmt.Printf("Name:   %s\n", p.Name)
	fmt.Printf("Age:    %d\n", p.Age)
	fmt.Printf("Email:  %s\n", p.Email)
	fmt.Println()
}

func main() {
	person1 := Person{Name: "Bob", Age: 30, Email: "bob@example.com"}
	person2 := Person{Name: "Alice", Age: 40, Email: "alice@example.com"}

	displayInformation(person1)
	displayInformation(person2)
}
