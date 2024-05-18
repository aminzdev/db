package db

import "errors"

var (
	ErrInternal      = errors.New("internal error")
	ErrNotFound      = errors.New("record not found")
	ErrAlreadyExists = errors.New("record already exists")
)
