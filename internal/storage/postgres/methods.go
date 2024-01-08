package postgres

import (
	"fmt"
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

func (p *PostgresDatabase) GetAlias(url string) (string, error) {
	const op = "storage.postgres.GetAlias"
	rows, err := p.Db.Query("SELECT alias FROM url WHERE url = $1", url)
	if err != nil {
		return "", fmt.Errorf("%s: %s", op, err.Error())
	}

	var alias string
	if rows.Next() {
		err := rows.Scan(&alias)
		if err != nil {
			return "", fmt.Errorf("%s: %s", op, err.Error())
		}
		return alias, nil
	}
	return "", fmt.Errorf("%s: %s", op, "Alias Not Found")

}

func (p *PostgresDatabase) DeleteUrl(alias string) error {
	const op = "storage.postgres.deleteUrl"
	_, err := p.Db.Exec("DELETE FROM url WHERE alias = $1", alias)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err.Error())
	}
	return nil
}
