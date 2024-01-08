package storage

import (
	"database/sql"
	"log/slog"
	"url-short/internal/config"
)

//type Storage struct {
//	Db *sql.DB
//}

// Storage is an interface for storage.
type Storage interface {
	// Connect connects to the database
	Connect(cfg *config.Config, log *slog.Logger) Storage
	// GetUrl returns url by alias
	GetUrl(alias string) (string, error)
	// SaveUrl saves url to database
	SaveUrl(urlToSave string, alias string) error
	// GetDb returns the database db field.
	GetDb() *sql.DB
	// DeleteUrl deletes url by alias from database
	DeleteUrl(alias string) error
}
