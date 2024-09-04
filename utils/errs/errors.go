package errs

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func NewNotFoundError() *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: "user not found",
	}
}

func NewUnexpectedError() *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}

func NewInvalidParameterError() *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: "invalid parameters",
	}
}

func NewInvalidParameterValueError() *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: "invalid parameter value",
	}
}

func NewInvalidRequestMethodError() *AppError {
	return &AppError{
		Code:    http.StatusMethodNotAllowed,
		Message: "invalid request method",
	}
}

func (e *AppError) ToJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)
	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		return
	}
}
