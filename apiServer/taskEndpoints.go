package apiServer

import (
	"log"
	"net/http"
	"strconv"
	"workerpool/taskManager"
)

// HandleEndpoints registers the endpoints to the server
func (o *ApiServer) HandleEndpoints(prefix string) {
	o.serveMux.HandleFunc(prefix+"/status", o.interceptError(o.status))
	o.serveMux.HandleFunc(prefix+"/enqueueTask", o.interceptError(o.enqueueTask))
}

// status endpoint
func (o *ApiServer) status(w http.ResponseWriter, r *http.Request) error {
	response := EnqueueTaskResponse{Status: "OK"}
	return WriteJson(w, response, http.StatusOK)
}

// enqueueTask endpoint adds a task to the TaskManager
func (o *ApiServer) enqueueTask(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return &ApiError{
			Message: "Invalid request method",
			Code:    http.StatusMethodNotAllowed,
		}
	}

	// Get taskID and workload duration from request parameters
	taskID := r.URL.Query().Get("taskID")
	if taskID == "" {
		return NewApiError("Missing taskID", http.StatusBadRequest)
	}

	workloadSeconds := r.URL.Query().Get("workload")
	if workloadSeconds == "" {
		return NewApiError("Missing workload parameter", http.StatusBadRequest)
	}

	// Convert workload to integer
	workloadDuration, err := strconv.Atoi(workloadSeconds)
	if err != nil || workloadDuration < 0 {
		return NewApiError("Invalid workload value", http.StatusBadRequest)
	}

	// Create a task with the specified workload duration
	task := &taskManager.ConcreteTask{
		ID:       taskID,
		Workload: workloadDuration, // Sleep for the given workload duration
	}

	// Add task to the TaskManager
	o.taskManager.AddTask(task)

	// Return success
	response := EnqueueTaskResponse{Status: "Task added to the queue"}
	return WriteJson(w, response, http.StatusOK)
}

// interceptor for error handling
func (o *ApiServer) interceptError(f func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			// Try to cast the error to an ApiError
			apiErr, ok := err.(*ApiError)

			if !ok {
				log.Printf("ApiError: %v\nInvalid type. ApiError structure is expected.\n%s", apiErr, apiErr.StackTrace)
				// Return a generic error response
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			// Log the API error and the stack trace
			log.Printf("ApiError: %v\nStack trace:\n%s", apiErr.Message, apiErr.StackTrace)

			// Send the ApiError message as an HTTP response
			http.Error(w, apiErr.Message, apiErr.Code)
		}
	}
}
