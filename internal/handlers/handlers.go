package handlers

import (
	"LinkShortener/internal/storage"
	"encoding/json"
	"net/http"
)

type Handler struct {
	db storage.Storage
}

func New(db storage.Storage) *Handler {
	return &Handler{db: db}
}

func (h *Handler) SendError(w http.ResponseWriter, message string, statusCode int) {
	h.SendJSON(w, statusCode, map[string]string{
		"status": "error",
		"error":  message,
	})
}

func (h *Handler) SendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
