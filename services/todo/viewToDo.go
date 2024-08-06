package todo

import (
	"fmt"
	"projects/config"
	"projects/utils/readers"
)

func viewToDo(currentUser string) {
	fmt.Println(config.STR_DECOR, "YOUR TODO LIST", config.STR_DECOR)
	for _, v := range readers.UserStore {
		if v.Username == currentUser {
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
