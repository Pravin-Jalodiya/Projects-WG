package course

import (
	"projects/models"
	"projects/utils/readers"
	"slices"
)

func Get(nums ...int) []models.Course {
	var courses []models.Course
	for _, val := range readers.Courses {
		if slices.Contains(nums, val.CID) {
			courses = append(courses, val)
		}
	}
	return courses
}
