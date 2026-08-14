package handlers

import (
	"net/http"
	"os"
)

func (h *Handler) Error(w http.ResponseWriter, r *http.Request) {
	html, err := os.ReadFile("./templates/404.html")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(html)
}
