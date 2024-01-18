package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"
)

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

// CheckUrlAndAlias check alias and url for validity
func CheckUrlAndAlias(url, alias string) error {
	const op = "handlers.CheckUrlAndAlias"

	err := CheckAlias(alias)
	if err != nil {
		if err, ok := err.(*ResponseError); ok {
			err.AddPath(op)
		}
		return err
	}

	err = CheckUrl(url)
	if err != nil {

		if err, ok := err.(*ResponseError); ok {
			err.AddPath(op)
		}
		return err
	}
	return nil
}

func CheckUrl(url string) error {
	const op = "handlers.CheckUrl"

	if len(url) <= 0 {
		err := NewErrResp(http.StatusBadRequest, op, "The url parameter is missing")
		return err
	}

	pattern := `^https?:\/\/[^\s\/$.?#].[^\s]*$`
	if !regexp.MustCompile(pattern).MatchString(url) {
		err := NewErrResp(http.StatusBadRequest, op, "Url format not valid: Url must start with 'http://' or 'https://'")
		return err
	}

	return nil
}

func CheckAlias(alias string) error {
	const op = "handlers.CheckAlias"

	if len(alias) <= 0 {
		err := NewErrResp(http.StatusBadRequest, op, "The alias parameter is missing or less than 1 character")
		return err
	}
	return nil
}
