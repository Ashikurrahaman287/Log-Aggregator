package log

import (
	"time"
)

// LogEntry represents a log entry
type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source"`
	Message   string    `json:"message"`
	Level     string    `json:"level"`
}
