package apiServer

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiError struct {
	Message string `json:"message"`
	Code    int    `json:"-"`
}

// NewApiError creates a new ApiError instance
func NewApiError(message string, code int) *ApiError {
	return &ApiError{Message: message, Code: code}
}

// Error implements the error interface by returning the error message
func (e *ApiError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func WriteJson(w http.ResponseWriter, data interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
