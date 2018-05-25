package model

import "errors"

var(
	ErrUserNotExist = errors.New("user not exist")
	ErrInvalidPassword = errors.New("passwd or username not right")
	ErrInvalidParams = errors.New("param is invalid")
	ErrUerExist = errors.New("user already exist")
	ErrInsertUser = errors.New("insert user error")
)