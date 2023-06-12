package handler

import "errors"

const ErrCodeValidationFailed = "validation_failed"

var (
	errInvalidName = errors.New("invalid name")
)
