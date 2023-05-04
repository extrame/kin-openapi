package errors

import "fmt"

//Operation errors
const (
	NoSuchHTTPMethod = 10000 + iota
	NoSuchOperationIsDefinedInPath
)

//Validation errors
const (
	NoResponseIsDefined = 20000 + iota
	NoResponseCodeInResponses
)

type Error struct {
	// The error message.
	Message string `json:"message"`
	// The error code.
	Code int `json:"code"`
}

func New(code int, message string) *Error {
	return &Error{Message: message, Code: code}
}

func Errorf(code int, format string, args ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, args...), Code: code}
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d]%s", e.Code, e.Message)
}
