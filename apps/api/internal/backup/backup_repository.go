package backup

import (
	"database/sql"

	"github.com/google/uuid"
)

type BackupRepository struct {
	db *sql.DB
}

func NewBackupRepository(db *sql.DB) *BackupRepository {
	return &BackupRepository{
		db: db,
	}
}

func (r *BackupRepository) CreateBackup(backup *Backup) error {
	_, err := r.db.Exec("INSERT INTO backups (id, connection_id, status, path, size, scheduled_time, completed_time, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		backup.ID, backup.ConnectionID, backup.Status, backup.Path, backup.Size, backup.ScheduledTime, backup.CompletedTime, backup.CreatedAt, backup.UpdatedAt)
	return err
}

func (r *BackupRepository) GetAllBackups(UserID uuid.UUID) ([]*Backup, error) {
	query := `
			SELECT 
					b.id, b.connection_id, b.status, b.path, b.size, 
					b.scheduled_time, b.completed_time, b.created_at, b.updated_at
			FROM 
					backups b
			INNER JOIN 
					connections c ON b.connection_id = c.id
			WHERE 
					c.user_id = $1;
	`

	rows, err := r.db.Query(query, UserID)
	if err != nil {
			return nil, err
	}
	defer rows.Close()

	backups := make([]*Backup, 0)
	for rows.Next() {
			backup := &Backup{}
			err := rows.Scan(
					&backup.ID, &backup.ConnectionID, &backup.Status, &backup.Path,
					&backup.Size, &backup.ScheduledTime, &backup.CompletedTime,
					&backup.CreatedAt, &backup.UpdatedAt,
			)
			if err != nil {
					return nil, err
			}
			backups = append(backups, backup)
	}

	if err = rows.Err(); err != nil {
			return nil, err
	}

	return backups, nil
}


func (r *BackupRepository) GetBackup(id string) (*Backup, error) {
	backup := &Backup{}
	err := r.db.QueryRow("SELECT id, connection_id, status, path, size, scheduled_time, completed_time, created_at, updated_at FROM backups WHERE id = $1", id).
		Scan(&backup.ID, &backup.ConnectionID, &backup.Status, &backup.Path, &backup.Size, &backup.ScheduledTime, &backup.CompletedTime, &backup.CreatedAt, &backup.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return backup, nil
}
