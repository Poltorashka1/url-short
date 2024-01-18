package postgres

import (
	"github.com/lib/pq"
	"net/http"
	handlers "url-short/internal/handlers/general"
	"url-short/internal/storage"
)

func (p *PostgresDatabase) SaveUrl(urlToSave string, alias string) error {
	const op = "storage.postgres.SaveUrl"

	_, err := p.Db.Exec("INSERT INTO url(url, alias) VALUES($1, $2)", urlToSave, alias)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return handlers.NewErrResp(http.StatusInternalServerError, op, "Alias already exists")
		}
		return handlers.NewErrResp(http.StatusInternalServerError, op, err.Error())
	}
	return nil
}

func (p *PostgresDatabase) GetUrl(alias string) (string, error) {
	const op = "storage.sqlite.GetUrl"

	rows, err := p.Db.Query("SELECT url FROM url WHERE alias = $1", alias)
	if err != nil {
		return "", handlers.NewErrResp(http.StatusInternalServerError, op, err.Error())
	}

	if rows.Next() {
		var url string
		if err = rows.Scan(&url); err != nil {
			return "", handlers.NewErrResp(http.StatusInternalServerError, op, err.Error())
		}
		return url, nil
	} else {
		return "", handlers.NewErrResp(http.StatusNotFound, op, "Url Not Found")
	}
}

func (p *PostgresDatabase) GetAlias(urlReq string) (storage.AllAliasList, error) {
	const op = "storage.postgres.GetAlias"
	bad := storage.AllAliasList{}

	// request to database to get all record with url
	rows, err := p.Db.Query("SELECT * FROM url WHERE url = $1", urlReq)
	if err != nil {
		return bad, handlers.NewErrResp(http.StatusInternalServerError, op, err.Error())
	}

	allAlias := make([]storage.Path, 0, 8)

	if rows.Next() {
		for {
			var id int
			var url string
			var alias string

			err := rows.Scan(&id, &url, &alias)
			if err != nil {
				return bad, handlers.NewErrResp(http.StatusInternalServerError, op, err.Error())
			}

			// append alias to data
			aliasData := storage.NewPath(id, url, alias)
			allAlias = append(allAlias, *aliasData)

			if !rows.Next() {
				break // exit the loop when there are no more rows
			}
		}
		return storage.AllAliasList{Url: urlReq, Alias: allAlias, Code: http.StatusOK}, nil
	}
	return bad, handlers.NewErrResp(http.StatusNotFound, op, "Url Not Found")
}

func (p *PostgresDatabase) DeleteUrl(alias string) error {
	const op = "storage.sqlite.deleteUrl"
	result, err := p.Db.Exec("DELETE FROM url WHERE alias = $1", alias)

	// check database error
	if err != nil {
		return handlers.NewErrResp(http.StatusInternalServerError, op, err.Error())
	}

	rows, err := result.RowsAffected()
	// check RowsAffected() error
	if err != nil {
		return handlers.NewErrResp(http.StatusInternalServerError, op, err.Error())
	}

	// check if there is no rows affected
	if rows == 0 {
		return handlers.NewErrResp(http.StatusInternalServerError, op, "Alias not found")
	}

	return nil
}
