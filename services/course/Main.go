package course

import (
	"fmt"
	"projects/config"
)

func Main(currentUser string) {
	var choice int
	for {
		fmt.Printf("\n%sMANAGE COURSES%s\n\nPlease select an option\n2. View course list\n3. Update course list\n4. Go back\n", config.STR_DECOR, config.STR_DECOR)
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {

		case 1:

		case 2:
			view(currentUser)

		case 3:
			updateCourseProgress(currentUser)

		case 4:
			return

		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}
