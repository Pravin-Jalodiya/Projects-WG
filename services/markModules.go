package services

import (
	"fmt"
	"projects/config"
	"projects/models"
	"projects/utils"
)

func markModules(currentUser string) {
	//fmt.Println(utils.UserStore)
	fmt.Print("Enter MID of the module to be marked as done: ")
	var MID float32
	fmt.Scan(&MID)
	var completedModules []models.Module
	for i, val1 := range utils.UserStore {
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
				utils.UserStore[i].ToDo[j].Modules = filteredModules
			}
		}
	}

	dailyStatusUpdate(currentUser, completedModules)

	fmt.Printf("Module %.1f marked as done\nDaily status updated\n", MID)
	_, err := utils.FWriterToDo(config.USER_FILE, utils.UserStore)
	if err != nil {
		fmt.Println("Error writing to file.")
		return
	}

}
