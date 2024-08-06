package todo

import (
	"fmt"
)

func Main(currentUser string) {
	var choice int
	for {
		fmt.Printf("\n-----------------TODO LIST-----------------\n\nPlease select an option\n1. View ToDo list\n2. Update Progress\n3. Go back\n")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {
		case 1:
			viewToDo(currentUser)

		case 2:
			updateProgress(currentUser)

		case 3:
			return

		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}
