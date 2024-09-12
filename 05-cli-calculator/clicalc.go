package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("CLI Calculator")
	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1 + num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1 - num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Cannot divide by zero")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1 / num2)
		}
	default:
		fmt.Println("You provided an invalid operator. Provide either +,-,* or /")
	}
}
