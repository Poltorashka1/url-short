package storage

import (
	"database/sql"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"log/slog"
	"url-short/internal/config"
)

type Storage struct {
	Db *sql.DB
}

// Connector connects to the database.
type Connector interface {
	Connect(cfg *config.Config, log *slog.Logger) *Storage
}

func (s *Storage) SaveUrl(urlToSave string, alias string) error {
	const op = "storage.sqlite.SaveUrl"
	stmt, err := s.Db.Prepare("INSERT INTO url(url, alias) VALUES(?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(urlToSave, alias)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return fmt.Errorf("%s: %s", op, "Url already exists")
		}
		return err
	}
	return nil
}
