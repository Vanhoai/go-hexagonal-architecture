package scylla

import (
	"domain/entities"
	"domain/repositories"
)

type notificationRepositoryImpl struct {
	baseRepositoryImpl[*entities.Notification]
}

func NewNotificationRepository() repositories.NotificationRepository {
	return &notificationRepositoryImpl{}
}

type NotificationRepositoryImpl struct {
	baseRepositoryImpl[*entities.Notification]
}
