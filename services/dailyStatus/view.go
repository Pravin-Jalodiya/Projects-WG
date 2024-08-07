package dailyStatus

import (
	"fmt"
	"github.com/fatih/color"
	"projects/config"
	"projects/utils/readers"
)

func view(currentUser string) {

	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	emptyEmoji := "🚫"
	dateEmoji := "📅"
	timeEmoji := "⏰"
	moduleEmoji := "📘"

	for _, user := range readers.UserStore {
		if user.Username == currentUser {
			if len(user.DailyStatus) == 0 {
				fmt.Println(yellow(emptyEmoji), yellow("Daily status is empty."))
				return
			}

			fmt.Println(blue(config.STR_DECOR), blue("DAILY STATUS"), blue(config.STR_DECOR))
			fmt.Println()

			for _, status := range user.DailyStatus {
				fmt.Printf("%s Date: %s\n", dateEmoji, status.Date)
				fmt.Printf("%s Time: %s\n", timeEmoji, status.Time)
				fmt.Println("Modules completed:")

				for _, topic := range status.TopicsCompleted {
					fmt.Printf("  %s Module ID: %.1f\n", moduleEmoji, topic.MID)
					fmt.Printf("    Module Title: %s\n", green(topic.Title))
				}
				fmt.Println()
			}
		}
	}
}
