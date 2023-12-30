package server

import (
	"log/slog"
	"net/http"
	"os"
	"sync"
	"url-short/internal/config"
)

type Server struct {
	Host string
	Port string
	mux  *http.ServeMux
}

// NewServer creates a new instance of the Server struct and returns a pointer to it.
func NewServer(cfg *config.Config) *Server {
	server := &Server{
		Host: cfg.Server.Host,
		Port: cfg.Server.Port,
		mux:  http.NewServeMux(),
	}

	return server
}

// Start starts the server and listens for HTTP requests. It exits the program if an error occurs.
func (s *Server) Start(log *slog.Logger, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Info("Server started on ", slog.String("host", s.Host), slog.String("port", s.Port))

	err := http.ListenAndServe(s.Host+":"+s.Port, s.mux)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

// AddRoute adds a new path to the server's mux.
func (s *Server) AddRoute(route string, fn http.HandlerFunc) {
	s.mux.HandleFunc(route, fn)
}
