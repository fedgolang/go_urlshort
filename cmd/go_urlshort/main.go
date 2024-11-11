package main

import (
	"bytes"
	"net/http"

	"github.com/fedgolang/go_urlshort/internal/lib/shortener"
	"github.com/fedgolang/go_urlshort/internal/server"
	"github.com/go-chi/chi"
)

const (
	shortUrlLen = 7
	storagePath = "./db/storage.db"
)

func postURL(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body) // Читаем данные из тела и запишем в буфер
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Возвращаем 400 при наличии ошибки
		return
	}

	shortUrl := shortener.RandomString(shortUrlLen) // Вызовем функцию RandomString для получения короткого урла

}

func main() {
	// Запускаем сервер
	r := chi.NewRouter()
	server.Sever(r)

	// Хендлер на добавление урла и его сокращения
	r.Post("/", postURL)

}
