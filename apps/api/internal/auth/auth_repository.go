package auth

import (
	"database/sql"
	"errors"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateUser(data User) error {
	_, err := r.db.Exec("INSERT INTO users (id, username, password, created_at) VALUES ($1, $2, $3, $4)", data.ID, data.Username, data.Password, data.CreatedAt)
	return err
}

func (r *AuthRepository) GetUserByUsername(username string) (*User, error) {
	var user User
	err := r.db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}
	return &user, nil
}
