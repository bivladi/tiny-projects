package pocketlog

import "io"

// Option defines a functional option to our logger
type Option func(*Logger)

func WithOutput(output io.Writer) Option {
	return func(logger *Logger) {
		logger.output = output
	}
}

func WithLimit(limit int) Option {
	return func(logger *Logger) {
		logger.limit = limit
	}
}
