package course

import (
	"fmt"
	"github.com/fatih/color"
	"projects/config"
	"projects/utils/readers"
)

func view(currentUser string) {

	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	emptyEmoji := "🚫"
	moduleEmoji := "📖"

	userFound := false

	for _, user := range readers.UserStore {
		if user.Username == currentUser {
			userFound = true
			if len(user.ToDo) == 0 {
				fmt.Println(yellow(emptyEmoji), yellow("Course list is empty. No pending work."))
				return
			}
			fmt.Println(blue(config.STR_DECOR), blue("YOUR COURSE LIST"), blue(config.STR_DECOR))
			for _, course := range user.ToDo {
				if len(course.Modules) > 0 {
					fmt.Printf("\nCourse ID: %d\n%s Course name: %s\n",
						course.CID, blue(""), course.Title)
					for _, module := range course.Modules {
						fmt.Printf("  %s Module ID: %.1f\tModule name: %s\n",
							moduleEmoji, module.MID, module.Title)
					}
				}
			}
		}
	}
	fmt.Println()
	if !userFound {
		fmt.Println(red("User not found."))
	}
}
