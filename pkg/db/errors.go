package db

import "errors"

var (
	// ErrUnknownDriver database driver unknown
	ErrUnknownDriver = errors.New("database driver unknown")
)
