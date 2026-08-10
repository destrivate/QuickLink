package handlers

import (
	"LinkShortener/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method!", http.StatusMethodNotAllowed)
		return
	}

	var data models.RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		h.SendError(w, "Invalid JSON.", http.StatusBadRequest)
		return
	}
	if data.Path == "" {
		h.SendError(w, "All fields are required.", http.StatusBadRequest)
		return
	}
	red := h.db.Get(data.Path)
	if red == nil {
		h.SendError(w, "NotFound", http.StatusNotFound)
		return
	}

	h.SendJSON(w, http.StatusOK, map[string]string{"redirected": strconv.Itoa(red.Redirected)})
}
