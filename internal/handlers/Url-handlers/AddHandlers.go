package Url_handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"regexp"
	"url-short/internal/handlers"
	"url-short/internal/storage"
)

// AddAliasForUrlHandler adds a new alias for url
func AddAliasForUrlHandler(db storage.Storage, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.Url-handlers.AddAliasForUrl"

		url := r.URL.Query().Get("url")
		alias := r.URL.Query().Get("alias")

		// check url and alias for validity
		err := CheckUrl(url, alias)
		if err != nil {
			errorData := handlers.NewErrorResponse(http.StatusBadRequest, fmt.Errorf("%s: %s", op, err.Error()).Error())

			// return error response
			handlers.EncodeJson(w, log, errorData)
			return
		}
		// save url in database
		err = db.SaveUrl(url, alias)
		if err != nil {
			errorData := handlers.NewErrorResponse(http.StatusConflict, fmt.Errorf("%s - Url already exists: %s", op, err.Error()).Error())

			// return error response
			handlers.EncodeJson(w, log, errorData)
			return
		}
		// return success response
		successData := handlers.NewSuccessResponse(http.StatusOK, "Success added new url")
		handlers.EncodeJson(w, log, successData)
	}

}

// CheckUrl check alias and url for validity
func CheckUrl(url, alias string) error {
	const op = "handlers.Url-handlers.CheckUrl"
	if len(url) <= 0 {
		return fmt.Errorf("%s: %s", op, "Length of url must be > 0")
	}
	if len(alias) <= 0 {
		return fmt.Errorf("%s: %s", op, "Length of alias must be > 0")
	}
	pattern := `^https?:\/\/[^\s\/$.?#].[^\s]*$`
	if !regexp.MustCompile(pattern).MatchString(url) {
		return fmt.Errorf("%s: %s", op, "Url Not Valid")
	}
	return nil
}
