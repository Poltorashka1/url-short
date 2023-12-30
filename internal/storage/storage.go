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
	const op = "storage.SaveUrl"
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

func (s *Storage) GetUrl(alias string) (string, error) {
	const op = "storage.storage.GetUrl"
	rows, err := s.Db.Query("SELECT url FROM url WHERE alias = $1", alias)
	//stmt, err := s.Db.Prepare("SELECT url FROM url WHERE alias = $1")
	//if err != nil {
	//	return "", fmt.Errorf("%s: %s", op, err.Error())
	//}
	//
	//rows, err := stmt.Query(fmt.Sprintf("'%s'", alias))
	//if err != nil {
	//	return "", fmt.Errorf("%s: %s", op, err.Error())
	//}

	if rows.Next() {
		var url string
		if err = rows.Scan(&url); err != nil {
			return "", fmt.Errorf("%s: %s", op, err.Error())
		}
		return url, nil
	} else {
		return "", fmt.Errorf("%s: %s", op, "Url not found")
	}
}
