package Url_handlers

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"url-short/internal/handlers"
	"url-short/internal/storage"
)

type AllUrl struct {
	AllUrl []Urls `json:"AllUrl"`
	Code   int    `json:"code"`
}

type Urls struct {
	Id    int    `json:"id"`
	Url   string `json:"url"`
	Alias string `json:"alias"`
}

// GetAllUrlHandler returns all url in json response.
func GetAllUrlHandler(db storage.Storage, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.GetAllUrl"
		data := AllUrl{AllUrl: make([]Urls, 0, 8)}

		// get db to direct request to database
		db := db.GetDb()
		//  get all url and write all to data
		err := GetAllUrl(db, &data)
		if err != nil {
			errorData := handlers.NewErrorResponse(http.StatusInternalServerError, fmt.Errorf("%s - Server Error: %s", op, err.Error()).Error())

			// return error response and log it
			log.Error(fmt.Sprintf("%s: %s", op, err.Error()))
			handlers.EncodeJson(w, log, errorData)
			return
		}
		data.Code = http.StatusOK
		// return data response with all url
		handlers.EncodeJson(w, log, data)

	}
}

// GetAllUrl gets all url from database.
func GetAllUrl(db *sql.DB, data *AllUrl) error {
	res, err := db.Query("SELECT * FROM url")
	if err != nil {
		return err
	}

	for res.Next() {
		var id int
		var url string
		var alias string
		err := res.Scan(&id, &url, &alias)
		if err != nil {
			return err
		}
		// append url to data
		data.AllUrl = append(data.AllUrl, Urls{Id: id, Url: url, Alias: alias})
	}
	return nil
}

type Path struct {
	Url   string `json:"url"`
	Alias string `json:"alias"`
	Code  int    `json:"code"`
}

func NewPath(url, alias string, code int) *Path {
	return &Path{
		Url:   url,
		Alias: alias,
		Code:  code,
	}
}

// GetUrlFromAliasHandler returns specified url.
func GetUrlFromAliasHandler(db storage.Storage, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.GetUrl"
		alias := r.URL.Query().Get("alias")

		// get specified url from database
		url, err := db.GetUrl(alias)
		if err != nil {
			errorData := handlers.NewErrorResponse(http.StatusNotFound, fmt.Errorf("%s: %s", op, err.Error()).Error()) //"Url not found", err.Error()

			// return error response
			handlers.EncodeJson(w, log, errorData)
			return
		}

		// return specified url
		data := NewPath(url, alias, http.StatusOK)
		handlers.EncodeJson(w, log, data)

	}
}

// GetAliasFromUrlHandler gets all alias from url.
func GetAliasFromUrlHandler(db storage.Storage, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.GetAliasFromUrl"
		url := r.URL.Query().Get("url")
		// get specified alias from database
		alias, err := db.GetAlias(url)

		if err != nil {
			errorData := handlers.NewErrorResponse(http.StatusNotFound, fmt.Errorf("%s: %s", op, err.Error()).Error())

			// return error response
			handlers.EncodeJson(w, log, errorData)
			return
		}

		// return specified alias
		data := NewPath(url, alias, http.StatusOK)
		handlers.EncodeJson(w, log, data)
	}
}
