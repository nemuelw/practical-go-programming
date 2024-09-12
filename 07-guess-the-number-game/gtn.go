package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	attempts := 0
	var guess int

	fmt.Println("Welcome to Guess the Number game")
	fmt.Println("I am thinking of a number between 1 and 100. Can you guess it?")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)
		attempts++

		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed the number %d with %d attempts.\n", target, attempts)
			break
		}
	}
}
