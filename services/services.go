package services

import (
	"fmt"
	"projects/middleware"
)

func Services() {
	var choice int
	for {
		fmt.Printf("\n-----------------Welcome to Interns Portal-----------------\n\nPlease select an option\n1. Manage ToDo list\n2. Daily Status\n3. Log out\n")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {
		case 1:
			toDo(middleware.ActiveUser)

		case 2:
			dailyStatus(middleware.ActiveUser)

		case 3:
			Logout()
			return

		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}
