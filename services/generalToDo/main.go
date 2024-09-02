package generalToDo

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
	addTaskEmoji := "➕"
	deleteTaskEmoji := "🗑️"
	backEmoji := "↩️"
	errorEmoji := "❌"

	for {
		fmt.Printf("\n%s%sMANAGE TODO%s%s\n\n%sPlease select an option:\n1. %s View ToDo list\n2. %s Add Task\n3. %s Delete Task\n4. %s Add Task via API\n5. %s Delete Task via API\n6. %s Go back\n",
			cyan("======"), cyan(" "), cyan("======"), cyan(" "),
			blue(""), viewEmoji, addTaskEmoji, deleteTaskEmoji, addTaskEmoji, deleteTaskEmoji, backEmoji)

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println(red(errorEmoji), red("Invalid input"))
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
			addTaskViaApi(currentUser)
		case 5:
			deleteTaskViaApi(currentUser)

		case 6:
			return

		default:
			fmt.Println(red(errorEmoji), red("Invalid selection. Please try again."))
		}
	}
}
