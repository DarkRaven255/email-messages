package domain

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrRecordNotFound      = errors.New("Record not found")
	ErrConflict            = errors.New("Your item already exist")
	ErrBadParamInput       = errors.New("Given parameter is not valid")
)