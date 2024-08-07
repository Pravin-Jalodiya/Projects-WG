package progress

import (
	"projects/models"
	"projects/utils/readers"
	"slices"
)

func Update(currentUser string, completedModules []models.Module) {
	for i, val1 := range readers.UserStore {
		if val1.Username == currentUser {
			for _, val2 := range completedModules {
				if !slices.Contains(val1.Progress.ModulesFinished, val2.MID) {
					readers.UserStore[i].Progress.ModulesFinished = append(readers.UserStore[i].Progress.ModulesFinished, val2.MID)
				}
			}
		}
	}
}
