package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("user have been exist")
	ErrorUserNotExist    = errors.New("user is not exist")
	ErrorInvalidPassword = errors.New("incorrect password")
	ErrorInvalidCommID   = errors.New("community id is valid")
	// Err
)
