package utils

import "net/http"

func GetHost(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}

	host := r.Host

	fullURL := scheme + "://" + host + "/"
	return fullURL
}
