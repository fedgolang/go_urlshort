package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/fedgolang/go_urlshort/internal/config"
	"github.com/fedgolang/go_urlshort/internal/storage"
)

func TestPostURL(t *testing.T) {
	cfg := config.MustLoad()

	storage, db := storage.NewStorage(cfg.StoragePath)
	defer db.Close()

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
				response:    `{"status":"ok"}`,
				contentType: "application/json",
			},
		},
	}
	for _, tests := range tests {
		t.Run(tests.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(tests.give.body)))

			w := httptest.NewRecorder()

		})
	}
}

func TestGetURL(t *testing.T) {
	type args struct {
		s *storage.Storage
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetURL(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
