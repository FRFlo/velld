package backup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dendianugerah/velld/internal/mail"
	"github.com/dendianugerah/velld/internal/notification"
	"github.com/dendianugerah/velld/internal/settings"
	"github.com/google/uuid"
)

func (s *BackupService) createFailureNotification(connID string, backupErr error) error {

	conn, err := s.connStorage.GetConnection(connID)
	if err != nil {
		log.Printf("Failed to get connection details: %v", err)
		return fmt.Errorf("failed to get connection details: %v", err)
	}

	if conn == nil {
		log.Printf("Connection not found: %s", connID)
		return fmt.Errorf("connection not found: %s", connID)
	}

	if conn.UserID == uuid.Nil {
		log.Printf("Invalid user ID for connection: %s", connID)
		return fmt.Errorf("invalid user ID for connection: %s", connID)
	}

	userSettings, err := s.settingsRepo.GetUserSettings(conn.UserID)
	if err != nil {
		log.Printf("Failed to get user settings: %v", err)
		return fmt.Errorf("failed to get user settings: %v", err)
	}

	if userSettings == nil {
		log.Printf("No settings found for user: %s", conn.UserID)
		return fmt.Errorf("no settings found for user: %s", conn.UserID)
	}

	metadata := map[string]interface{}{
		"connection_id": connID,
		"database_name": conn.DatabaseName,
		"database_type": conn.Type,
		"error":         backupErr.Error(),
		"timestamp":     time.Now().Format(time.RFC3339),
	}

	metadataJSON, _ := json.Marshal(metadata)

	// Create dashboard notification if enabled
	if userSettings.NotifyDashboard {
		notification := &notification.Notification{
			ID:        uuid.New(),
			UserID:    conn.UserID,
			Title:     "Backup Failed",
			Message:   fmt.Sprintf("Backup failed for database '%s': %v", conn.DatabaseName, backupErr),
			Type:      notification.BackupFailed,
			Status:    notification.StatusUnread,
			Metadata:  metadataJSON,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := s.notificationRepo.CreateNotification(notification); err != nil {
			fmt.Printf("Error creating dashboard notification: %v\n", err)
		}
	}

	// Send webhook notification if enabled
	if userSettings.NotifyWebhook && userSettings.WebhookURL != nil {
		go s.sendWebhookNotification(*userSettings.WebhookURL, metadata)
	}

	// Send email notification if enabled
	if userSettings.NotifyEmail && userSettings.Email != nil {
		log.Printf("Attempting to send email notification to: %s", *userSettings.Email)
		// Use separate goroutine for email to prevent blocking
		go func(emailAddr string, userSettings *settings.UserSettings, meta map[string]interface{}) {
			if err := s.sendEmailNotification(emailAddr, userSettings, meta); err != nil {
				log.Printf("Failed to send email notification: %v", err)
			}
		}(*userSettings.Email, userSettings, metadata)
	} else {
		log.Printf("Email notification skipped - enabled: %v, email configured: %v",
			userSettings.NotifyEmail, userSettings.Email != nil)
	}

	return nil
}

func (s *BackupService) sendWebhookNotification(webhookURL string, data map[string]interface{}) {
	body, _ := json.Marshal(data)
	_, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("Error sending webhook notification: %v\n", err)
	}
}

func (s *BackupService) sendEmailNotification(email string, settings *settings.UserSettings, data map[string]interface{}) error {
	if settings == nil {
		return fmt.Errorf("settings cannot be nil")
	}

	if settings.SMTPHost == nil || settings.SMTPUsername == nil ||
		settings.SMTPPassword == nil || settings.SMTPPort == nil {
		return fmt.Errorf("incomplete SMTP configuration")
	}

	decryptedPassword, err := s.cryptoService.Decrypt(*settings.SMTPPassword)
	if err != nil {
		return fmt.Errorf("failed to decrypt SMTP password: %v", err)
	}

	smtpConfig := &mail.SMTPConfig{
		Host:     *settings.SMTPHost,
		Port:     *settings.SMTPPort,
		Username: *settings.SMTPUsername,
		Password: decryptedPassword,
	}

	msg := &mail.Message{
		From:    *settings.SMTPUsername,
		To:      email,
		Subject: "Backup Failed",
		Body:    fmt.Sprintf("Backup failed for database '%s'. Error: %v", data["database_name"], data["error"]),
	}

	if err := mail.SendEmail(smtpConfig, msg); err != nil {
		fmt.Printf("Error sending email notification: %v\n", err)
	}

	return nil
}
