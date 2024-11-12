package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/fedgolang/go_urlshort/internal/lib/shortener"
	"github.com/fedgolang/go_urlshort/internal/storage"
	"github.com/go-chi/chi"
)

const (
	shortUrlLen = 7
)

func PostURL(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer

		_, err := buf.ReadFrom(r.Body) // Читаем данные из тела и запишем в буфер
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) // Возвращаем 400 при наличии ошибки
			return
		}

		shortUrl := shortener.RandomString(shortUrlLen) // Вызовем функцию RandomString для получения короткого урла

		err = s.AddUrls(buf.String(), shortUrl)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
	}
}

func GetURL(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUlr := chi.URLParam(r, "shortUrl")

		fullUrl, err := s.GetUrl(shortUlr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) // Возвращаем 400 при наличии ошибки
			return
		}

		w.Header().Set("Content-type", "text/plain")
		w.Header().Set("Location", fullUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)

	}
}
