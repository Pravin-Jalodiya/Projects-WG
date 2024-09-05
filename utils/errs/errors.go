package errs

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Code    int    `json:"status_code,omitempty"`
	Message string `json:"message"`
}

// NewNotFoundError creates an error with a custom message
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

// NewUnexpectedError creates an error with a custom message
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

// NewInvalidParameterError creates an error with a custom message
func NewInvalidParameterError(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

// NewInvalidParameterValueError creates an error with a custom message
func NewInvalidParameterValueError(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

// NewInvalidRequestMethodError creates an error with a custom message
func NewInvalidRequestMethodError(message string) *AppError {
	return &AppError{
		Code:    http.StatusMethodNotAllowed,
		Message: message,
	}
}

// NewUnauthorizedError creates an error with a custom message
func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
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
