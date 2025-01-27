package connection

import (
	"database/sql"

	"github.com/dendianugerah/velld/internal/common"
	"github.com/google/uuid"
)

type ConnectionRepository struct {
	db     *sql.DB
	crypto *common.EncryptionService
}

func NewConnectionRepository(db *sql.DB, crypto *common.EncryptionService) *ConnectionRepository {
	return &ConnectionRepository{
		db:     db,
		crypto: crypto,
	}
}

func (r *ConnectionRepository) Save(conn StoredConnection) error {
	username, err := r.crypto.Encrypt(conn.Username)
	if err != nil {
		return err
	}

	password, err := r.crypto.Encrypt(conn.Password)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO connections (
			id, name, type, host, port, username, password, 
			database_name, ssl, user_id, status, database_size
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
		)`

	_, err = r.db.Exec(
		query,
		conn.ID,
		conn.Name,
		conn.Type,
		conn.Host,
		conn.Port,
		username,
		password,
		conn.DatabaseName,
		conn.SSL,
		conn.UserID,
		conn.Status,
		conn.DatabaseSize,
	)

	return err
}

func (r *ConnectionRepository) GetConnection(id string) (*StoredConnection, error) {
	var conn StoredConnection
	var encryptedUsername, encryptedPassword string

	query := `SELECT * FROM connections WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&conn.ID,
		&conn.Name,
		&conn.Type,
		&conn.Host,
		&conn.Port,
		&encryptedUsername,
		&encryptedPassword,
		&conn.DatabaseName,
		&conn.SSL,
		&conn.DatabaseSize,
		&conn.CreatedAt,
		&conn.UpdatedAt,
		&conn.LastConnectedAt,
		&conn.UserID,
		&conn.Status,
	)
	if err != nil {
		return nil, err
	}

	conn.Username, err = r.crypto.Decrypt(encryptedUsername)
	if err != nil {
		return nil, err
	}

	conn.Password, err = r.crypto.Decrypt(encryptedPassword)
	if err != nil {
		return nil, err
	}

	return &conn, nil
}

func (r *ConnectionRepository) Update(conn StoredConnection) error {
	username, err := r.crypto.Encrypt(conn.Username)
	if err != nil {
		return err
	}

	password, err := r.crypto.Encrypt(conn.Password)
	if err != nil {
		return err
	}

	query := `
		UPDATE connections SET 
			name = $1, type = $2, host = $3, port = $4, 
			username = $5, password = $6, database_name = $7, 
			ssl = $8, database_size = $9, updated_at = NOW()
		WHERE id = $10`

	_, err = r.db.Exec(
		query,
		conn.Name,
		conn.Type,
		conn.Host,
		conn.Port,
		username,
		password,
		conn.DatabaseName,
		conn.SSL,
		conn.DatabaseSize,
		conn.ID,
	)

	return err
}

func (r *ConnectionRepository) ListByUserID(userID uuid.UUID) ([]ConnectionListItem, error) {
	query := `
		SELECT 
			c.id,
			c.name,
			c.type,
			c.host,
			c.status,
			c.database_size,
			b.completed_time as last_backup_time,
			COALESCE(bs.enabled, false) as backup_enabled,
			bs.cron_schedule,
			bs.retention_days
		FROM connections c
		LEFT JOIN backup_schedules bs ON c.id = bs.connection_id AND bs.enabled = true
		LEFT JOIN backups b ON c.id = b.connection_id
			AND b.completed_time = (
				SELECT MAX(completed_time)
				FROM backups
				WHERE connection_id = c.id
			)
		WHERE c.user_id = $1
		GROUP BY c.id, c.name, c.type, c.host, c.status, c.database_size, b.completed_time, bs.enabled, bs.cron_schedule, bs.retention_days
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var connections []ConnectionListItem
	for rows.Next() {
		var conn ConnectionListItem
		var lastBackupTime sql.NullString
		var cronSchedule sql.NullString
		var retentionDays sql.NullInt64

		err := rows.Scan(
			&conn.ID,
			&conn.Name,
			&conn.Type,
			&conn.Host,
			&conn.Status,
			&conn.DatabaseSize,
			&lastBackupTime,
			&conn.BackupEnabled,
			&cronSchedule,
			&retentionDays,
		)
		if err != nil {
			return nil, err
		}

		if lastBackupTime.Valid {
			conn.LastBackupTime = &lastBackupTime.String
		}
		if cronSchedule.Valid {
			conn.CronSchedule = &cronSchedule.String
		}
		if retentionDays.Valid {
			days := int(retentionDays.Int64)
			conn.RetentionDays = &days
		}

		connections = append(connections, conn)
	}
	return connections, nil
}

func (r *ConnectionRepository) UpdateLastConnected(id string) error {
	query := `UPDATE connections SET last_connected_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ConnectionRepository) UpdateStatus(id string, status string) error {
	query := `UPDATE connections SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(query, status, id)
	return err
}
