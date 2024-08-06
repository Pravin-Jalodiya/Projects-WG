package course

import (
	"fmt"
	"projects/utils/readers"
)

func View() {
	fmt.Println("-------------AVAILABLE COURSES-------------")
	fmt.Println()
	for _, course := range readers.Courses {
		fmt.Println("Course ID: ", course.CID)
		fmt.Println("Course Title: ", course.Title)
	}
}
