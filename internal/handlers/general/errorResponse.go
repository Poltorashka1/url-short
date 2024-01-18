package handlers

// TODO: refactor handlers errorResponse format to make it more readable

// ErrorResponse is an error response.

//type Error struct {
//	Code         int    `json:"code"`
//	Message      string `json:"message"`
//	ShortMessage string `json:"shortMessage"`
//}
//
//// NewErrorResponse create new error response.
//func NewErrorResponse(code int, message string) error {
//	return &Error{
//		Code:    code,
//		Message: message,
//	}
//
//}
//func (e *Error) Error() string {
//	return e.Message
//
//}

//	type ResponseError interface {
//		Error() string
//		AddPath(path string)
//	}

type ResponseError struct {
	Err error `json:"error"`
}

type ErrorV2 struct {
	Code    int    `json:"code"`
	Path    string `json:"path"`
	Message string `json:"message"`
}

func AddPath(err error, path string) {
	switch err.(type) {
	case *ResponseError:
		err.(*ResponseError).AddPath(path)
	}
}

func NewErrResp(code int, path string, message string) *ResponseError {
	return &ResponseError{
		&ErrorV2{
			Code:    code,
			Path:    path,
			Message: message,
		},
	}
}

func (e *ErrorV2) Error() string {
	return e.Message
}

func (e *ErrorV2) AddPath(path string) {
	e.Path = path + ": " + e.Path
}

func (e *ResponseError) AddPath(path string) {
	e.Err.(*ErrorV2).AddPath(path)
}
func (e *ResponseError) Error() string {
	return e.Err.Error()
}
