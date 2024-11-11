package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func Sever(r *chi.Mux) {
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
	fmt.Println("Сервер запущен, слушаю порт 8080")
}
