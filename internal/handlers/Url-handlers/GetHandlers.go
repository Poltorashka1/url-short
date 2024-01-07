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
}

type Urls struct {
	Id    int    `json:"id"`
	Url   string `json:"url"`
	Alias string `json:"alias"`
	Code  int    `json:"code"`
}

func All(db storage.Storage, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.GetAllUrl"
		data := AllUrl{AllUrl: make([]Urls, 0, 8)}
		w.Header().Set("Content-type", "application/json")

		db := db.GetDb()
		err := GetAllUrl(db, &data)
		if err != nil {
			data := handlers.ErrorResponse{Error: handlers.Error{
				Code:    http.StatusInternalServerError,
				Message: "Server Error",
				Details: err.Error(),
			}}

			log.Error(fmt.Sprintf("%s: %s", op, err.Error()))
			handlers.EncodeJson(w, log, data)
			return
		}
		handlers.EncodeJson(w, log, data)
	}
}

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
		//data.Ekz[url] = Ekz{Id: id, Url: url, Alias: alias, Code: http.StatusOK}
		data.AllUrl = append(data.AllUrl, Urls{Id: id, Url: url, Alias: alias, Code: http.StatusOK})
	}
	return nil
}

type Alias struct {
	Url   string `json:"url"`
	Alias string `json:"alias"`
	Code  int    `json:"code"`
}

// GetUrlFromAlias gets url from alias.
func GetUrlFromAlias(db storage.Storage, log *slog.Logger) http.HandlerFunc {
	const op = "handlers.GetUrl"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		url := r.URL.Query().Get("alias")

		res, err := db.GetUrl(url)
		if err != nil {
			data := handlers.ErrorResponse{Error: handlers.Error{
				Code:    http.StatusNotFound,
				Message: "Url not found",
				Details: err.Error(),
			}}
			handlers.EncodeJson(w, log, data)
			return
		}

		al := Alias{
			Url:   url,
			Alias: res,
			Code:  http.StatusOK,
		}
		handlers.EncodeJson(w, log, al)
	}
}
