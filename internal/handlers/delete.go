package handlers

import (
	"LinkShortener/internal/models"
	"LinkShortener/internal/worker"
	"encoding/json"
	"net/http"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method!", http.StatusMethodNotAllowed)
		return
	}
	var data models.RequestDelete
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		h.SendError(w, "Invalid JSON.", http.StatusBadRequest)
		return
	}
	if data.Path == "" || data.Password == "" {
		h.SendError(w, "All fields are required.", http.StatusBadRequest)
		return
	}

	if !h.chache.Del(data.Path, data.Password) {
		h.SendError(w, "Invalid referral code or password.", http.StatusBadRequest)
		return
	}
	h.worker.DelPathInChache <- worker.DelStruct{Path: data.Path, Pass: data.Password}

	h.SendJSON(w, http.StatusOK, map[string]string{"success": "Success."})
}
