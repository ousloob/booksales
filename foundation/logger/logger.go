// Package logger provides support for initializing the log system.
package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
)

// TracedIDFn represents a function that can return the trace id from the specified context.
type TracedIDFn func(ctx context.Context) string

// Logger represents a logger for logging information.
type Logger struct {
	handler    slog.Handler
	tracedIDFn TracedIDFn
}

// New constructs a new log for application use.
func New(w io.Writer, minLevel slog.Level, serviceName string, traceIDFn TracedIDFn) *Logger {

	// Convert the file name to just the name.ext when this key/value will be logged.
	f := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			if source, ok := a.Value.Any().(*slog.Source); ok {
				v := fmt.Sprintf("%s:%d", filepath.Base(source.File), source.Line)
				return slog.Attr{Key: "file", Value: slog.StringValue(v)}
			}
		}

		return a
	}

	// Construct the slog JSON handler for use.
	handler := slog.Handler(slog.NewJSONHandler(w, &slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: f,
	}))

	// Attributes to add to every log.
	attrs := []slog.Attr{
		{Key: "service", Value: slog.StringValue(serviceName)},
	}

	// Add those attributes and capture the final handler.
	handler = handler.WithAttrs(attrs)

	return &Logger{
		handler:    handler,
		tracedIDFn: traceIDFn,
	}
}
