package common

import (
	"context"
	"fmt"
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
