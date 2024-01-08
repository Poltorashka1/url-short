package main

import (
	"log/slog"
	"sync"
	"url-short/internal/config"
	"url-short/internal/handlers/Url-handlers"
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
	cfg := config.NewConfig("config/configV2.yaml", log)

	// connect to db (type depending on the config)
	db := SQLConnect(cfg, log)

	// init server
	ser := server.NewServer(cfg)

	// start server
	wg.Add(1)
	go ser.Start(log, &wg)
	// add routes
	initAllRoute(ser, log, db)

	wg.Wait()
}

// SQLConnect connects to the database, the database type depend on the config type.
func SQLConnect(cfg *config.Config, log *slog.Logger) storage.Storage {
	var sqlConnector storage.Storage

	switch cfg.Type {
	case "sqlite":
		sqlConnector = &sqlite.SqliteDatabase{}
	case "postgres":
		sqlConnector = &postgres.PostgresDatabase{}
	}

	db := sqlConnector.Connect(cfg, log)
	return db
}

// initAllRoute creates a new route in the server's mux.
func initAllRoute(ser *server.Server, log *slog.Logger, db storage.Storage) {
	ser.AddRoute("/", test.GetTestResult(log))
	ser.AddRoute("/all", Url_handlers.All(db, log))
	ser.AddRoute("/getUrl/", Url_handlers.GetUrlFromAlias(db, log))
	ser.AddRoute("/saveUrl/", Url_handlers.AddAliasForUrl(db, log))
}
