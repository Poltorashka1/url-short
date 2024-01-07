package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"url-short/internal/config"
	"url-short/internal/storage"
)

type PostgresDatabase struct {
	Db *sql.DB
}

// Connect connects to the postgres database.
func (p *PostgresDatabase) Connect(cfg *config.Config, log *slog.Logger) storage.Storage {
	dbConfig := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DatabaseConfig.Config["host"],
		cfg.DatabaseConfig.Config["port"],
		cfg.DatabaseConfig.Config["user"],
		cfg.DatabaseConfig.Config["password"],
		cfg.DatabaseConfig.Config["dbname"],
	)

	// dont output connection error
	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	// check connection, and return error if failed
	err = db.Ping()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Info("Connected to Postgres database")
	return &PostgresDatabase{Db: db}
}

func (p *PostgresDatabase) GetDb() *sql.DB {
	return p.Db
}
