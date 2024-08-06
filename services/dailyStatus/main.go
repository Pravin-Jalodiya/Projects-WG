package dailyStatus

import (
	"fmt"
)

func Main(currentUser string) {
	var choice int
	for {
		fmt.Printf("\n-----------------TODO LIST-----------------\n\nPlease select an option\n1. View dailyStatus status\n2. Go back\n")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {
		case 1:
			viewDailyStatus(currentUser)

		case 2:
			return

		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}
