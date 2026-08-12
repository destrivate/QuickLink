package handlers

import (
	"net/http"
)

func (h *Handler) Redict(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("value")
	if path == "" {
		http.Error(w, "Link code not specified", http.StatusBadRequest)
		return
	}
	original := h.chache.Get(path)
	if original == nil {
		original = h.db.Get(path)
		if original == nil {
			h.SendError(w, "NotFound", http.StatusNotFound)
			return
		}
		h.chache.Add(*original)
	}
	h.chache.AddR(path)

	h.worker.AddRedictInPath <- path
	http.Redirect(w, r, original.OriginalPath, http.StatusSeeOther)
}
