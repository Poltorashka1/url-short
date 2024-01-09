package main

import (
	"log/slog"
	"sync"
	"url-short/internal/config"
	"url-short/internal/handlers/Url-handlers"
	"url-short/internal/logger"
	"url-short/internal/server"
	"url-short/internal/storage"
	"url-short/internal/storage/postgres"
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
	//case "sqlite":
	//	sqlConnector = &sqlite.SqliteDatabase{}
	case "postgres":
		sqlConnector = &postgres.PostgresDatabase{}
	}

	db := sqlConnector.Connect(cfg, log)
	return db
}

// initAllRoute creates a new route in the server's mux.
func initAllRoute(ser *server.Server, log *slog.Logger, db storage.Storage) {
	ser.AddRoute("/all", Url_handlers.GetAllUrlHandler(db, log))
	ser.AddRoute("/getUrl/", Url_handlers.GetUrlFromAliasHandler(db, log))
	ser.AddRoute("/getAlias/", Url_handlers.GetAliasFromUrlHandler(db, log))
	ser.AddRoute("/saveUrl/", Url_handlers.AddAliasForUrlHandler(db, log))
	ser.AddRoute("/deleteUrl/", Url_handlers.DeleteUrlHandler(db, log))
}
