package main

import (
	"LinkShortener/internal/handlers"
	"LinkShortener/internal/storage"
	"LinkShortener/internal/worker"
	"net/http"
)

const (
	dbPath string = "./db/db.db"
)

func main() {

	db := storage.NewSQLStorage(dbPath)
	chacheDb := storage.NewMemoryStorage()
	worker := worker.NewWorker(db)
	go worker.AddR()
	go worker.DelPath()
	handler := handlers.New(db, worker, chacheDb)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{value}", handler.Redict)
	mux.HandleFunc("GET /qr/{value}", handler.Qr)
	mux.HandleFunc("POST /create", handler.CreatePath)
	mux.HandleFunc("POST /info", handler.Info)
	mux.HandleFunc("POST /delete", handler.Delete)
	mux.HandleFunc("GET /{$}", handlers.Home)
	http.ListenAndServe(":8080", mux)
}
