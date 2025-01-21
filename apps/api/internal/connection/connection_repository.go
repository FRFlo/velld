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
			c.database_size,
			b.completed_time as last_backup_time,
			COALESCE(bs.enabled, false) as backup_enabled
		FROM connections c
		LEFT JOIN backup_schedules bs ON c.id = bs.connection_id
		LEFT JOIN backups b ON c.id = b.connection_id
			AND b.completed_time = (
				SELECT MAX(completed_time)
				FROM backups
				WHERE connection_id = c.id
			)
		WHERE c.user_id = $1
		GROUP BY c.id, c.name, c.type, c.database_size, b.completed_time, bs.enabled
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

		err := rows.Scan(
			&conn.ID,
			&conn.Name,
			&conn.Type,
			&conn.DatabaseSize,
			&lastBackupTime,
			&conn.BackupEnabled,
		)
		if err != nil {
			return nil, err
		}

		if lastBackupTime.Valid {
			conn.LastBackupTime = &lastBackupTime.String
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

func (r *ConnectionRepository) GetStatsByUserID(userID uuid.UUID) (*ConnectionStats, error) {
	stats := &ConnectionStats{}

	err := r.db.QueryRow(`
		SELECT 
			COUNT(*) as total_connections,
			COALESCE(SUM(database_size), 0) as total_size,
			COALESCE(AVG(NULLIF(database_size, 0)), 0) as average_size,
			COUNT(CASE WHEN status = 'connected' THEN 1 END) as active_count,
			COUNT(CASE WHEN status = 'disconnected' THEN 1 END) as inactive_count,
			COUNT(CASE WHEN ssl = true THEN 1 END) as ssl_count,
			(COUNT(CASE WHEN ssl = true THEN 1 END) / COUNT(*)) * 100 as ssl_percentage
		FROM connections 
		WHERE user_id = $1
	`, userID).Scan(
		&stats.TotalConnections,
		&stats.TotalSize,
		&stats.AverageSize,
		&stats.ActiveCount,
		&stats.InactiveCount,
		&stats.SSLCount,
		&stats.SSLPercentage,
	)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *ConnectionRepository) UpdateStatus(id string, status string) error {
	query := `UPDATE connections SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(query, status, id)
	return err
}
