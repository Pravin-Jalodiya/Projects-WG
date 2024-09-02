package generalToDo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projects/config"
	"projects/models"
	"projects/utils/readers"
	"projects/utils/writers"
	"strings"
	"time"
)

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Username string `json:"username"`
		Task     string `json:"task"`
		LastDay  int    `json:"last_day"`
	}

	var todo models.DoList

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	todo.Task = strings.TrimSpace(request.Task)
	todo.Deadline = time.Now().AddDate(0, 0, request.LastDay)

	var userExist bool = false

	for i, val := range readers.UserStore {
		if val.Username == request.Username {
			userExist = true
			readers.UserStore[i].GeneralTodo = append(readers.UserStore[i].GeneralTodo, todo)
			break
		}
	}

	if !userExist {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}

	_, err = writers.FWriterToDo(config.USER_FILE, readers.UserStore)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintln(w, "Task added successfully.")
	if err != nil {
		return
	}
}
