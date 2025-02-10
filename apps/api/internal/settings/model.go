package settings

import (
	"time"

	"github.com/google/uuid"
)

type UserSettings struct {
	ID                uuid.UUID `json:"id"`
	UserID            uuid.UUID `json:"user_id"`
	NotifyDashboard   bool      `json:"notify_dashboard"`
	NotifyEmail       bool      `json:"notify_email"`
	NotifyWebhook     bool      `json:"notify_webhook"`
	WebhookURL        *string   `json:"webhook_url,omitempty"`
	Email             *string   `json:"email,omitempty"`
	SMTPHost          *string   `json:"smtp_host,omitempty"`
	SMTPPort          *int      `json:"smtp_port,omitempty"`
	SMTPUsername      *string   `json:"smtp_username,omitempty"`
	SMTPPassword      *string   `json:"smtp_password,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type UpdateSettingsRequest struct {
	NotifyDashboard   *bool   `json:"notify_dashboard,omitempty"`
	NotifyEmail       *bool   `json:"notify_email,omitempty"`
	NotifyWebhook     *bool   `json:"notify_webhook,omitempty"`
	WebhookURL        *string `json:"webhook_url,omitempty"`
	Email             *string `json:"email,omitempty"`
	SMTPHost          *string `json:"smtp_host,omitempty"`
	SMTPPort          *int    `json:"smtp_port,omitempty"`
	SMTPUsername      *string `json:"smtp_username,omitempty"`
	SMTPPassword      *string `json:"smtp_password,omitempty"`
}
