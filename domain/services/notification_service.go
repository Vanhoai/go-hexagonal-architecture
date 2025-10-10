package services

import "domain/repositories"

type NotificationService struct {
	repository repositories.NotificationRepository
}

func NewNotificationService(repository repositories.NotificationRepository) *NotificationService {
	return &NotificationService{repository}
}
