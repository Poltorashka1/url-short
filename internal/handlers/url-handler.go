package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"url-short/internal/storage"
)

func GetAllUrl(db storage.Storage, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.GetAllUrl"
		db = db.GetDb()
		res, err := db.Db.Query("SELECT * FROM url")
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
			_, err = w.Write([]byte(fmt.Sprintf("id: %d, url: %s, alias: %s\n", id, url, alias)))

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
