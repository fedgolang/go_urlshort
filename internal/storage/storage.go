package storage

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(dbPath string) (*Storage, *sql.DB) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		fmt.Println(err)
		return &Storage{}, nil
	}

	return &Storage{db: db}, db
}

func (s *Storage) AddUrls(fullUrl, shortUlr string) error {
	stmt, err := s.db.Prepare("INSERT INTO urls(FullUrl, ShortUrl) VALUES(?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(fullUrl, shortUlr)
	if err != nil {
		return err
	}

	return nil

}

func (s Storage) GetUrl(shortUrl string) (string, error) {
	stmt, err := s.db.Prepare("SELECT FullUrl FROM urls WHERE ShortUrl =?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var resURL string
	err = stmt.QueryRow(shortUrl).Scan(&resURL)
	if err != nil {
		return "", err
	}

	return resURL, nil
}
