package handlers

// TODO: refactor handlers errorResponse format to make it more readable

// ErrorResponse is an error response.
type ErrorResponse struct {
	Error `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponse create new error response.
func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Error: Error{
			Code:    code,
			Message: message,
		},
	}
}
