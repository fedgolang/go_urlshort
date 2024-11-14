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

		urlString := buf.String()

		urlExist, shortUrl, err := s.IsUrlAlreadyExist(urlString)
		if err != nil {
			fmt.Println(err)
			return
		}

		if urlExist {
			_, _ = w.Write([]byte("Сокращение для данного URL уже есть: " + shortUrl))
			return
		}

		shortUrl = shortener.RandomString(shortUrlLen) // Вызовем функцию RandomString для получения короткого урла

		err = s.AddUrls(urlString, shortUrl)
		if err != nil {
			fmt.Println(err)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("Сокращение успешно добавлено"))
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
