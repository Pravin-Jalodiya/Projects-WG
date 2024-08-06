package main

import (
	"fmt"
	"os"
	"projects/config"
	"projects/controllers"
)

func main() {
	var choice int
	for {
		fmt.Printf("\n%sBatch 4 Management System%s\n\nPlease select an option\n1. Sign Up\n2. Log In\n3. Exit\n", config.STR_DECOR, config.STR_DECOR)
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {
		case 1:
			controllers.SignUp()

		case 2:
			controllers.Login()

		case 3:
			os.Exit(0)

		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}
