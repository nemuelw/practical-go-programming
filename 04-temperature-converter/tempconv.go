package main

import (
	"fmt"
)

// celsius to fahrenheit: (celsius * 9/5) + 32
// fahrenheit to celsius: (fahrenheit - 32) * 5/9

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9/5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5/9
}

func main() {
	fmt.Println("Temperature Converter")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Print("Enter the conversion type (1 or 2): ")

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	if choice != 1 && choice != 2 {
		fmt.Println("Provide either 1 or 2")
		return
	}

	var temp float64
	fmt.Print("Enter the temperature value: ")
	_, err = fmt.Scanln(&temp)
	if err != nil {
		fmt.Println("Invalid temperature input")
		return
	}

	switch choice {
	case 1:
		result := celsiusToFahrenheit(temp)
		fmt.Printf("%.2f degrees celsius = %.2f fahrenheit\n", temp, result)
	case 2:
		result := fahrenheitToCelsius(temp)
		fmt.Printf("%.2f fahrenheit = %.2f degrees celsius\n", temp, result)
	}
}
