package todo

import (
	"fmt"
	"projects/utils/readers"
)

func viewToDo(currentUser string) {
	fmt.Println("----------------Your to do list----------------")
	for _, v := range readers.UserStore {
		//fmt.Println(v, v.Username)
		if v.Username == currentUser {
			for _, v := range v.ToDo {
				fmt.Printf("\nCourse ID : %d\nCourse name : %s\n", v.CID, v.Title)
				for _, v := range v.Modules {
					fmt.Printf("Module ID : %.1f\tModule name : %s\n", v.MID, v.Title)
				}
			}
		}
	}
}
