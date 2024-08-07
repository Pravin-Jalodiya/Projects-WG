package generalToDo

import (
	"fmt"
	"github.com/fatih/color"
	"projects/config"
	"projects/utils/readers"
)

// Define color functions
var (
	blue   = color.New(color.FgBlue).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

func view(currentUser string) {
	for _, v := range readers.UserStore {
		if v.Username == currentUser {
			if len(v.GeneralTodo) == 0 {
				fmt.Println(yellow("ToDo list is empty. No pending work."))
				return
			}

			longestTitle := 0

			for _, todo := range v.GeneralTodo {
				if len(todo.Task) > longestTitle {
					longestTitle = len(todo.Task)
				}
			}

			fmt.Println(blue(config.STR_DECOR), blue("YOUR TODO LIST"), blue(config.STR_DECOR))
			fmt.Printf("\n%sTask%*s\tDeadline%*s\n", cyan(""), longestTitle-(len("Task"))+4, "", 5, "")
			fmt.Println("-------------------------------------------------------------")

			for i, todo := range v.GeneralTodo {
				fmt.Printf("%d. %s%*s\t%s%*s\n", i+1, green(todo.Task), longestTitle-len(todo.Task)+1, "", cyan(todo.Deadline), 5, "")
			}
		}
	}
}
