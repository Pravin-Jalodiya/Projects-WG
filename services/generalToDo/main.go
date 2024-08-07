package generalToDo

import (
	"fmt"
	"projects/config"
)

func Main(currentUser string) {
	var choice int
	for {
		fmt.Printf("\n%sMANAGE TODO%s\n\nPlease select an option\n1. View ToDo list\n2. Add Task\n3. Delete Task\n4. Go back\n", config.STR_DECOR, config.STR_DECOR)
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		switch choice {
		case 1:
			view(currentUser)

		case 2:
			addTask(currentUser)

		case 3:
			deleteTask(currentUser)

		case 4:
			return

		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}
