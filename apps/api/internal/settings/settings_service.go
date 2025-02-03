package settings

import (
	"github.com/dendianugerah/velld/internal/common"
	"github.com/google/uuid"
)

type SettingsService struct {
	repo         *SettingsRepository
	cryptoService *common.EncryptionService
}

func NewSettingsService(repo *SettingsRepository, crypto *common.EncryptionService) *SettingsService {
	return &SettingsService{
		repo:         repo,
		cryptoService: crypto,
	}
}

func (s *SettingsService) GetUserSettings(userID uuid.UUID) (*UserSettings, error) {
	settings, err := s.repo.GetUserSettings(userID)
	if err != nil {
		return nil, err
	}

	// Remove sensitive data before returning
	settings.SMTPPassword = nil
	return settings, nil
}

func (s *SettingsService) UpdateUserSettings(userID uuid.UUID, req *UpdateSettingsRequest) (*UserSettings, error) {
	settings, err := s.repo.GetUserSettings(userID)
	if err != nil {
		return nil, err
	}

	// Update only provided fields
	if req.NotifyDashboard != nil {
		settings.NotifyDashboard = *req.NotifyDashboard
	}
	if req.NotifyEmail != nil {
		settings.NotifyEmail = *req.NotifyEmail
	}
	if req.NotifyWebhook != nil {
		settings.NotifyWebhook = *req.NotifyWebhook
	}
	if req.WebhookURL != nil {
		settings.WebhookURL = req.WebhookURL
	}
	if req.Email != nil {
		settings.Email = req.Email
	}
	if req.SMTPHost != nil {
		settings.SMTPHost = req.SMTPHost
	}
	if req.SMTPPort != nil {
		settings.SMTPPort = req.SMTPPort
	}
	if req.SMTPUsername != nil {
		settings.SMTPUsername = req.SMTPUsername
	}
	if req.SMTPPassword != nil {
		// Encrypt SMTP password before storing
		encryptedPass, err := s.cryptoService.Encrypt(*req.SMTPPassword)
		if err != nil {
			return nil, err
		}
		settings.SMTPPassword = &encryptedPass
	}

	if err := s.repo.UpdateUserSettings(settings); err != nil {
		return nil, err
	}

	return settings, nil
}
