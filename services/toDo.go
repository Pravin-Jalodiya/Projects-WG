package services

import (
	"fmt"
)

func toDo(currentUser string) {
	var choice int
	for {
		fmt.Printf("\n-----------------ToDo list-----------------\n\nPlease select an option\n1. View ToDo list\n2. Mark as done\n3. Go back\n")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {
		case 1:
			viewToDo(currentUser)

		case 2:
			markAsDone(currentUser)

		case 3:
			return

		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}
