package internal

import (
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) RegisterRoutes() {
	userHandler := NewUserHandler()

	s.mux.HandleFunc("/task/get", userHandler.GetTasks)
}

func (s *Server) Start(addr string) error {
	log.Println("Server running on", addr)
	return http.ListenAndServe(addr, s.mux)
}

func (s *Server) Router() *http.ServeMux {
	return s.mux
}
