package course

import (
	"bufio"
	"fmt"
	"os"
	"projects/config"
	"projects/utils/course"
	"projects/utils/readers"
	"projects/utils/writers"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func cidFound(cid int, currentUser string) bool {
	for _, user := range readers.UserStore {
		if user.Username == currentUser {
			for _, toDo := range user.ToDo {
				if toDo.CID == cid {
					return true
				}
			}
		}
	}
	return false
}

func registration(currentUser string) {

	course.View()

	courseChoice := make([]int, 0)
	for {
		fmt.Print("\nEnter CID of course you want to enroll in (put 0 to exit) : ")
		cidString, err := reader.ReadString('\n')
		cidString = strings.TrimSuffix(cidString, "\n")
		cidString = strings.TrimSpace(cidString)
		if err != nil {
			fmt.Println("Error reading age.")
			return
		} else {
			cid, err := strconv.Atoi(cidString)
			if err != nil {
				fmt.Println("Invalid Number")
				continue
			} else {
				if cid >= config.COURSE_FIRST && cid <= config.COURSE_LAST {
					if !cidFound(cid, currentUser) {
						courseChoice = append(courseChoice, cid)
					} else {
						fmt.Println("You are already enrolled in this course.")
					}
				} else if cid == 0 {
					break
				} else {
					fmt.Println("Invalid CID")
				}
			}
		}
	}

	userCourses := course.Get(courseChoice...)

	for i, val := range readers.UserStore {
		if val.Username == currentUser {
			readers.UserStore[i].ToDo = append(readers.UserStore[i].ToDo, userCourses...)
			readers.UserStore[i].Progress.Courses = append(readers.UserStore[i].Progress.Courses, userCourses...)
			break
		}
	}

	_, err := writers.FWriterToDo(config.USER_FILE, readers.UserStore)
	if err != nil {
		fmt.Println("Error writing to file.")
		return
	}

	fmt.Println("Registration complete.")

}
