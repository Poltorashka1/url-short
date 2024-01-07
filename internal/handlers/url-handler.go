package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
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
		//res := GetAllUrl(db, log)
		const op = "handlers.GetAllUrl"
		data := AllUrl{AllUrl: make([]Urls, 0, 8)}
		w.Header().Set("Content-type", "application/json")

		// debug, to be refactor
		db := db.GetDb()
		res, err := db.Query("SELECT * FROM url")
		if err != nil {
			err := fmt.Errorf("%s: %s", op, err.Error()).Error()
			log.Error(err)
		}
		for res.Next() {
			var id int
			var url string
			var alias string
			err := res.Scan(&id, &url, &alias)
			if err != nil {
				err := fmt.Errorf("%s: %s", op, err.Error()).Error()
				log.Error(err)
				return
			}
			//data.Ekz[url] = Ekz{Id: id, Url: url, Alias: alias, Code: http.StatusOK}
			data.AllUrl = append(data.AllUrl, Urls{Id: id, Url: url, Alias: alias, Code: http.StatusOK})
		}
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
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
			data := storage.ErrorResponse{Error: storage.Error{
				Code:    http.StatusNotFound,
				Message: "Url not found",
				Details: "",
			}}
			EncodeJson(w, log, data)
			return
		}

		al := Alias{
			Url:   url,
			Alias: res,
			Code:  http.StatusOK,
		}
		EncodeJson(w, log, al)
	}
}

// EncodeJson create json response.
func EncodeJson(w http.ResponseWriter, log *slog.Logger, data interface{}) {
	const op = "handlers.EncodeJson"
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		err := fmt.Sprintf("%s: %s", op, err.Error())
		log.Error(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GetAllUrl(db storage.Storage, log *slog.Logger) {

}
