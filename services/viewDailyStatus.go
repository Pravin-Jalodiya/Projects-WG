package services

import (
	"fmt"
	"projects/utils"
)

func viewDailyStatus(currentUser string) {
	fmt.Println("----------------Your daily status----------------")
	for _, v := range utils.UserStore {
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
