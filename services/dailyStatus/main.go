package dailyStatus

import (
	"fmt"
	"projects/config"
)

func Main(currentUser string) {
	var choice int
	for {
		fmt.Printf("\n%sDAILY STATUS%s\n\nPlease select an option\n1. View daily status\n2. Go back\n", config.STR_DECOR, config.STR_DECOR)
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input")
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
