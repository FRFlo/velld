package connection

import (
	"database/sql"

	"github.com/dendianugerah/velld/internal/common"
	"github.com/google/uuid"
)

type ConnectionStorage struct {
	db     *sql.DB
	crypto *common.EncryptionService
}

func NewConnectionStorage(db *sql.DB, crypto *common.EncryptionService) *ConnectionStorage {
	return &ConnectionStorage{
		db:     db,
		crypto: crypto,
	}
}

func (s *ConnectionStorage) SaveConnection(conn StoredConnection) error {
	// Encrypt sensitive data
	username, err := s.crypto.Encrypt(conn.Username)
	if err != nil {
		return err
	}

	password, err := s.crypto.Encrypt(conn.Password)
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

	_, err = s.db.Exec(
		query,
		conn.ID,
		conn.Name,
		conn.Type,
		conn.Host,
		conn.Port,
		username, // encrypted
		password, // encrypted
		conn.DatabaseName,
		conn.SSL,
		conn.UserID,
		conn.Status,
		conn.DatabaseSize,
	)

	return err
}

func (s *ConnectionStorage) GetConnection(id string) (*StoredConnection, error) {
	var conn StoredConnection
	var encryptedUsername, encryptedPassword string

	query := `SELECT * FROM connections WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(
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

	// Decrypt sensitive data
	conn.Username, err = s.crypto.Decrypt(encryptedUsername)
	if err != nil {
		return nil, err
	}

	conn.Password, err = s.crypto.Decrypt(encryptedPassword)
	if err != nil {
		return nil, err
	}

	return &conn, nil
}

func (s *ConnectionStorage) ListConnections(userID uuid.UUID) ([]StoredConnection, error) {
	query := `SELECT * FROM connections WHERE user_id = $1`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var connections []StoredConnection
	for rows.Next() {
		var conn StoredConnection
		var encryptedUsername, encryptedPassword string

		err := rows.Scan(
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

		// Decrypt sensitive data
		conn.Username, err = s.crypto.Decrypt(encryptedUsername)
		if err != nil {
			return nil, err
		}

		conn.Password, err = s.crypto.Decrypt(encryptedPassword)
		if err != nil {
			return nil, err
		}

		connections = append(connections, conn)
	}
	return connections, nil
}

func (s *ConnectionStorage) UpdateLastConnected(id string) error {
	query := `UPDATE connections SET last_connected_at = NOW() WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
