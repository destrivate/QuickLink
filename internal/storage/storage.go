package storage

import "LinkShortener/internal/models"

type Storage interface {
	Add(link models.Link) bool
	Get(link string) *models.Link
	AddR(link string) bool
	Del(link string, pass string) bool
}
