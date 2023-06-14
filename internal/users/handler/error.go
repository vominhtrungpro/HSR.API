package handler

import (
	"errors"
)

var (
	errInvalidUsername = errors.New("username invalid")
	errInvalidPassword = errors.New("password invalid")
	errInvalidEmail    = errors.New("email invalid")
)
