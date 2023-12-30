package main

import (
	"log/slog"
	"sync"
	"url-short/internal/config"
	"url-short/internal/handlers/test"
	"url-short/internal/logger"
	"url-short/internal/server"
	"url-short/internal/storage"
	"url-short/internal/storage/postgres"
	"url-short/internal/storage/sqlite"
)

var wg sync.WaitGroup

func main() {
	// init logger
	log := logger.SetupLogger()
	// init config
	cfg := config.NewConfig("config/configV1.yaml", log)

	// connect to db
	db := SQLConnect(cfg, log)

	_ = db
	// init and start server
	ser := server.NewServer(cfg)
	wg.Add(1)
	go ser.Start(log, &wg)

	// add routes
	initAllRoute(ser, log)

	wg.Wait()
}

// SQLConnect connects to the database, the database type depend on the config type.
func SQLConnect(cfg *config.Config, log *slog.Logger) *storage.Storage {
	var sqlConnector storage.Connector
	switch cfg.Type {
	case "sqlite":
		sqlConnector = &sqlite.SqliteConnector{}
	case "postgres":
		sqlConnector = &postgres.PostgresConnector{}
	}

	db := sqlConnector.Connect(cfg, log)
	return db
}

// initAllRoute creates a new route in the server's mux.
func initAllRoute(ser *server.Server, log *slog.Logger) {
	ser.AddRoute("/", test.GetTestResult(log))
}
