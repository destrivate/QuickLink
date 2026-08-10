package storage

import (
	"LinkShortener/internal/models"
	"database/sql"
	"errors"
	"time"

	_ "modernc.org/sqlite"
)

type SQLStorage struct {
	dbPath string
	db     *sql.DB
}

func NewSQLStorage(dbPath string) *SQLStorage {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(5 * time.Minute)
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil
	}

	s := &SQLStorage{
		dbPath: dbPath,
		db:     db,
	}

	if _, err := db.Exec("PRAGMA journal_mode=WAL;"); err != nil {
		db.Close()
		return nil
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
	query := `INSERT INTO links (path, original_url, password, redirected) VALUES (?, ?, ?, ?);`

	_, err := s.db.Exec(query, link.Path, link.OriginalPath, link.Password, link.Redirected)
	if err != nil {
		return false
	}
	return true
}

func (s *SQLStorage) Get(link string) *models.Link {
	query := `SELECT path, original_url, password, redirected FROM links WHERE path = ?;`

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
	query := `DELETE FROM links WHERE path = ? AND password = ?;`

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
	query := `UPDATE links SET redirected = redirected + 1 WHERE path = ?;`

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
