package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information
type Logger struct {
	threshold Level
	output    io.Writer
	limit     int
}

// New returns you a logger, ready to log at the required threshold
func New(threshold Level, opts ...Option) *Logger {
	logger := &Logger{threshold: threshold, output: os.Stdout, limit: 100}
	for _, configFunc := range opts {
		configFunc(logger)
	}
	return logger
}

// Debugf formats and prints message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	l.Logf(LevelDebug, format, args...)
}

// Infof formats and prints message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
	l.Logf(LevelInfo, format, args...)
}

// Errorf formats and prints message if the log level is error or higher
func (l *Logger) Errorf(format string, args ...any) {
	l.Logf(LevelError, format, args...)
}

// Logf prints the message to the output
// Add decorations here, if any
func (l *Logger) Logf(level Level, format string, args ...any) {
	if l.threshold > level {
		return
	}
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.limit <= 0 {
		l.limit = 100
	}
	message := fmt.Sprintf("%v %s \n", level, fmt.Sprintf(format, args...))
	if len(message) > l.limit {
		r := []rune(message)
		message = string(r[:l.limit])
	}
	_, _ = fmt.Fprint(l.output, message)
}
