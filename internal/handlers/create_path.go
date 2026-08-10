package handlers

import (
	"LinkShortener/internal/models"
	"encoding/json"
	"net/http"
)

func (h *Handler) Create_Path(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method!", http.StatusMethodNotAllowed)
		return
	}

	req := models.Link{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.SendError(w, "Invalid JSON.", http.StatusBadRequest)
		return
	}
	req.Redirected = 0

	if req.OriginalPath == "" || req.Path == "" || req.Password == "" {
		h.SendError(w, "All fields are required.", http.StatusBadRequest)
		return
	}

	if !h.db.Add(req) {
		h.SendError(w, "This identifier is busy!", http.StatusBadRequest)
		return
	}

	h.SendJSON(w, http.StatusCreated, map[string]string{"success": "Success."})

}
