package main

import (
	"fmt"
	"os"
	"projects/controllers"
)

func main() {
	var choice int
	for {
		fmt.Printf("\n-----------------Batch 4 Management System-----------------\n\nPlease select an option\n1. Sign Up\n2. Log In\n3. Exit\n")
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
