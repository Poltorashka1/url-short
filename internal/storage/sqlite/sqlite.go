package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
	"os"
	"url-short/internal/config"
	"url-short/internal/storage"
)

type SqliteConnector struct{}

// Connect connects to the sqlite database.
func (s *SqliteConnector) Connect(cfg *config.Config, log *slog.Logger) *storage.Storage {
	db, err := sql.Open("sqlite3", cfg.DatabaseConfig.Config["storagePath"])
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Info("Connected to SQLite database")

	err = createFirstTable(db)
	if err != nil {
		log.Error(err.Error())
	}

	return &storage.Storage{Db: db}
}

// createFirstTable creates first table in the database.
// Its test function is used for testing.
func createFirstTable(db *sql.DB) error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS url(id INTEGER PRIMARY KEY, url TEXT NOT NULL, alias TEXT NOT NULL UNIQUE)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
