package handlers

import (
	"LinkShortener/internal/utils"
	"net/http"

	"github.com/skip2/go-qrcode"
)

func (h *Handler) Qr(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("value")
	if path == "" {
		http.Error(w, "Link code not specified", http.StatusBadRequest)
		return
	}
	original := h.chache.Get(path)
	if original == nil {
		original = h.db.Get(path)
		if original == nil {
			http.NotFound(w, r)
			return
		}
	}
	png, _ := qrcode.Encode(utils.GetHost(r)+path, qrcode.Medium, 256)
	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}
