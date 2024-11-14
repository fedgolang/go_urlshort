package storage

import (
	"database/sql"
	"errors"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func (s Storage) IsUrlAlreadyExist(fullUrl string) (bool, string, error) {
	stmt, err := s.db.Prepare("SELECT shortUrl FROM urls WHERE fullUrl =?")
	if err != nil {
		return false, "", err
	}
	defer stmt.Close()

	var resURL string
	err = stmt.QueryRow(fullUrl).Scan(&resURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
		} else {
			return false, "", err
		}

	}
	if resURL != "" {
		return true, resURL, nil
	}

	return false, "", nil

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

func (s *Storage) DeleteUrlFull(fullUrl string) error {
	stmt, err := s.db.Prepare("DELETE FROM urls WHERE FullUrl =?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(fullUrl)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteUrlShort(shortUrl string) error {
	stmt, err := s.db.Prepare("DELETE FROM urls WHERE ShortUrl =?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(shortUrl)
	if err != nil {
		return err
	}

	return nil
}
