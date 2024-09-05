package generalToDo

import (
	"net/http"
	"projects/config"
	"projects/utils/errs"
	"projects/utils/logger"
	"projects/utils/readers"
	"projects/utils/writers"
	"strconv"
	"time"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Get the username from the URL path variable
	username := r.URL.Path[len("/api/todo/delete/"):]
	if username == "" {
		logger.Logger.Warnw("Username not provided in the request URL",
			"time", time.Now())
		errr := errs.NewInvalidParameterError("Username parameter is missing in the URL")
		errr.ToJSON(w)
		return
	}

	// Get the task_no from the query parameter
	taskNoStr := r.URL.Query().Get("task")
	taskNo, err := strconv.Atoi(taskNoStr)
	if err != nil || taskNo <= 0 {
		logger.Logger.Warnw("Invalid task number",
			"task_no", taskNoStr,
			"error", err,
			"time", time.Now())
		errr := errs.NewInvalidParameterValueError("Task number must be a positive integer")
		errr.ToJSON(w)
		return
	}

	var userExist = false
	for i, user := range readers.UserStore {
		if user.Username == username {
			userExist = true
			if len(user.GeneralTodo) == 0 || taskNo > len(user.GeneralTodo) {
				logger.Logger.Warnw("Task number out of range",
					"task_no", taskNo,
					"total_tasks", len(user.GeneralTodo),
					"time", time.Now())
				errr := errs.NewInvalidParameterValueError("Task number is out of range for the user’s todo list")
				errr.ToJSON(w)
				return
			}
			readers.UserStore[i].GeneralTodo = append(readers.UserStore[i].GeneralTodo[:taskNo-1], readers.UserStore[i].GeneralTodo[taskNo:]...)
			break
		}
	}

	if !userExist {
		logger.Logger.Warnw("User not found in UserStore",
			"username", username,
			"time", time.Now())
		errr := errs.NewNotFoundError("User with the given username does not exist")
		errr.ToJSON(w)
		return
	}

	_, err = writers.FWriterToDo(config.USER_FILE, readers.UserStore)
	if err != nil {
		logger.Logger.Errorw("Error writing updated user data to file",
			"error", err,
			"time", time.Now())
		errr := errs.NewUnexpectedError("Failed to write updated todo list to file")
		errr.ToJSON(w)
		return
	}

	logger.Logger.Infow("Task deleted successfully",
		"username", username,
		"task_no", taskNo,
		"time", time.Now())

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Task deleted successfully."))
	if err != nil {
		logger.Logger.Errorw("Error writing success response",
			"error", err,
			"time", time.Now())
		return
	}
}
