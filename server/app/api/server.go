package api

import (
	"encoding/json"
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

func (s *Server) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := s.Storage.GetAllTasks()
	if err != nil {
		http.Error(w, "Tasks not found", http.StatusNotFound)
	}

	w.Header().Add("content-type", "aplication/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
}

func (s *Server) getTaskById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	task, err := s.Storage.GetTaskById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("content-type", "aplication/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {}
func (s *Server) updateTask(w http.ResponseWriter, r *http.Request) {}
func (s *Server) deleteTask(w http.ResponseWriter, r *http.Request) {}
