package models

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("user with this email already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid email or password")
)
