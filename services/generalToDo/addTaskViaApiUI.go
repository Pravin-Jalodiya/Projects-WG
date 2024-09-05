package generalToDo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"projects/config"
	"strconv"
	"strings"
)

func addTaskViaApi(currentUser string) {
	fmt.Printf("\n%sADD TASK%s\n", config.STR_DECOR, config.STR_DECOR)
	var username, task string
	var lastDay int
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
		fmt.Print("Enter task: ")
		task, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		task = strings.TrimSpace(task)
		break
	}

	for {
		fmt.Print("Enter deadline (in days): ")
		deadline, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		deadline = strings.TrimSpace(deadline)
		lastDay, err = strconv.Atoi(deadline)
		if err != nil {
			fmt.Println("Enter a valid number.")
			continue
		}
		break
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"username": username,
		"task":     task,
		"last_day": lastDay,
	})
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/api/todo/update", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Task added successfully.")
	} else {
		fmt.Printf("Failed to add task. Status code: %d\n", resp.StatusCode)
	}
}
