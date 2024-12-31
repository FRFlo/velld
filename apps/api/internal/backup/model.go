package backup

import (
	"time"

	"github.com/google/uuid"
)

type Backup struct {
	ID            uuid.UUID  `json:"id"`
	ConnectionID  string     `json:"connection_id"`
	Size          int64      `json:"size"`
	Status        string     `json:"status"`
	Path          string     `json:"path"`
	ScheduledTime time.Time  `json:"scheduled_time"`
	StartedAt     time.Time  `json:"started_at"`
	CompletedTime *time.Time `json:"completed_time"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type BackupRequest struct {
	ConnectionID string `json:"connection_id"`
}

type BackupStats struct {
	TotalBackups    int     `json:"total_backups"`
	TotalSize       int64   `json:"total_size"`
	AverageSize     int64   `json:"average_size"`
	AverageDuration float64 `json:"average_duration"`
	SuccessRate     float64 `json:"success_rate"`
	LastBackupTime  string  `json:"last_backup_time"`
}
