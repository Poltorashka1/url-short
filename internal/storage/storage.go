package storage

import (
	"database/sql"
	"log/slog"
	"url-short/internal/config"
)

//type Storage struct {
//	Db *sql.DB
//}

// ErrorResponse is an error response.
type ErrorResponse struct {
	Error `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

// Storage is an interface for storage.
type Storage interface {
	Connect(cfg *config.Config, log *slog.Logger) Storage
	GetUrl(alias string) (string, error)
	SaveUrl(urlToSave string, alias string) error
	GetDb() *sql.DB
}
