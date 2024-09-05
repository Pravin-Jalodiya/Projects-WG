package generalToDo

import (
	"fmt"
	"net/http"
	"projects/config"
	"strconv"
	"strings"
)

func deleteTaskViaApi(currentUser string) {

	view(currentUser)

	fmt.Printf("\n%sDELETE TASK%s\n", config.STR_DECOR, config.STR_DECOR)
	var username string
	var taskNo int
	var err error

	for {
		fmt.Print("Enter username: ")
		username, err = reader.ReadString('\n')
		username = strings.TrimSuffix(username, "\n")
		username = strings.TrimSpace(username)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		if username != currentUser {
			fmt.Println("Enter your username only! ;(")
			continue
		}
		break
	}

	for {
		fmt.Print("Enter Task no: ")
		taskStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		taskStr = strings.TrimSpace(taskStr)
		taskNo, err = strconv.Atoi(taskStr)
		if err != nil {
			fmt.Println("Enter a valid number.")
			continue
		}
		break
	}

	url := fmt.Sprintf("http://localhost:8080/api/todo/update/%s?task=%d", username, taskNo)

	req, err := http.NewRequest("DELETE", url, nil) // No request body needed
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Task deleted successfully.")
	} else {
		fmt.Printf("Failed to delete task. Status code: %d\n", resp.StatusCode)
	}
}
