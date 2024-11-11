package storage

import (
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

func (s *Storage) addUrls(fullUrl, shortUlr string) error {
	stmt, err := s.db.Prepare("INSERT INTO urls(FullUrl, ShortUrl) VALUES(?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(fullUrl, shortUlr)
	if err != nil {
		return err
	}

	return nil

}
