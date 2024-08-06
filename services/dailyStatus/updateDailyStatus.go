package dailyStatus

import (
	"projects/models"
	"projects/utils/readers"
	time2 "time"
)

func UpdateStatus(currentUser string, completedModules []models.Module) {

	dateTime := time2.Now()
	dailyStatusUpdate := models.DailyStatus{
		Date:            dateTime.Format("2006-01-02"),
		Time:            dateTime.Format("15:04:05"),
		TopicsCompleted: completedModules,
	}

	for i, val1 := range readers.UserStore {
		if val1.Username == currentUser {
			readers.UserStore[i].DailyStatus = append(readers.UserStore[i].DailyStatus, dailyStatusUpdate)
		}
	}
}
