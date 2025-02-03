package notification

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) CreateNotification(n *Notification) error {
	now := time.Now().Format(time.RFC3339)
	_, err := r.db.Exec(`
		INSERT INTO notifications (
			id, user_id, title, message, type, status, metadata, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		n.ID, n.UserID, n.Title, n.Message, n.Type, n.Status, n.Metadata, now, now)
	return err
}

func (r *NotificationRepository) GetUserNotifications(opts NotificationListOptions) ([]*NotificationList, int, error) {
	whereClause := "WHERE user_id = $1"
	args := []interface{}{opts.UserID}
	argCount := 2

	if opts.Status != nil {
		whereClause += fmt.Sprintf(" AND status = $%d", argCount)
		args = append(args, *opts.Status)
		argCount++
	}

	if opts.Type != nil {
		whereClause += fmt.Sprintf(" AND type = $%d", argCount)
		args = append(args, *opts.Type)
		argCount++
	}

	// Get total count
	var total int
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM notifications %s", whereClause)
	if err := r.db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Get notifications
	query := fmt.Sprintf(`
		SELECT id, title, message, type, status, created_at
		FROM notifications
		%s
		ORDER BY created_at DESC
		LIMIT $%d OFFSET $%d
	`, whereClause, argCount, argCount+1)

	args = append(args, opts.Limit, opts.Offset)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var notifications []*NotificationList
	for rows.Next() {
		n := &NotificationList{}
		err := rows.Scan(&n.ID, &n.Title, &n.Message, &n.Type, &n.Status, &n.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		notifications = append(notifications, n)
	}

	return notifications, total, nil
}

func (r *NotificationRepository) MarkAsRead(userID uuid.UUID, notificationIDs []uuid.UUID) error {
	query := `
		UPDATE notifications 
		SET status = $1, updated_at = $2 
		WHERE user_id = $3 AND id = ANY($4)`

	_, err := r.db.Exec(query, StatusRead, time.Now().Format(time.RFC3339), userID, notificationIDs)
	return err
}

func (r *NotificationRepository) DeleteNotifications(userID uuid.UUID, notificationIDs []uuid.UUID) error {
	_, err := r.db.Exec(
		"DELETE FROM notifications WHERE user_id = $1 AND id = ANY($2)",
		userID, notificationIDs)
	return err
}
