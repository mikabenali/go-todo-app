package api

import (
	"encoding/json"
	"main/types"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := s.Storage.GetAllTasks()
	if err != nil {
		http.Error(w, "Tasks not found", http.StatusNotFound)
	}

	w.Header().Add("content-type", "aplication/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
}

func (s *Server) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	task, err := s.Storage.GetTaskById(objectId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("content-type", "aplication/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
}

func (s *Server) CreateTask(w http.ResponseWriter, r *http.Request) {
	newTask := types.TaskRequest{}
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := s.Storage.CreateTask(&newTask)
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

func (s *Server) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedTask := types.TaskRequest{}
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.Storage.UpdateTask(updatedTask, objectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func (s *Server) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := s.Storage.DeleteTask(objectId); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
