package controller

import "errors"

var (
	errUsernameExist     = errors.New("username exist")
	errEmailExist        = errors.New("email exist")
	errIncorrectPassword = errors.New("incorect password")
)
