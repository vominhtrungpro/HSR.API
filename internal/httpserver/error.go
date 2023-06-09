package httpserver

import (
	"fmt"
	"net/http"
)

const (
	// DefaultErrorCode represents the default code for Error
	DefaultErrorCode = "internal_error"
	// DefaultErrorDesc represents the default description for Error
	DefaultErrorDesc = "Something went wrong"
)

var (
	// ErrDefaultInternal represents the default internal server error
	ErrDefaultInternal = &Error{
		Status: http.StatusInternalServerError,
		Code:   DefaultErrorCode,
		Desc:   DefaultErrorDesc,
	}
)

// Error represents a handler error.
type Error struct {
	Status int    `json:"-"`
	Code   string `json:"error"`
	Desc   string `json:"error_description"`
}

// Error satisfies the error interface
func (err Error) Error() string {
	return fmt.Sprintf("Status: [%d], Code: [%s], Desc: [%s]", err.Status, err.Code, err.Desc)
}
