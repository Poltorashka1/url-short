package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
	"os"
	"url-short/internal/config"
	"url-short/internal/storage"
)

type SqliteDatabase struct {
	Db *sql.DB
}

// Connect connects to the sqlite database.
func (s *SqliteDatabase) Connect(cfg *config.Config, log *slog.Logger) storage.Storage {
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

	return &SqliteDatabase{Db: db}
}

// createFirstTable creates first table in the database.
// Its debug function is used for debugging.
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

func (s *SqliteDatabase) GetDb() *sql.DB {
	return s.Db
}
