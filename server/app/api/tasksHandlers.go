package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

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
