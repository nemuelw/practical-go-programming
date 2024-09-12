package main

import (
	"bufio"
	"fmt"
	"os"
)

var todos []string

func addToDo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the task description: ")
	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	todos = append(todos, task)
	fmt.Println("Task added successfully")
}

func listToDos() {
	if len(todos) == 0 {
		fmt.Println("No tasks found")
		return
	}
	for i, todo := range todos {
		fmt.Printf("%d. %s", i+1, todo)
	}
}

func main() {
	fmt.Println("To-Do list app")

	for {
		var choice int
		fmt.Println("1. Add To-Do item")
		fmt.Println("2. List To-Dos")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch choice {
		case 1:
			addToDo()
		case 2:
			listToDos()
		case 3:
			fmt.Println("Exiting ...")
			return
		default:
			fmt.Println("Invalid option! Try again.")
		}

		fmt.Println()
	}
}
