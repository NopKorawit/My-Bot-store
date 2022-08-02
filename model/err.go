package model

import "errors"

var (
	ErrRepository        = errors.New("repository error")
	ErrCodenotFound      = errors.New("code not found")
	ErrGoodAlreadyExists = errors.New("good already exists")
	ErrDuplicateROW      = errors.New("row is duplicated")
)
