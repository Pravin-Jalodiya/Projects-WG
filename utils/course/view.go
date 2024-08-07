package course

import (
	"fmt"
	"github.com/fatih/color"
	"projects/utils/readers"
)

func View() {

	//blue := color.New(color.FgBlue).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	courseEmoji := "📚"

	// Header
	fmt.Println(cyan("==================================================="))
	fmt.Println(cyan("                AVAILABLE COURSES                 "))
	fmt.Println(cyan("==================================================="))
	fmt.Println()

	// Check if there are any courses available
	if len(readers.Courses) == 0 {
		fmt.Println(green("No courses available at the moment."))
		return
	}

	// Display courses
	for _, course := range readers.Courses {
		fmt.Printf("%s Course ID: %d\n", courseEmoji, course.CID)
		fmt.Printf("  %s Course Title: %s\n", green(""), course.Title)
		fmt.Println()
	}

	// Footer
	fmt.Println(cyan("==================================================="))
}
