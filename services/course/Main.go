package course

import (
	"fmt"
	"github.com/fatih/color"
)

func Main(currentUser string) {
	var choice int

	red := color.New(color.FgRed).SprintFunc()
	//green := color.New(color.FgGreen).SprintFunc()
	//yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	registerEmoji := "📝" // Registration emoji
	viewEmoji := "👁️"    // View emoji
	updateEmoji := "🔄"   // Update emoji
	backEmoji := "↩️"    // Go back emoji
	errorEmoji := "❌"    // Error emoji

	for {
		fmt.Printf("\n%s%sMANAGE COURSES%s%s\n\n%sPlease select an option:\n1. %s Register for courses\n2. %s View course list\n3. %s Update course list\n4. %s Go back\n",
			cyan("======"), cyan(" "), cyan("======"), cyan(" "),
			blue(""), registerEmoji, viewEmoji, updateEmoji, backEmoji)

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println(red(errorEmoji), red("Invalid input:"), err)
			continue
		}

		switch choice {
		case 1:
			registration(currentUser)
		case 2:
			view(currentUser)

		case 3:
			updateCourseProgress(currentUser)

		case 4:
			return

		default:
			fmt.Println(red(errorEmoji), red("Invalid selection. Please try again."))
		}
	}
}
