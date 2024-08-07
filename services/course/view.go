package course

import (
	"fmt"
	"projects/config"
	"projects/utils/readers"
)

func view(currentUser string) {
	for _, v := range readers.UserStore {
		if v.Username == currentUser {
			if len(v.ToDo) == 0 {
				fmt.Println("Course list is empty. No pending work.")
				return
			}
			fmt.Println(config.STR_DECOR, "YOUR COURSE LIST", config.STR_DECOR)
			for _, v := range v.ToDo {
				if len(v.Modules) > 0 {
					fmt.Printf("\nCourse ID : %d\nCourse name : %s\n", v.CID, v.Title)
					for _, v := range v.Modules {
						fmt.Printf("Module ID : %.1f\tModule name : %s\n", v.MID, v.Title)
					}
				}
			}
		}
	}
}
