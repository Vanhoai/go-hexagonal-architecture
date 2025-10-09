package scylla

import (
	"server/domain/entities"
	"server/domain/repositories"
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
