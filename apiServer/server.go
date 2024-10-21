package apiServer

import (
	"log"
	"net/http"
	"workerpool/taskManager"
)

type ApiServer struct {
	serveMux    *http.ServeMux
	listenPort  string
	taskManager *taskManager.TaskManager // Reference to the TaskManager
}

// NewServer initializes a new ApiServer with a TaskManager
func NewServer(listenPort string, taskManager *taskManager.TaskManager) *ApiServer {
	serveMux := http.NewServeMux()

	return &ApiServer{
		serveMux:    serveMux,
		listenPort:  listenPort,
		taskManager: taskManager,
	}
}

// Start the server
func (o *ApiServer) Start() error {
	log.Println("Starting server on", o.listenPort)
	if err := http.ListenAndServe(o.listenPort, o.serveMux); err != nil {
		log.Println("Server error:", err)
		return err
	}
	return nil
}
