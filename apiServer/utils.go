package apiServer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

// ApiError structure now includes a stack trace
type ApiError struct {
	Message    string `json:"message"`
	Code       int    `json:"code"`
	StackTrace string `json:"stack_trace,omitempty"`
}

// NewApiError function to capture stack trace
func NewApiError(message string, code int) *ApiError {
	stackBuf := make([]byte, 1024)
	stackSize := runtime.Stack(stackBuf, false)
	return &ApiError{
		Message:    message,
		Code:       code,
		StackTrace: string(stackBuf[:stackSize]),
	}
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("Error: %s, Code: %d, StackTrace: %s", e.Message, e.Code, e.StackTrace)
}

func WriteJson(w http.ResponseWriter, data interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
