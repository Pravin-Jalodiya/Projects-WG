package services

import "fmt"

func markAsDone(currentUser string) {
	var choice int
	for {
		fmt.Printf("\n-----------------ToDo list-----------------\n\nPlease select an option\n1. Mark modules as complete\n2. Go back\n")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {
		case 1:
			markModules(currentUser)

		case 2:
			return

		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}
