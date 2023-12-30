package logger

import (
	"log/slog"
	"os"
	"url-short/internal/lib/colorLog"
)

// SetupLogger creates and configures a logger.
func SetupLogger() *slog.Logger {
	// Create a new logger with a text handler that writes to os.Stdout
	opts := colorLog.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := colorLog.NewPrettyHandler(os.Stdout, opts)
	log := slog.New(handler)

	//log := slog.New(
	//	slog.NewTextHandler(
	//		os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	//)
	return log

}
