package worker

import "LinkShortener/internal/storage"

type DelStruct struct {
	Path string
	Pass string
}

type Workers struct {
	AddRedictInPath chan string
	DelPathInChache chan DelStruct
	db              storage.Storage
}

func NewWorker(db storage.Storage) *Workers {
	return &Workers{AddRedictInPath: make(chan string, 1000), DelPathInChache: make(chan DelStruct, 1000), db: db}
}

func (w *Workers) AddR() {
	for r := range w.AddRedictInPath {
		w.db.AddR(r)
	}
}
func (w *Workers) DelPath() {
	for r := range w.DelPathInChache {
		w.db.Del(r.Path, r.Pass)
	}
}
