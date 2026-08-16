package main

import (
	"LinkShortener/internal/handlers"
	"LinkShortener/internal/storage"
	"LinkShortener/internal/worker"
	"fmt"
	"net/http"
)

const (
	connStr string = ""
)

func main() {

	db, err := storage.NewPostgresStorage(connStr)
	if err != nil {
		fmt.Println(err)
	}
	chacheDb := storage.NewMemoryStorage()
	worker := worker.NewWorker(db)
	go worker.AddAnalitic()
	go worker.DelPath()
	handler := handlers.New(db, worker, chacheDb)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{value}", handler.Redict)
	mux.HandleFunc("GET /qr/{value}", handler.Qr)
	mux.HandleFunc("POST /create", handler.CreatePath)
	mux.HandleFunc("POST /info", handler.Info)
	mux.HandleFunc("POST /delete", handler.Delete)
	mux.HandleFunc("GET /{$}", handler.Home)
	mux.HandleFunc("GET /404", handler.Error)
	http.ListenAndServe(":8080", mux)
}
