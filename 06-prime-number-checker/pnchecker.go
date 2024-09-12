package main

import (
	"fmt"
)

func isPrime(number int) bool {
	var factors int = 0
	for i := 1; i <= number; i++ {
		if number % i == 0 {
			factors++
		}
	}
	if factors == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	var number int
	fmt.Print("Enter the number to check: ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	if isPrime(number) {
		fmt.Printf("%d is a prime number\n", number)
	} else {
		fmt.Printf("%d is not a prime number\n", number)
	}
}
