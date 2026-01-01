package internal

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	repo TaskRepository
}

func NewUserHandler(repo TaskRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
