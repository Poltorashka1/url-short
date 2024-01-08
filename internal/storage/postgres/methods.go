package postgres

import (
	"fmt"
)

func (p *PostgresDatabase) SaveUrl(urlToSave string, alias string) error {
	const op = "storage.SaveUrl"

	_, err := p.Db.Exec("INSERT INTO url(url, alias) VALUES($1, $2)", urlToSave, alias)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresDatabase) GetUrl(alias string) (string, error) {
	const op = "storage.storage.GetUrl"
	rows, err := p.Db.Query("SELECT url FROM url WHERE alias = $1", alias)

	if rows.Next() {
		var url string
		if err = rows.Scan(&url); err != nil {
			return "", fmt.Errorf("%s: %s", op, err.Error())
		}
		return url, nil
	} else {
		return "", fmt.Errorf("%s: %s", op, "Url not found")
	}
}
