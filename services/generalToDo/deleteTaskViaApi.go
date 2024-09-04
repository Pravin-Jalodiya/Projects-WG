package generalToDo

import (
	"encoding/json"
	"net/http"
	"projects/config"
	"projects/utils/errs"
	"projects/utils/logger"
	"projects/utils/readers"
	"projects/utils/writers"
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

	var request struct {
		Username string `json:"username"`
		TaskNo   int    `json:"task_no"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Logger.Errorw("Error parsing request body",
			"error", err,
			"request", request,
			"time", time.Now())
		errr := errs.NewInvalidParameterError()
		errr.ToJSON(w)
		return
	}

	var userExist bool = false

	for i, user := range readers.UserStore {
		if user.Username == request.Username {
			userExist = true
			if len(user.GeneralTodo) == 0 || request.TaskNo <= 0 || request.TaskNo > len(user.GeneralTodo) {
				logger.Logger.Warnw("Invalid task number",
					"task_no", request.TaskNo,
					"total_tasks", len(user.GeneralTodo),
					"time", time.Now())
				errr := errs.NewInvalidParameterValueError()
				errr.ToJSON(w)
				return
			}
			readers.UserStore[i].GeneralTodo = append(readers.UserStore[i].GeneralTodo[:request.TaskNo-1], readers.UserStore[i].GeneralTodo[request.TaskNo:]...)
			break
		}
	}

	if !userExist {
		logger.Logger.Warnw("User not found",
			"username", request.Username,
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
		"username", request.Username,
		"task_no", request.TaskNo,
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
