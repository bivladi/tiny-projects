package pocketlog

// Logger is used to log information
type Logger struct {
	threshold Level
}

func New(threshold Level) *Logger {
	return &Logger{threshold: threshold}
}

// Debugf formats and prints message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	// implement me
}

// Infof formats and prints message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
	// implement me
}

// Errorf formats and prints message if the log level is error or higher
func (l *Logger) Errorf(format string, args ...any) {

}
