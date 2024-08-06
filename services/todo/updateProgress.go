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
	fmt.Print("Enter MID of the module to be marked as done: ")
	var MID float32
	fmt.Scan(&MID)
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

	dailyStatus.UpdateStatus(currentUser, completedModules)

	fmt.Printf("Module %.1f marked as done\nDaily status updated\n", MID)
	_, err := writers.FWriterToDo(config.USER_FILE, readers.UserStore)
	if err != nil {
		fmt.Println("Error writing to file.")
		return
	}
}
