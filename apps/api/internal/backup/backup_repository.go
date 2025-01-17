package backup

import (
	"database/sql"
	"fmt"
	"strings"

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
	_, err := r.db.Exec("INSERT INTO backups (id, connection_id, status, path, size, scheduled_time, started_time, completed_time, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		backup.ID, backup.ConnectionID, backup.Status, backup.Path, backup.Size, backup.ScheduledTime, backup.StartedTime, backup.CompletedTime, backup.CreatedAt, backup.UpdatedAt)
	return err
}

func (r *BackupRepository) GetAllBackups(UserID uuid.UUID) ([]*BackupList, error) {
	query := `
			SELECT 
					b.id, b.connection_id, c.type, b.status, b.path, b.size, 
					b.scheduled_time, b.started_time, b.completed_time, b.created_at, b.updated_at
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

	backups := make([]*BackupList, 0)
	for rows.Next() {
		backup := &BackupList{}
		err := rows.Scan(
			&backup.ID, &backup.ConnectionID, &backup.DatabaseType, &backup.Status, &backup.Path,
			&backup.Size, &backup.ScheduledTime, &backup.StartedTime, &backup.CompletedTime,
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

type BackupListOptions struct {
	UserID uuid.UUID
	Limit  int
	Offset int
	Search string
}

func (r *BackupRepository) GetAllBackupsWithPagination(opts BackupListOptions) ([]*BackupList, int, error) {
	whereClause := "WHERE c.user_id = $1"
	args := []interface{}{opts.UserID}
	argCount := 2

	if opts.Search != "" {
		whereClause += fmt.Sprintf(" AND (LOWER(b.path) LIKE $%d OR LOWER(b.status) LIKE $%d)", argCount, argCount)
		args = append(args, "%"+strings.ToLower(opts.Search)+"%")
		argCount++
	}

	// Get total count
	countQuery := fmt.Sprintf(`
		SELECT COUNT(*) 
		FROM backups b
		INNER JOIN connections c ON b.connection_id = c.id
		%s`, whereClause)

	var total int
	if err := r.db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Get paginated results
	query := fmt.Sprintf(`
		SELECT 
			b.id, b.connection_id, c.type, b.status, b.path, b.size, 
			b.scheduled_time, b.started_time, b.completed_time, b.created_at, b.updated_at
		FROM backups b
		INNER JOIN connections c ON b.connection_id = c.id
		%s
		ORDER BY b.created_at DESC
		LIMIT $%d OFFSET $%d
	`, whereClause, argCount, argCount+1)

	args = append(args, opts.Limit, opts.Offset)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	backups := make([]*BackupList, 0)
	for rows.Next() {
		backup := &BackupList{}
		err := rows.Scan(
			&backup.ID, &backup.ConnectionID, &backup.DatabaseType, &backup.Status, &backup.Path,
			&backup.Size, &backup.ScheduledTime, &backup.StartedTime, &backup.CompletedTime,
			&backup.CreatedAt, &backup.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		backups = append(backups, backup)
	}

	return backups, total, rows.Err()
}
