package storage

import (
	"log/slog"
	"url-short/internal/config"
)

//type Storage struct {
//	Db *sql.DB
//}

// Storage is an interface for storage.
type Storage interface {
	Connect(cfg *config.Config, log *slog.Logger) Storage
	GetUrl(alias string) (string, error)
	SaveUrl(urlToSave string, alias string) error
}
