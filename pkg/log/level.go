package log

import "fmt"

// Level type
type Level string

const (
	// LevelDebug usually only enabled when debugging
	LevelDebug Level = "debug"
	// LevelInfo general operational entries about what's going on inside the application
	LevelInfo Level = "info"
	// LevelWarn non-critical entries that deserve eyes
	LevelWarn Level = "warn"
	// LevelError used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	LevelError Level = "error"
	// LevelFatal Logs and then calls `logger.Exit(1)`
	LevelFatal Level = "fatal"
	// LevelPanic logs and then calls panic
	LevelPanic Level = "panic"
)

// String return string value of a Level constant
func (l *Level) String() string {
	return string(*l)
}

// ParseLevel takes a string level and returns the Level constant
func ParseLevel(level string) (Level, error) {
	switch level {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warn":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	case "fatal":
		return LevelFatal, nil
	case "panic":
		return LevelPanic, nil
	default:
		return "", fmt.Errorf("%v: %s", ErrUnknownLevel, level)
	}
}
