package main

import (
	"LinkShortener/internal/handlers"
	"LinkShortener/internal/storage"
	"net/http"
)

const (
	dbPath      string = "./db/db.db"
	storageType string = "mem" //sql || mem
)

func main() {
	var db storage.Storage = nil
	if storageType == "sql" {
		db = storage.NewSQLStorage(dbPath)
	}
	if storageType == "mem" {
		db = storage.NewMemoryStorage()
	}

	handler := handlers.New(db)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{value}", handler.Redict)
	mux.HandleFunc("GET /qr/{value}", handler.Qr)
	mux.HandleFunc("POST /create", handler.Create_Path)
	mux.HandleFunc("POST /info", handler.Info)
	mux.HandleFunc("POST /delete", handler.Delete)
	mux.HandleFunc("GET /{$}", handlers.Home)
	http.ListenAndServe(":8080", mux)
}
