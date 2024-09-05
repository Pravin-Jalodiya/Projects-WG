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

	// Get the username from the URL path variable
	username := r.URL.Path[len("/api/todo/update/"):]
	if username == "" {
		logger.Logger.Warnw("Username not provided",
			"time", time.Now())
		errr := errs.NewInvalidParameterError()
		errr.ToJSON(w)
		return
	}

	var request struct {
		Task    string `json:"task"`
		LastDay int    `json:"last_day"`
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
		if val.Username == username {
			userExist = true
			readers.UserStore[i].GeneralTodo = append(readers.UserStore[i].GeneralTodo, todo)
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

	todo.Deadline = ConvertToHHMMSS(todo.Deadline)

	logger.Logger.Infow("Task added successfully",
		"username", username,
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

// ConvertToHHMMSS returns it in dd/mm/yyyy hh:mm:ss format.
func ConvertToHHMMSS(t time.Time) time.Time {
	// Create a new time object with the same values but truncated to seconds precision
	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		0, // Nanoseconds are set to 0
		t.Location(),
	)
}
