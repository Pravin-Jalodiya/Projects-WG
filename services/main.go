package services

import (
	"fmt"
	"projects/middleware"
	"projects/services/dailyStatus"
	"projects/services/todo"
)

func Main() {
	var choice int
	for {
		fmt.Printf("\n-----------------WELCOME TO INTERNS PORTAL-----------------\n\nPlease select an option\n1. Register for courses\n2. Manage ToDo list\n3. Daily Status\n4. Log out\n")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {
		case 1:

		case 2:
			todo.Main(middleware.ActiveUser)

		case 3:
			dailyStatus.Main(middleware.ActiveUser)

		case 4:
			middleware.ActiveUser = ""
			fmt.Println("User Logged out")
			return

		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}
