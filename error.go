package server

import "fmt"

func NewError(code string, message string) error {
	return &GoServerError{
		Code:    code,
		Message: message,
	}
}

type GoServerError struct {
	Code    string
	Message string
}

func (err *GoServerError) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message)
}
