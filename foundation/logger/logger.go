// Package logger provides support for initializing the log system.
package logger

import (
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
)

// New constructs a new log for application use.
func New(w io.Writer, serviceName string, build string) *slog.Logger {

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
	handler := slog.Handler(slog.NewTextHandler(w, &slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: f,
	}))

	// Attributes to add to every log.
	attrs := []slog.Attr{
		{Key: "service", Value: slog.StringValue(serviceName)},
	}

	// Add those attributes and capture the final handler.
	handler = handler.WithAttrs(attrs)

	return slog.New(handler).With(serviceName, build)
}
