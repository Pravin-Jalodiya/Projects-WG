package services

import (
	"fmt"
	"projects/config"
	"projects/middleware"
	"projects/services/course"
	"projects/services/dailyStatus"
	"projects/services/todo"
)

func Main() {
	var choice int
	for {
		fmt.Printf("\n%sWELCOME TO INTERNS PORTAL%s\n\nPlease select an option\n1. Register for courses\n2. Manage ToDo list\n3. Daily Status\n4. Log out\n", config.STR_DECOR, config.STR_DECOR)
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {
		case 1:
			course.Registeration(middleware.ActiveUser)

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
