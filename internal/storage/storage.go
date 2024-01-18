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
	// GetAlias return alias by url
	GetAlias(url string) (AllAliasList, error)
	// GetDb returns the database db field.
	GetDb() *sql.DB
	// SaveUrl saves url to database
	SaveUrl(urlToSave string, alias string) error
	// DeleteUrl deletes url by alias from database
	DeleteUrl(alias string) error
}

// Path and AllAliasList are structs for json response for Url-handlers.GetAlias() and  Url-handlers.GetUrl()
type Path struct {
	Id    int    `json:"id"`
	Url   string `json:"url"`
	Alias string `json:"alias"`
}

type AllAliasList struct {
	Url   string `json:"url"`
	Alias []Path `json:"allAlias"`
	Code  int    `json:"code"`
}

func NewPath(id int, url, alias string) *Path {
	return &Path{
		Id:    id,
		Url:   url,
		Alias: alias,
	}
}

type StorageError struct {
	Path    string
	Message string
}

func NewStorageError(path, message string) *StorageError {
	return &StorageError{
		Path:    path,
		Message: message,
	}
}

func (e *StorageError) Error() string {
	return e.Message
}
