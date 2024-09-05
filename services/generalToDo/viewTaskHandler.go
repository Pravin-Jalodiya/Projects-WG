package generalToDo

import (
	"encoding/json"
	"net/http"
	"projects/models"
	"projects/utils/errs"
	"projects/utils/logger"
	"projects/utils/readers"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func ViewTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Extract username from the URL path variable using gorilla/mux
	vars := mux.Vars(r)
	username := vars["username"]

	if strings.TrimSpace(username) == "" {
		logger.Logger.Errorw("Missing username parameter",
			"username", username,
			"time", time.Now())
		err := errs.NewInvalidParameterValueError()
		err.ToJSON(w)
		return
	}

	// Find the user in the UserStore
	var user *models.UserData

	for _, u := range readers.UserStore {
		if u.Username == username {
			user = &u
			break
		}
	}

	if user == nil {
		logger.Logger.Warnw("User not found",
			"username", username,
			"time", time.Now())
		err := errs.NewNotFoundError()
		err.ToJSON(w)
		return
	}

	// If no tasks are found for the user
	if len(user.GeneralTodo) == 0 {
		logger.Logger.Infow("No todo tasks found for the user",
			"username", username,
			"time", time.Now())
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ToDo list is empty. No pending work."))
		if err != nil {
			logger.Logger.Errorw("Error writing response",
				"error", err,
				"time", time.Now())
		}
		return
	}

	// Check for the 'limit' query parameter
	limitParam := r.URL.Query().Get("limit")
	limit := len(user.GeneralTodo) // Default to showing all tasks
	if limitParam != "" {
		parsedLimit, err := strconv.Atoi(limitParam)
		if err != nil || parsedLimit <= 0 {
			logger.Logger.Warnw("Invalid limit parameter",
				"limit", limitParam,
				"username", username,
				"time", time.Now())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Invalid 'limit' parameter. It must be a positive integer."))
			return
		}
		if parsedLimit < limit {
			limit = parsedLimit
		}
	}

	// Return the limited user's todo list in JSON format
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(user.GeneralTodo[:limit])
	if err != nil {
		logger.Logger.Errorw("Error encoding todo list to JSON",
			"error", err,
			"time", time.Now())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Failed to retrieve ToDo list."))
	}
}
