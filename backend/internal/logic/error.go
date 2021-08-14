package logic

import "errors"

var (
	ErrLoginFailed      = errors.New("login failed")
	ErrInvalidSession   = errors.New("invalid session")
	ErrInternalError    = errors.New("something went wrong")
	ErrInvalidParameter = errors.New("invalid parameter")
)
