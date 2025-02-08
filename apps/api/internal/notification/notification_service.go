package notification

import (
	"github.com/google/uuid"
)

type NotificationService struct {
	repo *NotificationRepository
}

func NewNotificationService(repo *NotificationRepository) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) GetNotifications(userID uuid.UUID) ([]*NotificationList, error) {
	return s.repo.GetUserNotifications(userID)
}

func (s *NotificationService) MarkAsRead(userID uuid.UUID, notificationID uuid.UUID) error {
	return s.repo.MarkAsRead(userID, notificationID)
}

func (s *NotificationService) DeleteNotifications(userID uuid.UUID, notificationIDs []uuid.UUID) error {
	return s.repo.DeleteNotifications(userID, notificationIDs)
}
