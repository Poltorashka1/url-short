package Url_handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"url-short/internal/handlers/general"
	"url-short/internal/storage"
)

// DeleteUrlHandler delete alias from database
func DeleteUrlHandler(db storage.Storage, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.DeleteUrl"

		alias := r.URL.Query().Get("alias")
		fmt.Println(alias)

		err := handlers.CheckAlias(alias)
		if err != nil {
			handlers.AddPath(err, op)
			handlers.EncodeJson(w, log, err)
			return
		}

		err = db.DeleteUrl(alias)
		if err != nil {
			handlers.AddPath(err, op)
			handlers.EncodeJson(w, log, err)
			return
		}

		resultData := handlers.NewSuccessResponse(http.StatusOK, "Delete Success")
		handlers.EncodeJson(w, log, resultData)
	}
}
