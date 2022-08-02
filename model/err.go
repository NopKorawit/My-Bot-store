package model

import "errors"

var (
	ErrRepository        = errors.New("repository error")
	ErrCodenotFound      = errors.New("code not found")
	ErrGoodAlreadyExists = errors.New("good already exists")
	ErrDuplicateROW      = errors.New("row is duplicated")
	ErrNotNumber         = errors.New("this value not a numer")
	ErrGoodNotEnough     = errors.New("good not enough")
)
