package readers

import (
	"os"
	"projects/config"
)

func SyncCourseData() {
	Courses = FReaderCourses(config.COURSE_FILE, os.O_RDONLY)
}
