package postgres

import (
	"fmt"
	"net/http"
	"url-short/internal/storage"
)

func (p *PostgresDatabase) SaveUrl(urlToSave string, alias string) error {
	const op = "storage.postgres.SaveUrl"

	_, err := p.Db.Exec("INSERT INTO url(url, alias) VALUES($1, $2)", urlToSave, alias)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresDatabase) GetUrl(alias string) (string, error) {
	const op = "storage.postgres.GetUrl"
	rows, err := p.Db.Query("SELECT url FROM url WHERE alias = $1", alias)
	if err != nil {
		return "", fmt.Errorf("%s: %s", op, err.Error())
	}

	if rows.Next() {
		var url string
		err := rows.Scan(&url)
		if err != nil {
			return "", fmt.Errorf("%s: %s", op, err.Error())
		}
		return url, nil
	}
	return "", fmt.Errorf("%s: %s", op, "Url Not Found")

}

func (p *PostgresDatabase) GetAlias(urlReq string) (storage.AllAliasList, error) {
	const op = "storage.postgres.GetAlias"
	bad := storage.AllAliasList{}

	// request to database to get all record with url
	rows, err := p.Db.Query("SELECT * FROM url WHERE url = $1", urlReq)
	if err != nil {
		return bad, fmt.Errorf("%s: %s", op, err.Error())
	}

	// ToDo: refactor []Path to more fast way
	allAlias := make([]storage.Path, 0, 8)

	if rows.Next() {
		for {
			var id int
			var url string
			var alias string

			err := rows.Scan(&id, &url, &alias)
			if err != nil {
				return bad, fmt.Errorf("%s: %s", op, err.Error())
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
	return bad, fmt.Errorf("%s: %s", op, "Alias Not Found")
}

func (p *PostgresDatabase) DeleteUrl(alias string) error {
	const op = "storage.postgres.deleteUrl"
	res, err := p.Db.Exec("DELETE FROM url WHERE alias = $1", alias)
	if c, _ := res.RowsAffected(); c == 0 {
		return fmt.Errorf("%s: %s", op, "Url Not Found to delete")
	}

	if err != nil {
		return fmt.Errorf("%s: %s", op, err.Error())
	}
	return nil
}
