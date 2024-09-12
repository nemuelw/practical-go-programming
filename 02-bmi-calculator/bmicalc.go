package main

import (
	"fmt"
)

func main() {
	var weight, height float64

	fmt.Print("Enter the weight (kg): ")
	fmt.Scanln(&weight)
	fmt.Print("Enter the height (m): ")
	fmt.Scanln(&height)

	bmi := weight / (height * height)
	fmt.Printf("Your BMI is: %.2f\n", bmi)

	// underweight: bmi < 18.5
	// normal weight: bmi 18.5 - 24.9
	// overweight: bmi 25 - 29.9
	// obese: bmi >= 30
	if bmi < 18.5 {
		fmt.Println("Category: Underweight")
	} else if bmi >= 18.5 && bmi <= 24.9 {
		fmt.Println("Category: Normal weight")
	} else if bmi >= 25 && bmi <= 29.9 {
		fmt.Println("Category: Overweight")
	} else {
		fmt.Println("Category: Obese")
	}
}
