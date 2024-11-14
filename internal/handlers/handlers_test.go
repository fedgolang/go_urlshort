package handlers

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fedgolang/go_urlshort/internal/config"
	"github.com/fedgolang/go_urlshort/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "modernc.org/sqlite"
)

func prepare() (*storage.Storage, *sql.DB, error) {
	cfg := config.MustLoad()
	storage, db := storage.NewStorage(cfg.StoragePath)

	// Подготовим тестовые данные
	err := storage.AddUrls("https://github.com/fedgolang/go-rest-api/blob/todo-list/precode.go", "Test")
	if err != nil {
		return nil, db, err
	}

	return storage, db, nil
}

func TestPostURL(t *testing.T) {
	// Подготовим БД и тестовые данные
	storage, db, err := prepare()
	defer db.Close()
	require.NoError(t, err)

	type give struct {
		contentType string
		body        string
	}
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "URL already exist",
			give: give{
				contentType: "text/plain",
				body:        "https://github.com/fedgolang/go-rest-api/blob/todo-list/precode.go",
			},
			want: want{
				code:        200,
				response:    `Сокращение для данного URL уже есть: Test`,
				contentType: "text/plain",
			},
		},
		{
			name: "New url",
			give: give{
				contentType: "text/plain",
				body:        "https://ya.ru/",
			},
			want: want{
				code:        201,
				response:    `Сокращение успешно добавлено`,
				contentType: "text/plain",
			},
		},
	}

	// Запустим тесты
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(test.give.body)))
			w := httptest.NewRecorder()

			postURL := PostURL(storage)

			postURL(w, r)

			assert.Equal(t, w.Code, test.want.code)
			assert.Equal(t, w.Body.String(), test.want.response)

		})
	}

	// Почистим за собой тестовые данные, чтобы не флакать тесты
	for _, test := range tests {
		storage.DeleteUrlFull(test.give.body)
	}

}

func TestGetURL(t *testing.T) {
	// Подготовим БД и тестовые данные
	storage, db, err := prepare()
	defer db.Close()
	require.NoError(t, err)

	type give struct {
		shortUrl string
	}
	type want struct {
		code        int
		contentType string
		Location    string
	}
	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "Success",
			give: give{
				shortUrl: "Test",
			},
			want: want{
				code:        307,
				contentType: "text/plain",
				Location:    "https://github.com/fedgolang/go-rest-api/blob/todo-list/precode.go",
			},
		},
		{
			name: "Negative, url not in db",
			give: give{
				shortUrl: "some_url",
			},
			want: want{
				code:        400,
				contentType: "text/plain; charset=utf-8",
				Location:    "",
			},
		},
	}
	// Запустим тесты
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/"+test.give.shortUrl, nil)
			w := httptest.NewRecorder()

			GetURL := GetURL(storage)

			GetURL(w, r)

			assert.Equal(t, w.Code, test.want.code)
			assert.Equal(t, w.Result().Header.Get("Content-Type"), test.want.contentType)
			if test.want.Location != "" {
				assert.Equal(t, w.Result().Header.Get("Location"), test.want.Location)
			}

		})
	}
	// Почистим за собой тестовые данные, чтобы не флакать тесты
	for _, test := range tests {
		storage.DeleteUrlShort(test.give.shortUrl)
	}
}
