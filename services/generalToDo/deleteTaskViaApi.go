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
	if r.Method != http.MethodDelete {
		errr := errs.NewInvalidRequestMethodError()
		logger.Logger.Errorw("Invalid request method",
			"method", r.Method,
			"url", r.URL.Path,
			"time", time.Now())
		errr.ToJSON(w)
		return
	}

	// Get the username from the URL path variable
	username := r.URL.Path[len("/api/todo/update/"):]
	if username == "" {
		logger.Logger.Warnw("Username not provided",
			"time", time.Now())
		errr := errs.NewInvalidParameterError()
		errr.ToJSON(w)
		return
	}

	// Get the task_no from the query parameter
	taskNoStr := r.URL.Query().Get("task-no")
	taskNo, err := strconv.Atoi(taskNoStr)
	if err != nil || taskNo <= 0 {
		logger.Logger.Warnw("Invalid task number",
			"task_no", taskNoStr,
			"time", time.Now())
		errr := errs.NewInvalidParameterValueError()
		errr.ToJSON(w)
		return
	}

	var userExist bool = false
	for i, user := range readers.UserStore {
		if user.Username == username {
			userExist = true
			if len(user.GeneralTodo) == 0 || taskNo > len(user.GeneralTodo) {
				logger.Logger.Warnw("Invalid task number",
					"task_no", taskNo,
					"total_tasks", len(user.GeneralTodo),
					"time", time.Now())
				errr := errs.NewInvalidParameterValueError()
				errr.ToJSON(w)
				return
			}
			readers.UserStore[i].GeneralTodo = append(readers.UserStore[i].GeneralTodo[:taskNo-1], readers.UserStore[i].GeneralTodo[taskNo:]...)
			break
		}
	}

	if !userExist {
		logger.Logger.Warnw("User not found",
			"username", username,
			"time", time.Now())
		errr := errs.NewNotFoundError()
		errr.ToJSON(w)
		return
	}

	_, err = writers.FWriterToDo(config.USER_FILE, readers.UserStore)
	if err != nil {
		logger.Logger.Errorw("Error writing to file",
			"error", err,
			"time", time.Now())
		errr := errs.NewUnexpectedError()
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
		logger.Logger.Errorw("Error writing response",
			"error", err,
			"time", time.Now())
		return
	}
}
