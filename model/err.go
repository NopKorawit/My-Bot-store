package model

import "errors"

var (
	ErrRepository           = errors.New("repository error")
	ErrCodenotFound         = errors.New("code not found")
	ErrProductAlreadyExists = errors.New("Product already exists")
	ErrDuplicateROW         = errors.New("row is duplicated")
	ErrNotNumber            = errors.New("this value not a numer")
	ErrProductNotEnough     = errors.New("Product not enough")
)
