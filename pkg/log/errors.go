package log

import "errors"

var (
	// ErrUnknownLevel unknown log level
	ErrUnknownLevel = errors.New("unknown log level")
)
