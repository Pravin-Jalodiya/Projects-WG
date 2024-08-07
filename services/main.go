package services

import (
	"fmt"
	"projects/config"
	"projects/middleware"
	"projects/services/course"
	"projects/services/dailyStatus"
	"projects/services/generalToDo"
	"projects/services/progress"
)

func Main() {
	var choice int
	for {
		fmt.Printf("\n%sWELCOME TO INTERNS PORTAL%s\n\nPlease select an option\n1. Manage course list\n2. Manage ToDo list\n3. Daily status\n4. View progress\n5. Log out\n", config.STR_DECOR, config.STR_DECOR)
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {

		case 1:
			course.Main(middleware.ActiveUser)

		case 2:
			generalToDo.Main(middleware.ActiveUser)

		case 3:
			dailyStatus.Main(middleware.ActiveUser)

		case 4:
			progress.View(middleware.ActiveUser)

		case 5:
			middleware.ActiveUser = ""
			fmt.Println("User Logged out")
			return

		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}
