package config

type Config struct {
	StoragePath string
	HTTPAdress  string
}

func MustLoad() *Config {
	var cfg Config

	cfg.HTTPAdress = "localhost:8080"

	// Строим путь к файлу базы данных
	cfg.StoragePath = "../../db/urls.db"

	return &cfg
}
