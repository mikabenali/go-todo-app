package api

import (
	"main/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct{}

type Server struct {
	Config  Config
	Storage storage.Database
}

func New(config Config, storage storage.Database) *Server {
	return &Server{
		Config:  config,
		Storage: storage,
	}
}

func (s *Server) Start() error {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", s.getTasks)
		r.Get("/{id}", s.getTaskById)
		r.Post("/", s.createTask)
		r.Put("/{id}", s.updateTask)
		r.Delete("/{id}", s.deleteTask)
	})

	return http.ListenAndServe(":3000", r)
}
