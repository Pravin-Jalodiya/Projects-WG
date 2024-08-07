package dailyStatus

import (
	"fmt"
	"github.com/fatih/color"
)

func Main(currentUser string) {
	var choice int

	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	viewEmoji := "👁️"
	backEmoji := "↩️"
	errorEmoji := "❌"

	for {
		fmt.Printf("\n%s%sDAILY STATUS%s%s\n\n%sPlease select an option:\n1. %s View daily status\n2. %s Go back\n",
			cyan("======"), cyan(" "), cyan("======"), cyan(" "),
			blue(""), viewEmoji, backEmoji)

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println(red(errorEmoji), red("Invalid input"))
			continue
		}

		switch choice {
		case 1:
			view(currentUser)

		case 2:
			return

		default:
			fmt.Println(red(errorEmoji), red("Invalid selection. Please try again."))
		}
	}
}
