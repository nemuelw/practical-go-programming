package main

import (
        "fmt"
)

type Contact struct {
	Name string
	Phone int
}

var contacts []Contact

func addContact() {
    var name string
	var phone int

	fmt.Print("Enter the contact name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter the contact phone number: ")
	_, err = fmt.Scanln(&phone)
	if err != nil {
		fmt.Println(err)
		return
	}

	contact := Contact{Name: name, Phone: phone}
	contacts = append(contacts, contact)
	fmt.Println("Contact added successfully")
}

func listContacts() {
	if len(contacts) == 0 {
			fmt.Println("No contacts found")
			return
	}
	for i, contact := range contacts {
			fmt.Printf("%d. %s - %d\n", i+1, contact.Name, contact.Phone)
	}
}

func main() {
	fmt.Println("Contact list app")

	for {
			var choice int
			fmt.Println("1. Add contact")
			fmt.Println("2. List contacts")
			fmt.Println("3. Exit")
			fmt.Print("Enter your choice: ")
			_, err := fmt.Scanln(&choice)
			if err != nil {
					fmt.Println(err)
					continue
			}

			switch choice {
			case 1:
					addContact()
			case 2:
					listContacts()
			case 3:
					fmt.Println("Exiting ...")
					return
			default:
					fmt.Println("Invalid option! Try again.")
			}

			fmt.Println()
	}
}