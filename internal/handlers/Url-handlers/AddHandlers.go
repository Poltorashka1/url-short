package Url_handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"url-short/internal/handlers"
	"url-short/internal/storage"
)

// TODO: check url in AddAliasForUrl()

// AddAliasForUrlHandler adds a new alias for url
func AddAliasForUrlHandler(db storage.Storage, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.AddAliasForUrl"

		newUrl := r.URL.Query().Get("url")
		alias := r.URL.Query().Get("alias")

		// save url in database
		err := db.SaveUrl(newUrl, alias)
		if err != nil {
			errorData := handlers.NewErrorResponse(http.StatusInternalServerError, fmt.Errorf("%s - Url already exists: %s", op, err.Error()).Error())

			// return error response
			handlers.EncodeJson(w, log, errorData)
			//log.Error(fmt.Sprintf("%s: %s", op, err.Error()))
			return
		}
		// return success response
		successData := handlers.NewSuccessResponse(http.StatusOK, "Success added new url")
		handlers.EncodeJson(w, log, successData)
	}

}
