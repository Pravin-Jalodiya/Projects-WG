package generalToDo

import (
	"fmt"
	"projects/config"
	"projects/utils/readers"
	"projects/utils/writers"
	"strconv"
	"strings"
)

func deleteTask(currentUser string) {
	view(currentUser)

	for _, v := range readers.UserStore {
		if v.Username == currentUser {
			if len(v.GeneralTodo) == 0 {
				return
			}
			fmt.Printf("\n%sDELETE TASK%s\n", config.STR_DECOR, config.STR_DECOR)
			var taskNo int

			for {
				fmt.Print("Enter Task s.no: ")
				sno, err := reader.ReadString('\n')
				fmt.Println()
				if err != nil {
					fmt.Println("Error reading input. Try again,")
					continue
				}
				sno = strings.TrimSuffix(sno, "\n")
				sno = strings.TrimSpace(sno)
				taskNo, err = strconv.Atoi(sno)
				if err != nil {
					fmt.Println("Enter a valid number.")
					continue
				}
				break
			}

			for i, val := range readers.UserStore {
				if val.Username == currentUser {
					for j, _ := range val.GeneralTodo {
						if j+1 == taskNo {
							readers.UserStore[i].GeneralTodo = append(readers.UserStore[i].GeneralTodo[:j], readers.UserStore[i].GeneralTodo[j+1:]...)
							break
						}
					}
					break
				}
			}

			_, err := writers.FWriterToDo(config.USER_FILE, readers.UserStore)
			if err != nil {
				fmt.Println("Error writing to file.")
				return
			}

			fmt.Println("Task deleted successfully.")
		}
	}
}
