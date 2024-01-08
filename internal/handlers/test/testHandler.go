package test

import (
	"log/slog"
	"net/http"
)

// GetTestResult its test handler function.
func GetTestResult(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/plain; charset=utf-8")
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			log.Debug(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
