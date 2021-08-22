package logic

import "errors"

var (
	ErrLoginFailed           = errors.New("login failed")
	ErrInvalidSession        = errors.New("invalid session")
	ErrInternalError         = errors.New("something went wrong")
	ErrInvalidParameter      = errors.New("invalid parameter")
	ErrNotFound              = errors.New("resource not found")
	ErrInsufficientPrivilege = errors.New("insufficient privilege")
	ErrConflict              = errors.New("there was a conflict")
)
