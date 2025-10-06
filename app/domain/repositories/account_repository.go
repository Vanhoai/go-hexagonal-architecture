package repositories

import (
	"app/domain/entities"
	"context"
)

// IAccountRepository extends base repository with account-specific operations
type IAccountRepository interface {
	IBaseRepository[*entities.AccountEntity]

	FindByEmail(ctx context.Context, email string) (*entities.AccountEntity, error)
	FindByName(ctx context.Context, name string) ([]*entities.AccountEntity, error)
}
