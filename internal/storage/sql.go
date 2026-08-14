package storage

import (
	"LinkShortener/internal/models"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type SQLStorage struct {
	db *sql.DB
}

func NewSQLStorage(connStr string) *SQLStorage {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil
	}

	s := &SQLStorage{
		db: db,
	}

	if err := s.migrate(); err != nil {
		db.Close()
		return nil
	}

	return s
}

func (s *SQLStorage) migrate() error {

	query := `
	CREATE TABLE IF NOT EXISTS links (
		path TEXT PRIMARY KEY,
		original_url TEXT NOT NULL,
		password TEXT,
		redirected INTEGER DEFAULT 0
	);`
	_, err := s.db.Exec(query)
	return err
}

func (s *SQLStorage) Add(link models.Link) bool {
	query := `INSERT INTO links (path, original_url, password, redirected) VALUES ($1, $2, $3, $4);`

	_, err := s.db.Exec(query, link.Path, link.OriginalPath, link.Password, link.Redirected)
	if err != nil {
		return false
	}
	return true
}

func (s *SQLStorage) Get(link string) *models.Link {
	query := `SELECT path, original_url, password, redirected FROM links WHERE path = $1;`

	var l models.Link
	err := s.db.QueryRow(query, link).Scan(&l.Path, &l.OriginalPath, &l.Password, &l.Redirected)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return nil
	}
	return &l
}

func (s *SQLStorage) Del(link string, pass string) bool {
	query := `DELETE FROM links WHERE path = $1 AND password = $2;`

	res, err := s.db.Exec(query, link, pass)
	if err != nil {
		return false
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return false
	}
	return true
}

func (s *SQLStorage) AddR(link string) bool {
	query := `UPDATE links SET redirected = redirected + 1 WHERE path = $1;`

	res, err := s.db.Exec(query, link)
	if err != nil {
		return false
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return false
	}
	return true
}
