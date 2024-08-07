package generalToDo

import (
	"bufio"
	"fmt"
	"os"
	"projects/config"
	"projects/models"
	"projects/utils/readers"
	"projects/utils/writers"
	"strconv"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func addTask(currentUser string) {
	fmt.Printf("\n%sADD TASK%s\n", config.STR_DECOR, config.STR_DECOR)
	var todo models.DoList
	var deadline string
	var lastDay int
	var err error
	for {
		fmt.Print("Enter task: ")
		todo.Task, err = reader.ReadString('\n')
		fmt.Println()
		if err != nil {
			fmt.Println("Error reading input")
			continue
		}
		todo.Task = strings.TrimSuffix(todo.Task, "\n")
		todo.Task = strings.TrimSpace(todo.Task)
		break
	}

	for {
		fmt.Print("Enter deadline (in days): ")
		deadline, err = reader.ReadString('\n')
		fmt.Println()
		if err != nil {
			fmt.Println("Error reading input. Try again,")
			continue
		}
		deadline = strings.TrimSuffix(deadline, "\n")
		deadline = strings.TrimSpace(deadline)
		lastDay, err = strconv.Atoi(deadline)
		if err != nil {
			fmt.Println("Enter a valid number.")
			continue
		}
		break
	}

	todo.Deadline = time.Now().AddDate(0, 0, lastDay)

	for i, val := range readers.UserStore {
		if val.Username == currentUser {
			readers.UserStore[i].GeneralTodo = append(readers.UserStore[i].GeneralTodo, todo)
			break
		}
	}

	_, err = writers.FWriterToDo(config.USER_FILE, readers.UserStore)
	if err != nil {
		fmt.Println("Error writing to file.")
		return
	}

	fmt.Println("Task added successfully.")
}
