package handlers

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
