package common

import (
	"fmt"
	"time"
)

func ParseTime(timeStr string) (time.Time, error) {
	formats := []string{
		time.RFC3339,                // "2006-01-02T15:04:05Z07:00"
		"2006-01-02 15:04:05-07:00", // SQLite format with timezone
		"2006-01-02 15:04:05+07:00", // SQLite format with timezone
		"2006-01-02 15:04:05",       // SQLite format without timezone
	}

	var lastErr error
	for _, format := range formats {
		t, err := time.Parse(format, timeStr)
		if err == nil {
			return t, nil
		}
		lastErr = err
	}
	return time.Time{}, fmt.Errorf("could not parse time '%s': %v", timeStr, lastErr)
}
