package generalToDo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projects/config"
	"projects/utils/readers"
	"projects/utils/writers"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Username string `json:"username"`
		TaskNo   int    `json:"task_no"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	var userExist bool = false

	for i, user := range readers.UserStore {
		if user.Username == request.Username {
			userExist = true
			if len(user.GeneralTodo) == 0 || request.TaskNo <= 0 || request.TaskNo > len(user.GeneralTodo) {
				http.Error(w, "Invalid task number", http.StatusBadRequest)
				return
			}
			readers.UserStore[i].GeneralTodo = append(readers.UserStore[i].GeneralTodo[:request.TaskNo-1], readers.UserStore[i].GeneralTodo[request.TaskNo:]...)
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
	_, err = fmt.Fprintln(w, "Task deleted successfully.")
	if err != nil {
		return
	}
}
