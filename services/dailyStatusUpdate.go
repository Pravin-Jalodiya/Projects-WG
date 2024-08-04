package services

import (
	"projects/models"
	"projects/utils"
	time2 "time"
)

func dailyStatusUpdate(currentUser string, completedModules []models.Module) {

	dateTime := time2.Now()
	dailyStatusUpdate := models.DailyStatus{
		Date:            dateTime.Format("2006-01-02"),
		Time:            dateTime.Format("15:04:05"),
		TopicsCompleted: completedModules,
	}

	for i, val1 := range utils.UserStore {
		if val1.Username == currentUser {
			utils.UserStore[i].DailyStatus = append(utils.UserStore[i].DailyStatus, dailyStatusUpdate)
		}
	}
}
