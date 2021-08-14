// Package simplelogger implements a simplelogger. It doesn't do much.
package simplelogger

import (
	"fmt"
	"os"
)

// SimpleLogger is a simple logger. It doesn't do much.
type SimpleLogger struct{}

// New returns a SimpleLogger.
func New() *SimpleLogger {
	return &SimpleLogger{}
}

// Info logs an info.
func (l *SimpleLogger) Info(args ...interface{}) {
	fmt.Fprintln(os.Stdout, args...)
}

// Error logs an error.
func (l *SimpleLogger) Error(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

// Warn logs a warning.
func (l *SimpleLogger) Warn(args ...interface{}) {
	fmt.Fprintln(os.Stdout, args...)
}
