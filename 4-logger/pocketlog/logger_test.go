package pocketlog_test

import (
	"fmt"
	"learngo/logger/pocketlog"
	"testing"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output:
	// [DEBUG] Hello, world
}

const (
	debugMessage = "Debug message,"
	infoMessage  = "Info message,"
	errorMessage = "Error message,"
)

func TestLogger_DebugfInfofErrorf(t *testing.T) {
	tt := map[string]struct {
		level    pocketlog.Level
		expected string
	}{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: fmt.Sprintf("[DEBUG] %s \n[INFO] %s \n[ERROR] %s \n", debugMessage, infoMessage, errorMessage),
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: fmt.Sprintf("[INFO] %s \n[ERROR] %s \n", infoMessage, errorMessage),
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: fmt.Sprintf("[ERROR] %s \n", errorMessage),
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}
			testLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			testLogger.Debugf(debugMessage)
			testLogger.Infof(infoMessage)
			testLogger.Errorf(errorMessage)

			if tw.contents != tc.expected {
				t.Fatalf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

func TestLogger_LogfLimit(t *testing.T) {
	tt := map[string]struct {
		limit    int
		input    string
		expected int
	}{
		"message less than limit": {
			limit:    100,
			expected: len(fmt.Sprintf("[INFO] %s \n", infoMessage)),
		},
		"message bigger than limit": {
			limit:    5,
			expected: 5,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}
			testLogger := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(tw), pocketlog.WithLimit(tc.limit))
			testLogger.Logf(pocketlog.LevelInfo, infoMessage)
			if len(tw.contents) != tc.expected {
				t.Fatalf("wrong output expected %v got %q", tc.expected, tw.contents)
			}
		})
	}
}

// testWriter is a struct that implements io.Writer.
// We use it to validate that we can write to a specific output.
type testWriter struct {
	contents string
}

// Write implements the io.Writer interface.
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
