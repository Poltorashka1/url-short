package Url_handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"url-short/internal/handlers"
	"url-short/internal/storage"
)

// DeleteUrlHandler delete alias from database
func DeleteUrlHandler(db storage.Storage, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.DeleteUrl"

		alias := r.URL.Query().Get("alias")
		err := db.DeleteUrl(alias)
		if err != nil {
			errorData := handlers.NewErrorResponse(http.StatusNotFound, fmt.Errorf("%s: %s", op, err.Error()).Error())

			handlers.EncodeJson(w, log, errorData)
			return
		}

		data := handlers.NewSuccessResponse(http.StatusOK, "Delete Success")
		handlers.EncodeJson(w, log, data)
	}
}
