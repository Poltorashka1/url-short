package sqlite

import (
	"fmt"
)

func (s *SqliteDatabase) SaveUrl(urlToSave string, alias string) error {
	const op = "storage.SaveUrl"

	_, err := s.Db.Exec("INSERT INTO url(url, alias) VALUES(?, ?)", urlToSave, alias)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqliteDatabase) GetUrl(alias string) (string, error) {
	const op = "storage.storage.GetUrl"
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
