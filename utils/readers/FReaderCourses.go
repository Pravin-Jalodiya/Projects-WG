package readers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"projects/config"
	"projects/models"
)

var (
	Courses []models.Course
)

func init() {
	Courses = FReaderCourses(config.COURSE_FILE, os.O_RDONLY)
}

func FReaderCourses(f string, flag int) []models.Course {

	var courses []models.Course

	file, err := os.OpenFile(f, flag, 0644)
	if err != nil {
		fmt.Println("Error opening file")
		//log error
	}
	byteValue, _ := io.ReadAll(file)
	err = json.Unmarshal(byteValue, &courses)
	if err != nil {
		return nil
	}
	return courses
}
