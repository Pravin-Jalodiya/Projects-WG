package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"projects/controllers"
)

func main() {
	var choice int
	red := color.New(color.FgRed).SprintFunc()

	blue := color.New(color.FgBlue).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	signUpEmoji := "✍️"
	loginEmoji := "🔑"
	exitEmoji := "🚪"
	errorEmoji := "❌"

	for {
		fmt.Printf("\n%s%sBATCH 4 MANAGEMENT SYSTEM%s%s\n\n%sPlease select an option:\n1. %s Sign Up\n2. %s Log In\n3. %s Exit\n",
			cyan("======"), cyan(" "), cyan("======"), cyan(" "),
			blue(""), signUpEmoji, loginEmoji, exitEmoji)

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println(red(errorEmoji), red("Invalid input:"), err)
			continue
		}

		switch choice {
		case 1:
			controllers.SignUp()

		case 2:
			controllers.Login()

		case 3:
			fmt.Println(blue(exitEmoji), blue("Exiting..."))
			os.Exit(0)

		default:
			fmt.Println(red(errorEmoji), red("Invalid selection. Please try again."))
		}
	}
}
