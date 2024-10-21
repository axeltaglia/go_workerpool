package main

import (
	"fmt"
	"log/slog"
	"workerpool/apiServer"
)

func main() {
	fmt.Println("ApiServer started in address :8082...")
	server := apiServer.NewServer(":8082")

	server.HandleEndpoints("/api/tasks")

	if err := server.Start(); err != nil {
		slog.Error("server couldn't start in address %s", ":8082")
	}
}
