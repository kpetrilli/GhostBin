package log

import (
	"log/slog"
	"os"
)

// NewLogger creates a new logger and saves logs to a file.
func NewLogger(logFilePath string, text ...bool) *slog.Logger {
	// Open the log file for writing, creating it if it doesn't exist
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		slog.Error("failed to open log file: %v", "error", err)
		os.Exit(1)
	}

	// Check if the log format should be text or JSON
	if len(text) > 0 && text[0] {
		return slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}

	return slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
}
