package middleware

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"projects/models"
	"projects/utils/errs"
	"projects/utils/logger"
	"projects/utils/password"
	"projects/utils/readers"
	"time"
)

var ActiveUser string

func Auth(username string) {
	ActiveUser = username
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Logger.Errorw("Error decoding request body",
			"error", err,
			"time", time.Now())
		errr := errs.NewInvalidParameterError()
		errr.ToJSON(w)
		return
	}

	// Find the user by username
	user := findUserByUsername(request.Username)
	if user == nil {
		logger.Logger.Warnw("Invalid credentials",
			"username", request.Username,
			"time", time.Now())
		errr := errs.NewUnauthorizedError()
		errr.ToJSON(w)
		return
	}

	// Verify password
	if !password.VerifyPassword(request.Password, user.Password) {
		logger.Logger.Warnw("Invalid credentials",
			"username", request.Username,
			"time", time.Now())
		errr := errs.NewUnauthorizedError()
		errr.ToJSON(w)
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: request.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Second).Unix(), // Token expires in 30 seconds
		},
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		logger.Logger.Errorw("Error generating token",
			"error", err,
			"time", time.Now())
		errr := errs.NewUnexpectedError()
		errr.ToJSON(w)
		return
	}

	// Return the token
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"token": tokenString}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Logger.Errorw("Error encoding token response",
			"error", err,
			"time", time.Now())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Failed to generate token."))
	}
}

func findUserByUsername(username string) *models.UserData {
	for _, u := range readers.UserStore {
		if u.Username == username {
			return &u
		}
	}
	return nil
}
