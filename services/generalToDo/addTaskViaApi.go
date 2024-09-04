package generalToDo

import (
	"encoding/json"
	"net/http"
	"projects/config"
	"projects/models"
	"projects/utils/errs"
	"projects/utils/logger"
	"projects/utils/readers"
	"projects/utils/writers"
	"strings"
	"time"
)

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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
		Task     string `json:"task"`
		LastDay  int    `json:"last_day"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || strings.TrimSpace(request.Task) == "" || request.LastDay <= 0 {
		logger.Logger.Errorw("Error decoding request body or invalid parameters",
			"error", err,
			"request", request,
			"time", time.Now())
		errr := errs.NewInvalidParameterValueError()
		errr.ToJSON(w)
		return
	}

	var todo models.DoList
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

	logger.Logger.Infow("Task added successfully",
		"username", request.Username,
		"task", todo.Task,
		"deadline", todo.Deadline,
		"time", time.Now())

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Task added successfully."))
	if err != nil {
		logger.Logger.Errorw("Error writing response",
			"error", err,
			"time", time.Now())
		return
	}
}
