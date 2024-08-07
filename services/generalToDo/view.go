package generalToDo

import (
	"fmt"
	"projects/config"
	"projects/utils/readers"
)

func view(currentUser string) {
	for _, v := range readers.UserStore {
		if v.Username == currentUser {
			if len(v.GeneralTodo) == 0 {
				fmt.Println("ToDo list is empty. No pending work.")
				return
			}

			longestTitle := 0

			for _, todo := range v.GeneralTodo {
				if len(todo.Task) > longestTitle {
					longestTitle = len(todo.Task)
				}
			}

			fmt.Println(config.STR_DECOR, "YOUR TODO LIST", config.STR_DECOR)
			fmt.Printf("\nTASK%*s\tDEADLINE%*s\n", longestTitle-(len("Task"))+4, "", 5, "")
			fmt.Println("-------------------------------------------------------------")
			for i, v := range v.GeneralTodo {
				fmt.Printf("%d. %s%*s\t%s%*s\n", i+1, v.Task, longestTitle-len(v.Task)+1, "", v.Deadline, 5, "")
			}
		}
	}
}
