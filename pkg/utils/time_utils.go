package utils

import (
	"fmt"
	"time"
)

// FormatDuration formats a time.Duration into a human-readable string.
func FormatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	if h > 0 {
		return fmt.Sprintf("%dч %dмин", h, m)
	}
	return fmt.Sprintf("%dмин", m)
}

// ParseDuration parses a string into a time.Duration.
func ParseDuration(durationStr string) (time.Duration, error) {
	return time.ParseDuration(durationStr)
}

// GetCurrentTime returns the current time.
func GetCurrentTime() time.Time {
	return time.Now()
}

// GetElapsedTime calculates the elapsed time between two time.Time values.
func GetElapsedTime(start, end time.Time) time.Duration {
	return end.Sub(start)
}
