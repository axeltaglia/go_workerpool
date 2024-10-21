package apiServer

import (
	"log"
	"net/http"
	"runtime"
)

func (o *ApiServer) HandleEndpoints(prefix string) {
	o.serveMux.HandleFunc(prefix+"/status", interceptError(o.status))
}

func (o *ApiServer) status(w http.ResponseWriter, r *http.Request) error {
	response := EnqueueTaskResponse{Status: "OK"}
	return WriteJson(w, response, http.StatusOK)
}

func interceptError(f func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			// Capture the stack trace
			stackBuf := make([]byte, 1024)              // Initial size for the stack trace buffer
			stackSize := runtime.Stack(stackBuf, false) // Get the stack trace of the current goroutine

			// Try to cast the error to an ApiError
			apiErr, ok := err.(*ApiError)
			if !ok {
				// Log the error and the stack trace for debugging purposes
				log.Printf("Unexpected error: %v\nStack trace:\n%s", err, string(stackBuf[:stackSize]))

				// Return a generic error response
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			// Log the API error and the stack trace
			log.Printf("ApiError: %v\nStack trace:\n%s", apiErr, string(stackBuf[:stackSize]))

			// Send the ApiError message as an HTTP response
			http.Error(w, apiErr.Message, apiErr.Code)
		}
	}
}

/*
if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
*/
