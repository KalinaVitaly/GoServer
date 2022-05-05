package store

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrWriteFile      = errors.New("error record data")
	ErrFoundFile      = errors.New("error find file")
	ErrReadFile       = errors.New("error read file")
)
