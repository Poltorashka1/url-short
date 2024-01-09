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
		return fmt.Errorf("%s: %s", op, err.Error())
	}
	err = CheckUrl(url)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err.Error())
	}
	return nil
}

func CheckUrl(url string) error {
	const op = "handlers.CheckUrl"
	if len(url) <= 0 {
		return fmt.Errorf("%s: %s", op, "Length of url must be > 0")
	}
	pattern := `^https?:\/\/[^\s\/$.?#].[^\s]*$`
	if !regexp.MustCompile(pattern).MatchString(url) {
		return fmt.Errorf("%s: %s", op, "Url Not Valid")
	}
	return nil
}

func CheckAlias(alias string) error {
	const op = "handlers.CheckAlias"
	if len(alias) <= 0 {
		return fmt.Errorf("%s: %s", op, "Length of alias must be > 0")
	}
	return nil
}
