package server

import (
	"log/slog"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	Handler    http.Handler
}

func New(port string, router http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + port,
			Handler: router,
			// TODO configs
		},
	}
}

func (s *Server) Run() error {
	slog.Info("server is started", "ADDR", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
