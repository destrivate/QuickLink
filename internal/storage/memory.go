package storage

import (
	"LinkShortener/internal/models"
	"sync"
)

type MemoryStorage struct {
	Data []models.Link
	mu   sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Data: make([]models.Link, 0),
	}
}

func (m *MemoryStorage) Add(link models.Link) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.Data {
		if l.Path == link.Path {
			return false
		}
	}
	m.Data = append(m.Data, link)
	return true
}

func (m *MemoryStorage) Get(link string) *models.Link {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := range m.Data {
		if m.Data[i].Path == link {
			return &m.Data[i]
		}
	}
	return nil
}

func (m *MemoryStorage) Del(link string, pass string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, l := range m.Data {
		if l.Path == link {
			if l.Password == pass {
				m.Data = append(m.Data[:i], m.Data[i+1:]...)
				return true
			}
		}
	}
	return false
}
func (m *MemoryStorage) AddR(link string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := range m.Data {
		if m.Data[i].Path == link {
			m.Data[i].Redirected += 1
			return true
		}
	}
	return false
}
