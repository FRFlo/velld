package settings

import (
    "database/sql"
    "time"
    "github.com/google/uuid"
    "github.com/dendianugerah/velld/internal/common"
)

type SettingsRepository struct {
    db *sql.DB
}

func NewSettingsRepository(db *sql.DB) *SettingsRepository {
    return &SettingsRepository{db: db}
}

func (r *SettingsRepository) GetUserSettings(userID uuid.UUID) (*UserSettings, error) {
    settings := &UserSettings{}
    var createdAtStr, updatedAtStr string

    err := r.db.QueryRow(`
        SELECT id, user_id, notify_dashboard, notify_email, notify_webhook,
               webhook_url, email, smtp_host, smtp_port, smtp_username, 
               smtp_password, created_at, updated_at
        FROM user_settings
        WHERE user_id = $1`, userID).Scan(
        &settings.ID, &settings.UserID, &settings.NotifyDashboard,
        &settings.NotifyEmail, &settings.NotifyWebhook, &settings.WebhookURL,
        &settings.Email, &settings.SMTPHost, &settings.SMTPPort,
        &settings.SMTPUsername, &settings.SMTPPassword,
        &createdAtStr, &updatedAtStr)

    if err == sql.ErrNoRows {
        // Create default settings if none exist
        now := time.Now()
        settings = &UserSettings{
            ID:              uuid.New(),
            UserID:          userID,
            NotifyDashboard: true,
            CreatedAt:       now,
            UpdatedAt:       now,
        }
        return settings, r.CreateUserSettings(settings)
    }

    if err != nil {
        return nil, err
    }

    // Parse timestamps
    settings.CreatedAt, err = common.ParseTime(createdAtStr)
    if err != nil {
        return nil, err
    }

    settings.UpdatedAt, err = common.ParseTime(updatedAtStr)
    if err != nil {
        return nil, err
    }

    return settings, nil
}

func (r *SettingsRepository) CreateUserSettings(settings *UserSettings) error {
    _, err := r.db.Exec(`
        INSERT INTO user_settings (
            id, user_id, notify_dashboard, notify_email, notify_webhook,
            webhook_url, email, smtp_host, smtp_port, smtp_username, 
            smtp_password, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
        settings.ID, settings.UserID, settings.NotifyDashboard,
        settings.NotifyEmail, settings.NotifyWebhook, settings.WebhookURL,
        settings.Email, settings.SMTPHost, settings.SMTPPort,
        settings.SMTPUsername, settings.SMTPPassword,
        settings.CreatedAt, settings.UpdatedAt)
    return err
}

func (r *SettingsRepository) UpdateUserSettings(settings *UserSettings) error {
    settings.UpdatedAt = time.Now()
    _, err := r.db.Exec(`
        UPDATE user_settings SET
            notify_dashboard = $1, notify_email = $2, notify_webhook = $3,
            webhook_url = $4, email = $5, smtp_host = $6, smtp_port = $7,
            smtp_username = $8, smtp_password = $9, updated_at = $10
        WHERE user_id = $11`,
        settings.NotifyDashboard, settings.NotifyEmail, settings.NotifyWebhook,
        settings.WebhookURL, settings.Email, settings.SMTPHost, settings.SMTPPort,
        settings.SMTPUsername, settings.SMTPPassword, settings.UpdatedAt,
        settings.UserID)
    return err
}
