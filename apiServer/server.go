package apiServer

import (
	"log/slog"
	"net/http"
)

type ApiServer struct {
	serveMux   *http.ServeMux
	listenPort string
}

func NewServer(listenPort string) *ApiServer {
	serveMux := http.NewServeMux()

	return &ApiServer{
		serveMux:   serveMux,
		listenPort: listenPort,
	}
}

func (o *ApiServer) Start() error {
	slog.Info("Starting server on", o.listenPort)
	if err := http.ListenAndServe(o.listenPort, o.serveMux); err != nil {
		slog.Error("Server error:", err)
		return err
	}
	return nil
}
