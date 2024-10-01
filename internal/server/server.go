package server

import "github.com/go-chi/chi/v5"

type Server struct {
	router     chi.Router
	repository *repository.Repository
}

func NewServer() *Server {
	return &Server{
		router: chi.NewRouter(),
	}
}
