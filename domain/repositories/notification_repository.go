package repositories

import (
	"domain/entities"
)

type NotificationRepository interface {
	BaseRepository[*entities.Notification]
}
