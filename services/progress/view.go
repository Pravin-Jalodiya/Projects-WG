package progress

import (
	"fmt"
	"projects/config"
	"projects/utils/readers"
)

const (
	courseIDWidth    = 10
	courseTitleWidth = 20
	progressWidth    = 10
	moduleTitleWidth = 20
	moduleIDWidth    = 5
	statusWidth      = 5
)

func View(currentUser string) {

	for _, user := range readers.UserStore {
		if user.Username == currentUser {
			completedModules := user.Progress.ModulesFinished
			if len(user.Progress.Courses) == 0 {
				fmt.Println("Register to at least 1 course to check progress.")
				return
			}
			fmt.Println(config.STR_DECOR, "PROGRESS", config.STR_DECOR)
			for _, course := range user.Progress.Courses {
				fmt.Printf("Course ID: %d\nCourse Title: %s\n", course.CID, course.Title)
				fmt.Printf("Progress: %d%%\n\n", calculateProgress(currentUser, course.CID, completedModules))

				longestTitle := 0
				for _, module := range course.Modules {
					if len(module.Title) > longestTitle {
						longestTitle = len(module.Title)
					}
				}

				fmt.Printf("Module Title%*s\tModule ID%*s\tStatus%*s\n", longestTitle-len("Module Title")+4, "", moduleIDWidth, "", statusWidth, "")
				fmt.Println("---------------------------------------------------------")

				for i, module := range course.Modules {
					status := putEmoji(module.MID, completedModules)
					fmt.Printf("%d. %s%*s\t%.1f%*s\t%s%*s\n", i+1, module.Title, longestTitle-len(module.Title)+1, "", module.MID, moduleIDWidth, "", status, statusWidth, "")
				}
				fmt.Println()
			}
			break
		}
	}

}

func calculateProgress(user string, cid int, completedModule []float32) int {
	//traverse each course and then check MIDs against common completed module slices
	totalProgress := 0
	for _, val := range readers.UserStore {
		if val.Username == user {
			for _, v := range val.Progress.Courses {
				if v.CID == cid {
					for _, m := range v.Modules {
						if isCompleted(m.MID, completedModule) {
							totalProgress += m.Weightage
						}
					}
				}
			}
		}
	}

	return totalProgress
}

func putEmoji(mid float32, completedModules []float32) string {

	if isCompleted(mid, completedModules) {
		return "✅"
	}
	return ""
}

func isCompleted(mid float32, completedModules []float32) bool {
	for _, moduleID := range completedModules {
		if moduleID == mid {
			return true
		}
	}
	return false
}
