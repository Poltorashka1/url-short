package sqlite

import (
	"fmt"
)

func (s *SqliteDatabase) SaveUrl(urlToSave string, alias string) error {
	const op = "storage.sqlite.SaveUrl"

	_, err := s.Db.Exec("INSERT INTO url(url, alias) VALUES(?, ?)", urlToSave, alias)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqliteDatabase) GetUrl(alias string) (string, error) {
	const op = "storage.sqlite.GetUrl"
	rows, err := s.Db.Query("SELECT url FROM url WHERE alias = ?", alias)

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

func (s *SqliteDatabase) GetAlias(url string) (string, error) {
	const op = "storage.sqlite.GetAlias"
	rows, err := s.Db.Query("SELECT alias FROM url WHERE url = ?", url)
	if err != nil {
		return "", fmt.Errorf("%s: %s", op, err.Error())
	}

	var alias string
	for rows.Next() {
		err := rows.Scan(&alias)
		if err != nil {
			return "", fmt.Errorf("%s: %s", op, err.Error())
		}
	}
	return alias, nil
}

func (s *SqliteDatabase) DeleteUrl(alias string) error {
	const op = "storage.sqlite.deleteUrl"
	_, err := s.Db.Exec("DELETE FROM url WHERE alias = ?", alias)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err.Error())
	}
	return nil
}
