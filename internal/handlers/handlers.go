package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

// ErrorResponse is an error response.
type ErrorResponse struct {
	Error `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

// NewErrorResponse create new error response.
func NewErrorResponse(code int, message string, details string) *ErrorResponse {
	return &ErrorResponse{
		Error: Error{
			Code:    code,
			Message: message,
			Details: details,
		},
	}
}

// SuccessResponse is a success response.
type SuccessResponse struct {
	Success `json:"success"`
}

type Success struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewSuccessResponse create new success response.
func NewSuccessResponse(code int, message string) *SuccessResponse {
	return &SuccessResponse{Success{
		Code:    code,
		Message: message,
	}}
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
