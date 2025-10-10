package repositories

import (
	"context"
	"domain/entities"
)

type AccountRepository interface {
	BaseRepository[*entities.Account]

	FindByEmail(ctx context.Context, email string) (*entities.Account, error)
	FindByName(ctx context.Context, name string) ([]*entities.Account, error)
}
