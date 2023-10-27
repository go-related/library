// Code generated by goa v3.13.2, DO NOT EDIT.
//
// Zerolog logger implementation
//
// Command:
// $ goa gen github.com/jt/books/design

package log

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

// Logger is an adapted zerologger
type Logger struct {
	*zerolog.Logger
}

// New creates a new zerologger
func New(serviceName string, isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Str("service", serviceName).Logger()
	return &Logger{&logger}
}

// Log is called by the log middleware to log HTTP requests key values
func (logger *Logger) Log(keyvals ...interface{}) error {
	fields := FormatFields(keyvals)
	logger.Info().Fields(fields).Msgf("HTTP Request")
	return nil
}

// formatFields formats input keyvals
// ref: https://github.com/goadesign/goa/blob/v1/logging/logrus/adapter.go#L64
func FormatFields(keyvals []interface{}) map[string]interface{} {
	n := (len(keyvals) + 1) / 2
	res := make(map[string]interface{}, n)
	for i := 0; i < len(keyvals); i += 2 {
		k := keyvals[i]
		var v interface{}
		if i+1 < len(keyvals) {
			v = keyvals[i+1]
		}
		res[fmt.Sprintf("%v", k)] = v
	}
	return res
}
