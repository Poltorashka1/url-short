package Url_handlers

import (
	"database/sql"
	"log/slog"
	"net/http"
	"url-short/internal/handlers/general"
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

// TODO : refactor this handler

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
			handlers.AddPath(err, op)
			handlers.EncodeJson(w, log, err)
			return
		}
		data.Code = http.StatusOK
		// return data response with all url
		handlers.EncodeJson(w, log, data)

	}
}

// TODO : GetAllUrl() replace to storage.methods

// GetAllUrl gets all url from database.
func GetAllUrl(db *sql.DB, data *AllUrl) error {
	const op = "handlers.GetAllUrl"

	res, err := db.Query("SELECT * FROM url")
	if err != nil {
		return handlers.NewErrResp(http.StatusInternalServerError, op, err.Error())
	}

	for res.Next() {
		var id int
		var url string
		var alias string
		err := res.Scan(&id, &url, &alias)
		if err != nil {
			err := handlers.NewErrResp(http.StatusInternalServerError, op, err.Error())
			return err
		}
		// append url to data
		data.AllUrl = append(data.AllUrl, Urls{Id: id, Url: url, Alias: alias})
	}

	return nil
}

// TODO : replace struct to new place

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

		// check alias for validity
		alias := r.URL.Query().Get("alias")
		err := handlers.CheckAlias(alias)
		if err != nil {
			handlers.AddPath(err, op)
			handlers.EncodeJson(w, log, err)
			return
		}

		// get specified url from database
		url, err := db.GetUrl(alias)
		if err != nil {
			handlers.AddPath(err, op)
			handlers.EncodeJson(w, log, err)
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

		// check url for validity
		url := r.URL.Query().Get("url")
		err := handlers.CheckUrl(url)
		if err != nil {
			handlers.AddPath(err, op)
			handlers.EncodeJson(w, log, err)
			return
		}

		// get specified alias from database
		allAliasList, err := db.GetAlias(url)
		if err != nil {
			handlers.AddPath(err, op)
			handlers.EncodeJson(w, log, err)
			return
		}

		// return specified alias
		handlers.EncodeJson(w, log, allAliasList)
	}
}
