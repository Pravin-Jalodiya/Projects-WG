package dailyStatus

import (
	"fmt"
	"projects/utils/readers"
)

func viewDailyStatus(currentUser string) {
	fmt.Println("----------------Your dailyStatus status----------------")
	for _, v := range readers.UserStore {
		if v.Username == currentUser {
			for _, v := range v.DailyStatus {
				fmt.Printf("\nDate : %s\nTime : %s\nModules completed\n", v.Date, v.Time)
				for _, v := range v.TopicsCompleted {
					fmt.Printf("Module ID: %.1f\tModule Title: %s\n", v.MID, v.Title)
				}
			}
		}
	}
}
