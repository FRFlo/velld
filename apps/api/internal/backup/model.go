package backup

import (
	"time"

	"github.com/google/uuid"
)

type Backup struct {
	ID                   uuid.UUID  `json:"id"`
	ConnectionID         string     `json:"connection_id"`
	DatabaseName         string     `json:"database_name"`
	DatabaseSizeAtBackup int64      `json:"database_size_at_backup"`
	BackupSize           int64      `json:"backup_size"`
	CompressionRatio     float64    `json:"compression_ratio"`
	Status               string     `json:"status"`
	Path                 string     `json:"path"`
	FileName             string     `json:"file_name"`
	BackupType           string     `json:"backup_type"`
	ScheduledTime        time.Time  `json:"scheduled_time"`
	StartedAt            time.Time  `json:"started_at"`
	CompletedTime        *time.Time `json:"completed_time"`
	DurationSeconds      int        `json:"duration_seconds"`
	ErrorMessage         *string    `json:"error_message"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

type BackupRequest struct {
	ConnectionID string `json:"connection_id"`
	BackupType   string `json:"backup_type"`
}

type BackupStats struct {
	TotalBackups    int     `json:"total_backups"`
	TotalSize       int64   `json:"total_size"`
	AverageSize     int64   `json:"average_size"`
	AverageDuration float64 `json:"average_duration"`
	SuccessRate     float64 `json:"success_rate"`
	LastBackupTime  string  `json:"last_backup_time"`
}
