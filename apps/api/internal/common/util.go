package common

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	claims, ok := ctx.Value("user").(jwt.MapClaims)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid user claims")
	}

	userIDStr := fmt.Sprintf("%v", claims["user_id"])
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID format: %v", err)
	}

	return userID, nil
}

var DefaultBinaryPaths = map[string]map[string]string{
	"windows": {
		"postgresql": "C:\\Program Files\\PostgreSQL\\16\\bin",
		"mysql":      "C:\\Program Files\\MySQL\\MySQL Server 8.0\\bin",
		"mariadb":    "C:\\Program Files\\MariaDB 10.6\\bin",
		"mongodb":    "C:\\Program Files\\MongoDB\\Server\\6.0\\bin",
	},
	"linux": {
		"postgresql": "/usr/bin",
		"mysql":      "/usr/bin",
		"mariadb":    "/usr/bin",
		"mongodb":    "/usr/bin",
	},
	"darwin": {
		"postgresql": "/opt/homebrew/bin",
		"mysql":      "/opt/homebrew/bin",
		"mariadb":    "/opt/homebrew/bin",
		"mongodb":    "/opt/homebrew/bin",
	},
}


func GetDefaultBinaryPath(dbType string) string {
	if paths, ok := DefaultBinaryPaths[runtime.GOOS]; ok {
		return paths[dbType]
	}
	return ""
}
