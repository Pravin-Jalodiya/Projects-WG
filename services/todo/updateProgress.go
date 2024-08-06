package todo

import (
	"fmt"
	"projects/config"
	"projects/models"
	"projects/services/dailyStatus"
	"projects/utils/readers"
	"projects/utils/writers"
)

func updateProgress(currentUser string) {

	for _, user := range readers.UserStore {
		if user.Username == currentUser {
			for _, toDo := range user.ToDo {
				if len(toDo.Modules) == 0 {
					fmt.Println("Cannot update progress. ToDo list is empty!")
					return
				}

			}

		}
	}

	fmt.Println(config.STR_DECOR, "UPDATE PROGRESS", config.STR_DECOR)
	fmt.Println()
	fmt.Print("Enter MID of the module to be marked as done: ")
	var MID float32

	for {
		fmt.Scan(&MID)
		fmt.Println()
		var isValid = false
		for _, user := range readers.UserStore {
			if user.Username == currentUser {
				for _, course := range user.ToDo {
					for _, module := range course.Modules {
						if module.MID == MID {
							isValid = true
							break
						}
					}
					if isValid {
						break
					}
				}
			}
			if isValid {
				break
			}
		}
		if !isValid {
			fmt.Print("Invalid MID. Try again : ")
		} else {
			break
		}
	}

	var completedModules []models.Module
	for i, val1 := range readers.UserStore {
		if val1.Username == currentUser {
			for j, val2 := range val1.ToDo {
				var filteredModules []models.Module
				for _, val3 := range val2.Modules {
					if val3.MID != MID {
						filteredModules = append(filteredModules, val3)
					} else {
						completedModules = append(completedModules, val3)
					}
				}
				readers.UserStore[i].ToDo[j].Modules = filteredModules
			}
		}
	}

	if len(completedModules) > 0 {
		dailyStatus.UpdateStatus(currentUser, completedModules)
		fmt.Printf("Module %.1f marked as done\nDaily status updated\n", MID)
		_, err := writers.FWriterToDo(config.USER_FILE, readers.UserStore)
		if err != nil {
			fmt.Println("Error writing to file.")
			return
		}
	}
}
