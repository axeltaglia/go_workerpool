package main

import (
	"fmt"
	"workerpool/apiServer"
	"workerpool/taskManager"
)

func main() {
	fmt.Println("ApiServer started in address :8082...")

	// Initialize TaskManager with a buffer size of 10
	manager := taskManager.NewTaskManager(10)

	// Start the TaskManager worker to process tasks
	manager.Start()

	// Create the server and pass the TaskManager
	server := apiServer.NewServer(":8082", manager)

	// Register task-related endpoints
	server.HandleEndpoints("/api/tasks")

	// Start the API server
	if err := server.Start(); err != nil {
		fmt.Printf("server couldn't start in address %s", ":8082")
	}
}
