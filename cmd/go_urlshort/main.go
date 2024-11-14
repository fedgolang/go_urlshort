package main

import (
	"fmt"

	"github.com/fedgolang/go_urlshort/internal/config"
	"github.com/fedgolang/go_urlshort/internal/handlers"
	"github.com/fedgolang/go_urlshort/internal/server"
	"github.com/fedgolang/go_urlshort/internal/storage"
	"github.com/go-chi/chi"
	_ "modernc.org/sqlite"
)

func main() {
	cfg := config.MustLoad()
	r := chi.NewRouter()

	// Открываем коннект к БД
	storage, db := storage.NewStorage(cfg.StoragePath)
	defer db.Close()

	// Хендлер на добавление урла и его сокращения
	r.Post("/", handlers.PostURL(storage))

	// Хендлер на запрос полного урла по его сокращению
	r.Get("/{shortUrl}", handlers.GetURL(storage))

	// Запускаем сервер
	fmt.Println("Сервер запустил")
	server.Sever(r, cfg.HTTPAdress)

}
